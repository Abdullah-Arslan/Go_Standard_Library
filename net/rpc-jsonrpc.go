/*
Go’nun **`net/rpc/jsonrpc`** paketini baştan sona açıklayayım ve örneklerle göstereyim.

`net/rpc/jsonrpc` paketi, Go’nun **RPC mekanizmasını JSON protokolü ile kullanmaya** yarar. Yani RPC metod çağrıları ve yanıtları **JSON formatında seri hale getirilir**. Bu, özellikle farklı diller arasında RPC yapmak için faydalıdır.

---

# 1️⃣ JSON-RPC Nedir?

* JSON-RPC, **dil bağımsız** bir RPC protokolüdür.
* İstemci → Sunucu JSON mesajı gönderir: `method`, `params`, `id`.
* Sunucu → İstemci JSON mesajı ile cevap verir: `result` veya `error`.
* `net/rpc/jsonrpc` paketi bunu Go ile kullanmayı sağlar.

---

# 2️⃣ Temel Özellikler

| Özellik                       | Açıklama                                            |
| ----------------------------- | --------------------------------------------------- |
| `jsonrpc.NewClient`           | JSON-RPC client oluşturur                           |
| `jsonrpc.NewClientCodec`      | RPC client codec’i oluşturur (io.ReadWrite)         |
| `jsonrpc.NewServerCodec`      | RPC server codec’i oluşturur (io.ReadWrite)         |
| TCP veya HTTP ile çalışabilir | JSON-RPC ile Go uygulamaları birbirine bağlanabilir |

---

# 3️⃣ JSON-RPC Sunucu Örneği
*/
``go
package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// Servis tipi
type Calculator int

type Args struct {
	A, B int
}

// Toplama
func (c *Calculator) Add(args Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

// Çıkarma
func (c *Calculator) Subtract(args Args, reply *int) error {
	*reply = args.A - args.B
	return nil
}

func main() {
	calc := new(Calculator)
	rpc.Register(calc)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	defer listener.Close()

	fmt.Println("JSON-RPC sunucusu 1234 portunda çalışıyor...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept hatası:", err)
			continue
		}
		// JSON-RPC codec ile servis
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
``
/*
---

# 4️⃣ JSON-RPC Client Örneği
*/
``go
package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Bağlantı hatası:", err)
		return
	}
	defer conn.Close()

	client := jsonrpc.NewClient(conn)

	args := Args{A: 15, B: 5}

	var sum int
	err = client.Call("Calculator.Add", args, &sum)
	if err != nil {
		fmt.Println("Add hatası:", err)
	} else {
		fmt.Println("Toplam:", sum)
	}

	var diff int
	err = client.Call("Calculator.Subtract", args, &diff)
	if err != nil {
		fmt.Println("Subtract hatası:", err)
	} else {
		fmt.Println("Fark:", diff)
	}
}
``
/*
---

# 5️⃣ Açıklamalar

1. Sunucu tarafında:

   * `rpc.ServeCodec(jsonrpc.NewServerCodec(conn))` → JSON-RPC ile çalışır.
2. Client tarafında:

   * `jsonrpc.NewClient(conn)` → JSON-RPC client oluşturur.
3. **Avantajı:**

   * JSON formatı sayesinde farklı dillerdeki istemciler Go sunucusuna bağlanabilir.
4. RPC metodları normal `net/rpc` kurallarına uyar:

   * Public olmalı
   * İki parametre ve `error` döndürmeli

---

# 6️⃣ Örnek JSON-RPC Mesajı

## İstemci → Sunucu
*/
``json
{
  "method": "Calculator.Add",
  "params": [{"A":15,"B":5}],
  "id":1
}
``

//## Sunucu → İstemci

