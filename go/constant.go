/*
**Go’nun standart kütüphanesinde `constant` adında bir paket yok.**

Go’da `constant` kavramı var (sabit değerler, `const` anahtar kelimesi ile tanımlanır).
Ama senin kastettiğin büyük ihtimalle:

1. **`go/constant` paketi** → Go **sabit ifadeleri (constants)** üzerinde çalışmayı sağlar. Özellikle `go/types` ve `go/ast` ile birlikte kullanılır.
2. Yoksa sadece `const` kullanımı sorulmuş olabilir.

Ben sana **`go/constant` paketini** detaylı anlatayım, çünkü bu gerçekten bir standart paket ve sık sorulan bir konu.

---

# 📦 `go/constant` Paketi

* Paket yolu:
*/
  ``go
  import "go/constant"
  ``
/*
  * Görevi: **Go sabitlerini (constant expressions)** temsil etmek ve üzerinde işlem yapmaktır.
* Bu paket genellikle Go’nun derleyici araç zincirinde (compiler toolchain) veya statik analiz araçlarında kullanılır.

`constant.Value` adında bir tip vardır → bu, bir sabiti soyutlar.
Bir sabit: **bool, string, int, float, complex** olabilir.

---

## 1️⃣ Sabit Oluşturma (`MakeX` Fonksiyonları)

* `constant.MakeBool(true)` → boolean sabit üretir
* `constant.MakeString("merhaba")` → string sabit üretir
* `constant.MakeInt64(42)` → integer sabit üretir
* `constant.MakeFloat64(3.14)` → float sabit üretir
* `constant.MakeComplex(real, imag)` → karmaşık sayı sabiti üretir

📌 Örnek:
*/
``go
package main

import (
	"fmt"
	"go/constant"
)

func main() {
	v1 := constant.MakeBool(true)
	v2 := constant.MakeString("Go Sabiti")
	v3 := constant.MakeInt64(123)

	fmt.Println(v1, v2, v3) // true "Go Sabiti" 123
}
``
/*
---

## 2️⃣ Sabit Türünü Kontrol Etme

* `constant.BoolVal(x)` → bool değerini döner
* `constant.StringVal(x)` → string değerini döner
* `constant.Int64Val(x)` → int64 değerini döner
* `constant.Float64Val(x)` → float64 değerini döner
* `constant.Kind(x)` → sabitin tipini döner (`Bool`, `String`, `Int`, `Float`, `Complex`)

📌 Örnek:
*/

``go
package main

import (
	"fmt"
	"go/constant"
)

func main() {
	v := constant.MakeFloat64(2.71)

	fmt.Println("Tür:", v.Kind())          // Float
	fmt.Println("Değer:", constant.Float64Val(v)) // 2.71
}
``
/*
---

## 3️⃣ Aritmetik ve Mantıksal İşlemler

* `constant.BinaryOp(x, y, token)` → iki sabit arasında işlem yapar
* `constant.UnaryOp(token, x, prec)` → tekli işlem (örneğin `-x`)
* `constant.Compare(x, y, token)` → karşılaştırma

📌 Örnek:
*/

``go
package main

import (
	"fmt"
	"go/constant"
	"go/token"
)

func main() {
	a := constant.MakeInt64(10)
	b := constant.MakeInt64(3)

	// Toplama
	sum := constant.BinaryOp(a, token.ADD, b)
	fmt.Println("10 + 3 =", sum) // 13

	// Çarpma
	mul := constant.BinaryOp(a, token.MUL, b)
	fmt.Println("10 * 3 =", mul) // 30

	// Karşılaştırma
	eq := constant.Compare(a, token.EQL, b)
	fmt.Println("10 == 3 ?", eq) // false
}
``
/*
---

## 4️⃣ Karmaşık Sayılar (Complex Numbers)

`MakeComplex` ile complex sabit oluşturabilirsin.

📌 Örnek:
*/
``go
package main

import (
	"fmt"
	"go/constant"
)

func main() {
	real := constant.MakeInt64(2)
	imag := constant.MakeInt64(3)

	complexVal := constant.MakeComplex(real, imag)
	fmt.Println("Karmaşık sayı:", complexVal) // (2 + 3i)
}
``
/*
---

## 5️⃣ Rasyonel Sayılar (Big Numbers)

`go/constant` arka planda `math/big` kullanır. Çok büyük sayılarla da çalışabilir.

📌 Örnek:
*/
``go
package main

import (
	"fmt"
	"go/constant"
	"math/big"
)

