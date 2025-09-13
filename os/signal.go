/*
Go’da **`os/signal`** paketi, işletim sisteminden gelen **sinyalleri yakalamak ve yönetmek** için kullanılır.
UNIX/Linux sistemlerinde sinyaller; `SIGINT`, `SIGTERM`, `SIGHUP`, `SIGKILL` gibi olaylardır.
Örneğin:

* `CTRL+C` → `SIGINT` gönderir.
* `kill -TERM <pid>` → `SIGTERM` gönderir.
* Sunucu uygulamaları bu sinyalleri yakalayarak **temiz kapanış** (graceful shutdown) yapabilir.

---

# 🔹 Paket Fonksiyonları

### 1. `signal.Notify(c chan<- os.Signal, sig ...os.Signal)`

Belirtilen sinyalleri yakalar ve `c` kanalına gönderir.

### 2. `signal.Stop(c chan<- os.Signal)`

Bir kanala sinyal gönderimini durdurur.

### 3. `signal.Ignore(sig ...os.Signal)`

Belirtilen sinyalleri tamamen yok sayar.

### 4. `signal.Reset(sig ...os.Signal)`

Sinyal davranışını varsayılana döndürür.

---

# 🔹 Örnekler

## 1. Basit SIGINT Yakalama (`CTRL+C`)
*/
``go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Kanal oluştur
	sigChan := make(chan os.Signal, 1)

	// SIGINT (CTRL+C) ve SIGTERM yakala
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Program çalışıyor... Çıkmak için CTRL+C basın")

	// Sinyal bekle
	sig := <-sigChan
	fmt.Println("Sinyal alındı:", sig)

	fmt.Println("Program kapanıyor...")
}
``

//📌 Çalıştırıp `CTRL+C` yaptığında:

``
Program çalışıyor... Çıkmak için CTRL+C basın
Sinyal alındı: interrupt
Program kapanıyor...
``
/*
---

## 2. Graceful Shutdown (Web Sunucusu Örneği)
*
``go
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := &http.Server{Addr: ":8080"}

	// Sunucuyu arka planda başlat
	go func() {
		fmt.Println("Sunucu çalışıyor :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Sunucu hatası:", err)
		}
	}()

	// Sinyal yakalama
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	fmt.Println("🔴 Sinyal alındı, sunucu kapatılıyor...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Zorla kapatma:", err)
	}

	fmt.Println("Sunucu kapandı ✅")
}
``
/*
📌 Bu sayede sunucuya `CTRL+C` veya `kill` sinyali gönderildiğinde:

* Yeni bağlantılar kabul edilmez.
* Açık bağlantılar kapanır.
* Sonra program kapanır.

---

## 3. Belirli Sinyali Yoksayma (`signal.Ignore`)
*
``go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// SIGINT sinyalini yok say
	signal.Ignore(os.Interrupt)

	fmt.Println("CTRL+C çalışmayacak, kapatmak için kill -9 PID kullanın")

	for {
		fmt.Println("Çalışıyor...")
		time.Sleep(2 * time.Second)
	}
}
``
/*
📌 Bu durumda `CTRL+C` işe yaramaz, sadece `kill -9` ile durdurabilirsin.

---

## 4. Varsayılan Davranışı Geri Getirme (`signal.Reset`)
*/

