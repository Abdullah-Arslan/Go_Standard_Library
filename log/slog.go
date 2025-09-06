/*
Go 1.21 ile birlikte gelen **`log/slog` paketi**, eski `log` paketine göre çok daha gelişmiş bir **yapılandırılabilir logging sistemi** sunar.

Kısaca:

* **Structured Logging (yapılandırılmış loglama)** sağlar → yani log mesajları sadece string değil, **anahtar/değer çiftleri** ile tutulur.
* JSON veya Text formatında çıktı verebilir.
* Seviyeli loglama (INFO, WARN, ERROR vb.) destekler.
* Farklı **Handler**’lar ile çıktıyı dosya, stdout, stderr veya başka hedeflere yönlendirebilirsiniz.
* Performans açısından optimize edilmiştir.

---

# 📌 slog Paketinin Temel Yapısı

* **`Logger`** → log yazmak için kullanılan ana nesne.
* **`Handler`** → logların nereye ve nasıl yazılacağını belirler.
* **`Level`** → log seviyeleri (`Debug`, `Info`, `Warn`, `Error`).
* **`Attr`** → log mesajına eklenen key/value çiftleri.
* **`Record`** → log kaydının iç yapısı (genelde kullanıcı doğrudan kullanmaz).

---

# 📖 Önemli Fonksiyonlar ve Kullanımlar

## 1. Basit Kullanım (`slog.Info`, `slog.Error`)

Varsayılan logger (`TextHandler`, `os.Stderr`) kullanır.
*/
``go
package main

import (
	"log/slog"
)

func main() {
	slog.Info("Uygulama başlatıldı")
	slog.Warn("Disk alanı düşük", "kalan_mb", 120)
	slog.Error("Veritabanı hatası", "kod", 500, "detay", "bağlantı başarısız")
}
``
/*
📌 Çıktı (varsayılan `TextHandler`):

```
time=2025-09-06T13:30:00.123+03:00 level=INFO msg="Uygulama başlatıldı"
time=2025-09-06T13:30:00.123+03:00 level=WARN msg="Disk alanı düşük" kalan_mb=120
time=2025-09-06T13:30:00.123+03:00 level=ERROR msg="Veritabanı hatası" kod=500 detay="bağlantı başarısız"
```

---

## 2. JSON Formatında Loglama
*/
``go
package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Kullanıcı girişi", "user", "ahmet", "ip", "192.168.1.10")
}
``
/*
📌 Çıktı (JSON):

```json
{"time":"2025-09-06T13:31:00.456+03:00","level":"INFO","msg":"Kullanıcı girişi","user":"ahmet","ip":"192.168.1.10"}
```

---

## 3. Log Seviyelerini Kullanma
*/
``go
package main

import (
	"log/slog"
)

func main() {
	slog.Debug("Debug mesajı") // Varsayılan logger INFO seviyesinden başlar, bu görünmez
	slog.Info("Bilgilendirme mesajı")
	slog.Warn("Uyarı mesajı")
	slog.Error("Hata mesajı")
}
``
/*
📌 Not: `Debug` görünmez çünkü varsayılan `Level=Info`’dur.

---

## 4. Log Seviyesini Değiştirme
*/
``go
package main

import (
	"log/slog"
	"os"
)

func main() {
	// Seviyeyi DEBUG'e düşür
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	logger := slog.New(handler)

	logger.Debug("Debug mesajı artık görünüyor")
	logger.Info("Bilgi mesajı")
}
``
/*
---

## 5. `With` Kullanımı (Logger’a sabit alan ekleme)

Bir logger’a sabit alanlar ekleyebilirsiniz.
*/
``go
package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	userLogger := logger.With("modül", "kullanıcı", "versiyon", "1.0")
	userLogger.Info("Kullanıcı oluşturuldu", "id", 101, "isim", "Ayşe")
}
``
/*
📌 Çıktı:

```
time=2025-09-06T13:35:00.789+03:00 level=INFO msg="Kullanıcı oluşturuldu" modül=kullanıcı versiyon=1.0 id=101 isim=Ayşe
```

---

## 6. Dosyaya Log Yazma
*/
``go
package main

import (
	"log/slog"
	"os"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	logger := slog.New(slog.NewJSONHandler(file, nil))

	logger.Info("Sunucu başlatıldı", "port", 8080)
	logger.Error("Beklenmedik hata", "detay", "veritabanı bağlantısı")
}
``

//📌 `app.log` içeriği:

``json
{"time":"2025-09-06T13:36:00.111+03:00","level":"INFO","msg":"Sunucu başlatıldı","port":8080}
{"time":"2025-09-06T13:36:00.111+03:00","level":"ERROR","msg":"Beklenmedik hata","detay":"veritabanı bağlantısı"}
``
/*
---

## 7. Özel Handler Yazma

Kendi `Handler`’ınızı yazarak logları farklı sistemlere (ör. Elasticsearch, Kafka, Syslog) gönderebilirsiniz.

Basit örnek (her logu sadece `fmt.Println` ile yazdıran handler):
*/
``go
package main

