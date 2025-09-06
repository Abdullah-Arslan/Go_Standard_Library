/*
`expvar` paketi, **Go’nun standart kütüphanesinde bulunan** ve bir uygulamanın **runtime sırasında çalışma istatistiklerini (metrics) ve değişkenlerini dışarıya açmasını** sağlayan bir pakettir. Özellikle **HTTP üzerinden JSON formatında uygulama metrikleri** sunmak için kullanılır.

Bu paket, genellikle **uygulama izleme (monitoring)**, **debugging** veya **performans ölçümleri** amacıyla tercih edilir.

---

## 📌 Genel Özellikler

* `expvar` içindeki değişkenler **global olarak kayıt edilir** ve `http://host:port/debug/vars` endpointinden JSON formatında görüntülenebilir.
* Paket, sayısal değerler (`Int`, `Float`), metinler (`String`), haritalar (`Map`) ve özel değişkenler (`Var` arayüzü ile) tanımlamaya izin verir.
* `http.DefaultServeMux` üzerinde `/debug/vars` endpointi otomatik olarak ayarlanır (tabii ki `net/http/pprof` benzeri).

---

## 📌 Temel Türler

`expvar` paketinde en çok kullanılan türler:

1. **expvar.Int** → Atomik (eşzamanlı güvenli) integer sayaç.
2. **expvar.Float** → Atomik float64 değer.
3. **expvar.String** → Atomik string değer.
4. **expvar.Map** → String-key, `expvar.Var` value saklayan harita.
5. **expvar.Func** → Dinamik olarak hesaplanan değerleri döndürür.

---

## 📌 Kullanım Örnekleri

### 1. Basit Sayaç (Int)
*/
``go
package main

import (
	"expvar"
	"fmt"
	"net/http"
)

var requestCount = expvar.NewInt("request_count")

