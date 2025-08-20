// Elbette! Go (Golang) dilinde .zip dosyaları ile çalışmak için standart kütüphane olan archive/zip paketi kullanılır. Bu paket, .zip arşivlerini oluşturmak ve açmak için kullanılır.

// Aşağıda, iki ana kullanım senaryosu ile birlikte açıklamalı örnekler bulacaksın:

//1. .zip Dosyası Oluşturma (Dosyaları Arşivleme)
//✅ Kod:
package main

import (
    "archive/zip"
    "fmt"
    "io"
    "os"
)

func main() {
    // Zip dosyası oluştur
    zipFile, err := os.Create("example.zip")
    if err != nil {
        panic(err)
    }
    defer zipFile.Close()

    // Zip yazıcısını oluştur
    zipWriter := zip.NewWriter(zipFile)
    defer zipWriter.Close()

    // Arşive eklemek istediğimiz dosyayı aç
    fileToZip := "test.txt"
    file, err := os.Open(fileToZip)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // Dosya bilgilerini al
    info, err := file.Stat()
    if err != nil {
        panic(err)
    }

    // Header oluştur
    header, err := zip.FileInfoHeader(info)
    if err != nil {
        panic(err)
    }
    header.Name = fileToZip // zip içindeki adı

    // Yazıcıyı başlat
    writer, err := zipWriter.CreateHeader(header)
    if err != nil {
        panic(err)
    }

    // Dosya içeriğini zip'e kopyala
    _, err = io.Copy(writer, file)
    if err != nil {
        panic(err)
    }

    fmt.Println("ZIP dosyası oluşturuldu: example.zip")
}

//🔹 Bu örnek, test.txt adlı dosyayı example.zip adında sıkıştırır.
//🔹 Eğer birden fazla dosya eklemek istersen, bu işlemi döngüye alabilirsin.

//2. .zip Dosyası Açma (Dosya Çıkartma)
//✅ Kod:
package main

import (
    "archive/zip"
    "fmt"
    "io"
    "os"
    "path/filepath"
)

func main() {
    // Zip dosyasını aç
    r, err := zip.OpenReader("example.zip")
    if err != nil {
        panic(err)
    }
    defer r.Close()

    // Her bir dosyayı gez
    for _, f := range r.File {
        fmt.Println("Çıkarılıyor:", f.Name)

        // Hedef dosya yolu
        outPath := filepath.Join("output", f.Name)

        // Eğer bu bir klasörse, oluştur
        if f.FileInfo().IsDir() {
            os.MkdirAll(outPath, os.ModePerm)
            continue
        }

        // Dosyayı aç
        rc, err := f.Open()
        if err != nil {
            panic(err)
        }
        defer rc.Close()

        // Hedef klasörü oluştur
        os.MkdirAll(filepath.Dir(outPath), os.ModePerm)

        // Dosyayı oluştur
        outFile, err := os.Create(outPath)
        if err != nil {
            panic(err)
        }
        defer outFile.Close()

        // İçeriği kopyala
        _, err = io.Copy(outFile, rc)
        if err != nil {
            panic(err)
        }
    }

    fmt.Println("Tüm dosyalar çıkarıldı.")
}

//🔹 Bu örnek, example.zip dosyasındaki içeriği output/ klasörüne çıkartır.
//🔹 Zip içinde klasör varsa onları da doğru şekilde oluşturur.


// Ekstra: Zip Dosyasına Birden Fazla Dosya Ekleme

// Dosyaları döngüyle tek tek zip’e ekleyebilirsin:
//✅ Kod:
files := []string{"file1.txt", "file2.txt"}
for _, fileName := range files {
    // Yukarıdaki zip yazma kodunu bu döngü içinde uygula
}


// Kullanışlı Kütüphaneler (3. Parti)

// Go’nun archive/zip paketi çoğu iş için yeterlidir. Ancak daha ileri seviye işlemler için şu paketlere göz atabilirsin:

// github.com/mholt/archiver
// : Zip + diğer formatlar için kolay arayüz.

// github.com/alexmullins/zip
// : Şifreli zip destekler.