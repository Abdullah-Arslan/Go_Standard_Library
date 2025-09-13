/*
Go’nun **`net/rpc`** paketini baştan sona açıklayayım ve örneklerle anlatayım.

`net/rpc` paketi, **Remote Procedure Call (RPC)** yani **uzaktan prosedür çağrısı** yapmak için kullanılır. Bir program başka bir programın fonksiyonunu sanki kendi içinde çağırıyormuş gibi kullanabilir.

---

# 1️⃣ `net/rpc` Paketinin Temel Amaçları

* Fonksiyonları **uzaktaki bir sunucuda çalıştırmak**
* Fonksiyonların **parametre ve dönüş değerlerini seri hale getirip** (marshal) network üzerinden göndermek
* TCP, HTTP veya kendi transport layer’ını kullanmak
* Synchronous ve asynchronous RPC desteği sağlamak

---

# 2️⃣ RPC Sunucusu Örneği

## Adım 1: Servis Tipi Tanımlama

RPC servisleri, **public tipler** olmalı ve **public metodlara** sahip olmalı.
Metotlar **iki parametre ve bir error** döndürmeli.
*/
``go
package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

// Servis tipi
type Calculator int

// Toplama fonksiyonu
// Args ve Reply struct tipinde olmalı
type Args struct {
	A, B int
}

func (c *Calculator) Add(args Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

// Çıkarma fonksiyonu
func (c *Calculator) Subtract(args Args, reply *int) error {
	*reply = args.A - args.B
	return nil
}

// Bölme fonksiyonu (error örneği)
func (c *Calculator) Divide(args Args, reply *float64) error {
	if args.B == 0 {
		return errors.New("sıfıra bölme hatası")
	}
	*reply = float64(args.A) / float64(args.B)
	return nil
}
``
/*
---

## Adım 2: RPC Sunucusunu Başlat
*/
``go
func main() {
	calc := new(Calculator)
	rpc.Register(calc) // Servisi kaydet

	// TCP üzerinden dinle
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	defer listener.Close()

	fmt.Println("RPC sunucusu 1234 portunda çalışıyor...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept hatası:", err)
			continue
		}
		go rpc.ServeConn(conn) // Her bağlantı için ayrı goroutine
	}
}
``
/*
---

# 3️⃣ RPC Client Örneği
*/
``go
package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	A, B int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Bağlantı hatası:", err)
		return
	}
	defer client.Close()

	args := Args{A: 10, B: 5}

	// Add
	var sum int
	err = client.Call("Calculator.Add", args, &sum)
	if err != nil {
		fmt.Println("Hata:", err)
	} else {
		fmt.Println("Toplam:", sum)
	}

	// Divide
	var div float64
	err = client.Call("Calculator.Divide", args, &div)
	if err != nil {
		fmt.Println("Hata:", err)
	} else {
		fmt.Println("Bölme:", div)
	}
}
``
/*
---

# 4️⃣ Önemli Kurallar

1. RPC metodları **public olmalı** (baş harfi büyük).
2. RPC metodları **iki parametre ve bir error döndürmeli**:

   * İlk parametre: input (struct veya basic type)
   * İkinci parametre: output (pointer)
   * Return: error
3. Parametre ve dönüş değerleri **exported (public)** olmalı, aksi takdirde seri hale getirilemez.
4. Sunucu TCP veya HTTP üzerinden dinleyebilir.

   * `rpc.ServeConn(conn)` → TCP/Conn
   * `rpc.ServeHTTP()` → HTTP

---

# 5️⃣ HTTP Üzerinden RPC Örneği
*/
``go
package main

import (
	"net/http"
	"net/rpc"
)

func main() {
	calc := new(Calculator)
	rpc.Register(calc)
	rpc.HandleHTTP() // RPC endpoint'lerini HTTP ile aç

	http.ListenAndServe(":1234", nil)
}
``
/*
* Client tarafında `rpc.DialHTTP("tcp", "localhost:1234")` kullanılır.

---

# 6️⃣ Özet Tablo
*/

