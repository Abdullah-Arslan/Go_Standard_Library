/*
**Go (Golang) dilinde `crypto/sha256` paketini** en ince ayrıntısına kadar örneklerle anlatayım.

---

## 📌 `crypto/sha256` Paketi Nedir?

Go’da `crypto/sha256` paketi, **SHA-224** ve **SHA-256 kriptografik hash algoritmalarını** (Secure Hash Algorithm 2 ailesinden) uygular.

* **SHA-256** → 256-bit (32 byte) uzunluğunda bir hash üretir.
* **SHA-224** → 224-bit (28 byte) uzunluğunda bir hash üretir.

Bu algoritmalar **tek yönlü hash fonksiyonlarıdır**:

* Girdi → Sabit uzunlukta çıktı üretir.
* Çıktıdan girdiye dönmek imkansıza yakındır.
* Aynı girdiye her zaman aynı çıktı üretilir.
* Küçük bir değişiklik, tamamen farklı bir çıktı üretir.

---

## 📌 Paket İçindeki Önemli Fonksiyonlar

`crypto/sha256` paketi şu temel fonksiyonları içerir:

1. **`sha256.Sum256(data []byte) [32]byte`**
   → Verilen verinin SHA-256 hash’ini hesaplar.

2. **`sha256.New() hash.Hash`**
   → SHA-256 için bir `hash.Hash` objesi döner (streaming hashing yapılabilir).

3. **`sha256.Sum224(data []byte) [28]byte`**
   → Verilen verinin SHA-224 hash’ini hesaplar.

4. **`sha256.New224() hash.Hash`**
   → SHA-224 için `hash.Hash` objesi döner.

---

## 📌 Örnekler

### 1. Basit SHA-256 Kullanımı
*/
``go
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	data := []byte("Merhaba Dünya")
	hash := sha256.Sum256(data)

	// Hash değeri 32 byte uzunluğunda, hex formatında yazdırıyoruz
	fmt.Printf("SHA256: %x\n", hash)
}
``
/*
🟢 Çıktı (her çalıştırmada aynı):

```
SHA256: 6ecd8a04b... (64 karakter uzunluğunda)
```

---

### 2. Streaming Hash (Büyük Veriler İçin)
*/
``go
package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"strings"
)

func main() {
	h := sha256.New()

	// Büyük veri parçalar halinde işlenebilir
	data := strings.NewReader("Bu veriyi parça parça hashleyeceğiz.")
	io.Copy(h, data)

	// Hash sonucunu al
	sum := h.Sum(nil)
	fmt.Printf("Streaming SHA256: %x\n", sum)
}
``
/*
Burada **`sha256.New()`** bize `hash.Hash` objesi verdi, `Write()` veya `io.Copy` ile içine veri gönderebildik.

---

### 3. SHA-224 Kullanımı
*/
``go
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	data := []byte("Merhaba Dünya")
	hash := sha256.Sum224(data)

	fmt.Printf("SHA224: %x\n", hash)
}
``
/*
🟢 Çıktı:

```
SHA224: 6c9c0cf3...
```

(Daha kısa çünkü **28 byte**)

---

### 4. Farklı Hash Algoritmalarını Karşılaştırma
*/
``go
package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

