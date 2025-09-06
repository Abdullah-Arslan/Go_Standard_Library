/*
Go dilinde **`io/fs` paketi** (genellikle `fs` olarak da kullanılır), Go 1.16 ile gelen **dosya sistemi soyutlaması** paketidir. Bu paket, **dosya ve dizinlere erişimi standartlaştırmak** için geliştirilmiştir ve hem yerel dosya sistemi hem de gömülü veya sanal dosya sistemleriyle çalışabilir.

Aşağıda detaylı bir açıklama ve örnekler bulabilirsin.

---

# 📌 `io/fs` Paketi Nedir?

* **`fs` paketi**, dosya sistemi nesnelerini (dosya ve dizinleri) **soyut bir arayüz ile** temsil eder.
* Paket sayesinde hem **gerçek dosya sistemine**, hem de **embed edilmiş dosya sistemlerine** (`embed.FS`) aynı şekilde erişim sağlayabilirsiniz.
* **Okuma odaklı**dır; yazma veya değiştirme işlemleri için ayrı paketler gerekir (örn. `os`).

---

# 📌 Önemli Arayüzler

1. **`fs.FS`**

   * Temel dosya sistemi arayüzü.
   * Tek metod:

     ```go
     Open(name string) (fs.File, error)
     ```
   * Dosya veya dizin açar.

2. **`fs.File`**

   * Açılmış bir dosyayı temsil eder.
   * Arayüzler:

     * `io.Reader` → okuma
     * `io.Closer` → kapatma
     * `Stat() (fs.FileInfo, error)` → dosya bilgisi

3. **`fs.ReadDirFS`**

   * Dizinleri okuyabilen FS.
   * `ReadDir(name string) ([]fs.DirEntry, error)` ile dizin içeriğini listeler.

4. **`fs.StatFS`**

   * Dosya bilgisi almayı destekleyen FS.

5. **`fs.Sub`**

   * Bir alt dizini FS olarak kullanmayı sağlar.

6. **`fs.Glob`**

   * Pattern’e göre dosya eşleştirme sağlar (`*.txt`, vb).

---

# 📌 Önemli Tipler

* **`fs.FileInfo`** → Dosya adı, boyutu, izinler, dizin mi dosya mı gibi bilgiler.
* **`fs.DirEntry`** → Dizin girişlerini temsil eder (ad, dizin mi vs).

---

# 📌 Örnekler

## 1️⃣ Basit Dosya Okuma
*/
``go
package main

import (
    "fmt"
    "io/fs"
    "io/ioutil"
    "log"
    "os"
)

