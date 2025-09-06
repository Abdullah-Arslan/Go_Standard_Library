/*
**Go’nun `go/printer` paketini** en ince ayrıntısına kadar açıklayayım.

---

# 📌 `go/printer` Paketi Nedir?

`go/printer` paketi, Go kaynak kodunu (Go AST – Abstract Syntax Tree üzerinden elde edilen yapıları) **formatlı bir şekilde yazdırmaya** yarar.

👉 Yani bu paket, **Go kodunu programatik olarak üretmek** veya var olan Go kodunu AST üzerinden **güzel, okunabilir bir şekilde yeniden yazmak** için kullanılır.
👉 `gofmt` aracı da aslında arka planda bu paketi kullanır.

---

# 📦 Paketin İçeriği

`go/printer` paketinde temel olarak şu işlevler bulunur:

1. **`Config` yapısı** → Biçimlendirme seçeneklerini tutar.
2. **`Fprint` fonksiyonu** → AST’yi yazdırır.
3. **`CommentedNode` yapısı** → Kod ile yorumları ilişkilendirmek için kullanılır.

---

# 🔑 Önemli Tipler ve Fonksiyonlar

## 1. `printer.Config`

Kodun nasıl yazdırılacağını belirler.
*/
``go
type Config struct {
    Mode     Mode  // biçimlendirme modu
    Tabwidth int   // tab genişliği
    Indent   int   // girinti miktarı
}
``
/*
* **`Mode`** → Yazdırma biçimini belirler (`UseSpaces`, `TabIndent`, `SourcePos` gibi flag’ler).
* **`Tabwidth`** → Bir tab karakterinin boşluk olarak karşılığı.
* **`Indent`** → Ekstra girinti miktarı.

---

## 2. `printer.Fprint`

AST’yi verilen `io.Writer` içine biçimlendirerek yazar.
*/

``go
func Fprint(output io.Writer, fset *token.FileSet, node any) error
``
/*
* **`output`** → Hedef (ör: `os.Stdout`, `bytes.Buffer`)
* **`fset`** → `token.FileSet`, kodun pozisyon bilgilerini tutar.
* **`node`** → Yazdırılacak AST düğümü (ör: `ast.File`)

---

## 3. `printer.CommentedNode`

AST düğümünü yorumlarla ilişkilendirmeye yarar.
*/
``go
type CommentedNode struct {
    Node     any
    Comments []*ast.CommentGroup
}
``
/*
👉 Böylece `//` veya `/* */` gibi yorum satırları da çıktıya eklenir.

---

# 📝 Örnekler

## Örnek 1: Basit AST Yazdırma

Go kodunu parse edip yeniden yazdıralım:
*/
``go
package main

import (
    "go/parser"
    "go/printer"
    "go/token"
    "os"
)

func main() {
    // Go kodu string olarak
    src := package main
import "fmt"
func main() {
fmt.Println("Merhaba Dünya")
}

    // Token set oluştur
    fset := token.NewFileSet()

    // Kodun AST'sini parse et
    node, err := parser.ParseFile(fset, "example.go", src, parser.ParseComments)
    if err != nil {
        panic(err)
    }

    // AST’yi biçimlendirerek stdout’a yaz
    printer.Fprint(os.Stdout, fset, node)
}
``

//📌 Çıktı düzenlenmiş ve okunabilir bir kod olur:

``go
package main

import "fmt"

func main() {
    fmt.Println("Merhaba Dünya")
}
``
/*
---

## Örnek 2: Config ile Özelleştirme

Girintiyi boşluklarla ayarlayalım:
*/
``go
config := &printer.Config{Mode: printer.UseSpaces | printer.TabIndent, Tabwidth: 4}
config.Fprint(os.Stdout, fset, node)
``
/*
👉 Bu sayede `tab` yerine 4 boşluk kullanılır.

---

## Örnek 3: Yorumları da Yazdırma

Yorumları AST ile ilişkilendirelim:
*/
``go
cn := &printer.CommentedNode{
    Node:     node,
    Comments: node.Comments,
}

printer.Fprint(os.Stdout, fset, cn)
``
/*
📌 Böylece kodla birlikte `//` ve `/* */` yorumları da korunur.

