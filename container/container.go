/*
Golang’daki **`container`** paketi, standart kütüphanede bulunan ve temel veri yapılarının (listeler, yığınlar, öncelikli kuyruklar) implementasyonunu sağlayan bir pakettir.
`container` paketi tek başına kullanılmaz, bunun altında üç alt paket bulunur:

1. **`container/list`** → Çift yönlü bağlı liste (doubly linked list)
2. **`container/heap`** → Min-heap ve max-heap mantığında çalışan veri yapısı
3. **`container/ring`** → Dairesel liste (circular list)

Şimdi her birini **örneklerle** açıklayayım 👇

---

## 1. `container/list` (Bağlı Liste)

* Çift yönlü bağlı liste sağlar.
* Dinamik eleman ekleme ve silme işlemlerinde avantajlıdır.
* Kullanımı slice’dan farklıdır, çünkü elemanlara index ile erişilemez.

### Örnek:
*/

package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Yeni bir liste oluştur
	l := list.New()

	// Listenin sonuna ekleme
	l.PushBack("Go")
	l.PushBack("Python")

	// Listenin başına ekleme
	l.PushFront("Java")

	// Belirli bir elemanın arkasına ekleme
	element := l.PushBack("C++")
	l.InsertAfter("Rust", element)

	// Elemanları dolaşma
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// Eleman silme
	l.Remove(element)

	fmt.Println("Silindikten sonra:")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

/*
🔹 Çıktı:

```
Java
Go
Python
C++
Rust
Silindikten sonra:
Java
Go
Python
Rust
```

---

## 2. `container/heap` (Heap – Öncelikli Kuyruk)

* Heap, **öncelikli kuyruk (priority queue)** oluşturmak için kullanılır.
* Varsayılan olarak **min-heap** mantığıyla çalışır (küçük olan önce çıkar).
* Max-heap için karşılaştırma fonksiyonu ters yazılır.

### Örnek (Min-Heap):
*/
package main

import (
	"container/heap"
	"fmt"
)

// Öncelikli kuyruk için bir tip oluşturuyoruz
type IntHeap []int

// heap.Interface için gerekli metotlar
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-heap
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &IntHeap{3, 2, 5}
	heap.Init(h)         // Heap başlat
	heap.Push(h, 1)      // Eleman ekle
	heap.Push(h, 4)

	fmt.Println("Min:", (*h)[0]) // En küçük eleman

	for h.Len() > 0 {
		fmt.Println(heap.Pop(h)) // Küçükten büyüğe çıkarır
	}
}
/*

🔹 Çıktı:

```
Min: 1
1
2
3
4
5
```

---

## 3. `container/ring` (Dairesel Liste)

* **Ring buffer / circular list** implementasyonu sağlar.
* Elemanlar dairesel şekilde birbirine bağlanır, yani son elemana gelince tekrar başa döner.
* Döngüsel yapılar için uygundur (örn. round-robin algoritması).

### Örnek:
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

	// Elemanları dolaş
	r.Do(func(x any) {
		fmt.Print(x, " ")
	})
	fmt.Println()

	// İleri hareket etme
	fmt.Println("Şu an:", r.Value)       // 1
	fmt.Println("2 ileri:", r.Move(2).Value) // 3
	fmt.Println("Geri:", r.Move(-1).Value)   // 5
}

/*
🔹 Çıktı:

```
1 2 3 4 5 
Şu an: 1
2 ileri: 3
Geri: 5
```

---

## Özet

Go’daki **`container`** paketi 3 alt yapı sunar:

* **`list`** → Çift yönlü bağlı liste
* **`heap`** → Öncelikli kuyruk (min-heap/max-heap)
* **`ring`** → Dairesel liste

📌 `container` paketini slice veya map gibi çok sık görmezsin, ama düşük seviyeli ve özel veri yapıları gerektiğinde oldukça faydalıdır.

---

İstersen sana **bunların hepsini kapsayan tek bir örnek proje** yazayım (list, heap ve ring aynı programda). İstiyor musun?
*/