func main() {
    var fSys fs.FS = os.DirFS(".") // Geçerli dizin FS olarak

    file, err := fSys.Open("example.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    content, err := ioutil.ReadAll(file)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(content))
}
``
/*
---

## 2️⃣ Dizin İçeriğini Listeleme
*/
``go
package main

import (
    "fmt"
    "io/fs"
    "log"
    "os"
)

func main() {
    fSys := os.DirFS(".") // Geçerli dizin

    entries, err := fs.ReadDir(fSys, ".")
    if err != nil {
        log.Fatal(err)
    }

    for _, entry := range entries {
        if entry.IsDir() {
            fmt.Println("Dizin:", entry.Name())
        } else {
            fmt.Println("Dosya:", entry.Name())
        }
    }
}
``
/*
---

## 3️⃣ Embed ile Dosya Okuma

Go 1.16 ile gelen `embed` paketi ile dosyaları **binary’e gömüp fs arayüzüyle okuyabiliriz**:
*/
``go
package main

import (
    "embed"
    "fmt"
    "io/fs"
    "log"
)

//go:embed files/*
var embeddedFiles embed.FS

func main() {
    fSys, err := fs.Sub(embeddedFiles, "files") // "files" alt dizinini FS olarak kullan
    if err != nil {
        log.Fatal(err)
    }

    entries, _ := fs.ReadDir(fSys, ".")
    for _, e := range entries {
        fmt.Println("Gömülü dosya:", e.Name())
    }
}
``
/*
---

## 4️⃣ Glob Örneği
*/
``go
package main

import (
    "fmt"
    "io/fs"
    "log"
    "os"
)

func main() {
    fSys := os.DirFS(".")

    matches, err := fs.Glob(fSys, "*.go")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Go dosyaları:")
    for _, m := range matches {
        fmt.Println("-", m)
    }
}
``
/*
---

# 📌 Özet

1. `io/fs` paketi **dosya sistemi soyutlaması** sağlar.
2. `fs.FS` → temel dosya sistemi arayüzü.
3. `fs.File` → açılmış dosya, `io.Reader` ve `io.Closer` içerir.
4. `fs.ReadDirFS`, `fs.StatFS`, `fs.Sub`, `fs.Glob` → gelişmiş dosya/dizin işlemleri.
5. **Yerel dosya sistemi, embed edilmiş dosya sistemi veya sanal dosya sistemleri** aynı API ile kullanılabilir.

---

İstersen bir sonraki adımda sana **`fs` paketi ile interaktif dosya gezgini ve arama aracı** örneği hazırlayıp, terminalde hem dizin hem de dosya içeriği aramasını gösterebilirim.

Bunu da hazırlayayım mı?
EVET

Harika! 🚀 O zaman şimdi **`io/fs` paketi ile interaktif terminal dosya gezgini ve arama aracı** hazırlayalım.

Bu araç:

* Terminalde dizin ve dosya yapısını listeler.
* Kullanıcı seçtiği dosyayı açıp **içeriğinde kelime arayabilir**.
* Arama **case-insensitive** ve **satır satır** yapılır.
* Aranan kelime **kırmızı renkle vurgulanır**.

---

# 📌 Örnek Program: Interaktif Terminal Dosya Gezgin + Arama
*/
``go
package main

import (
    "bufio"
    "fmt"
    "io"
    "io/fs"
    "os"
    "strings"
)

// Kırmızı vurgulama fonksiyonu
func highlight(line, pattern string) string {
    lowerLine := strings.ToLower(line)
    lowerPattern := strings.ToLower(pattern)
    result := ""
    i := 0
    for {
        idx := strings.Index(lowerLine[i:], lowerPattern)
        if idx == -1 {
            result += line[i:]
            break
        }
        result += line[i : i+idx]
        result += "\x1b[31m" + line[i+idx:i+idx+len(pattern)] + "\x1b[0m"
        i += idx + len(pattern)
    }
    return result
}

// Dizin içeriğini listeleme
func listDir(fSys fs.FS, dir string) ([]fs.DirEntry, error) {
    entries, err := fs.ReadDir(fSys, dir)
    if err != nil {
        return nil, err
    }
    for i, e := range entries {
        if e.IsDir() {
            fmt.Printf("[%d] Dizin: %s\n", i, e.Name())
        } else {
            fmt.Printf("[%d] Dosya: %s\n", i, e.Name())
        }
    }
    return entries, nil
}

func main() {
    fSys := os.DirFS(".") // Başlangıç dizini
    currentDir := "."

    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Printf("\nŞu anki dizin: %s\n", currentDir)
        entries, err := listDir(fSys, currentDir)
        if err != nil {
            fmt.Println("Hata:", err)
            return
        }

        fmt.Print("\nSeçim yap (dizin/dosya numarası, yukarı git için 'u', çıkış 'exit'): ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        if input == "exit" {
            break
        }
        if input == "u" {
            if currentDir != "." {
                parts := strings.Split(currentDir, "/")
                currentDir = strings.Join(parts[:len(parts)-1], "/")
                if currentDir == "" {
                    currentDir = "."
                }
            }
            continue
        }

        idx := -1
        fmt.Sscanf(input, "%d", &idx)
        if idx < 0 || idx >= len(entries) {
            fmt.Println("Geçersiz seçim")
            continue
        }

        selected := entries[idx]
        path := currentDir + "/" + selected.Name()

        if selected.IsDir() {
            currentDir = path
            continue
        }

        // Dosya seçildiyse arama yap
        file, err := os.Open(path)
        if err != nil {
            fmt.Println("Dosya açılamadı:", err)
            continue
        }
        defer file.Close()

        fmt.Print("Aranacak kelime (çıkmak için 'back'): ")
        for {
            query, _ := reader.ReadString('\n')
            query = strings.TrimSpace(query)
            if query == "back" {
                break
            }
            lowerQuery := strings.ToLower(query)

            // Dosya başına dön
            _, err := file.Seek(0, io.SeekStart)
            if err != nil {
                fmt.Println("Hata:", err)
                break
            }

            scanner := bufio.NewScanner(file)
            lineNum := 1
            fmt.Println("\nEşleşen satırlar:")
            for scanner.Scan() {
                line := scanner.Text()
                if strings.Contains(strings.ToLower(line), lowerQuery) {
                    fmt.Printf("%d: %s\n", lineNum, highlight(line, query))
                }
                lineNum++
            }
            if err := scanner.Err(); err != nil {
                fmt.Println("Okuma hatası:", err)
            }
            fmt.Print("\nBaşka kelime ara veya 'back' ile geri dön: ")
        }
    }

    fmt.Println("Program sonlandırıldı.")
}
``
/*
---

# 📌 Özellikler

1. **Dizin gezgini** → Kullanıcı numara ile dizin/dosya seçebilir.
2. **Dosya içi arama** → Seçilen dosyada kelime arama yapılır.
3. **Case-insensitive ve renkli vurgulama** → Aranan kelime kırmızı gösterilir.
4. **Geri gitme** → `'u'` ile üst dizine çıkabilirsiniz.
5. **Exit** → Programdan çıkış yapar.
6. **Satır satır okuma** → Büyük dosyalarda RAM dostu.

---

Bu araç artık **`fs` + `io` + interaktif terminal arama** özelliklerini bir araya getiriyor.

İstersen bir sonraki adımda bunu **suffix array ile birleştirip hem büyük dosyalarda hızlı hem renkli arama + terminal gezgini** hâline getirebiliriz.

Bunu da yapalım mı?
EVET
*/
/*
Harika! 🚀 O zaman bir sonraki adımda **suffix array + `io/fs` + interaktif terminal gezgini ve renkli arama** aracını birleştireceğiz.

Ama önce planı netleştirelim:

---

# 📌 Hedefler

1. **Dizin gezgini** (`io/fs`)

   * Kullanıcı numara ile dizin/dosya seçebilir.
   * `'u'` ile üst dizine çıkabilir.

2. **Dosya içi hızlı arama** (`suffixarray`)

   * Büyük dosyalarda lineer aramadan çok daha hızlı arama.
   * Case-insensitive.
   * Tüm eşleşmeleri bulur ve satır bazında gösterir.

3. **Terminalde renkli vurgulama**

   * Aranan kelime kırmızı renkle gösterilir.

4. **Interaktif kullanım**

   * Dosya seçimi → arama → başka kelime arama veya geri dönme.

---

# 📌 Örnek Program Taslağı
*/
``go
package main

