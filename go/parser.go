/*
**Go’nun `go/parser` paketini** detaylıca anlatacağım.

---

# 📌 `go/parser` Paketi Nedir?

`go/parser`, Go kaynak kodunu **parse** (çözümleme) ederek **AST (Abstract Syntax Tree – Soyut Sözdizim Ağacı)** üretmeye yarar.

Bu paket, **Go compiler**’ın temel bileşenlerinden biridir ve genellikle:

* Kod analizi
* Statik analiz araçları (ör. `golint`)
* Kod dönüştürme (refactoring)
* Go dokümantasyon araçları (`go doc`)

gibi yerlerde kullanılır.

---

# 📂 Paket İçi Önemli Yapılar & Fonksiyonlar

### 🔹 1. `parser.ParseExpr`

* Tek bir Go ifadesini (expression) parse eder.
* Geriye `ast.Expr` döner.
*/
``go
package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main() {
	expr, err := parser.ParseExpr("a + b*3")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", expr) // AST düğümü yazdırır
}
``
/*
---

### 🔹 2. `parser.ParseFile`

* Bir Go dosyasını parse eder.
* Geriye `*ast.File` döner.
*/
``go
package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main() {
	src :=package main

import "fmt"

func main() {
    fmt.Println("Merhaba Dünya")
}

	// Token konumlarını takip etmek için FileSet
	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, "example.go", src, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	fmt.Println("Paket Adı:", file.Name.Name) // main
}
``
/*
---

### 🔹 3. `parser.ParseDir`

* Bir dizindeki **tüm Go dosyalarını** parse eder.
* Geriye `map[string]*ast.Package` döner.
*/
``go
package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	fset := token.NewFileSet()

	// Geçerli dizindeki dosyaları parse et
	pkgs, err := parser.ParseDir(fset, ".", nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	for name := range pkgs {
		fmt.Println("Bulunan paket:", name)
	}
}
``
/*
---

### 🔹 4. `parser.Mode` (Parse Modları)

`ParseFile` ve `ParseDir` fonksiyonlarında davranışı değiştirmek için kullanılır:

* `parser.PackageClauseOnly` → sadece `package` kısmını okur
* `parser.ImportsOnly` → sadece import satırlarını okur
* `parser.ParseComments` → yorum satırlarını da parse eder
* `parser.AllErrors` → tüm hataları raporlar

Örnek:
*/
``go
file, err := parser.ParseFile(fset, "test.go", nil, parser.ImportsOnly)
``
/*
---

# ⚡ Daha Kapsamlı Örnek: Fonksiyon İsimlerini Çekmek
*/
``go
package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"go/ast"
)

func main() {
	src := package main

import "fmt"

func hello() { fmt.Println("hi") }
func add(a int, b int) int { return a + b }


	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "example.go", src, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	// AST üzerinde gezinti
	ast.Inspect(file, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok {
			fmt.Println("Fonksiyon bulundu:", fn.Name.Name)
		}
		return true
	})
}
``

//📌 Çıktı:

``
Fonksiyon bulundu: hello
Fonksiyon bulundu: add
``
/*
---

# ✅ Özet

`go/parser` paketi:

* **ParseExpr** → Tek ifade (ör. `x + y*2`)
* **ParseFile** → Tek bir Go dosyası
* **ParseDir** → Bir dizindeki tüm Go dosyaları
* **parser.Mode** → Yorumları, importları veya tüm hataları dahil etmek için

kullanılır.

Kısaca, Go kodunu **AST** (soyut sözdizim ağacı) haline getirip üzerinde işlem yapmanı sağlar.

---

👉 İstersen sana `parser` + `ast` paketini birlikte kullanarak **Go kodundan otomatik fonksiyon dokümantasyonu çıkaran bir mini araç** da yazabilirim. İstiyor musun?
EVET
*/