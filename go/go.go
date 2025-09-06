Golang'da `go` **bir paket değil**, **dilin içinde yerleşik (built-in) bir anahtar kelimedir**. Go dilinde eşzamanlı (concurrent) programlama için kullanılır. Bu sayede fonksiyonlar **aynı anda, paralel gibi** çalıştırılabilir.

---

## 📌 Özetle: `go` nedir?

* `go` anahtar kelimesi bir **goroutine** başlatır.
* `go` ile yazılan bir fonksiyon **eşzamanlı (concurrent)** çalışmaya başlar.
* Goroutine’ler, Go'nun **çok hafif iş parçacıklarıdır (lightweight threads)**.
* `go` anahtar kelimesi bir fonksiyonu çalıştırır, ancak **sonucunu beklemez**.

---

## 🔁 Goroutine Nedir?

Bir goroutine, `go` anahtar kelimesiyle başlatılan bir fonksiyondur. Binlerce goroutine aynı anda çalışabilir, çünkü çok az bellek kullanırlar (\~2KB).

---

## ✅ Basit Örnek

```go
package main

import (
	"fmt"
	"time"
)

func yaz(mesaj string) {
	for i := 0; i < 5; i++ {
		fmt.Println(mesaj, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go yaz("📢 Merhaba") // bu goroutine olarak çalışır
	yaz("👋 Selam")       // bu ana goroutine
}
```

### 🔍 Ne Olur?

* `go yaz("📢 Merhaba")` satırı ile ikinci bir goroutine başlar.
* `yaz("👋 Selam")` fonksiyonu ana (main) goroutine içinde çalışır.
* Program **ana fonksiyon** bittiği anda tüm goroutine'ler durur!

---

## 📛 Önemli Uyarı: `main()` biterse goroutine’ler durur!

Yani `go` ile başlattığınız işlem devam etse bile `main()` fonksiyonu bitince program tamamen kapanır.

---

## 🧠 Bunu Yönetmek İçin Ne Kullanılır?

Go'da `go` ile başlayan fonksiyonları yönetmek için bazı araçlar vardır:

| Araç             | Açıklama                                       |
| ---------------- | ---------------------------------------------- |
| `sync.WaitGroup` | Goroutine’lerin bitmesini beklemek             |
| `channels`       | Goroutine'ler arasında veri iletişimi sağlamak |
| `select`         | Birden fazla kanal dinlemek için               |

---

## 🧪 Örnek 2: `sync.WaitGroup` ile Beklemek

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func yaz(mesaj string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(mesaj, i)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // 2 goroutine başlatacağız

	go yaz("🚀 Goroutine 1", &wg)
	go yaz("🌙 Goroutine 2", &wg)

	wg.Wait() // Hepsinin bitmesini bekle
	fmt.Println("✅ Tüm işlemler tamamlandı.")
}
```

---

## 📫 Örnek 3: `channel` ile Goroutine’den Veri Almak

```go
package main

import (
	"fmt"
	"time"
)