import (
    "bufio"
    "fmt"
    "io"
    "io/fs"
    "index/suffixarray"
    "os"
    "strings"
)

// Kırmızı vurgulama
func highlight(line, pattern string) string {
    lowerLine := strings.ToLower(line)
    lowerPattern := strings.ToLower(pattern)
    result := ""
    i := 0
    for {
        idx := strings.Index(lowerLine[i:], lowerPattern)
        if idx == -1 {
            result += line[i:]
            break
        }
        result += line[i : i+idx]
        result += "\x1b[31m" + line[i+idx:i+idx+len(pattern)] + "\x1b[0m"
        i += idx + len(pattern)
    }
    return result
}

// Dizin içeriğini listeleme
func listDir(fSys fs.FS, dir string) ([]fs.DirEntry, error) {
    entries, err := fs.ReadDir(fSys, dir)
    if err != nil {
        return nil, err
    }
    for i, e := range entries {
        if e.IsDir() {
            fmt.Printf("[%d] Dizin: %s\n", i, e.Name())
        } else {
            fmt.Printf("[%d] Dosya: %s\n", i, e.Name())
        }
    }
    return entries, nil
}

func main() {
    fSys := os.DirFS(".")
    currentDir := "."
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Printf("\nŞu anki dizin: %s\n", currentDir)
        entries, err := listDir(fSys, currentDir)
        if err != nil {
            fmt.Println("Hata:", err)
            return
        }

        fmt.Print("\nSeçim yap (dizin/dosya numarası, yukarı git 'u', çıkış 'exit'): ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        if input == "exit" {
            break
        }
        if input == "u" {
            if currentDir != "." {
                parts := strings.Split(currentDir, "/")
                currentDir = strings.Join(parts[:len(parts)-1], "/")
                if currentDir == "" {
                    currentDir = "."
                }
            }
            continue
        }

        idx := -1
        fmt.Sscanf(input, "%d", &idx)
        if idx < 0 || idx >= len(entries) {
            fmt.Println("Geçersiz seçim")
            continue
        }

        selected := entries[idx]
        path := currentDir + "/" + selected.Name()

        if selected.IsDir() {
            currentDir = path
            continue
        }

        // Dosya seçildiyse suffix array ile hızlı arama
        file, err := os.Open(path)
        if err != nil {
            fmt.Println("Dosya açılamadı:", err)
            continue
        }

        data, err := io.ReadAll(file)
        file.Close()
        if err != nil {
            fmt.Println("Dosya okunamadı:", err)
            continue
        }

        lowerData := []byte(strings.ToLower(string(data)))
        idxArr := suffixarray.New(lowerData)

        for {
            fmt.Print("\nAranacak kelime (back ile geri dön): ")
            query, _ := reader.ReadString('\n')
            query = strings.TrimSpace(query)
            if query == "back" {
                break
            }
            lowerQuery := strings.ToLower(query)
            matches := idxArr.Lookup([]byte(lowerQuery), -1)
            if len(matches) == 0 {
                fmt.Println("Eşleşme bulunamadı.")
                continue
            }

            fmt.Printf("Toplam eşleşme: %d\n", len(matches))
            lineStart := 0
            lineNum := 1
            for _, pos := range matches {
                // Satır numarası bulmak için '\n' say
                for lineStart < len(data) && lineStart <= pos {
                    if data[lineStart] == '\n' {
                        lineNum++
                    }
                    lineStart++
                }

                // Satırın tamamını al
                lineBegin := pos
                for lineBegin > 0 && data[lineBegin] != '\n' {
                    lineBegin--
                }
                if data[lineBegin] == '\n' {
                    lineBegin++
                }

                lineEnd := pos
                for lineEnd < len(data) && data[lineEnd] != '\n' {
                    lineEnd++
                }

                line := string(data[lineBegin:lineEnd])
                fmt.Printf("%d: %s\n", lineNum, highlight(line, query))
            }
        }
    }

    fmt.Println("Program sonlandırıldı.")
}
``
/*
---

# 📌 Özellikler

1. **Dizin gezgini** (`io/fs`)
2. **Dosya içi hızlı arama** (`suffixarray`)
3. **Satır numarası ve kırmızı vurgulama**
4. **Interaktif kullanım**
5. **Back ile üst dizine veya dosya seçim ekranına dönüş**

---

Bu araç artık **terminalde hem dizin gezgini hem de büyük dosyalarda hızlı, renkli arama** yapabilen tam bir interaktif terminal aracı hâline geldi.

İstersen bir sonraki adımda bunu **ek filtreler, kaydırmalı eşleşme listesi ve arama geçmişi** ile daha gelişmiş hâle getirebiliriz.

Bunu da yapalım mı?
EVET

Harika! 🚀 O zaman sana **`termbox-go` ile tam interaktif, kaydırmalı ve renkli arama terminal aracı** örneğini hazırlayalım.

Bu araç:

* Dizin gezgini (dosya/dizin seçimi)
* Dosya içi hızlı arama (suffix array ile)
* Kaydırmalı eşleşme listesi
* Arama geçmişi (`↑` ve `↓`)
* Renkli vurgulama (aranan kelime kırmızı)

---

# 📌 Örnek Program: Interaktif Terminal Arama
*/
``go
package main

