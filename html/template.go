/*
Go dilinde **`text/template` ve `html/template`** paketlerini detaylıca inceleyelim.

---

# 📦 `template` Paketi – Genel Bakış

Go’da **`template`** paketleri, dinamik içerik üretmek için kullanılır.
Temel olarak **şablon (template) + veri = sonuç metin** mantığıyla çalışır.

* `text/template` → Saf metin (text) üretmek için.
* `html/template` → HTML üretmek için, XSS saldırılarına karşı otomatik olarak güvenlik sağlar.

---

# 📌 Paket İçe Aktarma
*/
``go
import (
	"text/template"   // saf metin
	"html/template"   // HTML şablonları
)
``
/*
---

# 🔑 Temel Kavramlar

1. **Template** → Şablon, içinde {{.Field}} veya {{function}} gibi ifadeler içerir.
2. **Data** → Template’e gönderilen veri. Struct, map veya basit değişken olabilir.
3. **Execute** → Template + Data → Sonuç metin üretir.

---

# 🔧 Önemli Fonksiyonlar

### 1. `template.New(name string) *Template`

Yeni bir template oluşturur.
*/

``go
t := template.New("example")
``
/*
---

### 2. `template.Parse` / `template.ParseFiles` / `template.ParseGlob`

* Template içeriğini yükler veya dosyadan okur.
*/
``go
t, _ := template.New("example").Parse("Merhaba {{.Name}}!")
``
/*
---

### 3. `template.Execute(wr io.Writer, data interface{})`

* Template’i verilen veri ile render eder.
*/
``go
type User struct {
	Name string
}

func main() {
	t, _ := template.New("example").Parse("Merhaba {{.Name}}!")
	user := User{Name: "Ahmet"}
	t.Execute(os.Stdout, user)
}
``

//📌 Çıktı:

``
Merhaba Ahmet!
``
/*
---

# 📂 Örnek 1: Basit Text Template
*/
``go
package main

import (
	"os"
	"text/template"
)

type Product struct {
	Name  string
	Price float64
}

func main() {
	tmpl := `Ürün: {{.Name}}
Fiyat: ${{.Price}}`

	p := Product{"Laptop", 1250.50}
	t := template.Must(template.New("product").Parse(tmpl))
	t.Execute(os.Stdout, p)
}
``

//📌 Çıktı:

``
Ürün: Laptop
Fiyat: $1250.5
``
/*
---

# 📂 Örnek 2: HTML Template (`html/template`)
*/
``go
package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Name string
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := `<h1>Merhaba {{.Name}}</h1>`
	t := template.Must(template.New("index").Parse(tmpl))
	t.Execute(w, User{Name: "<script>alert('XSS')</script>"})
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
``

//📌 Çıktı:

``
<h1>Merhaba &lt;script&gt;alert('XSS')&lt;/script&gt;</h1>
``
/*
✅ HTML template otomatik olarak XSS’e karşı güvenlik sağlar.

---

# 🔧 Template Fonksiyonları

1. **Pipeline Kullanımı**
*/
``go
{{.Field | printf "%.2f"}}
``

//2. **If, Range, With**

``go
{{if .Show}}Göster{{end}}
{{range .Items}}{{.}}{{end}}
{{with .User}}Hoşgeldin, {{.Name}}{{end}}
``

//3. **Custom Functions**

``go
t := template.New("test").Funcs(template.FuncMap{
	"uppercase": strings.ToUpper,
})
``
/*
---

# 📌 Özet

* **`text/template`** → Saf metin üretmek için.
* **`html/template`** → HTML üretmek için, XSS güvenliği sağlar.
* **Temel Akış:**

  1. Template oluştur
  2. Template’i parse et
  3. Data ile execute et
* **Kontrol Yapıları:** `if`, `range`, `with`
* **Fonksiyon Desteği:** Pipeline ve custom functions

---

👉 İstersen ben sana **küçük bir web uygulaması örneği hazırlayayım**,

* Bir template dosyası kullanacak
* Dinamik veri gösterecek
* HTML güvenliği sağlayacak
  Böylece `html/template` kullanımını gerçek bir senaryoda görebilirsin.

Bunu yapayım mı?
EVET
*/
/*
Harika! 🚀 O zaman sana **Go ile küçük bir web uygulaması** hazırlayayım.
Bu uygulama, bir template dosyası kullanarak **dinamik HTML** gösterecek ve **XSS saldırılarına karşı güvenli olacak**.

---

## 📂 Proje: `web_template_example`

### 1️⃣ Dosya Yapısı
*/
``
web_template_example/
│
├── main.go
└── templates/
    └── index.html
