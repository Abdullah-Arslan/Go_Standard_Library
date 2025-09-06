/*
Go’nun standart kütüphanesinde **`log/syslog`** paketi vardır. Bu paket, uygulamanın loglarını **Unix/Linux sistemlerindeki syslog servisine** göndermek için kullanılır.

📌 Yani:

* Logları dosyaya veya ekrana yazmak yerine **sistemin merkezi log servisine** iletebilirsiniz.
* Syslog servisleri (`rsyslog`, `syslogd`, `journald`) logları `/var/log/*` altında tutar.
* Özellikle sunucularda ve mikro servislerde logların merkezi yönetimi için tercih edilir.

⚠️ Not: Bu paket **Windows’ta çalışmaz**, sadece Unix/Linux tabanlı sistemlerde kullanılabilir.

---

# 📌 `log/syslog` Paketinin Yapısı

### 1. **Syslog Seviyeleri (Priority)**

Syslog log seviyelerini tanımlar.

* `LOG_EMERG`   → Sistem kullanılamaz
* `LOG_ALERT`   → Hemen müdahale edilmesi gereken durum
* `LOG_CRIT`    → Kritik hata
* `LOG_ERR`     → Hata
* `LOG_WARNING` → Uyarı
* `LOG_NOTICE`  → Normal ama önemli durum
* `LOG_INFO`    → Bilgilendirme
* `LOG_DEBUG`   → Debug (ayrıntılı bilgi)

### 2. **Facility (Kaynak Uygulama Türü)**

Syslog mesajlarının hangi uygulamadan geldiğini belirtir.

* `LOG_KERN`   → Kernel mesajları
* `LOG_USER`   → Kullanıcı uygulamaları (en sık kullanılan)
* `LOG_MAIL`   → Mail sistemi
* `LOG_DAEMON` → Servis/daemon logları
* `LOG_AUTH`   → Yetkilendirme mesajları
* `LOG_LOCAL0` → Uygulamalar için özel alanlar (LOCAL0–LOCAL7)

---

# 📖 Örnekler

## 1. Basit Syslog Kullanımı
*/
``go
package main

import (
	"log"
	"log/syslog"
)

func main() {
	// Syslog'a bağlan (local syslog)
	writer, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL0, "go-app")
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	// Log yaz
	writer.Info("Uygulama başlatıldı")
	writer.Warning("Disk alanı düşük")
	writer.Err("Veritabanı hatası oluştu")
}
``

//📌 Bu loglar Linux’ta **/var/log/syslog** veya **/var/log/messages** dosyasında görünecektir:

``
Sep  6 14:20:00 myhost go-app: Uygulama başlatıldı
Sep  6 14:20:00 myhost go-app: Disk alanı düşük
Sep  6 14:20:00 myhost go-app: Veritabanı hatası oluştu
``
/*
---

## 2. Syslog + `log.Logger` Kullanımı

Varsayılan `log` paketini syslog ile kullanabiliriz.
*/
``go
package main

import (
	"log"
	"log/syslog"
)

func main() {
	// Syslog writer oluştur
	writer, err := syslog.New(syslog.LOG_DEBUG|syslog.LOG_LOCAL1, "go-service")
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	// log.Logger syslog'a yönlendirildi
	logger := log.New(writer, "[GoApp] ", 0)

	logger.Println("Debug mesajı")
	logger.Println("Hata bulunamadı")
}
``

//📌 Çıktı (syslog üzerinden):

``
Sep  6 14:25:00 myhost go-service[12345]: [GoApp] Debug mesajı
Sep  6 14:25:00 myhost go-service[12345]: [GoApp] Hata bulunamadı
``
/*
---

## 3. Uzak Syslog Sunucusuna Log Gönderme

Syslog, TCP veya UDP üzerinden uzak sunuculara log gönderebilir.
*7
``go
package main

import (
	"log"
	"log/syslog"
)

func main() {
	// Uzak syslog sunucusuna bağlan (UDP 514)
	writer, err := syslog.Dial("udp", "192.168.1.100:514", syslog.LOG_INFO|syslog.LOG_DAEMON, "go-remote")
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	writer.Info("Uzak syslog sunucusuna log gönderildi")
}
``
/*
📌 Bu durumda log mesajı **192.168.1.100** IP’sindeki syslog sunucusuna gider.

---

## 4. Farklı Önceliklerle Yazma
*/
``go
package main

