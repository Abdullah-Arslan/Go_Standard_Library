/*
Go dilinde **`html`** paketini detaylıca inceleyelim.

---

# 🌐 `html` Paketi – Genel Bakış

Go’nun standart kütüphanesindeki **`html`** paketi, HTML metinlerini işlemek ve güvenli hâle getirmek için kullanılır.
Özellikle **XSS (Cross-Site Scripting)** gibi saldırılara karşı veriyi temizlemek veya HTML karakterlerini encode/decode etmek için idealdir.

Paket URL: [`https://pkg.go.dev/html`](https://pkg.go.dev/html)

---

# 📦 Paket İçe Aktarma
*/
``go
import "html"
``
/*
---

# 🔑 Önemli Fonksiyonlar

### 1. `html.EscapeString(s string) string`

* HTML içinde özel karakterleri (`<`, `>`, `&`, `"`) **HTML entity** hâline çevirir.
* Web uygulamalarında **user input**u güvenli hâle getirmek için kullanılır.
*/
``go
package main

import (
	"fmt"
	"html"
)

func main() {
	input := `<script>alert("XSS")</script>`
	escaped := html.EscapeString(input)
	fmt.Println("Escaped HTML:", escaped)
}
``

//📌 Çıktı:

``
Escaped HTML: &lt;script&gt;alert(&quot;XSS&quot;)&lt;/script&gt;
``
/*
---

### 2. `html.UnescapeString(s string) string`

* HTML entity hâline gelmiş karakterleri **orijinal hâline geri döndürür**.
*/
``go
func main() {
	escaped := `&lt;div&gt;Hello &amp; World&lt;/div&gt;`
	unescaped := html.UnescapeString(escaped)
	fmt.Println("Unescaped HTML:", unescaped)
}
``

//📌 Çıktı:

``
Unescaped HTML: <div>Hello & World</div>
``
/*
---

# 📂 Kullanım Senaryoları

### ✅ 1. Web Formları için Güvenlik

Kullanıcıdan gelen input’u HTML sayfasında göstermek istiyorsun.

* **Kaçırılmamış** karakterler XSS açığı yaratabilir.
* `html.EscapeString` ile güvenli hâle getirebilirsin:
*/
``go
func renderUserInput(userInput string) string {
	return "<p>" + html.EscapeString(userInput) + "</p>"
}

func main() {
	user := `<script>alert("hack")</script>`
	safeHTML := renderUserInput(user)
	fmt.Println(safeHTML)
}
``

//📌 Çıktı:

``
<p>&lt;script&gt;alert(&quot;hack&quot;)&lt;/script&gt;</p>
``
/*
---

### ✅ 2. HTML Entity Çözme

* HTML sayfasından gelen verilerdeki entity’leri geri çevirmek için:
*/
``go
func main() {
	data := "Tom &amp; Jerry &lt;Cartoon&gt;"
	fmt.Println(html.UnescapeString(data))
}
``

//📌 Çıktı:

``
Tom & Jerry <Cartoon>
``
/*
---

### ✅ 3. Örnek: Basit HTML Encoder/Decoder
*/
``go
package main

import (
	"fmt"
	"html"
)

func main() {
	text := `<h1>Merhaba "Dünya" & Go!</h1>`

	escaped := html.EscapeString(text)
	fmt.Println("Escaped:", escaped)

	unescaped := html.UnescapeString(escaped)
	fmt.Println("Unescaped:", unescaped)
}
``

//📌 Çıktı:

``
Escaped: &lt;h1&gt;Merhaba &quot;Dünya&quot; &amp; Go!&lt;/h1&gt;
Unescaped: <h1>Merhaba "Dünya" & Go!</h1>
``
/*
---

# 📝 Özet

* **`html` paketi**, HTML karakterlerini güvenli hâle getirmek ve geri çözmek için kullanılır.
* Fonksiyonlar:

  * `EscapeString(s string)` → HTML özel karakterleri escape eder
  * `UnescapeString(s string)` → HTML entity’lerini orijinal hâline getirir
* Kullanım alanları: Web form güvenliği, XSS önleme, veri temizleme ve HTML parsing öncesi ön işleme.

---

👉 İstersen ben sana **html.EscapeString ile kullanıcı inputlarını otomatik güvenli hâle getiren küçük bir Go web uygulaması** örneği hazırlayabilirim. Bu sayede gerçek bir senaryo görmüş olursun. Bunu ister misin?
EVET
*/
/*
Harika! 🚀 O zaman sana **Go ile basit bir web uygulaması** hazırlayacağım.
Bu uygulama, kullanıcıdan alınan inputu **HTML olarak güvenli bir şekilde** gösterecek ve **XSS saldırılarına karşı korunacak**.

---

## 📂 Proje: `safe_web_input`

### `main.go`
*/
``go
package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	// Formu gösteren handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
			<!DOCTYPE html>
			<html>
			<head><title>Safe Input</title></head>
			<body>
			<h2>Kullanıcı Inputu Göster</h2>
			<form action="/show" method="post">
				<input type="text" name="user_input" placeholder="Bir şey yazın">
				<input type="submit" value="Göster">
			</form>
			</body>
			</html>
		`)
	})

	// Kullanıcı inputunu işleyen ve güvenli şekilde gösteren handler
	http.HandleFunc("/show", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Form verisi okunamadı", http.StatusBadRequest)
			return
		}
		input := r.FormValue("user_input")

		// HTML güvenli hâle getir
		safeInput := html.EscapeString(input)

		fmt.Fprintf(w, `
			<!DOCTYPE html>
			<html>
			<head><title>Güvenli Input</title></head>
			<body>
			<h2>Kullanıcının Girdiği Veri:</h2>
			<p>%s</p>
			<a href="/">Geri</a>
			</body>
			</html>
		`, safeInput)
	})

	fmt.Println("Sunucu çalışıyor: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
``
/*
---

## 📌 Kullanım

1. Programı çalıştır:
*/

``bash
go run main.go
``
/*
2. Tarayıcıdan `http://localhost:8080` adresine git.

3. Input kutusuna örneğin `<script>alert("XSS")</script>` yaz ve "Göster" butonuna bas.

📌 Çıktı:

```
<p>&lt;script&gt;alert(&quot;XSS&quot;)&lt;/script&gt;</p>
```

✅ Script çalışmaz, kullanıcı inputu güvenli bir şekilde gösterilir.

---

## 🔑 Özet

* `html.EscapeString` sayesinde **XSS saldırılarına karşı korunuruz**.
* Kullanıcıdan gelen tüm veriler HTML entity’ye çevrilir.
* Basit web formu ve güvenli gösterim için ideal bir örnek oldu.

---

İstersen ben bu projeyi bir adım daha ileri götürüp:

* **Textarea ve çok satırlı inputları** da güvenli gösterecek
* Ve **inputları bir liste halinde saklayıp tekrar gösterecek** şekilde geliştirebilirim.

Bunu ister misin?

*/
