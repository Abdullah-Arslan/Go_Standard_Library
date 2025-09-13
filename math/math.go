/*
Go’nun **`math` paketi**, matematiksel işlemler için kullanılan temel bir standart kütüphanedir. İçerisinde trigonometrik fonksiyonlardan logaritmalara, üs alma işlemlerinden mutlak değer hesaplamaya kadar birçok hazır fonksiyon ve sabit bulunur.

Şimdi sana bu paketi **tam kapsamlı** şekilde açıklayayım.

---

## 1. Paket Tanımı
*/
``go
import "math"
``
/*
Bu paket, **float64** tipinde sayılar üzerinde çalışır. Yani çoğu fonksiyonun parametreleri `float64`, dönüş tipleri de `float64`’tür. (Bazı istisnalar var: `math.MaxInt64` gibi sabitler.)

---

## 2. Önemli Sabitler (Constants)

`math` paketinde birçok matematiksel sabit tanımlıdır:

* **`math.Pi`** → π sayısı (3.14159...)
* **`math.E`** → Euler sayısı (2.71828...)
* **`math.Phi`** → Altın oran (1.618...)
* **`math.Sqrt2`** → √2
* **`math.Ln2`** → ln(2)
* **`math.Ln10`** → ln(10)
* **`math.MaxFloat64`**, **`math.SmallestNonzeroFloat64`** → float64 için limit değerler
* **`math.MaxInt64`**, **`math.MinInt64`** → integer sınırları

### Örnek:
*/
``go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println("Pi:", math.Pi)
    fmt.Println("Euler:", math.E)
    fmt.Println("Altın Oran:", math.Phi)
    fmt.Println("Sqrt(2):", math.Sqrt2)
}
``
/*
---

## 3. Temel Fonksiyonlar

### ✅ Mutlak Değer ve İşaret

* `math.Abs(x)` → |x|
* `math.Copysign(x, y)` → x’in mutlak değerini alıp, y’nin işaretini verir.
* `math.Signbit(x)` → sayı negatifse `true`, değilse `false`.
*/
``go
fmt.Println(math.Abs(-5))        // 5
fmt.Println(math.Copysign(3, -2)) // -3
fmt.Println(math.Signbit(-10))    // true
``
/*
---

### ✅ Karekök ve Kuvvet

* `math.Sqrt(x)` → √x
* `math.Pow(x, y)` → xʸ
* `math.Pow10(n)` → 10ⁿ
*/
``go
fmt.Println(math.Sqrt(16))   // 4
fmt.Println(math.Pow(2, 3))  // 8
fmt.Println(math.Pow10(3))   // 1000
``
/*
---

### ✅ Logaritma ve Üs

* `math.Exp(x)` → eˣ
* `math.Log(x)` → ln(x)
* `math.Log10(x)` → log₁₀(x)
* `math.Log2(x)` → log₂(x)
*/
``go
fmt.Println(math.Exp(1))     // 2.718...
fmt.Println(math.Log(10))    // 2.302...
fmt.Println(math.Log10(100)) // 2
fmt.Println(math.Log2(8))    // 3
``
/*
---

### ✅ Trigonometrik Fonksiyonlar

* `math.Sin(x)`, `math.Cos(x)`, `math.Tan(x)`
* `math.Asin(x)`, `math.Acos(x)`, `math.Atan(x)`
* `math.Atan2(y, x)` → koordinat açısı

👉 **Not:** Bu fonksiyonlar radyan cinsinden çalışır.
(Derece → radyan: `radian = derece * math.Pi / 180`)
*/

``go
rad := 90 * math.Pi / 180
fmt.Println(math.Sin(rad)) // 1
fmt.Println(math.Cos(rad)) // 0
``
/*
---

### ✅ Yuvarlama Fonksiyonları

* `math.Floor(x)` → aşağı yuvarla
* `math.Ceil(x)` → yukarı yuvarla
* `math.Round(x)` → en yakın tam sayıya yuvarla
* `math.Trunc(x)` → tam kısmı alır
* `math.Mod(x, y)` → x % y (float mod)
*/

``go
fmt.Println(math.Floor(3.7)) // 3
fmt.Println(math.Ceil(3.2))  // 4
fmt.Println(math.Round(3.5)) // 4
fmt.Println(math.Trunc(3.9)) // 3
fmt.Println(math.Mod(10, 3)) // 1
``
/*
---

### ✅ Min / Max / Hipotenüs

* `math.Min(x, y)`
* `math.Max(x, y)`
* `math.Hypot(x, y)` → √(x²+y²)
*/

