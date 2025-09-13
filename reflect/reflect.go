/*
👍 Go dilinde en güçlü ama en çok kafa karıştıran paketlerden biri **`reflect`** paketidir.
Sana şimdi bu paketi en baştan **tüm önemli fonksiyonları ve tipleriyle**, bol **örnek kodlarla** anlatacağım.

---

# 📦 Go `reflect` Paketi

## 1. Amaç

`reflect` paketi, **runtime reflection** (çalışma zamanı yansıma) sağlar.
Yani:

* Bir değişkenin **tipini** (type) ve **değerini** (value) runtime’da öğrenebilirsin.
* Struct alanlarına (fields) ve metodlarına erişebilirsin.
* Yeni değerler set edebilirsin (settable).
* Dinamik fonksiyon çağırabilirsin.

Kısaca: `reflect` → **derleme zamanında bilinmeyen şeyleri çalışma zamanında keşfetme**.

---

## 2. Temel Tipler

* **`reflect.Type`** → Bir değerin tipini temsil eder.
* **`reflect.Value`** → Bir değerin kendisini runtime’da temsil eder.

Bunları elde etmek için:

* `reflect.TypeOf(x)` → Tipini alır.
* `reflect.ValueOf(x)` → Değerini alır.

---

## 3. Basit Örnek
*/
``go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 42
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println("Tip:", t)           // int
	fmt.Println("Kind:", t.Kind())   // int
	fmt.Println("Değer:", v.Int())   // 42
}
``

//📌 Çıktı:

``
Tip: int
Kind: int
Değer: 42
``
/*
---

## 4. Type ve Kind Farkı

* **Type** → `int`, `string`, `[]int`, `map[string]int`, `MyStruct` gibi gerçek tip.
* **Kind** → daha genel kategori (`int`, `slice`, `map`, `struct`).

Örnek:
*/
``go
var y = []string{"a", "b"}
t := reflect.TypeOf(y)
fmt.Println(t)        // []string
fmt.Println(t.Kind()) // slice
``
/*
---

## 5. Struct Alanlarına Erişim
*/
``go
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{"Abdullah", 25}

	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)  // Alan bilgisi
		value := v.Field(i)  // Alan değeri
		fmt.Printf("Alan: %s, Tip: %s, Değer: %v\n", field.Name, field.Type, value)
	}
}
``

//📌 Çıktı:

``
Alan: Name, Tip: string, Değer: Abdullah
Alan: Age, Tip: int, Değer: 25
``
/*
---

## 6. Struct Tag Okuma

Go’da struct alanlarına **tag** eklenebilir.
`reflect` ile bunları runtime’da okuyabilirsin.
*/
``go
type Product struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func main() {
	t := reflect.TypeOf(Product{})
	field := t.Field(1)
	fmt.Println("Tag:", field.Tag.Get("json")) // title
}
``
/*
---

## 7. Set Edilebilirlik (Settable)

Varsayılan olarak `reflect.ValueOf(x)` ile gelen `Value` **readonly**’dir.
Eğer değiştirmek istiyorsan **pointer** göndermelisin.
*/
``go
func main() {
	x := 10
	v := reflect.ValueOf(&x).Elem() // Elem() pointer içindeki değeri çıkarır

	if v.CanSet() {
		v.SetInt(99)
	}
	fmt.Println("Yeni değer:", x) // 99
}
``
/*
---

## 8. Dinamik Fonksiyon Çağırma
*/
``go
func Hello(name string) {
	fmt.Println("Merhaba,", name)
}

func main() {
	v := reflect.ValueOf(Hello)
	args := []reflect.Value{reflect.ValueOf("Abdullah")}
	v.Call(args) // Hello("Abdullah")
}
``

//📌 Çıktı:

``
Merhaba, Abdullah
``
/*
---

## 9. Interface Dönüşümü

Bir `reflect.Value` tekrar interface’e dönüştürülebilir:
*/
``go
x := 3.14
v := reflect.ValueOf(x)
i := v.Interface() // interface{} döner
y := i.(float64)   // type assertion
fmt.Println(y + 1) // 4.14
``
/*
---

## 10. Yeni Değer Üretmek (`reflect.New`)
*/
``go
t := reflect.TypeOf(42) // int
v := reflect.New(t)     // *int
fmt.Println(v.Type())   // *int

