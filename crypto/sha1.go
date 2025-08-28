/*
**Go’da `crypto/sha1` paketi**ni en temelden tüm fonksiyonlarıyla açıklayayım.

---

## 📌 `crypto/sha1` Paketi Nedir?

* Go’nun **SHA-1 (Secure Hash Algorithm 1)** fonksiyonunu sağlayan paketidir.
* SHA-1, 160 bit (20 byte) uzunluğunda bir özet (hash) üretir.
* Kullanımı kolaydır, genellikle şu amaçlarla kullanılır:

  * Veri bütünlüğü kontrolü
  * Hash tabanlı indeksleme
  * Eski sistemlerle uyumluluk

⚠️ **Not:** SHA-1 **artık kriptografik olarak güvenli değildir**. Çakışmalar (collision attacks) bulunmuştur. Güvenli uygulamalar için **SHA-256 veya SHA-512** kullanılır (`crypto/sha256`, `crypto/sha512` paketleri).

---

## 📌 Önemli Fonksiyonlar

### 1. `sha1.New()`

Yeni bir hash nesnesi döner (`hash.Hash` arayüzünü uygular).
*/
``go
import (
	"crypto/sha1"
	"fmt"
)

func main() {
	h := sha1.New()
	h.Write([]byte("Merhaba Dünya"))
	sum := h.Sum(nil)
	fmt.Printf("SHA1: %x\n", sum)
}
``
/*
---

### 2. `sha1.Sum(data []byte) [20]byte`

Direkt olarak tek seferde SHA-1 hash hesaplar.
*/
``go
data := []byte("Merhaba Dünya")
hash := sha1.Sum(data)
fmt.Printf("SHA1: %x\n", hash)
``

//📌 Burada dönüş tipi `[20]byte` → yani sabit uzunlukta bir dizi. Eğer `[]byte` olarak kullanmak istersen:

``go
hashBytes := hash[:] 
``
/*
---

### 3. **Adım Adım Hashleme (Stream İşleme)**

SHA-1 büyük verilerde parça parça işlenebilir.
*/
``go
h := sha1.New()
h.Write([]byte("Merhaba "))
h.Write([]byte("Dünya"))
fmt.Printf("SHA1: %x\n", h.Sum(nil))
``
/*
---

### 4. **Dosya Hash Hesaplama**

Özellikle büyük dosyaların SHA-1 kontrolü için:
*/
``go
import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func main() {
	file, _ := os.Open("example.txt")
	defer file.Close()

	h := sha1.New()
	_, _ = io.Copy(h, file) // dosyayı parça parça okur
	fmt.Printf("SHA1: %x\n", h.Sum(nil))
}
``
/*
---

## 📌 Komple Örnek
*/
``go
package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func main() {
	// 1. Tek seferde hash hesaplama
	data := []byte("Merhaba SHA1")
	hash := sha1.Sum(data)
	fmt.Printf("Tek seferde SHA1: %x\n", hash)

	// 2. Stream halinde hash hesaplama
	h := sha1.New()
	h.Write([]byte("Merhaba "))
	h.Write([]byte("SHA1 Stream"))
	fmt.Printf("Stream SHA1: %x\n", h.Sum(nil))

	// 3. Dosya hash hesaplama
	file, _ := os.Open("example.txt")
	defer file.Close()
	h2 := sha1.New()
	io.Copy(h2, file)
	fmt.Printf("Dosya SHA1: %x\n", h2.Sum(nil))
}
``
/*
---

## 📌 SHA-1 Paket Özeti

* `sha1.New()` → yeni hash oluşturur (stream destekler).
* `sha1.Sum(data)` → tek adımda hash hesaplar.
* Hash sonucu her zaman **20 byte (160 bit)** uzunluğunda olur.
* Güvenli kriptografi için **kullanılmamalı**, sadece eski sistemlerle uyumluluk için kullanılmalı.

---

👉 İstersen sana bir de **SHA-1 ile dosya doğrulama programı** (örn: indirilen dosyanın hash’i beklenen hash ile aynı mı) örneği yazabilirim. İstiyor musun?
EVET
*/
/*
Harika 👍 O zaman sana **SHA-1 ile dosya doğrulama programı** örneğini yazayım.

Bu program şunları yapacak:

1. Bir dosyanın SHA-1 hash’ini hesaplayacak.
2. Kullanıcının verdiği “beklenen hash” ile karşılaştıracak.
3. Uyumluysa ✅ diyecek, değilse ❌ uyaracak.

---

## 📌 SHA-1 Dosya Doğrulama Programı
*/
``go
package main

import (
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
)

