/*
Go’da **`math/big`** paketi, çok büyük sayılarla çalışmaya yarar. Normalde Go’nun `int`, `uint`, `float64` gibi veri tiplerinin sınırları vardır (örneğin `int64` en fazla 2⁶³-1’e kadar değer alır). Ancak bazı durumlarda (kripto, bilimsel hesaplama, finans, faktöriyel, kombinasyon gibi konularda) çok daha büyük sayılarla çalışmak gerekir.

İşte **`math/big`** paketi bu işi çözer.

Bu paket ile:

* **Büyük tamsayılar (`big.Int`)**
* **Büyük rasyonel sayılar (`big.Rat`)**
* **Büyük kayan noktalı sayılar (`big.Float`)**

üzerinde işlem yapabilirsin.

---

## 1. `big.Int` (Büyük Tamsayılar)

`big.Int` ile istediğin büyüklükte tam sayılarla işlem yapabilirsin.
*/
``go
package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Büyük sayı tanımlama
	a := big.NewInt(12345678901234567890)
	b := big.NewInt(98765432109876543210)

	// Toplama
	sum := new(big.Int).Add(a, b)
	fmt.Println("Toplam:", sum)

	// Çarpma
	mul := new(big.Int).Mul(a, b)
	fmt.Println("Çarpım:", mul)

	// Bölme (tam sayı bölme)
	div := new(big.Int).Div(b, a)
	fmt.Println("Bölüm:", div)

	// Üs alma (a^5)
	exp := new(big.Int).Exp(a, big.NewInt(5), nil)
	fmt.Println("a^5:", exp)
}
``
/*
✅ Çıktıda çok büyük sayılar görürsün, çünkü sınır yoktur.

---

## 2. `big.Rat` (Rasyonel Sayılar)

Kesirli sayıları tam olarak saklamak için `big.Rat` kullanılır.
`float64` gibi kayma hatası (precision error) yaşamazsın.
*7
``go
package main

import (
	"fmt"
	"math/big"
)

func main() {
	// 1/3 sayısını temsil et
	r1 := big.NewRat(1, 3)
	// 2/5 sayısını temsil et
	r2 := big.NewRat(2, 5)

	// Toplama
	sum := new(big.Rat).Add(r1, r2)
	fmt.Println("1/3 + 2/5 =", sum)

	// Çarpma
	mul := new(big.Rat).Mul(r1, r2)
	fmt.Println("1/3 * 2/5 =", mul)

	// Bölme
	div := new(big.Rat).Quo(r1, r2)
	fmt.Println("(1/3) / (2/5) =", div)
}
``

//✅ Çıktı tam kesir şeklinde olur:

``
1/3 + 2/5 = 11/15
1/3 * 2/5 = 2/15
(1/3) / (2/5) = 5/6
``
/*
---

## 3. `big.Float` (Büyük Ondalıklı Sayılar)

Yüksek hassasiyetli ondalıklı sayılar için `big.Float` kullanılır.
Burada **precision (hassasiyet)** ayarlayabilirsin.
*/
``go
package main

import (
	"fmt"
	"math/big"
)

func main() {
	// 1.5 değeri
	f1 := big.NewFloat(1.5)
	// 2.75 değeri
	f2 := big.NewFloat(2.75)

	// Toplama
	sum := new(big.Float).Add(f1, f2)
	fmt.Println("Toplam:", sum)

	// Çarpma
	mul := new(big.Float).Mul(f1, f2)
	fmt.Println("Çarpım:", mul)

	// Bölme
	div := new(big.Float).Quo(f2, f1)
	fmt.Println("Bölüm:", div)

	// Üs alma için float kullanmak yerine int tabanlı yapmalısın
}
``

/7✅ Çıktı:

``
Toplam: 4.25
Çarpım: 4.125
Bölüm: 1.8333333333333333
``
/*
---

## 4. `big.Int` ile Faktöriyel Örneği

Büyük sayıların klasik kullanım alanı faktöriyel hesaplamadır.
*/
``go
package main

