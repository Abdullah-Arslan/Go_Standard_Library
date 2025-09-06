/*
Go’nun standart kütüphanesinde bulunan **`encoding/xml`** paketi, **XML (Extensible Markup Language)** verilerini **okumak (parse etmek)** ve **yazmak (encode etmek)** için kullanılır.

JSON için nasıl `encoding/json` varsa, XML için de **`encoding/xml`** vardır.

---

# 📌 Genel Özellikler

* **Struct ↔ XML dönüşümü** (`Marshal`, `Unmarshal`)
* **Tag desteği** (`xml:"element"`, `xml:"attr"` vs.)
* **Decoder/Encoder** → akış (stream) bazlı işlem
* **Token bazlı işleme** → düşük seviyede XML parçalamak mümkün

---

# 📌 Temel Fonksiyonlar

1. **`xml.Marshal(v any) ([]byte, error)`**
   → Go struct’ını XML’e çevirir.

2. **`xml.Unmarshal(data []byte, v any) error`**
   → XML verisini Go struct’ına çevirir.

3. **`xml.NewDecoder(r io.Reader)`**
   → Akış bazlı XML okuma.

4. **`xml.NewEncoder(w io.Writer)`**
   → XML yazma.

5. **`xml.Token` arayüzü**
   → Düşük seviyede XML elementlerini tek tek işler.

---

# 📌 Örnekler

---

## 1. Struct → XML (Marshal)
*/
``go
package main

import (
	"encoding/xml"
	"fmt"
)

type User struct {
	XMLName xml.Name `xml:"user"`
	ID      int      `xml:"id"`
	Name    string   `xml:"name"`
	Email   string   `xml:"email"`
}

