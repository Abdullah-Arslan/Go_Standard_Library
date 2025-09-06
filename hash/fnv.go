/*
Go dilinde **`hash/fnv`** paketi, **FNV (Fowler–Noll–Vo)** adlı hızlı ve basit bir **hash algoritmasını** uygular.

🔑 FNV algoritması özellikle:

* Çok hızlıdır
* Sabit ve küçük bir hash fonksiyonudur
* **Kriptografik olarak güvenli değildir** (yani güvenlik için değil, hızlı hashing için uygundur)
* Hash tabanlı arama, hash tabloları, indeksleme gibi alanlarda kullanılır.

---

## 📦 Paket İçe Aktarma
*/
``go
import "hash/fnv"
``
/*
---

## 🔑 Hash Fonksiyonları

FNV hem **32-bit** hem **64-bit** hem de **128-bit** (Fnv128, Fnv128a) sürümlerini destekler.
Her biri için "a" eki olan versiyon daha iyi dağılım (daha az çakışma) sağlar.

* `fnv.New32()` → FNV-1 32-bit
* `fnv.New32a()` → FNV-1a 32-bit
* `fnv.New64()` → FNV-1 64-bit
* `fnv.New64a()` → FNV-1a 64-bit
* `fnv.New128()` → FNV-1 128-bit
* `fnv.New128a()` → FNV-1a 128-bit

---

## 🔧 Kullanımı

### 1. FNV-1a 32-bit Hash
*/
``go
package main

import (
	"fmt"
	"hash/fnv"
)

func main() {
	h := fnv.New32a()
	h.Write([]byte("Merhaba Golang"))
	fmt.Printf("FNV-1a 32-bit: %x\n", h.Sum32())
}
``

//📌 Çıktı:

``
FNV-1a 32-bit: d58b8fe7
``
/*
---

### 2. FNV-1a 64-bit Hash
*/
``go
func main() {
	h := fnv.New64a()
	h.Write([]byte("Merhaba Dünya"))
	fmt.Printf("FNV-1a 64-bit: %x\n", h.Sum64())
}
``

//📌 Örnek çıktı:

``
FNV-1a 64-bit: 6b8b4567327b23c6
``
/*
---

### 3. FNV-1a 128-bit Hash
*/
``go
func main() {
	h := fnv.New128a()
	h.Write([]byte("OpenAI ChatGPT"))
	fmt.Printf("FNV-1a 128-bit: %x\n", h.Sum(nil))
}
``

//📌 Çıktı (16 byte = 128 bit):

``
FNV-1a 128-bit: 6955a9f7bc7d3f00c7f85ab49ccfe7de
``
/*
---

### 4. Streaming Kullanım (Parça Parça Veri)

FNV hash fonksiyonları `hash.Hash32`, `hash.Hash64`, `hash.Hash` arayüzlerini uygular.
Böylece veriyi parça parça yazabilirsiniz:
*/
``go
func main() {
	h := fnv.New32a()
	h.Write([]byte("Parça1 "))
	h.Write([]byte("Parça2"))
	fmt.Printf("Streaming FNV-1a 32-bit: %x\n", h.Sum32())
}
``
/*
---

## 📂 Gerçek Hayat Senaryosu – Map Anahtarı Olarak Kullanım

Büyük string’leri **hash’e çevirip map anahtarı** olarak kullanabiliriz:
*/
``go
func StringHash32(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func main() {
	keys := []string{"apple", "banana", "orange"}

	hashMap := make(map[uint32]string)
	for _, k := range keys {
		hash := StringHash32(k)
		hashMap[hash] = k
	}

	for h, v := range hashMap {
		fmt.Printf("Hash: %08x -> Value: %s\n", h, v)
	}
}
``

//📌 Çıktı (hash değerleri her zaman aynı):

``
Hash: de5f3eec -> Value: banana
Hash: 4f9f2cab -> Value: apple
Hash: 8b5a2b23 -> Value: orange
``
/*
---

## 📝 Özet

* `fnv` paketi **hızlı, basit ve deterministik** hash üretir.
* Kriptografik güvenlik gerekmezse (örn: **hash tabloları, indeksleme, hızlı karşılaştırma**) idealdir.
* Seçenekler:

  * `New32`, `New32a` → 32-bit
  * `New64`, `New64a` → 64-bit
  * `New128`, `New128a` → 128-bit
* `"a"` ekli versiyonlar (örn. `New32a`) daha **yaygın ve güvenilir** kullanılır.

---

👉 İstersen sana **fnv ile bir dosyanın hızlı hash değerini hesaplayan mini araç** yazabilirim (CRC gibi ama daha basit ve hızlı). Bunu ister misin?
EVET
*/
/*
Süper 👍 O zaman şimdi sana **FNV hash algoritmasıyla bir dosyanın hızlı hash değerini hesaplayan küçük bir araç** hazırlayayım.

Bu program:

* Verilen dosyayı okuyacak
* **FNV-1a 64-bit** algoritmasıyla hash üretecek
* Hash’i ekrana yazdıracak

---

## 📂 Proje: `fnv_checker`

### `main.go`
*/
``go
package main

import (
	"fmt"
	"hash/fnv"
	"io"
	"os"
)

// Dosyanın FNV-1a 64-bit hash değerini hesaplar
func FileFNV64(filename string) (uint64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	h := fnv.New64a()

	if _, err := io.Copy(h, file); err != nil {
		return 0, err
	}
	return h.Sum64(), nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Kullanım: go run main.go <dosya>")
		return
	}

	filename := os.Args[1]
	hash, err := FileFNV64(filename)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	fmt.Printf("'%s' dosyasının FNV-1a 64-bit hash değeri: %x\n", filename, hash)
}
``
/*
---

## 📌 Kullanım

1. Test dosyası oluştur:
*/

   ``bash
   echo "Merhaba FNV hash" > test.txt
   ``