*v.Interface().(*int) = 100
fmt.Println(*v.Interface().(*int)) // 100
``
/*
---

## 11. Map ve Slice ile Çalışma

### Map
*/
``go
m := map[string]int{"a": 1, "b": 2}
v := reflect.ValueOf(m)

for _, key := range v.MapKeys() {
	fmt.Println("Key:", key, "Value:", v.MapIndex(key))
}
``

//### Slice

``go
s := []int{10, 20, 30}
v := reflect.ValueOf(s)
for i := 0; i < v.Len(); i++ {
	fmt.Println("Index:", i, "Value:", v.Index(i))
}
``

---

//## 12. Struct Metotlarını Listeleme

``go
type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Println("Merhaba,", p.Name)
}

func main() {
	t := reflect.TypeOf(Person{})
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println("Metot:", m.Name, "Tip:", m.Type)
	}
}
``
/*
---

## 13. Özet – `reflect` ile Yapabileceklerin

✅ Tip öğrenme (`TypeOf`, `Kind`)
✅ Değer öğrenme (`ValueOf`)
✅ Struct alanlarını okuma (`Field`)
✅ Struct tag okuma (`Tag.Get`)
✅ Değer değiştirme (`CanSet`, `SetInt`, `SetString` vs.)
✅ Dinamik fonksiyon çağırma (`Call`)
✅ Yeni değer oluşturma (`New`)
✅ Map/Slice üzerinde dolaşma
✅ Metotlara erişim (`Method`)

---

🔑 **Sonuç:**
`reflect` paketi Go’da **dinamik davranışlar**, **JSON/ORM kütüphaneleri**, **dependency injection**, **unit test frameworkleri** gibi yerlerde kritik rol oynar.

---

👉 İstersen ben sana şimdi bunun üstüne **mini bir JSON serializer (yani kendi `encoding/json` benzeri kodunu)** sadece `reflect` ile yazabilirim. Böylece paketin gerçek hayatta nasıl kullanıldığını çok iyi görebilirsin.

Bunu da ister misin?
EVET
*/
/*
Harika 😊 o zaman sana **`reflect` kullanarak mini bir JSON serializer** (kendi `encoding/json` benzeri basit kütüphane) örneği hazırlayayım.
Böylece `reflect`’in **gerçek hayattaki kullanımını** göreceksin.

---

# 📦 Mini JSON Serializer (`reflect` ile)

## Kod
*/
``go
package main

import (
	"fmt"
	"reflect"
	"strings"
)

// Mini JSON Encoder
func ToJSON(v interface{}) string {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	if typ.Kind() != reflect.Struct {
		return "{}"
	}

	var fields []string
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		// JSON key (tag varsa tag kullan)
		key := field.Tag.Get("json")
		if key == "" {
			key = field.Name
		}

		// Değeri string olarak hazırla
		var strVal string
		switch value.Kind() {
		case reflect.String:
			strVal = fmt.Sprintf("\"%s\"", value.String())
		case reflect.Int, reflect.Int64, reflect.Int32:
			strVal = fmt.Sprintf("%d", value.Int())
		case reflect.Bool:
			strVal = fmt.Sprintf("%t", value.Bool())
		default:
			strVal = fmt.Sprintf("\"%v\"", value.Interface())
		}

		fields = append(fields, fmt.Sprintf("\"%s\": %s", key, strVal))
	}

	return "{" + strings.Join(fields, ", ") + "}"
}

// Örnek struct
type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

func main() {
	u := User{"Abdullah", 25, true}
	jsonStr := ToJSON(u)
	fmt.Println(jsonStr)
}
``

---

//## Çıktı

