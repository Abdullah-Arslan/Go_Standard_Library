/*
Go’nun **`net/smtp`** paketini baştan sona açıklayayım ve örneklerle göstereyim.

`net/smtp` paketi, **Simple Mail Transfer Protocol (SMTP)** kullanarak e-posta göndermek için kullanılır.

---

# 1️⃣ `net/smtp` Paketi Temel Özellikleri

| Fonksiyon / Tip        | Açıklama                                                  |
| ---------------------- | --------------------------------------------------------- |
| `smtp.SendMail`        | SMTP sunucusu üzerinden e-posta gönderir                  |
| `smtp.PlainAuth`       | SMTP kimlik doğrulama (username/password) sağlar          |
| `smtp.Client`          | SMTP sunucusu ile bağlantı ve daha detaylı kontrol sağlar |
| `Mail`, `Rcpt`, `Data` | SMTP protokolü üzerinden manuel e-posta gönderimi sağlar  |

---

# 2️⃣ Temel Kullanım: `smtp.SendMail`
*/
``go
package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	// SMTP sunucu bilgileri
	smtpHost := "smtp.example.com"
	smtpPort := "587"

	// Gönderici bilgileri
	from := "gonderici@example.com"
	password := "sifre123"

	// Alıcı
	to := []string{"alici@example.com"}

	// E-posta içeriği
	message := []byte("Subject: Merhaba!\r\n" +
		"\r\n" +
		"Bu bir test e-postasıdır.\r\n")

	// SMTP kimlik doğrulama
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// E-posta gönder
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println("Gönderim hatası:", err)
		return
	}

	fmt.Println("E-posta gönderildi!")
}
``
/*
✅ Açıklamalar:

* `smtp.SendMail` en basit yöntemdir, tek satırda e-posta gönderir.
* `PlainAuth` kullanıcı adı ve şifre ile kimlik doğrulama yapar.
* `message` başlık (`Subject`) ve gövde (`Body`) içerir.

---

# 3️⃣ Detaylı Kullanım: `smtp.Client`
*/
``go
package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	smtpHost := "smtp.example.com"
	smtpPort := "587"
	from := "gonderici@example.com"
	password := "sifre123"
	to := "alici@example.com"

	// Sunucuya bağlan
	client, err := smtp.Dial(smtpHost + ":" + smtpPort)
	if err != nil {
		fmt.Println("Dial hatası:", err)
		return
	}
	defer client.Close()

	// Kimlik doğrulama
	auth := smtp.PlainAuth("", from, password, smtpHost)
	if err := client.Auth(auth); err != nil {
		fmt.Println("Auth hatası:", err)
		return
	}

	// Gönderici ve alıcı
	if err := client.Mail(from); err != nil {
		fmt.Println("Mail hatası:", err)
		return
	}
	if err := client.Rcpt(to); err != nil {
		fmt.Println("Rcpt hatası:", err)
		return
	}

	// Mesaj gönderimi
	wc, err := client.Data()
	if err != nil {
		fmt.Println("Data hatası:", err)
		return
	}

	message := "Subject: Merhaba!\r\n" +
		"\r\n" +
		"Bu, smtp.Client ile gönderilen test e-postasıdır.\r\n"

	_, err = wc.Write([]byte(message))
	if err != nil {
		fmt.Println("Write hatası:", err)
		return
	}
	wc.Close()

	fmt.Println("E-posta gönderildi!")
}
``
/*
✅ Açıklamalar:

* `smtp.Dial` ile sunucuya bağlanılır.
* `Mail` ve `Rcpt` ile gönderen ve alıcı tanımlanır.
* `Data` ile mesaj yazılır.
* Bu yöntem **daha fazla kontrol ve logging** sağlar.

---

# 4️⃣ Önemli Notlar

1. **Portlar ve TLS:**

   * Genellikle SMTP sunucuları 587 (STARTTLS) veya 465 (SSL/TLS) kullanır.
   * `net/smtp` paketinde SSL/TLS için `tls.Dial` ve `smtp.NewClient` birlikte kullanılır.

2. **E-posta Başlıkları:**

   * `Subject`, `From`, `To`, `Cc`, `Bcc` gibi başlıkları mesaj içinde manuel eklemelisiniz.

3. **Alternatifler:**

   * `net/smtp` artık **deprecated** kabul ediliyor, modern projelerde [`github.com/jordan-wright/email`](https://github.com/jordan-wright/email) veya diğer SMTP kütüphaneleri tercih edilebilir.

---

İstersen sana bunu bir adım ileri götürüp **TLS + concurrent e-posta gönderimi + HTML e-posta + attachment** örneğini de gösterebilirim.

Bunu ister misin?
EVET
*/
/*
Harika! 👍 O zaman sana **Go ile TLS destekli, concurrent, HTML ve attachment destekli e-posta gönderimi** örneğini göstereyim.

Bu örnek, gerçek projelerde güvenli ve gelişmiş e-posta gönderimi için kullanılabilir.

---

## 1️⃣ Paketler
*/
``go
package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"sync"

	"github.com/jordan-wright/email"
)
``
/*
> Not: Burada [`github.com/jordan-wright/email`](https://github.com/jordan-wright/email) kütüphanesini kullanıyoruz. Bu, `net/smtp`’yi daha modern ve kolay bir şekilde kullanmamızı sağlıyor.

---

## 2️⃣ Temel Konfigürasyon
*/
``go
func main() {
	smtpHost := "smtp.example.com"
	smtpPort := 587
	username := "gonderici@example.com"
	password := "sifre123"

	// TLS yapılandırması
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true, // Test için; prod ortamda false olmalı
		ServerName:         smtpHost,
	}

	numClients := 5
	var wg sync.WaitGroup
