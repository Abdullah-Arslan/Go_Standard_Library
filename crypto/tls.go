/*
Go’da `crypto/tls` paketini kapsamlı ama pratik örneklerle açıklayayım. Aşağıda ne olduğundan başlayıp, önemli yapı/alanlara, güvenlik notlarına ve birden fazla çalışır örneğe (basit HTTPS sunucu, özel `tls.Config`, mutual TLS, istemci örneği) kadar her şeyi bulacaksın.

# `crypto/tls` nedir?

`crypto/tls` Go’nun TLS/SSL (Transport Layer Security) protokolünü uygulayan paketidir. TLS, ağ üzerinden güvenli (şifreli) iletişim sağlar — HTTPS, gRPC, SMTP TLS vb. için kullanılır. Go 1.13+ ile TLS 1.3 desteği vardır ve paket hem sunucu hem istemci tarafı API’leri sunar.

---

# Önemli türler / alanlar (kısa özet)

* `tls.Config` — TLS davranışını belirler (sertifika doğrulama, TLS sürümü, cipher suite’ler, SNI, client auth vs).
* `tls.Certificate` — sertifika + private key çifti (PEM’den yüklenir).
* `tls.Conn` — TLS bağlantısı (net.Conn sarılı hali).
* `tls.Listen` / `tls.Dial` — TLS üzerinden dinlemek / bağlanmak.
* `tls.X509KeyPair` — PEM formatındaki sertifika ve anahtar dosyalarından `tls.Certificate` üretir.

---

# Güvenlik/konfigürasyon açısından önemli `tls.Config` alanları

* `MinVersion`, `MaxVersion` (örn. `tls.VersionTLS12`, `tls.VersionTLS13`)
* `CipherSuites` (TLS1.2 için; TLS1.3 cipherlar sabittir ve ayrı kontrol edilmez)
* `PreferServerCipherSuites` (server tarafı için)
* `Certificates []tls.Certificate` (sunucunun sertifikaları)
* `GetCertificate` (SNI’ye göre dinamik sertifika seçmek için)
* `ClientAuth` (ör. `tls.RequireAndVerifyClientCert` — mutual TLS)
* `ClientCAs *x509.CertPool` (istemci sertifikalarını doğrulamak için CA havuzu)
* `RootCAs *x509.CertPool` (istemci tarafında sunucu sertifikalarını doğrulamak için)
* `InsecureSkipVerify bool` (uyarı: genelde `false`; true yapmayın — yerine `VerifyPeerCertificate` kullanın)
* `VerifyPeerCertificate func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error` — özel doğrulama

---

# Hızlı: self-signed sertifika oluşturma (test amaçlı)

(gerçek üretimde CA imzalı sertifika kullan)

``bash
# RSA 2048 + self-signed
openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pem -out cert.pem -days 365 \
  -subj "/C=TR/ST=Istanbul/L=Istanbul/O=Example/OU=Dev/CN=localhost"
``
/*
---

# 1) Basit HTTPS sunucu — `http.ListenAndServeTLS`

En hızlı yol:
*/
``go
package main

import (
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Merhaba TLS!"))
}

func main() {
	http.HandleFunc("/", hello)
	// cert.pem ve key.pem dosyalarınız olmalı
	log.Fatal(http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil))
}
``
/*
`http.ListenAndServeTLS` arkasında `tls.Config` oluşturur; hızlı test için uygundur.

---

# 2) Özel `tls.Config` kullanarak sunucu (SNI örneği, TLS versiyon sınırla)
*/
``go
package main

import (
	"crypto/tls"
	"log"
	"net"
)

func main() {
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatal(err)
	}

	cfg := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12, // TLS 1.2+'ı zorunlu kıl
		// GetCertificate: func(clientHello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		//     // SNI bazlı dinamik sertifika seçimi
		// },
		PreferServerCipherSuites: true,
	}

	ln, err := tls.Listen("tcp", ":8443", cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("accept:", err)
			continue
		}
		go func(c net.Conn) {
			defer c.Close()
			// TLS bağlantısıyla şimdi standart net.Conn kullanarak read/write yapılır
			c.Write([]byte("Merhaba TLS (raw)!\n"))
		}(conn)
	}
}
``
/*
---

# 3) TLS istemci: `tls.Dial` ve özelleştirilmiş CA havuzu

Sunucunuz kendi CA ile imzalandıysa, istemci RootCAs ayarlamalı:
*/
``go
package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
)

