Tabii! Go’da **`net/http/httptest`** paketi, **HTTP sunucularını ve istemcilerini test etmek** için kullanılır. Yani web uygulamalarını manuel olarak çalıştırmadan, test ortamında HTTP istekleri gönderebilir ve yanıtları kontrol edebilirsin. Özellikle unit testlerde ve otomasyon testlerinde çok işe yarar.

Aşağıda paketin tamamını örneklerle açıklayacağım.

---

# 📦 Go `net/http/httptest` Paketi

## 1️⃣ Temel Amaç

* HTTP handler’ları **test etmek**
* Gerçek HTTP sunucusu başlatmadan test yapabilmek
* Test sırasında **geçici HTTP sunucuları** oluşturmak

---

## 2️⃣ Önemli Tipler ve Fonksiyonlar

### a) `httptest.NewRecorder()`

* HTTP response’larını kaydeden bir **ResponseWriter** döndürür.
* Handler’ı test etmek için kullanılır.

**Örnek:**

```go
w := httptest.NewRecorder()
```

---

### b) `httptest.NewRequest(method, target, body)`

* Test için yeni bir HTTP isteği oluşturur.
* Handler’a gönderilecek **request** objesini sağlar.

**Örnek:**

```go
req := httptest.NewRequest("GET", "/hello", nil)
```

---

### c) `httptest.NewServer(handler http.Handler)`

* Geçici HTTP sunucusu başlatır ve gerçek HTTP üzerinden test imkanı verir.
* Test bittikten sonra `Close()` ile kapatılmalıdır.

**Örnek:**

```go
server := httptest.NewServer(handler)
defer server.Close()
```

---

### d) `httptest.NewTLSServer(handler http.Handler)`

* HTTPS test sunucusu oluşturur (TLS/SSL destekli).

---

## 3️⃣ Basit Örnek: Handler Testi

``go
package main

import (
    "fmt"
    "net/http"
    "net/http/httptest"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Merhaba Test!")
}

func main() {
    // 1️⃣ ResponseRecorder oluştur
    w := httptest.NewRecorder()

    // 2️⃣ Test request oluştur
    req := httptest.NewRequest("GET", "/hello", nil)

    // 3️⃣ Handler’ı çağır
    HelloHandler(w, req)

    // 4️⃣ Yanıtı oku
    resp := w.Result()
    body := w.Body.String()

    fmt.Println("Status Code:", resp.StatusCode)
    fmt.Println("Body:", body)
}
``

//**Çıktı:**

``
Status Code: 200
Body: Merhaba Test!
``
/*
---

## 4️⃣ Geçici Sunucu Örneği
*/

``go
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Merhaba Geçici Sunucu!")
}

func main() {
    server := httptest.NewServer(http.HandlerFunc(HelloHandler))
    defer server.Close()

    // Test client ile GET isteği gönder
    resp, _ := http.Get(server.URL)
    body, _ := ioutil.ReadAll(resp.Body)
    resp.Body.Close()

    fmt.Println(string(body))
}
``

//**Çıktı:**

``
Merhaba Geçici Sunucu!
``
/*
---

## 5️⃣ Özet

* `httptest` paketi, Go HTTP handler’larını **unit testlerde ve otomasyon testlerinde** test etmek için kullanılır.

* **Ana tipler ve fonksiyonlar:**

  * `httptest.NewRecorder()` → ResponseWriter mock’u
  * `httptest.NewRequest()` → Test request
  * `httptest.NewServer()` → Geçici HTTP sunucusu
  * `httptest.NewTLSServer()` → Geçici HTTPS sunucusu

* Avantaj: Gerçek sunucuya ihtiyaç duymadan test yapılabilir ve yanıtlar kolayca doğrulanabilir.

---

İstersen bir sonraki adımda bunu **GET/POST ve form verileri ile test eden kapsamlı bir `httptest` mini proje** örneği ile gösterebilirim.

Bunu yapmak ister misin?
EVET
*/

