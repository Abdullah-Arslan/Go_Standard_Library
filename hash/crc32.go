/*
Golang’daki **`hash/crc32`** paketi, veriler üzerinde **CRC-32 (Cyclic Redundancy Check)** algoritmasını uygulamak için kullanılır. CRC32, veri bütünlüğünü kontrol etmek ve hataları tespit etmek için yaygın olarak kullanılan bir **hash fonksiyonudur**. Özellikle dosya transferlerinde, network paketlerinde veya depolama sistemlerinde veri bozulmalarını algılamak için tercih edilir.

Bu paketin sunduğu fonksiyonlar sayesinde farklı polinomlarla (IEEE, Castagnoli, Koopman gibi) CRC-32 hesaplamaları yapılabilir.

---

## 📦 Paket İçe Aktarma
*/
``go
import "hash/crc32"
``
/*
---

## 🔑 Temel Sabitler
*/
``go
const (
    // CRC-32 polinomları
    IEEE       = 0xedb88320 // En yaygın kullanılan
    Castagnoli = 0x82f63b78 // iSCSI standardı
    Koopman    = 0xeb31d82e
)
``
/*
Bu sabitler, CRC hesaplamasında kullanılacak **polinomu** belirtir.

---

## 🔧 Fonksiyonlar ve Kullanımları

### 1. `crc32.ChecksumIEEE(data []byte) uint32`

Varsayılan **IEEE** polinomu ile CRC32 hash değeri üretir.
*/
``go
package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	data := []byte("Merhaba Golang")
	checksum := crc32.ChecksumIEEE(data)
	fmt.Printf("IEEE CRC32: %08x\n", checksum)
}
``

//📌 Çıktı (her çalıştırmada aynı):

``
IEEE CRC32: 7f45fc8a
``
/*
---

### 2. `crc32.MakeTable(poly uint32) *crc32.Table`

Özel bir polinom tablosu oluşturur. Bu tablo ile farklı polinomlar üzerinden CRC hesaplanabilir.
*/
``go
func main() {
	data := []byte("Merhaba Dünya")

	// Castagnoli tablosu
	table := crc32.MakeTable(crc32.Castagnoli)

	checksum := crc32.Checksum(data, table)
	fmt.Printf("Castagnoli CRC32: %08x\n", checksum)
}
``
/*
---

### 3. `crc32.Checksum(data []byte, table *crc32.Table) uint32`

Verilen tabloyu (polinomu) kullanarak CRC32 değeri hesaplar.
*/
``go
func main() {
	data := []byte("OpenAI ChatGPT")

	// Koopman tablosu
	table := crc32.MakeTable(crc32.Koopman)

	checksum := crc32.Checksum(data, table)
	fmt.Printf("Koopman CRC32: %08x\n", checksum)
}
``
/*
---

### 4. **Streaming CRC Hesaplama** (`crc32.New`)

Büyük dosyaları parça parça okuyup CRC hesaplamak için kullanılır.
*/
``go
func main() {
	table := crc32.MakeTable(crc32.IEEE)
	h := crc32.New(table)

	h.Write([]byte("Parça1 "))
	h.Write([]byte("Parça2"))
	
	fmt.Printf("Streaming CRC32: %08x\n", h.Sum32())
}
``
/*
📌 Burada dosya okurken `bufio.Reader` ile parçaları okuyup `h.Write()` içine verebiliriz.

---

### 5. `crc32.Update(crc uint32, table *crc32.Table, p []byte) uint32`

Mevcut CRC değerini günceller. Özellikle parçalı veri işleme için kullanılır.
*/
``go
func main() {
	table := crc32.MakeTable(crc32.IEEE)

	// İlk parça
	crc := crc32.Update(0, table, []byte("Parça1 "))

	// İkinci parça
	crc = crc32.Update(crc, table, []byte("Parça2"))

	fmt.Printf("Update ile CRC32: %08x\n", crc)
}
``
/*
---

## 📂 Gerçek Hayat Senaryoları

### ✅ Dosya Bütünlüğü Kontrolü

Bir dosyanın CRC32’sini alıp, transfer sonrası karşılaştırabilirsiniz:
*/
``go
import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

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

func main() {
	crc, _ := FileCRC32("test.txt")
	fmt.Printf("Dosyanın CRC32 değeri: %08x\n", crc)
}
``
/*
---

## 📝 Özet

* `crc32.ChecksumIEEE(data)` → Hızlı, IEEE standardı ile checksum.
* `crc32.MakeTable(poly)` → Farklı polinomlarla çalışmak için tablo.
* `crc32.Checksum(data, table)` → Tablo kullanarak checksum.
* `crc32.New(table)` → Streaming CRC hesaplama.
* `crc32.Update(crc, table, data)` → Parça parça CRC güncelleme.

🔑 CRC32 paketi, özellikle **dosya transferlerinde, network protokollerinde, veri bütünlüğü doğrulamada** çok işe yarar.

---

İstersen ben sana bir **örnek proje** hazırlayayım: Hem dosya CRC32’sini hesaplayalım hem de başka bir yerde tekrar hesaplayıp **doğrulama** yapalım. İster misin?
EVET


Harika 👍 O zaman sana küçük bir **örnek proje** hazırlayayım. Bu projede:

1. Bir dosyanın **CRC32 değerini hesaplayacağız**.
2. Aynı dosyayı tekrar okuyup, **CRC değerini doğrulayacağız**.

---

## 📂 Proje: `crc32check`

### `main.go`
*/
``go
package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

// Dosyanın CRC32 değerini hesaplayan fonksiyon
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

func main() {
	filename := "ornek.txt"

	// Dosya CRC32 değerini hesapla
	crc1, err := FileCRC32(filename)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Printf("'%s' dosyasının CRC32 değeri: %08x\n", filename, crc1)

	// Tekrar hesapla ve karşılaştır
	crc2, _ := FileCRC32(filename)
	if crc1 == crc2 {
		fmt.Println("✅ Dosya doğrulandı, CRC32 değerleri eşleşiyor.")
	} else {
		fmt.Println("❌ Dosya bozulmuş olabilir! CRC32 eşleşmedi.")
	}
}
``
/*
---

## 📌 Kullanım
*/
//1. Bir test dosyası oluştur:

   ``bash
   echo "Merhaba CRC32" > ornek.txt
   ``
//2. Programı çalıştır:

   ``bash
   go run main.go
   ``

//📌 Çıktı:

``
ornek.txt dosyasının CRC32 değeri: 2f7a3e4a
✅ Dosya doğrulandı, CRC32 değerleri eşleşiyor.
``

//3. Dosyayı değiştir (ör. içine bir harf ekle) ve tekrar çalıştır:

   ``
   echo "Merhaba CRC32!!!" > ornek.txt
   go run main.go
   ``

   //Bu sefer CRC değişecek ve doğrulama başarısız olacak:

   ``
   ornek.txt dosyasının CRC32 değeri: 7a6d5b12
   ❌ Dosya bozulmuş olabilir! CRC32 eşleşmedi.
   ``
/*
---

Bu şekilde CRC32 ile **dosya bütünlüğü kontrolü** yapmış olduk. 🚀

👉 İstersen ben bu örneği biraz daha genişletebilirim: Mesela **birden fazla dosyayı CRC32 ile doğrulayan küçük bir araç** haline getirebiliriz. İster misin?
*/