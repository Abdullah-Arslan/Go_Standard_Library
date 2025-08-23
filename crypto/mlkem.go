/*
Go’nun standart kütüphanesinde veya bilinen popüler paketlerinde **`mlkem`** adında bir paket bulunmuyor 📦❌.
Muhtemelen sen ya:

1. **Yazım hatası** yaptın (örneğin `mkem` → *Module Kem* ya da `mlkem` yerine başka bir şey kastettin),
2. Ya da **özel bir üçüncü parti kripto paketi** (örneğin NIST post-quantum yarışmasındaki **ML-KEM (Module Lattice KEM, eski adı Kyber)**) hakkında soruyorsun.

---

## 🔑 ML-KEM (Kyber) Nedir?

Eğer kastettiğin **ML-KEM**, bu aslında Go’nun değil, **NIST PQC (Post-Quantum Cryptography)** standardında yer alan **Kyber** algoritmasının resmi adı.

* ML-KEM = **Module Lattice Key Encapsulation Mechanism**
* NIST tarafından seçilen kuantuma dayanıklı KEM (Key Encapsulation Mechanism).
* Klasik RSA/ECDH yerine **kuantuma dayanıklı anahtar değişimi** sağlar.
* 2024 itibarıyla, birçok dilde (C, Rust, Python) uygulamaları var. Go’da da bazı üçüncü parti kütüphaneler mevcut.

---

## 📦 Go’da ML-KEM Kullanımı

Go’nun standart kütüphanesinde yok ama şu paketlerle kullanılabiliyor:

* [pq-crystals/kyber](https://github.com/pq-crystals/kyber) (resmi referans C kodu)
* Go binding’leri: örneğin [`github.com/kudelskisecurity/crystals-go`](https://github.com/kudelskisecurity/crystals-go)

Go ile tipik kullanım (üçüncü parti paket üzerinden) şöyle olabilir:
*/
``go
package main

import (
	"fmt"
	"log"

	"github.com/kudelskisecurity/crystals-go/kyber"
)

func main() {
	// Parametreler: Kyber512, Kyber768, Kyber1024 (güvenlik seviyeleri)
	scheme := kyber.Kyber512

	// Anahtar çifti üret
	pk, sk, err := scheme.GenerateKeyPair()
	if err != nil {
		log.Fatal(err)
	}

	// Kapsülleme (encapsulation): shared secret + ciphertext üret
	ct, ssEncap, err := scheme.Encapsulate(pk)
	if err != nil {
		log.Fatal(err)
	}

	// Açma (decapsulation): shared secret geri kazan
	ssDecap, err := scheme.Decapsulate(ct, sk)
	if err != nil {
		log.Fatal(err)
	}

	// Aynı shared secret elde edilmesi gerekir
	fmt.Printf("Encapsulated SS: %x\n", ssEncap)
	fmt.Printf("Decapsulated SS: %x\n", ssDecap)
	fmt.Println("Eşleşti mi?:", string(ssEncap) == string(ssDecap))
}
``
/*
📌 Burada:

* `Encapsulate` → Alıcı için şifreli bir **session key** üretir.
* `Decapsulate` → Alıcı, kendi gizli anahtarı ile aynı session key’i geri çıkarır.
* Bu şekilde kuantuma dayanıklı **anahtar değişimi** yapılmış olur.

---

## 📌 Özet

* Go **standart kütüphanesinde `mlkem` paketi yok**.
* Ama eğer kastettiğin **ML-KEM (Kyber)** ise:

  * Bu **kuantuma dayanıklı anahtar değişim algoritmasıdır**.
  * Go’da üçüncü parti kütüphanelerle kullanılabilir.
  * Temel fonksiyonları: `GenerateKeyPair`, `Encapsulate`, `Decapsulate`.

---

👉 Sana iki soru sorayım:
Senin sorduğun `mlkem` **Go’nun kendi paketi** mi sanmıştın, yoksa gerçekten **NIST ML-KEM (Kyber)** algoritmasını mı öğrenmek istiyorsun?
EVET
*/