import (
	"log/syslog"
)

func main() {
	writer, _ := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL0, "myapp")
	defer writer.Close()

	writer.Debug("Debug bilgisi")
	writer.Info("Bilgilendirme mesajı")
	writer.Notice("Önemli durum")
	writer.Warning("Uyarı mesajı")
	writer.Err("Hata mesajı")
	writer.Crit("Kritik hata")
	writer.Alert("Acil müdahale gerekli")
	writer.Emerg("Sistem kullanılamaz!")
}
``
/*
---

# 🔑 Özet

* `log/syslog` → logları Linux/Unix syslog servisine gönderir.
* **Priority (seviye)** + **Facility (uygulama türü)** kullanılır.
* `syslog.New()` → yerel syslog
* `syslog.Dial()` → uzak syslog
* `writer.Info()`, `writer.Err()` gibi metodlarla log yazılır.
* `log.New()` ile `syslog.Writer` birleştirilip standart logger gibi kullanılabilir.

---

👉 İstersen sana **slog paketini syslog ile birlikte kullanarak yapılandırılmış logları syslog’a nasıl gönderebileceğini** de gösterebilirim. Bunu ister misin?
EVET
*/
/*
Süper 👍
Şimdi sana **`slog` paketini `syslog` ile birlikte kullanmayı** göstereyim.
Böylece hem modern, yapılandırılmış loglamanın (slog) avantajlarını kullanacağız hem de logları **sistemin syslog servisine** göndereceğiz.

---

# 📌 `slog` + `syslog` Kullanımı

`log/slog` doğrudan syslog desteği vermez ama biz **özel bir `Handler`** yazarak `syslog.Writer` üzerinden logları syslog’a yönlendirebiliriz.

---

## 1. Syslog Handler Tanımı
*/
``go
package main

import (
	"context"
	"log"
	"log/slog"
	"log/syslog"
)

type SyslogHandler struct {
	writer *syslog.Writer
	level  slog.Level
}

// hangi seviyeler aktif
func (h *SyslogHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level
}

// log mesajını syslog'a gönder
func (h *SyslogHandler) Handle(_ context.Context, record slog.Record) error {
	msg := record.Message

	// seviyeye göre syslog fonksiyonu seç
	switch {
	case record.Level >= slog.LevelError:
		return h.writer.Err(msg)
	case record.Level >= slog.LevelWarn:
		return h.writer.Warning(msg)
	case record.Level >= slog.LevelInfo:
		return h.writer.Info(msg)
	default:
		return h.writer.Debug(msg)
	}
}

// WithAttrs ve WithGroup (şimdilik basit tutuyoruz)
func (h *SyslogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}
func (h *SyslogHandler) WithGroup(name string) slog.Handler {
	return h
}
``
/*
---

## 2. Kullanım Örneği
*/
``go
package main

import (
	"log"
	"log/slog"
	"log/syslog"
	"os"
)

