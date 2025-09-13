/*
## Go `plugin` Paketi Nedir?

Go’nun `plugin` paketi, **Linux** ve **macOS**’ta dinamik olarak derlenmiş paylaşımlı nesne dosyalarını (`.so`) çalıştırma zamanında yüklemenizi sağlar.
Bu sayede programınıza yeniden derleme yapmadan yeni işlevler veya semboller ekleyebilirsiniz.

📌 **Notlar:**

* Sadece **Linux** ve **macOS** üzerinde çalışır. Windows desteklenmez.
* `.so` dosyaları `go build -buildmode=plugin` ile üretilir.
* Yüklenen eklentiler (plugin) sadece bir kez yüklenebilir. Aynı dosyayı tekrar açmak aynı nesneyi döndürür.

---

## Temel Kavramlar

* \**plugin.Open(path string) (*plugin.Plugin, error)**
  Belirtilen `.so` dosyasını açar.

* \**(*plugin.Plugin).Lookup(symName string) (plugin.Symbol, error)**
  Açılan plugin içindeki fonksiyon veya değişkeni sembol adı ile bulur.

* **plugin.Symbol**
  `interface{}` türünde bir alias’tır. Genellikle `func` veya `*var` tipine **type assertion** yapılır.

---

## Örnek 1 – Basit Kullanım

### 1. Plugin Kodu (mathplugin.go)
*/
``go
package main

import "fmt"

// Dışarıya açılan fonksiyon
func Add(a, b int) int {
	return a + b
}

// Dışarıya açılan değişken
var PluginName = "Simple Math Plugin"

// Eklenti giriş noktası main paketi olmalı, ama main() fonksiyonu yazılmaz.
func init() {
	fmt.Println("Math Plugin yüklendi.")
}
``

//Bunu şu şekilde derleyin:

``bash
go build -buildmode=plugin -o mathplugin.so mathplugin.go
``

//### 2. Ana Program (main.go)

``go
package main

import (
	"fmt"
	"plugin"
)

func main() {
	// Plugin aç
	p, err := plugin.Open("mathplugin.so")
	if err != nil {
		panic(err)
	}

	// Fonksiyon sembolünü bul
	addSym, err := p.Lookup("Add")
	if err != nil {
		panic(err)
	}

	// Type assertion ile fonksiyon haline getir
	addFunc := addSym.(func(int, int) int)

	// Kullan
	fmt.Println("3 + 5 =", addFunc(3, 5))

	// Değişken sembolünü bul
	nameSym, err := p.Lookup("PluginName")
	if err != nil {
		panic(err)
	}
	name := nameSym.(*string)

	fmt.Println("Plugin Adı:", *name)
}
``

//### Çalıştırma

``bash
go run main.go
``

//📌 Çıktı:

``
Math Plugin yüklendi.
3 + 5 = 8
Plugin Adı: Simple Math Plugin
``
/*
---

## Örnek 2 – Farklı Plugin ile Dinamik Davranış

Birden fazla `.so` dosyası hazırlayabilir ve aynı ana programda farklı pluginleri seçerek yükleyebilirsiniz.

### greeter.go
*/
``go
package main

import "fmt"

func Greet(name string) {
	fmt.Println("Merhaba,", name)
}
``

//Derleme:

``bash
go build -buildmode=plugin -o greeter.so greeter.go
``

//### main.go

``go
package main

import (
	"plugin"
)

func main() {
	p, _ := plugin.Open("greeter.so")
	greetSym, _ := p.Lookup("Greet")
	greet := greetSym.(func(string))
	greet("Abdullah")
}
``

//Çıktı:

``
Merhaba, Abdullah
``
/*
---

## Örnek 3 – Plugin ile Strateji Deseni

Bir “plugin” mantığını **Strategy Pattern** gibi kullanabilirsiniz.
Örneğin farklı matematik işlemlerini farklı `.so` dosyaları olarak hazırlayabilir, çalışma zamanında seçebilirsiniz.

* `add.so` → toplama
* `mul.so` → çarpma
* `div.so` → bölme

Ana program, sadece “hangi .so dosyası seçildi” bilgisini bilerek doğru fonksiyonu çağırır.

---

## Örnek 4 – Hata Durumları
*/
``go
p, err := plugin.Open("notfound.so")
if err != nil {
	fmt.Println("Plugin bulunamadı:", err)
}