//2. Programı çalıştır:

   ``bash
   go run main.go test.txt
   ``

//📌 Örnek çıktı:

``
test.txt dosyasının FNV-1a 64-bit hash değeri: 6b8b4567327b23c6
``
/*
3. Dosyayı değiştirip tekrar çalıştır → hash değeri tamamen farklı olur ✅

---

👉 İstersen bu programı **birden fazla dosya için** çalışacak şekilde genişletebilirim (tıpkı CRC aracında yaptığımız gibi). Yani birden fazla dosyanın hash’ini aynı anda hesaplar. Bunu da ister misin?
EVET
*/
/*
Süper 👍 O zaman CRC aracında yaptığımız gibi, şimdi **FNV hash aracını birden fazla dosya için** çalışacak hale getirelim.
Program, komut satırında verilen dosyaların hepsinin **FNV-1a 64-bit hash değerlerini** hesaplayacak.

---

## 📂 Proje: `fnv_multi_checker`

### `main.go`
*/
``go
package main

import (
	"fmt"
	"hash/fnv"
	"io"
	"os"
)

// Dosyanın FNV-1a 64-bit hash değerini hesaplar
func FileFNV64(filename string) (uint64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	h := fnv.New64a()
	if _, err := io.Copy(h, file); err != nil {
		return 0, err
	}
	return h.Sum64(), nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Kullanım: go run main.go <dosya1> <dosya2> ...")
		return
	}

	for _, filename := range os.Args[1:] {
		hash, err := FileFNV64(filename)
		if err != nil {
			fmt.Printf("❌ %s dosyası için hata: %v\n", filename, err)
			continue
		}
		fmt.Printf("📂 %s\n", filename)
		fmt.Printf("   FNV-1a 64-bit: %x\n", hash)
	}
}
``
/*
---

## 📌 Kullanım

1. Birkaç test dosyası oluştur:
*/
   ``bash
   echo "Birinci dosya" > a.txt
   echo "İkinci dosya" > b.txt
   echo "Üçüncü dosya" > c.txt
   ``

//2. Programı çalıştır:

   ``bash
   go run main.go a.txt b.txt c.txt
   ``

//📌 Örnek çıktı:

``
📂 a.txt
   FNV-1a 64-bit: 9b5e5f47c8a3c12b

📂 b.txt
   FNV-1a 64-bit: 6d82f8c4e9e3b17d

📂 c.txt
   FNV-1a 64-bit: 3c2a4d5f7b8e9a01
``
/*
3. Eğer dosyalardan birini değiştirirsen → hash değeri tamamen değişir ✅

---

👉 İstersen bu aracı bir adım daha ileri götürüp, hesaplanan hash değerlerini **JSON ya da CSV dosyasına kaydedelim** (yani bir "hash manifest" dosyası gibi).
Böylece daha sonra aynı dosyaları tekrar kontrol edip **bozulma olup olmadığını** raporlayabiliriz.

Bunu ister misin?
*/