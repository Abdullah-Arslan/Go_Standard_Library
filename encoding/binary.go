/*
**Go’nun `encoding/binary` paketini** baştan sona örneklerle açıklayayım.

---

# 📌 Go `encoding/binary` Paketi

Go’daki **`encoding/binary`** paketi, **binary (ikili) verileri encode ve decode etmek** için kullanılır.
Yani:

* **struct → \[]byte** (encode)
* **\[]byte → struct** (decode)

Ayrıca sayıları **BigEndian** veya **LittleEndian** formatında byte dizisine çevirme imkanı sağlar.

---

## 🔎 Endianness Nedir?

* **BigEndian**: En anlamlı byte en başa yazılır. (network protokollerinde yaygındır)
* **LittleEndian**: En anlamlı byte en sona yazılır. (Intel CPU’larda yaygındır)

Örneğin sayı `0x12345678` olsun:

* **BigEndian** → `[12 34 56 78]`
* **LittleEndian** → `[78 56 34 12]`

---

## 📦 Paketi Import Etme
*/
``go
import (
    "encoding/binary"
    "bytes"
    "fmt"
)
``
/*
---

# 🚀 Örneklerle Kullanım

## 1. Bir Sayıyı Byte Dizisine Çevirmek
*/
``go
package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	var num uint32 = 305419896 // 0x12345678

	// BigEndian
	big := make([]byte, 4)
	binary.BigEndian.PutUint32(big, num)
	fmt.Println("BigEndian:", big)

	// LittleEndian
	little := make([]byte, 4)
	binary.LittleEndian.PutUint32(little, num)
	fmt.Println("LittleEndian:", little)
}
``
/*
🔹 Çıktı:

```
BigEndian: [18 52 86 120]
LittleEndian: [120 86 52 18]
```

---

## 2. Byte Dizisinden Sayı Okuma
*/
``go
package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	data := []byte{0x12, 0x34, 0x56, 0x78}

	// BigEndian
	num1 := binary.BigEndian.Uint32(data)
	fmt.Println("BigEndian:", num1)

	// LittleEndian
	num2 := binary.LittleEndian.Uint32(data)
	fmt.Println("LittleEndian:", num2)
}
``
/*
🔹 Çıktı:

```
BigEndian: 305419896
LittleEndian: 2018915346
```

---

## 3. Struct → Byte Dizisi (binary.Write)
*/
``go
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Person struct {
	Age   uint8
	Score uint16
}

func main() {
	buf := new(bytes.Buffer)

	p := Person{Age: 25, Score: 900}

	// Struct’ı LittleEndian ile yaz
	err := binary.Write(buf, binary.LittleEndian, p)
	if err != nil {
		panic(err)
	}

	fmt.Println("Struct to Bytes:", buf.Bytes())
}
``
/*
🔹 Çıktı:

```
Struct to Bytes: [25 132 3]
```

(25 = yaş, 900 = `0x0384` → `132 3` little endian)

---

## 4. Byte Dizisi → Struct (binary.Read)
*/
``go
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Person struct {
	Age   uint8
	Score uint16
}

func main() {
	// [25 yaş, 900 skor]
	data := []byte{25, 132, 3}
	buf := bytes.NewReader(data)

	var p Person
	err := binary.Read(buf, binary.LittleEndian, &p)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Decoded Struct: %+v\n", p)
}
``
/*
🔹 Çıktı:

```
Decoded Struct: {Age:25 Score:900}
```

---

## 5. bool / string Encode Etme (manuel)

`binary` paketi sadece fixed-size tipler (`int8, uint16, float64` vs.) için çalışır.
Eğer `string` veya `bool` encode etmek istiyorsak, kendimiz uzunluğu yazıp sonra veriyi byte’a çevirmeliyiz.
*/
``go
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	buf := new(bytes.Buffer)

	// String encode et (önce uzunluk, sonra data)
	str := "GoLang"
	binary.Write(buf, binary.LittleEndian, uint16(len(str)))
	buf.Write([]byte(str))

	// Decode
	var length uint16
	binary.Read(buf, binary.LittleEndian, &length)
	strBytes := make([]byte, length)
	buf.Read(strBytes)

	fmt.Println("Decoded String:", string(strBytes))
}
``
/*
🔹 Çıktı:

```
Decoded String: GoLang
```

---

## 6. binary.Size()

Bir struct’ın kaç byte tuttuğunu bulmak için:
*/

