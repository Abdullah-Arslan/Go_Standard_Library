/*
Go’da **`net/http/pprof`** paketi, **profiling (çalışan programın performans analizi)** yapmak için kullanılan bir pakettir.

Bu paket sayesinde çalışan programın:

* CPU kullanımını,
* Bellek kullanımını,
* Goroutine sayısını,
* Block (kilitlenme) durumlarını,
* Heap bilgilerini

**runtime** sırasında web üzerinden görebilirsin.

Kısacası: `http/pprof` paketi Go uygulamanın içine gömülen bir **profiling aracı**dır.

---

## 1. Paketi Kullanma

`import _ "net/http/pprof"` yazdığında, bu paket otomatik olarak bazı HTTP handler’ları **`/debug/pprof/`** altında açar.

### Basit Örnek
*/
``go
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Merhaba Dünya")
	})

	// 6060 portunda HTTP server açıyoruz
	fmt.Println("Server başladı: http://localhost:6060/")
	http.ListenAndServe(":6060", nil)
}
``
/*
👉 Bu program çalıştığında:

* `http://localhost:6060/debug/pprof/` → Profil sayfası
* `http://localhost:6060/debug/pprof/goroutine` → Goroutine dump
* `http://localhost:6060/debug/pprof/heap` → Heap dump
* `http://localhost:6060/debug/pprof/profile?seconds=30` → 30 saniyelik CPU profili

---

## 2. `pprof` Endpoint’leri

`http/pprof` şu handler’ları ekler:
*/

| Endpoint                    | Açıklama                                   |
| --------------------------- | ------------------------------------------ |
| `/debug/pprof/`             | Tüm profil türlerini listeler              |
| `/debug/pprof/cmdline`      | Programın çalıştırma argümanlarını döner   |
| `/debug/pprof/profile`      | CPU profili (varsayılan 30 saniye)         |
| `/debug/pprof/symbol`       | Programdaki semboller (fonksiyon isimleri) |
| `/debug/pprof/trace`        | Program yürütme izleri (execution trace)   |
| `/debug/pprof/goroutine`    | Aktif goroutine dump                       |
| `/debug/pprof/heap`         | Heap (bellek) bilgisi                      |
| `/debug/pprof/threadcreate` | Thread oluşturma bilgisi                   |
| `/debug/pprof/block`        | Blocking profili                           |

/*
---

## 3. Profil Alma ve Analiz Etme

### CPU Profili Alma

Terminalde:
*/
``bash
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=15
``
/*
Bu, 15 saniyelik CPU kullanım profilini toplar.

Ardından etkileşimli `pprof` arayüzü açılır:

* `top` → en çok CPU kullanan fonksiyonları gösterir
* `list FuncName` → belirli fonksiyonun satır satır CPU kullanımı
* `web` → grafiksel gösterim (Graphviz gerektirir)

---

### Heap Profili Alma
*/
``bash
go tool pprof http://localhost:6060/debug/pprof/heap
``
/*
Bu, belleğin nerede tahsis edildiğini gösterir.

---

### Goroutine Analizi

Tarayıcıdan:
*/
``
http://localhost:6060/debug/pprof/goroutine?debug=2
``
/*
Detaylı goroutine dump verir. Deadlock (kilitlenme) veya sızıntı (goroutine leak) bulmak için çok faydalıdır.

---

## 4. Kodla Manuel Profil Çıkarma

`pprof` sadece HTTP server değil, manuel olarak dosyaya profil çıkarmak için de kullanılabilir. Bunun için `runtime/pprof` paketi kullanılır (ama HTTP versiyonu daha pratiktir).

Örneğin CPU profilini dosyaya yazmak:
*/
``go
package main

import (
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Profil başlat
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// İşlem simülasyonu
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
	}
}
``

//Sonra analiz için:

``bash
go tool pprof cpu.prof
``
/*
---

## 5. Özet Tablo
*/

| Fonksiyon / Endpoint        | Açıklama                       |
| --------------------------- | ------------------------------ |
| `/debug/pprof/`             | Genel pprof ana sayfası        |
| `/debug/pprof/cmdline`      | Program argümanları            |
| `/debug/pprof/profile`      | CPU profili (varsayılan 30sn)  |
| `/debug/pprof/symbol`       | Semboller (fonksiyon isimleri) |
| `/debug/pprof/trace`        | Yürütme izleri                 |
| `/debug/pprof/goroutine`    | Goroutine bilgisi              |
| `/debug/pprof/heap`         | Heap bilgisi                   |
| `/debug/pprof/block`        | Block profili                  |
| `/debug/pprof/threadcreate` | Thread bilgisi                 |

