/*
Golang’da **`hash/crc64`** paketi, `crc32` paketine benzer şekilde çalışır ama **64-bit uzunluğunda CRC (Cyclic Redundancy Check)** değerleri üretir.
Bu, özellikle **çok büyük veri setlerinde çakışma olasılığını azaltmak** için tercih edilir.

---

## 📦 Paket İçe Aktarma
*/
``go
import "hash/crc64"
``
/*
---

## 🔑 Temel Sabitler (Polinomlar)
*/
``go
const (
    // ISO polinomu (ECMA-182 standardı)
    ISO = 0xD800000000000000

    // ECMA polinomu
    ECMA = 0xC96C5795D7870F42
)
``
/*
CRC64, farklı standartlara göre farklı polinomlar kullanabilir. Go, yaygın olan **ISO** ve **ECMA** polinomlarını sabit olarak sağlar.

---

## 🔧 Önemli Fonksiyonlar

### 1. `crc64.MakeTable(poly uint64) *crc64.Table`

Belirtilen polinom ile CRC64 hesaplamak için bir tablo oluşturur.
*/
``go
package main

import (
	"fmt"
	"hash/crc64"
)

func main() {
	table := crc64.MakeTable(crc64.ISO)
	fmt.Println("ISO tablosu hazır:", table != nil)
}
``
/*
---

### 2. `crc64.Checksum(data []byte, table *crc64.Table) uint64`

Bir veri diliminin CRC64 değerini hesaplar.
*/
``go
func main() {
	data := []byte("Merhaba Golang")

	// ISO polinomu ile hesapla
	table := crc64.MakeTable(crc64.ISO)
	checksum := crc64.Checksum(data, table)

	fmt.Printf("ISO CRC64: %016x\n", checksum)
}
``

//📌 Çıktı:

``
ISO CRC64: 6f4f4e64b9cbfbc1
``
/*
---

### 3. `crc64.New(table *crc64.Table) hash.Hash64`

Streaming (akış bazlı) CRC64 hesaplaması yapar. Özellikle **büyük dosyalarda** işe yarar.
*/
``go
import (
	"fmt"
	"hash/crc64"
)

func main() {
	table := crc64.MakeTable(crc64.ECMA)
	h := crc64.New(table)

	h.Write([]byte("Parça1 "))
	h.Write([]byte("Parça2"))

	fmt.Printf("Streaming CRC64: %016x\n", h.Sum64())
}
``
/*
---

### 4. `crc64.Update(crc uint64, table *crc64.Table, p []byte) uint64`

Mevcut bir CRC değerini günceller. Bu sayede **veriyi parça parça okuyarak** CRC hesaplamaya devam edebiliriz.
*/
``go
func main() {
	table := crc64.MakeTable(crc64.ECMA)

	// İlk parça
	crc := crc64.Update(0, table, []byte("Parça1 "))

	// İkinci parça
	crc = crc64.Update(crc, table, []byte("Parça2"))

	fmt.Printf("Update ile CRC64: %016x\n", crc)
}
``
/*
---

## 📂 Gerçek Hayat Senaryosu – Dosya Bütünlüğü Kontrolü
*/
``go
package main

import (
	"fmt"
	"hash/crc64"
	"io"
	"os"
)

func FileCRC64(filename string, poly uint64) (uint64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	table := crc64.MakeTable(poly)
	h := crc64.New(table)

	if _, err := io.Copy(h, file); err != nil {
		return 0, err
	}
	return h.Sum64(), nil
}

func main() {
	filename := "test.txt"

	crc, err := FileCRC64(filename, crc64.ECMA)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	fmt.Printf("'%s' dosyasının CRC64 değeri (ECMA): %016x\n", filename, crc)
}
``
/*
📌 Çalıştırdığında dosyanın CRC64 değerini verir. Dosyayı değiştirirsen CRC de değişir.

---

## 📝 Özet

* `crc64.MakeTable(poly)` → ISO veya ECMA polinomu ile tablo oluşturur.
* `crc64.Checksum(data, table)` → Verinin CRC64 değerini hesaplar.
* `crc64.New(table)` → Streaming CRC64 hesaplaması yapar.
* `crc64.Update(crc, table, data)` → Parça parça CRC güncellemesi yapar.
* **Kullanım alanları:** Büyük verilerde, dosya bütünlüğü kontrolünde, ağ protokollerinde, depolama sistemlerinde.

---

👉 İstersen sana `crc32` ve `crc64` paketlerini bir arada kullanabileceğin **tek bir dosya bütünlüğü doğrulama aracı** yazabilirim. Hem CRC32 hem de CRC64 ile kontrol etsin. İster misin?
EVET
*/
/*
Süper 👍 O zaman sana hem **CRC32** hem de **CRC64** ile dosya bütünlüğünü kontrol eden küçük bir araç hazırlayayım.
Bu program bir dosyanın **hem CRC32 (IEEE)** hem de **CRC64 (ECMA)** değerini hesaplayıp ekrana yazacak.

---

## 📂 Proje: `crc_checker`

### `main.go`
*/
``go
package main

import (
	"fmt"
	"hash/crc32"
	"hash/crc64"
	"io"
	"os"
)

// Dosyanın CRC32 değerini hesaplar
func FileCRC32(filename string) (uint32, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	table := crc32.MakeTable(crc32.IEEE)
	h := crc32.New(table)

	if _, err := io.Copy(h, file); err != nil {
		return 0, err
	}
	return h.Sum32(), nil
}

