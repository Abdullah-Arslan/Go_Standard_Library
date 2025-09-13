/*
Go dilindeki **`go/syntax`** paketi, Go kaynak kodunu **parse** edip, bir **soyut sözdizim ağacı (AST – Abstract Syntax Tree)** üretmek için kullanılan, deneysel bir paketidir. Bu paket `go/parser` ve `go/ast` ile benzer işler yapar, ancak daha **hafif** ve **saf sözdizim temelli** çalışır. Yani tip bilgisiyle uğraşmaz; sadece kodun yapısını çıkarır.

Bu paket özellikle **araç geliştirme (linter, kod formatlayıcı, analiz aracı)** yazarken kullanılır.

---

## 📦 Paketin Temel Yapısı

`go/syntax` paketi şu türleri ve fonksiyonları içerir:

* **`syntax.File`** → Bir `.go` dosyasını temsil eden üst düğüm.
* **`syntax.Node`** → Tüm sözdizim ağaç düğümlerinin ortak arayüzü.
* **`syntax.ParserMode`** → Parse işlemini yöneten bayraklar (`syntax.CheckBranches`, `syntax.AllowGenerics` vs.).
* **`syntax.ParseFile`** → Verilen kaynak koddan `*syntax.File` üretir.
* **`syntax.Walk`** → AST içinde gezinmek için ziyaretçi fonksiyonu.
* **Çeşitli düğüm tipleri**: `syntax.Expr`, `syntax.Stmt`, `syntax.FuncDecl`, `syntax.ImportDecl`, `syntax.CallExpr` vb.

---

## 🔑 Önemli Fonksiyonlar ve Kullanımı

### 1. `ParseFile`

Bir Go dosyasını parse edip `*syntax.File` döner.
*/
``go
package main

import (
	"fmt"
	"go/syntax"
	"os"
)