import (
	"fmt"
	"math/big"
)

func factorial(n int64) *big.Int {
	result := big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		result.Mul(result, big.NewInt(i))
	}
	return result
}

func main() {
	fmt.Println("50! =", factorial(50))
}
``
/*
✅ 50 faktöriyel gibi devasa bir sayı üretir.

---

## Özet:

* `big.Int` → sınırsız büyüklükte **tam sayı**
* `big.Rat` → kesin değerli **kesirli sayı**
* `big.Float` → yüksek hassasiyetli **ondalıklı sayı**

Bunların hepsi `math/big` paketindedir ve `Add`, `Sub`, `Mul`, `Div`, `Exp`, `GCD` gibi klasik matematiksel fonksiyonları destekler.

---

İstersen ben sana bu paketteki **tüm fonksiyonların detaylı bir tablo** şeklinde açıklamasını çıkarabilirim. Bunu PDF veya Excel formatında hazırlamamı ister misin?
EVET
*/
/*
Süper 👍 O zaman sana hazırladığım tabloyu **Markdown** formatında da veriyorum.

---

# Go `math/big` Paketi Özeti
*/

| Tip       | Fonksiyon / Metot | Açıklama                              | Örnek Kullanım                                |
| --------- | ----------------- | ------------------------------------- | --------------------------------------------- |
| **Int**   | `NewInt`          | Yeni bir Int nesnesi oluşturur        | `a := big.NewInt(123)`                        |
| Int       | `Add`             | Toplama işlemi yapar                  | `c := new(big.Int).Add(a, b)`                 |
| Int       | `Sub`             | Çıkarma işlemi yapar                  | `c := new(big.Int).Sub(a, b)`                 |
| Int       | `Mul`             | Çarpma işlemi yapar                   | `c := new(big.Int).Mul(a, b)`                 |
| Int       | `Div`             | Bölme işlemi yapar (tam sayı)         | `c := new(big.Int).Div(a, b)`                 |
| Int       | `Mod`             | Mod alma işlemi                       | `c := new(big.Int).Mod(a, b)`                 |
| Int       | `Exp`             | Üs alma işlemi (a^b)                  | `c := new(big.Int).Exp(a, b, nil)`            |
| Int       | `GCD`             | En büyük ortak bölen                  | `g := new(big.Int).GCD(nil, nil, a, b)`       |
| Int       | `Cmp`             | Karşılaştırma (-1, 0, 1)              | `res := a.Cmp(b)`                             |
| Int       | `Abs`             | Mutlak değer                          | `c := new(big.Int).Abs(a)`                    |
| Int       | `Neg`             | Negatifini alır                       | `c := new(big.Int).Neg(a)`                    |
| Int       | `SetString`       | String’den sayı oluşturur             | `a, _ := new(big.Int).SetString("12345", 10)` |
| **Rat**   | `NewRat`          | Yeni bir Rat (kesirli sayı) oluşturur | `r := big.NewRat(1, 3)`                       |
| Rat       | `Add`             | Kesir toplama                         | `c := new(big.Rat).Add(r1, r2)`               |
| Rat       | `Sub`             | Kesir çıkarma                         | `c := new(big.Rat).Sub(r1, r2)`               |
| Rat       | `Mul`             | Kesir çarpma                          | `c := new(big.Rat).Mul(r1, r2)`               |
| Rat       | `Quo`             | Kesir bölme                           | `c := new(big.Rat).Quo(r1, r2)`               |
| Rat       | `Inv`             | Tersini alır (1/x)                    | `c := new(big.Rat).Inv(r1)`                   |
| Rat       | `Cmp`             | Karşılaştırma (-1,0,1)                | `res := r1.Cmp(r2)`                           |
| **Float** | `NewFloat`        | Yeni bir Float oluşturur              | `f := big.NewFloat(1.23)`                     |
| Float     | `SetPrec`         | Hassasiyeti ayarlar (bit cinsinden)   | `f.SetPrec(200)`                              |
| Float     | `Add`             | Toplama işlemi                        | `c := new(big.Float).Add(f1, f2)`             |
| Float     | `Sub`             | Çıkarma işlemi                        | `c := new(big.Float).Sub(f1, f2)`             |
| Float     | `Mul`             | Çarpma işlemi                         | `c := new(big.Float).Mul(f1, f2)`             |
| Float     | `Quo`             | Bölme işlemi                          | `c := new(big.Float).Quo(f1, f2)`             |
| Float     | `Cmp`             | Karşılaştırma (-1,0,1)                | `res := f1.Cmp(f2)`                           |
| Float     | `Text`            | String’e çevirme                      | `s := f.Text('f', 10)`                        |

