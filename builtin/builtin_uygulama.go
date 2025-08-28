//  **`builtin` paketindeki türler, sabitler ve fonksiyonların hepsini örnekleyen** bir mini uygulama yazıyorum.

// Bu programda:

// * Bütün temel tipler kullanılacak
// * `iota`, `true/false` sabitleri gösterilecek
// * Tüm yerleşik fonksiyonlar (`len`, `cap`, `make`, `new`, `append`, `copy`, `delete`, `complex`, `real`, `imag`, `close`, `panic`, `recover`, `print`, `println`) örneklenecek

// ---

// ## 📝 Uygulama: `builtin_demo.go`

package main

import (
	"fmt"
)

func main() {
	fmt.Println("===== BUILTIN TÜRLER =====")
	var b bool = true
	var i int = 42
	var f float64 = 3.14
	var s string = "Merhaba"
	var r rune = 'ğ'
	var by byte = 255
	var c complex64 = complex(2, 3)
	var e error = fmt.Errorf("bir hata oluştu")

	fmt.Println(b, i, f, s, r, by, c, e)

	fmt.Println("\n===== SABİTLER =====")
	const (
		A = iota
		B
		C
	)
	fmt.Println("true:", true, "false:", false, "iota örnek:", A, B, C)

	fmt.Println("\n===== FONKSİYONLAR =====")

	// len & cap
	sl := []int{1, 2, 3}
	fmt.Println("len:", len(sl), "cap:", cap(sl))

	// make
	sl2 := make([]int, 2, 5)
	mp := make(map[string]int)
	ch := make(chan int, 2)
	fmt.Println("make slice:", sl2, "map:", mp, "chan cap:", cap(ch))

	// new
	ptr := new(int)
	fmt.Println("new int:", *ptr)
	*ptr = 99
	fmt.Println("new int atandı:", *ptr)

	// append
	sl = append(sl, 4, 5)
	fmt.Println("append:", sl)

	// copy
	src := []int{7, 8, 9}
	dst := make([]int, 3)
	copy(dst, src)
	fmt.Println("copy:", dst)

	// delete
	mp["a"] = 1
	mp["b"] = 2
	delete(mp, "a")
	fmt.Println("delete map:", mp)

	// complex, real, imag
	cmplx := complex(5, -4)
	fmt.Println("complex:", cmplx, "real:", real(cmplx), "imag:", imag(cmplx))

	// close (channel)
	ch <- 10
	ch <- 20
	close(ch)
	for v := range ch {
		fmt.Println("chan eleman:", v)
	}

	// panic & recover
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recover ile panik yakalandı:", r)
			}
		}()
		panic("Panikkk!")
	}()

	// print & println
	print("print fonksiyonu çıktı -> ")
	println(123, " ", true)
}

// ```

// ---

// ## 📌 Çalıştırıldığında Konsol Çıktısı (örnek)

// ```
// ===== BUILTIN TÜRLER =====
// true 42 3.14 Merhaba 287 255 (2+3i) bir hata oluştu

// ===== SABİTLER =====
// true: true false: false iota örnek: 0 1 2

// ===== FONKSİYONLAR =====
// len: 3 cap: 3
// make slice: [0 0] map: map[] chan cap: 2
// new int: 0
// new int atandı: 99
// append: [1 2 3 4 5]
// copy: [7 8 9]
// delete map: map[b:2]
// complex: (5-4i) real: 5 imag: -4
// chan eleman: 10
// chan eleman: 20
// recover ile panik yakalandı: Panikkk!
// print fonksiyonu çıktı -> 123  true
// ```

// ---

// ✅ Bu uygulama ile `builtin` paketinin **tamamını** canlı olarak görebilirsin.

// ---