---

## Örnek 4: AST Manipülasyonu Sonrası Yazdırma

AST üzerinde değişiklik yapıp tekrar yazdırabiliriz:
*/
``go
for _, decl := range node.Decls {
    if fn, ok := decl.(*ast.FuncDecl); ok {
        fn.Name.Name = "Hello" // main fonksiyonunun adını değiştir
    }
}

printer.Fprint(os.Stdout, fset, node)
``

//Çıktı:

``go
func Hello() {
    fmt.Println("Merhaba Dünya")
}
``
/*
---

# 🚀 Özet

* `go/printer` → Go kodunu **AST’den biçimlendirip yazdırmaya yarar**.
* `Fprint` → Ana fonksiyon, çıktıyı oluşturur.
* `Config` → Biçimlendirme ayarları (tab, boşluk, girinti).
* `CommentedNode` → AST düğümleriyle yorumları ilişkilendirir.
* `gofmt` aracı da aslında bu paketi kullanır.

---

👉 İstersen ben sana `printer` paketini kullanarak **kendi küçük `gofmt` aracımızı yazalım** mı?
EVET
*/
/*
Süper 👍 O zaman beraberce küçük bir **`gofmt` klonu** yazalım.
Bu program:

1. Bir `.go` dosyasını okuyacak
2. AST’ye parse edecek
3. `printer.Fprint` ile biçimlendirip ekrana yazdıracak
4. İstersek dosyanın üzerine yazabilecek

---

# 📂 Mini Gofmt – Kod
*/
``go
package main

import (
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
)

func main() {
	// Kullanım: go run main.go dosya.go
	if len(os.Args) < 2 {
		fmt.Println("Kullanım: gofmt <dosya.go>")
		return
	}

	filename := os.Args[1]

	// Dosyayı oku
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Dosya okunamadı:", err)
		return
	}

	// Token set oluştur
	fset := token.NewFileSet()

	// Kaynağı parse et (yorumlarla birlikte)
	node, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		fmt.Println("Parse hatası:", err)
		return
	}

	// Config oluştur → boşluklarla tab yerine yazdıralım
	config := &printer.Config{
		Mode:     printer.UseSpaces | printer.TabIndent,
		Tabwidth: 4,
		Indent:   0,
	}

	// Biçimlendirilmiş çıktıyı buffer’a yaz
	var output []byte
	buf := &printer.CommentedNode{Node: node, Comments: node.Comments}

	file, err := os.Create(filename + ".formatted.go")
	if err != nil {
		fmt.Println("Yeni dosya oluşturulamadı:", err)
		return
	}
	defer file.Close()

	// AST → dosya
	err = config.Fprint(file, fset, buf)
	if err != nil {
		fmt.Println("Yazdırma hatası:", err)
		return
	}

	fmt.Println("Biçimlendirilmiş kod", filename+".formatted.go", "dosyasına kaydedildi.")
}
``
/*
---

# 📌 Nasıl Çalışır?

1. Yukarıdaki kodu `mini_gofmt.go` diye kaydet.
2. Test için bozuk girintili bir dosya hazırla:

`örnek.go`:
*/
``go
package main
import "fmt"
func main(){
fmt.Println("Selam")
}
``

//3. Çalıştır:

``bash
go run mini_gofmt.go örnek.go
``

//4. Çıktı `örnek.go.formatted.go` dosyasına kaydedilir:

``go
package main

import "fmt"

func main() {
    fmt.Println("Selam")
}
``
/*
---

İstersen bu kodu biraz daha geliştirip:

* **dosyanın üzerine yazma** (tüm IDE’lerdeki `gofmt -w` gibi)
* **bir klasördeki tüm `.go` dosyalarını otomatik formatlama**
* **stdout’a yazma (gofmt gibi)**

özelliklerini ekleyebiliriz.

👉 İstiyor musun ben sana **gofmt -w benzeri (dosyanın üzerine yazan)** versiyonunu da yazayım?
*/