``go
fmt.Println(math.Min(2, 5))     // 2
fmt.Println(math.Max(2, 5))     // 5
fmt.Println(math.Hypot(3, 4))   // 5 (3-4-5 üçgeni)
``
/*
---

### ✅ Özel Fonksiyonlar

* `math.IsNaN(x)` → sayı NaN mı?
* `math.IsInf(x, sign)` → sayı ∞ mi? (sign: 1=+∞, -1=-∞, 0=her ikisi)
* `math.Inf(sign)` → ∞ üretir
*7
``go
x := math.Log(-1) // NaN
fmt.Println(math.IsNaN(x)) // true

y := math.Inf(1) 
fmt.Println(y)              // +Inf
fmt.Println(math.IsInf(y, 1)) // true
``
/*
---

## 4. Kullanım Senaryoları

1. **Fizik/Mühendislik hesaplamaları**

   * Sinüs, kosinüs, logaritma gibi bilimsel hesaplamalar.
2. **Finansal hesaplamalar**

   * Yuvarlama, faiz hesaplamaları, logaritmalar.
3. **Grafik ve oyun geliştirme**

   * Koordinat dönüşümleri, açı hesaplama (`Atan2`).
4. **Veri bilimi**

   * Normalizasyon, istatistiksel formüller.

---

## 5. Küçük Bir Uygulama Örneği
*7
``go
package main

import (
    "fmt"
    "math"
)

func main() {
    // Bir cismin eğimli atış açısında maksimum yüksekliği bulalım
    hız := 20.0  // m/s
    açı := 45.0 * math.Pi / 180 // dereceyi radyana çevirdik
    g := 9.81 // yerçekimi

    // Maksimum yükseklik formülü: (v^2 * sin^2(θ)) / (2*g)
    h := (math.Pow(hız*math.Sin(açı), 2)) / (2 * g)

    fmt.Printf("Maksimum Yükseklik: %.2f metre\n", h)
}
``
/*
---

✅ Özet:
`math` paketi, Go’nun **tüm temel matematiksel işlemlerini** sağlayan güçlü bir kütüphanedir. Sayısal hesaplamalar, bilimsel uygulamalar, mühendislik hesaplamaları, oyun ve grafik programlamada sıkça kullanılır.

---

İstersen sana `math` paketinin **tam fonksiyon listesini tablo halinde** de çıkarabilirim. Bunu ister misin?
EVET
*/
/*
Süper 👍 O zaman sana **Go `math` paketinin tamamını fonksiyon ve sabitleriyle tablo halinde** sunuyorum. Böylece hepsini tek bakışta görebilirsin.

---

# 📌 Go `math` Paketi – Fonksiyonlar & Sabitler Tablosu
*/

//## 🔹 1. Matematiksel Sabitler

| Sabit                         | Açıklama               | Değer                |
| ----------------------------- | ---------------------- | -------------------- |
| `math.Pi`                     | π sayısı               | 3.14159...           |
| `math.E`                      | Euler sayısı           | 2.71828...           |
| `math.Phi`                    | Altın oran             | 1.61803...           |
| `math.Sqrt2`                  | √2                     | 1.41421...           |
| `math.SqrtE`                  | √e                     | 1.64872...           |
| `math.SqrtPi`                 | √π                     | 1.77245...           |
| `math.Ln2`                    | ln(2)                  | 0.6931...            |
| `math.Ln10`                   | ln(10)                 | 2.3025...            |
| `math.Log2E`                  | log₂(e)                | 1.4427...            |
| `math.Log10E`                 | log₁₀(e)               | 0.43429...           |
| `math.MaxFloat64`             | float64 max değeri     | \~1.8e308            |
| `math.SmallestNonzeroFloat64` | float64 en küçük değer | \~5e-324             |
| `math.MaxInt64`               | int64 max              | 9223372036854775807  |
| `math.MinInt64`               | int64 min              | -9223372036854775808 |

---

## 🔹 2. Temel İşlemler

| Fonksiyon        | Açıklama                           | Örnek                     |
| ---------------- | ---------------------------------- | ------------------------- |
| `Abs(x)`         | Mutlak değer                       | `Abs(-5) → 5`             |
| `Copysign(x, y)` | x’in mutlak değeri + y’nin işareti | `Copysign(3, -2) → -3`    |
| `Dim(x, y)`      | (x-y) veya 0                       | `Dim(5, 3) → 2`           |
| `Signbit(x)`     | Negatif mi? (bool)                 | `Signbit(-2) → true`      |
| `IsNaN(x)`       | NaN kontrolü                       | `IsNaN(Log(-1)) → true`   |
| `IsInf(x, sign)` | Sonsuz mu?                         | `IsInf(Inf(1), 1) → true` |
| `NaN()`          | NaN döndürür                       | `NaN()`                   |
| `Inf(sign)`      | ±∞ üretir                          | `Inf(-1) → -Inf`          |

---

## 🔹 3. Karekök & Kuvvet

| Fonksiyon     | Açıklama | Örnek             |
| ------------- | -------- | ----------------- |
| `Sqrt(x)`     | √x       | `Sqrt(16) → 4`    |
| `Pow(x, y)`   | xʸ       | `Pow(2, 3) → 8`   |
| `Pow10(n)`    | 10ⁿ      | `Pow10(3) → 1000` |
| `Cbrt(x)`     | Küp kök  | `Cbrt(27) → 3`    |
| `Hypot(x, y)` | √(x²+y²) | `Hypot(3, 4) → 5` |

---

## 🔹 4. Logaritma & Üs Fonksiyonları