/*
---

👉 İstersen ben sana bu fonksiyonların **hepsi için çalışır küçük örnek programlar** da yazabilirim. Böylece her fonksiyonu canlı test edebilirsin. İstiyor musun?
EVET
*/
/*
Harika 😄 O zaman ben sana **`math/big` paketindeki tüm tipler (`Int`, `Rat`, `Float`) için küçük çalışır örnek programlar** hazırlayacağım. Her örnek, ilgili fonksiyonları gösterip kolayca çalıştırabileceğin şekilde olacak.

---

## 1️⃣ `big.Int` Örnekleri
*/
``go
package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1234567890)
	b := big.NewInt(9876543210)

	// Toplama
	sum := new(big.Int).Add(a, b)
	fmt.Println("Toplam:", sum)

	// Çıkarma
	sub := new(big.Int).Sub(b, a)
	fmt.Println("Çıkarma:", sub)

	// Çarpma
	mul := new(big.Int).Mul(a, b)
	fmt.Println("Çarpma:", mul)

	// Bölme
	div := new(big.Int).Div(b, a)
	fmt.Println("Bölme:", div)

	// Mod
	mod := new(big.Int).Mod(b, a)
	fmt.Println("Mod:", mod)

	// Üs alma
	exp := new(big.Int).Exp(a, big.NewInt(5), nil)
	fmt.Println("a^5:", exp)

	// Mutlak değer ve negatif alma
	neg := new(big.Int).Neg(a)
	fmt.Println("Negatif:", neg)
	abs := new(big.Int).Abs(neg)
	fmt.Println("Mutlak değer:", abs)

	// Karşılaştırma
	fmt.Println("a == b ?", a.Cmp(b) == 0)
	fmt.Println("a < b ?", a.Cmp(b) < 0)
	fmt.Println("a > b ?", a.Cmp(b) > 0)
}
``
/*
---

## 2️⃣ `big.Rat` Örnekleri
*/
``go
package main

import (
	"fmt"
	"math/big"
)

func main() {
	r1 := big.NewRat(1, 3)
	r2 := big.NewRat(2, 5)

	// Toplama
	sum := new(big.Rat).Add(r1, r2)
	fmt.Println("1/3 + 2/5 =", sum)

	// Çıkarma
	sub := new(big.Rat).Sub(r1, r2)
	fmt.Println("1/3 - 2/5 =", sub)

	// Çarpma
	mul := new(big.Rat).Mul(r1, r2)
	fmt.Println("1/3 * 2/5 =", mul)

	// Bölme
	div := new(big.Rat).Quo(r1, r2)
	fmt.Println("(1/3) / (2/5) =", div)

	// Tersini alma
	inv := new(big.Rat).Inv(r1)
	fmt.Println("1 / (1/3) =", inv)

	// Karşılaştırma
	fmt.Println("r1 == r2 ?", r1.Cmp(r2) == 0)
	fmt.Println("r1 < r2 ?", r1.Cmp(r2) < 0)
	fmt.Println("r1 > r2 ?", r1.Cmp(r2) > 0)
}
``
/*
---

## 3️⃣ `big.Float` Örnekleri
*/
``go
package main

