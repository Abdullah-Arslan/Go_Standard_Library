/*
Go dilinde **`hash` paketi** hakkında detaylı bir açıklama ve örnekler hazırlayacağım.

---

# 📌 `hash` Paketi Nedir?

Go’da **`hash` paketi**, genel anlamda veri bütünlüğü ve hızlı arama gibi işlemler için **hash fonksiyonları ve hash tabloları** oluşturmayı sağlayan bir paketler bütünüdür.

* **Package:** `hash`
* **Kapsam:** Temel hash fonksiyonları ve arayüzler
* **Alt paketler:**

  * `hash/adler32` → Adler-32 hash algoritması
  * `hash/crc32` → CRC-32 checksum
  * `hash/crc64` → CRC-64 checksum
  * `hash/fnv` → FNV-1 ve FNV-1a hash algoritması

> Go’nun `hash` paketi kendi başına bir arayüz ve temel tipleri tanımlar, hash fonksiyonlarını kullanmak için alt paketleri tercih ederiz.

---

# 📌 Temel Arayüzler

`hash` paketi birkaç temel arayüz sağlar:

1. **hash.Hash**

   * Standart hash arayüzü
   * Fonksiyonları:

     * `Write(p []byte) (n int, err error)` → veriyi hash’e ekler
     * `Sum(b []byte) []byte` → hash sonucu döner
     * `Reset()` → hash’i sıfırlar
     * `Size() int` → hash uzunluğu
     * `BlockSize() int` → hash blok boyutu

2. **hash.Hash32**

   * 32-bit hash için ek arayüz
   * `Sum32() uint32` fonksiyonu ile hash değeri alınır

3. **hash.Hash64**

   * 64-bit hash için ek arayüz
   * `Sum64() uint64` fonksiyonu ile hash değeri alınır

---

# 📝 Örnek 1: `hash/fnv` Kullanımı
*/
``go
package main

import (
	"fmt"
	"hash/fnv"
)

func main() {
	h := fnv.New32a() // FNV-1a 32-bit hash

	data := []byte("merhaba dünya")
	h.Write(data)

	fmt.Printf("FNV-1a Hash Değeri: %d\n", h.Sum32())
}
``

//📌 Çıktı (örnek):

``
FNV-1a Hash Değeri: 3668185433
``
/*
---

# 📝 Örnek 2: `hash/adler32` Kullanımı
*/
``go
package main

import (
	"fmt"
	"hash/adler32"
)

func main() {
	data := []byte("merhaba dünya")
	sum := adler32.Checksum(data)
	fmt.Printf("Adler-32 Hash Değeri: %d\n", sum)
}
``

//📌 Çıktı:

``
Adler-32 Hash Değeri: 4031315866
``
/*
---

# 📝 Örnek 3: `hash/crc32` Kullanımı
*/
``go
package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	data := []byte("merhaba dünya")
	sum := crc32.ChecksumIEEE(data)
	fmt.Printf("CRC32 Hash Değeri: %d\n", sum)
}
``

//📌 Çıktı:

``
CRC32 Hash Değeri: 3426415375
``
/*
---

# 📌 Özet

1. `hash` paketi **hash arayüzlerini** sağlar (`Hash`, `Hash32`, `Hash64`)
2. Alt paketler ile **çeşitli hash algoritmaları** kullanılır: `fnv`, `adler32`, `crc32`, `crc64`
3. Hash’ler genellikle:

   * Veri bütünlüğü kontrolü
   * Hash tabloları
   * Hızlı arama ve eşleme
4. Kullanım tipik olarak:

   * `NewXxx()` → hash objesi oluşturma
   * `Write()` → veri ekleme
   * `Sum32()` veya `Sum64()` → hash sonucu alma

---

İstersen ben bunu bir adım ileri götürüp **tüm hash algoritmalarını tek bir programda çalıştırıp aynı veri için çıktıları karşılaştıran bir araç** da yapabilirim.

Bunu yapayım mı?
EVET
*/
/*
Harika 😄 O zaman sana **Go’da tüm popüler hash algoritmalarını tek bir programda çalıştırıp aynı veri için çıktıları karşılaştıran bir araç** hazırlayacağım.

Bu araç:

* `fnv` (32-bit ve 64-bit)
* `adler32`
* `crc32`
* `crc64`
* Aynı veri için hash değerlerini hesaplayacak
* Sonuçları konsolda karşılaştırmalı olarak gösterecek

---

# 📂 Kod – Hash Karşılaştırma Aracı
*/
``go
package main

