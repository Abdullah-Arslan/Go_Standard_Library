/*
Go 1.18 ile gelen **`net/netip`** paketini detaylıca açıklayayım.

`net/netip` paketi, **IP adreslerini ve IP ağlarını temsil etmek, işlemek ve karşılaştırmak** için geliştirilmiş modern bir pakettir. Eski `net` paketine kıyasla **daha hızlı ve daha güvenli**dır çünkü immutable tipler ve değer semantiği kullanır.

---

# 1️⃣ Paket Amaçları

* IP adreslerini temsil etmek (`IPv4` ve `IPv6`)
* IP ağlarını temsil etmek (`CIDR` formatları)
* IP adreslerini karşılaştırmak ve manipüle etmek
* Daha verimli ve tip güvenli adres işleme

---

# 2️⃣ Temel Tipler
*/

| Tip                | Açıklama                                                   |
| ------------------ | ---------------------------------------------------------- |
| `netip.Addr`       | Bir IP adresini temsil eder. Immutable’dır.                |
| `netip.Prefix`     | Bir IP adresi bloğunu temsil eder (örn: `192.168.1.0/24`). |
| `netip.AddrPort`   | IP + Port ikilisi.                                         |
| `netip.IPPortFrom` | IP ve portu birleştirip AddrPort oluşturur.                |

/*
---

# 3️⃣ `netip.Addr` Kullanımı

### IP adresi oluşturma
*/
``go
package main

import (
	"fmt"
	"net/netip"
)

func main() {
	ip, err := netip.ParseAddr("192.168.1.100")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	fmt.Println("IP:", ip)
	fmt.Println("IsValid:", ip.IsValid())   // Geçerli mi?
	fmt.Println("Is4:", ip.Is4())           // IPv4 mü?
	fmt.Println("Is6:", ip.Is6())           // IPv6 mı?
	fmt.Println("Unmap:", ip.Unmap())       // IPv4-mapped IPv6 varsa dönüştürür
}
``
/*
---

# 4️⃣ `netip.Prefix` Kullanımı

### CIDR (IP ağı) oluşturma
*/
``go
prefix, err := netip.ParsePrefix("192.168.1.0/24")
if err != nil {
    fmt.Println("Hata:", err)
    return
}

fmt.Println("Prefix:", prefix)
fmt.Println("Network:", prefix.Masked())       // Ağ adresi
fmt.Println("Bits:", prefix.Bits())            // Prefix uzunluğu
fmt.Println("Contains 192.168.1.5?", prefix.Contains(netip.MustParseAddr("192.168.1.5")))
``
/*
* `Contains` → IP’nin bu ağ bloğuna ait olup olmadığını kontrol eder.

---

# 5️⃣ `netip.AddrPort` Kullanımı
*/
``go
ip := netip.MustParseAddr("192.168.1.10")
addrPort := netip.AddrPortFrom(ip, 8080)

fmt.Println("AddrPort:", addrPort)
fmt.Println("IP:", addrPort.Addr())
fmt.Println("Port:", addrPort.Port())
``
/*
* `AddrPort` tipinde IP + port ikilisi güvenli ve immutable şekilde tutulur.

---

# 6️⃣ IP Karşılaştırma ve Fonksiyonlar
*/
``go
ip1 := netip.MustParseAddr("192.168.1.1")
ip2 := netip.MustParseAddr("192.168.1.2")

fmt.Println(ip1 == ip2)                // false
fmt.Println(ip1.Compare(ip2))          // -1 : ip1 < ip2, 0 : eşit, 1 : ip1 > ip2
fmt.Println(ip1.Next())                 // Sonraki IP adresini döndürür
fmt.Println(ip2.Prev())                 // Önceki IP adresi
``
/*
---

# 7️⃣ Tüm IP’yi Listeleme Örneği

CIDR içerisindeki tüm IP adreslerini döngü ile gezebiliriz:
*/
``go
prefix := netip.MustParsePrefix("192.168.1.0/30")

for ip := prefix.Masked(); prefix.Contains(ip); ip = ip.Next() {
    fmt.Println(ip)
}
``

//Çıktı:

``
192.168.1.0
192.168.1.1
192.168.1.2
192.168.1.3
``

---

//# 8️⃣ Özet Tablo

