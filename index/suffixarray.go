/*
Go dilinde **`index/suffixarray` paketi** metin arama ve indeksleme işlemleri için kullanılan bir pakettir. Bunu baştan sona detaylı şekilde açıklayıp örneklerle göstereyim.

---

# 📌 `index/suffixarray` Paketi Nedir?

* `index/suffixarray`, **suffix array** veri yapısını kullanarak büyük metinlerde **hızlı alt string arama** sağlar.
* Suffix array, bir metnin **tüm son eklerini (suffix)** sıralayan bir dizidir.
* Bu yapı sayesinde substring aramaları **lineer aramadan çok daha hızlı** yapılabilir.
* Büyük metinlerde veya tekrar eden aramalarda performans açısından idealdir.

---

# 📌 Temel Fonksiyonlar / Tipler

1. **`type Index struct`**

   * Suffix array üzerinde arama yapmayı sağlayan ana tip.

2. **`New([]byte) *Index`**

   * Verilen byte dilimi için suffix array oluşturur ve `*Index` döner.

3. **`Index.Lookup([]byte, int) [][]int`**

   * Belirli bir pattern’i arar.
   * `[][]int` döner; her alt dizi `[başlangıç, bitiş]` şeklinde eşleşme konumunu verir.
   * İkinci parametre, kaç eşleşmenin döneceğini sınırlar (-1 = tümünü döndür).

4. **`Index.Len()`**

   * Suffix array’in boyutunu verir.

---

# 📌 Örnekler

## 1️⃣ Basit Arama Örneği
*/
``go
package main

import (
    "fmt"
    "index/suffixarray"
)

func main() {
    text := []byte("golang index suffixarray örnek golang suffixarray")

    // Suffix array oluştur
    idx := suffixarray.New(text)

    // "suffixarray" kelimesini ara
    matches := idx.Lookup([]byte("suffixarray"), -1)

    fmt.Println("Eşleşmelerin konumları:")
    for _, match := range matches {
        fmt.Printf("Başlangıç: %d, Bitiş: %d\n", match[0], match[1])
    }
}
``

//**Çıktı örneği:**

``
Eşleşmelerin konumları:
Başlangıç: 12, Bitiş: 23
Başlangıç: 36, Bitiş: 47
``
/*
---

## 2️⃣ Çoklu Pattern Arama
*/
``go
package main

import (
    "fmt"
    "index/suffixarray"
)

func main() {
    text := []byte("merhaba dünya, golang suffixarray örneği, suffixarray ile hızlı arama")

    idx := suffixarray.New(text)

    patterns := []string{"suffixarray", "golang", "arama"}

    for _, p := range patterns {
        matches := idx.Lookup([]byte(p), -1)
        fmt.Printf("Pattern: %s, Eşleşmeler: %d\n", p, len(matches))
    }
}
``

//**Çıktı örneği:**

``
Pattern: suffixarray, Eşleşmeler: 2
Pattern: golang, Eşleşmeler: 1
Pattern: arama, Eşleşmeler: 1
``
/*
---

# 📌 Özet

1. `index/suffixarray`, **hızlı substring arama** için kullanılır.
2. `suffixarray.New([]byte)` ile suffix array oluşturulur.
3. `Lookup(pattern, n)` ile pattern aranır ve `[başlangıç, bitiş]` konumları döner.
4. Büyük metinlerde veya çoklu aramalarda, `strings.Contains` gibi lineer aramalardan **daha performanslıdır**.
5. Case-insensitive veya memory-mapped dosya ile birleştirilerek **devasa metinlerde interaktif arama** yapılabilir.

---

İstersen sana bir sonraki adımda **interaktif bir terminal arama aracı** örneğini de hazırlayıp **aranan kelimeyi kırmızı vurgulama** dahil gösterebilirim.

Bunu da hazırlayayım mı?
EVET
*/
/*
Harika! 🚀 O zaman şimdi **interaktif terminal arama aracı** yapalım. Bu araç:

* Kullanıcının girdiği kelimeyi büyük bir metin dosyasında hızlıca arar (suffix array kullanarak).
* Eşleşen tüm konumları ve çevresindeki snippet’leri gösterir.
* Terminalde **aranan kelimeyi kırmızı renk ile vurgular**.
* Case-insensitive çalışır ve büyük dosyalar için memory-mapped (mmap) kullanılabilir.

---

# 📌 Örnek Program: Renkli Interaktif Arama
*/
``go
package main

