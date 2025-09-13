/*
Golang’da standart kütüphane içinde **`net/mail`** paketi bulunur.
Bu paket, e-posta (mail) mesajlarını **parçalamak (parse etmek)**, **header** bilgilerini okumak ve **adresleri işlemek** için kullanılır.

📌 Önemli nokta: `net/mail` paketi **mail göndermez**. Mail göndermek için `net/smtp` gibi paketler kullanılır. `net/mail` yalnızca mail formatını çözümlemek ve adreslerle çalışmak için vardır.

Şimdi adım adım tüm paketi açıklayayım.

---

# 📦 `net/mail` Paketinin İçeriği

### 1. `Address` Yapısı

Bir e-posta adresini temsil eder.
*/
``go
type Address struct {
    Name    string // Görünen ad (örn: "Ali Veli")
    Address string // Gerçek mail adresi (örn: "ali@example.com")
}
``

//🔹 Örnek:

``go
package main

import (
	"fmt"
	"net/mail"
)

func main() {
	addr := &mail.Address{
		Name:    "Ali Veli",
		Address: "ali@example.com",
	}
	fmt.Println(addr.String()) // "Ali Veli <ali@example.com>"
}
``
/*
---

### 2. `ParseAddress`

Bir string ifadesini alıp `Address` nesnesine dönüştürür.
*/
``go
addr, _ := mail.ParseAddress("Ali Veli <ali@example.com>")
fmt.Println("İsim:", addr.Name)
fmt.Println("Adres:", addr.Address)
``
/*
---

### 3. `ParseAddressList`

Birden fazla adresi ayrıştırır.
*/
``go
addresses, _ := mail.ParseAddressList("Ali <ali@example.com>, Veli <veli@example.com>")
for _, a := range addresses {
    fmt.Println(a.Name, a.Address)
}
``
/*
---

### 4. `Message` Yapısı

Bir e-posta mesajını temsil eder:
*/
``go
type Message struct {
    Header Header      // Mail header bilgileri
    Body   io.Reader   // Mail gövdesi (mesaj içeriği)
}
``
/*
---

### 5. `Header` Tipi

Mail başlıklarını (`From`, `To`, `Subject`, `Date`, …) tutar.

* `Get(key string)` → Header değerini alır.
* `AddressList(key string)` → Bir header’daki adresleri listeler.
* `Date()` → `Date` header’ını zaman tipine çevirir.

🔹 Örnek:
*/
``go
msg, _ := mail.ReadMessage(strings.NewReader(
	"From: Ali <ali@example.com>\r\n" +
	"To: Veli <veli@example.com>\r\n" +
	"Subject: Selam\r\n" +
	"\r\n" +
	"Merhaba Veli!"))

from, _ := msg.Header.AddressList("From")
to, _ := msg.Header.AddressList("To")
subject := msg.Header.Get("Subject")

fmt.Println("Kimden:", from[0].String())
fmt.Println("Kime:", to[0].String())
fmt.Println("Konu:", subject)
``
/*
---

### 6. `ReadMessage`

Bir `io.Reader` (dosya, string, network) üzerinden mail mesajı okur ve `Message` döner.
*/
``go
r := strings.NewReader("Subject: Test\r\n\r\nMerhaba!")
msg, _ := mail.ReadMessage(r)

fmt.Println("Subject:", msg.Header.Get("Subject"))
body, _ := io.ReadAll(msg.Body)
fmt.Println("Body:", string(body))
``
/*
---

### 7. `ParseDate`

RFC 5322 uyumlu bir tarih string’ini `time.Time` nesnesine dönüştürür.
*/
``go
t, _ := mail.ParseDate("Mon, 02 Jan 2006 15:04:05 -0700")
fmt.Println("Tarih:", t)
``
/*
---

# 📌 Özet Fonksiyonlar & Yapılar
*/

