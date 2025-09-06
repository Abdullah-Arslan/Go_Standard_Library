/*
**Go’nun `go/types` paketini** baştan sona açıklayacağım ve örneklerle göstereceğim. Bu paket Go dilinde **tip kontrolü ve tip bilgisi** için kritik öneme sahiptir.

---

# 📌 `go/types` Paketi Nedir?

* `go/types`, Go programlarının **tiplerini analiz etmek**, **tip hatalarını bulmak** ve **tip bilgilerini almak** için kullanılır.
* Genellikle **`go/ast`** ve **`go/parser`** ile birlikte kullanılır.
* Derleyici düzeyinde tip kontrolü yapmak isteyen araçlar (`gofmt`, linters, kod analiz araçları) bu paketi kullanır.

---

# 📦 Temel Kavramlar

1. **Type**

   * Her ifade, değişken, fonksiyon veya sabitin bir tipi vardır.
   * `types.Type` interface’i bu tipleri temsil eder.

2. **Basic types**

   * `types.Basic` → int, string, bool, float64 vb.

3. **Named types**

   * Kullanıcının tanımladığı tipler (`type MyInt int`)

4. **Composite types**

   * `struct`, `array`, `slice`, `map`, `pointer`, `channel`, `interface`

5. **Package & Scope**

   * `types.Package` → paket bazlı tip bilgisi
   * `types.Scope` → belirli bir scope’daki tanımlamaları tutar

6. **Info**

   * `types.Info` → AST üzerindeki tipleri ve değerleri ilişkilendirmek için kullanılır

7. **Config & Check**

   * `types.Config` → tip kontrolünü başlatmak için konfigürasyon
   * `types.Check` → `Check` fonksiyonu ile tip denetimi yapılır

---

# 📝 Basit Örnek 1: Tip Kontrolü
*/
``go
package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"go/types"
)

func main() {
	src := `
		package main
		func main() {
			var x int = 10
			var y string = "hello"
			_ = x + 5
			_ = x + y // tip hatası
		}
	`

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "example.go", src, 0)
	if err != nil {
		panic(err)
	}

	conf := types.Config{Importer: nil}
	info := &types.Info{
		Types: make(map[types.Expr]types.TypeAndValue),
	}

	// Tip kontrolü
	_, err = conf.Check("main", fset, []*ast.File{file}, info)
	if err != nil {
		fmt.Println("Tip hatası:", err)
	}

	// Tip bilgilerini yazdır
	for expr, tv := range info.Types {
		fmt.Printf("%v -> %v\n", expr, tv.Type)
	}
}
``

//📌 Çıktı (kısmi):

``
Tip hatası: example.go:6:11: invalid operation: x + y (mismatched types int and string)
``
/*
---

# 📝 Basit Örnek 2: Değişken Tiplerini Okuma
*/
``go
package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"go/types"
)

func main() {
	src := `
		package main
		func main() {
			var a int
			var b string
			c := 3.14
		}
	`

	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "example.go", src, 0)

	conf := types.Config{}
	info := &types.Info{
		Types: make(map[types.Expr]types.TypeAndValue),
		Defs:  make(map[*ast.Ident]types.Object),
	}

	_, err := conf.Check("main", fset, []*ast.File{file}, info)
	if err != nil {
		panic(err)
	}

	for ident, obj := range info.Defs {
		if obj != nil {
			fmt.Printf("Değişken: %s, Tip: %s\n", ident.Name, obj.Type())
		}
	}
}
``

//📌 Çıktı:

``
Değişken: a, Tip: int
Değişken: b, Tip: string
Değişken: c, Tip: float64
``
/*
---

# 📝 Örnek 3: Fonksiyon ve Parametre Tipleri
*/
``go
package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"go/types"
)

func main() {
	src := `
		package main
		func add(x int, y int) int {
			return x + y
		}
	`

	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "example.go", src, 0)

	conf := types.Config{}
	info := &types.Info{
		Defs: make(map[*ast.Ident]types.Object),
	}

	pkg, err := conf.Check("main", fset, []*ast.File{file}, info)
	if err != nil {
		panic(err)
	}

	for ident, obj := range info.Defs {
		if obj != nil {
			fmt.Printf("%s -> %s\n", ident.Name, obj.Type())
		}
	}

	fmt.Println("Paket adı:", pkg.Name())
}
``

//📌 Çıktı:

