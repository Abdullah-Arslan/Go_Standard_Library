O zaman  **Go `crypto` kütüphanesinin tüm alt paketlerini alfabetik sırayla**, her birini **ne işe yaradığını ve örnek kullanımını** vererek anlatacağım.
Bu, komple bir **cheatsheet** olacak 🚀

---

# 📘 Go `crypto` Paketi – Tüm Alt Paketler

---

## 🔹 1. `crypto/aes`

AES (Advanced Encryption Standard) blok şifreleme algoritması.
*/
``
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	key := []byte("examplekey123456") // 16, 24, 32 byte olmalı
	plaintext := []byte("Merhaba AES")

	block, _ := aes.NewCipher(key)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	io.ReadFull(rand.Reader, iv)

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	fmt.Printf("Şifreli: %x\n", ciphertext)
}
``
/*
---

## 🔹 2. `crypto/cipher`

Blok şifreleri için arayüzler (CBC, CFB, CTR, GCM).

👉 AES, DES vb. ile birlikte kullanılır.
*/
``
// AES-GCM örneği
block, _ := aes.NewCipher(key)
gcm, _ := cipher.NewGCM(block)
nonce := make([]byte, gcm.NonceSize())
io.ReadFull(rand.Reader, nonce)
ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
``
/*
---

## 🔹 3. `crypto/des`

DES ve 3DES (Triple DES). Günümüzde güvenli sayılmaz.
*/
``
package main

import (
	"crypto/des"
	"fmt"
)

func main() {
	key := []byte("8bytekey")
	block, _ := des.NewCipher(key)
	fmt.Println("Blok boyutu:", block.BlockSize())
}
``
/*
---

## 🔹 4. `crypto/dsa`

DSA (Digital Signature Algorithm). Yerini genelde ECDSA aldı.
*/
``
 Genelde doğrudan kullanılmaz, `crypto/x509` ile birlikte çalışır.
``
/*
---

## 🔹 5. `crypto/ecdsa`

Elliptic Curve Digital Signature Algorithm.
*/
``
package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	msg := []byte("ECDSA test")
	hash := sha256.Sum256(msg)

	r, s, _ := ecdsa.Sign(rand.Reader, priv, hash[:])
	valid := ecdsa.Verify(&priv.PublicKey, hash[:], r, s)

	fmt.Println("İmza doğrulandı mı?", valid)
}
``
/*
---

## 🔹 6. `crypto/ed25519`

Modern, hızlı imzalama algoritması.
*/
``
package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
)

func main() {
	pub, priv, _ := ed25519.GenerateKey(rand.Reader)
	msg := []byte("Ed25519 örneği")

	sig := ed25519.Sign(priv, msg)
	ok := ed25519.Verify(pub, msg, sig)

	fmt.Println("Doğrulandı mı?", ok)
}
``
/*
---

## 🔹 7. `crypto/elliptic`

ECDSA’da kullanılan eğriler (P224, P256, P384, P521).
*/
``
curve := elliptic.P256()
x, y := curve.ScalarBaseMult([]byte{1,2,3})
fmt.Println("Nokta:", x, y)
``
/*
---

## 🔹 8. `crypto/hmac`

HMAC (Hash-based Message Authentication Code).
*/
``
h := hmac.New(sha256.New, []byte("secret"))
h.Write([]byte("mesaj"))
fmt.Printf("HMAC: %x\n", h.Sum(nil))
``
/*
---

## 🔹 9. `crypto/md5`

MD5 hash fonksiyonu (güvenli değil!).
*/
``
sum := md5.Sum([]byte("test"))
fmt.Printf("MD5: %x\n", sum)
``

/*
---

## 🔹 10. `crypto/rand`

Kriptografik güvenli rastgele sayı üretici.
*/
``
b := make([]byte, 16)
rand.Read(b)
fmt.Printf("Random: %x\n", b)
``
/*
---

## 🔹 11. `crypto/rc4`

RC4 stream cipher (güvenlik için önerilmez).
*/
``
c, _ := rc4.NewCipher([]byte("key"))
msg := []byte("hello")
c.XORKeyStream(msg, msg)
fmt.Printf("Şifreli: %x\n", msg)
``

/*
---

## 🔹 12. `crypto/rsa`

RSA şifreleme ve imzalama.
*/

