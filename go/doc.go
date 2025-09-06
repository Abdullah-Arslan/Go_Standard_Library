/*
**`go/doc` paketini** detaylıca anlatacağım.

---

# 📦 `go/doc` Paketi

`go/doc` paketi, Go kaynak kodunu (AST üzerinden) **insanlara okunabilir şekilde belgeler**.
Aslında bu paket, `godoc` aracının temelini oluşturur.

Kısaca:

* `go/ast` ile parse edilen Go kodunu alır
* İçindeki paketleri, tipleri, fonksiyonları, sabitleri, değişkenleri inceler
* Bunlardan **belgelendirme nesneleri** (`doc.Package`, `doc.Type`, `doc.Func` …) üretir

---

## 🔑 Temel Yapılar

### 1. `doc.New`
*/

``go
func New(pkg *ast.Package, importPath string, mode Mode) *Package
``
/*
Bir `ast.Package`’ten `*doc.Package` üretir.
Yani: önce `go/parser` ile parse ➝ sonra `doc.New` ile belgeleme nesnesi.

---

### 2. `doc.Package`

Bir Go paketini belgeler. İçinde:

* `Name` → paket adı
* `Doc` → paket dokümantasyon yorumu
* `Consts`, `Vars`, `Funcs`, `Types` → paket üyeleri
* `Imports` → import edilen paketler

---

### 3. `doc.Type`

Bir Go tipini belgeler. İçinde:

* `Name` → tip adı
* `Doc` → yorum
* `Methods` → tipin metodları
* `Funcs` → o tipe bağlı fonksiyonlar

---

### 4. `doc.Func`

Bir fonksiyonu belgeler. İçinde:

* `Name` → fonksiyon adı
* `Doc` → yorum

---

### 5. `doc.Value`

Sabit (`const`) ve değişken (`var`) bildirilerini belgeler.

---

### 6. `doc.Examples`

Kaynak kod içindeki `Example` fonksiyonlarını döndürür (`godoc`’ta görülen örnekler için).

---

## 📌 Basit Örnek: Paket Dokümantasyonu

Örnek dosya (`example.go`):
*/
``go
// Package mathutils küçük matematiksel yardımcı fonksiyonlar sağlar.
package mathutils

// Square bir sayının karesini döndürür.
func Square(x int) int {
    return x * x
}

// Cube bir sayının küpünü döndürür.
func Cube(x int) int {
    return x * x * x
}
``

//Analiz programı (`analyze_doc.go`):

``go
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/doc"
	"log"
)