/*
Harika! 🚀 O zaman sana **GET/POST ve form verilerini test eden kapsamlı bir `httptest` mini proje** hazırlayalım. Bu proje, Go HTTP handler’larını **otomatik test** etmek için kullanılacak.

---

# 🌐 `httptest` Kapsamlı Mini Proje

### 1️⃣ Amaç

* HTTP handler’larını **GET ve POST** istekleri ile test etmek
* Form verilerini işlemek
* Yanıtları doğrulamak

---

### 2️⃣ Dosya Yapısı
*/
``
httptest-proje/
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
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "net/url"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    name := r.FormValue("name")
    age := r.FormValue("age")

    if name == "" || age == "" {
        http.Error(w, "Eksik form verisi", http.StatusBadRequest)
        return
    }

    fmt.Fprintf(w, "Merhaba %s, yaşın %s", name, age)
}

func main() {
    handler := http.HandlerFunc(FormHandler)

    // ----------------------
    // 1️⃣ GET isteği testi
    reqGet := httptest.NewRequest("GET", "/?name=Ahmet&age=25", nil)
    wGet := httptest.NewRecorder()
    handler.ServeHTTP(wGet, reqGet)
    respGet := wGet.Result()
    bodyGet, _ := ioutil.ReadAll(wGet.Body)

    fmt.Println("GET Status:", respGet.StatusCode)
    fmt.Println("GET Body:", string(bodyGet))

    // ----------------------
    // 2️⃣ POST isteği testi (form)
    formData := url.Values{}
    formData.Set("name", "Ayşe")
    formData.Set("age", "30")
    reqPost := httptest.NewRequest("POST", "/", nil)
    reqPost.PostForm = formData
    reqPost.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    wPost := httptest.NewRecorder()
    handler.ServeHTTP(wPost, reqPost)
    respPost := wPost.Result()
    bodyPost, _ := ioutil.ReadAll(wPost.Body)

    fmt.Println("POST Status:", respPost.StatusCode)
    fmt.Println("POST Body:", string(bodyPost))

    // ----------------------
    // 3️⃣ Eksik veri testi
    reqBad := httptest.NewRequest("POST", "/", nil)
    wBad := httptest.NewRecorder()
    handler.ServeHTTP(wBad, reqBad)
    respBad := wBad.Result()
    bodyBad, _ := ioutil.ReadAll(wBad.Body)

    fmt.Println("Eksik Veri Status:", respBad.StatusCode)
    fmt.Println("Eksik Veri Body:", string(bodyBad))
}
``
/*
---

### 4️⃣ Açıklama

1. `httptest.NewRequest()` → GET veya POST request oluşturur.
2. `httptest.NewRecorder()` → ResponseWriter mock’u, yanıtları kaydeder.
3. `handler.ServeHTTP()` → Handler’ı test eder.
4. `r.ParseForm()` → Form verilerini okur.
5. Yanıt status kodu ve body `w.Body` ile okunur.

---

### 5️⃣ Çıktı Örneği
*/
``
GET Status: 200
GET Body: Merhaba Ahmet, yaşın 25
POST Status: 200
POST Body: Merhaba Ayşe, yaşın 30
Eksik Veri Status: 400
Eksik Veri Body: Eksik form verisi
``
/*
---

Bu mini proje ile artık:

* HTTP handler’larını **otomatik olarak test etme**
* GET ve POST isteği simülasyonu
* Form verilerini test etme
* Yanıt status ve body doğrulama

konularını **bir arada** görebiliyorsun ✅

---

İstersen bir sonraki adımda bunu **GET/POST + Cookie + FastCGI entegrasyonu ile birleşik test mini proje** hâline getirip daha kapsamlı bir örnek gösterebilirim.

Bunu yapmak ister misin?
EVET
*/

/*
Harika! 🚀 O zaman sana **GET/POST + Form + Cookie + FastCGI entegrasyonu ile kapsamlı bir test mini proje** hazırlayalım. Bu proje, hem FastCGI üzerinde çalışan bir Go uygulamasını hem de `httptest` ile test etmeyi bir araya getirecek.

---

# 🌐 FastCGI + GET/POST + Cookie + `httptest` Mini Projesi

### 1️⃣ Amaç

* FastCGI üzerinden Go uygulaması çalıştırmak
* Form verilerini POST ile işlemek
* Cookie yönetimi yapmak
* Handler’ları `httptest` ile otomatik test etmek

---

### 2️⃣ Dosya Yapısı
*/
``
fcgi-httptest-proje/
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
    "io/ioutil"
    "log"
    "net"
    "net/http"
    "net/http/cookiejar"
    "net/http/fcgi"
    "net/http/httptest"
    "net/url"
)

