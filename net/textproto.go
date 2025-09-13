/*
Go’nun **`net/textproto`** paketini baştan sona açıklayayım ve örneklerle göstereyim.

`net/textproto` paketi, **metin tabanlı protokollerle (HTTP, SMTP, POP3, IMAP, FTP vb.) çalışmayı kolaylaştıran bir paket**tir.

---

# 1️⃣ `net/textproto` Paketi Nedir?

* Metin tabanlı protokollerde **satır satır veri okuma/yazma** işini kolaylaştırır.
* **Header yönetimi** (ör. HTTP header) sağlar.
* TCP veya diğer bağlantılar üzerinden metin tabanlı protokolleri uygularken kullanılır.

---

# 2️⃣ Temel Özellikler

| Fonksiyon / Tip          | Açıklama                                 |
| ------------------------ | ---------------------------------------- |
| `textproto.Conn`         | Metin tabanlı bağlantı yönetimi sağlar   |
| `Reader.ReadLine`        | Satır satır okuma yapar                  |
| `Reader.ReadMIMEHeader`  | Header alanlarını parse eder             |
| `Writer.PrintfLine`      | Satır yazma işlemi yapar                 |
| `CanonicalMIMEHeaderKey` | Header key’lerini standart forma çevirir |

---

# 3️⃣ Basit Örnek: Satır Satır Veri Okuma/Yazma
*/
``go
package main

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
)

func main() {
	// TCP server
	ln, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println("Listen hatası:", err)
		return
	}
	defer ln.Close()
	fmt.Println("Server 9090 portunda çalışıyor...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept hatası:", err)
			continue
		}

		go func(c net.Conn) {
			defer c.Close()

			tpConn := textproto.NewConn(c)

			// Satır okuma
			line, err := tpConn.ReadLine()
			if err != nil {
				fmt.Println("ReadLine hatası:", err)
				return
			}
			fmt.Println("Alınan satır:", line)

			// Satır yazma
			tpConn.PrintfLine("Merhaba, aldım: %s", line)
		}(conn)
	}
}
``
/*
---

# 4️⃣ Header Okuma Örneği (HTTP/SMTP tarzı)
*/
``go
package main

import (
	"fmt"
	"net"
	"net/textproto"
	"strings"
)

func main() {
	raw := "Host: example.com\r\nUser-Agent: GoClient\r\nContent-Type: text/plain\r\n\r\n"
	r := textproto.NewReader(strings.NewReader(raw))

	mimeHeader, err := r.ReadMIMEHeader()
	if err != nil {
		fmt.Println("Header okuma hatası:", err)
		return
	}

	for k, v := range mimeHeader {
		fmt.Printf("%s: %s\n", k, v)
	}
}
``

//✅ Çıktı:

``
Host: [example.com]
User-Agent: [GoClient]
Content-Type: [text/plain]
``
/*
* `ReadMIMEHeader()` HTTP ve SMTP header’larını **map\[string]\[]string** olarak döndürür.

---

# 5️⃣ Writer Örneği
*/
``go
package main

import (
	"net"
	"net/textproto"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:9090")
	defer conn.Close()

	tpConn := textproto.NewConn(conn)
	tpConn.PrintfLine("Merhaba Server!")
	response, _ := tpConn.ReadLine()
	println("Server’dan gelen:", response)
}
``
/*
---

# 6️⃣ Özet

* `net/textproto` metin tabanlı protokoller için **satır okuma/yazma + header yönetimi** sağlar.
* `Conn`, `Reader`, `Writer` yapıları sayesinde TCP üzerinden kolay protokol geliştirebilirsin.
* SMTP, POP3, IMAP, HTTP gibi protokollerde sıklıkla kullanılır.
* Header parsing (`ReadMIMEHeader`) ve satır bazlı yazma (`PrintfLine`) en temel kullanım alanlarıdır.

---

İstersen sana bunu bir adım ileri götürüp **SMTP/POP3 protokol simülasyonu** yaparak `textproto` ile **gerçek e-posta sunucu protokolü uygulaması** örneğini de gösterebilirim.

Bunu ister misin?
EVET
*/
/*
Harika! 👍 O zaman sana **Go ile `net/textproto` kullanarak basit bir SMTP sunucu simülasyonu** örneğini göstereyim. Bu örnek, gerçek e-posta sunucularının nasıl çalıştığını anlamanı sağlar.

---

## 1️⃣ Basit SMTP Sunucu Simülasyonu
*/
``go
package main