``go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)

	go func() {
		for sig := range c {
			fmt.Println("Yakalanan sinyal:", sig)
			fmt.Println("Şimdi resetleniyor, bir sonraki CTRL+C programı sonlandıracak")
			signal.Reset(syscall.SIGINT) // Varsayılana döndür
		}
	}()

	select {}
}
``
/*
📌 İlk `CTRL+C` sadece mesaj basar.
İkinci `CTRL+C` programı **normal şekilde kapatır**.

---

# 🔹 Özet

* **`signal.Notify`** → Sinyalleri kanala gönderir.
* **`signal.Stop`** → Kanala gönderimi durdurur.
* **`signal.Ignore`** → Sinyali yok sayar.
* **`signal.Reset`** → Varsayılan davranışı geri getirir.

Bu paket genellikle:

* **Sunucuların graceful shutdown** yapması,
* **Daemon/servis yönetimi**,
* **Özel sinyal yakalama** (örn. `SIGHUP` ile config dosyası yeniden yükleme)

senaryolarında çok kullanılır ✅

---

👉 İstersen sana `signal` paketinin fonksiyonlarını tablo halinde (Fonksiyon + Açıklama + Örnek Kod) **Excel/PDF** formatında hazırlayabilirim.
Bunu da ister misin?
EVET

👉 İstersen buna ek olarak bir de gerçek dünya mini proje (ör. config dosyası değişince SIGHUP ile yeniden yükleme) örneği yazayım mı?
EVET

Süper 👍 Şimdi sana `os/signal` paketinin **gerçek dünya senaryosunda** nasıl kullanılacağını göstereyim.

Bu örnekte:

* Bir **config dosyası** (`config.txt`) olacak.
* Program ilk açıldığında bu dosyayı okuyacak.
* Eğer `SIGHUP` sinyali (`kill -HUP <pid>`) gelirse config dosyasını **yeniden yükleyecek**.
* Böylece programı kapatmadan ayarları değiştirebilirsin (örneğin web sunucularında sık kullanılan bir teknik).

---

# 🔹 Örnek: SIGHUP ile Config Dosyası Yeniden Yükleme

## 📄 `main.go`
*/
``go
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Config yükleme fonksiyonu
func loadConfig(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content string
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}
	return content, scanner.Err()
}

func main() {
	configFile := "config.txt"

	// İlk yükleme
	config, err := loadConfig(configFile)
	if err != nil {
		fmt.Println("Config okunamadı:", err)
		return
	}
	fmt.Println("🔹 İlk config yüklendi:\n", config)

	// Sinyal kanalı
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	for {
		sig := <-sigChan
		switch sig {
		case syscall.SIGHUP:
			fmt.Println("🔄 SIGHUP alındı → config yeniden yükleniyor...")
			config, err = loadConfig(configFile)
			if err != nil {
				fmt.Println("Config okunamadı:", err)
			} else {
				fmt.Println("Yeni config:\n", config)
			}
		case syscall.SIGINT, syscall.SIGTERM:
			fmt.Println("🔴 Program kapanıyor (sinyal:", sig, ")")
			return
		}
	}
}
``
/*
---

## 📄 `config.txt`
*/

``
APP_NAME=MyApp
APP_PORT=8080
DEBUG=true
``
/*
---

## 🖥 Kullanım

1. Programı çalıştır:
*/
   ``bash
   go run main.go
   ``

   //Çıktı:

   ``
   🔹 İlk config yüklendi:
   APP_NAME=MyApp
   APP_PORT=8080
   DEBUG=true
   ``
/*
2. Config dosyasını değiştir (örneğin `DEBUG=false` yap).

3. Programı kapatmadan config’i yenile:
*/
   ``bash
   kill -HUP <pid>
   ``

   //Çıktı:

   ``
   🔄 SIGHUP alındı → config yeniden yükleniyor...
   Yeni config:
   APP_NAME=MyApp
   APP_PORT=8080
   DEBUG=false
   ``
/*
---

✅ Böylece `os/signal` paketi ile **runtime sırasında ayarları değiştirme** imkânın oldu.
Bu yöntem özellikle **web sunucuları, daemon’lar, log seviyesi değiştirme, cache temizleme** gibi durumlarda çok kullanılır.

👉 İstersen bu örneği biraz daha ileri götürüp, **HTTP sunucusu ile birlikte (SIGHUP ile log level / ayar değişimi)** çalışan bir proje de yapabilirim. İster misin?
EVET
*/
/*
Süper 👍 Şimdi sana **os/signal + HTTP sunucusu + canlı config değişimi** örneğini hazırlayayım.

Bu projede:

* Küçük bir **HTTP sunucusu** açacağız (`localhost:8080`).
* Bir `config.txt` dosyasından **ayarlar** (ör. `APP_PORT`, `LOG_LEVEL`) yüklenecek.
* Program çalışırken `kill -HUP <pid>` gönderirsen → **config dosyası yeniden okunacak**.
* HTTP sunucusu yeni config ile çalışmaya devam edecek (program kapanmadan).

---

# 🔹 Örnek: SIGHUP ile Canlı Config Yenileme + HTTP Sunucusu

## 📄 `main.go`
*/
``go
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

