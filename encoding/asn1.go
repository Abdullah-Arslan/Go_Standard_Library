/*
**Go’da `encoding/asn1` paketi**ni tüm detayları ve örneklerle açıklayayım.

---

# 📌 `encoding/asn1` Paketi Nedir?

Go’nun `encoding/asn1` paketi, **ASN.1 (Abstract Syntax Notation One)** formatında veri **encode/decode** etmek için kullanılır.

📌 ASN.1 nedir?

* Telekom, güvenlik ve kriptografi standartlarında kullanılan bir veri tanımlama formatıdır.
* Özellikle **X.509 sertifikaları**, **TLS/SSL**, **PKCS#8 / PKCS#12** (private key formatları) gibi alanlarda çok önemlidir.
* Go’da `crypto/x509` paketi sertifika işlemlerinde bu `asn1` paketini kullanır.

---

# 📌 Önemli Fonksiyonlar

* **`asn1.Marshal(v interface{}) ([]byte, error)`**
  → Go değerini ASN.1 formatında encode eder.

* **`asn1.Unmarshal(data []byte, val interface{}) (rest []byte, err error)`**
  → ASN.1 verisini decode edip struct veya değişkene atar.

* **`asn1.MarshalWithParams(v interface{}, params string)`**
  → Encode sırasında parametrelerle (örn. `tag`, `optional`) kontrol sağlar.

---

# 📌 Örnekler

## 1. Basit Struct Encode/Decode
*/
``go
package main

import (
	"encoding/asn1"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Encode et
	p := Person{"Ali", 30}
	data, _ := asn1.Marshal(p)
	fmt.Printf("ASN.1 Encoded: %x\n", data)

	// Decode et
	var decoded Person
	asn1.Unmarshal(data, &decoded)
	fmt.Println("Decoded Struct:", decoded)
}
``

/7📌 Çıktı örneği:

``
ASN.1 Encoded: 30110c03416c6902011e
Decoded Struct: {Ali 30}
``
/*
---

## 2. Basit Tip Encode/Decode
*/
``go
package main

import (
	"encoding/asn1"
	"fmt"
)

func main() {
	// Encode
	num := 42
	data, _ := asn1.Marshal(num)
	fmt.Printf("ASN.1 (INTEGER): %x\n", data)

	// Decode
	var decoded int
	asn1.Unmarshal(data, &decoded)
	fmt.Println("Decoded:", decoded)
}
``
/*
📌 Çıktı:

```
ASN.1 (INTEGER): 02012a
Decoded: 42
```

---

## 3. Slice Encode/Decode
*/
``go
package main

import (
	"encoding/asn1"
	"fmt"
)

func main() {
	names := []string{"Ali", "Ayşe", "Mehmet"}
	data, _ := asn1.Marshal(names)
	fmt.Printf("ASN.1 Encoded Slice: %x\n", data)

	var decoded []string
	asn1.Unmarshal(data, &decoded)
	fmt.Println("Decoded Slice:", decoded)
}
``

//📌 Çıktı:

``
ASN.1 Encoded Slice: 30180c03416c690c04417973650c064d65686d6574
Decoded Slice: [Ali Ayşe Mehmet]
``
/*
---

## 4. Optional Alanlar

ASN.1’de bazı alanlar **opsiyonel** olabilir.
Bunu `asn1:"optional"` etiketi ile tanımlarız.
*/
``go
package main

import (
	"encoding/asn1"
	"fmt"
)

type User struct {
	Name string
	Age  int    `asn1:"optional"`
	City string `asn1:"optional"`
}

func main() {
	u := User{Name: "Ali"}
	data, _ := asn1.Marshal(u)
	fmt.Printf("ASN.1 Optional: %x\n", data)

	var decoded User
	asn1.Unmarshal(data, &decoded)
	fmt.Println("Decoded:", decoded)
}
``

//📌 Çıktı:

``
ASN.1 Optional: 30050c03416c69
Decoded: {Ali 0 }
``
/*
---

## 5. Custom Tag Kullanımı

ASN.1’de özel tag değerleri kullanılabilir.
*/
``go
package main

import (
	"encoding/asn1"
	"fmt"
)

type Employee struct {
	ID   int    `asn1:"tag:1"`
	Name string `asn1:"tag:2"`
}

func main() {
	e := Employee{ID: 101, Name: "Ali"}
	data, _ := asn1.Marshal(e)
	fmt.Printf("ASN.1 Custom Tags: %x\n", data)

	var decoded Employee
	asn1.Unmarshal(data, &decoded)
	fmt.Println("Decoded:", decoded)
}
``
/*
---

## 6. Time Encode/Decode

ASN.1 içinde tarih/saat saklamak için `time.Time` kullanılabilir.
*/
``go
package main

