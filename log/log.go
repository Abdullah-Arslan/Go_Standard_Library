/*
Golang’daki **`log`** paketi, program çalışırken **log (kayıt) mesajları** üretmek için kullanılan standart kütüphanedir. Özellikle hata ayıklama, uygulama durumunu takip etme ve raporlama için kullanılır. Bu paket, mesajları belirli bir formatta yazdırır, tarih-saat ekler, standart çıkışa (`os.Stdout`) veya hata çıkışına (`os.Stderr`) yönlendirebilir. Ayrıca farklı **Logger** nesneleri oluşturarak dosyalara veya farklı yazıcı hedeflerine loglama yapılabilir.

Şimdi sana bu paketi **tam kapsamlı** olarak açıklayayım:

---

# 📌 `log` Paketinin Temel Özellikleri

* **Mesaj yazdırma**: `Print`, `Printf`, `Println`
* **Hata yazdırma** (çıkış kodu 1 ile biter): `Fatal`, `Fatalf`, `Fatalln`
* **Panik oluşturma**: `Panic`, `Panicf`, `Panicln`
* **Çıkış biçimini değiştirme**: `SetFlags`, `SetPrefix`
* **Özel logger oluşturma**: `log.New`

---

# 📖 Fonksiyonlar ve Örnekler

## 1. `log.Print`, `log.Println`, `log.Printf`

Normal log mesajı yazdırır.
*/
``go
package main

import (
	"log"
)

func main() {
	log.Print("Bu bir log mesajıdır.")            // düz yazı
	log.Println("Bu bir log satırıdır.")          // satır sonuna \n ekler
	log.Printf("Log formatlı mesaj: %d + %d", 2, 3) // printf gibi çalışır
}
``
/*
📌 Çıktı (örnek):

```
2025/09/06 13:10:50 Bu bir log mesajıdır.
2025/09/06 13:10:50 Bu bir log satırıdır.
2025/09/06 13:10:50 Log formatlı mesaj: 2 + 3
```

---

## 2. `log.Fatal`, `log.Fatalln`, `log.Fatalf`

Mesajı yazdırır ve programı **hemen sonlandırır** (`os.Exit(1)`).
*/
``go
package main

import (
	"log"
)

func main() {
	log.Fatal("Ciddi hata: program sonlanıyor!") 
	// Buradan sonrası çalışmaz
	log.Println("Bu satır çalışmaz.")
}
``
/*
---

## 3. `log.Panic`, `log.Panicln`, `log.Panicf`

Mesajı yazdırır ve ardından `panic` tetikler.
*/
``go
package main

import (
	"log"
)

func main() {
	log.Panic("Beklenmedik hata!") 
	// panic oluştuğu için aşağısı çalışmaz
}
``

/7📌 Çıktı:

``
2025/09/06 13:12:00 Beklenmedik hata!
panic: Beklenmedik hata!
``
/*
---

## 4. `log.SetPrefix`

Log mesajlarının başına özel bir ön ek ekler.
*/
``go
package main

import (
	"log"
)

func main() {
	log.SetPrefix("[UYGULAMA] ")
	log.Println("Başlatılıyor...")
}
``
/*
📌 Çıktı:

```
[UYGULAMA] 2025/09/06 13:13:00 Başlatılıyor...
```

---

## 5. `log.SetFlags`

Tarih ve saat biçimlerini ayarlamak için kullanılır.

| Flag                | Açıklama                        |
| ------------------- | ------------------------------- |
| `log.Ldate`         | YYYY/MM/DD tarihi ekler         |
| `log.Ltime`         | HH\:MM\:SS saati ekler          |
| `log.Lmicroseconds` | Mikro saniyeleri ekler          |
| `log.Llongfile`     | Tam dosya yolu + satır numarası |
| `log.Lshortfile`    | Kısa dosya adı + satır numarası |
| `log.LUTC`          | UTC zaman dilimi                |
*/

``go
package main

import (
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Loglama ayarlandı")
}
``
/*
📌 Çıktı:

```
2025/09/06 13:14:10 main.go:10: Loglama ayarlandı
```

---

## 6. Özel Logger Oluşturma (`log.New`)

`log.New` ile kendi **Logger** nesnenizi oluşturup, farklı hedeflere (örneğin dosya) yazdırabilirsiniz.
*/
``go
package main

import (
	"log"
	"os"
)

func main() {
	// log.txt dosyasına yaz
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "[DOSYA LOG] ", log.Ldate|log.Ltime)

	logger.Println("Bu mesaj dosyaya yazıldı")
}
``

//📌 `log.txt` içeriği:

``
[DOSYA LOG] 2025/09/06 13:15:10 Bu mesaj dosyaya yazıldı
``
/*
---

# 🔑 Özet

* `Print`, `Fatal`, `Panic` → farklı şiddette loglar
* `SetPrefix`, `SetFlags` → log biçimlendirme
* `New` → özel log hedefleri (dosya, socket, vb.)
* Varsayılan olarak loglar **stderr**’e yazılır.

---

👉 İstersen sana **log paketini kullanarak hem dosyaya hem ekrana aynı anda log yazdırmayı** da gösterebilirim. İstiyor musun?
EVET


Harika 👍 O zaman sana **hem ekrana hem de dosyaya aynı anda log yazdırmayı** göstereyim.

Bunu yapmak için `io.MultiWriter` kullanırız. Bu fonksiyon birden fazla `io.Writer`’ı birleştirir. Böylece hem `os.Stdout` (ekran) hem de dosya aynı anda yazma hedefi olur.

---

## 📌 Örnek: Ekrana + Dosyaya Log Yazma
*/
``go
package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// log.txt dosyasını aç (yoksa oluştur, varsa ekle)
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Hem ekrana hem dosyaya yazacak writer
	multiWriter := io.MultiWriter(os.Stdout, file)

	// Özel logger
	logger := log.New(multiWriter, "[UYGULAMA] ", log.Ldate|log.Ltime|log.Lshortfile)

	logger.Println("Uygulama başlatıldı")
	logger.Println("Veritabanına bağlanılıyor...")
	logger.Println("İşlem tamamlandı")
}
``
/*
---

## 📌 Çıktı (ekranda göreceğin):

```
[UYGULAMA] 2025/09/06 13:20:15 main.go:20: Uygulama başlatıldı
[UYGULAMA] 2025/09/06 13:20:15 main.go:21: Veritabanına bağlanılıyor...
[UYGULAMA] 2025/09/06 13:20:15 main.go:22: İşlem tamamlandı
```

---

## 📌 `log.txt` Dosyasındaki içerik:

```
[UYGULAMA] 2025/09/06 13:20:15 main.go:20: Uygulama başlatıldı
[UYGULAMA] 2025/09/06 13:20:15 main.go:21: Veritabanına bağlanılıyor...
[UYGULAMA] 2025/09/06 13:20:15 main.go:22: İşlem tamamlandı
```

---

👉 Yani artık loglar **aynı anda hem dosyada saklanıyor hem de konsolda görünüyor** ✅

İstersen sana bu sistemi biraz daha geliştirip **farklı seviyeler (INFO, WARNING, ERROR)** için ayrı log dosyaları oluşturmayı da gösterebilirim. İster misin?
*/