/*
**Go’nun `go/format` paketini** baştan sona, örneklerle açıklayayım.

---

# 📦 `go/format` Paketi

`go/format`, Go kaynak kodunu **standart Go biçimlendirme kurallarına göre** düzenlemek için kullanılan pakettir.
Yani aslında `gofmt` aracının arkasındaki kütüphane budur.

👉 Paket yolu:
*/
``go
import "go/format"
``
/*
---

## 🔑 Önemli Fonksiyonlar

### 1. `format.Source`
*/
``go
func Source(src []byte) ([]byte, error)
``
/*
* Parametre: Go kodu (byte slice olarak).
* Çıktı: Biçimlendirilmiş kod.
* Eğer kod **derlenebilir (parse edilebilir)** değilse hata döner.

📌 Örnek:
*/
``go
package main

import (
	"fmt"
	"go/format"
)

func main() {
	src := []byte(`package main
import "fmt"
func main(){fmt.Println("Merhaba")}`)

	pretty, err := format.Source(src)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(pretty))
}
``

//📤 Çıktı:

``go
package main

import "fmt"

func main() {
	fmt.Println("Merhaba")
}
``
/*
👉 Kod otomatik olarak `gofmt` kurallarına uygun hale getirildi.

---

### 2. `format.Node`
*/
``go
func Node(dst io.Writer, fset *token.FileSet, node ast.Node) error
``
/*
* `ast.Node`: Go kodunu temsil eden AST (Abstract Syntax Tree) düğümü.
* `token.FileSet`: Pozisyon bilgileri (AST için gerekli).
* `dst`: Biçimlendirilmiş çıktının yazılacağı hedef (`os.Stdout`, `bytes.Buffer`, dosya vs.).

📌 Örnek:
*/
``go
package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
)

func main() {
	src := `package main; import "fmt"; func main(){fmt.Println("Selam")}`

	// Parse et
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "example.go", src, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	// AST’yi formatla
	var buf bytes.Buffer
	err = format.Node(&buf, fset, file)
	if err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
}
``

//📤 Çıktı:

``go
package main

import "fmt"

func main() {
	fmt.Println("Selam")
}
``
/*
---

## 🔧 Kullanım Senaryoları

1. **Kod Biçimlendirme Aracı Yazmak**
   `go/format` kullanarak kendine özel bir `gofmt` aracı geliştirebilirsin.
   Örneğin bir editör eklentisi.

2. **AST ile Kod Üretme**

   * `go/ast` + `go/printer` ile kod üretiyorsun.
   * Çıktıyı `go/format` ile düzenlersen, sonuç **temiz ve standart** hale gelir.

3. **Kod Analizi ve Refaktör**

   * `go/parser` ile kodu parse et
   * AST üzerinde değişiklik yap
   * `go/format` ile yeniden yazdır

---

## ⚡ Mini Örnek: Dosya Biçimlendirme
*/
``go
package main

import (
	"fmt"
	"go/format"
	"os"
)

func main() {
	data, err := os.ReadFile("bad.go") // biçimsiz Go dosyası
	if err != nil {
		panic(err)
	}

	pretty, err := format.Source(data)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("bad.go", pretty, 0644) // dosyayı temiz yaz
	if err != nil {
		panic(err)
	}

	fmt.Println("Dosya biçimlendirildi ✔️")
}
``
/*
---

## 📝 Özet

* `go/format` = `gofmt`’in kütüphanesi.
* `Source`: ham Go kodunu biçimlendirir.
* `Node`: AST’yi biçimlendirir ve `io.Writer`’a yazar.
* Kod üretimi, refaktör araçları ve editör eklentileri için çok faydalıdır.

---

👉 İstersen sana **`go/ast` + `go/format` kullanarak bir fonksiyon ekleyen refaktör aracı** da yazabilirim.
Böylece bir `.go` dosyasına otomatik kod ekleyip `go/format` ile temizlenmiş halde kaydederiz.

Bunu ister misin?
EVET
*/
/*
Harika 👍 O zaman sana **`go/ast` + `go/format` kullanarak bir dosyaya otomatik fonksiyon ekleyen refaktör aracı** yazalım 🚀

---

# 🛠️ Refaktör Aracı: Go Dosyasına Fonksiyon Ekleme

📂 Proje yapısı:
*/
``
refactor/
 ├── main.go      // refaktör aracı
 └── sample.go    // üzerinde değişiklik yapılacak dosya
``
/*
---

## `sample.go` (başlangıç dosyası)
*/
``go
package sample

import "fmt"

func Hello() {
    fmt.Println("Hello from original function")
}
``
/*
---

## `main.go` (refaktör aracı)
*/
``go
package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	filename := "sample.go"

	// Dosyayı oku
	src, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// Parse et
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, src, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	// Yeni fonksiyon AST düğümü oluştur
	newFunc := &ast.FuncDecl{
		Name: ast.NewIdent("AddedFunction"), // fonksiyon adı
		Type: &ast.FuncType{                 // dönüş tipi ve parametreler
			Params:  &ast.FieldList{},
			Results: nil,
		},
		Body: &ast.BlockStmt{ // gövde
			List: []ast.Stmt{
				&ast.ExprStmt{
					X: &ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   ast.NewIdent("fmt"),
							Sel: ast.NewIdent("Println"),
						},
						Args: []ast.Expr{
							&ast.BasicLit{
								Kind:  token.STRING,
								Value: `"Hello from AddedFunction"`,
							},
						},
					},
				},
			},
		},
	}

	// Dosya AST’sine yeni fonksiyonu ekle
	file.Decls = append(file.Decls, newFunc)

	// Biçimlendir ve yaz
	var buf bytes.Buffer
	err = format.Node(&buf, fset, file)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filename, buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ sample.go dosyasına AddedFunction eklendi!")
}
``
/*
---

## 🚀 Çalıştırma
*/
``bash
go run main.go
``
/*
---

## `sample.go` (son hali)
*/
``go
package sample

import "fmt"

func Hello() {
	fmt.Println("Hello from original function")
}

func AddedFunction() {
	fmt.Println("Hello from AddedFunction")
}
``
/*
---

✅ Böylece AST üzerinden **fonksiyon ekleyip `go/format` ile temiz yazdırdık**.
Bu yaklaşımı genişletip:

* Yeni **struct ekleme**
* Var olan fonksiyonun içine **statement ekleme**
* Otomatik **paket import etme**

gibi refaktör işlemleri yapılabiliyor.

---

👉 İster misin, bir sonraki adımda bu aracı geliştirip **var olan fonksiyonların içine yeni satır ekleyen bir versiyonunu** yazalım?
*/