func main() {
	// Sunucuyu imzalayan CA'yı yükle
	caCert, err := ioutil.ReadFile("ca_cert.pem")
	if err != nil { log.Fatal(err) }
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caCert)

	cfg := &tls.Config{
		RootCAs:    pool,
		ServerName: "server.example.com", // SNI ve sertifika CN/ SAN ile eşleşmeli
		MinVersion: tls.VersionTLS12,
	}

	conn, err := tls.Dial("tcp", "server.example.com:443", cfg)
	if err != nil { log.Fatal(err) }
	defer conn.Close()
	// conn üzerinde Read/Write yapabilirsiniz
}
``
/*
---

# 4) HTTP istemcisinde `tls.Config` kullanmak (özellikle sertifika doğrulama)
*/
``go
package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	caCert, _ := ioutil.ReadFile("ca_cert.pem")
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caCert)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: pool,
			// InsecureSkipVerify: true, // asla üretimde kullanmayın
		},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get("https://server.example.com/resource")
	if err != nil { log.Fatal(err) }
	defer resp.Body.Close()
	// ...
}
``
/*
---

# 5) Mutual TLS (istemci sertifikası doğrulama)

Sunucu, istemcinin sertifikasını doğrulamak için `ClientAuth` ve `ClientCAs` kullanır:
*/
``go
package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Sunucu sertifikası
	cert, _ := tls.LoadX509KeyPair("server_cert.pem", "server_key.pem")

	// İstemciyi imzalayan CA
	clientCA, _ := ioutil.ReadFile("client_ca.pem")
	clientPool := x509.NewCertPool()
	clientPool.AppendCertsFromPEM(clientCA)

	cfg := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    clientPool,
		ClientAuth:   tls.RequireAndVerifyClientCert, // mutual TLS
		MinVersion:   tls.VersionTLS12,
	}

	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: cfg,
		Handler:   http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// İstemci sertifikası r.TLS.PeerCertificates içinde
			if len(r.TLS.PeerCertificates) > 0 {
				subject := r.TLS.PeerCertificates[0].Subject
				w.Write([]byte("Hoşgeldin: " + subject.CommonName))
			} else {
				w.Write([]byte("İstemci sertifikası yok"))
			}
		}),
	}
	log.Fatal(server.ListenAndServeTLS("", "")) // sertifika zaten TLSConfig içinde
}
``
/*
---

# 6) `InsecureSkipVerify` tehlikesi ve `VerifyPeerCertificate`

* `InsecureSkipVerify: true` → istemci sunucu sertifikasını doğramaz. Üretimde **sakın kullanma**.
* Eğer özel doğrulama (ör. certificate pinning) yapmak istiyorsan `InsecureSkipVerify: true` + `VerifyPeerCertificate` kombinasyonu kullanılır — ama dikkatli ol:
*/
``go
cfg := &tls.Config{
	InsecureSkipVerify: true, // sertifika zincirinin otomatik doğrulamasını pas geç
	VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
		// Burada rawCerts[0] ile sertifika parse edip kendi doğrulamanı yap
		// örn: sertifika fingerprint kontrolü (pinning)
		return nil
	},
}
``
/*
Yani `InsecureSkipVerify` tek başına tehlikelidir; yalnızca `VerifyPeerCertificate` ile güvenli bir doğrulama yapıyorsanız anlamlı olur.

---

# 7) Oturum (session ticket), yeniden kullanılabilirlik ve performans

* Go TLS oturum ticket’larını ve yeniden kullanım mekanizmalarını destekler, ama detaylar otomatik yönetilir.
* TLS 1.3 performans ve güvenlik açısından iyidir (0-RTT bazı riskler taşır; dikkatli kullan).

---

# 8) Örnek: Sunucu tarafında SNI’ye göre dinamik sertifika seçme
*/
``go
cfg := &tls.Config{
	GetCertificate: func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		switch hello.ServerName {
		case "foo.example.com":
			return tls.LoadX509KeyPair("foo_cert.pem", "foo_key.pem")
		case "bar.example.com":
			return tls.LoadX509KeyPair("bar_cert.pem", "bar_key.pem")
		default:
			return tls.LoadX509KeyPair("default_cert.pem", "default_key.pem")
		}
	},
}
``