| Konsept                                    | Açıklama                                              |
| ------------------------------------------ | ----------------------------------------------------- |
| `rpc.Register`                             | Servisi kaydeder                                      |
| `rpc.ServeConn`                            | TCP bağlantısı üzerinden RPC sunar                    |
| `rpc.ServeHTTP`                            | HTTP üzerinden RPC sunar                              |
| `rpc.Dial`                                 | TCP ile sunucuya bağlanır                             |
| `rpc.DialHTTP`                             | HTTP ile sunucuya bağlanır                            |
| `client.Call("Type.Method", args, &reply)` | RPC metodunu çağırır                                  |
| RPC metod kuralları                        | public, iki parametre (input, \*output), return error |

/*
---

✅ **Kısaca:**

* `net/rpc` Go’da **basit ve hızlı bir RPC mekanizması** sağlar.
* TCP veya HTTP üzerinden çalışabilir.
* Parametreler ve dönüş değerleri struct veya temel tipler olmalıdır.
* Büyük ve dağıtık sistemlerde temel bir RPC protokolü olarak kullanılabilir.

---

İstersen sana bunu bir adım daha ileri götürüp **concurrent RPC client ve server ile performans testi** örneğini de gösterebilirim. Bunu ister misin?
EVET
*/
/*
Harika 👍 o zaman sana **concurrent RPC client ve server örneğini** göstereyim. Bu örnekle:

* Sunucu aynı anda birden fazla client çağrısını işleyebilecek,
* Client’lar eş zamanlı (concurrent) olarak RPC çağrıları yapacak,
* Böylece **performans ve concurrency testi** yapabiliriz.

---

## 1️⃣ RPC Sunucu Kodu (Concurrent)
*/
``go
package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
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

// Bölme
func (c *Calculator) Divide(args Args, reply *float64) error {
	if args.B == 0 {
		return errors.New("sıfıra bölme hatası")
	}
	*reply = float64(args.A) / float64(args.B)
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

	fmt.Println("Concurrent RPC sunucusu çalışıyor :1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept hatası:", err)
			continue
		}
		go rpc.ServeConn(conn) // Her client için ayrı goroutine
	}
}
``
/*
* Sunucu her client için ayrı goroutine açıyor → **concurrent**.

---

## 2️⃣ Concurrent RPC Client Kodu
*/
``go
package main

import (
	"fmt"
	"net/rpc"
	"sync"
)

type Args struct {
	A, B int
}

func main() {
	numClients := 5       // eş zamanlı client sayısı
	var wg sync.WaitGroup

	for i := 1; i <= numClients; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			client, err := rpc.Dial("tcp", "localhost:1234")
			if err != nil {
				fmt.Printf("Client %d bağlanamadı: %v\n", id, err)
				return
			}
			defer client.Close()

			args := Args{A: id * 10, B: id * 2}

			var sum int
			err = client.Call("Calculator.Add", args, &sum)
			if err != nil {
				fmt.Printf("Client %d Add hatası: %v\n", id, err)
			} else {
				fmt.Printf("Client %d Toplam: %d\n", id, sum)
			}

			var div float64
			err = client.Call("Calculator.Divide", args, &div)
			if err != nil {
				fmt.Printf("Client %d Divide hatası: %v\n", id, err)
			} else {
				fmt.Printf("Client %d Bölme: %.2f\n", id, div)
			}

		}(i)
	}

	wg.Wait()
	fmt.Println("Tüm concurrent client işlemleri tamamlandı.")
}
``
/*
---

## 3️⃣ Açıklamalar

1. **Sunucu tarafında:**

   * `rpc.ServeConn(conn)` her bağlantı için ayrı goroutine başlatır.
   * Böylece **birden fazla client eş zamanlı** hizmet alabilir.

2. **Client tarafında:**

   * `numClients` kadar goroutine başlatılır.
   * Her goroutine RPC sunucusuna bağlanır ve `Add` ile `Divide` çağrılarını yapar.
   * `sync.WaitGroup` ile tüm client’ların bitmesi beklenir.

3. **Performans Testi:**

   * Çok sayıda goroutine açıp RPC çağrısı yaparak **CPU ve concurrency testleri** yapılabilir.
   * `pprof` ile sunucu tarafında **CPU ve memory kullanımı** ölçülebilir.

---

## 4️⃣ Örnek Çıktı
*/
``
Client 1 Toplam: 12
Client 1 Bölme: 5.00
Client 2 Toplam: 24
Client 2 Bölme: 5.00
Client 3 Toplam: 36
Client 3 Bölme: 5.00
...
Tüm concurrent client işlemleri tamamlandı.
``
/*
---

✅ Bu yöntemle:

* RPC sunucusu **concurrent client** çağrılarını güvenle yönetir.
* Client tarafında eş zamanlı çağrılarla performans ölçümü yapılabilir.
* `pprof` kullanarak **CPU/memory profili** çıkarılabilir ve darboğazlar tespit edilebilir.

---

İstersen sana bir sonraki adım olarak **RPC + Timeout + Context yönetimi ile güvenli ve zaman aşımı kontrollü concurrent RPC** örneğini de gösterebilirim. Bunu ister misin?
EVET
*/
/*
Harika 👍 o zaman sana **RPC + Timeout + Context yönetimi** ile güvenli ve zaman aşımı kontrollü concurrent RPC örneğini göstereyim.

Bu yöntemle:

* RPC çağrıları **timeout** ile sınırlandırılır,
* Uzun süren veya takılan işlemler kontrol altına alınır,
* Client tarafında **context ile iptal** edilebilir.

---

## 1️⃣ RPC Sunucu Kodu (Basit)

Sunucu tarafında timeout yönetimi yok, sadece RPC metodları var:
*/
``go
package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"time"
)

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