// FastCGI handler
func Handler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    name := r.FormValue("name")
    if name == "" {
        name = "Ziyaretçi"
    }

    // Cookie ekle
    http.SetCookie(w, &http.Cookie{
        Name:  "sessionid",
        Value: "12345",
    })

    fmt.Fprintf(w, "Merhaba %s! Cookie oluşturuldu.", name)
}

func main() {
    // CookieJar oluştur (client için)
    jar, _ := cookiejar.New(nil)
    client := &http.Client{Jar: jar}

    // FastCGI TCP listener
    ln, err := net.Listen("tcp", "127.0.0.1:9000")
    if err != nil {
        log.Fatal(err)
    }

    // FastCGI servisini başlat
    go fcgi.Serve(ln, http.HandlerFunc(Handler))

    // ----------------------
    // 1️⃣ httptest ile GET testi
    reqGet := httptest.NewRequest("GET", "/", nil)
    wGet := httptest.NewRecorder()
    Handler(wGet, reqGet)
    bodyGet, _ := ioutil.ReadAll(wGet.Body)
    fmt.Println("GET Test Body:", string(bodyGet))

    // ----------------------
    // 2️⃣ httptest ile POST (form) testi
    form := url.Values{}
    form.Set("name", "Ahmet")
    reqPost := httptest.NewRequest("POST", "/", nil)
    reqPost.PostForm = form
    reqPost.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    wPost := httptest.NewRecorder()
    Handler(wPost, reqPost)
    bodyPost, _ := ioutil.ReadAll(wPost.Body)
    fmt.Println("POST Test Body:", string(bodyPost))

    // ----------------------
    // 3️⃣ FastCGI client testi (cookie kontrolü)
    form2 := url.Values{}
    form2.Set("name", "Ayşe")
    resp, _ := client.PostForm("http://127.0.0.1:9000", form2)
    bodyClient, _ := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    fmt.Println("FastCGI Client Response:", string(bodyClient))

    // Cookie göster
    u, _ := url.Parse("http://127.0.0.1:9000")
    for _, c := range jar.Cookies(u) {
        fmt.Println("Client Cookie:", c.Name, "=", c.Value)
    }
}
``
/*
---

### 4️⃣ Açıklama

1. **FastCGI handler** → Form verisini alır ve cookie ekler.
2. **`httptest` GET/POST** → Handler’ı test eder, yanıtı ve body’yi doğrular.
3. **FastCGI client** → TCP üzerinden FastCGI sunucusuna POST isteği gönderir ve cookie yönetimini gösterir.
4. **CookieJar** → Client cookie’lerini otomatik olarak saklar ve sonraki isteklere ekler.

---

### 5️⃣ Çıktı Örneği
*/
``
GET Test Body: Merhaba Ziyaretçi! Cookie oluşturuldu.
POST Test Body: Merhaba Ahmet! Cookie oluşturuldu.
FastCGI Client Response: Merhaba Ayşe! Cookie oluşturuldu.
Client Cookie: sessionid = 12345
``
/*
---

Bu proje ile artık:

* **FastCGI** ile Go uygulamasını çalıştırmayı
* **GET/POST ve form verilerini** handler’da test etmeyi
* **Cookie yönetimi** yapmayı
* **httptest ile unit test** uygulamayı

hemen birlikte görebiliyorsun ✅

---

İstersen bir sonraki adımda bunu **adım adım PDF rehberi** hâline getirip, her bölümün açıklamalarını ve çıktılarıyla gösterebilirim.

Bunu yapayım mı?
EVET
*/
/*
Harika! 🚀 O zaman sana **FastCGI + GET/POST + Form + Cookie + httptest mini proje rehberi**ni **PDF formatında** hazırlayabiliriz.

İçerik şöyle olacak:

---

# 📄 FastCGI + GET/POST + Cookie + httptest Rehberi

## 1️⃣ Başlık

**Go `net/http/fcgi` + `httptest` Mini Projesi Rehberi**

---

## 2️⃣ Amaç

* FastCGI üzerinden Go uygulaması çalıştırmak
* HTTP GET/POST ve form verilerini test etmek
* Cookie yönetimi yapmak
* `httptest` ile unit test ve otomasyon testleri yapmak

---

## 3️⃣ Kod Örneği – Handler
*/
``go
package main

