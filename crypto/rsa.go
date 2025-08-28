/*
**Go dilinde `crypto/rsa` paketi**ni en temelden, örneklerle ve detaylarıyla açıklayayım.

---

## 📌 `crypto/rsa` Paketi Nedir?

* Go’nun **RSA (Rivest–Shamir–Adleman)** algoritmasını uygulayan paketidir.
* RSA **asimetrik şifreleme** algoritmasıdır:

  * **Public Key (açık anahtar)** → şifreleme için
  * **Private Key (özel anahtar)** → çözme için
* Ayrıca imzalama (**digital signature**) ve doğrulama (**verification**) işlemleri için de kullanılır.
* Genelde **`crypto/rand`** ile birlikte rastgelelik sağlamak için ve **`crypto/sha256`** gibi hash fonksiyonları ile birlikte çalışır.

📌 **RSA günümüzde güvenlidir**, fakat doğru kullanmak gerekir. Mesela **OAEP padding** ile şifreleme, **PSS padding** ile imzalama önerilir.

---

## 📌 Anahtar Fonksiyonlar

### 1. `rsa.GenerateKey(rand io.Reader, bits int)`

* Yeni RSA özel/public key çifti üretir.
* `bits` → genellikle 2048 veya 4096 seçilir.
*/
``go
privKey, err := rsa.GenerateKey(rand.Reader, 2048)
if err != nil {
	panic(err)
}
pubKey := &privKey.PublicKey
``
/*
---

### 2. **Şifreleme – Çözme**

#### `rsa.EncryptOAEP(hash hash.Hash, rand io.Reader, pub *PublicKey, msg []byte, label []byte)`

* Açık anahtar ile OAEP padding kullanarak şifreleme yapar.
* `label` genelde `nil` bırakılır.

#### `rsa.DecryptOAEP(hash hash.Hash, rand io.Reader, priv *PrivateKey, ciphertext []byte, label []byte)`

* Özel anahtar ile şifreyi çözer.
*/
``go
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	// Anahtar üret
	privKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	pubKey := &privKey.PublicKey

	message := []byte("Merhaba RSA!")

	// Şifreleme
	label := []byte("")
	hash := sha256.New()
	ciphertext, _ := rsa.EncryptOAEP(hash, rand.Reader, pubKey, message, label)
	fmt.Printf("Şifreli veri: %x\n", ciphertext)

	// Çözme
	plaintext, _ := rsa.DecryptOAEP(hash, rand.Reader, privKey, ciphertext, label)
	fmt.Printf("Çözülen mesaj: %s\n", plaintext)
}
``
/*
---

### 3. **İmzalama – Doğrulama**

#### `rsa.SignPSS(rand io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte, opts *rsa.PSSOptions)`

* Mesajın hash’ini özel anahtarla imzalar.

#### `rsa.VerifyPSS(pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte, opts *rsa.PSSOptions)`

* İmzayı doğrular.
*/
``go
package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	// Anahtar üret
	privKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	pubKey := &privKey.PublicKey

	message := []byte("Bu bir imza testidir.")

	// Mesajı hashle
	hashed := sha256.Sum256(message)

	// İmzalama
	signature, _ := rsa.SignPSS(rand.Reader, privKey, crypto.SHA256, hashed[:], nil)
	fmt.Printf("İmza: %x\n", signature)

	// Doğrulama
	err := rsa.VerifyPSS(pubKey, crypto.SHA256, hashed[:], signature, nil)
	if err != nil {
		fmt.Println("İmza doğrulama BAŞARISIZ:", err)
	} else {
		fmt.Println("✅ İmza doğrulama BAŞARILI")
	}
}
``
/*
---

### 4. **PKCS#1v15 Modu** (Daha eski, günümüzde önerilmiyor ama destekleniyor)

* **Şifreleme**: `rsa.EncryptPKCS1v15`
* **Çözme**: `rsa.DecryptPKCS1v15`
* **İmzalama**: `rsa.SignPKCS1v15`
* **Doğrulama**: `rsa.VerifyPKCS1v15`

Örnek:
*/
``go
ciphertext, _ := rsa.EncryptPKCS1v15(rand.Reader, pubKey, message)
plaintext, _ := rsa.DecryptPKCS1v15(rand.Reader, privKey, ciphertext)
``
/*
📌 Bu mod backward-compatibility için vardır, modern uygulamalarda **OAEP** ve **PSS** kullanılmalı.

---

### 5. **Anahtar Yönetimi (PEM formatında kaydetme/okuma)**

RSA anahtarları genelde `.pem` dosyalarında saklanır. Bunun için `encoding/pem` ve `crypto/x509` paketleri kullanılır.
*/
``go
// Private Key'i PEM formatında export etme
privBytes := x509.MarshalPKCS1PrivateKey(privKey)
pem.Encode(os.Stdout, &pem.Block{
	Type:  "RSA PRIVATE KEY",
	Bytes: privBytes,
})
``
/*
---

## 📌 Özet

* `rsa.GenerateKey()` → RSA anahtar üretir.
* `rsa.EncryptOAEP()` / `rsa.DecryptOAEP()` → güvenli şifreleme/çözme.
* `rsa.SignPSS()` / `rsa.VerifyPSS()` → güvenli imzalama/doğrulama.
* `PKCS1v15` fonksiyonları eski yöntemler için vardır.
* Anahtarlar genelde `.pem` dosyalarında saklanır.

---

👉 İstersen ben sana bir **RSA ile dosya şifreleme/çözme** örneği de yazabilirim (mesela `.txt` dosyasını şifrele, sonra çöz). Bunu da ister misin?
EVET
*/
/*
Süper 👍 O zaman sana **Go’da RSA ile dosya şifreleme/çözme** örneğini yazayım.

RSA doğrudan **çok büyük dosyaları** şifrelemeye uygun değildir (performans ve limitler nedeniyle).
Genelde şu yöntem uygulanır:

* Küçük mesaj/anahtarlar RSA ile şifrelenir,
* Büyük dosyalar simetrik (AES gibi) ile şifrelenir.

Ama öğrenme amacıyla basitçe RSA ile bir `.txt` dosyası okuyup **şifreleme/çözme** yapalım.

---

## 📌 RSA Dosya Şifreleme / Çözme Örneği
*/
``go
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
)

