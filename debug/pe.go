/*
Go’daki **`debug/pe`** paketini baştan sona örneklerle anlatayım.

---

# 📌 `debug/pe` Paketi Nedir?

Go’nun **`debug/pe`** paketi, **Windows PE (Portable Executable)** dosyalarını okumak ve incelemek için kullanılır.
PE formatı, Windows’taki çalıştırılabilir dosyaların (`.exe`), dinamik kütüphanelerin (`.dll`), sürücülerin (`.sys`) temel formatıdır.

Go’daki `debug/pe` paketi ile:

* **PE Header** bilgilerini okuyabilir,
* **Section tablolarını** listeleyebilir,
* **Import / Export tablolarını** çıkarabilir,
* **Sembol tablosunu** (eski COFF tarzı) çözümleyebilirsiniz.

Bu paket, ELF (`debug/elf`) ve Mach-O (`debug/macho`) paketlerine benzer şekilde, dosya analiz araçları geliştirmek için kullanılır.

---

# 📌 Temel Yapılar

`debug/pe` paketinde öne çıkan yapılar şunlardır:

* **`File`** → PE dosyasını temsil eder
* **`FileHeader`** → COFF (Common Object File Format) header bilgileri
* **`OptionalHeader32/64`** → PE header (32-bit / 64-bit farkı)
* **`Section`** → Section tabloları (`.text`, `.data`, `.rdata`, `.rsrc` vb.)
* **`ImportDirectory`** → Import edilen DLL’ler ve fonksiyonlar
* **`ExportDirectory`** → Export edilen semboller
* **`Symbol`** → Eski COFF sembolleri

---

# 📌 PE Dosyasını Açmak
*/
``go
package main

import (
    "debug/pe"
    "fmt"
    "log"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatalf("Kullanım: %s <pe-dosyası>\n", os.Args[0])
    }

    file := os.Args[1]
    f, err := pe.Open(file)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    fmt.Printf("Machine: 0x%x\n", f.FileHeader.Machine)
    fmt.Printf("Number of Sections: %d\n", f.FileHeader.NumberOfSections)
    fmt.Printf("TimeDateStamp: %d\n", f.FileHeader.TimeDateStamp)
}
``
/*
📌 Örnek çıktı (`notepad.exe` için):

```
Machine: 0x14c
Number of Sections: 6
TimeDateStamp: 1658791234
```

---

# 📌 Section’ları Listeleme
*/
``go
for _, sec := range f.Sections {
    fmt.Printf("Section: %-8s VirtAddr=0x%x Size=%d RawDataSize=%d\n",
        sec.Name, sec.VirtualAddress, sec.VirtualSize, sec.Size)
}
``
/*
Örnek çıktı:

```
Section: .text    VirtAddr=0x1000 Size=52345 RawDataSize=53248
Section: .rdata   VirtAddr=0xe000 Size=10432 RawDataSize=11264
Section: .data    VirtAddr=0x11000 Size=2048 RawDataSize=1024
Section: .rsrc    VirtAddr=0x12000 Size=8192 RawDataSize=8192
```

---

# 📌 Optional Header (32-bit veya 64-bit)
*/
``go
switch opt := f.OptionalHeader.(type) {
case *pe.OptionalHeader32:
    fmt.Printf("32-bit Entry Point: 0x%x\n", opt.AddressOfEntryPoint)
    fmt.Printf("ImageBase: 0x%x\n", opt.ImageBase)
case *pe.OptionalHeader64:
    fmt.Printf("64-bit Entry Point: 0x%x\n", opt.AddressOfEntryPoint)
    fmt.Printf("ImageBase: 0x%x\n", opt.ImageBase)
}
``
/*
---

# 📌 Import Tabloları (DLL Bağımlılıkları)
*/

``go
imps, err := f.ImportedLibraries()
if err == nil {
    fmt.Println("Kullanılan DLL'ler:")
    for _, lib := range imps {
        fmt.Println(" -", lib)
    }
}

funcs, err := f.ImportedSymbols()
if err == nil {
    fmt.Println("\nImport edilen fonksiyonlar:")
    for _, fn := range funcs {
        fmt.Println(" -", fn)
    }
}
``
/*
Örnek çıktı:

```
Kullanılan DLL'ler:
 - KERNEL32.dll
 - USER32.dll
 - GDI32.dll

Import edilen fonksiyonlar:
 - KERNEL32.dll!ExitProcess
 - KERNEL32.dll!CreateFileW
 - USER32.dll!MessageBoxW
```
/*
---

# 📌 Export Tabloları (DLL’den Dışa Açılan Fonksiyonlar)
*/
``go
if f.Export != nil {
    fmt.Println("Export edilen semboller:")
    for _, exp := range f.Export.Functions {
        fmt.Printf("Ordinal: %d, Address: 0x%x\n", exp.Ordinal, exp.Address)
    }
}
``
/*
Örnek çıktı (`user32.dll` için):

```
Export edilen semboller:
Ordinal: 100, Address: 0x12345
Ordinal: 101, Address: 0x12400
```

---

# 📌 Sembol Tablosu (Eski COFF Tarzı)
*/
``go
if len(f.Symbols) > 0 {
    fmt.Println("COFF sembolleri:")
    for _, sym := range f.Symbols {
        fmt.Printf("Name: %s, Value: %d, Section: %d\n", sym.Name, sym.Value, sym.SectionNumber)
    }
}
``
/*
---

# 📌 Tam Örnek – PE İnceleme Aracı
*/
``go
package main

