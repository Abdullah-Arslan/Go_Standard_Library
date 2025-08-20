/*
Go’da **`bzip2` paketi**, **Bzip2 sıkıştırma algoritmasını** kullanmak için standart kütüphanedeki pakettir.

* Paket: `compress/bzip2`
* Ama önemli bir detay: Go’nun standart `compress/bzip2` paketi **sadece okuma (`Reader`) için** uygundur; yazma/sıkıştırma (`Writer`) için standart Go kütüphanesinde destek yoktur.
* Bzip2, DEFLATE veya gzip’e göre **daha yüksek sıkıştırma oranı** sağlar ancak genellikle biraz daha yavaştır.

---

# 📦 `compress/bzip2` Paketinin Temel Yapısı

* `bzip2.NewReader(r io.Reader) io.Reader` → bir `io.Reader` döndürür.
* Bu reader, **Bzip2 ile sıkıştırılmış veriyi açmak** için kullanılır.

> ⚠️ Yazma/sıkıştırma (`Writer`) Go standart kütüphanesinde yok; bunu kullanmak için genellikle **3. parti paketler** (`github.com/dsnet/compress/bzip2`) tercih edilir.

---

## 1️⃣ Bzip2 Okuma Örneği
*/

package main

import (
	"bytes"
	"compress/bzip2"
	"fmt"
	"io"
)

func main() {
	// Örnek bzip2 ile sıkıştırılmış veri (örnek olarak byte slice)
	// Gerçek uygulamada genellikle dosya açılır
	data := []byte{
		0x42, 0x5a, 0x68, 0x39, 0x31, 0x41, 0x59, 0x26, // BZh91AY&...
		// ... bzip2 sıkıştırılmış veri devam eder
	}

	reader := bzip2.NewReader(bytes.NewReader(data))
	out := new(bytes.Buffer)
	_, err := io.Copy(out, reader)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Println("Açılmış veri:", out.String())
}

/*
> Not: Yukarıdaki veri sadece bir örnek. Gerçek veri genellikle `.bz2` dosyasından okunur.

---

## 2️⃣ Bzip2 Dosya Açma Örneği
*/

package main

import (
	"compress/bzip2"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("example.bz2")
	if err != nil {
		fmt.Println("Dosya açılamadı:", err)
		return
	}
	defer f.Close()

	reader := bzip2.NewReader(f)
	out := new(bytes.Buffer)
	_, err = io.Copy(out, reader)
	if err != nil {
		fmt.Println("Okuma hatası:", err)
		return
	}

	fmt.Println("Bzip2 Açılmış Veri:")
	fmt.Println(out.String())
}

/*
---

## 3️⃣ Özet

| Özellik        | Açıklama                                           |
| -------------- | -------------------------------------------------- |
| Paket          | `compress/bzip2`                                   |
| Desteklenen    | Okuma (`Reader`)                                   |
| Sıkıştırma     | Hayır, standart Go’da yok                          |
| Kullanım alanı | `.bz2` dosyalarını açmak                           |
| Alternatif     | `github.com/dsnet/compress/bzip2` (okuma ve yazma) |

---

### ⚡ Öneri

Eğer **Bzip2 ile hem sıkıştırma hem açma** yapmak istiyorsan, standart paket yerine `github.com/dsnet/compress/bzip2` gibi üçüncü taraf paketleri kullanmak gerekir.

---

İstersen sana **Go’da hem okuma hem yazma yapabilen Bzip2 tam demo uygulaması** hazırlayabilirim, üçüncü parti paket kullanarak.

Bunu hazırlayayım mı? Evet
*/