| Fonksiyon / Yapı     | Açıklama                                      |
| -------------------- | --------------------------------------------- |
| `Address`            | Bir e-posta adresi (İsim + Adres).            |
| `Address.String()`   | Adresi `Name <email>` formatında döner.       |
| `ParseAddress`       | Tek bir e-posta adresi ayrıştırır.            |
| `ParseAddressList`   | Virgülle ayrılmış adres listesini ayrıştırır. |
| `Message`            | Bir mail mesajı (`Header` + `Body`).          |
| `ReadMessage`        | Bir `io.Reader`’dan mail mesajı okur.         |
| `Header`             | Mail başlıklarını tutar.                      |
| `Header.Get`         | Header değerini alır.                         |
| `Header.AddressList` | Header içindeki adresleri döner.              |
| `Header.Date`        | Date header’ını `time.Time` tipine çevirir.   |
| `ParseDate`          | String tarihi `time.Time` tipine çevirir.     |

/*
---

# 🎯 Kısaca

* `net/mail`, **mail parse etme ve adreslerle çalışma** için kullanılır.
* **Mail göndermez**, sadece RFC 5322 uyumlu mesajları ayrıştırır.
* Daha çok **SMTP ile gelen mailleri okumak** veya **dosyadan parse etmek** için kullanılır.

---

İstersen ben sana bunun üstüne bir de **SMTP ile mail gönderme** kısmını (`net/smtp` paketi ile) örnekleyebilirim. Onu da ister misin?
EVET
*/
/*
Harika 👍 o zaman şimdi sana Go’da **`net/smtp`** paketi ile e-posta gönderme örneğini göstereyim.

---

## 1. Temel SMTP Mail Gönderme
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
	smtpPort := "587" // veya 465 (SSL)
	username := "your_email@example.com"
	password := "your_password"

	// Gönderen ve alıcı bilgileri
	from := "your_email@example.com"
	to := []string{"recipient@example.com"}

	// Mail içeriği (RFC 5322 formatına uygun)
	message := []byte("Subject: Selam!\r\n" +
		"\r\n" +
		"Merhaba, bu bir Go mail örneğidir.\r\n")

	// SMTP kimlik doğrulaması
	auth := smtp.PlainAuth("", username, password, smtpHost)

	// Mail gönder
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	fmt.Println("Mail gönderildi!")
}
``
/*
---

## 2. Açıklamalar

* `smtpHost` ve `smtpPort`: SMTP sunucusunun adresi ve portu. Örnek: Gmail için `smtp.gmail.com:587`.
* `username` ve `password`: SMTP login bilgisi. Gmail’de “App Password” kullanmalısın.
* `from`: Gönderen e-posta adresi.
* `to`: Alıcı adresleri slice olarak verilir.
* `message`: Mail başlığı ve gövdesini içerir.

  * `Subject: ...` mutlaka ilk satırda olmalı.
  * Başlık ve gövde arasında boş bir satır olmalı (`\r\n`).

---

## 3. HTML Mail Gönderme
*/
``go
message := []byte("Subject: HTML Örnek\r\n" +
	"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
	"\r\n" +
	"<h1>Merhaba!</h1><p>Bu bir HTML mail örneğidir.</p>")