func main() {
	fset := token.NewFileSet()
	// Dosyayı parse et
	file, err := parser.ParseFile(fset, "example.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	// AST'ten paket oluştur
	pkg := &ast.Package{
		Name:  file.Name.Name,
		Files: map[string]*ast.File{"example.go": file},
	}

	// doc.Package oluştur
	docPkg := doc.New(pkg, "./mathutils", 0)

	fmt.Println("📦 Paket:", docPkg.Name)
	fmt.Println("📖 Açıklama:", docPkg.Doc)

	fmt.Println("\nFonksiyonlar:")
	for _, f := range docPkg.Funcs {
		fmt.Printf("  %s → %s\n", f.Name, f.Doc)
	}
}
``

//Çıktı:

``
📦 Paket: mathutils
📖 Açıklama: Package mathutils küçük matematiksel yardımcı fonksiyonlar sağlar.

Fonksiyonlar:
  Square → Square bir sayının karesini döndürür.
  Cube → Cube bir sayının küpünü döndürür.
``
/*
---

## 📌 Örnek: Tip & Metot Dokümantasyonu
*/
``go
package shapes

// Shape tüm şekiller için temel arayüzdür.
type Shape interface {
    Area() float64
}

// Rectangle dikdörtgeni temsil eder.
type Rectangle struct {
    Width, Height float64
}

// Area dikdörtgenin alanını döndürür.
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
``

//Analiz:

``go
docPkg := doc.New(pkg, "./shapes", 0)

fmt.Println("Tipler:")
for _, t := range docPkg.Types {
    fmt.Printf("Type %s → %s\n", t.Name, t.Doc)
    for _, m := range t.Methods {
        fmt.Printf("  Method %s → %s\n", m.Name, m.Doc)
    }
}
``

//Çıktı:

``
Tipler:
Type Shape → Shape tüm şekiller için temel arayüzdür.
Type Rectangle → Rectangle dikdörtgeni temsil eder.
  Method Area → Area dikdörtgenin alanını döndürür.
``
/*
---

## 📌 `doc.Examples`

Kaynak dosyada:
*/
``go
// ExampleSquare Square fonksiyonunun nasıl çalıştığını gösterir.
func ExampleSquare() {
    fmt.Println(Square(5))
    // Output: 25
}
``

//Kodda:

``go
examples := doc.Examples(file)
for _, ex := range examples {
    fmt.Printf("Örnek: %s, açıklama: %s\n", ex.Name, ex.Doc)
}
``

/*
---

# ✅ Özet

* `go/doc`, Go kodunun **AST temelli belgelendirme aracı**dır.
* Temel tipleri: `Package`, `Type`, `Func`, `Value`.
* `doc.New` → `ast.Package`’ten belgeler çıkarır.
* `doc.Examples` → test amaçlı `Example` fonksiyonlarını toplar.
* `godoc` aracı bu paketi kullanır.

---

👉 İstersen ben sana küçük bir **mini-godoc aracı** yazabilirim: Belirttiğin klasörü tarayıp `go/doc` ile dokümantasyon çıkaran. Bunu ister misin?
EVET
*/
/*
Harika 🙌 o zaman sana **mini-godoc aracı** hazırlayalım.

Bu araç:

1. Verilen dizindeki tüm `.go` dosyalarını tarar
2. `go/parser` ile parse eder
3. `go/doc` ile paket dokümantasyonu, tipler, fonksiyonlar ve örnekleri toplar
4. Konsola **okunabilir bir formatta** yazdırır

---

# 📌 Kod: `minigodoc.go`
*/
``go
package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"go/doc"
	"go/ast"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	root := "./myapp" // analizi yapılacak proje dizini

	fmt.Println("🔍 Mini-godoc başlatılıyor...\n")

	pkgFiles := make(map[string]*ast.File)
	fset := token.NewFileSet()

	// 1️⃣ Tüm .go dosyalarını tara
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
			if err != nil {
				fmt.Println("Dosya parse edilemedi:", path, err)
				return nil
			}
			pkgFiles[path] = file
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	if len(pkgFiles) == 0 {
		fmt.Println("Hiç .go dosyası bulunamadı.")
		return
	}

	// 2️⃣ AST -> doc.Package
	astPkg := &ast.Package{
		Name:  "", // paket adı otomatik alınacak
		Files: pkgFiles,
	}

	docPkg := doc.New(astPkg, root, 0)

	// 3️⃣ Paket bilgisi
	fmt.Printf("📦 Paket: %s\n", docPkg.Name)
	fmt.Printf("📖 Açıklama:\n%s\n\n", docPkg.Doc)

	// 4️⃣ Sabitler ve değişkenler
	if len(docPkg.Consts) > 0 {
		fmt.Println("🔹 Sabitler:")
		for _, c := range docPkg.Consts {
			fmt.Printf("  %s → %s\n", namesList(c.Names), c.Doc)
		}
		fmt.Println()
	}

	if len(docPkg.Vars) > 0 {
		fmt.Println("🔹 Değişkenler:")
		for _, v := range docPkg.Vars {
			fmt.Printf("  %s → %s\n", namesList(v.Names), v.Doc)
		}
		fmt.Println()
	}

	// 5️⃣ Fonksiyonlar
	if len(docPkg.Funcs) > 0 {
		fmt.Println("🔹 Fonksiyonlar:")
		for _, f := range docPkg.Funcs {
			fmt.Printf("  %s → %s\n", f.Name, f.Doc)
		}
		fmt.Println()
	}

	// 6️⃣ Tipler ve metodlar
	if len(docPkg.Types) > 0 {
		fmt.Println("🔹 Tipler:")
		for _, t := range docPkg.Types {
			fmt.Printf("  Type %s → %s\n", t.Name, t.Doc)
			for _, m := range t.Methods {
				fmt.Printf("    Method %s → %s\n", m.Name, m.Doc)
			}
		}
		fmt.Println()
	}

	// 7️⃣ Örnekler
	if len(docPkg.Examples) > 0 {
		fmt.Println("🔹 Örnekler:")
		for _, ex := range docPkg.Examples {
			fmt.Printf("  Example %s → %s\n", ex.Name, ex.Doc)
		}
		fmt.Println()
	}
}

