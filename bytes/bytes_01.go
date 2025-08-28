//  Go’da **`bytes` paketi**, **byte dilimleri (`[]byte`) üzerinde çalışmayı kolaylaştıran bir standart kütüphanedir**.

// * Temel amacı: `[]byte` üzerinde **okuma, yazma, arama, değiştirme ve karşılaştırma** işlemlerini kolay ve verimli yapmak.
// * `strings` paketi ile mantık olarak benzer ama `strings` **string** ile çalışırken, `bytes` **byte slice (`[]byte`)** ile çalışır.
// * Bu paket, genellikle I/O işlemleri, metin işleme ve network programlamada kullanılır.

// ---

// # 📦 `bytes` Paketindeki Temel Yapılar ve Fonksiyonlar

// ### 1️⃣ **`bytes.Buffer`**

// * Bir **dinamik byte tamponu**dır.
// * Yazma (`Write`) ve okuma (`Read`) işlemlerini yönetir.
// * Hafıza kopyalamadan verimli string/byte yönetimi sağlar.

// #### Örnek:


package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer

	buf.WriteString("Merhaba ")
	buf.Write([]byte("Dünya!"))

	fmt.Println(buf.String()) // Merhaba Dünya!
}


// #### Önemli Metodlar:

// * `Write(p []byte)` → byte slice yazar
// * `WriteString(s string)` → string yazar
// * `Read(p []byte)` → byte slice okur
// * `Bytes()` → mevcut byte slice’i döner
// * `String()` → string’e çevirir
// * `Reset()` → tamponu temizler

// ---

// ### 2️⃣ **`bytes.Reader`**

// * Bir byte slice’ı **okunabilir bir veri kaynağı** gibi davranmasını sağlar (`io.Reader`, `io.Seeker`).

// #### Örnek:

b := []byte("Merhaba Dünya")
r := bytes.NewReader(b)

p := make([]byte, 7)
n, _ := r.Read(p)
fmt.Println(string(p[:n])) // Merhaba


// #### Özellikler:

// * `Read(p []byte)` → okuma
// * `Seek(offset int64, whence int)` → okuma konumunu değiştirir
// * `Len()` → kalan uzunluk
// * `Size()` → toplam boyut

// ---

// ### 3️⃣ **`bytes.Equal`**

// * İki byte slice’in eşit olup olmadığını kontrol eder


a := []byte("abc")
b := []byte("abc")
fmt.Println(bytes.Equal(a, b)) // true


// ---

// ### 4️⃣ **`bytes.Compare`**

// * İki byte slice’i sözlük sırasına göre karşılaştırır
// * Döner: `-1` (a\<b), `0` (a=b), `1` (a>b)


fmt.Println(bytes.Compare([]byte("a"), []byte("b"))) // -1


// ---

// ### 5️⃣ **`bytes.Contains`, `bytes.Index`, `bytes.HasPrefix`, `bytes.HasSuffix`**

// * `Contains`: Alt byte dizisi var mı
// * `Index`: Alt byte dizisinin konumu
// * `HasPrefix`, `HasSuffix`: Başlangıç/bitiş kontrolü

b := []byte("Merhaba Dünya")
fmt.Println(bytes.Contains(b, []byte("Dünya"))) // true
fmt.Println(bytes.Index(b, []byte("Dünya")))    // 8
fmt.Println(bytes.HasPrefix(b, []byte("Mer")))  // true
fmt.Println(bytes.HasSuffix(b, []byte("ya")))   // true


// ---

// ### 6️⃣ **`bytes.Split`, `bytes.Join`**

// * `Split` → Byte slice’i ayırır
// * `Join` → Slice’leri birleştirir


b := []byte("a,b,c")
parts := bytes.Split(b, []byte(","))
fmt.Println(parts) // [[97] [98] [99]]

joined := bytes.Join(parts, []byte("-"))
fmt.Println(string(joined)) // a-b-c


// ---

// ### 7️⃣ **`bytes.Replace`**

// * Byte slice içindeki alt slice’i değiştirir

b := []byte("Merhaba Dünya")
newB := bytes.Replace(b, []byte("Dünya"), []byte("Go"), -1)
fmt.Println(string(newB)) // Merhaba Go


// ---

// ### 8️⃣ **`bytes.Repeat`**

// * Byte slice’i n kez tekrarlar


b := []byte("abc")
fmt.Println(string(bytes.Repeat(b, 3))) // abcabcabc

// ```

// ---

// ### 9️⃣ **`bytes.Trim`, `bytes.TrimSpace`**

// * Baştaki ve sondaki karakterleri veya boşlukları temizler

b := []byte("  Merhaba  ")
fmt.Println(string(bytes.TrimSpace(b))) // Merhaba
fmt.Println(string(bytes.Trim(b, " M"))) //erhaba


// ---

// ### 10️⃣ **`bytes.ToLower`, `bytes.ToUpper`**

// * Büyük/küçük harf dönüşümü

b := []byte("MerHaBa")
fmt.Println(string(bytes.ToLower(b))) // merhaba
fmt.Println(string(bytes.ToUpper(b))) // MERHABA

// ```

// ---

// ### 11️⃣ **`bytes.Buffer` ile `io.Writer` ve `io.Reader`**

// * `bytes.Buffer`, `Reader`, `Writer` gibi yapıların tüm `io` arayüzlerini desteklemesi sayesinde **I/O ile uyumlu çalışır**.

// ```go
// buf := bytes.NewBufferString("Merhaba")
// var out []byte
// buf.Read(out) // io.Reader olarak kullanılabilir
// ```

// ---

// ## 🚀 Özet

// **`bytes` paketi**, byte slice ile çalışmayı kolaylaştırır:

| Yapı / Fonksiyon          | Açıklama                             |
| ------------------------- | ------------------------------------ |
| `Buffer`                  | Dinamik byte tamponu, okuma/yazma    |
| `Reader`                  | Byte slice’ı okunabilir hale getirir |
| `Equal`                   | Karşılaştırma                        |
| `Compare`                 | Sıralama karşılaştırması             |
| `Contains`                | Alt diziyi kontrol eder              |
| `Index`                   | Alt dizinin konumu                   |
| `HasPrefix` / `HasSuffix` | Başlangıç/bitiş kontrolü             |
| `Split` / `Join`          | Bölme / birleştirme                  |
| `Replace`                 | Değiştirme                           |
| `Repeat`                  | Tekrarlama                           |
| `Trim` / `TrimSpace`      | Temizleme                            |
| `ToLower` / `ToUpper`     | Harf dönüşümü                        |

---

