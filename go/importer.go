/*
**`go/importer`** paketini detaylıca anlatayım.

---

# 📦 `go/importer` Paketi Nedir?

Go’nun **`go/importer`** paketi, `go/types` paketiyle birlikte çalışır.
Görevi: **import edilen paketlerin tip bilgilerini yüklemek**.

Normalde `go/types` yalnızca verdiğin `AST` üzerinden tip kontrolü yapar. Ama eğer kodda `import "fmt"` gibi başka paketler varsa, bunların tip bilgilerini bilmesi gerekir. İşte burada **`importer`** devreye girer.

---

# 🔑 Önemli Fonksiyonlar

### 1. `importer.Default()`

* Kullanılan Go sürümüne uygun varsayılan importer döndürür.
* Genelde bu yeterlidir.

### 2. `importer.For(compiler string, lookup Lookup) types.Importer`

* Belirli bir compiler için importer döndürür (`"gc"`, `"gccgo"`).
* `lookup` fonksiyonu ile paket bulma işini özelleştirebilirsin.

### 3. `type Lookup func(path string) (io.ReadCloser, error)`

* `For` fonksiyonuna verdiğin özel bir lookup fonksiyonu, paket bilgisini senin istediğin yerden (ör. önbellek, dosya, network) getirebilir.

---

# 📝 Basit Örnek: `fmt` paketini import edip tiplerini yazdırma
*/
``go
package main

import (
	"fmt"
	"go/importer"
	"go/types"
)

func main() {
	// Varsayılan importer al
	imp := importer.Default()

	// fmt paketini yükle
	pkg, err := imp.Import("fmt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Paket adı:", pkg.Name())
	fmt.Println("Dosya yolu:", pkg.Path())

	// İçindeki exported (dışa açık) tanımları yazdıralım
	scope := pkg.Scope()
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		fmt.Printf("%s: %s\n", name, obj.Type())
	}
}
``

//### 📌 Çıktı (kısaltılmış):

``
Paket adı: fmt
Dosya yolu: fmt
Errorf: func(format string, a ...interface{}) error
Println: func(a ...interface{}) (n int, err error)
Sprint: func(a ...interface{}) string
...
``
/*
Yani `importer` sayesinde Go’nun kendi paketlerindeki tip bilgilerine erişebiliyoruz.

---

# 📝 Örnek 2: `gc` derleyicisi için özel importer
*/
``go
package main

import (
	"fmt"
	"go/importer"
	"go/types"
)

func main() {
	// "gc" compiler için importer al
	imp := importer.For("gc", nil)

	// math paketini yükle
	pkg, err := imp.Import("math")
	if err != nil {
		panic(err)
	}

	fmt.Println("Paket:", pkg.Path())
	fmt.Println("Scope:")
	for _, name := range pkg.Scope().Names() {
		obj := pkg.Scope().Lookup(name)
		fmt.Println("-", name, ":", obj.Type())
	}
}
``
/*
Bu örnekte `math` paketindeki tüm fonksiyonları, sabitleri, tipleri görebiliriz.

---

# 📝 Örnek 3: `types.Config` ile birlikte kullanmak

En çok kullanılan senaryo şudur:
AST parse edilir, sonra `go/types.Config` tip denetiminde `Importer` alanına `importer.Default()` verilir. Böylece `fmt`, `math` gibi paketlerin tipleri de tanınır.
*/
``go
package main

import (
	"fmt"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
)

func main() {
	src := `package main
import "fmt"
func main() {
	fmt.Println("Hello")
}`

	// Parse et
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "example.go", src, 0)
	if err != nil {
		panic(err)
	}

	// Tip kontrolü
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		Uses: make(map[*types.Ident]types.Object),
	}
	_, err = conf.Check("main", fset, []*ast.File{f}, info)
	if err != nil {
		panic(err)
	}

	// fmt.Println’ın tipini bul
	for ident, obj := range info.Uses {
		if ident.Name == "Println" {
			fmt.Println("fmt.Println tipi:", obj.Type())
		}
	}
}
``
/*
### 📌 Çıktı:
*/
``
fmt.Println tipi: func(a ...interface{}) (n int, err error)
``
/*
---

# 🔑 Özet

* `go/importer` paketi, Go’nun standart veya 3. parti paketlerinin **tip bilgilerini yüklemek** için kullanılır.
* `importer.Default()` → en yaygın kullanım.
* `importer.For("gc", nil)` → `gc` compiler için.
* Genellikle `go/types` ile birlikte kullanılır, AST içindeki `import`ları çözmek için.

---

👉 İstersen sana `importer` paketini kullanarak **bir paketin tüm fonksiyonlarını JSON olarak döndüren bir analiz aracı** yazayım. İster misin?
EVET
*/
/*
Harika 👍 O zaman sana **`go/importer` paketini kullanarak paket analiz aracı** yazacağım.
Bu araç:

* Verilen bir Go paketini (ör. `"fmt"`, `"math"`) import edecek,
* İçindeki **fonksiyon, sabit, tip ve değişkenleri** çıkaracak,
* Bunları JSON formatında ekrana basacak.

---

## 📝 Kod: Paket Analiz Aracı
*/
``go
package main

import (
	"encoding/json"
	"fmt"
	"go/importer"
	"go/types"
	"os"
)

