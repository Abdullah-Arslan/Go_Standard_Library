/*
Go dilinde `path/filepath` paketi, dosya ve dizin yollarını **işletim sistemi bağımsız** bir şekilde işlemek için kullanılır. Yani Windows, Linux ve macOS gibi sistemlerde farklı ayracı (`\` veya `/`) dikkate alarak güvenli ve doğru çalışır. Özellikle Go projelerinde dosya bulma, uzantı kontrolü, dizin oluşturma ve yolları birleştirme gibi işlemler için sıkça tercih edilir.

Aşağıda bu paketin tüm önemli fonksiyonlarını açıklayacağım ve her biri için örnek kodlar vereceğim:

---

# 📂 `filepath` Paketi Fonksiyonları

## 1. `filepath.Abs(path string) (string, error)`

Verilen yolu **mutlak yola** çevirir.
*/
``go
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := "docs/readme.txt"
	abs, _ := filepath.Abs(path)
	fmt.Println("Mutlak Yol:", abs)
}
``
/*
---

## 2. `filepath.Base(path string) string`

Yolun **en son öğesini** döndürür.
*/
``go
func main() {
	fmt.Println(filepath.Base("/home/user/file.txt")) // file.txt
	fmt.Println(filepath.Base("/home/user/"))        // user
}
``
/*
---

## 3. `filepath.Clean(path string) string`

Yolu temizler (`.`, `..`, fazlalık `/`) kaldırır.
*/
``go
func main() {
	fmt.Println(filepath.Clean("/home/user/../docs//file.txt"))
	// Çıktı: /home/docs/file.txt
}
``
/*
---

## 4. `filepath.Dir(path string) string`

Yolun **dizin kısmını** döndürür.
*/
``go
func main() {
	fmt.Println(filepath.Dir("/home/user/file.txt")) // /home/user
}
``
/*
---

## 5. `filepath.Ext(path string) string`

Dosya uzantısını döndürür.
*/
``go
func main() {
	fmt.Println(filepath.Ext("report.pdf")) // .pdf
}
``
/*
---

## 6. `filepath.Join(elem ...string) string`

Parçaları güvenli şekilde birleştirir.
*/
``go
func main() {
	path := filepath.Join("home", "user", "docs", "file.txt")
	fmt.Println(path) // home/user/docs/file.txt (Linux)
}
``
/*
---

## 7. `filepath.Match(pattern, name string) (bool, error)`

Bir yolun pattern ile eşleşip eşleşmediğini kontrol eder (`*`, `?` destekler).
*/
``go
func main() {
	match, _ := filepath.Match("*.txt", "notes.txt")
	fmt.Println(match) // true
}
``
/*
---

## 8. `filepath.Glob(pattern string) ([]string, error)`

Belirli bir pattern’e uyan dosyaları bulur.
*/
``go
func main() {
	files, _ := filepath.Glob("*.go")
	fmt.Println("Go Dosyaları:", files)
}
``
/*
---

## 9. `filepath.Rel(basepath, targpath string) (string, error)`

İki yol arasındaki **göreceli farkı** bulur.
*/
``go
func main() {
	rel, _ := filepath.Rel("/home/user", "/home/user/docs/file.txt")
	fmt.Println(rel) // docs/file.txt
}
``
/*
---

## 10. `filepath.Split(path string) (dir, file string)`

Bir yolu **dizin ve dosya adı** olarak ayırır.
*/
``go
func main() {
	dir, file := filepath.Split("/home/user/file.txt")
	fmt.Println("Dizin:", dir)  // /home/user/
	fmt.Println("Dosya:", file) // file.txt
}
``
/*
---

## 11. `filepath.VolumeName(path string) string`

Windows için disk sürücüsünü döndürür.
*/
``go
func main() {
	fmt.Println(filepath.VolumeName("C:\\Users\\file.txt")) // C:
}
``
/*
---

## 12. `filepath.WalkDir(root string, fn fs.WalkDirFunc) error`

Bir dizini **derinlemesine dolaşır**.
*/
``go
import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println("Bulunan:", path)
		return nil
	})
}
``
/*
---

## 13. `filepath.Walk(root string, walkFn filepath.WalkFunc) error` (Eskisi)

`WalkDir`’in eski versiyonu, `os.FileInfo` ile çalışır.

---

## 14. `filepath.ToSlash(path string) string`

Yoldaki ayraçları `/` ile değiştirir.
*/
``go
func main() {
	fmt.Println(filepath.ToSlash(`C:\Users\Admin\file.txt`))
	// C:/Users/Admin/file.txt
}
``
/*
---

## 15. `filepath.FromSlash(path string) string`

Tersine çevirir (`/` → sistem ayracı).
*/
``go
func main() {
	fmt.Println(filepath.FromSlash("C:/Users/Admin/file.txt"))
	// Windows: C:\Users\Admin\file.txt
}
``
/*
---

## 16. `filepath.IsAbs(path string) bool`

Yolun **mutlak** olup olmadığını kontrol eder.
*/
``go
func main() {
	fmt.Println(filepath.IsAbs("/usr/local/bin")) // true
	fmt.Println(filepath.IsAbs("docs/readme.md")) // false
}
``
/*
---

# ✅ Özet

* `filepath` paketi dosya yollarını işletim sistemi bağımsız şekilde işler.
* **Dosya adı / uzantısı alma:** `Base`, `Dir`, `Ext`, `Split`
* **Yol birleştirme / temizleme:** `Join`, `Clean`, `Abs`, `Rel`
* **Pattern eşleştirme / listeleme:** `Match`, `Glob`
* **Dizin tarama:** `Walk`, `WalkDir`
* **Slash dönüştürme:** `ToSlash`, `FromSlash`

---

👉 İstersen sana bu fonksiyonların **hepsi için tek tek çalışan bir Go demo projesi** hazırlayıp sırayla kodları verebilirim. İstiyor musun?
EVET
*/
/*
Harika 👍 O zaman sana **`filepath` paketinin tüm fonksiyonlarını kapsayan bir demo projesi** hazırlayayım. Bu proje tek bir `main.go` dosyası olacak ve her fonksiyonu **ayrı örnek** olarak gösterecek. Çalıştırdığında sırayla bütün fonksiyonların çıktısını görebileceksin.

---

# 📂 Demo Proje: `filepath_demo/main.go`
*/
``go
package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	// 1. Abs
	abs, _ := filepath.Abs("docs/readme.txt")
	fmt.Println("Abs   :", abs)

	// 2. Base
	fmt.Println("Base  :", filepath.Base("/home/user/file.txt"))

	// 3. Clean
	fmt.Println("Clean :", filepath.Clean("/home/user/../docs//file.txt"))

	// 4. Dir
	fmt.Println("Dir   :", filepath.Dir("/home/user/file.txt"))

	// 5. Ext
	fmt.Println("Ext   :", filepath.Ext("report.pdf"))

	// 6. Join
	fmt.Println("Join  :", filepath.Join("home", "user", "docs", "file.txt"))

	// 7. Match
	match, _ := filepath.Match("*.txt", "notes.txt")
	fmt.Println("Match :", match)

	// 8. Glob
	files, _ := filepath.Glob("*.go")
	fmt.Println("Glob  :", files)

	// 9. Rel
	rel, _ := filepath.Rel("/home/user", "/home/user/docs/file.txt")
	fmt.Println("Rel   :", rel)

	// 10. Split
	dir, file := filepath.Split("/home/user/file.txt")
	fmt.Println("Split :", "Dizin:", dir, "Dosya:", file)

	// 11. VolumeName (özellikle Windows için anlamlıdır)
	fmt.Println("VolumeName:", filepath.VolumeName("C:\\Users\\file.txt"))

	// 12. WalkDir
	fmt.Println("WalkDir:")
	filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println("  ", path)
		return nil
	})

	// 13. Walk (eski sürüm, FileInfo ile çalışır)
	fmt.Println("Walk:")
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println("  ", path)
		return nil
	})

	// 14. ToSlash
	fmt.Println("ToSlash :", filepath.ToSlash(`C:\Users\Admin\file.txt`))

	// 15. FromSlash
	fmt.Println("FromSlash:", filepath.FromSlash("C:/Users/Admin/file.txt"))

	// 16. IsAbs
	fmt.Println("IsAbs (/usr/local/bin):", filepath.IsAbs("/usr/local/bin"))
	fmt.Println("IsAbs (docs/readme.md):", filepath.IsAbs("docs/readme.md"))
}
``
/*
---

# ▶️ Çalıştırma
*/
``sh
go run main.go
``

//Örnek çıktı (Linux’ta, bulunduğun klasöre göre değişir):

``
Abs   : /home/abdullah/projects/docs/readme.txt
Base  : file.txt
Clean : /home/docs/file.txt
Dir   : /home/user
Ext   : .pdf
Join  : home/user/docs/file.txt
Match : true
Glob  : [main.go]
Rel   : docs/file.txt
Split : Dizin:/home/user/ Dosya:file.txt
VolumeName: C:
WalkDir:
   .
   ./main.go
Walk:
   .
   ./main.go
ToSlash : C:/Users/Admin/file.txt
FromSlash: C:\Users\Admin\file.txt
IsAbs (/usr/local/bin): true
IsAbs (docs/readme.md): false
``

---

👉 İstersen ben bu projeyi senin için **PDF, Excel veya Markdown tablo** halinde (fonksiyon + açıklama + örnek) hazırlayıp verebilirim. Böylece offline çalışırken referansın olur.

Bunu ister misin?
EVET

