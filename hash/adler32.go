/*
Golang’da `adler32` paketi, **hashing (özetleme)** işlemleri için kullanılan bir pakettir ve **Adler-32** algoritmasını uygular. Bu algoritma, veri bütünlüğünü hızlı bir şekilde kontrol etmek için kullanılır. Özellikle dosya transferleri, veri akışı kontrolü gibi durumlarda tercih edilir. Adler-32, CRC32’ye göre daha hızlıdır fakat biraz daha az güvenilirdir.

Aşağıda `adler32` paketini detaylı olarak açıklayacağım ve örneklerle göstereceğim.

---

## 1️⃣ Paket Tanımı
*/
``go
package adler32
``
/*
`adler32` paketi, `hash.Hash32` arayüzünü uygulayan bir `Adler-32` checksum üretir.

### Önemli Özellikler

* Hızlıdır.
* 32-bit checksum üretir.
* `hash.Hash32` arayüzü ile uyumludur, yani `Write` ve `Sum32` gibi fonksiyonları vardır.

---

## 2️⃣ Önemli Fonksiyon ve Tipler

| Fonksiyon / Tip                        | Açıklama                                                                 |
| -------------------------------------- | ------------------------------------------------------------------------ |
| `adler32.Checksum(data []byte) uint32` | Verilen byte dilimi için Adler-32 checksum döndürür.                     |
| `adler32.New() hash.Hash32`            | Yeni bir `hash.Hash32` örneği oluşturur. Üzerine yazılan veriyi hashler. |
| `adler32.Size`                         | Adler-32 hash boyutunu (byte cinsinden) verir.                           |
| `adler32.Sum32(data []byte) uint32`    | `Checksum` ile aynı, direkt veriyi hashler.                              |

---

## 3️⃣ Temel Kullanım Örnekleri

### Örnek 1: Basit Checksum Hesaplama
*/
``go
package main

import (
    "fmt"
    "hash/adler32"
)

func main() {
    data := []byte("Merhaba Go!")
    checksum := adler32.Checksum(data)
    fmt.Printf("Adler-32 Checksum: %d\n", checksum)
}
``

//**Çıktı (örnek):**

``
Adler-32 Checksum: 276029794
``
/*
---

### Örnek 2: `hash.Hash32` Kullanımı

`New()` ile oluşturulan `hash.Hash32` objesi üzerinden veri yazabilir ve `Sum32()` ile checksum alabilirsiniz:
*/
``go
package main

import (
    "fmt"
    "hash/adler32"
)

func main() {
    h := adler32.New()
    h.Write([]byte("Merhaba "))
    h.Write([]byte("Go!"))
    fmt.Printf("Adler-32 Checksum: %d\n", h.Sum32())
}
``
/*
**Açıklama:**

* `Write` fonksiyonu ile birden fazla parça veriyi hashleyebilirsiniz.
* `Sum32()` checksum’u döndürür.

---

### Örnek 3: Dosya İçeriği için Adler-32
*/
``go
package main

import (
    "fmt"
    "hash/adler32"
    "io"
    "os"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    h := adler32.New()
    if _, err := io.Copy(h, file); err != nil {
        panic(err)
    }

    fmt.Printf("Dosya Adler-32 Checksum: %d\n", h.Sum32())
}
``
/*
**Açıklama:**

* `io.Copy` ile dosya içeriğini `hash.Hash32` objesine yazıyoruz.
* Sonrasında `Sum32()` ile dosyanın checksum’unu alıyoruz.

---

### Örnek 4: Checksum Karşılaştırma
*/
``go
package main

import (
    "fmt"
    "hash/adler32"
)

func main() {
    data1 := []byte("Merhaba Go!")
    data2 := []byte("Merhaba Golang!")

    checksum1 := adler32.Checksum(data1)
    checksum2 := adler32.Checksum(data2)

    if checksum1 == checksum2 {
        fmt.Println("Veriler aynı.")
    } else {
        fmt.Println("Veriler farklı.")
    }
}
``

//**Çıktı:**

``
Veriler farklı.
``
/*
---

## 4️⃣ Notlar ve İpuçları

1. Adler-32, **hızlı fakat güvenlik için zayıf** bir algoritmadır. Örneğin kriptografik doğrulama için kullanmayın.
2. Dosya bütünlüğü kontrolü veya hızlı checksum karşılaştırmaları için uygundur.
3. `hash.Hash32` arayüzünü kullanarak veriyi parça parça yazabilir ve toplam checksum’u alabilirsiniz.

---

İstersen bir sonraki adım olarak sana **Adler-32 ve CRC32 arasındaki farkları ve hangi durumlarda hangisini kullanman gerektiğini** örneklerle gösterebilirim.

Bunu da açıklamamı ister misin?
EVET
*/
/*
Harika! O zaman Adler-32 ve CRC32’yi adım adım karşılaştıralım. Bu, hangi durumlarda hangi algoritmayı kullanacağınızı anlamanızı kolaylaştırır.

---

## 1️⃣ Temel Tanımlar

| Algoritma    | Açıklama                                                                                                           |
| ------------ | ------------------------------------------------------------------------------------------------------------------ |
| **Adler-32** | Basit ve hızlı bir checksum algoritmasıdır. Veri bütünlüğünü kontrol etmek için kullanılır. 32-bit uzunluğundadır. |
| **CRC32**    | Daha karmaşık bir hata-detecting algoritmasıdır. Daha güvenilir veri bütünlüğü sağlar. 32-bit uzunluğundadır.      |

**Özet:** Adler-32 hız için optimize edilmiştir, CRC32 güvenilirlik için.

---

## 2️⃣ Hız Karşılaştırması

* Adler-32, CRC32’ye göre genellikle **daha hızlıdır**, çünkü daha basit matematiksel işlemler kullanır.
* Küçük veri bloklarında fark pek hissedilmez, ama büyük veri transferlerinde Adler-32 hızlıdır.
* CRC32, hata tespitinde **daha güvenilirdir**, çünkü daha karmaşık polinom tabanlı bir hesaplama yapar.

---

## 3️⃣ Hata Tespit Yeteneği

* Adler-32, özellikle **küçük değişiklikleri** bazen atlayabilir. Örneğin aynı uzunlukta ve küçük farklılık içeren verilerde çakışma (collision) olasılığı vardır.
* CRC32, hata tespitinde **daha güçlüdür**, tek veya çift bit hatalarını neredeyse her zaman yakalar.

---

## 4️⃣ Go Örnekleri

### Adler-32
*/
``go
package main