func main() {
	// Büyük bir sayı
	bigInt := big.NewInt(1)
	bigInt.Lsh(bigInt, 100) // 2^100

	val := constant.MakeFromLiteral(bigInt.String(), token.INT, 0)
	fmt.Println("Büyük sayı:", val)
}
``
/*
---

## 6️⃣ Literallerden Sabit Yapmak

* `constant.MakeFromLiteral(lit, token, zero)`
  → string literali parse edip sabit üretir.

📌 Örnek:
*/
``go
package main

import (
	"fmt"
	"go/constant"
	"go/token"
)

func main() {
	v := constant.MakeFromLiteral("42", token.INT, 0)
	fmt.Println("Literal sabit:", v) // 42
}
``
/*
---

# 🚀 Özet Tablo

| Fonksiyon / Tip                          | Açıklama                                   |
| ---------------------------------------- | ------------------------------------------ |
| `constant.MakeBool(true)`                | Bool sabit oluştur                         |
| `constant.MakeString("x")`               | String sabit oluştur                       |
| `constant.MakeInt64(42)`                 | Int sabit oluştur                          |
| `constant.MakeFloat64(3.14)`             | Float sabit oluştur                        |
| `constant.MakeComplex(a, b)`             | Complex sabit oluştur                      |
| `constant.MakeFromLiteral(lit, kind, 0)` | String literali sabite dönüştür            |
| `constant.Kind(v)`                       | Sabitin türünü döner                       |
| `constant.Int64Val(v)`                   | int64 değerini döner                       |
| `constant.Float64Val(v)`                 | float64 değerini döner                     |
| `constant.StringVal(v)`                  | string değerini döner                      |
| `constant.BinaryOp(a, op, b)`            | İkili işlem (toplama, çarpma, çıkarma vb.) |
| `constant.UnaryOp(op, x, prec)`          | Tekli işlem (ör. `-x`)                     |
| `constant.Compare(a, op, b)`             | Karşılaştırma                              |

---

✅ Yani `go/constant` paketi, **Go sabitlerini derleyici gibi analiz etmek ve üzerinde işlem yapmak için** kullanılır.
Genellikle `go/types`, `go/ast` ve `go/parser` ile birlikte **kod analiz araçları** yazarken çok işine yarar.

---

İstersen sana küçük bir **örnek araç** yazayım:
→ Bir Go dosyasını parse etsin, içindeki `const` tanımlarını bulsun, ve `go/constant` ile bunların değerlerini göstersin.

Bunu ister misin?
EVET
*/
/*
Harika 👍 o zaman sana **küçük bir araç** yazalım:

Bu araç:

1. Bir Go dosyasını `parser` ile parse edecek
2. İçindeki `const` tanımlarını bulacak (`ast.GenDecl`)
3. `go/constant` ile değerlerini çözecek
4. Ekrana yazdıracak

---

## 📌 Örnek Go Dosyası (`example.go`)
*/
``go
package main

const Pi = 3.14159
const Hello = "Merhaba"
const Flag = true
const Big = 1 << 20 // 1 milyon civarı
``
/*
---

## 📌 Analiz Programı (`analyzer.go`)
*/
``go
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"go/constant"
	"log"
)

func main() {
	// Dosya konumu
	filename := "example.go"

	// Parse et
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	// Tip bilgisi çözümleyici
	conf := types.Config{Importer: nil}
	info := &types.Info{
		Defs: make(map[*ast.Ident]types.Object),
	}
	_, err = conf.Check("", fset, []*ast.File{file}, info)
	if err != nil {
		log.Fatal(err)
	}

	// AST içindeki const'ları tara
	for ident, obj := range info.Defs {
		if obj == nil {
			continue
		}
		if cnst, ok := obj.(*types.Const); ok {
			val := cnst.Val() // go/constant.Value
			// Sabitin türünü öğren
			kind := val.Kind()

			fmt.Printf("Const %s (%s) = %v\n", ident.Name, kind, val)

			// Eğer sayıysa normal değeri alalım
			switch kind {
			case constant.Int:
				i, _ := constant.Int64Val(val)
				fmt.Printf("  -> int64 değer: %d\n", i)
			case constant.Float:
				f, _ := constant.Float64Val(val)
				fmt.Printf("  -> float64 değer: %f\n", f)
			case constant.String:
				fmt.Printf("  -> string değer: %s\n", constant.StringVal(val))
			case constant.Bool:
				fmt.Printf("  -> bool değer: %t\n", constant.BoolVal(val))
			}
		}
	}
}
``
/*
---

## 📌 Çalıştırma
*/
``bash
go run analyzer.go
``

