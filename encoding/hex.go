/*
**Go’nun `encoding/hex` (kısaca `hex`) paketini** en ince ayrıntısına kadar örneklerle anlatayım.

---

# 📦 Go `hex` Paketi

Go’nun `encoding/hex` paketi, **ikili (binary) veriyi hexadecimal (onaltılık) formata** çevirmek ve tekrar eski haline döndürmek için kullanılır.

* **Hexadecimal (onaltılık)** → 16 tabanlı sayı sistemi.

  * 0–9 ve A–F harfleri kullanılır.
  * Örn: `255 (decimal)` = `FF (hex)`
* **Kullanım Alanları**:

  * Hash değerlerini (MD5, SHA256 vb.) string olarak göstermek.
  * İkili veriyi daha okunabilir hale getirmek.
  * Kriptografik anahtarlar, byte dizilerini depolamak / loglamak.

---

## 🔑 Paket Fonksiyonları

`encoding/hex` paketindeki en önemli fonksiyonlar:

* `hex.Encode(dst, src []byte) int` → `src` verisini hex olarak `dst`'ye yazar.
* `hex.EncodeToString(src []byte) string` → `src`'yi hex string olarak döner.
* `hex.Decode(dst, src []byte) (int, error)` → Hex string’ini byte slice’a çevirir.
* `hex.DecodeString(s string) ([]byte, error)` → Hex string’ini byte slice’a döner.
* `hex.Dumper(w io.Writer)` → Hex dump için writer döner (debug amaçlı).

---

## 📌 Basit Encode/Decode Örneği
*/
``go
package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	// Binary data (byte slice)
	data := []byte("Merhaba")

	// Encode to hex string
	hexStr := hex.EncodeToString(data)
	fmt.Println("Hex:", hexStr)

	// Decode back to []byte
	decoded, err := hex.DecodeString(hexStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decoded:", string(decoded))
}
``
/*
🔹 Çıktı:

```
Hex: 4d657268616261
Decoded: Merhaba
```

---

## 📌 hex.Encode ve hex.Decode Kullanımı

`hex.Encode` → hedef bir `[]byte` buffer’a yazar.
`hex.Decode` → hex string’i byte slice’a çevirir.
*/
``go
package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	src := []byte("GoLang")

	// Encode
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	fmt.Println("Hex:", string(dst))

	// Decode
	decoded := make([]byte, hex.DecodedLen(len(dst)))
	n, err := hex.Decode(decoded, dst)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decoded:", string(decoded[:n]))
}
``
/*
🔹 Çıktı:

```
Hex: 476f4c616e67
Decoded: GoLang
```

---

## 📌 Hex Dumper Kullanımı

Hex dump, binary veriyi hem hex hem de ASCII karşılığıyla gösterir (debug amaçlı).
*/
``go
package main

import (
	"encoding/hex"
	"os"
)

func main() {
	data := []byte("Merhaba Hex Dumper! 1234567890")

	dumper := hex.Dumper(os.Stdout)
	defer dumper.Close()

	dumper.Write(data)
}
``
/*
🔹 Çıktı (örnek):

```
00000000  4d 65 72 68 61 62 61 20 48 65 78 20 44 75 6d 70  |Merhaba Hex Dump|
00000010  65 72 21 20 31 32 33 34 35 36 37 38 39 30         |er! 1234567890|
```

---

## 📌 Hatalı Hex String Decode

Eğer geçersiz bir hex string verilirse `hex.DecodeString` hata döner.
*/
``go
package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	badHex := "XYZ123" // Geçersiz karakter

	_, err := hex.DecodeString(badHex)
	if err != nil {
		fmt.Println("Hata:", err)
	}
}
``
/*
🔹 Çıktı:

```
Hata: encoding/hex: invalid byte: U+0058 'X'
```

---

## 📌 Özet

* `hex` paketi → **binary <-> hexadecimal dönüşümü** için kullanılır.
* Kullanım senaryoları:

  * Hash değerlerini string olarak göstermek.
  * Şifreleme anahtarlarını saklamak.
  * Binary verileri loglamak.
* Fonksiyonlar:

  * `EncodeToString`, `DecodeString` → string odaklı.
  * `Encode`, `Decode` → `[]byte` buffer ile çalışır.
  * `Dumper` → debug için güzel bir yardımcıdır.

---

👉 İstersen ben sana **SHA-256 hash + hex paketini birlikte kullanma** örneği de gösterebilirim (kriptografide çok yaygın). Bunu da ekleyeyim mi?
EVET
*/
/*
Harika 👍 O zaman şimdi **`crypto/sha256` + `encoding/hex`** birlikte kullanımını göstereyim.

Bu çok yaygın bir senaryo:

* Bir veriyi **hash’lemek** (ör. SHA-256 ile).
* Hash çıktısı `[]byte` olur → okunabilir hale getirmek için **hex string’e** çevrilir.

---

## 📌 SHA-256 Hash + Hex Örneği
*/
``go
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	// Hashlenecek veri
	data := "Merhaba Hex ve SHA256!"

	// SHA256 hash hesaplama
	hash := sha256.Sum256([]byte(data))

	// []byte → hex string
	hexHash := hex.EncodeToString(hash[:])

	fmt.Println("Orijinal veri :", data)
	fmt.Println("SHA-256 (hex) :", hexHash)

	// Hex string → []byte (geri çevirme)
	bytes, err := hex.DecodeString(hexHash)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decoded byte uzunluğu:", len(bytes))
}
``
/*
---

🔹 Örnek Çıktı:

```
Orijinal veri : Merhaba Hex ve SHA256!
SHA-256 (hex) : 56f0cc34623c6bcd3268e098fe8ce58d75c84736f786e1a90c4cda44a68ecf7d
Decoded byte uzunluğu: 32
```

---

## 📌 Nerelerde Kullanılır?

* ✅ Kullanıcı parolalarını hashleyip saklamak (ama SHA-256 tek başına değil, `bcrypt` veya `scrypt` ile birlikte).
* ✅ API anahtarlarını veya token’ları string formatta göstermek.
* ✅ Dosya bütünlüğü kontrolü (file checksum).
* ✅ Blockchain teknolojilerinde (Bitcoin, Ethereum) transaction hash’leri.

---

👉 Eğer istersen ben sana **bir dosyanın SHA-256 hash’ini hesaplayıp hex olarak yazdıran** bir örnek de verebilirim. Bu da gerçek hayatta çok kullanılır (örn: indirilen dosyanın sağlamlık kontrolü). Onu da ekleyeyim mi?
*/