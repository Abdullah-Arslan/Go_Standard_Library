// Golang Quotedprintable Examples
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
