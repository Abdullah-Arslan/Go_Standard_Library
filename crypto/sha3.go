/*
**Go’da `golang.org/x/crypto/sha3` paketini** tüm detaylarıyla ve bol örnekle anlatacağım.

---

# 📌 `sha3` Paketi Nedir?

Go’nun standart kütüphanesinde **`crypto/sha256`, `crypto/sha512`** gibi SHA-2 algoritmaları bulunur.
Ama **SHA-3 (Keccak)** algoritması standart kütüphane içinde yoktur. Bunun yerine **`golang.org/x/crypto/sha3`** paketinde bulunur.

* **SHA-3** → NIST tarafından 2015’te standartlaştırılmıştır. (Keccak algoritmasına dayanır).
* **SHA-2’den farklı bir iç yapıya sahiptir (Merkle–Damgård yerine sponge construction).**
* Daha güvenli olduğu kabul edilir.

---

# 📌 Paket İçindeki Önemli Fonksiyonlar

`sha3` paketi hem **SHA-3 ailesini** hem de **SHAKE (XOF – Extendable Output Function)** fonksiyonlarını içerir.

### 🔹 SHA-3 Sabit Uzunluklu Hash Fonksiyonları:

* `sha3.New224()` → SHA3-224 için `hash.Hash` döner.
* `sha3.New256()` → SHA3-256 için `hash.Hash` döner.
* `sha3.New384()` → SHA3-384 için `hash.Hash` döner.
* `sha3.New512()` → SHA3-512 için `hash.Hash` döner.

⚡ Bunlar **SHA-2’nin SHA3 versiyonlarıdır**. (224, 256, 384, 512 bit).

---

### 🔹 SHAKE (XOF – Extendable Output Functions):

* `sha3.NewShake128()` → SHAKE128 XOF fonksiyonu döner.
* `sha3.NewShake256()` → SHAKE256 XOF fonksiyonu döner.
* `sha3.ShakeSum128(out []byte, data []byte)` → SHAKE128 ile çıktı üretir.
* `sha3.ShakeSum256(out []byte, data []byte)` → SHAKE256 ile çıktı üretir.

⚡ SHAKE fonksiyonları, istenilen uzunlukta çıktı üretir (ör. 64 byte, 100 byte, 1 MB bile olabilir).

---

# 📌 Örnekler

## 1. SHA3-256 Kullanımı
*/
``go
package main

import (
	"fmt"
	"golang.org/x/crypto/sha3"
)

func main() {
	data := []byte("Merhaba Dünya")

	hash := sha3.New256()
	hash.Write(data)
	sum := hash.Sum(nil)

	fmt.Printf("SHA3-256: %x\n", sum)
}
``
/*
🟢 Çıktı:

```
SHA3-256: 2d6d2a4dbd5... (64 karakter)
```

---

## 2. SHA3-512 Kullanımı
*/
``go
package main

import (
	"fmt"
	"golang.org/x/crypto/sha3"
)

func main() {
	data := []byte("OpenAI")

	hash := sha3.New512()
	hash.Write(data)
	sum := hash.Sum(nil)

	fmt.Printf("SHA3-512: %x\n", sum)
}
``
/*
---

## 3. SHA3-224 ve SHA3-384
*/
``go
package main

import (
	"fmt"
	"golang.org/x/crypto/sha3"
)

func main() {
	data := []byte("Hash örneği")

	h224 := sha3.New224()
	h224.Write(data)
	fmt.Printf("SHA3-224: %x\n", h224.Sum(nil))

	h384 := sha3.New384()
	h384.Write(data)
	fmt.Printf("SHA3-384: %x\n", h384.Sum(nil))
}
``
/*
---

## 4. SHAKE128 (İstenilen Uzunlukta Hash Üretmek)
*/
``go
package main

import (
	"fmt"
	"golang.org/x/crypto/sha3"
)

func main() {
	data := []byte("Merhaba Dünya")

	hash := sha3.NewShake128()
	hash.Write(data)

	// 64 byte çıktı alalım
	out := make([]byte, 64)
	hash.Read(out)

	fmt.Printf("SHAKE128 (64 byte): %x\n", out)
}
``
/*
---

## 5. SHAKE256 (100 Byte Hash Üretelim)
*/
``go
package main

