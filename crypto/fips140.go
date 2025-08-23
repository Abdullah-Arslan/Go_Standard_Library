/*
İşte Go’daki **`crypto/fips140`** modülü hakkında kapsamlı ve güncel bir açıklama.

---

## 1. `crypto/fips140` Paketi Nedir?

Go’da doğrudan "fips140 paketi" denildiğinde kast edilen şey, aslında standart kütüphanedeki gizli (`internal`) bir kriptografik modül olup, FIPS 140-3 standartlarına uygun algoritmaları sağlar. Kamuya açık bir API içeren tek dış paket ise **`crypto/fips140`** paketidir ve sadece bir fonksiyon sunar: `Enabled()`.

Bu fonksiyon, Go uygulamanızın şu anda FIPS 140-3 modunda çalışıp çalışmadığını bildirir ([pkg.go.dev][1]).

---

## 2. FIPS 140-3 Desteği Go’ya Nasıl Geldi?

### &#x20;Go 1.24 ve Üzeri

Go 1.24 sürümünden itibaren, Go’nun kriptografik kütüphaneleri (ör. `crypto/ecdsa`, `crypto/rand`) **yerel olarak FIPS 140-3 moduna uygun (validated)** algoritmalar içeren **Go Cryptographic Module** tabanlı şekilde çalışacak şekilde güncellendi ([go.dev][2]).

### FIPS Modunu Etkinleştirme

* Çalışma zamanında FIPS modu aktif hale getirilebilir:

  ```bash
  export GODEBUG="fips140=on"
  ```
* Daha sıkı bir mod olan `"only"`, FIPS uyumlu olmayan algoritmaları kullanıldığında hata verir veya panik çıkar ([go.dev][2]).

### `GOFIPS140` Build Constraint

* `GOFIPS140=v1.0.0` gibi bir ayar ile, derleme zamanında **Go Cryptographic Module v1.0.0** kullanılacak şekilde binari oluşturulabilir. Bu modül, CMVP’ye gönderilmiş ve FIPS 140-3 kapsamındaki algoritmaları içerir ([go.dev][3]).

---

## 3. `crypto/fips140.Enabled()` Nasıl Çalışır?
*/
``go
import "crypto/fips140"

func main() {
    if fips140.Enabled() {
        fmt.Println("FIPS 140-3 modu aktif.")
    } else {
        fmt.Println("FIPS modu devre dışı.")
    }
}
``
/*
Bu fonksiyon sadece çalışma zamanında modun aktif olup olmadığını bildirir ve `GODEBUG="fips140=on"` gibi ayarlara dayanır ([pkg.go.dev][1]).

---

## 4. Faydalar & Notlar

| Avantaj                        | Açıklama                                                                                                                                                         |
| ------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Yerel Go modülü                | Tümü Go dili ile yazılmış, cgo veya OpenSSL gerektirmiyor ([go.dev][3], [Microsoft for Developers][4]).                                                          |
| Kendiliğinden Kontroller       | Başlangıçta bütünlük testi (checksum), key self-test, DRBG tabanlı rastgelelik gibi FIPS gereksinimleri yerine getiriliyor ([go.dev][2]).                        |
| Standart Kitaplıkla Bütünleşik | `crypto/tls`, `crypto/rand`, `crypto/ecdsa` gibi paketler FIPS moduna göre davranıyor (örneğin TLS sadece FIPS-onaylı algoritmaları kabul ediyor) ([go.dev][2]). |

> Önemli: FIPS modu yasal ve regülasyon gereksinimlerini tek başına karşılamaz. Uygulamanızın ve ortamınızın da uyumlu olması gerekir ([go.dev][2]).

---

## Özet

* **`crypto/fips140`** paketi sadece `Enabled()` fonksiyonunu sunar, FIPS modunun aktifliğini kontrol eder.
* **Go 1.24+** sürümleri, yerel **FIPS 140-3 uyumlu mod** ile kriptografi sağlar.
* **`GODEBUG=fips140=on`** veya **`GOFIPS140` build ayarları** ile bu modu aktif hale getirebilirsiniz.
* Tüm bu mekanizma Go’nun standart kütüphaneleriyle entegre çalışır ve cgo veya üçüncü parti kütüphane gerektirmez.

---

Dilersen, bu modun nasıl heyecan verici bir yönde kullanılabileceğini gösteren bir **demo HTTP sunucusu örneği** hazırlayabilirim — örneğin `crypto/tls` ile FIPS-onaylı cipher suite’lerle çalışan bir TLS sunucusu. Bunu ister misin?

[1]: https://pkg.go.dev/crypto/fips140?utm_source=chatgpt.com "crypto/fips140"
[2]: https://go.dev/doc/security/fips140?utm_source=chatgpt.com "FIPS 140-3 Compliance"
[3]: https://go.dev/blog/fips140?utm_source=chatgpt.com "The FIPS 140-3 Go Cryptographic Module"
[4]: https://devblogs.microsoft.com/go/go-1-24-fips-update/?utm_source=chatgpt.com "Microsoft build of Go 1.24 FIPS changes"
EVET
*/

