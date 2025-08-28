/*
**Go’daki `encoding` paketi**ni en ince ayrıntısına kadar açıklayayım.

---

# 📌 `encoding` Paketi Nedir?

Go dilinde `encoding` **bir üst (parent) pakettir**.
Yani kendi başına kullanılabilen bir işlevi yoktur ama altında pek çok alt paket barındırır.

Bu alt paketler, **farklı veri formatlarını kodlamak (encode) ve çözmek (decode)** için kullanılır.

---

# 📌 `encoding` Alt Paketleri

`encoding` dizini altında bulunan önemli paketler şunlardır:

| Paket              | Açıklama                                                                   |
| ------------------ | -------------------------------------------------------------------------- |
| `encoding/ascii85` | ASCII85 (Base85) kodlaması                                                 |
| `encoding/asn1`    | ASN.1 (Abstract Syntax Notation One) kodlaması (sertifikalarda kullanılır) |
| `encoding/base32`  | Base32 kodlama/çözme                                                       |
| `encoding/base64`  | Base64 kodlama/çözme                                                       |
| `encoding/binary`  | Sayıları (int, float vb.) ikili (binary) formatta okuma/yazma              |
| `encoding/csv`     | CSV (Comma-Separated Values) dosya formatı okuma/yazma                     |
| `encoding/gob`     | Go’ya özel binary serileştirme formatı                                     |
| `encoding/hex`     | Hexadecimal (16’lık sayı sistemi) kodlama                                  |
| `encoding/json`    | JSON verilerini encode/decode etme                                         |
| `encoding/pem`     | PEM formatı (sertifikalarda, anahtar dosyalarında kullanılır)              |
| `encoding/xml`     | XML verilerini encode/decode etme                                          |

---

# 📌 Şimdi Tek Tek Örneklerle Açıklayalım

## 1. `encoding/ascii85`

**Base85 kodlaması**. Daha az yer kaplar.
*/
``go
package main

import (
	"encoding/ascii85"
	"fmt"
)

func main() {
	data := []byte("Merhaba Dünya")
	encoded := make([]byte, ascii85.MaxEncodedLen(len(data)))
	n := ascii85.Encode(encoded, data)
	fmt.Println("Encoded:", string(encoded[:n]))

	decoded := make([]byte, len(data))
	nd, _, _ := ascii85.Decode(decoded, encoded[:n], true)
	fmt.Println("Decoded:", string(decoded[:nd]))
}
``
/*
---

## 2. `encoding/asn1`

Genellikle **X.509 sertifikaları** için kullanılır.
*/
``go
package main

import (
	"encoding/asn1"
	"fmt"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}

	p := Person{"Ali", 30}
	data, _ := asn1.Marshal(p)
	fmt.Println("ASN.1 Encoded:", data)

	var decoded Person
	asn1.Unmarshal(data, &decoded)
	fmt.Println("Decoded Struct:", decoded)
}
``
/*
---

## 3. `encoding/base32`
*/
``go
package main

import (
	"encoding/base32"
	"fmt"
)

func main() {
	msg := "Merhaba"
	encoded := base32.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println("Base32 Encoded:", encoded)

	decoded, _ := base32.StdEncoding.DecodeString(encoded)
	fmt.Println("Decoded:", string(decoded))
}
``
/*
---

## 4. `encoding/base64`
*/
``go
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	msg := "Merhaba"
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println("Base64 Encoded:", encoded)

	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Println("Decoded:", string(decoded))
}
``
/*
---

## 5. `encoding/binary`

Sayısal değerleri **little-endian / big-endian** olarak yazıp okumak için.
*/
``go
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	var buf bytes.Buffer
	num := uint32(123456)

	// Yaz
	binary.Write(&buf, binary.LittleEndian, num)
	fmt.Println("Binary data:", buf.Bytes())

	// Oku
	var result uint32
	binary.Read(&buf, binary.LittleEndian, &result)
	fmt.Println("Decoded number:", result)
}
``
/*
---

## 6. `encoding/csv`
*/
``go
package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func main() {
	csvData := "name,age\nAli,30\nAyşe,25"
	r := csv.NewReader(strings.NewReader(csvData))
	records, _ := r.ReadAll()
	fmt.Println("CSV:", records)
}
``
/*
---

## 7. `encoding/gob`

Go’ya özel **struct serileştirme**.
*/
``go
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	var buf bytes.Buffer

	// Encode
	enc := gob.NewEncoder(&buf)
	enc.Encode(Person{"Ali", 30})

	// Decode
	var p Person
	dec := gob.NewDecoder(&buf)
	dec.Decode(&p)

	fmt.Println("Decoded:", p)
}
``
/*
---

## 8. `encoding/hex`
*/
``go
package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	data := []byte("Merhaba")
	encoded := hex.EncodeToString(data)
	fmt.Println("Hex:", encoded)

	decoded, _ := hex.DecodeString(encoded)
	fmt.Println("Decoded:", string(decoded))
}
``
/*
---

## 9. `encoding/json`

En çok kullanılan paketlerden biri 🚀
*/
``go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	u := User{"Ali", 30}
	data, _ := json.Marshal(u)
	fmt.Println("JSON:", string(data))

	var decoded User
	json.Unmarshal(data, &decoded)
	fmt.Println("Decoded:", decoded)
}
``
/*
---

## 10. `encoding/pem`

PEM formatı (sertifikalar, private keyler için).
*/
``go
package main

