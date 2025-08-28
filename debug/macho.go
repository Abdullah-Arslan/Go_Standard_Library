/*
Go’daki **`debug/macho`** paketini baştan sona örneklerle anlatayım.

---

# 📌 `debug/macho` Paketi Nedir?

Go’nun **`debug/macho`** paketi, **Mach-O (Mach Object)** formatındaki dosyaları okumak için kullanılır.
Mach-O, **macOS** ve **iOS** sistemlerinde kullanılan standart çalıştırılabilir dosya formatıdır.
Yani Linux’ta ELF, Windows’ta PE/COFF neyse, macOS tarafında **Mach-O** odur.

Go’daki `debug/macho` paketi, ELF için `debug/elf`, Windows için `debug/pe` gibi **Mach-O dosyalarının header, section, segment, sembol tablolarını** çözümlemenizi sağlar.

---

# 📌 Önemli Yapılar

`debug/macho` ile açılan bir dosya şu bilgileri içerir:

* **`FileHeader`** → Mach-O dosyasının temel bilgileri (magic, CPU tipi, flags, dosya tipi)
* **`Load`** → Yükleme komutları (segmentler, dylib bağımlılıkları, vs.)
* **`Segment`** → Segment bilgileri (kod, veri, stack, vs.)
* **`Section`** → Segment içindeki bölümler (`__TEXT.__text`, `__DATA.__data`, vs.)
* **`Symtab`** → Sembol tablosu (fonksiyon, global değişkenler)
* **`Dysymtab`** → Dinamik semboller (paylaşımlı kütüphane bağları)
* **`Dylib`** → Kullanılan kütüphaneler (örn: `/usr/lib/libSystem.B.dylib`)

---

# 📌 Mach-O Dosyasını Açma
*/
``go
package main

import (
    "debug/macho"
    "fmt"
    "log"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatalf("Kullanım: %s <mach-o dosyası>\n", os.Args[0])
    }

    file := os.Args[1]
    f, err := macho.Open(file)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    fmt.Printf("CPU: %v\n", f.Cpu)
    fmt.Printf("Dosya tipi: %v\n", f.Type)
    fmt.Printf("Bayraklar: 0x%x\n", f.Flags)
}
``

/*
📌 Çıktı örneği (`/bin/ls` için macOS’ta):

```
CPU: x86_64
Dosya tipi: EXECUTE
Bayraklar: 0x2000
```

---

# 📌 Segmentleri Listeleme

Mach-O dosyasında segmentler (ör. `__TEXT`, `__DATA`) vardır.
*/
``go
for _, seg := range f.Segments {
    fmt.Printf("Segment: %-10s VMAddr=0x%x FileOff=0x%x Size=%d\n",
        seg.Name, seg.Addr, seg.Offset, seg.Filesz)
}
``
7?
Örnek çıktı:

```
Segment: __TEXT      VMAddr=0x100000000 FileOff=0x0 Size=4096
Segment: __DATA      VMAddr=0x100001000 FileOff=0x1000 Size=2048
```

---

# 📌 Section’ları Listeleme

Her segment içinde **section** bulunur.
*/

``go
for _, sec := range f.Sections {
    fmt.Printf("Section: %-20s Segment=%s Addr=0x%x Size=%d\n",
        sec.Name, sec.Seg, sec.Addr, sec.Size)
}
``
/*
Örnek çıktı:

```
Section: __text               Segment=__TEXT Addr=0x100001000 Size=1536
Section: __cstring            Segment=__TEXT Addr=0x100002000 Size=256
Section: __data               Segment=__DATA Addr=0x100003000 Size=128
```

---

# 📌 Sembol Tablosunu Okuma

Mach-O dosyaları fonksiyon ve sembol tablolarını içerir.
*/

``go
if f.Symtab != nil {
    for _, sym := range f.Symtab.Syms {
        fmt.Printf("Sembol: %-30s Value=0x%x Type=%x\n", sym.Name, sym.Value, sym.Type)
    }
}
``
/*
Örnek çıktı:

```
Sembol: _main                         Value=0x1000010f0 Type=0xf
Sembol: _printf                       Value=0x0 Type=0x1
```

---

# 📌 Dinamik Kütüphaneleri Listeleme

Mach-O çalıştırılabilirleri hangi dylib’lere bağlıdır?
*/

``go
for _, dylib := range f.Dylibs {
    fmt.Println("Dylib:", dylib.Name)
}
``
/*
Örnek çıktı:

```
Dylib: /usr/lib/libSystem.B.dylib
Dylib: /System/Library/Frameworks/CoreFoundation.framework/CoreFoundation
```

---

# 📌 Tam Örnek – Mach-O İnceleme Aracı
*/
``go
package main

import (
    "debug/macho"
    "fmt"
    "log"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatalf("Kullanım: %s <mach-o dosyası>\n", os.Args[0])
    }

    file := os.Args[1]
    f, err := macho.Open(file)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    fmt.Println("== MACH-O HEADER ==")
    fmt.Printf("CPU: %v, Type: %v, Flags: 0x%x\n", f.Cpu, f.Type, f.Flags)

    fmt.Println("\n== SEGMENTS ==")
    for _, seg := range f.Segments {
        fmt.Printf("Segment: %-10s Addr=0x%x Size=%d\n", seg.Name, seg.Addr, seg.Filesz)
    }

    fmt.Println("\n== SECTIONS ==")
    for _, sec := range f.Sections {
        fmt.Printf("Section: %-20s Segment=%s Addr=0x%x Size=%d\n",
            sec.Name, sec.Seg, sec.Addr, sec.Size)
    }

    fmt.Println("\n== SYMBOLS ==")
    if f.Symtab != nil {
        for _, sym := range f.Symtab.Syms {
            fmt.Printf("%-30s 0x%x\n", sym.Name, sym.Value)
        }
    } else {
        fmt.Println("Sembol tablosu bulunamadı.")
    }

    fmt.Println("\n== DYLIBS ==")
    for _, dylib := range f.Dylibs {
        fmt.Println(dylib.Name)
    }
}
``

//📌 Kullanım:

``bash
go run macho_reader.go /bin/ls
``
/*
---

# 📌 Özet

`debug/macho` paketi ile:

* ✅ Mach-O dosyalarının **header** bilgilerini okuyabilirsiniz
* ✅ Segment ve section bilgilerini listeleyebilirsiniz
* ✅ Sembol tablosunu ve dinamik sembolleri alabilirsiniz
* ✅ Hangi **dylib**’lere bağlı olduğunu öğrenebilirsiniz

Bu sayede `otool -l`, `nm`, `otool -L` gibi macOS komutlarının Go ile yazılmış mini versiyonunu yapabilirsin.

---

👉 İstersen sana Linux’taki `readelf` benzeri ama macOS için bir **`otool` klonu** (Go ile yazılmış Mach-O analiz aracı) kodlayabilirim. Bunu ister misin?
*/