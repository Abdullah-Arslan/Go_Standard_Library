/*
Go’da **`compress/lzw` paketi**, **Lempel-Ziv-Welch (LZW) algoritması** ile veri sıkıştırma ve açma işlemleri yapmak için kullanılır.

* **LZW algoritması**: Tekrar eden veri dizilerini tablolayarak sıkıştırma yapar.
* **Kullanım alanları**: GIF dosyaları, TIFF formatları ve bazı eski sıkıştırma formatları.
* **Avantaj**: Basit, hızlı ve tekrarlayan verilerde iyi sıkıştırma sağlar.

---

# 📦 `compress/lzw` Paketinin Temel Yapısı

### 1️⃣ Writer (Sıkıştırma)

```go
func NewWriter(w io.Writer, order Order, litWidth int) *Writer
```

* `w io.Writer`: Sıkıştırılmış verinin yazılacağı hedef
* `order Order`: Bit sıralama (`LSB` veya `MSB`)
* `litWidth int`: Literal bit genişliği (genellikle 8)

### 2️⃣ Reader (Açma)

```go
func NewReader(r io.Reader, order Order, litWidth int) io.ReadCloser
```

* `r io.Reader`: LZW ile sıkıştırılmış veri
* `order Order`: Bit sıralama (`LSB` veya `MSB`)
* `litWidth int`: Literal bit genişliği
* Döndürdüğü `io.ReadCloser` ile veriyi açabilirsiniz

---

# 🔹 Örnekler

### 1️⃣ Basit Sıkıştırma ve Açma
*/

package main

import (
	"bytes"
	"compress/lzw"
	"fmt"
	"io"
)

func main() {
	data := []byte("Merhaba Go! compress/lzw paketi örneği.")

	// Sıkıştırma
	var buf bytes.Buffer
	writer := lzw.NewWriter(&buf, lzw.LSB, 8)
	writer.Write(data)
	writer.Close()
	fmt.Println("Sıkıştırılmış boyut:", buf.Len())

	// Açma
	reader := lzw.NewReader(&buf, lzw.LSB, 8)
	defer reader.Close()
	out := new(bytes.Buffer)
	io.Copy(out, reader)
	fmt.Println("Açılmış veri:", out.String())
}

/*
---

### 2️⃣ MSB ve LSB Farkı

```go
// LSB (Least Significant Bit) ve MSB (Most Significant Bit) farklı bit sıralaması sağlar
writerLSB := lzw.NewWriter(&buf, lzw.LSB, 8)
writerMSB := lzw.NewWriter(&buf, lzw.MSB, 8)
```

* LSB ve MSB, sıkıştırılan verinin bit sıralamasını belirler.
* Genellikle TIFF ve GIF formatları bu sıralamalara göre değişir.

---

### 3️⃣ Reader ve Writer Arasında Doğrudan Veri Transferi
*/

input := []byte("Uzun bir metin bloğu LZW ile sıkıştırılacak.")

var compressed bytes.Buffer
w := lzw.NewWriter(&compressed, lzw.LSB, 8)
w.Write(input)
w.Close()

r := lzw.NewReader(&compressed, lzw.LSB, 8)
decompressed := new(bytes.Buffer)
io.Copy(decompressed, r)
r.Close()

fmt.Println("Açılmış veri:", decompressed.String())

/*
---

# ⚡ Özet

| Yapı / Fonksiyon          | Açıklama                             |
| ------------------------- | ------------------------------------ |
| `lzw.NewWriter`           | Sıkıştırıcı oluşturur                |
| `lzw.NewReader`           | Açıcı oluşturur                      |
| `Writer.Write`            | Veriyi sıkıştırır                    |
| `Writer.Close`            | Sıkıştırmayı bitirir                 |
| `Reader.Read` / `io.Copy` | Veriyi açar                          |
| `Order`                   | Bit sıralaması (`LSB` veya `MSB`)    |
| `litWidth`                | Literal bit genişliği (genellikle 8) |

---

`compress/lzw` paketi genellikle **GIF veya TIFF gibi formatlarda** kullanılır ve **basit, hızlı sıkıştırma** sağlar.

---

İstersen sana **interaktif bir LZW demo uygulaması** hazırlayıp kullanıcı metin girip LSB veya MSB seçebilsin ve sıkıştırma/sonucu görebilsin şeklinde bir örnek yazabilirim.

Bunu yapayım mı? Evet
*/
//-----------------------------
/*
Harika! O zaman Go’da **interaktif bir LZW demo uygulaması** hazırlayalım:

* Kullanıcı metin girecek
* LSB veya MSB bit sıralamasını seçebilecek
* Program girilen metni LZW ile sıkıştıracak ve açacak
* Sıkıştırılmış boyutu ve açılmış veriyi gösterecek

---

## 📝 Uygulama: `lzw_cli_demo.go`

*/
package main

import (
	"bufio"
	"bytes"
	"compress/lzw"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Bir metin girin: ")
	inputStr, _ := reader.ReadString('\n')
	inputStr = strings.TrimSpace(inputStr)
	data := []byte(inputStr)

	for {
		fmt.Println("\nHangi bit sıralamasını kullanmak istersiniz?")
		fmt.Println("1. LSB (Least Significant Bit)")
		fmt.Println("2. MSB (Most Significant Bit)")
		fmt.Println("3. Çıkış")
		fmt.Print("Seçiminiz: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)

		switch choiceStr {
		case "1":
			runLZW(data, lzw.LSB)
		case "2":
			runLZW(data, lzw.MSB)
		case "3":
			fmt.Println("Çıkış yapılıyor...")
			return
		default:
			fmt.Println("Geçersiz seçim!")
		}
	}
}

///////////////////// Fonksiyon /////////////////////

func runLZW(data []byte, order lzw.Order) {
	var buf bytes.Buffer
	writer := lzw.NewWriter(&buf, order, 8)
	writer.Write(data)
	writer.Close()

	fmt.Println("Sıkıştırılmış boyut:", buf.Len())

	reader := lzw.NewReader(&buf, order, 8)
	defer reader.Close()
	out := new(bytes.Buffer)
	io.Copy(out, reader)

	fmt.Println("Açılmış veri:", out.String())
}

/*
---

### 📌 Açıklamalar

1. Kullanıcıdan bir metin alınır (`bufio.Reader`)
2. Menü ile **LSB veya MSB bit sıralaması** seçilir
3. Seçilen sıralama ile:

   * LZW Writer oluşturulur ve veri sıkıştırılır
   * Sıkıştırılmış boyut ekrana yazılır
   * LZW Reader ile veri açılır ve ekrana yazılır
4. Çıkış seçeneği ile program sonlanır

---

### 📌 Örnek Çalışma

```
Bir metin girin: Merhaba Go LZW CLI!

Hangi bit sıralamasını kullanmak istersiniz?
1. LSB (Least Significant Bit)
2. MSB (Most Significant Bit)
3. Çıkış
Seçiminiz: 1
Sıkıştırılmış boyut: 44
Açılmış veri: Merhaba Go LZW CLI!

Seçiminiz: 2
Sıkıştırılmış boyut: 44
Açılmış veri: Merhaba Go LZW CLI!

Seçiminiz: 3
Çıkış yapılıyor...
```

---

Bu uygulama ile kullanıcı:

* LZW algoritmasını deneyimleyebilir
* LSB ve MSB arasındaki farkları görebilir
* Sıkıştırılmış boyut ve açılmış veriyi anında görebilir

---
*/
