**Go (Golang) `net/http` paketi**ni **tam kapsamlı** şekilde anlatacağım. Bu paket Go’nun en önemli paketlerinden biridir çünkü HTTP tabanlı **web sunucusu** ve **istemci** yazmak için temel araçları sağlar.

---

# 📌 `net/http` Paketi Nedir?

`net/http` paketi, Go dilinde HTTP protokolü üzerinde **sunucu** (server) ve **istemci** (client) uygulamaları geliştirmek için kullanılır.

* **Sunucu tarafında**: HTTP request kabul etme, handler (işleyici) fonksiyonları yazma, route (yönlendirme) yapma.
* **İstemci tarafında**: HTTP request gönderme (GET, POST, PUT, DELETE vb.), response (cevap) alma.
* **Ekstra**: Cookie, Header, URL parametreleri, Form işlemleri, Proxy, Timeout gibi konular da desteklenir.

---

# 📌 Temel Yapılar

`http` paketinin en çok kullanılan yapıları:

* `http.ListenAndServe` → Sunucu başlatır.
* `http.HandleFunc` → Bir route’a handler (işleyici fonksiyon) bağlar.
* `http.Request` → İstemciden gelen HTTP isteğini temsil eder.
* `http.ResponseWriter` → Sunucu cevabını yazmak için kullanılır.
* `http.Client` → HTTP istekleri göndermek için istemci.
* `http.ServeMux` → Bir router (çoklu handler yöneticisi).


**`net/http` paketinin tamamını (fonksiyonlar, tipler, arayüzler, sabitler) tek tek açıklayayım ve her biri için örnekler vereyim.**

Bu paket çok geniş, ama ben sana **tüm önemli bileşenlerini** sistematik bir şekilde anlatacağım.

---

# 📦 Go `net/http` Paketi – Kapsamlı Açıklama

## 🔹 1. Sunucu Başlatma

### `http.ListenAndServe`

* Belirtilen adreste (host\:port) HTTP sunucusu başlatır.
* İkinci parametre `Handler'dır (nil → `DefaultServeMux` kullanılır).

**Örnek:**

```go
http.ListenAndServe(":8080", nil)
```

---

### `http.ListenAndServeTLS`

* HTTPS sunucusu başlatır.
* Sertifika (`certFile`) ve özel anahtar (`keyFile`) ister.

**Örnek:**

```go
http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
```

---

## 🔹 2. Handler’lar

### `http.HandleFunc`

* Bir route’a fonksiyon bağlar.

**Örnek:**

```go
http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Merhaba")
})
```

---

### `http.Handle`

* Bir `Handler` objesini route’a bağlar.

**Örnek:**

```go
fs := http.FileServer(http.Dir("./public"))
http.Handle("/static/", fs)
```

---

### `http.Handler` Arayüzü

* Tüm handler’lar şu arayüzü uygular:

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

**Örnek:**

```go
type myHandler struct{}

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Ben özel handler’ım")
}

http.ListenAndServe(":8080", myHandler{})
```

---

## 🔹 3. Router (`ServeMux`)

### `http.ServeMux`

* Route’ları yöneten router’dır.

**Örnek:**

```go
mux := http.NewServeMux()
mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Ana sayfa")
})
mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hakkında")
})
http.ListenAndServe(":8080", mux)
```

---

## 🔹 4. İstek (`http.Request`)

### Yapı

`http.Request` gelen isteği temsil eder.
Önemli alanlar:

* `Method` → "GET", "POST"
* `URL` → istek URL bilgisi
* `Header` → HTTP başlıkları
* `Body` → istek içeriği

**Örnek:**

```go
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Metod:", r.Method)
    fmt.Println("URL:", r.URL.Path)
    fmt.Println("Header:", r.Header.Get("User-Agent"))
}
```

---

## 🔹 5. Cevap (`http.ResponseWriter`)

* Sunucudan istemciye veri göndermek için kullanılır.

**Örnek:**

```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK) // 200
    w.Header().Set("Content-Type", "text/plain")
    fmt.Fprintln(w, "Merhaba Dünya")
}
``

---

## 🔹 6. Yardımcı Fonksiyonlar

### `http.Error`