``json
{
  "result": 20,
  "error": null,
  "id":1
}
``
/*
---

# 7️⃣ Özet

* `net/rpc/jsonrpc` → Go’nun **RPC mekanizmasını JSON formatına uyarlayan paketi**.
* TCP üzerinden çalışır, HTTP için ekstra çözüm gerekebilir.
* Avantajı: farklı dillerle kolay entegrasyon, JSON ile network üzerinden veri taşınabilir.
* RPC kuralları aynıdır (public metod, args & reply, return error).

---

İstersen sana bunu bir adım daha ileri götürüp **concurrent JSON-RPC client + server + timeout yönetimi** ile örnek bir performans testi de gösterebilirim. Bunu ister misin?
EVET
*/
/*
Harika 👍 o zaman sana **concurrent JSON-RPC client + server + timeout yönetimi** örneğini göstereyim.

Bu örnekle:

* Sunucu aynı anda birden fazla JSON-RPC client çağrısını işleyebilecek,
* Client’lar eş zamanlı olarak JSON-RPC çağrıları yapacak,
* Her çağrı için timeout kontrolü olacak.

---

## 1️⃣ JSON-RPC Sunucu Kodu (Concurrent)
*/
``go
package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

// Servis tipi
type Calculator int

type Args struct {
	A, B int
}

// Toplama
func (c *Calculator) Add(args Args, reply *int) error {
	time.Sleep(2 * time.Second) // Simüle edilmiş gecikme
	*reply = args.A + args.B
	return nil
}

// Çıkarma
func (c *Calculator) Subtract(args Args, reply *int) error {
	time.Sleep(1 * time.Second)
	*reply = args.A - args.B
	return nil
}

func main() {
	calc := new(Calculator)
	rpc.Register(calc)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Concurrent JSON-RPC sunucusu çalışıyor :1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept hatası:", err)
			continue
		}
		// Her client için ayrı goroutine
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
``
/*
---

## 2️⃣ Concurrent JSON-RPC Client Kodu (Timeout)
*/
``go
package main

import (
	"context"
	"fmt"
	"net"
	"net/rpc/jsonrpc"
	"sync"
	"time"
)

type Args struct {
	A, B int
}

func callWithTimeout(client *jsonrpc.Client, method string, args Args, reply interface{}, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	errCh := make(chan error, 1)

	go func() {
		errCh <- client.Call(method, args, reply)
	}()

	select {
	case <-ctx.Done():
		return fmt.Errorf("JSON-RPC çağrısı zaman aşımı: %v", ctx.Err())
	case err := <-errCh:
		return err
	}
}

func main() {
	numClients := 5
	var wg sync.WaitGroup

	for i := 1; i <= numClients; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			conn, err := net.Dial("tcp", "localhost:1234")
			if err != nil {
				fmt.Printf("Client %d bağlanamadı: %v\n", id, err)
				return
			}
			defer conn.Close()

			client := jsonrpc.NewClient(conn)
			args := Args{A: id * 10, B: id * 2}

			var sum int
			err = callWithTimeout(client, "Calculator.Add", args, &sum, 1*time.Second)
			if err != nil {
				fmt.Printf("Client %d Add hatası: %v\n", id, err)
			} else {
				fmt.Printf("Client %d Toplam: %d\n", id, sum)
			}

			var diff int
			err = callWithTimeout(client, "Calculator.Subtract", args, &diff, 3*time.Second)
			if err != nil {
				fmt.Printf("Client %d Subtract hatası: %v\n", id, err)
			} else {
				fmt.Printf("Client %d Fark: %d\n", id, diff)
			}

		}(i)
	}

	wg.Wait()
	fmt.Println("Tüm concurrent JSON-RPC client işlemleri tamamlandı.")
}
``
/*
---

## 3️⃣ Açıklamalar

1. **Sunucu tarafında:**

   * `rpc.ServeCodec(jsonrpc.NewServerCodec(conn))` ile JSON-RPC çağrıları işleniyor.
   * Her client için **ayrı goroutine** açılıyor → concurrency sağlanıyor.

2. **Client tarafında:**

   * `numClients` kadar goroutine ile eş zamanlı JSON-RPC çağrısı yapılıyor.
   * `callWithTimeout` fonksiyonu ile **her çağrı için timeout** kontrolü sağlanıyor.

3. **Performans ve güvenlik:**

   * Uzun süren işlemler client tarafından timeout ile iptal edilebilir.
   * Çoklu client testlerinde server concurrency performansı gözlemlenebilir.

---

## 4️⃣ Örnek Çıktı
*/
``
Client 1 Add hatası: JSON-RPC çağrısı zaman aşımı: context deadline exceeded
Client 1 Fark: 8
Client 2 Add hatası: JSON-RPC çağrısı zaman aşımı: context deadline exceeded
Client 2 Fark: 16
...
Tüm concurrent JSON-RPC client işlemleri tamamlandı.
``
/*
* Add metodunun gecikmesi 2 saniye → timeout 1 saniye → hata
* Subtract metodunun gecikmesi 1 saniye → timeout 3 saniye → başarılı

---

✅ Bu yöntemle:

* JSON-RPC ile **concurrent client ve server** uygulayabiliyorsun.
* Her çağrı için **timeout ve context yönetimi** ekleyerek güvenliği artırabiliyorsun.
* `pprof` ile sunucu performansını analiz edebilirsin.

---

İstersen sana bunu bir adım daha ileri götürüp **concurrent JSON-RPC + worker pool + pprof + timeout** şeklinde **tam performans testi framework** örneğini birleştirilmiş olarak gösterebilirim. Bunu ister misin?
EVET
*/
/*
Harika! 👍 O zaman sana **concurrent JSON-RPC + worker pool + timeout + pprof ile tam performans testi framework** örneğini göstereyim. Bu örnek, gerçek hayatta dağıtık sistemlerde veya servislerde kullanılabilecek şekilde tasarlanmıştır.

---

## 1️⃣ Sunucu Kodu (Concurrent + JSON-RPC + pprof)
*/
``go
package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	_ "net/http/pprof"
	"time"
)