``
/*
---

## 3️⃣ Concurrent HTML + Attachment Gönderimi
*/
``go
	for i := 1; i <= numClients; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			e := email.NewEmail()
			e.From = fmt.Sprintf("Gönderici <%s>", username)
			e.To = []string{fmt.Sprintf("alici%d@example.com", id)}
			e.Subject = fmt.Sprintf("Merhaba %d!", id)
			e.HTML = []byte(fmt.Sprintf("<h1>Merhaba Client %d!</h1><p>Bu bir HTML test e-postasıdır.</p>", id))
			e.Text = []byte("Bu bir test e-postasıdır.")
			e.AttachFile("test.txt") // test.txt mevcut olmalı

			// TLS üzerinden gönder
			err := e.SendWithTLS(fmt.Sprintf("%s:%d", smtpHost, smtpPort),
				smtp.PlainAuth("", username, password, smtpHost),
				tlsconfig)
			if err != nil {
				log.Printf("Client %d gönderim hatası: %v\n", id, err)
				return
			}

			log.Printf("Client %d e-posta gönderildi!\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("Tüm e-postalar gönderildi!")
}
``
/*
---

## 4️⃣ Açıklamalar

1. **Concurrent Gönderim:**

   * `numClients` kadar goroutine ile eş zamanlı e-posta gönderiliyor.
   * Worker pool veya kanallar ile daha büyük ölçek yönetilebilir.

2. **HTML ve Text:**

   * `e.HTML` → HTML içeriği
   * `e.Text` → düz metin (fallback için)

3. **Attachment:**

   * `e.AttachFile("dosya.txt")` ile dosya eklenebilir.

4. **TLS Güvenliği:**

   * `tls.Config` ile SMTP sunucusuna güvenli bağlantı kuruluyor.
   * `InsecureSkipVerify: false` olmalı prod ortamda.

5. **Modern Paket:**

   * `net/smtp` yerine `github.com/jordan-wright/email` kullanmak, HTML, attachment ve TLS yönetimini kolaylaştırıyor.

---

## 5️⃣ Örnek Çıktı
*/
``
2025/09/11 12:45: Client 1 e-posta gönderildi!
2025/09/11 12:45: Client 2 e-posta gönderildi!
2025/09/11 12:45: Client 3 e-posta gönderildi!
...
Tüm e-postalar gönderildi!
``
/*
* E-postalar eş zamanlı gönderiliyor.
* HTML ve attachment destekli.
* TLS ile güvenli bağlantı sağlanıyor.

---

Bu yapı üzerine bir adım daha ekleyip **SMTP rate limit, retry mekanizması ve logging + pprof ile performans takibi** ekleyebiliriz.

Bunu da ister misin?
EVET
*/
/*
Harika! 👍 O zaman sana **Go ile TLS destekli, concurrent e-posta gönderimi + rate limiting + retry mekanizması + logging + pprof ile performans takibi** örneğini göstereyim.

Bu örnek, gerçek dünya projelerinde **yüksek hacimli e-posta gönderimi** için kullanışlıdır.

---

## 1️⃣ Paketler ve Hazırlık
*/
``go
package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
	_ "net/http/pprof"
	"net/http"

	"github.com/jordan-wright/email"
)
``
/*
* `pprof` ile CPU/memory/goroutine profilini alabiliriz (`http://localhost:6060/debug/pprof/`).

---

## 2️⃣ Temel Konfigürasyon
*/
``go
func main() {
	// pprof HTTP server
	go func() {
		log.Println("pprof server :6060 çalışıyor")
		http.ListenAndServe(":6060", nil)
	}()

	smtpHost := "smtp.example.com"
	smtpPort := 587
	username := "gonderici@example.com"
	password := "sifre123"

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}

	numClients := 10
	rateLimit := 3 // saniyede 3 e-posta
	retryCount := 3
	var wg sync.WaitGroup

	sem := make(chan struct{}, rateLimit) // rate limiting kanal