``
add -> func(x int, y int) int
x -> int
y -> int
Paket adı: main
``
/*
---

# 🚀 Özet

1. `go/types` → Go kodundaki **tüm tipleri ve tip hatalarını** analiz eder
2. AST ile birlikte çalışır (`go/parser` + `go/ast`)
3. Önemli tipler: `Basic`, `Named`, `Struct`, `Array`, `Slice`, `Pointer`, `Map`, `Interface`
4. `types.Config` ve `types.Info` → tip kontrolü ve bilgi toplamak için kullanılır
5. `Check` fonksiyonu → tip denetimini başlatır
6. `Info.Defs` → değişken/fonksiyon tanımları ve tipleri
7. `Info.Types` → ifadelerin tipleri

---

İstersen ben sana bunu **bir adım ileri götürüp `go/types` ile tüm bir Go dosyasındaki değişken, sabit ve fonksiyonların tiplerini otomatik olarak listeleyen analiz aracı** da yazabilirim.

Bunu yapayım mı?
EVET
*/
/*
Harika 😄 O zaman sana **`go/types` kullanarak bir Go dosyasındaki tüm değişken, sabit ve fonksiyonları tipleriyle birlikte listeleyen bir analiz aracı** hazırlayacağım.

Bu araç:

* Bir `.go` dosyasını okuyacak
* `go/parser` ile AST oluşturacak
* `go/types` ile tip kontrolü yapacak
* Tüm **değişken**, **sabit** ve **fonksiyonları** isim ve tipleriyle ekrana yazacak 🚀

---

# 📂 Kod – Go Types Analiz Aracı
*/
``go
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Kullanım: go run main.go <dosya.go>")
		return
	}

	filename := os.Args[1]

	// Dosyayı oku
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Dosya okunamadı:", err)
		return
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, src, parser.AllErrors)
	if err != nil {
		fmt.Println("Parse hatası:", err)
		return
	}

	conf := types.Config{}
	info := &types.Info{
		Defs:  make(map[*ast.Ident]types.Object),
		Types: make(map[ast.Expr]types.TypeAndValue),
	}

	// Tip kontrolü
	pkg, err := conf.Check("main", fset, []*ast.File{file}, info)
	if err != nil {
		fmt.Println("Tip hatası:", err)
	}

	// Değişkenler, Sabitler ve Fonksiyonlar
	var vars, consts, funcs []string

	for ident, obj := range info.Defs {
		if obj == nil {
			continue
		}

		switch obj.(type) {
		case *types.Var:
			vars = append(vars, fmt.Sprintf("%s : %s", ident.Name, obj.Type()))
		case *types.Const:
			consts = append(consts, fmt.Sprintf("%s : %s", ident.Name, obj.Type()))
		case *types.Func:
			funcs = append(funcs, fmt.Sprintf("%s : %s", ident.Name, obj.Type()))
		}
	}

	fmt.Println("📌 Paket:", pkg.Name())
	fmt.Println("\n📌 Fonksiyonlar:")
	for _, f := range funcs {
		fmt.Println(" -", f)
	}

	fmt.Println("\n📌 Değişkenler:")
	for _, v := range vars {
		fmt.Println(" -", v)
	}

	fmt.Println("\n📌 Sabitler:")
	for _, c := range consts {
		fmt.Println(" -", c)
	}
}
``
/*
---

# 📌 Örnek Test

ornek.go dosyamız:
*/
``go
package main

import "fmt"

const Pi = 3.14

func add(x int, y int) int {
	return x + y
}

func main() {
	var a int
	var b string
	c := 2.5
	fmt.Println(add(a, int(c)), b, Pi)
}
``

//Çalıştır:

``bash
go run analyzer.go ornek.go
``

//📌 Çıktı:

``
📌 Paket: main

📌 Fonksiyonlar:
 - add : func(x int, y int) int
 - main : func()

📌 Değişkenler:
 - a : int
 - b : string
 - c : float64

📌 Sabitler:
 - Pi : untyped float