func main() {
	// syslog writer oluştur
	writer, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL0, "go-slog-app")
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	// özel handler'ı slog ile bağla
	handler := &SyslogHandler{writer: writer, level: slog.LevelInfo}
	logger := slog.New(handler)

	// test logları
	logger.Info("Uygulama başlatıldı", "port", 8080)
	logger.Warn("Disk alanı düşük", "kalan_mb", 200)
	logger.Error("Veritabanı hatası", "kod", 500)
}
``
/*
---

## 📌 Syslog Üzerindeki Çıktı

Linux’ta `tail -f /var/log/syslog` (veya CentOS’ta `/var/log/messages`) ile görebilirsin:
*/
``
Sep  6 14:45:00 myhost go-slog-app: Uygulama başlatıldı
Sep  6 14:45:00 myhost go-slog-app: Disk alanı düşük
Sep  6 14:45:00 myhost go-slog-app: Veritabanı hatası
``
/*
---

# 🔑 Özet

* `slog` → yapılandırılmış modern loglama
* `syslog.Writer` → logları sistem syslog’una gönderme
* Özel bir `Handler` yazarak slog loglarını syslog’a aktarabiliyoruz.
* Bu sayede hem **modern log API’si (slog)** kullanıyoruz hem de **merkezi log yönetimi (syslog)** avantajını elde ediyoruz.

---

👉 İstersen sana bir de **aynı anda slog loglarını hem syslog’a hem de JSON dosyasına yazdırma** örneğini gösterebilirim. Bunu da ister misin?
EVET
*/
/*
Süper 👍
Şimdi sana **aynı anda slog loglarını hem syslog’a hem de JSON dosyasına** nasıl yazabileceğini göstereyim.

Burada yapacağımız şey:

* `SyslogHandler` (az önce yazdığımız özel handler) → logları syslog’a gönderecek
* `JSONHandler` → logları `app.json` dosyasına yazacak
* `slog.NewMultiHandler` → ikisini birleştirecek

---

# 📌 Örnek: slog → Syslog + JSON Dosyası
*/
``go
package main

import (
	"context"
	"log"
	"log/slog"
	"log/syslog"
	"os"
)

// Özel Syslog Handler
type SyslogHandler struct {
	writer *syslog.Writer
	level  slog.Level
}

func (h *SyslogHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *SyslogHandler) Handle(_ context.Context, record slog.Record) error {
	msg := record.Message

	switch {
	case record.Level >= slog.LevelError:
		return h.writer.Err(msg)
	case record.Level >= slog.LevelWarn:
		return h.writer.Warning(msg)
	case record.Level >= slog.LevelInfo:
		return h.writer.Info(msg)
	default:
		return h.writer.Debug(msg)
	}
}

func (h *SyslogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	// basit örnekte attr'ları syslog'a aktarmıyoruz
	return h
}
func (h *SyslogHandler) WithGroup(name string) slog.Handler {
	return h
}

func main() {
	// 1. Syslog writer
	sysWriter, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL0, "go-multi-app")
	if err != nil {
		log.Fatal(err)
	}
	defer sysWriter.Close()

	// 2. Syslog handler
	sysHandler := &SyslogHandler{writer: sysWriter, level: slog.LevelInfo}

	// 3. JSON dosyası
	file, err := os.OpenFile("app.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	jsonHandler := slog.NewJSONHandler(file, &slog.HandlerOptions{Level: slog.LevelDebug})

	// 4. MultiHandler (syslog + json)
	multiHandler := slog.NewMultiHandler(sysHandler, jsonHandler)

	// 5. Logger
	logger := slog.New(multiHandler)

	// 6. Loglar
	logger.Info("Uygulama başlatıldı", "port", 8080)
	logger.Warn("Disk alanı düşük", "kalan_mb", 200)
	logger.Error("Veritabanı hatası", "kod", 500)
}
``
/*
---

## 📌 Syslog’ta Görünecek (Text format)
*/
``text
Sep  6 14:55:00 myhost go-multi-app: Uygulama başlatıldı
Sep  6 14:55:00 myhost go-multi-app: Disk alanı düşük
Sep  6 14:55:00 myhost go-multi-app: Veritabanı hatası
``
/*
---

## 📌 `app.json` Dosyasında Görünecek (JSON format)
*/
``json
{"time":"2025-09-06T14:55:00.111+03:00","level":"INFO","msg":"Uygulama başlatıldı","port":8080}
{"time":"2025-09-06T14:55:00.111+03:00","level":"WARN","msg":"Disk alanı düşük","kalan_mb":200}
{"time":"2025-09-06T14:55:00.111+03:00","level":"ERROR","msg":"Veritabanı hatası","kod":500}
``
/*
---

✅ Artık loglar **syslog’a** gidiyor (merkezi yönetim için) ve aynı anda **JSON dosyasında** saklanıyor (analiz için).

---

👉 İstersen sana bir de **bu sistemi daha da geliştirip, INFO ve üstünü syslog’a, sadece ERROR seviyelerini JSON dosyasına yazacak şekilde** ayarlamayı gösterebilirim. İstiyor musun?
*/