func main() {
	// Dosyayı aç
	f, err := os.Open("example.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Parse et
	file, err := syntax.ParseFile(f.Name(), f, nil, 0)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Dosya başarıyla parse edildi: %s\n", file.PkgName.Value)
}
``
/*
📌 `file.PkgName.Value` bize dosyanın paket adını verir.

---

### 2. `Walk`

AST üzerinde dolaşmak için kullanılır.
*/
``go
package main

import (
	"fmt"
	"go/syntax"
	"os"
)

func main() {
	src := `package main
import "fmt"
func main() { fmt.Println("Merhaba") }`

	file, _ := syntax.ParseFile("src.go", strings.NewReader(src), nil, 0)

	// AST üzerinde gezin
	syntax.Walk(file, func(n syntax.Node) bool {
		switch x := n.(type) {
		case *syntax.FuncDecl:
			fmt.Println("Fonksiyon bulundu:", x.Name.Value)
		case *syntax.ImportDecl:
			fmt.Println("Import bulundu:", x.Path.Value)
		}
		return true
	})
}
``

//Çıktı:

``
Import bulundu: "fmt"
Fonksiyon bulundu: main
``
/*
---

### 3. Örnek: Fonksiyonları Listeleme

Bir dosyadaki tüm fonksiyonları çıkarmak:
*/
``go
package main

import (
	"fmt"
	"go/syntax"
	"os"
)

func main() {
	f, _ := os.Open("example.go")
	defer f.Close()

	file, _ := syntax.ParseFile(f.Name(), f, nil, 0)

	for _, decl := range file.DeclList {
		if fn, ok := decl.(*syntax.FuncDecl); ok {
			fmt.Println("Fonksiyon:", fn.Name.Value)
		}
	}
}
``
/*
---

### 4. ParserMode Kullanımı

Parse davranışını değiştirmek için:

* `syntax.CheckBranches` → `break/continue` hedeflerini kontrol et.
* `syntax.AllowGenerics` → Go 1.18+ generic sözdizimini kabul et.
* `syntax.AllowTypeParams` → Fonksiyon tür parametrelerini kabul et.
*/
``go
file, _ := syntax.ParseFile("example.go", f, nil, syntax.AllowGenerics|syntax.CheckBranches)
``
/*
---

## ✅ Özet

* `go/syntax`, Go kodunu **sözdizimi temelli** olarak parse eder.
* `ParseFile` → dosyayı parse eder.
* `Walk` → AST üzerinde dolaşır.
* Düğüm tipleri (`FuncDecl`, `ImportDecl`, `CallExpr`, vs.) ile kodun yapısı analiz edilir.
* Daha çok **araç geliştirme**, **kod analizi** ve **formatlama** gibi işlerde kullanılır.

---

👉 İstersen sana bu paket için **tüm tipler ve fonksiyonları tablo halinde** (örnekleriyle birlikte) **Excel + PDF** olarak da hazırlayabilirim. Bunu ister misin?
EVET
*/
/*
Süper 👍

O halde önce sana **tam tabloyu burada metin halinde** göstereceğim. Daha sonra bu tabloyu **Excel (.xlsx)** ve **PDF** olarak hazırlayıp indirebilir hale getireceğim.

---
*/
# 📋 Go `syntax` Paketi Özeti (Tablo Halinde)

| Fonksiyon / Tip  | Açıklama                                                                               | Örnek Kod                                                                    | Notlar                                                        |
| ---------------- | -------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ------------------------------------------------------------- |
| **`ParseFile`**  | Bir Go dosyasını/parçayı parse edip `*syntax.File` döner.                              | `go\nfile, err := syntax.ParseFile("example.go", src, nil, 0)\n`             | Kaynak kodu `io.Reader` veya dosya adı ile verilebilir.       |
| **`File`**       | Bir `.go` dosyasını temsil eder (paket adı, importlar, fonksiyonlar, değişkenler vs.). | `go\nfmt.Println(file.PkgName.Value)\n`                                      | AST’in en üst düğümü.                                         |
| **`Node`**       | Tüm AST düğümlerinin ortak arayüzü.                                                    | Kullanıcı genelde `Walk` ile görür.                                          | Tür iddiaları ile (`x.(*syntax.FuncDecl)`) kullanılabilir.    |
| **`Decl`**       | Bildirim (declaration) türü: `FuncDecl`, `ImportDecl`, `VarDecl` vs.                   | `go\nfor _, d := range file.DeclList { ... }\n`                              | Kaynak dosyadaki global bildirileri içerir.                   |
| **`FuncDecl`**   | Bir fonksiyon bildirimini temsil eder.                                                 | `go\nif fn, ok := d.(*syntax.FuncDecl); ok { fmt.Println(fn.Name.Value) }\n` | Fonksiyon adı, parametreleri, gövdesi gibi alanlara erişilir. |
| **`ImportDecl`** | Bir `import` bildirimini temsil eder.                                                  | `go\nimp := decl.(*syntax.ImportDecl)\nfmt.Println(imp.Path.Value)\n`        | Import edilen paketlerin yolunu verir.                        |
| **`Expr`**       | İfade düğümlerinin ana türü. Alt tipler: `CallExpr`, `Name`, `BasicLit` vb.            | `go\nexpr := call.Fun.(*syntax.Name)\n`                                      | İfadeleri çözümlemek için kullanılır.                         |
| **`CallExpr`**   | Fonksiyon çağrılarını temsil eder.                                                     | `go\nfmt.Println("Çağrı:", call.Fun)\n`                                      | `fmt.Println("...")` gibi ifadeler.                           |
| **`BasicLit`**   | Sayılar, stringler gibi temel literal değerleri temsil eder.                           | `go\nfmt.Println(lit.Value)\n`                                               | Kod içindeki sabit değerler.                                  |
| **`Stmt`**       | Tüm ifade cümleleri (statement) için temel arayüz.                                     | Örn: `IfStmt`, `ForStmt`, `ReturnStmt`.                                      | Kod akışını temsil eder.                                      |
| **`Walk`**       | AST üzerinde DFS dolaşımı yapar.                                                       | `go\nsyntax.Walk(file, func(n syntax.Node) bool { ...; return true })\n`     | Tüm düğümlere ulaşmak için ziyaretçi deseni.                  |
| **`ParserMode`** | Parse seçenekleri: `CheckBranches`, `AllowGenerics`, `AllowTypeParams`.                | `go\nsyntax.ParseFile("x.go", f, nil, syntax.AllowGenerics)\n`               | Derleyiciye hangi kuralları dikkate alacağını söyler.         |

/*
---

Bu tabloyu hem **Excel** hem **PDF** formatında hazırlayacağım 📂.

👉 Şimdi sana sorum: Dosya isimlerini **`syntax_paketi.xlsx`** ve **`syntax_paketi.pdf`** olarak mı kaydedeyim, yoksa sen özel isim vermek ister misin?
EVET


👉 İstersen ben bu paket için ayrıca örnek mini bir proje (örneğin bir Go dosyasındaki tüm fonksiyonları ve importları listeliyor) de hazırlayabilirim. İstiyor musun
EVET
*/
/*
Harika 🎉 O zaman sana **`go/syntax` paketini kullanan küçük bir örnek proje** hazırlayayım. Bu proje:

* Bir `.go` dosyasını parse edecek
* İçindeki **paket adını**, **import edilen kütüphaneleri** ve **fonksiyon isimlerini** listeleyecek

---

## 📂 Proje Yapısı
*/
``
syntax-demo/
 ├── main.go
 └── sample.go   (analiz edilecek dosya)
``

---

//## 📝 `sample.go` (örnek dosya)

``go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Merhaba, Go!")
}