| Fonksiyon / Tip                       | Açıklama                                       |
| ------------------------------------- | ---------------------------------------------- |
| `netip.ParseAddr`                     | String’den IP adresi oluşturur                 |
| `netip.MustParseAddr`                 | Hata döndürmeden IP oluşturur (panic on error) |
| `netip.Addr.IsValid`                  | IP geçerli mi?                                 |
| `netip.Addr.Is4` / `Is6`              | IPv4/IPv6 kontrolü                             |
| `netip.ParsePrefix`                   | CIDR’den Prefix oluşturur                      |
| `netip.Prefix.Contains`               | IP bu ağda mı?                                 |
| `netip.AddrPortFrom`                  | IP + Port oluşturur                            |
| `AddrPort.Addr()` / `AddrPort.Port()` | IP ve port bilgisini alır                      |
| `Addr.Compare`                        | IP adreslerini karşılaştırır                   |
| `Addr.Next()` / `Addr.Prev()`         | Sıradaki / önceki IP                           |
| `Prefix.Masked()`                     | Ağ adresini verir                              |

/*
---

✅ **Kısaca:**
`net/netip`, Go’da **IP adreslerini ve ağlarını güvenli, hızlı ve tip güvenli şekilde** yönetmek için modern pakettir.
Eski `net.IP` ve `net.IPNet` tiplerinin yerine geçmesi önerilir.

---

İstersen sana bunun üzerine **IPv4 ve IPv6 karışık bir listeyi `netip` ile filtreleme ve sıralama örneği** de gösterebilirim. Bunu ister misin?
EVET
*/
/*
Harika 👍 o zaman sana **IPv4 ve IPv6 karışık bir IP listesini `net/netip` ile filtreleme ve sıralama** örneğini göstereyim.

---

## 1️⃣ Örnek Kod
*/
``go
package main

import (
	"fmt"
	"net/netip"
	"sort"
)

func main() {
	// Karışık IP listesi (IPv4 ve IPv6)
	ipStrings := []string{
		"192.168.1.10",
		"2001:db8::1",
		"10.0.0.5",
		"fe80::1",
		"172.16.0.1",
		"8.8.8.8",
	}

	var ips []netip.Addr

	// Stringleri netip.Addr tipine dönüştür
	for _, s := range ipStrings {
		ip, err := netip.ParseAddr(s)
		if err == nil {
			ips = append(ips, ip)
		}
	}

	// IPv4 ve IPv6 ayrı listelere ayır
	var ipv4List, ipv6List []netip.Addr
	for _, ip := range ips {
		if ip.Is4() {
			ipv4List = append(ipv4List, ip)
		} else if ip.Is6() {
			ipv6List = append(ipv6List, ip)
		}
	}

	// Her listeyi sırala (Compare ile)
	sort.Slice(ipv4List, func(i, j int) bool {
		return ipv4List[i].Compare(ipv4List[j]) < 0
	})

	sort.Slice(ipv6List, func(i, j int) bool {
		return ipv6List[i].Compare(ipv6List[j]) < 0
	})

	// Sonuçları yazdır
	fmt.Println("IPv4 Listesi:")
	for _, ip := range ipv4List {
		fmt.Println(ip)
	}

	fmt.Println("\nIPv6 Listesi:")
	for _, ip := range ipv6List {
		fmt.Println(ip)
	}
}
``
/*
---

## 2️⃣ Açıklamalar

1. **`netip.ParseAddr`** ile string IP’leri tip güvenli `netip.Addr`’a dönüştürdük.
2. **`ip.Is4()` ve `ip.Is6()`** ile IPv4 ve IPv6 adreslerini ayırdık.
3. **`Compare` fonksiyonu** ile IP’leri sıraladık.

   * IPv4: küçük → büyük
   * IPv6: küçük → büyük
4. Sonuçta **IPv4 ve IPv6 listeleri ayrı ve sıralı** hale geldi.

---

## 3️⃣ Örnek Çıktı
*/
``
IPv4 Listesi:
8.8.8.8
10.0.0.5
172.16.0.1
192.168.1.10

IPv6 Listesi:
2001:db8::1
fe80::1
``
/*
---

✅ Bu yöntemle karışık IP listelerini kolayca **ayırabilir, sıralayabilir ve filtreleyebilirsin**.
`net/netip` paketi sayesinde eski `net.IP`’e göre hem **daha hızlı** hem **daha güvenli** çalışıyoruz.

---

İstersen bir sonraki adım olarak sana **CIDR filtreleme ve IP aralıklarını kontrol etme örneğini** de gösterebilirim. Bunu da ister misin?
EVET
*/
/*
Harika 👍 o zaman sana **CIDR filtreleme ve IP aralıklarını kontrol etme** örneğini `net/netip` ile göstereyim.

---

## 1️⃣ Örnek Kod: IP’leri CIDR ile Filtreleme
*/
``go
package main

import (
	"fmt"
	"net/netip"
)

