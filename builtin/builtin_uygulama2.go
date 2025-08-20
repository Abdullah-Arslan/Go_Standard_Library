// Builtin paketinin her fonksiyon için ayrı küçük demolar haline getirebilirim.
// (örneğin demoLen(), demoMake() gibi).

// O zaman her `builtin` fonksiyonunu/hususunu ayrı birer **küçük demo fonksiyonuna** bölelim.

// Bu sayede:

// * Kod daha düzenli olur
// * Tek tek çağırıp denemek kolaylaşır
// * Hepsi tek dosyada ama fonksiyon fonksiyon ayrılmış olur

// ---

// ## 📝 Uygulama: `builtin_demos.go`

package main

import (
	"fmt"
)

func main() {
	fmt.Println("===== BUILTIN DEMOLAR =====")

	demoTypes()
	demoConstants()

	demoLenCap()
	demoMake()
	demoNew()
	demoAppend()
	demoCopy()
	demoDelete()
	demoComplex()
	demoChannel()
	demoPanicRecover()
	demoPrint()
}

// //////////////////// DEMOLAR ////////////////////

func demoTypes() {
	fmt.Println("\n--- DEMO: Types ---")
	var b bool = true
	var i int = 42
	var f float64 = 3.14
	var s string = "Merhaba"
	var r rune = 'ğ'
	var by byte = 255
	var c complex64 = complex(2, 3)
	var e error = fmt.Errorf("bir hata oluştu")

	fmt.Println(b, i, f, s, r, by, c, e)
}

func demoConstants() {
	fmt.Println("\n--- DEMO: Constants ---")
	const (
		A = iota
		B
		C
	)
	fmt.Println("true:", true, "false:", false, "iota:", A, B, C)
}

func demoLenCap() {
	fmt.Println("\n--- DEMO: len & cap ---")
	sl := []int{1, 2, 3}
	fmt.Println("slice:", sl, "len:", len(sl), "cap:", cap(sl))
}

func demoMake() {
	fmt.Println("\n--- DEMO: make ---")
	sl := make([]int, 2, 5)
	mp := make(map[string]int)
	ch := make(chan int, 3)
	fmt.Println("slice:", sl, "map:", mp, "chan cap:", cap(ch))
}

func demoNew() {
	fmt.Println("\n--- DEMO: new ---")
	ptr := new(int)
	fmt.Println("başlangıç:", *ptr)
	*ptr = 99
	fmt.Println("atandı:", *ptr)
}

func demoAppend() {
	fmt.Println("\n--- DEMO: append ---")
	sl := []int{1, 2}
	sl = append(sl, 3, 4)
	fmt.Println("append sonrası:", sl)
}

func demoCopy() {
	fmt.Println("\n--- DEMO: copy ---")
	src := []int{7, 8, 9}
	dst := make([]int, 3)
	copy(dst, src)
	fmt.Println("src:", src, "dst:", dst)
}

func demoDelete() {
	fmt.Println("\n--- DEMO: delete ---")
	m := map[string]int{"a": 1, "b": 2}
	delete(m, "a")
	fmt.Println("map:", m)
}

func demoComplex() {
	fmt.Println("\n--- DEMO: complex, real, imag ---")
	c := complex(5, -4)
	fmt.Println("complex:", c, "real:", real(c), "imag:", imag(c))
}

func demoChannel() {
	fmt.Println("\n--- DEMO: close(channel) ---")
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20
	close(ch)
	for v := range ch {
		fmt.Println("chan eleman:", v)
	}
}

func demoPanicRecover() {
	fmt.Println("\n--- DEMO: panic & recover ---")
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recover ile yakalandı:", r)
			}
		}()
		panic("Panikkkk!")
	}()
}

func demoPrint() {
	fmt.Println("\n--- DEMO: print & println ---")
	print("print çıktı -> ")
	println(123, " ", true)
}

// ```

// ---

// ## 📌 Çalıştırıldığında Konsol Çıktısı (özet)

// ```
// ===== BUILTIN DEMOLAR =====

// --- DEMO: Types ---
// true 42 3.14 Merhaba 287 255 (2+3i) bir hata oluştu

// --- DEMO: Constants ---
// true: true false: false iota: 0 1 2

// --- DEMO: len & cap ---
// slice: [1 2 3] len: 3 cap: 3

// --- DEMO: make ---
// slice: [0 0] map: map[] chan cap: 3

// --- DEMO: new ---
// başlangıç: 0
// atandı: 99

// --- DEMO: append ---
// append sonrası: [1 2 3 4]

// --- DEMO: copy ---
// src: [7 8 9] dst: [7 8 9]

// --- DEMO: delete ---
// map: map[b:2]

// --- DEMO: complex, real, imag ---
// complex: (5-4i) real: 5 imag: -4

// --- DEMO: close(channel) ---
// chan eleman: 10
// chan eleman: 20

// --- DEMO: panic & recover ---
// recover ile yakalandı: Panikkkk!

// --- DEMO: print & println ---
// print çıktı -> 123  true
// ```

// ---

// 👉 Böylece her `builtin` özelliği **küçük, bağımsız demolar** halinde oldu.
