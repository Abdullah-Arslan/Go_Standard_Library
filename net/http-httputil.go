/*
Golang’daki **`net/http/httputil`** paketi, `net/http` paketi ile birlikte HTTP istekleri ve yanıtlarını daha kolay analiz etmek, hata ayıklamak ve proxy gibi ileri seviye kullanım senaryolarında işimizi kolaylaştırmak için geliştirilmiş yardımcı araçlar içerir.

Bu paket özellikle:

* HTTP isteklerini ve yanıtlarını kolay okunabilir hale getirmek,
* Reverse proxy uygulamaları yazmak,
* HTTP dump (kaydetme) işlemleri yapmak,
* Geliştirme sırasında debug etmek

gibi işlevler için kullanılır.

---

## Paketin Sağladığı Temel Fonksiyonlar ve Yapılar

### 1. `DumpRequest` ve `DumpRequestOut`

Bir HTTP isteğini (`http.Request`) ham formatta (raw HTTP) string veya byte slice olarak döndürür.

* **`DumpRequest(req *http.Request, body bool)`** → Sunucuya gelen istekleri dump etmek için.
* **`DumpRequestOut(req *http.Request, body bool)`** → İstemciden giden istekleri dump etmek için.

**Örnek: DumpRequest**

``go
package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Gelen isteği dump et
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, "DumpRequest hatası", http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump)) // Terminalde isteği göster
	w.Write([]byte("İstek alındı"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
``
/*
Tarayıcıdan `http://localhost:8080/test?x=1` istek atarsan terminalde raw HTTP isteğini görürsün.

---

### 2. `DumpResponse`

Bir HTTP yanıtını (`http.Response`) ham formatta döndürür.

**Örnek: DumpResponse**
*/
``go
package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
}
``
/*
Bu kod `https://httpbin.org/get` yanıtını raw HTTP response formatında gösterir.

---

### 3. `ReverseProxy`

Bir reverse proxy oluşturmanı sağlar.

* Proxy, gelen istekleri başka bir sunucuya yönlendirir.
* `httputil.NewSingleHostReverseProxy(target *url.URL)` fonksiyonu sık kullanılır.

**Örnek: Reverse Proxy**
*/
``go
package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	target, _ := url.Parse("https://httpbin.org")
	proxy := httputil.NewSingleHostReverseProxy(target)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r) // gelen isteği hedefe yönlendir
	})

	http.ListenAndServe(":8080", nil)
}
``
/*
Şimdi `http://localhost:8080/get` adresine gittiğinde, aslında istek `https://httpbin.org/get` adresine gider.

---

### 4. `ClientConn` ve `ServerConn` (daha eski kullanım)

Bu yapılar HTTP üzerinden düşük seviyeli client-server bağlantılarını işlemek için kullanılıyordu.
Ama **Go 1.8’den sonra deprecated (kullanımı önerilmiyor)**. Modern uygulamalarda `http.Transport` ve `http.Server` kullanılır.

**Örnek (deprecated)**:
*/
``go
// Bu kod sadece eski API örneği olarak gösteriliyor
// Artık pek kullanılmıyor.
``
/*
---

### 5. `DumpRequestOut` (özellikle client tarafında)

Bu fonksiyon, client tarafından yapılacak bir isteğin dump’ını almak için kullanılır.

**Örnek: DumpRequestOut**
*/
``go
package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	req, _ := http.NewRequest("GET", "https://httpbin.org/get", nil)
	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
}
``

---

## Özet Tablo

| Fonksiyon / Struct          | Açıklama                                      | Kullanım Alanı                 |
| --------------------------- | --------------------------------------------- | ------------------------------ |
| `DumpRequest`               | HTTP isteğini raw formatta döndürür           | Sunucu tarafı debug            |
| `DumpRequestOut`            | İstemciden çıkan isteği raw formatta döndürür | Client debug                   |
| `DumpResponse`              | HTTP yanıtını raw formatta döndürür           | Response debug                 |
| `NewSingleHostReverseProxy` | Reverse proxy oluşturur                       | Proxy servisleri               |
| `ReverseProxy` struct       | Proxy davranışını özelleştirmek için          | Load balancer / Gateway        |
| `ClientConn` (deprecated)   | Düşük seviye client bağlantısı                | Yerine `http.Transport` kullan |
| `ServerConn` (deprecated)   | Düşük seviye server bağlantısı                | Yerine `http.Server` kullan    |