func main() {
	// IP listesi
	ipStrings := []string{
		"192.168.1.10",
		"192.168.1.20",
		"10.0.0.5",
		"172.16.0.1",
		"8.8.8.8",
		"2001:db8::1",
	}

	// CIDR prefix
	cidr := "192.168.1.0/24"
	prefix, err := netip.ParsePrefix(cidr)
	if err != nil {
		fmt.Println("CIDR hatalı:", err)
		return
	}

	fmt.Println("CIDR filtresi:", cidr)
	fmt.Println("Bu aralıkta olan IP’ler:")

	// IP’leri filtrele
	for _, s := range ipStrings {
		ip, err := netip.ParseAddr(s)
		if err != nil {
			continue
		}

		if prefix.Contains(ip) {
			fmt.Println(ip)
		}
	}
}
``
/*
---

## 2️⃣ Açıklamalar

1. **`netip.ParsePrefix`** ile CIDR’yi `netip.Prefix` tipine dönüştürdük.
2. **`Prefix.Contains(ip)`** ile IP’nin bu CIDR aralığında olup olmadığını kontrol ettik.
3. Bu yöntemle hem IPv4 hem IPv6 CIDR aralıklarını kolayca kontrol edebilirsin.

---

## 3️⃣ Örnek Çıktı
*/
``
CIDR filtresi: 192.168.1.0/24
Bu aralıkta olan IP’ler:
192.168.1.10
192.168.1.20
``
/*
---

✅ Avantajları:

* IPv4 ve IPv6 destekli.
* `net/netip` ile hızlı ve tip güvenli.
* Büyük IP listelerini veya ağ bloklarını filtrelemek için ideal.
* `Compare`, `Next`, `Prev` ile aralık üzerinde daha gelişmiş işlemler yapılabilir.

---

İstersen bir sonraki adım olarak sana **IP aralıklarını birleştirme ve boşta kalan IP’leri tespit etme** örneğini gösterebilirim. Bunu da ister misin?
Harika 👍 o zaman sana **IP aralıklarını birleştirme ve boşta kalan IP’leri tespit etme** örneğini `net/netip` ile göstereyim.

---

## 1️⃣ Örnek Kod: IP Aralıklarını Birleştirme ve Boş IP’leri Bulma
*/
``go
package main

import (
	"fmt"
	"net/netip"
)

func main() {
	// CIDR listesi
	cidrs := []string{
		"192.168.1.0/26",  // 192.168.1.0 - 192.168.1.63
		"192.168.1.64/27", // 192.168.1.64 - 192.168.1.95
		"192.168.1.100/30",// 192.168.1.100 - 192.168.1.103
	}

	// Parse edilmiş prefix listesi
	var prefixes []netip.Prefix
	for _, c := range cidrs {
		p, err := netip.ParsePrefix(c)
		if err == nil {
			prefixes = append(prefixes, p)
		}
	}

	// Aralıkları sırala
	for i := 0; i < len(prefixes)-1; i++ {
		for j := i + 1; j < len(prefixes); j++ {
			if prefixes[i].Masked().Compare(prefixes[j].Masked()) > 0 {
				prefixes[i], prefixes[j] = prefixes[j], prefixes[i]
			}
		}
	}

	// Boş IP’leri bul
	fmt.Println("Boş IP aralıkları:")
	for i := 0; i < len(prefixes)-1; i++ {
		endIP := prefixes[i].Masked().Next()
		for endIP.Compare(prefixes[i].Broadcast()) <= 0 {
			endIP = endIP.Next()
		}

		startNext := prefixes[i+1].Masked()
		for ip := endIP; ip.Compare(startNext) < 0; ip = ip.Next() {
			fmt.Println(ip)
		}
	}
}
``
/*
---

## 2️⃣ Açıklamalar

1. **`ParsePrefix`** ile CIDR’leri parse ettik.
2. Prefix’leri **başlangıç IP’sine göre sıraladık**.
3. Her aralıktan sonra **bir sonraki aralığın başlangıcına kadar boş IP’leri** hesapladık.
4. `Next()` ve `Broadcast()` fonksiyonlarıyla IP aralıklarını gezdik.

---

## 3️⃣ Örnek Çıktı (IP’leri tek tek listelemek yerine mantığı gösterir)
*/
``
Boş IP aralıkları:
192.168.1.96
192.168.1.97
192.168.1.98
192.168.1.99
``
/*
* `192.168.1.96-99` arası CIDR’ler tarafından kapsanmıyor, yani boş IP’ler.

---

✅ **Avantajlar:**

* Ağ yönetimi ve IP planlamasında çok kullanışlı.
* IPv4 ve IPv6 destekli.
* `net/netip` immutable ve tip güvenli olduğu için büyük IP bloklarında güvenle kullanılabilir.

---

İstersen bir sonraki adım olarak sana **IP adreslerini CIDR’lere otomatik bölme (subnetting) ve boş IP’leri optimize etme** örneğini de gösterebilirim. Bunu ister misin?
*/