import (
	"fmt"
	"math/big"
)

func main() {
	f1 := big.NewFloat(1.23456789)
	f2 := big.NewFloat(2.34567891)

	// Hassasiyeti arttır
	f1.SetPrec(200)
	f2.SetPrec(200)

	// Toplama
	sum := new(big.Float).Add(f1, f2)
	fmt.Println("Toplam:", sum)

	// Çıkarma
	sub := new(big.Float).Sub(f2, f1)
	fmt.Println("Çıkarma:", sub)

	// Çarpma
	mul := new(big.Float).Mul(f1, f2)
	fmt.Println("Çarpma:", mul)

	// Bölme
	div := new(big.Float).Quo(f2, f1)
	fmt.Println("Bölme:", div)

	// String'e çevirme
	fmt.Println("f1:", f1.Text('f', 20))
	fmt.Println("f2:", f2.Text('f', 20))

	// Karşılaştırma
	fmt.Println("f1 == f2 ?", f1.Cmp(f2) == 0)
	fmt.Println("f1 < f2 ?", f1.Cmp(f2) < 0)
	fmt.Println("f1 > f2 ?", f1.Cmp(f2) > 0)
}
``
/*
---

💡 Bu üç örnek seti sayesinde `math/big` paketindeki **ana tiplerin ve temel fonksiyonların** kullanımını tamamen görebilirsin.

İstersen ben bir sonraki adım olarak sana **bu üç tip için bir “faktöriyel, büyük kesir ve yüksek hassasiyetli float hesaplama” mini uygulaması** da hazırlayabilirim. Bu şekilde gerçek bir senaryoda nasıl kullanılacağını görebilirsin.

Bunu hazırlayayım mı?
EVET
*/
/*
Harika 😄 O zaman sana **`math/big` paketini gerçek bir senaryoda kullanan mini uygulama** hazırlayacağım. Bu uygulama üç tip için örnek olacak:

1. **Faktöriyel hesaplama (`big.Int`)**
2. **Büyük kesir toplama ve çarpma (`big.Rat`)**
3. **Yüksek hassasiyetli float hesaplama (`big.Float`)**

---
*/
``go
package main

import (
	"fmt"
	"math/big"
)

// Faktöriyel hesaplama
func factorial(n int64) *big.Int {
	result := big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		result.Mul(result, big.NewInt(i))
	}
	return result
}

func main() {
	// ==========================
	// 1️⃣ Faktöriyel (big.Int)
	// ==========================
	n := int64(30)
	fmt.Printf("%d! = %v\n\n", n, factorial(n))

	// ==========================
	// 2️⃣ Büyük Kesirler (big.Rat)
	// ==========================
	r1 := big.NewRat(123456789, 987654321)
	r2 := big.NewRat(987654321, 123456789)

	sumRat := new(big.Rat).Add(r1, r2)
	mulRat := new(big.Rat).Mul(r1, r2)
	fmt.Println("r1 + r2 =", sumRat)
	fmt.Println("r1 * r2 =", mulRat)
	fmt.Println()

	// ==========================
	// 3️⃣ Yüksek Hassasiyetli Float (big.Float)
	// ==========================
	f1 := big.NewFloat(1.123456789123456789)
	f2 := big.NewFloat(2.987654321987654321)

	// Hassasiyeti 200 bit yap
	f1.SetPrec(200)
	f2.SetPrec(200)

	sumFloat := new(big.Float).Add(f1, f2)
	mulFloat := new(big.Float).Mul(f1, f2)

	fmt.Println("f1 + f2 =", sumFloat.Text('f', 30))
	fmt.Println("f1 * f2 =", mulFloat.Text('f', 30))
}
``
/*
---

### ✅ Uygulamanın Çalışması

* **Faktöriyel**: Çok büyük bir tamsayı üretir (30! gibi)
* **Kesirli Sayılar**: Kesirleri tam olarak toplar ve çarpar, kayan nokta hatası olmaz
* **Float**: Çok yüksek hassasiyetle ondalıklı sayılarla işlem yapar

---

İstersen ben bunu bir adım ileri götürüp **kullanıcıdan sayı alıp faktöriyel, kesir ve float hesaplama yapabilen interaktif bir mini konsol uygulaması** hâline de getirebilirim.

Bunu da hazırlayayım mı?
EVET
*/
/*
Harika 😄 O zaman sana **kullanıcıdan veri alıp interaktif olarak `math/big` ile hesaplama yapan mini uygulama** hazırlıyorum. Bu uygulama üç bölümü kapsayacak:

1. **Faktöriyel (`big.Int`)**
2. **Kesir toplama ve çarpma (`big.Rat`)**
3. **Hassas float toplama ve çarpma (`big.Float`)**

---
*/
``go
package main