* Hata mesajı ve status kod döner.
*/
```go
http.Error(w, "Bir hata oluştu", http.StatusInternalServerError)
``
/*
### `http.Redirect`

* İstemciyi başka sayfaya yönlendirir.
*/
```go
http.Redirect(w, r, "/login", http.StatusFound)
``
/*
### `http.NotFound`

* 404 döndürür.
*/
```go
http.NotFound(w, r)
``
/*
---

## 🔹 7. Dosya Servisi

### `http.FileServer`

* Bir klasördeki dosyaları sunar.
*/
```go
fs := http.FileServer(http.Dir("./public"))
http.Handle("/", fs)
``
/*
### `http.ServeFile`

* Tek bir dosya gönderir.
*/
`go
http.ServeFile(w, r, "index.html")
``
/*
---

## 🔹 8. Cookie

### `http.Cookie` Yapısı
*/
`go
cookie := &http.Cookie{
    Name:  "kullanici",
    Value: "Ahmet",
}
http.SetCookie(w, cookie)
`

//### Okuma:

`go
cookie, err := r.Cookie("kullanici")
if err == nil {
    fmt.Println(cookie.Value)
}
``
/*
---

## 🔹 9. Client (İstemci)

### `http.Get`
*/
`go
resp, _ := http.Get("https://httpbin.org/get")
defer resp.Body.Close()
``

//### `http.Post`

``go
resp, _ := http.Post("https://httpbin.org/post", "application/json", strings.NewReader(`{"a":1}`))
``
/*
### `http.Client`

* Özel ayarlar (timeout, proxy vb.)
*/
```go
client := &http.Client{Timeout: 5 * time.Second}
resp, _ := client.Get("https://httpbin.org/get")

/*
### `http.NewRequest`
*/
go
req, _ := http.NewRequest("GET", "https://httpbin.org/get", nil)
req.Header.Set("User-Agent", "GoClient")
resp, _ := http.DefaultClient.Do(req)

/*
---

## 🔹 10. `http.Server`

* Daha fazla kontrol için kullanılır (timeout, TLS, custom handler).
*/
go
srv := &http.Server{
    Addr:         ":8080",
    Handler:      http.DefaultServeMux,
    ReadTimeout:  5 * time.Second,
    WriteTimeout: 10 * time.Second,
}
srv.ListenAndServe()

/*
---

## 🔹 11. Sabitler

* `http.StatusOK` (200)
* `http.StatusNotFound` (404)
* `http.StatusInternalServerError` (500)
* … ve tüm HTTP status kodları.

**Örnek:**
*/

go
w.WriteHeader(http.StatusCreated) // 201

/*
---

# 🎯 Özet

`net/http` paketi ile:

* Sunucu: `ListenAndServe`, `HandleFunc`, `ServeMux`
* İstek/Cevap: `Request`, `ResponseWriter`
* Dosya sunma: `FileServer`, `ServeFile`
* İstemci: `Get`, `Post`, `Client`, `NewRequest`
* Yardımcılar: `Error`, `Redirect`, `NotFound`
* Cookie işlemleri
* Timeout ve özel `http.Server` ayarları

hepsi yapılabiliyor ✅

---

👉 İstersen bu **tam anlatımı (örneklerle birlikte)** de **PDF + Excel tablo** halinde hazırlayıp sana verebilirim. Böylece tüm fonksiyonları offline kaynak olarak saklayabilirsin. İstiyor musun?

---------------------------------------------------------------------------------------------
BAŞKA BİR AÇIKLAMA
*/
/*
**Go (Golang) `net/http` paketi**ni **tam kapsamlı** şekilde anlatacağım. Bu paket Go’nun en önemli paketlerinden biridir çünkü HTTP tabanlı **web sunucusu** ve **istemci** yazmak için temel araçları sağlar.

---

# 📌 `net/http` Paketi Nedir?

`net/http` paketi, Go dilinde HTTP protokolü üzerinde **sunucu** (server) ve **istemci** (client) uygulamaları geliştirmek için kullanılır.

* **Sunucu tarafında**: HTTP request kabul etme, handler (işleyici) fonksiyonları yazma, route (yönlendirme) yapma.
* **İstemci tarafında**: HTTP request gönderme (GET, POST, PUT, DELETE vb.), response (cevap) alma.
* **Ekstra**: Cookie, Header, URL parametreleri, Form işlemleri, Proxy, Timeout gibi konular da desteklenir.

---

# 📌 Temel Yapılar

`http` paketinin en çok kullanılan yapıları:

* `http.ListenAndServe` → Sunucu başlatır.
* `http.HandleFunc` → Bir route’a handler (işleyici fonksiyon) bağlar.
* `http.Request` → İstemciden gelen HTTP isteğini temsil eder.
* `http.ResponseWriter` → Sunucu cevabını yazmak için kullanılır.
* `http.Client` → HTTP istekleri göndermek için istemci.
* `http.ServeMux` → Bir router (çoklu handler yöneticisi).

---

# 📌 1. Basit HTTP Sunucusu
*/
go
package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Merhaba, Go HTTP Server!")
}

