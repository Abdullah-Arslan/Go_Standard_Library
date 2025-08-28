/*
Şimdi Go’daki **ECDH (Elliptic Curve Diffie-Hellman) paketi ve kullanımını** detaylı olarak ele alalım.

---

# 📌 ECDH Nedir?

* **ECDH (Elliptic Curve Diffie-Hellman)**, **elliptic curve tabanlı anahtar değişim protokolü**dür.
* Amaç: İki tarafın **güvenli bir ortak gizli anahtar (shared secret)** oluşturmasıdır.
* ECDH kendisi **şifreleme yapmaz**, sadece **ortak anahtar üretir**.
* Genellikle AES gibi simetrik şifreleme algoritmaları ile birlikte kullanılır.

---

# 📌 Go’da ECDH

* Go 1.20+ ile birlikte **`crypto/ecdh`** paketi geldi
* Bu paket, **elliptic curve key exchange** işlemleri için modern bir API sağlar

## Önemli Tipler ve Fonksiyonlar

| Tip / Fonksiyon                        | Açıklama                                          |
| -------------------------------------- | ------------------------------------------------- |
| `ecdh.PrivateKey`                      | ECDH özel anahtar                                 |
| `ecdh.PublicKey`                       | ECDH genel anahtar                                |
| `ecdh.P256()`                          | P-256 elliptic curve (NIST)                       |
| `ecdh.GenerateKey(curve, rand.Reader)` | Anahtar çifti oluşturur                           |
| `PrivateKey.PublicKey()`               | Genel anahtarı alır                               |
| `PrivateKey.ECDH(peerPublic)`          | Peer’in public key’i ile **shared secret** üretir |

> Not: Go 1.20 öncesinde ECDH için üçüncü parti paketler veya `crypto/elliptic` kullanılırdı.

---

# 📌 1️⃣ Basit ECDH Örneği
*/

``go
package main

import (
	"crypto/ecdh"
	"crypto/rand"
	"fmt"
)

func main() {
	// 1️⃣ Alice anahtar çifti
	alicePriv, _ := ecdh.P256().GenerateKey(rand.Reader)
	alicePub := alicePriv.PublicKey()

	// 2️⃣ Bob anahtar çifti
	bobPriv, _ := ecdh.P256().GenerateKey(rand.Reader)
	bobPub := bobPriv.PublicKey()

	// 3️⃣ Ortak gizli anahtarlar
	aliceShared, _ := alicePriv.ECDH(bobPub)
	bobShared, _ := bobPriv.ECDH(alicePub)

	fmt.Printf("Alice Shared Key: %x\n", aliceShared)
	fmt.Printf("Bob Shared Key:   %x\n", bobShared)

	// Aynı mı kontrol
	if string(aliceShared) == string(bobShared) {
		fmt.Println("Ortak anahtarlar eşleşiyor ✅")
	} else {
		fmt.Println("Ortak anahtarlar eşleşmiyor ❌")
	}
}
``
/*
---

# 📌 Açıklamalar

1. **`GenerateKey`** → elliptic curve tabanlı private/public key çifti oluşturur
2. **`PublicKey()`** → private key’den public key elde edilir
3. **`ECDH(peerPublic)`** → peer’in public key’i ile ortak gizli anahtar üretir
4. Alice ve Bob farklı key çifti ile başlar ama **ECDH sayesinde aynı ortak anahtar elde edilir**

---

# 📌 Kullanım Senaryoları

1. Ortak anahtar oluşturma → AES gibi simetrik şifreleme için
2. TLS, SSH, VPN ve güvenli protokollerde anahtar değişimi
3. Mesajlaşma uygulamalarında **end-to-end encryption (E2EE)**

---

✅ Özetle:

* `crypto/ecdh` → ECDH key exchange API’si
* Anahtar üret, paylaş, `ECDH()` ile ortak anahtar üret
* Şifreleme için simetrik algoritmalarla birleştir

---

İstersen bir sonraki adımda **ECDH ile AES kullanarak gerçek bir mesaj şifreleme ve deşifreleme örneği** yapabilirim.

Bunu yapayım mı?
EVET


Harika 😄 O zaman şimdi Go’da **ECDH ile AES kullanarak mesaj şifreleme ve deşifreleme örneği** yapalım.

Bu örnek:

* Alice ve Bob ECDH ile ortak anahtar oluşturacak
* Ortak anahtar AES-256 için kullanılacak
* AES-CBC modunda şifreleme ve deşifreleme yapılacak

---

# 📌 `main.go` – ECDH + AES Mini Örneği
*/