func encryptFile(filename string, pub *rsa.PublicKey) ([]byte, error) {
	// Dosya oku
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// RSA OAEP ile şifrele
	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, data, nil)
	if err != nil {
		return nil, err
	}

	return ciphertext, nil
}

func decryptFile(ciphertext []byte, priv *rsa.PrivateKey) ([]byte, error) {
	// RSA OAEP ile çöz
	hash := sha256.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, priv, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func main() {
	// RSA anahtar çifti üret
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Anahtar üretilemedi:", err)
		os.Exit(1)
	}
	pubKey := &privKey.PublicKey

	// Örnek plaintext dosyası oluştur
	_ = ioutil.WriteFile("plain.txt", []byte("Merhaba RSA! Bu dosya RSA ile şifrelendi."), 0644)

	// Dosyayı şifrele
	ciphertext, err := encryptFile("plain.txt", pubKey)
	if err != nil {
		fmt.Println("Şifreleme hatası:", err)
		os.Exit(1)
	}
	_ = ioutil.WriteFile("encrypted.bin", ciphertext, 0644)
	fmt.Println("✅ Dosya şifrelendi → encrypted.bin")

	// Şifrelenmiş dosyayı çöz
	decrypted, err := decryptFile(ciphertext, privKey)
	if err != nil {
		fmt.Println("Çözme hatası:", err)
		os.Exit(1)
	}
	_ = ioutil.WriteFile("decrypted.txt", decrypted, 0644)
	fmt.Println("✅ Dosya çözüldü → decrypted.txt")
}
``
/*
---

## 📌 Çalışma Akışı

1. Program `plain.txt` dosyası oluşturur.
2. RSA OAEP ile şifreler → `encrypted.bin` dosyasına yazar.
3. `encrypted.bin` tekrar çözülür → `decrypted.txt` dosyasına yazılır.

### Örnek Çıktı

```
✅ Dosya şifrelendi → encrypted.bin
✅ Dosya çözüldü → decrypted.txt
```

👉 `decrypted.txt` içeriği, `plain.txt` ile **birebir aynı** olur.

---

⚠️ **Not:**
RSA yalnızca \~200 byte civarında veri şifreleyebilir (2048-bit anahtar için). Daha büyük dosyaları şifrelemek için **hibrit şifreleme** yapılır (dosya AES ile şifrelenir, AES anahtarı RSA ile şifrelenir).

---

İstersen sana **AES + RSA hibrit şifreleme sistemi** (büyük dosyaları güvenle şifrelemek için) de yazabilirim. İstiyor musun?
*/