func main() {
	http.HandleFunc("/", helloHandler) // Route tanımı
	fmt.Println("Sunucu 8080 portunda çalışıyor...")
	http.ListenAndServe(":8080", nil) // Sunucuyu başlat
}


//👉 Çalıştır:


go run main.go

/*
Tarayıcıda [http://localhost:8080](http://localhost:8080) açtığında `"Merhaba, Go HTTP Server!" yazar.

---

# 📌 2. Parametreli Route
*/
``go
func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // URL query parametresi ?name=Ahmet
	if name == "" {
		name = "Ziyaretçi"
	}
	fmt.Fprintf(w, "Merhaba %s!", name)
}

func main() {
	http.HandleFunc("/greet", greetHandler)
	http.ListenAndServe(":8080", nil)
}
``
/*
👉 Örnek: `http://localhost:8080/greet?name=Ahmet`
Çıktı: **Merhaba Ahmet!**

---

# 📌 3. Form ile POST İşlemi
*/
``go
func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		name := r.FormValue("name")
		age := r.FormValue("age")
		fmt.Fprintf(w, "Ad: %s, Yaş: %s", name, age)
	} else {
		fmt.Fprintln(w, `<form method="POST">
			Ad: <input name="name">
			Yaş: <input name="age">
			<input type="submit">
		</form>`)
	}
}

func main() {
	http.HandleFunc("/form", formHandler)
	http.ListenAndServe(":8080", nil)
}
``
/*
---

# 📌 4. JSON Response Döndürme
*/
``go
import (
	"encoding/json"
)

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"mesaj": "Merhaba Dünya"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/json", jsonHandler)
	http.ListenAndServe(":8080", nil)
}
``

//👉 `curl http://localhost:8080/json` çıktısı:


``json
{"mesaj":"Merhaba Dünya"}
``
/*
---

# 📌 5. HTTP Client (İstemci)
*/
``go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Status:", resp.Status)
	fmt.Println("Response:", string(body))
}
``
/*
---

# 📌 6. Özel Router (`ServeMux`)
*/
``go
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ana Sayfa")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hakkında")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)

	http.ListenAndServe(":8080", mux)
}
``
/*
---

# 📌 7. Middleware (Ara Katman)
*/
``go
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("İstek:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Merhaba Middleware!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	
	// Middleware ekle
	loggedMux := loggingMiddleware(mux)
	http.ListenAndServe(":8080", loggedMux)
}
``
/*
---

# 📌 8. Cookie Kullanımı
*/
``go
func setCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "kullanici",
		Value: "Ahmet",
	})
	fmt.Fprintln(w, "Cookie ayarlandı")
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("kullanici")
	if err != nil {
		fmt.Fprintln(w, "Cookie bulunamadı")
		return
	}
	fmt.Fprintf(w, "Cookie değeri: %s", cookie.Value)
}

func main() {
	http.HandleFunc("/set", setCookie)
	http.HandleFunc("/get", getCookie)
	http.ListenAndServe(":8080", nil)
}
``
/*
---

# 📌 9. Dosya Sunmak (Static File Server)
*/
``go
func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	fmt.Println("Sunucu http://localhost:8080 üzerinde çalışıyor")
	http.ListenAndServe(":8080", nil)
}
``
/*
👉 `./public` klasöründeki dosyalar tarayıcıya sunulur.

---

# 📌 10. Timeout ve Özel Ayarlar
*/
``go
func main() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: http.DefaultServeMux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Merhaba Timeout")
	})

	srv.ListenAndServe()
}
``
/*
---

# 📌 Özet

`net/http` paketi ile:

* **Sunucu** (handler, router, middleware, cookie, JSON, form)
* **İstemci** (GET, POST, Header, Timeout, Proxy)
* **Static dosya sunumu**
* **Özel ayarlar ve middleware**

hepsi kolayca yapılabilir. 🚀

---

İstersen ben bu `http` paketi için **tam bir tablo** hazırlayıp sana **PDF veya Excel** formatında verebilirim. Böylece tüm fonksiyonları ve örnekleri offline da kullanabilirsin. İstiyor musun?
EVET

Bu dosyalarda artık her fonksiyon/tip için çalışır durumda küçük Go kod parçaları da eklenmiş durumda. 🚀

İstersen bundan sonraki adımda sana mini bir proje (küçük web uygulaması) üzerinden http paketinin tüm özelliklerini uygulamalı olarak gösterebilirim. Bunu ister misin?
EVET

👉 İstersen bir sonraki adımda bu mini projeyi adım adım çalıştırma rehberi (kurulum + test senaryoları + tarayıcıda ne göreceğin) şeklinde de hazırlayabilirim. İstiyor musun?
EVET
*/