func namesList(names []string) string {
	return strings.Join(names, ", ")
}
``
/*
---

# ⚙️ Kullanım

1. Örnek proje:
*/
``
myapp/
│── constants.go
│── math.go
│── shapes.go
``

2. Araç dosyasını çalıştır:
*/
``bash
go run minigodoc.go
``
/*
---

# 🖥 Beklenen Çıktı (örnek)
*/
``
🔍 Mini-godoc başlatılıyor...

📦 Paket: myapp
📖 Açıklama:
Package myapp küçük örnek paket

🔹 Sabitler:
  Pi → Pi sayısal sabittir
  Flag → Flag boolean sabittir

🔹 Fonksiyonlar:
  Square → Square bir sayının karesini döndürür.
  Cube → Cube bir sayının küpünü döndürür.

🔹 Tipler:
  Type Rectangle → Dikdörtgen tipi
    Method Area → Dikdörtgenin alanını döndürür.

🔹 Örnekler:
  ExampleSquare → Square fonksiyon örneği
``
/*
---

✅ Bu araç ile artık bir proje klasörünü **mini-godoc** gibi tarayabilir, tüm sabitleri, değişkenleri, tipleri, fonksiyonları ve örnekleri kolayca görebilirsin.

---

İstersen bir sonraki adımda bu aracı **HTML çıktısı üreten** bir hale de getirebilirim, tıpkı gerçek `godoc` gibi.

Bunu yapalım mı?
EVET
*/
/*
Harika 🙌 o zaman mini-godoc aracımızı **HTML çıktısı üreten bir hale** getirelim.

Ama önce önemli noktaları özetleyelim:

* Artık araç, proje dizinini tarayacak
* `go/doc` ile paket, tip, fonksiyon ve sabitleri okuyacak
* HTML formatında çıktı üretecek (basit bir stil ile)
* Sonrasında tarayıcıda açılabilir olacak

---

# 📌 Kod: `minigodoc_html.go`
*/
``go
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/doc"
	"go/token"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type FuncDoc struct {
	Name, Doc string
}

type TypeDoc struct {
	Name    string
	Doc     string
	Methods []FuncDoc
}

type ConstVarDoc struct {
	Names []string
	Doc   string
}

type PackageDoc struct {
	Name     string
	Doc      string
	Consts   []ConstVarDoc
	Vars     []ConstVarDoc
	Funcs    []FuncDoc
	Types    []TypeDoc
	Examples []FuncDoc
}