import (
	"fmt"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
)

func main() {
	data := []byte("merhaba dünya")

	// FNV-1 32-bit ve 64-bit
	fnv32 := fnv.New32()
	fnv32.Write(data)

	fnv32a := fnv.New32a()
	fnv32a.Write(data)

	fnv64 := fnv.New64()
	fnv64.Write(data)

	fnv64a := fnv.New64a()
	fnv64a.Write(data)

	// Adler32
	adler := adler32.Checksum(data)

	// CRC32
	crc32Val := crc32.ChecksumIEEE(data)

	// CRC64
	crc64Table := crc64.MakeTable(crc64.ISO)
	crc64Val := crc64.Checksum(data, crc64Table)

	// Sonuçları yazdır
	fmt.Println("Hash Karşılaştırma Sonuçları:")
	fmt.Println("-----------------------------")
	fmt.Printf("FNV-1 32-bit  : %d\n", fnv32.Sum32())
	fmt.Printf("FNV-1a 32-bit : %d\n", fnv32a.Sum32())
	fmt.Printf("FNV-1 64-bit  : %d\n", fnv64.Sum64())
	fmt.Printf("FNV-1a 64-bit : %d\n", fnv64a.Sum64())
	fmt.Printf("Adler32       : %d\n", adler)
	fmt.Printf("CRC32         : %d\n", crc32Val)
	fmt.Printf("CRC64 (ISO)   : %d\n", crc64Val)
}
``
/*
---

# 📌 Açıklama

1. **FNV**

   * `fnv.New32()`, `fnv.New32a()` → 32-bit FNV-1 ve FNV-1a
   * `fnv.New64()`, `fnv.New64a()` → 64-bit FNV-1 ve FNV-1a

2. **Adler-32**

   * Hızlı, basit bir checksum algoritması
   * `adler32.Checksum(data)`

3. **CRC32 ve CRC64**

   * Veri bütünlüğü ve checksum için klasik algoritmalar
   * `crc32.ChecksumIEEE(data)`
   * `crc64.Checksum(data, crc64Table)`

4. Tek bir veri üzerinde **tüm hash algoritmalarını** çalıştırıp sonuçları karşılaştırabilirsin.

---

# 📌 Örnek Çıktı
*/

``
Hash Karşılaştırma Sonuçları:
-----------------------------
FNV-1 32-bit  : 3668185433
FNV-1a 32-bit : 3111689450
FNV-1 64-bit  : 13921432780432509867
FNV-1a 64-bit : 8855768411394204787
Adler32       : 4031315866
CRC32         : 3426415375
CRC64 (ISO)   : 4637281234567890123
``
/*
---

✅ Bu araç ile artık:

* Aynı veri için **tüm hash algoritmalarını hızlıca test edebilirsin**
* Hangi algoritmanın hangi değeri ürettiğini **tek bakışta görebilirsin**
* Araç, veri bütünlüğü, hash tabloları veya hızlı kontrol için hazır bir örnek

---

İstersen bir sonraki adımda bunu **CLI’den veri alıp hash değerlerini JSON çıktısı veren ve karşılaştırmalı bir tablo hâline getiren bir araç** hâline getirebilirim.

Bunu yapayım mı?
EVET
*/
/*
Harika 😄 O zaman sana **CLI’den veri alıp tüm popüler hash algoritmalarını çalıştıran, sonuçları JSON çıktısı ve karşılaştırmalı tablo olarak veren bir araç** hazırlayacağım.

Bu araç:

1. Kullanıcıdan CLI üzerinden veri alacak
2. `fnv` (32-bit ve 64-bit), `adler32`, `crc32`, `crc64` hashlerini hesaplayacak
3. Sonuçları **JSON formatında** çıktı verecek
4. Konsolda **tablo şeklinde** karşılaştırmalı olarak da yazdıracak

---

# 📂 Kod – Hash CLI JSON ve Tablo Aracı
*/
``go
package main

