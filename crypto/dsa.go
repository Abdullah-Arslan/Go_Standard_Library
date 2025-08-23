/*
Şimdi Go’daki **`crypto/dsa`** paketini detaylıca ele alalım.

---

# 📌 `crypto/dsa` Paketi Nedir?

* Go’nun standart kütüphanesinde bulunur: `import "crypto/dsa"`
* **DSA (Digital Signature Algorithm)** için fonksiyonlar sağlar
* Amaç: **veri bütünlüğünü ve kimlik doğrulamasını** sağlamak
* DSA, **asenkron imza algoritması**dır (RSA gibi, fakat sadece imza için kullanılır)
* `crypto/rand` ve `math/big` ile birlikte çalışır

> Not: DSA günümüzde eski sayılır, genellikle **ECDSA** veya **RSA** tercih edilir.

---

# 📌 Temel Tipler ve Fonksiyonlar

| Tip / Fonksiyon          | Açıklama                         |
| ------------------------ | -------------------------------- |
| `dsa.Parameters`         | DSA parametreleri (P, Q, G)      |
| `dsa.PrivateKey`         | Özel anahtar, `PublicKey` içerir |
| `dsa.PublicKey`          | Genel anahtar                    |
| `dsa.GenerateParameters` | Parametreleri oluşturur          |
| `dsa.GenerateKey`        | Özel ve genel anahtar oluşturur  |
| `dsa.Sign`               | Mesaj imzalar                    |
| `dsa.Verify`             | Mesaj imzasını doğrular          |

---

# 📌 1️⃣ Basit DSA Örneği
*/

``go
package main

import (
	"crypto/dsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	var params dsa.Parameters

	// Parametreleri oluştur
	dsa.GenerateParameters(&params, rand.Reader, dsa.L1024N160)

	// Özel ve genel anahtar oluştur
	priv := new(dsa.PrivateKey)
	priv.Parameters = params
	dsa.GenerateKey(priv, rand.Reader)
	pub := &priv.PublicKey

	// Mesaj
	message := []byte("Hello Go DSA")
	hash := sha256.Sum256(message)

	// İmza oluştur
	r, s, _ := dsa.Sign(rand.Reader, priv, hash[:])
	fmt.Printf("İmza: r=%s, s=%s\n", r.String(), s.String())

	// İmza doğrula
	valid := dsa.Verify(pub, hash[:], r, s)
	fmt.Println("İmza geçerli mi?", valid)
}
``
/*
---

# 📌 Açıklamalar

1. **`dsa.GenerateParameters`**

   * DSA parametreleri (P, Q, G) oluşturur
   * `L1024N160`, `L2048N224`, `L2048N256`, `L3072N256` seçenekleri var

2. **`dsa.GenerateKey`**

   * Özel ve genel anahtarları oluşturur

3. **`dsa.Sign` ve `dsa.Verify`**

   * `Sign`: mesaj hash’ini imzalar
   * `Verify`: imzanın geçerliliğini kontrol eder

4. **Hash Fonksiyonu**

   * DSA direkt mesajı imzalamaz, **hash kullanılır** (`sha1`, `sha256`)

---

# 📌 Özet

* `crypto/dsa` → DSA imza algoritması
* Kullanım: **Digital Signature**
* Eski ve artık sınırlı kullanımı var, modern uygulamalarda **ECDSA veya RSA** tercih edilir
* Temel akış:

  1. Parametre oluştur (`dsa.GenerateParameters`)
  2. Anahtar oluştur (`dsa.GenerateKey`)
  3. Mesajı hash’le
  4. İmzala (`dsa.Sign`)
  5. Doğrula (`dsa.Verify`)

---

İstersen ben sana bir sonraki adımda **DSA ile dosya veya metin imzalayan ve doğrulayan mini bir CLI uygulaması** örneği hazırlayabilirim.

Bunu yapayım mı?
EVET
*/

/*
Harika 😄 O zaman sana Go’da **DSA ile metin imzalayan ve doğrulayan mini bir CLI uygulaması** hazırlayalım.

---

# 📌 `main.go` – DSA CLI Uygulaması
*/

``go
package main

import (
	"bufio"
	"crypto/dsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Go DSA CLI ===")
	fmt.Print("İmzalanacak mesajı girin: ")
	message, _ := reader.ReadString('\n')
	message = strings.TrimSpace(message)

	// 1️⃣ Parametreleri oluştur
	var params dsa.Parameters
	dsa.GenerateParameters(&params, rand.Reader, dsa.L1024N160)

	// 2️⃣ Özel ve genel anahtar oluştur
	priv := new(dsa.PrivateKey)
	priv.Parameters = params
	dsa.GenerateKey(priv, rand.Reader)
	pub := &priv.PublicKey

	// 3️⃣ Mesaj hash
	hash := sha256.Sum256([]byte(message))

	// 4️⃣ Mesajı imzala
	r, s, _ := dsa.Sign(rand.Reader, priv, hash[:])
	fmt.Printf("\nİmza oluşturuldu:\nr = %s\ns = %s\n", r.String(), s.String())

	// 5️⃣ İmza doğrula
	valid := dsa.Verify(pub, hash[:], r, s)
	if valid {
		fmt.Println("İmza geçerli ✅")
	} else {
		fmt.Println("İmza geçersiz ❌")
	}

	// 6️⃣ Doğrulama testi (kullanıcı değiştirme)
	fmt.Print("\nMesajı değiştirip doğrulamak ister misiniz? (e/h): ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	if choice == "e" {
		fmt.Print("Yeni mesajı girin: ")
		newMsg, _ := reader.ReadString('\n')
		newMsg = strings.TrimSpace(newMsg)
		newHash := sha256.Sum256([]byte(newMsg))
		valid = dsa.Verify(pub, newHash[:], r, s)
		if valid {
			fmt.Println("İmza geçerli ✅")
		} else {
			fmt.Println("İmza geçersiz ❌")
		}
	}
}
``
/*
---

# 📌 Özellikler

1. Kullanıcıdan mesaj alır
2. DSA parametreleri ve anahtarları oluşturur
3. Mesajı SHA-256 ile hash’ler
4. `dsa.Sign` ile imzalar
5. `dsa.Verify` ile imzayı doğrular
6. İsteğe bağlı olarak kullanıcı mesajı değiştirip imza geçerliliğini test edebilir

---

# 📌 Örnek Çalıştırma

```
=== Go DSA CLI ===
İmzalanacak mesajı girin: Merhaba Go
İmza oluşturuldu:
r = 12345678901234567890
s = 98765432109876543210
İmza geçerli ✅

Mesajı değiştirip doğrulamak ister misiniz? (e/h): e
Yeni mesajı girin: Merhaba Dünya
İmza geçersiz ❌
```

---

Bu CLI ile artık **DSA imzalama ve doğrulama** adımlarını interaktif olarak görebiliyorsun.

İstersen ben bir sonraki adımda **DSA + dosya imzalama/okuma desteği** ekleyip **tam dosya imza aracı** haline getirebilirim.

Bunu yapayım mı?
*/