// Global config yapısı
type Config struct {
	AppName  string
	AppPort  string
	LogLevel string
}

var (
	config     Config
	configFile = "config.txt"
	mu         sync.RWMutex // eşzamanlı erişim için kilit
)

// Config yükleme fonksiyonu
func loadConfig(filename string) (Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cfg := Config{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key, value := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
		switch key {
		case "APP_NAME":
			cfg.AppName = value
		case "APP_PORT":
			cfg.AppPort = value
		case "LOG_LEVEL":
			cfg.LogLevel = value
		}
	}
	return cfg, scanner.Err()
}

// HTTP handler
func handler(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()
	fmt.Fprintf(w, "App: %s\nPort: %s\nLogLevel: %s\n",
		config.AppName, config.AppPort, config.LogLevel)
}

func main() {
	// İlk config yükleme
	var err error
	config, err = loadConfig(configFile)
	if err != nil {
		fmt.Println("Config okunamadı:", err)
		return
	}
	fmt.Println("🔹 Config yüklendi:", config)

	// HTTP server
	http.HandleFunc("/", handler)
	go func() {
		fmt.Println("🌐 Sunucu çalışıyor http://localhost:" + config.AppPort)
		if err := http.ListenAndServe(":"+config.AppPort, nil); err != nil {
			fmt.Println("Sunucu hatası:", err)
		}
	}()

	// Sinyal yakalama
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	for {
		sig := <-sigChan
		switch sig {
		case syscall.SIGHUP:
			fmt.Println("🔄 SIGHUP alındı → config yeniden yükleniyor...")
			newConfig, err := loadConfig(configFile)
			if err != nil {
				fmt.Println("Config okunamadı:", err)
				continue
			}
			mu.Lock()
			config = newConfig
			mu.Unlock()
			fmt.Println("✅ Yeni config:", config)
		case syscall.SIGINT, syscall.SIGTERM:
			fmt.Println("🔴 Program kapanıyor (sinyal:", sig, ")")
			return
		}
	}
}
``

---

//## 📄 `config.txt`

``
APP_NAME=MyApp
APP_PORT=8080
LOG_LEVEL=DEBUG
``
/*
---

## 🖥 Kullanım

1. Programı çalıştır:
*/

   ``bash
   go run main.go
   ``

   ``
   🔹 Config yüklendi: {MyApp 8080 DEBUG}
   🌐 Sunucu çalışıyor http://localhost:8080
   ``

//2. Tarayıcıda aç: [http://localhost:8080](http://localhost:8080)

   ``
   App: MyApp
   Port: 8080
   LogLevel: DEBUG
   ``
/*
3. `config.txt` dosyasını değiştir (örneğin `LOG_LEVEL=INFO` yap).

4. PID bul:
*/
   ``bash
   ps aux | grep main
   ``

   //ardından:

   ``bash
   kill -HUP <pid>
   ``

   //Konsolda:

   ``
   🔄 SIGHUP alındı → config yeniden yükleniyor...
   ✅ Yeni config: {MyApp 8080 INFO}
   ``
/*
5. Tarayıcıyı yenile → Yeni config değerini görürsün.

---

✅ Böylece canlı çalışan HTTP sunucun, kapanmadan **config reload** yapabiliyor.
Bu teknik, özellikle **Nginx, Apache, Docker daemon, Go microservices** gibi yazılımların kullandığı yöntemdir.

---

👉 İstiyor musun, ben bu projeyi daha da ileri götürüp **birden fazla config dosyasını izleyen (hot-reload)** bir sürüm hazırlayayım?
*/