import (
    "debug/pe"
    "fmt"
    "log"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatalf("Kullanım: %s <pe-dosyası>\n", os.Args[0])
    }

    file := os.Args[1]
    f, err := pe.Open(file)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    fmt.Println("== PE HEADER ==")
    fmt.Printf("Machine: 0x%x, Sections: %d, Timestamp: %d\n",
        f.FileHeader.Machine, f.FileHeader.NumberOfSections, f.FileHeader.TimeDateStamp)

    fmt.Println("\n== OPTIONAL HEADER ==")
    switch opt := f.OptionalHeader.(type) {
    case *pe.OptionalHeader32:
        fmt.Printf("32-bit Entry=0x%x ImageBase=0x%x\n", opt.AddressOfEntryPoint, opt.ImageBase)
    case *pe.OptionalHeader64:
        fmt.Printf("64-bit Entry=0x%x ImageBase=0x%x\n", opt.AddressOfEntryPoint, opt.ImageBase)
    }

    fmt.Println("\n== SECTIONS ==")
    for _, sec := range f.Sections {
        fmt.Printf("%-8s VirtAddr=0x%x Size=%d RawSize=%d\n",
            sec.Name, sec.VirtualAddress, sec.VirtualSize, sec.Size)
    }

    fmt.Println("\n== IMPORTED LIBRARIES ==")
    if libs, err := f.ImportedLibraries(); err == nil {
        for _, lib := range libs {
            fmt.Println(" -", lib)
        }
    }

    fmt.Println("\n== IMPORTED SYMBOLS ==")
    if funcs, err := f.ImportedSymbols(); err == nil {
        for _, fn := range funcs {
            fmt.Println(" -", fn)
        }
    }

    fmt.Println("\n== EXPORTED SYMBOLS ==")
    if f.Export != nil {
        for _, exp := range f.Export.Functions {
            fmt.Printf("Ordinal: %d, Address: 0x%x\n", exp.Ordinal, exp.Address)
        }
    }
}
``

//📌 Kullanım:

``bash
go run pereader.go C:\Windows\System32\notepad.exe
``
/*
---

# 📌 Özet

✅ `debug/pe` paketi ile:

* PE header bilgilerini alabilirsiniz
* Section tablolarını inceleyebilirsiniz
* Import edilen DLL ve fonksiyonları çıkarabilirsiniz
* Export edilen sembolleri görebilirsiniz
* Eski COFF sembollerini listeleyebilirsiniz

Böylece Windows’taki **`dumpbin`**, **`objdump`** veya **PE Explorer** gibi araçların temel işlevlerini Go ile yazabilirsiniz.

---

👉 İstersen sana bunun üstüne **mini bir PE Explorer** (örneğin `dumpbin` benzeri) yazabilirim. İster misin?
EVET
*/
/*
Harika 🎯 O zaman sana Go ile **mini bir `PE Explorer`** yazalım.
Bu araç, Windows’ta **`dumpbin`** veya **PE Explorer** gibi çalışacak:

* PE header bilgilerini gösterecek,
* Section’ları listeleyecek,
* Import edilen DLL’leri ve fonksiyonları gösterecek,
* Export edilen fonksiyonları listeleyecek.

---

# 📌 `pe_explorer.go`
*/
``go
package main