func main() {
	u := User{ID: 1, Name: "Ali", Email: "ali@example.com"}

	data, err := xml.Marshal(u)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
``
/*
**Çıktı:**

```xml
<user>
  <id>1</id>
  <name>Ali</name>
  <email>ali@example.com</email>
</user>
```

---

## 2. XML → Struct (Unmarshal)
*/
``go
package main

import (
	"encoding/xml"
	"fmt"
)

type Product struct {
	XMLName xml.Name `xml:"product"`
	Name    string   `xml:"name"`
	Price   float64  `xml:"price"`
}

func main() {
	data := []byte(`
	<product>
		<name>Laptop</name>
		<price>14999.90</price>
	</product>`)

	var p Product
	err := xml.Unmarshal(data, &p)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Ürün: %s, Fiyat: %.2f\n", p.Name, p.Price)
}
``
/*
**Çıktı:**

```
Ürün: Laptop, Fiyat: 14999.90
```

---

## 3. XML Attribute Kullanımı
*/
``go
package main

import (
	"encoding/xml"
	"fmt"
)

type Book struct {
	XMLName xml.Name `xml:"book"`
	Title   string   `xml:"title,attr"`
	Author  string   `xml:"author"`
}

func main() {
	data := []byte(`<book title="Go Programming"><author>John Doe</author></book>`)

	var b Book
	if err := xml.Unmarshal(data, &b); err != nil {
		panic(err)
	}

	fmt.Printf("Kitap: %s, Yazar: %s\n", b.Title, b.Author)
}
``
/*
**Çıktı:**

```
Kitap: Go Programming, Yazar: John Doe
```

---

## 4. Liste (Slice) ile Çalışma
*/
``go
package main

import (
	"encoding/xml"
	"fmt"
)

type Library struct {
	XMLName xml.Name `xml:"library"`
	Books   []Book   `xml:"book"`
}

type Book struct {
	Title string `xml:"title"`
}

func main() {
	data := []byte(`
	<library>
		<book><title>Go</title></book>
		<book><title>Rust</title></book>
	</library>`)

	var lib Library
	_ = xml.Unmarshal(data, &lib)

	for _, b := range lib.Books {
		fmt.Println("Kitap:", b.Title)
	}
}
``
/*
**Çıktı:**

```
Kitap: Go
Kitap: Rust
```

---

## 5. Akış (Stream) Bazlı XML Okuma (`Decoder`)

Büyük XML dosyalarında tek seferde `Unmarshal` etmek yerine, satır satır (token bazlı) okumak daha verimlidir.
*/
``go
package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

func main() {
	input := `
	<library>
		<book><title>Go</title></book>
		<book><title>Rust</title></book>
	</library>`

	decoder := xml.NewDecoder(strings.NewReader(input))

	for {
		tok, err := decoder.Token()
		if err != nil {
			break
		}
		switch t := tok.(type) {
		case xml.StartElement:
			fmt.Println("Başlangıç:", t.Name.Local)
		case xml.EndElement:
			fmt.Println("Bitiş:", t.Name.Local)
		case xml.CharData:
			text := strings.TrimSpace(string(t))
			if text != "" {
				fmt.Println("Metin:", text)
			}
		}
	}
}
``
/*
**Çıktı:**

```
Başlangıç: library
Başlangıç: book
Başlangıç: title
Metin: Go
Bitiş: title
Bitiş: book
Başlangıç: book
Başlangıç: title
Metin: Rust
Bitiş: title
Bitiş: book
Bitiş: library
```

---

## 6. XML Yazma (`Encoder`)
*7
``go
package main

import (
	"encoding/xml"
	"os"
)

type Person struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

func main() {
	p := Person{"Ahmet", 30}

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ") // formatlı yaz
	_ = enc.Encode(p)
}
``

/7**Çıktı:**

``xml
<Person>
  <name>Ahmet</name>
  <age>30</age>
</Person>
``
/*
---

# 📌 Özet

* **`xml.Marshal`** → Struct → XML
* **`xml.Unmarshal`** → XML → Struct
* **`xml.Decoder`** → Akış (stream) okuma
* **`xml.Encoder`** → XML yazma
* **Tag desteği** → `xml:"field"`, `xml:"field,attr"`
* **Token bazlı işleme** → Düşük seviye XML işlemek için

---

👉 İstersen sana **`encoding/xml` ile RSS Feed okuma ve parse etme** (gerçek dünya örneği) gösterebilirim. Görmek ister misin?
EVET
*/
/*
Harika 👍 o zaman sana **`encoding/xml` ile bir RSS feed okuma ve parse etme** örneği göstereyim.

RSS (Really Simple Syndication) aslında **XML tabanlı bir haber akışı**dır. Haber siteleri, bloglar veya podcast servisleri genelde RSS üzerinden içerik paylaşır.

---

# 📌 Örnek: RSS Feed Parse Etme

## 1. RSS Yapısı (Örnek)

Bir RSS feed dosyası genelde şöyle görünür:
*/
``xml
<rss version="2.0">
  <channel>
    <title>Go Haberler</title>
    <link>https://golang.org</link>
    <description>Go diliyle ilgili güncel haberler</description>
    <item>
      <title>Go 1.23 çıktı!</title>
      <link>https://golang.org/doc/go1.23</link>
      <description>Yeni özellikler...</description>
    </item>
    <item>
      <title>GoCon 2025</title>
      <link>https://gocon.org/2025</link>
      <description>Konferans detayları...</description>
    </item>
  </channel>
</rss>
``
/*
---

## 2. Go Koduyla Parse Etmek
*/

``go
package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// RSS struct yapıları
type RSS struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

func main() {
	// Normalde bunu bir URL'den http.Get ile çekebilirsin
	data := `
	<rss version="2.0">
	  <channel>
	    <title>Go Haberler</title>
	    <link>https://golang.org</link>
	    <description>Go diliyle ilgili güncel haberler</description>
	    <item>
	      <title>Go 1.23 çıktı!</title>
	      <link>https://golang.org/doc/go1.23</link>
	      <description>Yeni özellikler...</description>
	    </item>
	    <item>
	      <title>GoCon 2025</title>
	      <link>https://gocon.org/2025</link>
	      <description>Konferans detayları...</description>
	    </item>
	  </channel>
	</rss>`

	var rss RSS
	err := xml.Unmarshal([]byte(data), &rss)
	if err != nil {
		panic(err)
	}

	// Kanal bilgisi
	fmt.Println("Kanal:", rss.Channel.Title)
	fmt.Println("Link:", rss.Channel.Link)
	fmt.Println("Açıklama:", rss.Channel.Description)
	fmt.Println()

	// Haberler
	for i, item := range rss.Channel.Items {
		fmt.Printf("%d) %s\n", i+1, item.Title)
		fmt.Println("   Link:", item.Link)
		fmt.Println("   Açıklama:", strings.TrimSpace(item.Description))
	}
}
``
/*
---

## 3. Beklenen Çıktı

```
Kanal: Go Haberler
Link: https://golang.org
Açıklama: Go diliyle ilgili güncel haberler

1) Go 1.23 çıktı!
   Link: https://golang.org/doc/go1.23
   Açıklama: Yeni özellikler...
2) GoCon 2025
   Link: https://gocon.org/2025
   Açıklama: Konferans detayları...
```

---

# 📌 Özet

* `encoding/xml` sayesinde RSS feed gibi **XML tabanlı içerikleri kolayca parse edebilirsin.**
* Yapı (struct) tanımlarken **tag’ler** (`xml:"tag"`) kullanılır.
* RSS’de tipik olarak:

  * `<channel>` → kanal bilgisi
  * `<item>` → haber/podcast/entry
* Gerçek dünyada bu veri genelde **`http.Get`** ile internetten alınır, biz burada sabit XML string kullandık.

---

👉 İstersen sana bu örneği **gerçek bir RSS URL’si (ör. BBC News, Go Blog, Haber siteleri)** üzerinden `http.Get` ile çalışan bir sürümünü de gösterebilirim. Görmek ister misin?
EVET
*/
/*
Süper 👍 şimdi sana **gerçek bir RSS feed’i** `encoding/xml` ile **http.Get** üzerinden parse etmeyi göstereyim.

Ben örnek olarak **Go Blog RSS feed** (`https://blog.golang.org/feed.atom`) veya **BBC RSS feed** (`http://feeds.bbci.co.uk/news/rss.xml`) kullanacağım.

---

# 📌 Örnek: Gerçek RSS Feed Okuma
*/

``go
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// RSS Feed yapısı
type RSS struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

func main() {
	// BBC News RSS feed
	url := "http://feeds.bbci.co.uk/news/rss.xml"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Tüm XML'i oku
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var rss RSS
	err = xml.Unmarshal(data, &rss)
	if err != nil {
		panic(err)
	}

	// Kanal bilgisi
	fmt.Println("Kanal:", rss.Channel.Title)
	fmt.Println("Link:", rss.Channel.Link)
	fmt.Println("Açıklama:", rss.Channel.Description)
	fmt.Println()

	// İlk 5 haberi yazdıralım
	for i, item := range rss.Channel.Items {
		if i >= 5 {
			break
		}
		fmt.Printf("%d) %s\n", i+1, item.Title)
		fmt.Println("   Link:", item.Link)
		fmt.Println("   Açıklama:", strings.TrimSpace(item.Description))
	}
}
``
/*
---

# 📌 Çalıştırınca Ne Olur?

Programı çalıştırdığında, BBC’nin RSS feed’inden **canlı haber başlıkları** gelir:

```
Kanal: BBC News - Home
Link: http://www.bbc.co.uk/news/
Açıklama: BBC News - Home

1) Hurricane impacts Florida coast
   Link: http://www.bbc.co.uk/news/world-...
   Açıklama: BBC News provides trusted World and UK news...
2) Stock markets rally amid tech growth
   Link: http://www.bbc.co.uk/news/business-...
   Açıklama: Investors optimistic about AI-driven growth...
...
```

---

# 📌 Özet

* `http.Get` ile RSS XML’i indirdik.
* `encoding/xml.Unmarshal` ile `struct` içine parse ettik.
* İlk 5 haberi ekrana bastık.

---

👉 İstersen ben şimdi bunun bir **podcast RSS parser** versiyonunu da yapabilirim (ör. `<enclosure>` ile ses dosyalarını da çekmek için). Onu da görmek ister misin?
*/