/*
Harika 🚀 Şimdi sana bu **mini projeyi adım adım çalıştırma rehberi** hazırlayayım. Böylece hiç takılmadan projeyi kurup test edebilirsin.

---

# 📘 `net/http` Mini Proje Çalıştırma Rehberi

## 1️⃣ Ortam Hazırlığı

* Bilgisayarında **Go** yüklü olmalı (sürüm ≥ 1.18 önerilir).
  Kontrol etmek için:
*/
``bash
go version
``
/*
Eğer yüklü değilse: [https://go.dev/dl/](https://go.dev/dl/) adresinden indir.

---

## 2️⃣ Proje Klasörünü Oluştur
*/

``bash
mkdir go-http-proje
cd go-http-proje
``
/*
---

## 3️⃣ Dosya Yapısı
*/
``
go-http-proje/
│── main.go
│── public/
│    └── index.html
``
/*
📄 `main.go` → Sana verdiğim mini proje kodunu buraya koy.

📄 `public/index.html` → Statik dosya için basit bir HTML:
*/

``html
<!DOCTYPE html>
<html>
<head><title>Static Örnek</title></head>
<body><h1>Merhaba, ben statik dosyayım!</h1></body>
</html>
``
/*
---

## 4️⃣ Projeyi Çalıştır

Terminalde:
*/

``bash
go run main.go
``

//Çıktı:

``
Sunucu http://localhost:8080 üzerinde çalışıyor...
``
/*
---

## 5️⃣ Test Senaryoları

🔹 **Ana Sayfa**
Adres: [http://localhost:8080/](http://localhost:8080/)
Görünen:
*/
``
Hoşgeldin! /form, /json, /setcookie, /getcookie adreslerini dene.
``
/*
---

🔹 **Form İşlemi**
Adres: [http://localhost:8080/form](http://localhost:8080/form)

1. Ad ve yaş gir.
2. POST sonrası:
*/
``
Ad: Ahmet, Yaş: 25
``
/*
---

🔹 **JSON API**
Adres: [http://localhost:8080/json](http://localhost:8080/json)
JSON çıktısı:
*/
``json
{"mesaj":"Merhaba JSON!"}
``
/*
---

🔹 **Cookie Ayarlama**
Adres: [http://localhost:8080/setcookie](http://localhost:8080/setcookie)
Sonuç:
*/
``
Cookie ayarlandı!
``
/*
---

🔹 **Cookie Okuma**
Adres: [http://localhost:8080/getcookie](http://localhost:8080/getcookie)
Sonuç:
*/
``
Cookie değeri: Ahmet
``
/*
---

🔹 **Static Dosya Sunumu**
Adres: [http://localhost:8080/static/](http://localhost:8080/static/)
Görünen:
*/
``
Merhaba, ben statik dosyayım!
``
/*
---

## 6️⃣ Middleware Log Çıktısı

Terminalde her istek loglanır, örneğin:
*/
``
İstek: GET /form
İstek: POST /form
İstek: GET /json
``
/*
---

## ✅ Özet

Bu rehberle:

* `net/http` ile **web server** kurmayı,
* **Form POST işlemi**,
* **JSON API**,
* **Cookie kullanımı**,
* **Statik dosya servisi**,
* **Middleware**

gibi özellikleri uygulamalı gördün. 🚀

---

👉 İstersen bu **çalıştırma rehberini de PDF + Excel** olarak hazırlayayım. Böylece proje kodu, test adımları ve senaryolar tek dosyada elinde olur. İstiyor musun?
*/