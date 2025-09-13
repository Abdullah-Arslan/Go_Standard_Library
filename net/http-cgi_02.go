/*
Go’da **`net/http/cgi` paketi**, Go programlarının **CGI (Common Gateway Interface)** protokolünü kullanarak çalışmasını sağlar. Yani Go uygulaman, bir web sunucusu üzerinden CGI script gibi çalışabilir. Bunu daha iyi anlamak için detaylıca açıklayalım ve örnekler verelim.

---

# 📦 Go `net/http/cgi` Paketi

## 1. Temel Amaç

* CGI protokolü üzerinden gelen HTTP isteklerini işlemek.
* CGI ortam değişkenlerini (`REQUEST_METHOD`, `QUERY_STRING`, `CONTENT_TYPE`, vb.) okuyarak bir `http.Request` nesnesi oluşturmak.
* Programın çıktısını HTTP cevabı olarak web sunucusuna geri göndermek.

CGI kullanımı genellikle **Apache, Nginx veya IIS** gibi web sunucularında uygulanır.

---

## 2. Önemli Tipler ve Fonksiyonlar

### a) `cgi.Handler`

* CGI scriptlerini çalıştırmak için kullanılan `http.Handler`.
* Özellikler:
*/
  ``go
  type Handler struct {
      Path       string   // Çalıştırılacak CGI programı/script
      Dir        string   // Çalışma dizini
      Env        []string // Ekstra environment değişkenleri
      Args       []string // Ekstra argümanlar
      InheritEnv []string // Ortamdan aktarılacak env değişkenleri
  }
  ``

//**Örnek:**

``go
package main

import (
    "net/http"
    "net/http/cgi"
)

func main() {
    handler := &cgi.Handler{
        Path: "/usr/lib/cgi-bin/test.cgi", 
        Dir:  "/usr/lib/cgi-bin",
    }
    http.Handle("/cgi-bin/", handler)
    http.ListenAndServe(":8080", nil)
}
``
/*
---

### b) `(*Handler).ServeHTTP`

* `http.Handler` arayüzü gibi çalışır ve CGI scripti çağırır.

**Örnek:**
*/
``go
handler.ServeHTTP(w, r)
``
/*
---

### c) `cgi.Request()`

* CGI ortam değişkenlerinden `*http.Request` oluşturur.

**Örnek:**
*/
``go
package main

import (
    "fmt"
    "net/http/cgi"
)

func main() {
    r, err := cgi.Request()
    if err != nil {
        panic(err)
    }
    fmt.Println("Method:", r.Method)
    fmt.Println("Path:", r.URL.Path)
}
``
/*
---

### d) `cgi.Serve(handler http.Handler)`

* Verilen handler’ı CGI uyumlu çalıştırır.
* HTTP isteğini handler’a iletir ve çıktıyı CGI üzerinden web sunucusuna yollar.

**Örnek:**
*/
``go
package main

import (
    "fmt"
    "net/http"
    "net/http/cgi"
)

func main() {
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Merhaba CGI Dünya!")
    })
    cgi.Serve(handler)
}
``
/*
---

### e) Form CGI Örneği

* CGI üzerinden gelen form verilerini okuma:
*/
``go
package main

import (
    "fmt"
    "net/http"
    "net/http/cgi"
)

func main() {
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()
        name := r.FormValue("name")
        if name == "" {
            fmt.Fprintln(w, "Merhaba Ziyaretçi!")
        } else {
            fmt.Fprintf(w, "Merhaba %s!\n", name)
        }
    })
    cgi.Serve(handler)
}
``

//CGI ile çalıştırıldığında:

``
http://example.com/cgi-bin/hello?name=Ahmet
``

//Sonuç:

``
Merhaba Ahmet!
``
/*
---

## 3. Özet

* `net/http/cgi` paketi Go uygulamalarını CGI script gibi çalıştırmaya yarar.
* Önemli bileşenler:

  * **`cgi.Handler`** → CGI script çağırır.
  * **`(*Handler).ServeHTTP`** → Handler’ı CGI uyumlu çalıştırır.
  * **`cgi.Request()`** → CGI ortamından `http.Request` oluşturur.
  * **`cgi.Serve(handler)`** → Handler’ı CGI uyumlu şekilde çalıştırır.
* Kullanım: Apache/Nginx altında Go programlarını CGI olarak çalıştırmak.

---

İstersen bir sonraki adımda sana net/http/cgi paketinin küçük bir mini proje örneği de gösterebilirim, böylece canlı olarak nasıl çalıştığını görebilirsin.

Bunu yapayım mı?
EVET
*/
/*
Harika! 🚀 O zaman sana **`net/http/cgi` paketini kullanan mini bir proje** hazırlayalım. Bu proje, CGI üzerinden form verilerini okuyacak ve cevap döndürecek.

---

# 🌐 CGI Mini Proje – Örnek

### 1️⃣ Dosya Yapısı
*/
``
cgi-mini-proje/
│── hello.go          # Go CGI programı
│── public/           # Statik dosyalar (opsiyonel)
``
/*
---

### 2️⃣ `hello.go` – CGI Programı
*/
``go
package main

import (
    "fmt"
    "net/http"
    "net/http/cgi"
)

func main() {
    // CGI handler oluştur
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()
        name := r.FormValue("name")
        if name == "" {
            fmt.Fprintln(w, "Merhaba Ziyaretçi!")
        } else {
            fmt.Fprintf(w, "Merhaba %s!\n", name)
        }
    })

    // CGI üzerinden çalıştır
    cgi.Serve(handler)
}
``
/*
---

### 3️⃣ Çalıştırma

1. Go programını derle:
*/