---
/*
# 9 Faydalı ipuçları / en iyi uygulamalar

* Üretimde **TLS 1.2 veya 1.3** kullan; tercihen `MinVersion: tls.VersionTLS12` ve `MaxVersion: tls.VersionTLS13`.
* `InsecureSkipVerify = false` tut; özel doğrulama gerekiyorsa `VerifyPeerCertificate` kullan.
* Sertifikaları otomatik yenilemek için Let's Encrypt + `certmagic` / `autocert` gibi kütüphaneler kullanılabilir.
* Kendi CA / self-signed test sertifikalarında istemci tarafında `RootCAs` eklemeyi unutma.
* Parola/anahtar dosyalarını disk üzerinde güvenli tut (dosya izinleri).
* Cipher suite’leri elle seçerken dikkatli ol (genelde varsayılanlar modern, güvenlidir).
* `http.Server` ile TLS kullanırken `TLSConfig` ayarlayarak ileri seviye kontrol sağla.

---

# 10) Hızlı referans — sık kullanılan fonksiyonlar

* `tls.LoadX509KeyPair(certFile, keyFile)` — PEM’den `tls.Certificate` oluşturur.
* `tls.Listen("tcp", addr, cfg)` — TLS ile dinler.
* `tls.Dial("tcp", addr, cfg)` — TLS ile bağlanır.
* `http.ListenAndServeTLS(addr, certFile, keyFile, handler)` — hızlı HTTPS sunucu.

---

# Örnek: Minimal hem sunucu hem istemci (çalışır örnekler)

Sunucu (http.ListenAndServeTLS kullanan):

``go
// see previous ListenAndServeTLS example
``

İstemci (`http.Client` TLS config):

``go
// see previous HTTP client example with custom Transport
``
/*
(Üstteki bölümlerde tam kod blokları mevcut — doğrudan o örnekleri kullanabilirsin.)

---

## Özet

* `crypto/tls` Go’da TLS/SSL için ana pakettir; hem istemci hem sunucu tarafı API’si sunar.
* `tls.Config` ile davranışı detaylıca kontrol edersin: TLS versiyon, cipher suite, sertifika seçimi, client auth, özel doğrulama vs.
* Üretimde CA-imzalı sertifikalar ve TLS 1.2/1.3 tercih et; `InsecureSkipVerify` kullanma.
* Mutual TLS, SNI, certificate pinning gibi gelişmiş senaryolar `tls.Config` üzerinden implement edilir.

---

İstersen şimdi:

* Bu örneklerden herhangi birini tam çalışır hâle getirecek (sunucu+istemci) **tam projeyi** oluşturayım (ör: self-signed cert üretme + sunucu + istemci).
* Ya da `http.Server` + Let's Encrypt otomatik sertifika yenileme örneği hazırlayayım.
  Hangisini istersin?
EVET
*/

/*
# Go TLS Full Project

Bu proje: **self-signed CA ile imzalanmış sunucu sertifikası üretme**, **opsiyonel mutual TLS** ve hem sunucu hem istemci örnekleri içerir. Aşağıdaki dosyalar canvas içinde yer alıyor — kopyala/çalıştır veya terminaldeki talimatları izle.
*/
---

## Dosya: `README.md`

````markdown
# go-tls-full-project

Bu proje, Go ile TLS kullanarak güvenli bir sunucu ve istemci örneği içerir. Ayrıca `scripts/generate_certs.sh` ile bir CA oluşturur, bu CA ile sunucu ve (opsiyonel) istemci sertifikaları imzalar.

### İçindekiler
- `scripts/generate_certs.sh` - CA, server ve client sertifikalarını oluşturur.
- `server/main.go` - TLS sunucusu. İki mod:
  - Normal TLS (server cert sadece)
  - Mutual TLS (istemciden sertifika doğrulama) — `MUTUAL=true` ortam değişkeni ile aktifleşir
- `client/main.go` - TLS istemcisi. İki mod:
  - Server doğrulama only (CA trust)
  - Mutual TLS (istemci sertifikası sunar) — `--mutual` flag ile

### Gereksinimler
- Go 1.18+
- OpenSSL (sertifika oluşturmak için)

### Hızlı kullanım
1. Sertifikaları üret:

``bash
chmod +x scripts/generate_certs.sh
./scripts/generate_certs.sh
``

2. Sunucuyu çalıştır (normal TLS):

``bash
cd server
go run main.go
# veya
MUTUAL=false go run main.go
```

//3. İstemci (server doğrulama):

``bash
cd client
go run main.go --ca ../scripts/ca/ca.pem
``
/*
4. Mutual TLS testi (istemci sertifikası ile):

Önce sunucuyu mutual modda başlat:

``bash
MUTUAL=true go run main.go
``

Sonra istemciyi mutual flag ile çalıştır:
*/
```bash
go run main.go --ca ../scripts/ca/ca.pem --cert ../scripts/ca/client_cert.pem --key ../scripts/ca/client_key.pem --mutual
``
/*
``text
Not: Sunucu 8443 portunda dinler. Eğer port kullanımında sorun varsa değiştirebilirsin.
``
*/


