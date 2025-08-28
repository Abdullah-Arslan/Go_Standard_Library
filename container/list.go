/*
Go’da **`container/list`** paketi, **çift yönlü bağlı liste (doubly linked list)** veri yapısını sağlar.

Slice’lardan farkı:

* Slice’lar **index tabanlı** çalışır, `list` ise **bağlı elemanlarla** çalışır.
* `list` ile **ortadan eleman eklemek/silmek O(1)** karmaşıklığındadır (ama elemana ulaşmak için sıralı gezinmek gerekir).
* Özellikle **FIFO (queue)**, **LIFO (stack)** veya **sıralı veri ekleme/çıkarma** işlemlerinde avantajlıdır.

---

## 🔹 `list` Paketinin Temel Yapısı

* `list.List` → Liste yapısı
* `list.Element` → Listenin her elemanını temsil eder (`Value` alanı veriyi taşır)

📌 Önemli metodlar:

* `list.New()` → Yeni liste oluşturur
* `l.PushFront(v)` → Listenin başına eleman ekler
* `l.PushBack(v)` → Listenin sonuna eleman ekler
* `l.InsertBefore(v, e)` → Belirtilen elemandan önce ekleme
* `l.InsertAfter(v, e)` → Belirtilen elemandan sonra ekleme
* `l.Remove(e)` → Belirtilen elemanı siler
* `l.Front()` → İlk elemana erişim
* `l.Back()` → Son elemana erişim
* `e.Next()` → Sonraki elemana git
* `e.Prev()` → Önceki elemana git

---

## 🔹 Basit Örnek

*/
package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Yeni liste oluştur
	l := list.New()

	// Eleman ekleme
	l.PushBack("Go")        // sona
	l.PushBack("Python")
	l.PushFront("Java")     // başa
	l.PushBack("Rust")

	// Elemanları yazdır
	fmt.Println("Liste:")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// Ortaya eleman ekleme
	element := l.Front().Next() // Java'nın ardından Go var
	l.InsertAfter("C++", element)

	// Eleman silme
	l.Remove(element) // Go silinir

	fmt.Println("\nGüncellenmiş Liste:")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

/*
### 📌 Çıktı

```
Liste:
Java
Go
Python
Rust

Güncellenmiş Liste:
Java
C++
Python
Rust
```

---

## 🔹 Stack (Yığın) Örneği

*/
package main

import (
	"container/list"
	"fmt"
)

func main() {
	stack := list.New()

	// Push
	stack.PushBack(10)
	stack.PushBack(20)
	stack.PushBack(30)

	// Pop (sondan çıkar)
	for stack.Len() > 0 {
		e := stack.Back()
		fmt.Println("Pop:", e.Value)
		stack.Remove(e)
	}
}
/*

### 📌 Çıktı

```
Pop: 30
Pop: 20
Pop: 10
```

---

## 🔹 Queue (Kuyruk) Örneği
*/

package main

import (
	"container/list"
	"fmt"
)

func main() {
	queue := list.New()

	// Enqueue
	queue.PushBack("A")
	queue.PushBack("B")
	queue.PushBack("C")

	// Dequeue (baştan çıkar)
	for queue.Len() > 0 {
		e := queue.Front()
		fmt.Println("Dequeue:", e.Value)
		queue.Remove(e)
	}
}

/*
### 📌 Çıktı

```
Dequeue: A
Dequeue: B
Dequeue: C
```

---

## 🔹 Özet

Go’daki **`container/list`** paketi:

* Çift yönlü bağlı liste implementasyonu sağlar.
* **Eklemek/çıkarmak O(1)**, ama **arama O(n)**’dir.
* FIFO, LIFO ve **ortadan ekleme/silme işlemleri** için idealdir.
* Slice kadar yaygın kullanılmaz ama **özel veri yapıları** gerektiğinde oldukça faydalıdır.

---

👉 İstersen ben sana `list` paketini kullanarak **gerçek bir uygulama senaryosu** (mesela **tarayıcı geri-ileri geçmişi** ya da **LRU Cache**) örneği de gösterebilirim. İster misin?
*/
/*
O zaman  `container/list` paketini kullanarak **gerçek bir uygulama senaryosu** göstereyim:

👉 **LRU Cache (Least Recently Used Cache)**
Bu yapı çok kullanılan bir **cache algoritmasıdır**:

* Kapasiteyi dolduran en eski (en az kullanılan) eleman silinir.
* Yeni elemanlar eklenir.
* Kullanılan eleman liste başına taşınır.

Bunu `container/list` ile yapmak çok kolaydır çünkü:

* `list` sayesinde **en son kullanılanı başa alma** işlemi O(1)’dir.
* `map` ile birlikte kullanıldığında **arama O(1)** olur.

---

## 🚀 Örnek: LRU Cache
*/

package main

import (
	"container/list"
	"fmt"
)

// LRU Cache yapısı
type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

// Liste elemanı için struct
type entry struct {
	key   int
	value int
}

// Yeni cache oluştur
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