import (
	"fmt"
	"math/big"
)

// Faktöriyel hesaplama
func factorial(n int64) *big.Int {
	result := big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		result.Mul(result, big.NewInt(i))
	}
	return result
}

func main() {
	var choice int
	fmt.Println("Math/Big Mini Uygulaması")
	fmt.Println("1: Faktöriyel Hesaplama (big.Int)")
	fmt.Println("2: Kesir İşlemleri (big.Rat)")
	fmt.Println("3: Hassas Float İşlemleri (big.Float)")
	fmt.Print("Seçiminiz (1-3): ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		var n int64
		fmt.Print("Faktöriyel için bir sayı girin: ")
		fmt.Scanln(&n)
		fmt.Printf("%d! = %v\n", n, factorial(n))
	case 2:
		var num1, den1, num2, den2 int64
		fmt.Println("Birinci kesri girin:")
		fmt.Print("Pay: "); fmt.Scanln(&num1)
		fmt.Print("Payda: "); fmt.Scanln(&den1)
		fmt.Println("İkinci kesri girin:")
		fmt.Print("Pay: "); fmt.Scanln(&num2)
		fmt.Print("Payda: "); fmt.Scanln(&den2)

		r1 := big.NewRat(num1, den1)
		r2 := big.NewRat(num2, den2)

		sum := new(big.Rat).Add(r1, r2)
		mul := new(big.Rat).Mul(r1, r2)

		fmt.Println("Toplam:", sum)
		fmt.Println("Çarpım:", mul)
	case 3:
		var f1Str, f2Str string
		fmt.Print("Birinci float sayıyı girin: ")
		fmt.Scanln(&f1Str)
		fmt.Print("İkinci float sayıyı girin: ")
		fmt.Scanln(&f2Str)

		f1, _, err1 := big.ParseFloat(f1Str, 10, 200, big.ToNearestEven)
		f2, _, err2 := big.ParseFloat(f2Str, 10, 200, big.ToNearestEven)

		if err1 != nil || err2 != nil {
			fmt.Println("Hatalı giriş!")
			return
		}

		sum := new(big.Float).Add(f1, f2)
		mul := new(big.Float).Mul(f1, f2)

		fmt.Println("Toplam:", sum.Text('f', 30))
		fmt.Println("Çarpım:", mul.Text('f', 30))
	default:
		fmt.Println("Geçersiz seçim!")
	}
}
``
/*
---

### 📝 Kullanım Notları

1. Program çalışınca kullanıcıya **1-3 arası seçim** sorulur.
2. **1:** Faktöriyel için tamsayı girişi alınır ve sonucu gösterir.
3. **2:** İki kesir için pay ve payda girilir, toplam ve çarpım hesaplanır.
4. **3:** İki ondalıklı sayı string olarak alınır, yüksek hassasiyetle toplama ve çarpma yapılır.

---

İstersen ben bunu bir adım daha ileri götürüp **tüm işlemleri aynı anda çalıştırıp sonuçları tek seferde gösteren tam bir “math/big hesaplama paneli”** hâline getirebilirim. Bunu da hazırlayayım mı?
EVET
*/
/*
Harika 😄 O zaman sana **tüm işlemleri aynı anda yapan, interaktif bir “math/big hesaplama paneli”** hazırlayacağım. Kullanıcı bir kere sayıları giriyor, program hem faktöriyel, hem kesir işlemleri, hem de hassas float hesaplamalarını tek seferde gösteriyor.

---
*/
``go
package main

