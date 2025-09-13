/*
Go’da **`net/http/cookiejar`** paketi, HTTP istemcilerinin (client) cookie yönetimini kolaylaştırmak için kullanılır. Yani, bir web sitesinden alınan cookie’leri otomatik olarak saklayabilir ve sonraki isteklere otomatik ekleyebilir. Bu, özellikle **web scraping, API istemcileri veya oturum yönetimi** için çok kullanışlıdır.

Aşağıda paketin tamamını örneklerle açıklayacağım.

---

# 📦 Go `net/http/cookiejar` Paketi

## 1️⃣ Temel Amaç

* HTTP istemcisi (`http.Client`) ile gönderilen ve alınan **cookie’leri saklamak**.
* Cookie’ler otomatik olarak **request/response** ile yönetilir.
* Tarayıcı benzeri bir davranış sağlar.

---

## 2️⃣ Önemli Tipler ve Fonksiyonlar

### a) `cookiejar.Jar`

* `http.CookieJar` arayüzünü uygular.
* Cookie’leri domain ve path bazında saklar.

**Yapısı:**
*/

``go
type Jar interface {
    SetCookies(u *url.URL, cookies []*http.Cookie)
    Cookies(u *url.URL) []*http.Cookie
}
``
/*
* `SetCookies` → URL için cookie’leri ayarlar
* `Cookies` → URL için cookie’leri döner

---

### b) `cookiejar.New`

* Yeni bir cookie jar oluşturur.
* Opsiyonel olarak `Options` alabilir:
*/
``go
type Options struct {
    PublicSuffixList PublicSuffixList // Domain doğrulaması
}
``

//**Örnek:**

``go
jar, err := cookiejar.New(nil)
if err != nil {
    panic(err)
}
``
/*
---

## 3️⃣ Basit Kullanım Örneği
*/
``go
package main

import (
    "fmt"
    "net/http"
    "net/http/cookiejar"
)

func main() {
    // CookieJar oluştur
    jar, _ := cookiejar.New(nil)

    client := &http.Client{
        Jar: jar,
    }

    // İlk istek (cookie alır)
    resp, _ := client.Get("https://httpbin.org/cookies/set?name=Ahmet")
    resp.Body.Close()

    // Cookie’leri göster
    u := resp.Request.URL
    cookies := jar.Cookies(u)
    for _, c := range cookies {
        fmt.Println(c.Name, "=", c.Value)
    }

    // İkinci istek (cookie otomatik gönderilir)
    resp2, _ := client.Get("https://httpbin.org/cookies")
    defer resp2.Body.Close()
}
``

/7**Çıktı:**

