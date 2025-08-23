/*
Go (Golang) `bytes` paketini** en baştan, tüm fonksiyonlarıyla ve örneklerle açıklayayım.

---

# 🔑 `bytes` Paketi Nedir?

Go’daki [`bytes`](https://pkg.go.dev/bytes) paketi, **byte slice (`[]byte`)** üzerinde işlem yapmak için araçlar sağlar.

* `strings` paketi string işlemleri için neyse,
* `bytes` paketi de `[]byte` işlemleri için odur.

✅ Metin arama, karşılaştırma, değiştirme, bölme, birleştirme
✅ Okuma/yazma için buffer yönetimi
✅ Byte seviyesinde manipülasyon

---

# 📦 Paketin İçeriği

`bytes` paketinde 2 ana kısım var:

1. **Fonksiyonlar** (`Contains`, `Split`, `Replace`, `Compare` vs.)
2. **Tipler**

   * `Buffer` → verimli string/byte birleştirme ve yazma
   * `Reader` → byte slice üzerinde okuma

---

# 🔧 Fonksiyonlar ve Örnekler

### 1. Arama Fonksiyonları
*/
``go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	data := []byte("Merhaba Dünya")

	// Belirli bir alt string içeriyor mu?
	fmt.Println(bytes.Contains(data, []byte("Dünya"))) // true
	fmt.Println(bytes.ContainsAny(data, "xyz"))        // false

	// Belirli byte ile başlıyor/bitiyor mu?
	fmt.Println(bytes.HasPrefix(data, []byte("Merh"))) // true
	fmt.Println(bytes.HasSuffix(data, []byte("ya")))   // true

	// İlk index
	fmt.Println(bytes.Index(data, []byte("Dünya"))) // 8
	fmt.Println(bytes.LastIndex(data, []byte("a"))) // 12
}
``
/*
---

### 2. Karşılaştırma Fonksiyonları
?7
``go
a := []byte("abc")
b := []byte("abd")

fmt.Println(bytes.Equal(a, b))     // false
fmt.Println(bytes.Compare(a, b))   // -1 (a < b)
fmt.Println(bytes.EqualFold([]byte("Go"), []byte("gO"))) // true (case-insensitive)
``
/*
---

### 3. Bölme ve Birleştirme
*/
``go
s := []byte("go,java,python")

parts := bytes.Split(s, []byte(","))
fmt.Println(parts) // [[103 111] [106 97 118 97] [112 121 116 104 111 110]]

joined := bytes.Join(parts, []byte(" | "))
fmt.Println(string(joined)) // "go | java | python"
``

//---

//### 4. Değiştirme

``go
text := []byte("golang golang golang")

res := bytes.Replace(text, []byte("golang"), []byte("GO"), 2) 
fmt.Println(string(res)) // "GO GO golang"

res2 := bytes.ReplaceAll(text, []byte("golang"), []byte("GO"))
fmt.Println(string(res2)) // "GO GO GO"
``
/*
---

### 5. Trim (Kırpma)
*/
``go
s := []byte("   hello   ")

fmt.Println(string(bytes.TrimSpace(s)))   // "hello"
fmt.Println(string(bytes.Trim(s, " h")))  // "ello"
fmt.Println(string(bytes.TrimPrefix(s, []byte("   ")))) // "hello   "
fmt.Println(string(bytes.TrimSuffix(s, []byte("   ")))) // "   hello"
``
/*
---

### 6. Dönüştürme
*/
``go
fmt.Println(bytes.ToUpper([]byte("golang"))) // "GOLANG"
fmt.Println(bytes.ToLower([]byte("GoLANG"))) // "golang"
fmt.Println(bytes.Title([]byte("merhaba dünya"))) // "Merhaba Dünya"
``
/*
---

# 🔧 `Buffer` Tipi

`bytes.Buffer`, dinamik olarak büyüyen bir byte dizisidir.
Çok sık **string birleştirme** ve **I/O işlemleri** için kullanılır.

### Örnek: Yazma ve Okuma
*/
``go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer

	buf.WriteString("Merhaba")
	buf.WriteString(" Dünya")

	fmt.Println(buf.String()) // "Merhaba Dünya"

	// Okuma
	b := make([]byte, 8)
	n, _ := buf.Read(b)
	fmt.Printf("Okunan: %s (%d byte)\n", b[:n], n)
}
``
/*
---

# 🔧 `Reader` Tipi

`bytes.Reader`, bir `[]byte` üzerinde okuma yapmak için kullanılır.
`strings.NewReader`’in byte karşılığıdır.
*/
``go
package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	data := []byte("Merhaba Dünya")
	r := bytes.NewReader(data)

	b := make([]byte, 7)
	n, _ := r.Read(b)
	fmt.Printf("Okunan: %s (%d byte)\n", b[:n], n)

	// Seek ile pozisyon değiştirme
	r.Seek(8, io.SeekStart)
	b2 := make([]byte, 5)
	r.Read(b2)
	fmt.Printf("Kalan: %s\n", b2)
}
``
/*
📌 Çıktı:

```
Okunan: Merhaba (7 byte)
Kalan: Dünya
```

---

# 📌 Özet

* `bytes` paketi → `[]byte` üzerinde `strings` paketine benzer fonksiyonlar sağlar.
* Önemli fonksiyonlar:

  * **Arama:** `Contains`, `HasPrefix`, `Index`
  * **Karşılaştırma:** `Equal`, `Compare`, `EqualFold`
  * **Bölme/Birleştirme:** `Split`, `Join`
  * **Değiştirme:** `Replace`, `ReplaceAll`
  * **Trim/Dönüştürme:** `TrimSpace`, `ToUpper`, `ToLower`
* Tipler:

  * **`Buffer`** → dinamik yazma/okuma
  * **`Reader`** → `[]byte` üzerinde okuma

---

İstersen sana bir sonraki adımda, `bytes.Buffer` ve `bytes.Reader` ile **dosya okuma/yazma simülasyonu** yapmayı gösterebilirim. İster misin?
*/