// Bölme
func (c *Calculator) Divide(args Args, reply *float64) error {
	if args.B == 0 {
		return errors.New("sıfıra bölme hatası")
	}
	time.Sleep(1 * time.Second)
	*reply = float64(args.A) / float64(args.B)
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

	fmt.Println("RPC sunucusu çalışıyor :1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept hatası:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
``
/*
---

## 2️⃣ Client Kodu (Timeout + Context)
*/
``go
package main

import (
	"context"
	"fmt"
	"net/rpc"
	"time"
)

type Args struct {
	A, B int
}

func callWithTimeout(client *rpc.Client, method string, args Args, reply interface{}, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	errCh := make(chan error, 1)

	go func() {
		errCh <- client.Call(method, args, reply)
	}()

	select {
	case <-ctx.Done():
		return fmt.Errorf("RPC çağrısı zaman aşımı: %v", ctx.Err())
	case err := <-errCh:
		return err
	}
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Bağlantı hatası:", err)
		return
	}
	defer client.Close()

	args := Args{A: 10, B: 5}

	var sum int
	err = callWithTimeout(client, "Calculator.Add", args, &sum, 1*time.Second)
	if err != nil {
		fmt.Println("Add hatası:", err)
	} else {
		fmt.Println("Toplam:", sum)
	}

	var div float64
	err = callWithTimeout(client, "Calculator.Divide", args, &div, 3*time.Second)
	if err != nil {
		fmt.Println("Divide hatası:", err)
	} else {
		fmt.Println("Bölme:", div)
	}
}
``
/*
---

## 3️⃣ Açıklamalar

1. `context.WithTimeout` kullanarak RPC çağrısına **zaman aşımı** ekledik.
2. Goroutine içinde `client.Call` çağrısı yapılıyor ve sonucu bir channel’a gönderiliyor.
3. `select` ile:

   * Eğer context timeout olursa → "zaman aşımı" hatası
   * Eğer RPC başarılı olursa → normal dönüş
4. Böylece uzun süren işlemler **client tarafında iptal edilebilir**.

---

## 4️⃣ Örnek Çıktı
*/
``
Add hatası: RPC çağrısı zaman aşımı: context deadline exceeded
Bölme: 2.00
``
/*
* `Add` metodunun süresi 2 saniye → timeout 1 saniye → hata veriyor
* `Divide` metodunun süresi 1 saniye → timeout 3 saniye → başarılı

---

✅ **Avantajlar:**

* Büyük ve uzun süren RPC çağrılarında **client tarafında kontrol** sağlar
* Concurrency ile birlikte **her çağrı için ayrı timeout** belirlenebilir
* `context` iptal mekanizması ile sistem kaynakları korunur

---

İstersen bir sonraki adım olarak sana **concurrent RPC + timeout + worker pool + pprof ile performans testi** örneğini birleştirilmiş şekilde gösterebilirim. Bunu ister misin?
*/