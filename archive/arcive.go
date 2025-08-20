// Go programlama dilinde, standart kütüphane içinde sıkıştırılmış dosyalarla çalışmak için kullanılan bazı önemli paketler şunlardır:

// -archive/tar

// -archive/zip

// -compress/gzip, compress/zlib, vs.

// Bu yanıt, özellikle archive/tar ve archive/zip paketlerinin kullanımına odaklanacak ve örneklerle açıklayacaktır.


// 📦 1. archive/tar — TAR Arşivleriyle Çalışma

// archive/tar, .tar formatındaki arşiv dosyalarını oluşturmak (write) ve açmak (read) için kullanılır. .tar.gz gibi dosyaları işlemek için compress/gzip paketiyle birlikte kullanılır.

package main

import (
    "archive/tar"
    "os"
)

func main() {
    file, err := os.Create("output.tar")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    tw := tar.NewWriter(file)
    defer tw.Close()

    // Arşive eklenecek dosya bilgileri
    content := []byte("Hello, this is file content!")
    hdr := &tar.Header{
        Name: "file1.txt",
        Mode: 0600,
        Size: int64(len(content)),
    }

    // Header ve içeriği yaz
    if err := tw.WriteHeader(hdr); err != nil {
        panic(err)
    }

    if _, err := tw.Write(content); err != nil {
        panic(err)
    }
}



//✅ Örnek: .tar Arşivini Açma

package main

import (
    "archive/tar"
    "fmt"
    "io"
    "os"
)

func main() {
    file, err := os.Open("output.tar")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    tr := tar.NewReader(file)

    for {
        hdr, err := tr.Next()
        if err == io.EOF {
            break
        }
        if err != nil {
            panic(err)
        }

        fmt.Println("Dosya:", hdr.Name)
        content, _ := io.ReadAll(tr)
        fmt.Println("İçerik:", string(content))
    }
}


// 📦 2. archive/zip — ZIP Arşivleriyle Çalışma

// archive/zip, .zip formatındaki dosyaları okumak ve oluşturmak için kullanılır. Geniş platform desteği nedeniyle sık tercih edilir.

// ✅ Örnek: .zip Arşivi Oluşturma

package main

import (
    "archive/zip"
    "bytes"
    "os"
)

func main() {
    file, err := os.Create("output.zip")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    zw := zip.NewWriter(file)
    defer zw.Close()

    // Dosya ekle
    f, err := zw.Create("file1.txt")
    if err != nil {
        panic(err)
    }

    _, err = f.Write([]byte("This is file content inside zip"))
    if err != nil {
        panic(err)
    }
}



//Örnek: .zip Arşivini Açma

package main

import (
    "archive/zip"
    "fmt"
    "io"
    "os"
)

func main() {
    r, err := zip.OpenReader("output.zip")
    if err != nil {
        panic(err)
    }
    defer r.Close()

    for _, f := range r.File {
        fmt.Println("Dosya:", f.Name)
        rc, err := f.Open()
        if err != nil {
            panic(err)
        }
        content, _ := io.ReadAll(rc)
        fmt.Println("İçerik:", string(content))
        rc.Close()
    }
}


//🔄 tar + gzip ile .tar.gz Oluşturma

package main

import (
    "archive/tar"
    "compress/gzip"
    "os"
)

func main() {
    file, _ := os.Create("output.tar.gz")
    defer file.Close()

    gw := gzip.NewWriter(file)
    defer gw.Close()

    tw := tar.NewWriter(gw)
    defer tw.Close()

    content := []byte("This is compressed file")
    hdr := &tar.Header{
        Name: "file.txt",
        Mode: 0600,
        Size: int64(len(content)),
    }

    tw.WriteHeader(hdr)
    tw.Write(content)
}