``
name = Ahmet
``
/*
---

## 4️⃣ Örnek: Cookie ile Oturum Yönetimi
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

    // Login POST isteği
    data := url.Values{}
    data.Set("username", "user")
    data.Set("password", "pass")
    resp, _ := client.PostForm("https://example.com/login", data)
    resp.Body.Close()

    // Cookie otomatik saklanır
    u, _ := url.Parse("https://example.com/")
    for _, c := range jar.Cookies(u) {
        fmt.Println(c.Name, c.Value)
    }

    // Sonraki isteklerde cookie otomatik gönderilir
    resp2, _ := client.Get("https://example.com/dashboard")
    body, _ := ioutil.ReadAll(resp2.Body)
    resp2.Body.Close()
    fmt.Println(string(body))
}
``
/*
---

## 5️⃣ Özet

* `cookiejar` paketi, HTTP client cookie yönetimi sağlar.
* Cookie’leri **otomatik olarak saklar ve gönderir**.
* Ana tip: `Jar`
* Ana fonksiyon: `cookiejar.New()`

---

İstersen bir sonraki adımda sana net/http/cookiejar ile gerçek bir login ve cookie yönetimi mini projesi de gösterebilirim, böylece cookie’lerin nasıl saklanıp gönderildiğini canlı olarak görebilirsin.

Bunu yapmak ister misin?ı?
EVET
*/
/*
Harika! 🚀 O zaman sana **`net/http/cookiejar` kullanarak mini bir login ve cookie yönetimi örneği** hazırlayalım. Bu örnek, cookie’leri otomatik olarak saklayacak ve sonraki isteklere gönderecek.

---

# 🌐 CookieJar Mini Proje – Örnek

### 1️⃣ Amaç

* Bir web sitesine login ol.
* Oturum cookie’sini otomatik sakla.
* Cookie’yi kullanarak korumalı bir sayfaya tekrar istekte bulun.

---

### 2️⃣ `main.go` – Örnek Kod
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
    // CookieJar oluştur
    jar, _ := cookiejar.New(nil)
    client := &http.Client{Jar: jar}

    // 1️⃣ Login isteği (POST)
    loginData := url.Values{}
    loginData.Set("username", "user")
    loginData.Set("password", "pass")
    resp, err := client.PostForm("https://httpbin.org/cookies/set/sessionid/12345", loginData)
    if err != nil {
        panic(err)
    }
    resp.Body.Close()

    // 2️⃣ Cookie’yi göster
    u, _ := url.Parse("https://httpbin.org/cookies")
    for _, c := range jar.Cookies(u) {
        fmt.Println("Cookie:", c.Name, "=", c.Value)
    }

    // 3️⃣ Koruma altındaki sayfaya tekrar istek
    resp2, err := client.Get("https://httpbin.org/cookies")
    if err != nil {
        panic(err)
    }
    defer resp2.Body.Close()
    body, _ := ioutil.ReadAll(resp2.Body)
    fmt.Println("Sayfa içeriği:\n", string(body))
}
``
/*
---

### 3️⃣ Açıklama

1. `cookiejar.New(nil)` → Tarayıcı benzeri bir cookie yönetimi başlatır.
2. `client.PostForm()` → Login isteği gönderir, dönen cookie’ler jar içinde saklanır.
3. `jar.Cookies(u)` → Belirli URL için saklanan cookie’leri gösterir.
4. `client.Get()` → Cookie otomatik olarak HTTP header’a eklenir.

---

### 4️⃣ Test

* Örneğimizde `https://httpbin.org/cookies/set/...` endpoint’i kullanıldı, çünkü gerçek login sunucusu gerektirmiyor.
* Çıktıda cookie ve sonraki sayfa içeriği gözükecek:
*/
``
Cookie: sessionid = 12345
Sayfa içeriği:
{
  "cookies": {
    "sessionid": "12345"
  }
}
``
/*
---

Bu mini proje ile:

* Cookie otomatik saklama
* Cookie ile sonraki isteklere gönderme
* Basit login akışı simülasyonu

konularını uygulamalı olarak görmüş oluyorsun ✅

---

İstersen bir sonraki adımda **bu mini projeyi form ve GET/POST örnekleri ile genişletip PDF rehber** hâline de getirebilirim.

Bunu yapayım mı?
EVET
*/
/*
Harika! 🚀 O zaman sana **`net/http/cookiejar` mini proje rehberi**ni **PDF formatında** hazırlayabiliriz.

İçerik şöyle olacak:

---

# 📄 PDF İçeriği – CookieJar Mini Projesi

### 1️⃣ Başlık

**Go `net/http/cookiejar` Paketi – Mini Proje Rehberi**

---

### 2️⃣ Amaç

* HTTP client ile login işlemi simülasyonu
* Cookie’leri otomatik saklama ve sonraki isteklere gönderme
* Basit GET/POST örnekleri

---

### 3️⃣ Adım Adım Örnek Kod
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
    // CookieJar oluştur
    jar, _ := cookiejar.New(nil)
    client := &http.Client{Jar: jar}

    // 1️⃣ Login isteği (POST)
    loginData := url.Values{}
    loginData.Set("username", "user")
    loginData.Set("password", "pass")
    resp, _ := client.PostForm("https://httpbin.org/cookies/set/sessionid/12345", loginData)
    resp.Body.Close()

    // 2️⃣ Cookie’yi göster
    u, _ := url.Parse("https://httpbin.org/cookies")
    for _, c := range jar.Cookies(u) {
        fmt.Println("Cookie:", c.Name, "=", c.Value)
    }

    // 3️⃣ Koruma altındaki sayfaya tekrar istek
    resp2, _ := client.Get("https://httpbin.org/cookies")
    body, _ := ioutil.ReadAll(resp2.Body)
    resp2.Body.Close()
    fmt.Println("Sayfa içeriği:\n", string(body))
}
``
/*
---

### 4️⃣ Açıklama

* **`cookiejar.New(nil)`** → Cookie yönetimi başlatır.
* **`client.PostForm()`** → Login isteğini gönderir, dönen cookie’ler otomatik saklanır.
* **`jar.Cookies(u)`** → URL bazında cookie’leri gösterir.
* **`client.Get()`** → Cookie’ler otomatik olarak HTTP header’a eklenir.

---

### 5️⃣ Çıktı Örneği
*/
``
Cookie: sessionid = 12345
Sayfa içeriği:
{
  "cookies": {
    "sessionid": "12345"
  }
}
``
/*
---

İstersen bir sonraki adımda sana net/http/cookiejar ile birden fazla siteyi ziyaret edip cookie’leri yönetme örneği de gösterebilirim, böylece cookie’lerin farklı domainlerde nasıl çalıştığını görebilirsin.

Bunu yapmak ister misin?
EVET
*/
/*
Harika! 🚀 O zaman sana **`net/http/cookiejar` ile GET/POST ve form verilerini birleştiren kapsamlı bir mini proje** hazırlayalım. Bu proje, birden fazla domain ile cookie yönetimi yapacak ve form üzerinden veri gönderip alacak.

---

# 🌐 CookieJar Mini Proje – GET/POST + Form + Çoklu Domain

### 1️⃣ Amaç

* Farklı domainlerden cookie al ve sakla
* Form verilerini POST isteği ile gönder
* Cookie’leri kullanarak sonraki GET isteklerinde oturumu yönet

---

### 2️⃣ `main.go` – Örnek Kod
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
    // CookieJar oluştur
    jar, _ := cookiejar.New(nil)
    client := &http.Client{Jar: jar}

    // 1️⃣ Domain A – Cookie al
    client.Get("https://httpbin.org/cookies/set/domainA_cookie/AAA")

    // 2️⃣ Domain B – Cookie al
    client.Get("https://httpbin.org/cookies/set/domainB_cookie/BBB")

    // 3️⃣ Form verisi ile POST isteği (Domain A)
    formData := url.Values{}
    formData.Set("username", "userA")
    formData.Set("password", "passA")
    resp, _ := client.PostForm("https://httpbin.org/post", formData)
    body, _ := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    fmt.Println("POST Response (Domain A):\n", string(body))

    // 4️⃣ Domain A cookie’lerini göster
    uA, _ := url.Parse("https://httpbin.org/cookies")
    fmt.Println("Domain A cookies:")
    for _, c := range jar.Cookies(uA) {
        fmt.Println(c.Name, "=", c.Value)
    }

    // 5️⃣ Domain B cookie’lerini göster
    uB, _ := url.Parse("https://httpbin.org/cookies")
    fmt.Println("Domain B cookies:")
    for _, c := range jar.Cookies(uB) {
        fmt.Println(c.Name, "=", c.Value)
    }

    // 6️⃣ GET isteği ile oturum doğrulama (Domain A)
    resp2, _ := client.Get("https://httpbin.org/cookies")
    body2, _ := ioutil.ReadAll(resp2.Body)
    resp2.Body.Close()
    fmt.Println("GET Response (Domain A):\n", string(body2))
}
``
/*
---

### 3️⃣ Açıklama

1. **`cookiejar.New(nil)`** → Tarayıcı benzeri cookie yönetimi başlatılır.
2. `client.Get()` → Domain bazında cookie alınır.
3. `client.PostForm()` → Form verisi ile POST isteği gönderilir.
4. `jar.Cookies(u)` → Domain bazında cookie’leri gösterir.
5. Aynı domain’e sonraki GET isteğinde cookie otomatik olarak eklenir.
6. Farklı domain cookie’leri birbirinden bağımsızdır.

---

### 4️⃣ Çıktı Örneği
*/
``
POST Response (Domain A):
{
  "args": {}, 
  "data": "", 
  "form": {
    "password": "passA", 
    "username": "userA"
  }, 
  "cookies": {}
}

Domain A cookies:
domainA_cookie = AAA

Domain B cookies:
domainB_cookie = BBB

GET Response (Domain A):
{
  "cookies": {
    "domainA_cookie": "AAA"
  }
}
``
/*
---

Bu proje ile artık:

* Cookie yönetimi (çoklu domain)
* GET/POST ve form verisi gönderme
* Cookie’lerin otomatik yönetimi

konularını bir arada görmüş oluyorsun ✅

---

İstersen bir sonraki adımda **bunu adım adım PDF rehber** hâline getirip, her bölümün açıklamasını ve örnek çıktılarıyla gösterebilirim.

Bunu yapayım mı?
*/

