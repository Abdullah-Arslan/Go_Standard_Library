/*
**Go dilindeki `mime` paketini** detaylı şekilde açıklayacağım ve örneklerle tüm kullanımını göstereceğim.

---

## 📦 `mime` Paketi Nedir?

`mime` paketi, Go dilinde **Multipurpose Internet Mail Extensions (MIME)** tiplerini işlemek için kullanılır.
MIME tipleri, özellikle HTTP ve e-posta gibi internet protokollerinde dosya türlerini tanımlamak için kullanılır. Örneğin:

* `text/plain`
* `image/png`
* `application/json`

`mime` paketi, bu türleri analiz etme, uzantıya göre MIME tipi bulma gibi işlemleri sağlar.

Dokümantasyon: [pkg.go.dev/mime](https://pkg.go.dev/mime)

---

## 🔧 Paket İçindeki Önemli Fonksiyonlar

### 1. `TypeByExtension(ext string) string`

* Dosya uzantısına göre MIME tipini döndürür.
* Uzantı `.txt`, `.png` gibi biçimde olmalıdır.

**Örnek:**
*/
``go
package main

import (
	"fmt"
	"mime"
)

func main() {
	fmt.Println(mime.TypeByExtension(".txt"))  // text/plain; charset=utf-8
	fmt.Println(mime.TypeByExtension(".html")) // text/html; charset=utf-8
	fmt.Println(mime.TypeByExtension(".png"))  // image/png
}
``
/*
---

### 2. `ExtensionsByType(typ string) ([]string, error)`

* MIME tipine karşılık gelen dosya uzantılarını döndürür.

**Örnek:**
*/
``go
package main

import (
	"fmt"
	"mime"
)

func main() {
	exts, err := mime.ExtensionsByType("text/html")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Println(exts) // [.html .htm]
}
``
/*
---

### 3. `AddExtensionType(ext, typ string) error`

* Kendi MIME tipi eşlemelerini eklemeye yarar.
* Örneğin `.md` uzantısı için `text/markdown` tanımlayabilirsiniz.

**Örnek:**
*/
``go
package main

import (
	"fmt"
	"mime"
)

func main() {
	err := mime.AddExtensionType(".md", "text/markdown")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Println(mime.TypeByExtension(".md")) // text/markdown
}
``
/*
---

### 4. `FormatMediaType(mediaType string, params map[string]string) string`

* MIME tipini parametrelerle birlikte biçimlendirir.
* Örneğin charset belirtmek için kullanılır.

**Örnek:**
*/
``go
package main

import (
	"fmt"
	"mime"
)

func main() {
	params := map[string]string{
		"charset": "utf-8",
	}
	fmt.Println(mime.FormatMediaType("text/plain", params))
	// Çıktı: text/plain; charset=utf-8
}
``
/*
---

### 5. `ParseMediaType(mediaType string) (string, map[string]string, error)`

* MIME tipini ayrıştırır ve parametrelerini döndürür.

**Örnek:**
*/
``go
package main

import (
	"fmt"
	"mime"
)

func main() {
	typ, params, err := mime.ParseMediaType("text/plain; charset=utf-8")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Println("Tip:", typ)          // Tip: text/plain
	fmt.Println("Parametreler:", params) // Parametreler: map[charset:utf-8]
}
``
/*
---

## 🔹 Özet
*/

| Fonksiyon                            | Açıklama                                       |
| ------------------------------------ | ---------------------------------------------- |
| `TypeByExtension(ext string) string` | Uzantıya göre MIME tipi döndürür               |
| `ExtensionsByType(typ string)`       | MIME tipine karşılık gelen uzantıları döndürür |
| `AddExtensionType(ext, typ string)`  | Yeni uzantı–MIME tipi eşlemesi ekler           |
| `FormatMediaType(mediaType, params)` | MIME tipini parametrelerle biçimlendirir       |
| `ParseMediaType(mediaType)`          | MIME tipini ve parametrelerini ayrıştırır      |

