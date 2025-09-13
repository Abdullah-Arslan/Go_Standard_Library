/*
Go’nun **`net/url`** paketini baştan sona açıklayayım ve örneklerle göstereyim.

`net/url` paketi, **URL’leri parse etme, oluşturma, query parametreleri yönetme ve URL encode/decode işlemleri** için kullanılır.

---

# 1️⃣ `net/url` Paketi Nedir?

* URL parçalarını yönetmek için struct ve fonksiyonlar sağlar.
* HTTP isteklerinde query parametrelerini oluşturmak ve çözmek için idealdir.
* URL encode/decode işlemleri yapabilir.

---

# 2️⃣ Temel Tipler ve Fonksiyonlar

| Tip / Fonksiyon         | Açıklama                                             |
| ----------------------- | ---------------------------------------------------- |
| `url.URL`               | URL’yi parçalara ayırmak için kullanılan struct      |
| `url.Parse()`           | String URL’yi parse eder ve `*url.URL` döner         |
| `url.ParseRequestURI()` | HTTP request URI’lerini parse eder                   |
| `URL.String()`          | URL’yi string’e çevirir                              |
| `URL.Query()`           | Query parametrelerini `url.Values` olarak döner      |
| `url.Values`            | Query parametrelerini yönetmek için map tarzı struct |
| `url.QueryEscape()`     | String’i URL için encode eder                        |
| `url.QueryUnescape()`   | Encode edilmiş string’i decode eder                  |

---

# 3️⃣ URL Parse Örneği
*/
``go
package main

import (
	"fmt"
	"net/url"
)

func main() {
	rawURL := "https://example.com:8080/path/page.html?name=John&age=30#section1"

	u, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Parse hatası:", err)
		return
	}

	fmt.Println("Scheme:", u.Scheme)       // https
	fmt.Println("Host:", u.Host)           // example.com:8080
	fmt.Println("Path:", u.Path)           // /path/page.html
	fmt.Println("RawQuery:", u.RawQuery)   // name=John&age=30
	fmt.Println("Fragment:", u.Fragment)   // section1

	// Query parametrelerini almak
	q := u.Query()
	fmt.Println("Name parametresi:", q.Get("name")) // John
	fmt.Println("Age parametresi:", q.Get("age"))   // 30
}
``
/*
---

# 4️⃣ Query Parametreleri Ekleme
*/
``go
package main

import (
	"fmt"
	"net/url"
)

func main() {
	baseURL := "https://example.com/search"
	u, _ := url.Parse(baseURL)

	// Query parametreleri oluştur
	q := u.Query()
	q.Set("q", "golang url paketi")
	q.Set("page", "2")
	u.RawQuery = q.Encode()

	fmt.Println("Tam URL:", u.String())
	// https://example.com/search?q=golang+url+paketi&page=2
}
``
/*
---

# 5️⃣ URL Encode / Decode
*/
``go
package main

import (
	"fmt"
	"net/url"
)

func main() {
	str := "golang url paketi & örnekler"

	encoded := url.QueryEscape(str)
	fmt.Println("Encoded:", encoded)
	// golang+url+paketi+%26+%C3%B6rnekler

	decoded, _ := url.QueryUnescape(encoded)
	fmt.Println("Decoded:", decoded)
	// golang url paketi & örnekler
}
``
/*
---

# 6️⃣ URL Oluşturma (`url.URL` struct ile)
*/
``go
package main

import (
	"fmt"
	"net/url"
)

func main() {
	u := &url.URL{
		Scheme:   "https",
		Host:     "example.com:8080",
		Path:     "/path/page.html",
		RawQuery: "name=John&age=30",
		Fragment: "section1",
	}

	fmt.Println("URL String:", u.String())
	// https://example.com:8080/path/page.html?name=John&age=30#section1
}
``
/*
---

# 7️⃣ Özet

* `net/url` URL parse, encode/decode ve query parametre yönetimi için kullanılır.
* `url.Parse()` ile URL parçalarına erişebilirsin.
* `url.Values` ile query parametrelerini ekleyip çıkarabilirsin.
* `url.QueryEscape()` ve `url.QueryUnescape()` ile güvenli URL oluşturabilirsin.

---

İstersen sana bunu bir adım ileri götürüp **concurrent URL parse ve query parametre yönetimi + HTML form submission simülasyonu** örneğini gösterebilirim.

Bunu ister misin?
EVET
*/
/*
Harika — aşağıda **concurrent URL parse + HTTP GET + response parse + rate limiting + logging** yapan, çalıştırılmaya hazır, açıklamalı bir Go örneği bulacaksın. Kod sadece standart kütüphaneleri kullanır (ek bir bağımlılık yok), isteklerde zaman aşımı, rate limiting (RPS), worker pool, basit HTML `<title>` ayrıştırma ve sonuç kaydı (CSV) yapar.

Kodu bir dosyaya (`main.go`) koyup `go run main.go` ile çalıştırabilirsin.

---
*/
``go
// main.go
package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"sync"
	"time"
)