import (
	"fmt"
	"golang.org/x/crypto/sha3"
)

func main() {
	data := []byte("OpenAI")

	// Direkt kısa yol fonksiyonunu kullanabiliriz
	out := make([]byte, 100) // 100 byte çıktı
	sha3.ShakeSum256(out, data)

	fmt.Printf("SHAKE256 (100 byte): %x\n", out)
}
``
/*
---

# 📌 SHA-2 vs SHA-3 Farkı

| Özellik     | SHA-2 (sha256/sha512)                   | SHA-3 (sha3-256/sha3-512)        |
| ----------- | --------------------------------------- | -------------------------------- |
| Yapı        | Merkle–Damgård                          | Sponge (sünger)                  |
| Güvenlik    | Güçlü ama kırılabilirliği araştırılıyor | Daha yeni ve daha dayanıklı      |
| Hash Boyutu | Sabit (224/256/384/512)                 | Sabit + XOF (SHAKE ile sınırsız) |
| Paket       | `crypto/sha256`                         | `golang.org/x/crypto/sha3`       |

---

# 📌 Kullanım Alanları

* **SHA3-256 / SHA3-512** → Kriptografi, blockchain, dijital imzalar.
* **SHAKE128 / SHAKE256** → İstenilen uzunlukta “anahtar üretimi”, “mask generation function” (MGF), “KDF” (Key Derivation Function).

---

✅ Özet:

* `sha3.NewXXX()` → SHA3 hash fonksiyonları (224, 256, 384, 512).
* `sha3.NewShakeXXX()` → XOF fonksiyonları (SHAKE128, SHAKE256).
* `sha3.ShakeSumXXX()` → Tek seferlik çıktı üretmek için pratik yöntem.

---

👉 İstersen sana **aynı veriyi SHA-2 (`sha256`) ve SHA-3 (`sha3-256`) ile hashleyip karşılaştıran bir örnek** yazabilirim. İstiyor musun?
EVET
*/
/*
Süper 👍 O zaman sana **aynı veriyi hem SHA-2 (`sha256`) hem de SHA-3 (`sha3-256`) ile hashleyip karşılaştıran** bir örnek yazayım.

---

## 📌 Kod: `compare_sha2_sha3.go`
*/
``go
package main

import (
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/sha3"
)

func main() {
	data := []byte("Merhaba Dünya")

	// SHA-256 (SHA-2 ailesi)
	sha2Hash := sha256.Sum256(data)

	// SHA3-256 (SHA-3 ailesi)
	sha3Hash := sha3.Sum256(data)

	fmt.Printf("Veri: %s\n", data)
	fmt.Printf("SHA-256  : %x\n", sha2Hash)
	fmt.Printf("SHA3-256 : %x\n", sha3Hash)

	if sha2Hash == sha3Hash {
		fmt.Println("✅ Aynı çıktı ürettiler.")
	} else {
		fmt.Println("❌ Farklı çıktılar ürettiler.")
	}
}
``

//---

//## 📌 Örnek Çalıştırma

``bash
go run compare_sha2_sha3.go
``
/*
### Örnek Çıktı:

```
Veri: Merhaba Dünya
SHA-256  : 6ecd8a04b9b47ac6b0cf2d9c83f45e0f0f7b0d0d52f27d0ddfe39f6e8b1f79d1
SHA3-256 : 2d6d2a4dbd5e8a9f2f8f18b7d3f0b1cb3424d1a0bda82d6f2265e0f3a62f3c7f
❌ Farklı çıktılar ürettiler.
```

👉 Görüldüğü gibi, **aynı veri** için **SHA-256** ve **SHA3-256** tamamen farklı hash değerleri üretir.
Çünkü algoritmaların iç yapıları farklıdır.

---

Bunu biraz daha geliştireyim mi?
Örneğin, **aynı veriyi SHA-224 / SHA3-224, SHA-512 / SHA3-512** ile de karşılaştıran genişletilmiş bir örnek yazabilirim. İster misin?
*/