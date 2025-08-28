/*
Go’nun **`debug/plan9obj`** paketini **temelden ileriye** doğru adım adım açıklayayım.

---

## 📌 `debug/plan9obj` Nedir?

Go’daki `debug/plan9obj` paketi, **Plan 9 nesne dosyalarını (`.obj`) okumak için** kullanılan bir standart kütüphanedir.

Plan 9 işletim sistemi (Bell Labs’in geliştirdiği Unix benzeri bir sistem) kendi özel nesne dosya formatına sahiptir. Bu format, ELF (Linux), Mach-O (macOS) ve PE (Windows) dosyalarına benzer ama çok daha basittir.

Go, `debug/plan9obj` ile bu dosyaların:

* Başlıklarını,
* Segmentlerini,
* Sembollerini

okumamıza izin verir.

> Yani: ELF → `debug/elf`, Mach-O → `debug/macho`, PE → `debug/pe`, Plan 9 → `debug/plan9obj`.

---

## 📦 Paket İçeriği

`debug/plan9obj` şu tipleri ve fonksiyonları içerir:

* **`Open(file string)`** → Bir Plan 9 `.obj` dosyasını açar.
* **`File`** → Açılan dosyayı temsil eder. İçinde:

  * `FileHeader` → Dosya hakkında temel bilgiler
  * `Sections` → Tüm bölümler (kod, veri vs.)
  * `Symbols` → Semboller (fonksiyon/ değişken isimleri)
* **`Section`** → Dosya içindeki tek bir bölüm (ör. `.text`, `.data`)
* **`Symbol`** → Fonksiyon/değişken gibi bir sembol.
* **`OpenFile(r io.ReaderAt)`** → Dosya yerine bir reader’dan okuma imkanı verir.

---

## 📂 Plan 9 Nesne Dosya Yapısı

Bir Plan 9 `.obj` dosyası genellikle şunları içerir:

* **Header** → format ve makine türü bilgisi
* **Sections** → kod/veri blokları
* **Symbols** → sembol tablosu

---

## 📝 Örnek 1 – Basitçe Dosya Açma
*/
``go
package main

import (
	"debug/plan9obj"
	"fmt"
	"log"
)

func main() {
	f, err := plan9obj.Open("program.obj")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Printf("Makine türü: %v\n", f.FileHeader.Magic) // dosya formatını gösterir
}
``
/*
📌 Bu kod bir `.obj` dosyası açar ve **magic number** (dosya formatı türü) bilgisini yazdırır.

---

## 📝 Örnek 2 – Bölümleri Listeleme
*/
``go
for _, sec := range f.Sections {
	fmt.Printf("Bölüm: %s, Offset: %d, Boyut: %d\n",
		sec.Name, sec.Offset, sec.Size)
}
``
/*
👉 Çıktı şöyle olabilir:

```
Bölüm: .text, Offset: 64, Boyut: 1200
Bölüm: .data, Offset: 1264, Boyut: 300
```

---

## 📝 Örnek 3 – Sembolleri Listeleme
*/
``go
for _, sym := range f.Symbols {
	fmt.Printf("Sembol: %s, Adres: %#x, Tür: %d\n",
		sym.Name, sym.Value, sym.Type)
}
``
/*
👉 Çıktı:

```
Sembol: main, Adres: 0x1000, Tür: 1
Sembol: printf, Adres: 0x2000, Tür: 2
```

---

## 📝 Örnek 4 – Bölüm İçeriğini Okuma
*/
``go
text := f.Section(".text")
	if text != nil {
		data := make([]byte, text.Size)
		_, err := text.ReadAt(data, 0)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(".text bölümü ilk 16 byte: %x\n", data[:16])
	}
