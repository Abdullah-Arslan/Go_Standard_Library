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

/*
---

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

/*
⚠️ **Dikkat:** Büyük veri veya config taşımak için **kötü pratik**.

---

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