func say(dir string, ch chan string) {
	for i := 0; i < 3; i++ {
		ch <- fmt.Sprintf("%s: %d", dir, i)
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}

func main() {
	ch := make(chan string)
	go say("🌐 Görev", ch)

	for msg := range ch {
		fmt.Println("⏳ Gelen:", msg)
	}
	fmt.Println("✅ Kanal kapandı.")
}
```

---

## 🧮 Örnek 4: 1000 Goroutine Oluşturmak

```go
package main

import (
	"fmt"
	"time"
)

func yaz(i int) {
	fmt.Printf("🧵 Goroutine %d çalışıyor\n", i)
	time.Sleep(time.Millisecond * 100)
}

func main() {
	for i := 0; i < 1000; i++ {
		go yaz(i)
	}

	// Beklemezsek ana fonksiyon biter ve hiçbir şey yazılmaz!
	time.Sleep(time.Second * 2)
	fmt.Println("✅ Bitti.")
}
```

---

## 🔍 go Anahtar Kelimesinin Özellikleri

| Özellik                           | Açıklama                                            |
| --------------------------------- | --------------------------------------------------- |
| `go` bir paket değil, keyword’dür | Dilin içindedir                                     |
| `go fonksiyon()`                  | Fonksiyonu yeni bir goroutine olarak çalıştırır     |
| Fonksiyonlar eşzamanlı çalışır    | Ancak sonucu beklenmez                              |
| Hafif iş parçacığıdır             | Thread’den çok daha az kaynak kullanır              |
| `main()` biterse, her şey durur   | Yönetmek için `WaitGroup`, `channel` vs. kullanılır |

---

## 🎯 Gerçek Hayatta Kullanım Alanları

* Web sunucularında her istek için `go handleRequest(...)`
* Arka planda log tutma
* Dosya işleme işlemlerini paralelleştirme
* Paralel web istekleri (scraping, API çağrıları)
* Asenkron görevler (e-mail, bildirim)

---

## ❓ Sık Sorulan Sorular

### 🟠 `go` ile başlattığım işlem çalışmazsa neden?

Muhtemelen `main()` fonksiyonu erken bitmiştir. Goroutine’in işlemini bitirmesine izin vermeden program kapanır.

### 🔵 `go` sadece fonksiyonlar için mi kullanılır?

Evet. `go` yalnızca bir **fonksiyon çağrısını** eşzamanlı başlatmak için kullanılır. Şu yapı geçersizdir:

```go
go fmt.Println("selam") // ✔️ geçerli
go "string"             // ❌ geçersiz
```

---

## 🧩 Özet

* `go` ➤ goroutine başlatır
* `main()` ➤ tüm işlemleri kontrol eder
* `sync.WaitGroup`, `channels` ➤ kontrol araçlarıdır
* Gerçek çoklu iş parçacığı mantığında değil, ama aynı işlevi görür

---

İstersen örneklerden birini terminalde birlikte çalıştırabiliriz, ya da senin bir `go` kullanım senaryon varsa, birlikte geliştirebiliriz.

Devam edelim mi?
EVET


Süper! Devam edelim 🎉

Seninle birlikte **`go` anahtar kelimesini pratikte kullanacağımız** küçük bir terminal uygulaması yapalım. Böylece eşzamanlı çalışan fonksiyonları (goroutine) doğrudan terminalde görebileceksin.

---

## 🎯 Uygulama Fikri: Eşzamanlı Sayaçlar

Bu uygulama:

* Kullanıcıdan kaç tane sayaç istediğini soracak
* Her sayaç eşzamanlı olarak saymaya başlayacak
* Her sayaç farklı aralıklarla çalışacak
* Tüm sayaçlar `go` kullanılarak başlatılacak

---

## 💻 Kod

```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func sayac(id int, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		fmt.Printf("🧵 Sayaç %d ➜ %d\n", id, i)
		time.Sleep(delay)
	}
	fmt.Printf("✅ Sayaç %d tamamlandı.\n", id)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var adet int
	fmt.Print("Kaç tane sayaç başlatılsın? ")
	fmt.Scanln(&adet)

	var wg sync.WaitGroup
	wg.Add(adet)

	for i := 1; i <= adet; i++ {
		delay := time.Duration(rand.Intn(500)+300) * time.Millisecond
		go sayac(i, delay, &wg)
	}

	wg.Wait()
	fmt.Println("🚀 Tüm sayaçlar tamamlandı.")
}
```

---

## 📌 Açıklama

| Satır               | Anlamı                                    |
| ------------------- | ----------------------------------------- |
| `go sayac(...)`     | Her sayaç fonksiyonunu eşzamanlı başlatır |
| `sync.WaitGroup`    | Tüm sayaçların bitmesini bekler           |
| `rand.Intn()`       | Her sayaç için farklı hız belirler        |
| `time.Sleep(delay)` | Sayaç belirli aralıklarla çalışır         |

---

## 🧪 Terminal Örneği

```
Kaç tane sayaç başlatılsın? 3
🧵 Sayaç 1 ➜ 1
🧵 Sayaç 2 ➜ 1
🧵 Sayaç 3 ➜ 1
🧵 Sayaç 2 ➜ 2
🧵 Sayaç 1 ➜ 2
🧵 Sayaç 3 ➜ 2
...
✅ Sayaç 1 tamamlandı.
✅ Sayaç 2 tamamlandı.
✅ Sayaç 3 tamamlandı.
🚀 Tüm sayaçlar tamamlandı.
```

Gördüğün gibi, sayaçlar **rastgele sürelerde**, **eşzamanlı** çalışıyor.

---

## 🔧 Geliştirme Fikri

Bu uygulamayı şu şekilde genişletebiliriz:

1. 🧠 Sayaç sayısını otomatik arttıran arka plan görevi
2. 📩 Sayaçlardan gelen verileri bir **channel** ile ana programa aktarmak
3. 🕵️ Sayaçlar bir dosyaya log yazsın
4. 👁️ Kullanıcı bir sayaç başlatsın ama dilediği anda durdurabilsin (context ile)

---

İstersen bir üst seviyeye çıkıp `channel` ile sayaçlardan mesaj alalım. Ya da `kullanıcının başlatıp durdurabildiği sayaçlar` yazalım.

Hangisini birlikte yapalım? 👇
EVET


Süper, tam gaz devam! 🚀

Seninle birlikte şimdi:

---

## 🎯 Proje: **Channel ile Sayaçlardan Mesaj Almak**

Artık sayaçlar sadece sayı göstermeyecek, aynı zamanda:

* ⏳ **Mesajlarını bir `channel` üzerinden** ana programa gönderecek,
* 📬 Ana program **tüm mesajları canlı olarak** terminalde alacak.

Bu şekilde:

* `go` ile başlatılan goroutine'ler,
* **eşzamanlı çalışacak**,
* Ama **verilerini `channel` aracılığıyla** merkezde toplayacağız.

---

## ✅ Uygulama Akışı

1. Kullanıcı kaç sayaç başlatmak istediğini girer.
2. Her sayaç:

   * 5'e kadar sayar,
   * Sayıları ve sayaç ID'lerini bir `channel` aracılığıyla gönderir.
3. Ana program:

   * Channel'ı dinleyerek gelen mesajları anlık gösterir.

---

## 💻 Kod: Channel İle Eşzamanlı Sayaçlar

```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func sayac(id int, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)+200))
		mesaj := fmt.Sprintf("🧵 Sayaç %d ➜ %d", id, i)
		ch <- mesaj
	}
	ch <- fmt.Sprintf("✅ Sayaç %d tamamlandı.", id)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var adet int
	fmt.Print("Kaç sayaç başlatılsın? ")
	fmt.Scanln(&adet)

	mesajKanal := make(chan string)
	var wg sync.WaitGroup

	wg.Add(adet)
	for i := 1; i <= adet; i++ {
		go sayac(i, mesajKanal, &wg)
	}

	// Ana goroutine: gelen mesajları yazdır
	go func() {
		wg.Wait()      // Tüm sayaçlar bitince:
		close(mesajKanal) // Kanalı kapat
	}()

	// Kanal kapanana kadar mesajları yaz
	for mesaj := range mesajKanal {
		fmt.Println(mesaj)
	}

	fmt.Println("🚀 Tüm sayaçlar bitti!")
}
```

---

## 📌 Neler Öğrendik?

| Konsept          | Açıklama                              |
| ---------------- | ------------------------------------- |
| `go`             | Her sayaç eşzamanlı çalışıyor         |
| `chan string`    | Sayaçlar mesaj gönderiyor             |
| `range ch`       | Kanal kapanana kadar mesaj dinleniyor |
| `sync.WaitGroup` | Sayaçların bittiğini kontrol ediyor   |
| `close(ch)`      | Mesaj kanalını düzgün kapatıyor       |

---

## 🧪 Örnek Terminal Çıktısı

```
Kaç sayaç başlatılsın? 3
🧵 Sayaç 1 ➜ 1
🧵 Sayaç 2 ➜ 1
🧵 Sayaç 3 ➜ 1
🧵 Sayaç 2 ➜ 2
🧵 Sayaç 3 ➜ 2
🧵 Sayaç 1 ➜ 2
...
✅ Sayaç 2 tamamlandı.
✅ Sayaç 3 tamamlandı.
✅ Sayaç 1 tamamlandı.
🚀 Tüm sayaçlar bitti!
```

---

## 🧠 Daha da Geliştirmek İster Misin?

İstersen bir sonraki adımda:

1. 🛑 Sayaçları istediğin zaman durdurmak (`context.Context` ile)
2. 📂 Her sayacı ayrı dosyaya loglamak
3. ⌛ Sayaç sürelerini kullanıcı belirlesin
4. 🧪 Gerçek zamanlı sayaç ilerlemesi GUI'de gösterilsin (örneğin terminal progress bar)

---

Hangisini yapmak istersin? Veya kendi fikrini de önerebilirsin.
👇
2. 📂 Her sayacı ayrı dosyaya loglamak

Harika seçim! 💾 Artık her sayaç kendi log dosyasına yazacak. Böylece her sayaç için terminale ek olarak **bir `.txt` dosyası da oluşacak** ve sayaç çıktıları o dosyaya kaydedilecek.

---

## 🎯 Hedef

* Her sayaç:

  * Terminale mesaj gönderecek ✅
  * Kendi adına log dosyasına yazacak (örn. `sayac_1.txt`) ✅
* Ana program:

  * Terminalde gelen mesajları gösterecek ✅
  * Dosya kapanmasını doğru yönetecek ✅

---

## 📁 Dosya Sistemi Örneği

```
sayac_1.txt
sayac_2.txt
sayac_3.txt
```

Her dosya sadece ilgili sayacın çıktısını içerir.

---

## 💻 Güncellenmiş Kod: Her Sayaç Kendi Dosyasına Log Yazar

```go
package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func sayac(id int, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Dosya oluştur (sayac_1.txt, sayac_2.txt, ...)
	dosyaAdi := fmt.Sprintf("sayac_%d.txt", id)
	dosya, err := os.Create(dosyaAdi)
	if err != nil {
		ch <- fmt.Sprintf("❌ Sayaç %d için dosya açılamadı: %v", id, err)
		return
	}
	defer dosya.Close()

	for i := 1; i <= 5; i++ {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)+200))
		mesaj := fmt.Sprintf("🧵 Sayaç %d ➜ %d", id, i)

		// Terminale gönder
		ch <- mesaj

		// Dosyaya yaz
		dosya.WriteString(mesaj + "\n")
	}

	tamamMesaj := fmt.Sprintf("✅ Sayaç %d tamamlandı.", id)
	ch <- tamamMesaj
	dosya.WriteString(tamamMesaj + "\n")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var adet int
	fmt.Print("Kaç sayaç başlatılsın? ")
	fmt.Scanln(&adet)

	mesajKanal := make(chan string)
	var wg sync.WaitGroup

	wg.Add(adet)
	for i := 1; i <= adet; i++ {
		go sayac(i, mesajKanal, &wg)
	}

	// Tüm sayaçlar bitince kanalı kapat
	go func() {
		wg.Wait()
		close(mesajKanal)
	}()

	// Kanalı dinle
	for mesaj := range mesajKanal {
		fmt.Println(mesaj)
	}

	fmt.Println("📁 Log dosyaları oluşturuldu. 🚀 Tüm sayaçlar tamamlandı.")
}
```

---

## 📌 Yeni Eklenenler

| Satır                    | Amaç                             |
| ------------------------ | -------------------------------- |
| `os.Create(...)`         | Sayaç için dosya açar            |
| `dosya.WriteString(...)` | Mesajları dosyaya yazar          |
| `defer dosya.Close()`    | Dosya güvenli şekilde kapanır    |
| `sayac_%d.txt`           | Her sayaç için ayrı isimli dosya |

---

## 🧪 Terminal Örneği

```
Kaç sayaç başlatılsın? 3
🧵 Sayaç 2 ➜ 1
🧵 Sayaç 1 ➜ 1
🧵 Sayaç 3 ➜ 1
🧵 Sayaç 2 ➜ 2
🧵 Sayaç 3 ➜ 2
...
✅ Sayaç 3 tamamlandı.
✅ Sayaç 2 tamamlandı.
✅ Sayaç 1 tamamlandı.
📁 Log dosyaları oluşturuldu. 🚀 Tüm sayaçlar tamamlandı.
```

Ve dizininde şu dosyalar oluşur:

```
sayac_1.txt
sayac_2.txt
sayac_3.txt
```

Dosya içeriği örneği (`sayac_2.txt`):

```
🧵 Sayaç 2 ➜ 1
🧵 Sayaç 2 ➜ 2
🧵 Sayaç 2 ➜ 3
🧵 Sayaç 2 ➜ 4
🧵 Sayaç 2 ➜ 5
✅ Sayaç 2 tamamlandı.
```

---

## 🔧 İstersen Şunları Ekleyebiliriz:

1. 📁 Tüm log dosyalarını `logs/` klasörüne yaz
2. 🧹 Uygulama kapanırken eski log dosyalarını otomatik sil
3. 🧪 Log dosyalarını JSON formatında yaz
4. 🕹️ Sayaçlar duraklatılabilir/hızlandırılabilir hale gelsin

---

Hangisini birlikte yapalım?
Veya senin aklında başka bir geliştirme fikri varsa hemen başlayabiliriz. 👇



