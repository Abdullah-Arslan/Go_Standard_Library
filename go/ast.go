Golang’daki **`ast` (Abstract Syntax Tree)** paketi, Go kaynak kodunu bir **soyut sözdizim ağacı (AST)** yapısında temsil etmek için kullanılır. Yani, Go kodunu parçalayarak programın **değişkenlerini, fonksiyonlarını, ifadelerini, tiplerini, paketlerini** bir ağaç yapısında tutar.

AST yapısı, **kod analizi, kod dönüştürme, linter yazma, statik analiz araçları, code generator** geliştirme gibi amaçlarla kullanılır.

`ast` paketi, genelde şu paketlerle birlikte kullanılır:

* **`go/parser`** → Go kaynak dosyasını parse eder (AST üretir).
* **`go/token`** → Token konumlarını (satır, sütun) tutar.
* **`go/format`** → AST’ten tekrar formatlı Go kodu üretir.

---

## 🔹 `ast` Paketindeki Temel Yapılar

`ast` paketi, Go dilinin **gramerini temsil eden struct’lar** içerir. Önemli olanlardan bazıları:

* **`ast.File`** → Bir Go dosyasını temsil eder.
* **`ast.Package`** → Birden fazla dosyadan oluşan bir paketi temsil eder.
* **`ast.FuncDecl`** → Fonksiyon bildirimlerini temsil eder.
* **`ast.GenDecl`** → Genel bildirimleri (import, var, const, type) tutar.
* **`ast.Ident`** → Bir değişken/fonksiyon/tip ismi (identifier).
* **`ast.BasicLit`** → Sayılar, stringler, bool gibi temel literal’ler.
* **`ast.CallExpr`** → Fonksiyon çağrısı.
* **`ast.AssignStmt`** → Atama ifadesi.
* **`ast.IfStmt`, `ast.ForStmt`, `ast.SwitchStmt`** → Kontrol yapıları.
* **`ast.CompositeLit`** → Slice, struct, map literal ifadeleri.
* **`ast.Inspect`** → AST üzerinde dolaşmayı kolaylaştırır.

---

## 🔹 Basit Örnek – Kod Parse Etme ve AST Basma

```go
package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"go/ast"
)

func main() {
	src := `package main
import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(add(2, 3))
}`

	// Token set oluştur (dosya konumları için)
	fset := token.NewFileSet()

	// Kaynak kodu parse et → AST üret
	node, err := parser.ParseFile(fset, "example.go", src, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	// Paket ismi
	fmt.Println("Paket:", node.Name.Name)

	// AST içindeki fonksiyonları gezelim
	for _, decl := range node.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			fmt.Println("Fonksiyon:", fn.Name.Name)
		}
	}
}
```

### Çıktı:

```
Paket: main
Fonksiyon: add
Fonksiyon: main
```

---

## 🔹 AST İçinde Dolaşmak (`ast.Inspect`)

```go
ast.Inspect(node, func(n ast.Node) bool {
    switch x := n.(type) {
    case *ast.FuncDecl:
        fmt.Println("Fonksiyon bulundu:", x.Name.Name)
    case *ast.CallExpr:
        if fun, ok := x.Fun.(*ast.Ident); ok {
            fmt.Println("Fonksiyon çağrısı:", fun.Name)
        }
    }
    return true
})
```

### Çıktı:

```
Fonksiyon bulundu: add
Fonksiyon bulundu: main
Fonksiyon çağrısı: Println
Fonksiyon çağrısı: add
```

---

## 🔹 AST Manipülasyonu (Kod Üretme)

AST üzerinde değişiklik yapıp tekrar **Go koduna çevirmek** mümkündür.

```go
package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

func main() {
	src := `package main
import "fmt"
func main() {
	fmt.Println("Hello")
}`

	fset := token.NewFileSet()
	node, _ := parser.ParseFile(fset, "", src, parser.ParseComments)

	// AST üzerinde dolaş ve string literal değiştir
	ast.Inspect(node, func(n ast.Node) bool {
		if bl, ok := n.(*ast.BasicLit); ok && bl.Kind == token.STRING {
			bl.Value = `"Merhaba Dünya"`
		}
		return true
	})

	// Değiştirilmiş kodu yazdır
	printer.Fprint(os.Stdout, fset, node)
}
```

### Çıktı:

```go
package main
import "fmt"
func main() {
	fmt.Println("Merhaba Dünya")
}
```

---

## 🔹 AST ile Yeni Fonksiyon Ekleme

