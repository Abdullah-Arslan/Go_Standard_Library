

// Go’da **`archive`** aslında **üst paket**tir. Bunun altında **iki alt paket** vardır:

// * **`archive/zip`** → ZIP formatı ile çalışır
// * **`archive/tar`** → TAR formatı ile çalışır

// Yani `archive` tek başına bir iş yapmaz, sadece bu iki paketi içerir. Şimdi ikisini ayrı ayrı **tam detay + örneklerle** açıklayayım:

// ---

// # 📦 `archive/tar` Paketi

// TAR (Tape Archive) formatı, özellikle Linux/Unix dünyasında kullanılır (`.tar`, `.tar.gz`).
// Go’nun `archive/tar` paketi **TAR dosyası oluşturma ve okuma** işlemlerini yapar.

// ### ✅ Önemli Yapılar

// * `tar.Writer` → TAR dosyası oluşturur, dosya ekler.
// * `tar.Reader` → TAR dosyasından dosya okur.
// * `tar.Header` → Dosyanın meta bilgilerini (isim, boyut, izinler vs.) tutar.

// ### 🔹 Örnek: TAR dosyası oluşturma


package main

import (
	"archive/tar"
	"fmt"
	"os"
	"time"
)

func main() {
	// 1. TAR dosyası oluştur
	file, _ := os.Create("arsiv.tar")
	defer file.Close()

	// 2. TAR writer
	tw := tar.NewWriter(file)
	defer tw.Close()

	// 3. Dosya başlığı (Header)
	header := &tar.Header{
		Name:    "ornek.txt",
		Size:    int64(len("Merhaba TAR!")),
		Mode:    0600,
		ModTime: time.Now(),
	}

	// 4. Header yaz
	tw.WriteHeader(header)

	// 5. İçeriği yaz
	tw.Write([]byte("Merhaba TAR!"))

	fmt.Println("arsiv.tar oluşturuldu ✅")
}
// ```

// ---

// ### 🔹 Örnek: TAR dosyasını okuma


package main

import (
	"archive/tar"
	"fmt"
	"os"
)

func main() {
	// TAR dosyasını aç
	file, _ := os.Open("arsiv.tar")
	defer file.Close()

	// TAR reader
	tr := tar.NewReader(file)

	// İçindeki dosyaları gez
	for {
		header, err := tr.Next()
		if err != nil {
			break
		}

		fmt.Println("Dosya:", header.Name)
		buf := make([]byte, header.Size)
		tr.Read(buf)
		fmt.Println("İçerik:", string(buf))
	}
}
// ```

// ---

// # 📦 `archive/zip` Paketi

// ZIP formatı Windows ve yaygın olarak kullanılan sıkıştırma formatıdır.
// Go’nun `archive/zip` paketi ZIP dosyası **oluşturma ve açma** işlemleri yapar.

// ### ✅ Önemli Yapılar

// * `zip.Writer` → ZIP dosyası oluşturur, dosya ekler.
// * `zip.Reader` → ZIP dosyası açar, dosyaları okur.
// * `zip.File` → ZIP içindeki her dosya hakkında bilgi tutar.

// ---

// ### 🔹 Örnek: ZIP dosyası oluşturma


package main

import (
	"archive/zip"
	"fmt"
	"os"
)

func main() {
	// ZIP dosyası oluştur
	zipFile, _ := os.Create("arsiv.zip")
	defer zipFile.Close()

	// ZIP writer
	zw := zip.NewWriter(zipFile)
	defer zw.Close()

	// ZIP içine dosya ekle
	f, _ := zw.Create("ornek.txt")
	f.Write([]byte("Merhaba ZIP!"))

	fmt.Println("arsiv.zip oluşturuldu ✅")
}
// ```

// ---

// ### 🔹 Örnek: ZIP dosyasını okuma


package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func main() {
	// ZIP dosyasını aç
	r, _ := zip.OpenReader("arsiv.zip")
	defer r.Close()

	// ZIP içindeki dosyaları gez
	for _, f := range r.File {
		fmt.Println("Dosya:", f.Name)

		rc, _ := f.Open()
		defer rc.Close()

		io.Copy(os.Stdout, rc) // İçeriği ekrana yaz
		fmt.Println()
	}
}

// ```

// ---

// # 🚀 Özet

// * **`archive/tar`** → TAR dosyalarını oluşturur ve okur.

//   * `tar.Writer` → Yazma
//   * `tar.Reader` → Okuma
//   * `tar.Header` → Dosya bilgisi

// * **`archive/zip`** → ZIP dosyalarını oluşturur ve okur.

//   * `zip.Writer` → Yazma
//   * `zip.Reader` → Okuma
//   * `zip.File` → Dosya bilgisi

// ---
