/*
Go’daki **`debug/gosym`** paketini baştan sona, örneklerle anlatayım.

---

# 📌 `debug/gosym` Paketi Nedir?

Go’nun **`debug/gosym`** paketi, derlenmiş Go programlarının **sembol tablolarını** ve **debug bilgilerini** okumak için kullanılır.
Özellikle **eski formatlardaki Go derlenmiş ikili dosyalar** için (Go 1.10 öncesi `6l`, `8l`, `5l` linkleyicileri) tasarlanmıştır.
Günümüzde modern Go binary’lerinde daha çok **`debug/dwarf`** ve **`debug/elf` / `debug/macho` / `debug/pe`** kullanılsa da, `gosym` hâlâ işlevseldir çünkü:

* **Fonksiyon isimlerini** çözümlemeye (PC → fonksiyon adı eşleştirmesi) yarar.
* **Sembolleri** ve **satır tablolarını** okumaya yarar.
* Profiling, debugging, stack trace çözümleme gibi işlerde kullanılabilir.

---

# 📌 Temel Kavramlar

`gosym` üç ana yapı etrafında çalışır:

1. **`Sym`** → Bir sembol (ör. fonksiyon veya değişken)
2. **`LineTable`** → Program counter (PC) adresi ile kaynak dosya/satır ilişkisini tutar
3. **`Table`** → Sembol tablosu (`Sym` listesi) + `LineTable`

---

# 📌 Önemli Tipler ve Fonksiyonlar

### 1. `Sym`

Bir Go sembolünü temsil eder.
*/

``go
type Sym struct {
    Value  uint64 // Bellekteki adres
    Type   byte   // Sembol tipi (T = text, D = data)
    Name   string // Fonksiyon / değişken adı
    GoType uint64 // Go tipi (varsa)
}
``
/*
---

### 2. `LineTable`

Adres ↔ Dosya/Satır eşlemesini tutar.
*/

``go
func (lt *LineTable) PCToLine(pc uint64) (file string, line int, fn *Sym)
``
/*
* `pc` adresinden dosya adı, satır numarası ve fonksiyon sembolünü verir.

---

### 3. `Table`

Sembollerin ve `LineTable`’ın birleşimidir.
*/

``go
type Table struct {
    Syms  []Sym
    Funcs []Func
    LT    *LineTable
}
``

//`Table` oluşturma:

``go
tab, err := gosym.NewTable(symtabData, lineTable)
``
/*
---

# 📌 Örnek 1 – Basit Sembol Tablosu Okuma

Aşağıdaki örnekte, derlenmiş bir Go binary’sinden sembolleri okuyoruz:
*/

``go
package main

import (
    "debug/gosym"
    "debug/elf"
    "fmt"
    "log"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatalf("Kullanım: %s <binary>\n", os.Args[0])
    }

    bin := os.Args[1]
    f, err := elf.Open(bin)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    // .gosymtab ve .gopclntab sectionlarını bul
    symtabSec := f.Section(".gosymtab")
    pclnSec := f.Section(".gopclntab")
    if symtabSec == nil || pclnSec == nil {
        log.Fatal("Bu binary .gosymtab veya .gopclntab içermiyor.")
    }

    symtabData, _ := symtabSec.Data()
    pclnData, _ := pclnSec.Data()

    // LineTable oluştur
    lineTable := gosym.NewLineTable(pclnData, f.Section(".text").Addr)

    // Sembol tablosu oluştur
    table, err := gosym.NewTable(symtabData, lineTable)
    if err != nil {
        log.Fatal(err)
    }

    // Fonksiyonları yazdır
    for _, fn := range table.Funcs {
        fmt.Printf("Fonksiyon: %-30s Start=0x%x End=0x%x\n",
            fn.Sym.Name, fn.Sym.Value, fn.End)
    }
}
``

//📌 Çalıştır:

``bash
go build -o myprog main.go
go run gosym_reader.go ./myprog
``
/*
Örnek çıktı:

```
Fonksiyon: main.main                   Start=0x45e000 End=0x45e120
Fonksiyon: main.helper                 Start=0x45e130 End=0x45e190
Fonksiyon: runtime.main                Start=0x401000 End=0x401320
...
```

---

# 📌 Örnek 2 – Adresten Satır Bilgisi Bulma
*/

``go
pc := table.LookupFunc("main.main").Sym.Value
file, line, fn := table.LT.PCToLine(pc)
fmt.Printf("Adres 0x%x -> %s:%d (fonksiyon %s)\n", pc, file, line, fn.Name)
``
/*
Çıktı:

```
Adres 0x45e000 -> /home/user/main.go:10 (fonksiyon main.main)
```

---

# 📌 Örnek 3 – Sembolleri Tek Tek Listeleme
*/

