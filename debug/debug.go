/*
Golang’daki **`debug`** paketi aslında **standart kütüphanedeki bir üst seviye dizin**’dir.
Yani `debug` tek başına kullanılmaz, fakat içinde farklı alt paketler bulunur. Bu alt paketler genellikle **binary dosyaları (ELF, Mach-O, PE)**, **profil dosyaları (dwarf, plan9obj)** ve hata ayıklama (debugging) ile ilgili şeyleri analiz etmek için kullanılır.

`debug` altındaki paketler şunlardır:

* **`debug/dwarf`** → DWARF debugging bilgilerini okumak için.
* **`debug/elf`** → ELF formatlı (Linux/Unix) dosyaları incelemek için.
* **`debug/gosym`** → Go derleyicisinin ürettiği sembol tablolarını çözmek için.
* **`debug/macho`** → Mach-O formatlı (macOS, iOS) dosyaları incelemek için.
* **`debug/pe`** → PE formatlı (Windows Portable Executable) dosyaları incelemek için.
* **`debug/plan9obj`** → Plan 9 işletim sistemi için kullanılan object file formatını incelemek için.

---

Şimdi her birini tek tek açıklayayım ve örneklerle göstereyim:

---

## 1. `debug/dwarf`

DWARF, derlenmiş dosyalarda hata ayıklama bilgilerini (değişkenler, tipler, fonksiyonlar, satır numaraları vb.) tutan bir formattır.

**Örnek: ELF dosyasındaki DWARF bilgilerini okuma**
*/

``go
package main

import (
	"debug/dwarf"
	"debug/elf"
	"fmt"
	"os"
)

func main() {
	// Bir ELF dosyasını açıyoruz (örneğin /bin/ls)
	file, err := elf.Open("/bin/ls")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	defer file.Close()

	// DWARF bilgilerini al
	data, err := file.DWARF()
	if err != nil {
		fmt.Println("DWARF bulunamadı:", err)
		return
	}

	// Tip bilgilerini gezelim
	r := data.Reader()
	for {
		e, err := r.Next()
		if err != nil {
			fmt.Println("Hata:", err)
			break
		}
		if e == nil {
			break
		}
		fmt.Println("Entry:", e.Tag)
	}
}
``
/*
---

## 2. `debug/elf`

ELF (Executable and Linkable Format), Linux/Unix dünyasında kullanılan dosya formatıdır.

**Örnek: ELF başlık bilgilerini okuma**
*/
``go
package main

import (
	"debug/elf"
	"fmt"
)

func main() {
	file, err := elf.Open("/bin/ls")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("ELF Class:", file.Class)
	fmt.Println("ELF Type:", file.Type)
	fmt.Println("Machine:", file.Machine)

	// Bölümleri listele
	for _, section := range file.Sections {
		fmt.Println("Section:", section.Name, "Size:", section.Size)
	}
}
``
/*
---

## 3. `debug/gosym`

Go programlarının sembol tablolarını çözümlemek için kullanılır. Özellikle **`runtime/pprof`** ve **stack trace analizlerinde** iş görür.

**Örnek: Sembolleri çözme**
*/
``go
package main

import (
	"debug/gosym"
	"fmt"
)

func main() {
	// Örnek sembol tablo (normalde binary'den okunur)
	table := []byte("\x00\x00...") // Gerçek örnek için binary okunmalı
	lineTab := &gosym.LineTable{Data: table, SymTab: 0}

	tableData := &gosym.Table{
		LineTable: lineTab,
	}

	fmt.Println("Semboller:", tableData.Syms)
}
``
/*
*(Gerçek kullanımda ELF veya PE dosyasıyla birlikte kullanılır.)*

---

## 4. `debug/macho`

Mach-O (Mach Object File), macOS ve iOS’ta kullanılan dosya formatıdır.

**Örnek: Mach-O dosyası okuma**
*/

``go
package main

import (
	"debug/macho"
	"fmt"
)

