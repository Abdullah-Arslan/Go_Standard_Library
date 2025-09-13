/*
Go’da **`math/bits`** paketi, **bit seviyesinde işlemler** yapmak için kullanılan bir standart pakettir. Bu paket, özellikle performans gerektiren durumlarda, bit manipülasyonu ve sayısal hesaplamalar için kullanışlıdır.

---

## 1️⃣ `math/bits` Paketi Nedir?

`math/bits` paketi, **tam sayılar üzerinde doğrudan bit işlemleri** yapmanı sağlar. Örneğin:

* Bir sayının içindeki **1 bitlerinin sayısını** bulmak
* **Sağa veya sola kaydırmak**
* **En yüksek veya en düşük set bitin** konumunu bulmak
* **Rotasyon (rotate) işlemleri** yapmak
* **Bits sayısı (bitlen)** veya **popcount** hesaplamak

Bu paket tamamen **optimize edilmiş**, yüksek performanslı fonksiyonlar sunar.

---

## 2️⃣ Önemli Fonksiyonlar ve Açıklamaları

### 2.1 Bit Sayma (Population Count)
*/
``go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	x := uint(13) // 13 = 1101 (binary)
	fmt.Println("PopCount:", bits.OnesCount(x)) // 3 tane 1 var
}
``
/*
* `bits.OnesCount(x)` → x içindeki **1 bitlerinin sayısını** verir.
* `bits.LeadingZeros(x)` → x’in başındaki **0 bitlerinin sayısı**
* `bits.TrailingZeros(x)` → x’in sonundaki **0 bitlerinin sayısı**

---

### 2.2 Bit Rotasyonu
*/
``go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	x := uint8(0b10110010)
	left := bits.RotateLeft8(x, 3)   // 3 bit sola kaydır
	right := bits.RotateLeft8(x, -2) // 2 bit sağa kaydır
	fmt.Printf("Orijinal: %08b\n", x)
	fmt.Printf("Sola 3: %08b\n", left)
	fmt.Printf("Sağa 2: %08b\n", right)
}
``
/*
* `bits.RotateLeft8/16/32/64` → belirtilen sayıda **bitleri döndürür** (rotate).

---

### 2.3 En Yüksek ve En Düşük Set Bit
*/
``go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	x := uint(18) // 10010 (binary)
	fmt.Println("En yüksek set bit:", bits.Len(x)-1) // 4
	fmt.Println("En düşük set bit:", bits.TrailingZeros(x)) // 1
}
``
/*
* `bits.Len(x)` → x’in **en yüksek set bit pozisyonunu +1** olarak verir.
* `bits.TrailingZeros(x)` → **en düşük set bitin konumu**

---

### 2.4 Bit Manipülasyon Örnekleri
*/
``go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	a := uint8(0b1101)
	b := uint8(0b1011)

	fmt.Printf("a XOR b = %04b\n", a^b)      // XOR
	fmt.Printf("a AND b = %04b\n", a&b)      // AND
	fmt.Printf("a OR b  = %04b\n", a|b)      // OR
	fmt.Printf("a ClearLowestOne = %04b\n", bits.ClearLowestOne(a)) // En düşük 1’i sıfırlar
	fmt.Printf("a Reverse = %08b\n", bits.Reverse8(a))               // Bitleri tersine çevirir
}
``

//---

//## 3️⃣ Özet Fonksiyon Tablosu