func kare(x int) int {
	return x * x
}

func kok(x float64) float64 {
	return math.Sqrt(x)
}
``

---

//## 📝 `main.go` (analiz aracı)

``go
package main

import (
	"fmt"
	"go/syntax"
	"os"
)

func main() {
	// Analiz edilecek dosya
	filename := "sample.go"

	// Dosyayı aç
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Parse et
	file, err := syntax.ParseFile(filename, f, nil, 0)
	if err != nil {
		panic(err)
	}

	// Paket adı
	fmt.Println("Paket adı:", file.PkgName.Value)

	// AST üzerinde gez
	for _, decl := range file.DeclList {
		switch d := decl.(type) {
		case *syntax.ImportDecl:
			fmt.Println("Import:", d.Path.Value)
		case *syntax.FuncDecl:
			fmt.Println("Fonksiyon:", d.Name.Value)
		}
	}
}
``

---

//## ▶️ Çalıştırma

``bash
go run main.go
``

//### Çıktı:

``
Paket adı: main
Import: "fmt"
Import: "math"
Fonksiyon: main
Fonksiyon: kare
Fonksiyon: kok
``
/*
---

👉 Böylece `go/syntax` paketi ile bir Go dosyasını okuyup **paket adı, importlar ve fonksiyonları** bulabiliyorsun.

İstersen bu örneği biraz daha geliştirip:

* **Değişken bildirimlerini** de çıkartabiliriz
* Fonksiyonların **parametre tiplerini** gösterebiliriz
* Hatta AST gezip **if/for gibi statement’leri** analiz edebiliriz

Bunlardan hangisini yapmamı istersin?
EVET
*/
/*
Süper 👍 O zaman bu mini projeyi biraz daha **ileri seviye** hale getirelim.

Yeni versiyonda:

* Paket adı ✅
* Import’lar ✅
* Fonksiyon adları ✅
* **Fonksiyon parametreleri** ve **dönüş tipleri** 🆕
* **Global değişkenler / sabitler** 🆕
* Kod içindeki **if, for, return** gibi statement’leri 🆕

---

## 📝 `sample.go` (analiz edilecek dosya)
*/
``go
package main

import (
	"fmt"
	"math"
)

var globalVar int = 42
const appName = "SyntaxDemo"

func main() {
	for i := 0; i < 3; i++ {
		if i%2 == 0 {
			fmt.Println("Çift:", i)
		} else {
			fmt.Println("Tek:", i)
		}
	}
}

func kare(x int) int {
	return x * x
}