---

## Dosya: `scripts/generate_certs.sh`

``bash
#!/usr/bin/env bash
set -euo pipefail

OUT=./scripts/ca
mkdir -p "$OUT"
cd "$OUT"

echo "Creating CA..."
# CA private key
openssl genrsa -out ca.key 4096
# CA self-signed cert
openssl req -x509 -new -nodes -key ca.key -sha256 -days 3650 \
  -subj "/C=TR/ST=Istanbul/L=Istanbul/O=ExampleCA/OU=Dev/CN=Example Root CA" \
  -out ca.pem

echo "Creating server key & CSR..."
openssl genrsa -out server_key.pem 2048
openssl req -new -key server_key.pem -subj "/C=TR/ST=Istanbul/L=Istanbul/O=ExampleServer/OU=Dev/CN=localhost" -out server.csr

# Create server cert signed by CA (include SAN localhost)
cat > server_ext.cnf <<EOF
basicConstraints=CA:FALSE
subjectAltName=DNS:localhost,IP:127.0.0.1
keyUsage = digitalSignature, keyEncipherment
extendedKeyUsage = serverAuth
EOF

openssl x509 -req -in server.csr -CA ca.pem -CAkey ca.key -CAcreateserial -out server_cert.pem -days 365 -sha256 -extfile server_ext.cnf

# Optional: create client cert (for mutual TLS tests)
echo "Creating client key & CSR..."
openssl genrsa -out client_key.pem 2048
openssl req -new -key client_key.pem -subj "/C=TR/ST=Istanbul/L=Istanbul/O=ExampleClient/OU=Dev/CN=client.local" -out client.csr

cat > client_ext.cnf <<EOF
basicConstraints=CA:FALSE
keyUsage = digitalSignature, keyEncipherment
extendedKeyUsage = clientAuth
EOF

openssl x509 -req -in client.csr -CA ca.pem -CAkey ca.key -CAcreateserial -out client_cert.pem -days 365 -sha256 -extfile client_ext.cnf

# cleanup
rm -f server.csr client.csr server_ext.cnf client_ext.cnf ca.srl

echo "Certs generated in: $OUT"
ls -la "$OUT"
``

/*
---

## Dosya: `server/main.go`
*/
``go
package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":8443", "listen address")
	certFile := flag.String("cert", "../scripts/ca/server_cert.pem", "server cert PEM")
	keyFile := flag.String("key", "../scripts/ca/server_key.pem", "server key PEM")
	caFile := flag.String("ca", "../scripts/ca/ca.pem", "CA cert to verify clients (for mutual TLS)")
	flag.Parse()

	mutual := false
	if v := os.Getenv("MUTUAL"); v == "true" || v == "1" {
		mutual = true
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.TLS != nil && len(r.TLS.PeerCertificates) > 0 {
			fmt.Fprintf(w, "Hello, mutual TLS client CN=%s\n", r.TLS.PeerCertificates[0].Subject.CommonName)
			return
		}
		w.Write([]byte("Hello, TLS world!\n"))
	})

	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Fatalf("failed to load server key pair: %v", err)
	}

	cfg := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}

	if mutual {
		// load CA pool for client cert verification
		caPEM, err := ioutil.ReadFile(*caFile)
		if err != nil {
			log.Fatalf("failed to read CA file: %v", err)
		}
		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM(caPEM) {
			log.Fatalf("failed to append CA cert")
		}
		cfg.ClientCAs = pool
		cfg.ClientAuth = tls.RequireAndVerifyClientCert
		log.Println("Mutual TLS: enabled (require client cert)")
	} else {
		log.Println("Mutual TLS: disabled")
	}

	server := &http.Server{
		Addr:      *addr,
		Handler:   mux,
		TLSConfig: cfg,
	}

	log.Printf("Listening on %s (TLS)\n", *addr)
	log.Fatal(server.ListenAndServeTLS("", "")) // certs come from TLSConfig
}
``
/*
---

## Dosya: `client/main.go`
*/
``go
package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	caFile := flag.String("ca", "../scripts/ca/ca.pem", "CA cert to trust")
	certFile := flag.String("cert", "", "client cert PEM (for mutual TLS)")
	keyFile := flag.String("key", "", "client key PEM (for mutual TLS)")
	url := flag.String("url", "https://localhost:8443/", "server URL")
	mutual := flag.Bool("mutual", false, "use client cert (mutual TLS)")
	flag.Parse()

	// load CA
	caPEM, err := ioutil.ReadFile(*caFile)
	if err != nil {
		log.Fatalf("failed to read CA file: %v", err)
	}
	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(caPEM) {
		log.Fatalf("failed to append CA cert")
	}

	cfg := &tls.Config{
		RootCAs:    pool,
		MinVersion: tls.VersionTLS12,
		// ServerName: "localhost", // set if CN/SAN differs
	}

	if *mutual {
		if *certFile == "" || *keyFile == "" {
			log.Fatalln("mutual mode requires --cert and --key")
		}
		cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("failed to load client key pair: %v", err)
		}
		cfg.Certificates = []tls.Certificate{cert}
	}

	tr := &http.Transport{TLSClientConfig: cfg}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(*url)
	if err != nil {
		log.Fatalf("GET error: %v", err)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("status: %s\n", resp.Status)
	fmt.Printf("body:\n%s\n", string(b))

	os.Exit(0)
}
``