``
/*
---

### 2️⃣ `templates/index.html`
*/
``html
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Hoşgeldin</title>
</head>
<body>
    <h1>Merhaba, {{.Name}}!</h1>

    <p>Mesaj: {{.Message}}</p>
</body>
</html>
``
/*
> Dikkat: Kullanıcı inputu `html/template` ile güvenli hale gelir.

---

### 3️⃣ `main.go`
*/
``go
package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Name    string
	Message string
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	data := PageData{
		Name:    r.URL.Query().Get("name"),
		Message: r.URL.Query().Get("message"),
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Template render hatası", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)

	log.Println("Sunucu çalışıyor: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
``
/*
---

## 📌 Kullanım

1. Projeyi çalıştır:
*/
``bash
go run main.go
``

//2. Tarayıcıdan şu URL’e git:

``
http://localhost:8080/?name=Ahmet&message=<script>alert('XSS')</script>
``

//📌 Çıktı:

``
Merhaba, Ahmet!
Mesaj: &lt;script&gt;alert('XSS')&lt;/script&gt;
``
/*
✅ Script çalışmaz, HTML template otomatik olarak escape eder.

---

### 🔑 Özet

* `html/template` paketi XSS güvenliği sağlar.
* Template dosyası ile **ayrılmış HTML** ve Go kodu oluşturmak kolaydır.
* Query parametreleri veya form inputları güvenli bir şekilde gösterilebilir.
* Dinamik veri göstermek için struct veya map kullanılabilir.

---

İstersen bunu bir adım daha ileri götürüp **form üzerinden kullanıcı inputu alıp liste halinde gösterecek bir uygulama** da yapabiliriz.
Böylece gelen tüm inputlar güvenli bir şekilde saklanır ve gösterilir.

Bunu yapayım mı?
EVET
*/
/*
Harika! 😄 O zaman şimdi sana **form üzerinden kullanıcı inputu alıp liste halinde gösterecek güvenli bir web uygulaması** hazırlayacağım.

Bu uygulama:

* Kullanıcıdan isim ve mesaj alacak
* Verileri güvenli bir şekilde listeleyecek (XSS korumalı)
* HTML template kullanacak

---

## 📂 Proje: `safe_message_board`

### 1️⃣ Dosya Yapısı
*/
``
safe_message_board/
│
├── main.go
└── templates/
    └── index.html
``
/*
---

### 2️⃣ `templates/index.html`
*/
``html
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Mesaj Panosu</title>
</head>
<body>
    <h1>Mesaj Panosu</h1>

    <form action="/" method="post">
        <input type="text" name="name" placeholder="İsim" required>
        <input type="text" name="message" placeholder="Mesaj" required>
        <input type="submit" value="Gönder">
    </form>

    <h2>Gönderilen Mesajlar:</h2>
    <ul>
    {{range .Messages}}
        <li><strong>{{.Name}}:</strong> {{.Message}}</li>
    {{else}}
        <li>Henüz mesaj yok.</li>
    {{end}}
    </ul>
</body>
</html>
``
/*
> Not: `html/template` kullanıldığı için tüm inputlar otomatik olarak escape edilir.

---

### 3️⃣ `main.go`
*/
``go
package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"
)

type Message struct {
	Name    string
	Message string
}

type PageData struct {
	Messages []Message
}

var (
	messages []Message
	mu       sync.Mutex
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err == nil {
			name := r.FormValue("name")
			message := r.FormValue("message")

			mu.Lock()
			messages = append(messages, Message{Name: name, Message: message})
			mu.Unlock()
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	data := PageData{Messages: messages}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Template render hatası", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Sunucu çalışıyor: http://localhost:8080")
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

3. Formdan birkaç mesaj gönder:
*/
``
İsim: Ahmet
Mesaj: Merhaba, <script>alert('XSS')</script>
``

//4. Gönderilen mesajlar güvenli bir şekilde listelenir:

``
Ahmet: Merhaba, &lt;script&gt;alert('XSS')&lt;/script&gt;
``
/*
✅ Script çalışmaz, XSS önlenir.

---

### 🔑 Özellikler

* Kullanıcı inputu **otomatik escape edilir** (`html/template`).
* **Mutex (`sync.Mutex`)** kullanılarak eşzamanlı erişim güvenliği sağlanır.
* Tüm mesajlar hafızada saklanır ve sayfa yenilendiğinde gösterilir.
* Basit ama gerçek bir **güvenli mesaj panosu** uygulaması.

---

İstersen bir adım daha ileri götürüp:

* Mesajları **JSON veya CSV dosyasına kaydederek kalıcı hâle getirebiliriz**.
* Sunucu yeniden başlatıldığında veriler kaybolmaz.

Bunu da yapalım mı?
EVET
*/

/*
Harika! 🚀 O zaman mevcut mesaj panosunu **kalıcı hâle getirecek şekilde geliştirelim**.
Mesajlar artık **JSON dosyasına kaydedilecek ve sunucu yeniden başlatıldığında geri yüklenecek**.

---

## 📂 Proje: `persistent_message_board`

### 1️⃣ Dosya Yapısı

```
persistent_message_board/
│
├── main.go
└── templates/
    └── index.html
```

> Template aynı: `index.html` değişmeyecek.

---

### 2️⃣ `main.go`
*/
``go
package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"sync"
)

type Message struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type PageData struct {
	Messages []Message
}

var (
	messages []Message
	mu       sync.Mutex
	dataFile = "messages.json"
)

// Mesajları JSON dosyasından yükle
func loadMessages() {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			messages = []Message{}
			return
		}
		log.Fatal("Dosya açma hatası:", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&messages); err != nil {
		log.Println("JSON decode hatası:", err)
		messages = []Message{}
	}
}