/*
---

✅ Sonuç:
`net/http/pprof` paketi, Go programlarını **canlı izlemek, performans sorunlarını tespit etmek, bellek ve CPU analizleri yapmak için** çok güçlü bir araçtır.

---

İstersen sana ben bunun için **örnek bir senaryo** (örneğin bellek sızıntısı yapan kod yazıp `pprof` ile tespit etme) da gösterebilirim. İster misin?
EVET
*/
/*
Süper 👍 o zaman şimdi sana **bellek sızıntısı (memory leak) senaryosu** oluşturup, bunu `http/pprof` ile nasıl yakalayabileceğini göstereyim.

---

## 🔴 1. Bellek Sızıntısı Olan Kod

Aşağıdaki kodda, sürekli slice içine veri ekleniyor ama asla temizlenmiyor. Bu da **heap bellek kullanımının artmasına** sebep oluyor:
*
``go
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var leaks [][]byte

func leakMemory() {
	for {
		// 1 MB veri ekle
		data := make([]byte, 1024*1024)
		leaks = append(leaks, data)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// HTTP handler (pprof için)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Bellek sızıntısı örneği çalışıyor...")
	})

	// Leak fonksiyonunu ayrı goroutine’de başlat
	go leakMemory()

	// Pprof HTTP server
	fmt.Println("pprof aktif: http://localhost:6060/debug/pprof/")
	http.ListenAndServe(":6060", nil)
}
``
/*
👉 Bu program çalıştığında:

* Bellek sürekli artacak.
* `/debug/pprof/heap` ile görebileceksin.

---

## 🔍 2. Bellek Profilini Alma

Terminalde:
*/
``bash
go tool pprof http://localhost:6060/debug/pprof/heap
``
/*
Ardından pprof konsolu açılır:

* `top` → hangi fonksiyonların bellek kullandığını gösterir
* `list leakMemory` → `leakMemory` fonksiyonunda nerede bellek kullanıldığını satır satır gösterir
* `web` → grafiksel olarak bellek dağılımını gösterir (Graphviz gerekir)

Örnek çıktı:
*/
``
Showing nodes accounting for 512.00MB, 100% of 512.00MB total
      flat  flat%   sum%        cum   cum%
  512.00MB   100%   100%   512.00MB   100%  main.leakMemory
``
/*
👉 Burada `main.leakMemory` fonksiyonunun sürekli bellek ayırdığı net şekilde görülüyor.

---

## 🟢 3. Sorunu Düzeltmek

Bellek sızıntısını engellemek için slice sürekli büyümemeli. Örneğin sadece son 10 öğeyi tutalım:
*/
``go
func leakMemoryFixed() {
	for {
		data := make([]byte, 1024*1024)
		leaks = append(leaks, data)
		if len(leaks) > 10 { // sadece 10 öğeyi sakla
			leaks = leaks[1:]
		}
		time.Sleep(500 * time.Millisecond)
	}
}
``
/*
Artık `heap` profilini aldığında bellek tüketimi sabitlenmiş olacak.

---

## 📌 Özet

* `http/pprof` ile çalışan Go programını canlı izleyebiliriz.
* Bellek sızıntısını **heap profili** ile görebiliriz.
* `go tool pprof` → `top`, `list`, `web` komutlarıyla hangi fonksiyonun sızıntıya sebep olduğunu tespit edebiliriz.

---

İstersen ben sana **CPU darboğazı (infinite loop veya yoğun hesaplama)** örneği yazıp `pprof` ile analiz etmeyi de gösterebilirim. Onu da ister misin?
EVET
*/
/*
Harika 👍 şimdi sana **CPU darboğazı (bottleneck)** senaryosu göstereyim ve bunu `http/pprof` ile analiz edelim.

---

## 🔴 1. CPU’yu Yoran Kod

Aşağıdaki örnekte program, sürekli **asal sayıları hesaplıyor** ve CPU’yu zorluyor:
*/
``go
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

// Asal sayı kontrolü (CPU yoğun işlem)
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func heavyComputation() {
	num := 2
	for {
		_ = isPrime(num) // CPU sürekli asal sayı hesaplıyor
		num++
	}
}

