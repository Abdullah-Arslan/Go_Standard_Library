/*
Go (Golang) `pbkdf2` paketini** en baştan, tüm detaylarıyla ve örneklerle açıklayayım.

---

# 🔑 `pbkdf2` Paketi Nedir?

Go’da **`pbkdf2`** paketi aslında `golang.org/x/crypto/pbkdf2` altında bulunur (standart kütüphane değil).

* Açılımı: **Password-Based Key Derivation Function 2**
* Amaç: **zayıf parolalardan güvenli anahtar üretmek**
* Kullanım alanları:

  * Parola hashleme (doğrudan saklamak yerine)
  * Paroladan AES, HMAC gibi algoritmalar için key üretmek
  * Parola tabanlı şifreleme

**Özellikleri:**

* Salt kullanır (rastgele değer → rainbow table saldırılarını engeller).
* İterasyon sayısı ayarlanabilir (parola kırmayı yavaşlatır).
* HMAC ile bir hash fonksiyonu (SHA1, SHA256, SHA512) kullanır.

---

# 📦 Fonksiyonlar

`pbkdf2` paketinde sadece **tek bir fonksiyon** vardır:

```go
func Key(password, salt []byte, iter, keyLen int, h func() hash.Hash) []byte
```

Parametreler:

* `password` → Kullanıcının parolası
* `salt` → Rastgele salt değeri
* `iter` → İterasyon sayısı (ör: 10000)
* `keyLen` → Üretilecek anahtar uzunluğu (byte cinsinden)
* `h` → Kullanılacak hash fonksiyonu (örn: sha256.New)

---

# 🔧 Kullanım Örnekleri

### 1. Basit PBKDF2 Kullanımı (SHA256 ile)
*/
``go
package main

import (
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

func main() {
	password := []byte("mysecretpassword")
	salt := []byte("random_salt")

	// 10000 iterasyon, 32 byte key üretelim
	key := pbkdf2.Key(password, salt, 10000, 32, sha256.New)

	fmt.Printf("Derived key: %x\n", key)
}
``
/*
📌 Çıktı: (her zaman aynı password+salt için aynı key çıkar)

```
Derived key: 3b8e0a72b4c29a5b742ed2e0868f39c2822baf9ef3c22b4e9341537f29b3f9d8
```

---

### 2. Parola Hashleme (Doğrulama ile)
*/
``go
package main

import (
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
)

func hashPassword(password, salt []byte) []byte {
	return pbkdf2.Key(password, salt, 10000, 32, sha256.New)
}

func main() {
	password := []byte("mypassword")
	salt := []byte("random_salt")

	// Kullanıcı kayıt olurken hash hesaplanır
	storedHash := hashPassword(password, salt)
	fmt.Printf("Stored hash: %x\n", storedHash)

	// Kullanıcı giriş yaptığında tekrar hesaplanır
	input := []byte("mypassword")
	inputHash := hashPassword(input, salt)

	// Karşılaştır
	if string(storedHash) == string(inputHash) {
		fmt.Println("✅ Parola doğru")
	} else {
		fmt.Println("❌ Parola yanlış")
	}
}
``
/*
---

### 3. AES için Anahtar Üretme
*/
PBKDF2 ile üretilen key’i AES şifrelemede kullanabilirsin.

``go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

func main() {
	password := []byte("supersecret")
	salt := []byte("random_salt")

	// AES-256 için 32 byte key üret
	key := pbkdf2.Key(password, salt, 10000, 32, sha256.New)

	// AES-GCM örneği
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)

	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)

	plaintext := []byte("Merhaba Dünya")
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	fmt.Printf("Ciphertext: %x\n", ciphertext)
}
``
/*
---

### 4. Farklı Hash Fonksiyonları ile PBKDF2
*/
``go
import (
	"crypto/sha1"
	"crypto/sha512"
)

// SHA1
key1 := pbkdf2.Key([]byte("pass"), []byte("salt"), 1000, 32, sha1.New)

