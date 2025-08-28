/*
Go’daki **`debug/elf`** paketi, ELF (Executable and Linkable Format) dosyalarını okumak ve analiz etmek için kullanılır.
ELF formatı, Linux ve Unix tabanlı sistemlerde kullanılan standart çalıştırılabilir dosya, paylaşımlı kütüphane (`.so`), çekirdek modülü ve obje dosyası (`.o`) formatıdır.

Go’nun `debug/elf` paketi sayesinde bir ELF dosyasının **header** bilgilerini, **section** (bölüm) ve **segment** (program header) yapılarını, sembol tablolarını, string tablolarını ve diğer metadata bilgilerini okuyabilirsiniz.

Şimdi paketi baştan sona açıklayalım. 🚀

---

## 📌 Paket İmport
*/
``go
import (
    "debug/elf"
    "fmt"
    "log"
    "os"
)
``
/*
---

## 📌 ELF Dosyasını Açmak

`elf.Open(filename)` ile ELF dosyası açılır.
*/
``go
func main() {
    f, err := elf.Open("mybinary")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    fmt.Println("ELF Dosyası Açıldı:", f.FileHeader)
}
``
/*
📌 `f.FileHeader` size ELF dosyasının temel **header** bilgilerini verir:

* ELF class (32-bit / 64-bit)
* Endianness (little / big endian)
* Type (EXEC = executable, DYN = shared object, REL = relocatable)
* Machine (işlemci mimarisi, x86\_64, ARM, RISC-V vs.)
* Version

---

## 📌 ELF Header Bilgilerini Okumak

```go
fmt.Printf("Class: %v\n", f.Class)
fmt.Printf("Data: %v\n", f.Data)
fmt.Printf("Type: %v\n", f.Type)
fmt.Printf("Machine: %v\n", f.Machine)
```

Örnek çıktı:

```
Class: ELFCLASS64
Data: ELFDATA2LSB
Type: ET_EXEC
Machine: EM_X86_64
```

---

## 📌 Section’ları Listelemek

ELF dosyası **section** bölümlerine ayrılmıştır (kod, veri, semboller vs.).
*/

``go
for _, section := range f.Sections {
    fmt.Printf("Section: %-20s Size: %d bytes\n", section.Name, section.Size)
}
``
/*
Örnek:

```
Section: .text                Size: 12345 bytes
Section: .data                Size: 512 bytes
Section: .rodata              Size: 1024 bytes
Section: .symtab              Size: 2048 bytes
Section: .strtab              Size: 512 bytes
```

---

## 📌 Section İçeriğini Okumak
*/
``go
sec := f.Section(".text")
data, err := sec.Data()
if err != nil {
    log.Fatal(err)
}
fmt.Printf(".text section %d byte içeriyor.\n", len(data))
``
/*
---

## 📌 Program Header (Segment) Bilgileri

Segmentler, çalışma zamanında belleğe nasıl yükleneceğini tanımlar.
*/
``go
for _, prog := range f.Progs {
    fmt.Printf("Type: %v Offset: 0x%x Vaddr: 0x%x Filesz: %d Memsz: %d\n",
        prog.Type, prog.Off, prog.Vaddr, prog.Filesz, prog.Memsz)
}
``
/*
---

## 📌 Sembol Tablosunu Okumak

ELF dosyaları **semboller** içerir (fonksiyon, global değişken vs.).
*/
``go
symbols, err := f.Symbols()
if err != nil {
    log.Fatal(err)
}

for _, sym := range symbols {
    fmt.Printf("Name: %-30s Value: 0x%x Size: %d\n", sym.Name, sym.Value, sym.Size)
}
``
/*
Örnek çıktı:

``
Name: main                        Value: 0x401000 Size: 50
Name: printf                      Value: 0x0 Size: 0
``

---

## 📌 Dinamik Semboller (`.dynsym`)

Paylaşımlı kütüphanelerde (`.so`) dinamik semboller olur.
*/
``go
dynSymbols, err := f.DynamicSymbols()
if err == nil {
    for _, sym := range dynSymbols {
        fmt.Printf("DynSymbol: %s\n", sym.Name)
    }
}
``
/*
---

## 📌 Not Section (Metadata)

ELF dosyaları bazen ek metadata içerir (ör. `GNU Build ID`).
*/
``go
for _, note := range f.Notes {
    fmt.Printf("Note: Name=%s Type=%d Desc=%x\n", note.Name, note.Type, note.Desc)
}
``
/*
---

## 📌 String Table

String tabloları, sembol isimleri için kullanılır.
*/
``go
strSec := f.Section(".strtab")
strData, _ := strSec.Data()
fmt.Println("String Table Size:", len(strData))
``
/*
---

## 📌 Tam Örnek – ELF İnceleme Aracı
*/

``go
package main

import (
    "debug/elf"
    "fmt"
    "log"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatalf("Kullanım: %s <elf-dosyası>\n", os.Args[0])
    }

    f, err := elf.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    fmt.Println("== ELF HEADER ==")
    fmt.Printf("Class: %v, Data: %v, Type: %v, Machine: %v\n",
        f.Class, f.Data, f.Type, f.Machine)

    fmt.Println("\n== SECTIONS ==")
    for _, sec := range f.Sections {
        fmt.Printf("%-20s %d bytes\n", sec.Name, sec.Size)
    }

    fmt.Println("\n== SYMBOLS ==")
    if syms, err := f.Symbols(); err == nil {
        for _, s := range syms {
            fmt.Printf("%-30s 0x%x (%d bytes)\n", s.Name, s.Value, s.Size)
        }
    }
}
``

