// Go dilinde (Golang) bytes konusu, özellikle []byte türü (byte slice) etrafında döner. 
// Bytes, temel olarak veriyi ham (raw) formda temsil eden 8-bitlik (1 byte) birimlerdir. 
// Genellikle string ve dosya işlemlerinde, ağ iletişiminde, şifrelemede veya binary veri işleme durumlarında kullanılır.

// 1. Bytes Nedir?

// byte: Go dilinde byte aslında uint8 ile aynı anlama gelir. Yani, 0-255 arasında değer alabilen bir tam sayıdır.

// []byte: Byte dizisi, yani byte slice. Bu, birden fazla byte’ın sıralı koleksiyonudur. Stringlere benzer ama daha düşük seviyede, daha esnek veri yapısıdır.

// 2. Bytes ve String Arasındaki İlişki

// Stringler Go’da UTF-8 kodlamasında karakterler içerir, aslında arka planda []byte dizisidir fakat değiştirilemez (immutable).

// []byte ise değiştirilebilir (mutable) olduğu için stringe göre daha esnektir.

// string → []byte dönüşümü ve tersi sık yapılır.

// 3. Bytes Kullanımına Basit Örnekler
// Örnek 1: String’i byte slice’a dönüştürmek

// Byte Slice Nedir?

// Go dilinde byte slice, aslında []byte tipinde bir dilimdir (slice). byte türü, Go’da 8 bitlik işaretsiz tam sayıyı (uint8) temsil eder. 
// byte genellikle ham veri (binary data) veya metin verisi (string) işlemlerinde kullanılır.
package main

import "fmt"

func main() {
    s := "Merhaba Go"
    b := []byte(s)  // stringi byte slice'a çeviriyoruz

    fmt.Println(b)        // [77 101 114 104 97 98 97 32 71 111]
    fmt.Println(string(b)) // tekrar stringe çeviriyoruz -> Merhaba Go
}


// Örnek 3: bytes paketinden faydalanmak

// Go standart kütüphanesinde bytes adında bir paket vardır. 
// Bu paket byte slice'lar üzerinde pratik fonksiyonlar sunar.

package main

import (
    "bytes"
    "fmt"
)

func main() {
    b := []byte("golang")

    // Contains - içinde var mı?
    fmt.Println(bytes.Contains(b, []byte("lan"))) // true

    // Compare - karşılaştırma
    fmt.Println(bytes.Compare([]byte("a"), []byte("b"))) // -1 (a < b)

    // ToUpper
    fmt.Println(string(bytes.ToUpper(b))) // GOLANG
}


// Örnek 4: Byte buffer kullanımı

// bytes.Buffer yapısı, byte dilimini yazmak, okumak için kullanılır. Performanslı ve kolaydır.
package main

import (
    "bytes"
    "fmt"
)

func main() {
    var buf bytes.Buffer

    buf.WriteString("Merhaba ")
    buf.Write([]byte("Dünya"))

    fmt.Println(buf.String())  // Merhaba Dünya
}

// Go’da Byte Buffer dediğimiz şey, genellikle bytes.Buffer tipiyle ifade edilen, baytları (byte) dinamik olarak tutup işleyebileceğimiz bir yapıdır.

// bytes.Buffer Nedir?

// bytes.Buffer, Go standart kütüphanesindeki bytes paketinde bulunan bir türdür.

// İçinde byte dilimi (slice) tutar ve bu byte’lar üzerinde okuma, yazma, ekleme, temizleme gibi işlemleri kolayca yapmamıza olanak sağlar.

// Dinamik olarak büyüyebilir, yani içine yazdıkça büyür.

// io.Reader ve io.Writer arayüzlerini uygular, bu sayede çok sayıda Go fonksiyonu ile uyumlu çalışır.

// Ne İşe Yarar?

// Metin ya da binary veriyi arabellek olarak tutup üzerinde işlem yapmak için.

// Birden fazla küçük parçayı tek bir byte dizisinde birleştirmek için.

// Ağ, dosya veya başka bir kaynaktan okuma/yazma işlemlerini tamponlamak için.

package main

import (
    "bytes"
    "fmt"
)

func main() {
    var buffer bytes.Buffer

    // Buffer içine yazıyoruz
    buffer.Write([]byte("Merhaba "))
    buffer.WriteString("Dünya!")

    // Buffer içeriğini string olarak alıyoruz
    fmt.Println(buffer.String()) // Merhaba Dünya!

    // Buffer içeriğini byte slice olarak da alabiliriz
    b := buffer.Bytes()
    fmt.Println(b) // [77 101 114 104 97 98 97 32 68 252 110 121 97 33]
}

// 4. Byte Slice Avantajları

// String'den farklı olarak değiştirilebilir.

// Binary veri (dosya, network paketleri vs.) kolay işlenir.

// Go'nun standart kütüphanesinde ve birçok üçüncü taraf kütüphanede temel yapı olarak kullanılır.

// Özet:

// byte = uint8, yani 0-255 arası tam sayı.

// []byte = byte dizisi, değiştirilebilir, binary veriler için ideal.

// String ve []byte arasında dönüşüm çok kolay.

// bytes paketi, byte slice işlemlerini kolaylaştırır.

// bytes.Buffer performanslı veri okuma/yazma sağlar.