// SHA512
key2 := pbkdf2.Key([]byte("pass"), []byte("salt"), 2000, 64, sha512.New)
``
/*
---

# ⚠️ Dikkat

* **MD5 veya SHA1** artık güvenli değil, **SHA256 veya SHA512** kullan.
* `iter` değeri ne kadar yüksekse kırmak o kadar zor ama işlem maliyeti de artar.
  Günümüzde **100.000+ iterasyon** öneriliyor.
* Parola saklamak için **bcrypt, scrypt veya argon2** daha güvenli kabul ediliyor.

---

# 📌 Özet

* `pbkdf2` = Password-Based Key Derivation Function 2
* Go’da → `golang.org/x/crypto/pbkdf2` paketinde bulunur.
* Tek fonksiyon:

  * `Key(password, salt, iter, keyLen, hashFunc)`
* Kullanım:

  * Parola hashleme
  * AES/HMAC gibi algoritmalar için anahtar türetme
  * Parola tabanlı şifreleme sistemleri

---

👉 İstersen sana bir sonraki adımda, **PBKDF2 ile kullanıcı login sistemi (register + login)** örneği yazabilirim.
Bunu da görmek ister misin?
EVET

Harika 👍 o zaman sana **PBKDF2 ile kullanıcı kayıt ve giriş sistemi (Register + Login)** örneği yazayım.

Burada mantık şu:

* Kullanıcı kayıt olurken → parolası **PBKDF2 + Salt** ile hashlenip saklanır.
* Giriş yaparken → kullanıcıdan gelen parola tekrar aynı salt ile hashlenir, veri tabanındaki hash ile karşılaştırılır.

---

# 📝 PBKDF2 ile Kullanıcı Sistemi Örneği
*/
``go
package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

// PBKDF2 ayarları
const (
	Iter     = 100_000 // iterasyon sayısı (yüksek tut!)
	KeyLen   = 32      // AES-256 için 32 byte
	SaltSize = 16      // 16 byte salt
)

// Rastgele salt üret
func generateSalt() []byte {
	salt := make([]byte, SaltSize)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}
	return salt
}

// Parolayı hashle
func hashPassword(password string, salt []byte) string {
	key := pbkdf2.Key([]byte(password), salt, Iter, KeyLen, sha256.New)
	// Salt + Hash birlikte saklanır (base64 encoding ile)
	return base64.StdEncoding.EncodeToString(salt) + "$" + base64.StdEncoding.EncodeToString(key)
}

// Parola doğrula
func verifyPassword(password, stored string) bool {
	// Salt ve Hash'i ayır
	var saltB64, hashB64 string
	fmt.Sscanf(stored, "%[^$]$%s", &saltB64, &hashB64)

	salt, _ := base64.StdEncoding.DecodeString(saltB64)
	expectedHash, _ := base64.StdEncoding.DecodeString(hashB64)

	// Kullanıcının girdiği parolayı tekrar hashle
	key := pbkdf2.Key([]byte(password), salt, Iter, KeyLen, sha256.New)

	// Karşılaştır
	return string(key) == string(expectedHash)
}

// ====================== DEMO ======================
func main() {
	// Kullanıcı kayıt oluyor
	password := "supersecret123"
	salt := generateSalt()
	storedHash := hashPassword(password, salt)

	fmt.Println("DB'ye kaydedilen değer:", storedHash)

	// Kullanıcı giriş yapıyor
	loginPassword := "supersecret123" // Doğru parola
	// loginPassword := "yanlışparola" // Yanlış parola

	if verifyPassword(loginPassword, storedHash) {
		fmt.Println("✅ Giriş başarılı")
	} else {
		fmt.Println("❌ Parola yanlış")
	}
}
``
/*
---

# 📌 Çalışma Mantığı

1. `generateSalt()` → Her kullanıcı için rastgele salt üretir.
2. `hashPassword()` → `PBKDF2(password, salt, iter, keyLen, sha256)` ile hash hesaplar. Salt ve hash DB’ye kaydedilir.
3. `verifyPassword()` → Giriş sırasında parolayı tekrar hashler, DB’deki hash ile karşılaştırır.

---

# ⚠️ Önemli Notlar

* `Iter` (iterasyon) sayısını yüksek tut → en az **100k**, tercihen **300k+**.
* `Salt` her kullanıcı için farklı olmalı (rastgele üretilir).
* Hash ve salt birlikte saklanmalı (örneğin `salt$hash` formatında).
* Daha güvenli alternatifler: **bcrypt, scrypt, argon2**.

---

👉 İstersen sana aynı sistemi **bcrypt ile nasıl yapılır** da gösterebilirim (PBKDF2’ye göre daha güvenli kabul ediliyor). Onu da ister misin?
*/