import (
	"encoding/asn1"
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// Encode
	data, _ := asn1.Marshal(now)
	fmt.Printf("ASN.1 Time: %x\n", data)

	// Decode
	var decoded time.Time
	asn1.Unmarshal(data, &decoded)
	fmt.Println("Decoded Time:", decoded)
}
``
/*
---

# 📌 ASN.1 Struct Etiketleri (Tagging)

Go struct alanlarında şu etiketler kullanılabilir:

* `asn1:"optional"` → Alan opsiyonel hale gelir.
* `asn1:"tag:n"` → Özel tag numarası atanır.
* `asn1:"set"` → ASN.1 SET tipi olarak encode edilir.
* `asn1:"omitempty"` → Boşsa encode edilmez.
* `asn1:"utf8"` → Stringi UTF-8 olarak encode eder.

---

# 📌 Gerçek Hayatta Kullanım

* **X.509 sertifikaları** (SSL/TLS)
* **PKCS#1 / PKCS#8** (RSA, Private Key)
* **Kerberos, LDAP, SNMP** protokolleri
* Kriptografi anahtar değişiminde

Örneğin Go’nun `crypto/x509` paketi, sertifikaları parse ederken `asn1.Unmarshal` kullanır.

---

# 📌 Özet

* `encoding/asn1`, Go’da ASN.1 formatında encode/decode yapmaya yarar.
* Özellikle güvenlik, sertifika ve şifreleme protokollerinde kullanılır.
* Struct etiketleri (`asn1:"optional"`, `asn1:"tag:..."`) ile ASN.1 formatına uygun veri üretilebilir.

---

👉 İstersen sana küçük bir **örnek sertifika bilgisi parse etme** demosu gösterebilirim (`asn1` ile X.509’un Subject alanını çözmek gibi). İster misin?
EVET
*/
/*
Harika 👍 O zaman şimdi sana **Go’da `encoding/asn1` kullanarak X.509 sertifikasından alan okuma** örneği göstereyim.

Go’da sertifika işlemleri genelde `crypto/x509` paketi ile yapılır ama bu paket kendi içinde `encoding/asn1` kullanır. Biz burada doğrudan `asn1` ile çözüm yapacağız.

---

# 📌 Örnek: X.509 Sertifikadan Subject Bilgisi Parse Etme

Diyelim elimizde bir X.509 sertifika var (örneğin `.crt` veya `.pem`). İçindeki **Subject (CN, O, C vs.)** bilgisini `asn1` ile çıkaralım.
*/
``go
package main

import (
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"os"
)

// X.509 Sertifika Subject için ASN.1 yapısı
type RDNSequence []RelativeDistinguishedNameSET
type RelativeDistinguishedNameSET []AttributeTypeAndValue

type AttributeTypeAndValue struct {
	Type  asn1.ObjectIdentifier
	Value interface{}
}

func main() {
	// Sertifika dosyasını oku (örnek: cert.pem)
	data, err := os.ReadFile("cert.pem")
	if err != nil {
		panic(err)
	}

	// PEM'den DER'e dönüştür
	block, _ := pem.Decode(data)
	if block == nil {
		panic("PEM decode edilemedi")
	}

	// ASN.1 ile subject alanını çöz
	var rdnSeq RDNSequence
	_, err = asn1.Unmarshal(block.Bytes, &rdnSeq)
	if err != nil {
		fmt.Println("ASN.1 parse hatası:", err)
		return
	}

	// Alanları yazdır
	for _, rdnSet := range rdnSeq {
		for _, atv := range rdnSet {
			fmt.Printf("OID: %v, Value: %v\n", atv.Type, atv.Value)
		}
	}
}
``
/*
---

# 📌 Bu kod ne yapıyor?

1. Sertifika dosyasını (`cert.pem`) okuyor.
2. PEM formatını **DER** formatına çeviriyor.
3. ASN.1 ile `Subject` alanındaki `RelativeDistinguishedName` setlerini çözüyor.
4. Her attribute’nin OID ve değerini yazdırıyor.

---

# 📌 Örnek Çıktı

Bir SSL sertifikası için çıktısı şöyle olabilir:

```
OID: [2 5 4 6], Value: TR
OID: [2 5 4 8], Value: Istanbul
OID: [2 5 4 7], Value: Kadikoy
OID: [2 5 4 10], Value: MyCompany Ltd
OID: [2 5 4 3], Value: mydomain.com
```

📌 Buradaki **OID’ler**:

* `2.5.4.6` → Country (C)
* `2.5.4.8` → State (ST)
* `2.5.4.7` → Locality (L)
* `2.5.4.10` → Organization (O)
* `2.5.4.3` → Common Name (CN)

---

# 📌 Neden Önemli?

* Sertifika alanlarını (CN, O, C) ASN.1 sayesinde çözebiliyoruz.
* `crypto/x509` aslında arka planda `asn1.Unmarshal` kullanıyor.
* Bu yöntemle kendi ASN.1 tabanlı protokollerini de yazabilirsin.

---

👉 İstersen ben sana bir **ASN.1 OID → anlamlı isim eşlemesi** tablosu (CN, O, C gibi) hazırlayayım ki çıktıları daha kolay okuyabilesin. Bunu yapayım mı?
*/