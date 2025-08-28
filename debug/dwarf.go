/*
Go’nun standart kütüphanesinde **`debug/dwarf`** paketi bulunur.
Bu paket, **DWARF (Debug With Arbitrary Record Format)** denilen bir debug formatını okumak ve işlemek için kullanılır.

---

## 📌 DWARF Nedir?

DWARF, **derleyicilerin** binary’lere gömdüğü hata ayıklama bilgisidir.
İçinde şunlar bulunur:

* Fonksiyon, tip, değişken bilgileri
* Kaynak kod satır numaraları
* Struct ve method tanımları
* Symbol tablosu

👉 Özetle: Bir binary dosyası, **DWARF sayesinde kaynak kod seviyesinde debug edilebilir**.

Linux’ta ELF dosyalarının (`/bin/ls` gibi) içine DWARF bilgisi gömülebilir.
Go’nun `debug/dwarf` paketi bu bilgiyi okumaya yarar.

---

## 📦 `debug/dwarf` Paketindeki Önemli Yapılar

* **`Data`** → DWARF verilerinin tamamını temsil eder.
* **`Reader`** → DWARF içindeki `Entry`’leri gezmek için kullanılır.
* **`Entry`** → Her DWARF kaydını temsil eder (ör: tip, fonksiyon, değişken).
* **`Field`** → Entry içindeki bir attribute (örn. `DW_AT_name`, `DW_AT_type`).
* **`Tag`** → Entry’nin tipini belirtir (örn. `DW_TAG_subprogram` = fonksiyon).

---

## 1️⃣ ELF dosyasından DWARF bilgisi okuma
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
	// ELF dosyasını aç
	file, err := elf.Open("/bin/ls")
	if err != nil {
		fmt.Println("Hata:", err)
		os.Exit(1)
	}
	defer file.Close()

	// DWARF verisini oku
	d, err := file.DWARF()
	if err != nil {
		fmt.Println("DWARF bilgisi yok:", err)
		return
	}

	// Reader ile DWARF içindeki entry’leri gez
	r := d.Reader()
	for {
		entry, err := r.Next()
		if err != nil {
			fmt.Println("Okuma hatası:", err)
			break
		}
		if entry == nil {
			break
		}

		fmt.Println("Tag:", entry.Tag)
		for _, f := range entry.Field {
			fmt.Printf("  %s = %v\n", f.Attr, f.Val)
		}
	}
}
``
/*
👉 Bu kod `/bin/ls` içindeki DWARF bilgilerini okur (debug sembolleri varsa).

---

## 2️⃣ DWARF verisini analiz etme (tip ve fonksiyon bulma)
*/
``go
package main

import (
	"debug/dwarf"
	"debug/elf"
	"fmt"
)

func main() {
	f, _ := elf.Open("/bin/ls")
	defer f.Close()

	d, _ := f.DWARF()
	r := d.Reader()

	for {
		e, _ := r.Next()
		if e == nil {
			break
		}

		switch e.Tag {
		case dwarf.TagSubprogram:
			name := e.Val(dwarf.AttrName)
			fmt.Println("Fonksiyon:", name)
		case dwarf.TagVariable:
			name := e.Val(dwarf.AttrName)
			fmt.Println("Değişken:", name)
		case dwarf.TagStructType:
			name := e.Val(dwarf.AttrName)
			fmt.Println("Struct:", name)
		}
	}
}
``
/*
👉 Çıktı: binary içinde tanımlı fonksiyon, global değişken ve struct isimlerini gösterir.

---

## 3️⃣ DWARF Entry’den bilgi çekme

`Entry.Val(attr)` ile belirli bir alan alınabilir.
Örneğin:

* `dwarf.AttrName` → değişken/fonksiyon ismi
* `dwarf.AttrType` → tip bilgisi
* `dwarf.AttrLocation` → bellekteki konum
*/
``go
if e.Tag == dwarf.TagVariable {
    fmt.Println("Değişken adı:", e.Val(dwarf.AttrName))
    fmt.Println("Tip ID:", e.Val(dwarf.AttrType))
}
``
/*
---

## 4️⃣ DWARF Reader Kullanımı

DWARF `Reader`, tıpkı iterator gibi çalışır.

* `Next()` → sonraki entry
* `Seek()` → belirli offset’e git
* `SkipChildren()` → child entry’leri atla

Örnek:
*/
``go
r := d.Reader()
for {
    e, _ := r.Next()
    if e == nil {
        break
    }
    fmt.Println("Tag:", e.Tag, "Offset:", e.Offset)
}
``
/*
---

## 5️⃣ Özet

👉 **`debug/dwarf` paketi** şunları yapmanı sağlar:

* ELF/Mach-O/PE dosyalarındaki **DWARF debug bilgilerini** okumak
* Binary içindeki **fonksiyon, değişken, struct, tip bilgilerini** çıkarmak
* Debugger benzeri bir uygulama geliştirmek

Kullanım senaryoları:

* Profiling ve debugging araçları geliştirmek
* Derlenmiş binary’den kaynak kod bilgisi çıkarmak
* Reverse engineering ve güvenlik araştırmaları

---

✍️ İstersen sana küçük bir **DWARF explorer aracı** yazabilirim:

* Dosya yolunu parametre olarak alır
* İçindeki fonksiyonları, değişkenleri, struct’ları listeler

Bunu ister misin?
EVET
*/