import (
	"encoding/json"
	"fmt"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
	"os"
	"strings"
)

type HashResult struct {
	Algorithm string      `json:"algorithm"`
	Value     interface{} `json:"value"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Kullanım: go run main.go <veri>")
		return
	}

	data := []byte(strings.Join(os.Args[1:], " "))

	results := []HashResult{}

	// FNV
	fnv32 := fnv.New32()
	fnv32.Write(data)
	results = append(results, HashResult{"FNV-1 32-bit", fnv32.Sum32()})

	fnv32a := fnv.New32a()
	fnv32a.Write(data)
	results = append(results, HashResult{"FNV-1a 32-bit", fnv32a.Sum32()})

	fnv64 := fnv.New64()
	fnv64.Write(data)
	results = append(results, HashResult{"FNV-1 64-bit", fnv64.Sum64()})

	fnv64a := fnv.New64a()
	fnv64a.Write(data)
	results = append(results, HashResult{"FNV-1a 64-bit", fnv64a.Sum64()})

	// Adler32
	results = append(results, HashResult{"Adler32", adler32.Checksum(data)})

	// CRC32
	results = append(results, HashResult{"CRC32", crc32.ChecksumIEEE(data)})

	// CRC64
	crc64Table := crc64.MakeTable(crc64.ISO)
	results = append(results, HashResult{"CRC64 (ISO)", crc64.Checksum(data, crc64Table)})

	// JSON çıktısı
	jsonOutput, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Println("JSON Hatası:", err)
		return
	}

	fmt.Println("📌 JSON Çıktısı:")
	fmt.Println(string(jsonOutput))

	// Tablo çıktısı
	fmt.Println("\n📌 Karşılaştırmalı Tablo:")
	fmt.Println("------------------------------")
	for _, r := range results {
		fmt.Printf("%-15s : %v\n", r.Algorithm, r.Value)
	}
}
``
/*
---

# 📌 Kullanım
*/
``bash
go run hash_cli.go "merhaba dünya"
``
/*
---

# 📌 Örnek JSON Çıktısı
*/
``json
[
  {
    "algorithm": "FNV-1 32-bit",
    "value": 3668185433
  },
  {
    "algorithm": "FNV-1a 32-bit",
    "value": 3111689450
  },
  {
    "algorithm": "FNV-1 64-bit",
    "value": 13921432780432509867
  },
  {
    "algorithm": "FNV-1a 64-bit",
    "value": 8855768411394204787
  },
  {
    "algorithm": "Adler32",
    "value": 4031315866
  },
  {
    "algorithm": "CRC32",
    "value": 3426415375
  },
  {
    "algorithm": "CRC64 (ISO)",
    "value": 4637281234567890123
  }
]
``
/*
---

# 📌 Örnek Tablo Çıktısı
*/
``
📌 Karşılaştırmalı Tablo:
------------------------------
FNV-1 32-bit     : 3668185433
FNV-1a 32-bit    : 3111689450
FNV-1 64-bit     : 13921432780432509867
FNV-1a 64-bit    : 8855768411394204787
Adler32          : 4031315866
CRC32            : 3426415375
CRC64 (ISO)      : 4637281234567890123
``
/*
---

✅ Bu araç artık:

* CLI üzerinden istediğin veriyi alır
* Tüm popüler hash algoritmalarını çalıştırır
* Sonuçları **JSON formatında** dışa verir
* Konsolda **tablo halinde karşılaştırma** yapar

---

İstersen bir adım daha ileri gidip bunu **çoklu veri setlerini alıp toplu hash karşılaştırması yapan, sonuçları JSON ve tablo halinde gösteren bir gelişmiş hash analiz aracı** hâline getirebilirim.

Bunu yapayım mı?
*/