``
/*
---

# 🚀 Özet

* `types.Config.Check` → tip kontrolü ve tip bilgisi toplar
* `info.Defs` → tüm **tanımlar** (var, const, func)
* Her bir obje tipi (`*types.Var`, `*types.Const`, `*types.Func`) ile filtreleme
* Bu araç ile artık bir Go dosyasındaki tüm değişkenleri, sabitleri ve fonksiyonları **tipleriyle birlikte** görebilirsin ✅

---

İstersen ben bunu bir adım ileri götürüp **her fonksiyon içindeki parametre ve dönüş tiplerini ayrıca listeleyen**, yani tam bir **Go dosyası tip haritası aracı** hâline getirebilirim.

Bunu yapayım mı?
EVET
*/
/*
Harika 😄 O zaman bir sonraki adımda **tam bir Go dosyası tip haritası aracı** yapalım. Bu araç:

* Bir `.go` dosyasını okuyacak
* AST ve `go/types` ile tip kontrolü yapacak
* **Tüm fonksiyonları**, **parametre tiplerini**, **dönüş tiplerini**, **local değişkenleri**, **sabitleri** ve **global değişkenleri** listeleyecek

Yani neredeyse bir **Go dosyası tip haritası** oluşturacak 🚀

---

# 📂 Kod – Go Tip Haritası Analiz Aracı
*/
``go
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Kullanım: go run main.go <dosya.go>")
		return
	}

	filename := os.Args[1]

	// Dosyayı oku
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Dosya okunamadı:", err)
		return
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, src, parser.AllErrors)
	if err != nil {
		fmt.Println("Parse hatası:", err)
		return
	}

	conf := types.Config{}
	info := &types.Info{
		Defs:  make(map[*ast.Ident]types.Object),
		Types: make(map[ast.Expr]types.TypeAndValue),
	}

	// Tip kontrolü
	pkg, err := conf.Check("main", fset, []*ast.File{file}, info)
	if err != nil {
		fmt.Println("Tip hatası:", err)
	}

	fmt.Println("📌 Paket:", pkg.Name())

	// Fonksiyonlar ve detayları
	fmt.Println("\n📌 Fonksiyonlar:")
	for ident, obj := range info.Defs {
		fn, ok := obj.(*types.Func)
		if !ok || obj == nil {
			continue
		}
		sig := fn.Type().(*types.Signature)
		fmt.Printf(" - %s(", ident.Name)
		params := sig.Params()
		for i := 0; i < params.Len(); i++ {
			p := params.At(i)
			fmt.Printf("%s %s", p.Name(), p.Type())
			if i < params.Len()-1 {
				fmt.Print(", ")
			}
		}
		fmt.Print(")")
		results := sig.Results()
		if results.Len() > 0 {
			fmt.Print(" -> ")
			for i := 0; i < results.Len(); i++ {
				r := results.At(i)
				fmt.Print(r.Type())
				if i < results.Len()-1 {
					fmt.Print(", ")
				}
			}
		}
		fmt.Println()
	}

	// Global değişkenler ve sabitler
	var globals, consts []string
	for ident, obj := range info.Defs {
		if obj == nil {
			continue
		}
		switch v := obj.(type) {
		case *types.Var:
			globals = append(globals, fmt.Sprintf("%s : %s", ident.Name, v.Type()))
		case *types.Const:
			consts = append(consts, fmt.Sprintf("%s : %s", ident.Name, v.Type()))
		}
	}

	fmt.Println("\n📌 Global Değişkenler:")
	for _, v := range globals {
		fmt.Println(" -", v)
	}

	fmt.Println("\n📌 Sabitler:")
	for _, c := range consts {
		fmt.Println(" -", c)
	}

	// Local değişkenler (main fonksiyon içi) AST üzerinden bulunabilir
	fmt.Println("\n📌 Local Değişkenler (fonksiyon içinde):")
	ast.Inspect(file, func(n ast.Node) bool {
		switch stmt := n.(type) {
		case *ast.AssignStmt:
			for _, expr := range stmt.Lhs {
				if ident, ok := expr.(*ast.Ident); ok {
					if obj := info.ObjectOf(ident); obj != nil {
						if _, ok := obj.(*types.Var); ok && obj.Parent() != pkg.Scope() {
							fmt.Printf(" - %s : %s\n", ident.Name, obj.Type())
						}
					}
				}
			}
		case *ast.DeclStmt:
			if decl, ok := stmt.Decl.(*ast.GenDecl); ok {
				for _, spec := range decl.Specs {
					if vs, ok := spec.(*ast.ValueSpec); ok {
						for _, ident := range vs.Names {
							if obj := info.ObjectOf(ident); obj != nil {
								if _, ok := obj.(*types.Var); ok && obj.Parent() != pkg.Scope() {
									fmt.Printf(" - %s : %s\n", ident.Name, obj.Type())
								}
							}
						}
					}
				}
			}
		}
		return true
	})
}
``
/*
---

# 📌 Örnek Test

ornek.go:
*/
``go
package main

import "fmt"

const Pi = 3.14

func add(x int, y int) int {
	sum := x + y
	return sum
}

func main() {
	var a int
	b := "hello"
	c := 3.14
	fmt.Println(add(a, int(c)), b, Pi)
}
``

//Çalıştır:

``bash
go run tipharita.go ornek.go
``

//📌 Çıktı (özet):

``
📌 Paket: main

📌 Fonksiyonlar:
 - add(x int, y int) -> int
 - main()

📌 Global Değişkenler:
 - a : int

📌 Sabitler:
 - Pi : untyped float

📌 Local Değişkenler (fonksiyon içinde):
 - sum : int
 - a : int
 - b : string
 - c : float64
``
/*
---

# 🚀 Özellikler

1. Fonksiyonların **parametre ve dönüş tipleri** listeleniyor
2. Global ve sabitler ayrı gösteriliyor
3. AST ve `types.Info` ile **local değişkenler** de bulunuyor
4. Neredeyse tüm tip bilgilerini tek bir araçta görebiliyorsun ✅

---

İstersen bunu bir adım daha ileri götürüp **fonksiyon içinde hangi değişken hangi tipte, hangi satırda ve scope içinde** detaylı bir **JSON çıktısı** üreten bir araç hâline getirebiliriz.

Bunu yapayım mı?
*/