import (
	"context"
	"fmt"
	"log/slog"
)

type SimpleHandler struct{}

func (h *SimpleHandler) Enabled(_ context.Context, level slog.Level) bool {
	return true
}

func (h *SimpleHandler) Handle(_ context.Context, record slog.Record) error {
	fmt.Printf("[%s] %s\n", record.Level, record.Message)
	return nil
}

func (h *SimpleHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *SimpleHandler) WithGroup(name string) slog.Handler {
	return h
}

func main() {
	logger := slog.New(&SimpleHandler{})
	logger.Info("Selam dünya")
	logger.Error("Hata oluştu")
}
``
/*
📌 Çıktı:

```
[INFO] Selam dünya
[ERROR] Hata oluştu
```

---

# 🔑 Özet

* `slog` → modern, yapılandırılabilir log paketi.
* **Seviyeler**: `Debug`, `Info`, `Warn`, `Error`.
* **Handler**’lar sayesinde farklı formatlar (JSON/Text) ve farklı hedefler (stdout, dosya, özel sistemler).
* **With** → sabit alanlar ekleme.
* Performans ve okunabilirlik açısından `log` paketinden çok daha güçlü.

---

👉 İstersen sana **slog ile aynı anda hem JSON formatlı dosya logu hem de konsol için text logu nasıl alınır** gösterebilirim. İster misin?
EVET
*/
/*
Süper 👍
Şimdi sana **aynı anda hem ekrana (text formatında) hem de dosyaya (JSON formatında) log yazdırmayı** göstereyim. Bunu yapmak için `io.MultiWriter` benzeri bir mantıkla **birden fazla handler’ı birleştireceğiz**.

---

# 📌 Örnek: slog ile Hem Konsol (Text) Hem Dosya (JSON) Loglama
*/
``go
package main

import (
	"log/slog"
	"os"
)

func main() {
	// 1. Dosyayı aç
	file, err := os.OpenFile("app.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 2. Konsol için Text handler
	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})

	// 3. Dosya için JSON handler
	jsonHandler := slog.NewJSONHandler(file, &slog.HandlerOptions{Level: slog.LevelDebug})

	// 4. İki handler’ı birleştiren MultiHandler
	handler := slog.NewMultiHandler(textHandler, jsonHandler)

	// 5. Logger oluştur
	logger := slog.New(handler)

	// 6. Log yaz
	logger.Info("Uygulama başlatıldı", "port", 8080)
	logger.Warn("Disk alanı düşük", "kalan_mb", 512)
	logger.Error("Beklenmedik hata", "kod", 500, "detay", "veritabanı bağlantısı")
}
``
/*
---

## 📌 Konsolda Görülecek (Text format)

```
time=2025-09-06T13:45:00.123+03:00 level=INFO msg="Uygulama başlatıldı" port=8080
time=2025-09-06T13:45:00.123+03:00 level=WARN msg="Disk alanı düşük" kalan_mb=512
time=2025-09-06T13:45:00.123+03:00 level=ERROR msg="Beklenmedik hata" kod=500 detay="veritabanı bağlantısı"
```

---

## 📌 `app.json` İçeriği

```json
{"time":"2025-09-06T13:45:00.123+03:00","level":"INFO","msg":"Uygulama başlatıldı","port":8080}
{"time":"2025-09-06T13:45:00.123+03:00","level":"WARN","msg":"Disk alanı düşük","kalan_mb":512}
{"time":"2025-09-06T13:45:00.123+03:00","level":"ERROR","msg":"Beklenmedik hata","kod":500,"detay":"veritabanı bağlantısı"}
```

---

✅ Artık loglar **ekranda okunabilir text formatında**, aynı anda da **dosyada JSON formatında** tutuluyor. Bu sayede:

* **Developer (sen)** konsolda kolayca okuyabilir.
* **Sistem** (örneğin bir log toplayıcı) JSON veriyi işleyebilir.

---

👉 İstersen sana bir de **slog ile INFO ve üstü konsola, sadece ERROR seviyelerini ayrı bir hata.log dosyasına yazdırmayı** gösterebilirim. İstiyor musun?
EVET
*/
/*
Harika 👍
Şimdi sana **slog ile farklı seviyelerde logları farklı yerlere yazdırmayı** göstereyim:

* **INFO ve üstü (INFO, WARN, ERROR)** → konsola (text formatında)
* **Sadece ERROR logları** → `hata.log` dosyasına (JSON formatında)

Bunu yaparken `slog.NewLevelVar` ve `HandlerOptions` ile seviyeleri kontrol edeceğiz.

---

# 📌 Örnek: Farklı Seviyeler için Farklı Çıkışlar
*/
``go
package main

import (
	"log/slog"
	"os"
)

