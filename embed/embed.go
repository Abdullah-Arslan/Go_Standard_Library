/*
**Go (Golang) dilinde `embed` paketini** en ince ayrıntısına kadar anlatayım.

---

# 📌 `embed` Paketi Nedir?

Go 1.16 ile birlikte gelen `embed` paketi, **dosyaları veya dizinleri Go binary’sinin içine gömmek** için kullanılır.
Böylece dağıttığınız program tek bir dosya olur, harici olarak CSS, HTML, resim, config gibi dosyaları taşımanız gerekmez.

`embed` aslında **özel bir derleme-time direktifi** (`//go:embed`) ile çalışır. Yani:

* Kaynak dosyalar, Go koduna **derleme sırasında** gömülür.
* Çalışma zamanında (`runtime`) dosya sistemine ihtiyaç duyulmaz.

Desteklediği türler:

* `string` → Tek dosya içeriğini string olarak almak için
* `[]byte` → Tek dosya içeriğini `byte` dizisi olarak almak için
* `embed.FS` → Birden çok dosya/dizini sanal dosya sistemi (read-only) olarak almak için

---

# 📌 Kullanımı

## 1. Tek dosyayı `string` olarak gömmek
*/
``go
package main

import (
	_ "embed"
	"fmt"
)

//go:embed hello.txt
var hello string

func main() {
	fmt.Println(hello)
}
``
/*
👉 `hello.txt` içeriği direkt `hello` değişkenine gömülür.
Örn. `hello.txt` → `Merhaba Dünya` yazıyorsa, program ekrana `"Merhaba Dünya"` basar.

---

## 2. Tek dosyayı `[]byte` olarak gömmek
?7
``go
package main

import (
	_ "embed"
	"fmt"
)

//go:embed logo.png
var logo []byte

func main() {
	fmt.Println("Logo boyutu:", len(logo), "byte")
}
``

/*
👉 Burada `logo.png` dosyasını `[]byte` içinde tutuyoruz.
Resmi HTTP response olarak dönebilir, disk’e tekrar yazabilir ya da memory’den işleyebiliriz.

---

## 3. Birden fazla dosyayı `embed.FS` ile gömmek
*/
``go
package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
)

//go:embed templates/*
var templates embed.FS

func main() {
	entries, _ := fs.ReadDir(templates, "templates")
	for _, e := range entries {
		fmt.Println("Bulunan dosya:", e.Name())
	}

	// Tek dosya okumak
	data, _ := templates.ReadFile("templates/index.html")
	fmt.Println(string(data))

	// Disk’e yazmak
	os.WriteFile("copy_index.html", data, 0644)
}
``
/*
👉 `templates` klasörü altındaki **tüm dosyalar** binary içine gömülür.
Böylece dışarıya HTML dosyaları gönderebiliriz.

---

## 4. Birden fazla dosya deseniyle gömmek
*/
``go
package main

import (
	"embed"
	"fmt"
)

//go:embed static/*.css static/*.js
var staticFiles embed.FS

func main() {
	css, _ := staticFiles.ReadFile("static/style.css")
	fmt.Println("CSS içeriği:", string(css))
}
``
/*
👉 Birden fazla desen kullanabilirsin. `*.css` ve `*.js` dosyaları gömülür.

---

## 5. HTTP sunucusunda gömülü dosya kullanmak
*/

``go
package main

import (
	"embed"
	"net/http"
)

//go:embed static
var static embed.FS

func main() {
	// "static" klasörünü /assets yoluna bağla
	fs := http.FileServer(http.FS(static))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.ListenAndServe(":8080", nil)
}
``

/*
👉 Burada `static/` klasöründeki tüm dosyalar (CSS, JS, resimler) **binary içine gömülür** ve
`http://localhost:8080/assets/...` üzerinden servis edilir.

---

# 📌 Önemli Notlar

1. `//go:embed` **yorum satırı gibi görünür ama özel bir direktiftir**. Mutlaka değişkenin üstünde olmalı.
2. Gömülen dosyalar **salt okunur** (read-only).
3. Dosya yolları **Go dosyasının bulunduğu klasöre göre** ayarlanır.
4. `embed.FS` → sadece okuma işlemleri için kullanılır, yazma desteklenmez.
5. Büyük dosyalar (örn. 500MB video) binary boyutunu çok artırır, dikkatli olunmalı.

---

# 📌 Gerçek Hayatta Kullanım Senaryoları

* ✅ HTML şablonlarını gömmek (web app dağıtırken)
* ✅ CSS, JS, resim gibi static dosyaları tek binary’ye paketlemek
* ✅ Config dosyalarını yanında taşıma derdinden kurtulmak
* ✅ CLI tool’larında default template veya örnek dosyalar sunmak

---

📌 Özetle:
`embed` paketi, **Go uygulamanıza dosya ve dizinleri doğrudan gömerek** tek bir dağıtılabilir binary oluşturmanızı sağlar.

---

İstersen sana küçük bir **mini web uygulaması** örneği hazırlayayım, `embed` ile HTML + CSS gömüp tarayıcıda çalıştıracak şekilde. İster misin?
EVET
*/

/*
Süper 👍 O zaman sana **Go’da `embed` ile HTML + CSS gömülmüş küçük bir web uygulaması** örneği yapayım.
Bu uygulama:

* `index.html` ve `style.css` dosyalarını binary içine gömecek,
* HTTP sunucusunda tarayıcıya servis edecek,
* Tek bir `.exe` (veya Linux/Mac’te binary) dosya olarak çalışacak.

---

# 📌 Proje Yapısı

```
myapp/
 ├─ main.go
 ├─ templates/
 │   └─ index.html
 └─ static/
     └─ style.css
```

---

# 📌 Kod: `main.go`
*/
``go
package main

