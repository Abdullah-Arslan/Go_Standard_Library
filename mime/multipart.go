/*
**Go dilindeki `mime/multipart` paketini** detaylı şekilde açıklayacağım ve örneklerle tüm kullanımını göstereceğim.

---

## 📦 `mime/multipart` Paketi Nedir?

`mime/multipart` paketi, Go’da **MIME çok parçalı (multipart) mesajlarını** işlemek için kullanılır.
Bu paket özellikle **HTTP dosya yüklemeleri (file upload)** ve **e-posta ekleri** gibi senaryolarda kullanılır.

* HTTP’de `multipart/form-data` formatındaki istekleri işler.
* Form verilerini ve dosya eklerini ayrıştırabilir.
* Dosya ekleri oluşturup göndermeye de olanak sağlar.

Dokümantasyon: [pkg.go.dev/mime/multipart](https://pkg.go.dev/mime/multipart)

---

## 🔧 Paket İçindeki Temel Yapılar ve Fonksiyonlar

### 1. `Reader` ve `Writer` Yapıları

* **`multipart.Reader`**: Çok parçalı verileri okumak için kullanılır.
* **`multipart.Writer`**: Çok parçalı veri oluşturmak (form ve dosya ekleri) için kullanılır.

---

### 2. `NewReader(r io.Reader, boundary string) *Reader`

* Verilen `io.Reader` ve boundary kullanarak multipart veriyi okur.

**Örnek:**
*/
``go
package main

import (
	"fmt"
	"log"
	"mime/multipart"
	"strings"
)

func main() {
	const data = `--boundary
Content-Disposition: form-data; name="field1"

value1
--boundary--`

	r := multipart.NewReader(strings.NewReader(data), "boundary")
	part, err := r.NextPart()
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	n, _ := part.Read(buf)
	fmt.Println("Part Data:", string(buf[:n]))
}
``
/*
---

### 3. `NewWriter(w io.Writer) *Writer`

* Çok parçalı veri oluşturmak için kullanılır.
* Form alanları ve dosya ekleri eklenebilir.

**Örnek:**
*/
``go
package main

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
)

func main() {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	// Form alanı ekleme
	fw, err := w.CreateFormField("name")
	if err != nil {
		log.Fatal(err)
	}
	fw.Write([]byte("Golang"))

	// Dosya ekleme
	fw2, err := w.CreateFormFile("file", "example.txt")
	if err != nil {
		log.Fatal(err)
	}
	fw2.Write([]byte("Hello World"))

	w.Close()
	fmt.Println("Multipart Data:\n", b.String())
	fmt.Println("Boundary:", w.Boundary())
}
``
/*
---

### 4. `CreateFormFile` ve `CreateFormField`

* **`CreateFormFile(fieldname, filename string) (io.Writer, error)`**: Form-data içine dosya ekler.
* **`CreateFormField(fieldname string) (io.Writer, error)`**: Form-data içine normal alan ekler.

---

### 5. `NextPart()` – Bölümleri Okuma

* `multipart.Reader` üzerinden sıradaki bölümü okur.
* Bölümler dosya veya form alanı olabilir.

**Örnek:**
*/
``go
package main

import (
	"fmt"
	"log"
	"mime/multipart"
	"strings"
)

func main() {
	const data = `--boundary
Content-Disposition: form-data; name="username"

Alice
--boundary
Content-Disposition: form-data; name="file"; filename="hello.txt"

Hello Go
--boundary--`

	r := multipart.NewReader(strings.NewReader(data), "boundary")
	for {
		part, err := r.NextPart()
		if err != nil {
			break
		}
		buf := make([]byte, 1024)
		n, _ := part.Read(buf)
		fmt.Printf("Part Name: %s, Data: %s\n", part.FormName(), string(buf[:n]))
	}
}
``
/*
---

### 🔹 Özet Fonksiyon ve Yapılar
*/

| Fonksiyon / Yapı                                    | Açıklama                                                                       |
| --------------------------------------------------- | ------------------------------------------------------------------------------ |
| `multipart.NewReader(r io.Reader, boundary string)` | Çok parçalı veriyi okumak için reader oluşturur                                |
| `multipart.NewWriter(w io.Writer)`                  | Çok parçalı veri yazmak için writer oluşturur                                  |
| `CreateFormFile`                                    | Form-data içine dosya ekler                                                    |
| `CreateFormField`                                   | Form-data içine normal alan ekler                                              |
| `NextPart()`                                        | Multipart reader ile sıradaki bölümü okur                                      |
| `Part`                                              | Multipart içindeki tek bir bölümü temsil eder (`FormName`, `FileName`, `Read`) |

/*
---

İstersen ben şimdi **tüm bu `mime/multipart` örneklerini tek bir çalışma seti hâline getirip, eksiksiz ve çalıştırılabilir Go kodu** olarak hazırlayabilirim.

Bunu yapmamı ister misin?
EVET
*/

//Golang Multipart Examples 
package main

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"strings"
)

func main() {
	// 1️⃣ Multipart Reader Örneği
	fmt.Println("--- Multipart Reader Örneği ---")
	const data = `--boundary
Content-Disposition: form-data; name="field1"

value1
--boundary--`

	r := multipart.NewReader(strings.NewReader(data), "boundary")
	part, err := r.NextPart()
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	n, _ := part.Read(buf)
	fmt.Println("Part Data:", string(buf[:n]))

	// 2️⃣ Multipart Writer Örneği
	fmt.Println("--- Multipart Writer Örneği ---")
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	// Form alanı ekleme
	fw, err := w.CreateFormField("name")
	if err != nil {
		log.Fatal(err)
	}
	fw.Write([]byte("Golang"))

	// Dosya ekleme
	fw2, err := w.CreateFormFile("file", "example.txt")
	if err != nil {
		log.Fatal(err)
	}
	fw2.Write([]byte("Hello World"))

	w.Close()
	fmt.Println("Multipart Data:\n", b.String())
	fmt.Println("Boundary:", w.Boundary())

	// 3️⃣ NextPart ile tüm bölümleri okuma
	fmt.Println("--- NextPart ile Tüm Bölümleri Okuma ---")
	const multiData = `--boundary
Content-Disposition: form-data; name="username"

Alice
--boundary
Content-Disposition: form-data; name="file"; filename="hello.txt"

Hello Go
--boundary--`

	r2 := multipart.NewReader(strings.NewReader(multiData), "boundary")
	for {
		part, err := r2.NextPart()
		if err != nil {
			break
		}
		buf := make([]byte, 1024)
		n, _ := part.Read(buf)
		fmt.Printf("Part Name: %s, Data: %s\n", part.FormName(), string(buf[:n]))
	}
}

/*
Tamamdır! Tüm `mime/multipart` örneklerini **çalıştırılabilir tek bir Go dosyası** olarak hazırladım. Artık bu dosyayı çalıştırarak:

* Multipart reader ile tek parça okuma
* Multipart writer ile form ve dosya ekleme
* `NextPart` ile tüm bölümleri okuma

işlemlerini deneyimleyebilirsin.
*/