import (
	"fmt"
	"net"
	"net/textproto"
	"strings"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	tp := textproto.NewConn(conn)
	tp.PrintfLine("220 SimpleSMTP Server Hazır") // Sunucu başlangıç mesajı

	for {
		line, err := tp.ReadLine()
		if err != nil {
			fmt.Println("ReadLine hatası:", err)
			return
		}

		fmt.Println("İstemciden gelen:", line)
		cmd := strings.ToUpper(line)

		switch {
		case strings.HasPrefix(cmd, "HELO"):
			tp.PrintfLine("250 Merhaba, %s", line[5:])
		case strings.HasPrefix(cmd, "MAIL FROM:"):
			tp.PrintfLine("250 OK")
		case strings.HasPrefix(cmd, "RCPT TO:"):
			tp.PrintfLine("250 OK")
		case strings.HasPrefix(cmd, "DATA"):
			tp.PrintfLine("354 Mesajı gönder, bitirmek için . kullan")
			var msgLines []string
			for {
				dataLine, err := tp.ReadLine()
				if err != nil {
					fmt.Println("ReadLine hatası:", err)
					return
				}
				if dataLine == "." {
					break
				}
				msgLines = append(msgLines, dataLine)
			}
			fmt.Println("Alınan mesaj:\n", strings.Join(msgLines, "\n"))
			tp.PrintfLine("250 Mesaj alındı")
		case strings.HasPrefix(cmd, "QUIT"):
			tp.PrintfLine("221 Güle güle")
			return
		default:
			tp.PrintfLine("500 Komut anlaşılmadı")
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":2525")
	if err != nil {
		fmt.Println("Listen hatası:", err)
		return
	}
	defer ln.Close()
	fmt.Println("Simple SMTP Server 2525 portunda çalışıyor...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept hatası:", err)
			continue
		}
		go handleClient(conn)
	}
}
``
/*
---

## 2️⃣ Açıklamalar

1. **`textproto.Conn` kullanımı:**

   * Satır satır veri okuma/yazma (`ReadLine`, `PrintfLine`) yapılıyor.
   * SMTP protokolü metin tabanlı olduğu için `textproto` ideal.

2. **SMTP Komutları:**

   * `HELO` → Selamlaşma
   * `MAIL FROM:` → Gönderen adresi
   * `RCPT TO:` → Alıcı adresi
   * `DATA` → E-posta içeriği
   * `QUIT` → Bağlantıyı kapatma

3. **Mesaj Alma:**

   * `DATA` komutundan sonra `.` ile biten mesaj satırlarını alıyoruz.

4. **Concurrent Client:**

   * Her yeni client için `go handleClient(conn)` ile goroutine açılıyor → concurrency sağlanıyor.

---

## 3️⃣ Test Etme

Terminalden test etmek için:
*/
``bash
telnet localhost 2525
``

//Komutları sırayla deneyebilirsin:

``
HELO localhost
MAIL FROM:<gonderici@example.com>
RCPT TO:<alici@example.com>
DATA
Subject: Test
Merhaba, bu test mesajıdır.
.
QUIT
``

//Sunucuda şu çıktıyı görürsün:

``
İstemciden gelen: HELO localhost
İstemciden gelen: MAIL FROM:<gonderici@example.com>
İstemciden gelen: RCPT TO:<alici@example.com>
İstemciden gelen: DATA
Alınan mesaj:
Subject: Test
Merhaba, bu test mesajıdır.
İstemciden gelen: QUIT
``
/*
---

## 4️⃣ Özet

* `net/textproto` metin tabanlı protokoller için ideal.
* SMTP protokolünü satır bazlı olarak simüle ettik.
* Concurrency sayesinde birden fazla client aynı anda bağlanabilir.
* Bu yapı üzerinde kolayca **POP3/IMAP veya JSON-RPC protokolü** simülasyonu da yapılabilir.

---

İstersen bir sonraki adımda bunu **SMTP sunucusu + concurrent client + HTML e-posta + attachment + logging** şeklinde geliştirilmiş bir tam simülasyon örneğine çevirebilirim.

Bunu ister misin?
EVET
*/
/*
Harika! 👍 O zaman sana **Go ile gelişmiş bir SMTP sunucu simülasyonu** göstereyim. Bu örnek:

* Concurrent client destekli,
* HTML ve düz metin e-posta içeriği destekli,
* Attachment simülasyonu ve logging özellikli,
* `net/textproto` ile satır bazlı protokol işleyişini gösteren bir yapı olacak.

---

## 1️⃣ Gelişmiş SMTP Sunucu Simülasyonu
*/
``go
package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/textproto"
	"strings"
	"sync"
	"time"
)

type EmailMessage struct {
	From    string
	To      []string
	Subject string
	Text    string
	HTML    string
	Attach  []string
}

var emailLog []EmailMessage
var mu sync.Mutex

