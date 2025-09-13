/*
Go’da **`path`** paketi, dosya yollarını **işlemek ve manipüle etmek** için kullanılan basit ama güçlü bir standart kütüphane paketidir. Bu paket, özellikle **Unix-benzeri (forward slash `/`) yollar** üzerinde çalışır. Eğer işletim sistemi bağımlı (Windows `\`, Unix `/`) yollarla çalışmak isterseniz, onun için **`path/filepath`** paketini kullanmanız gerekir.

Şimdi sana `path` paketinin tamamını fonksiyon fonksiyon açıklayayım ve her biri için örnekler vereyim. 🚀

---

# 📦 `path` Paketi Genel Yapısı
*/
``go
import "path"
``
/*
* Yolları birleştirme (`Join`)
* Normalleştirme (`Clean`)
* Yol parçalama (`Split`, `Base`, `Dir`, `Ext`)
* Göreli yol hesaplama (`Rel`)
* Ön ek kontrol (`HasPrefix`, `IsAbs`)
* Yolları eşleştirme (`Match`)

---

# 🔹 Fonksiyonlar ve Örnekler

### 1. `path.Clean(path string) string`

Yolu **normalize** eder, gereksiz `.` ve `..` kısımlarını kaldırır.
*/
``go
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Clean("/foo/bar/../baz/./qux"))
	// Çıktı: /foo/baz/qux
}
``
/*
---

### 2. `path.Join(elem ...string) string`

Parçaları birleştirir ve temizler.
*/
``go
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Join("foo", "bar", "baz"))
	// Çıktı: foo/bar/baz
}
``
/*
---

### 3. `path.Split(path string) (dir, file string)`

Yolu **dizin** ve **dosya** olarak böler.
*/

``go
package main

import (
	"fmt"
	"path"
)

func main() {
	dir, file := path.Split("/home/user/docs/file.txt")
	fmt.Println("Dir:", dir)   // Dir: /home/user/docs/
	fmt.Println("File:", file) // File: file.txt
}
``
/*
---

### 4. `path.Base(path string) string`

Son elemanı (dosya veya dizin) döner.
*/
``go
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Base("/home/user/docs/file.txt"))
	// Çıktı: file.txt
}
``
/*
---

### 5. `path.Dir(path string) string`

Yalnızca **dizin kısmını** döner.
*/
``go
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("/home/user/docs/file.txt"))
	// Çıktı: /home/user/docs
}
``
/*
---

### 6. `path.Ext(path string) string`

Dosya uzantısını alır.
*/
``go
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Ext("report.pdf"))
	// Çıktı: .pdf
}
``
/*
---

### 7. `path.IsAbs(path string) bool`

Yolun **mutlak yol** olup olmadığını kontrol eder.
*/
``go
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.IsAbs("/home/user")) // true
	fmt.Println(path.IsAbs("docs/file"))  // false
}
``
/*
---

### 8. `path.Rel(basepath, targpath string) (string, error)`

İki yol arasındaki **göreli yolu** hesaplar.
*/
``go
package main

import (
	"fmt"
	"path"
)

func main() {
	rel, _ := path.Rel("/home/user", "/home/user/docs/file.txt")
	fmt.Println(rel)
	// Çıktı: docs/file.txt
}
``
/*
---

### 9. `path.HasPrefix(p, prefix string) bool` (⚠️ DEPRECATED)

Go 1.12’den sonra **`strings.HasPrefix`** kullanılması önerilir.
Eskiden bir yolun ön ekini kontrol etmek için vardı.
*/
``go
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.HasPrefix("/foo/bar", "/foo"))
	// Çıktı: true
}
``
/*
---

### 10. `path.Match(pattern, name string) (matched bool, err error)`

Unix-style glob eşleşmesi yapar (`*`, `?`, `[...]`).
*/
``go
package main

import (
	"fmt"
	"path"
)

func main() {
	ok, _ := path.Match("*.txt", "notes.txt")
	fmt.Println(ok) // true

	ok, _ = path.Match("data-?.csv", "data-1.csv")
	fmt.Println(ok) // true

	ok, _ = path.Match("data-?.csv", "data-12.csv")
	fmt.Println(ok) // false
}
``

---

# 📌 Özet Tablo

| Fonksiyon          | Açıklama                                  |
| ------------------ | ----------------------------------------- |
| `Clean`            | Yolu normalize eder (`.`, `..` temizler). |
| `Join`             | Yol parçalarını birleştirir.              |
| `Split`            | Dizin ve dosyayı ayırır.                  |
| `Base`             | Yolun son elemanını döner.                |
| `Dir`              | Dizin kısmını döner.                      |
| `Ext`              | Dosya uzantısını alır.                    |
| `IsAbs`            | Mutlak yol mu kontrol eder.               |
| `Rel`              | Göreli yol hesaplar.                      |
| `HasPrefix` (eski) | Yolun önekini kontrol eder.               |
| `Match`            | Glob eşleştirme yapar.                    |