func main() {
	// 1. Hata dosyasını aç
	errorFile, err := os.OpenFile("hata.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer errorFile.Close()

	// 2. Konsol için handler (INFO ve üstü görünsün)
	consoleHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	// 3. Hata dosyası için handler (sadece ERROR yazsın)
	errorHandler := slog.NewJSONHandler(errorFile, &slog.HandlerOptions{
		Level: slog.LevelError,
	})

	// 4. MultiHandler ile birleştir
	handler := slog.NewMultiHandler(consoleHandler, errorHandler)

	// 5. Logger oluştur
	logger := slog.New(handler)

	// 6. Test logları
	logger.Debug("Debug mesajı - görünmez") // Konsolda da dosyada da çıkmaz
	logger.Info("Bilgilendirme mesajı", "işlem", "başlatma") // Sadece konsolda
	logger.Warn("Uyarı mesajı", "disk", "düşük") // Sadece konsolda
	logger.Error("Kritik hata!", "kod", 500, "detay", "veritabanı") // Konsolda + hata.log dosyasında
}
``
/*
---

## 📌 Konsolda Görülecek (Text format)

```
time=2025-09-06T13:55:00.123+03:00 level=INFO msg="Bilgilendirme mesajı" işlem=başlatma
time=2025-09-06T13:55:00.123+03:00 level=WARN msg="Uyarı mesajı" disk=düşük
time=2025-09-06T13:55:00.123+03:00 level=ERROR msg="Kritik hata!" kod=500 detay=veritabanı
```

---

## 📌 `hata.log` Dosyası (JSON format)
*/
``json
{"time":"2025-09-06T13:55:00.123+03:00","level":"ERROR","msg":"Kritik hata!","kod":500,"detay":"veritabanı"}
``
/*
---

✅ Böylece:

* Konsolda **INFO, WARN, ERROR** seviyeleri okunabilir şekilde duruyor.
* Dosyada sadece **kritik hatalar (ERROR)** JSON formatında saklanıyor.

---

👉 İstersen sana bir sonraki adımda **slog ile logları günlük dosyalara ayırmayı (ör. app-2025-09-06.log)** da gösterebilirim. İstiyor musun?
EVET
*/
/*
Harika 👍
Şimdi sana **slog ile günlük dosyalara log kaydı tutmayı** göstereyim.

📌 Ama önemli bir nokta:
Go’nun standart `slog` paketinde doğrudan “daily log rotation” (günlük log döndürme) özelliği **yoktur**.
Bunu yapmak için:

1. Dosya adını tarih bazlı verip her gün yeni dosya açabiliriz.
2. Daha gelişmiş ihtiyaçlar için `lumberjack` gibi harici paketler kullanılır.

Ben burada **saf slog + os paketleriyle günlük log dosyası** nasıl yapılır onu göstereceğim.

---

# 📌 Örnek: Günlük Dosyalara Log Yazma
*/
``go
package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"
)

func createDailyLogger() *slog.Logger {
	// Bugünün tarihini dosya adı olarak kullan
	today := time.Now().Format("2006-01-02") // YYYY-MM-DD
	filename := fmt.Sprintf("app-%s.log", today)

	// Dosyayı aç
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	// JSON formatlı handler
	jsonHandler := slog.NewJSONHandler(file, &slog.HandlerOptions{Level: slog.LevelInfo})

	// Logger döndür
	return slog.New(jsonHandler)
}

func main() {
	logger := createDailyLogger()

	logger.Info("Uygulama başlatıldı", "port", 8080)
	logger.Warn("Disk alanı azaldı", "kalan_mb", 256)
	logger.Error("Bağlantı hatası", "kod", 500)

	// Örnek olarak başka bir gün için log alalım (test amaçlı)
	// Normalde cron job ya da uygulama restartı ile yeni dosya açılır
	time.Sleep(2 * time.Second)
}
``
/*
---

## 📌 Dosya çıktıları

Bugün **6 Eylül 2025** ise dosya adı:

```
app-2025-09-06.log
```

İçerik:

```json
{"time":"2025-09-06T14:05:00.111+03:00","level":"INFO","msg":"Uygulama başlatıldı","port":8080}
{"time":"2025-09-06T14:05:00.111+03:00","level":"WARN","msg":"Disk alanı azaldı","kalan_mb":256}
{"time":"2025-09-06T14:05:00.111+03:00","level":"ERROR","msg":"Bağlantı hatası","kod":500}
```

---

# 🔑 Özet

* Dosya adını `YYYY-MM-DD` formatıyla oluşturursak günlük log dosyaları elde ederiz.
* Uygulama her açıldığında `app-2025-09-06.log`, `app-2025-09-07.log` gibi dosyalara yazmaya başlar.
* Daha gelişmiş “otomatik döndürme” (belirli boyuta gelince yeni dosya, eski dosyaları silme) için harici kütüphaneler (`gopkg.in/natefinch/lumberjack.v2`) kullanılır.

---

👉 İstersen sana bir de **slog + lumberjack kullanarak boyut sınırına göre otomatik dönen log dosyaları** örneğini de gösterebilirim. Görmek ister misin?
*/