func handleClient(conn net.Conn) {
	defer conn.Close()
	tp := textproto.NewConn(conn)
	tp.PrintfLine("220 AdvancedSMTP Server Hazır")

	var currentEmail EmailMessage

	for {
		line, err := tp.ReadLine()
		if err != nil {
			fmt.Println("ReadLine hatası:", err)
			return
		}
		cmd := strings.ToUpper(line)
		fmt.Println("İstemciden gelen:", line)

		switch {
		case strings.HasPrefix(cmd, "HELO"):
			tp.PrintfLine("250 Merhaba, %s", line[5:])
		case strings.HasPrefix(cmd, "MAIL FROM:"):
			currentEmail.From = line[10:]
			tp.PrintfLine("250 OK")
		case strings.HasPrefix(cmd, "RCPT TO:"):
			currentEmail.To = append(currentEmail.To, line[8:])
			tp.PrintfLine("250 OK")
		case strings.HasPrefix(cmd, "DATA"):
			tp.PrintfLine("354 Mesajı gönder, bitirmek için . kullan")
			var msgLines []string
			for {
				dataLine, err := tp.ReadLine()
				if err != nil {
					fmt.Println("ReadLine hatası:", err)
					return
				}
				if dataLine == "." {
					break
				}
				msgLines = append(msgLines, dataLine)
			}

			// Basit HTML ve Subject ayrıştırması
			for _, l := range msgLines {
				if strings.HasPrefix(strings.ToUpper(l), "SUBJECT:") {
					currentEmail.Subject = l[8:]
				} else if strings.HasPrefix(strings.ToUpper(l), "HTML:") {
					currentEmail.HTML = l[5:]
				} else if strings.HasPrefix(strings.ToUpper(l), "TEXT:") {
					currentEmail.Text = l[5:]
				} else if strings.HasPrefix(strings.ToUpper(l), "ATTACH:") {
					currentEmail.Attach = append(currentEmail.Attach, l[7:])
					// Simülasyon: attachment içeriği okuma
					content, _ := ioutil.ReadFile(l[7:])
					fmt.Printf("Attachment %s içeriği: %s\n", l[7:], string(content))
				}
			}

			// Logla
			mu.Lock()
			emailLog = append(emailLog, currentEmail)
			mu.Unlock()

			tp.PrintfLine("250 Mesaj alındı")
			currentEmail = EmailMessage{} // Reset
		case strings.HasPrefix(cmd, "QUIT"):
			tp.PrintfLine("221 Güle güle")
			return
		default:
			tp.PrintfLine("500 Komut anlaşılmadı")
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":2525")
	if err != nil {
		fmt.Println("Listen hatası:", err)
		return
	}
	defer ln.Close()
	fmt.Println("Advanced SMTP Server 2525 portunda çalışıyor...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept hatası:", err)
			continue
		}
		go handleClient(conn)
	}
}
``
/*
---

## 2️⃣ Özellikler

1. **Concurrent Client:**

   * Her client için `go handleClient(conn)` ile goroutine açılıyor → concurrency.

2. **HTML ve Text E-posta:**

   * `HTML:` ve `TEXT:` satırları ile içerik ayrıştırılıyor.

3. **Attachment Simülasyonu:**

   * `ATTACH:` ile dosya okunup loglanıyor (gerçek e-posta gibi).

4. **Logging:**

   * Tüm e-postalar `emailLog` slice’ında tutuluyor.
   * `sync.Mutex` ile concurrency güvenliği sağlanıyor.

5. **textproto Kullanımı:**

   * `ReadLine` ve `PrintfLine` ile satır bazlı protokol işleyişi.
   * SMTP komutları simüle edildi.

---

## 3️⃣ Test Örneği

Telnet ile test edebilirsin:
*/
``bash
telnet localhost 2525
``

//Komut örneği:

``
HELO localhost
MAIL FROM:<gonderici@example.com>
RCPT TO:<alici@example.com>
DATA
SUBJECT: Test Mail
TEXT: Merhaba, bu bir test mesajıdır.
HTML: <h1>Merhaba!</h1>
ATTACH: test.txt
.
QUIT
``

//Sunucuda:

``
İstemciden gelen: HELO localhost
İstemciden gelen: MAIL FROM:<gonderici@example.com>
İstemciden gelen: RCPT TO:<alici@example.com>
İstemciden gelen: DATA
Attachment test.txt içeriği: Bu test dosyasıdır.
...
``
/*
---

## 4️⃣ Özet

* `net/textproto` ile **gerçekçi SMTP protokolü** simüle edildi.
* **HTML, Text ve Attachment** destekli mesajlar alınabiliyor.
* **Concurrent Client** desteği ile birden fazla kullanıcı aynı anda bağlanabilir.
* **Logging ve mutex** ile güvenli mesaj yönetimi sağlanıyor.

---

İstersen bunu bir adım daha ileri götürüp **concurrent client + rate limit + retry + logging + pprof** ile **tam e-posta sunucu performans testi frameworkü** haline getirebilirim.

Bunu ister misin?
*/