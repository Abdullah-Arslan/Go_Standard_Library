/*
Go dilinde **`io` paketi**, giriş/çıkış işlemlerini (input/output) soyutlamak ve standartlaştırmak için kullanılan temel bir pakettir. Bu paketi baştan sona detaylı şekilde açıklayıp örneklerle göstereyim.

---

# 📌 `io` Paketi Nedir?

* `io` paketi, **okuma ve yazma işlemleri için temel arayüzleri** ve yardımcı fonksiyonları sağlar.
* Dosya, network, bellek gibi farklı kaynaklar üzerinde **ortak bir API ile** işlem yapılmasını sağlar.
* `io` paketi **arabellekler, kopyalama, sınırlandırma, çoklu yazma/okuma** gibi işlemleri de destekler.

---

# 📌 Önemli Arayüzler

1. **`io.Reader`**

   * Veri okuma için kullanılır.
   * Tek bir metod vardır:

     ```go
     Read(p []byte) (n int, err error)
     ```
   * Örnek: Dosya, string, network gibi kaynaklardan veri okunabilir.

2. **`io.Writer`**

   * Veri yazmak için kullanılır.
   * Tek bir metod vardır:

     ```go
     Write(p []byte) (n int, err error)
     ```
   * Örnek: Dosya, buffer, network gibi hedeflere veri yazabilir.

3. **`io.Closer`**

   * Kaynağı kapatmak için kullanılır.

     ```go
     Close() error
     ```

4. **`io.Seeker`**

   * Okuma/yazma konumunu değiştirmek için kullanılır.
*/
     ``go
     Seek(offset int64, whence int) (int64, error)
     ``
/*
5. **`io.ReadWriter`**

   * Hem `Reader` hem `Writer` arayüzlerini kapsar.

6. **`io.ReaderAt` / `io.WriterAt`**

   * Belirli bir konumdan okuma/yazma.

7. **`io.ReadCloser` / `io.WriteCloser`**

   * Hem okuma/yazma hem de kapatma işlemleri.

---

# 📌 Önemli Fonksiyonlar

1. **`io.Copy(dst io.Writer, src io.Reader) (written int64, err error)`**

   * `src`'tan `dst`'ye veri kopyalar.

2. **`io.CopyN(dst io.Writer, src io.Reader, n int64)`**

   * Sadece `n` byte kopyalar.

3. **`io.ReadFull(r io.Reader, buf []byte)`**

   * Tam buffer dolana kadar okuma yapar.

4. **`io.LimitReader(r io.Reader, n int64)`**

   * Sadece `n` byte okuyabilen reader döndürür.

5. **`io.MultiReader(r ...io.Reader)`**

   * Birden fazla reader’ı ardışık olarak okur.

6. **`io.MultiWriter(w ...io.Writer)`**

   * Aynı veriyi birden fazla yazıcıya yazar.

---

# 📌 Örnekler

## 1️⃣ Basit Reader Örneği
*/
``go
package main

import (
    "fmt"
    "io"
    "strings"
)

func main() {
    r := strings.NewReader("Merhaba Go io paketi!")

    buf := make([]byte, 8)
    for {
        n, err := r.Read(buf)
        if err == io.EOF {
            break
        }
        fmt.Printf("Okunan: %s\n", buf[:n])
    }
}
``

//**Çıktı örneği:**

``
Okunan: Merhaba 
Okunan: Go io pa
Okunan: keti!
``
/*
---

## 2️⃣ Dosya Kopyalama (`io.Copy`) Örneği
*/
``go
package main

import (
    "io"
    "log"
    "os"
)