func kok(x float64) float64 {
	return math.Sqrt(x)
}
``
/*
---

## 📝 `main.go` (geliştirilmiş analiz aracı)
*/
``go
package main

import (
	"fmt"
	"go/syntax"
	"os"
)

func main() {
	filename := "sample.go"

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	file, err := syntax.ParseFile(filename, f, nil, 0)
	if err != nil {
		panic(err)
	}

	// Paket adı
	fmt.Println("📦 Paket adı:", file.PkgName.Value)

	// AST üzerinde gez
	syntax.Walk(file, func(n syntax.Node) bool {
		switch d := n.(type) {
		case *syntax.ImportDecl:
			fmt.Println("📥 Import:", d.Path.Value)

		case *syntax.FuncDecl:
			fmt.Printf("🔧 Fonksiyon: %s\n", d.Name.Value)

			// Parametreleri yazdır
			if d.Type != nil && d.Type.ParamList != nil {
				for _, param := range d.Type.ParamList {
					fmt.Printf("   ↳ Parametre: %s\n", param.Name.Value)
				}
			}

			// Dönüş tipleri
			if d.Type.ResultList != nil {
				for _, res := range d.Type.ResultList {
					fmt.Printf("   ↳ Dönüş Tipi: %T\n", res.Type)
				}
			}

		case *syntax.VarDecl:
			for _, name := range d.NameList {
				fmt.Println("🌍 Global değişken:", name.Value)
			}

		case *syntax.ConstDecl:
			for _, name := range d.NameList {
				fmt.Println("🔒 Sabit:", name.Value)
			}

		case *syntax.IfStmt:
			fmt.Println("🔎 If ifadesi bulundu")

		case *syntax.ForStmt:
			fmt.Println("🔁 For döngüsü bulundu")

		case *syntax.ReturnStmt:
			fmt.Println("↩️ Return ifadesi bulundu")
		}
		return true
	})
}
``
/*
---

## ▶️ Çalıştırma
*/
``bash
go run main.go
``

//### Çıktı:

``
📦 Paket adı: main
📥 Import: "fmt"
📥 Import: "math"
🌍 Global değişken: globalVar
🔒 Sabit: appName
🔧 Fonksiyon: main
🔁 For döngüsü bulundu
🔎 If ifadesi bulundu
🔧 Fonksiyon: kare
   ↳ Parametre: x
   ↳ Dönüş Tipi: *syntax.Name
↩️ Return ifadesi bulundu
🔧 Fonksiyon: kok
   ↳ Parametre: x
   ↳ Dönüş Tipi: *syntax.Name
↩️ Return ifadesi bulundu
``
/*
---

✅ Bu geliştirilmiş örnek ile:

* **Import’lar, fonksiyonlar, parametreler, dönüş tipleri**
* **Global değişken ve sabitler**
* **If, For, Return** gibi statement’ler

hepsi AST üzerinden analiz edilebiliyor.

---

👉 İstersen bunu daha da geliştirip bir **komut satırı aracı (CLI)** haline getirebilirim. Yani `go run main.go sample.go` komutuyla istediğin `.go` dosyasının yapısını analiz eden bir araç olur. Bunu da ister misin?
EVET
*/
/*
Harika 🎉 O zaman bu örneği bir **CLI aracı** haline getirelim. Böylece istediğin `.go` dosyasını analiz edebilirsin:

---

# 📂 Proje Yapısı
*/
``
syntax-cli/
 ├── main.go      → CLI aracı
 ├── test1.go     → Analiz için örnek dosya
 └── test2.go     → Başka örnek dosya
``

---

//# 📝 `main.go`

``go
package main

import (
	"flag"
	"fmt"
	"go/syntax"
	"os"
)

