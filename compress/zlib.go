/*
Go’da **`compress/zlib` paketi**, **ZLIB sıkıştırma formatını** kullanmak için tasarlanmış bir standart kütüphane paketidir.

* ZLIB, **DEFLATE algoritmasını temel alır**, yani `compress/flate` ile aynı sıkıştırma algoritmasını kullanır ama bir **header ve checksum** ekleyerek veri bütünlüğünü sağlar.
* Kullanım alanı: Dosya sıkıştırma, network veri sıkıştırma (HTTP), bellek tasarrufu.

---

# 📦 `compress/zlib` Paketinin Temel Yapısı

### 1️⃣ Writer (Sıkıştırma)

```go
func NewWriter(w io.Writer) *Writer
```

* `w io.Writer`: Sıkıştırılmış verinin yazılacağı hedef
* Döndürdüğü `Writer` ile veriyi sıkıştırabilirsiniz

```go
func (w *Writer) Write(p []byte) (int, error)
func (w *Writer) Close() error
```

* `Write` → veriyi sıkıştırır
* `Close` → sıkıştırmayı tamamlar ve flush eder

---

### 2️⃣ Reader (Açma)

```go
func NewReader(r io.Reader) (io.ReadCloser, error)
```

* `r io.Reader`: ZLIB ile sıkıştırılmış veri
* Döndürdüğü `io.ReadCloser` ile veriyi açabilirsiniz

---

# 🔹 Örnekler

### 1️⃣ Basit Sıkıştırma ve Açma
*/

package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
)

func main() {
	data := []byte("Merhaba Go! compress/zlib paketi örneği.")

	// Sıkıştırma
	var buf bytes.Buffer
	writer := zlib.NewWriter(&buf)
	writer.Write(data)
	writer.Close()
	fmt.Println("Sıkıştırılmış boyut:", buf.Len())

	// Açma
	reader, _ := zlib.NewReader(&buf)
	defer reader.Close()
	out := new(bytes.Buffer)
	io.Copy(out, reader)
	fmt.Println("Açılmış veri:", out.String())
}


// ---

// ### 2️⃣ Dosyaya Yazma ve Okuma


import "os"

func main() {
	data := []byte("Dosyaya yazılacak zlib verisi.")

	// Sıkıştırılmış dosya oluşturma
	file, _ := os.Create("data.zlib")
	writer := zlib.NewWriter(file)
	writer.Write(data)
	writer.Close()
	file.Close()

	// Dosyayı açma ve açma işlemi
	f, _ := os.Open("data.zlib")
	reader, _ := zlib.NewReader(f)
	out := new(bytes.Buffer)
	io.Copy(out, reader)
	reader.Close()
	f.Close()

	fmt.Println("Açılmış veri:", out.String())
}


// ---

// ### 3️⃣ Reader ve Writer Arasında Doğrudan Veri Transferi


input := []byte("Uzun bir metin bloğu zlib ile sıkıştırılacak.")

var compressed bytes.Buffer
writer := zlib.NewWriter(&compressed)
writer.Write(input)
writer.Close()

reader, _ := zlib.NewReader(&compressed)
decompressed := new(bytes.Buffer)
io.Copy(decompressed, reader)
reader.Close()

fmt.Println("Açılmış veri:", decompressed.String())

/*
---

# ⚡ Özet

| Yapı / Fonksiyon          | Açıklama                           |
| ------------------------- | ---------------------------------- |
| `zlib.NewWriter`          | Sıkıştırıcı oluşturur              |
| `Writer.Write`            | Veriyi sıkıştırır                  |
| `Writer.Close`            | Sıkıştırmayı bitirir ve flush eder |
| `zlib.NewReader`          | Açıcı oluşturur                    |
| `Reader.Read` / `io.Copy` | Veriyi açar                        |

---

### 📌 Notlar

* `compress/zlib` **gzip ile benzerdir**, fakat **farklı header ve checksum** kullanır.
* `compress/zlib` kullanımı **flate algoritmasını temel alır**, dolayısıyla `compress/flate` ile uyumludur.
* Network veya dosya sıkıştırmalarında çok yaygın kullanılır.

---

İstersen sana **zlib için interaktif bir demo uygulaması** hazırlayıp, kullanıcı metin girip sıkıştırma ve açma işlemini görsün ve boyut farklarını karşılaştırsın şeklinde bir örnek de yazabilirim.

Bunu hazırlayayım mı? Evet
*/
//----------------------
/*
Harika! O zaman Go’da **interaktif bir zlib demo uygulaması** hazırlayalım:

* Kullanıcı metin girecek
* Program girilen metni zlib ile sıkıştıracak ve açacak
* Sıkıştırılmış boyutu ve açılmış veriyi gösterecek

---

## 📝 Uygulama: `zlib_cli_demo.go`
*/

package main

import (
	"bufio"
	"bytes"
	"compress/zlib"
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
		fmt.Println("\nSeçenekler:")
		fmt.Println("1. Zlib ile sıkıştır ve aç")
		fmt.Println("2. Çıkış")
		fmt.Print("Seçiminiz: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)

		switch choiceStr {
		case "1":
			runZlib(data)
		case "2":
			fmt.Println("Çıkış yapılıyor...")
			return
		default:
			fmt.Println("Geçersiz seçim!")
		}
	}
}

///////////////////// Fonksiyon /////////////////////

func runZlib(data []byte) {
	var buf bytes.Buffer

	// Sıkıştırma
	writer := zlib.NewWriter(&buf)
	writer.Write(data)
	writer.Close()

	fmt.Println("Sıkıştırılmış boyut:", buf.Len())

	// Açma
	reader, _ := zlib.NewReader(&buf)
	defer reader.Close()
	out := new(bytes.Buffer)
	io.Copy(out, reader)

	fmt.Println("Açılmış veri:", out.String())
}

/*
---

### 📌 Açıklamalar

1. Kullanıcıdan metin alınır (`bufio.Reader`)
2. Menü ile zlib sıkıştırma seçeneği sunulur
3. Seçilen işlemler yapılır:

   * Zlib Writer ile veri sıkıştırılır
   * Sıkıştırılmış boyut gösterilir
   * Zlib Reader ile veri açılır ve ekrana yazılır
4. Çıkış seçeneği ile program sonlanır

---

### 📌 Örnek Çalışma

```
Bir metin girin: Merhaba Go Zlib CLI!

Seçenekler:
1. Zlib ile sıkıştır ve aç
2. Çıkış
Seçiminiz: 1
Sıkıştırılmış boyut: 46
Açılmış veri: Merhaba Go Zlib CLI!

Seçiminiz: 2
Çıkış yapılıyor...
```

---

Bu uygulama ile kullanıcı:

* Zlib algoritmasını deneyimleyebilir
* Sıkıştırılmış boyutu ve açılmış veriyi görebilir

---
*/