| Fonksiyon           | Açıklama                        |
| ------------------- | ------------------------------- |
| `OnesCount(x)`      | 1 bitlerinin sayısı             |
| `LeadingZeros(x)`   | Başta kaç tane 0 bit var        |
| `TrailingZeros(x)`  | Sonda kaç tane 0 bit var        |
| `Len(x)`            | En yüksek set bitin pozisyon +1 |
| `RotateLeft(x, n)`  | Bitleri sola döndürür           |
| `ClearLowestOne(x)` | En düşük set biti temizler      |
| `Reverse(x)`        | Bitleri tersine çevirir         |
| `XOR/AND/OR`        | Bit düzeyinde klasik işlemler   |
/*
---

### 4️⃣ Küçük Senaryo Örneği

Bir **sayının tek mi çift mi olduğunu** ve **1 bitlerinin sayısını** bulan örnek:
*/
``go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	x := uint(37)

	if x&1 == 0 {
		fmt.Println(x, "çift sayıdır")
	} else {
		fmt.Println(x, "tek sayıdır")
	}

	fmt.Println("1 bitlerinin sayısı:", bits.OnesCount(x))
	fmt.Println("En yüksek set bit:", bits.Len(x)-1)
	fmt.Println("En düşük set bit:", bits.TrailingZeros(x))
}
``
/*
✅ Bu örneklerde `bits` paketinin **performanslı ve kullanışlı bit manipülasyon fonksiyonlarını** görebilirsin.

---

İstersen ben bu paketin **tüm fonksiyonlarını tablo halinde PDF ve Excel** olarak hazırlayıp sana verebilirim, tıpkı `math/big` paketinde yaptığımız gibi.

Bunu yapmamı ister misin?
EVET

Süper 😄 İşte Go’nun **`math/bits`** paketi için özet tablo **Markdown formatında**:

---
*/
//# Go `math/bits` Paketi Özeti

