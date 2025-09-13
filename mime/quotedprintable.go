/*
**Go dilindeki `mime/quotedprintable` paketini** detaylı bir şekilde açıklayacağım ve örneklerle tüm kullanımını göstereceğim.

---

## 📦 `mime/quotedprintable` Paketi Nedir?

`mime/quotedprintable` paketi, **quoted-printable (QP) kodlamasını** okumak ve yazmak için kullanılır.

* Quoted-printable, e-posta mesajlarında veya MIME verilerinde **ASCII karakterleri korurken, özel karakterleri güvenli bir şekilde temsil etmek** için kullanılır.
* Özellikle 8-bit karakterlerin veya uzun satırların HTTP veya SMTP üzerinden iletilmesini sağlar.

Dokümantasyon: [pkg.go.dev/mime/quotedprintable](https://pkg.go.dev/mime/quotedprintable)

---

## 🔧 Paket İçindeki Temel Yapılar ve Fonksiyonlar

### 1. `Reader`

* `quotedprintable.NewReader(r io.Reader) io.Reader`
* Quoted-printable veriyi çözmek için bir reader döndürür.

**Örnek:**
*/
``go
package main

import (
	"fmt"
	"io"
	"mime/quotedprintable"
	"strings"
)

func main() {
	const qpData = "Hello=20World=21"
	r := quotedprintable.NewReader(strings.NewReader(qpData))
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Println(buf.String()) // Çıktı: Hello World!
}
``
/*
---

### 2. `Writer`

* `quotedprintable.NewWriter(w io.Writer) io.WriteCloser`
* Quoted-printable veriyi yazmak için kullanılır.

**Örnek:**
*/
``go
package main

import (
	"fmt"
	"mime/quotedprintable"
	"strings"
)

func main() {
	var sb strings.Builder
	w := quotedprintable.NewWriter(&sb)
	w.Write([]byte("Hello World! Çalışıyor")) // Özel karakterleri QP ile kodlar
	w.Close()
	fmt.Println(sb.String())
}
``
/*
---

### 3. Kullanım Adımları

1. **Okuma (Decode)**: `quotedprintable.NewReader` ile QP veriyi decode edersiniz.
2. **Yazma (Encode)**: `quotedprintable.NewWriter` ile QP formatında veri oluşturabilirsiniz.
3. Writer kapatılmadan veriyi tamamlamaz; `Close()` mutlaka çağrılmalı.

---

### 🔹 Örnek: Dosya ile QP Kodlama ve Çözme
*/
``go
package main

import (
	"fmt"
	"io"
	"mime/quotedprintable"
	"strings"
)

func main() {
	// Encode
	var sb strings.Builder
	w := quotedprintable.NewWriter(&sb)
	w.Write([]byte("Merhaba Dünya!"))
	w.Close()
	encoded := sb.String()
	fmt.Println("Encoded:", encoded)

	// Decode
	r := quotedprintable.NewReader(strings.NewReader(encoded))
	buf := new(strings.Builder)
	io.Copy(buf, r)
	fmt.Println("Decoded:", buf.String())
}
``

//---

### 🔹 Özet Fonksiyon ve Yapılar

| Yapı / Fonksiyon         | Açıklama                               |
| ------------------------ | -------------------------------------- |
| `NewReader(r io.Reader)` | Quoted-printable veriyi okur (decode)  |
| `NewWriter(w io.Writer)` | Quoted-printable veriyi yazar (encode) |
| `Write`                  | Writer’a veri yazar                    |
| `Close`                  | Writer’ı kapatır, kodlamayı tamamlar   |

/*
---

İstersen ben şimdi bu **tüm örnekleri tek bir çalıştırılabilir Go dosyası** hâline getirip, eksiksiz ve hazır bir set hazırlayabilirim.

Bunu yapmamı ister misin?
EVET
*/
//Golang Quotedprintable Examples

package main

import (
	"fmt"
	"io"
	"mime/quotedprintable"
	"strings"
)

func main() {
	// 1️⃣ Quoted-Printable Reader Örneği
	fmt.Println("--- Quoted-Printable Reader Örneği ---")
	const qpData = "Hello=20World=21"
	r := quotedprintable.NewReader(strings.NewReader(qpData))
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Println("Decoded:", buf.String()) // Hello World!

	// 2️⃣ Quoted-Printable Writer Örneği
	fmt.Println("--- Quoted-Printable Writer Örneği ---")
	var sb strings.Builder
	w := quotedprintable.NewWriter(&sb)
	w.Write([]byte("Hello World! Çalışıyor"))
	w.Close()
	fmt.Println("Encoded:", sb.String())

	// 3️⃣ Encode ve Decode Birlikte
	fmt.Println("--- Encode ve Decode ---")
	var sb2 strings.Builder
	w2 := quotedprintable.NewWriter(&sb2)
	w2.Write([]byte("Merhaba Dünya!"))
	w2.Close()
	encoded := sb2.String()
	fmt.Println("Encoded:", encoded)

	r2 := quotedprintable.NewReader(strings.NewReader(encoded))
	decodedBuf := new(strings.Builder)
	io.Copy(decodedBuf, r2)
	fmt.Println("Decoded:", decodedBuf.String())
}
/*
Tamamdır! Tüm `mime/quotedprintable` örneklerini **çalıştırılabilir tek bir Go dosyası** hâline getirdim. Artık bu dosya ile:

* Quoted-printable veri okuma (decode)
* Quoted-printable veri yazma (encode)
* Encode ve decode işlemlerini birlikte deneme

işlemlerini tek seferde deneyebilirsin.
*/