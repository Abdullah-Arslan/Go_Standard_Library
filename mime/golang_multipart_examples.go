/*
İstersen ben şimdi **tüm bu `mime/multipart` örneklerini tek bir çalışma seti hâline getirip, eksiksiz ve çalıştırılabilir Go kodu** olarak hazırlayabilirim.

Bunu yapmamı ister misin?
*/

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
