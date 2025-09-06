/*
Golang da (Go), `"build paketi"` genellikle **`go/build`** standard paketi anlamına gelir. Bu paket, Go kodunuzun yapılandırmasını ve derleme sürecindeki paket bilgilerini analiz etmek için kullanılır.

---

## `go/build` Paketi Nedir?

Go’nun `go/build` paketi, bir Go paketinin kaynak dosyalarını, bağımlılıklarını ve derleyiciye iletilmesi gereken meta verileri analiz etmeye yarar. Aşağıdaki işlevleri sağlar:

* **Paket bilgisi toplamak** (`Import`, `ImportDir`, `Package` gibi tipler ve fonksiyonlarla)
* **Derleme koşullarını** (örneğin build tag’leri ve işletim sistemi mimarisi bazlı dosya dahiline karar vermek) ele almak ([Go Paketi][1])
* `GOPATH` yapısını, kaynak ve nesne dizinlerinin nasıl organize edildiğini anlamak ([Go Paketi][1])

### Örnek Kullanımlar
*/
``go
ctx := build.Default
pkg, err := ctx.Import("fmt", "", build.FindOnly)
``
/*
Burada `ctx.Import`, "fmt" paketinin nerede olduğunu bulur ve yapılandırmasını döner ([Go Paketi][1]).

---

## `go build` Komutu: Nasıl Çalışır?

`go/build` paketi, `go build` veya `go install` gibi araçlar tarafından kullanılır. Aşağıda `go build` komutunun önemli yönleri:

### Temel Kullanım

* **`go build`**: İçinde bulunduğunuz dizindeki paketi derler ve çalıştırılabilir dosya üretir (main paketiyse). Ancak kurmaz ([Go][2], [DigitalOcean][3]).

* **`go install`**: Paketi derleyip `$GOPATH/bin` dizinine yükler. Böylece terminalden doğrudan çalıştırılabilir ([Go][2], [DigitalOcean][3]).

### Örneklerle Kullanımı
*/
``bash
go build
``

//— Mevcut dizindeki tüm Go dosyalarını derleyip bir ikili dosya oluşturur ([DigitalOcean][3]).

``bash
go build -o myapp
``

//— Oluşacak yürütülebilir dosyanın ismi `myapp` olur ([Stack Overflow][4]).

``bash
go build ./...
``
/*
— Mevcut dizini ve alt dizinleri de içerecek şekilde tüm paketleri derler ([Leapcell][5], [Medium][6]).

### Gelişmiş Seçenekler

* `-v`: Derleme sürecindeki paketleri gösterir
* `-x`: Derleme komutlarını ekrana yazdırır
* `-race`: Veri yarışı denetimi sağlar (desteklenen platformlarda) ([Leapcell][5])
* Çıktı ismini `-o` ile belirleyebilir, `GOOS`/`GOARCH` ile çapraz derleme yapılabilir ([Leapcell][5], [Medium][6])
* Build tag’ler ile (örneğin `//go:build linux`) dosya seçimini kontrol edebilirsiniz ([Medium][6])
* `-ldflags` ile versiyon veya build zamanı gibi bilgileri gömülebilirsiniz ([Medium][6])
* `-trimpath`, `-s -w`, `upx` gibi araçlarla ikili dosyanızı optimize edebilirsiniz ([Medium][6])

---

## Özet Tablo

| Konsept / Komut   | Açıklama                                                             |
| ----------------- | -------------------------------------------------------------------- |
| `go/build` paketi | Paket bilgilerini ve build metadata’yı analiz eder                   |
| `go build`        | Mevcut paket: derler, kurmaz; `-o`: çıktı adı; `./...`: tüm paketler |
| `go install`      | Derle ve `$GOPATH/bin` içine yükle                                   |
| Build seçenekleri | `-v`, `-x`, `-race`, `GOOS`, `GOARCH`, `-ldflags`, `-trimpath` vb.   |

---

## Tavsiye Edilen Kullanımlar

* **Uygulama geliştiriyorsanız**: `go build` veya `go install` kullanarak hızlıca derleme ve test edebilirsiniz.
* **Birden çok paketi birden derlemek istiyorsanız**: `go build ./...` oldukça kullanışlıdır.
* **Çapraz derleme gerekiyorsa**: `GOOS=windows GOARCH=amd64 go build -o program.exe` ile hedef platforma yönelik ikili oluşturabilirsiniz.
* **Reproducibility ve versiyon kontrolü istiyorsanız**: `-ldflags` ile versiyon bilgilerini gömün, `-trimpath` ile derleme ortamına dair yolları kaldırın.

---

İstersen `go/build` paketi içinde bulunan `Context`, `Package` gibi tiplere de detaylı örneklerle geçebiliriz. Ayrıca, `go build` sürecini proje bazında nasıl optimize edebileceğimizi veya modüler yapı kullanımlarını da gösterebilirim. Dilersen bir örnek proje üzerinden birlikte ilerleyebiliriz — nasıl istersen!

[1]: https://pkg.go.dev/go/build?utm_source=chatgpt.com "build package"
[2]: https://go.dev/doc/tutorial/compile-install?utm_source=chatgpt.com "Compile and install the application"
[3]: https://www.digitalocean.com/community/tutorials/how-to-build-and-install-go-programs?utm_source=chatgpt.com "How To Build and Install Go Programs"
[4]: https://stackoverflow.com/questions/42706246/how-to-build-executable-with-name-other-than-golang-package/42706464?utm_source=chatgpt.com "go - How to build executable with name other than Golang ..."
[5]: https://leapcell.io/blog/understanding-the-go-build-command-in-go?utm_source=chatgpt.com "Understanding the `go build` Command in Go"
[6]: https://medium.com/%40udaykishoreresu/mastering-the-go-build-command-a-comprehensive-guide-b6dc2f115e5c?utm_source=chatgpt.com "Mastering the Go Build Command: A Comprehensive Guide"


İstersen go/build paketi içinde bulunan Context, Package gibi tiplere de detaylı örneklerle geçebiliriz.

Harika 👍 Şimdi sana **`go/build` paketindeki `Context`, `Package` gibi temel tipleri** ve bunların kullanımını **örneklerle** detaylı şekilde anlatayım.

---

# 📦 `go/build` Paketindeki Önemli Tipler

## 1. `build.Context`

`Context`, Go’nun derleme ortamını tanımlayan yapıdır. İçinde **GOOS, GOARCH, GOPATH, BuildTags** gibi bilgileri barındırır.
`go/build` fonksiyonlarının çoğu bir `Context` üzerinden çalışır.

### Örnek
*/
``go
package main