// Dosya için SHA-1 hesapla
func fileSHA1(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	h := sha1.New()
	if _, err := io.Copy(h, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

func main() {
	// Komut satırı parametreleri
	filename := flag.String("file", "", "Kontrol edilecek dosya")
	expected := flag.String("hash", "", "Beklenen SHA1 hash (hex formatında)")
	flag.Parse()

	if *filename == "" || *expected == "" {
		fmt.Println("Kullanım: go run main.go -file=dosya.txt -hash=beklenenhash")
		os.Exit(1)
	}

	// Dosyanın SHA1 hash’ini hesapla
	actual, err := fileSHA1(*filename)
	if err != nil {
		fmt.Println("Hata:", err)
		os.Exit(1)
	}

	// Karşılaştır
	fmt.Println("Dosya:", *filename)
	fmt.Println("Beklenen:", *expected)
	fmt.Println("Hesaplanan:", actual)

	if actual == *expected {
		fmt.Println("✅ Hash eşleşti, dosya DOĞRU.")
	} else {
		fmt.Println("❌ Hash eşleşmedi, dosya DEĞİŞTİRİLMİŞ olabilir!")
	}
}
``
/*
---

## 📌 Kullanım

1. Bir dosya oluştur (`example.txt`)

```txt
Merhaba Dünya
```

2. Dosyanın SHA-1 hash’ini öğrenmek için:

```bash
sha1sum example.txt
```

📌 (Linux/Mac’te hazır gelir. Windows için `CertUtil -hashfile example.txt SHA1` kullanılabilir.)

Örneğin çıktı:

```
b1f5c8f2b019a78cddc6a64d3bb1a72f9f2e41a3  example.txt
```

3. Programı çalıştır:
*/
``bash
go run main.go -file=example.txt -hash=b1f5c8f2b019a78cddc6a64d3bb1a72f9f2e41a3
``
/*
### Örnek Çıktı

``
Dosya: example.txt
Beklenen: b1f5c8f2b019a78cddc6a64d3bb1a72f9f2e41a3
Hesaplanan: b1f5c8f2b019a78cddc6a64d3bb1a72f9f2e41a3
✅ Hash eşleşti, dosya DOĞRU.
``
*/
/*
---

👉 İstersen ben sana bu programı **SHA256 ve SHA512 destekli** hale getirip, parametreyle (`-algo=sha256`) seçilebilir hale de getirebilirim. Bunu ister misin?
EVET
*/


/*
Süper 👍 O zaman sana **çok algoritmalı dosya doğrulama programı** yazıyorum.

Bu sürümde:

* `-file` → doğrulanacak dosya
* `-hash` → beklenen hash değeri (hex formatında)
* `-algo` → kullanılacak algoritma (`sha1`, `sha256`, `sha512`)

---

## 📌 Çok Algoritmalı Dosya Doğrulama Programı
*/
``go
package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
)

// Dosya için hash hesapla (algoritmaya göre)
func fileHash(filename, algo string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var h hash.Hash
	switch algo {
	case "sha1":
		h = sha1.New()
	case "sha256":
		h = sha256.New()
	case "sha512":
		h = sha512.New()
	default:
		return "", fmt.Errorf("desteklenmeyen algoritma: %s", algo)
	}

	if _, err := io.Copy(h, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

func main() {
	// Komut satırı parametreleri
	filename := flag.String("file", "", "Kontrol edilecek dosya")
	expected := flag.String("hash", "", "Beklenen hash (hex formatında)")
	algo := flag.String("algo", "sha256", "Algoritma: sha1, sha256, sha512")
	flag.Parse()

	if *filename == "" || *expected == "" {
		fmt.Println("Kullanım: go run main.go -file=dosya.txt -hash=beklenenhash -algo=sha256")
		os.Exit(1)
	}

	// Dosyanın hash’ini hesapla
	actual, err := fileHash(*filename, *algo)
	if err != nil {
		fmt.Println("Hata:", err)
		os.Exit(1)
	}

	// Karşılaştır
	fmt.Println("Dosya:", *filename)
	fmt.Println("Algoritma:", *algo)
	fmt.Println("Beklenen:", *expected)
	fmt.Println("Hesaplanan:", actual)

	if actual == *expected {
		fmt.Println("✅ Hash eşleşti, dosya DOĞRU.")
	} else {
		fmt.Println("❌ Hash eşleşmedi, dosya DEĞİŞTİRİLMİŞ olabilir!")
	}
}
``
/*
---

## 📌 Kullanım

### 1. SHA1 ile kontrol
*/
``bash
go run main.go -file=example.txt -hash=b1f5c8f2b019a78cddc6a64d3bb1a72f9f2e41a3 -algo=sha1
``

//### 2. SHA256 ile kontrol

``bash
sha256sum example.txt
go run main.go -file=example.txt -hash=<çıkan_hash> -algo=sha256
``

//### 3. SHA512 ile kontrol

``bash
sha512sum example.txt
go run main.go -file=example.txt -hash=<çıkan_hash> -algo=sha512
``
/*
---

📌 Artık aynı programı hem **SHA1** hem **SHA256** hem de **SHA512** ile dosya doğrulamak için kullanabiliyorsun 🚀

👉 İstersen ben buna **otomatik hash hesaplama modu** da ekleyebilirim (yani sadece `-file` ver, program tüm algoritmalarla hash’leri göstersin). Bunu ister misin?
*/