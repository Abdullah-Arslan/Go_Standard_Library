/*
**Go dilinde `math/rand` paketini** baştan sona anlatacağım.

---

# 📌 `math/rand` Paketi Nedir?

Go dilindeki `math/rand` paketi, **rastgele sayılar** üretmek için kullanılır.
Bu paket:

* Tamsayı, ondalıklı sayı, dizi, permütasyon gibi rastgele değerler üretir.
* Pseudo-random sayı üreticisi kullanır (deterministik → aynı seed ile aynı sonuç çıkar).
* Eğer gerçek rastgelelik gerekirse `crypto/rand` paketi kullanılmalıdır.

📦 Import edilmesi:
*/
``go
import (
    "fmt"
    "math/rand"
    "time"
)
``
/*
---

# 📌 Temel Kavram: Seed

* `rand.Seed(n int64)` → Rastgele sayı üreticisine başlangıç değeri verir.
* Aynı `seed` değeri ile çalıştırıldığında **aynı sonuçları** üretir.
*/
``go
rand.Seed(time.Now().UnixNano()) // farklı her çalıştırmada farklı sonuç
fmt.Println(rand.Intn(100))      // 0–99 arası sayı
``
/*
---

# 📌 `rand` Paketindeki Fonksiyonlar

## 🎲 1. Tamsayı Fonksiyonları

### 🔹 `Int() int`

Rastgele bir `int` üretir.
*/
``go
rand.Seed(time.Now().UnixNano())
fmt.Println(rand.Int())
``
/*
---

### 🔹 `Intn(n int) int`

0 ile `n-1` arasında sayı döndürür.
*/
``go
fmt.Println(rand.Intn(10)) // 0–9 arası sayı
``
/*
---

### 🔹 `Int31(), Int31n(n int32) int32`

32-bit tamsayı üretir.
*/
``go
fmt.Println(rand.Int31())
fmt.Println(rand.Int31n(100)) // 0–99
``
/*
---

### 🔹 `Int63(), Int63n(n int64) int64`

63-bit tamsayı üretir.
*/
``go
fmt.Println(rand.Int63())
fmt.Println(rand.Int63n(1000)) // 0–999
``
/*
---

### 🔹 `Uint32(), Uint64()`

Pozitif 32-bit veya 64-bit tamsayı üretir.
*/
``go
fmt.Println(rand.Uint32())
fmt.Println(rand.Uint64())
``
/*
---

## 🎲 2. Ondalıklı Sayılar

### 🔹 `Float32() float32`

0.0 ile 1.0 arasında `float32` üretir.
*/
``go
fmt.Println(rand.Float32())
``
/*
---

### 🔹 `Float64() float64`

0.0 ile 1.0 arasında `float64` üretir.
*/
``go
fmt.Println(rand.Float64())
``
/*
---

### 🔹 `NormFloat64() float64`

Normal dağılıma (ortalama 0, varyans 1) göre sayı üretir.
*/
``go
fmt.Println(rand.NormFloat64())
``
/*
---

### 🔹 `ExpFloat64() float64`

Üstel dağılıma göre sayı üretir.
*/
``go
fmt.Println(rand.ExpFloat64())
``
/*
---

## 🎲 3. Dizi / Permütasyon Fonksiyonları

### 🔹 `Perm(n int) []int`

0’dan `n-1`’e kadar sayıların rastgele permütasyonunu döndürür.
*/

``go
fmt.Println(rand.Perm(5)) // örn: [3 1 4 0 2]
``
/*
---

### 🔹 `Shuffle(n int, swap func(i, j int))`

Dizi veya slice içeriğini rastgele karıştırır.
*/
``go
arr := []string{"a", "b", "c", "d", "e"}
rand.Shuffle(len(arr), func(i, j int) {
    arr[i], arr[j] = arr[j], arr[i]
})
fmt.Println(arr)
``
/*
---

## 🎲 4. Kaynak (Source) Kullanımı

`rand.NewSource(seed int64)` ile **özel kaynak** oluşturulabilir.
`rand.New(source)` ile **bağımsız random generator** üretilebilir.
*/
``go
src := rand.NewSource(time.Now().UnixNano())
r := rand.New(src)

