/*
🚀 Şimdi sana **Go’da `crypto/hkdf` paketi**ni ayrıntılı olarak açıklayayım.

---

# 📌 `crypto/hkdf` Paketi Nedir?

Go’daki [`crypto/hkdf`](https://pkg.go.dev/crypto/hkdf) paketi, **HMAC-based Key Derivation Function (HKDF)** algoritmasını uygular.

HKDF, **RFC 5869**’da tanımlanmış bir KDF (Key Derivation Function) olup, genellikle:

* TLS, IPsec, Noise protokolü gibi kriptografik protokollerde,
* Tek bir güçlü “master key”den farklı alt anahtarlar türetmek için
  kullanılır.

Go’da `crypto/hkdf` paketi, HKDF’yi bir **`io.Reader`** gibi uygular. Yani, `hkdf.New` fonksiyonu bir `io.Reader` döndürür; bu reader’dan `Read` ederek istediğin kadar türetilmiş anahtar üretebilirsin.

---
*/
//# 📚 Temel Fonksiyonlar

//### 1. `hkdf.New`

``go
func New(hash func() hash.Hash, secret, salt, info []byte) io.Reader
``
/*
* **hash** → kullanılacak hash fonksiyonu (`sha256.New`, `sha512.New`, vb.)
* **secret** → ana giriş anahtarı (IKM, Input Keying Material)
* **salt** → opsiyonel ek rastgelelik (yoksa boş array)
* **info** → bağlam/veri (anahtarın hangi amaç için kullanılacağını belirtir, ör. “TLS key expansion”)

Döndürdüğü `io.Reader` üzerinden **sınırsız uzunlukta key stream** üretebilirsin.

---

# 🔑 Örnekler

### Örnek 1: Basit HKDF ile 32 byte anahtar türetme
*/
``go
package main

import (
    "crypto/hkdf"
    "crypto/sha256"
    "fmt"
    "io"
)

func main() {
    secret := []byte("supersecretkey")
    salt := []byte("randomsalt")
    info := []byte("example-context")

    // HKDF Reader oluştur
    hkdfReader := hkdf.New(sha256.New, secret, salt, info)

    key := make([]byte, 32) // 32 byte = 256 bit key
    if _, err := io.ReadFull(hkdfReader, key); err != nil {
        panic(err)
    }

    fmt.Printf("Derived key: %x\n", key)
}
``
&*
📌 Çıktı her çalıştırıldığında farklı olur (salt değişirse), 32 byte uzunluğunda bir anahtar elde edilir.

---

### Örnek 2: Birden fazla anahtar türetme
*/

``go
package main

import (
    "crypto/hkdf"
    "crypto/sha512"
    "fmt"
    "io"
)

func main() {
    masterKey := []byte("shared-master-key")
    salt := []byte("protocol-salt")
    info := []byte("handshake data")

    reader := hkdf.New(sha512.New, masterKey, salt, info)

    // Aynı kaynaktan iki ayrı key türetelim
    key1 := make([]byte, 16) // 128-bit
    key2 := make([]byte, 16)

    if _, err := io.ReadFull(reader, key1); err != nil {
        panic(err)
    }
    if _, err := io.ReadFull(reader, key2); err != nil {
        panic(err)
    }

    fmt.Printf("Key1: %x\n", key1)
    fmt.Printf("Key2: %x\n", key2)
}
``
/*
📌 Aynı `io.Reader` üzerinden okudukça ardışık key blokları elde edersin. Böylece tek bir master key’den **Key1, Key2, Key3…** türetilebilir.

---

### Örnek 3: TLS benzeri key expansion

Gerçek TLS gibi düşünelim: bir master secret var ve hem client hem server için ayrı anahtarlar türetmek istiyoruz.
?7

``go
package main

import (
    "crypto/hkdf"
    "crypto/sha256"
    "fmt"
    "io"
)

func main() {
    masterSecret := []byte("TLS-master-secret")
    salt := []byte("tls-salt")

    clientInfo := []byte("client key expansion")
    serverInfo := []byte("server key expansion")

    clientReader := hkdf.New(sha256.New, masterSecret, salt, clientInfo)
    serverReader := hkdf.New(sha256.New, masterSecret, salt, serverInfo)

    clientKey := make([]byte, 16)
    serverKey := make([]byte, 16)

    io.ReadFull(clientReader, clientKey)
    io.ReadFull(serverReader, serverKey)

    fmt.Printf("ClientKey: %x\n", clientKey)
    fmt.Printf("ServerKey: %x\n", serverKey)
}
``
/*
📌 Görüldüğü gibi `info` alanını farklı vererek aynı master secret’tan **bağımsız anahtarlar** elde ediyoruz. Bu HKDF’nin en büyük avantajı: **bağlama (context) duyarlı güvenli key derivation**.

---

# ⚡ Önemli Noktalar

1. `salt` güvenliği artırır. Yoksa tüm sıfırlar (`nil`) kullanılır, ama tavsiye edilen **her zaman rastgele salt** kullanmaktır.
2. `info` aynı secret’tan farklı amaçlar için anahtar türetmek için çok kullanışlıdır (ör: “AES key”, “HMAC key”).
3. HKDF deterministiktir: aynı `(hash, secret, salt, info)` parametreleri → aynı türetilmiş key.
4. `io.Reader` yaklaşımı sayesinde istediğin uzunlukta key stream üretilebilir (ör: 1 KB bile).

---

# ✅ Özet

* `crypto/hkdf` → RFC 5869 tabanlı HKDF implementasyonu.
* Kullanımı: `hkdf.New(hash, secret, salt, info)` → `io.Reader` döner.
* `io.ReadFull` ile istediğin kadar key üretebilirsin.
* `salt` ve `info` kullanarak bağlam ayrımı yapabilirsin.
* Kriptografik protokollerde **master key’den türetilmiş alt anahtarlar** için kullanılır.

---

👉 İstersen sana **HKDF’yi AES şifreleme ile birleştiren bir örnek** de gösterebilirim (örn. master key → AES-256 key türet → metin şifrele). İstiyor musun?
EVET
*/

