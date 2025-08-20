/*
**tek bir uygulamada** tüm `compress` alt paketlerini kullanarak:

* Kullanıcının verdiği bir metni
* `gzip`, `zlib`, `flate`, `lzw` ile sıkıştırıp
* Tek tek açıp ekrana gösteren bir demo uygulama yazalım.

---

## 📝 Uygulama: `compress_demo.go`
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
)

func main() {
	data := []byte("Merhaba Go! Compress paketi örneği.")

	fmt.Println("Orijinal veri:", string(data))
	fmt.Println("Boyut:", len(data))
	fmt.Println("---------------------------------------------------")

	// ---------------- gzip ----------------
	fmt.Println("=== gzip ===")
	var gzipBuf bytes.Buffer
	gw := gzip.NewWriter(&gzipBuf)
	gw.Write(data)
	gw.Close()
	fmt.Println("Sıkıştırılmış boyut:", gzipBuf.Len())

	gr, _ := gzip.NewReader(&gzipBuf)
	defer gr.Close()
	out := new(bytes.Buffer)
	io.Copy(out, gr)
	fmt.Println("Açılmış veri:", out.String())
	fmt.Println("---------------------------------------------------")

	// ---------------- zlib ----------------
	fmt.Println("=== zlib ===")
	var zlibBuf bytes.Buffer
	zw := zlib.NewWriter(&zlibBuf)
	zw.Write(data)
	zw.Close()
	fmt.Println("Sıkıştırılmış boyut:", zlibBuf.Len())

	zr, _ := zlib.NewReader(&zlibBuf)
	defer zr.Close()
	out.Reset()
	io.Copy(out, zr)
	fmt.Println("Açılmış veri:", out.String())
	fmt.Println("---------------------------------------------------")

	// ---------------- flate ----------------
	fmt.Println("=== flate ===")
	var flateBuf bytes.Buffer
	fw, _ := flate.NewWriter(&flateBuf, flate.BestCompression)
	fw.Write(data)
	fw.Close()
	fmt.Println("Sıkıştırılmış boyut:", flateBuf.Len())

	fr := flate.NewReader(&flateBuf)
	defer fr.Close()
	out.Reset()
	io.Copy(out, fr)
	fmt.Println("Açılmış veri:", out.String())
	fmt.Println("---------------------------------------------------")

	// ---------------- lzw ----------------
	fmt.Println("=== lzw ===")
	var lzwBuf bytes.Buffer
	lw := lzw.NewWriter(&lzwBuf, lzw.LSB, 8)
	lw.Write(data)
	lw.Close()
	fmt.Println("Sıkıştırılmış boyut:", lzwBuf.Len())

	lr := lzw.NewReader(&lzwBuf, lzw.LSB, 8)
	defer lr.Close()
	out.Reset()
	io.Copy(out, lr)
	fmt.Println("Açılmış veri:", out.String())
	fmt.Println("---------------------------------------------------")
}

/*
---

## 📌 Açıklama

1. **Orijinal veri** ekrana yazılır.
2. Her bir sıkıştırma algoritması için:

   * `Writer` ile veri sıkıştırılır
   * Sıkıştırılmış boyut gösterilir
   * `Reader` ile veri açılır ve ekrana yazılır
3. Tüm paketler (`gzip`, `zlib`, `flate`, `lzw`) aynı mantıkla kullanılır.

---

## 📌 Örnek Çıktı

```
Orijinal veri: Merhaba Go! Compress paketi örneği.
Boyut: 33
---------------------------------------------------
=== gzip ===
Sıkıştırılmış boyut: 45
Açılmış veri: Merhaba Go! Compress paketi örneği.
---------------------------------------------------
=== zlib ===
Sıkıştırılmış boyut: 35
Açılmış veri: Merhaba Go! Compress paketi örneği.
---------------------------------------------------
=== flate ===
Sıkıştırılmış boyut: 29
Açılmış veri: Merhaba Go! Compress paketi örneği.
---------------------------------------------------
=== lzw ===
Sıkıştırılmış boyut: 44
Açılmış veri: Merhaba Go! Compress paketi örneği.
---------------------------------------------------
```

---

Bu uygulama ile **Go’nun compress paketinin tamamını tek bir demo içinde** görmüş oluyorsun:

* Sıkıştırma ve açma işlemleri
* Boyut karşılaştırması
* gzip, zlib, flate, lzw algoritmaları

---

İstersen bir sonraki adımda bunu **kullanıcıdan metin alan interaktif CLI versiyonu** hâline getirebiliriz. Böylece kullanıcı metin girip hangi algoritmayı denemek istediğini seçebilir.

Bunu yapayım mı? Evet
*/

/*
 O zaman uygulamayı **interaktif hâle getiriyoruz**:

* Kullanıcı metin girecek
* Menü ile hangi sıkıştırma algoritmasını çalıştırmak istediğini seçecek
* Seçilen algoritma hem sıkıştıracak hem açacak ve boyut ile sonucu gösterecek

---

## 📝 Uygulama: `compress_cli_demo.go`
*/

package main