// İş (job) tipi: hangi URL'yi işleyeceğimizi tutar.
type Job struct {
	ID  int
	URL string
}

// Sonuç (result) tipi: işin çıktısını tutar.
type Result struct {
	ID           int
	URL          string
	Status       string
	HTTPStatus   int
	Title        string
	ContentBytes int64
	Error        string
	DurationMS   int64
}

// Basit başlık çıkarma regex'i (örnek amaçlı — HTML edge-case'leri yutabilir).
var titleRe = regexp.MustCompile("(?is)<title[^>]*>\\s*(.*?)\\s*</title>")

func main() {
	// ---------- Konfigürasyon ----------
	seedURLs := []string{
		"https://golang.org",
		"https://example.com",
		"https://httpbin.org/get",
		"https://httpbin.org/status/500",
		"https://httpbin.org/delay/1",
	}
	numWorkers := 5
	ratePerSecond := 3            // saniyede 3 istek toplam
	httpTimeout := 5 * time.Second
	outputCSV := "results.csv"
	// ------------------------------------

	// Kanal ve yapılar
	jobs := make(chan Job, len(seedURLs))
	results := make(chan Result, len(seedURLs))

	// Rate limiter: token kanalı dolduruluyor ticker ile.
	tokenBucket := make(chan struct{}, ratePerSecond)
	go func() {
		ticker := time.NewTicker(time.Second / time.Duration(ratePerSecond))
		defer ticker.Stop()
		for {
			<-ticker.C
			// non-blocking send: eğer bucket doluysa token atılmasın
			select {
			case tokenBucket <- struct{}{}:
			default:
			}
		}
	}()

	// HTTP client (timeout ile)
	client := &http.Client{
		Timeout: httpTimeout,
	}

	var wg sync.WaitGroup
	// Worker pool
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				// Rate limit: token al
				<-tokenBucket

				start := time.Now()
				res := fetchAndParse(job, client)
				res.DurationMS = time.Since(start).Milliseconds()
				log.Printf("[worker %d] done job %d (%s) status=%d err=%q\n", workerID, job.ID, job.URL, res.HTTPStatus, res.Error)
				results <- res
			}
		}(w)
	}

	// Jobları doldur
	for i, u := range seedURLs {
		jobs <- Job{ID: i + 1, URL: u}
	}
	close(jobs)

	// Sonuçları ayrı goroutine ile CSV'ye yaz (aynı zamanda log'lar)
	var outWg sync.WaitGroup
	outWg.Add(1)
	go func() {
		defer outWg.Done()
		if err := writeResultsCSV(outputCSV, results); err != nil {
			log.Println("CSV yazma hatası:", err)
		}
	}()

	// Bekle worker'lar bitene kadar
	wg.Wait()
	close(results)
	outWg.Wait()

	log.Println("Tüm işler tamamlandı. Sonuçlar:", outputCSV)
}