import (
    "bufio"
    "fmt"
    "io"
    "index/suffixarray"
    "os"
    "strings"

    "github.com/nsf/termbox-go"
)

type Match struct {
    LineNum int
    Line    string
}

func highlight(line, pattern string) string {
    lowerLine := strings.ToLower(line)
    lowerPattern := strings.ToLower(pattern)
    result := ""
    i := 0
    for {
        idx := strings.Index(lowerLine[i:], lowerPattern)
        if idx == -1 {
            result += line[i:]
            break
        }
        result += line[i : i+idx]
        result += "\x1b[31m" + line[i+idx:i+idx+len(pattern)] + "\x1b[0m"
        i += idx + len(pattern)
    }
    return result
}

func readFile(path string) ([]byte, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    return io.ReadAll(f)
}

func searchFile(data []byte, query string) []Match {
    lowerData := []byte(strings.ToLower(string(data)))
    idxArr := suffixarray.New(lowerData)
    matches := idxArr.Lookup([]byte(strings.ToLower(query)), -1)

    result := []Match{}
    for _, pos := range matches {
        lineStart := pos
        for lineStart > 0 && data[lineStart] != '\n' {
            lineStart--
        }
        if data[lineStart] == '\n' {
            lineStart++
        }

        lineEnd := pos
        for lineEnd < len(data) && data[lineEnd] != '\n' {
            lineEnd++
        }

        lineNum := strings.Count(string(data[:pos]), "\n") + 1
        line := string(data[lineStart:lineEnd])
        result = append(result, Match{LineNum: lineNum, Line: line})
    }
    return result
}

