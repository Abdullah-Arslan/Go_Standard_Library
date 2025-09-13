/*
Şimdi senin **`goweb` blog projen** üzerinden, statik dosyalar (CSS, JS, images, uploads) için `path/filepath` kullanımını göstereyim. Bu senin daha önce yaşadığın MIME ve `template.ParseFiles` hatalarına da birebir çözüm olacak.

---

# ⚡ Neden `filepath`?

* Projeyi **Linux (Docker)** ve **Windows (senin geliştirme ortamın)** üzerinde çalıştırıyorsun.
* Eğer `path` (Unix-only) kullanırsan, Windows’ta `\` yerine `/` sorun çıkarabilir.
* Bu yüzden **`filepath.Join`** kullanarak platform bağımsız yollar üretmelisin.

---

# 🔹 Örnek: Template Dosyalarını Yükleme

Senin `template.ParseFiles("templates/index.html")` hatası şuradan kaynaklanıyor olabilir:

* Container içinde **çalışma dizini** farklı (`/app`)
* `./templates/...` yolunu bulamıyor

Doğru kullanım 👇
*/
``go
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func main() {
	// Template dosyasının yolunu güvenli oluştur
	tmplPath := filepath.Join("templates", "index.html")

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		panic(fmt.Sprintf("Template bulunamadı: %v", err))
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	fmt.Println("Server çalışıyor: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
``
/*
📌 Burada `filepath.Join("templates", "index.html")` →

* Linux: `templates/index.html`
* Windows: `templates\index.html`

Ama her iki sistemde de **çalışır**. ✅

---

# 🔹 Örnek: Statik Dosya (CSS/JS/Images)

Sen daha önce `httprouter.ServeFiles("/assets/*filepath", http.Dir("./assets"))` yapmıştın.
Onu **`filepath` ile güvenli** hale getirebilirsin:
*/
``go
package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	// Statik dosyaların kök klasörünü ayarla
	assetsPath := filepath.Join(".", "assets")
	router.ServeFiles("/assets/*filepath", http.Dir(assetsPath))

	// Uploads klasörü
	uploadsPath := filepath.Join(".", "uploads")
	router.ServeFiles("/uploads/*filepath", http.Dir(uploadsPath))

	fmt.Println("Server çalışıyor: http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
``
/*
📌 Böylece:

* `/assets/css/style.css` → `./assets/css/style.css` (Linux & Windows uyumlu)
* `/uploads/img.png` → `./uploads/img.png`

Ve artık MIME hataları yaşamazsın çünkü Go otomatik `Content-Type` belirler.

---

# 🔹 Örnek: Config ile Merkezileştirme

Senin projende `config.Routes()` vardı. Orada yolları `filepath` ile çözebilirsin:
*/
``go
package config

import (
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

func Routes() *httprouter.Router {
	router := httprouter.New()

	// Tüm statik yolları tek yerde yönet
	staticDirs := map[string]string{
		"/assets/":  "assets",
		"/uploads/": "uploads",
	}

	for route, dir := range staticDirs {
		dirPath := filepath.Join(".", dir)
		router.ServeFiles(route+"*filepath", http.Dir(dirPath))
	}

	return router
}
``
/*
---

# 🔹 Docker İçinde Yol Çözümü

Senin `Dockerfile` → `/app` içinde çalışıyor.
O yüzden template/assets yollarını **göreli (`./`)** değil, **mutlak (`/app/...`)** çözmek daha sağlam olur.
*/
``go
workDir := "/app"
tmplPath := filepath.Join(workDir, "templates", "index.html")
``

//veya Go içinde **`os.Getwd()`** ile çalışma dizinini alabilirsin:

``go
cwd, _ := os.Getwd()
tmplPath := filepath.Join(cwd, "templates", "index.html")
``
/*
---

# ✅ ÖZET

* **Template hatası** → `filepath.Join` ile yolları oluştur, `os.Getwd()` ile çalışma dizinini doğrula
* **CSS/JS MIME hatası** → `router.ServeFiles("/assets/*filepath", http.Dir(filepath.Join(".", "assets")))`
* **Uploads/images** → aynı mantık
* **Docker uyumluluğu** → `/app` köküne göre yol kur

---

Abdullah, istersen ben sana şimdi **tam `goweb` projesi için `main.go` + `config/routes.go` dosyalarını `filepath` ile düzeltilmiş şekilde hazır** verebilirim. İstiyor musun?
EVET
*/
/*
Harika 👍 O zaman sana `goweb` projen için **tam `main.go` ve `config/routes.go`** dosyalarını `path/filepath` ile düzeltilmiş, Docker ve Windows/Linux uyumlu olacak şekilde hazırlayayım.

---

# 📌 `main.go`
*/
``go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"goweb/config"
)

func main() {
	// Çalışma dizinini bul
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Çalışma dizini alınamadı: %v", err)
	}
	fmt.Println("Çalışma dizini:", cwd)

	// Router'ı yükle
	router := config.Routes()

	// Örnek ana sayfa (template)
	tmplPath := filepath.Join(cwd, "templates", "index.html")
	if _, err := os.Stat(tmplPath); os.IsNotExist(err) {
		log.Fatalf("Template bulunamadı: %s", tmplPath)
	}

	// Server başlat
	fmt.Println("Server çalışıyor: http://0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
``
/*
---

# 📌 `config/routes.go`
*/

``go
package config

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

func Routes() *httprouter.Router {
	router := httprouter.New()

	// Statik dizinler
	staticDirs := map[string]string{
		"/assets/":  "assets",
		"/uploads/": "uploads",
	}

	cwd, _ := os.Getwd()
	for route, dir := range staticDirs {
		dirPath := filepath.Join(cwd, dir)
		router.ServeFiles(route+"*filepath", http.Dir(dirPath))
	}

	// Örnek index route
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		cwd, _ := os.Getwd()
		tmplPath := filepath.Join(cwd, "templates", "index.html")

		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			http.Error(w, "Template yüklenemedi: "+err.Error(), 500)
			return
		}

		tmpl.Execute(w, nil)
	})

	return router
}
``
/*
---

# ✅ Bu Kod Ne Sağlıyor?

1. **Windows & Linux uyumu** → Tüm yollar `filepath.Join` ile oluşturuluyor.
2. **Docker uyumu** → Container içinde `/app/templates/index.html` sorunsuz bulunur.
3. **Statik dosyalar** → `assets` ve `uploads` otomatik serve ediliyor.

   * `/assets/css/style.css` → `./assets/css/style.css`
   * `/uploads/image.png` → `./uploads/image.png`
4. **Template hataları engelleniyor** → `os.Getwd()` + `filepath.Join` ile doğru dosya yolu garanti.

---

👉 Abdullah, şimdi sana önerim:

1. `templates/index.html` dosyanı mutlaka **proje kökünde** `/templates` klasöründe tut.
2. `assets/` (css/js/images) ve `uploads/` klasörlerini **aynı seviyede** koy.
3. `docker-compose` ile mount ederken, proje kökünü `/app` içine kopyala.

---

İstersen ben sana bunun için **Dockerfile ve docker-compose.yml dosyalarını da `filepath` uyumlu olacak şekilde tam sürüm** olarak verebilirim. Hazırlamamı ister misin?
EVET
*/
/*
Süper 👍 O zaman sana hem **Dockerfile** hem de **docker-compose.yml** dosyalarını, `goweb` projen için **tam, çalışır ve `filepath` uyumlu** şekilde hazırlayayım.

Bu sayede:

* Go blog uygulaman (`goweb`) → `:8080`
* MySQL DB → `:3306`
* phpMyAdmin → `:8081`
  tam uyumlu olacak.

---

# 📌 `Dockerfile`
*/

``dockerfile
# ---- Builder Stage ----
FROM golang:1.24-alpine AS builder

WORKDIR /app

# go mod ve bağımlılıkları indir
COPY go.mod go.sum ./
RUN go mod download

# tüm projeyi kopyala
COPY . .

# binary üret
RUN go build -o goweb .

# ---- Final Stage ----
FROM alpine:3.12

WORKDIR /app

# runtime bağımlılıkları
RUN apk add --no-cache ca-certificates tzdata

# builder’dan binary kopyala
COPY --from=builder /app/goweb .
# templates, assets, uploads klasörlerini de kopyala
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/assets ./assets
COPY --from=builder /app/uploads ./uploads

EXPOSE 8080

CMD ["./goweb"]
``
/*
✅ Burada `filepath.Join` kullanımı Go kodunda olduğu için Linux/Windows fark etmez, container içinde `/app/...` altında çalışır.

---

# 📌 `docker-compose.yml`
*/
``yaml
version: "3.9"

services:
  goweb:
    build: .
    container_name: goweb-blog
    ports:
      - "8080:8080"
    depends_on:
      - db
    volumes:
      - ./templates:/app/templates
      - ./assets:/app/assets
      - ./uploads:/app/uploads
    networks:
      - goweb-net

  db:
    image: mysql:9.4
    container_name: goweb-db
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: rootpass
      MYSQL_DATABASE: goweb
      MYSQL_USER: goweb
      MYSQL_PASSWORD: gowebpass
    ports:
      - "3306:3306"
    volumes:
      - db_data:/var/lib/mysql
    networks:
      - goweb-net

  phpmyadmin:
    image: phpmyadmin/phpmyadmin
    container_name: goweb-phpmyadmin
    restart: always
    environment:
      PMA_HOST: db
      PMA_USER: goweb
      PMA_PASSWORD: gowebpass
    ports:
      - "8081:80"
    depends_on:
      - db
    networks:
      - goweb-net

volumes:
  db_data:

networks:
  goweb-net:
    driver: bridge
``
/*
---

# ✅ Bu Yapı Ne Sağlıyor?

1. **Go Web Blog (`goweb`)**

   * `http://localhost:8080`
   * Kod, `templates/`, `assets/`, `uploads/` klasörleriyle beraber mount ediliyor
   * `filepath` sayesinde Windows + Linux uyumlu

2. **MySQL DB**

   * `localhost:3306`
   * kullanıcı: `goweb` | şifre: `gowebpass` | db: `goweb`

3. **phpMyAdmin**

   * `http://localhost:8081`
   * giriş: `goweb / gowebpass`

4. **Volume & Network**

   * `db_data` → kalıcı MySQL verileri
   * `goweb-net` → konteynerler arası bağlantı

---

# 🚀 Çalıştırma Adımları
*/
``sh
docker-compose up --build
``
/*
📌 Sonra:

* `http://localhost:8080` → Blog çalışıyor
* `http://localhost:8081` → phpMyAdmin arayüzü
* `docker logs goweb-blog` → Template & DB hatası olmamalı

---

👉 Abdullah, şimdi sana bonus olarak:
**Go kodunda `.env` dosyası ile DB bilgilerini environment variable’dan okuyan bir yapı** da ekleyebilirim.
Böylece `docker-compose.yml` ile `.env` senkronize çalışır.

İster misin `.env` destekli hale getireyim?
*/