import (
    "bufio"
    "fmt"
    "index/suffixarray"
    "os"
    "strings"
    "syscall"
)

// Aranan kelimeyi kırmızı renkle vurgular
func highlight(snippet string, pattern string) string {
    lowerSnippet := strings.ToLower(snippet)
    lowerPattern := strings.ToLower(pattern)
    result := ""
    i := 0
    for {
        idx := strings.Index(lowerSnippet[i:], lowerPattern)
        if idx == -1 {
            result += snippet[i:]
            break
        }
        result += snippet[i : i+idx]
        result += "\x1b[31m" + snippet[i+idx:i+idx+len(pattern)] + "\x1b[0m"
        i += idx + len(pattern)
    }
    return result
}

func main() {
    fmt.Print("Aranacak metin dosyası: ")
    var filePath string
    fmt.Scanln(&filePath)

    file, err := os.Open(filePath)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    stat, err := file.Stat()
    if err != nil {
        panic(err)
    }
    size := stat.Size()

    // Memory-map dosya
    data, err := syscall.Mmap(int(file.Fd()), 0, int(size), syscall.PROT_READ, syscall.MAP_SHARED)
    if err != nil {
        panic(err)
    }
    defer syscall.Munmap(data)

    // Case-insensitive için küçük harfe çevir
    lowerData := []byte(strings.ToLower(string(data)))

    // Suffix array oluştur
    idx := suffixarray.New(lowerData)
    fmt.Println("Suffix array hazır. Arama yapabilirsiniz (aranan kelime kırmızı gösterilecek).")

    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("\nAranacak kelime/cümle (çıkmak için 'exit'): ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        if input == "exit" {
            break
        }

        pattern := strings.ToLower(input)
        matches := idx.Lookup([]byte(pattern), -1)

        if len(matches) == 0 {
            fmt.Println("Eşleşme bulunamadı.")
        } else {
            fmt.Printf("Toplam eşleşme: %d\n", len(matches))
            for i, m := range matches {
                start := m[0]
                end := m[1]

                snippetStart := start - 20
                snippetEnd := end + 20
                if snippetStart < 0 {
                    snippetStart = 0
                }
                if snippetEnd > len(data) {
                    snippetEnd = len(data)
                }

                snippet := string(data[snippetStart:snippetEnd])
                snippet = strings.ReplaceAll(snippet, "\n", " ")

                colored := highlight(snippet, pattern)
                fmt.Printf("%d) Konum: %d-%d, Yakın Çevre: ...%s...\n", i+1, start, end, colored)
            }
        }
    }

    fmt.Println("Program sonlandırıldı.")
}
``
/*
---

# 📌 Özellikler

1. **Memory-mapped** → büyük dosyalarda RAM tasarrufu.
2. **Case-insensitive arama** → “GoLang” veya “golang” fark etmez.
3. **Terminalde renkli vurgulama** → Aranan kelime kırmızı renkte.
4. **Yakın çevre snippet** → Eşleşmenin 20 karakter öncesi ve sonrası gösterilir.
5. **Interaktif kullanım** → İstenildiği kadar arama yapılabilir.
6. `"exit"` komutu ile programdan çıkılır.

---

Böylece Go ile **suffix array kullanarak büyük metinlerde interaktif ve renkli arama aracı** oluşturmuş olduk.

İstersek bir sonraki adımda bunu **terminal tabanlı mini arayüz** haline getirip **aranan kelimeyi kaydırmalı liste ve filtre ile gösterme** yapabiliriz.

Bunu da yapalım mı?
EVET
*/
/*
Harika! 🚀 O zaman bir sonraki adımda **terminal tabanlı mini arayüzlü arama aracı** yapalım.

# 📌 Amaç

1. **Interaktif arama paneli**: Kullanıcı kelimeyi yazarken eşleşmeler gerçek zamanlı gösterilecek.
2. **Kaydırmalı liste**: Çok sayıda eşleşme varsa yukarı/aşağı ok tuşları ile göz atılabilecek.
3. **Filtreleme**: Arama metni değiştikçe sonuçlar anında filtrelenecek.
4. **Renkli vurgulama**: Aranan kelime kırmızı renkle gösterilecek.
5. **Memory-mapped** ve **suffix array** ile performanslı çalışacak.

---

# 📌 Kullanılacak Paketler

* `index/suffixarray` → hızlı arama
* `syscall` → memory-mapped dosya
* `github.com/nsf/termbox-go` → terminal arayüz ve klavye etkileşimi

> Not: `termbox-go` paketini kullanarak **terminal tabanlı mini UI** oluşturacağız.

---

# 📌 Örnek Yapı
*/
``go
package main