type PackageInfo struct {
	Name    string              `json:"name"`
	Path    string              `json:"path"`
	Objects map[string]string   `json:"objects"`
}

func analyzePackage(pkgPath string) (*PackageInfo, error) {
	// Varsayılan importer kullan
	imp := importer.Default()

	// Paketi import et
	pkg, err := imp.Import(pkgPath)
	if err != nil {
		return nil, err
	}

	info := &PackageInfo{
		Name:    pkg.Name(),
		Path:    pkg.Path(),
		Objects: make(map[string]string),
	}

	// Scope’taki tüm dışa açık tanımları topla
	scope := pkg.Scope()
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		info.Objects[name] = obj.Type().String()
	}

	return info, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Kullanım: go run main.go <paket_adi>")
		return
	}

	pkgPath := os.Args[1]

	info, err := analyzePackage(pkgPath)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	// JSON formatında yazdır
	data, _ := json.MarshalIndent(info, "", "  ")
	fmt.Println(string(data))
}
``
/*
---

## ▶️ Kullanım

Terminalde:
*/
``bash
go run main.go fmt
``

//### 📌 Çıktı (kısaltılmış):

``json
{
  "name": "fmt",
  "path": "fmt",
  "objects": {
    "Errorf": "func(format string, a ...interface{}) error",
    "Println": "func(a ...interface{}) (n int, err error)",
    "Sprint": "func(a ...interface{}) string",
    "Formatter": "interface{Format(f State, c rune)}",
    "State": "interface{Write(b []byte) (n int, err error); Width() (wid int, ok bool); Precision() (prec int, ok bool); Flag(c int) bool}"
  }
}
``
/*
Bu şekilde JSON formatında istediğin paketin tip bilgisini çıkarabilirsin 🚀

---

👉 İstersen ben bu aracı biraz daha geliştirip:

* **Yalnızca fonksiyonları filtreleyelim**,
* **Parametre ve dönüş tiplerini ayrı ayrı gösterelim**

ister misin?
EVET
*/
/*
Süper 👍 Şimdi `go/importer` + `go/types` ile yazdığımız aracı geliştirip, **paket fonksiyonlarını ayrı ayrı çözümleyelim**.

Bu sürümde:

* Sadece fonksiyonları çıkarıyoruz,
* Her fonksiyonun **adı**, **parametre listesi** ve **dönüş değerleri** ayrı ayrı listeleniyor.

---

## 📝 Gelişmiş Kod
*/
``go
package main

import (
	"encoding/json"
	"fmt"
	"go/importer"
	"go/types"
	"os"
)

type FuncInfo struct {
	Name       string   `json:"name"`
	Parameters []string `json:"parameters"`
	Results    []string `json:"results"`
}

type PackageFuncs struct {
	Name  string     `json:"name"`
	Path  string     `json:"path"`
	Funcs []FuncInfo `json:"funcs"`
}

func analyzePackageFuncs(pkgPath string) (*PackageFuncs, error) {
	// Paket import et
	imp := importer.Default()
	pkg, err := imp.Import(pkgPath)
	if err != nil {
		return nil, err
	}

	info := &PackageFuncs{
		Name:  pkg.Name(),
		Path:  pkg.Path(),
		Funcs: []FuncInfo{},
	}

	scope := pkg.Scope()
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)

		// Sadece fonksiyonları seçelim
		if fn, ok := obj.(*types.Func); ok {
			sig := fn.Type().(*types.Signature)

			// Parametreler
			params := []string{}
			for i := 0; i < sig.Params().Len(); i++ {
				p := sig.Params().At(i)
				params = append(params, p.Type().String())
			}

			// Dönüş değerleri
			results := []string{}
			for i := 0; i < sig.Results().Len(); i++ {
				r := sig.Results().At(i)
				results = append(results, r.Type().String())
			}

			info.Funcs = append(info.Funcs, FuncInfo{
				Name:       name,
				Parameters: params,
				Results:    results,
			})
		}
	}

	return info, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Kullanım: go run main.go <paket_adi>")
		return
	}

	pkgPath := os.Args[1]

	info, err := analyzePackageFuncs(pkgPath)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	data, _ := json.MarshalIndent(info, "", "  ")
	fmt.Println(string(data))
}
``
/*
---

## ▶️ Kullanım
*/
``bash
go run main.go fmt
``

//### 📌 Çıktı (kısaltılmış):

``json
{
  "name": "fmt",
  "path": "fmt",
  "funcs": [
    {
      "name": "Errorf",
      "parameters": [
        "string",
        "...interface{}"
      ],
      "results": [
        "error"
      ]
    },
    {
      "name": "Println",
      "parameters": [
        "...interface{}"
      ],
      "results": [
        "int",
        "error"
      ]
    },
    {
      "name": "Sprintf",
      "parameters": [
        "string",
        "...interface{}"
      ],
      "results": [
        "string"
      ]
    }
  ]
}
``
/*
---

✅ Artık elimizde **paket fonksiyonlarını detaylı gösteren bir analiz aracı** var.
Bunu kullanarak istediğin paketin fonksiyon imzalarını görebilirsin.

---

👉 İstersen bir sonraki adımda bu aracı genişletip **tipleri (struct, interface) ve sabitleri (const) de çıkaracak** hale getirelim mi?
*/