/*
Harika 👍 O zaman sana **DWARF Explorer** adında küçük bir Go aracı hazırlayayım.
Bu araç:

* Parametre olarak bir **Go veya ELF binary** alır
* İçindeki **fonksiyonları, global değişkenleri, struct tiplerini** listeler
* `debug/dwarf` paketini kullanır

---

## 📂 Proje Yapısı

```
dwarf-explorer/
 ├── main.go
```

---

## 📌 `main.go`
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
	if len(os.Args) < 2 {
		fmt.Println("Kullanım: go run main.go <binary_dosyası>")
		return
	}

	binaryPath := os.Args[1]

	// ELF dosyasını aç
	f, err := elf.Open(binaryPath)
	if err != nil {
		fmt.Println("ELF dosyası açılamadı:", err)
		return
	}
	defer f.Close()

	// DWARF verisini oku
	d, err := f.DWARF()
	if err != nil {
		fmt.Println("DWARF bilgisi bulunamadı:", err)
		return
	}

	fmt.Println("📂 Analiz edilen dosya:", binaryPath)
	fmt.Println("=======================================")

	// Reader ile DWARF gezelim
	r := d.Reader()
	for {
		entry, err := r.Next()
		if err != nil {
			fmt.Println("Okuma hatası:", err)
			break
		}
		if entry == nil {
			break
		}

		switch entry.Tag {
		case dwarf.TagSubprogram:
			if name := entry.Val(dwarf.AttrName); name != nil {
				fmt.Println("🔹 Fonksiyon:", name)
			}
		case dwarf.TagVariable:
			if name := entry.Val(dwarf.AttrName); name != nil {
				fmt.Println("🔸 Global Değişken:", name)
			}
		case dwarf.TagStructType:
			if name := entry.Val(dwarf.AttrName); name != nil {
				fmt.Println("🏗️ Struct:", name)
			}
		}
	}
}
``
/*
---

## 🚀 Kullanım

### 1. Aracı çalıştır
*/
``bash
go run main.go /bin/ls
``
/*
### 2. Çıktı örneği (DWARF bilgisi olan bir ELF binary için)

``
📂 Analiz edilen dosya: /bin/ls
=======================================
🔹 Fonksiyon: main
🔹 Fonksiyon: usage
🔸 Global Değişken: optind
🏗️ Struct: stat
``

---

## 📌 Özet

Bu küçük araç ile:

* ELF dosyası açılır
* İçindeki **DWARF section** okunur
* Fonksiyonlar, global değişkenler ve struct tipleri listelenir

---

👉 İstersen ben buna bir de **Dockerfile** ekleyeyim, böylece DWARF Explorer’ı Docker üzerinden de kullanabilirsin. Ekleyeyim mi?
EVET
*/
/*
Süper 👍 şimdi sana **DWARF Explorer** aracı için bir **Dockerfile** ekliyorum.
Böylece makinede Go kurulu olmasa bile ELF binary’leri Docker içinde analiz edebilirsin.

---

## 📂 Proje Yapısı

```
dwarf-explorer/
 ├── main.go
 ├── Dockerfile
```

---
*/
//## 📌 `Dockerfile`

``dockerfile
# 1. Builder aşaması: Go ile derle
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Kaynak kodları kopyala
COPY . .

# Eğer go.mod yoksa oluştur
RUN go mod init dwarf-explorer || true

# Binary'yi build et
RUN go build -o dwarf-explorer main.go

# 2. Final aşama: küçük runtime image
FROM alpine:latest

WORKDIR /root/

# Binary'yi kopyala
COPY --from=builder /app/dwarf-explorer .

# Varsayılan komut
ENTRYPOINT ["./dwarf-explorer"]
```
/*
---

## 🚀 Kullanım

### 1. Docker imajını build et
*/
``bash
docker build -t dwarf-explorer .
``
/*
### 2. ELF binary analiz et

Host’taki ELF dosyasını container içine mount edip çalıştır:
*/
``bash
docker run --rm -v /bin/ls:/bin/ls dwarf-explorer /bin/ls
``
/*
### 3. Örnek çıktı

```
📂 Analiz edilen dosya: /bin/ls
=======================================
🔹 Fonksiyon: main
🔹 Fonksiyon: usage
🔸 Global Değişken: optind
🏗️ Struct: stat
```

---

✅ Bu şekilde DWARF Explorer’ı hem native olarak hem de Docker içinden çalıştırabilirsin.

👉 İstersen ben bu aracı daha da geliştirip, **JSON formatında çıktı veren** bir versiyon da yazabilirim (mesela `--json` parametresi ile). İstiyor musun?
*/