fmt.Println(r.Intn(100))
fmt.Println(r.Float64())
``
/*
Bu şekilde aynı anda birden fazla bağımsız rastgele sayı üretici kullanılabilir.

---
*/
# 📌 Özet Tablo

| Kategori               | Fonksiyonlar                                                            |
| ---------------------- | ----------------------------------------------------------------------- |
| **Tamsayı**            | `Int`, `Intn`, `Int31`, `Int31n`, `Int63`, `Int63n`, `Uint32`, `Uint64` |
| **Ondalıklı**          | `Float32`, `Float64`, `NormFloat64`, `ExpFloat64`                       |
| **Dizi / Permütasyon** | `Perm`, `Shuffle`                                                       |
| **Seed / Source**      | `Seed`, `NewSource`, `New`                                              |
/*
---

# ✅ Sonuç

* `math/rand` paketi, Go’da **rastgele sayı üretmek için** kullanılır.
* Deterministiktir (aynı seed → aynı sonuç).
* İstatistiksel dağılımlar (normal, üstel) için destek içerir.
* Diziler üzerinde permütasyon/karıştırma işlemleri yapılabilir.

---

İstersen ben sana bu bilgileri yine **PDF** ve **Excel tablo** halinde hazırlayabilirim. İstiyor musun?
Süper 👍 Şimdi sana **`math/rand` paketindeki her fonksiyon için ayrı çalıştırılabilir Go örnekleri** yazacağım.
Her bir örnek kendi `main.go` dosyası gibi bağımsız çalıştırılabilir olacak.

---

# 📌 `math/rand` Paketindeki Fonksiyonlar ve Örnekler

---

## 1. `Int()`
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
	fmt.Println(rand.Int()) // Rastgele int
}
``
/*
---

## 2. `Intn(n int)`
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
	fmt.Println(rand.Intn(10)) // 0–9 arasında sayı
}
``
/*
---

## 3. `Int31(), Int31n(n int32)`
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
	fmt.Println(rand.Int31())     // Rastgele int32
	fmt.Println(rand.Int31n(100)) // 0–99 arası int32
}
``
/*
---

## 4. `Int63(), Int63n(n int64)`
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
	fmt.Println(rand.Int63())     // Rastgele int64
	fmt.Println(rand.Int63n(100)) // 0–99 arası int64
}
``
/*
---

## 5. `Uint32(), Uint64()`
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
	fmt.Println(rand.Uint32()) // Rastgele uint32
	fmt.Println(rand.Uint64()) // Rastgele uint64
}
``
/*
---

## 6. `Float32(), Float64()`
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
	fmt.Println(rand.Float32()) // 0.0–1.0 arasında float32
	fmt.Println(rand.Float64()) // 0.0–1.0 arasında float64
}
``
/*
---

## 7. `NormFloat64()`
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
	fmt.Println(rand.NormFloat64()) // Normal dağılım (ortalama 0, varyans 1)
}
``
/*
---

## 8. `ExpFloat64()`
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
	fmt.Println(rand.ExpFloat64()) // Üstel dağılım
}
``
/*
---

## 9. `Perm(n int)`
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
	fmt.Println(rand.Perm(5)) // [3 1 4 0 2] gibi bir permütasyon
}
``
/*
---

## 10. `Shuffle(n int, swap func(i, j int))`
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
	arr := []string{"a", "b", "c", "d", "e"}

	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	fmt.Println(arr) // Karıştırılmış slice
}
``
/*
---

## 11. `Seed(seed int64)`
*/
``go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(42) // aynı seed → aynı sonuç
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
}
``
/*
---

## 12. `NewSource(seed int64)` ve `New(src Source)`
*/
``go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	fmt.Println(r.Intn(100)) // 0–99
	fmt.Println(r.Float64()) // 0.0–1.0 arasında float
}
``
/*
---

# ✅ Özet

Yukarıdaki örneklerde `math/rand` paketindeki **her fonksiyon** için ayrı çalıştırılabilir Go kodlarını gördün.

---

👉 İstersen ben sana bu örneklerin tamamını **tek bir PDF dökümanı** veya **bir zip dosyası içinde `.go` dosyaları** halinde de hazırlayabilirim. Hangisini tercih edersin?
*/