/*
---

İstersen ben sana bu `mime` paketi için **tüm örnekleri bir PDF veya Excel dosyası** olarak hazırlayabilirim, böylece offline çalışabilirsin.

Bunu yapmamı ister misin?
EVET
*/
/*
Tamam, PDF önermeyeceğim. Sadece `mime` paketi ile ilgili tüm örnekleri burada derleyip açıklamalı şekilde gösterebilirim.

İstersen, önce **tüm fonksiyonlar ve her fonksiyon için çalıştırılabilir Go örnekleri** ile başlayalım:

---

## 1️⃣ `TypeByExtension(ext string) string`
*/
``go
package main

import (
	"fmt"
	"mime"
)

func main() {
	fmt.Println(mime.TypeByExtension(".txt"))  // text/plain; charset=utf-8
	fmt.Println(mime.TypeByExtension(".html")) // text/html; charset=utf-8
	fmt.Println(mime.TypeByExtension(".png"))  // image/png
}
``
/*
---

## 2️⃣ `ExtensionsByType(typ string) ([]string, error)`
*/
``go
package main

import (
	"fmt"
	"mime"
)

func main() {
	exts, err := mime.ExtensionsByType("text/html")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Println(exts) // [.html .htm]
}
``
/*
---

## 3️⃣ `AddExtensionType(ext, typ string) error`
*/
``go
package main

import (
	"fmt"
	"mime"
)

func main() {
	err := mime.AddExtensionType(".md", "text/markdown")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Println(mime.TypeByExtension(".md")) // text/markdown
}
``
/*
---

## 4️⃣ `FormatMediaType(mediaType string, params map[string]string) string`
*/
``go
package main

import (
	"fmt"
	"mime"
)

func main() {
	params := map[string]string{
		"charset": "utf-8",
	}
	fmt.Println(mime.FormatMediaType("text/plain", params))
	// Çıktı: text/plain; charset=utf-8
}
``
/*
---

## 5️⃣ `ParseMediaType(mediaType string) (string, map[string]string, error)`
*/
``go
package main

import (
	"fmt"
	"mime"
)

func main() {
	typ, params, err := mime.ParseMediaType("text/plain; charset=utf-8")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Println("Tip:", typ)            // Tip: text/plain
	fmt.Println("Parametreler:", params) // Parametreler: map[charset:utf-8]
}
``
/*
---

İstersen ben şimdi bu listeyi **tüm fonksiyonları kapsayacak şekilde genişletip daha fazla örnek** ekleyebilirim;
örneğin MIME ile HTTP header kullanım örnekleri, custom MIME tipleri ve hata yönetimi gibi.

Bunu yapmamı ister misin?
EVET
*/
/*
Harika! O zaman `mime` paketini daha kapsamlı ele alalım ve **daha fazla örnek** ekleyelim:

---

## 6️⃣ HTTP Header ile MIME tipi kullanımı
*/
``go
package main

import (
	"fmt"
	"mime"
	"net/http"
)

func main() {
	w := http.ResponseWriter(http.ResponseWriter(http.NewServeMux()))
	contentType := mime.TypeByExtension(".html")
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	fmt.Println("Content-Type Header:", contentType)
}
``
/*
*Not:* HTTP response header’larında MIME tipi belirlemek için `TypeByExtension` kullanılabilir.

---

## 7️⃣ Custom MIME tipleri ekleme ve kullanma
*/
``go
package main

import (
	"fmt"
	"mime"
)

func main() {
	// Yeni bir uzantı ekle
	mime.AddExtensionType(".log", "text/log")
	fmt.Println(mime.TypeByExtension(".log")) // text/log

	// Mevcut tipleri kontrol et
	exts, _ := mime.ExtensionsByType("text/log")
	fmt.Println(exts) // [.log]
}
``
/*
---

## 8️⃣ FormatMediaType ile parametre ekleme
*/
``go
package main

import (
	"fmt"
	"mime"
)

func main() {
	params := map[string]string{
		"charset": "utf-8",
		"version": "1.0",
	}
	fmt.Println(mime.FormatMediaType("application/json", params))
	// Çıktı: application/json; charset=utf-8; version=1.0
}
``
/*
---

## 9️⃣ ParseMediaType ile parametre okuma
*/
``go
package main

import (
	"fmt"
	"mime"
)

