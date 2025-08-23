/*
Go (Golang) `crypto/md5` paketini** en baştan tüm detaylarıyla ve örneklerle açıklayayım.

---

# 🔑 `crypto/md5` Paketi Nedir?

Go’daki [`crypto/md5`](https://pkg.go.dev/crypto/md5) paketi, **MD5 (Message-Digest Algorithm 5)** hash algoritmasını uygular.

* MD5 → 128 bit (16 byte) uzunluğunda bir özet (hash) üretir.
* Çoğunlukla **dosya bütünlüğü kontrolü (checksum)** için kullanılır.
* **Kriptografik güvenlik için önerilmez**, çünkü **çakışmalar (collision)** bulunmuştur.
  → Yani güvenlik açısından SHA-256 veya SHA-512 tercih edilmelidir.

---

# 📦 Paketin İçeriği

`crypto/md5` paketi `hash.Hash` arayüzünü implemente eder.
İçinde şu fonksiyonlar vardır:

1. **`func New() hash.Hash`**
   → Yeni bir MD5 hash hesaplayıcı döner.

2. **`func Sum(data []byte) [16]byte`**
   → Verilen verinin MD5 hash’ini tek seferde hesaplar.

---

# 🔧 Kullanım Örnekleri

### 1. `md5.Sum` ile Tek Seferde Hash Alma
*/
``go
package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	data := []byte("Merhaba Dünya")
	hash := md5.Sum(data)

	// md5.Sum -> [16]byte döner
	fmt.Printf("MD5: %x\n", hash)
}
``
/*
📌 Çıktı:

```
MD5: 7ddf32e17a6ac5ce04a8ecbf782ca509
```

---

### 2. `md5.New` ile Adım Adım Hash Hesaplama

Büyük veriler veya akış (stream) işlenirken kullanılır.
*/

``go
package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

func main() {
	h := md5.New()

	// Veri parça parça yazılabilir
	io.WriteString(h, "Merhaba ")
	io.WriteString(h, "Dünya")

	// Hash sonucu
	hash := h.Sum(nil)

	fmt.Printf("MD5: %x\n", hash)
}
``
/*
📌 Çıktı:

``
MD5: 7ddf32e17a6ac5ce04a8ecbf782ca509
``

---

### 3. Dosya Hash’i Hesaplama

Dosyanın bütünlüğünü kontrol etmek için sık kullanılır.
*/
``go
package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	h := md5.New()

	// Dosyayı hash’e yaz
	if _, err := io.Copy(h, file); err != nil {
		panic(err)
	}

	// Sonucu al
	hash := h.Sum(nil)
	fmt.Printf("Dosyanın MD5 hash'i: %x\n", hash)
}
``
/*
---

### 4. Hex String Olarak Hash Döndürme

Hash sonucu genelde hex string olarak saklanır.
*/
``go
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	data := []byte("golang")
	hash := md5.Sum(data)

	// Hex string
	hashStr := hex.EncodeToString(hash[:])
	fmt.Println("MD5 Hash:", hashStr)
}
``
/*
📌 Çıktı:

```
MD5 Hash: 8ab2d368b8dc7f2c6bc64c94d6ed94d9
```

---

### 5. Kullanıcı Parola Hashleme (Önerilmez ⚠️)

Eskiden MD5 parolalar için kullanılırdı, fakat **güvenli değildir**.
Sadece eğitim amaçlı göstereyim:
*/
``go
func hashPassword(password string) string {
	hash := md5.Sum([]byte(password))
	return fmt.Sprintf("%x", hash)
}
``
/*
📌 Gerçek uygulamalarda: `bcrypt`, `scrypt`, `argon2` gibi güçlü algoritmalar kullanılmalıdır.

---

# 📌 Özet

* `crypto/md5` → MD5 hash algoritmasını uygular.
* `md5.Sum(data)` → Tek seferde hash alır.
* `md5.New()` → `hash.Hash` döner, akış (stream) verilerde kullanılır.
* MD5 hızlıdır ama **kriptografik güvenlik için uygun değildir**.
* Daha güvenli alternatifler: **SHA-256 (`crypto/sha256`)**, **SHA-512 (`crypto/sha512`)**.

---

İstersen sana bir sonraki adımda, **MD5 ile dosya kontrolü (örn: indirdiğin bir ISO dosyasının hash karşılaştırması)** örneğini de gösterebilirim. Bunu ister misin?
EVET


Harika 👍 o zaman sana **MD5 ile dosya bütünlüğü (checksum) kontrolü** örneğini göstereyim.

Bu yöntem genellikle:

* Bir dosya indirdiğinde (örn. ISO, ZIP, EXE)
* Yayıncının verdiği **MD5 checksum** ile
* Senin indirdiğin dosyanın MD5 hash’ini karşılaştırarak dosyanın bozulmadığından emin olmak için kullanılır.

---

# 🛠 MD5 ile Dosya Kontrolü Örneği
*/
``go
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// Dosyanın MD5 hash'ini hesapla
func fileMD5(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	h := md5.New()

	// Dosyayı hash fonksiyonuna kopyala
	if _, err := io.Copy(h, file); err != nil {
		return "", err
	}

	// Hash sonucunu hex string olarak döndür
	return hex.EncodeToString(h.Sum(nil)), nil
}

func main() {
	// Kontrol edilecek dosya
	filePath := "example.iso"

	// Yayıncının verdiği resmi MD5 checksum (örnek)
	expectedMD5 := "5eb63bbbe01eeed093cb22bb8f5acdc3"

	// Bizim hesapladığımız MD5
	actualMD5, err := fileMD5(filePath)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	fmt.Println("Beklenen MD5:", expectedMD5)
	fmt.Println("Hesaplanan MD5:", actualMD5)

	// Karşılaştırma
	if actualMD5 == expectedMD5 {
		fmt.Println("✅ Dosya bütünlüğü doğrulandı. (Bozulmamış)")
	} else {
		fmt.Println("❌ Dosya bütünlüğü bozulmuş veya değiştirilmiş!")
	}
}
``
/*
---

# 📌 Çalışma Mantığı

1. Dosya `io.Copy` ile hash fonksiyonuna aktarılır.
2. `md5.New()` → `hash.Hash` oluşturur.
3. Hash sonucu hex string’e çevrilir.
4. Beklenen MD5 değeriyle karşılaştırılır.

---

# 🎯 Kullanım Senaryosu

* Bir ISO dosyası indiriyorsun: `ubuntu.iso`
* Web sitesinde şu yazıyor:

```
MD5: d41d8cd98f00b204e9800998ecf8427e
```

* Senin hesapladığın değer aynıysa → ✅ dosya sağlamdır.
* Farklıysa → ❌ dosya bozulmuş ya da oynanmıştır.

---

Bunu biraz daha ileri götürüp istersen ben sana aynı örneği **SHA256 ile yapmayı** gösterebilirim (günümüzde MD5 yerine SHA256 tercih ediliyor). Bunu da görmek ister misin?
*/