func handler(w http.ResponseWriter, r *http.Request) {
	requestCount.Add(1) // her istek geldiğinde artır
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
``
/*
➡️ Çalıştırdıktan sonra `http://localhost:8080/debug/vars` adresine gittiğinde JSON çıktısı görürsün:
*/

``json
{
  "cmdline": ["./myapp"],
  "memstats": { ... },
  "request_count": 5
}
``
/*
---

### 2. Float Kullanımı
*/

``go
var temperature = expvar.NewFloat("temperature")

func updateTemperature(newVal float64) {
	temperature.Set(newVal)
}
``

//JSON çıktısı:

``json
{
  "temperature": 36.6
}
``
/*
---

### 3. String Kullanımı
*/

``go
var version = expvar.NewString("app_version")

func main() {
	version.Set("1.0.3")
	http.ListenAndServe(":8080", nil)
}
``

//JSON çıktısı:

``json
{
  "app_version": "1.0.3"
}
``
/*
---

### 4. Map Kullanımı
*/

``go
var stats = expvar.NewMap("stats")

func handler(w http.ResponseWriter, r *http.Request) {
	stats.Add("hits", 1)
	stats.Add("users", 2)
	fmt.Fprintln(w, "Stats updated")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
``

//JSON çıktısı:

``json
{
  "stats": {
    "hits": 3,
    "users": 6
  }
}
``
/*
---

### 5. Dinamik Değerler (Func)

`expvar.Func`, bir fonksiyonun **her çağrıldığında** değer hesaplamasını sağlar:
*/

``go
var uptime = expvar.NewFunc("uptime", func() interface{} {
	return "42s"
})
``

//JSON çıktısı:

``json
{
  "uptime": "42s"
}
``
/*
---

## 📌 expvar Paketinin Sağladığı Otomatik Değerler

`expvar` paketi çalıştığında bazı değişkenleri **otomatik** olarak sağlar:

* **cmdline** → Programın başlatılma argümanları.
* **memstats** → Go’nun `runtime.MemStats` bilgisini JSON formatında verir.
  (Heap kullanımı, GC istatistikleri vs.)

---

## 📌 Özet

* `expvar`, Go uygulamalarında **istatistik ve metrikleri HTTP üzerinden JSON** olarak yayınlamak için kullanılır.
* Kullanım alanları:

  * İstek sayaçları (request\_count)
  * Uptime
  * Versiyon bilgisi
  * Memory kullanımı
  * Dinamik metrikler

Özellikle **prometheus, grafana gibi monitoring sistemlerine entegre etmek** için idealdir.

---

İstersen ben sana `expvar` paketinin tüm fonksiyonlarını (API dokümantasyonu gibi) tek tek açıklayayım ve örneklerle göstereyim. İster misin?
EVET
*/

/*
Harika 👍 O zaman sana **`expvar` paketinin tamamını** (fonksiyonlar, tipler, arayüzler) tek tek açıklayayım ve örneklerle göstereyim.

---

# 📖 Go `expvar` Paketi – Tüm Fonksiyonlar & Tipler

`import "expvar"`

---

## 1. **Var Arayüzü**
*/

``go
type Var interface {
	String() string
}
``
/*
* `expvar` içinde saklanan her şey bu arayüzü uygular.
* Yani her değişken JSON formatında string döndürmelidir.
* `Int`, `Float`, `String`, `Map`, `Func` hepsi `Var` arayüzünü uygular.

---

## 2. **Publish / Get / Do**

### `expvar.Publish(name string, v Var)`

* Bir `Var`’ı global tabloya ekler.
* `/debug/vars` endpointinde görünmesini sağlar.
*/

``go
expvar.Publish("active_users", expvar.NewInt("ignored"))
``
/*
> `ignored` ismi kullanılmaz, önemli olan `Publish`’te verilen addır.

---

### `expvar.Get(name string) Var`

* Daha önce eklenmiş değişkeni getirir.
*/

``go
v := expvar.Get("active_users")
fmt.Println(v.String())
``
/*
---

### `expvar.Do(f func(KeyValue))`

* Tüm kayıtlı değişkenler üzerinde döner.
*/

``go
expvar.Do(func(kv expvar.KeyValue) {
	fmt.Printf("%s = %s\n", kv.Key, kv.Value)
})
``
/*
---

## 3. **KeyValue Yapısı**
*/

``go
type KeyValue struct {
	Key   string
	Value Var
}
``

/*
* `Do` fonksiyonu ile elde edilen her değişken `KeyValue` tipindedir.

---

## 4. **Hazır Türler**

### a) **expvar.Int**

Atomik (goroutine güvenli) integer sayaçtır.
*/

``go
var requests = expvar.NewInt("requests")

func handler(w http.ResponseWriter, r *http.Request) {
	requests.Add(1)
}
``
/*
Metotlar:

* `Add(delta int64)`
* `Set(value int64)`
* `Value() int64`
* `String() string`

---

### b) **expvar.Float**

Atomik float64 saklar.
*/

``go
var loadAvg = expvar.NewFloat("load_avg")

func main() {
	loadAvg.Set(0.85)
}
``
/*
Metotlar:

* `Add(delta float64)`
* `Set(value float64)`
* `Value() float64`
* `String() string`

---

### c) **expvar.String**

Atomik string saklar.
*/

``go
var version = expvar.NewString("version")

func main() {
	version.Set("2.1.0")
}
``
/*
Metotlar:

* `Set(value string)`
* `Value() string`
* `String() string`

---

### d) **expvar.Map**

String-key, `Var` value saklayan haritadır.
*/

``go
var stats = expvar.NewMap("stats")

func main() {
	stats.Add("hits", 1)
	stats.Add("errors", 2)
}
``
/*
Metotlar:

* `Add(key string, delta int64)` → Key yoksa otomatik Int oluşturur.
* `Set(key string, v Var)` → Kendi `Var`’ını ekle.
* `Get(key string) Var`
* `Do(func(KeyValue))` → Map içindeki değerler üzerinde dön.
* `String() string`

---

### e) **expvar.Func**

Her çağrıldığında fonksiyon çalıştırır ve sonucu döner.
*/

``go
expvar.Publish("uptime", expvar.Func(func() interface{} {
	return time.Since(startTime).String()
}))
``

/*
---

## 5. **Hazır Yardımcı Fonksiyonlar**

### `expvar.NewInt(name string) *Int`

Yeni bir integer değişken oluşturur ve global tabloya ekler.

### `expvar.NewFloat(name string) *Float`

Yeni float değişken oluşturur.

### `expvar.NewString(name string) *String`

Yeni string değişken oluşturur.

### `expvar.NewMap(name string) *Map`

Yeni map oluşturur.

### `expvar.NewFunc(name string, f func() interface{}) Func`

Dinamik değer üretmek için kullanılır.

---

## 6. **expvar’ın Otomatik Yayınladığı Değerler**

* `cmdline` → Programın başlatılma argümanları
* `memstats` → `runtime.MemStats` bilgileri (heap, GC, allocation vs.)

Örnek JSON:
*/
``json
{
  "cmdline": ["./app"],
  "memstats": {
    "Alloc": 123456,
    "HeapSys": 987654,
    ...
  },
  "requests": 10,
  "load_avg": 0.85,
  "version": "2.1.0",
  "stats": {
    "hits": 5,
    "errors": 1
  }
}
``
/*
---

## 7. **Tam Uygulama Örneği**
*/
``go
package main

import (
	"expvar"
	"fmt"
	"net/http"
	"time"
)

var (
	requests = expvar.NewInt("requests")
	errors   = expvar.NewInt("errors")
	version  = expvar.NewString("version")
	stats    = expvar.NewMap("stats")
	start    = time.Now()
)

func main() {
	// versiyon bilgisi
	version.Set("1.0.0")

	// uptime
	expvar.Publish("uptime", expvar.Func(func() interface{} {
		return time.Since(start).String()
	}))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requests.Add(1)
		stats.Add("hits", 1)
		fmt.Fprintln(w, "Hello, World!")
	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		errors.Add(1)
		stats.Add("errors", 1)
		http.Error(w, "Something went wrong!", 500)
	})

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
``

//➡️ `http://localhost:8080/debug/vars` çıktısı:

``json
{
  "cmdline": ["./myapp"],
  "memstats": {...},
  "requests": 5,
  "errors": 2,
  "version": "1.0.0",
  "stats": {
    "hits": 5,
    "errors": 2
  },
  "uptime": "1m23s"
}
``
/*
---

✅ Böylece `expvar` paketinin **tüm API’sini** ve nasıl kullanıldığını gördük.
Bu bilgiyi istersen ben sana **PDF olarak bir cheatsheet / özet** hazırlayıp verebilirim. İstiyor musun?
*/