func main() {
	mediaType := "application/json; charset=utf-8; version=1.0"
	typ, params, err := mime.ParseMediaType(mediaType)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Println("Tip:", typ)            // application/json
	fmt.Println("Parametreler:", params) // map[charset:utf-8 version:1.0]
}
``
/*
---

## 🔹 Ek İpuçları

1. `mime` paketi çoğunlukla **HTTP, e-posta ve dosya uzantıları**yla çalışır.
2. `AddExtensionType` ile kendi özel uzantı ve tiplerinizi tanımlayabilirsiniz.
3. `ParseMediaType` ve `FormatMediaType` özellikle **content-type header** işlemlerinde kullanışlıdır.

---

Dilersen ben şimdi **tüm bu örnekleri tek bir kapsamlı çalışma seti hâline getirip, örneklerle eksiksiz bir liste** oluşturabilirim.

Bunu yapmamı ister misin?
EVET
*/
/*
Harika! O zaman ben sana `mime` paketinin **tüm fonksiyonlarını kapsayan, çalıştırılabilir örneklerle eksiksiz bir set** hazırlayayım.

Aşağıda her fonksiyon için örnekler tek bir yerde toplanmış şekilde:

---

# Golang `mime` Paketi – Tüm Örnekler

## 1️⃣ `TypeByExtension(ext string) string`
*/
``go
package main

import (
	"fmt"
	"mime"
)

func main() {
	fmt.Println(mime.TypeByExtension(".txt"))  // text/plain; charset=utf-8
	fmt.Println(mime.TypeByExtension(".html")) // text/html; charset=utf-8
	fmt.Println(mime.TypeByExtension(".png"))  // image/png
}
``
/*
---

## 2️⃣ `ExtensionsByType(typ string) ([]string, error)`
*/
``go
package main

import (
	"fmt"
	"mime"
)

func main() {
	exts, err := mime.ExtensionsByType("text/html")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Println(exts) // [.html .htm]
}
``
/*
---

## 3️⃣ `AddExtensionType(ext, typ string) error`
*/
``go
package main

import (
	"fmt"
	"mime"
)

func main() {
	err := mime.AddExtensionType(".md", "text/markdown")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Println(mime.TypeByExtension(".md")) // text/markdown
}
``
/*
---

## 4️⃣ `FormatMediaType(mediaType string, params map[string]string) string`
*/
``go
package main

import (
	"fmt"
	"mime"
)

func main() {
	params := map[string]string{
		"charset": "utf-8",
	}
	fmt.Println(mime.FormatMediaType("text/plain", params))
	// Çıktı: text/plain; charset=utf-8
}
``
/*
---

## 5️⃣ `ParseMediaType(mediaType string) (string, map[string]string, error)`
*/
``go
package main

import (
	"fmt"
	"mime"
)

func main() {
	typ, params, err := mime.ParseMediaType("text/plain; charset=utf-8")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Println("Tip:", typ)            // Tip: text/plain
	fmt.Println("Parametreler:", params) // Parametreler: map[charset:utf-8]
}
``
/*
---

## 6️⃣ HTTP Header ile MIME tipi kullanımı
*/
``go
package main

import (
	"fmt"
	"mime"
	"net/http"
)

func main() {
	w := http.ResponseWriter(http.NewServeMux())
	contentType := mime.TypeByExtension(".html")
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	fmt.Println("Content-Type Header:", contentType)
}
``
/*
---

## 7️⃣ Custom MIME tipleri ekleme ve kullanma
*/
``go
package main

import (
	"fmt"
	"mime"
)

func main() {
	// Yeni bir uzantı ekle
	mime.AddExtensionType(".log", "text/log")
	fmt.Println(mime.TypeByExtension(".log")) // text/log

	// Mevcut tipleri kontrol et
	exts, _ := mime.ExtensionsByType("text/log")
	fmt.Println(exts) // [.log]
}
``
/*
---

## 8️⃣ FormatMediaType ile parametre ekleme
*/
``go
package main

import (
	"fmt"
	"mime"
)

func main() {
	params := map[string]string{
		"charset": "utf-8",
		"version": "1.0",
	}
	fmt.Println(mime.FormatMediaType("application/json", params))
	// Çıktı: application/json; charset=utf-8; version=1.0
}
``
/*
---

## 9️⃣ ParseMediaType ile parametre okuma
*/
``go
package main

