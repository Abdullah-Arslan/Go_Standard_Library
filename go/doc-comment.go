/*
**`go/doc/comment`** paketini en ince ayrıntısına kadar anlatayım.

---

# 📦 go/doc/comment` Paketi Nedir?

go/doc/comment` paketi, **Go kaynak kodundaki yorumları (comments)** daha yapılandırılmış (structured) bir hale getirmek için kullanılır.

Özellikle:

* `//` veya `/* */` ile yazılmış açıklamaları **parçalar** (cümle, kod bloğu, bağlantı vs.).
* GoDoc ve `go doc` araçlarının **nasıl daha okunabilir dokümantasyon ürettiğini** sağlayan altyapıyı sunar.
* Linkler, kod örnekleri, paragraflar ve listeler gibi yapıları ayırır.

---

# 📌 Temel Bileşenler

go/doc/comment` paketi, metinleri **parse edip formatlayan** bir API sunar.

## 1. comment.Parser

* Bir string (yorum metni) alır.
* İçindeki yapıları (`Doc`, `Block`, `Text`, `Link`, `Heading` vb.) çıkarır.

### Örnek
*/
```go
package main

import (
	"fmt"
	"go/doc/comment"
)

func main() {
	text := `This is a sample function.

It does the following:
- Parse input
- Return result

For more info: https://golang.org


	parser := comment.Parser{}
	doc := parser.Parse(text)

	fmt.Printf("Parsed Doc: %+v\n", doc)
}
``
/*
🔹 Çıktıda `doc.Content` içinde paragraflar, listeler, linkler ayrı ayrı görünecek.

---

## 2. `comment.Doc`

* Parse edilmiş doküman yapısını temsil eder.
* İçinde:

  * **Summary** → İlk cümle (kısa açıklama)
  * **Content** → Tüm doküman (paragraf, kod, liste, link vs.)
  * **Blocks** → İç bloklar

---

## 3. `comment.Block` ve Alt Tipleri

* `Text` → Normal metin
* `Code` → Kod blokları
* `List` → Liste maddeleri (`- item`)
* `Heading` → Başlıklar (`# Title`)
* `Link` → URL veya referans

### Örnek: Kod bloğu parse etme
*/
``go
text := `Example usage:

    fmt.Println("Hello")`

doc := parser.Parse(text)
for _, block := range doc.Content {
    switch b := block.(type) {
    case *comment.Paragraph:
        fmt.Println("Paragraph:", b.Text)
    case *comment.Code:
        fmt.Println("Code Block:", b.Text)
    }
}
``
/*
---

## 4. `comment.Printer`

* Parse edilen `Doc` nesnesini tekrar **string** haline getirir.
* Bu, yorumların biçimli bir şekilde geri yazdırılmasını sağlar.

### Örnek
*/

``go
doc := parser.Parse("This is a comment.\n\nVisit https://golang.org for docs.")
printer := comment.Printer{}
out := printer.Comment(doc)
fmt.Println(out)
``
/*
---

## 5. `comment.Renderer`

* Yorumları **HTML** veya başka formatlara dönüştürmek için kullanılır.
* Örn: GoDoc web sitesinde HTML çıktısı üretir.

### Örnek (HTML render)
*/
``go
renderer := comment.Renderer{}
html := renderer.Render(doc)
fmt.Println(string(html))
``
/*
---

# 🎯 Komple Örnek

Bir fonksiyonun açıklamasını parse edip, hem **metin** hem de **HTML** olarak yazdıralım:
*/
```go
package main

import (
	"fmt"
	"go/doc/comment"
)

func main() {
	src := `PrintHello prints "Hello, world".

Example:

    fmt.Println("Hello, world")

See also: https://golang.org


	// Parse et
	parser := comment.Parser{}
	doc := parser.Parse(src)

	// Normal yazdır
	printer := comment.Printer{}
	fmt.Println("Plain Output:")
	fmt.Println(printer.Comment(doc))

	// HTML yazdır
	renderer := comment.Renderer{}
	fmt.Println("\nHTML Output:")
	fmt.Println(string(renderer.Render(doc)))
}
``
/*
### 📌 Çıktı:

**Plain Output:**
*/
``
PrintHello prints "Hello, world".

Example:

    fmt.Println("Hello, world")

See also: https://golang.org
``

//**HTML Output (kısaltılmış):**