``bash
go build -o hello.cgi hello.go
``

//2. CGI script’in çalıştırılabilir olduğundan emin ol:


``bash
chmod +x hello.cgi
``
/*
3. Web sunucuna (`/usr/lib/cgi-bin/` gibi bir klasöre) kopyala.
4. Tarayıcıdan test et:
*/

``
http://localhost/cgi-bin/hello.cgi?name=Ahmet
``
/*
Çıktı:
*/
``
Merhaba Ahmet!
``
/*
---

### 4️⃣ Açıklama

* `cgi.Serve(handler)` ile HTTP isteği CGI üzerinden `handler` fonksiyonuna iletiliyor.
* `r.ParseForm()` ile GET/POST parametreleri okunabiliyor.
* CGI script’in stdout’una yazılan her şey, web sunucusu tarafından HTTP cevabı olarak kullanıcıya gönderiliyor.

---

### 5️⃣ Opsiyonel: Form ile Deneme

Basit bir HTML form:
*/
``html
<form action="/cgi-bin/hello.cgi" method="GET">
    Adınız: <input name="name">
    <input type="submit">
</form>
``
/*
* Form submit edildiğinde CGI script çalışır ve isim ile birlikte mesaj döner.

---

Bu mini proje sayesinde:

* `net/http/cgi` kullanımı
* GET/POST verilerini okuma
* CGI script üzerinden HTTP cevabı gönderme

konularını uygulamalı olarak görebilirsin.

---

İstersen bir sonraki adımda sana net/http/cgi ve net/http farklarını örneklerle gösterebilirim, böylece CGI kullanımının ne zaman gerekli olduğunu net olarak anlayabilirsin.

Bunu yapmak ister misin?
EVET
*/
/*
Harika! 🚀 O zaman sana **`net/http` ve `net/http/cgi` paketlerinin farklarını** örneklerle açıklayayım. Böylece hangi durumda hangi paketi kullanman gerektiğini görebileceksin.

---

# 🌐 `net/http` vs `net/http/cgi`
*/

| Özellik / Kriter | `net/http`                                                                                                                                | `net/http/cgi`                                                                                                                        |
| ---------------- | ----------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| Amaç             | Go programının kendi HTTP sunucusunu açması                                                                                               | Go programını CGI script olarak çalıştırmak                                                                                           |
| Sunucu           | Go doğrudan TCP portu dinler (`ListenAndServe`)                                                                                           | Web sunucusu CGI modülü üzerinden çağırır                                                                                             |
| HTTP İsteği      | `http.Request` ile normal şekilde alınır                                                                                                  | CGI ortam değişkenlerinden `http.Request` oluşturulur                                                                                 |
| Çıktı            | `ResponseWriter` direkt TCP üzerinden istemciye gider                                                                                     | `ResponseWriter` stdout’a yazar, web sunucusu bunu HTTP cevabı olarak yollar                                                          |
| Kullanım Alanı   | Standalone Go web uygulamaları                                                                                                            | Apache/Nginx/IIS üzerinde CGI uyumlu Go uygulamaları                                                                                  |
| Performans       | Daha hızlı (direkt TCP, thread/process yok)                                                                                               | CGI olarak çalıştığı için her istek ayrı process oluşturabilir, daha yavaş                                                            |
| Kod Örneği       | `go\nhttp.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){fmt.Fprintln(w, "Merhaba")})\nhttp.ListenAndServe(":8080", nil)\n` | `go\nhandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){fmt.Fprintln(w, "Merhaba CGI")})\ncgi.Serve(handler)\n` |

/*
---

## Örnek Senaryolar

### 1️⃣ `net/http` Örneği – Standalone Sunucu
*/
``go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Merhaba Standalone HTTP!")
    })
    http.ListenAndServe(":8080", nil)
}
``
/*
* Tarayıcıdan `http://localhost:8080/` adresine gittiğinde çalışır.
* Go kendi sunucusunu açar, başka web sunucusuna gerek yok.

---

### 2️⃣ `net/http/cgi` Örneği – CGI Script
*/
``go
package main

import (
    "fmt"
    "net/http"
    "net/http/cgi"
)

func main() {
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Merhaba CGI HTTP!")
    })
    cgi.Serve(handler)
}
``
/*
* CGI script olarak web sunucusuna konulur (`/usr/lib/cgi-bin/hello.cgi` gibi).
* Tarayıcıdan `http://localhost/cgi-bin/hello.cgi` adresiyle çalışır.
* Web sunucusu CGI modülü isteği Go programına iletir ve stdout’u HTTP cevabı olarak döner.

---

## Özet

* **`net/http`** → Standalone web uygulamaları için ideal. TCP portunu Go kendisi açar.
* **`net/http/cgi`** → Mevcut web sunucularında CGI script olarak çalıştırmak için kullanılır.
* Performans açısından `net/http` genellikle daha hızlıdır, `cgi` daha çok uyumluluk içindir.
* Kod yazma mantığı benzerdir, fark sadece **sunucu/CGI ortam yönetimi**ndedir.

---

İstersen ben bunu da **PDF veya Excel tablosu** olarak hazırlayıp, örnekleri ve farkları tek dosyada gösterebilirim.

Bunu yapayım mı?
*/

