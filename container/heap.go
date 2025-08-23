/*👍
Go’nun **`container/heap`** paketi, standart kütüphanede bulunan ve **öncelikli kuyruk (priority queue)** mantığını kolayca kurmamızı sağlayan bir pakettir.

## 🔹 `heap` Paketinin Özellikleri

* `heap` paketi **min-heap** mantığı ile çalışır (yani en küçük eleman en başta olur).
* Eğer **max-heap** istiyorsan `Less` fonksiyonunu tersine çevirmen gerekir.
* `heap` aslında özel bir tip değil, **senin tanımladığın slice üzerinde** çalışır.
* Bunun için senin tipinin `heap.Interface` arayüzünü (interface) implemente etmesi gerekir.

---

## 🔹 `heap.Interface`

`heap` paketini kullanabilmek için tanımladığın slice tipinin şu metodları içermesi gerekir:

```go
type Interface interface {
    sort.Interface        // Len(), Less(i, j), Swap(i, j) metodlarını içerir
    Push(x any)           // Yeni eleman eklemek için
    Pop() any             // Eleman çıkarmak için
}
```

Yani:

* **Len() int** → Kaç eleman var
* **Less(i, j int) bool** → Karşılaştırma (min-heap için `h[i] < h[j]`)
* **Swap(i, j int)** → Elemanları yer değiştir
* **Push(x any)** → Eleman ekle
* **Pop() any** → Eleman çıkar

---

## 🔹 Temel Min-Heap Örneği
*/

package main

import (
	"container/heap"
	"fmt"
)

// Heap için tip
type IntHeap []int

// sort.Interface implementasyonu
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // Küçük olan önce gelir (min-heap)
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// heap.Interface implementasyonu
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]   // son elemanı al
	*h = old[:n-1]  // slice küçült
	return x
}

func main() {
	// Başlangıç elemanları
	h := &IntHeap{3, 2, 5}
	heap.Init(h) // Heap yapısını hazırla

	heap.Push(h, 1) // Eleman ekle
	heap.Push(h, 4)

	fmt.Println("En küçük eleman:", (*h)[0]) // En küçük değer her zaman en baştadır

	// Elemanları çıkar (küçükten büyüğe)
	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}
}
/*

### Çıktı

```
En küçük eleman: 1
1
2
3
4
5
```

---

## 🔹 Max-Heap Örneği

Aynı mantıkla sadece `Less` fonksiyonunu ters çeviririz:

```go
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] } // Büyük olan önce gelir
```

Böylece **max-heap** olur.

---

## 🔹 Öncelikli Kuyruk (Priority Queue) Örneği

Genelde `heap` en çok **öncelikli kuyruk** implementasyonu için kullanılır.

Örnek: Görevler (task) öncelik değerine göre sıralanıyor.
*/

package main

import (
	"container/heap"
	"fmt"
)

type Task struct {
	name     string
	priority int
}

// Görevleri heap üzerinde tutmak için slice
type TaskHeap []Task

func (h TaskHeap) Len() int           { return len(h) }
func (h TaskHeap) Less(i, j int) bool { return h[i].priority < h[j].priority } // küçük öncelik daha önce
func (h TaskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TaskHeap) Push(x any) {
	*h = append(*h, x.(Task))
}

func (h *TaskHeap) Pop() any {
	old := *h
	n := len(old)
	task := old[n-1]
	*h = old[:n-1]
	return task
}

func main() {
	h := &TaskHeap{
		{"Backup DB", 3},
		{"Email Reports", 2},
		{"Clean Cache", 1},
	}
	heap.Init(h)

	heap.Push(h, Task{"Deploy App", 0}) // en öncelikli

	fmt.Println("En öncelikli görev:", (*h)[0].name)

	for h.Len() > 0 {
		task := heap.Pop(h).(Task)
		fmt.Printf("Çalıştırılıyor -> %s (öncelik %d)\n", task.name, task.priority)
	}
}