import (
    "debug/pe"
    "fmt"
    "log"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatalf("Kullanım: %s <exe/dll dosyası>\n", os.Args[0])
    }

    path := os.Args[1]
    f, err := pe.Open(path)
    if err != nil {
        log.Fatalf("Dosya açılamadı: %v", err)
    }
    defer f.Close()

    fmt.Println("===== PE HEADER =====")
    fmt.Printf("Machine: 0x%x\n", f.FileHeader.Machine)
    fmt.Printf("Sections: %d\n", f.FileHeader.NumberOfSections)
    fmt.Printf("TimeDateStamp: %d\n", f.FileHeader.TimeDateStamp)
    fmt.Printf("Characteristics: 0x%x\n", f.FileHeader.Characteristics)

    fmt.Println("\n===== OPTIONAL HEADER =====")
    switch oh := f.OptionalHeader.(type) {
    case *pe.OptionalHeader32:
        fmt.Printf("32-bit EntryPoint: 0x%x\n", oh.AddressOfEntryPoint)
        fmt.Printf("ImageBase: 0x%x\n", oh.ImageBase)
        fmt.Printf("Subsystem: %d\n", oh.Subsystem)
    case *pe.OptionalHeader64:
        fmt.Printf("64-bit EntryPoint: 0x%x\n", oh.AddressOfEntryPoint)
        fmt.Printf("ImageBase: 0x%x\n", oh.ImageBase)
        fmt.Printf("Subsystem: %d\n", oh.Subsystem)
    }

    fmt.Println("\n===== SECTIONS =====")
    for _, sec := range f.Sections {
        fmt.Printf("Name: %-8s VirtAddr=0x%x Size=%d RawSize=%d\n",
            sec.Name, sec.VirtualAddress, sec.VirtualSize, sec.Size)
    }

    fmt.Println("\n===== IMPORTED LIBRARIES =====")
    if libs, err := f.ImportedLibraries(); err == nil {
        for _, lib := range libs {
            fmt.Println(" -", lib)
        }
    } else {
        fmt.Println("Import tablosu yok.")
    }

    fmt.Println("\n===== IMPORTED SYMBOLS =====")
    if funcs, err := f.ImportedSymbols(); err == nil {
        for _, fn := range funcs {
            fmt.Println(" -", fn)
        }
    } else {
        fmt.Println("Import edilen fonksiyon yok.")
    }

    fmt.Println("\n===== EXPORTED SYMBOLS =====")
    if f.Export != nil {
        for _, exp := range f.Export.Functions {
            fmt.Printf("Ordinal: %d Address: 0x%x\n", exp.Ordinal, exp.Address)
        }
    } else {
        fmt.Println("Export tablosu yok.")
    }
}
``
/*
---

# 📌 Kullanım

Örneğin Windows’ta `notepad.exe` dosyasını inceleyelim:
*/

``powershell
go run pe_explorer.go C:\Windows\System32\notepad.exe
``
/*
Örnek çıktı (kısaltılmış):

``
===== PE HEADER =====
Machine: 0x14c
Sections: 6
TimeDateStamp: 1661234567
Characteristics: 0x102

===== OPTIONAL HEADER =====
32-bit EntryPoint: 0x1410
ImageBase: 0x400000
Subsystem: 2

===== SECTIONS =====
Name: .text    VirtAddr=0x1000 Size=52345 RawSize=53248
Name: .rdata   VirtAddr=0xe000 Size=10432 RawSize=11264
Name: .data    VirtAddr=0x11000 Size=2048 RawSize=1024
Name: .rsrc    VirtAddr=0x12000 Size=8192 RawSize=8192

===== IMPORTED LIBRARIES =====
 - KERNEL32.dll
 - USER32.dll
 - GDI32.dll

===== IMPORTED SYMBOLS =====
 - KERNEL32.dll!ExitProcess
 - USER32.dll!MessageBoxW

===== EXPORTED SYMBOLS =====
Export tablosu yok.
```
*/
/*
---

👉 Bu kodu biraz daha geliştirip,

* **JSON çıktısı** verecek,
* veya **GUI (ör. Fyne, Gio, Ebiten)** ekleyip görselleştirecek şekilde yapabiliriz.

İstiyor musun, sana bu PE Explorer’ı **JSON çıktılı bir araç** ya da **GUI tabanlı bir program** haline getireyim?
EVET
*/

/*
Süper 🎯 O zaman biz yazdığımız **`PE Explorer`**’ı biraz daha **profesyonel** hale getirelim:

* Çıktıyı **JSON formatında** verecek,
* İstersek çıktıyı ekrana yazacak, istersek dosyaya kaydedecek,
* Daha sonra bu JSON’u başka bir uygulama (ör. web arayüzü, GUI) kolayca kullanabilecek.

---

# 📌 `pe_explorer_json.go`
*/
``go
package main

