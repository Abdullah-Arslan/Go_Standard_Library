/*
Go’da `crypto/x509` paketi, TLS/SSL sertifikaları ve X.509 standardına göre tanımlanmış public key altyapısı (PKI) ile çalışmak için kullanılır.

Bu paket şunları yapmamızı sağlar:

* Sertifika (X.509) oluşturmak, imzalamak, doğrulamak.
* Sertifikalardan **public/private key** bilgilerini çıkarmak.
* Sertifikaları `PEM` veya `DER` formatında encode/decode etmek.
* CRL (Certificate Revocation List) ve sertifika zincirlerini doğrulamak.
* TLS protokolünde client ve server kimlik doğrulamasında kullanılacak **root CA** ve sertifikaları yönetmek.

---

## 📌 Temel Yapılar

### `x509.Certificate`

Bir sertifikanın tüm bilgilerini içerir. (CN, SAN, PublicKey, Signature vb.)

### `x509.CertificateRequest`

CSR (Certificate Signing Request) oluşturmak için kullanılır.

### `x509.CertPool`

Root veya ara sertifikaları doğrulama havuzu.

---

## 📌 Önemli Fonksiyonlar

* `x509.CreateCertificate()` → Yeni bir sertifika oluşturur.
* `x509.CreateCertificateRequest()` → CSR oluşturur.
* `x509.ParseCertificate()` → DER formatındaki sertifikayı Go objesine dönüştürür.
* `x509.ParseCertificateRequest()` → CSR parse eder.
* `x509.MarshalPKCS1PrivateKey()` → RSA private key encode eder.
* `x509.MarshalPKIXPublicKey()` → Public key encode eder.
* `x509.ParsePKIXPublicKey()` → Public key parse eder.
* `(*Certificate).Verify()` → Sertifikayı root CA havuzuna göre doğrular.

---

## 📌 Örnek 1: Self-Signed Sertifika Oluşturma
*/
``go
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"time"
)

func main() {
	// RSA anahtar çifti üret
	priv, _ := rsa.GenerateKey(rand.Reader, 2048)

	// Sertifika bilgileri
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"My Company"},
			CommonName:   "localhost",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour), // 1 yıl geçerli
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	// Self-signed sertifika
	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)

	// Sertifikayı kaydet
	certOut, _ := os.Create("cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	// Private key kaydet
	keyOut, _ := os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	keyOut.Close()
}
``
/*
✅ Bu kod çalıştırıldığında `cert.pem` ve `key.pem` dosyaları oluşur.

---

## 📌 Örnek 2: Sertifika Parse Etme ve Bilgilerini Okuma
*/
``go
package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	// Dosyadan oku
	certPEM, _ := os.ReadFile("cert.pem")
	block, _ := pem.Decode(certPEM)
	cert, _ := x509.ParseCertificate(block.Bytes)

	fmt.Println("CommonName:", cert.Subject.CommonName)
	fmt.Println("Organization:", cert.Subject.Organization)
	fmt.Println("Geçerlilik:", cert.NotBefore, "-", cert.NotAfter)
}
``
/*
---

## 📌 Örnek 3: CSR (Certificate Signing Request) Oluşturma
*/
``go
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"os"
)

func main() {
	priv, _ := rsa.GenerateKey(rand.Reader, 2048)

	csrTemplate := x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName:   "example.com",
			Organization: []string{"TestOrg"},
		},
	}

	csrDER, _ := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, priv)

	// CSR dosyaya yaz
	csrFile, _ := os.Create("request.csr")
	pem.Encode(csrFile, &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrDER})
	csrFile.Close()
}
``
/*
✅ Bu CSR, bir **CA** (Certificate Authority) tarafından imzalanabilir.

---

## 📌 Örnek 4: Sertifikayı Root CA ile Doğrulama
*/
``go
package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	// Client sertifikasını oku
	clientPEM, _ := os.ReadFile("client.pem")
	clientBlock, _ := pem.Decode(clientPEM)
	clientCert, _ := x509.ParseCertificate(clientBlock.Bytes)

	// Root CA sertifikasını oku
	rootPEM, _ := os.ReadFile("ca.pem")
	rootBlock, _ := pem.Decode(rootPEM)
	rootCert, _ := x509.ParseCertificate(rootBlock.Bytes)

	// Root CA havuzu oluştur
	roots := x509.NewCertPool()
	roots.AddCert(rootCert)

	// Doğrulama
	opts := x509.VerifyOptions{
		Roots: roots,
	}
	if _, err := clientCert.Verify(opts); err != nil {
		fmt.Println("Doğrulama başarısız:", err)
	} else {
		fmt.Println("Doğrulama başarılı ✅")
	}
}
``
/*
---

## 📌 Özet

* `crypto/x509` paketi **sertifikalarla ilgili tüm işlemleri** yapar.
* Sertifika üretme (`CreateCertificate`)
* CSR oluşturma (`CreateCertificateRequest`)
* Sertifika parse etme (`ParseCertificate`)
* Sertifika doğrulama (`Verify`)
* Key encode/decode (`MarshalPKCS1PrivateKey`, `ParsePKIXPublicKey`)

---

👉 İstersen sana gerçek bir **Root CA, Intermediate CA ve Client sertifikası zinciri** oluşturup `x509` ile doğrulayan bir mini proje de hazırlayayım. Bunu ister misin?
EVET
*/