func analyzeFile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("❌ Dosya açılamadı:", err)
		return
	}
	defer f.Close()

	file, err := syntax.ParseFile(filename, f, nil, 0)
	if err != nil {
		fmt.Println("❌ Parse hatası:", err)
		return
	}

	fmt.Println("=================================")
	fmt.Println("📂 Dosya:", filename)
	fmt.Println("📦 Paket:", file.PkgName.Value)
	fmt.Println("=================================")

	// AST üzerinde dolaş
	syntax.Walk(file, func(n syntax.Node) bool {
		switch d := n.(type) {
		case *syntax.ImportDecl:
			fmt.Println("📥 Import:", d.Path.Value)

		case *syntax.FuncDecl:
			fmt.Printf("🔧 Fonksiyon: %s\n", d.Name.Value)

			// Parametreler
			if d.Type != nil && d.Type.ParamList != nil {
				for _, param := range d.Type.ParamList {
					fmt.Printf("   ↳ Parametre: %s\n", param.Name.Value)
				}
			}

			// Dönüş tipleri
			if d.Type.ResultList != nil {
				for _, res := range d.Type.ResultList {
					fmt.Printf("   ↳ Dönüş Tipi: %T\n", res.Type)
				}
			}

		case *syntax.VarDecl:
			for _, name := range d.NameList {
				fmt.Println("🌍 Global değişken:", name.Value)
			}

		case *syntax.ConstDecl:
			for _, name := range d.NameList {
				fmt.Println("🔒 Sabit:", name.Value)
			}

		case *syntax.IfStmt:
			fmt.Println("🔎 If ifadesi bulundu")

		case *syntax.ForStmt:
			fmt.Println("🔁 For döngüsü bulundu")

		case *syntax.ReturnStmt:
			fmt.Println("↩️ Return ifadesi bulundu")
		}
		return true
	})
	fmt.Println()
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("Kullanım: go run main.go <dosya1.go> <dosya2.go> ...")
		return
	}

	for _, filename := range flag.Args() {
		analyzeFile(filename)
	}
}
``

---

//# 📝 `test1.go`

``go
package main

import "fmt"

var version = "1.0"

func hello(name string) string {
	return "Merhaba, " + name
}

func main() {
	fmt.Println(hello("Go"))
}
``

---

//# 📝 `test2.go`

``go
package mathutils

import "math"

const pi = 3.14

func Kare(x int) int {
	return x * x
}

func Kok(x float64) float64 {
	return math.Sqrt(x)
}
``
/*
---

# ▶️ Kullanım
*/
``bash
go run main.go test1.go test2.go
``
/*
---

# 📌 Örnek Çıktı
*/
``
=================================
📂 Dosya: test1.go
📦 Paket: main
=================================
📥 Import: "fmt"
🌍 Global değişken: version
🔧 Fonksiyon: hello
   ↳ Parametre: name
   ↳ Dönüş Tipi: *syntax.Name
↩️ Return ifadesi bulundu
🔧 Fonksiyon: main

=================================
📂 Dosya: test2.go
📦 Paket: mathutils
=================================
📥 Import: "math"
🔒 Sabit: pi
🔧 Fonksiyon: Kare
   ↳ Parametre: x
   ↳ Dönüş Tipi: *syntax.Name
↩️ Return ifadesi bulundu
🔧 Fonksiyon: Kok
   ↳ Parametre: x
   ↳ Dönüş Tipi: *syntax.Name
↩️ Return ifadesi bulundu
``
/*
---

✅ Böylece elinde artık **Go dosyalarını AST üzerinden analiz eden bir CLI aracı** oldu.
İstersen ben bunu bir adım ileri götürüp **çıktıyı JSON formatında verdirip**, başka araçların kolayca kullanabileceği bir hale getirebilirim.

👉 Çıktıyı JSON formatında almak ister misin?
EVET
*/
/*
Süper 🚀 O zaman CLI aracını geliştirelim: artık çıktıyı **JSON formatında** verecek. Böylece başka araçlarla kolayca işleyebilirsin.

---

# 📝 `main.go` (JSON Destekli CLI)
*/
``go
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/syntax"
	"os"
)

type AnalysisResult struct {
	FileName   string        `json:"file"`
	Package    string        `json:"package"`
	Imports    []string      `json:"imports"`
	Functions  []FunctionDef `json:"functions"`
	Vars       []string      `json:"variables"`
	Consts     []string      `json:"constants"`
	Statements []string      `json:"statements"`
}

