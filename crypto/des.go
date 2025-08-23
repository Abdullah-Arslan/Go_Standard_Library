/*
Şimdi Go’da **`crypto/des`** paketini detaylı olarak ele alalım. Bu paket **DES (Data Encryption Standard) algoritmasını** ve varyasyonlarını sağlar.

---

# 📌 `crypto/des` Paketi Nedir?

* Go’nun standart kütüphanesinde bulunur: `import "crypto/des"`
* Amaç: **DES, Triple DES (3DES) blok şifreleme algoritmalarını** uygulamak
* Tipik kullanım: `crypto/cipher` ile birlikte **blok şifreleme modlarında (CBC, CFB, OFB, CTR)** çalışır
* DES algoritması **64-bit blok boyutuna** sahiptir

---

# 📌 Temel Fonksiyonlar ve Tipler

| Fonksiyon / Tip                                        | Açıklama                                        |
| ------------------------------------------------------ | ----------------------------------------------- |
| `NewCipher(key []byte) (cipher.Block, error)`          | DES blok şifresi oluşturur (8 byte anahtar)     |
| `NewTripleDESCipher(key []byte) (cipher.Block, error)` | 3DES şifreleme (16 veya 24 byte anahtar)        |
| `BlockSize()`                                          | Blok boyutunu döner (DES: 8 byte, 3DES: 8 byte) |

> Not: DES **güvenlik açısından zayıf** sayılır. 3DES biraz daha güvenlidir. Modern uygulamalarda AES tercih edilir.

---

# 📌 1️⃣ DES Örneği (CBC Modu)
*/

``go
package main

import (
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

func main() {
	key := []byte("12345678")           // 8 byte = DES
	plaintext := []byte("Hello DES!!") // 8 byte blok

	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// CBC mod için IV
	iv := []byte("abcdefgh") // 8 byte
	mode := cipher.NewCBCEncrypter(block, iv)

	ciphertext := make([]byte, len(plaintext))
	mode.CryptBlocks(ciphertext, plaintext)
	fmt.Printf("DES-CBC şifreli: %x\n", ciphertext)

	// Decrypt
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(ciphertext))
	decrypter.CryptBlocks(decrypted, ciphertext)
	fmt.Printf("DES-CBC deşifreli: %s\n", decrypted)
}
``
/*
> DES blok boyutu 8 byte olduğu için mesaj uzunluğu **8’in katı olmalıdır**.
> Eğer değilse padding (PKCS5/PKCS7) uygulanmalıdır.

---

# 📌 2️⃣ 3DES Örneği (Triple DES)
*/

``go
package main

import (
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

func main() {
	key := []byte("123456781234567812345678") // 24 byte = 3DES
	plaintext := []byte("Hello TripleDES!")  // 16 byte

	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(err)
	}

	iv := []byte("abcdefgh") // 8 byte IV
	mode := cipher.NewCBCEncrypter(block, iv)

	ciphertext := make([]byte, len(plaintext))
	mode.CryptBlocks(ciphertext, plaintext)
	fmt.Printf("3DES-CBC şifreli: %x\n", ciphertext)

	// Decrypt
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(ciphertext))
	decrypter.CryptBlocks(decrypted, ciphertext)
	fmt.Printf("3DES-CBC deşifreli: %s\n", decrypted)
}
``
/*
---

# 📌 Özet

1. `des.NewCipher(key)` → DES blok şifreleyici oluşturur (8 byte anahtar)
2. `des.NewTripleDESCipher(key)` → 3DES oluşturur (16 veya 24 byte)
3. DES blok boyutu **8 byte**
4. `crypto/cipher` ile **CBC, CFB, CTR, OFB** modlarında kullanılır
5. DES tek başına artık güvenli değil → sadece eski sistemlerle uyum için önerilir
6. Modern uygulamalarda **AES** tercih edilmelidir

---

İstersen sana bir sonraki adımda **DES ve 3DES’i birlikte CBC ve CFB modlarıyla şifreleyip çözen mini bir uygulama** hazırlayabilirim.

Bunu yapayım mı?
EVET
*/

/*
Harika 👍 O zaman sana Go’da **DES ve 3DES’i CBC ve CFB modlarıyla şifreleyip çözen mini bir uygulama** hazırlayalım.

---

# 📌 `main.go` – DES ve 3DES Mini Şifreleme Uygulaması
*/
``go
package main