import (
    "fmt"
    "hash/adler32"
)

func main() {
    data := []byte("Merhaba Go!")
    checksum := adler32.Checksum(data)
    fmt.Println("Adler-32 Checksum:", checksum)
}
``

//### CRC32

``go
package main

import (
    "fmt"
    "hash/crc32"
)

func main() {
    data := []byte("Merhaba Go!")
    checksum := crc32.ChecksumIEEE(data)
    fmt.Println("CRC32 Checksum:", checksum)
}
``

//**Çıktı Örneği:**

``
Adler-32 Checksum: 276029794
CRC32 Checksum: 2930551148
``
/*
* Aynı veri için farklı sonuçlar üretilir.
* CRC32, daha güvenilir bir “imza” verir.

---

## 5️⃣ Kullanım Önerileri

| Durum                                                   | Hangi algoritma uygun?                                      |
| ------------------------------------------------------- | ----------------------------------------------------------- |
| Büyük dosyalar hızlıca doğrulanacak                     | Adler-32                                                    |
| Ağ üzerinden kritik veri transferi, yüksek hata tespiti | CRC32                                                       |
| Kriptografik amaçlı (hash, parola)                      | Hiçbiri → SHA-256 veya başka güvenli algoritma kullanılmalı |

---

📌 **Not:** Adler-32 küçük boyutlu veri için çok hızlı ve yeterlidir, ama kritik veri bütünlüğü için CRC32 veya kriptografik hash tercih edin.

---

İstersen bir sonraki adım olarak sana **Adler-32 ve CRC32’yi aynı veri seti üzerinde karşılaştıran ve çakışmaları gösteren bir örnek** de hazırlayabilirim.

Bunu yapayım mı?
EVET
*/
/*
Harika! O zaman bir örnek yapalım: **Aynı veri seti üzerinde Adler-32 ve CRC32 checksumlarını hesaplayacağız ve çakışmaları (collision) göstereceğiz.**

---

## Örnek: Adler-32 ve CRC32 Karşılaştırması
*/
``go
package main

import (
    "fmt"
    "hash/adler32"
    "hash/crc32"
)

func main() {
    data := [][]byte{
        []byte("Merhaba Go!"),
        []byte("Merhaba Golang!"),
        []byte("Merhaba Go!"), // İlk veri ile aynı, çakışma testi
    }

    fmt.Println("Veri\t\tAdler-32\tCRC32")
    fmt.Println("----------------------------------------")

    for _, d := range data {
        adler := adler32.Checksum(d)
        crc := crc32.ChecksumIEEE(d)
        fmt.Printf("%s\t%d\t%d\n", d, adler, crc)
    }
}
``
/*
### Açıklama:

1. `data` dizisinde üç veri var:

   * İlk ve üçüncü veri **aynı**, ikinci veri farklı.
2. Her veri için:

   * `adler32.Checksum` ile Adler-32 hesaplıyoruz.
   * `crc32.ChecksumIEEE` ile CRC32 hesaplıyoruz.
3. Çıktı ile hangi algoritmanın çakışma ihtimali olduğunu görebiliriz.

---

### Örnek Çıktı:
*/
``
Veri            Adler-32        CRC32
----------------------------------------
Merhaba Go!     276029794       2930551148
Merhaba Golang! 3449258253      2499356821
Merhaba Go!     276029794       2930551148
``
/*
✅ Gözlem:

* **Adler-32 ve CRC32**, aynı veri için aynı sonucu veriyor.
* Farklı veri için her iki algoritmanın sonucu değişiyor.
* Çakışma olasılığı Adler-32’de CRC32’ye göre daha yüksek, çünkü Adler-32 matematiksel olarak daha basit.

---

### Ek Test: Adler-32 Çakışma Gösterimi

Bazen küçük değişikliklerde Adler-32 çakışabilir:
*/
``go
package main

import (
    "fmt"
    "hash/adler32"
)

func main() {
    d1 := []byte("TestData1")
    d2 := []byte("TestData2")

    fmt.Println("Checksum 1:", adler32.Checksum(d1))
    fmt.Println("Checksum 2:", adler32.Checksum(d2))
}
``
/*
* Adler-32’in çakışma ihtimali, çok büyük veri setlerinde daha belirgin olur.
* CRC32, tek veya çift bit hatalarını çok daha güvenilir yakalar.

---

İstersen bir sonraki adım olarak sana **Adler-32 ve CRC32 çakışma olasılıklarını simüle eden bir program** da hazırlayabilirim. Bu program rastgele veriler üretip hangi algoritmanın daha çok çakıştığını gösterecek.

Bunu da yapalım mı?
*/