func main() {
	// Pprof endpointlerini aç
	fmt.Println("pprof aktif: http://localhost:6060/debug/pprof/")
	go heavyComputation() // CPU’yu zorlayan fonksiyonu ayrı goroutine’de çalıştır

	http.ListenAndServe(":6060", nil)
}
``
/*
👉 Bu program çalıştığında CPU kullanımın **%100’e yakın çıkar**.

---

## 🔍 2. CPU Profilini Alma

Terminalden:
*/
``bash
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=20
``
/*
* 20 saniyelik CPU profili toplar.
* Sonra etkileşimli `pprof` konsolu açılır.

---

## 🛠 3. CPU Analizi Komutları

`pprof` içinde:

* `top`
*/
  ``
  Showing nodes accounting for 20s, 100% of 20s total
      flat  flat%   sum%        cum   cum%
      20s   100%   100%       20s   100%  main.isPrime
  ``

  /*
  👉 Burada tüm CPU’nun `main.isPrime` fonksiyonunda harcandığını görürüz.

* `list isPrime`
  Satır satır hangi döngüde CPU harcandığını gösterir.

* `web`
  Fonksiyon çağrılarını grafiksel görselleştirme (Graphviz gerekir).

---

## 🟢 4. Sorunu Düzeltmek

Bu tür hesaplamalarda:

* Daha iyi algoritma (ör. asal sayılar için Eratosthenes Eleği)
* Paralel işleme (goroutine + worker pool)
* CPU profiline bakarak hangi fonksiyonun darboğaz olduğunu optimize etme

kullanılır.

Örneğin asal kontrolünü optimize etmek:
*/
``go
func isPrimeFast(n int) bool {
	if n < 2 {
		return false
	}
	if n%2 == 0 && n != 2 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
``
/*
👉 Bu versiyon çok daha az CPU harcar.

---

## 📌 Özet

* `http/pprof` → `profile` endpoint’i ile **CPU darboğazlarını** yakalayabiliriz.
* `top`, `list`, `web` komutları hangi fonksiyonun CPU yaktığını gösterir.
* Sonuçlara göre algoritma veya mimariyi optimize ederiz.

---

İstersen sana şimdi **goroutine leak (çok fazla goroutine açılması) senaryosu** da gösterebilirim. Onu da ister misin?
EVET
*/
/*
Süper 👍 şimdi de sana **goroutine leak (goroutine sızıntısı)** örneği göstereyim ve bunu `http/pprof` ile nasıl yakalayabileceğimizi anlatayım.

---

## 🔴 1. Goroutine Leak Oluşturan Kod

Aşağıdaki örnekte, her gelen HTTP isteğinde **yeni bir goroutine açılıyor ama asla kapanmıyor**. Bu da zamanla goroutine sayısının artmasına sebep olur.
*/
``go
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func leakyHandler(w http.ResponseWriter, r *http.Request) {
	// Sonsuz çalışan goroutine (kapanmıyor)
	go func() {
		for {
			time.Sleep(time.Second) // sürekli uyuyor ama kapanmıyor
		}
	}()
	fmt.Fprintln(w, "Yeni goroutine açıldı!")
}

func main() {
	http.HandleFunc("/", leakyHandler)

	// pprof endpointlerini aktif et
	fmt.Println("Server başladı: http://localhost:6060/")
	http.ListenAndServe(":6060", nil)
}
``
/*
👉 Her HTTP isteği attığında yeni goroutine açılır. Bir süre sonra yüzlerce/thousands goroutine çalışmaya başlar.

---

## 🔍 2. Goroutine Sayısını İzleme

Tarayıcıdan veya terminalden:
*/
``
http://localhost:6060/debug/pprof/goroutine?debug=1
``

//veya

``
go tool pprof http://localhost:6060/debug/pprof/goroutine
``

//Örneğin:

``
goroutine profile: total 1500
  1498 @ 0x45d3a5 0x47b4b5 0x47b498 0x482ef1
  ...
``
/*
👉 Burada **binlerce goroutine** olduğunu görürsün.

---

## 🟢 3. Sorunu Düzeltme

Sorun şu: handler içinde **sonsuz döngülü goroutine** başlatılıyor.
Çözüm:

* Gereksiz goroutine açmamak,
* Context ile kontrol etmek (goroutine, istek bitince kapanmalı).

Doğru versiyon:
*/
``go
package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func safeHandler(w http.ResponseWriter, r *http.Request) {
	// İstekle ilişkili context
	ctx := r.Context()

	// Bu goroutine context bitince kapanacak
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine kapandı")
				return
			case <-time.After(time.Second):
				fmt.Println("Goroutine çalışıyor...")
			}
		}
	}()

	fmt.Fprintln(w, "Goroutine açıldı ama güvenli şekilde kapanacak.")
}

func main() {
	http.HandleFunc("/", safeHandler)

	fmt.Println("Server başladı: http://localhost:6060/")
	http.ListenAndServe(":6060", nil)
}
``
/*
👉 Artık istek bittiğinde `ctx.Done()` tetiklenir ve goroutine kapanır → leak olmaz.

---

## 📌 Özet

* Goroutine leak, **sürekli yeni goroutine açıp kapatmamak** yüzünden olur.
* `pprof` ile `goroutine` profilini alarak sayıyı görebilirsin.
* Çözüm: **Context**, **kanallar** veya **timeout** ile goroutine yaşam döngüsünü kontrol etmek.

---

İstersen sana son olarak **pprof verilerini canlı grafikte (ör. Flamegraph)** görüntülemeyi nasıl yapabileceğini de gösterebilirim. Onu da ister misin?
*/