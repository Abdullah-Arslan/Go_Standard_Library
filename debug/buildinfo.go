/*
Go’nun standart kütüphanesinde **`runtime/debug`** altında bulunan **`buildinfo`** yapısı ve fonksiyonları vardır.
Ayrıca **Go 1.18** ile beraber gelen **`debug/buildinfo`** paketi (ayrı) vardır.

Bu paket, **Go derleyicisinin binary içine gömdüğü derleme bilgilerini** okumaya yarar.
Örneğin:

* Go versiyonu
* Modül bilgileri (go.mod’dan)
* Derlendiği path
* VCS (git) bilgileri (branch, commit, zaman)

Bunu hem kendi binary içinde (`debug.ReadBuildInfo`) hem de dışarıdaki başka bir binary’de (`debug/buildinfo` paketi) kullanabilirsin.

---

# 1️⃣ `runtime/debug.ReadBuildInfo`

Bu, **çalışan programın kendi derleme bilgilerini** verir.

**Örnek: kendi build info’yu okuma**
*/
``go
package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("Build info bulunamadı")
		return
	}

	fmt.Println("Go Versiyonu:", info.GoVersion)
	fmt.Println("Path:", info.Path)

	fmt.Println("Ana Modül:", info.Main.Path, info.Main.Version)
	fmt.Println("Bağımlılıklar:")
	for _, dep := range info.Deps {
		fmt.Printf("  %s %s\n", dep.Path, dep.Version)
	}
}
``
/*
👉 Bunu çalıştırdığında senin projenin **go.mod** içindeki bağımlılıkları ve Go sürümü listelenir.

---

# 2️⃣ `debug/buildinfo` Paketi

Bu paket ise **herhangi bir Go binary’sini açıp build info’yu okumaya** yarar.
Yani `/bin/ls` gibi **Go ile derlenmiş bir programın** build bilgilerini görebilirsin.

**Örnek: dışarıdaki binary’den build info okuma**
*/
``go
package main

import (
	"debug/buildinfo"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Kullanım: go run main.go <go-binary>")
		return
	}

	file := os.Args[1]
	info, err := buildinfo.ReadFile(file)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	fmt.Println("Go Versiyonu:", info.GoVersion)
	fmt.Println("Path:", info.Path)

	fmt.Println("Ana Modül:", info.Main.Path, info.Main.Version)
	fmt.Println("Bağımlılıklar:")
	for _, dep := range info.Deps {
		fmt.Printf("  %s %s\n", dep.Path, dep.Version)
	}
}
``

//👉 Bu örneği çalıştırıp kendi derlediğin Go binary’ini parametre olarak verirsen build info’yu göreceksin:

``bash
go build -o app main.go
go run reader.go ./app
``
/*
---

# 3️⃣ Build Info İçinde Neler Olur?

Go binary’sine gömülen build metadata şunları içerir:

* **Go Versiyonu** → binary’nin hangi Go sürümü ile derlendiği.
* **Path** → ana modülün import yolu.
* **Main Modül** → uygulamanın `go.mod` içindeki ana modülü ve versiyonu.
* **Dependencies** → bağımlı modüller ve versiyonları.
* **Settings** → Git commit, VCS bilgisi, `-ldflags` parametreleri, mod cache bilgileri.

Örnek çıktı:

``
Go Versiyonu: go1.22.1
Path: example.com/myapp
Ana Modül: example.com/myapp v0.0.0
Bağımlılıklar:
  github.com/gorilla/mux v1.8.0
  golang.org/x/sys v0.10.0
``

---

# 4️⃣ Özet

* `runtime/debug.ReadBuildInfo()` → **çalışan programın kendi build info’sunu** verir.
* `debug/buildinfo.ReadFile()` → **başka bir Go binary’sinin build info’sunu** okur.
* Kullanım alanı:

  * Versiyon bilgisini CLI uygulamalarında `--version` komutuyla göstermek
  * Deployment sırasında hangi commit/branch ile build edildiğini öğrenmek
  * Go binary’lerinin modül bağımlılıklarını kontrol etmek

---

👉 İstersen ben sana hem `runtime/debug` hem de `debug/buildinfo` kullanan **tek bir proje** hazırlayayım:

* `go run main.go` → kendi build info’sunu göstersin.
* `go run main.go ./app` → dışarıdaki bir binary’yi analiz etsin.

Bunu ister misin?
EVET
*/
/*
Süper 👍 o zaman sana **tek bir Go projesi** hazırlayayım.
Bu proje hem:

* **kendi çalıştığı programın build info’sunu** (`runtime/debug.ReadBuildInfo`)
* hem de **dışarıdaki başka bir Go binary’sini** (`debug/buildinfo.ReadFile`)

okuyabilecek.

---

## 📂 Proje Yapısı

```
buildinfo-inspector/
 ├── main.go
```

---

## 📌 `main.go`
*/
``go
package main

import (
	"debug/buildinfo"
	"fmt"
	"os"
	"runtime/debug"
)