// Get -> cache’ten değer al
func (l *LRUCache) Get(key int) int {
	if elem, found := l.cache[key]; found {
		l.list.MoveToFront(elem) // En son kullanılan olarak işaretle
		return elem.Value.(*entry).value
	}
	return -1 // Bulunamadı
}

// Put -> cache’e değer ekle
func (l *LRUCache) Put(key int, value int) {
	if elem, found := l.cache[key]; found {
		// Güncelle ve en öne taşı
		elem.Value.(*entry).value = value
		l.list.MoveToFront(elem)
		return
	}

	// Yeni eleman ekle
	newElem := l.list.PushFront(&entry{key, value})
	l.cache[key] = newElem

	// Kapasite aşıldıysa en eskiyi sil
	if l.list.Len() > l.capacity {
		oldest := l.list.Back()
		if oldest != nil {
			l.list.Remove(oldest)
			delete(l.cache, oldest.Value.(*entry).key)
		}
	}
}

func main() {
	cache := NewLRUCache(3)

	cache.Put(1, 10)
	cache.Put(2, 20)
	cache.Put(3, 30)

	fmt.Println(cache.Get(1)) // 10 (bulundu)
	cache.Put(4, 40)          // kapasite aşıldı, en eski (2) silinir

	fmt.Println(cache.Get(2)) // -1 (silindi)
	fmt.Println(cache.Get(3)) // 30
	fmt.Println(cache.Get(4)) // 40
}
/*
---

## 📌 Çıktı

```
10
-1
30
40
```

---

✅ Bu örnekte:

* **`list`** → Elemanları kullanım sırasına göre tuttu.
* **`map`** → Elemanlara hızlı erişim sağladı.
* Kapasite dolunca `list.Back()` ile **en eski eleman O(1) sürede silindi**.

---

👉 İstersen aynı `list` paketini kullanarak bir **tarayıcı geçmişi (geri-ileri butonlarıyla gezinme)** örneği de gösterebilirim. Bunu da ister misin?
*/
/*
O zaman şimdi  `container/list` paketini kullanarak **tarayıcı geçmişi (geri–ileri gezinme)** uygulaması yapalım.

📌 Senaryo:

* Kullanıcı web sayfalarına giriyor → yeni sayfa **stack mantığında** tutuluyor.
* **Back** (geri) butonuna basınca, sayfa “geri listesine” atılıyor.
* **Forward** (ileri) butonuna basınca, geri listeden ileriye geçiliyor.

---

## 🚀 Örnek: Tarayıcı Geçmişi
*/

package main

import (
	"container/list"
	"fmt"
)

type Browser struct {
	current string     // Şu anki sayfa
	back    *list.List // Geri listesi
	forward *list.List // İleri listesi
}

// Yeni browser oluştur
func NewBrowser() *Browser {
	return &Browser{
		back:    list.New(),
		forward: list.New(),
	}
}

// Yeni sayfaya git
func (b *Browser) Visit(url string) {
	if b.current != "" {
		b.back.PushBack(b.current) // mevcut sayfayı geri listesine ekle
	}
	b.current = url
	b.forward.Init() // ileri listesi temizlenir
	fmt.Println("Ziyaret edildi:", url)
}

// Geri git
func (b *Browser) Back() {
	if b.back.Len() == 0 {
		fmt.Println("Geri gidilecek sayfa yok.")
		return
	}
	b.forward.PushFront(b.current)       // şimdiki sayfayı ileri listesine at
	last := b.back.Back()                // geri listesinin sonunu al
	b.current = last.Value.(string)      // oraya git
	b.back.Remove(last)                  // geri listesinden sil
	fmt.Println("Geri gidildi:", b.current)
}

// İleri git
func (b *Browser) Forward() {
	if b.forward.Len() == 0 {
		fmt.Println("İleri gidilecek sayfa yok.")
		return
	}
	b.back.PushBack(b.current)           // şimdiki sayfayı geri listesine at
	first := b.forward.Front()           // ileri listesinin başını al
	b.current = first.Value.(string)     // oraya git
	b.forward.Remove(first)              // ileri listesinden sil
	fmt.Println("İleri gidildi:", b.current)
}

func main() {
	b := NewBrowser()

	b.Visit("google.com")
	b.Visit("golang.org")
	b.Visit("github.com")

	b.Back()    // geri -> golang.org
	b.Back()    // geri -> google.com
	b.Forward() // ileri -> golang.org
	b.Visit("stackoverflow.com") // yeni ziyaret, ileri listesi sıfırlanır
	b.Back()    // geri -> golang.org
}

/*
---

## 📌 Çıktı

```
Ziyaret edildi: google.com
Ziyaret edildi: golang.org
Ziyaret edildi: github.com
Geri gidildi: golang.org
Geri gidildi: google.com
İleri gidildi: golang.org
Ziyaret edildi: stackoverflow.com
Geri gidildi: golang.org
```

---

✅ Burada:

* **back listesi (`list.List`)** → Geri gittiğimiz sayfaları saklıyor.
* **forward listesi (`list.List`)** → İleri gitmek için bekleyen sayfaları saklıyor.
* Yeni sayfa ziyaret edilince **ileri listesi temizleniyor** (tarayıcıda olduğu gibi).

*/