//Çalıştırma:

``bash
go run elfreader.go /bin/ls
``

Örnek çıktı:
/*
```
== ELF HEADER ==
Class: ELFCLASS64, Data: ELFDATA2LSB, Type: ET_EXEC, Machine: EM_X86_64

== SECTIONS ==
.text                 130432 bytes
.data                 2048 bytes
...

== SYMBOLS ==
main                  0x401080 (120 bytes)
puts                  0x0 (0 bytes)
...
```

---

✅ Özetle `debug/elf` paketi ile:

* ELF header bilgilerini alabilir,
* Section ve Segment bilgilerini inceleyebilir,
* Sembol tablolarını ve string tablolarını okuyabilir,
* Dinamik semboller ve not bilgilerini görüntüleyebilirsiniz.

---

İstersen senin için bir **mini `readelf` aracı** (Linux’taki `readelf` komutunun Go ile yapılmış hali) kodlayabilirim. İstiyor musun?
EVET
*/

/*
Harika 🚀 O zaman sana Go ile yazılmış **mini bir `readelf` aracı** yapayım.
Bu araç, Linux’taki `readelf -h`, `readelf -S`, `readelf -s` komutlarının basit bir versiyonu olacak.

---

# 📌 Go `readelf` Aracı
*/
``go
package main

import (
	"debug/elf"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Kullanım: %s <elf-dosyası>\n", os.Args[0])
	}

	file := os.Args[1]
	f, err := elf.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// === ELF HEADER ===
	fmt.Println("== ELF HEADER ==")
	fmt.Printf("Class:   %v\n", f.Class)
	fmt.Printf("Data:    %v\n", f.Data)
	fmt.Printf("Type:    %v\n", f.Type)
	fmt.Printf("Machine: %v\n", f.Machine)
	fmt.Printf("Version: %v\n", f.Version)
	fmt.Printf("Entry:   0x%x\n", f.Entry)

	// === SECTIONS ===
	fmt.Println("\n== SECTIONS ==")
	for i, sec := range f.Sections {
		fmt.Printf("[%2d] %-20s Addr=0x%x Size=%d bytes\n",
			i, sec.Name, sec.Addr, sec.Size)
	}

	// === PROGRAM HEADERS ===
	fmt.Println("\n== PROGRAM HEADERS ==")
	for i, prog := range f.Progs {
		fmt.Printf("[%2d] Type=%-10v Offset=0x%x Vaddr=0x%x Filesz=%d Memsz=%d Flags=%v\n",
			i, prog.Type, prog.Off, prog.Vaddr, prog.Filesz, prog.Memsz, prog.Flags)
	}

	// === SYMBOLS ===
	fmt.Println("\n== SYMBOLS ==")
	if syms, err := f.Symbols(); err == nil {
		for i, s := range syms {
			fmt.Printf("[%3d] %-30s Value=0x%x Size=%d Info=%d\n",
				i, s.Name, s.Value, s.Size, s.Info)
		}
	} else {
		fmt.Println("Sembol tablosu bulunamadı.")
	}

	// === DYNAMIC SYMBOLS ===
	fmt.Println("\n== DYNAMIC SYMBOLS ==")
	if dynSyms, err := f.DynamicSymbols(); err == nil {
		for i, s := range dynSyms {
			fmt.Printf("[%3d] %-30s Value=0x%x Size=%d Info=%d\n",
				i, s.Name, s.Value, s.Size, s.Info)
		}
	} else {
		fmt.Println("Dinamik semboller bulunamadı.")
	}
}
``
/*
---

# 📌 Kullanım

Kodu `readelf.go` olarak kaydet:
*/
``bash
go run readelf.go /bin/ls
``
/*
Örnek çıktı:

```
== ELF HEADER ==
Class:   ELFCLASS64
Data:    ELFDATA2LSB
Type:    ET_EXEC
Machine: EM_X86_64
Version: EV_CURRENT
Entry:   0x401080

== SECTIONS ==
[ 0]                         Addr=0x0 Size=0 bytes
[ 1] .interp                 Addr=0x400238 Size=28 bytes
[ 2] .note.gnu.build-id      Addr=0x400254 Size=36 bytes
[ 3] .text                   Addr=0x401000 Size=132432 bytes
...

== PROGRAM HEADERS ==
[ 0] Type=PT_PHDR     Offset=0x40 Vaddr=0x400040 Filesz=392 Memsz=392 Flags=R
[ 1] Type=PT_INTERP   Offset=0x238 Vaddr=0x400238 Filesz=28 Memsz=28 Flags=R
...

== SYMBOLS ==
[  0]                                Value=0x0 Size=0 Info=0
[  1] _start                         Value=0x401080 Size=0 Info=0
[  2] main                           Value=0x402000 Size=128 Info=0
...

== DYNAMIC SYMBOLS ==
[  0] libc.so.6                      Value=0x0 Size=0 Info=0
[  1] printf                         Value=0x0 Size=0 Info=0
...
```

---

✅ Bu araç ile bir ELF dosyasının:

* **Header bilgisini** (`readelf -h`)
* **Section listesini** (`readelf -S`)
* **Program header (segment) bilgilerini** (`readelf -l`)
* **Sembol tablolarını** (`readelf -s`, `readelf -Ws`)

inceleyebilirsin.

---

İstiyorsan buna **`readelf` gibi opsiyon ekleyelim** (örn: `-h` sadece header göstersin, `-s` sadece sembolleri göstersin). Bunu da ister misin?
*/