func main() {
    srcFile, err := os.Open("source.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer srcFile.Close()

    dstFile, err := os.Create("destination.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer dstFile.Close()

    written, err := io.Copy(dstFile, srcFile)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("%d byte kopyalandı\n", written)
}
``
/*
---

## 3️⃣ MultiWriter Örneği (Birden fazla yere yazma)
*7
``go
package main

import (
    "io"
    "os"
)

func main() {
    f1, _ := os.Create("file1.txt")
    f2, _ := os.Create("file2.txt")
    defer f1.Close()
    defer f2.Close()

    mw := io.MultiWriter(f1, f2)
    mw.Write([]byte("Merhaba, her iki dosyaya da yazıldı!"))
}
``
/*
---

## 4️⃣ LimitReader Örneği
*/
``go
package main

import (
    "fmt"
    "io"
    "strings"
)

func main() {
    r := strings.NewReader("Bu uzun bir metin örneğidir.")
    lr := io.LimitReader(r, 10)

    buf := make([]byte, 20)
    n, _ := lr.Read(buf)
    fmt.Printf("Sadece 10 byte okundu: %s\n", buf[:n])
}
``

//**Çıktı:**

``
Sadece 10 byte okundu: Bu uzun bi
``
/*
---

# 📌 Özet

* `io` paketi **Go’nun giriş/çıkış temeli**dir.
* Dosya, bellek, network gibi tüm kaynaklar için **ortak Reader/Writer arayüzleri** sağlar.
* `io.Copy`, `io.MultiWriter`, `io.LimitReader` gibi yardımcı fonksiyonlar ile **çoklu, sınırlı ve verimli veri işleme** yapılabilir.
* `io.Reader` ve `io.Writer` arayüzlerini kullanarak **soyut ve modüler kod yazmak** mümkün.

---

İstersen bir sonraki adımda sana **`io` paketi ile interaktif terminal dosya okuma ve filtreleme örneği** de hazırlayabilirim; yani kullanıcı girdiği metni satır satır okuyup arama ve yazdırma yapabilir.

Bunu da hazırlayayım mı?
EVET
*/
/*
Harika! 🚀 O zaman şimdi **`io` paketi ile interaktif terminal dosya okuma ve filtreleme** örneğini hazırlayalım.

Bu araç:

* Kullanıcının girdiği **aranacak kelimeyi** satır satır metin dosyasında arar.
* Eşleşen satırları terminalde gösterir.
* Arama **case-insensitive** yapılır.
* `io.Reader` ve `bufio.Scanner` ile verimli satır okuma sağlanır.

---

# 📌 Örnek Program: Terminalde Satır Satır Arama
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

func main() {
    fmt.Print("Aranacak dosya: ")
    var filePath string
    fmt.Scanln(&filePath)

    file, err := os.Open(filePath)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    reader := bufio.NewReader(file)

    for {
        fmt.Print("\nAranacak kelime (çıkmak için 'exit'): ")
        input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
        input = strings.TrimSpace(input)
        if input == "exit" {
            break
        }
        lowerInput := strings.ToLower(input)

        // Dosya başına dön
        _, err := file.Seek(0, io.SeekStart)
        if err != nil {
            panic(err)
        }
        reader.Reset(file)

        lineNum := 1
        fmt.Println("\nEşleşen satırlar:")
        for {
            line, err := reader.ReadString('\n')
            if err != nil && err != io.EOF {
                panic(err)
            }

            if strings.Contains(strings.ToLower(line), lowerInput) {
                fmt.Printf("%d: %s", lineNum, line)
            }

            if err == io.EOF {
                break
            }
            lineNum++
        }
    }

    fmt.Println("Program sonlandırıldı.")
}
``
/*
---

# 📌 Açıklama

1. **`os.Open`** → Dosyayı açar.
2. **`bufio.NewReader`** → Satır satır okuma sağlar.
3. **`file.Seek(0, io.SeekStart)`** → Her aramada dosya başına dönülür.
4. **`strings.ToLower`** → Case-insensitive arama için hem satır hem input küçük harfe çevrilir.
5. Kullanıcı `"exit"` yazana kadar interaktif arama devam eder.

---

# 📌 Örnek Çıktı

Dosya içeriği:
*/
``
Merhaba Go
IO paketi örnek
golang io arayüzleri
Dosya okuma ve yazma
``

//Kullanıcı `io` kelimesini ararsa:

``
Eşleşen satırlar:
2: IO paketi örnek
3: golang io arayüzleri
``
/*
---

Böylece **`io` paketi + bufio ile interaktif arama ve filtreleme** yapılmış oldu.

İstersen bir sonraki adımda bunu **aranan kelimeyi terminalde renkli vurgulayan versiyona** da dönüştürebiliriz.

Bunu da yapalım mı?
EVET
*/
/*
Harika!  O zaman şimdi **`io` paketi ile interaktif ve renkli vurgulu terminal arama aracı** yapalım.

Bu araç:

* Dosyayı satır satır okur (`io.Reader` + `bufio.Scanner`).
* Kullanıcı arama kelimesini girer.
* Eşleşen satırları terminalde gösterir ve **aranan kelimeyi kırmızı renkle vurgular**.
* Case-insensitive arama yapar.

---

# 📌 Örnek Program: Renkli Vurgulu Terminal Arama
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

// Aranan kelimeyi kırmızı renkle vurgulama
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

func main() {
    fmt.Print("Aranacak dosya: ")
    var filePath string
    fmt.Scanln(&filePath)

    file, err := os.Open(filePath)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    reader := bufio.NewReader(file)

    for {
        fmt.Print("\nAranacak kelime (çıkmak için 'exit'): ")
        input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
        input = strings.TrimSpace(input)
        if input == "exit" {
            break
        }

        lowerInput := strings.ToLower(input)

        // Dosya başına dön
        _, err := file.Seek(0, io.SeekStart)
        if err != nil {
            panic(err)
        }
        reader.Reset(file)

        lineNum := 1
        fmt.Println("\nEşleşen satırlar:")
        for {
            line, err := reader.ReadString('\n')
            if err != nil && err != io.EOF {
                panic(err)
            }

            if strings.Contains(strings.ToLower(line), lowerInput) {
                fmt.Printf("%d: %s", lineNum, highlight(line, input))
            }

            if err == io.EOF {
                break
            }
            lineNum++
        }
    }

    fmt.Println("Program sonlandırıldı.")
}
``
/*
---

# 📌 Özellikler

1. **Satır satır okuma** → RAM kullanımı minimum.
2. **Case-insensitive arama** → Aranan kelime büyük/küçük fark etmez.
3. **Terminal renkli vurgulama** → Aranan kelime kırmızı renkte gösterilir.
4. **Interaktif kullanım** → Kullanıcı istediği kadar arama yapabilir.
5. `"exit"` ile çıkış yapılır.

---

Örnek çıktı:

Dosya içeriği:
*/
``
Merhaba Go
IO paketi örnek
golang io arayüzleri
Dosya okuma ve yazma
``

A//ranan kelime: `io`

``
Eşleşen satırlar:
2: [31mIO[0m paketi örnek
3: golang [31mio[0m arayüzleri
``
/*
---

Bu sayede artık **`io` paketi ile interaktif, satır bazlı ve renkli arama** yapabiliyoruz.

İstersen bir sonraki adımda bunu **dosya çok büyükse memory-mapped ve suffix array ile birleştirip hem hızlı hem renkli arama** haline getirebiliriz.

Bunu da yapalım mı?
*/