``
/*
---

## 3️⃣ Concurrent E-posta Gönderimi + Retry + Rate Limiting
*/
``go
	for i := 1; i <= numClients; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			sem <- struct{}{}           // rate limiting
			defer func() { <-sem }()    // release semaphore

			for attempt := 1; attempt <= retryCount; attempt++ {
				e := email.NewEmail()
				e.From = fmt.Sprintf("Gönderici <%s>", username)
				e.To = []string{fmt.Sprintf("alici%d@example.com", id)}
				e.Subject = fmt.Sprintf("Merhaba %d!", id)
				e.HTML = []byte(fmt.Sprintf("<h1>Merhaba Client %d!</h1>", id))
				e.Text = []byte("Bu bir test e-postasıdır.")

				err := e.SendWithTLS(fmt.Sprintf("%s:%d", smtpHost, smtpPort),
					smtp.PlainAuth("", username, password, smtpHost),
					tlsConfig)

				if err != nil {
					log.Printf("Client %d attempt %d hatası: %v", id, attempt, err)
					time.Sleep(500 * time.Millisecond) // retry bekleme
					continue
				} else {
					log.Printf("Client %d e-posta gönderildi!", id)
					break
				}
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Tüm e-postalar gönderildi!")
}
``
/*
---

## 4️⃣ Açıklamalar

1. **Concurrent Gönderim + Worker Pool:**

   * Her client için ayrı goroutine → concurrency sağlanıyor.

2. **Rate Limiting:**

   * `sem := make(chan struct{}, rateLimit)` → aynı anda `rateLimit` kadar e-posta gönderilebilir.

3. **Retry Mekanizması:**

   * Hata olursa `retryCount` kadar tekrar gönderim denemesi yapılıyor.
   * `time.Sleep` ile kısa bekleme ekleniyor.

4. **Logging:**

   * Gönderim ve hatalar loglanıyor → debugging kolay.

5. **pprof ile Profil:**

   * `http://localhost:6060/debug/pprof/` üzerinden CPU/memory/goroutine profili alınabilir.

6. **TLS:**

   * Güvenli bağlantı sağlanıyor (`SendWithTLS`).

---

## 5️⃣ Örnek Çıktı
*/
``
2025/09/11 13:10: Client 1 e-posta gönderildi!
2025/09/11 13:10: Client 2 e-posta gönderildi!
2025/09/11 13:10: Client 3 attempt 1 hatası: smtp: TLS handshake timeout
2025/09/11 13:10: Client 3 e-posta gönderildi!
...
Tüm e-postalar gönderildi!
``
/*
* Rate limiting sayesinde aynı anda sadece `rateLimit` kadar e-posta gönderiliyor.
* Retry ile başarısız denemeler tekrar deneniyor.
* pprof ile performans analizini yapabilirsin.

---

Bu yapı ile artık **yüksek hacimli, güvenli, retry ve rate-limit kontrollü e-posta gönderimi** yapabilirsin.

İstersen sana bunu bir adım daha ileri götürüp **concurrent JSON-RPC + SMTP + pprof + worker pool + retry + timeout** şeklinde birleşik bir framework örneğini gösterebilirim.

Bunu ister misin?
*/