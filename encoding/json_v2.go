/*
Golang’ın **`encoding/json/v2`** paketi, Go 1.23 ile gelen **deneysel (experimental)** bir pakettir.
Bu paket aslında uzun zamandır kullanılan **`encoding/json`** paketinin modernleştirilmiş ve daha güçlü bir sürümüdür.
Amaç, eski `encoding/json`’da eksik kalan veya verimsiz çalışan noktaları iyileştirmek, daha esnek API’ler sunmak ve performansı artırmaktır.

---

## 1. Genel Özellikler

* **Daha katı hata kontrolü**: JSON formatı ile Go tipleri arasındaki uyuşmazlıkları daha iyi raporlar.
* **Daha esnek API**: `Unmarshal` ve `Marshal` dışında **`Decoder` / `Encoder`** geliştirilmiş versiyonlar sunar.
* **Performans geliştirmeleri**: Özellikle büyük JSON verilerinde daha hızlı çalışır.
* **Daha doğru tip eşleştirme**: Örneğin sayıların `float64` veya `int64`’e çevrilmesi daha kontrollüdür.
* **`jsontext` entegrasyonu**: JSON’un metin tabanlı token seviyesinde işlenmesine izin verir.

`json/v2` şu anda **deneysel** olduğundan **kararlı bir sürüm değildir**. API ileride değişebilir.

---

## 2. Kullanım

Normal `encoding/json` gibi import edilir, sadece yolu değişiktir:
*/

``go
import "encoding/json/v2"
``
/*
---

## 3. Temel Kullanımlar

### 3.1 Struct → JSON (Marshal)
*/

``go
package main