import (
	"debug/pe"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type PEInfo struct {
	FileHeader struct {
		Machine          uint16 `json:"machine"`
		NumberOfSections uint16 `json:"number_of_sections"`
		TimeDateStamp    uint32 `json:"time_date_stamp"`
		Characteristics  uint16 `json:"characteristics"`
	} `json:"file_header"`

	OptionalHeader interface{} `json:"optional_header"`

	Sections []SectionInfo `json:"sections"`

	ImportedLibraries []string `json:"imported_libraries"`
	ImportedSymbols   []string `json:"imported_symbols"`
	ExportedSymbols   []ExportInfo `json:"exported_symbols"`
}

type SectionInfo struct {
	Name      string `json:"name"`
	VirtAddr  uint32 `json:"virtual_address"`
	VirtSize  uint32 `json:"virtual_size"`
	RawSize   uint32 `json:"raw_size"`
}

type ExportInfo struct {
	Ordinal uint16 `json:"ordinal"`
	Address uint32 `json:"address"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Kullanım: %s <exe/dll dosyası> [json_output_file]\n", os.Args[0])
	}

	path := os.Args[1]
	f, err := pe.Open(path)
	if err != nil {
		log.Fatalf("Dosya açılamadı: %v", err)
	}
	defer f.Close()

	info := PEInfo{}

	// File Header
	info.FileHeader.Machine = f.FileHeader.Machine
	info.FileHeader.NumberOfSections = f.FileHeader.NumberOfSections
	info.FileHeader.TimeDateStamp = f.FileHeader.TimeDateStamp
	info.FileHeader.Characteristics = f.FileHeader.Characteristics

	// Optional Header
	switch oh := f.OptionalHeader.(type) {
	case *pe.OptionalHeader32:
		info.OptionalHeader = map[string]interface{}{
			"entry_point": fmt.Sprintf("0x%x", oh.AddressOfEntryPoint),
			"image_base":  fmt.Sprintf("0x%x", oh.ImageBase),
			"subsystem":   oh.Subsystem,
		}
	case *pe.OptionalHeader64:
		info.OptionalHeader = map[string]interface{}{
			"entry_point": fmt.Sprintf("0x%x", oh.AddressOfEntryPoint),
			"image_base":  fmt.Sprintf("0x%x", oh.ImageBase),
			"subsystem":   oh.Subsystem,
		}
	}

	// Sections
	for _, sec := range f.Sections {
		info.Sections = append(info.Sections, SectionInfo{
			Name:     sec.Name,
			VirtAddr: sec.VirtualAddress,
			VirtSize: sec.VirtualSize,
			RawSize:  sec.Size,
		})
	}

	// Imports
	if libs, err := f.ImportedLibraries(); err == nil {
		info.ImportedLibraries = libs
	}
	if funcs, err := f.ImportedSymbols(); err == nil {
		info.ImportedSymbols = funcs
	}

	// Exports
	if f.Export != nil {
		for _, exp := range f.Export.Functions {
			info.ExportedSymbols = append(info.ExportedSymbols, ExportInfo{
				Ordinal: exp.Ordinal,
				Address: exp.Address,
			})
		}
	}

	// JSON encode
	jsonData, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		log.Fatal("JSON kodlama hatası:", err)
	}

	// Ekrana yaz
	fmt.Println(string(jsonData))

	// Eğer dosya ismi verildiyse dosyaya kaydet
	if len(os.Args) >= 3 {
		err := os.WriteFile(os.Args[2], jsonData, 0644)
		if err != nil {
			log.Fatal("JSON dosyaya yazılamadı:", err)
		}
		fmt.Println("JSON çıktısı kaydedildi:", os.Args[2])
	}
}
``
/*
---

# 📌 Kullanım

1. `notepad.exe` veya `calc.exe` gibi bir Windows PE dosyası seç:
*/
``powershell
go run pe_explorer_json.go C:\Windows\System32\notepad.exe
``

2. JSON çıktısı ekranda:

``json
{
  "file_header": {
    "machine": 332,
    "number_of_sections": 6,
    "time_date_stamp": 1661234567,
    "characteristics": 258
  },
  "optional_header": {
    "entry_point": "0x1410",
    "image_base": "0x400000",
    "subsystem": 2
  },
  "sections": [
    {
      "name": ".text",
      "virtual_address": 4096,
      "virtual_size": 52345,
      "raw_size": 53248
    },
    {
      "name": ".rdata",
      "virtual_address": 57344,
      "virtual_size": 10432,
      "raw_size": 11264
    }
  ],
  "imported_libraries": ["KERNEL32.dll", "USER32.dll"],
  "imported_symbols": [
    "KERNEL32.dll!ExitProcess",
    "USER32.dll!MessageBoxW"
  ],
  "exported_symbols": []
}
``

//3. Eğer JSON dosyasına kaydetmek istersen:

``powershell
go run pe_explorer_json.go C:\Windows\System32\notepad.exe notepad.json
``
/*
---

✅ Artık elimizde bir **JSON tabanlı PE Explorer** var.
👉 Bunu web servis (REST API) yapabiliriz ya da GUI’ye (ör. Fyne / Gio) aktarabiliriz.

İstersen ben sana bu JSON PE Explorer’ı bir **web arayüzü (React + Go backend)** ile entegre edip tarayıcıda dosya seçip analiz edebileceğin şekilde de hazırlayabilirim. Bunu ister misin?
*/