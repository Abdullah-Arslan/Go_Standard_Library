/*
Go’da **`net/http/httptrace`** paketi, HTTP client isteklerinin **detaylı yaşam döngüsünü izlemek** için kullanılır. Yani bir HTTP isteği yapılırken **DNS çözümleme, TCP bağlanma, TLS handshake, ilk byte alınması, redirect ve response tamamlama gibi olayları** takip edebilirsin. Bu paket özellikle **performans ölçümü, debugging ve network analizi** için çok faydalıdır.

Aşağıda paketin tamamını örneklerle açıklıyorum.

---

# 📦 Go `net/http/httptrace` Paketi

## 1️⃣ Temel Amaç

* HTTP client isteklerinin **adım adım takip edilmesi**
* DNS lookup, TCP connect, TLS handshake, first byte, response complete gibi olayları izlemek
* Performans analizleri ve debugging yapmak

---

## 2️⃣ Önemli Tipler ve Fonksiyonlar

### a) `ClientTrace`

* HTTP isteği sırasında **farklı olaylar için callback’ler** sağlar.
* Örneğin:

  * `DNSStart`, `DNSDone` → DNS çözümleme başlangıcı ve bitişi
  * `ConnectStart`, `ConnectDone` → TCP bağlantı başlangıcı ve bitişi
  * `GotConn` → TCP bağlantısı alındığında
  * `GotFirstResponseByte` → İlk byte alındığında

---

### b) `httptrace.WithClientTrace(req *http.Request, trace *ClientTrace)`

* HTTP request’e trace ekler.
* Trace callback’leri, istek boyunca tetiklenir.

---

## 3️⃣ Basit Örnek
*/
``go
package main

import (
    "fmt"
    "net/http"
    "net/http/httptrace"
    "time"
)

func main() {
    req, _ := http.NewRequest("GET", "https://httpbin.org/get", nil)

    trace := &httptrace.ClientTrace{
        DNSStart: func(info httptrace.DNSStartInfo) {
            fmt.Println("DNS çözümleme başladı:", info.Host)
        },
        DNSDone: func(info httptrace.DNSDoneInfo) {
            fmt.Println("DNS çözümleme tamamlandı:", info.Addrs)
        },
        ConnectStart: func(network, addr string) {
            fmt.Println("TCP bağlantısı başlıyor:", network, addr)
        },
        ConnectDone: func(network, addr string, err error) {
            fmt.Println("TCP bağlantısı tamamlandı:", network, addr, "Hata:", err)
        },
        GotFirstResponseByte: func() {
            fmt.Println("İlk byte alındı!")
        },
    }

    req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

    client := &http.Client{}
    start := time.Now()
    resp, _ := client.Do(req)
    defer resp.Body.Close()
    fmt.Println("HTTP Durum Kodu:", resp.StatusCode)
    fmt.Println("İstek Süresi:", time.Since(start))
}
``
/*
---

### 4️⃣ Örnek Çıktı
*/
``
DNS çözümleme başladı: httpbin.org
DNS çözümleme tamamlandı: [34.204.17.123 34.204.17.124]
TCP bağlantısı başlıyor: tcp 34.204.17.123:443
TCP bağlantısı tamamlandı: tcp 34.204.17.123:443 Hata: <nil>
İlk byte alındı!
HTTP Durum Kodu: 200
İstek Süresi: 120ms
``
/*
---

## 5️⃣ Özet

* `httptrace` paketi, HTTP client isteklerinin **detaylı olaylarını izlemeye** yarar.

* Kullanımı:

  1. `ClientTrace` ile event callback’leri tanımlanır.
  2. `httptrace.WithClientTrace(req.Context(), trace)` ile request’e trace eklenir.
  3. `client.Do(req)` ile HTTP isteği yapılır ve callback’ler tetiklenir.

* Faydaları:

  * Performans analizi
  * DNS veya TCP problemlerini tespit etme
  * Debugging

---

İstersen bir sonraki adımda bunu **GET/POST ve redirect takibi ile kombine edilmiş kapsamlı bir httptrace mini proje örneği** ile gösterebilirim.

Bunu yapmak ister misin?
EVET
*/
/*
Harika! 🚀 O zaman sana **GET/POST + Redirect + DNS/TCP/TLS + httptrace mini proje** hazırlayalım. Bu proje, HTTP isteklerinin **tüm yaşam döngüsünü izleyecek** ve performans analizini gösterecek.

---

# 🌐 `httptrace` Kapsamlı Mini Proje

### 1️⃣ Amaç

* HTTP GET ve POST isteklerini izlemek
* Redirect ve form verilerini test etmek
* DNS, TCP, TLS ve first byte olaylarını takip etmek

---

### 2️⃣ Dosya Yapısı
*/
``
httptrace-proje/
│── main.go
``
/*
---

### 3️⃣ `main.go` – Örnek Kod
*/
``go
package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/http/httptrace"
    "net/url"
    "time"
)

