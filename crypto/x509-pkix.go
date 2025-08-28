Harika bir noktaya geldik 👍
`crypto/x509/pkix` paketi, Go’nun TLS/sertifika tarafında sıkça kullandığımız **X.509 Public Key Infrastructure (PKIX)** veri yapılarını tanımlar.

---

# 📦 Go `crypto/x509/pkix` Paketi

Bu paket, aslında **sertifika oluşturma ve doğrulamada kullanılan temel ASN.1 (Abstract Syntax Notation One) veri yapıları** için tipler sağlar.
Doğrudan sertifika imzalamaz, şifreleme yapmaz. Daha çok `crypto/x509` ile birlikte kullanılır.

---

## 🔑 İçindeki Önemli Tipler

### 1. **`pkix.Name`**

* Sertifikadaki **Subject (özne)** veya **Issuer (veren)** alanlarını tanımlar.
* `Organization`, `Country`, `Locality`, `CommonName` gibi alanlar içerir.

```go
package main

import (
	"crypto/x509/pkix"
	"fmt"
)

func main() {
	name := pkix.Name{
		CommonName:   "example.com",
		Organization: []string{"Example Corp"},
		Country:      []string{"TR"},
		Locality:     []string{"Istanbul"},
	}

	// String() metodunu kullanabiliriz
	fmt.Println("Sertifika adı:", name.String())
}
```

📌 Çıktı:

```
Sertifika adı: CN=example.com,O=Example Corp,C=TR,L=Istanbul
```

---

### 2. **`pkix.Extension`**

* X.509 sertifikalardaki **uzantıları (extensions)** temsil eder.
* Örn: Key Usage, Basic Constraints, Subject Alternative Name (SAN).

```go
package main

import (
	"crypto/x509/pkix"
	"encoding/asn1"
	"fmt"
)

func main() {
	// Subject Alternative Name (SAN) OID: 2.5.29.17
	oidSAN := asn1.ObjectIdentifier{2, 5, 29, 17}

	ext := pkix.Extension{
		Id:       oidSAN,
		Critical: false,
		Value:    []byte{0x30, 0x00}, // ASN.1 boş dizi
	}

	fmt.Println("Extension OID:", ext.Id)
	fmt.Println("Extension critical mi?:", ext.Critical)
}
```

---

### 3. **`pkix.AlgorithmIdentifier`**

* Sertifika ve imzalarda kullanılan algoritmaları belirtir.
* Örn: SHA256WithRSA, ECDSAWithSHA384.

```go
package main

import (
	"crypto/x509/pkix"
	"encoding/asn1"
	"fmt"
)

func main() {
	// SHA256 With RSA Encryption OID: 1.2.840.113549.1.1.11
	oidSHA256WithRSA := asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 11}

	alg := pkix.AlgorithmIdentifier{
		Algorithm:  oidSHA256WithRSA,
		Parameters: asn1.RawValue{Tag: 5}, // NULL parametresi
	}

	fmt.Println("Algoritma OID:", alg.Algorithm)
}
```

---

### 4. **`pkix.RDNSequence` ve `pkix.RelativeDistinguishedNameSET`**

* Distinguished Name (DN) alanlarının ASN.1 temsilleri.
* `pkix.Name` bunları daha okunabilir hale getiriyor, ama düşük seviyede `RDNSequence` kullanılabilir.

```go
package main

import (
	"crypto/x509/pkix"
	"encoding/asn1"
	"fmt"
)

func main() {
	rdn := pkix.RDNSequence{
		[]pkix.AttributeTypeAndValue{
			{
				Type:  asn1.ObjectIdentifier{2, 5, 4, 3}, // CN
				Value: "example.com",
			},
		},
	}

	der, _ := asn1.Marshal(rdn)
	fmt.Printf("RDN ASN.1 kodu: %x\n", der)
}
```

---

## 📌 Özet

`crypto/x509/pkix` paketi şunlar için kullanılır:

* ✅ Sertifika ad bilgilerini (`pkix.Name`) tanımlamak
* ✅ Sertifika uzantılarını (`pkix.Extension`) temsil etmek
* ✅ Kriptografik algoritma tanımlarını (`pkix.AlgorithmIdentifier`) saklamak
* ✅ Düşük seviye ASN.1 DN yapıları (`RDNSequence`) oluşturmak

Genellikle biz **`x509.Certificate`** oluştururken `pkix.Name` gibi tipleri doğrudan kullanıyoruz. Yani `x509/pkix`, `crypto/x509` paketinin “altyapı yapı taşı”.

---

📍 Şimdi sana dilersen **x509 + pkix ile Root CA ve Client sertifikası üretip doğrulayan tam bir Go projesi (Docker destekli)** gösterebilirim. İstiyor musun?
EVET

Projeyi oluşturdum ve canvas’a ekledim ✅

İstersen şimdi:

* Sana Docker ile nasıl build/run yapacağını adım adım yazayım, **veya**
* Hemen `make build && make run` komut setini verip başlatmana yardımcı olayım.

Hangisini tercih ediyorsun? EVET
Docker-File ile olan

Tamamdır 🚀 Şimdi sana Docker üzerinden bu `x509/pkix` projesini nasıl build edip çalıştıracağını adım adım göstereyim:

---

## 1️⃣ Docker Image Build Et

Projenin kök klasöründe (Dockerfile’ın bulunduğu yerde) şu komutu çalıştır:

```bash
docker build -t go-x509-pkix .
```

📌 Bu komut:

* `go-x509-pkix` adında bir image oluşturacak.
* İçinde sertifika oluşturma scripti (`scripts/generate_certs.sh`) çalıştırmaya hazır olacak.
* `server` ve `client` binary’lerini build edecek.

---

## 2️⃣ Container Çalıştır

Server’i ayağa kaldırmak için:

```bash
docker run -it --rm -p 8443:8443 go-x509-pkix server
```

Client’i çalıştırmak için (ayrı bir terminalde):

```bash
docker run -it --rm go-x509-pkix client
```

📌 Böylece client, container içindeki `certs/` klasöründeki root CA’ya güvenerek server’a TLS ile bağlanacak.
Mesaj göreceksin:

```
Server dinliyor: 8443
Client: Sunucudan cevap alındı: Merhaba, TLS istemcisi!
```

---

## 3️⃣ Docker Compose (Opsiyonel, tek komutla client+server)

Eğer `docker-compose.yml` dosyası da varsa sadece şu komutu çalıştırman yeterli:

```bash
docker-compose up --build
```

📌 Bu durumda hem `server` hem de `client` otomatik çalışacak.

---

👉 İster misin ben sana **`docker-compose logs` çıktısını nasıl görüp debug yapacağını** da göstereyim?