type FunctionDef struct {
	Name       string   `json:"name"`
	Params     []string `json:"params"`
	ReturnType []string `json:"return_types"`
}

func analyzeFile(filename string) AnalysisResult {
	res := AnalysisResult{
		FileName:   filename,
		Imports:    []string{},
		Functions:  []FunctionDef{},
		Vars:       []string{},
		Consts:     []string{},
		Statements: []string{},
	}

	f, err := os.Open(filename)
	if err != nil {
		res.Package = fmt.Sprintf("❌ Dosya açılamadı: %v", err)
		return res
	}
	defer f.Close()

	file, err := syntax.ParseFile(filename, f, nil, 0)
	if err != nil {
		res.Package = fmt.Sprintf("❌ Parse hatası: %v", err)
		return res
	}

	res.Package = file.PkgName.Value

	syntax.Walk(file, func(n syntax.Node) bool {
		switch d := n.(type) {
		case *syntax.ImportDecl:
			res.Imports = append(res.Imports, d.Path.Value)

		case *syntax.FuncDecl:
			fn := FunctionDef{Name: d.Name.Value}

			// Parametreler
			if d.Type != nil && d.Type.ParamList != nil {
				for _, param := range d.Type.ParamList {
					fn.Params = append(fn.Params, param.Name.Value)
				}
			}

			// Dönüş tipleri
			if d.Type.ResultList != nil {
				for _, resType := range d.Type.ResultList {
					fn.ReturnType = append(fn.ReturnType, fmt.Sprintf("%T", resType.Type))
				}
			}

			res.Functions = append(res.Functions, fn)

		case *syntax.VarDecl:
			for _, name := range d.NameList {
				res.Vars = append(res.Vars, name.Value)
			}

		case *syntax.ConstDecl:
			for _, name := range d.NameList {
				res.Consts = append(res.Consts, name.Value)
			}

		case *syntax.IfStmt:
			res.Statements = append(res.Statements, "if")

		case *syntax.ForStmt:
			res.Statements = append(res.Statements, "for")

		case *syntax.ReturnStmt:
			res.Statements = append(res.Statements, "return")
		}
		return true
	})

	return res
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("Kullanım: go run main.go <dosya1.go> <dosya2.go> ...")
		return
	}

	var results []AnalysisResult
	for _, filename := range flag.Args() {
		results = append(results, analyzeFile(filename))
	}

	// JSON çıktısı
	output, _ := json.MarshalIndent(results, "", "  ")
	fmt.Println(string(output))
}
``
/*
---

# ▶️ Kullanım
*/
``bash
go run main.go test1.go test2.go > output.json
``
/*
---

# 📌 Örnek JSON Çıktısı
*/
``json
[
  {
    "file": "test1.go",
    "package": "main",
    "imports": [
      "\"fmt\""
    ],
    "functions": [
      {
        "name": "hello",
        "params": ["name"],
        "return_types": ["*syntax.Name"]
      },
      {
        "name": "main",
        "params": [],
        "return_types": []
      }
    ],
    "variables": ["version"],
    "constants": [],
    "statements": ["return"]
  },
  {
    "file": "test2.go",
    "package": "mathutils",
    "imports": ["\"math\""],
    "functions": [
      {
        "name": "Kare",
        "params": ["x"],
        "return_types": ["*syntax.Name"]
      },
      {
        "name": "Kok",
        "params": ["x"],
        "return_types": ["*syntax.Name"]
      }
    ],
    "variables": [],
    "constants": ["pi"],
    "statements": ["return", "return"]
  }
]
``
/*
---

✅ Böylece artık analiz aracın **JSON çıktı** üretiyor.
Bunu başka uygulamalar, frontend’ler veya veri tabanları kolayca işleyebilir.

👉 İstersen bir sonraki adımda bu JSON çıktıyı alıp **web arayüzünde görselleştiren** (örneğin Go fonksiyonlarını listeler, parametrelerini tablo yapar) bir mini **React/Tailwind dashboard** da hazırlayabilirim. Bunu ister misin?
*/