```go
newFunc := &ast.FuncDecl{
	Name: ast.NewIdent("hello"),
	Type: &ast.FuncType{
		Params: &ast.FieldList{},
	},
	Body: &ast.BlockStmt{
		List: []ast.Stmt{
			&ast.ExprStmt{
				X: &ast.CallExpr{
					Fun:  ast.NewIdent("println"),
					Args: []ast.Expr{ast.NewIdent(`"Hello from new function!"`)},
				},
			},
		},
	},
}

node.Decls = append(node.Decls, newFunc)
```

Böylece mevcut AST’ye yeni fonksiyon eklenir.

---

## 🔹 Özet

* `ast` paketi, Go kaynak kodunu **ağaç yapısı** olarak temsil eder.
* Genelde `go/parser`, `go/token`, `go/printer`, `go/format` ile birlikte çalışır.
* Kullanım alanları:

  * **Kod analizi** (ör. lint, güvenlik tarayıcıları)
  * **Kod dönüştürme** (ör. otomatik refactor)
  * **Kod üretme** (ör. code generator)
  * **Araç geliştirme** (ör. `go vet`, `gofmt`)

---

👉 İstersen sana **`ast` paketindeki tüm struct’ları tek tek açıklayıp her birine örnekler verebilirim** (ör. `FuncDecl`, `GenDecl`, `IfStmt`, `ForStmt`, `SwitchStmt`, `AssignStmt` vs.). Bunu ister misin?
EVET

Harika 👍 O zaman sana **Go’nun `ast` paketindeki tüm önemli struct ve tipleri** tek tek açıklayacağım ve **her biri için küçük örnekler** vereceğim. Bu biraz uzun olacak ama böylece `ast` paketini **tam anlamıyla kavrayabileceksin**.

---

# 📘 Go `ast` Paketi – Tüm Yapılar

## 🔹 1. `ast.File`

Bir Go dosyasını temsil eder. İçinde **package adı, importlar, fonksiyonlar, değişkenler** vardır.

```go
type File struct {
    Name    *Ident        // Paket adı
    Decls   []Decl        // Bildirimler (fonksiyon, değişken, tip)
    Scope   *Scope        // Tanımlanan isimlerin kapsamı
    Imports []*ImportSpec // Import edilen paketler
}
```

**Örnek:**

```go
src := `package main
import "fmt"
func main() {}`
fset := token.NewFileSet()
f, _ := parser.ParseFile(fset, "", src, 0)
fmt.Println("Paket adı:", f.Name.Name) // "main"
fmt.Println("Import sayısı:", len(f.Imports)) // 1
```

---

## 🔹 2. `ast.Package`

Birden fazla dosyadan oluşan Go paketini temsil eder.

```go
type Package struct {
    Name  string
    Files map[string]*File
}
```

**Örnek:**

```go
pkgs, _ := parser.ParseDir(fset, ".", nil, 0)
for name := range pkgs {
    fmt.Println("Paket bulundu:", name)
}
```

---

## 🔹 3. `ast.FuncDecl`

Fonksiyon bildirimini temsil eder.

```go
type FuncDecl struct {
    Name *Ident      // Fonksiyon adı
    Type *FuncType   // Parametreler ve dönüş tipi
    Body *BlockStmt  // Fonksiyon gövdesi
}
```

**Örnek:**

```go
ast.Inspect(f, func(n ast.Node) bool {
    if fn, ok := n.(*ast.FuncDecl); ok {
        fmt.Println("Fonksiyon adı:", fn.Name.Name)
    }
    return true
})
```

---

## 🔹 4. `ast.GenDecl`

Genel bildirimler: **import, var, const, type**

```go
type GenDecl struct {
    Tok   token.Token // IMPORT, CONST, TYPE, VAR
    Specs []Spec
}
```

**Örnek:**

```go
const pi = 3.14
type User struct{}
var x int
```

Bu kodda her biri `GenDecl` ile temsil edilir.

---

## 🔹 5. `ast.Ident`

Bir **identifier** yani isim (değişken, fonksiyon, tip adı).

```go
x := 42
```

Burada `x` → `ast.Ident`

```go
ast.Inspect(f, func(n ast.Node) bool {
    if id, ok := n.(*ast.Ident); ok {
        fmt.Println("Identifier:", id.Name)
    }
    return true
})
```

---

## 🔹 6. `ast.BasicLit`

Basit literal değerler (**string, int, float, bool, rune**).

```go
x := "Merhaba"
y := 123
```