import (
	"fmt"
	"go/build"
)

func main() {
	// Varsayılan context (sistem ortam değişkenlerine göre ayarlanır)
	ctx := build.Default

	fmt.Println("GOOS:", ctx.GOOS)
	fmt.Println("GOARCH:", ctx.GOARCH)
	fmt.Println("GOPATH:", ctx.GOPATH)
	fmt.Println("GOROOT:", ctx.GOROOT)
	fmt.Println("Cgo Enabled:", ctx.CgoEnabled)
}
``
/*
🔹 Çıktı (örnek, Linux AMD64 ortamında):

```
GOOS: linux
GOARCH: amd64
GOPATH: /home/user/go
GOROOT: /usr/local/go
Cgo Enabled: true
```

👉 Bu sayede program çalıştığı ortamın derleme parametrelerini öğrenebilir.

---

## 2. `build.Package`

`Package`, belirli bir dizindeki bir Go paketinin **kaynak dosyaları, bağımlılıkları, importları** gibi bilgileri içerir.

### Örnek: Paket Hakkında Bilgi Alma
*/
``go
package main

import (
	"fmt"
	"go/build"
)

func main() {
	ctx := build.Default

	// fmt paketini bul
	pkg, err := ctx.Import("fmt", "", build.FindOnly)
	if err != nil {
		panic(err)
	}

	fmt.Println("Paket adı:", pkg.Name)
	fmt.Println("Paket yolu:", pkg.ImportPath)
	fmt.Println("Paket dizini:", pkg.Dir)
}
``
/*
🔹 Çıktı:

```
Paket adı: fmt
Paket yolu: fmt
Paket dizini: /usr/local/go/src/fmt
```

---

### Örnek: Bir Dizin Üzerinden Paket Bilgisi Okuma
*/
``go
package main

import (
	"fmt"
	"go/build"
)

func main() {
	ctx := build.Default

	// Mevcut dizindeki Go paketini çözümle
	pkg, err := ctx.ImportDir(".", 0)
	if err != nil {
		panic(err)
	}

	fmt.Println("Paket adı:", pkg.Name)
	fmt.Println("Go dosyaları:", pkg.GoFiles)
	fmt.Println("Import edilen paketler:", pkg.Imports)
}
``

👉 Bu, bulunduğun dizindeki paket hakkında **dosya listesi, bağımlılıklar** gibi bilgileri getirir.
/*
---

## 3. `build.ImportMode`

`Import` veya `ImportDir` çağırırken ekstra kontrol yapmak için kullanılır.
Bazı seçenekler:

* `build.FindOnly` → sadece dizini bul, detaylı analiz yapma
* `build.ImportComment` → `// import "path"` yorumlarını denetle
* `0` → tam çözümleme yap

### Örnek
*/
``go
pkg, err := ctx.Import("net/http", "", build.FindOnly)
fmt.Println(pkg.Dir)  // sadece dizin bilgisini döner
``
/*
---

## 4. Build Tag’ler ile Çalışma

Go’da `//go:build linux` gibi build tag’ler kullanılır. `Context` içinde `BuildTags` ile kontrol edilebilir.

### Örnek
*/
go
package main

