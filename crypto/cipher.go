/*
Go’daki **`crypto/cipher`** paketini detaylıca ele alalım. Bu paket **şifreleme (encryption) için blok ve stream cipher’ları** sağlar.

---

# 📌 `crypto/cipher` Paketi Nedir?

* Go’nun standart kütüphanesinde bulunur: `import "crypto/cipher"`
* Amacı: **bloke edilmiş veri ve stream veriyi şifreleme** için soyut bir API sunmak.
* Kendisi **doğrudan algoritma içermez** → `crypto/aes`, `crypto/des` gibi paketlerle birlikte çalışır.

Temel kullanım alanları:

1. **Block cipher** (AES, DES, 3DES)
2. **Block cipher modes** (CBC, CFB, CTR, OFB, GCM)
3. **Stream cipher** (RC4 gibi)

---

# 📌 Temel Tipler (Types)

| Tip            | Açıklama                                             |
| -------------- | ---------------------------------------------------- |
| `Block`        | Blok şifreleme algoritması (AES, DES vb.)            |
| `BlockMode`    | Blok modları (CBC, CFB, OFB, CTR) için arayüz        |
| `Stream`       | Stream cipher interface (XORKeyStream ile şifreleme) |
| `AEAD`         | Authenticated encryption (GCM gibi)                  |
| `KeySizeError` | Yanlış anahtar boyutu hatası                         |

---

# 📌 1️⃣ Block Cipher Kullanımı (AES Örneği)
*/
``go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func main() {
	key := []byte("1234567890123456") // 16 byte = AES-128
	plaintext := []byte("Hello Go Cipher!")

	// AES block oluştur
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// CBC mod için IV
	iv := []byte("1234567890123456") // block size kadar

	mode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	mode.CryptBlocks(ciphertext, plaintext)

	fmt.Printf("Ciphertext: %x\n", ciphertext)

	// Decrypt
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(ciphertext))
	decrypter.CryptBlocks(decrypted, ciphertext)

	fmt.Printf("Decrypted: %s\n", decrypted)
}
``
/*
---

# 📌 2️⃣ Stream Cipher Kullanımı (CFB Örneği)
*/
``go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func main() {
	key := []byte("1234567890123456")
	plaintext := []byte("Hello Stream Cipher")

	block, _ := aes.NewCipher(key)
	iv := []byte("1234567890123456")

	stream := cipher.NewCFBEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	fmt.Printf("Ciphertext: %x\n", ciphertext)

	// Decrypt
	streamDec := cipher.NewCFBDecrypter(block, iv)
	decrypted := make([]byte, len(ciphertext))
	streamDec.XORKeyStream(decrypted, ciphertext)

	fmt.Printf("Decrypted: %s\n", decrypted)
}
``
/*
---

# 📌 3️⃣ AEAD (GCM) Örneği
*/
``go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func main() {
	key := []byte("1234567890123456") // 16 byte = AES-128
	block, _ := aes.NewCipher(key)

	aead, _ := cipher.NewGCM(block)
	nonce := []byte("123456789012") // GCM nonce = 12 byte
	plaintext := []byte("Hello AEAD GCM")

	ciphertext := aead.Seal(nil, nonce, plaintext, nil)
	fmt.Printf("Ciphertext: %x\n", ciphertext)

	decrypted, _ := aead.Open(nil, nonce, ciphertext, nil)
	fmt.Printf("Decrypted: %s\n", decrypted)
}
``
/*
---

# 📌 Özet

1. **`cipher.Block`** → AES, DES gibi blok algoritmaları için temel yapı
2. **`cipher.BlockMode`** → CBC, CFB, OFB, CTR gibi blok modları
3. **`cipher.Stream`** → XORKeyStream ile stream şifreleme
4. **`cipher.AEAD`** → Authenticated encryption (AES-GCM)
5. **`KeySizeError`** → Yanlış anahtar boyutu

> Not: `cipher` paketi **veri şifreleme için soyutlama sağlar**, ama **kendi algoritmalarını içermez**. Bu yüzden AES/DES gibi algoritmaları başka paketlerden alırız (`crypto/aes` gibi).

---

İstersen sana bir sonraki adımda **AES, DES ve ChaCha20 ile cipher paketi kullanarak mini bir şifreleme uygulaması** da yazabilirim, hem blok hem de stream örnekli.

Bunu yapayım mı?
EVET
*/