func main() {
    // Form verisi
    form := url.Values{}
    form.Set("username", "Ahmet")
    form.Set("password", "1234")
    body := bytes.NewBufferString(form.Encode())

    // Request oluştur
    req, _ := http.NewRequest("POST", "https://httpbin.org/post", body)
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    // Trace callback’leri
    trace := &httptrace.ClientTrace{
        DNSStart: func(info httptrace.DNSStartInfo) {
            fmt.Println("DNS çözümleme başladı:", info.Host)
        },
        DNSDone: func(info httptrace.DNSDoneInfo) {
            fmt.Println("DNS çözümleme tamamlandı:", info.Addrs)
        },
        ConnectStart: func(network, addr string) {
            fmt.Println("TCP bağlantısı başlıyor:", network, addr)
        },
        ConnectDone: func(network, addr string, err error) {
            fmt.Println("TCP bağlantısı tamamlandı:", network, addr, "Hata:", err)
        },
        TLSHandshakeStart: func() {
            fmt.Println("TLS handshake başladı")
        },
        TLSHandshakeDone: func(state httptrace.TLSConnectionState, err error) {
            fmt.Println("TLS handshake tamamlandı, Hata:", err)
        },
        GotFirstResponseByte: func() {
            fmt.Println("İlk byte alındı!")
        },
        GotConn: func(info httptrace.GotConnInfo) {
            fmt.Println("TCP bağlantısı kullanıldı, reuse:", info.Reused)
        },
        WroteRequest: func(info httptrace.WroteRequestInfo) {
            fmt.Println("Request yazıldı, Hata:", info.Err)
        },
    }

    // Trace request’e ekle
    req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

    client := &http.Client{
        CheckRedirect: func(req *http.Request, via []*http.Request) error {
            fmt.Println("Redirect algılandı:", req.URL)
            return nil // redirect devam edebilir
        },
    }

    start := time.Now()
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Hata:", err)
        return
    }
    defer resp.Body.Close()

    bodyResp, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("HTTP Durum Kodu:", resp.StatusCode)
    fmt.Println("Response Body:", string(bodyResp))
    fmt.Println("Toplam süre:", time.Since(start))
}
``
/*
---

### 4️⃣ Açıklama

1. **Form verisi** POST isteği ile gönderiliyor.
2. `httptrace.ClientTrace` ile DNS, TCP, TLS handshake, ilk byte ve redirect gibi olaylar izleniyor.
3. `CheckRedirect` callback’i ile redirect’ler loglanıyor.
4. `client.Do(req)` ile istek gönderiliyor ve trace callback’leri tetikleniyor.
5. Yanıt ve toplam süre ekrana yazdırılıyor.

---

### 5️⃣ Örnek Çıktı
*/
``
DNS çözümleme başladı: httpbin.org
DNS çözümleme tamamlandı: [34.204.17.123 34.204.17.124]
TCP bağlantısı başlıyor: tcp 34.204.17.123:443
TCP bağlantısı tamamlandı: tcp 34.204.17.123:443 Hata: <nil>
TLS handshake başladı
TLS handshake tamamlandı, Hata: <nil>
TCP bağlantısı kullanıldı, reuse: false
Request yazıldı, Hata: <nil>
İlk byte alındı!
HTTP Durum Kodu: 200
Response Body: { ... JSON ... }
Toplam süre: 220ms
``
/*
---

Bu proje ile artık:

* HTTP GET/POST ve form verilerini gönderebilirsin
* DNS, TCP, TLS handshake ve redirect’leri detaylı izleyebilirsin
* İstek performansını ölçebilir ve debug yapabilirsin ✅

---

İstersen bir sonraki adımda **bunu PDF veya Excel tablosu** hâline getirip, tüm callback’leri ve örnek çıktılarıyla görselleştirebilirim.

Bunu yapayım mı?
*/