import (
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

func main() {
	message := []byte("Hello Go DES/3DES Cipher!")

	fmt.Println("Orijinal mesaj:", string(message))

	// ---------------- DES ----------------
	desKey := []byte("12345678") // 8 byte
	desIV := []byte("abcdefgh")  // 8 byte

	desBlock, _ := des.NewCipher(desKey)

	// DES-CBC
	cbcCipher := cipher.NewCBCEncrypter(desBlock, desIV)
	cbcEncrypted := make([]byte, len(message))
	copyCbcMsg := message[:8] // DES blok 8 byte, örnek amaçlı
	cbcCipher.CryptBlocks(cbcEncrypted[:8], copyCbcMsg)
	fmt.Printf("DES-CBC şifreli (ilk blok): %x\n", cbcEncrypted[:8])

	// DES-CFB
	cfbCipher := cipher.NewCFBEncrypter(desBlock, desIV)
	cfbEncrypted := make([]byte, len(message))
	cfbCipher.XORKeyStream(cfbEncrypted, message)
	fmt.Printf("DES-CFB şifreli: %x\n", cfbEncrypted)

	// ---------------- 3DES ----------------
	tripleKey := []byte("123456781234567812345678") // 24 byte
	tripBlock, _ := des.NewTripleDESCipher(tripleKey)
	tripIV := []byte("abcdefgh") // 8 byte

	// 3DES-CBC
	tripCbc := cipher.NewCBCEncrypter(tripBlock, tripIV)
	tripCbcEncrypted := make([]byte, len(message))
	copyTripMsg := message[:16] // 3DES blok örneği
	tripCbc.CryptBlocks(tripCbcEncrypted[:16], copyTripMsg)
	fmt.Printf("3DES-CBC şifreli (ilk iki blok): %x\n", tripCbcEncrypted[:16])

	// 3DES-CFB
	tripCfb := cipher.NewCFBEncrypter(tripBlock, tripIV)
	tripCfbEncrypted := make([]byte, len(message))
	tripCfb.XORKeyStream(tripCfbEncrypted, message)
	fmt.Printf("3DES-CFB şifreli: %x\n", tripCfbEncrypted)

	// ---------------- Decryption Örnek ----------------
	// DES-CFB deşifreleme
	desCfbDecrypter := cipher.NewCFBDecrypter(desBlock, desIV)
	desCfbDecrypted := make([]byte, len(cfbEncrypted))
	desCfbDecrypter.XORKeyStream(desCfbDecrypted, cfbEncrypted)
	fmt.Println("DES-CFB deşifreli:", string(desCfbDecrypted))

	// 3DES-CFB deşifreleme
	tripCfbDecrypter := cipher.NewCFBDecrypter(tripBlock, tripIV)
	tripCfbDecrypted := make([]byte, len(tripCfbEncrypted))
	tripCfbDecrypter.XORKeyStream(tripCfbDecrypted, tripCfbEncrypted)
	fmt.Println("3DES-CFB deşifreli:", string(tripCfbDecrypted))
}
``
/*
---

# 📌 Açıklamalar

1. **DES ve 3DES** için ayrı anahtar ve IV kullandık
2. **CBC modunda**: Blok blok şifreleme (padding gerekebilir, örnek amaçlı sadece ilk blok)
3. **CFB modunda**: Stream benzeri şifreleme, mesaj boyutu fark etmez
4. **Decryption**: CBC ve CFB modları için ayrı deşifreleme objesi oluşturulur

> Not: CBC modunda mesaj uzunluğu mutlaka blok boyutunun katı olmalı. Gerçek uygulamada **PKCS5/PKCS7 padding** uygulanır.

---

✅ Bu örnekle artık:

* DES ve 3DES’i **CBC ve CFB modlarıyla** kullanabiliyorsun
* Hem şifreleme hem deşifreleme yapabiliyorsun
* Blok ve stream modları arasındaki farkı görebiliyorsun

---

İstersen bir sonraki adımda bu uygulamayı **CLI üzerinden kullanıcıdan mesaj alan, mod ve algoritmayı seçebilen interaktif bir şifreleme aracı** hâline getirebilirim.

Bunu yapayım mı?
EVET
*/

/*
Harika 😄 O zaman şimdi Go’da **CLI tabanlı interaktif bir şifreleme aracı** yapalım.

Bu araç:

* Kullanıcıdan mesaj alacak
* Şifreleme algoritmasını seçecek: DES, 3DES veya AES
* Modu seçecek: CBC, CFB, GCM (AES için), CBC, CFB (DES/3DES için)
* Mesajı şifreleyecek ve deşifreleyecek

---

# 📌 `main.go` – Interaktif Şifreleme CLI
*/
``go
package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Go Mini Cipher CLI ===")
	fmt.Print("Mesajınızı girin: ")
	message, _ := reader.ReadString('\n')
	message = message[:len(message)-1] // '\n' kaldır

	fmt.Print("Algoritma seçin (DES / 3DES / AES): ")
	alg, _ := reader.ReadString('\n')
	alg = alg[:len(alg)-1]

	fmt.Print("Mod seçin (CBC / CFB / GCM-AES sadece): ")
	mod, _ := reader.ReadString('\n')
	mod = mod[:len(mod)-1]

	var ciphertext, decrypted []byte

	switch alg {
	case "DES":
		key := []byte("12345678")   // 8 byte
		iv := []byte("abcdefgh")    // 8 byte
		block, _ := des.NewCipher(key)

		switch mod {
		case "CBC":
			ciphertext = make([]byte, 8) // örnek amaçlı ilk blok
			cbc := cipher.NewCBCEncrypter(block, iv)
			copyMsg := []byte(message)
			if len(copyMsg) > 8 {
				copyMsg = copyMsg[:8]
			}
			cbc.CryptBlocks(ciphertext, copyMsg)

			decrypted = make([]byte, 8)
			cbcDec := cipher.NewCBCDecrypter(block, iv)
			cbcDec.CryptBlocks(decrypted, ciphertext)

		case "CFB":
			ciphertext = make([]byte, len(message))
			cfb := cipher.NewCFBEncrypter(block, iv)
			cfb.XORKeyStream(ciphertext, []byte(message))

			decrypted = make([]byte, len(ciphertext))
			cfbDec := cipher.NewCFBDecrypter(block, iv)
			cfbDec.XORKeyStream(decrypted, ciphertext)
		}

	case "3DES":
		key := []byte("123456781234567812345678") // 24 byte
		iv := []byte("abcdefgh")                 // 8 byte
		block, _ := des.NewTripleDESCipher(key)

		switch mod {
		case "CBC":
			ciphertext = make([]byte, 16) // ilk 2 blok
			copyMsg := []byte(message)
			if len(copyMsg) > 16 {
				copyMsg = copyMsg[:16]
			}
			cbc := cipher.NewCBCEncrypter(block, iv)
			cbc.CryptBlocks(ciphertext, copyMsg)

			decrypted = make([]byte, 16)
			cbcDec := cipher.NewCBCDecrypter(block, iv)
			cbcDec.CryptBlocks(decrypted, ciphertext)

		case "CFB":
			ciphertext = make([]byte, len(message))
			cfb := cipher.NewCFBEncrypter(block, iv)
			cfb.XORKeyStream(ciphertext, []byte(message))

			decrypted = make([]byte, len(ciphertext))
			cfbDec := cipher.NewCFBDecrypter(block, iv)
			cfbDec.XORKeyStream(decrypted, ciphertext)
		}

	case "AES":
		key := []byte("1234567890123456") // 16 byte
		block, _ := aes.NewCipher(key)
		iv := []byte("1234567890123456") // 16 byte

		switch mod {
		case "CBC":
			ciphertext = make([]byte, len(message))
			cbc := cipher.NewCBCEncrypter(block, iv)
			copyMsg := []byte(message)
			if len(copyMsg)%aes.BlockSize != 0 {
				copyMsg = append(copyMsg, make([]byte, aes.BlockSize-len(copyMsg)%aes.BlockSize)...)
			}
			cbc.CryptBlocks(ciphertext, copyMsg)

			decrypted = make([]byte, len(ciphertext))
			cbcDec := cipher.NewCBCDecrypter(block, iv)
			cbcDec.CryptBlocks(decrypted, ciphertext)

		case "CFB":
			ciphertext = make([]byte, len(message))
			cfb := cipher.NewCFBEncrypter(block, iv)
			cfb.XORKeyStream(ciphertext, []byte(message))

			decrypted = make([]byte, len(ciphertext))
			cfbDec := cipher.NewCFBDecrypter(block, iv)
			cfbDec.XORKeyStream(decrypted, ciphertext)

		case "GCM":
			aead, _ := cipher.NewGCM(block)
			nonce := iv[:aead.NonceSize()]
			ciphertext = aead.Seal(nil, nonce, []byte(message), nil)
			decrypted, _ = aead.Open(nil, nonce, ciphertext, nil)
		}
	}

	fmt.Printf("\nŞifreli: %x\n", ciphertext)
	fmt.Printf("Deşifreli: %s\n", decrypted)
}
``
/*
---

# 📌 Kullanım Örneği

``
=== Go Mini Cipher CLI ===
Mesajınızı girin: Hello World
Algoritma seçin (DES / 3DES / AES): AES
Mod seçin (CBC / CFB / GCM-AES sadece): CFB

Şifreli: 3a1f5c...
Deşifreli: Hello World
``

---

# 📌 Özellikler

* DES, 3DES, AES algoritmalarını destekler
* CBC, CFB modları (AES için GCM de ekledik)
* Hem şifreleme hem de deşifreleme yapar
* Kullanıcıdan **interaktif** olarak mesaj, algoritma ve modu seçme imkanı

---

İstersen ben bir sonraki adımda **ChaCha20 ve AES-CTR gibi modern stream cipher’ları da ekleyip tam bir “Go Cipher Toolkit” CLI uygulaması** haline getirebilirim.

Bunu yapayım mı?
*/