``
/*
* `Content-Type: text/html` ile HTML formatında mail gönderebilirsin.

---

## 4. Birden Fazla Alıcıya Mail
*/
``go
to := []string{"ali@example.com", "veli@example.com"}
err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
``
/*
* Mail tüm alıcılara gönderilir.

---

## 5. TLS Kullanımı (SSL)

Güvenli bağlantı gerekiyorsa `smtp.SendMail` yerine `tls` ile bağlantı açmak gerekir.
Bu biraz daha ileri seviye ama basit kullanım için çoğu SMTP sunucusu `587` portunda STARTTLS ile çalışır.

---

✅ Özet:

* `net/mail` → mail içeriğini **okuma ve parse etme** için.
* `net/smtp` → mail **gönderme** için.
* `smtp.SendMail` ile basitçe mail gönderebilir, header ve body’yi kendin oluşturabilirsin.

---

İstersen sana bunu **pprof ile CPU ve memory izleyerek toplu mail gönderme (ör. binlerce mail gönderirken performans analizi)** örneğini de gösterebilirim. Bunu ister misin?
EVET
*/
/*
Harika 👍 o zaman sana Go’da **binlerce mail gönderirken performans analizi** için `pprof` kullanımını göstereyim.

Bu örnekte:

* `net/smtp` ile çok sayıda mail gönderilecek.
* `http/pprof` ile CPU ve memory profili alınacak.
* Böylece **performans darboğazları ve bellek sızıntıları** tespit edilebilecek.

---

## 1. Örnek Kod: Toplu Mail Gönderme + Pprof
*/
``go
package main

import (
	"fmt"
	"net/http"
	"net/smtp"
	_ "net/http/pprof"
	"time"
)

// SMTP ayarları (örnek)
var smtpHost = "smtp.example.com"
var smtpPort = "587"
var username = "your_email@example.com"
var password = "your_password"
var from = "your_email@example.com"

// Mail gönderen fonksiyon
func sendMail(to string, subject string, body string) error {
	msg := []byte("Subject: " + subject + "\r\n" +
		"\r\n" + body)

	auth := smtp.PlainAuth("", username, password, smtpHost)
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
}

// Toplu mail gönder
func bulkMail(toList []string) {
	for _, to := range toList {
		err := sendMail(to, "Deneme Maili", "Bu bir Go pprof test mailidir.")
		if err != nil {
			fmt.Println("Hata:", err)
		}
	}
}

func main() {
	// Pprof endpointlerini başlat
	go func() {
		fmt.Println("Pprof aktif: http://localhost:6060/debug/pprof/")
		http.ListenAndServe(":6060", nil)
	}()

	// Örnek alıcı listesi (binlerce mail için çoğalt)
	toList := []string{"ali@example.com", "veli@example.com", "ayse@example.com"}

	// Toplu mail gönderimini başlat
	for i := 0; i < 1000; i++ { // 1000 tur mail gönder
		bulkMail(toList)
		time.Sleep(100 * time.Millisecond) // kısa bekleme
	}
}
``
/*
---

## 2. Profil Alma

### CPU Profili

Terminalde:
*/
``bash
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
``
/*
* 30 saniyelik CPU kullanımı profili alır.
* `top`, `list` ve `web` komutları ile hangi fonksiyonların en çok CPU kullandığını görebilirsin.

### Heap Profili
*/
``bash
go tool pprof http://localhost:6060/debug/pprof/heap
``
/*
* Bellek kullanımını görebilirsin.
* Eğer mail gönderiminde buffer veya goroutine leak varsa burada ortaya çıkar.

---

## 3. İpuçları

1. **Goroutine sayısı**:

   * Çok fazla eş zamanlı `sendMail` çağrısı goroutine leak veya yoğun CPU yaratabilir.
   * Çözüm: Worker pool kullanmak.

2. **Rate limit (SMTP sınırlamaları)**:

   * SMTP sunucuları aşırı hızlı mail gönderimini kısıtlayabilir.
   * `time.Sleep` veya throttling ile yavaşlatmak gerekir.

3. **Memory sızıntısı kontrolü**:

   * Her mail için oluşturulan byte slice’lar belleği artırır.
   * Profil ile `bulkMail` fonksiyonunun memory kullanımını kontrol et.

---

✅ Sonuç:

* Bu yöntemle **binlerce mail gönderirken CPU ve memory analizini** gerçek zamanlı yapabilirsin.
* `pprof` → darboğaz ve sızıntıları tespit etmede çok kullanışlıdır.
* Böylece sistem kaynaklarını optimize ederek toplu mail gönderimlerini güvenli hale getirebilirsin.

---

İstersen sana **bu örneği worker pool + concurrency ile optimize edip performansı arttırma** versiyonunu da gösterebilirim. Onu da ister misin?
EVET
*/
/*
Harika 👍 o zaman şimdi sana **worker pool ile toplu mail gönderme** örneğini göstereyim. Bu yöntem:

* CPU ve bellek kullanımını dengeler,
* SMTP sunucusuna aşırı yüklenmeyi engeller,
* `pprof` ile performansı analiz etmeyi kolaylaştırır.

---

## 1. Worker Pool ile Mail Gönderme Örneği
*/
``go
package main