Kodda `"Merhaba"` ve `123` → `BasicLit`

---

## 🔹 7. `ast.CallExpr`

Fonksiyon çağrısını temsil eder.

```go
fmt.Println("Selam")
```

Burada:

* `fmt.Println` → `Fun`
* `"Selam"` → `Args[0]`

```go
ast.Inspect(f, func(n ast.Node) bool {
    if call, ok := n.(*ast.CallExpr); ok {
        fmt.Println("Fonksiyon çağrısı bulundu")
    }
    return true
})
```

---

## 🔹 8. `ast.AssignStmt`

Atama ifadeleri.

```go
x := 10
y = 20
```

Bunlar → `AssignStmt`

---

## 🔹 9. Kontrol Yapıları

### `ast.IfStmt`

```go
if x > 0 { fmt.Println("Pozitif") }
```

### `ast.ForStmt`

```go
for i := 0; i < 5; i++ { fmt.Println(i) }
```

### `ast.SwitchStmt`

```go
switch x {
case 1: fmt.Println("Bir")
}
```

Her biri `ast` içinde farklı struct olarak tutulur.

---

## 🔹 10. `ast.CompositeLit`

Literal struct, slice, map, array değerleri.

```go
u := User{"Ahmet", 25}
s := []int{1, 2, 3}
```

Bunlar → `CompositeLit`

---

## 🔹 11. `ast.StructType`, `ast.InterfaceType`

```go
type User struct {
    Name string
}
type Speaker interface {
    Speak()
}
```

Bunlar → `StructType` ve `InterfaceType`

---

## 🔹 12. `ast.ReturnStmt`

Fonksiyon içindeki `return` ifadeleri.

```go
func add(a, b int) int {
    return a + b
}
```

---

## 🔹 13. `ast.Inspect`

AST üzerinde dolaşmak için kullanılan yardımcı fonksiyon.

```go
ast.Inspect(f, func(n ast.Node) bool {
    switch x := n.(type) {
    case *ast.FuncDecl:
        fmt.Println("Fonksiyon:", x.Name.Name)
    case *ast.IfStmt:
        fmt.Println("If ifadesi bulundu")
    }
    return true
})
```

---

# 📌 Özet

* `ast.File`, `ast.Package` → dosya/paket yapısı
* `ast.FuncDecl`, `ast.GenDecl` → bildirimler
* `ast.Ident`, `ast.BasicLit`, `ast.CallExpr` → ifadeler
* `ast.AssignStmt`, `ast.IfStmt`, `ast.ForStmt` → kontrol/atama
* `ast.CompositeLit`, `ast.StructType`, `ast.InterfaceType` → tipler
* `ast.ReturnStmt`, `ast.BlockStmt` → gövde yapıları
* `ast.Inspect` → gezinti

---

👉 İstersen ben sana **küçük bir program** yazabilirim: Bu program bir Go dosyasını parse edip AST’den **fonksiyonların parametrelerini, dönüş tiplerini, değişken bildirimlerini** çıkarsın. Böylece hepsini canlı görebiliriz.

Bunu ister misin?
EVET

Süper 👍 O zaman sana **bir analiz aracı** yazalım. Bu araç:

* Bir Go kaynağını parse edecek
* `ast` kullanarak fonksiyonları bulacak
* Parametrelerini ve dönüş tiplerini listeleyecek
* Değişken bildirimlerini gösterecek

---

# 📘 Örnek Kod – AST Analiz Aracı

```go
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func main() {
	src := `package main
import "fmt"

var g = 42

type User struct {
	Name string
	Age  int
}

func add(a int, b int) int {
	return a + b
}

func hello(name string) {
	fmt.Println("Hello", name)
}