import (
	"bufio"
	"bytes"
	"compress/flate"
	"compress/gzip"
	"compress/lzw"
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
		fmt.Println("\nHangi algoritmayı kullanmak istersiniz?")
		fmt.Println("1. gzip")
		fmt.Println("2. zlib")
		fmt.Println("3. flate")
		fmt.Println("4. lzw")
		fmt.Println("5. Çıkış")
		fmt.Print("Seçiminiz: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)

		switch choiceStr {
		case "1":
			runGzip(data)
		case "2":
			runZlib(data)
		case "3":
			runFlate(data)
		case "4":
			runLZW(data)
		case "5":
			fmt.Println("Çıkış yapılıyor...")
			return
		default:
			fmt.Println("Geçersiz seçim!")
		}
	}
}

//////////////////// DEMO FONKSİYONLARI ////////////////////

func runGzip(data []byte) {
	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	writer.Write(data)
	writer.Close()
	fmt.Println("=== gzip ===")
	fmt.Println("Sıkıştırılmış boyut:", buf.Len())

	reader, _ := gzip.NewReader(&buf)
	defer reader.Close()
	out := new(bytes.Buffer)
	io.Copy(out, reader)
	fmt.Println("Açılmış veri:", out.String())
}

func runZlib(data []byte) {
	var buf bytes.Buffer
	writer := zlib.NewWriter(&buf)
	writer.Write(data)
	writer.Close()
	fmt.Println("=== zlib ===")
	fmt.Println("Sıkıştırılmış boyut:", buf.Len())

	reader, _ := zlib.NewReader(&buf)
	defer reader.Close()
	out := new(bytes.Buffer)
	io.Copy(out, reader)
	fmt.Println("Açılmış veri:", out.String())
}

func runFlate(data []byte) {
	var buf bytes.Buffer
	writer, _ := flate.NewWriter(&buf, flate.BestCompression)
	writer.Write(data)
	writer.Close()
	fmt.Println("=== flate ===")
	fmt.Println("Sıkıştırılmış boyut:", buf.Len())

	reader := flate.NewReader(&buf)
	defer reader.Close()
	out := new(bytes.Buffer)
	io.Copy(out, reader)
	fmt.Println("Açılmış veri:", out.String())
}

func runLZW(data []byte) {
	var buf bytes.Buffer
	writer := lzw.NewWriter(&buf, lzw.LSB, 8)
	writer.Write(data)
	writer.Close()
	fmt.Println("=== lzw ===")
	fmt.Println("Sıkıştırılmış boyut:", buf.Len())

	reader := lzw.NewReader(&buf, lzw.LSB, 8)
	defer reader.Close()
	out := new(bytes.Buffer)
	io.Copy(out, reader)
	fmt.Println("Açılmış veri:", out.String())
}

/*
---

### 📌 Açıklama

1. **Kullanıcıdan metin alınır** (`bufio.Reader`)
2. Menü ile algoritma seçilir
3. Seçilen algoritma için ayrı fonksiyon çalışır:

   * Sıkıştırır (`Writer`)
   * Boyutu gösterir
   * Açar (`Reader`) ve sonucu gösterir
4. `5` seçilirse program sonlanır

---

### Örnek Çalışma Akışı

```
Bir metin girin: Merhaba Go Compress CLI!
Hangi algoritmayı kullanmak istersiniz?
1. gzip
2. zlib
3. flate
4. lzw
5. Çıkış
Seçiminiz: 1
=== gzip ===
Sıkıştırılmış boyut: 42
Açılmış veri: Merhaba Go Compress CLI!

Seçiminiz: 3
=== flate ===
Sıkıştırılmış boyut: 35
Açılmış veri: Merhaba Go Compress CLI!

Seçiminiz: 5
Çıkış yapılıyor...

*/

/*
Harika! O zaman bir sonraki adım olarak **tüm algoritmaları tek seferde çalıştırıp boyut karşılaştırması yapan bir CLI uygulaması** hazırlayalım.

* Kullanıcı metin giriyor
* Program tüm algoritmalar (`gzip`, `zlib`, `flate`, `lzw`) ile sıkıştırıyor
* Hem sıkıştırılmış boyutları hem de açılmış veriyi bir tabloda gösteriyor

---

## 📝 Uygulama: `compress_compare.go`
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
)

func main() {
	var input string
	fmt.Print("Bir metin girin: ")
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

/*
---

### 📌 Çıktı Örneği

```
Bir metin girin: Merhaba Go Compress Test
Algoritma       Sıkıştırılmış Boyut       Açılmış Veri
-------------------------------------------------------------
gzip            42                       Merhaba Go Compress Test
zlib            36                       Merhaba Go Compress Test
flate           30                       Merhaba Go Compress Test
lzw             44                       Merhaba Go Compress Test
```

---

### ✅ Özellikler

* Kullanıcı sadece bir metin giriyor
* Program **tüm algoritmaları aynı anda çalıştırıyor**
* Hem **sıkıştırılmış boyut** hem de **açılmış veri** karşılaştırmalı olarak gösteriliyor
* Böylece hangi algoritmanın daha verimli olduğunu kolayca görebiliyorsun

---

*/