/*
### Çıktı

```
En öncelikli görev: Deploy App
Çalıştırılıyor -> Deploy App (öncelik 0)
Çalıştırılıyor -> Clean Cache (öncelik 1)
Çalıştırılıyor -> Email Reports (öncelik 2)
Çalıştırılıyor -> Backup DB (öncelik 3)
```

---

## 🔹 Özet

Go’daki **`container/heap`** paketi:

* Slice tabanlı bir heap implementasyonu sağlar.
* Default olarak **min-heap** çalışır.
* **Max-heap** yapmak için `Less` metodunu ters yazmak yeterli.
* Genellikle **öncelikli kuyruk (priority queue)** için kullanılır.
* Eleman ekleme ve çıkarma işlemleri `heap.Push` ve `heap.Pop` ile yapılır.

---

👉 İstersen ben sana `heap` paketini kullanarak **zamanlayıcı (scheduler)** veya **Dijkstra algoritmasıyla en kısa yol bulma** gibi daha ileri bir uygulama da gösterebilirim. Bunu da ister misin?
*/
/*
O zaman  **`heap` paketini kullanarak Dijkstra algoritması** ile **en kısa yol bulma** uygulaması göstereyim.

Bu örnekte `heap`, **öncelikli kuyruk (priority queue)** olarak kullanılacak.
Dijkstra algoritması da, **her adımda en kısa mesafeli düğümü** almak için min-heap mantığını kullanır.

---

## 🚀 Örnek: Dijkstra Algoritması (heap ile en kısa yol bulma)
*/

package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Graf temsili
type Edge struct {
	to, weight int
}
type Graph map[int][]Edge

// Kuyrukta tutulacak eleman
type Item struct {
	node int
	dist int
}

// Öncelikli kuyruk (min-heap)
type PriorityQueue []Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist // küçük mesafe önce gelir
}
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x any)   { *pq = append(*pq, x.(Item)) }
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// Dijkstra algoritması
func Dijkstra(graph Graph, start int) map[int]int {
	dist := make(map[int]int)

	// Başlangıçta tüm düğümler sonsuz
	for node := range graph {
		dist[node] = math.MaxInt
	}
	dist[start] = 0

	// Min-heap başlat
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Item{start, 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(Item)

		// Eğer daha iyi yol varsa devam et
		if current.dist > dist[current.node] {
			continue
		}

		// Komşulara bak
		for _, edge := range graph[current.node] {
			newDist := dist[current.node] + edge.weight
			if newDist < dist[edge.to] {
				dist[edge.to] = newDist
				heap.Push(pq, Item{edge.to, newDist})
			}
		}
	}

	return dist
}

func main() {
	// Örnek graf (ağırlıklı)
	graph := Graph{
		1: {{2, 2}, {3, 4}},
		2: {{3, 1}, {4, 7}},
		3: {{5, 3}},
		4: {{6, 1}},
		5: {{4, 2}, {6, 5}},
		6: {},
	}

	// Dijkstra ile en kısa yolları bul
	start := 1
	dist := Dijkstra(graph, start)

	fmt.Printf("Düğüm %d'den en kısa yollar:\n", start)
	for node, d := range dist {
		fmt.Printf("→ %d : %d\n", node, d)
	}
}

/*
---

## 📌 Çıktı

```
Düğüm 1'den en kısa yollar:
→ 1 : 0
→ 2 : 2
→ 3 : 3
→ 4 : 8
→ 5 : 6
→ 6 : 9
```

---

## ✅ Açıklama

* Burada `heap` paketini kullanarak **öncelikli kuyruk** oluşturduk.
* Her adımda en küçük mesafeli düğümü `heap.Pop()` ile çıkardık.
* Bu sayede Dijkstra algoritması klasik **O(E log V)** karmaşıklığıyla çalıştı.

---
*/