/*
---

👉 Yani `http/httputil` paketi, Go’nun **HTTP debug aracı + proxy aracı** diyebiliriz.

---

Özetlemek gerekirse:

http/httputil paketi özellikle debug ve reverse proxy senaryolarında kullanılır.

En çok kullanılanları:

DumpRequest, DumpRequestOut, DumpResponse → istek ve yanıtları ham (raw HTTP) olarak görebilmek.

NewSingleHostReverseProxy → basit bir reverse proxy kurmak.

İstersen ben sana ReverseProxy’yi özelleştirme (örneğin: header ekleme, response’u değiştirme, loglama) örnekleri de yazabilirim.

Bunu ister misin?
EVET
*/
/*
Harika 👌 o zaman `httputil.ReverseProxy`’yi özelleştirme örneklerini adım adım göstereyim.

Normalde `httputil.NewSingleHostReverseProxy` kullanınca sadece **gelen isteği aynen hedef sunucuya yönlendirir**. Ama çoğu durumda:

* Header eklemek/değiştirmek,
* Yanıtı değiştirmek,
* Loglamak

gerekebilir. İşte bu noktada `ReverseProxy` özelleştirilebilir.

---

## 1. **İstek Üzerinde Değişiklik Yapma**

`Director` fonksiyonu ile hedefe gitmeden önce isteği değiştirebilirsin.
*/
``go
package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	target, _ := url.Parse("https://httpbin.org")
	proxy := httputil.NewSingleHostReverseProxy(target)

	// Director fonksiyonunu özelleştirelim
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)

		// Header ekle
		req.Header.Set("X-Proxy-Header", "BenimProxy")

		// Query parametre ekle
		q := req.URL.Query()
		q.Add("extra", "1")
		req.URL.RawQuery = q.Encode()
	}

	http.ListenAndServe(":8080", proxy)
}
``
/*
👉 `http://localhost:8080/get` çağırdığında, hedef sunucuya ek header ve query parametre ile gidecektir.

---

## 2. **Yanıtı Değiştirme**

`ModifyResponse` fonksiyonunu kullanabilirsin. Bu fonksiyon hedef sunucudan gelen yanıtı yakalar.
*/
``go
package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	target, _ := url.Parse("https://httpbin.org")
	proxy := httputil.NewSingleHostReverseProxy(target)

	// Yanıtı değiştir
	proxy.ModifyResponse = func(resp *http.Response) error {
		// Response body'yi oku
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		resp.Body.Close()

		// İçeriği değiştir
		newBody := bytes.ReplaceAll(body, []byte("httpbin"), []byte("BENİM_PROXY"))

		// Yeni body'yi ayarla
		resp.Body = io.NopCloser(bytes.NewReader(newBody))
		resp.ContentLength = int64(len(newBody))
		resp.Header.Set("Content-Length", string(len(newBody)))

		return nil
	}

	http.ListenAndServe(":8080", proxy)
}
``
/*
👉 Artık yanıt gövdesinde `"httpbin"` geçen yerler `"BENİM_PROXY"` olarak değiştirilecektir.

---

## 3. **Hata Yönetimi (ErrorHandler)**

Hedef sunucuya ulaşılamadığında kendi özel hata mesajını verebilirsin.
*/
``go
package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	target, _ := url.Parse("https://yanlis-adres.org") // geçersiz adres
	proxy := httputil.NewSingleHostReverseProxy(target)

	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		fmt.Println("Proxy hatası:", err)
		http.Error(w, "Sunucuya ulaşılamıyor. Lütfen sonra tekrar deneyin.", http.StatusBadGateway)
	}

	http.ListenAndServe(":8080", proxy)
}
``
/*
👉 Hedef sunucuya ulaşılamadığında kullanıcıya özel mesaj döner.

---

## 4. **Tam Özelleştirilmiş Proxy**

Hem `Director`, hem `ModifyResponse`, hem de `ErrorHandler` aynı anda kullanılabilir.
*/
``go
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	target, _ := url.Parse("https://httpbin.org")
	proxy := httputil.NewSingleHostReverseProxy(target)

	// İsteği özelleştir
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.Header.Set("X-Custom-Header", "ProxyKullaniyorum")
	}

	// Yanıtı özelleştir
	proxy.ModifyResponse = func(resp *http.Response) error {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		resp.Body.Close()

		newBody := bytes.ReplaceAll(body, []byte("httpbin"), []byte("PROXY_DEGISTIRDI"))
		resp.Body = io.NopCloser(bytes.NewReader(newBody))
		resp.ContentLength = int64(len(newBody))
		resp.Header.Set("Content-Length", fmt.Sprint(len(newBody)))
		return nil
	}

	// Hata yönetimi
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		fmt.Println("Proxy hatası:", err)
		http.Error(w, "Proxy üzerinden erişim hatası.", http.StatusBadGateway)
	}

	http.ListenAndServe(":8080", proxy)
}
``
/*
---

✅ Görüldüğü gibi `httputil.ReverseProxy`, sadece basit bir proxy değil, **tamamen özelleştirilebilir bir HTTP gateway** gibi çalışabiliyor.

---

İstersen sana **load balancer** (örneğin gelen istekleri birden fazla backend’e dağıtma) örneği de gösterebilirim. İstiyor musun?
EVET
*/
/*
Süper 👍 o zaman şimdi sana `httputil.ReverseProxy` kullanarak **basit bir load balancer (yük dengeleyici)** örneği göstereyim.

Load balancer’ın amacı:

* Gelen istekleri **birden fazla backend sunucuya** dağıtmak.
* Örneğin: `http://localhost:8080` → backend1, backend2, backend3 arasında dönüşümlü olarak yönlendirsin.

---

## 1. **Round-Robin Load Balancer**

Bu örnekte gelen her isteği sırayla farklı backend’e yönlendireceğiz.
*/
``go
package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

