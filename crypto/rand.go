/*
Şimdi sana **Go dilinde `math/rand` paketi**ni en temelden tüm detaylarıyla açıklayayım.

Go’da rastgele sayı üretmek için genellikle `math/rand` paketi kullanılır.

---

## 📌 `math/rand` Paketi Nedir?

* **Pseudo-Random Number Generator (PRNG)** yani **sözde rastgele sayı üreticisi** sağlar.
* Gerçek rastgelelik sağlamaz, başlangıç değerine (seed) bağlı olarak her zaman aynı sayılar üretilir.
* Eğer `rand.Seed()` fonksiyonuyla farklı bir seed verilmezse, **her çalıştırmada aynı sayı dizisini üretir.**

Eğer **kriptografik olarak güvenli rastgele sayılar** isteniyorsa, `crypto/rand` paketi kullanılmalıdır.

---

## 📌 Önemli Fonksiyonlar ve Kullanımları

### 1. `rand.Seed(seed int64)`

Rastgele sayı üreteci için başlangıç değerini ayarlar.
*/
``go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Her çalıştırmada farklı sonuç için
	fmt.Println(rand.Int())         // Rastgele int
}
``
/*
📌 Eğer `Seed` verilmezse, her çalıştırmada **aynı sayı** çıkar.

---

### 2. `rand.Int()`

0 ile `math.MaxInt` arasında rastgele bir `int` döner.
*/
``go
fmt.Println(rand.Int()) 
``
/*
---

### 3. `rand.Intn(n int)`

`0 <= x < n` olacak şekilde rastgele bir int döner.
*/
``go
fmt.Println(rand.Intn(10)) // 0 ile 9 arasında sayı
``
/*
---

### 4. `rand.Int31()` ve `rand.Int31n(n int32)`

* `Int31()`: 0 ile 2³¹-1 arasında sayı döner.
* `Int31n(n)`: 0 ile n-1 arasında `int32` döner.
*/
``go
fmt.Println(rand.Int31())
fmt.Println(rand.Int31n(100))
``
/*
---

### 5. `rand.Int63()` ve `rand.Int63n(n int64)`

* `Int63()`: 0 ile 2⁶³-1 arasında sayı döner.
* `Int63n(n)`: 0 ile n-1 arasında `int64` döner.
*/
``go
fmt.Println(rand.Int63())
fmt.Println(rand.Int63n(1000))
``
/*
---

### 6. `rand.Float32()` ve `rand.Float64()`

* `Float32()`: 0.0 ≤ f < 1.0 aralığında `float32`.
* `Float64()`: 0.0 ≤ f < 1.0 aralığında `float64`.
*/
``go
fmt.Println(rand.Float32())
fmt.Println(rand.Float64())
``
/*
---

### 7. `rand.Perm(n int)`

`0` ile `n-1` arası sayıları **rastgele karışık** bir slice döner.
*/
``go
fmt.Println(rand.Perm(5)) // Örn: [3 0 4 1 2]
``
/*
---

### 8. `rand.Shuffle(n int, swap func(i, j int))`

Elemanları rastgele karıştırır.
*/
``go
arr := []string{"elma", "armut", "muz", "çilek"}
rand.Shuffle(len(arr), func(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
})
fmt.Println(arr)
``
/*
---

### 9. `rand.NormFloat64()`

Normal dağılıma (Gaussian distribution, ortalama=0, std=1) göre sayı döner.
*/
``go
fmt.Println(rand.NormFloat64())
``
/*
---

### 10. `rand.ExpFloat64()`

Üstel dağılıma (exponential distribution, ortalama=1) göre sayı döner.
*/
``go
fmt.Println(rand.ExpFloat64())
``
/*
---

### 11. `rand.New(src rand.Source)`

Kendi rastgele üreteç objeni oluşturabilirsin.
*/
``go
src := rand.NewSource(time.Now().UnixNano())
r := rand.New(src)

fmt.Println(r.Intn(100))
fmt.Println(r.Float64())
``

//---

//## 📌 Komple Örnek

``go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 1. Rastgele int
	fmt.Println("Int:", rand.Int())

	// 2. Belirli aralıkta int
	fmt.Println("Intn(10):", rand.Intn(10))

	// 3. Float sayılar
	fmt.Println("Float32:", rand.Float32())
	fmt.Println("Float64:", rand.Float64())

	// 4. Permütasyon
	fmt.Println("Perm(5):", rand.Perm(5))

	// 5. Shuffle
	arr := []int{1, 2, 3, 4, 5}
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	fmt.Println("Shuffled:", arr)

	// 6. Normal dağılım
	fmt.Println("Normal:", rand.NormFloat64())

	// 7. Üstel dağılım
	fmt.Println("Exponential:", rand.ExpFloat64())
}
``
/*
---

## 📌 Özet

* `math/rand`: Sözde rastgele sayı üretimi için.
* `rand.Seed()` kullanılmazsa aynı sayılar tekrar eder.
* Sayılar, `Int`, `Intn`, `Float64`, `Perm`, `Shuffle`, `NormFloat64`, `ExpFloat64` gibi farklı yöntemlerle üretilebilir.
* **Kripto güvenliği için `crypto/rand` kullanılmalı.**

---

👉 İstersen sana `math/rand` ile **zar atma simülasyonu** veya **kart destesini karıştırma** örneği de yazabilirim. İstiyor musun?
EVET

Süper 👍 O zaman sana **iki örnek** göstereyim:

---

# 🎲 Örnek 1: Zar Atma Simülasyonu
*/
``go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Her çalıştırmada farklı olsun

	// 1 ile 6 arasında zar simülasyonu
	dice := rand.Intn(6) + 1
	fmt.Println("Zar sonucu:", dice)
}
``
/*
📌 `rand.Intn(6)` bize **0–5** arası değer döner.
Buna `+1` ekleyince **1–6** arasında zar simülasyonu yapılır.

---

# 🃏 Örnek 2: Kart Deste Karıştırma
*/
``go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Basit iskambil destesi (sadece sembol + rakam)
	deck := []string{
		"♠A", "♠2", "♠3", "♠4", "♠5", "♠6", "♠7", "♠8", "♠9", "♠10", "♠J", "♠Q", "♠K",
		"♥A", "♥2", "♥3", "♥4", "♥5", "♥6", "♥7", "♥8", "♥9", "♥10", "♥J", "♥Q", "♥K",
		"♦A", "♦2", "♦3", "♦4", "♦5", "♦6", "♦7", "♦8", "♦9", "♦10", "♦J", "♦Q", "♦K",
		"♣A", "♣2", "♣3", "♣4", "♣5", "♣6", "♣7", "♣8", "♣9", "♣10", "♣J", "♣Q", "♣K",
	}

	// Deste karıştırma
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	// Karıştırılmış deste
	fmt.Println("Karıştırılmış deste:")
	for i, card := range deck {
		fmt.Printf("%2d: %s\n", i+1, card)
	}
}
``
/*
📌 Burada `rand.Shuffle()` fonksiyonunu kullanarak **destedeki kartların sırasını rastgele karıştırdık**.
Her çalıştırmada farklı bir dizilim olur.

---

👉 İstersen ben sana **çok zar atan bir simülasyon (örneğin 5 zar atıp toplamını bulan)** ya da **kart oyununda oyunculara kart dağıtan** bir örnek de yazabilirim. Hangisini yapalım?
*/