``
priv, _ := rsa.GenerateKey(rand.Reader, 2048)
msg := []byte("RSA mesajı")
hash := sha256.Sum256(msg)
sig, _ := rsa.SignPKCS1v15(rand.Reader, priv, 0, hash[:])
fmt.Println("İmza:", sig)
``

/*
---

## 🔹 13. `crypto/sha1`

SHA-1 hash (artık güvenli değil).
*/
``
h := sha1.Sum([]byte("test"))
fmt.Printf("SHA1: %x\n", h)
``
/*
---

## 🔹 14. `crypto/sha256`

SHA-224 ve SHA-256.
*/

``
h := sha256.Sum256([]byte("test"))
fmt.Printf("SHA256: %x\n", h)
``
/*
---

## 🔹 15. `crypto/sha512`

SHA-384 ve SHA-512.
*/
``
h := sha512.Sum512([]byte("test"))
fmt.Printf("SHA512: %x\n", h)
``

/*
---

## 🔹 16. `crypto/subtle`

Zaman sabitli karşılaştırma (timing attack önlemek için).
*/
``
ok := subtle.ConstantTimeCompare([]byte("a"), []byte("a")) == 1
fmt.Println("Sonuç:", ok)
``
/*
---

## 🔹 17. `crypto/tls`

TLS (SSL) protokolü implementasyonu.
*/
``
cer, _ := tls.LoadX509KeyPair("server.crt", "server.key")
config := &tls.Config{Certificates: []tls.Certificate{cer}}
ln, _ := tls.Listen("tcp", ":8443", config)
fmt.Println("TLS sunucu başladı:", ln.Addr())
``

/*
---

## 🔹 18. `crypto/x509`

X.509 sertifikaları işlemek.
*/
``
block, _ := pem.Decode(certPEMData)
cert, _ := x509.ParseCertificate(block.Bytes)
fmt.Println("Sertifika CN:", cert.Subject.CommonName)
``

/*
---

## 🔹 19. `crypto/x509/pkix`

PKIX yapıları (sertifika Subject, Issuer gibi alanlar).
*/
``
pkixName := pkix.Name{CommonName: "example.com"}
fmt.Println("CN:", pkixName.CommonName)
``
/*
---

# ✅ Özet Tablosu

| Paket              | Açıklama                            |
| ------------------ | ----------------------------------- |
| `crypto/aes`       | AES blok şifre                      |
| `crypto/cipher`    | Şifreleme modları (CBC, CTR, GCM)   |
| `crypto/des`       | DES, 3DES                           |
| `crypto/dsa`       | DSA imza                            |
| `crypto/ecdsa`     | ECDSA imza                          |
| `crypto/ed25519`   | Ed25519 imza                        |
| `crypto/elliptic`  | Eliptik eğri matematiği             |
| `crypto/hmac`      | HMAC                                |
| `crypto/md5`       | MD5 hash (zayıf)                    |
| `crypto/rand`      | Kripto güvenli rastgele sayı        |
| `crypto/rc4`       | RC4 cipher (zayıf)                  |
| `crypto/rsa`       | RSA şifreleme & imza                |
| `crypto/sha1`      | SHA-1 hash (zayıf)                  |
| `crypto/sha256`    | SHA-224, SHA-256                    |
| `crypto/sha512`    | SHA-384, SHA-512                    |
| `crypto/subtle`    | Timing-attack güvenli karşılaştırma |
| `crypto/tls`       | TLS protokolü                       |
| `crypto/x509`      | Sertifika işleme                    |
| `crypto/x509/pkix` | Sertifika alan yapıları             |

---

🔐 Böylece Go’nun **crypto paketinin tüm alt paketlerini** ve her birinin **örneklerini** görmüş olduk.

👉 İstersen ben sana bunların içinden **en çok kullanılan kombinasyonları (örn. AES + HMAC, RSA + SHA256, TLS + X509)** gerçek dünyada nasıl kullanıldığını gösteren **mini projeler** yazabilirim. İster misin?
*/
/*
👉 İstersen ben sana bunların içinden **en çok kullanılan kombinasyonları (örn. AES + HMAC, RSA + SHA256, TLS + X509)** gerçek dünyada nasıl kullanıldığını gösteren **mini projeler** yazabilirim. İster misin?

O zaman  **Go `crypto` kütüphanesinin en çok kullanılan kombinasyonlarını** göstereyim.
Bunlar gerçek dünyada **şifreleme, imzalama, TLS, sertifika işlemleri** için kullanılan hazır “mini proje” örnekleri olacak 🚀

---

# 📘 1. AES + HMAC → Güvenli Mesaj Şifreleme

AES ile veriyi şifrele, HMAC ile bütünlüğünü doğrula.
*/
``
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
)

