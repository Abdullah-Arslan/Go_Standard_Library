//  `bufio`’nun tüm yapılarının mantığını gösterecek **tam bir uygulama** yazalım:

// 📌 **Senaryo:**

// * `girdi.txt` dosyasını satır satır okuyacağız (`bufio.Scanner`)
// * Her satırın başına satır numarası ekleyeceğiz (`bufio.Writer`)
// * Yeni dosyaya (`cikti.txt`) yazacağız
// * Aynı zamanda yazma ve okuma işlemlerini `bufio.ReadWriter` ile yapacağız.

// ---

// ## 📝 Tam Uygulama: `main.go`

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 1. Girdi dosyasını aç
	girdi, err := os.Open("girdi.txt")
	if err != nil {
		fmt.Println("Dosya açılamadı:", err)
		return
	}
	defer girdi.Close()

	// 2. Çıktı dosyasını oluştur
	cikti, err := os.Create("cikti.txt")
	if err != nil {
		fmt.Println("Dosya oluşturulamadı:", err)
		return
	}
	defer cikti.Close()

	// 3. bufio yapıları
	scanner := bufio.NewScanner(girdi)                        // Satır satır okuyacak
	writer := bufio.NewWriter(cikti)                          // Yazacak
	rw := bufio.NewReadWriter(bufio.NewReader(girdi), writer) // Okuma + Yazma

	// 4. Satır satır oku ve yeni dosyaya yaz
	satirNo := 1
	for scanner.Scan() {
		satir := scanner.Text()
		yeniSatir := fmt.Sprintf("%d: %s\n", satirNo, satir)

		// bufio.Writer kullanarak yaz
		_, err := rw.WriteString(yeniSatir)
		if err != nil {
			fmt.Println("Yazma hatası:", err)
			return
		}

		satirNo++
	}

	// 5. Tampondaki veriyi dosyaya aktar
	err = rw.Flush()
	if err != nil {
		fmt.Println("Flush hatası:", err)
		return
	}

	// 6. Hata kontrolü (scanner için)
	if err := scanner.Err(); err != nil {
		fmt.Println("Okuma hatası:", err)
		return
	}

	fmt.Println("İşlem tamamlandı ✅ cikti.txt dosyası oluşturuldu.")
}

// ```

// ---

// ## 📂 Örnek Çalışma

// **`girdi.txt` içeriği:**

// ```
// Go çok hızlıdır
// bufio paketi verimlidir
// Dosya işlemleri kolaydır
// ```

// **`cikti.txt` sonucu:**

// ```
// 1: Go çok hızlıdır
// 2: bufio paketi verimlidir
// 3: Dosya işlemleri kolaydır
// ```

// ---

// Bu örnekte:

// * `bufio.Scanner` ile **satır satır okuduk** ✅
// * `bufio.Writer` ile **dosyaya yazdık** ✅
// * `bufio.ReadWriter` kullanarak ikisini birleştirdik ✅

// ---