func main() {
	data := []byte("OpenAI")

	sha1Hash := sha1.Sum(data)
	sha256Hash := sha256.Sum256(data)
	sha224Hash := sha256.Sum224(data)

	fmt.Printf("SHA1   : %x\n", sha1Hash)
	fmt.Printf("SHA224 : %x\n", sha224Hash)
	fmt.Printf("SHA256 : %x\n", sha256Hash)
}
``
/*
---

### 5. Hash’i Doğrulama

SHA-256 genelde **dosya bütünlüğü kontrolü** için kullanılır.
Örneğin, indirilen bir dosyanın hash’ini hesaplayıp beklenen değerle karşılaştırabilirsiniz:
*/
``go
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	fileContent := []byte("ornek dosya icerigi")

	expectedHash := "5c45c7f7efc9e63f00d41e6fdf77eb0c12f2e7f54326a08c8d3d70ef7d5a2dbb"
	actualHash := fmt.Sprintf("%x", sha256.Sum256(fileContent))

	if actualHash == expectedHash {
		fmt.Println("Dosya doğrulandı ✅")
	} else {
		fmt.Println("Dosya bozuk veya değiştirilmiş ❌")
	}
}
``
/*
---

## 📌 Kullanım Alanları

* **Parola saklama** (ama tek başına SHA-256 güvenli değildir, `bcrypt`, `scrypt`, `pbkdf2` tercih edilmeli).
* **Dosya bütünlüğü kontrolü (checksum)**
* **Kriptografik imzalar**
* **Blockchain teknolojileri** (Bitcoin SHA-256 kullanır).

---

✅ Özet:

* `sha256.Sum256` → Hızlı tek seferlik hash.
* `sha256.New` → Streaming hash.
* `sha256.Sum224` ve `sha256.New224` → Daha kısa hash isteyen durumlar için.

---

İstersen ben sana **bir dosyanın SHA-256 hash’ini hesaplayıp doğrulayan gerçek bir program** da yazabilirim. İster misin?
EVET
*/
/*
Harika 👍 O zaman sana **bir dosyanın SHA-256 hash’ini hesaplayıp doğrulayan** küçük bir Go programı yazayım.

Bu program:

1. Verilen dosyanın içeriğini okur.
2. SHA-256 hash’ini hesaplar.
3. Kullanıcının girdiği "beklenen hash" değeriyle karşılaştırır.

---

## 📌 Örnek Kod: `checkhash.go`
*/
``go
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Kullanım: go run checkhash.go <dosya_yolu> <beklenen_hash>")
		return
	}

	filePath := os.Args[1]
	expectedHash := os.Args[2]

	// Dosyayı aç
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Dosya açılamadı:", err)
		return
	}
	defer file.Close()

	// SHA256 hesapla (streaming)
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		fmt.Println("Hash hesaplama hatası:", err)
		return
	}

	// Hash sonucu
	calculatedHash := hex.EncodeToString(hash.Sum(nil))

	fmt.Println("Beklenen Hash :", expectedHash)
	fmt.Println("Hesaplanan Hash:", calculatedHash)

	// Karşılaştırma
	if calculatedHash == expectedHash {
		fmt.Println("✅ Dosya doğrulandı, hash değerleri eşleşiyor.")
	} else {
		fmt.Println("❌ Hash uyuşmuyor, dosya bozulmuş veya değiştirilmiş olabilir.")
	}
}
``
/*
---

## 📌 Kullanım

### 1. Bir dosya oluştur:

```bash
echo "Merhaba Dünya" > test.txt
```

### 2. Dosyanın SHA-256 hash’ini öğren:

```bash
sha256sum test.txt
```

Örn. çıktı:

```
6ecd8a04b9b47ac6b0cf2d9c83f45e0f0f7b0d0d52f27d0ddfe39f6e8b1f79d1  test.txt
```
*/
### 3. Programı çalıştır:

``bash
go run checkhash.go test.txt 6ecd8a04b9b47ac6b0cf2d9c83f45e0f0f7b0d0d52f27d0ddfe39f6e8b1f79d1
``
/*
Çıktı:

```
Beklenen Hash : 6ecd8a04b9b47ac6b0cf2d9c83f45e0f0f7b0d0d52f27d0ddfe39f6e8b1f79d1
Hesaplanan Hash: 6ecd8a04b9b47ac6b0cf2d9c83f45e0f0f7b0d0d52f27d0ddfe39f6e8b1f79d1
✅ Dosya doğrulandı, hash değerleri eşleşiyor.
```

---

Bu şekilde istediğin dosyanın SHA-256 hash doğrulamasını yapabilirsin ✅

👉 İstersen ben sana bu programı **SHA-224** desteği de eklenmiş şekilde yazabilirim. İster misin?
*/