// Backend adresleri
var backends = []string{
	"http://localhost:9001",
	"http://localhost:9002",
	"http://localhost:9003",
}

var counter uint64

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Round-robin algoritması
		index := atomic.AddUint64(&counter, 1)
		target := backends[int(index)%len(backends)]

		url, _ := url.Parse(target)
		proxy := httputil.NewSingleHostReverseProxy(url)

		// Hangi backend’e yönlendirildiğini logla
		w.Header().Set("X-Backend", url.Host)

		proxy.ServeHTTP(w, r)
	})

	http.ListenAndServe(":8080", nil)
}
``
/*
👉 Kullanım:

* 3 tane basit HTTP sunucu aç (`:9001`, `:9002`, `:9003`)
* Load balancer’ı `:8080`’de çalıştır.
* Tarayıcıdan `http://localhost:8080`’e arka arkaya istek atarsan, istekler sırayla 9001 → 9002 → 9003 → 9001 … şeklinde dağıtılır.

---

## 2. **Rastgele (Random) Load Balancer**

Bazen round-robin yerine rastgele sunucu seçmek isteyebilirsin.
*/
``go
package main

import (
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

var backends = []string{
	"http://localhost:9001",
	"http://localhost:9002",
	"http://localhost:9003",
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		target := backends[rand.Intn(len(backends))]

		url, _ := url.Parse(target)
		proxy := httputil.NewSingleHostReverseProxy(url)

		w.Header().Set("X-Backend", url.Host)

		proxy.ServeHTTP(w, r)
	})

	http.ListenAndServe(":8080", nil)
}
``
/*
👉 İstekler rastgele bir backend’e yönlendirilir.

---

## 3. **Health Check ile Load Balancer**

Bir backend sunucusu çökerse diğerlerini kullanmaya devam edebilmek için “health check” mantığı ekleyebiliriz.
*/
``go
package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
	"time"
)

var backends = []string{
	"http://localhost:9001",
	"http://localhost:9002",
	"http://localhost:9003",
}

var alive = make([]bool, len(backends))
var counter uint64

