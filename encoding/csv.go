/*
Go’nun standart kütüphanesinde bulunan **`encoding/csv`** paketi, **CSV (Comma-Separated Values)** dosyalarını **okumak ve yazmak** için kullanılır.

CSV dosyaları genellikle:

* Excel çıktıları
* Veritabanı exportları
* Veri analizi için kullanılan datasetler
* Konfigürasyon ve basit tablolama işlerinde

kullanılır.

---

# 📦 `encoding/csv` Paketi

Paketi şu şekilde import ediyoruz:

```go
import "encoding/csv"
```

CSV formatı:

* Satırlar `\n` ile ayrılır.
* Kolonlar **virgül (`,`)** ile ayrılır. (Ama istenirse değiştirilebilir, örn. `;` veya `\t`)
* Hücrelerde özel karakterler varsa (`,` veya `\n`), hücreler **tırnak içinde (`"`)** yazılır.

---

# 🔹 1. CSV Dosyası Okuma

Diyelim ki `data.csv` dosyamız şöyle:

```csv
id,name,age
1,Ahmet,25
2,Mehmet,30
3,Ayşe,22
```

👉 Kod:
*/
``go
package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Tüm satırları oku
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// Ekrana yazdır
	for i, row := range records {
		fmt.Printf("Satır %d: %v\n", i, row)
	}
}
``
/*
✅ Çıktı:

```
Satır 0: [id name age]
Satır 1: [1 Ahmet 25]
Satır 2: [2 Mehmet 30]
Satır 3: [3 Ayşe 22]
```

---

# 🔹 2. Satır Satır CSV Okuma (Büyük Dosyalar İçin)

Büyük CSV dosyalarında `ReadAll` kullanmak hafızayı şişirir. Bunun yerine satır satır `Read()` kullanabiliriz.
*/

``go
reader := csv.NewReader(file)
for {
	record, err := reader.Read()
	if err != nil {
		break
	}
	fmt.Println(record)
}
``
/*
---

# 🔹 3. CSV Dosyası Yazma
*/

``go
package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("output.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush() // buffer’daki verileri dosyaya yaz

	data := [][]string{
		{"id", "name", "age"},
		{"1", "Ahmet", "25"},
		{"2", "Mehmet", "30"},
		{"3", "Ayşe", "22"},
	}

	for _, row := range data {
		if err := writer.Write(row); err != nil {
			panic(err)
		}
	}
}
``
/*
📂 `output.csv` çıktısı:

```csv
id,name,age
1,Ahmet,25
2,Mehmet,30
3,Ayşe,22
```

---

# 🔹 4. Ayracı Değiştirmek (Örn. `;` ile ayırmak)

Bazı CSV dosyalarında ayırıcı `,` değil `;` olabilir. Bu durumda:
*/

``go
reader := csv.NewReader(file)
reader.Comma = ';'
``

//Aynı şekilde yazarken de:

``go
writer := csv.NewWriter(file)
writer.Comma = ';'
``
/*
---

# 🔹 5. Özel Karakterler & Tırnak Kullanımı
*/
``go
data := [][]string{
	{"id", "name", "comment"},
	{"1", "Ahmet", "Merhaba, nasılsın?"},
	{"2", "Mehmet", "Go \"harika\" bir dil!"},
}
``
/*
Yazıldığında `comment` kolonundaki virgül ve çift tırnak otomatik olarak korunur:

```csv
id,name,comment
1,Ahmet,"Merhaba, nasılsın?"
2,Mehmet,"Go ""harika"" bir dil!"
```

---

# 🔹 6. Struct ile CSV Kullanımı

Go’nun `encoding/csv` paketi doğrudan struct desteği sunmaz. Ama kolayca kendimiz mapleyebiliriz:
*/
``go
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	file, _ := os.Open("data.csv")
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	var users []User
	for i, row := range records {
		if i == 0 { // başlık satırını atla
			continue
		}
		id, _ := strconv.Atoi(row[0])
		age, _ := strconv.Atoi(row[2])
		users = append(users, User{
			ID:   id,
			Name: row[1],
			Age:  age,
		})
	}

	for _, u := range users {
		fmt.Printf("%+v\n", u)
	}
}
``
/*
✅ Çıktı:

```
{ID:1 Name:Ahmet Age:25}
{ID:2 Name:Mehmet Age:30}
{ID:3 Name:Ayşe Age:22}
```

---

# 📌 Özet

* `csv.NewReader` → CSV dosyasını okur.

  * `ReadAll()` → hepsini oku
  * `Read()` → satır satır oku
* `csv.NewWriter` → CSV dosyasına yaz.
* `Comma` → ayırıcıyı değiştirme (`;`, `\t`, vb.)
* `Flush()` → yazarken veriyi dosyaya aktarmayı unutma.
* Struct ile çalışmak için `strconv` gibi paketlerle dönüştürmek gerekir.

---

👉 İstersen sana **CSV → JSON dönüştürücü** küçük bir program da yazabilirim (mesela CSV verilerini alıp JSON’a çeviren). Bunu ister misin?
EVET
*/
/*
Harika 👍 O zaman sana **CSV → JSON dönüştürücü** küçük bir Go programı göstereyim.

---

## 📂 Örnek CSV (`data.csv`)

```csv
id,name,age
1,Ahmet,25
2,Mehmet,30
3,Ayşe,22
```

---

## 💻 Go Programı: CSV → JSON
*/
``go
package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// CSV dosyasını aç
	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// CSV okuyucu
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var users []User

	// İlk satır başlık, o yüzden 1’den başlıyoruz
	for i, row := range records {
		if i == 0 {
			continue
		}
		id, _ := strconv.Atoi(row[0])
		age, _ := strconv.Atoi(row[2])
		users = append(users, User{
			ID:   id,
			Name: row[1],
			Age:  age,
		})
	}

	// JSON’a çevir
	jsonData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		panic(err)
	}

	// JSON çıktısını dosyaya yaz
	if err := os.WriteFile("output.json", jsonData, 0644); err != nil {
		panic(err)
	}

	// Konsola da yazdır
	fmt.Println(string(jsonData))
}
``
/*
---

## ✅ Çıktı (`output.json`)
*/
``json
[
  {
    "id": 1,
    "name": "Ahmet",
    "age": 25
  },
  {
    "id": 2,
    "name": "Mehmet",
    "age": 30
  },
  {
    "id": 3,
    "name": "Ayşe",
    "age": 22
  }
]
``
/*
---

Böylece CSV verilerini JSON’a çevirmiş olduk 🚀
Bu programı biraz geliştirip **API’ye bağlayıp CSV yükleyeni JSON olarak dönen bir servis** de yapabiliriz.

👉 İstersen sana bir **mini REST API** örneği göstereyim: CSV dosyasını upload edip JSON response dönen. Bunu ister misin?
*/