func main() {
	file, err := macho.Open("/bin/ls")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("CPU:", file.Cpu)
	fmt.Println("Type:", file.Type)

	for _, load := range file.Loads {
		fmt.Printf("Load: %T\n", load)
	}
}
``
/*
---

## 5. `debug/pe`

PE (Portable Executable), Windows çalıştırılabilir dosya formatıdır.

**Örnek: PE dosyasını inceleme**
*/

``go
package main

import (
	"debug/pe"
	"fmt"
)

func main() {
	file, err := pe.Open("C:\\Windows\\System32\\notepad.exe")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("Machine:", file.FileHeader.Machine)
	fmt.Println("Number of Sections:", file.FileHeader.NumberOfSections)

	for _, sec := range file.Sections {
		fmt.Println("Section:", sec.Name, "Size:", sec.Size)
	}
}
``


/*
---

## 6. `debug/plan9obj`

Plan 9 işletim sistemi için kullanılan **Plan9 object file** formatını analiz eder.

**Örnek: Plan9 object dosyasını açma**
*/
``go
package main

import (
	"debug/plan9obj"
	"fmt"
)

func main() {
	file, err := plan9obj.Open("program.o")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("Magic:", file.Magic)
	for _, s := range file.Sections {
		fmt.Println("Section:", s.Name, "Size:", s.Size)
	}
}
``
/*
---

## Özet

👉 `debug` paketi doğrudan kullanılmaz, **altındaki alt paketler binary/debug bilgilerini analiz etmek için** vardır.

* **ELF, PE, Mach-O, Plan9** → Farklı işletim sistemlerinin dosya formatları.
* **DWARF** → Debugging bilgileri.
* **GoSym** → Go’nun sembol tabloları.

---

İstersen ben sana bu paketleri tek tek **bir örnek proje** şeklinde hazırlayıp, aynı anda ELF, PE, Mach-O dosyalarını okuyan bir Go programı yapabilirim. İster misin?
EVET
*/

/*
Harika 👍 O zaman sana **tek bir Go projesi** hazırlayayım.
Bu proje:

* Linux’ta **ELF**,
* Windows’ta **PE**,
* macOS’te **Mach-O**,
* varsa **DWARF** ve **Go sembolleri** de

dosya hakkında bilgi çıkaracak.

---

## 📂 Proje Yapısı

```
debug-inspector/
 ├── main.go
```

---
*/
//## 📌 `main.go`

``go
package main

import (
	"debug/dwarf"
	"debug/elf"
	"debug/gosym"
	"debug/macho"
	"debug/pe"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Kullanım: go run main.go <dosya_yolu>")
		return
	}

	path := os.Args[1]
	ext := filepath.Ext(path)

	fmt.Println("Dosya:", path)
	fmt.Println("İşletim sistemi:", runtime.GOOS)

	// ELF (Linux/Unix)
	if ext == "" || ext == ".elf" {
		if f, err := elf.Open(path); err == nil {
			defer f.Close()
			fmt.Println("==== ELF Bilgileri ====")
			fmt.Println("Class:", f.Class, "Machine:", f.Machine, "Type:", f.Type)
			for _, sec := range f.Sections {
				fmt.Println("Section:", sec.Name, "Size:", sec.Size)
			}

			// DWARF varsa oku
			if d, err := f.DWARF(); err == nil {
				printDWARF(d)
			}
		}
	}

	// PE (Windows)
	if ext == ".exe" || ext == ".dll" {
		if f, err := pe.Open(path); err == nil {
			defer f.Close()
			fmt.Println("==== PE Bilgileri ====")
			fmt.Println("Machine:", f.FileHeader.Machine)
			for _, sec := range f.Sections {
				fmt.Println("Section:", sec.Name, "Size:", sec.Size)
			}
		}
	}

	// Mach-O (macOS)
	if ext == "" || ext == ".macho" {
		if f, err := macho.Open(path); err == nil {
			defer f.Close()
			fmt.Println("==== Mach-O Bilgileri ====")
			fmt.Println("CPU:", f.Cpu, "Type:", f.Type)
			for _, sec := range f.Sections {
				fmt.Println("Section:", sec.Name, "Size:", sec.Size)
			}
		}
	}

	// Go Sembolleri (opsiyonel)
	if data, err := os.ReadFile(path); err == nil {
		lt := &gosym.LineTable{Data: data, SymTab: 0}
		table := &gosym.Table{LineTable: lt}
		if len(table.Syms) > 0 {
			fmt.Println("==== Go Sembolleri ====")
			for _, sym := range table.Syms {
				fmt.Println("Symbol:", sym.Name)
			}
		}
	}
}