/*
Süper — aşağıda **Go (1.24+) ile FIPS-140-3 modunda çalışan örnek bir TLS sunucusu** var: önce hangi ortam/derleme ayarlarını kullanman gerektiğini, sonra kodu, nasıl çalıştıracağını ve neler olup bittiğini adım adım anlattım.

Kısa özet (temel noktalar, kaynaklarla):

* Go 1.24+ ile gelen Go Cryptographic Module sayesinde programını **FIPS 140-3 moduna** alabilirsin. Bu çalışma zamanı modunu `GODEBUG="fips140=on"` ile açarsın; `fips140=only` daha katıdır ve FIPS-uyumlu olmayan algoritma kullanılırsa hata/panik üretebilir. ([Go][1], [Microsoft for Developers][2])
* `crypto/fips140.Enabled()` ile programın FIPS modunda olup olmadığını öğrenebilirsin. ([Go][3])
* `crypto/tls` paketi FIPS modunda yalnızca FIPS-onaylı protokolleri/algoritmaları pazarlık edecektir; bu kısıtlama otomatik uygulanır (ve Go sürümüne / seçilen Go Cryptographic Module sürümüne göre değişebilir). ([Go Packages][4], [Go][1])

---

## 1) Nasıl çalıştırılır (önerilen adımlar)

1. Go sürümü **1.24 veya daha yeni** olmalı. (FIPS mod desteği 1.24+). ([Go][1])
2. Terminalde (runtime FIPS modu):
*/

``bash
export GODEBUG="fips140=on"
# veya daha katı:
# export GODEBUG="fips140=only"
``
/*
3. Sunucuyu çalıştır: `go run fips_tls_server.go` (kod aşağıda).
4. Tarayıcı/`curl` ile bağlan: `curl -k https://localhost:8443/` (kendi ürettiğimiz self-signed sertifika olduğu için `-k`/insecure gerekli).

> Opsiyonel: build-time seçimi yapmak istiyorsan `GOFIPS140` ile özel Go Cryptographic Module versiyonunu seçebilirsin (örneğin CI ile). ([tip.golang.org][5], [Microsoft for Developers][2])

---

## 2) Tam kod — `fips_tls_server.go`
*/
``go
// fips_tls_server.go
// Minimal TLS server that prefers FIPS-approved settings.
// Requires Go 1.24+. Run with: GODEBUG="fips140=on" go run fips_tls_server.go
package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/fips140" // to check status at runtime (optional)
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"
)

func main() {
	// Info: is FIPS mode enabled right now?
	fmt.Println("crypto/fips140.Enabled() ->", fips140.Enabled())

	// create a short-lived self-signed ECDSA P-256 cert (P-256 is FIPS-allowed)
	certPEM, keyPEM, err := makeSelfSignedECDSACert()
	if err != nil {
		log.Fatal(err)
	}

	// tls.Certificate from PEM
	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		log.Fatal(err)
	}

	// TLS config: explicitly restrict to strong, FIPS-typical suites.
	// In FIPS mode crypto/tls will further restrict what it actually negotiates.
	tlsCfg := &tls.Config{
		Certificates: []tls.Certificate{cert},

		// Prefer modern TLS versions. In FIPS mode crypto/tls may restrict.
		MinVersion: tls.VersionTLS12,
		MaxVersion: tls.VersionTLS13,

		// Prefer server cipher order (affects some TLS1.2 negotiations).
		PreferServerCipherSuites: true,

		// FIPS-friendly TLS 1.2 suites (ECDHE + AES-GCM). TLS1.3 is handled internally.
		// Note: crypto/tls may silently ignore suites not allowed in FIPS mode.
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			// TLS 1.3 suites are negotiated implicitly; you don't list them here.
		},

		// Prefer P-256 and P-384 curves (both FIPS-allowed depending on policy).
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.CurveP384,
		},
	}

	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsCfg,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "hello — FIPS mode: %v\n", fips140.Enabled())
		}),
	}

	log.Println("Listening on https://localhost:8443 — press Ctrl+C to stop")
	// Use TLSConfig certs (we pass empty certFile/keyFile because tlsCfg holds certs)
	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatal(err)
	}
}

