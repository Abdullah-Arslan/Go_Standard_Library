/*
Golang’da **`ioutil` paketi**, Go 1.16 öncesinde dosya ve veri okuma/yazma işlemleri için kullanılan bir yardımcı paketti. Ancak **Go 1.16 ve sonrasında** çoğu fonksiyonu `io` ve `os` paketlerine taşındı, `ioutil` artık **deprecated** sayılır. Yine de anlaşılması için detaylı şekilde açıklayalım.

---

## 1️⃣ `ioutil` Paketi Nedir?

`ioutil`, “input/output utility” anlamına gelir ve temel olarak **dosya okuma, yazma ve geçici dosya/dizin oluşturma** işlemlerini kolaylaştırır.

**Öne çıkan fonksiyonları:**

| Fonksiyon                                                         | Açıklama                                           |
| ----------------------------------------------------------------- | -------------------------------------------------- |
| `ReadFile(filename string) ([]byte, error)`                       | Dosyanın tamamını okur ve byte slice olarak döner. |
| `WriteFile(filename string, data []byte, perm os.FileMode) error` | Byte slice’ı dosyaya yazar.                        |
| `ReadDir(dirname string) ([]os.FileInfo, error)`                  | Dizindeki tüm dosya ve dizinleri listeler.         |
| `TempFile(dir, pattern string) (*os.File, error)`                 | Geçici dosya oluşturur.                            |
| `TempDir(dir, pattern string) (string, error)`                    | Geçici dizin oluşturur.                            |
| `NopCloser(r io.Reader) io.ReadCloser`                            | Sadece Reader’dan ReadCloser oluşturur.            |

> ⚠️ Go 1.16 ve sonrası için:
>
> * `ioutil.ReadFile` → `os.ReadFile`
> * `ioutil.WriteFile` → `os.WriteFile`
> * `ioutil.ReadDir` → `os.ReadDir`
> * `ioutil.TempFile` → `os.CreateTemp`
> * `ioutil.TempDir` → `os.MkdirTemp`

---

## 2️⃣ Örnekler

### a) Dosya okuma (`ReadFile`)
*/
``go
package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    data, err := ioutil.ReadFile("example.txt")
    if err != nil {
        fmt.Println("Hata:", err)
        return
    }
    fmt.Println(string(data))
}
``
/*
---

### b) Dosya yazma (`WriteFile`)
*/
``go
package main

import (
    "io/ioutil"
)

func main() {
    content := []byte("Merhaba Go!")
    err := ioutil.WriteFile("output.txt", content, 0644)
    if err != nil {
        panic(err)
    }
}
``
/*
---

### c) Dizin okuma (`ReadDir`)
*/
``go
package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    files, err := ioutil.ReadDir(".")
    if err != nil {
        panic(err)
    }
    for _, f := range files {
        fmt.Println(f.Name(), f.IsDir())
    }
}
``
/*
---

### d) Geçici dosya oluşturma (`TempFile`)
*/
``go
package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    tmpFile, err := ioutil.TempFile("", "example-*.txt")
    if err != nil {
        panic(err)
    }
    defer tmpFile.Close()

    fmt.Println("Geçici dosya oluşturuldu:", tmpFile.Name())
}
``
/*
---

### e) Reader’dan ReadCloser oluşturma (`NopCloser`)
*/
``go
package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "strings"
)

func main() {
    r := strings.NewReader("Merhaba Go!")
    rc := ioutil.NopCloser(r)
    defer rc.Close()

    data, _ := io.ReadAll(rc)
    fmt.Println(string(data))
}
``
/*
---

## 3️⃣ Özet

* `ioutil` **dosya/dizin okuma-yazma ve geçici dosya/dizin oluşturmayı kolaylaştırır.**
* **Go 1.16 sonrası** çoğu fonksiyonu `os` ve `io` paketleriyle değiştirilmiştir.
* Modern Go kodu yazarken artık `ioutil` yerine `os.ReadFile`, `os.WriteFile`, `os.ReadDir` vb. kullanmak önerilir.

---

İstersen bir sonraki adımda sana **`ioutil` ve yeni `os/io` fonksiyonları ile örnek interaktif dosya işlemleri uygulaması** da gösterebilirim.

Bunu yapayım mı?
EVET
*/
/*
Harika! 🚀 O zaman sana **`ioutil` ve modern `os/io` fonksiyonları kullanarak interaktif dosya işlemleri uygulaması** örneği hazırlayayım.

Bu örnek, **dizin gezgini, dosya okuma ve yazma, geçici dosya oluşturma** işlemlerini terminal üzerinden interaktif yapmanı sağlayacak.

---

# 📌 Örnek: Interaktif Dosya İşlemleri
*/
``go
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