func printDWARF(d *dwarf.Data) {
	fmt.Println("==== DWARF Bilgileri ====")
	r := d.Reader()
	for {
		e, err := r.Next()
		if err != nil || e == nil {
			break
		}
		fmt.Println("Tag:", e.Tag)
	}
}
``
/*
---

## 🚀 Kullanım
*/

//Linux’ta:

``bash
go run main.go /bin/ls
``

//Windows’ta:

``powershell
go run main.go C:\Windows\System32\notepad.exe
``

//macOS’te:

``bash
go run main.go /bin/ls
``
/*
---

## 📌 Özet

Bu proje:

* **ELF / PE / Mach-O** dosya başlıklarını ve bölümlerini listeler.
* **DWARF** bilgisi varsa okur.
* **Go sembollerini** çözmeyi dener.

---

👉 İstersen ben buna **Dockerfile** de ekleyeyim, böylece Linux üzerinde build edip çalıştırabilirsin. Eklememi ister misin?
EVET
*/

/*
Süper 👍 O zaman bu projeyi **Docker ile paketleyelim**.
Docker sayesinde hem senin makinede Go kurulu olmasına gerek kalmaz, hem de farklı işletim sistemleri için kolayca build alabilirsin.

---

## 📂 Proje Yapısı

```
debug-inspector/
 ├── main.go
 ├── Dockerfile
```

---

## 📌 `Dockerfile`
*/
``dockerfile
# 1. Go imajını kullanıyoruz
FROM golang:1.22-alpine AS builder

# 2. Çalışma dizini
WORKDIR /app

# 3. Kaynak kodu kopyala
COPY . .

# 4. Go mod init (eğer go.mod yoksa)
RUN go mod init debug-inspector || true

# 5. Build
RUN go build -o inspector main.go

# 6. Daha küçük final imaj
FROM alpine:latest

WORKDIR /root/

# Builder'dan binary kopyala
COPY --from=builder /app/inspector .

# Varsayılan komut
ENTRYPOINT ["./inspector"]
``
/*
---

## 🚀 Build ve Çalıştırma

1. **Docker imajını build et**
*/
``bash
docker build -t debug-inspector .
``

//2. **Bir binary dosyayı analiz et** (örneğin `/bin/ls`)

``bash
docker run --rm -v /bin/ls:/bin/ls debug-inspector /bin/ls
``

//> Burada `-v /bin/ls:/bin/ls` ile host’taki dosyayı container içine mount ettik.

Windows için:

``powershell
docker run --rm -v C:\Windows\System32\notepad.exe:/notepad.exe debug-inspector /notepad.exe
``

//macOS için:

``bash
docker run --rm -v /bin/ls:/bin/ls debug-inspector /bin/ls
``
/*
---

## 📌 Sonuç

Bu Docker imajı ile:

* Hangi sistemde olursan ol **ELF / PE / Mach-O** dosyalarını analiz edebilirsin.
* İmaj küçük (Alpine tabanlı).
* DWARF ve Go sembollerini de gösterebiliyor.

---

👉 İstersen ben bu Docker imajını **multi-platform build** (Linux, Windows, macOS binary’leri aynı anda üretmek için) ayarlayabilirim. Bunu da ister misin?
*