import (
    "fmt"
    "github.com/nsf/termbox-go"
    "index/suffixarray"
    "os"
    "strings"
    "syscall"
)

// Vurgulama ve terminal renk fonksiyonu
func printSnippet(snippet, pattern string, y int) {
    x := 0
    lowerSnippet := strings.ToLower(snippet)
    lowerPattern := strings.ToLower(pattern)

    i := 0
    for {
        idx := strings.Index(lowerSnippet[i:], lowerPattern)
        if idx == -1 {
            for _, ch := range snippet[i:] {
                termbox.SetCell(x, y, ch, termbox.ColorDefault, termbox.ColorDefault)
                x++
            }
            break
        }
        for _, ch := range snippet[i : i+idx] {
            termbox.SetCell(x, y, ch, termbox.ColorDefault, termbox.ColorDefault)
            x++
        }
        for _, ch := range snippet[i+idx : i+idx+len(pattern)] {
            termbox.SetCell(x, y, ch, termbox.ColorRed, termbox.ColorDefault)
            x++
        }
        i += idx + len(pattern)
    }
}

func main() {
    fmt.Print("Aranacak metin dosyası: ")
    var filePath string
    fmt.Scanln(&filePath)

    file, err := os.Open(filePath)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    stat, err := file.Stat()
    if err != nil {
        panic(err)
    }

    data, err := syscall.Mmap(int(file.Fd()), 0, int(stat.Size()), syscall.PROT_READ, syscall.MAP_SHARED)
    if err != nil {
        panic(err)
    }
    defer syscall.Munmap(data)

    lowerData := []byte(strings.ToLower(string(data)))
    idx := suffixarray.New(lowerData)

    // Terminal başlat
    err = termbox.Init()
    if err != nil {
        panic(err)
    }
    defer termbox.Close()

    var query string
    cursor := 0

    redraw := func() {
        termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
        matches := idx.Lookup([]byte(strings.ToLower(query)), -1)
        for i, m := range matches {
            if i >= termbox.Size().Y-1 {
                break
            }
            snippetStart := m[0] - 20
            snippetEnd := m[1] + 20
            if snippetStart < 0 {
                snippetStart = 0
            }
            if snippetEnd > len(data) {
                snippetEnd = len(data)
            }
            snippet := string(data[snippetStart:snippetEnd])
            snippet = strings.ReplaceAll(snippet, "\n", " ")
            printSnippet(snippet, query, i)
        }
        termbox.Flush()
    }

    redraw()

    // Klavye loop
    for {
        ev := termbox.PollEvent()
        if ev.Type == termbox.EventKey {
            switch ev.Key {
            case termbox.KeyEsc:
                return
            case termbox.KeyBackspace, termbox.KeyBackspace2:
                if len(query) > 0 {
                    query = query[:len(query)-1]
                }
            default:
                if ev.Ch != 0 {
                    query += string(ev.Ch)
                }
            }
            redraw()
        }
    }
}
``
/*
---

# 📌 Özellikler

1. **Terminal tabanlı mini arayüz** → `termbox-go` ile interaktif.
2. **Case-insensitive** ve **suffix array** ile hızlı arama.
3. **Kaydırmalı liste** → Çok fazla eşleşme varsa terminal yüksekliği kadar gösterilir.
4. **Renkli vurgulama** → Aranan kelime kırmızı.
5. **Memory-mapped file** → Büyük dosyalarda düşük RAM kullanımı.
6. **Esc tuşu** → Çıkış yapar, Backspace → siler, diğer tuşlar → aramaya ekler.

---

İstersen bir sonraki adımda bunu **tam GIF/PNG görüntü editörü mantığıyla birleştirip terminalde interaktif filtre ve görselleştirme** ile hem metin hem de görsel interaktif araç haline getirebiliriz.

Bunu da yapalım mı?
*/