func main() {
	// Eğer argüman verilmezse kendi build info'yu oku
	if len(os.Args) < 2 {
		fmt.Println("==== Çalışan Programın Build Info'su ====")
		selfInfo()
		return
	}

	// Eğer argüman verilirse dışarıdaki binary'yi oku
	file := os.Args[1]
	fmt.Println("==== Dışarıdaki Binary Build Info ====")
	readOtherBinary(file)
}

// Çalışan programın build info'sunu oku
func selfInfo() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("Build info bulunamadı")
		return
	}

	fmt.Println("Go Versiyonu:", info.GoVersion)
	fmt.Println("Path:", info.Path)
	fmt.Println("Ana Modül:", info.Main.Path, info.Main.Version)

	fmt.Println("Bağımlılıklar:")
	for _, dep := range info.Deps {
		fmt.Printf("  %s %s\n", dep.Path, dep.Version)
	}

	fmt.Println("Ayarlar (Settings):")
	for _, s := range info.Settings {
		fmt.Printf("  %s = %s\n", s.Key, s.Value)
	}
}

// Dışarıdaki binary'den build info oku
func readOtherBinary(file string) {
	info, err := buildinfo.ReadFile(file)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	fmt.Println("Go Versiyonu:", info.GoVersion)
	fmt.Println("Path:", info.Path)
	fmt.Println("Ana Modül:", info.Main.Path, info.Main.Version)

	fmt.Println("Bağımlılıklar:")
	for _, dep := range info.Deps {
		fmt.Printf("  %s %s\n", dep.Path, dep.Version)
	}

	fmt.Println("Ayarlar (Settings):")
	for _, s := range info.Settings {
		fmt.Printf("  %s = %s\n", s.Key, s.Value)
	}
}
``
/*
---

## 🚀 Kullanım

### 1. Kendi build info’sunu görmek
*/
``bash
go run main.go
``
/*
Örnek çıktı:

`
==== Çalışan Programın Build Info'su ====
Go Versiyonu: go1.22.1
Path: example.com/myapp
Ana Modül: example.com/myapp v0.0.0
Bağımlılıklar:
  github.com/gorilla/mux v1.8.0
Ayarlar (Settings):
  vcs = git
  vcs.revision = a1b2c3d4
  vcs.time = 2025-08-27T09:00:00Z
`


---

### 2. Başka bir binary’nin build info’sunu görmek

Önce kendi binary’ni build et:
*/
``bash
go build -o app main.go
``

//Sonra başka bir programdan oku:

``bash
go run main.go ./app
``
/*
---

✅ Böylece hem **kendi derleme bilgini** hem de **dışarıdaki Go binary’lerinin build info’sunu** tek projede okuyabilirsin.

---

👉 İstersen ben buna **Dockerfile** da ekleyeyim, böylece Go yüklü olmadan bile build info okuyabilirsin. Ekleyeyim mi?
EVET
*/

/*
Süper 👍 o zaman bu **buildinfo-inspector** projesine bir **Dockerfile** ekleyelim.
Böylece senin makinede Go kurulu olmasa bile binary’lerin build info’sunu okuyabilirsin.

---

## 📂 Proje Yapısı

```
buildinfo-inspector/
 ├── main.go
 ├── Dockerfile
``

---

## 📌 `Dockerfile`
*/

``dockerfile
# 1. Builder aşaması: Go ile derle
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Kodları kopyala
COPY . .

# Eğer go.mod yoksa otomatik oluştur
RUN go mod init buildinfo-inspector || true

# Binary'yi üret
RUN go build -o inspector main.go

# 2. Final aşaması: Küçük Alpine imajı
FROM alpine:latest

WORKDIR /root/

# Binary'yi kopyala
COPY --from=builder /app/inspector .

# Varsayılan komut
ENTRYPOINT ["./inspector"]
``
/*
---

## 🚀 Kullanım

### 1. Docker imajını build et
*/

``bash
docker build -t buildinfo-inspector .
``

//### 2. Kendi build info’nu görmek

``bash
docker run --rm buildinfo-inspector
``

/*
### 3. Başka bir Go binary’sini incelemek

Örnek: host’taki `./app` binary’sini container’a mount edip okutalım:
*/
``bash
docker run --rm -v $(pwd)/app:/app buildinfo-inspector /app
``

//Windows’ta:

``powershell
docker run --rm -v ${PWD}\app:/app buildinfo-inspector /app
``
/*
---

## 📌 Özet

* Bu Dockerfile sayesinde senin bilgisayarda Go kurulu olmasa bile `buildinfo` okuyabilirsin.
* Hem kendi binary’ni hem de başka Go binary’lerini analiz edebilirsin.
* Küçük boyutlu (Alpine tabanlı) bir imaj oluşturduk.

---

👉 İstersen ben sana bu Docker imajına ek olarak bir **multi-platform build (Linux, Windows, macOS)** ayarı da ekleyebilirim, böylece her işletim sistemi için binary üretebiliriz. Bunu ister misin?
*/
