/*
Go’da **`compress` paketi**, veri sıkıştırma ve açma işlemleri için kullanılan bir üst pakettir.

* `compress` paketi **tek başına kullanılmaz**, altında çeşitli alt paketler vardır:

  * **`compress/gzip`** → GZIP formatı
  * **`compress/zlib`** → ZLIB formatı
  * **`compress/flate`** → DEFLATE algoritması
  * **`compress/lzw`** → Lempel-Ziv-Welch algoritması
* Bu paketler genellikle **dosya sıkıştırma, network veri sıkıştırma veya bellek tasarrufu** amaçlı kullanılır.

---

# 📦 `compress` Paketinin Alt Paketleri ve Kullanımı

---

### 1️⃣ **`compress/gzip`**

* En yaygın kullanılan paket. `.gz` dosyalarını oluşturur ve okur.
* `gzip.Writer` → Sıkıştırır (write)
* `gzip.Reader` → Açar (read)
* `gzip.Header` → Dosya bilgisi tutar

#### Örnek: GZIP ile sıkıştırma
*/

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
)

func main() {
	data := []byte("Merhaba Go, compress/gzip örneği!")

	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)

	writer.Write(data)
	writer.Close()

	fmt.Println("Sıkıştırılmış boyut:", buf.Len())

	// Açma
	reader, _ := gzip.NewReader(&buf)
	defer reader.Close()

	out := new(bytes.Buffer)
	out.ReadFrom(reader)
	fmt.Println("Açılmış veri:", out.String())
}

/*
---

### 2️⃣ **`compress/zlib`**

* ZLIB formatını destekler (`deflate` + header).
* `zlib.NewWriter` → Sıkıştırır
* `zlib.NewReader` → Açar
*/

package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
)

func main() {
	data := []byte("Merhaba Go, compress/zlib örneği!")

	var buf bytes.Buffer
	writer := zlib.NewWriter(&buf)
	writer.Write(data)
	writer.Close()

	fmt.Println("Sıkıştırılmış boyut:", buf.Len())

	reader, _ := zlib.NewReader(&buf)
	defer reader.Close()

	out := new(bytes.Buffer)
	out.ReadFrom(reader)
	fmt.Println("Açılmış veri:", out.String())
}

/*
---

 3️⃣ **`compress/flate`**

* DEFLATE algoritması temelinde sıkıştırma yapar.
* `flate.NewWriter` → Sıkıştırır
* `flate.NewReader` → Açar

*/
package main

import (
	"bytes"
	"compress/flate"
	"fmt"
	"io"
)

func main() {
	data := []byte("Merhaba Go, compress/flate örneği!")

	var buf bytes.Buffer
	writer, _ := flate.NewWriter(&buf, flate.BestCompression)
	writer.Write(data)
	writer.Close()

	fmt.Println("Sıkıştırılmış boyut:", buf.Len())

	reader := flate.NewReader(&buf)
	defer reader.Close()
	out := new(bytes.Buffer)
	io.Copy(out, reader)
	fmt.Println("Açılmış veri:", out.String())
}

/*
---

### 4️⃣ **`compress/lzw`**

* Lempel-Ziv-Welch algoritmasını uygular.
* Genellikle TIFF veya GIF dosyalarında kullanılır.
* `lzw.NewWriter` → Sıkıştırır
* `lzw.NewReader` → Açar

*/

package main

import (
	"bytes"
	"compress/lzw"
	"fmt"
	"io"
)

func main() {
	data := []byte("Merhaba Go, compress/lzw örneği!")

	var buf bytes.Buffer
	writer := lzw.NewWriter(&buf, lzw.LSB, 8)
	writer.Write(data)
	writer.Close()

	fmt.Println("Sıkıştırılmış boyut:", buf.Len())

	reader := lzw.NewReader(&buf, lzw.LSB, 8)
	defer reader.Close()

	out := new(bytes.Buffer)
	io.Copy(out, reader)
	fmt.Println("Açılmış veri:", out.String())
}


/* ---

# 🚀 Özet

| Paket            | Açıklama                              |
| ---------------- | ------------------------------------- |
| `compress/gzip`  | GZIP formatı, en yaygın               |
| `compress/zlib`  | ZLIB formatı, deflate + header        |
| `compress/flate` | DEFLATE algoritması, temel sıkıştırma |
| `compress/lzw`   | LZW algoritması, TIFF/GIF için        |

 **Genel Mantık:**

 1. `Writer` → veri sıkıştırır
 2. `Reader` → veri açar
 3. Alt paketler algoritma ve format açısından farklıdır ama kullanım şekli benzerdir.

*/---
