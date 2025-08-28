// Go (Golang) dilinde .tar dosyalarıyla çalışmak, arşivleme ve sıkıştırma işlemleri için yaygın olarak yapılır. 
// Go’nun archive/tar standard kütüphanesi bu işler için uygundur. 
// Aşağıda hem .tar arşivi oluşturma, hem de .tar arşivini çıkarma örneklerini adım adım açıklamalı olarak veriyorum.

// 📦 1. .tar Arşivi Oluşturmak (Dosyaları .tar haline getirmek)
// ➕ Amaç:

// Belirli dosyaları alıp archive.tar isimli bir arşiv dosyası oluşturmak.

//✅ Kod:
package main

import (
    "archive/tar"
    "fmt"
    "io"
    "os"
)

func main() {
    // 1. Arşiv dosyasını oluştur
    tarFile, err := os.Create("archive.tar")
    if err != nil {
        panic(err)
    }
    defer tarFile.Close()

    // 2. Tar yazıcı oluştur
    tarWriter := tar.NewWriter(tarFile)
    defer tarWriter.Close()

    // 3. Arşive eklemek istediğimiz dosyaları tanımlayalım
    files := []string{"file1.txt", "file2.txt"}

    for _, fileName := range files {
        file, err := os.Open(fileName)
        if err != nil {
            panic(err)
        }
        defer file.Close()

        // Dosya bilgilerini al
        stat, err := file.Stat()
        if err != nil {
            panic(err)
        }

        // Header (başlık) yaz
        header := &tar.Header{
            Name: fileName,
            Size: stat.Size(),
            Mode: int64(stat.Mode()),
        }
        err = tarWriter.WriteHeader(header)
        if err != nil {
            panic(err)
        }

        // Dosya içeriğini yaz
        _, err = io.Copy(tarWriter, file)
        if err != nil {
            panic(err)
        }
        fmt.Println("Eklendi:", fileName)
    }
}

//📁 Gerekli: file1.txt ve file2.txt dosyalarının önceden mevcut olması gerekir.



// 📂 2. .tar Arşivini Açmak (Dosyaları çıkarmak)
// ➕ Amaç:

// Bir .tar dosyasındaki dosyaları açmak ve klasöre çıkarmak.

// ✅ Kod:

package main

import (
    "archive/tar"
    "fmt"
    "io"
    "os"
)

func main() {
    // 1. Arşiv dosyasını aç
    tarFile, err := os.Open("archive.tar")
    if err != nil {
        panic(err)
    }
    defer tarFile.Close()

    // 2. Tar okuyucu oluştur
    tarReader := tar.NewReader(tarFile)

    // 3. İçeriği sırayla oku
    for {
        header, err := tarReader.Next()
        if err == io.EOF {
            break // bitti
        }
        if err != nil {
            panic(err)
        }

        fmt.Println("Çıkarılıyor:", header.Name)

        outFile, err := os.Create(header.Name)
        if err != nil {
            panic(err)
        }
        defer outFile.Close()

        _, err = io.Copy(outFile, tarReader)
        if err != nil {
            panic(err)
        }
    }
}

// 🔧 İpucu:

// .tar.gz gibi sıkıştırılmış dosyalar için ayrıca compress/gzip kütüphanesini de kullanman gerekir.

// Bu örnekler .tar içindir, .tar.gz desteği istersen ayrıca örnek sağlayabilirim.