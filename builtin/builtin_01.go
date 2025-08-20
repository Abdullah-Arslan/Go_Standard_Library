

// Go’da **`builtin`** adında özel bir **paket** vardır.

// * Bu paket aslında **Go diline gömülü (built-in) türleri, sabitleri ve fonksiyonları** içerir.
// * Yani `int`, `string`, `len()`, `make()`, `append()`, `panic()` gibi şeyler aslında **`builtin` paketinde tanımlıdır**.
// * Ama biz kod yazarken `import "builtin"` diye eklemeyiz → çünkü bu paket derleyici tarafından otomatik olarak kullanılır.
// * Kaynağına bakabilirsin: [https://pkg.go.dev/builtin](https://pkg.go.dev/builtin)

// Kısaca:
// 👉 **`builtin` paketi = Go’nun çekirdeğine gömülü olan temel araçların tanım yeri.**

// ---

// # 📦 `builtin` Paketinde Neler Var?


//  1. **Temel Türler**

// * `bool`, `byte`, `rune`
// * `int`, `int8`, `int16`, `int32`, `int64`
// * `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`
// * `float32`, `float64`
// * `complex64`, `complex128`
// * `string`, `error`

// ✅ Örnek:


package main

import "fmt"

func main() {
	var a bool = true
	var b int = 42
	var c float64 = 3.14
	var d string = "Merhaba"
	var e error = fmt.Errorf("bir hata oluştu")

	fmt.Println(a, b, c, d, e)
}

// ---

//  2. **Sabitler**

// * `true`, `false`
// * `iota` → artan sabit üretici

// Go’da iota özel bir sabit tanımlama aracıdır.

// iota, const blokları içinde kullanılır ve otomatik artan tamsayılar üretir.

// İlk satırda 0 değerinden başlar, her satırda 1 artar.

// Genellikle enum (numaralandırma) benzeri değerler tanımlamak için kullanılır

// ✅ Örnek:


package main

import "fmt"

func main() {
	const (
		A = iota // 0
		B        // 1
		C        // 2
	)

	fmt.Println(true, false, A, B, C)
}


//---

//  3. **Yerleşik Fonksiyonlar**

// `builtin` paketinde tanımlı olan fonksiyonlar şunlardır:

// 🔹 `len(v)` → Uzunluk döner


fmt.Println(len("Merhaba"))       // 7
fmt.Println(len([]int{1,2,3,4})) // 4


// 🔹 `cap(v)` → Kapasite döner


s := make([]int, 2, 5)
fmt.Println(len(s), cap(s)) // 2 5


// 🔹 `make()` → Slice, map, channel oluşturur


sl := make([]int, 3)
mp := make(map[string]int)
ch := make(chan int)

fmt.Println(sl, mp, ch)


// 🔹 `new()` → Bellekte sıfır değerli bir değişken oluşturur (pointer döner)


x := new(int)
fmt.Println(*x) // 0
*x = 42
fmt.Println(*x) // 42

// 🔹 `append(slice, elems...)` → Slice’a eleman ekler


s := []int{1, 2}
s = append(s, 3, 4)
fmt.Println(s) // [1 2 3 4]


// 🔹 `copy(dst, src)` → Slice kopyalar


src := []int{1,2,3}
dst := make([]int, 3)
copy(dst, src)
fmt.Println(dst) // [1 2 3]
```

//🔹 `delete(map, key)` → Map’ten eleman siler


m := map[string]int{"a": 1, "b": 2}
delete(m, "a")
fmt.Println(m) // map[b:2]


//  🔹 `complex(r, i)` → Karmaşık sayı oluşturur

//  🔹 `real(c)` → Karmaşık sayının reel kısmı

//  🔹 `imag(c)` → Karmaşık sayının sanal kısmı


c := complex(2, 3) // 2+3i
fmt.Println(real(c), imag(c)) // 2 3


// #### 🔹 `close(ch)` → Kanal kapatır

// #### 🔹 `panic(v)` → Panik başlatır (hata)

// #### 🔹 `recover()` → Panikten kurtarır


func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panik yakalandı:", r)
		}
	}()
	panic("Bir hata oldu")
}

// ```

// #### 🔹 `print()` ve `println()`

// * Debug amaçlı basit çıktı verir (fmt gibi değildir).

// ---

// # 🚀 Özet

// * **`builtin` paketi** = Go’nun temel yapı taşları
// * İçinde şunlar var:

//   * Temel türler (int, string, error, …)
//   * Sabitler (true, false, iota)
//   * Yerleşik fonksiyonlar (`len`, `cap`, `make`, `append`, `copy`, `delete`, `complex`, `real`, `imag`, `close`, `panic`, `recover`, `print`, `println`)

// 👉 Bunları kullanmak için `import "builtin"` yazmana gerek yok. Zaten **her Go dosyasında otomatik olarak kullanılabilir.**

// ---