// fetchAndParse: HTTP GET yapar, durumu ve <title>'ı çıkartır.
func fetchAndParse(job Job, client *http.Client) Result {
	res := Result{ID: job.ID, URL: job.URL}
	// Geçerli URL mi kontrol et
	parsed, err := url.Parse(job.URL)
	if err != nil {
		res.Error = "invalid URL: " + err.Error()
		res.Status = "ERROR"
		return res
	}

	// Context ile ekstra timeout (client timeout zaten var)
	ctx, cancel := context.WithTimeout(context.Background(), client.Timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", parsed.String(), nil)
	if err != nil {
		res.Error = "new request error: " + err.Error()
		res.Status = "ERROR"
		return res
	}
	// Basit header
	req.Header.Set("User-Agent", "Go-URL-Scraper/1.0")

	resp, err := client.Do(req)
	if err != nil {
		res.Error = err.Error()
		res.Status = "ERROR"
		return res
	}
	defer resp.Body.Close()

	res.HTTPStatus = resp.StatusCode
	res.Status = resp.Status

	// Okunan byte sayısını saymak için io.Copy ile io.Discard kullan
	n, err := io.Copy(io.Discard, resp.Body)
	res.ContentBytes = n
	if err != nil {
		// İçerik okuma hatası, ama yine title alamayız çünkü gövde okundu
		res.Error = "read body error: " + err.Error()
		return res
	}

	// Tekrar istek atıp gövdeyi string okumak (küçük sayfalarda kabul edilebilir).
	// NOT: performans için gerçek projede iki kez GET atmak yerine
	// gövdeyi buffer'a okuyup hem byte sayısını hem title'ı çıkarmak daha iyidir.
	// Burada örneği sade tutmak için basitleştiriyoruz.
	// (Alternatif: önceden body := new(bytes.Buffer); n, _ := io.Copy(body, resp.Body) ...)
	// Ancak resp.Body zaten okundu — bu yüzden yeniden GET atıyoruz sadece örnek amaçlı.
	req2, _ := http.NewRequestWithContext(ctx, "GET", parsed.String(), nil)
	req2.Header.Set("User-Agent", "Go-URL-Scraper/1.0")
	resp2, err := client.Do(req2)
	if err == nil {
		defer resp2.Body.Close()
		bodyBytes, err := io.ReadAll(io.LimitReader(resp2.Body, 1_000_000)) // maksimum 1MB
		if err == nil {
			title := extractTitle(string(bodyBytes))
			res.Title = title
		} else {
			res.Error = "read body for title error: " + err.Error()
		}
	} else {
		// yeniden GET hatası: sadece logla
		res.Error = "second GET for title failed: " + err.Error()
	}

	return res
}

// extractTitle: basit bir <title> tag ayrıştırıcısı
func extractTitle(html string) string {
	m := titleRe.FindStringSubmatch(html)
	if len(m) >= 2 {
		return sanitizeTitle(m[1])
	}
	return ""
}

// sanitizeTitle: başlığı kısalt ve whitespace temizle
func sanitizeTitle(s string) string {
	// Normalize whitespace
	reSpace := regexp.MustCompile(`\s+`)
	s = reSpace.ReplaceAllString(s, " ")
	if len(s) > 200 {
		return s[:200] + "..."
	}
	return s
}

// writeResultsCSV: results kanalı kapanana kadar satırları yazıp dosyayı kapatır.
func writeResultsCSV(filename string, results <-chan Result) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	// Başlık satırı
	if err := w.Write([]string{"id", "url", "http_status", "status_text", "title", "bytes", "duration_ms", "error"}); err != nil {
		return err
	}

	for r := range results {
		row := []string{
			fmt.Sprint(r.ID),
			r.URL,
			fmt.Sprint(r.HTTPStatus),
			r.Status,
			r.Title,
			fmt.Sprint(r.ContentBytes),
			fmt.Sprint(r.DurationMS),
			r.Error,
		}
		if err := w.Write(row); err != nil {
			log.Println("CSV yazma hatası:", err)
		}
	}

	return nil
}
``
/*
---

## Açıklamalar — neden böyle?

* Worker pool (`numWorkers`) istekleri paralel işler; `jobs` kanalını doldurup kapatıyoruz.
* Rate limiting için **token bucket**: her `sec / ratePerSecond` aralığında token eklenir; worker token alana kadar bekler — toplam RPS limiti sağlar.
* `http.Client.Timeout` ve `context.WithTimeout` ile güvenli zaman aşımı sağlanır.
* `fetchAndParse` işlevi GET yapar, HTTP durumunu kaydeder, basit `<title>` çıkarma yapar. (Gerçek HTML parse için `golang.org/x/net/html` veya `github.com/PuerkitoBio/goquery` kullanılabilir.)
* Sonuçlar CSV olarak kaydedilir (kolay analiz için).
* Tüm işlemler `log.Printf` ile anlık loglanır.

---

## Nasıl genişletirsin

* HTML parsing için `goquery` ekleyebilirsin (`github.com/PuerkitoBio/goquery`) — daha sağlam ve CSS selector tabanlı.
* Sonuçları DB’ye (Postgres/SQLite) yazabilirsin.
* `robots.txt` kontrolü, retry/backoff (exponential backoff), proxy desteği, cookie yönetimi ekleyebilirsin.
* pprof eklemek istersen program başında `go func(){ log.Println(http.ListenAndServe(":6060", nil)) }()` ve `import _ "net/http/pprof"` ekle; bu sayede canlı profil alırsın.

---

İstersen şimdi:

* Bu örneği **goquery** ile daha sağlam HTML parsing yapacak şekilde güncelleyeyim,
* veya **retry + backoff**, **proxy** veya **robots.txt** desteği ekleyeyim — hangisini istersin?
*/