``go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdh"
	"crypto/rand"
	"fmt"
)

func main() {
	message := []byte("Hello Go ECDH + AES!")

	// ---------------- Alice anahtar çifti ----------------
	alicePriv, _ := ecdh.P256().GenerateKey(rand.Reader)
	alicePub := alicePriv.PublicKey()

	// ---------------- Bob anahtar çifti ----------------
	bobPriv, _ := ecdh.P256().GenerateKey(rand.Reader)
	bobPub := bobPriv.PublicKey()

	// ---------------- Ortak gizli anahtar ----------------
	aliceShared, _ := alicePriv.ECDH(bobPub)
	bobShared, _ := bobPriv.ECDH(alicePub)

	fmt.Printf("Alice Shared Key: %x\n", aliceShared)
	fmt.Printf("Bob Shared Key:   %x\n", bobShared)

	// AES için 32 byte key üret (P256 ECDH output genellikle 32 byte)
	aesKey := aliceShared[:32] // Alice veya Bob fark etmez

	// ---------------- AES-CBC şifreleme ----------------
	block, _ := aes.NewCipher(aesKey)
	iv := make([]byte, aes.BlockSize)
	rand.Read(iv)

	ciphertext := make([]byte, len(message))
	mode := cipher.NewCBCEncrypter(block, iv)

	// Mesaj uzunluğunu blok boyutuna göre ayarla (padding)
	padLen := aes.BlockSize - len(message)%aes.BlockSize
	paddedMsg := append(message, make([]byte, padLen)...)

	mode.CryptBlocks(ciphertext, paddedMsg)
	fmt.Printf("\nAES-CBC şifreli: %x\n", ciphertext)

	// ---------------- AES-CBC deşifreleme ----------------
	decrypted := make([]byte, len(ciphertext))
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypter.CryptBlocks(decrypted, ciphertext)

	// Padding’i kaldır
	decrypted = decrypted[:len(message)]
	fmt.Printf("AES-CBC deşifreli: %s\n", decrypted)
}
``
/*
---

# 📌 Açıklamalar

1. **ECDH ile ortak anahtar**: Alice ve Bob farklı key çifti kullanıyor ama `ECDH()` ile aynı anahtarı üretiyorlar
2. **AES-CBC modunda şifreleme**: Ortak anahtar AES key olarak kullanılıyor
3. **Padding**: CBC modunda blok boyutu (16 byte) uyumu için mesaj padding uygulanıyor
4. **Deşifreleme**: CBC deşifreleme ile padding kaldırılarak orijinal mesaj elde ediliyor

---

# 📌 Çıktı Örneği

```
Alice Shared Key: 1f2a3b4c5d6e7f8a...
Bob Shared Key:   1f2a3b4c5d6e7f8a...
AES-CBC şifreli: 8d5f2a9b4c1e3f...
AES-CBC deşifreli: Hello Go ECDH + AES!
```

---

✅ Artık ECDH ile **güvenli bir ortak anahtar** üretip, AES ile **mesaj şifreleme ve deşifreleme** yapabiliyorsun.

İstersen bir sonraki adımda bunu **CLI tabanlı interaktif bir ECDH + AES şifreleme aracı** hâline getirebiliriz, kullanıcı mesaj ve mod seçebilecek.

Bunu yapayım mı?
*/