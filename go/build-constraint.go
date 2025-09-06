/*
Şimdi sana **Go’nun `build/constraint` paketini** detaylıca anlatayım. Bu paket, **Go 1.17** ile birlikte geldi ve **build tag (yapı etiketleri)** mekanizmasını parse etmek / kontrol etmek için kullanılıyor.

Eskiden build tag’ler `// +build linux darwin` gibi yazılırdı. Artık modern yazımı:
*/
``go
go:build linux || darwin
``
/*
şeklinde. İşte `build/constraint` paketi bu ifadeleri **parse edip yorumlamak** için var.

---

# 📦 `build/constraint` Paketi

📍 Paket yolu:
*/

``go
import "go/build/constraint"
``
/*
## 1️⃣ Temel Kavramlar

* **Build tag (etiket)**: Derleme sırasında hangi dosyanın seçileceğini belirler.
  Örn: `//go:build linux` → sadece Linux’ta derlenir.
* `constraint.Expr`: Bir build constraint ifadesini temsil eder. (ör: `linux && amd64`)
* `Parse()`: Bir build constraint stringini çözümler.
* `Eval()`: Belirli etiketler için constraint ifadesini değerlendirir.

---

## 2️⃣ Önemli Tipler ve Fonksiyonlar

### 🔹 `constraint.Expr`

Bir build constraint ifadesini soyutlar. Bu aslında bir interface:

* `AndExpr`, `OrExpr`, `NotExpr` gibi alt tipleri vardır.
* `Eval(tags map[string]bool) bool` → Etiketler setine göre ifadenin doğruluğunu döner.

---

### 🔹 `constraint.Parse(line string) (Expr, error)`

Bir `//go:build` satırını parse edip `Expr` döner.

📌 Örnek:
*/
``go
expr, err := constraint.Parse("linux && amd64")
if err != nil {
    panic(err)
}
fmt.Println(expr.String()) // "linux && amd64"
``
/*
---

### 🔹 `Expr.Eval(tags map[string]bool) bool`

Bir constraint ifadesini verilen etiket kümesine göre değerlendirir.

📌 Örnek:
*/
``go
package main

import (
	"fmt"
	"go/build/constraint"
)

func main() {
	expr, _ := constraint.Parse("linux && amd64")

	// Etiket seti: linux + amd64 aktif
	tags := map[string]bool{
		"linux": true,
		"amd64": true,
	}
	fmt.Println(expr.Eval(tags)) // true

	// Etiket seti: windows + amd64
	tags2 := map[string]bool{
		"windows": true,
		"amd64":   true,
	}
	fmt.Println(expr.Eval(tags2)) // false
}
``
/*
---

### 🔹 `constraint.And`, `constraint.Or`, `constraint.Not`

Program içinde manuel olarak constraint oluşturmak için yardımcı fonksiyonlardır.

📌 Örnek:
*/

``go
package main

import (
	"fmt"
	"go/build/constraint"
)

func main() {
	linux := constraint.Tag("linux")
	amd64 := constraint.Tag("amd64")

	expr := constraint.And(linux, amd64) // linux && amd64

	fmt.Println(expr.String()) // "linux && amd64"

	tags := map[string]bool{"linux": true, "amd64": true}
	fmt.Println(expr.Eval(tags)) // true
}
``
/*
---

## 3️⃣ Gerçekçi Senaryo

Diyelim ki elimizde bir dosya var:
*/
``go
//go:build linux || windows

package main
``
*/
//Bunu programla analiz edelim:

``go
package main

import (
	"fmt"
	"go/build/constraint"
)

func main() {
	expr, _ := constraint.Parse("linux || windows")

	fmt.Println("Build ifadesi:", expr.String())

	tags := map[string]bool{"linux": true}
	fmt.Println("Linux için geçerli mi?", expr.Eval(tags))

	tags2 := map[string]bool{"darwin": true}
	fmt.Println("Darwin için geçerli mi?", expr.Eval(tags2))
}
``
/*
🖥 Çıktı:

```
Build ifadesi: linux || windows
Linux için geçerli mi? true
Darwin için geçerli mi? false
```

---

## 4️⃣ Özet Tablo

| Fonksiyon / Tip          | Açıklama                              |   |                |
| ------------------------ | ------------------------------------- | - | -------------- |
| `constraint.Parse(line)` | `//go:build` satırını parse eder      |   |                |
| `constraint.Expr`        | Build constraint ifadesi (interface)  |   |                |
| `Expr.String()`          | İnsan okunabilir string karşılığı     |   |                |
| `Expr.Eval(tags)`        | Verilen etiketlere göre boolean sonuç |   |                |
| `constraint.Tag(name)`   | Basit bir tag oluşturur               |   |                |
| `constraint.And(a,b)`    | `a && b` constraint                   |   |                |
| `constraint.Or(a,b)`     | \`a                                   |   | b\` constraint |
| `constraint.Not(a)`      | `!a` constraint                       |   |                |

---

✅ Yani `build/constraint` paketini kullanarak Go dosyalarının **hangi platformlarda derleneceğini** analiz edebilir, özel build sistemleri veya kod analiz araçları yazabilirsin.

---

İstersen ben sana küçük bir **araç** yazayım:
→ Bir dizindeki `.go` dosyalarının `//go:build` satırlarını okuyup, `build/constraint` ile hangi platformda derleneceğini ekrana yazdırsın.

Bunu ister misin?
EVET
*/

