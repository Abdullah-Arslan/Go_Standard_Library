/*
**Go dilindeki `net` paketini** detaylı şekilde açıklayacağım ve örneklerle kullanımını göstereceğim.

---

## 📦 `net` Paketi Nedir?

`net` paketi, Go dilinde **ağ (network) programlaması** yapmak için kullanılan temel pakettir.

* TCP, UDP gibi bağlantılar kurmayı sağlar.
* IP adresleri, hostname çözümleme, bağlantı açma ve dinleme gibi işlemleri içerir.
* HTTP veya daha üst seviye protokoller `net/http` paketini kullanır ama altta hep `net` paketi çalışır.

Dokümantasyon: [pkg.go.dev/net](https://pkg.go.dev/net)

---

## 🔧 Paket İçindeki Temel Fonksiyon ve Yapılar

### 1️⃣ `Dial(network, address string) (Conn, error)`

* Belirtilen network ve adrese bağlanır (`tcp`, `udp`, `ip`).
* Örnek: TCP ile bağlanma
*/
``go
package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Bağlantı başarılı:", conn.RemoteAddr())
}
``
/*
---

### 2️⃣ `Listen(network, address string) (Listener, error)`

* Belirtilen network ve adreste gelen bağlantıları dinler.
* Örnek: TCP sunucu
*/
``go
package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	defer ln.Close()
	fmt.Println("Sunucu dinleniyor :8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Hata:", err)
			continue
		}
		fmt.Println("Yeni bağlantı:", conn.RemoteAddr())
		conn.Close()
	}
}
``
/*
---

### 3️⃣ `ResolveIPAddr(network, address string) (*IPAddr, error)`

* IP adresi çözümler (`tcp4`, `tcp6`, `ip`).
* Örnek:
*/
``go
package main

import (
	"fmt"
	"net"
)

func main() {
	ipAddr, err := net.ResolveIPAddr("ip", "google.com")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Println("Google IP:", ipAddr)
}
``
/*
---

### 4️⃣ `LookupHost(host string) ([]string, error)`

* Bir hostun IP adreslerini döndürür.
*/
``go
package main

import (
	"fmt"
	"net"
)

func main() {
	ips, err := net.LookupHost("google.com")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Println("Google IP adresleri:", ips)
}
``
/*
---

### 5️⃣ `IP` ve `IPNet` Yapıları

* `IP`: IP adresini temsil eder (`net.IP`).
* `IPNet`: IP ağı ve maskesi (`net.IPNet`).

**Örnek:**
*/
``go
package main

import (
	"fmt"
	"net"
)

func main() {
	ip := net.ParseIP("192.168.1.1")
	fmt.Println("IP:", ip)

	_, ipNet, err := net.ParseCIDR("192.168.1.0/24")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Println("IP Network:", ipNet)
}
``
/*
---

### 6️⃣ UDP Bağlantısı
*/
``go
package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:8081")
	conn, _ := net.DialUDP("udp", nil, addr)
	defer conn.Close()
	conn.Write([]byte("Merhaba UDP"))
	fmt.Println("UDP mesajı gönderildi")
}
``
/*
---

### 🔹 Özet Fonksiyon ve Yapılar
*/

| Fonksiyon / Yapı      | Açıklama                                         |
| --------------------- | ------------------------------------------------ |
| `Dial`                | TCP/UDP/IP bağlantısı açar                       |
| `Listen`              | Belirtilen port ve networkte dinleme yapar       |
| `Accept`              | Gelen bağlantıyı kabul eder                      |
| `ResolveIPAddr`       | Hostname’den IP çözer                            |
| `LookupHost`          | Hostname’in IP adreslerini döndürür              |
| `ParseIP`             | String’den IP nesnesi oluşturur                  |
| `ParseCIDR`           | CIDR formatından IP ve network bilgisi oluşturur |
| `Conn`                | Açılan bağlantıyı temsil eder                    |
| `Listener`            | Dinlenen portu temsil eder                       |
| `UDPAddr` / `DialUDP` | UDP bağlantısı için kullanılır                   |
/*
---

İstersen ben şimdi bu **tüm `net` paketi örneklerini tek bir çalıştırılabilir Go dosyası** hâline getirip eksiksiz bir set hazırlayabilirim.

Bunu yapmamı ister misin?
EVET
*/
//Golang Net Examples
package main

import (
	"fmt"
	"net"
)

func main() {
	// 1️⃣ TCP Dial Örneği
	fmt.Println("--- TCP Dial Örneği ---")
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		fmt.Println("Hata:", err)
	} else {
		defer conn.Close()
		fmt.Println("Bağlantı başarılı:", conn.RemoteAddr())
	}

	// 2️⃣ TCP Sunucu (Listen) Örneği
	fmt.Println("--- TCP Sunucu (Listen) Örneği ---")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Hata:", err)
	} else {
		defer ln.Close()
		fmt.Println("Sunucu dinleniyor :8080")

		// Sadece tek bir bağlantı kabul etme örneği
		go func() {
			conn, err := ln.Accept()
			if err != nil {
				fmt.Println("Hata:", err)
				return
			}
			fmt.Println("Yeni bağlantı:", conn.RemoteAddr())
			conn.Close()
		}()
	}

	// 3️⃣ IP Adres Çözümleme (ResolveIPAddr)
	fmt.Println("--- IP Adres Çözümleme ---")
	ipAddr, err := net.ResolveIPAddr("ip", "google.com")
	if err != nil {
		fmt.Println("Hata:", err)
	} else {
		fmt.Println("Google IP:", ipAddr)
	}

	// 4️⃣ Host Lookup
	fmt.Println("--- Host Lookup ---")
	ips, err := net.LookupHost("google.com")
	if err != nil {
		fmt.Println("Hata:", err)
	} else {
		fmt.Println("Google IP adresleri:", ips)
	}

	// 5️⃣ IP ve IPNet
	fmt.Println("--- IP ve IPNet ---")
	ip := net.ParseIP("192.168.1.1")
	fmt.Println("IP:", ip)

	_, ipNet, err := net.ParseCIDR("192.168.1.0/24")
	if err != nil {
		fmt.Println("Hata:", err)
	} else {
		fmt.Println("IP Network:", ipNet)
	}

	// 6️⃣ UDP Bağlantısı
	fmt.Println("--- UDP Bağlantısı ---")
	addr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:8081")
	connUDP, _ := net.DialUDP("udp", nil, addr)
	defer connUDP.Close()
	connUDP.Write([]byte("Merhaba UDP"))
	fmt.Println("UDP mesajı gönderildi")
}
/*
Tamamdır! Tüm `net` paketi örneklerini **çalıştırılabilir tek bir Go dosyası** olarak hazırladım. Bu dosya ile:

* TCP bağlantısı açma (`Dial`)
* TCP sunucu oluşturma (`Listen`)
* IP çözümleme ve host lookup
* IP ve CIDR işlemleri
* UDP bağlantısı

gibi temel ağ işlemlerini tek seferde deneyebilirsin.
*/