func main() {
	root := "./myapp" // analiz edilecek proje dizini
	pkgFiles := make(map[string]*ast.File)
	fset := token.NewFileSet()

	// Tüm .go dosyalarını tara
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
			if err != nil {
				fmt.Println("Dosya parse edilemedi:", path, err)
				return nil
			}
			pkgFiles[path] = file
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if len(pkgFiles) == 0 {
		fmt.Println("Hiç .go dosyası bulunamadı.")
		return
	}

	// AST -> doc.Package
	astPkg := &ast.Package{
		Name:  "", // paket adı otomatik alınacak
		Files: pkgFiles,
	}
	docPkg := doc.New(astPkg, root, 0)

	p := PackageDoc{
		Name: docPkg.Name,
		Doc:  docPkg.Doc,
	}

	// Sabitler
	for _, c := range docPkg.Consts {
		p.Consts = append(p.Consts, ConstVarDoc{Names: c.Names, Doc: c.Doc})
	}

	// Değişkenler
	for _, v := range docPkg.Vars {
		p.Vars = append(p.Vars, ConstVarDoc{Names: v.Names, Doc: v.Doc})
	}

	// Fonksiyonlar
	for _, f := range docPkg.Funcs {
		p.Funcs = append(p.Funcs, FuncDoc{Name: f.Name, Doc: f.Doc})
	}

	// Tipler
	for _, t := range docPkg.Types {
		var methods []FuncDoc
		for _, m := range t.Methods {
			methods = append(methods, FuncDoc{Name: m.Name, Doc: m.Doc})
		}
		p.Types = append(p.Types, TypeDoc{Name: t.Name, Doc: t.Doc, Methods: methods})
	}

	// Örnekler
	for _, ex := range docPkg.Examples {
		p.Examples = append(p.Examples, FuncDoc{Name: ex.Name, Doc: ex.Doc})
	}

	// HTML çıktısı oluştur
	tmpl := template.Must(template.New("doc").Parse(htmlTemplate))
	file, err := os.Create("doc.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = tmpl.Execute(file, p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ HTML dokümantasyon oluşturuldu: doc.html")
}