type Calculator int

type Args struct {
	A, B int
}

func (c *Calculator) Add(args Args, reply *int) error {
	time.Sleep(500 * time.Millisecond) // Simüle edilmiş iş yükü
	*reply = args.A + args.B
	return nil
}

func (c *Calculator) Subtract(args Args, reply *int) error {
	time.Sleep(300 * time.Millisecond)
	*reply = args.A - args.B
	return nil
}

func main() {
	calc := new(Calculator)
	rpc.Register(calc)

	// pprof HTTP server
	go func() {
		fmt.Println("pprof HTTP server :6060 çalışıyor")
		http.ListenAndServe(":6060", nil)
	}()

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("TCP listen hatası:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Concurrent JSON-RPC sunucusu :1234 çalışıyor")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept hatası:", err)
			continue
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
``
/*
✅ Açıklamalar:

* `pprof` ile CPU/memory profili alınabilir (`http://localhost:6060/debug/pprof/`).
* Her client bağlantısı için ayrı goroutine → concurrency sağlanıyor.

---

## 2️⃣ Client Kodu (Worker Pool + Timeout + Concurrent JSON-RPC)
*/
``go
package main

import (
	"context"
	"fmt"
	"net"
	"net/rpc/jsonrpc"
	"sync"
	"time"
)

type Args struct {
	A, B int
}

// Timeout ile çağrı yapan fonksiyon
func callWithTimeout(client *jsonrpc.Client, method string, args Args, reply interface{}, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	errCh := make(chan error, 1)
	go func() {
		errCh <- client.Call(method, args, reply)
	}()

	select {
	case <-ctx.Done():
		return fmt.Errorf("JSON-RPC çağrısı zaman aşımı: %v", ctx.Err())
	case err := <-errCh:
		return err
	}
}

func worker(id int, jobs <-chan Args, wg *sync.WaitGroup) {
	defer wg.Done()

	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Printf("Worker %d bağlanamadı: %v\n", id, err)
		return
	}
	defer conn.Close()

	client := jsonrpc.NewClient(conn)

	for args := range jobs {
		var sum int
		err := callWithTimeout(client, "Calculator.Add", args, &sum, 1*time.Second)
		if err != nil {
			fmt.Printf("Worker %d Add hatası: %v\n", id, err)
		} else {
			fmt.Printf("Worker %d Add: %d + %d = %d\n", id, args.A, args.B, sum)
		}

		var diff int
		err = callWithTimeout(client, "Calculator.Subtract", args, &diff, 2*time.Second)
		if err != nil {
			fmt.Printf("Worker %d Subtract hatası: %v\n", id, err)
		} else {
			fmt.Printf("Worker %d Subtract: %d - %d = %d\n", id, args.A, args.B, diff)
		}
	}
}