// makeSelfSignedECDSACert creates a self-signed ECDSA P-256 cert (DER PEMs).
func makeSelfSignedECDSACert() (certPEM []byte, keyPEM []byte, err error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	serial, _ := rand.Int(rand.Reader, big.NewInt(1<<62))
	tmpl := x509.Certificate{
		SerialNumber: serial,
		Subject: pkix.Name{
			CommonName:   "localhost",
			Organization: []string{"Example Org"},
		},
		NotBefore:             time.Now().Add(-time.Minute),
		NotAfter:              time.Now().Add(24 * time.Hour),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		DNSNames:              []string{"localhost"},
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, &priv.PublicKey, priv)
	if err != nil {
		return nil, nil, err
	}

	// cert PEM
	certPEM = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	// key PEM (PKCS8 or EC PRIVATE KEY; EC PRIVATE KEY is fine)
	b, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		return nil, nil, err
	}
	keyPEM = pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: b})
	return certPEM, keyPEM, nil
}
``
/*
**Ne yapıyor bu kod?**

* Program başında `crypto/fips140.Enabled()` ile FIPS mod durumunu yazdırır. ([Go][3])
* ECDSA P-256 ile self-signed sertifika üretiyor (P-256 FIPS tarafından yaygın olarak desteklenen bir eğridir; kurum politikana göre tercih değişebilir).
* `tls.Config` içinde TLS versiyonlarını, `CipherSuites` ve `CurvePreferences` ile güçlü (ECDHE + AES-GCM) tercihini koyuyor. **Not:** Go’nun `crypto/tls` paketi FIPS modunda kendi kurallarını uygular; yani eğer bir suite FIPS dışıysa o suite otomatik göz ardı edilir. ([Go Packages][4], [Go][1])

---

## 3) Önemli notlar ve tavsiyeler

* **Go sürümü**: Bu mekanizma (yerel FIPS mod) Go 1.24+ ile gelmiştir; dolayısıyla sistemdeki Go versiyonunu kontrol et. ([Go][1])
* **Kendi ortam uyumluluğu**: FIPS uyumluluğu yalnızca `GODEBUG`/`GOFIPS140` değişkenlerini ayarlamakla bitmez — kuruluşunun güvenlik süreçleri, sertifika zinciri, işletim sistemi yapılandırması ve CMVP gereksinimleri de göz önünde bulundurulmalıdır. ([Go][1], [csrc.nist.gov][6])
* **`fips140=only`** modu daha katıdır: FIPS olmayan algoritma çağrısı yapılırsa hata/panik üretir; `fips140=on` daha esnektir (uygunsuz algoritmaları yoksayar veya TLS müzakeresinden çıkarır). Bu davranış Go sürümleri arasında olası küçük farklar gösterebilir; test et. ([Microsoft for Developers][2], [Medium][7])
* **TLS 1.3**: TLS 1.3 cipher-suite yönetimi farklı çalışır — Go `crypto/tls` kendi uygun TLS1.3 listelemesini yapar; FIPS modu TLS1.3 tarafında da uygun seçimler yapacaktır (sürüm/donanım ve seçilen Go Cryptographic Module etkili). ([Go Packages][4])
* **Testetmeden üretime geçme**: FIPS modunda davranış farklı olabilir (ör. bazı ciphers/curves reddedilebilir). Üretime almadan önce test ortamında `GODEBUG` ile kapsamlı test yap. Birçok proje `fips140=only` ile test edip uyumsuz noktalarda kod değişimi yapıyor. ([GitHub][8])

---

## 4) İleri adımlar (isteğe bağlı, yapabilirim)

* Sunucuya **mutual TLS** (mTLS) ekleyip FIPS uyumlu client sertifikalarını zorlamak.
* `GOFIPS140` ile build-time module seçimi ve bir CI örneği (Dockerfile + build args) hazırlamak.
* Sistemdeki sertifika ve anahtar yönetimini (HSM / PKCS#11) nasıl entegre edebileceğini göstermek.
* `fips140=only` kullanılarak ortaya çıkabilecek uyumsuz noktaları bulup çözümlemeye yardımcı olacak test senaryoları oluşturmak.

Hangisini istersin — mTLS mi, Docker + GOFIPS140 CI örneği mi, yoksa `fips140=only` ile test senaryoları mı hazırlayayım?

[1]: https://go.dev/blog/fips140?utm_source=chatgpt.com "The FIPS 140-3 Go Cryptographic Module"
[2]: https://devblogs.microsoft.com/go/go-1-24-fips-update/?utm_source=chatgpt.com "Microsoft build of Go 1.24 FIPS changes"
[3]: https://go.dev/doc/security/fips140?utm_source=chatgpt.com "FIPS 140-3 Compliance"
[4]: https://pkg.go.dev/crypto/tls?utm_source=chatgpt.com "tls package - crypto/tls"
[5]: https://tip.golang.org/doc/go1.24?utm_source=chatgpt.com "Go 1.24 Release Notes"
[6]: https://csrc.nist.gov/csrc/media/Projects/cryptographic-module-validation-program/documents/fips%20140-3/FIPS%20140-3%20IG.pdf?utm_source=chatgpt.com "Implementation Guidance for FIPS 140-3"
[7]: https://medium.com/%40moeghifar/go-1-24-the-game-changing-security-and-cryptography-release-6ee495742be6?utm_source=chatgpt.com "Go 1.24: The Game-Changing Security and Cryptography ..."
[8]: https://github.com/quic-go/quic-go/issues/4894?utm_source=chatgpt.com "FIPS Compliance Issues with Go 1.24 #4894"
*/