func main() {
	x := 10
	y := 20
	fmt.Println(add(x, y))
}`

	// Token set (konum bilgisi için)
	fset := token.NewFileSet()

	// Parse et
	node, err := parser.ParseFile(fset, "example.go", src, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	fmt.Println("📦 Paket:", node.Name.Name)
	fmt.Println()

	// Fonksiyonları bulalım
	fmt.Println("🔹 Fonksiyonlar:")
	for _, decl := range node.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			params := []string{}
			if fn.Type.Params != nil {
				for _, field := range fn.Type.Params.List {
					for _, name := range field.Names {
						params = append(params, fmt.Sprintf("%s %s", name.Name, exprToString(field.Type)))
					}
				}
			}

			results := []string{}
			if fn.Type.Results != nil {
				for _, field := range fn.Type.Results.List {
					results = append(results, exprToString(field.Type))
				}
			}

			fmt.Printf("  • %s(%s)", fn.Name.Name, strings.Join(params, ", "))
			if len(results) > 0 {
				fmt.Printf(" → %s", strings.Join(results, ", "))
			}
			fmt.Println()
		}
	}

	// Değişken bildirimlerini bulalım
	fmt.Println("\n🔹 Değişkenler:")
	ast.Inspect(node, func(n ast.Node) bool {
		if decl, ok := n.(*ast.GenDecl); ok && decl.Tok == token.VAR {
			for _, spec := range decl.Specs {
				if vs, ok := spec.(*ast.ValueSpec); ok {
					for _, name := range vs.Names {
						fmt.Println("  •", name.Name)
					}
				}
			}
		}
		if assign, ok := n.(*ast.AssignStmt); ok {
			for _, lhs := range assign.Lhs {
				if id, ok := lhs.(*ast.Ident); ok {
					fmt.Println("  •", id.Name)
				}
			}
		}
		return true
	})
}

// Yardımcı: Tipleri stringe çevir
func exprToString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return exprToString(t.X) + "." + t.Sel.Name
	default:
		return fmt.Sprintf("%T", expr)
	}
}
```

---

# ▶️ Çalıştırınca Çıktı

```
📦 Paket: main

🔹 Fonksiyonlar:
  • add(a int, b int) → int
  • hello(name string)
  • main()

🔹 Değişkenler:
  • g
  • x
  • y
```

---

👉 Bu aracı genişletip:

* **Struct ve interface tiplerini** listeleyebiliriz,
* **If, For, Switch gibi kontrol yapıları** tespit edebiliriz,
* Hatta **fonksiyon çağrılarını** çıkartabiliriz.

Bunu da ister misin?
* **If, For, Switch gibi kontrol yapıları** tespit edebiliriz,

Tamamdır! 😎 O zaman mevcut AST analiz aracımızı genişletelim ve **`if`, `for`, `switch`** gibi kontrol yapılarını tespit edecek hale getirelim. Böylece kodun **akış kontrolünü de çıkarabiliriz**.

---

# 📘 AST Kontrol Yapıları Tespit Aracı

```go
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func main() {
	src := `package main
import "fmt"

func main() {
	x := 10
	y := 20

	if x > y {
		fmt.Println("x büyük")
	} else {
		fmt.Println("y büyük")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	switch x {
	case 10:
		fmt.Println("On")
	default:
		fmt.Println("Diğer")
	}
}`

	// Token set oluştur
	fset := token.NewFileSet()

	// Parse et
	node, err := parser.ParseFile(fset, "example.go", src, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	fmt.Println("📦 Paket:", node.Name.Name)
	fmt.Println()

	// Fonksiyonları bul
	fmt.Println("🔹 Fonksiyonlar:")
	for _, decl := range node.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			fmt.Println("  •", fn.Name.Name)
		}
	}

	// Kontrol yapılarını bul
	fmt.Println("\n🔹 Kontrol Yapıları:")
	ast.Inspect(node, func(n ast.Node) bool {
		switch stmt := n.(type) {
		case *ast.IfStmt:
			fmt.Println("  • If ifadesi bulundu")
		case *ast.ForStmt:
			fmt.Println("  • For döngüsü bulundu")
		case *ast.RangeStmt:
			fmt.Println("  • For-range döngüsü bulundu")
		case *ast.SwitchStmt:
			fmt.Println("  • Switch ifadesi bulundu")
		}
		return true
	})
}
```

---

# ▶️ Çıktı

```
📦 Paket: main

🔹 Fonksiyonlar:
  • main

🔹 Kontrol Yapıları:
  • If ifadesi bulundu
  • For döngüsü bulundu
  • Switch ifadesi bulundu
```

---

## 🔹 Açıklama

* **`ast.IfStmt`** → `if` ifadeleri
* **`ast.ForStmt`** → klasik `for` döngüleri
* **`ast.RangeStmt`** → `for k, v := range …` döngüleri
* **`ast.SwitchStmt`** → `switch` ifadeleri
* `ast.Inspect` ile tüm AST’de dolaşıyoruz ve tip kontrolü ile tespit ediyoruz.

---

💡 İstersen bir adım daha ileri gidip **her kontrol yapısının konumunu (satır ve sütun) ve içerisindeki değişkenleri** de listeleyebiliriz. Böylece **tam bir statik analiz aracı** olur.

Bunu da yapmamı ister misin?