const htmlTemplate = `
<!DOCTYPE html>
<html lang="tr">
<head>
<meta charset="UTF-8">
<title>Mini GoDoc - {{.Name}}</title>
<style>
body { font-family: sans-serif; margin: 20px; }
h1 { color: #2c3e50; }
h2 { color: #34495e; margin-top: 20px; }
pre { background: #ecf0f1; padding: 10px; }
</style>
</head>
<body>
<h1>Paket: {{.Name}}</h1>
<p>{{.Doc}}</p>

{{if .Consts}}
<h2>Sabitler</h2>
<ul>
{{range .Consts}}
<li><b>{{join .Names ", "}}</b>: {{.Doc}}</li>
{{end}}
</ul>
{{end}}

{{if .Vars}}
<h2>Değişkenler</h2>
<ul>
{{range .Vars}}
<li><b>{{join .Names ", "}}</b>: {{.Doc}}</li>
{{end}}
</ul>
{{end}}

{{if .Funcs}}
<h2>Fonksiyonlar</h2>
<ul>
{{range .Funcs}}
<li><b>{{.Name}}</b>: {{.Doc}}</li>
{{end}}
</ul>
{{end}}

{{if .Types}}
<h2>Tipler</h2>
<ul>
{{range .Types}}
<li><b>{{.Name}}</b>: {{.Doc}}
	{{if .Methods}}
	<ul>
	{{range .Methods}}
	<li><b>{{.Name}}</b>: {{.Doc}}</li>
	{{end}}
	</ul>
	{{end}}
</li>
{{end}}
</ul>
{{end}}

{{if .Examples}}
<h2>Örnekler</h2>
<ul>
{{range .Examples}}
<li><b>{{.Name}}</b>: {{.Doc}}</li>
{{end}}
</ul>
{{end}}

</body>
</html>


func join(list []string, sep string) string {
	return strings.Join(list, sep)
}

func init() {
	template.FuncMap{"join": join}
}
``
/*
---

# ⚙️ Kullanım

1. Proje klasörünü `./myapp` olarak hazırla.
2. `minigodoc_html.go` dosyasını çalıştır:
*/
``bash
go run minigodoc_html.go
``
/*
3. Çıktı: `doc.html` dosyası oluşur.
4. Tarayıcıda açarak tüm paket dokümantasyonunu görebilirsin.

---

✅ Böylece artık mini-godoc aracımız **HTML formatında çıktı** üretiyor.

* Paket açıklaması
* Sabitler & değişkenler
* Fonksiyonlar
* Tipler ve metodlar
* Örnekler

hepsini şık bir şekilde görebiliyorsun.

---

İstersen bir sonraki adımda bunu **CSS ve renklerle daha profesyonel görünüm**e taşıyabiliriz, tıpkı gerçek `godoc` gibi.

Bunu yapalım mı?
EVET
*/
/*
Harika 🙌 o zaman mini-godoc aracımızı **daha profesyonel bir görünüm**e taşıyalım.

Amaç:

* HTML çıktısını **modern ve okunabilir** yapalım
* Başlıklar, kod blokları, sabit/fonksiyon tipleri için **renkler** kullanalım
* `CSS` ile stil ekleyelim, tıpkı gerçek `godoc` gibi

---

# 📌 Güncellenmiş HTML Template (`minigodoc_html_styled.go`)
*/
```go
const htmlTemplate = `
<!DOCTYPE html>
<html lang="tr">
<head>
<meta charset="UTF-8">
<title>Mini GoDoc - {{.Name}}</title>
<style>
body { font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif; margin: 20px; line-height: 1.6; color: #333; background: #f9f9f9;}
h1 { color: #2c3e50; border-bottom: 2px solid #2c3e50; padding-bottom: 5px; }
h2 { color: #34495e; margin-top: 25px; border-bottom: 1px solid #bdc3c7; padding-bottom: 3px;}
h3 { color: #16a085; margin-top: 15px; }
ul { list-style: none; padding-left: 0; }
li { margin: 5px 0; }
b { color: #2980b9; }
pre { background: #2c3e50; color: #ecf0f1; padding: 10px; border-radius: 5px; overflow-x: auto;}
code { background: #ecf0f1; padding: 2px 4px; border-radius: 3px; color: #c0392b;}
.section { background: #fff; padding: 15px; border-radius: 5px; box-shadow: 0px 2px 5px rgba(0,0,0,0.1); margin-bottom: 20px;}
</style>
</head>
<body>

<h1>Paket: {{.Name}}</h1>
<div class="section">
<p>{{.Doc}}</p>
</div>

{{if .Consts}}
<h2>Sabitler</h2>
<div class="section">
<ul>
{{range .Consts}}
<li><b>{{join .Names ", "}}</b>: {{.Doc}}</li>
{{end}}
</ul>
</div>
{{end}}

{{if .Vars}}
<h2>Değişkenler</h2>
<div class="section">
<ul>
{{range .Vars}}
<li><b>{{join .Names ", "}}</b>: {{.Doc}}</li>
{{end}}
</ul>
</div>
{{end}}

{{if .Funcs}}
<h2>Fonksiyonlar</h2>
<div class="section">
<ul>
{{range .Funcs}}
<li><b>{{.Name}}</b>: {{.Doc}}</li>
{{end}}
</ul>
</div>
{{end}}

{{if .Types}}
<h2>Tipler</h2>
<div class="section">
<ul>
{{range .Types}}
<li><b>{{.Name}}</b>: {{.Doc}}
	{{if .Methods}}
	<ul>
	{{range .Methods}}
	<li><b>{{.Name}}</b>: {{.Doc}}</li>
	{{end}}
	</ul>
	{{end}}
</li>
{{end}}
</ul>
</div>
{{end}}

{{if .Examples}}
<h2>Örnekler</h2>
<div class="section">
<ul>
{{range .Examples}}
<li><b>{{.Name}}</b>: <pre>{{.Doc}}</pre></li>
{{end}}
</ul>
</div>
{{end}}

</body>
</html>
`
```
/*
---

# 🔹 Yeni Özellikler

1. **Modern font ve renkler**
2. Her bölüm (`Consts`, `Vars`, `Funcs`, `Types`, `Examples`) **kutular içinde**
3. Kod blokları `pre` ile renkli ve kaydırılabilir
4. Tipler, fonksiyonlar ve sabitler **renkli başlıklar**
5. Genel olarak **godoc tarzı görünüm**

---

# ⚙️ Kullanım

1. `minigodoc_html_styled.go` dosyasını projenin kök dizininde çalıştır:

```bash
go run minigodoc_html_styled.go
```

2. `doc.html` dosyası oluşacak.
3. Tarayıcıda açtığında **güzel, profesyonel bir GoDoc görünümü** göreceksin.

---

İstersen bir sonraki adımda **multi-package projeleri** de destekleyecek şekilde geliştirebiliriz:

* Proje dizinindeki tüm alt paketleri tarayacak
* Her paketi ayrı sekmede veya sayfada gösterecek

Bunu yapalım mı?
*/