import (
	"fmt"
	"net/http"
	"net/smtp"
	_ "net/http/pprof"
	"sync"
	"time"
)

var (
	smtpHost = "smtp.example.com"
	smtpPort = "587"
	username = "your_email@example.com"
	password = "your_password"
	from     = "your_email@example.com"
)

// Mail gönderen fonksiyon
func sendMail(to string, subject string, body string) error {
	msg := []byte("Subject: " + subject + "\r\n" +
		"\r\n" + body)

	auth := smtp.PlainAuth("", username, password, smtpHost)
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
}

// Worker pool fonksiyonu
func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for to := range jobs {
		err := sendMail(to, "Deneme Maili", fmt.Sprintf("Merhaba, bu mail worker %d tarafından gönderildi.", id))
		if err != nil {
			fmt.Printf("Worker %d hata: %v\n", id, err)
		} else {
			fmt.Printf("Worker %d mail gönderdi: %s\n", id, to)
		}
		time.Sleep(50 * time.Millisecond) // SMTP rate limit önlemi
	}
}

func main() {
	// Pprof server
	go func() {
		fmt.Println("Pprof aktif: http://localhost:6060/debug/pprof/")
		http.ListenAndServe(":6060", nil)
	}()

	// Örnek alıcı listesi
	toList := []string{
		"ali@example.com",
		"veli@example.com",
		"ayse@example.com",
		"mehmet@example.com",
	}

	// Worker sayısı
	numWorkers := 3

	// Job kanalı
	jobs := make(chan string, len(toList))

	// WaitGroup ile tüm worker’ları bekle
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Job’ları kanala ekle
	for _, to := range toList {
		for i := 0; i < 100; i++ { // Her alıcıya 100 mail gönder
			jobs <- to
		}
	}
	close(jobs)

	// Tüm worker’lar bitene kadar bekle
	wg.Wait()
	fmt.Println("Tüm mailler gönderildi!")
}
``
/*
---

## 2. Açıklamalar

1. **Worker Pool Mantığı**:

   * `numWorkers` kadar goroutine çalışıyor.
   * `jobs` kanalından mail adreslerini alıyor ve gönderiyor.
   * Böylece çok fazla eş zamanlı SMTP bağlantısı oluşmuyor.

2. **Rate Limiting**:

   * `time.Sleep(50ms)` ile her mail gönderimi arasında kısa gecikme ekleniyor.
   * Bu SMTP sunucusunu aşırı yüklenmekten korur.

3. **Pprof ile Analiz**:

   * CPU ve memory kullanımını gerçek zamanlı olarak izleyebilirsin:
*/
     ``
     http://localhost:6060/debug/pprof/
     go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
     go tool pprof http://localhost:6060/debug/pprof/heap
     ``
/*
4. **Scalability**:

   * `numWorkers` artırılarak daha fazla concurrency sağlanabilir.
   * Alıcı listesi büyüdüğünde bile sistem stabil kalır.

---

✅ Bu yöntemle **binlerce mail gönderirken performansı yönetebilir ve pprof ile analiz yapabilirsin**.

İstersen sana bunu bir adım daha ileri götürüp **weighted worker pool** (farklı worker’lar farklı sayıda mail gönderiyor) şeklinde optimize edilmiş bir versiyon da gösterebilirim. Bunu da ister misin?
*/