/*
Çok güzel bir soru 👍
Şimdi sana **Go’nun `build/constraint` paketini** detaylıca anlatayım. Bu paket, **Go 1.17** ile birlikte geldi ve **build tag (yapı etiketleri)** mekanizmasını parse etmek / kontrol etmek için kullanılıyor.

Eskiden build tag’ler `// +build linux darwin` gibi yazılırdı. Artık modern yazımı:
*/

``go
go:build linux || darwin
``
/*
şeklinde. İşte `build/constraint` paketi bu ifadeleri **parse edip yorumlamak** için var.

---

# 📦 `build/constraint` Paketi

📍 Paket yolu:
*/

``go
import "go/build/constraint"
``
/*
## 1️⃣ Temel Kavramlar

* **Build tag (etiket)**: Derleme sırasında hangi dosyanın seçileceğini belirler.
  Örn: `//go:build linux` → sadece Linux’ta derlenir.
* `constraint.Expr`: Bir build constraint ifadesini temsil eder. (ör: `linux && amd64`)
* `Parse()`: Bir build constraint stringini çözümler.
* `Eval()`: Belirli etiketler için constraint ifadesini değerlendirir.

---

## 2️⃣ Önemli Tipler ve Fonksiyonlar

### 🔹 `constraint.Expr`

Bir build constraint ifadesini soyutlar. Bu aslında bir interface:

* `AndExpr`, `OrExpr`, `NotExpr` gibi alt tipleri vardır.
* `Eval(tags map[string]bool) bool` → Etiketler setine göre ifadenin doğruluğunu döner.

---

### 🔹 `constraint.Parse(line string) (Expr, error)`

Bir `//go:build` satırını parse edip `Expr` döner.

📌 Örnek:
*/

``go
expr, err := constraint.Parse("linux && amd64")
if err != nil {
    panic(err)
}
fmt.Println(expr.String()) // "linux && amd64"
``
/*
---

### 🔹 `Expr.Eval(tags map[string]bool) bool`

Bir constraint ifadesini verilen etiket kümesine göre değerlendirir.

📌 Örnek:
*/
``go
package main

import (
	"fmt"
	"go/build/constraint"
)

func main() {
	expr, _ := constraint.Parse("linux && amd64")

	// Etiket seti: linux + amd64 aktif
	tags := map[string]bool{
		"linux": true,
		"amd64": true,
	}
	fmt.Println(expr.Eval(tags)) // true

	// Etiket seti: windows + amd64
	tags2 := map[string]bool{
		"windows": true,
		"amd64":   true,
	}
	fmt.Println(expr.Eval(tags2)) // false
}
``
/*
---

### 🔹 `constraint.And`, `constraint.Or`, `constraint.Not`

Program içinde manuel olarak constraint oluşturmak için yardımcı fonksiyonlardır.

📌 Örnek:
*/
``go
package main

import (
	"fmt"
	"go/build/constraint"
)

func main() {
	linux := constraint.Tag("linux")
	amd64 := constraint.Tag("amd64")

	expr := constraint.And(linux, amd64) // linux && amd64

	fmt.Println(expr.String()) // "linux && amd64"

	tags := map[string]bool{"linux": true, "amd64": true}
	fmt.Println(expr.Eval(tags)) // true
}
``
/*
---

## 3️⃣ Gerçekçi Senaryo

Diyelim ki elimizde bir dosya var:
*/

``go
//go:build linux || windows

package main
``

//Bunu programla analiz edelim:

``go
package main

import (
	"fmt"
	"go/build/constraint"
)

func main() {
	expr, _ := constraint.Parse("linux || windows")

	fmt.Println("Build ifadesi:", expr.String())

	tags := map[string]bool{"linux": true}
	fmt.Println("Linux için geçerli mi?", expr.Eval(tags))

	tags2 := map[string]bool{"darwin": true}
	fmt.Println("Darwin için geçerli mi?", expr.Eval(tags2))
}
``