sym, err := p.Lookup("MissingFunc")
if err != nil {
	fmt.Println("Sembol bulunamadı:", err)
}
``
/*
📌 Bu durumda:

* Eğer dosya yoksa → `plugin.Open` hata döner.
* Eğer sembol yoksa → `Lookup` hata döner.

---

## Önemli Sınırlamalar

* **Windows desteği yok**.
* **Sadece `main` paketinden export edilen semboller dışarı açılır.**
* Pluginler tekrar tekrar yüklenmez (cache vardır).
* API değişirse uyumsuzluk çıkabilir, çünkü `.so` Go sürümüne bağımlıdır.

---

✅ Özetle:
Go’nun `plugin` paketi, `.so` dosyalarını açıp içindeki fonksiyon/değişkenleri dinamik olarak kullanmanızı sağlar. Bu, özellikle **genişletilebilir yazılımlar**, **modüler uygulamalar** veya **plugin tabanlı mimariler** için çok güçlü bir özelliktir.

---

İstersen sana bu `plugin` paketini **“gerçek dünyada blog projen” için admin modülüne plugin eklemek** tarzında bir demo hazırlayayım mı?
EVET
*/
/*
Harika 👍 O zaman sana **Go blog projen için plugin tabanlı admin modülü** örneği hazırlayayım.
Böylece blog’unu tekrar derlemeden yeni admin özellikleri ekleyebileceksin.

---

# 📦 Blog için Plugin Tabanlı Admin Modülü

## 1. Admin Plugin Örneği (`admin_stats.go`)
*/
``go
package main

import "fmt"

// Admin paneline eklenecek fonksiyon
func ShowStats() {
	fmt.Println("Blog istatistikleri:")
	fmt.Println("- Toplam Kullanıcı: 128")
	fmt.Println("- Toplam Yazı: 52")
	fmt.Println("- Toplam Kategori: 7")
}
``

//Bunu **plugin olarak derle**:

``bash
go build -buildmode=plugin -o admin_stats.so admin_stats.go
``
/*
---

## 2. Ana Blog Kodu (`main.go`)

Senin blogunda `httprouter` kullandığını biliyorum, bu yüzden plugin çağrısını **/admin/plugin** route’una ekliyorum:
*/
``go
package main