import (
	"embed"
	"fmt"
	"net/http"
)

// HTML şablonları
//go:embed templates/*
var templates embed.FS

// Statik dosyalar (CSS, JS, resimler)
//go:embed static/*
var static embed.FS

func main() {
	// HTML endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := templates.ReadFile("templates/index.html")
		if err != nil {
			http.Error(w, "Sayfa bulunamadı", 500)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(data)
	})

	// Statik dosyalar (CSS vs.)
	fs := http.FileServer(http.FS(static))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("🚀 Sunucu çalışıyor: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
``
/*
---

# 📌 Kod: `templates/index.html`
*/

``html
<!DOCTYPE html>
<html lang="tr">
<head>
    <meta charset="UTF-8">
    <title>Embed Örneği</title>
    <link rel="stylesheet" href="/static/style.css">
</head>
<body>
    <h1>Merhaba Dünya!</h1>
    <p>Bu sayfa <b>embed paketi</b> ile binary içine gömüldü 🎉</p>
</body>
</html>
``
/*
---

# 📌 Kod: `static/style.css`
*/
``css
body {
    font-family: Arial, sans-serif;
    background-color: #f0f4f8;
    text-align: center;
    margin-top: 50px;
}
h1 {
    color: #007acc;
}
p {
    color: #333;
}
``
/*
---

# 📌 Çalıştırma

Terminalde:

```bash
go run main.go
```

veya binary almak için:

```bash
go build -o myapp
./myapp
```

👉 Tarayıcıda [http://localhost:8080](http://localhost:8080) açtığında HTML + CSS gömülü şekilde gelir.
Artık dışarıda dosya taşıma derdi yok, tek bir binary yeterli ✅

---

İstersen ben sana bu örneği biraz daha geliştirip, içine **birden fazla HTML sayfası** ekleyeyim (örneğin `/about` sayfası). İster misin?
EVET
*/
/*
Harika 👌 O zaman `embed` ile **birden fazla HTML sayfası** ve CSS kullanabileceğimiz küçük bir web uygulaması yapalım.

Bu uygulamada:

* `index.html` → Anasayfa
* `about.html` → Hakkında sayfası
* `style.css` → Ortak stil dosyası
  hepsi binary içine gömülecek.

---

# 📌 Proje Yapısı

```
myapp/
 ├─ main.go
 ├─ templates/
 │   ├─ index.html
 │   └─ about.html
 └─ static/
     └─ style.css
```
*/
/*
---

# 📌 Kod: `main.go`
*/
``go
package main

import (
	"embed"
	"fmt"
	"net/http"
)

// HTML şablonları
//go:embed templates/*
var templates embed.FS

// Statik dosyalar (CSS, JS, resimler)
//go:embed static/*
var static embed.FS

// Yardımcı fonksiyon: HTML dosyası oku ve cevap döndür
func renderHTML(w http.ResponseWriter, filename string) {
	data, err := templates.ReadFile("templates/" + filename)
	if err != nil {
		http.Error(w, "Sayfa bulunamadı", 404)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(data)
}

func main() {
	// Anasayfa
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderHTML(w, "index.html")
	})

	// Hakkında sayfası
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		renderHTML(w, "about.html")
	})

	// Statik dosyalar (CSS vs.)
	fs := http.FileServer(http.FS(static))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("🚀 Sunucu çalışıyor: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
``
/*
---

# 📌 Kod: `templates/index.html`
*/
``html
<!DOCTYPE html>
<html lang="tr">
<head>
    <meta charset="UTF-8">
    <title>Ana Sayfa</title>
    <link rel="stylesheet" href="/static/style.css">
</head>
<body>
    <h1>Merhaba Dünya!</h1>
    <p>Bu sayfa <b>embed paketi</b> ile binary içine gömüldü 🎉</p>
    <a href="/about">Hakkında</a>
</body>
</html>
``
/*
---

# 📌 Kod: `templates/about.html`
*/
``html
<!DOCTYPE html>
<html lang="tr">
<head>
    <meta charset="UTF-8">
    <title>Hakkında</title>
    <link rel="stylesheet" href="/static/style.css">
</head>
<body>
    <h1>Hakkında</h1>
    <p>Bu küçük web uygulaması Go'nun <code>embed</code> paketi ile yazıldı.</p>
    <a href="/">Ana Sayfa</a>
</body>
</html>
``
/*
---

# 📌 Kod: `static/style.css`
*/

``css
body {
    font-family: Arial, sans-serif;
    background-color: #f0f4f8;
    text-align: center;
    margin-top: 50px;
}
h1 {
    color: #007acc;
}
p {
    color: #333;
}
a {
    display: inline-block;
    margin-top: 20px;
    padding: 10px 20px;
    background: #007acc;
    color: white;
    text-decoration: none;
    border-radius: 8px;
}
a:hover {
    background: #005f99;
}
``
/*
---

# 📌 Çalıştırma

```bash
go run main.go
```

👉 Tarayıcıda aç:

* [http://localhost:8080](http://localhost:8080) → Ana Sayfa
* [http://localhost:8080/about](http://localhost:8080/about) → Hakkında

Tüm HTML ve CSS dosyaları binary’ye gömülü 💯

---

👉 İstersen ben bu yapıyı bir adım daha ileri götürüp, `html/template` paketini de işin içine katayım. Böylece sayfalara **dinamik veri** gönderebilirsin. Bunu da ister misin?
*/