---

## Dosya: `Makefile` (opsiyonel)

``makefile
.PHONY: certs server client client-mutual

certs:
	chmod +x scripts/generate_certs.sh
	./scripts/generate_certs.sh

server:
	cd server && go run main.go

server-mutual:
	cd server && MUTUAL=true go run main.go

client:
	cd client && go run main.go --ca ../scripts/ca/ca.pem

client-mutual:
	cd client && go run main.go --ca ../scripts/ca/ca.pem --cert ../scripts/ca/client_cert.pem --key ../scripts/ca/client_key.pem --mutual
``
/*
---

### Son notlar

* Bu proje test/dev amaçlıdır. Üretimde sertifika yönetimi için Let's Encrypt, cert-manager veya HashiCorp Vault gibi çözümler tercih edin.
* `scripts/generate_certs.sh` OpenSSL kullanır; Windows kullanıcıları WSL veya benzeri bir ortamda çalıştırabilir.

---

Başka bir isteğin olursa (örneğin: Dockerfile ekleme, systemd servis, veya Let's Encrypt entegrasyonu) söyle, ona göre güncelleyeyim.
EVET
*/
# Go TLS Full Project

Bu proje **Docker** ve **Docker Compose** kullanılarak container içinde çalıştırılabilir hale getirildi.

## 📌 Dockerfile

Proje kök dizinine `Dockerfile` ekleyin:

``dockerfile
# 1. Go image kullanılıyor
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Go mod dosyalarını kopyala
COPY go.mod go.sum ./
RUN go mod download

# Server ve Client kodlarını kopyala
COPY server/ ./server/
COPY client/ ./client/

# Sertifika üretme scriptini kopyala
COPY scripts/ ./scripts/

# OpenSSL ekle (sertifika üretimi için)
RUN apk add --no-cache openssl bash

# Sertifikaları oluştur
RUN bash scripts/generate_certs.sh

# Server binary oluştur
RUN cd server && go build -o /server

# Client binary oluştur
RUN cd client && go build -o /client

# 2. Final image
FROM alpine:latest
WORKDIR /app

# Sertifikaları kopyala
COPY --from=builder /app/server/server.crt ./
COPY --from=builder /app/server/server.key ./

# Binaryleri kopyala
COPY --from=builder /server ./server
COPY --from=builder /client ./client

# Varsayılan olarak server çalıştırılır
CMD ["./server"]
``

/*
---

## 📌 docker-compose.yml

Proje kök dizinine `docker-compose.yml` ekleyin:
*/
``yaml
version: '3.9'

services:
  server:
    build: .
    container_name: tls_server
    ports:
      - "8443:8443"
    command: ["./server"]

  client:
    build: .
    container_name: tls_client
    depends_on:
      - server
    network_mode: host
    command: ["./client"]
``
/*
---

## 📌 Çalıştırma

### 1. Tüm servisleri başlat

```bash
docker-compose up --build
```

### 2. Çalışan containerları görmek için

```bash
docker ps
```

### 3. Server ve Client loglarını takip etmek için

```bash
docker-compose logs -f
```

---

✅ Artık hem **TLS Server** hem de **TLS Client**, `docker-compose` ile aynı anda çalışıyor.

İstersen client ile server arasında **iki yönlü TLS (mutual TLS)** desteğini de ekleyelim. Bunu ister misin?
*/