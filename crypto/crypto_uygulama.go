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