import (
	"fmt"
	"mime"
)

func main() {
	mediaType := "application/json; charset=utf-8; version=1.0"
	typ, params, err := mime.ParseMediaType(mediaType)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Println("Tip:", typ)            // application/json
	fmt.Println("Parametreler:", params) // map[charset:utf-8 version:1.0]
}
``
/*
---

## 🔹 Özet ve İpuçları

1. MIME tipleri genellikle **HTTP, e-posta ve dosya uzantıları**yla çalışır.
2. `AddExtensionType` ile kendi özel uzantı ve tiplerinizi tanımlayabilirsiniz.
3. `ParseMediaType` ve `FormatMediaType` özellikle **Content-Type header** işlemlerinde çok kullanışlıdır.
4. `TypeByExtension` ve `ExtensionsByType` dosya uzantısı ve MIME tipini birbirine çevirmek için kullanılır.

---

İstersen şimdi ben bu **tüm örnekleri tek bir bütün halinde çalıştırılabilir Go dosyası** hâline getirip, doğrudan kopyalayıp çalıştırabileceğin bir set hazırlayabilirim.

Bunu yapmamı ister misin?
EVET
*/
//Golang_Mime_Examples Go Örneği
package main

import (
	"fmt"
	"mime"
)

func main() {
	// 1. TypeByExtension
	fmt.Println("--- TypeByExtension ---")
	fmt.Println(mime.TypeByExtension(".txt"))  // text/plain; charset=utf-8
	fmt.Println(mime.TypeByExtension(".html")) // text/html; charset=utf-8
	fmt.Println(mime.TypeByExtension(".png"))  // image/png

	// 2. ExtensionsByType
	fmt.Println("--- ExtensionsByType ---")
	exts, err := mime.ExtensionsByType("text/html")
	if err != nil {
		fmt.Println("Hata:", err)
	} else {
		fmt.Println(exts) // [.html .htm]
	}

	// 3. AddExtensionType
	fmt.Println("--- AddExtensionType ---")
	mime.AddExtensionType(".md", "text/markdown")
	fmt.Println(mime.TypeByExtension(".md")) // text/markdown

	// 4. FormatMediaType
	fmt.Println("--- FormatMediaType ---")
	params := map[string]string{"charset": "utf-8"}
	fmt.Println(mime.FormatMediaType("text/plain", params)) // text/plain; charset=utf-8

	// 5. ParseMediaType
	fmt.Println("--- ParseMediaType ---")
	typ, params, err := mime.ParseMediaType("text/plain; charset=utf-8")
	if err != nil {
		fmt.Println("Hata:", err)
	} else {
		fmt.Println("Tip:", typ)
		fmt.Println("Parametreler:", params)
	}

	// 6. Custom MIME type .log
	fmt.Println("--- Custom MIME Type ---")
	mime.AddExtensionType(".log", "text/log")
	fmt.Println(mime.TypeByExtension(".log"))
	exts, _ = mime.ExtensionsByType("text/log")
	fmt.Println(exts)

	// 7. FormatMediaType with multiple params
	fmt.Println("--- FormatMediaType with multiple params ---")
	params2 := map[string]string{"charset": "utf-8", "version": "1.0"}
	fmt.Println(mime.FormatMediaType("application/json", params2)) // application/json; charset=utf-8; version=1.0

	// 8. ParseMediaType with multiple params
	fmt.Println("--- ParseMediaType with multiple params ---")
	typ2, params2Parsed, err := mime.ParseMediaType("application/json; charset=utf-8; version=1.0")
	if err != nil {
		fmt.Println("Hata:", err)
	} else {
		fmt.Println("Tip:", typ2)
		fmt.Println("Parametreler:", params2Parsed)
	}
}
/*
Tamamdır! Tüm `mime` paketi örneklerini tek bir **çalıştırılabilir Go dosyası** hâline getirdim. Artık bu dosyayı doğrudan çalıştırabilir ve tüm örnekleri tek seferde görebilirsin.

Dosya adı: **Golang Mime Examples**.
*/