/*
Anladım 👍 o zaman senin sorduğun şey aslında **ML-KEM (Kyber)** yani NIST’in **kuantuma dayanıklı (post-quantum)** algoritması.
Go’nun kendi **standart kütüphanesinde `mlkem` paketi yok**, ama üçüncü parti kütüphaneler ile kullanabiliyoruz.

---

# 🔑 ML-KEM (Kyber) Nedir?

* Açılımı: **Module Lattice Key Encapsulation Mechanism**
* Eski adı: **Kyber**
* NIST tarafından **kuantuma dayanıklı anahtar değişim (KEM)** algoritması olarak seçildi.
* RSA veya ECDH gibi klasik yöntemlerin yerini alması için tasarlandı.
* 3 güvenlik seviyesi var:

  * **Kyber512** (AES-128 güvenliğine denk)
  * **Kyber768** (AES-192 güvenliğine denk)
  * **Kyber1024** (AES-256 güvenliğine denk)

---

# 📦 Go’da ML-KEM Kullanımı

Go için resmi `crypto/mlkem` paketi yok ama bazı üçüncü parti kütüphaneler var. Örneğin:

* [`github.com/kudelskisecurity/crystals-go`](https://github.com/kudelskisecurity/crystals-go)

---

# 🔧 Örnek Kullanım (Kyber512)
*/
``go
package main

import (
	"fmt"
	"log"

	"github.com/kudelskisecurity/crystals-go/kyber"
)

func main() {
	// Güvenlik seviyesi seçelim (Kyber512, Kyber768, Kyber1024)
	scheme := kyber.Kyber512

	// 1. Anahtar çifti üret
	pk, sk, err := scheme.GenerateKeyPair()
	if err != nil {
		log.Fatal("Anahtar üretilemedi:", err)
	}

	// 2. Kapsülleme (Encapsulation) -> Shared secret ve ciphertext üret
	ct, ssEncap, err := scheme.Encapsulate(pk)
	if err != nil {
		log.Fatal("Encapsulation hatası:", err)
	}

	// 3. Açma (Decapsulation) -> Shared secret geri elde edilir
	ssDecap, err := scheme.Decapsulate(ct, sk)
	if err != nil {
		log.Fatal("Decapsulation hatası:", err)
	}

	// 4. Sonuçları yazdır
	fmt.Printf("Encapsulated Shared Secret: %x\n", ssEncap)
	fmt.Printf("Decapsulated Shared Secret: %x\n", ssDecap)
	fmt.Println("Aynı mı?:", string(ssEncap) == string(ssDecap))
}
``
/*
---

# 📌 Çalışma Mantığı

1. **GenerateKeyPair** → `pk` (public key) + `sk` (secret key)
2. **Encapsulate(pk)** → Gönderen taraf, alıcıya göndereceği şifreli session key (`ct`) ve kendi shared secret (`ssEncap`) üretir.
3. **Decapsulate(ct, sk)** → Alıcı taraf, kendi secret key’i ile aynı shared secret’ı (`ssDecap`) çıkarır.
4. Sonuç: İki taraf aynı shared secret’a sahiptir → güvenli anahtar değişimi.

---

# ⚠️ Notlar

* MD5, SHA1 gibi eski algoritmalar **kuantum saldırılara karşı zayıftır**, bu yüzden ML-KEM gibi yeni nesil algoritmalar geliştirildi.
* Go’da resmi destek ileride `crypto/mlkem` altında gelebilir, ama şimdilik üçüncü parti kütüphaneler kullanılıyor.

---

👉 İstersen sana **ML-KEM ile TLS benzeri bir güvenli iletişim örneği** (mesaj şifreleme ve çözme) de gösterebilirim. Görmek ister misin?
*/