import (
	"encoding/json/v2"
	"fmt"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

func main() {
	u := User{ID: 1, Name: "Ali"}

	// Struct'tan JSON'a çevirme
	data, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
``
/*
**Çıktı:**

```json
{"id":1,"name":"Ali"}
```

(`omitempty` sayesinde `Email` boş olduğunda JSON’a yazılmaz.)

---

### 3.2 JSON → Struct (Unmarshal)
*/

``go
package main

import (
	"encoding/json/v2"
	"fmt"
)

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	jsonData := `{"name":"Laptop","price":14999.90}`

	var p Product
	err := json.Unmarshal([]byte(jsonData), &p)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", p)
}
``
/*
**Çıktı:**

```
{Name:Laptop Price:14999.9}
```

---

### 3.3 `Decoder` ile JSON Akışı Okuma

Büyük dosyalarda JSON’u tek seferde `Unmarshal` etmek bellek dostu değildir.
Bunun için **`Decoder`** kullanılır.
*/

``go
package main

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

func main() {
	data := `{"id":1,"name":"Ali"}{"id":2,"name":"Veli"}`

	dec := json.NewDecoder(strings.NewReader(data))

	for {
		var obj map[string]any
		err := dec.Decode(&obj)
		if err != nil {
			break
		}
		fmt.Println(obj)
	}
}
``
/*
**Çıktı:**

```
map[id:1 name:Ali]
map[id:2 name:Veli]
```

---

### 3.4 `Encoder` ile JSON Yazma
*/

``go
package main

import (
	"encoding/json/v2"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	people := []Person{
		{"Ahmet", 30},
		{"Mehmet", 25},
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ") // JSON'u daha okunabilir yaz

	enc.Encode(people)
}
``

//**Çıktı:**

``json
[
  {
    "name": "Ahmet",
    "age": 30
  },
  {
    "name": "Mehmet",
    "age": 25
  }
]
``
/*
---

## 4. `Unmarshal` Hata Kontrolü

`json/v2`, `json`’dan farklı olarak **daha katı** davranır.
Örneğin, JSON’da bilinmeyen alan varsa hata fırlatabilir:
*/

``go
package main

import (
	"encoding/json/v2"
	"fmt"
)

type Config struct {
	Port int `json:"port"`
}

func main() {
	data := `{"port":8080, "host":"localhost"}`

	var c Config
	err := json.Unmarshal([]byte(data), &c)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Println(c)
}
``

/*
**Çıktı (hata verir):**

```
Hata: json: unknown field "host"
```

Eski `encoding/json` bunu sessizce yoksayardı, `json/v2` ise daha güvenlidir.

---

## 5. Token Bazlı İşleme (`jsontext` ile)

`json/v2`, alt seviyede **`jsontext`** paketiyle çalışır.
Bu sayede JSON’u **token token** (ör. `{`, `}`, `"key"`, `value`) işleyebilirsiniz.
*/

``go
package main

import (
	"encoding/json/jsontext"
	"fmt"
)

func main() {
	data := []byte(`{"a":1,"b":true,"c":"hello"}`)

	dec := jsontext.NewDecoder(data)

	for {
		tok, err := dec.ReadToken()
		if err != nil {
			break
		}
		fmt.Printf("Token: %v\n", tok)
	}
}
``
/*
**Çıktı:**

```
Token: {  
Token: a  
Token: 1  
Token: b  
Token: true  
Token: c  
Token: hello  
Token: }  
```

---

## 6. Özet

* `encoding/json/v2`, Go’nun **yeni nesil JSON işleme paketi**dir.
* `encoding/json` ile çok benzerdir, ama daha **katı**, **güvenli** ve **hızlıdır**.
* Temel fonksiyonlar:

  * `Marshal` → Struct → JSON
  * `Unmarshal` → JSON → Struct
  * `Encoder` / `Decoder` → Akış (stream) bazlı işleme
  * `jsontext` → Token bazlı düşük seviye işleme
* Henüz **deneysel** olduğundan, resmi Go sürümlerinde API değişiklikleri olabilir.

---

👉 İstersen sana `json/v2` ile eski `encoding/json` arasındaki **farkları tablo halinde** çıkarayım. İstiyor musun?
EVET
*/

/*
Harika 👍 O zaman sana **`encoding/json`** ile **`encoding/json/v2`** arasındaki farkları tablo halinde çıkarıyorum:

---

# `encoding/json` vs `encoding/json/v2` Karşılaştırma Tablosu

| Özellik                   | `encoding/json` (Eski)                                     | `encoding/json/v2` (Yeni)                                 |
| ------------------------- | ---------------------------------------------------------- | --------------------------------------------------------- |
| **Durum**                 | Kararlı (Go 1’den beri var)                                | **Deneysel** (Go 1.23+ ile geldi, API değişebilir)        |
| **Performans**            | Orta seviye, büyük JSON’larda yavaş                        | Daha hızlı, bellek kullanımı daha verimli                 |
| **Hata Yönetimi**         | Çoğu hatayı yutar (ör. bilinmeyen alanları sessizce atlar) | Daha **katı**: bilinmeyen alanlarda hata döner            |
| **Sayıların İşlenmesi**   | Varsayılan olarak `float64` kullanır                       | Daha doğru tip dönüşümleri (`int64`, `float64`)           |
| **API**                   | `Marshal`, `Unmarshal`, `Encoder`, `Decoder`               | Aynı ama daha gelişmiş, `jsontext` ile alt seviye kontrol |
| **Bilinmeyen Alanlar**    | Yoksayılır                                                 | Varsayılan olarak **hata fırlatır**                       |
| **Struct Tag Desteği**    | `json:"field,omitempty"`                                   | Aynı, ama bazı ek kurallar daha katı işleniyor            |
| **Bozuk JSON**            | Hata mesajları bazen belirsiz                              | Daha açıklayıcı hata mesajları                            |
| **Token Seviyesi İşleme** | Sınırlı (`Decoder.Token`)                                  | Gelişmiş: `jsontext` ile düşük seviye token işleme        |
| **Kapsam**                | Yıllardır projelerde kullanılıyor, güvenli                 | Yeni olduğu için henüz yaygın değil                       |
| **Gelecek**               | Desteklenmeye devam edecek ama değişmeyecek                | Go’nun gelecekteki JSON standardı olma adayı              |

---

## Örnek Fark

### `encoding/json`
*/

``go
package main

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	Port int `json:"port"`
}

func main() {
	data := `{"port":8080, "host":"localhost"}`
	var c Config
	_ = json.Unmarshal([]byte(data), &c)
	fmt.Println(c) // {8080}  (host alanı sessizce yok sayılır)
}
``

//### `encoding/json/v2`

``go
package main

import (
	"encoding/json/v2"
	"fmt"
)

type Config struct {
	Port int `json:"port"`
}

func main() {
	data := `{"port":8080, "host":"localhost"}`
	var c Config
	err := json.Unmarshal([]byte(data), &c)
	fmt.Println("Hata:", err) // unknown field "host"
}
``
/*
---

✅ Yani özetle: `json/v2` → **daha güvenli, hızlı, katı**
`encoding/json` → **daha esnek ama bazı hataları saklayan**

---

Sana istersen ben şimdi **bir tablo daha çıkarabilirim: Hangi senaryoda `encoding/json`, hangi senaryoda `json/v2` kullanmalısın** diye. İstiyor musun?
EVET
*/

/*
Süper 👍 O zaman sana **hangi durumda `encoding/json`, hangi durumda `encoding/json/v2` kullanmalısın** diye bir tablo çıkarıyorum:

---

# Kullanım Senaryoları: `encoding/json` vs `encoding/json/v2`

| Senaryo                                  | `encoding/json` (Eski)                                                  | `encoding/json/v2` (Yeni)                                             |
| ---------------------------------------- | ----------------------------------------------------------------------- | --------------------------------------------------------------------- |
| **Kararlılık / Üretim**                  | ✅ Çok kararlı, 10+ yıldır projelerde güvenle kullanılıyor.              | ⚠️ Deneysel, API değişebilir, uzun vadede yeni standarda dönüşebilir. |
| **Hızlı prototip geliştirme**            | ✅ Daha esnek (bilinmeyen alanları yok sayar, hızlı ilerlersin).         | ❌ Katı hata kontrolü yüzünden prototipte yavaşlatabilir.              |
| **Büyük JSON dosyaları**                 | ⚠️ Daha yavaş, bellek kullanımında verimsiz olabilir.                   | ✅ Daha hızlı ve bellek dostu (stream & token işleme daha gelişmiş).   |
| **Hassas veri / Güvenlik**               | ❌ Bilinmeyen alanları sessizce atlar → güvenlik açığı olabilir.         | ✅ Bilinmeyen alanlarda hata fırlatır, daha güvenlidir.                |
| **JSON → Struct tip dönüşümü**           | ⚠️ Bazen tip uyuşmazlıklarını fark etmez (`int` yerine `float64` atar). | ✅ Daha doğru tip dönüşümleri ve daha net hata mesajları.              |
| **Backward compatibility (geri uyumlu)** | ✅ Eski Go sürümleriyle uyumlu.                                          | ❌ Sadece Go 1.23+ ve ilerisi.                                         |
| **Detaylı hata mesajları**               | ❌ Hata mesajları kısa ve belirsiz olabilir.                             | ✅ Hata mesajları daha açıklayıcı.                                     |
| **Token bazlı işleme (low-level)**       | ⚠️ `Decoder.Token()` sınırlı.                                           | ✅ `jsontext` ile çok güçlü token işleme var.                          |
| **Topluluk & Destek**                    | ✅ Çok yaygın, her kütüphanede kullanılıyor.                             | ⚠️ Henüz yaygın değil, topluluk desteği düşük.                        |

---

## Özet

* Eğer **bugün prod ortamında stabil bir şey** geliştiriyorsan → **`encoding/json`** kullan.
* Eğer **geleceğe hazırlık yapıyorsan, güvenlik/hassasiyet kritikse** ve Go 1.23+ kullanıyorsan → **`json/v2`** dene.
* Büyük JSON verilerini stream etmek veya token bazlı işlemek istiyorsan → **`json/v2` + `jsontext`** tercih et.

---

👉 İstersen sana **aynı JSON dosyasını hem `encoding/json` hem `json/v2` ile parse eden bir mini proje** (örnek kıyaslama) hazırlayayım. Böylece farkları canlı olarak görebilirsin. İstiyor musun?
EVET
*/

/*
Harika 👍 O zaman sana **aynı JSON verisini hem `encoding/json` hem de `encoding/json/v2` ile parse eden bir mini proje** hazırlayayım. Bu proje sayesinde farkları canlı olarak görebilirsin.

---

# 📌 Mini Proje: `encoding/json` vs `encoding/json/v2`

## Proje Yapısı

```
json-compare/
│
├── main.go
└── go.mod
```

---

## `go.mod`
?*/

``go
module json-compare

go 1.23
``
/*
---

## `main.go`
*/

``go
package main

import (
	"encoding/json"
	jsonv2 "encoding/json/v2"
	"fmt"
)

type Config struct {
	Port int    `json:"port"`
	Host string `json:"host"`
}

func main() {
	// JSON verimiz bilinmeyen bir alan içeriyor ("debug")
	data := `{"port":8080, "host":"localhost", "debug":true}`

	fmt.Println("=== encoding/json (eski) ===")
	var cfg1 Config
	err := json.Unmarshal([]byte(data), &cfg1)
	if err != nil {
		fmt.Println("Hata:", err)
	} else {
		fmt.Printf("Sonuç: %+v\n", cfg1)
	}

	fmt.Println("\n=== encoding/json/v2 (yeni) ===")
	var cfg2 Config
	err = jsonv2.Unmarshal([]byte(data), &cfg2)
	if err != nil {
		fmt.Println("Hata:", err)
	} else {
		fmt.Printf("Sonuç: %+v\n", cfg2)
	}
}
``
/*
---

## Beklenen Çıktı

### `encoding/json`

```
=== encoding/json (eski) ===
Sonuç: {Port:8080 Host:localhost}
```

(`debug` alanı sessizce yok sayılır.)

### `encoding/json/v2`

```
=== encoding/json/v2 (yeni) ===
Hata: json: unknown field "debug"
```

(`debug` alanı bilinmediği için hata fırlatır.)

---

✅ Bu basit örnek şunu gösteriyor:

* `encoding/json` → Daha **esnek** ama hataları gizleyebilir.
* `encoding/json/v2` → Daha **katı** ama güvenlik açısından daha sağlam.

---

👉 İstersen ben bu projeyi biraz daha büyütüp **Decoder ile stream işleme** de ekleyebilirim (büyük JSON dosyalarını nasıl farklı işlediklerini görebilirsin). İster misin?
*/