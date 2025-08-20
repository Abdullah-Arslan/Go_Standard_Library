// Örnekte: `cmp.Compare`, `cmp.Less` ve `cmp.Or` fonksiyonlarının hepsini kapsayan **gerçek hayata yakın bir örnek** yazayım.

// Senaryo:

// * Bir `User` listemiz var.
// * Kullanıcıların `Name`, `Age`, `City` bilgileri var.
// * `cmp.Compare` ile yaşa göre sıralama yapacağız.
// * `cmp.Less` ile şehre göre sıralama yapacağız.
// * `cmp.Or` ile boş gelen kullanıcı adını **"Guest"** yapacağız.

// ---

// ## 📌 Kapsamlı Örnek

package main

import (
	"cmp"
	"fmt"
	"slices"
)

// User yapısı
type User struct {
	Name string
	Age  int
	City string
}

// Kullanıcı adı boşsa "Guest" döndür
func (u User) DisplayName() string {
	return cmp.Or(u.Name, "Guest")
}

func main() {
	users := []User{
		{"Alice", 25, "Istanbul"},
		{"", 30, "Ankara"}, // Name boş
		{"Charlie", 20, "Izmir"},
		{"Bob", 30, "Bursa"},
	}

	fmt.Println("👉 Orijinal liste:")
	for _, u := range users {
		fmt.Printf("%s (%d) - %s\n", u.DisplayName(), u.Age, u.City)
	}

	// 1. Yaşa göre sıralama (küçükten büyüğe)
	slices.SortFunc(users, func(a, b User) int {
		return cmp.Compare(a.Age, b.Age)
	})

	fmt.Println("\n👉 Yaşa göre sıralı liste:")
	for _, u := range users {
		fmt.Printf("%s (%d) - %s\n", u.DisplayName(), u.Age, u.City)
	}

	// 2. Şehre göre alfabetik sıralama
	slices.SortFunc(users, func(a, b User) int {
		// cmp.Less ile doğrudan bool döndüğü için Compare'e çeviriyoruz
		if cmp.Less(a.City, b.City) {
			return -1
		}
		if cmp.Less(b.City, a.City) {
			return 1
		}
		return 0
	})

	fmt.Println("\n👉 Şehre göre alfabetik sıralı liste:")
	for _, u := range users {
		fmt.Printf("%s (%d) - %s\n", u.DisplayName(), u.Age, u.City)
	}
}

//```

// ---

// ## 📌 Çıktı

// ```
// 👉 Orijinal liste:
// Alice (25) - Istanbul
// Guest (30) - Ankara
// Charlie (20) - Izmir
// Bob (30) - Bursa

// 👉 Yaşa göre sıralı liste:
// Charlie (20) - Izmir
// Alice (25) - Istanbul
// Guest (30) - Ankara
// Bob (30) - Bursa

// 👉 Şehre göre alfabetik sıralı liste:
// Guest (30) - Ankara
// Bob (30) - Bursa
// Alice (25) - Istanbul
// Charlie (20) - Izmir
// ```

// ---

// Bu örnekte:

// * `cmp.Or` → boş isimleri `"Guest"` yaptı.
// * `cmp.Compare` → yaşları sıraladı.
// * `cmp.Less` → şehirleri alfabetik olarak sıraladı.

// ---