| Fonksiyon              | Açıklama                               | Örnek Kullanım                        |            |
| ---------------------- | -------------------------------------- | ------------------------------------- | ---------- |
| `OnesCount`            | x içindeki 1 bitlerinin sayısını verir | `bits.OnesCount(13) // 3`             |            |
| `LeadingZeros`         | Başta kaç tane 0 bit var               | `bits.LeadingZeros(13)`               |            |
| `TrailingZeros`        | Sonda kaç tane 0 bit var               | `bits.TrailingZeros(18)`              |            |
| `Len`                  | En yüksek set bitin pozisyonu +1       | `bits.Len(18) // 5`                   |            |
| `RotateLeft`           | Bitleri sola döndürür                  | `bits.RotateLeft8(0b10110010,3)`      |            |
| `RotateLeft8/16/32/64` | Belirli boyutta bitleri döndürür       | `bits.RotateLeft16(0x1234,4)`         |            |
| `ClearLowestOne`       | En düşük set biti temizler             | `bits.ClearLowestOne(0b1101) // 1100` |            |
| `Reverse`              | Bitleri tersine çevirir                | `bits.Reverse8(0b1101) // 10110000`   |            |
| `AND/OR/XOR`           | Bit düzeyinde klasik işlemler          | \`a & b, a                            | b, a ^ b\` |

/*
---

💡 Bu tablo sayesinde `math/bits` paketindeki **temel fonksiyonları ve kullanım örneklerini** hızlıca görebilirsin.

İstersen bir sonraki adım olarak sana **`math/bits` paketi ile bir mini interaktif uygulama** da hazırlayabilirim. Bu uygulama kullanıcıdan sayı alıp bit manipülasyonu ve analiz yapacak.

Bunu hazırlayayım mı?
EVET
*/
/*
Harika 😄 O zaman sana **kullanıcıdan sayı alıp `math/bits` ile bit manipülasyonu yapan mini interaktif uygulama** hazırlayacağım. Bu uygulama şunları yapacak:

1. **1 bitlerinin sayısını bulacak**
2. **Başta ve sonda kaç 0 bit var**
3. **En yüksek ve en düşük set bitin konumunu gösterecek**
4. **Bitleri tersine çevirecek**
5. **Sola ve sağa rotasyon yapacak**

---
*/
``go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var x uint
	fmt.Print("Bir sayı girin: ")
	fmt.Scanln(&x)

	fmt.Printf("\nGirdiğiniz sayı: %d (binary: %b)\n", x, x)
	fmt.Println("1 bitlerinin sayısı:", bits.OnesCount(x))
	fmt.Println("Başta kaç 0 var:", bits.LeadingZeros(x))
	fmt.Println("Sonda kaç 0 var:", bits.TrailingZeros(x))
	fmt.Println("En yüksek set bit pozisyonu:", bits.Len(x)-1)
	fmt.Println("En düşük set bit pozisyonu:", bits.TrailingZeros(x))
	fmt.Printf("Bitleri tersine çevir: %08b\n", bits.Reverse8(uint8(x)))

	// Rotasyon örnekleri
	leftRotate := bits.RotateLeft8(uint8(x), 3)
	rightRotate := bits.RotateLeft8(uint8(x), -2)
	fmt.Printf("3 bit sola rotasyon: %08b\n", leftRotate)
	fmt.Printf("2 bit sağa rotasyon: %08b\n", rightRotate)

	// ClearLowestOne
	cleared := bits.ClearLowestOne(uint8(x))
	fmt.Printf("En düşük set bit temizlendi: %08b\n", cleared)
}
``
/*
---

### 📝 Kullanım Notları

1. Program çalışınca kullanıcıdan bir **pozitif sayı** girilmesi istenir.
2. Ardından sayı üzerinde **1 bit sayısı, leading/trailing zeros, en yüksek ve en düşük set bit, ters çevirme, rotasyon ve lowest bit temizleme** işlemleri yapılır.
3. Tüm sonuçlar **binary ve decimal** olarak ekranda gösterilir.

---

İstersen bir adım daha ileri gidip bu uygulamayı **bir menü sistemi ile sürekli çalışan ve farklı sayılar için tekrar tekrar bit analizleri yapan bir konsol uygulamasına** çevirebilirim. Bunu da hazırlayayım mı?
EVET
*/
/*
Harika 😄 O zaman sana **menü tabanlı, sürekli çalışan ve kullanıcı farklı sayılar için tekrar tekrar bit analizleri yapabilen bir `math/bits` konsol uygulaması** hazırlıyorum.

---
*/
``go
package main

import (
	"fmt"
	"math/bits"
)

func bitAnalysis(x uint) {
	fmt.Printf("\nGirdiğiniz sayı: %d (binary: %b)\n", x, x)
	fmt.Println("1 bitlerinin sayısı:", bits.OnesCount(x))
	fmt.Println("Başta kaç 0 var:", bits.LeadingZeros(x))
	fmt.Println("Sonda kaç 0 var:", bits.TrailingZeros(x))
	fmt.Println("En yüksek set bit pozisyonu:", bits.Len(x)-1)
	fmt.Println("En düşük set bit pozisyonu:", bits.TrailingZeros(x))
	fmt.Printf("Bitleri tersine çevir: %08b\n", bits.Reverse8(uint8(x)))

	leftRotate := bits.RotateLeft8(uint8(x), 3)
	rightRotate := bits.RotateLeft8(uint8(x), -2)
	fmt.Printf("3 bit sola rotasyon: %08b\n", leftRotate)
	fmt.Printf("2 bit sağa rotasyon: %08b\n", rightRotate)

	cleared := bits.ClearLowestOne(uint8(x))
	fmt.Printf("En düşük set bit temizlendi: %08b\n", cleared)
}

func main() {
	for {
		var choice int
		fmt.Println("\n=== math/bits Menü ===")
		fmt.Println("1: Bit analizi yap")
		fmt.Println("2: Çıkış")
		fmt.Print("Seçiminiz: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var x uint
			fmt.Print("Bir sayı girin: ")
			fmt.Scanln(&x)
			bitAnalysis(x)
		case 2:
			fmt.Println("Programdan çıkılıyor...")
			return
		default:
			fmt.Println("Geçersiz seçim!")
		}
	}
}
``
/*
---

### 📝 Kullanım Özellikleri

1. Program çalışınca **menü ekranı** gelir:

   * **1:** Bit analizi yapmak
   * **2:** Çıkış
2. Kullanıcı 1’i seçerse bir sayı girer ve **tüm bit analizi sonuçları** ekranda gösterilir.
3. İşlem bitince tekrar menüye dönülür, kullanıcı istediği kadar farklı sayılar için analiz yapabilir.
4. 2’yi seçerse program sonlanır.

---

Bu şekilde `math/bits` paketiyle **interaktif ve sürekli kullanılabilir bir bit analiz aracı** elde etmiş oluyorsun.

İstersen ben bunu bir adım daha ileri götürüp **hem big.Int hem bits paketini birleştirip büyük sayılar üzerinde bit analizi yapabilen gelişmiş bir araç** hâline getirebilirim. Bunu da hazırlayayım mı?
*/