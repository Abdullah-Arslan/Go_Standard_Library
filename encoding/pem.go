/*
Go’nun standart kütüphanesinde bulunan **`encoding/pem`** paketi, **PEM (Privacy-Enhanced Mail) formatında kodlanmış verilerle** çalışmaya yarar.

PEM, aslında **Base64 ile kodlanmış ikili (binary) veriyi** içerir ve genelde şu alanlarda kullanılır:

* **Sertifikalar** (X.509 → `.crt`, `.cer`)
* **Anahtarlar** (RSA, EC → `.pem`, `.key`)
* **TLS / SSL** dosyaları
* **PKCS#7, PKCS#8** içerikleri

PEM dosyalarının genel yapısı şöyledir:

```
-----BEGIN TYPE-----
(base64-encoded data)
-----END TYPE-----
```

Örneğin bir RSA private key dosyası:

```
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAuYj...
-----END RSA PRIVATE KEY-----
```

---

## 📌 `pem` Paketinin Fonksiyonları

`encoding/pem` paketi üç temel şey sunar:

1. **`pem.Block` yapısı** → Tek bir PEM bloğunu temsil eder.
*/
   ``go
   type Block struct {
       Type    string            // Örn: "CERTIFICATE", "RSA PRIVATE KEY"
       Headers map[string]string // Opsiyonel metadata
       Bytes   []byte            // Gerçek binary içerik (base64 decode edilmiş)
   }
   ``
/*
2. **`pem.Decode(data []byte) (*pem.Block, rest []byte)`**

   * Bir PEM verisini çözer.
   * İlk bloğu döndürür, geri kalanı `rest` içinde kalır.

3. **`pem.Encode(w io.Writer, b *pem.Block) error`**

   * Bir `pem.Block`’u yazar (PEM formatına çevirir).

4. **`pem.EncodeToMemory(b *pem.Block) []byte`**

   * Bir `pem.Block`’u bellekte PEM formatında döndürür.

---

## 📌 Örnekler

### 1. PEM Dosyasını Okuma (Decode)
*/

``go
package main

import (
	"encoding/pem"
	"fmt"
)

func main() {
	pemData := []byte(`
-----BEGIN CERTIFICATE-----
MIIBkTCCATmgAwIBAgIUEXWJt...
-----END CERTIFICATE-----
`)

	block, rest := pem.Decode(pemData)
	if block == nil {
		fmt.Println("PEM bloğu çözülemedi")
		return
	}

	fmt.Println("Tip:", block.Type)
	fmt.Println("Boyut (byte):", len(block.Bytes))
	fmt.Println("Kalan veri:", len(rest))
}
``
/*
🔹 Bu kod `CERTIFICATE` tipini bulur ve **Base64 çözülmüş ham byte verisini** `block.Bytes` içinde verir.

---

### 2. Bir PEM Bloğu Yazma (Encode)
*/
``go
package main

import (
	"encoding/pem"
	"fmt"
)

func main() {
	data := []byte("Merhaba PEM!")

	block := &pem.Block{
		Type:  "MESSAGE",
		Bytes: data,
	}

	pemBytes := pem.EncodeToMemory(block)
	fmt.Println(string(pemBytes))
}
``
/*
**Çıktı:**

```
-----BEGIN MESSAGE-----
TWVySGFiYSBQRU0h
-----END MESSAGE-----
```

(`Merhaba PEM!` → Base64 → PEM formatı)

---

### 3. Birden Fazla PEM Bloğu Çözmek

PEM dosyalarında birden fazla blok olabilir (örn. sertifika zinciri).
*/
``go
package main

import (
	"encoding/pem"
	"fmt"
)

func main() {
	data := []byte(`
-----BEGIN CERTIFICATE-----
AAA...
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
BBB...
-----END CERTIFICATE-----
`)

	for {
		block, rest := pem.Decode(data)
		if block == nil {
			break
		}
		fmt.Println("Bulunan tip:", block.Type)
		data = rest
	}
}
``
/*
---

### 4. PEM Başlıkları (Headers) Kullanmak

PEM bloklarında ekstra metadata da olabilir:
*/
``go
package main

import (
	"encoding/pem"
	"fmt"
)

func main() {
	block := &pem.Block{
		Type: "MY DATA",
		Headers: map[string]string{
			"App":    "Go Example",
			"Author": "AA",
		},
		Bytes: []byte("secret info"),
	}

	pemBytes := pem.EncodeToMemory(block)
	fmt.Println(string(pemBytes))
}
``
/*
**Çıktı:**

```
-----BEGIN MY DATA-----
App: Go Example
Author: AA

c2VjcmV0IGluZm8=
-----END MY DATA-----
```

---

### 5. PEM → Sertifika Parse Etme (`x509` ile birlikte)

`pem` genellikle **`crypto/x509`** paketi ile birlikte kullanılır.
Örneğin sertifika okumak:
*/