/*
---

✅ `path` paketi: **sadece Unix-benzeri yollar** için.
✅ OS bağımlı (Windows dahil) yollarla çalışacaksan → **`path/filepath`** kullan.

---

İstersen ben şimdi sana bunun **`path/filepath`** versiyonunu da (Windows + Linux uyumlu) aynı şekilde örneklerle açıklayabilirim. Onu da ister misin?
EVET

Süper 👍 O zaman sana **`path` vs `path/filepath`** farklarını hem tablo hem de mini proje örneği ile göstereyim. 🚀

---
*/

# 📌 `path` vs `path/filepath` Karşılaştırma Tablosu

| Özellik                 | `path`                             | `path/filepath`                             |
| ----------------------- | ---------------------------------- | ------------------------------------------- |
| Platform uyumu          | Sadece **Unix tarzı** yollar (`/`) | **OS bağımlı** (Windows `\`, Linux/Mac `/`) |
| Kullanım alanı          | URL, web path, network path işleme | Dosya sistemi yolları (platform bağımsız)   |
| Örnek `Join("a","b")`   | `a/b`                              | Linux: `a/b`, Windows: `a\b`                |
| Örnek `Clean("a/../b")` | `b`                                | Windows’ta da aynı şekilde normalize eder   |
| `Glob`, `WalkDir`       | ❌ Yok                              | ✅ Var                                       |
| `ToSlash` / `FromSlash` | ❌ Yok                              | ✅ Var (Windows/Linux uyumu)                 |

/*
---

# 🔹 Uygulamalı Karşılaştırma

Aşağıdaki örnek kod, hem `path` hem de `path/filepath` paketini karşılaştırır 👇
*/
``go
package main

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
)

func main() {
	fmt.Println("Go çalıştığı OS:", runtime.GOOS)
	fmt.Println("--- path paketi ---")
	fmt.Println("Join:", path.Join("foo", "bar", "baz"))
	fmt.Println("Clean:", path.Clean("/foo/bar/../baz"))
	fmt.Println("Ext:", path.Ext("report.pdf"))

	fmt.Println("\n--- filepath paketi ---")
	fmt.Println("Join:", filepath.Join("foo", "bar", "baz"))
	fmt.Println("Clean:", filepath.Clean("foo/../bar/baz"))
	fmt.Println("Ext:", filepath.Ext("report.pdf"))

	abs, _ := filepath.Abs("docs/file.txt")
	fmt.Println("Abs:", abs)

	rel, _ := filepath.Rel("/home/user", "/home/user/docs/file.txt")
	fmt.Println("Rel:", rel)

	// Platform bağımlı dönüşüm
	fmt.Println("ToSlash:", filepath.ToSlash("C:\\Users\\Admin\\file.txt"))
	fmt.Println("FromSlash:", filepath.FromSlash("C:/Users/Admin/file.txt"))
}
``
/*
---

# 🖥️ Örnek Çalıştırma Sonuçları

### Linux / MacOS’ta:
*/
``
Go çalıştığı OS: linux
--- path paketi ---
Join: foo/bar/baz
Clean: /foo/baz
Ext: .pdf

--- filepath paketi ---
Join: foo/bar/baz
Clean: bar/baz
Ext: .pdf
Abs: /home/kullanici/proje/docs/file.txt
Rel: docs/file.txt
ToSlash: C:/Users/Admin/file.txt
FromSlash: C:/Users/Admin/file.txt
``

//### Windows’ta:

``
Go çalıştığı OS: windows
--- path paketi ---
Join: foo/bar/baz
Clean: /foo/baz
Ext: .pdf

--- filepath paketi ---
Join: foo\bar\baz
Clean: bar\baz
Ext: .pdf
Abs: C:\Users\Abdullah\Proje\docs\file.txt
Rel: docs\file.txt
ToSlash: C:/Users/Admin/file.txt
FromSlash: C:\Users\Admin\file.txt
``
/*
---

# 🔹 Mini Proje: Platform Bağımsız Dosya Tarayıcı

Bu küçük program çalıştığı klasördeki **.go dosyalarını bulur**.
*/
``go
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// Geçerli klasördeki tüm .go dosyalarını ara
	files, _ := filepath.Glob("*.go")

	fmt.Println("Bulunan Go dosyaları:")
	for _, f := range files {
		abs, _ := filepath.Abs(f)
		fmt.Println(" -", abs)
	}
}
``
/*
📌 Linux/Mac’te `/home/user/proje/main.go`
📌 Windows’ta `C:\Users\Abdullah\Proje\main.go` şeklinde çıktı verir.

---

✅ ÖZET:

* `path` → web URL/path işlemleri
* `filepath` → dosya sistemi (platform bağımsız)
* Proje geliştirirken **diskteki dosyalar** için → `filepath` kullanmalısın.

---

Abdullah, istersen ben sana bunun üzerine **gerçek bir blog projesinde (senin `goweb`) statik dosya (CSS/JS/images) yollarını `filepath` ile nasıl çözebileceğini** de gösterebilirim. İster misin?
EVET
*/