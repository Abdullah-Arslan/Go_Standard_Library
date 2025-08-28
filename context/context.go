/*
Go’da **`context` paketi**, birden fazla goroutine’in yaşam süresini yönetmek, iptal etmek, zaman aşımı koymak ve belirli request-scope verileri taşımak için kullanılan standart bir pakettir. Özellikle **HTTP sunucularında, uzun süren işlemlerde, API çağrılarında** ve **işlem zincirlerinde** çok önemlidir.

---

# 📌 Context Paketinin Genel Amacı

* Bir işlemi veya alt işlemleri **iptal** edebilmek (cancellation).
* İşlemlere **zaman sınırı (timeout/deadline)** koymak.
* Fonksiyonlar arası küçük verileri **taşımak** (request ID, kullanıcı bilgisi gibi).
* Bir request/işlem tamamlandığında, ilgili alt işlemleri de **otomatik sonlandırmak**.

---

# 📦 Paket Fonksiyonları ve Tipleri

`context` paketinin en çok kullanılan fonksiyon ve tipleri şunlardır:

1. **`context.Background()`**

   * En üst düzey (root) context’tir. Hiçbir deadline veya iptal sinyali yoktur.
   * Genelde uygulama giriş noktasında (main, server start vb.) kullanılır.

2. **`context.TODO()`**

   * Nerede hangi context’i kullanacağınızı bilmediğiniz durumlarda geçici olarak kullanılır.
   * Daha sonra `Background` veya başka context ile değiştirilir.

3. **`context.WithCancel(parent)`**

   * Bir context üretir. `cancel()` çağrıldığında, o context ve alt contextler iptal edilir.

4. **`context.WithTimeout(parent, süre)`**

   * Belirtilen süre sonra otomatik olarak iptal olur.

5. **`context.WithDeadline(parent, zaman)`**

   * Belirtilen kesin tarihte otomatik iptal olur.

6. **`context.WithValue(parent, key, value)`**

   * Context içine key-value şeklinde veri taşır. (Genelde küçük metadata için)

7. **`ctx.Done()`**

   * Bir kanal döner. Context iptal edilirse bu kanal kapanır.

8. **`ctx.Err()`**

   * Context neden iptal oldu? (`context.Canceled` veya `context.DeadlineExceeded`)

9. **`ctx.Value(key)`**

   * Context’ten değer okur.

---

# 🔎 Örnekler ile Açıklama

### 1. Basit `Background` ve `TODO` kullanımı
*/

package main

import (
	"context"
	"fmt"
)

func main() {
	ctx1 := context.Background() // Root context
	ctx2 := context.TODO()       // Daha sonra belirlenecek

	fmt.Println("Background:", ctx1)
	fmt.Println("TODO:", ctx2)
}


// ---

// ### 2. `WithCancel` ile işlemi iptal etmek


package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		// uzun süren işlem
		for {
			select {
			case <-ctx.Done():
				fmt.Println("İşlem iptal edildi:", ctx.Err())
				return
			default:
				fmt.Println("Çalışıyor...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	cancel() // işlemi iptal et
	time.Sleep(1 * time.Second)
}


// ---

// ### 3. `WithTimeout` ile otomatik iptal


package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("İş tamamlandı.")
	case <-ctx.Done():
		fmt.Println("Timeout:", ctx.Err())
	}
}

// ---

// ### 4. `WithDeadline` kullanımı


package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("İş tamamlandı.")
	case <-ctx.Done():
		fmt.Println("Deadline geçti:", ctx.Err())
	}
}
/*

---

### 5. `WithValue` ile veri taşıma
*/
package main

import (
	"context"
	"fmt"
)

func main() {
	// request ID taşımak örneği
	ctx := context.WithValue(context.Background(), "requestID", "REQ-12345")

	processRequest(ctx)
}

func processRequest(ctx context.Context) {
	reqID := ctx.Value("requestID")
	fmt.Println("İşlem Request ID ile çalışıyor:", reqID)
}
/*

⚠️ Not: `WithValue`, büyük veri taşımak için **kötü bir pratik**. Genellikle sadece küçük metadata için kullanılır.

---

### 6. HTTP Sunucusunda Context Kullanımı

Go’nun HTTP paketinde her request için bir context vardır.
Eğer client bağlantıyı kapatırsa, context otomatik iptal olur.
*/

package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Println("Request başladı")
	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "İşlem bitti")
	case <-ctx.Done():
		fmt.Println("İstemci bağlantıyı kapattı:", ctx.Err())
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
/*

---

# 📌 Özet

* `context` Go’da **goroutine yönetimi** için çok kritik bir pakettir.
* Fonksiyon zincirlerinde request bazlı iptal, timeout ve veri taşıma sağlar.
* Özellikle **HTTP server**, **gRPC**, **veritabanı işlemleri**, **uzun süren işler** için standarttır.

---

İstersen sana **context paketindeki her fonksiyon ve metodun kullanımını tek tek örnekleyen tam bir “cheatsheet”** çıkarayım. İster misin?
*/
/*
 O zaman  **Go `context` paketi için tam kapsamlı bir “cheatsheet”** hazırlayayım.
Bu listede paketin **tüm fonksiyonlarını, methodlarını, hata tiplerini** göreceksin ve her biri için **örnek kod** olacak.

---

# 📘 Go `context` Cheatsheet

---

## 1. `context.Background()`

🔹 Root (temel) context oluşturur.
Genelde uygulama başlangıcında kullanılır.
*/

