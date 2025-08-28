/*
Go’nun standart kütüphanesinde bulunan **`container/ring`** paketi, **dairesel (circular) liste** veri yapısını sağlar.

📌 **Circular list (ring)**:

* Elemanlar birbirine halka şeklinde bağlıdır.
* Son elemana geldiğinde tekrar başa döner.
* Özellikle **round-robin (tur usulü)** algoritmalarında ve **sonsuz döngüsel veri** tutmada kullanılır.

---

# 🔹 Ring Paketinin Yapısı

`ring.Ring` yapısı şu özelliklere sahiptir:

### Yapı Alanları:

* **Value any** → Düğüme eklenen değer
* \**Next() *Ring** → Sonraki elemana geç
* \**Prev() *Ring** → Önceki elemana geç
* **Len() int** → Halka uzunluğunu döndürür
* **Do(func(any))** → Halkadaki her eleman için verilen fonksiyonu uygular
* **Move(n int)** → İleri (+n) veya geri (-n) hareket eder

### Fonksiyonlar:

* `ring.New(n)` → `n` elemanlı bir halka oluşturur (her düğümün Value’su nil olur)
* `r.Link(s)` → İki halkayı birleştirir
* `r.Unlink(n)` → `r`’den başlayarak `n` elemanı siler ve yeni halka döner

---

# 🔹 Basit Örnek
*/

package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// 5 elemanlı dairesel liste oluştur
	r := ring.New(5)

	// Elemanları doldur
	for i := 0; i < r.Len(); i++ {
		r.Value = i + 1
		r = r.Next()
	}

	// Elemanları yazdır
	fmt.Println("Ring elemanları:")
	r.Do(func(x any) {
		fmt.Print(x, " ")
	})
}
/*

### 📌 Çıktı

```
Ring elemanları:
1 2 3 4 5
```

---

# 🔹 Move (İleri / Geri Hareket)
*/

package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(5)
	for i := 0; i < r.Len(); i++ {
		r.Value = i + 1
		r = r.Next()
	}

	// Başlangıç noktası
	fmt.Println("Şu an:", r.Value)

	// 2 ileri git
	fmt.Println("2 ileri:", r.Move(2).Value)

	// 1 geri git
	fmt.Println("1 geri:", r.Move(-1).Value)
}
/*

### 📌 Çıktı

```
Şu an: 1
2 ileri: 3
1 geri: 5
```

---

# 🔹 Link & Unlink

`Link` → İki halkayı birleştirir
`Unlink` → Halkadan eleman siler
*/
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// 3 elemanlı halka
	r1 := ring.New(3)
	for i := 0; i < 3; i++ {
		r1.Value = i + 1
		r1 = r1.Next()
	}

	// 2 elemanlı halka
	r2 := ring.New(2)
	for i := 0; i < 2; i++ {
		r2.Value = i + 100
		r2 = r2.Next()
	}

	// r1 ve r2'yi birleştir
	r1.Link(r2)

	fmt.Println("Birleştirilmiş halka:")
	r1.Do(func(x any) {
		fmt.Print(x, " ")
	})

	// 2 elemanı sil (Unlink)
	removed := r1.Unlink(2)
	fmt.Println("\nSilinenler:")
	removed.Do(func(x any) {
		fmt.Print(x, " ")
	})

	fmt.Println("\nKalan halka:")
	r1.Do(func(x any) {
		fmt.Print(x, " ")
	})
}
/*

### 📌 Çıktı

```
Birleştirilmiş halka:
1 2 3 100 101 
Silinenler:
100 101 
Kalan halka:
1 2 3 
```

---

# 🔹 Gerçek Kullanım: Round-Robin Scheduler

`ring` en çok **round-robin** tarzı görev dağıtımında kullanılır.
*/

package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// 3 CPU çekirdeği
	cores := ring.New(3)
	cores.Value = "CPU-1"
	cores.Next().Value = "CPU-2"
	cores.Next().Next().Value = "CPU-3"

	// 6 görev, round-robin ile dağıt
	r := cores
	for i := 1; i <= 6; i++ {
		fmt.Printf("Görev %d -> %s\n", i, r.Value)
		r = r.Next()
	}
}

/*
### 📌 Çıktı

```
Görev 1 -> CPU-1
Görev 2 -> CPU-2
Görev 3 -> CPU-3
Görev 4 -> CPU-1
Görev 5 -> CPU-2
Görev 6 -> CPU-3
```

---

# 🔹 Özet

Go’daki **`container/ring`** paketi:

* **Dairesel liste** (circular linked list) implementasyonu sağlar.
* **Move, Do, Link, Unlink** metodlarıyla güçlüdür.
* Özellikle **round-robin**, **buffer yapıları** ve **sonsuz döngüler** için çok uygundur.

---

👉 İstersen sana `ring` kullanarak **gerçek bir uygulama senaryosu** (örneğin **Müzik Çalar Playlist** veya **Oyunlarda sıra tabanlı tur sistemi**) örneği yazabilirim. Hangisini istersin?
*/

/*
Süper 👍 O zaman sana `container/ring` paketini kullanarak **Müzik Çalar Playlist** örneği göstereyim 🎵

📌 Senaryo:

* Playlist bir **dairesel liste (ring)** olarak tutulur.
* Kullanıcı **Next (sonraki şarkı)** ve **Prev (önceki şarkı)** ile gezinebilir.
* Liste döngüsel olduğu için son şarkıdan sonra yine başa dönülür.

---

## 🚀 Örnek: Müzik Çalar Playlist
*/

package main

import (
	"container/ring"
	"fmt"
)

type Playlist struct {
	songs *ring.Ring
	curr  *ring.Ring
}

// Yeni playlist oluştur
func NewPlaylist(tracks []string) *Playlist {
	r := ring.New(len(tracks))
	for _, t := range tracks {
		r.Value = t
		r = r.Next()
	}
	return &Playlist{songs: r, curr: r}
}

// Şu an çalan şarkı
func (p *Playlist) Current() string {
	return p.curr.Value.(string)
}

// Sonraki şarkıya geç
func (p *Playlist) Next() {
	p.curr = p.curr.Next()
}

// Önceki şarkıya dön
func (p *Playlist) Prev() {
	p.curr = p.curr.Prev()
}

func main() {
	// Playlist oluştur
	tracks := []string{"Shape of You", "Believer", "Numb", "Havana"}
	player := NewPlaylist(tracks)

	fmt.Println("Şu an çalan:", player.Current())

	player.Next()
	fmt.Println("Sonraki şarkı:", player.Current())

	player.Next()
	fmt.Println("Sonraki şarkı:", player.Current())

	player.Prev()
	fmt.Println("Önceki şarkı:", player.Current())

	// Döngüsel çalışma
	player.Next()
	player.Next()
	player.Next()
	player.Next()
	fmt.Println("Döngü tamamlandı, tekrar başa dönüldü:", player.Current())
}
/*

---

## 📌 Çıktı

```
Şu an çalan: Shape of You
Sonraki şarkı: Believer
Sonraki şarkı: Numb
Önceki şarkı: Believer
Döngü tamamlandı, tekrar başa dönüldü: Shape of You
```

---

✅ Burada `container/ring`:

* Playlist’i **dairesel liste** olarak tuttu.
* `Next()` ve `Prev()` sayesinde kolayca şarkılar arasında dolaşıldı.
* Son şarkıdan sonra **tekrar başa dönüldü** (circular structure avantajı).

*/