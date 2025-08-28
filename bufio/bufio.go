// Go dilinde bufio paketi, buffered (arabellekli) I/O işlemleri için kullanılan standart bir pakettir. 
// Bu paket, özellikle dosya veya ağdan veri okurken ve yazarken performansı artırmak için kullanılır. 
// Bufio paketi, verileri satır satır veya belirli bir boyutta okuma/yazma gibi işlemleri kolaylaştırır.

// 📦 bufio Paketinin Temel Amacı:

// I/O işlemlerinde daha az sistem çağrısı (syscall) yapmak.

// Küçük küçük veri okumak/işlemek yerine, büyük bloklar halinde okumak/yazmak.

// Okuma ve yazma işlemlerinde ek işlevsellik sağlamak (örneğin: satır satır okuma, belirli karaktere kadar okuma).

// Aşağıda bufio paketini örneklerle açıklayacağım:

// 📦 bufio Paketini Kullanmak

// Öncelikle paketi içeri aktar:

import (
    "bufio"
    "fmt"
    "os"
)

//🔹 1. bufio.NewReader ile Dosyadan Satır Satır Okuma

package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("ornek.txt")
    if err != nil {
        fmt.Println("Hata:", err)
        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)

    for {
        line, err := reader.ReadString('\n') // satır sonuna kadar oku
        if err != nil {
            break
        }
        fmt.Print("Okunan satır: ", line)
    }
}

// Açıklama:

// bufio.NewReader(file) → Dosya için bir buffered reader oluşturur.

// ReadString('\n') → Satır sonuna kadar okur.


//🔹 2. bufio.NewScanner ile Satır Satır Okuma (Kolay Yol)

package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("ornek.txt")
    if err != nil {
        fmt.Println("Hata:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        fmt.Println("Satır:", scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Okuma hatası:", err)
    }
}
// Açıklama:

// bufio.NewScanner → Dosya içeriğini satır satır okumak için pratik bir yoldur.

// scanner.Text() → Okunan satırın içeriğini verir.


//🔹 3. bufio.NewWriter ile Dosyaya Yazma

package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("yaz.txt")
    if err != nil {
        fmt.Println("Hata:", err)
        return
    }
    defer file.Close()

    writer := bufio.NewWriter(file)
    
    _, err = writer.WriteString("Merhaba, Dünya!\n")
    if err != nil {
        fmt.Println("Yazma hatası:", err)
        return
    }

    writer.Flush() // buffer’ı dosyaya yaz
}
// Açıklama:

// bufio.NewWriter(file) → Buffered writer oluşturur.

// writer.WriteString(...) → Buffer'a yazar.

// writer.Flush() → Buffer’daki verileri dosyaya aktarır (yoksa yazılmaz!).


//🔹 4. bufio.Reader.ReadBytes Kullanımı

reader := bufio.NewReader(file)
line, err := reader.ReadBytes('\n')
fmt.Println("Satır (byte dizisi):", line)

//Bu, satırı byte dilimi ([]byte) olarak almanı sağlar. Özellikle binary veri ile çalışırken kullanılır.


| Kullanım           | Açıklama                                |
| ------------------ | --------------------------------------- |
| `bufio.NewReader`  | Buffered okuma                          |
| `ReadString`       | Belirli bir ayırıcıya kadar okur        |
| `ReadBytes`        | Belirli bir ayırıcıya kadar byte dizisi |
| `bufio.NewScanner` | Satır satır okuma (kolay yol)           |
| `bufio.NewWriter`  | Buffered yazma                          |
| `Flush`            | Yazma işlemini tamamlar                 |