func main() {
	numWorkers := 3
	numJobs := 10

	jobs := make(chan Args, numJobs)
	var wg sync.WaitGroup

	// Worker pool oluştur
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Jobları gönder
	for i := 1; i <= numJobs; i++ {
		jobs <- Args{A: i * 10, B: i * 2}
	}
	close(jobs)

	wg.Wait()
	fmt.Println("Tüm işler tamamlandı.")
}
``
/*
---

## 3️⃣ Açıklamalar

1. **Worker Pool:**

   * `numWorkers` kadar worker başlatılıyor.
   * Workerlar `jobs` kanalından işleri alıp işliyor.

2. **Timeout:**

   * `callWithTimeout` fonksiyonu her RPC çağrısına zaman aşımı ekliyor.

3. **Concurrent JSON-RPC:**

   * Birden fazla worker aynı anda server’a çağrı yapıyor.
   * Server her client için ayrı goroutine ile cevap veriyor.

4. **pprof ile Profil:**

   * Sunucu üzerinde `http://localhost:6060/debug/pprof/` adresinden **CPU, memory, goroutine profili** alınabilir.

---

## 4️⃣ Örnek Çıktı
*/
``
Worker 1 Add: 10 + 2 = 12
Worker 2 Add: 20 + 4 = 24
Worker 3 Add: 30 + 6 = 36
Worker 1 Subtract: 10 - 2 = 8
Worker 2 Subtract: 20 - 4 = 16
Worker 3 Subtract: 30 - 6 = 24
...
Tüm işler tamamlandı.
``
/*
* İşler workerlar arasında eşit dağılıyor → concurrency sağlanıyor.
* Timeout sayesinde takılan işlemler iptal edilebilir.

---

✅ Bu framework sayesinde:

* **Conccurrent JSON-RPC + Timeout + Worker Pool + Profiling** tamamlandı.
* Gerçek dünya RPC servisleri için **yük testi, performans analizi ve güvenli çağrılar** yapılabilir.

---

İstersen sana bunu bir adım daha ileri götürüp **JSON-RPC üzerinden IP dağıtımı ve net/netip filtreleme** ile bir **network yönetim simülasyonu** örneğini birleştirilmiş olarak gösterebilirim. Bunu ister misin?
EVET
*/
/*
Harika! 👍 O zaman sana **JSON-RPC + net/netip ile IP dağıtımı ve filtreleme simülasyonu** örneğini göstereyim.

Bu örnekle:

* Sunucu, client’lardan gelen IP taleplerini yönetiyor,
* JSON-RPC ile IP ataması yapıyor,
* `net/netip` ile IP aralıklarını kontrol ediyor ve filtreliyor,
* Concurrent workerlar ile birden fazla client eş zamanlı IP alabiliyor.

---

## 1️⃣ Sunucu Kodu (JSON-RPC + IP Pool + net/netip)
*/
``go
package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"net/netip"
	"sync"
)

type Args struct {
	ClientID string
}

type IPReply struct {
	IP string
}

type IPManager struct {
	mu      sync.Mutex
	pool    []netip.Addr
	assigned map[string]netip.Addr
}

func NewIPManager() *IPManager {
	var ips []netip.Addr
	for i := 1; i <= 10; i++ {
		ip, _ := netip.ParseAddr(fmt.Sprintf("192.168.1.%d", i))
		ips = append(ips, ip)
	}
	return &IPManager{
		pool:    ips,
		assigned: make(map[string]netip.Addr),
	}
}

