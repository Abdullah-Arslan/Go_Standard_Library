/*
👍 **Go’nun `encoding/json` paketini** en ince ayrıntısına kadar, örneklerle birlikte anlatayım.

---

# 📦 Go `encoding/json` Paketi

Go’nun `encoding/json` paketi, **JSON (JavaScript Object Notation)** formatında veriyi **encode (Go → JSON)** ve **decode (JSON → Go)** etmemizi sağlar.

* **JSON nedir?**

  * Hafif bir veri formatıdır.
  * İnsan tarafından okunabilir.
  * Dilde bağımsızdır (Go, Python, Java, JS her yerde).
* **Kullanım Alanı:**

  * Web API’leri (REST, GraphQL, gRPC JSON).
  * Konfigürasyon dosyaları.
  * Veri paylaşımı (Go ↔ başka diller).

---

## 🔑 Temel Fonksiyonlar

* `json.Marshal(v interface{}) ([]byte, error)` → Go verisini JSON’a çevirir.
* `json.MarshalIndent(v, prefix, indent string) ([]byte, error)` → Daha okunabilir JSON üretir (indentli).
* `json.Unmarshal(data []byte, v interface{}) error` → JSON verisini Go nesnesine çevirir.
* `json.NewEncoder(w io.Writer).Encode(v)` → JSON’u direkt stream yazar (ör. dosya, network).
* `json.NewDecoder(r io.Reader).Decode(v)` → Stream’den JSON okur.

---

## 📌 Basit Örnek (Struct → JSON ve JSON → Struct)
*/
``go
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Struct → JSON
	p := Person{Name: "Ahmet", Age: 30}
	jsonData, _ := json.Marshal(p)
	fmt.Println("JSON:", string(jsonData))

	// JSON → Struct
	var p2 Person
	json.Unmarshal(jsonData, &p2)
	fmt.Println("Struct:", p2)
}
``
/*
🔹 Çıktı:

```
JSON: {"Name":"Ahmet","Age":30}
Struct: {Ahmet 30}
```

---

## 📌 `MarshalIndent` (Güzelleştirilmiş JSON)
*/
``go
package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name  string
	Price float64
}

func main() {
	p := Product{Name: "Laptop", Price: 12999.90}
	jsonData, _ := json.MarshalIndent(p, "", "  ")
	fmt.Println(string(jsonData))
}
``

//🔹 Çıktı:

``json
{
  "Name": "Laptop",
  "Price": 12999.9
}
``
/*
---

## 📌 Struct Tag Kullanımı (`json:"..."`)

Struct alanlarının JSON karşılığını **etiketlerle (`tag`)** değiştirebiliriz.
*/
``go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"-"`           // JSON’a dahil edilmez
	Email    string `json:"email,omitempty"` // boşsa dahil edilmez
}

func main() {
	u := User{Username: "admin", Password: "12345"}

	data, _ := json.MarshalIndent(u, "", "  ")
	fmt.Println(string(data))
}
``

//🔹 Çıktı:

``json
{
  "username": "admin"
}
``
/*
---

## 📌 Slice ve Map ile JSON
*/
``go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Slice → JSON
	langs := []string{"Go", "Python", "JavaScript"}
	data, _ := json.Marshal(langs)
	fmt.Println(string(data)) // ["Go","Python","JavaScript"]

	// Map → JSON
	m := map[string]int{"elma": 5, "armut": 10}
	data2, _ := json.Marshal(m)
	fmt.Println(string(data2)) // {"armut":10,"elma":5}
}
``
/*
---

## 📌 `interface{}` ile Dinamik JSON Okuma

JSON’un yapısı önceden bilinmiyorsa `map[string]interface{}` veya `interface{}` ile çözülebilir.
*/

``go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `{"username":"ali","age":25,"active":true}`

	var result map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &result)

	fmt.Println("Kullanıcı adı:", result["username"])
	fmt.Println("Yaş:", result["age"])
	fmt.Println("Aktif mi:", result["active"])
}
``
/*
---

## 📌 JSON Encode/Decode (Stream - Dosya veya Network)

### JSON’u Dosyaya Yazmak
*/

``go
package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Mehmet", Age: 40}
	file, _ := os.Create("person.json")
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // okunabilir olsun
	encoder.Encode(p)
}
``
/*
### Dosyadan JSON Okumak
*/
``go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	file, _ := os.Open("person.json")
	defer file.Close()

	var p Person
	decoder := json.NewDecoder(file)
	decoder.Decode(&p)

	fmt.Println("Dosyadan okunan:", p)
}
``
/*
---

## 📌 `RawMessage` Kullanımı

Bazen JSON’un bir kısmını ham string olarak saklamak isteyebiliriz. Bunun için `json.RawMessage` kullanılır.
*/
``go
package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Status string
	Data   json.RawMessage
}

func main() {
	jsonStr := `{"Status":"ok","Data":{"Name":"Ayşe","Age":28}}`

	var r Response
	json.Unmarshal([]byte(jsonStr), &r)

	fmt.Println("Status:", r.Status)
	fmt.Println("Ham JSON Data:", string(r.Data))
}
``
/*
---

## 📌 Özet

* `encoding/json` paketi JSON verisiyle çalışmak için kullanılır.
* 🔄 **Marshal / Unmarshal**: Go <-> JSON dönüşümleri.
* 🎯 Struct `tag` → JSON alan isimlerini kontrol eder.
* 📂 Encoder / Decoder → dosya ve network üzerinden stream JSON işleme.
* 🔧 `RawMessage` ve `interface{}` → dinamik JSON işlemleri.

---

👉 İstersen ben sana bir de **REST API’den JSON veri çekme ve struct’a map etme** örneği gösterebilirim. Bu, gerçek hayatta en sık kullanılan senaryodur. Onu da ekleyeyim mi?
EVET
*/
/*
Harika 👍 Şimdi sana **REST API’den JSON veri çekme ve Go struct’a map etme** örneğini göstereyim.
Bu, gerçek hayatta `encoding/json` paketinin **en sık kullanıldığı senaryodur** 🚀

---

# 📌 REST API’den JSON Çekme ve Struct’a Map Etme

Dış bir API’den JSON verisi çekip Go içinde kullanacağız.
Bunun için:

* `net/http` → API’den veri çekmek için,
* `encoding/json` → JSON → Go struct dönüşümü için.

---

## 🌍 Örnek API (JSON Placeholder)

JSONPlaceholder sahte REST API’sini kullanacağız:
👉 `https://jsonplaceholder.typicode.com/users/1`

