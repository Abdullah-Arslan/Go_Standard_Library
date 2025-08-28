//  **ZIP** hem de **TAR** için bir uygulama yazalım.
// Sen sadece bir klasör vereceksin, program:

// 1. O klasörü **ZIP** ve **TAR** dosyası haline getirecek.
// 2. Sonra bu arşivleri açıp içindekileri gösterecek.

// ---

// ## 📝 Tam Uygulama: `main.go`

package main

import (
	"archive/tar"
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func main() {
	klasor := "ornek_klasor"

	// ZIP oluştur
	err := zipOlustur("arsiv.zip", klasor)
	if err != nil {
		fmt.Println("ZIP hata:", err)
		return
	}
	fmt.Println("arsiv.zip oluşturuldu ✅")

	// TAR oluştur
	err = tarOlustur("arsiv.tar", klasor)
	if err != nil {
		fmt.Println("TAR hata:", err)
		return
	}
	fmt.Println("arsiv.tar oluşturuldu ✅")

	// ZIP aç
	fmt.Println("\nZIP içeriği:")
	zipAc("arsiv.zip")

	// TAR aç
	fmt.Println("\nTAR içeriği:")
	tarAc("arsiv.tar")
}

// //////////////////// ZIP ////////////////////

func zipOlustur(zipDosya, klasor string) error {
	zipFile, err := os.Create(zipDosya)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zw := zip.NewWriter(zipFile)
	defer zw.Close()

	filepath.Walk(klasor, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		relPath, _ := filepath.Rel(klasor, path)
		f, _ := zw.Create(relPath)

		src, _ := os.Open(path)
		defer src.Close()
		io.Copy(f, src)
		return nil
	})

	return nil
}

func zipAc(zipDosya string) {
	r, err := zip.OpenReader(zipDosya)
	if err != nil {
		fmt.Println("ZIP açma hata:", err)
		return
	}
	defer r.Close()

	for _, f := range r.File {
		fmt.Println("📄", f.Name)
		rc, _ := f.Open()
		defer rc.Close()
		io.Copy(os.Stdout, rc)
		fmt.Println()
	}
}

// //////////////////// TAR ////////////////////

func tarOlustur(tarDosya, klasor string) error {
	tarFile, err := os.Create(tarDosya)
	if err != nil {
		return err
	}
	defer tarFile.Close()

	tw := tar.NewWriter(tarFile)
	defer tw.Close()

	filepath.Walk(klasor, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		relPath, _ := filepath.Rel(klasor, path)
		header := &tar.Header{
			Name:    relPath,
			Size:    info.Size(),
			Mode:    int64(info.Mode()),
			ModTime: time.Now(),
		}

		tw.WriteHeader(header)

		src, _ := os.Open(path)
		defer src.Close()
		io.Copy(tw, src)
		return nil
	})

	return nil
}

func tarAc(tarDosya string) {
	file, err := os.Open(tarDosya)
	if err != nil {
		fmt.Println("TAR açma hata:", err)
		return
	}
	defer file.Close()

	tr := tar.NewReader(file)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("TAR okuma hata:", err)
			break
		}

		fmt.Println("📄", header.Name)
		buf := make([]byte, header.Size)
		tr.Read(buf)
		fmt.Println(string(buf))
	}
}

// ```

// ---

// ## 📂 Örnek Kullanım

// Diyelim `ornek_klasor/` içinde şu dosyalar var:

// ```
// ornek_klasor/
//  ├─ a.txt   (içinde: Merhaba A)
//  ├─ b.txt   (içinde: Merhaba B)
// ```

// Program çalıştırıldığında:

// * `arsiv.zip` ve `arsiv.tar` oluşturulur
// * Ekranda şu görünür:

// ```
// arsiv.zip oluşturuldu ✅
// arsiv.tar oluşturuldu ✅

// ZIP içeriği:
// 📄 a.txt
// Merhaba A
// 📄 b.txt
// Merhaba B

// TAR içeriği:
// 📄 a.txt
// Merhaba A
// 📄 b.txt
// Merhaba B
// ```

// ---