//📤 Çıktı:

``
Const Pi (Float) = 3.14159
  -> float64 değer: 3.141590
Const Hello (String) = "Merhaba"
  -> string değer: Merhaba
Const Flag (Bool) = true
  -> bool değer: true
Const Big (Int) = 1048576
  -> int64 değer: 1048576
``
/*
---

✅ Böylece elimizde bir **Go sabit analiz aracı** oldu.
Bu araç, `go/parser` + `go/types` + `go/constant` paketlerini birlikte kullanarak `const` tanımlarını buldu ve değerlerini çözdü.

---

👉 İstersen ben bu aracı biraz daha geliştirip:

* Bir **paketin tüm `.go` dosyalarını** taramasını,
* İçindeki bütün sabitleri bir tablo halinde çıkarmasını
  sağlayabilirim.

Bunu ister misin?
EVET
*/
/*
Harika 🙌 o zaman araçları biraz daha geliştirip **tüm paketi tarayan ve tüm sabitleri tablo halinde gösteren bir analiz aracı** yazalım.

---

# 📌 Gelişmiş Analiz Aracı (`analyze_constants.go`)

Bu araç:

1. Verilen proje dizinindeki tüm `.go` dosyalarını bulur
2. Her dosyayı parse eder
3. Tüm `const` tanımlarını `go/constant` ile çözer
4. Sonucu **tablo halinde** ekrana yazdırır
*/
``go
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"go/constant"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	root := "./myapp" // analiz etmek istediğin proje dizini

	fmt.Println("🔍 Go sabitleri analizi başlıyor...\n")

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			analyzeFile(path)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func analyzeFile(filename string) {
	fmt.Printf("📄 Dosya: %s\n", filename)

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		fmt.Println("  Hata:", err)
		return
	}

	conf := types.Config{Importer: nil}
	info := &types.Info{
		Defs: make(map[*ast.Ident]types.Object),
	}
	_, err = conf.Check("", fset, []*ast.File{file}, info)
	if err != nil {
		fmt.Println("  Tip kontrol hatası:", err)
		return
	}

	for ident, obj := range info.Defs {
		if obj == nil {
			continue
		}
		if cnst, ok := obj.(*types.Const); ok {
			val := cnst.Val() // go/constant.Value
			kind := val.Kind()
			fmt.Printf("  Const %-15s %-8s = %v\n", ident.Name, kind, val)

			switch kind {
			case constant.Int:
				i, _ := constant.Int64Val(val)
				fmt.Printf("    -> int64 değer: %d\n", i)
			case constant.Float:
				f, _ := constant.Float64Val(val)
				fmt.Printf("    -> float64 değer: %f\n", f)
			case constant.String:
				fmt.Printf("    -> string değer: %s\n", constant.StringVal(val))
			case constant.Bool:
				fmt.Printf("    -> bool değer: %t\n", constant.BoolVal(val))
			case constant.Complex:
				fmt.Printf("    -> complex değer: %v\n", val)
			}
		}
	}
	fmt.Println()
}
``
/*
---

# ⚙️ Kullanım

1. Projeyi hazırla, örneğin `myapp/` dizini içinde `.go` dosyaları olsun:
*/
``go
// myapp/constants.go
package myapp

const Pi = 3.14159
const Hello = "Merhaba"
const Flag = true
const Big = 1 << 20
``

//2. Araç dosyasını çalıştır:

``bash
go run analyze_constants.go
``
/*
---

# 🖥 Beklenen Çıktı
*/
``
🔍 Go sabitleri analizi başlıyor...

📄 Dosya: ./myapp/constants.go
  Const Pi              Float    = 3.14159
    -> float64 değer: 3.141590
  Const Hello           String   = "Merhaba"
    -> string değer: Merhaba
  Const Flag            Bool     = true
    -> bool değer: true
  Const Big             Int      = 1048576
    -> int64 değer: 1048576
``
/*
---

✅ Bu araç ile artık:

* Projedeki tüm `.go` dosyalarını tarayabiliyorsun
* Tüm `const` tanımlarını görebiliyorsun
* Tiplerini ve değerlerini rahatça inceleyebiliyorsun

---

İstersen bir sonraki adımda bunu **build tag’leri ile birleştirip**, sadece **aktif platformda derlenecek dosyalardaki sabitleri** analiz edecek şekilde geliştirebiliriz.

Bunu yapalım mı?
*/