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
/*
O zaman az önceki örneği **gerçek hayattaki bir CPU Scheduler (işlemci görev planlayıcı)** gibi tasarlayalım.

📌 Senaryomuz:

* CPU’ya gelen **işler (tasks)** bir **görev kuyruğuna (list)** kaydedilir.
* İşler **öncelik (heap)** kullanılarak sıralanır. Yani acil işler önce işlenir.
* CPU, işleri **round-robin (ring)** algoritmasıyla mevcut çekirdekler arasında dağıtır.

---

## Örnek: CPU Scheduler
*/

package main

import (
	"container/heap"
	"container/list"
	"container/ring"
	"fmt"
)

///////////////////////
// TASK ve HEAP (Priority Queue)
///////////////////////

type Task struct {
	name     string
	priority int // Küçük sayı = yüksek öncelik
}

type TaskHeap []Task

func (h TaskHeap) Len() int           { return len(h) }
func (h TaskHeap) Less(i, j int) bool { return h[i].priority < h[j].priority }
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
	// 1) Görev kuyruğu (list)
	taskList := list.New()
	taskList.PushBack(Task{"Kritik Güvenlik Yaması", 1})
	taskList.PushBack(Task{"Veritabanı Yedekleme", 3})
	taskList.PushBack(Task{"Log Temizleme", 5})
	taskList.PushBack(Task{"E-posta Raporlama", 4})
	taskList.PushBack(Task{"Makine Öğrenmesi Eğitimi", 2})

	fmt.Println("=== Gelen Görevler (list) ===")
	for e := taskList.Front(); e != nil; e = e.Next() {
		t := e.Value.(Task)
		fmt.Printf("Görev: %-25s Öncelik: %d\n", t.name, t.priority)
	}

	// 2) Öncelikli kuyruğa (heap) yükle
	h := &TaskHeap{}
	heap.Init(h)

	for e := taskList.Front(); e != nil; e = e.Next() {
		heap.Push(h, e.Value.(Task))
	}

	// 3) CPU çekirdeklerini (ring) oluştur
	cores := ring.New(3) // 3 çekirdekli CPU
	cores.Value = "CPU-Core-1"
	cores.Next().Value = "CPU-Core-2"
	cores.Next().Next().Value = "CPU-Core-3"

	// 4) Görevleri çekirdeklere round-robin ile dağıt
	fmt.Println("\n=== CPU Scheduler ===")
	core := cores
	for h.Len() > 0 {
		task := heap.Pop(h).(Task)
		fmt.Printf("[%-10s] çalıştırıyor -> %-25s (öncelik %d)\n",
			core.Value, task.name, task.priority)
		core = core.Next()
	}
}

/*
---

## Çıktı

```
=== Gelen Görevler (list) ===
Görev: Kritik Güvenlik Yaması      Öncelik: 1
Görev: Veritabanı Yedekleme        Öncelik: 3
Görev: Log Temizleme               Öncelik: 5
Görev: E-posta Raporlama           Öncelik: 4
Görev: Makine Öğrenmesi Eğitimi    Öncelik: 2

=== CPU Scheduler ===
[CPU-Core-1] çalıştırıyor -> Kritik Güvenlik Yaması      (öncelik 1)
[CPU-Core-2] çalıştırıyor -> Makine Öğrenmesi Eğitimi    (öncelik 2)
[CPU-Core-3] çalıştırıyor -> Veritabanı Yedekleme        (öncelik 3)
[CPU-Core-1] çalıştırıyor -> E-posta Raporlama           (öncelik 4)
[CPU-Core-2] çalıştırıyor -> Log Temizleme               (öncelik 5)
```

---

✅ Burada:

* **list** → gelen görev kuyruğunu temsil etti.
* **heap** → önceliğe göre hangi işin önce yapılacağını seçti.
* **ring** → görevlerin çekirdekler arasında round-robin ile dağıtılmasını sağladı.

---
*/