// Backend sağlıklı mı kontrol et
func healthCheck() {
	for i, backend := range backends {
		resp, err := http.Get(backend + "/health")
		if err != nil || resp.StatusCode != 200 {
			alive[i] = false
		} else {
			alive[i] = true
		}
		if resp != nil {
			resp.Body.Close()
		}
	}
}

func getNextAliveBackend() string {
	for i := 0; i < len(backends); i++ {
		index := int(atomic.AddUint64(&counter, 1) % uint64(len(backends)))
		if alive[index] {
			return backends[index]
		}
	}
	return "" // hiçbir backend sağlıklı değilse
}

func main() {
	// Her 5 saniyede bir health check
	go func() {
		for {
			healthCheck()
			time.Sleep(5 * time.Second)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		target := getNextAliveBackend()
		if target == "" {
			http.Error(w, "Hiçbir backend ayakta değil!", http.StatusServiceUnavailable)
			return
		}

		url, _ := url.Parse(target)
		proxy := httputil.NewSingleHostReverseProxy(url)

		w.Header().Set("X-Backend", url.Host)
		proxy.ServeHTTP(w, r)
	})

	fmt.Println("Load balancer çalışıyor :8080")
	http.ListenAndServe(":8080", nil)
}
``
/*
👉 Burada backend sunucuların `/health` endpoint’i çalışıyor olmalı ve `200 OK` dönmeli. Sağlıksız olanlar otomatik devre dışı bırakılır.

---

✅ Yani `httputil.ReverseProxy`, sadece proxy değil, **çok basit load balancer’lar** yazmak için de kullanılabiliyor.

---

İstersen sana bu load balancer’ı **daha gelişmiş hale getirip (örneğin ağırlıklı dağıtım: %70 backend1, %30 backend2)** örnek de yazabilirim. Bunu ister misin?
EVET
*/

/*
Harika 👍 o zaman şimdi sana **Weighted Load Balancer (Ağırlıklı Yük Dengeleyici)** örneğini göstereyim.

Bu modelde:

* Backend sunuculara farklı **ağırlıklar (weight)** verilir.
* Örneğin:

  * Backend1 → %70
  * Backend2 → %20
  * Backend3 → %10
* Böylece gelen istekler bu oranlara göre dağıtılır.

---

## **Weighted Load Balancer Örneği**
*/
``go
package main

import (
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

// Backend yapısı
type Backend struct {
	URL    string
	Weight int // Yük oranı (ağırlık)
}

// Backend listesi
var backends = []Backend{
	{"http://localhost:9001", 70}, // %70
	{"http://localhost:9002", 20}, // %20
	{"http://localhost:9003", 10}, // %10
}

// Ağırlıklı seçim yapmak için slice
var weightedPool []Backend

// Weighted pool hazırlama
func buildWeightedPool() {
	for _, b := range backends {
		for i := 0; i < b.Weight; i++ {
			weightedPool = append(weightedPool, b)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	buildWeightedPool()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Rastgele bir backend seç ama ağırlığa göre
		chosen := weightedPool[rand.Intn(len(weightedPool))]

		url, _ := url.Parse(chosen.URL)
		proxy := httputil.NewSingleHostReverseProxy(url)

		// Hangi backend seçildiğini response header’a ekleyelim
		w.Header().Set("X-Backend", url.Host)

		proxy.ServeHTTP(w, r)
	})

	http.ListenAndServe(":8080", nil)
}
``
/*
---

### 📝 Nasıl Çalışır?

* `buildWeightedPool()` fonksiyonu ağırlıkları baz alarak bir havuz oluşturuyor.

  * Backend1 (`9001`) → 70 defa eklendi.
  * Backend2 (`9002`) → 20 defa eklendi.
  * Backend3 (`9003`) → 10 defa eklendi.
* Gelen istekler bu havuzdan **rastgele** seçiliyor.
* Sonuçta uzun vadede:

  * %70 `9001`,
  * %20 `9002`,
  * %10 `9003` istek alır.

---

✅ Böylece kendi basit **weighted round robin load balancer** yazmış olduk.

---

İstersen ben sana bu load balancer’a **health check (sağlıklı olmayan backend’i devre dışı bırakma)** özelliğini de ekleyebilirim. Bunu ister misin?
*/