``
/*
👉 `.text` (kod) bölümünün ham byte’larını okur. Bu sayede disassembly yapılabilir.

---

## 📊 Kullanım Alanları

* **Derleyici ve Linker** yazarken,
* **Binary analiz araçları** geliştirirken,
* Plan 9 sistemindeki **programların obj dosyalarını incelemek** için.

---

## 🎯 Özet

* `debug/plan9obj` → Plan 9 `.obj` dosyalarını okumak için.
* **Ana özellikleri**: Başlık, bölüm, sembol tablosu okuma.
* ELF, Mach-O, PE paketlerinin kardeşi.
* Kullanımı oldukça benzer, ama Plan 9’a özgü.

---

👉 İstersen ben sana bu paketi kullanarak **küçük bir "Plan 9 obj explorer" CLI** (mesela `plan9objdump`) yazabilirim. Dosya verildiğinde bölümler + semboller JSON olarak çıksın. İstiyor musun?
EVET
*/

/*
Harika! O zaman sana **Go ile Plan 9 obj dosyalarını analiz eden küçük bir CLI aracı** hazırlayacağım. Bu araç:

* `.obj` dosyasını açacak,
* Bölümleri ve sembolleri okuyacak,
* JSON formatında çıktı verecek (ekrana veya dosyaya yazabilecek).

---

# 📌 `plan9obj_explorer.go`
*/
``go
package main

import (
	"debug/plan9obj"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ObjInfo struct {
	FileHeader struct {
		Magic uint32 `json:"magic"`
	} `json:"file_header"`
	Sections []SectionInfo `json:"sections"`
	Symbols  []SymbolInfo  `json:"symbols"`
}

type SectionInfo struct {
	Name   string `json:"name"`
	Offset int64  `json:"offset"`
	Size   int64  `json:"size"`
}

type SymbolInfo struct {
	Name  string `json:"name"`
	Value uint32 `json:"value"`
	Type  uint8  `json:"type"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Kullanım: %s <plan9.obj> [json_output_file]\n", os.Args[0])
	}

	path := os.Args[1]
	f, err := plan9obj.Open(path)
	if err != nil {
		log.Fatalf("Dosya açılamadı: %v", err)
	}
	defer f.Close()

	info := ObjInfo{}
	info.FileHeader.Magic = f.FileHeader.Magic

	for _, sec := range f.Sections {
		info.Sections = append(info.Sections, SectionInfo{
			Name:   sec.Name,
			Offset: sec.Offset,
			Size:   sec.Size,
		})
	}

	for _, sym := range f.Symbols {
		info.Symbols = append(info.Symbols, SymbolInfo{
			Name:  sym.Name,
			Value: sym.Value,
			Type:  sym.Type,
		})
	}

	// JSON encode
	jsonData, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		log.Fatal("JSON kodlama hatası:", err)
	}

	// Ekrana yaz
	fmt.Println(string(jsonData))

	// Eğer ikinci argüman verildiyse dosyaya kaydet
	if len(os.Args) >= 3 {
		outPath := os.Args[2]
		err := os.WriteFile(outPath, jsonData, 0644)
		if err != nil {
			log.Fatal("JSON dosyaya yazılamadı:", err)
		}
		fmt.Println("JSON çıktısı kaydedildi:", outPath)
	}
}
``
/*
---

# 📌 Kullanım

1. **Obj dosyasını analiz et ve ekrana yaz:**
*/
``bash
go run plan9obj_explorer.go hello.obj
``

//2. **JSON çıktısını dosyaya kaydet:**

``bash
go run plan9obj_explorer.go hello.obj output.json
``

---

//# 📌 Örnek Çıktı

``json
{
  "file_header": {
    "magic": 127947
  },
  "sections": [
    {
      "name": ".text",
      "offset": 64,
      "size": 1024
    },
    {
      "name": ".data",
      "offset": 1088,
      "size": 256
    }
  ],
  "symbols": [
    {
      "name": "main",
      "value": 0,
      "type": 1
    },
    {
      "name": "printf",
      "value": 128,
      "type": 2
    }
  ]
}
``
/*
---

Bu araçla artık **Plan 9 obj dosyalarını hızlıca analiz edip JSON olarak işleyebilirsin**.

İstersen bir sonraki adım olarak bunu **basit bir web arayüzü** ile entegre edebiliriz, tıpkı PE Explorer Web örneğinde olduğu gibi.

Bunu yapalım mı?
*/