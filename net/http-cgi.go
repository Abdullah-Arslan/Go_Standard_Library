/*
Go’da **`net/http/cgi` paketi**, **CGI (Common Gateway Interface)** protokolünü kullanarak web uygulamaları çalıştırmanı sağlar.

Normalde **`net/http`** paketi ile Go doğrudan HTTP sunucusu açar.
Ama **`http/cgi`** ile Go programı, **bir web sunucusunun (Apache, Nginx, IIS vs.) CGI modülü üzerinden çağrılabilir hale gelir.**

---

# 📦 `net/http/cgi` Paketi – Kapsamlı Açıklama

## 🔹 Temel Amaç

* CGI ortam değişkenlerini (`REQUEST_METHOD`, `QUERY_STRING`, `CONTENT_TYPE`, vs.) okuyarak **HTTP isteklerini `http.Request` nesnesine dönüştürür**.
* Programın çıktısını **HTTP cevabı olarak web sunucusuna geri gönderir**.

Bu sayede, Go ile yazdığın program **CGI script** gibi çalışır.

---

## 🔹 Önemli Tipler ve Fonksiyonlar

### 1. `cgi.Handler`

* Bir CGI çalıştırıcısıdır.
* `net/http.Handler` arayüzünü uygular.
*/
``go
type Handler struct {
    Path string   // Çalıştırılacak CGI binary yolu
    Dir  string   // Çalıştırma klasörü
    Env  []string // Ekstra environment değişkenleri
    Args []string // Ekstra argümanlar
    InheritEnv []string // Ortamdan aktarılacak env değişkenleri
}
``
/*
---

### 2. `(*Handler) ServeHTTP`

* `http.Handler` gibi çalışır, HTTP isteğini alır, CGI programını çalıştırır.

**Örnek:**
*/
``go
package main

import (
    "net/http"
    "net/http/cgi"
)

func main() {
    handler := &cgi.Handler{
        Path: "/usr/lib/cgi-bin/test.cgi", // çalıştırılacak script
        Dir:  "/usr/lib/cgi-bin",          // çalışma dizini
    }

    http.Handle("/cgi-bin/", handler)
    http.ListenAndServe(":8080", nil)
}
``
/*
Burada:

* `/cgi-bin/` isteği geldiğinde Go, `test.cgi` dosyasını çalıştırır.
* CGI çıktısı HTTP response olarak gönderilir.

---

### 3. `cgi.Request()`

* Ortam değişkenlerinden (`os.Environ`) ve stdin’den gelen veriden bir `*http.Request` üretir.
* CGI uyumlu çalışmak için kullanılır.

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
👉 Bu kod, CGI ortamında çalıştırıldığında isteği doğru şekilde alır.

---

### 4. `cgi.Serve(handler http.Handler)`

* Kolaylık fonksiyonu:

  * `cgi.Request()` ile `*http.Request` oluşturur.
  * Belirtilen `handler`’a iletir.
  * Çıktıyı CGI uyumlu şekilde stdout’a yazar.

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
Bu program CGI olarak çalıştırıldığında, **web sunucusundan gelen istek** `handler` fonksiyonuna aktarılır ve yanıt stdout üzerinden sunucuya döner.

---

## 🔹 CGI Çalışma Mantığı

1. Web sunucusu (ör. Apache) bir CGI programı çağırır.
2. CGI ortam değişkenlerini (`REQUEST_METHOD`, `QUERY_STRING`, vb.) ayarlar.
3. HTTP isteğinin body’sini stdin üzerinden CGI programına yollar.
4. CGI programı (Go kodun) `cgi.Request()` ile isteği okur.
5. `http.ResponseWriter` ile yazılan her şey stdout’a basılır.
6. Web sunucusu stdout’u alır ve istemciye HTTP cevabı olarak yollar.

---

## 🔹 Gerçekçi Örnek – Form CGI
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
/*
📌 Bu CGI programı:

* `http://example.com/cgi-bin/hello?name=Ahmet` isteğinde →
*/

``
Merhaba Ahmet!
``
/*
---

## 🔹 Özet

* `net/http/cgi` paketi, Go programlarını **CGI script** gibi çalıştırmaya yarar.
* Ana bileşenler:

  * **`cgi.Handler`** → CGI script çağırır (proxy gibi).
  * **`cgi.Request`** → CGI ortamından `http.Request` oluşturur.
  * **`cgi.Serve`** → HTTP handler’ı CGI uyumlu şekilde çalıştırır.
* Kullanım senaryosu: **Apache/Nginx altında CGI uyumlu Go uygulamaları çalıştırmak**.

---

👉 İstersen ben sana bunun da **PDF + Excel tablo (fonksiyonlar + örnekler)** formatını hazırlayabilirim.
Böylece `http/cgi` paketini de offline kaynak olarak saklayabilirsin. İstiyor musun?
EVET
*/