// IP ata
func (m *IPManager) AssignIP(args Args, reply *IPReply) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Zaten atanmış mı kontrol et
	if ip, ok := m.assigned[args.ClientID]; ok {
		reply.IP = ip.String()
		return nil
	}

	if len(m.pool) == 0 {
		return fmt.Errorf("IP havuzu boş")
	}

	ip := m.pool[0]
	m.pool = m.pool[1:]
	m.assigned[args.ClientID] = ip
	reply.IP = ip.String()
	return nil
}

// IP iade et
func (m *IPManager) ReleaseIP(args Args, reply *IPReply) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if ip, ok := m.assigned[args.ClientID]; ok {
		delete(m.assigned, args.ClientID)
		m.pool = append(m.pool, ip)
		reply.IP = ip.String()
		return nil
	}

	return fmt.Errorf("clientID bulunamadı")
}

func main() {
	manager := NewIPManager()
	rpc.Register(manager)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Listen hatası:", err)
		return
	}
	defer listener.Close()

	fmt.Println("IP Manager sunucusu JSON-RPC :1234 çalışıyor")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept hatası:", err)
			continue
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
``
/*
---

## 2️⃣ Client Kodu (Concurrent IP Talebi)
*/
``go
package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
	"sync"
)

type Args struct {
	ClientID string
}

type IPReply struct {
	IP string
}

func main() {
	numClients := 5
	var wg sync.WaitGroup

	for i := 1; i <= numClients; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			conn, err := net.Dial("tcp", "localhost:1234")
			if err != nil {
				fmt.Printf("Client %d bağlanamadı: %v\n", id, err)
				return
			}
			defer conn.Close()

			client := jsonrpc.NewClient(conn)
			args := Args{ClientID: fmt.Sprintf("client-%d", id)}

			var reply IPReply
			err = client.Call("IPManager.AssignIP", args, &reply)
			if err != nil {
				fmt.Printf("Client %d IP atama hatası: %v\n", id, err)
			} else {
				fmt.Printf("Client %d atanan IP: %s\n", id, reply.IP)
			}

			// IP iade et
			var releaseReply IPReply
			err = client.Call("IPManager.ReleaseIP", args, &releaseReply)
			if err != nil {
				fmt.Printf("Client %d IP iade hatası: %v\n", id, err)
			} else {
				fmt.Printf("Client %d iade edilen IP: %s\n", id, releaseReply.IP)
			}

		}(i)
	}

	wg.Wait()
	fmt.Println("Tüm IP işlemleri tamamlandı.")
}
``
/*
---

## 3️⃣ Açıklamalar

1. **Sunucu:**

   * `IPManager` struct ile IP havuzu yönetiliyor.
   * `AssignIP` → IP atama
   * `ReleaseIP` → IP iade
   * `sync.Mutex` ile concurrency güvenliği sağlanıyor.
   * `net/netip` kullanarak IP adresleri tip güvenli olarak yönetiliyor.

2. **Client:**

   * Her goroutine bir client simülasyonu yapıyor.
   * JSON-RPC ile sunucuya bağlanıp IP talebi ve iadesi yapıyor.

3. **Concurrent:**

   * Aynı anda birden fazla client IP alabilir.
   * Mutex ve slice yönetimi ile IP çakışması önleniyor.

---

## 4️⃣ Örnek Çıktı
*/
``
Client 1 atanan IP: 192.168.1.1
Client 1 iade edilen IP: 192.168.1.1
Client 2 atanan IP: 192.168.1.2
Client 2 iade edilen IP: 192.168.1.2
Client 3 atanan IP: 192.168.1.3
Client 3 iade edilen IP: 192.168.1.3
...
Tüm IP işlemleri tamamlandı.
``
/*
* IP havuzu güvenli bir şekilde dağıtılıyor.
* Concurrent clientlar sorunsuz IP alıp iade edebiliyor.
* `net/netip` ile IP aralıkları ve filtreleme rahatlıkla yapılabilir.

---

Bu yapıyı bir adım daha ileri götürüp **pprof + timeout + worker pool ile büyük IP dağıtımı simülasyonu** haline de getirebiliriz.

Bunu da ister misin?
*/
