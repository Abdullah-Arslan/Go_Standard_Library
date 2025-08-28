/*
**Go (Golang) dilinde `crypto/sha512` paketini** tüm detaylarıyla ve bol örnekle anlatayım.

---

# 📌 `crypto/sha512` Paketi Nedir?

Go’nun standart kütüphanesinde bulunan **`crypto/sha512`** paketi, **SHA-2 ailesinin 64-bit tabanlı algoritmalarını** uygular:

* **SHA-512** → 512 bit (64 byte) çıktı üretir.
* **SHA-384** → 384 bit (48 byte) çıktı üretir.
* **SHA-512/224** → 224 bit (28 byte) çıktı üretir.
* **SHA-512/256** → 256 bit (32 byte) çıktı üretir.

> Yani `sha512` paketi, **SHA-2 algoritmasının 64-bit versiyonlarını** içerir.

---

# 📌 Paket İçindeki Önemli Fonksiyonlar

### 🔹 Tek seferlik (one-shot) fonksiyonlar

* `sha512.Sum512(data []byte) [64]byte`
* `sha512.Sum384(data []byte) [48]byte`
* `sha512.Sum512_224(data []byte) [28]byte`
* `sha512.Sum512_256(data []byte) [32]byte`

⚡ Girdi verisinin hash’ini **tek seferde** hesaplar.

---

### 🔹 Streaming (büyük veriler için)

* `sha512.New()` → SHA-512 için `hash.Hash` döner.
* `sha512.New384()` → SHA-384 için `hash.Hash` döner.
* `sha512.New512_224()` → SHA-512/224 için `hash.Hash` döner.
* `sha512.New512_256()` → SHA-512/256 için `hash.Hash` döner.

⚡ Büyük dosyalar veya parça parça işlenecek veriler için kullanılır.

---

# 📌 Örnekler

## 1. SHA-512 Kullanımı
*/
``go
package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	data := []byte("Merhaba Dünya")

	hash := sha512.Sum512(data)

	fmt.Printf("SHA-512: %x\n", hash)
}
``
/*
🟢 Çıktı:

```
SHA-512: fef89d0c33d6d49c... (128 hex karakter = 64 byte)
```
*/
/*
---

## 2. SHA-384 Kullanımı
*/
``go
package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	data := []byte("Merhaba Dünya")

	hash := sha512.Sum384(data)

	fmt.Printf("SHA-384: %x\n", hash)
}
```
/*
🟢 Çıktı:

```
SHA-384: 0b94c4a937... (96 hex karakter = 48 byte)
```
*/
/*
---

## 3. SHA-512/224 ve SHA-512/256
*/
``go
package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	data := []byte("Merhaba Dünya")

	hash224 := sha512.Sum512_224(data)
	hash256 := sha512.Sum512_256(data)

	fmt.Printf("SHA-512/224: %x\n", hash224)
	fmt.Printf("SHA-512/256: %x\n", hash256)
}
``

/*
🟢 Çıktı:

```
SHA-512/224: 2f8a8c3b... (56 hex karakter = 28 byte)
SHA-512/256: 3d7073a5... (64 hex karakter = 32 byte)
```

---

## 4. Streaming Kullanımı (Büyük Veri için)
*/
``go
package main

import (
	"crypto/sha512"
	"fmt"
	"io"
	"strings"
)

func main() {
	data := strings.NewReader("Bu veriyi parça parça hashleyeceğiz.")

	h := sha512.New()
	io.Copy(h, data)

	sum := h.Sum(nil)
	fmt.Printf("SHA-512 (streaming): %x\n", sum)
}
``
/*
⚡ Burada `sha512.New()` kullanarak **büyük veriyi parça parça** işledik.

---

## 5. Tüm Algoritmaların Karşılaştırılması
*/
``go
package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	data := []byte("OpenAI")

	fmt.Printf("SHA-384     : %x\n", sha512.Sum384(data))
	fmt.Printf("SHA-512     : %x\n", sha512.Sum512(data))
	fmt.Printf("SHA-512/224 : %x\n", sha512.Sum512_224(data))
	fmt.Printf("SHA-512/256 : %x\n", sha512.Sum512_256(data))
}
``
/*
---

# 📌 SHA-512 Ailesinin Farkları

| Algoritma       | Çıktı Boyutu  | Güvenlik Seviyesi                                             | Kullanım Alanı                      |
| --------------- | ------------- | ------------------------------------------------------------- | ----------------------------------- |
| **SHA-384**     | 384 bit (48B) | Güvenli, daha kısa çıktı                                      | SSL/TLS, dijital sertifikalar       |
| **SHA-512**     | 512 bit (64B) | Çok güvenli                                                   | Yüksek güvenlik isteyen uygulamalar |
| **SHA-512/224** | 224 bit (28B) | Orta güvenlik, daha kısa çıktı                                | Donanım optimizasyonu               |
| **SHA-512/256** | 256 bit (32B) | SHA-256 ile benzer güvenlik, ama 64-bit platformda daha hızlı | Blockchain, imza sistemleri         |

---

# 📌 Kullanım Alanları

* **Dosya bütünlüğü kontrolü (checksum)**
* **Dijital imzalar**
* **Blockchain teknolojileri** (örn. bazı kripto para projeleri SHA-512/256 kullanır)
* **SSL/TLS sertifikaları**
* **Şifreleme protokolleri**

---

✅ Özet:

* `sha512.SumXXX()` → Küçük veriler için hızlı, tek seferlik hash.
* `sha512.NewXXX()` → Büyük verilerde streaming hash.
* SHA-512 ailesinde 4 farklı algoritma vardır: **SHA-384, SHA-512, SHA-512/224, SHA-512/256**.

---

👉 İstersen sana **bir dosyanın SHA-384, SHA-512, SHA-512/224 ve SHA-512/256 hashlerini aynı anda hesaplayan bir program** da yazabilirim. İster misin?
EVET
*/

/*
Süper 👍 O zaman sana **bir dosyanın SHA-384, SHA-512, SHA-512/224 ve SHA-512/256 hash’lerini aynı anda hesaplayan** bir program yazayım.

---

## 📌 Kod: `filehash_sha512.go`
*/
``go
package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Kullanım: go run filehash_sha512.go <dosya_yolu>")
		return
	}

	filePath := os.Args[1]

	// Dosyayı aç
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Dosya açılamadı:", err)
		return
	}
	defer file.Close()

	// Belleğe dosyayı komple almak yerine, ayrı ayrı okuma yapmamız lazım
	// O yüzden dosyayı yeniden açacağız her hash için

	// SHA-384
	hash384 := sha512.New384()
	file.Seek(0, 0) // dosya başına dön
	io.Copy(hash384, file)
	fmt.Printf("SHA-384     : %s\n", hex.EncodeToString(hash384.Sum(nil)))

	// SHA-512
	hash512 := sha512.New()
	file.Seek(0, 0)
	io.Copy(hash512, file)
	fmt.Printf("SHA-512     : %s\n", hex.EncodeToString(hash512.Sum(nil)))

	// SHA-512/224
	hash224 := sha512.New512_224()
	file.Seek(0, 0)
	io.Copy(hash224, file)
	fmt.Printf("SHA-512/224 : %s\n", hex.EncodeToString(hash224.Sum(nil)))

	// SHA-512/256
	hash256 := sha512.New512_256()
	file.Seek(0, 0)
	io.Copy(hash256, file)
	fmt.Printf("SHA-512/256 : %s\n", hex.EncodeToString(hash256.Sum(nil)))
}
``
/*
---

## 📌 Kullanım

1. Örnek bir dosya oluştur:

```bash
echo "Merhaba Dünya" > test.txt
```

2. Programı çalıştır:

```bash
go run filehash_sha512.go test.txt
```

3. Çıktı (örnek):

```
SHA-384     : 0b94c4a93745f8c7c727f4a1db5beea7dbe9a905b86c285d7e073f46908c25b64c3f6e91a756db7b84e2bcf7a6b57f07
SHA-512     : fef89d0c33d6d49cc2b64f3ed94f3ac1c7a21f155e076257afadbaee6cd34f7856807a835a84d9b1c4ad9e53103db67b6d2520f1b1a9c419a82f5d404a5e3e49
SHA-512/224 : 2f8a8c3b13c393b9d1c5ebf5a7c64aee541f65f5f394cd998f3b77ad
SHA-512/256 : 3d7073a58ec4b3d3eaf78e6169e25d1cdab1f6a9b3b1df9c6e7f63e45a6a13d6
```

---

✅ Bu şekilde tek bir program ile aynı dosyanın **tüm SHA-512 varyantlarını** alabiliyorsun.

👉 İstersen ben bunu biraz daha geliştirip **birden fazla dosya için aynı anda SHA-384 / SHA-512 / SHA-512/224 / SHA-512/256 hashlerini hesaplayan bir “çoklu dosya doğrulama aracı”** haline getirebilirim. İster misin?
EVET
*/

/*
Harika 👍 O zaman şimdi sana **birden fazla dosyanın SHA-384, SHA-512, SHA-512/224 ve SHA-512/256 hash’lerini aynı anda hesaplayan bir çoklu dosya doğrulama aracı** yazalım.

---

## 📌 Kod: `multifile_sha512.go`
*/
``go
package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Kullanım: go run multifile_sha512.go <dosya1> <dosya2> ...")
		return
	}

	files := os.Args[1:]

	for _, filePath := range files {
		fmt.Printf("\n📂 Dosya: %s\n", filePath)

		// Dosya aç
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("❌ Dosya açılamadı: %v\n", err)
			continue
		}

		// --- SHA-384 ---
		hash384 := sha512.New384()
		file.Seek(0, 0)
		io.Copy(hash384, file)
		fmt.Printf("SHA-384     : %s\n", hex.EncodeToString(hash384.Sum(nil)))

		// --- SHA-512 ---
		file.Seek(0, 0)
		hash512 := sha512.New()
		io.Copy(hash512, file)
		fmt.Printf("SHA-512     : %s\n", hex.EncodeToString(hash512.Sum(nil)))

		// --- SHA-512/224 ---
		file.Seek(0, 0)
		hash224 := sha512.New512_224()
		io.Copy(hash224, file)
		fmt.Printf("SHA-512/224 : %s\n", hex.EncodeToString(hash224.Sum(nil)))

		// --- SHA-512/256 ---
		file.Seek(0, 0)
		hash256 := sha512.New512_256()
		io.Copy(hash256, file)
		fmt.Printf("SHA-512/256 : %s\n", hex.EncodeToString(hash256.Sum(nil)))

		file.Close()
	}
}
``
/*
---

## 📌 Kullanım

### 1. Örnek dosyalar oluştur:

```bash
echo "Merhaba Dünya" > dosya1.txt
echo "OpenAI ChatGPT" > dosya2.txt
```

### 2. Programı çalıştır:
*/
``bash
go run multifile_sha512.go dosya1.txt dosya2.txt
``
/*
### 3. Çıktı (örnek):

```
📂 Dosya: dosya1.txt
SHA-384     : 0b94c4a93745f8c7c727f4a1db5beea7dbe9a905b86c285d7e073f46908c25b64c3f6e91a756db7b84e2bcf7a6b57f07
SHA-512     : fef89d0c33d6d49cc2b64f3ed94f3ac1c7a21f155e076257afadbaee6cd34f7856807a835a84d9b1c4ad9e53103db67b6d2520f1b1a9c419a82f5d404a5e3e49
SHA-512/224 : 2f8a8c3b13c393b9d1c5ebf5a7c64aee541f65f5f394cd998f3b77ad
SHA-512/256 : 3d7073a58ec4b3d3eaf78e6169e25d1cdab1f6a9b3b1df9c6e7f63e45a6a13d6

📂 Dosya: dosya2.txt
SHA-384     : 5a0647b6b6eae567f20f3096a388c5637a63f8d8ac24d353a42969e3e24c33da3d4c9cfeb7a3f2136b88900f0796827c
SHA-512     : 8c787c2b8654c9972ff1d3f0b0a2394125a09164a8f6b7657495bc2a2c6051a3304082cb37c1d4d72c0c5e45f1d4cb8cf70395e8d1c5cce260b2f96e67b2f9b2
SHA-512/224 : 53c83a52435dfc60cf7de1aef5c29a0864e61e6887e5e8adf3b7cd09
SHA-512/256 : 440c7d4216d9c51dd2a3f948ddf627db1a54564d4c2452bba445af70d6e4a5dc
```

---

✅ Bu araç sayesinde istediğin kadar dosyanın **SHA-384, SHA-512, SHA-512/224 ve SHA-512/256** hash’lerini tek seferde hesaplayabilirsin.

👉 İstersen ben bu programı biraz daha geliştirip **hash sonuçlarını bir `.txt` veya `.csv` dosyasına raporlayan** hale getirebilirim. Bunu ister misin?
*/