import (
	"encoding/pem"
	"fmt"
)

func main() {
	block := &pem.Block{
		Type:  "MESSAGE",
		Bytes: []byte("Merhaba Dünya"),
	}

	pemData := pem.EncodeToMemory(block)
	fmt.Println(string(pemData))

	decoded, _ := pem.Decode(pemData)
	fmt.Println("Decoded:", string(decoded.Bytes))
}
``
/*
---

## 11. `encoding/xml`
*/

``go
package main

import (
	"encoding/xml"
	"fmt"
)

type User struct {
	XMLName xml.Name `xml:"user"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
}

func main() {
	u := User{Name: "Ali", Age: 30}
	data, _ := xml.MarshalIndent(u, "", "  ")
	fmt.Println(string(data))

	var decoded User
	xml.Unmarshal(data, &decoded)
	fmt.Println("Decoded:", decoded)
}
``
/*
---

# 📌 Özet

Go’daki `encoding` paketi, **veri formatlarını encode/decode etmek için alt paketler sunar**:

* Metin tabanlı: `json`, `xml`, `csv`, `asn1`, `pem`
* İkili/özel: `base32`, `base64`, `hex`, `ascii85`, `binary`, `gob`

Bunları kullanarak dosya, network veya API veri alışverişlerinde farklı formatları kolayca işleyebilirsin.

---

👉 İstersen ben sana **bunların hepsini tek bir örnek proje içinde birleştireyim** (örneğin bir API hem JSON hem XML hem de CSV üretsin). Böyle bir **mini demo projesi** ister misin?
EVET
*/

/*
Süper 👍 O zaman sana **tek bir Go projesi** hazırlayayım:

Bu proje bir **HTTP API** olacak ve aynı `User` verisini:

* `/json` → JSON olarak,
* `/xml` → XML olarak,
* `/csv` → CSV olarak,
* `/base64` → Base64 encode edilmiş olarak,
* `/hex` → Hex encode edilmiş olarak döndürecek.

Böylece `encoding` paketlerinin çoğunu **tek uygulamada** görmüş olacağız 🚀

---

# 📌 Proje Yapısı

```
myapp/
 └─ main.go
```

---

# 📌 Kod: `main.go`
*/
``go
package main

import (
	"encoding/base64"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
)

// Ortak veri
type User struct {
	Name string `json:"name" xml:"name"`
	Age  int    `json:"age"  xml:"age"`
}

// Global veri
var users = []User{
	{"Ali", 30},
	{"Ayşe", 25},
	{"Mehmet", 40},
}

func main() {
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/xml", xmlHandler)
	http.HandleFunc("/csv", csvHandler)
	http.HandleFunc("/base64", base64Handler)
	http.HandleFunc("/hex", hexHandler)

	fmt.Println("🚀 Sunucu çalışıyor: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// JSON handler
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := json.MarshalIndent(users, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// XML handler
func xmlHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := xml.MarshalIndent(users, "", "  ")
	w.Header().Set("Content-Type", "application/xml")
	w.Write(data)
}

// CSV handler
func csvHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/csv")
	writer := csv.NewWriter(w)
	writer.Write([]string{"Name", "Age"})
	for _, u := range users {
		writer.Write([]string{u.Name, fmt.Sprintf("%d", u.Age)})
	}
	writer.Flush()
}

// Base64 handler
func base64Handler(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(users) // JSON'u Base64'e çeviriyoruz
	encoded := base64.StdEncoding.EncodeToString(data)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(encoded))
}

// Hex handler
func hexHandler(w http.ResponseWriter, r *http.Request) {
	data := []byte("Merhaba Dünya") // örnek sabit veri
	encoded := hex.EncodeToString(data)
	decoded, _ := hex.DecodeString(encoded)

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Encoded HEX: %s\nDecoded: %s", encoded, string(decoded))
}
``
/*
---

# 📌 Çalıştırma
*/
``bash
go run main.go
``
/*
Tarayıcıda veya `curl` ile test et:

* [http://localhost:8080/json](http://localhost:8080/json) → JSON formatı
* [http://localhost:8080/xml](http://localhost:8080/xml) → XML formatı
* [http://localhost:8080/csv](http://localhost:8080/csv) → CSV formatı
* [http://localhost:8080/base64](http://localhost:8080/base64) → Base64 encoded JSON
* [http://localhost:8080/hex](http://localhost:8080/hex) → Hex encode/decode örneği

---

📌 Böylece tek proje içinde `encoding/json`, `encoding/xml`, `encoding/csv`, `encoding/base64`, `encoding/hex` kullandık ✅

---

👉 İstersen bu örneğe bir de **`gob` ile struct encode/decode** ekleyeyim (örneğin `POST /gob` ile binary formatta gönderip çözsün). Onu da ister misin?
*/