ctx := context.Background()
fmt.Println(ctx) // <context.Background>

/*
---

## 2. `context.TODO()`

🔹 Henüz hangi context kullanılacağını bilmediğinde placeholder olarak kullanılır.
*/

ctx := context.TODO()
fmt.Println(ctx) // <context.TODO>

/*
---

## 3. `context.WithCancel(parent)`

🔹 Parent context’ten türetilir.
🔹 `cancel()` çağrılınca, bu context ve alt contextleri iptal olur.
*/

ctx, cancel := context.WithCancel(context.Background())

go func() {
    <-ctx.Done()
    fmt.Println("İptal edildi:", ctx.Err())
}()

time.Sleep(time.Second)
cancel()

/*
---

## 4. `context.WithTimeout(parent, duration)`

🔹 Belirlenen süre sonunda otomatik iptal olur.
🔹 `cancel()` manuel çağrılabilir (her zaman `defer cancel()` yazılır).
*/

ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

select {
case <-time.After(3 * time.Second):
    fmt.Println("Tamamlandı")
case <-ctx.Done():
    fmt.Println("Timeout:", ctx.Err()) // DeadlineExceeded
}

//---

/*

## 5. `context.WithDeadline(parent, time)`

🔹 Belirlenen **kesin tarihte** iptal olur.

*/
deadline := time.Now().Add(2 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()

<-ctx.Done()
fmt.Println("Deadline:", ctx.Err())

/*
---

## 6. `context.WithValue(parent, key, value)`

🔹 Context içine metadata ekler.
🔹 Küçük ve sabit veriler için kullanılmalı (örn. request ID).
*/

ctx := context.WithValue(context.Background(), "userID", 42)

userID := ctx.Value("userID")
fmt.Println("UserID:", userID)


//⚠️ **Dikkat:** Büyük veri veya config taşımak için **kötü pratik**.

//---
/*
## 7. `ctx.Done()`

🔹 `chan struct{}` döner.
🔹 Context iptal olduğunda bu kanal kapanır.
*/

ctx, cancel := context.WithCancel(context.Background())

go func() {
    <-ctx.Done()
    fmt.Println("İşlem durdu:", ctx.Err())
}()

cancel()

/*
---

## 8. `ctx.Err()`

🔹 Context iptal edildiğinde sebebi döner:

* `context.Canceled` → manuel iptal edildi
* `context.DeadlineExceeded` → timeout veya deadline doldu
*/

ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()

<-ctx.Done()
fmt.Println("Hata:", ctx.Err()) // DeadlineExceeded

/*
---

## 9. `ctx.Value(key)`

🔹 `WithValue` ile saklanan veriyi çeker.
*/

ctx := context.WithValue(context.Background(), "traceID", "XYZ-123")
fmt.Println("Trace:", ctx.Value("traceID"))

/*
---

## 10. Özel Hata Tipleri

* `context.Canceled` → İşlem iptal edildi
* `context.DeadlineExceeded` → Deadline veya timeout geçti
*/

ctx, cancel := context.WithCancel(context.Background())
cancel()

if errors.Is(ctx.Err(), context.Canceled) {
    fmt.Println("İş iptal edildi")
}

/*
---

# ⚡ Komple Örnek: İşlem Zinciri Yönetimi

Aşağıdaki örnek, **cancel + timeout + value** kombinasyonunu gösteriyor:
*/

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Ana context
	ctx := context.Background()

	// Request ID ekle
	ctx = context.WithValue(ctx, "reqID", "REQ-999")

	// Timeout ekle
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	doWork(ctx)
}

func doWork(ctx context.Context) {
	fmt.Println("Başladı, Request ID:", ctx.Value("reqID"))

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("İş bitti")
	case <-ctx.Done():
		fmt.Println("İptal:", ctx.Err())
	}
}

/*

📌 Çıktı:


Başladı, Request ID: REQ-999
İptal: context deadline exceeded


---

# ✅ Özet (Cheatsheet Tablosu)

| Fonksiyon / Method          | Açıklama                   |
| --------------------------- | -------------------------- |
| `Background()`              | Root context               |
| `TODO()`                    | Geçici placeholder         |
| `WithCancel(parent)`        | İptal edilebilir context   |
| `WithTimeout(parent, süre)` | Süre dolunca iptal         |
| `WithDeadline(parent, t)`   | Belirli zamanda iptal      |
| `WithValue(parent, k, v)`   | Metadata ekleme            |
| `ctx.Done()`                | İptal kanalı               |
| `ctx.Err()`                 | İptal sebebi               |
| `ctx.Value(k)`              | Veri okuma                 |
| `context.Canceled`          | Manuel iptal hatası        |
| `context.DeadlineExceeded`  | Süre/Deadline doldu hatası |

---

Böylece `context` paketinin **tamamını örneklerle** özetlemiş olduk 🚀
*/