func encryptAndSign(key, msg []byte) ([]byte, []byte) {
	block, _ := aes.NewCipher(key)
	iv := make([]byte, aes.BlockSize)
	io.ReadFull(rand.Reader, iv)

	ciphertext := make([]byte, len(msg))
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext, msg)

	// HMAC ile imzala
	mac := hmac.New(sha256.New, key)
	mac.Write(ciphertext)
	signature := mac.Sum(nil)

	// iv + ciphertext döndür
	return append(iv, ciphertext...), signature
}

func main() {
	key := []byte("examplekey123456") // 16, 24, 32 byte olabilir
	message := []byte("Gizli Mesaj")

	ciphertext, signature := encryptAndSign(key, message)

	fmt.Printf("Şifreli veri: %x\n", ciphertext)
	fmt.Printf("HMAC: %x\n", signature)
}
``
/*
📌 **Kullanım Alanı:** Güvenli dosya saklama, token imzalama.

---

# 📘 2. RSA + SHA256 → İmzalama & Doğrulama

RSA anahtarıyla imza at ve doğrula.
*/
``
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	// Anahtar üret
	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	pub := &priv.PublicKey

	message := []byte("Bu bir mesajdır")
	hash := sha256.Sum256(message)

	// İmzala
	sig, _ := rsa.SignPKCS1v15(rand.Reader, priv, 0, hash[:])
	fmt.Printf("İmza: %x\n", sig)

	// Doğrula
	err := rsa.VerifyPKCS1v15(pub, 0, hash[:], sig)
	if err != nil {
		fmt.Println("İmza geçersiz")
	} else {
		fmt.Println("İmza doğrulandı ✔")
	}
}
``
/*
📌 **Kullanım Alanı:** Dijital imzalar, lisans doğrulama, yazılım imzalama.

---

# 📘 3. TLS Sunucu + X.509 Sertifika

Self-signed sertifika ile basit bir HTTPS server.
*/
``
package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "TLS ile güvenli bağlantı!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	server := &http.Server{
		Addr:    ":8443",
		Handler: mux,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}

	fmt.Println("HTTPS sunucu 8443 portunda çalışıyor...")
	// server.crt ve server.key dosyaları olmalı
	err := server.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		panic(err)
	}
}
``
/*
📌 **Kullanım Alanı:** HTTPS API, güvenli web uygulamaları.

---

# 📘 4. Ed25519 → Modern İmzalama

Hızlı ve güvenli imza algoritması.
*/
``
package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
)

func main() {
	pub, priv, _ := ed25519.GenerateKey(rand.Reader)

	msg := []byte("Merhaba Ed25519")
	sig := ed25519.Sign(priv, msg)

	if ed25519.Verify(pub, msg, sig) {
		fmt.Println("İmza doğrulandı ✔")
	} else {
		fmt.Println("İmza geçersiz ❌")
	}
}
``
/*
📌 **Kullanım Alanı:** SSH anahtarları, blockchain imzaları, modern API güvenliği.

---

# 📘 5. AES-GCM → Authenticated Encryption

AES-GCM, hem **şifreleme hem doğrulama** yapar (HMAC gerekmez).
*/
``
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	key := []byte("examplekey123456")
	plaintext := []byte("AES-GCM deneme")

	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)

	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	fmt.Printf("Şifreli + Auth: %x\n", ciphertext)

	// Çözme
	nonce, data := ciphertext[:gcm.NonceSize()], ciphertext[gcm.NonceSize():]
	plaintext2, _ := gcm.Open(nil, nonce, data, nil)
	fmt.Println("Çözüldü:", string(plaintext2))
}
``
/*
📌 **Kullanım Alanı:** Güvenli mesajlaşma, dosya şifreleme.

---

# ✅ Sonuç

Go `crypto` ile en yaygın kombinasyonlar:

1. **AES + HMAC** → Dosya/token güvenliği
2. **RSA + SHA256** → Dijital imza
3. **TLS + X.509** → HTTPS, güvenli ağ
4. **Ed25519** → Modern hızlı imzalama
5. **AES-GCM** → Hem şifreleme hem bütünlük doğrulama

---
*/