import (
	"fmt"
	"net/http"
	"plugin"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	// Ana sayfa
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintln(w, "Blog Ana Sayfa")
	})

	// Plugin tabanlı admin endpoint
	router.GET("/admin/plugin", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// Plugin yükle
		p, err := plugin.Open("./admin_stats.so")
		if err != nil {
			http.Error(w, "Plugin yüklenemedi: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Plugin fonksiyonunu bul
		sym, err := p.Lookup("ShowStats")
		if err != nil {
			http.Error(w, "Fonksiyon bulunamadı: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Type assertion
		showStats := sym.(func())

		// Plugin çalıştır
		fmt.Fprintln(w, "Plugin çalıştırılıyor...")
		showStats()
	})

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", router)
}
``
/*
---

## 3. Çalıştırma

1. Önce plugin’i derle:
*/
   ``bash
   go build -buildmode=plugin -o admin_stats.so admin_stats.go
   ``

//2. Sonra blog server’ı çalıştır:

   ``bash
   go run main.go
   ``
/*
3. Tarayıcıdan aç:
   👉 [http://localhost:8080/admin/plugin](http://localhost:8080/admin/plugin)

📌 Çıktı (tarayıcıda ve terminalde):
*/
``
Plugin çalıştırılıyor...
Blog istatistikleri:
- Toplam Kullanıcı: 128
- Toplam Yazı: 52
- Toplam Kategori: 7
``
/*
---

## 4. Avantajı

* Yeni bir **admin özelliği** (örneğin `admin_users.go`, `admin_logs.go`) hazırlayıp `.so` olarak derlersin.
* Ana uygulaman **değişmeden** sadece `.so` dosyasını ekleyerek blog’unu genişletirsin.
* Docker konteynerine sadece yeni `.so` dosyasını kopyalaman yeterli olur.

---

👉 İstersen ben sana hemen **ikinci bir plugin** (örneğin `admin_users.go`) yazayım, ve aynı `/admin/plugin?mod=users` parametresiyle hangi plugin’in yükleneceğini seçelim. Böylece **çoklu plugin yönetimi** mantığını da gösteririm.

Onu da ister misin?
EVET
*/
/*
Süper 👍 Şimdi sana **çoklu plugin yönetimi** için örnek hazırlıyorum.
Amaç şu olacak:

* `/admin/plugin?mod=stats` → `admin_stats.so` çalışsın
* `/admin/plugin?mod=users` → `admin_users.so` çalışsın
* Gelecekte yeni plugin eklemek için sadece `.so` dosyası eklemen yeterli olsun

---

# 📦 Çoklu Plugin Yönetimi

## 1. Kullanıcı Plugin’i (`admin_users.go`)
*/
``go
package main

import "fmt"

// Admin paneline eklenecek kullanıcı listesi fonksiyonu
func ShowUsers() {
	fmt.Println("Kullanıcı Listesi:")
	fmt.Println("- Abdullah")
	fmt.Println("- Ayşe")
	fmt.Println("- Mehmet")
}
``

//Bunu derle:

``bash
go build -buildmode=plugin -o admin_users.so admin_users.go
``
/*
---

## 2. Güncellenmiş Ana Kod (`main.go`)
*/
``go
package main

import (
	"fmt"
	"net/http"
	"plugin"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	// Ana sayfa
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintln(w, "Blog Ana Sayfa")
	})

	// Dinamik plugin çağırma
	router.GET("/admin/plugin", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// URL parametresi (örn: ?mod=stats)
		mod := r.URL.Query().Get("mod")
		if mod == "" {
			http.Error(w, "Lütfen ?mod= parametresini girin (ör: stats, users)", http.StatusBadRequest)
			return
		}

		// Plugin dosya adı
		file := fmt.Sprintf("./admin_%s.so", mod)

		// Plugin aç
		p, err := plugin.Open(file)
		if err != nil {
			http.Error(w, "Plugin yüklenemedi: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Fonksiyon adı sabit olsun: ShowStats veya ShowUsers değil → "Run"
		sym, err := p.Lookup("Run")
		if err != nil {
			http.Error(w, "Run fonksiyonu bulunamadı: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Type assertion
		runFunc := sym.(func(http.ResponseWriter))

		// Plugin çalıştır
		fmt.Fprintf(w, "Plugin (%s) çalıştırılıyor...\n\n", mod)
		runFunc(w)
	})

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", router)
}
``
/*
---

## 3. Plugin’leri Düzenleme

Bütün pluginlerde **tek tip bir fonksiyon** olsun (`Run`).
Örneğin:

### `admin_stats.go` (güncellenmiş)
*/
``go
package main

import (
	"fmt"
	"net/http"
)

func Run(w http.ResponseWriter) {
	fmt.Fprintln(w, "Blog istatistikleri:")
	fmt.Fprintln(w, "- Toplam Kullanıcı: 128")
	fmt.Fprintln(w, "- Toplam Yazı: 52")
	fmt.Fprintln(w, "- Toplam Kategori: 7")
}
``

//Derle:

``bash
go build -buildmode=plugin -o admin_stats.so admin_stats.go
``
/*
---

### `admin_users.go` (güncellenmiş)
*/
``go
package main

import (
	"fmt"
	"net/http"
)

func Run(w http.ResponseWriter) {
	fmt.Fprintln(w, "Kullanıcı Listesi:")
	fmt.Fprintln(w, "- Abdullah")
	fmt.Fprintln(w, "- Ayşe")
	fmt.Fprintln(w, "- Mehmet")
}
``

//Derle:

``bash
go build -buildmode=plugin -o admin_users.so admin_users.go
``
/*
---

## 4. Çalıştırma

1. Pluginleri derle:
*/
   ``bash
   go build -buildmode=plugin -o admin_stats.so admin_stats.go
   go build -buildmode=plugin -o admin_users.so admin_users.go
   ``

//2. Server başlat:

   ``bash
   go run main.go
   ``
/*
3. Tarayıcıdan test et:

* [http://localhost:8080/admin/plugin?mod=stats](http://localhost:8080/admin/plugin?mod=stats)
*/
  ``
  Plugin (stats) çalıştırılıyor...

  Blog istatistikleri:
  - Toplam Kullanıcı: 128
  - Toplam Yazı: 52
  - Toplam Kategori: 7
  ``
/*
* [http://localhost:8080/admin/plugin?mod=users](http://localhost:8080/admin/plugin?mod=users)
*/
  ``
  Plugin (users) çalıştırılıyor...

  Kullanıcı Listesi:
  - Abdullah
  - Ayşe
  - Mehmet
  ``
/*
---

## 5. Avantajı

* Yeni bir plugin (`admin_logs.go`, `admin_settings.go`) eklemek için:
  → sadece `Run(w http.ResponseWriter)` fonksiyonlu bir Go dosyası yaz,
  → `.so` dosyası olarak derle,
  → parametre ile çağır.

📌 **Sonuç:** Blog’un admin paneli artık **tamamen plugin tabanlı** hale geldi 🚀

---

İstersen ben sana şimdi bu yapıyı **Docker Compose ile entegre edilmiş şekilde** (yani plugin dosyalarını container’a mount ederek) gösterebilirim. Böylece blog container’ını yeniden build etmeden yeni admin modülleri ekleyebilirsin.

Bunu da ister misin?
EVET
*/
/*
Süper 👍 O zaman şimdi sana **Docker Compose + Plugin yönetimi** entegrasyonunu göstereyim.
Böylece blog container’ını tekrar build etmeden, sadece yeni `.so` dosyası kopyalayarak admin modülleri ekleyebileceksin.

---

# 📦 Go Blog + Plugin + Docker Compose

## 1. Proje Yapısı
*/
``bash
goweb/
│── main.go
│── admin_stats.go
│── admin_users.go
│── Dockerfile
│── docker-compose.yml
│── plugins/            # plugin .so dosyaları buraya
``

---

//## 2. Dockerfile (Multi-stage)

``dockerfile
# ---- Builder ----
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Uygulama kaynaklarını kopyala
COPY . .

# Ana binary üret
RUN go build -o goweb main.go

# ---- Final ----
FROM alpine:3.12

WORKDIR /app
COPY --from=builder /app/goweb .

# Plugins klasörünü mount edeceğiz (docker-compose ile)
VOLUME ["/app/plugins"]

EXPOSE 8080
CMD ["./goweb"]
``

---

//## 3. docker-compose.yml

``yaml
version: "3.8"

services:
  blog:
    build: .
    container_name: goweb-blog
    ports:
      - "8080:8080"
    volumes:
      - ./plugins:/app/plugins   # plugin dosyalarını mount et
``
/*
---

## 4. main.go (Plugin Destekli)
*/
``go
package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"plugin"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	// Ana sayfa
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintln(w, "Go Blog Ana Sayfa")
	})

	// Plugin çağırma
	router.GET("/admin/plugin", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		mod := r.URL.Query().Get("mod")
		if mod == "" {
			http.Error(w, "Lütfen ?mod= parametresini girin (ör: stats, users)", http.StatusBadRequest)
			return
		}

		// Plugin dosya yolu
		file := filepath.Join("./plugins", fmt.Sprintf("admin_%s.so", mod))

		// Plugin yükle
		p, err := plugin.Open(file)
		if err != nil {
			http.Error(w, "Plugin yüklenemedi: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Run fonksiyonunu bul
		sym, err := p.Lookup("Run")
		if err != nil {
			http.Error(w, "Run fonksiyonu bulunamadı: "+err.Error(), http.StatusInternalServerError)
			return
		}

		run := sym.(func(http.ResponseWriter))
		fmt.Fprintf(w, "Plugin (%s) çalıştırılıyor...\n\n", mod)
		run(w)
	})

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", router)
}
``

---

## 5. Plugin Dosyaları

### admin\_stats.go

``go
package main

import (
	"fmt"
	"net/http"
)

func Run(w http.ResponseWriter) {
	fmt.Fprintln(w, "Blog istatistikleri:")
	fmt.Fprintln(w, "- Kullanıcı: 128")
	fmt.Fprintln(w, "- Yazı: 52")
	fmt.Fprintln(w, "- Kategori: 7")
}
``

//### admin\_users.go

``go
package main

import (
	"fmt"
	"net/http"
)

func Run(w http.ResponseWriter) {
	fmt.Fprintln(w, "Kullanıcı Listesi:")
	fmt.Fprintln(w, "- Abdullah")
	fmt.Fprintln(w, "- Ayşe")
	fmt.Fprintln(w, "- Mehmet")
}
``
/*
---

## 6. Plugin Derleme

Container’a girmeden önce pluginleri **lokalde derleyip `plugins/` içine koy**:
*/
``bash
go build -buildmode=plugin -o plugins/admin_stats.so admin_stats.go
go build -buildmode=plugin -o plugins/admin_users.so admin_users.go
``
/*
---

## 7. Çalıştırma
*/
``bash
docker-compose up --build
``
/*
Tarayıcıdan test et:

* [http://localhost:8080/admin/plugin?mod=stats](http://localhost:8080/admin/plugin?mod=stats)
* [http://localhost:8080/admin/plugin?mod=users](http://localhost:8080/admin/plugin?mod=users)

---

## 8. Avantaj

* Ana `goweb` container **tekrar build edilmeden** yeni plugin eklenebilir.
* Sadece `.so` dosyasını `plugins/` klasörüne koyman yeterli.
* Bu sayede **admin paneli dinamik olarak genişletilebilir**. 🚀

---

👉 İstersen ben sana şimdi bunun üstüne **plugin hot-reload (çalışırken yeniden yükleme)** mekanizması da gösterebilirim.
Yani `.so` dosyasını güncellediğinde, server yeniden başlatılmadan değişiklik yansısın.

Onu da ister misin?
*/