func listDir(dir string) []os.DirEntry {
    entries, err := os.ReadDir(dir)
    if err != nil {
        fmt.Println("Hata:", err)
        return nil
    }
    for i, e := range entries {
        typ := "Dosya"
        if e.IsDir() {
            typ = "Dizin"
        }
        fmt.Printf("[%d] %s: %s\n", i, typ, e.Name())
    }
    return entries
}

func readFile(path string) {
    data, err := os.ReadFile(path)
    if err != nil {
        fmt.Println("Dosya okunamadı:", err)
        return
    }
    fmt.Println("=== Dosya İçeriği ===")
    fmt.Println(string(data))
    fmt.Println("=====================")
}

func writeFile(path string, content string) {
    err := os.WriteFile(path, []byte(content), 0644)
    if err != nil {
        fmt.Println("Dosya yazılamadı:", err)
        return
    }
    fmt.Println("Dosya başarıyla kaydedildi:", path)
}

func createTempFile() {
    tmpFile, err := os.CreateTemp("", "example-*.txt")
    if err != nil {
        fmt.Println("Geçici dosya oluşturulamadı:", err)
        return
    }
    defer tmpFile.Close()
    fmt.Println("Geçici dosya oluşturuldu:", tmpFile.Name())
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    currentDir := "."

    for {
        fmt.Println("\nMevcut dizin:", currentDir)
        entries := listDir(currentDir)

        fmt.Print("\nSeçim yap (numara), 'u' üst dizin, 'r' dosya oku, 'w' dosya yaz, 't' geçici dosya, 'exit': ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        switch input {
        case "exit":
            return
        case "u":
            if currentDir != "." {
                parts := strings.Split(currentDir, "/")
                currentDir = strings.Join(parts[:len(parts)-1], "/")
                if currentDir == "" {
                    currentDir = "."
                }
            }
            continue
        case "t":
            createTempFile()
            continue
        case "r":
            fmt.Print("Dosya adı gir: ")
            fname, _ := reader.ReadString('\n')
            fname = strings.TrimSpace(fname)
            readFile(currentDir + "/" + fname)
            continue
        case "w":
            fmt.Print("Dosya adı gir: ")
            fname, _ := reader.ReadString('\n')
            fname = strings.TrimSpace(fname)
            fmt.Print("Yazılacak içerik gir: ")
            content, _ := reader.ReadString('\n')
            content = strings.TrimSpace(content)
            writeFile(currentDir+"/"+fname, content)
            continue
        default:
            idx := -1
            fmt.Sscanf(input, "%d", &idx)
            if idx >= 0 && idx < len(entries) {
                if entries[idx].IsDir() {
                    currentDir = currentDir + "/" + entries[idx].Name()
                } else {
                    fmt.Println("Seçilen bir dosya, dizin seçmelisiniz veya 'r' ile okuyun.")
                }
            } else {
                fmt.Println("Geçersiz seçim")
            }
        }
    }
}
``
/*
---

# 📌 Özellikler

1. **Dizin gezgini** → Numara ile dizin seçimi.
2. **Dosya okuma** → `os.ReadFile` kullanımı.
3. **Dosya yazma** → `os.WriteFile` ile içerik ekleme.
4. **Geçici dosya oluşturma** → `os.CreateTemp`.
5. **Üst dizine çıkma** → `'u'` ile.
6. **Terminal tabanlı interaktif kullanım** → Numara veya komut ile işlem.

---

Bu örnek sayesinde artık hem `ioutil` fonksiyonlarını hem de modern `os/io` fonksiyonlarını kullanarak **interaktif dosya işlemleri** yapabilirsin.

İstersen bir sonraki adımda bunu **dizin gezgini + çoklu dosya arama + arama sonuçları kaydırma + renkli vurgulama** ile birleştirip tam **terminal grep ve dosya yöneticisi uygulaması** hâline getirebilirim.

Bunu da yapalım mı?
*/