``go
for _, sym := range table.Syms {
    fmt.Printf("Sembol: %-30s Addr=0x%x Type=%c\n", sym.Name, sym.Value, sym.Type)
}
``
/*
Örnek çıktı:

```
Sembol: main.main                     Addr=0x45e000 Type=T
Sembol: main.helper                   Addr=0x45e130 Type=T
Sembol: runtime.printstring           Addr=0x401200 Type=T
Sembol: fmt.Println                    Addr=0x0 Type=U
```

---

# 📌 `gosym` Kullanım Alanları

* Go binary’sinden **fonksiyon isimlerini çözmek**
* Profil/trace araçları (ör. `pprof`) yazmak
* Stack trace’lerde adresleri **kaynak dosya + satıra dönüştürmek**
* `dlv` (Delve debugger) gibi debug araçlarında

---

✅ **Özetle**:

* `debug/gosym`, Go binary’lerindeki sembolleri ve satır tablolarını okumak için kullanılan eski ama yararlı bir pakettir.
* `.gosymtab` ve `.gopclntab` section’larını okuyarak sembollere ulaşır.
* Fonksiyon isimlerini çözme, stack trace adreslerini dosya/satır ile eşleştirme gibi işlerde kullanılır.

---

İstersen ben sana `gosym` ile yazılmış **mini bir stacktrace çözücü** (adres → dosya.go\:line + fonksiyon adı) yapabilirim. Bunu ister misin?
EVET

Süper 🚀 O zaman sana **`gosym` ile stack trace çözücü** bir mini araç yazayım.
Bu araç, elimizdeki bir Go binary’si (`.exe`, `./myprog`) için adresleri alacak ve o adreslerin hangi **dosya.go\:satır** ve **fonksiyon**a denk geldiğini gösterecek.

---

# 📌 Mini Stack Trace Çözücü (`addr2line` benzeri)
*/

``go
package main

import (
	"bufio"
	"debug/elf"
	"debug/gosym"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Kullanım: %s <binary>\n", os.Args[0])
	}

	// ELF dosyasını aç
	file := os.Args[1]
	f, err := elf.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Gerekli sectionları bul
	symtabSec := f.Section(".gosymtab")
	pclnSec := f.Section(".gopclntab")
	textSec := f.Section(".text")

	if symtabSec == nil || pclnSec == nil || textSec == nil {
		log.Fatal("ELF dosyası .gosymtab / .gopclntab / .text sectionlarını içermiyor.")
	}

	// Section verilerini al
	symtabData, _ := symtabSec.Data()
	pclnData, _ := pclnSec.Data()

	// LineTable ve Table oluştur
	lineTable := gosym.NewLineTable(pclnData, textSec.Addr)
	table, err := gosym.NewTable(symtabData, lineTable)
	if err != nil {
		log.Fatal(err)
	}

	// Kullanıcıdan adres al
	fmt.Println("Adresleri hex formatında gir (ör: 0x45e000). Çıkış için boş satır bırak.")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if line == "" {
			break
		}

		// Hex string → uint64 çevir
		addr, err := strconv.ParseUint(line, 0, 64)
		if err != nil {
			fmt.Println("Geçersiz adres:", err)
			continue
		}

		// Adresi çöz
		file, lno, fn := table.LT.PCToLine(addr)
		if fn != nil {
			fmt.Printf("0x%x => %s:%d (fonksiyon: %s)\n", addr, file, lno, fn.Name)
		} else {
			fmt.Printf("0x%x => çözüm bulunamadı\n", addr)
		}
	}
}
``
/*
---

# 📌 Kullanım

1. Örnek bir Go programı derle:
*/

``go
package main

import "fmt"

func helper() {
	fmt.Println("yardımcı fonksiyon")
}

func main() {
	helper()
	fmt.Println("merhaba dünya")
}
``

//Kaydet: `program.go`

``bash
go build -o myprog program.go
``

//2. Çözücüyü çalıştır:

``bash
go run stacktrace_resolver.go ./myprog
``
/*
3. Adres gir (örneğin `objdump -d myprog` veya `nm myprog` ile bulduğun bir adres):

```
Adresleri hex formatında gir (ör: 0x45e000). Çıkış için boş satır bırak.
> 0x45e000
0x45e000 => /home/user/program.go:4 (fonksiyon: main.helper)
> 0x45e050
0x45e050 => /home/user/program.go:8 (fonksiyon: main.main)
>
```

---

✅ Böylece elimizde küçük bir **Go `addr2line` aracı** oldu.

* ELF dosyasını açıyor,
* `.gosymtab` + `.gopclntab` sectionlarını kullanıyor,
* Adresleri dosya, satır ve fonksiyon adına çeviriyor.

---

👉 İstersen ben bu aracı daha da geliştirip `panic` çıktısını (stack trace) **otomatik parse edip çözümleyen** bir versiyon da yapabilirim. Bunu ister misin?
*/