// Mesajları JSON dosyasına kaydet
func saveMessages() {
	file, err := os.Create(dataFile)
	if err != nil {
		log.Println("JSON kaydetme hatası:", err)
		return
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(messages); err != nil {
		log.Println("JSON encode hatası:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err == nil {
			name := r.FormValue("name")
			message := r.FormValue("message")

			mu.Lock()
			messages = append(messages, Message{Name: name, Message: message})
			saveMessages()
			mu.Unlock()
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := PageData{Messages: messages}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Template render hatası", http.StatusInternalServerError)
	}
}

func main() {
	loadMessages()

	http.HandleFunc("/", handler)
	log.Println("Sunucu çalışıyor: http://localhost:8080")
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
2. Tarayıcıdan `http://localhost:8080` adresine git ve mesaj gönder.

3. Sunucuyu kapatıp tekrar başlatsan bile mesajlar **JSON dosyasından yüklenecek**.

📌 Örnek `messages.json`:
*/

``json
[
  {"name":"Ahmet","message":"Merhaba &lt;script&gt;alert('XSS')&lt;/script&gt;"},
  {"name":"Ayşe","message":"Selam Dünya!"}
]
``

/*
---

### 🔑 Özellikler

* Kullanıcı inputları **HTML güvenli** (`html/template`)
* **Kalıcı mesajlar:** JSON dosyasına kaydediliyor
* **Eşzamanlı erişim güvenliği:** `sync.Mutex`
* Basit ama **gerçek bir kalıcı güvenli mesaj panosu**

---

İstersen bir sonraki adımda bunu **FNV veya maphash ile mesajlara hızlı hash ekleyip, bozulma kontrolü** yapacak hâle de getirebiliriz. Böylece mesajların değiştirilip değiştirilmediğini kontrol edebiliriz.

Bunu yapalım mı?
EVET
*/

/*
Harika! 😄 O zaman mevcut **kalıcı mesaj panosunu** bir adım daha ileri götürelim ve **mesajlara FNV veya maphash hash** ekleyelim.

Bu sayede:

* Her mesajın **benzersiz hash değeri** olacak
* Mesajlar üzerinde değişiklik olup olmadığını **hash ile kontrol edebileceğiz**

---

## 📂 Proje: `hashed_message_board`

### 1️⃣ Dosya Yapısı

```
hashed_message_board/
│
├── main.go
└── templates/
    └── index.html
```

> Template’i hafifçe hash sütunu ekleyerek güncelleyelim:

### `templates/index.html`
*/
``html
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Mesaj Panosu (Hashli)</title>
</head>
<body>
    <h1>Mesaj Panosu</h1>

    <form action="/" method="post">
        <input type="text" name="name" placeholder="İsim" required>
        <input type="text" name="message" placeholder="Mesaj" required>
        <input type="submit" value="Gönder">
    </form>

    <h2>Gönderilen Mesajlar:</h2>
    <table border="1" cellpadding="5">
        <tr>
            <th>İsim</th>
            <th>Mesaj</th>
            <th>Hash (FNV-1a 64-bit)</th>
        </tr>
        {{range .Messages}}
        <tr>
            <td>{{.Name}}</td>
            <td>{{.Message}}</td>
            <td>{{.Hash}}</td>
        </tr>
        {{else}}
        <tr><td colspan="3">Henüz mesaj yok.</td></tr>
        {{end}}
    </table>
</body>
</html>
``
/*
---

### 2️⃣ `main.go`
*/
``go
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"hash/fnv"
	"log"
	"net/http"
	"os"
	"sync"
)

type Message struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Hash    string `json:"hash"`
}

type PageData struct {
	Messages []Message
}

var (
	messages []Message
	mu       sync.Mutex
	dataFile = "messages.json"
)

// FNV-1a 64-bit hash hesapla
func computeHash(name, message string) string {
	h := fnv.New64a()
	h.Write([]byte(name + ":" + message))
	return fmt.Sprintf("%x", h.Sum64())
}

// Mesajları JSON dosyasından yükle
func loadMessages() {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			messages = []Message{}
			return
		}
		log.Fatal("Dosya açma hatası:", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&messages); err != nil {
		log.Println("JSON decode hatası:", err)
		messages = []Message{}
	}
}