/*
 O zaman  **tek bir Go projesi** içinde hem `list`, hem `heap`, hem de `ring` kullanımını gösteren kapsamlı bir örnek yazayım.

Bu örnekte senaryo şu olacak:

* **Bir görev listesi (`list`)** tutacağız.
* Görevlerin önceliğini **öncelikli kuyruk (`heap`)** ile yöneteceğiz.
* Görevleri çalıştırırken **round-robin (`ring`)** algoritmasıyla döngüsel sırada çalıştıracağız.

---

## Kapsamlı Örnek: Görev Yöneticisi
*/
package main

import (
	"container/heap"
	"container/list"
	"container/ring"
	"fmt"
)

///////////////////////
// HEAP (Priority Queue)
///////////////////////

type Task struct {
	name     string
	priority int
}

type TaskHeap []Task

func (h TaskHeap) Len() int           { return len(h) }
func (h TaskHeap) Less(i, j int) bool { return h[i].priority < h[j].priority } // Küçük öncelikli
func (h TaskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TaskHeap) Push(x any) {
	*h = append(*h, x.(Task))
}

func (h *TaskHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

///////////////////////
// MAIN
///////////////////////

func main() {
	// 1) LIST -> Görev listesi
	taskList := list.New()
	taskList.PushBack(Task{"Backup DB", 2})
	taskList.PushBack(Task{"Email Reports", 3})
	taskList.PushBack(Task{"Clean Cache", 1})

	fmt.Println("=== Görev Listesi (list) ===")
	for e := taskList.Front(); e != nil; e = e.Next() {
		t := e.Value.(Task)
		fmt.Printf("Görev: %s (öncelik %d)\n", t.name, t.priority)
	}

	// 2) HEAP -> Öncelikli kuyruk
	h := &TaskHeap{}
	heap.Init(h)

	for e := taskList.Front(); e != nil; e = e.Next() {
		heap.Push(h, e.Value.(Task))
	}

	fmt.Println("\n=== Öncelikli Kuyruk (heap) ===")
	for h.Len() > 0 {
		task := heap.Pop(h).(Task)
		fmt.Printf("Çalıştırılıyor -> %s (öncelik %d)\n", task.name, task.priority)
	}

	// 3) RING -> Round-robin görev çalıştırma
	fmt.Println("\n=== Round-Robin (ring) ===")
	r := ring.New(3)
	r.Value = "Server-1"
	r.Next().Value = "Server-2"
	r.Next().Next().Value = "Server-3"

	start := r
	for i := 0; i < 6; i++ { // 6 tur döndürelim
		fmt.Printf("Görev çalıştırılıyor -> %v\n", start.Value)
		start = start.Next()
	}
}

/*
---

## Çıktı

```
=== Görev Listesi (list) ===
Görev: Backup DB (öncelik 2)
Görev: Email Reports (öncelik 3)
Görev: Clean Cache (öncelik 1)

=== Öncelikli Kuyruk (heap) ===
Çalıştırılıyor -> Clean Cache (öncelik 1)
Çalıştırılıyor -> Backup DB (öncelik 2)
Çalıştırılıyor -> Email Reports (öncelik 3)

=== Round-Robin (ring) ===
Görev çalıştırılıyor -> Server-1
Görev çalıştırılıyor -> Server-2
Görev çalıştırılıyor -> Server-3
Görev çalıştırılıyor -> Server-1
Görev çalıştırılıyor -> Server-2
Görev çalıştırılıyor -> Server-3
```

---

✅ Böylece tek bir programda `list`, `heap` ve `ring` paketlerinin hepsini kullanmış olduk.
Bu yapılar sayesinde:

* **list** → dinamik görev listesi,
* **heap** → öncelik bazlı seçim,
* **ring** → round-robin çalıştırma mantığı kurulmuş oldu.

---

İstersen ben bunu sana **gerçek hayat senaryosu** (örneğin "CPU Scheduler" veya "Job Queue Manager") şeklinde daha da geliştirebilirim. İster misin?
*/