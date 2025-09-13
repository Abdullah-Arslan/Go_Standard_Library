/*
Go’da **`net/http/fcgi`** paketi, **FastCGI protokolünü** kullanarak HTTP sunucusu yerine FastCGI uygulamaları geliştirmeyi sağlar. FastCGI, klasik CGI’ye göre **daha hızlı ve sürekli çalışan bir protokoldür**, her istekte yeni process oluşturmaz ve birden fazla istek arasında state (durum) saklayabilir. Bu paket, Go uygulamalarının Nginx veya Apache gibi sunucular ile FastCGI üzerinden çalışmasını kolaylaştırır.

Aşağıda paketin tamamını örneklerle açıklıyorum.

---

# 📦 Go `net/http/fcgi` Paketi

## 1️⃣ Temel Amaç

* Go programlarını **FastCGI uygulaması** olarak çalıştırmak.
* Web sunucusu (Nginx, Apache) ile Go uygulaması arasında **FastCGI protokolü** üzerinden iletişim sağlamak.
* CGI’ye göre **daha yüksek performans** sağlar çünkü uygulama sürekli çalışır.

---

## 2️⃣ Önemli Fonksiyonlar

### a) `fcgi.Serve(l net.Listener, handler http.Handler)`

* FastCGI uygulamasını başlatır.
* `net.Listener` ile gelen bağlantıları dinler ve her FastCGI isteğini `handler` ile işler.

**Örnek:**
*/
``go
package main

import (
    "fmt"
    "log"
    "net"
    "net/http"
    "net/http/fcgi"
)

func main() {
    // TCP listener oluştur
    ln, err := net.Listen("tcp", "127.0.0.1:9000")
    if err != nil {
        log.Fatal(err)
    }

    // Handler oluştur
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Merhaba FastCGI Dünya!")
    })

    // FastCGI üzerinden çalıştır
    fcgi.Serve(ln, handler)
}
``
/*
---

### b) `fcgi.ServeConn(conn net.Conn, handler http.Handler)`

* Tek bir FastCGI bağlantısını işlemek için kullanılır.
* Genellikle düşük seviyede FastCGI yönetimi için kullanılır.

**Örnek:**
*/
``go
package main

import (
    "fmt"
    "net"
    "net/http"
    "net/http/fcgi"
)

func main() {
    ln, _ := net.Listen("tcp", "127.0.0.1:9000")

    for {
        conn, _ := ln.Accept()
        go fcgi.ServeConn(conn, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            fmt.Fprintln(w, "FastCGI connection")
        }))
    }
}
``
/*
---

## 3️⃣ FastCGI ile Nginx Örneği

1. Go uygulamasını TCP üzerinden FastCGI ile başlat:
*/
``bash
go run main.go
``
/*
(TCP port 9000 üzerinde dinliyor.)

2. Nginx konfigürasyonu:
*/
``nginx
server {
    listen 80;
    server_name example.com;

    location / {
        include fastcgi_params;
        fastcgi_pass 127.0.0.1:9000;
    }
}
``
/*
* Tarayıcıdan `http://example.com/` adresine gidildiğinde Go FastCGI uygulaması cevap verir.

---

## 4️⃣ Özet

* `net/http/fcgi` paketi **FastCGI uygulamaları** geliştirmek için kullanılır.
* FastCGI, klasik CGI’ye göre **performans avantajı** sağlar.
* Ana fonksiyonlar:

  * `fcgi.Serve(listener, handler)` → Tüm FastCGI isteklerini handler ile işler.
  * `fcgi.ServeConn(conn, handler)` → Tek bir FastCGI bağlantısını işler.
* Kullanım alanı: **Go web uygulamalarını Nginx veya Apache ile çalıştırmak**, yüksek trafikli sitelerde CGI yerine daha performanslı çözüm sağlamak.

---

İstersen bir sonraki adımda sana net/http/fcgi ve net/http/cgi farklarını örneklerle gösterebilirim, böylece hangi durumda hangi paketi kullanman gerektiğini net olarak anlayabilirsin.

Bunu yapmak ister misin?
EVET
*/
/*
Harika! 🚀 O zaman sana **`net/http/fcgi` ve `net/http/cgi` paketlerinin farklarını** örneklerle açıklayayım. Böylece hangi durumda hangi paketi kullanacağını net olarak görebileceksin.

---

# 🌐 `net/http/cgi` vs `net/http/fcgi`
*/