// Mesajları JSON dosyasına kaydet
func saveMessages() {
	file, err := os.Create(dataFile)
	if err != nil {
		log.Println("JSON kaydetme hatası:", err)
		return
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(messages); err != nil {
		log.Println("JSON encode hatası:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err == nil {
			name := r.FormValue("name")
			message := r.FormValue("message")

			hash := computeHash(name, message)

			mu.Lock()
			messages = append(messages, Message{Name: name, Message: message, Hash: hash})
			saveMessages()
			mu.Unlock()
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := PageData{Messages: messages}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Template render hatası", http.StatusInternalServerError)
	}
}

func main() {
	loadMessages()

	http.HandleFunc("/", handler)
	log.Println("Sunucu çalışıyor: http://localhost:8080")
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
2. Tarayıcıdan `http://localhost:8080` adresine git ve mesaj gönder.

3. Her mesajın yanında **benzersiz FNV-1a hash** görünecek.

📌 Örnek `messages.json`:
?7

``json
[
  {"name":"Ahmet","message":"Merhaba!","hash":"6b8b4567327b23c6"},
  {"name":"Ayşe","message":"Selam Dünya","hash":"1a2b3c4d5e6f7890"}
]
``
/*
---

### 🔑 Özellikler

* Mesajlar **HTML güvenli** (`html/template`)
* Mesajlar **JSON dosyasında kalıcı**
* Her mesajın **FNV-1a 64-bit hash’i** var → veri bütünlüğü kontrolü yapılabilir
* Basit ve güvenli bir **hash destekli mesaj panosu**

---

İstersen bunu bir adım daha ileri götürüp, **maphash kullanarak da hash hesaplayan ve seed ile güvenliği artıran versiyon** hazırlayabilirim.
Böylece hash’ler saldırılara karşı daha güvenli olur.

Bunu da yapalım mı?
*/