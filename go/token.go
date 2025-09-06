/*
**Go’nun `go/token` paketini** en ince ayrıntısına kadar açıklayayım.

---

# 📌 `go/token` Paketi Nedir?

`go/token` paketi, **Go kaynak kodunu analiz etmek için kullanılan token (sözcük) bilgilerini** tutar.
AST (Abstract Syntax Tree) ve parser işlemleri için bir **temel yapı taşıdır**.

👉 Kısacası: `go/token`, Go kodundaki **konumları (dosya, satır, sütun)** ve **token türlerini (if, for, +, =, identifier, vs.)** temsil eder.

Bu paket genelde şunlarla birlikte kullanılır:

* `go/scanner` (Go kodunu tokenize eder)
* `go/parser` (AST oluşturur)
* `go/ast` (AST üzerinde işlem yapar)

---

# 📦 Paket İçeriği

## 1. Token Türleri (`token.Token`)

Go’daki tüm sözcük ve sembollerin listesi burada tanımlanır:

* **Identifiers** → `IDENT` (değişken, fonksiyon isimleri)
* **Keywords** → `BREAK, CASE, FUNC, RETURN, IF, FOR, PACKAGE, IMPORT` …
* **Literals** → `INT, FLOAT, IMAG, CHAR, STRING`
* **Operators & Delimiters** → `+ - * / % & | ^ << >> &^ ++ -- == != < <= > >= && || <- ( ) [ ] { } , ; : = := ...`
* **Special** → `ILLEGAL`, `EOF`, `COMMENT`

Her token’ın string karşılığı vardır (`token.ADD.String()` → `"+"`).

---

## 2. Konum Yönetimi

`go/token`’in en önemli kısmı **kaynak kod konumları**dır.

### `Position`

Bir dosya içindeki konum bilgisini tutar.
*/
``go
type Position struct {
    Filename string // dosya adı
    Offset   int    // dosya başından byte cinsinden uzaklık
    Line     int    // satır numarası
    Column   int    // sütun numarası
}
``

### `Pos`

AST içindeki soyut konum tipidir (sayısal değer).
Kendi başına bilgi taşımaz; `FileSet` içinde çözülür.

### `FileSet`

Bir veya daha fazla dosyanın pozisyonlarını yönetir.

``go
fset := token.NewFileSet()
file := fset.AddFile("example.go", -1, len(src))
``
/*
---

# 📝 Örnekler

## Örnek 1: Token Türlerini Yazdırma
*/
``go
package main

import (
	"fmt"
	"go/token"
)

func main() {
	// Bazı token türlerini yazdıralım
	fmt.Println(token.ADD, token.ADD.String())   // + "+"
	fmt.Println(token.IF, token.IF.String())     // if "if"
	fmt.Println(token.RETURN, token.RETURN.String()) // return "return"
	fmt.Println(token.INT, token.INT.String())   // int "int"
}
``

//📌 Çıktı:

``
43 +
4 if
15 return
5 int
``
/*
---

## Örnek 2: Scanner ile Tokenize Etme
*/
``go
package main

import (
	"fmt"
	"go/scanner"
	"go/token"
)

func main() {
	src := `package main
	func add(x int, y int) int { return x + y }`

	// FileSet oluştur
	fset := token.NewFileSet()
	file := fset.AddFile("example.go", -1, len(src))

	// Scanner başlat
	var s scanner.Scanner
	s.Init(file, []byte(src), nil, scanner.ScanComments)

	// Tokenları sırayla yazdır
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%-8s %-10q %v\n", tok, lit, fset.Position(pos))
	}
}
``

//📌 Çıktı (kısaltılmış):

```
package "package" example.go:1:1
IDENT   "main"   example.go:1:9
func    "func"   example.go:2:2
IDENT   "add"    example.go:2:7
(       "("      example.go:2:10
IDENT   "x"      example.go:2:11
int     "int"    example.go:2:13

```
/*
---

## Örnek 3: Position Kullanımı
*/
``go
package main