``go
package main

import (
	"encoding/binary"
	"fmt"
)

type Example struct {
	A int32
	B int16
}

func main() {
	fmt.Println("Struct size:", binary.Size(Example{})) // 8 byte
}
``
/*
---

# 📌 Özet

* **`binary.BigEndian` / `binary.LittleEndian`** → Endianness yönetir.
* **`PutUint16/32/64`** → Sayıyı byte dizisine çevirir.
* **`Uint16/32/64`** → Byte dizisini sayıya çevirir.
* **`binary.Write`** → Struct → Bytes
* **`binary.Read`** → Bytes → Struct
* **`binary.Size`** → Struct’ın boyutunu döner.

---

👉 İstersen sana bununla ilgili **küçük bir dosya formatı (ör. binary protokol)** oluşturup, hem yazma hem okuma örneği de yapabilirim.

Bunu ister misin?
EVET
*/
/*
Süper 👍 O zaman sana **küçük bir Binary Protokol** örneği yapalım.
Burada şunu yapacağız:

1. 📄 **Kendi binary dosya formatımızı tanımlayacağız**
2. ✍️ Struct’ı byte dizisine yazacağız
3. 📥 Dosyadan geri okuyacağız

---

# 🔹 1. Binary Formatımız

Her kaydı şöyle tasarlayalım:

| Alan      | Tip      | Açıklama        |
| --------- | -------- | --------------- |
| `ID`      | `uint32` | Kullanıcı ID’si |
| `Age`     | `uint8`  | Yaş             |
| `NameLen` | `uint16` | İsim uzunluğu   |
| `Name`    | `[]byte` | İsim (UTF-8)    |

---

# 🔹 2. Kod (Binary Yaz & Oku)
*/

``go
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

type User struct {
	ID   uint32
	Age  uint8
	Name string
}

// Binary dosyaya yaz
func WriteUserToFile(filename string, user User) error {
	buf := new(bytes.Buffer)

	// ID ve Age yaz
	if err := binary.Write(buf, binary.LittleEndian, user.ID); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, user.Age); err != nil {
		return err
	}

	// İsim uzunluğu ve isim byte’larını yaz
	nameBytes := []byte(user.Name)
	nameLen := uint16(len(nameBytes))
	if err := binary.Write(buf, binary.LittleEndian, nameLen); err != nil {
		return err
	}
	if _, err := buf.Write(nameBytes); err != nil {
		return err
	}

	// Dosyaya kaydet
	return os.WriteFile(filename, buf.Bytes(), 0644)
}

// Binary dosyadan oku
func ReadUserFromFile(filename string) (User, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return User{}, err
	}
	buf := bytes.NewReader(data)

	var u User

	// ID ve Age oku
	if err := binary.Read(buf, binary.LittleEndian, &u.ID); err != nil {
		return u, err
	}
	if err := binary.Read(buf, binary.LittleEndian, &u.Age); err != nil {
		return u, err
	}

	// İsim uzunluğunu oku
	var nameLen uint16
	if err := binary.Read(buf, binary.LittleEndian, &nameLen); err != nil {
		return u, err
	}

	// İsmi oku
	nameBytes := make([]byte, nameLen)
	if _, err := buf.Read(nameBytes); err != nil {
		return u, err
	}
	u.Name = string(nameBytes)

	return u, nil
}

func main() {
	user := User{ID: 12345, Age: 30, Name: "Abdullah"}

	// Dosyaya yaz
	if err := WriteUserToFile("user.dat", user); err != nil {
		panic(err)
	}
	fmt.Println("Kullanıcı dosyaya yazıldı.")

	// Dosyadan oku
	readUser, err := ReadUserFromFile("user.dat")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Dosyadan okunan: %+v\n", readUser)
}
``
/*
---

# 🔹 3. Örnek Çalıştırma

```
Kullanıcı dosyaya yazıldı.
Dosyadan okunan: {ID:12345 Age:30 Name:Abdullah}
```

Ve `user.dat` dosyasını hex editörde açarsan şunu görürsün:

```
39 30 00 00 1e 08 00 41 62 64 75 6c 6c 61 68
```

* `39 30 00 00` → `12345` (LittleEndian)
* `1e` → `30` yaş
* `08 00` → isim uzunluğu 8
* `41 62 64 75 6c 6c 61 68` → “Abdullah”

---

👉 Böylece kendi **mini binary protokolümüzü** yazmış olduk 🚀

İstersen ben bunu genişletip **birden fazla kullanıcıyı (array of struct)** binary dosyaya yazıp tekrar okumayı da gösterebilirim. Bunu ister misin?
EVET
*/

