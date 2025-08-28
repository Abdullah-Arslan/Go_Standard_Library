/*
Go’da **`compress/flate` paketi**, **DEFLATE algoritması** ile veri sıkıştırma ve açma işlemleri yapmak için kullanılır.

* **DEFLATE**: Lempel-Ziv (LZ77) tabanlı bir sıkıştırma algoritmasıdır ve gzip, zlib gibi formatların temelini oluşturur.
* **Avantaj**: Hızlı sıkıştırma ve açma, düşük bellek kullanımı.

Go’daki `compress/flate` paketi, sıkıştırma ve açma işlemlerini **`Writer`** ve **`Reader`** aracılığıyla yapar.

---

# 📦 `compress/flate` Paketinin Temel Yapısı

### 1️⃣ Writer (Sıkıştırma)

```go
func NewWriter(w io.Writer, level int) (*Writer, error)
```

* `w io.Writer`: sıkıştırılmış verinin yazılacağı hedef
* `level int`: sıkıştırma seviyesi (0-9 veya `flate.BestSpeed`, `flate.BestCompression`, `flate.DefaultCompression`)

### 2️⃣ Reader (Açma)

```go
func NewReader(r io.Reader) io.ReadCloser
```

* `r io.Reader`: DEFLATE ile sıkıştırılmış veri
* Döndürdüğü `io.ReadCloser` ile veriyi açabilirsiniz

---

# 🔹 Örnekler

### 1️⃣ Basit Sıkıştırma ve Açma
*/

package main

import (
	"bytes"
	"compress/flate"
	"fmt"
	"io"
)

func main() {
	data := []byte("Merhaba Go! compress/flate paketi örneği.")

	// Sıkıştırma
	var buf bytes.Buffer
	writer, _ := flate.NewWriter(&buf, flate.BestCompression)
	writer.Write(data)
	writer.Close()
	fmt.Println("Sıkıştırılmış boyut:", buf.Len())

	// Açma
	reader := flate.NewReader(&buf)
	defer reader.Close()
	out := new(bytes.Buffer)
	io.Copy(out, reader)
	fmt.Println("Açılmış veri:", out.String())
}


//---

//### 2️⃣ Farklı Sıkıştırma Seviyeleri


levels := []int{
	flate.NoCompression,
	flate.BestSpeed,
	flate.BestCompression,
	flate.DefaultCompression,
}

for _, level := range levels {
	var buf bytes.Buffer
	writer, _ := flate.NewWriter(&buf, level)
	writer.Write(data)
	writer.Close()
	fmt.Printf("Seviye %d boyut: %d\n", level, buf.Len())
}


// ---

// ### 3️⃣ Reader ve Writer Arasında Doğrudan Veri Transferi


input := []byte("Uzun bir metin bloğu burada yer alıyor. DEFLATE ile sıkıştırılacak.")

var compressed bytes.Buffer
w, _ := flate.NewWriter(&compressed, flate.BestCompression)
w.Write(input)
w.Close()

r := flate.NewReader(&compressed)
decompressed := new(bytes.Buffer)
io.Copy(decompressed, r)
r.Close()

fmt.Println("Açılmış veri:", decompressed.String())

/*
---

# ⚡ Özet

| Yapı / Fonksiyon          | Açıklama                                                              |
| ------------------------- | --------------------------------------------------------------------- |
| `flate.NewWriter`         | Sıkıştırıcı oluşturur                                                 |
| `flate.NewReader`         | Açıcı oluşturur                                                       |
| `Writer.Write`            | Veriyi sıkıştırır                                                     |
| `Writer.Close`            | Sıkıştırmayı bitirir                                                  |
| `Reader.Read` / `io.Copy` | Veriyi açar                                                           |
| Sıkıştırma Seviyeleri     | `NoCompression`, `BestSpeed`, `BestCompression`, `DefaultCompression` |

---

`compress/flate` paketi genellikle **gzip ve zlib paketlerinin altında da kullanılır**, çünkü bu paketler DEFLATE algoritmasını temel alır.

---

*/