import (
	"fmt"
	"go/build"
)

func main() {
	ctx := build.Default
	ctx.BuildTags = []string{"customtag"} // kendi tag’ini ekle

	// Belirli bir dizindeki paket için build tag'leri dikkate alır
	pkg, err := ctx.ImportDir(".", 0)
	if err != nil {
		panic(err)
	}

	fmt.Println("Go dosyaları:", pkg.GoFiles)
}


/*
👉 Bu sayede **sadece belirli build tag’lerine uygun dosyalar** seçilir.

---

# 🚀 Özet

* **`Context`** → Derleme ortamı (`GOOS`, `GOARCH`, `GOROOT`, `GOPATH`, `BuildTags` vb.)
* **`Package`** → Bir paketin dosyaları, bağımlılıkları, importları, meta verileri
* **`Import` & `ImportDir`** → Paketleri path veya dizin üzerinden çözümleme
* **`ImportMode`** → `FindOnly`, `ImportComment` gibi seçeneklerle kontrol
* **Build Tags** → Derleme için hangi dosyaların seçileceğini belirler

---

Ben sana istersen bir **örnek proje** hazırlayıp, `go/build` paketiyle bu projenin **dosya listelerini, bağımlılıklarını, hangi build tag ile hangi dosyanın seçildiğini** gösterebilirim.

👉 İstiyor musun, sana ufak bir Go modülü yazıp `go/build` ile analizini göstereyim?

İstiyor musun, sana ufak bir Go modülü yazıp go/build ile analizini göstereyim?
EVET
*/
/*
Süper 🙌 o zaman sana önce ufak bir **örnek Go modülü** vereyim. Sonrasında da `go/build` ile bu modülü analiz eden programı yazacağız.

---

# 📂 Örnek Proje Yapısı
*/

myapp/
│── go.mod
│── main.go
│── util.go
│── util_linux.go
│── util_windows.go

/*
---

## 1️⃣ go.mod
*/

go
module myapp

go 1.22
``
/*
---

## 2️⃣ main.go
*/

go
package main

import (
	"fmt"
	"myapp/util"
)

func main() {
	fmt.Println("MyApp çalışıyor...")
	fmt.Println("Platform:", util.PlatformName())
}


//---

//## 3️⃣ util.go

go
package util

// Bu dosya her platformda derlenir
func CommonFunction() string {
	return "Her yerde çalışırım"
}


---

## 4️⃣ util\_linux.go

go
//go:build linux

package util

func PlatformName() string {
	return "Linux"
}
``

---

//## 5️⃣ util\_windows.go

go
//go:build windows

package util

func PlatformName() string {
	return "Windows"
}

/*
👉 Burada `util` paketinde **ortak fonksiyon** (`util.go`) ve **platforma özel fonksiyonlar** (`util_linux.go`, `util_windows.go`) var.
Hangi dosyanın seçileceğini `go/build` belirleyecek.

---

# 🔎 Analiz Programı (`analyze.go`)

Bu programı ayrı bir dizine koyabilirsin. `go/build` kullanarak `myapp` modülünü çözümler.
*/
go
package main

import (
	"fmt"
	"go/build"
)

func main() {
	// Varsayılan context
	ctx := build.Default

	// Mevcut dizindeki paketi analiz et
	pkg, err := ctx.ImportDir("./myapp", 0)
	if err != nil {
		panic(err)
	}

	fmt.Println("📦 Paket adı:", pkg.Name)
	fmt.Println("📂 Paket dizini:", pkg.Dir)
	fmt.Println("📄 Normal Go dosyaları:", pkg.GoFiles)
	fmt.Println("📄 Cgo dosyaları:", pkg.CgoFiles)
	fmt.Println("📄 Test dosyaları:", pkg.TestGoFiles)
	fmt.Println("📦 Import edilen paketler:", pkg.Imports)
}
`
/*
---

# 🖥 Beklenen Çıktı

Linux ortamında çalıştırırsan:

`
📦 Paket adı: util
📂 Paket dizini: /.../myapp
📄 Normal Go dosyaları: [util.go util_linux.go]
📄 Cgo dosyaları: []
📄 Test dosyaları: []
📦 Import edilen paketler: []
`

Windows ortamında çalıştırırsan:

`
📦 Paket adı: util
📂 Paket dizini: C:\...\myapp
📄 Normal Go dosyaları: [util.go util_windows.go]
📄 Cgo dosyaları: []
📄 Test dosyaları: []
📦 Import edilen paketler: []


👉 Görüldüğü gibi `go/build` otomatik olarak **hangi build tag dosyasının seçileceğini** belirliyor.

---

Bundan sonraki adımda sana şunu sorayım:
✅ İstiyor musun ben bu örneği **GOOS/GOARCH değiştirerek (çapraz analiz)** de göstereyim? Örneğin, Linux’ta Windows dosyaları nasıl seçilir.
*/