``
{"name": "Abdullah", "age": 25, "active": true}
``
/*
---

## Mantık

* `reflect.TypeOf(u)` → struct’ın alanlarını öğreniyoruz.
* `reflect.ValueOf(u)` → alan değerlerini alıyoruz.
* `Field(i).Tag.Get("json")` → JSON tag’ini okuyoruz.
* Tipine göre (`Kind`) string, int, bool değerleri dönüştürüyoruz.
* `{"key": value}` şeklinde string birleştiriyoruz.

---

📌 Bu sadece **basit bir serializer** — gerçek `encoding/json` kütüphanesi:

* Slice, Map, Pointer, Interface, Nil kontrolü
* Nested Struct
* Float formatlama
* Özel Encoder interface’leri

gibi çok daha fazla şeyi kapsıyor.

Ama bu küçük örnek sayesinde `reflect` paketinin JSON gibi kütüphanelerde **neden kritik** olduğunu gördün.

---

👉 İstersen ben sana bir adım daha ileri gidip **reflect ile struct → SQL INSERT query generator** (yani ORM mantığı) da yazabilirim.
Böylece veritabanı katmanında da `reflect` kullanımını görmüş olursun.

Onu da ister misin?
EVET
*/
/*
Süper 👍 Şimdi sana `reflect` paketi kullanarak **mini bir ORM benzeri SQL INSERT generator** yazalım.
Amaç şu olacak:

* Struct alacağız.
* Alan isimleri → tablo sütunları
* Alan değerleri → SQL VALUES
* Struct tag varsa (`db:"colname"`) onu kullanacağız.

---

# 📦 Reflect ile SQL INSERT Generator

## Kod
*/
``go
package main

import (
	"fmt"
	"reflect"
	"strings"
)

// SQL Insert Generator
func ToSQLInsert(table string, v interface{}) string {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	if typ.Kind() != reflect.Struct {
		return ""
	}

	var cols []string
	var vals []string

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		// Sütun adı (tag varsa onu kullan)
		col := field.Tag.Get("db")
		if col == "" {
			col = field.Name
		}
		cols = append(cols, col)

		// Değer stringe çevrilir
		switch value.Kind() {
		case reflect.String:
			vals = append(vals, fmt.Sprintf("'%s'", value.String()))
		case reflect.Int, reflect.Int64, reflect.Int32:
			vals = append(vals, fmt.Sprintf("%d", value.Int()))
		case reflect.Bool:
			if value.Bool() {
				vals = append(vals, "TRUE")
			} else {
				vals = append(vals, "FALSE")
			}
		default:
			vals = append(vals, fmt.Sprintf("'%v'", value.Interface()))
		}
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s);",
		table,
		strings.Join(cols, ", "),
		strings.Join(vals, ", "),
	)

	return query
}

// Örnek struct
type User struct {
	ID     int    `db:"id"`
	Name   string `db:"username"`
	Age    int    `db:"age"`
	Active bool   `db:"is_active"`
}

func main() {
	u := User{ID: 1, Name: "Abdullah", Age: 25, Active: true}
	sql := ToSQLInsert("users", u)
	fmt.Println(sql)
}
``
/*
---

## Çıktı
*/
``
INSERT INTO users (id, username, age, is_active) VALUES (1, 'Abdullah', 25, TRUE);
``
/*
---

## Mantık

1. `reflect.TypeOf(v)` → struct’ın alanlarını aldık.
2. `reflect.ValueOf(v)` → değerleri aldık.
3. Eğer `db:"colname"` tag’i varsa sütun adı olarak onu kullandık.
4. String değerleri `'...'`, int değerleri sayı olarak, bool değerleri `TRUE/FALSE` çevirdik.
5. Sonunda dinamik bir SQL cümlesi ürettik.

---

## Geliştirme Fikirleri

* `UPDATE` query generator eklenebilir.
* Nested struct desteği eklenebilir.
* `omitempty` gibi tag’lerle null kontrolü yapılabilir.
* `reflect.New()` ile boş struct oluşturulup scan yapılabilir (SELECT için).

---

🔑 Görüldüğü gibi `reflect`, **ORM mantığının kalbidir**. GORM, sqlx gibi popüler kütüphaneler tam olarak bu mantıkla çalışıyor 🚀

---

👉 İstersen bir adım daha ileri gidip sana **reflect ile dependency injection container** örneği gösterebilirim (yani hangi servisin hangi struct’a enjekte edileceğini runtime’da belirleme).

Onu da ister misin?
*/