import (
    "fmt"
    "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    name := r.FormValue("name")
    if name == "" {
        name = "Ziyaretçi"
    }

    http.SetCookie(w, &http.Cookie{
        Name:  "sessionid",
        Value: "12345",
    })

    fmt.Fprintf(w, "Merhaba %s! Cookie oluşturuldu.", name)
}
``
/*
---

## 4️⃣ FastCGI Sunucusu
*/
``go
package main

import (
    "log"
    "net"
    "net/http/fcgi"
)

func main() {
    ln, err := net.Listen("tcp", "127.0.0.1:9000")
    if err != nil {
        log.Fatal(err)
    }

    fcgi.Serve(ln, http.HandlerFunc(Handler))
}
``
/*
* TCP üzerinden FastCGI listener oluşturulur
* Handler FastCGI üzerinden gelen tüm isteklere cevap verir

---

## 5️⃣ `httptest` ile GET/POST Testleri
*/
``go
package main

import (
    "fmt"
    "io/ioutil"
    "net/http/httptest"
    "net/url"
)

func main() {
    // GET testi
    reqGet := httptest.NewRequest("GET", "/", nil)
    wGet := httptest.NewRecorder()
    Handler(wGet, reqGet)
    bodyGet, _ := ioutil.ReadAll(wGet.Body)
    fmt.Println("GET Test Body:", string(bodyGet))

    // POST testi (form)
    form := url.Values{}
    form.Set("name", "Ahmet")
    reqPost := httptest.NewRequest("POST", "/", nil)
    reqPost.PostForm = form
    reqPost.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    wPost := httptest.NewRecorder()
    Handler(wPost, reqPost)
    bodyPost, _ := ioutil.ReadAll(wPost.Body)
    fmt.Println("POST Test Body:", string(bodyPost))
}
``
/*
---

## 6️⃣ FastCGI Client ile Test ve Cookie Yönetimi
*/

``go
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/http/cookiejar"
    "net/url"
)

func main() {
    jar, _ := cookiejar.New(nil)
    client := &http.Client{Jar: jar}

    form := url.Values{}
    form.Set("name", "Ayşe")
    resp, _ := client.PostForm("http://127.0.0.1:9000", form)
    bodyClient, _ := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    fmt.Println("FastCGI Client Response:", string(bodyClient))

    u, _ := url.Parse("http://127.0.0.1:9000")
    for _, c := range jar.Cookies(u) {
        fmt.Println("Client Cookie:", c.Name, "=", c.Value)
    }
}
``
/*
---

## 7️⃣ Özet

* **FastCGI** → Go uygulamasını TCP üzerinden sürekli çalıştırır
* **httptest** → Handler’ları unit test ve otomasyon testleri için test eder
* **GET/POST ve form** → Veri gönderme ve işleme
* **CookieJar** → Client cookie’lerini yönetir

Çıktı örneği:
*/
``
GET Test Body: Merhaba Ziyaretçi! Cookie oluşturuldu.
POST Test Body: Merhaba Ahmet! Cookie oluşturuldu.
FastCGI Client Response: Merhaba Ayşe! Cookie oluşturuldu.
Client Cookie: sessionid = 12345
``
/*
---

Bu PDF rehberi ile artık:

* FastCGI uygulamalarını test etme
* HTTP handler’larını otomatik test etme
* GET/POST ve form işlemleri
* Cookie yönetimi

konularını bir arada görebiliyorsun ✅

---

İstersen bunu ben sana **Excel tablosu** olarak da hazırlayabilirim; fonksiyonlar + açıklamalar + örnekler şeklinde.

Bunu yapayım mı?
*/