``html
<p>PrintHello prints "Hello, world".</p>
<p>Example:</p>
<pre>fmt.Println("Hello, world")</pre>
<p>See also: <a href="https://golang.org">https://golang.org</a></p>
``
/*
---

# ✅ Özet

* `go/doc/comment` → Go kodundaki yorumları (docstring) **parçalamak, analiz etmek ve yeniden biçimlendirmek** için kullanılır.
* **Ana tipler**:

  * `Parser` → Yorumları parse eder
  * `Doc` → Yorumun yapısı
  * `Printer` → Yorumları düz metin olarak biçimler
  * `Renderer` → HTML gibi formatlara dönüştürür
* Kullanım alanı → `godoc`, `go doc` ve özel dokümantasyon araçları

---

👉 İstersen ben sana `doc/comment` ile kendi **mini-godoc HTML render aracı** yazayım (senin modülündeki yorumları otomatik HTML’e döksün).

Bunu ister misin?
EVET
*/
/*
Süper 👍 O zaman sana **mini-godoc HTML render aracı** yapalım.

Bu araç:

1. Belirli bir `.go` dosyasını açacak.
2. İçindeki yorum satırlarını (`//`, `/* */`) okuyacak.
3. `go/doc/comment` paketi ile parse edecek.
4. Çıktıyı **HTML** olarak yazdıracak.

---

# 📝 Örnek Proje

📂 Proje yapısı:
*/
``
mini-godoc/
 ├── main.go        // aracımız
 └── sample.go      // örnek yorumlar içeren dosya
``
/*
---

## `sample.go` (örnek dosya)
*/
``go
package sample

// PrintHello prints "Hello, world".
//
// Example:
//
//    fmt.Println("Hello, world")
//
// See also: https://golang.org
func PrintHello() {
    println("Hello, world")
}
``
/*
---

## `main.go` (mini-godoc aracı)
*/
``go
package main

import (
	"fmt"
	"go/doc/comment"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	// Dosya yolunu parametre al
	filename := "sample.go"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	// Go dosyasını parse et (sadece yorumları almak için)
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// Tüm yorumları tek string haline getir
	var allComments string
	for _, cgroup := range file.Comments {
		allComments += cgroup.Text() + "\n"
	}

	// comment paketi ile parse et
	parser := comment.Parser{}
	doc := parser.Parse(allComments)

	// HTML render et
	renderer := comment.Renderer{}
	html := renderer.Render(doc)

	fmt.Println("==== HTML Output ====")
	fmt.Println(string(html))
}
``
/*
---

## 🚀 Çalıştırma
*/
``bash
go run main.go sample.go
``

//### Örnek Çıktı:

``html
<p>PrintHello prints "Hello, world".</p>
<p>Example:</p>
<pre>fmt.Println("Hello, world")</pre>
<p>See also: <a href="https://golang.org">https://golang.org</a></p>
``
/*
---

✅ Böylece kendi **mini-godoc** aracımızı yazmış olduk.
İstersen bunu genişletip:

* Paket seviyesindeki yorumları
* Fonksiyon açıklamalarını
* Struct / Method açıklamalarını

da çıkarabiliriz.

👉 İstiyor musun, ben bunu geliştirip tüm **paket dokümantasyonunu HTML olarak** çıkaran bir versiyonunu yazayım?
EVET
*/
/*
Harika 🚀 Şimdi bir üst seviyeye çıkalım ve **go/doc/comment** paketini kullanarak komple bir paket için HTML dokümantasyon çıkaran bir **mini-godoc aracı** yazalım.

---

# 📂 Proje Yapısı

```
mini-godoc/
 ├── main.go       // aracımız
 ├── sample/
 │    └── sample.go
