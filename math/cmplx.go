/*
**Go dilinde `cmplx` paketini** en baştan sona kadar açıklayayım.

---

## 📌 `cmplx` Paketi Nedir?

`cmplx` paketi, Go dilinde **karmaşık sayılar (complex numbers)** üzerinde matematiksel işlemler yapabilmek için kullanılan standart kütüphanedir.

Go dilinde **karmaşık sayılar** `complex64` ve `complex128` tiplerinde bulunur:

* `complex64` → `float32` + `float32i`
* `complex128` → `float64` + `float64i`

`cmplx` paketi özellikle `complex128` tipinde çalışır ve trigonometrik, logaritmik, üstel, kök alma gibi fonksiyonlar sağlar.

---

## 📌 Paket İmport Edilmesi
*/
``go
import (
    "fmt"
    "math/cmplx"
)
``
/*
---

## 📌 `cmplx` Paketindeki Fonksiyonlar ve Örnekler

### 1. **Temel Fonksiyonlar**

#### 🔹 `Abs(x complex128) float64`

Bir karmaşık sayının **mutlak değerini (büyüklüğünü)** döndürür.
Formül: `√(real² + imag²)`
*/
``go
z := complex(3, 4) // 3 + 4i
fmt.Println(cmplx.Abs(z)) // 5
``
/*
---

#### 🔹 `Phase(x complex128) float64`

Karmaşık sayının açısını (radyan cinsinden) döndürür.
Formül: `atan2(imag, real)`
*/
``go
z := complex(1, 1)
fmt.Println(cmplx.Phase(z)) // 0.785398163 (π/4)
``
/*
---

#### 🔹 `Polar(x complex128) (r, θ float64)`

Bir karmaşık sayıyı **kutupsal koordinatlara** çevirir.
*/
``go
z := complex(1, 1)
r, θ := cmplx.Polar(z)
fmt.Println(r, θ) // 1.414213562 0.785398163
``
/*
---

#### 🔹 `Rect(r, θ float64) complex128`

Kutupsal koordinatları (`r, θ`) tekrar **dikdörtgen (kartesyen)** forma çevirir.
?7
``go
z := cmplx.Rect(2, cmplx.Pi/4)
fmt.Println(z) // (1.414213562 + 1.414213562i)
``
/*
---

### 2. **Üstel ve Logaritmik Fonksiyonlar**

#### 🔹 `Exp(x complex128) complex128`

`e^x` değerini döndürür.
?7
``go
z := complex(0, math.Pi)
fmt.Println(cmplx.Exp(z)) // -1+0i
``
/*
---

#### 🔹 `Log(x complex128) complex128`

Karmaşık logaritma (`ln(x)`).
*/
``go
z := complex(1, 1)
fmt.Println(cmplx.Log(z)) // (0.34657 + 0.78539i)
``
/*
---

#### 🔹 `Log10(x complex128) complex128`

10 tabanında logaritma.
*/
``go
z := complex(100, 0)
fmt.Println(cmplx.Log10(z)) // (2+0i)
``
/*
---

#### 🔹 `Sqrt(x complex128) complex128`

Karmaşık karekök.
*/
``go
z := complex(-1, 0)
fmt.Println(cmplx.Sqrt(z)) // (0+1i)
``
/*
---

### 3. **Trigonometrik Fonksiyonlar**

* `Sin(x complex128) complex128`
* `Cos(x complex128) complex128`
* `Tan(x complex128) complex128`
*/
``go
z := complex(0, math.Pi/2)
fmt.Println(cmplx.Sin(z)) // (0+2.301298902i)
fmt.Println(cmplx.Cos(z)) // (2.509178478+0i)
fmt.Println(cmplx.Tan(z)) // (0+0.917152335i)
``
/*
---

### 4. **Ters Trigonometrik Fonksiyonlar**

* `Asin(x complex128) complex128`
* `Acos(x complex128) complex128`
* `Atan(x complex128) complex128`
*/
``go
z := complex(1, 0)
fmt.Println(cmplx.Asin(z)) // (1.570796327+0i)
fmt.Println(cmplx.Acos(z)) // (0+0i)
fmt.Println(cmplx.Atan(z)) // (0.785398163+0i)
``
/*
---

### 5. **Hiperbolik Fonksiyonlar**

* `Sinh(x complex128) complex128`
* `Cosh(x complex128) complex128`
* `Tanh(x complex128) complex128`
*/
``go
z := complex(1, 1)
fmt.Println(cmplx.Sinh(z)) // (0.6349639148+1.29845758i)
fmt.Println(cmplx.Cosh(z)) // (0.8337300251+0.9888977058i)
fmt.Println(cmplx.Tanh(z)) // (1.08392333e-19+0.76159415595i)
``
/*
---

### 6. **Ters Hiperbolik Fonksiyonlar**

* `Asinh(x complex128) complex128`
* `Acosh(x complex128) complex128`
* `Atanh(x complex128) complex128`
*/
``go
z := complex(1, 0)
fmt.Println(cmplx.Asinh(z)) // (0.881373587+0i)
fmt.Println(cmplx.Acosh(z)) // (0+1.047197551i)
fmt.Println(cmplx.Atanh(z)) // (0.549306144+0i)
``
/*
---

### 7. **Sabitler**

* `cmplx.Pi` → π (pi sayısı, `3.14159...`)
*/
``go
fmt.Println(cmplx.Pi) // 3.141592653589793
``

//---

## 📌 Özet Tablo

| Fonksiyon Kategorisi   | Fonksiyonlar                    |
| ---------------------- | ------------------------------- |
| **Temel**              | `Abs`, `Phase`, `Polar`, `Rect` |
| **Üstel / Logaritmik** | `Exp`, `Log`, `Log10`, `Sqrt`   |
| **Trigonometrik**      | `Sin`, `Cos`, `Tan`             |
| **Ters Trigonometrik** | `Asin`, `Acos`, `Atan`          |
| **Hiperbolik**         | `Sinh`, `Cosh`, `Tanh`          |
| **Ters Hiperbolik**    | `Asinh`, `Acosh`, `Atanh`       |
| **Sabitler**           | `Pi`                            |
/*
---

## ✅ Sonuç

`cmplx` paketi, Go’da **karmaşık sayılarla bilimsel ve mühendislik hesaplamaları** yapabilmek için kullanılır. Trigonometriden logaritmaya, kök almadan kutupsal dönüşüme kadar birçok fonksiyon içerir.

---

İstersen ben sana bu bilgileri **PDF** ve **Excel tablo** halinde hazırlayıp verebilirim. Offline çalışırken kolay olur. İstiyor musun?
EVET
*/