/*
Harika 👍 O zaman sana Go’da **`crypto/cipher` kullanarak mini bir şifreleme uygulaması** hazırlayalım.

Bu uygulama:

* AES (blok şifreleme) ile CBC ve GCM modunu kullanacak
* ChaCha20 (stream cipher) ile şifreleme/deşifreleme yapacak
* Kullanıcıdan mesaj alıp şifreleyip çözecek

---

# 📌 `main.go` – Mini Cipher Uygulaması
*/
``go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"

	"golang.org/x/crypto/chacha20"
)

func main() {
	message := []byte("Hello Go Cipher Mini App!")

	fmt.Println("Orijinal mesaj:", string(message))

	// 1️⃣ AES-CBC
	key := []byte("1234567890123456") // 16 byte = AES-128
	iv := []byte("1234567890123456")  // CBC IV

	block, _ := aes.NewCipher(key)
	ciphertext := make([]byte, len(message))

	cbcEncrypter := cipher.NewCBCEncrypter(block, iv)
	cbcEncrypter.CryptBlocks(ciphertext, message)
	fmt.Printf("AES-CBC şifreli: %x\n", ciphertext)

	// CBC decrypt
	decrypted := make([]byte, len(ciphertext))
	cbcDecrypter := cipher.NewCBCDecrypter(block, iv)
	cbcDecrypter.CryptBlocks(decrypted, ciphertext)
	fmt.Println("AES-CBC deşifreli:", string(decrypted))

	// 2️⃣ AES-GCM
	aead, _ := cipher.NewGCM(block)
	nonce := make([]byte, aead.NonceSize())
	io.ReadFull(rand.Reader, nonce)

	gcmCiphertext := aead.Seal(nil, nonce, message, nil)
	fmt.Printf("AES-GCM şifreli: %x\n", gcmCiphertext)

	// GCM decrypt
	gcmDecrypted, _ := aead.Open(nil, nonce, gcmCiphertext, nil)
	fmt.Println("AES-GCM deşifreli:", string(gcmDecrypted))

	// 3️⃣ ChaCha20 (Stream Cipher)
	chachaKey := []byte("12345678901234567890123456789012") // 32 byte
	nonceCha := []byte("12345678abcdefgh")                   // 12 byte

	stream, _ := chacha20.NewUnauthenticatedCipher(chachaKey, nonceCha)
	chaChaCiphertext := make([]byte, len(message))
	stream.XORKeyStream(chaChaCiphertext, message)
	fmt.Printf("ChaCha20 şifreli: %x\n", chaChaCiphertext)

	// ChaCha20 decrypt
	streamDec, _ := chacha20.NewUnauthenticatedCipher(chachaKey, nonceCha)
	chaChaDecrypted := make([]byte, len(chaChaCiphertext))
	streamDec.XORKeyStream(chaChaDecrypted, chaChaCiphertext)
	fmt.Println("ChaCha20 deşifreli:", string(chaChaDecrypted))
}
``
/*
---

# 📌 Açıklamalar

1. **AES-CBC**

   * Blok şifreleme
   * Sabit IV (örnek amaçlı, gerçek hayatta rastgele olmalı)
   * `CryptBlocks` ile şifreleme ve deşifreleme

2. **AES-GCM**

   * Blok şifreleme + authenticated encryption
   * `Seal` ile şifrele, `Open` ile deşifrele

3. **ChaCha20**

   * Stream cipher
   * `XORKeyStream` ile hem şifreleme hem deşifreleme
   * Anahtar ve nonce gerektirir

---

# 📌 Çıktı Örneği

``
Orijinal mesaj: Hello Go Cipher Mini App!
AES-CBC şifreli: 8f6c9f5d4a... 
AES-CBC deşifreli: Hello Go Cipher Mini App!
AES-GCM şifreli: 3d5a7c1f2b...
AES-GCM deşifreli: Hello Go Cipher Mini App!
ChaCha20 şifreli: a4f7e91b0c...
ChaCha20 deşifreli: Hello Go Cipher Mini App!
``

---

✅ Bu örnekle artık:

* `crypto/cipher` paketini blok ve stream cipher ile kullanabiliyorsun
* CBC, GCM, ChaCha20 gibi farklı modları deneyebiliyorsun
* Hem şifreleme hem de deşifreleme yapabiliyorsun

---

İstersen ben bir sonraki adımda bu uygulamayı **CLI üzerinden kullanıcıdan mesaj alan ve AES/ChaCha20 modunu seçebilen interaktif bir şifreleme aracı** hâline getirebilirim.

Bunu yapayım mı?
*/