//🖥 Çıktı:

``
Build ifadesi: linux || windows
Linux için geçerli mi? true
Darwin için geçerli mi? false
``
/*
---

## 4️⃣ Özet Tablo

| Fonksiyon / Tip          | Açıklama                              |   |                |
| ------------------------ | ------------------------------------- | - | -------------- |
| `constraint.Parse(line)` | `//go:build` satırını parse eder      |   |                |
| `constraint.Expr`        | Build constraint ifadesi (interface)  |   |                |
| `Expr.String()`          | İnsan okunabilir string karşılığı     |   |                |
| `Expr.Eval(tags)`        | Verilen etiketlere göre boolean sonuç |   |                |
| `constraint.Tag(name)`   | Basit bir tag oluşturur               |   |                |
| `constraint.And(a,b)`    | `a && b` constraint                   |   |                |
| `constraint.Or(a,b)`     | \`a                                   |   | b\` constraint |
| `constraint.Not(a)`      | `!a` constraint                       |   |                |

---

✅ Yani `build/constraint` paketini kullanarak Go dosyalarının **hangi platformlarda derleneceğini** analiz edebilir, özel build sistemleri veya kod analiz araçları yazabilirsin.

---

İstersen ben sana küçük bir **araç** yazayım:
→ Bir dizindeki `.go` dosyalarının `//go:build` satırlarını okuyup, `build/constraint` ile hangi platformda derleneceğini ekrana yazdırsın.

Bunu ister misin?
EVET
*/
/*
Süper 👍 o zaman `build/constraint` + `go/build` paketlerini **birlikte** kullanan bir araç yazalım.

Bu araç:

1. Bir dizindeki `.go` dosyalarını bulacak.
2. `//go:build` satırlarını `constraint.Parse` ile analiz edecek.
3. Aynı zamanda `go/build` paketiyle o dizinin **aktif ortamda** derlenecek dosyalarını gösterecek.

---

# 📝 Kod: `analyze_full.go`
*/

``go
package main

import (
	"bufio"
	"fmt"
	"go/build"
	"go/build/constraint"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	root := "./myapp" // analiz etmek istediğin proje dizini

	fmt.Println("🔍 Build constraint analizi başlıyor...\n")

	// 1️⃣ Dosyaları tarayalım ve build tag'leri çözümleyelim
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			checkBuildConstraints(path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Hata:", err)
	}

	// 2️⃣ go/build kullanarak aktif ortamda derlenecek dosyaları listeleyelim
	fmt.Println("\n⚙️ go/build analizi:")
	ctx := build.Default
	pkg, err := ctx.ImportDir(root, 0)
	if err != nil {
		fmt.Println("Import hatası:", err)
		return
	}

	fmt.Println("📦 Paket adı:", pkg.Name)
	fmt.Println("📂 Dizini:", pkg.Dir)
	fmt.Println("📄 Bu ortamda derlenecek Go dosyaları:", pkg.GoFiles)
}
``
/*
---

# 🔎 Nasıl Çalışır?

Örneğin proje `myapp/` şu dosyalara sahipse:
*/
``go
// util_linux.go
//go:build linux
package util
``

``go
// util_windows.go
//go:build windows
package util
``

``go
// util.go
package util
``
/*
---

# 🖥 Çıktı (Linux ortamında):

```
🔍 Build constraint analizi başlıyor...

📄 ./myapp/util_linux.go
   Build ifadesi: linux
   Tags map[amd64:true linux:true] => true
   Tags map[amd64:true windows:true] => false
   Tags map[arm64:true darwin:true] => false

📄 ./myapp/util_windows.go
   Build ifadesi: windows
   Tags map[amd64:true linux:true] => false
   Tags map[amd64:true windows:true] => true
   Tags map[arm64:true darwin:true] => false

⚙️ go/build analizi:
📦 Paket adı: util
📂 Dizini: /home/user/myapp
📄 Bu ortamda derlenecek Go dosyaları: [util.go util_linux.go]
```

👉 Burada `go/build` paketi sayesinde **aktif ortamda hangi dosyaların seçileceğini** görebiliyoruz.
(`util.go` + `util_linux.go` derlenmiş, `util_windows.go` atlanmış.)

---

İstersen ben bu aracı biraz daha geliştirip:

* `GOOS` / `GOARCH` değerlerini **manuel seçme** imkanı ekleyebilirim (örneğin Linux’tayken Windows için hangi dosyalar derlenirdi?).

Bunu ister misin?
*/