/*
Harika! O zaman sana **HKDF ile AES-256-GCM şifreleme/deşifreleme** yapan bir örnek hazırlayayım. 🚀

Bu örnekte:

* HKDF ile bir **master key**den 32 byte (256-bit) AES key türeteceğiz.
* AES-GCM ile **mesajı şifreleyeceğiz ve doğrulama tag’ını kontrol edeceğiz**.
* Salt ve info ile güvenliği artıracağız.

---
*/
//# 🔑 `hkdf_aes_gcm.go`

``go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hkdf"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
)

// HKDF ile 32 byte key türet
func deriveKey(masterKey, salt, info []byte) ([]byte, error) {
	h := hkdf.New(sha256.New, masterKey, salt, info)
	key := make([]byte, 32) // AES-256
	if _, err := io.ReadFull(h, key); err != nil {
		return nil, err
	}
	return key, nil
}

// AES-GCM ile şifreleme
func encrypt(plaintext, key []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, err
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	return ciphertext, nonce, nil
}

// AES-GCM ile deşifreleme
func decrypt(ciphertext, nonce, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func main() {
	// Master key
	masterKey := []byte("supersecret-master-key")
	salt := []byte("random-salt")
	info := []byte("aes-256-gcm-example")

	// 1️⃣ HKDF ile AES key türet
	key, err := deriveKey(masterKey, salt, info)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Derived AES-256 key: %x\n", key)

	// 2️⃣ Mesajı şifrele
	plaintext := []byte("Merhaba, HKDF + AES-GCM örneği!")
	ciphertext, nonce, err := encrypt(plaintext, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ciphertext: %x\nNonce: %x\n", ciphertext, nonce)

	// 3️⃣ Mesajı deşifrele
	decrypted, err := decrypt(ciphertext, nonce, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypted: %s\n", decrypted)
}
``
/*
---

# 🔹 Çalışma Mantığı

1. **HKDF** ile master key → AES-256 key (32 byte) türetiyoruz.
2. **AES-GCM** ile şifreleme yapıyoruz: GCM nonce + authentication tag içeriyor.
3. Şifrelenmiş veri + nonce → deşifrelenip doğrulanıyor.

Bu şekilde: **tek bir master key’den güvenli şekilde birçok alt anahtar** türetebilir ve farklı veri bloklarını güvenle şifreleyebilirsin.

---

İstersen bir sonraki adımda **bu HKDF + AES-GCM mantığını bir CLI’ye çevirip dosya şifreleyip/deşifreleyen bir araç** hâline getirebilirim.

Bunu yapayım mı?
EVET

Harika! O zaman sana **HKDF + AES-256-GCM tabanlı dosya şifreleme/deşifreleme CLI aracı** hazırlayalım. 🚀

Bu CLI üç komutla çalışacak:

1. **keygen** → Master key oluşturur ve dosyaya kaydeder.
2. **encrypt** → Dosyayı master key’den türetilmiş AES key ile şifreler.
3. **decrypt** → Şifrelenmiş dosyayı master key’den türetilmiş AES key ile açar.

---
*/

# 🔹 `hkdf_aes_cli.go`

``go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hkdf"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// HKDF ile AES-256 key türet
func deriveKey(masterKey, salt, info []byte) ([]byte, error) {
	h := hkdf.New(sha256.New, masterKey, salt, info)
	key := make([]byte, 32) // AES-256
	if _, err := io.ReadFull(h, key); err != nil {
		return nil, err
	}
	return key, nil
}

// AES-GCM ile şifreleme
func encrypt(plaintext, key []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, err
	}
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	return ciphertext, nonce, nil
}

// AES-GCM ile deşifreleme
func decrypt(ciphertext, nonce, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return aesgcm.Open(nil, nonce, ciphertext, nil)
}

// Master key oluştur ve dosyaya yaz
func keygen(filename string) {
	key := make([]byte, 32) // 256-bit master key
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(filename, []byte(hex.EncodeToString(key)), 0600); err != nil {
		panic(err)
	}
	fmt.Println("Master key oluşturuldu:", filename)
}

// Dosyayı şifrele
func encryptFile(masterKeyFile, inputFile, outputFile string) {
	masterKeyHex, err := ioutil.ReadFile(masterKeyFile)
	if err != nil {
		panic(err)
	}
	masterKey, _ := hex.DecodeString(string(masterKeyHex))

	plaintext, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	salt := []byte("file-encryption-salt")
	info := []byte(inputFile)

	key, _ := deriveKey(masterKey, salt, info)
	ciphertext, nonce, _ := encrypt(plaintext, key)

	// Dosya formatı: nonce | ciphertext
	final := append(nonce, ciphertext...)
	if err := ioutil.WriteFile(outputFile, final, 0644); err != nil {
		panic(err)
	}
	fmt.Println("Şifreleme tamamlandı:", outputFile)
}

// Dosyayı deşifrele
func decryptFile(masterKeyFile, inputFile, outputFile string) {
	masterKeyHex, err := ioutil.ReadFile(masterKeyFile)
	if err != nil {
		panic(err)
	}
	masterKey, _ := hex.DecodeString(string(masterKeyHex))

	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	salt := []byte("file-encryption-salt")
	info := []byte(outputFile) // info aynen encrypt'taki gibi olmalı

	key, _ := deriveKey(masterKey, salt, info)

	block, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(block)
	nonceSize := aesgcm.NonceSize()

	if len(data) < nonceSize {
		panic("geçersiz dosya formatı")
	}

	nonce := data[:nonceSize]
	ciphertext := data[nonceSize:]

	plaintext, err := decrypt(ciphertext, nonce, key)
	if err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(outputFile, plaintext, 0644); err != nil {
		panic(err)
	}
	fmt.Println("Deşifreleme tamamlandı:", outputFile)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Kullanım:")
		fmt.Println("  keygen <master_key_file>")
		fmt.Println("  encrypt <master_key_file> <input_file> <output_file>")
		fmt.Println("  decrypt <master_key_file> <input_file> <output_file>")
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "keygen":
		if len(os.Args) != 3 {
			fmt.Println("keygen <master_key_file>")
			return
		}
		keygen(os.Args[2])
	case "encrypt":
		if len(os.Args) != 5 {
			fmt.Println("encrypt <master_key_file> <input_file> <output_file>")
			return
		}
		encryptFile(os.Args[2], os.Args[3], os.Args[4])
	case "decrypt":
		if len(os.Args) != 5 {
			fmt.Println("decrypt <master_key_file> <input_file> <output_file>")
			return
		}
		decryptFile(os.Args[2], os.Args[3], os.Args[4])
	default:
		fmt.Println("Bilinmeyen komut:", cmd)
	}
}
``
/*
---

# 🔹 Kullanım Örneği
*/
``bash
# 1️⃣ Master key oluştur
go run hkdf_aes_cli.go keygen master.key

# 2️⃣ Dosyayı şifrele
echo "Merhaba, HKDF + AES-GCM CLI örneği!" > mesaj.txt
go run hkdf_aes_cli.go encrypt master.key mesaj.txt mesaj.enc

# 3️⃣ Dosyayı deşifrele
go run hkdf_aes_cli.go decrypt master.key mesaj.enc mesaj_decrypted.txt

# Kontrol et
cat mesaj_decrypted.txt
``
/*
📌 **Notlar:**

* Salt ve info ile HKDF, dosyaya özel key türetiyor.
* AES-GCM nonce dosyanın başına ekleniyor.
* Master key gizli tutulmalı (`0600` izinli dosya).
* Info parametresi encrypt/decrypt aşamasında aynı olmalı; örnekte dosya adı kullanılıyor.

---

Bunu daha ileriye taşıyıp:

* Farklı info/salt kombinasyonlarıyla bir **key hierarchy** oluşturabiliriz,
* Hatta **CLI’ye parola tabanlı master key türetme** ekleyebiliriz (PBKDF2/Argon2 ile).

Bunu ister misin?
*/