import (
	"fmt"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	file := fset.AddFile("test.go", -1, 100)

	// Satır ekle
	file.SetLinesForContent([]byte("line1\nline2\nline3"))

	// 8. byte’ın konumunu öğrenelim
	pos := file.Pos(8)
	position := fset.Position(pos)

	fmt.Println("Offset:", position.Offset)
	fmt.Println("Line:", position.Line)
	fmt.Println("Column:", position.Column)
}
``

//📌 Çıktı:

``
Offset: 8
Line: 2
Column: 3
``
/*
---

## Örnek 4: Parser ile Kullanımı

`go/parser` ve `go/token` birlikte kullanılarak AST çıkarılabilir:
*/
``go
package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main() {
	src := `package main; var x = 42`

	fset := token.NewFileSet()
	node, _ := parser.ParseFile(fset, "example.go", src, 0)

	// Dosya bilgisi
	fmt.Println("Paket adı:", node.Name.Name)
	fmt.Println("Konum:", fset.Position(node.Pos()))
}
``

//📌 Çıktı:

``
Paket adı: main
Konum: example.go:1:1
``
/*
---

# 🚀 Özet

* `go/token` → Go kaynak kodundaki **token türlerini ve konumlarını** temsil eder.
* **Token türleri** → `IDENT`, `INT`, `STRING`, `IF`, `FOR`, `+`, `-`, vs.
* **Konum bilgisi** → `Pos` (soyut), `Position` (dosya, satır, sütun), `FileSet` (birden fazla dosya yönetir).
* Kullanım alanı → Parser, AST, statik analiz, kod üretimi, gofmt benzeri araçlar.

---

👉 İstersen ben sana `go/token` paketini kullanarak **küçük bir kod analiz aracı** (örneğin: bir `.go` dosyasındaki tüm `if`, `for`, `func` anahtar kelimelerini ve yerlerini bulan) yazabilirim. İstiyor musun?
EVET
*/
/*
Süper 👍 O zaman beraber **`go/token` + `go/scanner`** kullanarak küçük bir **kod analiz aracı** yapalım.

📌 Bu araç:

* Bir `.go` dosyasını okuyacak
* Token’ları tarayacak
* Tüm **`if`**, **`for`** ve **`func`** anahtar kelimelerini bulacak
* Satır & sütun bilgisiyle birlikte ekrana yazacak

---

# 📂 Kod – Go Kod Analiz Aracı
*/
``go
package main

import (
	"fmt"
	"go/scanner"
	"go/token"
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

	// FileSet oluştur
	fset := token.NewFileSet()
	file := fset.AddFile(filename, -1, len(src))

	// Scanner başlat
	var s scanner.Scanner
	s.Init(file, src, nil, scanner.ScanComments)

	// Tokenları tara
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}

		// if, for, func arayalım
		if tok == token.IF || tok == token.FOR || tok == token.FUNC {
			fmt.Printf("%s bulundu: %q → %v\n", tok, lit, fset.Position(pos))
		}
	}
}
``
/*
---

# 📌 Örnek Test

`ornek.go` dosyamız:
*/
``go
package main

import "fmt"

func main() {
    for i := 0; i < 3; i++ {
        if i%2 == 0 {
            fmt.Println("çift:", i)
        }
    }
}
``

//Çalıştır:

``bash
go run main.go ornek.go
``

/7📌 Çıktı:

``
func bulundu: "func" → ornek.go:5:1
for bulundu: "for" → ornek.go:6:5
if bulundu: "if" → ornek.go:7:9
``
/*
---

# 🚀 Özet

* `go/token` → **token türleri + konum bilgisi** sağladı.
* `go/scanner` → Kaynak kodu token’lara ayırdı.
* Program, `if`, `for`, `func` anahtar kelimelerini ve **dosya\:satır\:sütun** bilgilerini buldu.

---

👉 İstersen bu aracı **geliştirip:**

* Bir dosyadaki tüm **değişken adlarını** toplama
* Fonksiyonların **başlangıç ve bitiş satırlarını** gösterme
* Kodun **karmaşıklığını ölçme (kaç tane if, for, switch var)**

gibi özellikler ekleyebiliriz.

Bunlardan hangisini yapalım?
EVET
*/
/*
Süper 👍 O zaman sana **tek bir araç** yazıyorum:

📌 Bu araç şunları yapacak:

1. Tüm **fonksiyonları** bulacak → başlangıç & bitiş satırlarını gösterecek
2. Tüm **değişken adlarını** çıkaracak
3. **Kontrol yapıları** (`if`, `for`, `switch`) sayısını hesaplayacak

Böylece küçük bir **Go kod analiz aracı** yapmış olacağız 🚀

---

# 📂 Kod – Go Kod Analiz Aracı
*/
``go
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
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

	// FileSet oluştur
	fset := token.NewFileSet()

	// Parse et
	node, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		fmt.Println("Parse hatası:", err)
		return
	}

	// Sayaclar
	ifCount, forCount, switchCount := 0, 0, 0
	var vars []string
	var funcs []string

	// AST üzerinde gezin
	ast.Inspect(node, func(n ast.Node) bool {
		switch stmt := n.(type) {
		case *ast.FuncDecl:
			// Fonksiyon başlangıç & bitiş satırlarını al
			start := fset.Position(stmt.Pos())
			end := fset.Position(stmt.End())
			funcs = append(funcs, fmt.Sprintf("%s [%d:%d - %d:%d]",
				stmt.Name.Name, start.Line, start.Column, end.Line, end.Column))

		case *ast.AssignStmt: // Değişken atamaları
			for _, expr := range stmt.Lhs {
				if ident, ok := expr.(*ast.Ident); ok {
					vars = append(vars, ident.Name)
				}
			}

		case *ast.ValueSpec: // Var tanımları
			for _, ident := range stmt.Names {
				vars = append(vars, ident.Name)
			}

		case *ast.IfStmt:
			ifCount++
		case *ast.ForStmt:
			forCount++
		case *ast.SwitchStmt:
			switchCount++
		}
		return true
	})

	// 📌 Sonuçları yazdır
	fmt.Println("Fonksiyonlar:")
	for _, f := range funcs {
		fmt.Println("  -", f)
	}

	fmt.Println("\nDeğişkenler:")
	for _, v := range vars {
		fmt.Println("  -", v)
	}

	fmt.Println("\nKontrol Yapıları:")
	fmt.Printf("  if: %d\n", ifCount)
	fmt.Printf("  for: %d\n", forCount)
	fmt.Printf("  switch: %d\n", switchCount)
}
``
/*
---

# 📌 Örnek Test

`ornek.go` dosyamız:
*/
``go
package main

import "fmt"

func add(a, b int) int {
    result := a + b
    if result > 10 {
        fmt.Println("Büyük")
    }
    return result
}

func main() {
    x := 5
    y := 7
    total := add(x, y)
    for i := 0; i < 3; i++ {
        switch i {
        case 0:
            fmt.Println("Sıfır")
        default:
            fmt.Println("Diğer")
        }
    }
    fmt.Println("Toplam:", total)
}
``

/7Çalıştır:

``bash
go run analyzer.go ornek.go
``

//📌 Çıktı:

``
Fonksiyonlar:
  - add [5:1 - 11:2]
  - main [13:1 - 25:2]

Değişkenler:
  - result
  - x
  - y
  - total
  - i

Kontrol Yapıları:
  if: 1
  for: 1
  switch: 1
``
/*
---

# 🚀 Özet

* **Fonksiyon analizi** → Başlangıç ve bitiş satırlarını bulduk ✅
* **Değişken analizi** → Tüm değişken adlarını listeledik ✅
* **Kontrol yapıları** → if, for, switch saydık ✅

---

👉 İstersen bunu daha da geliştirip:

* Hangi **fonksiyon içinde** kaç tane `if`, `for` olduğunu gösterebiliriz
* Fonksiyonların **parametre sayısını** yazabiliriz
* Bir **JSON raporu** oluşturup saklayabiliriz

Bunlardan hangisini yapmamı istersin?
*/