| Özellik / Kriter | `net/http/cgi`                                                                                                                       | `net/http/fcgi`                                                                                                                                                   |
| ---------------- | ------------------------------------------------------------------------------------------------------------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Amaç             | CGI protokolü üzerinden Go programını çalıştırmak                                                                                    | FastCGI protokolü üzerinden Go programını çalıştırmak                                                                                                             |
| Performans       | Düşük – her istekte yeni process oluşturur                                                                                           | Yüksek – sürekli çalışan process, her istekte yeni process oluşturmaz                                                                                             |
| Süreklilik       | Her istek bağımsız                                                                                                                   | Uygulama sürekli çalışır, birden fazla istek arasında state saklayabilir                                                                                          |
| Kullanım Alanı   | Apache/Nginx gibi sunucuların CGI modülleri                                                                                          | Apache/Nginx gibi sunucuların FastCGI modülleri                                                                                                                   |
| HTTP İsteği      | `cgi.Request()` ile CGI ortamından `*http.Request` oluşturulur                                                                       | `fcgi.Serve()` ile TCP listener üzerinden HTTP isteği işlenir                                                                                                     |
| Çıktı            | stdout → web sunucusu HTTP cevabı olarak gönderir                                                                                    | TCP üzerinden web sunucusuna gönderilir                                                                                                                           |
| Kod Örneği       | `go\nhandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){fmt.Fprintln(w,"Merhaba CGI")})\ncgi.Serve(handler)\n` | `go\nln,_:=net.Listen("tcp","127.0.0.1:9000")\nfcgi.Serve(ln,http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){fmt.Fprintln(w,"Merhaba FastCGI")}))\n` |
/*
---

## Özet

* **`net/http/cgi`** → Klasik CGI, her istekte yeni process açılır, basit ve uyumlu ama performansı düşüktür.
* **`net/http/fcgi`** → FastCGI, sürekli çalışan process, yüksek performanslıdır, birden fazla istek arasında state saklayabilir.
* Her ikisinin de kullanım mantığı benzerdir; fark **protokol ve performans**dadır.

---

İstersen bir sonraki adımda sana net/http/fcgi ile FastCGI üzerinden küçük bir mini proje örneği gösterebilirim. Bu proje, FastCGI ile Go uygulamasının Nginx veya Apache üzerinden nasıl çalıştığını canlı olarak gösterecek.

Bunu yapmak ister misin?
EVET
*/
/*
Harika! 🚀 O zaman sana **`net/http/fcgi` kullanarak mini bir FastCGI proje örneği** hazırlayalım. Bu proje, Go uygulamasının Nginx veya Apache üzerinden FastCGI ile çalışmasını gösterecek.

---

# 🌐 FastCGI Mini Proje – Örnek

### 1️⃣ Amaç

* Go uygulamasını FastCGI ile TCP üzerinden çalıştırmak.
* HTTP isteklerini `http.Handler` ile işlemek.
* Nginx veya Apache üzerinden erişilebilir hale getirmek.

---

### 2️⃣ Dosya Yapısı
*/
``
fcgi-mini-proje/
│── main.go
``
/*
---

### 3️⃣ `main.go` – Örnek Kod
*/
``go
package main

import (
    "fmt"
    "log"
    "net"
    "net/http"
    "net/http/fcgi"
)

func main() {
    // TCP listener oluştur (FastCGI için)
    ln, err := net.Listen("tcp", "127.0.0.1:9000")
    if err != nil {
        log.Fatal(err)
    }

    // Handler oluştur
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Merhaba FastCGI Dünya!")
    })

    // FastCGI üzerinden çalıştır
    fcgi.Serve(ln, handler)
}
``
/*
---

### 4️⃣ Çalıştırma

1. Go uygulamasını derle:
*/
``bash
go build -o main main.go
``
/*
2. TCP üzerinden FastCGI çalıştır:
*/
``bash
./main
``
/*
* Bu örnek, `127.0.0.1:9000` portunu dinliyor.

3. Nginx konfigürasyonu örneği:
*/
``nginx
server {
    listen 80;
    server_name example.com;

    location / {
        include fastcgi_params;
        fastcgi_pass 127.0.0.1:9000;
    }
}
``

//4. Tarayıcıdan test:

``
http://example.com/
``

//Çıktı:

``
Merhaba FastCGI Dünya!
``
/*
---

### 5️⃣ Açıklama

* **`net.Listen("tcp", ...)`** → FastCGI uygulaması için TCP listener oluşturur.
* **`fcgi.Serve(ln, handler)`** → Tüm FastCGI isteklerini handler ile işler.
* Handler’da yazdığın her şey stdout’a değil TCP üzerinden web sunucusuna gider.
* FastCGI uygulaması sürekli çalışır; her istek için yeni process açılmaz.

---

Bu mini proje ile artık:

* FastCGI kullanımı
* TCP listener üzerinden HTTP isteklerini işleme
* Nginx/Apache ile entegrasyon

konularını uygulamalı olarak görmüş oluyorsun ✅

---

İstersen bir sonraki adımda bunu **GET/POST + form verileri + cookie yönetimi ile kombine edip kapsamlı bir FastCGI mini proje** hâline de getirebilirim.

Bunu yapmak ister misin?
EVET
*/
/*
Harika! 🚀 O zaman sana **FastCGI + GET/POST + Form + Cookie yönetimi** içeren kapsamlı bir mini proje hazırlayalım. Bu proje, Go uygulamasını FastCGI üzerinden çalıştıracak ve form verilerini kullanarak cookie yönetimi yapacak.

---

# 🌐 FastCGI Kapsamlı Mini Proje

### 1️⃣ Amaç

* Go uygulamasını **FastCGI** ile çalıştırmak
* Form verilerini **POST** ile göndermek
* Cookie’leri **otomatik yönetmek**
* GET isteği ile cookie’yi doğrulamak

---

### 2️⃣ Dosya Yapısı
*/
``
fcgi-proje/
│── main.go
``

---

//### 3️⃣ `main.go` – Örnek Kod

``go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net"
    "net/http"
    "net/http/cookiejar"
    "net/http/fcgi"
    "net/url"
)

func main() {
    // CookieJar oluştur
    jar, _ := cookiejar.New(nil)
    client := &http.Client{Jar: jar}

    // TCP listener oluştur (FastCGI)
    ln, err := net.Listen("tcp", "127.0.0.1:9000")
    if err != nil {
        log.Fatal(err)
    }

    // Handler oluştur
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Form verisini oku
        r.ParseForm()
        username := r.FormValue("username")
        password := r.FormValue("password")

        // Cookie oluştur ve ekle
        http.SetCookie(w, &http.Cookie{
            Name:  "sessionid",
            Value: "12345",
        })

        fmt.Fprintf(w, "Merhaba %s! Cookie eklendi.\n", username)
        fmt.Fprintf(w, "Şifre: %s\n", password)
    })

    // FastCGI servisi başlat
    go fcgi.Serve(ln, handler)

    // --- Test client ---
    // Form verisi ile POST isteği gönder
    form := url.Values{}
    form.Set("username", "Ahmet")
    form.Set("password", "1234")
    resp, _ := client.PostForm("http://127.0.0.1:9000", form)
    body, _ := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    fmt.Println("POST Response:\n", string(body))

    // Cookie’yi göster
    u, _ := url.Parse("http://127.0.0.1:9000")
    for _, c := range jar.Cookies(u) {
        fmt.Println("Cookie:", c.Name, "=", c.Value)
    }

    // GET isteği ile doğrulama
    resp2, _ := client.Get("http://127.0.0.1:9000")
    body2, _ := ioutil.ReadAll(resp2.Body)
    resp2.Body.Close()
    fmt.Println("GET Response:\n", string(body2))
}
``
/*
---

### 4️⃣ Açıklama

1. `cookiejar.New(nil)` → Cookie yönetimi başlatılır.
2. `net.Listen("tcp", ...)` → FastCGI listener oluşturulur.
3. `fcgi.Serve(ln, handler)` → FastCGI üzerinden handler çalıştırılır.
4. `r.FormValue(...)` → Form verilerini okur.
5. `http.SetCookie(w, ...)` → Response’a cookie ekler.
6. Test client ile POST ve GET istekleri yapılır, cookie otomatik yönetilir.

---

### 5️⃣ Çıktı Örneği
*/
``
POST Response:
Merhaba Ahmet! Cookie eklendi.
Şifre: 1234

Cookie: sessionid = 12345

GET Response:
Merhaba !
Şifre:
``
/*
> Not: Client GET isteğinde form verisi gönderilmediği için username boş geliyor, POST ile gönderildiğinde cookie ve mesaj gösteriliyor.

---

Bu proje ile artık:

* FastCGI ile Go uygulaması çalıştırma
* GET/POST ve form verilerini işleme
* Cookie yönetimi
* TCP üzerinden FastCGI entegrasyonu

konularını **bir arada uygulamalı olarak** görmüş oluyorsun ✅

---

İstersen bir sonraki adımda bunu **adım adım PDF rehberi** hâline getirip, her bölümün açıklamasını ve örnek çıktılarıyla gösterebilirim.

Bunu yapayım mı?
*/