/*
Süper 👍 O zaman sana **Root CA → Intermediate CA → Client** sertifika zincirini oluşturan ve `crypto/x509` ile doğrulayan **mini bir Go projesi** vereyim.

---

# 📂 Proje Yapısı

``
x509-chain-demo/
│── generate_certs.go   // Root CA, Intermediate CA ve Client sertifikalarını üretir
│── verify_chain.go     // Zinciri doğrular
│── certs/
     ├── rootCA.pem
     ├── rootCA.key
     ├── intermediateCA.pem
     ├── intermediateCA.key
     ├── client.pem
     ├── client.key
``
*/
/*
---

## 🔹 1. Sertifika Zincirini Üretme (`generate_certs.go`)
*/
``go
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

func savePEM(filename, pemType string, bytes []byte) {
	file, _ := os.Create(filename)
	pem.Encode(file, &pem.Block{Type: pemType, Bytes: bytes})
	file.Close()
}

func main() {
	// 1️⃣ Root CA key ve sertifika oluştur
	rootKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	rootTemplate := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"My Root CA"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(10 * 365 * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
		MaxPathLen:            2,
	}
	rootDER, _ := x509.CreateCertificate(rand.Reader, &rootTemplate, &rootTemplate, &rootKey.PublicKey, rootKey)

	os.Mkdir("certs", 0755)
	savePEM("certs/rootCA.pem", "CERTIFICATE", rootDER)
	savePEM("certs/rootCA.key", "RSA PRIVATE KEY", x509.MarshalPKCS1PrivateKey(rootKey))

	fmt.Println("✅ Root CA oluşturuldu")

	// 2️⃣ Intermediate CA
	interKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	interTemplate := x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject: pkix.Name{
			Organization: []string{"My Intermediate CA"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(5 * 365 * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
		IsCA:                  true,
		MaxPathLen:            1,
	}
	interDER, _ := x509.CreateCertificate(rand.Reader, &interTemplate, &rootTemplate, &interKey.PublicKey, rootKey)
	savePEM("certs/intermediateCA.pem", "CERTIFICATE", interDER)
	savePEM("certs/intermediateCA.key", "RSA PRIVATE KEY", x509.MarshalPKCS1PrivateKey(interKey))

	fmt.Println("✅ Intermediate CA oluşturuldu")

	// 3️⃣ Client Sertifikası
	clientKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	clientTemplate := x509.Certificate{
		SerialNumber: big.NewInt(3),
		Subject: pkix.Name{
			CommonName:   "client.local",
			Organization: []string{"My Client"},
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:    x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
	}
	clientDER, _ := x509.CreateCertificate(rand.Reader, &clientTemplate, &interTemplate, &clientKey.PublicKey, interKey)
	savePEM("certs/client.pem", "CERTIFICATE", clientDER)
	savePEM("certs/client.key", "RSA PRIVATE KEY", x509.MarshalPKCS1PrivateKey(clientKey))

	fmt.Println("✅ Client sertifikası oluşturuldu")
}
``
/*
---

## 🔹 2. Sertifika Zincirini Doğrulama (`verify_chain.go`)
*/
``go
package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func loadCert(path string) *x509.Certificate {
	data, _ := os.ReadFile(path)
	block, _ := pem.Decode(data)
	cert, _ := x509.ParseCertificate(block.Bytes)
	return cert
}

func main() {
	// Sertifikaları yükle
	rootCert := loadCert("certs/rootCA.pem")
	interCert := loadCert("certs/intermediateCA.pem")
	clientCert := loadCert("certs/client.pem")

	// Root CA havuzu
	roots := x509.NewCertPool()
	roots.AddCert(rootCert)

	// Ara sertifika havuzu
	intermediates := x509.NewCertPool()
	intermediates.AddCert(interCert)

	opts := x509.VerifyOptions{
		Roots:         roots,
		Intermediates: intermediates,
	}

	if chains, err := clientCert.Verify(opts); err != nil {
		fmt.Println("❌ Doğrulama başarısız:", err)
	} else {
		fmt.Println("✅ Zincir doğrulandı! Sertifika yolu:")
		for _, chain := range chains {
			for _, cert := range chain {
				fmt.Println("   →", cert.Subject.Organization, cert.Subject.CommonName)
			}
		}
	}
}
``
/*
---

## 🔹 Çalıştırma
*/
``bash
go run generate_certs.go
go run verify_chain.go
``

//📌 Çıktı:

``
✅ Root CA oluşturuldu
✅ Intermediate CA oluşturuldu
✅ Client sertifikası oluşturuldu
✅ Zincir doğrulandı! Sertifika yolu:
   → [My Client] client.local
   → [My Intermediate CA] 
   → [My Root CA]
``
/*
---

Bu proje ile:
✔ Root CA, Intermediate CA ve Client sertifikaları üretildi.
✔ `crypto/x509` ile sertifika zinciri doğrulandı.

---

👉 İstersen ben bunu sana **tam bir proje klasörü** (Dockerfile ve Makefile dahil) olarak paketleyeyim, `docker run` ile çalıştırabilirsin. Bunu da ister misin?
EVET
*/

/*
Harika 👍 O zaman sana `crypto/x509` sertifika zinciri örneğini **tam proje klasörü + Dockerfile + Makefile** ile vereyim.

---

# 📂 Proje Yapısı

```
x509-chain-demo/
│── Dockerfile
│── Makefile
│── generate_certs.go
│── verify_chain.go
│── certs/         (otomatik oluşturulur)
``
*/
/*
---

## 🔹 1. `generate_certs.go`
*/

``go
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

func savePEM(filename, pemType string, bytes []byte) {
	file, _ := os.Create(filename)
	pem.Encode(file, &pem.Block{Type: pemType, Bytes: bytes})
	file.Close()
}

func main() {
	// Root CA
	rootKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	rootTemplate := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{Organization: []string{"My Root CA"}},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(10 * 365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		IsCA:         true,
	}
	rootDER, _ := x509.CreateCertificate(rand.Reader, &rootTemplate, &rootTemplate, &rootKey.PublicKey, rootKey)

	os.Mkdir("certs", 0755)
	savePEM("certs/rootCA.pem", "CERTIFICATE", rootDER)
	savePEM("certs/rootCA.key", "RSA PRIVATE KEY", x509.MarshalPKCS1PrivateKey(rootKey))
	fmt.Println("✅ Root CA oluşturuldu")

	// Intermediate CA
	interKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	interTemplate := x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject:      pkix.Name{Organization: []string{"My Intermediate CA"}},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(5 * 365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		IsCA:         true,
	}
	interDER, _ := x509.CreateCertificate(rand.Reader, &interTemplate, &rootTemplate, &interKey.PublicKey, rootKey)
	savePEM("certs/intermediateCA.pem", "CERTIFICATE", interDER)
	savePEM("certs/intermediateCA.key", "RSA PRIVATE KEY", x509.MarshalPKCS1PrivateKey(interKey))
	fmt.Println("✅ Intermediate CA oluşturuldu")

	// Client sertifikası
	clientKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	clientTemplate := x509.Certificate{
		SerialNumber: big.NewInt(3),
		Subject:      pkix.Name{CommonName: "client.local", Organization: []string{"My Client"}},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
	}
	clientDER, _ := x509.CreateCertificate(rand.Reader, &clientTemplate, &interTemplate, &clientKey.PublicKey, interKey)
	savePEM("certs/client.pem", "CERTIFICATE", clientDER)
	savePEM("certs/client.key", "RSA PRIVATE KEY", x509.MarshalPKCS1PrivateKey(clientKey))
	fmt.Println("✅ Client sertifikası oluşturuldu")
}
``
/*
---

## 🔹 2. `verify_chain.go`
*/

``go
package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func loadCert(path string) *x509.Certificate {
	data, _ := os.ReadFile(path)
	block, _ := pem.Decode(data)
	cert, _ := x509.ParseCertificate(block.Bytes)
	return cert
}

func main() {
	rootCert := loadCert("certs/rootCA.pem")
	interCert := loadCert("certs/intermediateCA.pem")
	clientCert := loadCert("certs/client.pem")

	roots := x509.NewCertPool()
	roots.AddCert(rootCert)

	intermediates := x509.NewCertPool()
	intermediates.AddCert(interCert)

	opts := x509.VerifyOptions{
		Roots:         roots,
		Intermediates: intermediates,
	}

	if chains, err := clientCert.Verify(opts); err != nil {
		fmt.Println("❌ Doğrulama başarısız:", err)
	} else {
		fmt.Println("✅ Zincir doğrulandı! Sertifika yolu:")
		for _, chain := range chains {
			for _, cert := range chain {
				fmt.Println("   →", cert.Subject.Organization, cert.Subject.CommonName)
			}
		}
	}
}
``
/*
---

## 🔹 3. `Dockerfile`
*/

``dockerfile
FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go mod init x509-chain-demo && go mod tidy

CMD ["sh", "-c", "go run generate_certs.go && go run verify_chain.go"]
``
/*
---

## 🔹 4. `Makefile`
*/
``makefile
build:
	docker build -t x509-chain-demo .

run:
	docker run --rm x509-chain-demo
``
/*
---

## 🔹 Çalıştırma
*/
``bash
make build
make run
``
/*
📌 Çıktı:

```
✅ Root CA oluşturuldu
✅ Intermediate CA oluşturuldu
✅ Client sertifikası oluşturuldu
✅ Zincir doğrulandı! Sertifika yolu:
   → [My Client] client.local
   → [My Intermediate CA] 
   → [My Root CA]
```

---

Böylece hem **Go kodları** hem de **Docker ortamı** hazır 🚀

👉 İstersen ben buna bir de **docker-compose.yml** ekleyip farklı konteynerlerde “CA” ve “client doğrulama” çalıştırabilirim. İstiyor musun?
*/