``go
package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func main() {
	certPEM := []byte(`
-----BEGIN CERTIFICATE-----
MIIBkTCCATmgAwIBAgIUEXWJt...
-----END CERTIFICATE-----
`)

	block, _ := pem.Decode(certPEM)
	if block == nil || block.Type != "CERTIFICATE" {
		fmt.Println("Sertifika bulunamadı")
		return
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Println("Parse hatası:", err)
		return
	}

	fmt.Println("Subject:", cert.Subject)
	fmt.Println("Issuer:", cert.Issuer)
}
``
/*
---

## 📌 Özet

* **`encoding/pem`** paketi → PEM formatındaki verileri **okumak** ve **yazmak** için kullanılır.
* Temel fonksiyonlar:

  * `Decode` → PEM bloğunu çözmek
  * `Encode` / `EncodeToMemory` → PEM bloğu oluşturmak
* Genelde **TLS sertifikaları, private/public anahtarlar** ile çalışırken kullanılır.
* Sıklıkla **`crypto/x509`** ile birlikte kullanılır.

---

👉 İstersen sana `pem` + `x509` + `tls` kullanarak **Go ile bir HTTPS sunucusu kurma (sertifikayı PEM’den okuma)** örneği gösterebilirim. İstiyor musun?
EVET
*/

/*
Süper 👍 o zaman sana **`pem` + `x509` + `tls`** kullanarak Go’da bir **HTTPS sunucusu** nasıl kurulur göstereyim.

Bu örnekte:

1. **Sahte (self-signed) sertifika** ve anahtar (`cert.pem`, `key.pem`) oluşturacağız.
2. `pem` ve `x509` ile bu dosyaları okuyacağız.
3. `tls.Config` ile HTTPS server ayağa kaldıracağız.

---

## 1️⃣ Sertifika ve Anahtar Oluşturma

Linux/Mac/WSL’de terminalde şu komut ile kendimize sahte bir sertifika ve anahtar üretebiliriz:

``bash
openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pem -out cert.pem -days 365
``

Bu işlemden sonra iki dosyamız olacak:

* `cert.pem` → Sertifika
* `key.pem` → Private Key

---

## 2️⃣ Go Kodu: HTTPS Sunucusu
*/

``go
package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Sertifikayı oku
	certPEM, err := ioutil.ReadFile("cert.pem")
	if err != nil {
		panic(err)
	}
	keyPEM, err := ioutil.ReadFile("key.pem")
	if err != nil {
		panic(err)
	}

	// PEM decode et → CERTIFICATE
	certBlock, _ := pem.Decode(certPEM)
	if certBlock == nil || certBlock.Type != "CERTIFICATE" {
		panic("Sertifika çözülemedi")
	}
	cert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		panic(err)
	}
	fmt.Println("Sertifika subject:", cert.Subject)

	// PEM decode et → PRIVATE KEY
	keyBlock, _ := pem.Decode(keyPEM)
	if keyBlock == nil {
		panic("Private key çözülemedi")
	}

	// Sertifika + anahtarı TLS için yükle
	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}

	// TLS Config
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
	}

	// Basit handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Merhaba! HTTPS sunucusu çalışıyor 🚀")
	})

	// HTTPS server başlat
	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}

	fmt.Println("HTTPS sunucu https://localhost:8443 adresinde başladı...")
	err = server.ListenAndServeTLS("", "") // "" çünkü tlsConfig içinde verdik
	if err != nil {
		panic(err)
	}
}
``
/*
---

## 3️⃣ Çalıştırma

Programı çalıştır:

```bash
go run main.go
```

Tarayıcıdan şuraya git:

👉 [https://localhost:8443](https://localhost:8443)

⚠️ Tarayıcı “Bu sertifika güvenilir değil” diye uyarı verecektir, çünkü **self-signed** (kendi imzaladığımız) bir sertifika kullandık. Normalde production’da **CA (Certificate Authority)** tarafından imzalanmış sertifika gerekir.

---

## 📌 Özet

* `pem` → sertifika & anahtar dosyalarını çözmek (Base64 decode + metadata)
* `x509` → çözülmüş veriyi sertifika yapısına dönüştürmek
* `tls` → HTTPS server/client oluşturmak için kullanılır

---

👉 İstersen ben şimdi aynı sertifikayı **client tarafında** `pem` ile okuyup, sadece bu sertifika ile doğrulama yapan **TLS client** örneği de gösterebilirim. Görmek ister misin?
*/