```

---

## `sample/sample.go` (örnek test paketi)
*/
``go
// Package sample küçük bir örnek pakettir.
//
// Bu paket, basit fonksiyon ve tiplerle godoc test etmek için yazılmıştır.
package sample

import "fmt"

// Message varsayılan mesajdır.
const Message = "Hello, Go!"

// PrintMessage ekrana varsayılan mesajı yazar.
func PrintMessage() {
    fmt.Println(Message)
}

// Greeter selam vermek için kullanılan bir yapıdır.
type Greeter struct {
    Name string
}

// Greet verilen isme selam döndürür.
func (g Greeter) Greet() string {
    return "Merhaba, " + g.Name
}
``
/*
---

## `main.go` (mini-godoc)
*/
``go
package main

import (
	"fmt"
	"go/doc"
	"go/doc/comment"
	"go/parser"
	"go/token"
	"html/template"
	"os"
)

const htmlTemplate = `
<!DOCTYPE html>
<html lang="tr">
<head>
<meta charset="UTF-8">
<title>{{.Name}} Paketi Dokümantasyonu</title>
<style>
body { font-family: Arial, sans-serif; margin: 20px; line-height: 1.6; background: #f8f9fa; color: #212529;}
h1 { color: #0d6efd; }
h2 { margin-top: 30px; color: #0d6efd; border-bottom: 2px solid #0d6efd; padding-bottom: 5px; }
pre { background: #212529; color: #f8f9fa; padding: 10px; border-radius: 5px; }
.section { background: #fff; padding: 15px; margin-bottom: 20px; border-radius: 8px; box-shadow: 0 2px 5px rgba(0,0,0,0.1);}
</style>
</head>
<body>
<h1>Paket: {{.Name}}</h1>
<div class="section">{{.DocHTML}}</div>

{{if .Consts}}
<h2>Sabitler</h2>
<div class="section">
{{range .Consts}}
<b>{{.Names}}</b><br>{{.DocHTML}}<br><br>
{{end}}
</div>
{{end}}

{{if .Funcs}}
<h2>Fonksiyonlar</h2>
<div class="section">
{{range .Funcs}}
<b>{{.Name}}</b><br>{{.DocHTML}}<br><br>
{{end}}
</div>
{{end}}

{{if .Types}}
<h2>Tipler</h2>
<div class="section">
{{range .Types}}
<b>{{.Name}}</b><br>{{.DocHTML}}
	{{if .Methods}}
	<ul>
	{{range .Methods}}
	<li><b>{{.Name}}</b>: {{.DocHTML}}</li>
	{{end}}
	</ul>
	{{end}}
<br>
{{end}}
</div>
{{end}}

</body>
</html>


func main() {
	// Paket yolunu al (örnek: ./sample)
	pkgPath := "./sample"
	if len(os.Args) > 1 {
		pkgPath = os.Args[1]
	}

	// Paket dosyalarını parse et
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, pkgPath, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// İlk bulunan paketi al
	var pkg *doc.Package
	for _, astPkg := range pkgs {
		pkg = doc.New(astPkg, "./", doc.AllDecls)
		break
	}

	// comment.Renderer ile HTML üretici
	renderer := comment.Renderer{}

	// Yardımcı struct
	type DocItem struct {
		Name    string
		Names   string
		DocHTML template.HTML
	}

	// Veri modeli
	data := struct {
		Name   string
		DocHTML template.HTML
		Consts []DocItem
		Funcs  []DocItem
		Types  []struct {
			Name    string
			DocHTML template.HTML
			Methods []DocItem
		}
	}{
		Name:    pkg.Name,
		DocHTML: template.HTML(renderer.Render(pkg.Doc)),
	}

	// Sabitler
	for _, c := range pkg.Consts {
		data.Consts = append(data.Consts, DocItem{
			Names:   fmt.Sprint(c.Names),
			DocHTML: template.HTML(renderer.Render(c.Doc)),
		})
	}

	// Fonksiyonlar
	for _, f := range pkg.Funcs {
		data.Funcs = append(data.Funcs, DocItem{
			Name:    f.Name,
			DocHTML: template.HTML(renderer.Render(f.Doc)),
		})
	}

	// Tipler ve metodlar
	for _, t := range pkg.Types {
		typ := struct {
			Name    string
			DocHTML template.HTML
			Methods []DocItem
		}{
			Name:    t.Name,
			DocHTML: template.HTML(renderer.Render(t.Doc)),
		}
		for _, m := range t.Methods {
			typ.Methods = append(typ.Methods, DocItem{
				Name:    m.Name,
				DocHTML: template.HTML(renderer.Render(m.Doc)),
			})
		}
		data.Types = append(data.Types, typ)
	}

	// HTML dosyasını oluştur
	tmpl := template.Must(template.New("doc").Parse(htmlTemplate))
	f, _ := os.Create("doc.html")
	defer f.Close()
	tmpl.Execute(f, data)

	fmt.Println("Dokümantasyon üretildi: doc.html")
}
``
/*
---

# 🚀 Kullanım
*/
``bash
go run main.go ./sample
``
/*
Oluşan `doc.html` dosyasını tarayıcıda açtığında, **paketin tüm sabitleri, fonksiyonları, tipleri ve metodları** renkli, modern bir arayüzle görebileceksin 🎉

---

👉 İstersen bu aracı **multi-package proje** desteğiyle genişletebilirim (birden fazla paket tarayıp her biri için ayrı sayfa veya menü).

Bunu ister misin?
*/