/*
Harika 👍 O zaman şimdi tek bir kullanıcı yerine **birden fazla kullanıcıyı (slice of struct)** binary dosyaya yazıp tekrar okuyalım.

---

# 🔹 Çoklu Kullanıcıları Binary Dosyaya Yazma & Okuma
*/

``go
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

type User struct {
	ID   uint32
	Age  uint8
	Name string
}

// Tek kullanıcı yazma
func writeUser(buf *bytes.Buffer, user User) error {
	// ID, Age
	if err := binary.Write(buf, binary.LittleEndian, user.ID); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, user.Age); err != nil {
		return err
	}

	// NameLen + Name
	nameBytes := []byte(user.Name)
	nameLen := uint16(len(nameBytes))
	if err := binary.Write(buf, binary.LittleEndian, nameLen); err != nil {
		return err
	}
	if _, err := buf.Write(nameBytes); err != nil {
		return err
	}
	return nil
}

// Tek kullanıcı okuma
func readUser(buf *bytes.Reader) (User, error) {
	var u User
	if err := binary.Read(buf, binary.LittleEndian, &u.ID); err != nil {
		return u, err
	}
	if err := binary.Read(buf, binary.LittleEndian, &u.Age); err != nil {
		return u, err
	}
	var nameLen uint16
	if err := binary.Read(buf, binary.LittleEndian, &nameLen); err != nil {
		return u, err
	}
	nameBytes := make([]byte, nameLen)
	if _, err := buf.Read(nameBytes); err != nil {
		return u, err
	}
	u.Name = string(nameBytes)
	return u, nil
}

// Çoklu kullanıcıları dosyaya yaz
func WriteUsers(filename string, users []User) error {
	buf := new(bytes.Buffer)

	// İlk önce kaç kullanıcı olduğunu yaz
	count := uint32(len(users))
	if err := binary.Write(buf, binary.LittleEndian, count); err != nil {
		return err
	}

	// Kullanıcıları sırayla yaz
	for _, u := range users {
		if err := writeUser(buf, u); err != nil {
			return err
		}
	}
	return os.WriteFile(filename, buf.Bytes(), 0644)
}

// Çoklu kullanıcıları dosyadan oku
func ReadUsers(filename string) ([]User, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewReader(data)

	// Kullanıcı sayısını oku
	var count uint32
	if err := binary.Read(buf, binary.LittleEndian, &count); err != nil {
		return nil, err
	}

	// Kullanıcıları sırayla oku
	users := make([]User, 0, count)
	for i := uint32(0); i < count; i++ {
		u, err := readUser(buf)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func main() {
	users := []User{
		{ID: 1, Age: 25, Name: "Ahmet"},
		{ID: 2, Age: 30, Name: "Mehmet"},
		{ID: 3, Age: 22, Name: "Ayşe"},
	}

	// Dosyaya yaz
	if err := WriteUsers("users.dat", users); err != nil {
		panic(err)
	}
	fmt.Println("Kullanıcılar dosyaya yazıldı.")

	// Dosyadan oku
	readUsers, err := ReadUsers("users.dat")
	if err != nil {
		panic(err)
	}

	fmt.Println("Dosyadan okunan kullanıcılar:")
	for _, u := range readUsers {
		fmt.Printf("ID=%d Age=%d Name=%s\n", u.ID, u.Age, u.Name)
	}
}
``

/*
---

# 🔹 Örnek Çalıştırma

```
Kullanıcılar dosyaya yazıldı.
Dosyadan okunan kullanıcılar:
ID=1 Age=25 Name=Ahmet
ID=2 Age=30 Name=Mehmet
ID=3 Age=22 Name=Ayşe
```

---

📂 `users.dat` dosyasına bakarsan:

* İlk 4 byte → kullanıcı sayısı (`3`)
* Sonrası → her bir `User` struct’ın binary hali

---

👉 Bu yöntemle aslında kendi **mini veritabanımızı / custom binary protokolümüzü** yazmış olduk 🚀

İstersen bir sonraki adımda sana bunu **TCP soket üzerinden göndermeyi** (client-server arasında binary iletişim) gösterebilirim. Onu da ister misin?
*/