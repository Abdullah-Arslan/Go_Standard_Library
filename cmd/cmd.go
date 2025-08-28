// Go 1.21 ile birlikte standart kütüphaneye cmp paketi eklendi. 
// Bu paket, değerleri karşılaştırmayı daha güvenli, okunabilir ve kolay hale getiriyor. 
// Eskiden if a < b { ... } gibi manuel kontroller veya özel fonksiyonlar yazmak gerekiyordu, ama cmp ile daha düzenli yapılabiliyor.

// Ben sana tüm fonksiyonlarını tek tek anlatayım ve her birine örnek vereyim:

// 📦 cmp Paketi

// cmp paketi üç temel fonksiyon içerir:

// 1-cmp.Compare(a, b)

// 2-cmp.Less(a, b)

// 3-cmp.Or(a, b, ...)

// Bunlar generic fonksiyonlardır, yani sayı, string gibi ordered türlerde çalışır.

// 1️⃣ cmp.Compare(a, b)

// İki değeri karşılaştırır ve int döner:

// -1 → a < b

// 0 → a == b

// +1 → a > b

//📌 Örnek:

package main

import (
	"fmt"
	"cmp"
)

func main() {
	fmt.Println(cmp.Compare(5, 10))   // -1  (5 < 10)
	fmt.Println(cmp.Compare(10, 10))  // 0   (10 == 10)
	fmt.Println(cmp.Compare(20, 10))  // 1   (20 > 10)

	// String örneği
	fmt.Println(cmp.Compare("apple", "banana")) // -1 ("apple" < "banana")
}


// 2️⃣ cmp.Less(a, b)

// Daha okunabilir bir a < b fonksiyonudur. true/false döner.
//📌 Örnek:
package main

import (
	"fmt"
	"cmp"
)

func main() {
	fmt.Println(cmp.Less(3, 7))   // true
	fmt.Println(cmp.Less(7, 3))   // false
	fmt.Println(cmp.Less("a", "b")) // true
}

//Bu, özellikle sıralama algoritmalarında çok kullanışlıdır.
import (
	"slices"
	"cmp"
	"fmt"
)

func main() {
	nums := []int{5, 3, 9, 1}
	slices.SortFunc(nums, cmp.Compare) // Compare ile sıralama
	fmt.Println(nums) // [1 3 5 9]
}


// 3️⃣ cmp.Or(a, b, c, ...)

// Birden fazla değeri öncelik sırasına göre döndürür. İlk sıfır olmayan (zero olmayan) değeri döner.

// Zero değer:

// int → 0

// string → ""

// bool → false

// 📌 Örnek:

package main

import (
	"fmt"
	"cmp"
)

func main() {
	fmt.Println(cmp.Or(0, 0, 5, 10))     // 5
	fmt.Println(cmp.Or("", "hello", "world")) // "hello"
	fmt.Println(cmp.Or(false, true))     // true
}

//Bu, özellikle varsayılan değer atama için çok kullanışlıdır:

func GetUserName(name string) string {
	return cmp.Or(name, "Guest") // Eğer name boşsa "Guest" döner
}

func main() {
	fmt.Println(GetUserName(""))      // Guest
	fmt.Println(GetUserName("Alice")) // Alice
}

| Fonksiyon     | Açıklama                                | Örnek                     |
| ------------- | --------------------------------------- | ------------------------- |
| `cmp.Compare` | a ile b’yi karşılaştırır → -1,0,1 döner | `cmp.Compare(5,10) // -1` |
| `cmp.Less`    | a < b kontrolü yapar → bool döner       | `cmp.Less(3,7) // true`   |
| `cmp.Or`      | İlk sıfır olmayan değeri döner          | `cmp.Or("", "X") // "X"`  |