| Fonksiyon  | Açıklama | Örnek                   |
| ---------- | -------- | ----------------------- |
| `Exp(x)`   | eˣ       | `Exp(1) → 2.718...`     |
| `Expm1(x)` | eˣ - 1   | `Expm1(1) → 1.718...`   |
| `Log(x)`   | ln(x)    | `Log(10) → 2.302...`    |
| `Log10(x)` | log₁₀(x) | `Log10(100) → 2`        |
| `Log2(x)`  | log₂(x)  | `Log2(8) → 3`           |
| `Log1p(x)` | ln(1+x)  | `Log1p(0.5) → 0.405...` |

---

## 🔹 5. Trigonometrik Fonksiyonlar

| Fonksiyon     | Açıklama        | Örnek               |
| ------------- | --------------- | ------------------- |
| `Sin(x)`      | sin(x)          | `Sin(Pi/2) → 1`     |
| `Cos(x)`      | cos(x)          | `Cos(Pi) → -1`      |
| `Tan(x)`      | tan(x)          | `Tan(Pi/4) → 1`     |
| `Asin(x)`     | arcsin(x)       | `Asin(1) → Pi/2`    |
| `Acos(x)`     | arccos(x)       | `Acos(0) → Pi/2`    |
| `Atan(x)`     | arctan(x)       | `Atan(1) → Pi/4`    |
| `Atan2(y, x)` | Koordinat açısı | `Atan2(1,1) → Pi/4` |

---

## 🔹 6. Hiperbolik Fonksiyonlar

| Fonksiyon  | Açıklama                |
| ---------- | ----------------------- |
| `Sinh(x)`  | Hiperbolik sinüs        |
| `Cosh(x)`  | Hiperbolik kosinüs      |
| `Tanh(x)`  | Hiperbolik tanjant      |
| `Asinh(x)` | Ters hiperbolik sinüs   |
| `Acosh(x)` | Ters hiperbolik kosinüs |
| `Atanh(x)` | Ters hiperbolik tanjant |

---

## 🔹 7. Yuvarlama Fonksiyonları

| Fonksiyon         | Açıklama                   | Örnek                    |
| ----------------- | -------------------------- | ------------------------ |
| `Floor(x)`        | Aşağı yuvarla              | `Floor(3.7) → 3`         |
| `Ceil(x)`         | Yukarı yuvarla             | `Ceil(3.2) → 4`          |
| `Round(x)`        | En yakın tam sayı          | `Round(3.5) → 4`         |
| `Trunc(x)`        | Tam kısmı al               | `Trunc(3.9) → 3`         |
| `Mod(x, y)`       | Mod alma (float)           | `Mod(10, 3) → 1`         |
| `Remainder(x, y)` | IEEE remainder             | `Remainder(10, 3) → 1`   |
| `Modf(x)`         | Tam & kesir kısmını ayırır | `Modf(3.14) → (3, 0.14)` |

---

## 🔹 8. Min / Max / Diğer

| Fonksiyon          | Açıklama                          | Örnek           |
| ------------------ | --------------------------------- | --------------- |
| `Min(x, y)`        | Minimum                           | `Min(2, 5) → 2` |
| `Max(x, y)`        | Maksimum                          | `Max(2, 5) → 5` |
| `Fabs(x)`          | Mutlak değer (Abs ile aynı)       |                 |
| `Frexp(f)`         | Sayıyı mantissa & üs olarak böler |                 |
| `Ldexp(frac, exp)` | frac × 2^exp                      |                 |

/*
---

✅ Böylece `math` paketindeki **tüm fonksiyonlar ve sabitler** özetlenmiş oldu.

---

İstersen ben sana bu tabloyu **PDF** veya **Excel** formatında hazırlayıp verebilirim. Böylece offline çalışırken de kolayca kullanabilirsin. İstiyor musun?
EVET
*/
/*
Süper 👍 O zaman sana hazırladığım tabloyu **Markdown** formatında da veriyorum.

---
*/
//# Go `math/big` Paketi Özeti

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

---
/*
👉 İstersen ben sana bu fonksiyonların **hepsi için çalışır küçük örnek programlar** da yazabilirim. Böylece her fonksiyonu canlı test edebilirsin. İstiyor musun?
EVET
*/

/*
Harika 😄 O zaman ben sana **`math/big` paketindeki tüm tipler (`Int`, `Rat`, `Float`) için küçük çalışır örnek programlar** hazırlayacağım. Her örnek, ilgili fonksiyonları gösterip kolayca çalıştırabileceğin şekilde olacak.

---

## 1️⃣ `big.Int` Örnekleri
*7
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
*/

/*
İstersen ben bunu bir adım daha ileri götürüp tüm işlemleri aynı anda çalıştırıp sonuçları tek seferde gösteren tam bir “math/big hesaplama paneli” hâline getirebilirim. Bunu da hazırlayayım mı?
EVET
*/

Harika 😄 O zaman sana **tüm işlemleri aynı anda yapan, interaktif bir “math/big hesaplama paneli”** hazırlayacağım. Kullanıcı bir kere sayıları giriyor, program hem faktöriyel, hem kesir işlemleri, hem de hassas float hesaplamalarını tek seferde gösteriyor.

---

```go
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
```

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