import (
	"fmt"
	"math/big"
)

// Faktöriyel hesaplama
func factorial(n int64) *big.Int {
	result := big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		result.Mul(result, big.NewInt(i))
	}
	return result
}

func main() {
	fmt.Println("=== Math/Big Hesaplama Paneli ===")

	// 1️⃣ Faktöriyel
	var n int64
	fmt.Print("Faktöriyel için bir sayı girin: ")
	fmt.Scanln(&n)
	fact := factorial(n)
	fmt.Printf("%d! = %v\n\n", n, fact)

	// 2️⃣ Kesir İşlemleri
	var num1, den1, num2, den2 int64
	fmt.Println("Birinci kesri girin:")
	fmt.Print("Pay: "); fmt.Scanln(&num1)
	fmt.Print("Payda: "); fmt.Scanln(&den1)
	fmt.Println("İkinci kesri girin:")
	fmt.Print("Pay: "); fmt.Scanln(&num2)
	fmt.Print("Payda: "); fmt.Scanln(&den2)

	r1 := big.NewRat(num1, den1)
	r2 := big.NewRat(num2, den2)

	sumRat := new(big.Rat).Add(r1, r2)
	mulRat := new(big.Rat).Mul(r1, r2)

	fmt.Println("Kesir Toplamı:", sumRat)
	fmt.Println("Kesir Çarpımı:", mulRat)
	fmt.Println()

	// 3️⃣ Hassas Float İşlemleri
	var f1Str, f2Str string
	fmt.Print("Birinci float sayıyı girin: ")
	fmt.Scanln(&f1Str)
	fmt.Print("İkinci float sayıyı girin: ")
	fmt.Scanln(&f2Str)

	f1, _, err1 := big.ParseFloat(f1Str, 10, 200, big.ToNearestEven)
	f2, _, err2 := big.ParseFloat(f2Str, 10, 200, big.ToNearestEven)

	if err1 != nil || err2 != nil {
		fmt.Println("Hatalı float giriş!")
		return
	}

	sumFloat := new(big.Float).Add(f1, f2)
	mulFloat := new(big.Float).Mul(f1, f2)

	fmt.Println("Float Toplam:", sumFloat.Text('f', 30))
	fmt.Println("Float Çarpım:", mulFloat.Text('f', 30))

	fmt.Println("\n=== Hesaplama Tamamlandı ===")
}
``
/*
---

### 📝 Kullanım Notları

1. Kullanıcı önce **faktöriyel** için bir sayı girer.
2. Ardından iki kesir girer (**pay ve payda**).
3. Son olarak iki **ondalıklı sayı** girer.
4. Program tüm sonuçları **tek seferde** gösterir:

   * Faktöriyel sonucu
   * Kesirlerin toplamı ve çarpımı
   * Float sayıların yüksek hassasiyetli toplamı ve çarpımı

---

Bu panel sayesinde `math/big` paketinin tüm tiplerini **gerçek bir senaryoda aynı anda kullanabilirsin**.

İstersen bir sonraki adım olarak bunu **menü tabanlı ve tekrar işlem yapabilen sürekli çalışan bir konsol uygulaması** hâline de getirebilirim. Bunu da hazırlayayım mı?

*/