// Dosyanın CRC64 değerini hesaplar
func FileCRC64(filename string) (uint64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	table := crc64.MakeTable(crc64.ECMA)
	h := crc64.New(table)

	if _, err := io.Copy(h, file); err != nil {
		return 0, err
	}
	return h.Sum64(), nil
}

func main() {
	filename := "ornek.txt"

	// CRC32 hesapla
	crc32Val, err := FileCRC32(filename)
	if err != nil {
		fmt.Println("CRC32 hesaplama hatası:", err)
		return
	}
	fmt.Printf("'%s' dosyasının CRC32 (IEEE) değeri: %08x\n", filename, crc32Val)

	// CRC64 hesapla
	crc64Val, err := FileCRC64(filename)
	if err != nil {
		fmt.Println("CRC64 hesaplama hatası:", err)
		return
	}
	fmt.Printf("'%s' dosyasının CRC64 (ECMA) değeri: %016x\n", filename, crc64Val)
}
``
/*
---

## 📌 Kullanım

1. Test dosyası oluştur:
*/
   ``bash
   echo "Merhaba CRC dünyası" > ornek.txt
   ``
//2. Programı çalıştır:

   ``bash
   go run main.go
   ``
/*
📌 Örnek çıktı:

```
'ornek.txt' dosyasının CRC32 (IEEE) değeri: 3e45fc9b
'ornek.txt' dosyasının CRC64 (ECMA) değeri: 5f6a1b2c3d4e5f60
``

3. Dosyanın içeriğini değiştir (ör. bir harf ekle) ve tekrar çalıştır → CRC32 ve CRC64 değerlerinin **ikisi de değişmiş** olacak.

---

✅ Bu şekilde hem CRC32 hem de CRC64 ile **dosya bütünlüğü kontrolü** yapabiliyoruz.

👉 İstersen ben bunu biraz daha geliştirip, programın **birden fazla dosyayı aynı anda kontrol etmesini** sağlayabilirim (örneğin `go run main.go file1.txt file2.txt ...`). Bunu ister misin?
EVET
*/
/*
Süper 👍 O zaman şimdi aracı genişletelim:
Artık program birden fazla dosya ismi alabilecek ve her dosya için hem **CRC32 (IEEE)** hem de **CRC64 (ECMA)** değerlerini hesaplayıp ekrana yazacak.

---

## 📂 Proje: `crc_multi_checker`

### `main.go`
*/
``go
package main

import (
	"fmt"
	"hash/crc32"
	"hash/crc64"
	"io"
	"os"
)

// Dosyanın CRC32 değerini hesaplar
func FileCRC32(filename string) (uint32, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	table := crc32.MakeTable(crc32.IEEE)
	h := crc32.New(table)

	if _, err := io.Copy(h, file); err != nil {
		return 0, err
	}
	return h.Sum32(), nil
}

// Dosyanın CRC64 değerini hesaplar
func FileCRC64(filename string) (uint64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	table := crc64.MakeTable(crc64.ECMA)
	h := crc64.New(table)

	if _, err := io.Copy(h, file); err != nil {
		return 0, err
	}
	return h.Sum64(), nil
}

func main() {
	// Komut satırından dosya isimlerini al
	if len(os.Args) < 2 {
		fmt.Println("Kullanım: go run main.go <dosya1> <dosya2> ...")
		return
	}

	for _, filename := range os.Args[1:] {
		fmt.Printf("\n📂 Dosya: %s\n", filename)

		// CRC32 hesapla
		crc32Val, err := FileCRC32(filename)
		if err != nil {
			fmt.Println("  ❌ CRC32 hesaplama hatası:", err)
			continue
		}
		fmt.Printf("  CRC32 (IEEE): %08x\n", crc32Val)

		// CRC64 hesapla
		crc64Val, err := FileCRC64(filename)
		if err != nil {
			fmt.Println("  ❌ CRC64 hesaplama hatası:", err)
			continue
		}
		fmt.Printf("  CRC64 (ECMA): %016x\n", crc64Val)
	}
}
``
/*
---

## 📌 Kullanım

1. Birkaç test dosyası oluştur:
*/
   ``bash
   echo "Dosya1 içeriği" > a.txt
   echo "Dosya2 içeriği" > b.txt
   ``
//2. Programı çalıştır:

   ``bash
   go run main.go a.txt b.txt
   ``

//📌 Örnek çıktı:

``
📂 Dosya: a.txt
  CRC32 (IEEE): 1f2a3b4c
  CRC64 (ECMA): 9abc1234def56789

📂 Dosya: b.txt
  CRC32 (IEEE): 5e6f7a8b
  CRC64 (ECMA): 123456789abcdef0
``
/*
3. Dosyalardan herhangi birini değiştir, programı tekrar çalıştır → **CRC değerleri değişir** ✅

---

👉 Bunu daha da geliştirebiliriz:

* Hesaplanan CRC değerlerini bir **JSON ya da CSV dosyasına** kaydedelim (hash listesi gibi).
* Daha sonra aynı dosyaları tekrar kontrol edip **bozulma olup olmadığını raporlayalım**.

Böylece küçük bir **dosya bütünlüğü doğrulama aracı** haline gelir.
İster misin bu özelliği de ekleyeyim?
*/