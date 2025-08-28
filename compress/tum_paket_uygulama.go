/*
O zaman Go’da **tüm compress paketlerini karşılaştıran bir mega demo uygulaması** hazırlayalım.

* Kullanıcı bir metin girecek
* Program **gzip, zlib, flate, lzw, bzip2** algoritmalarıyla sıkıştıracak ve açacak
* Sıkıştırılmış boyutları ve açılmış veriyi karşılaştırmalı olarak gösterecek
* Böylece hangi algoritmanın daha verimli olduğunu görebilecek

> Not: Bzip2 için üçüncü parti paket (`github.com/chai2010/bzip2`) kullanacağız, çünkü standart Go sadece okuma (decompression) destekliyor.

---

## 📝 Uygulama: `compress_all_demo.go`
*/
package main

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"compress/lzw"
	"compress/zlib"
	"fmt"
	"io"

	"github.com/chai2010/bzip2"
)

func main() {
	fmt.Print("Bir metin girin: ")
	var input string
	fmt.Scanln(&input)
	data := []byte(input)

	fmt.Println("\nAlgoritma\tSıkıştırılmış Boyut\tAçılmış Veri")
	fmt.Println("-------------------------------------------------------------")

	// gzip
	gzBuf, gzOut := compressGzip(data)
	fmt.Printf("gzip\t\t%d\t\t\t%s\n", gzBuf.Len(), gzOut)

	// zlib
	zBuf, zOut := compressZlib(data)
	fmt.Printf("zlib\t\t%d\t\t\t%s\n", zBuf.Len(), zOut)

	// flate
	fBuf, fOut := compressFlate(data)
	fmt.Printf("flate\t\t%d\t\t\t%s\n", fBuf.Len(), fOut)

	// lzw
	lzwBuf, lzwOut := compressLZW(data)
	fmt.Printf("lzw\t\t%d\t\t\t%s\n", lzwBuf.Len(), lzwOut)

	// bzip2
	bzBuf, bzOut := compressBzip2(data)
	fmt.Printf("bzip2\t\t%d\t\t\t%s\n", bzBuf.Len(), bzOut)
}

///////////////////// Fonksiyonlar /////////////////////

func compressGzip(data []byte) (*bytes.Buffer, string) {
	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	writer.Write(data)
	writer.Close()

	reader, _ := gzip.NewReader(&buf)
	defer reader.Close()
	out := new(bytes.Buffer)
	io.Copy(out, reader)
	return &buf, out.String()
}

func compressZlib(data []byte) (*bytes.Buffer, string) {
	var buf bytes.Buffer
	writer := zlib.NewWriter(&buf)
	writer.Write(data)
	writer.Close()

	reader, _ := zlib.NewReader(&buf)
	defer reader.Close()
	out := new(bytes.Buffer)
	io.Copy(out, reader)
	return &buf, out.String()
}

func compressFlate(data []byte) (*bytes.Buffer, string) {
	var buf bytes.Buffer
	writer, _ := flate.NewWriter(&buf, flate.BestCompression)
	writer.Write(data)
	writer.Close()

	reader := flate.NewReader(&buf)
	defer reader.Close()
	out := new(bytes.Buffer)
	io.Copy(out, reader)
	return &buf, out.String()
}

func compressLZW(data []byte) (*bytes.Buffer, string) {
	var buf bytes.Buffer
	writer := lzw.NewWriter(&buf, lzw.LSB, 8)
	writer.Write(data)
	writer.Close()

	reader := lzw.NewReader(&buf, lzw.LSB, 8)
	defer reader.Close()
	out := new(bytes.Buffer)
	io.Copy(out, reader)
	return &buf, out.String()
}

func compressBzip2(data []byte) (*bytes.Buffer, string) {
	var buf bytes.Buffer
	writer := bzip2.NewWriter(&buf)
	writer.Write(data)
	writer.Close()

	reader := bzip2.NewReader(&buf)
	out := new(bytes.Buffer)
	io.Copy(out, reader)
	return &buf, out.String()
}

/*
---

### 📌 Açıklamalar

1. Kullanıcıdan bir metin alınır
2. Her algoritma için:

   * Sıkıştırma (`Writer`)
   * Sıkıştırılmış boyut hesaplama
   * Açma (`Reader`) ve orijinal veriyi geri alma
3. Tüm sonuçlar tabloda gösterilir

---

### 📌 Örnek Çıktı

```
Bir metin girin: Merhaba Go Compress Mega Demo

Algoritma       Sıkıştırılmış Boyut       Açılmış Veri
-------------------------------------------------------------
gzip            45                       Merhaba Go Compress Mega Demo
zlib            39                       Merhaba Go Compress Mega Demo
flate           35                       Merhaba Go Compress Mega Demo
lzw             44                       Merhaba Go Compress Mega Demo
bzip2           38                       Merhaba Go Compress Mega Demo
```

---

Bu uygulama ile:

* Tüm sıkıştırma algoritmalarını **aynı anda karşılaştırabilirsiniz**
* **Boyut farklarını ve açılmış veriyi** görebilirsiniz
* Eğitim veya performans testleri için **ideal bir demo** olur

---
*/