Dönen JSON örneği:
*/
``json
{
  "id": 1,
  "name": "Leanne Graham",
  "username": "Bret",
  "email": "Sincere@april.biz"
}
``
/*
---

## 📝 Go Kodu
*/
``go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	// 1️⃣ API isteği
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 2️⃣ Body oku
	body, _ := ioutil.ReadAll(resp.Body)

	// 3️⃣ JSON → Struct
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		panic(err)
	}

	// 4️⃣ Sonuç
	fmt.Println("Kullanıcı ID   :", user.ID)
	fmt.Println("Adı            :", user.Name)
	fmt.Println("Kullanıcı Adı  :", user.Username)
	fmt.Println("E-posta        :", user.Email)
}
``
/*
---

## 🔹 Çalıştırınca Çıktı

```
Kullanıcı ID   : 1
Adı            : Leanne Graham
Kullanıcı Adı  : Bret
E-posta        : Sincere@april.biz
```

---

## 📌 Notlar

* Struct içindeki `json:"alan"` tag’leri JSON’daki alan isimlerini eşleştirmek için **zorunludur** (aksi halde büyük/küçük harf farklılıklarından dolayı eşleşmeyebilir).
* Eğer JSON içinde bazı alanları kullanmak istemiyorsak struct’a eklemeyebiliriz (JSON parser otomatik atlar).
* `http.Get` yerine gerçek uygulamalarda **`http.Client`** kullanmak daha profesyoneldir (timeout, retry gibi özelliklerle).

---

👉 Şimdi sana sorayım: İstersen ben bir adım öteye geçip **REST API’den birden fazla kullanıcı listesi çekip slice’a (`[]User`) map etme** örneğini de gösterebilirim. Bunu ister misin?
EVET
*/

/*
Süper 👍 Şimdi REST API’den **birden fazla kullanıcıyı JSON olarak çekip `[]User` slice’ına map etme** örneğini göstereyim.

---

# 📌 REST API’den JSON Liste Çekme (`[]User`)

API:
👉 `https://jsonplaceholder.typicode.com/users`

Bu endpoint bize bir **kullanıcı listesi** döner (10 kullanıcı).

Örnek JSON (ilk 2 kullanıcı):
*/
``json
[
  {
    "id": 1,
    "name": "Leanne Graham",
    "username": "Bret",
    "email": "Sincere@april.biz"
  },
  {
    "id": 2,
    "name": "Ervin Howell",
    "username": "Antonette",
    "email": "Shanna@melissa.tv"
  }
]
``
/*
---

## 📝 Go Kodu
*/
``go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	// 1️⃣ API isteği
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 2️⃣ Body oku
	body, _ := ioutil.ReadAll(resp.Body)

	// 3️⃣ JSON → []User
	var users []User
	err = json.Unmarshal(body, &users)
	if err != nil {
		panic(err)
	}

	// 4️⃣ Sonuçları yazdır
	for _, u := range users {
		fmt.Printf("ID: %d | Ad: %s | Kullanıcı: %s | Email: %s\n", u.ID, u.Name, u.Username, u.Email)
	}
}
``
/*
---

## 🔹 Çalıştırınca Çıktı (kısaltılmış)

```
ID: 1 | Ad: Leanne Graham | Kullanıcı: Bret | Email: Sincere@april.biz
ID: 2 | Ad: Ervin Howell | Kullanıcı: Antonette | Email: Shanna@melissa.tv
ID: 3 | Ad: Clementine Bauch | Kullanıcı: Samantha | Email: Nathan@yesenia.net
...
```

---

## 📌 Özet

* Tek bir obje için `var user User` → `json.Unmarshal`.
* Liste (array) için `var users []User` → `json.Unmarshal`.
* `[]User` slice’ına direkt JSON array’i map edebiliriz.

---

👉 İstersen buradan bir adım öteye gidip **Go ile küçük bir REST API server yazıp JSON response döndürme** örneğini de gösterebilirim. Bu sefer client değil, **server tarafı** örneği olur. Onu da ister misin?
*/