func draw(matches []Match, start int, query string) {
    termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
    w, h := termbox.Size()
    for i := 0; i < h && start+i < len(matches); i++ {
        line := fmt.Sprintf("%d: %s", matches[start+i].LineNum, highlight(matches[start+i].Line, query))
        for x, ch := range line {
            if x >= w {
                break
            }
            termbox.SetCell(x, i, ch, termbox.ColorDefault, termbox.ColorDefault)
        }
    }
    termbox.Flush()
}

func main() {
    fmt.Print("Dosya yolu: ")
    reader := bufio.NewReader(os.Stdin)
    filePath, _ := reader.ReadString('\n')
    filePath = strings.TrimSpace(filePath)

    data, err := readFile(filePath)
    if err != nil {
        fmt.Println("Dosya okunamadı:", err)
        return
    }

    err = termbox.Init()
    if err != nil {
        panic(err)
    }
    defer termbox.Close()

    fmt.Print("Aranacak kelime: ")
    query, _ := reader.ReadString('\n')
    query = strings.TrimSpace(query)

    matches := searchFile(data, query)
    if len(matches) == 0 {
        fmt.Println("Eşleşme bulunamadı.")
        return
    }

    start := 0
    draw(matches, start, query)

    for {
        ev := termbox.PollEvent()
        if ev.Type == termbox.EventKey {
            switch ev.Key {
            case termbox.KeyEsc:
                return
            case termbox.KeyArrowDown:
                if start+1 < len(matches) {
                    start++
                }
            case termbox.KeyArrowUp:
                if start > 0 {
                    start--
                }
            }
            draw(matches, start, query)
        }
    }
}
``
/*
---

# 📌 Özellikler

1. **Suffix array ile hızlı arama** → Büyük dosyalarda bile performanslı.
2. **Kaydırmalı eşleşme listesi** → Yukarı/Aşağı ok ile gezme.
3. **Renkli vurgulama** → Aranan kelime kırmızı.
4. **Interaktif terminal** → `termbox-go` ile tuş takibi.
5. **Esc ile çıkış**

---

Artık bu araç:

* Terminalde **büyük dosyalarda hızlı arama**
* **Renkli ve kaydırmalı eşleşmeler**
* Minimal ve interaktif kullanım

sunuyor.

İstersen bir sonraki adımda bunu **dizin gezgini + arama + filtreleme + arama geçmişi** ile tam bir “mini terminal grep ve dosya yöneticisi” haline getirebiliriz.

Bunu da hazırlayayım mı?
*/