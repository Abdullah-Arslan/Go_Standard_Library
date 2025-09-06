/*
Go’nun **`image`** paketi, standart kütüphanede bulunan ve görseller (resimler) ile çalışmaya yarayan temel pakettir. Bu paket sayesinde bitmap resimler (PNG, JPEG gibi) oluşturabilir, düzenleyebilir, piksellere erişebilir ve çeşitli renk modelleri ile çalışabilirsin.

Aşağıda sana **`image` paketini en baştan sona kadar, alt tipleriyle ve örneklerle** açıklayacağım.

---

# 📌 `image` Paketi Nedir?

Go’da `image` paketi, resimleri temsil eden veri yapıları ve arabirimler içerir.
Resimlerin **pikselleri** ve **renk modelleri** ile çalışmayı sağlar.

Bir resim (`image.Image`) şu özellikleri taşır:

* **Renk modeli (ColorModel)**: RGB, RGBA, Gray vb.
* **Sınır (Bounds)**: Görselin boyutları (dikdörtgen).
* **At(x, y)**: Belirli bir koordinattaki pikselin rengi.

---

# 📌 Önemli Tipler ve Fonksiyonlar

### 1. `image.Image` Arayüzü

Go’daki tüm resim tipleri bu arayüzü uygular.
*/

``go
type Image interface {
    ColorModel() color.Model   // Renk modeli
    Bounds() Rectangle         // Görselin boyutları
    At(x, y int) color.Color   // (x,y) piksel rengini döner
}
``

/*
### 2. `image.Rectangle`

Resmin kapsadığı dikdörtgen alanı temsil eder.
*/

``go
rect := image.Rect(0, 0, 100, 50) // 100x50 boyutunda dikdörtgen
``
/*
### 3. `image.Point`

Bir noktayı (x,y) koordinatlarıyla ifade eder.

### 4. Renk Modelleri (`color` paketiyle birlikte kullanılır)

* `color.RGBA`
* `color.Gray`
* `color.NRGBA`
* `color.Alpha`
* `color.CMYK` (ekstra)

### 5. Hazır Resim Tipleri

* `image.RGBA`
* `image.NRGBA`
* `image.Gray`
* `image.Gray16`
* `image.Alpha`
* `image.Alpha16`
* `image.Paletted` (palet tabanlı)

### 6. Resim Oluşturma
*/
``go
img := image.NewRGBA(image.Rect(0, 0, 200, 100)) // 200x100 RGBA görsel
``
/*
---

# 📌 Örnekler

## 🎨 1. Basit Görsel Oluşturma ve Kaydetme
*/
``go
package main

import (
    "image"
    "image/color"
    "image/png"
    "os"
)

func main() {
    // 200x100 boyutunda yeni RGBA görsel
    img := image.NewRGBA(image.Rect(0, 0, 200, 100))

    // Arka planı beyaz yap
    white := color.RGBA{255, 255, 255, 255}
    for x := 0; x < 200; x++ {
        for y := 0; y < 100; y++ {
            img.Set(x, y, white)
        }
    }

    // Ortaya kırmızı dikdörtgen çiz
    red := color.RGBA{255, 0, 0, 255}
    for x := 50; x < 150; x++ {
        for y := 25; y < 75; y++ {
            img.Set(x, y, red)
        }
    }

    // Dosyaya kaydet
    file, _ := os.Create("output.png")
    defer file.Close()
    png.Encode(file, img)
}
``
/*
👉 Bu kod çalıştırıldığında **beyaz arkaplan üzerinde kırmızı dikdörtgen** içeren bir `output.png` dosyası oluşturulur.

---

## 🎨 2. Bir Görseli Okuma ve Piksel Renklerini Alma
*/
``go
package main

import (
    "fmt"
    "image"
    "image/png"
    "os"
)

func main() {
    // Dosya aç
    file, _ := os.Open("input.png")
    defer file.Close()

    // Decode et
    img, _, _ := image.Decode(file)

    // Görsel boyutları
    bounds := img.Bounds()
    fmt.Println("Boyut:", bounds.Dx(), "x", bounds.Dy())

    // Orta pikselin rengini al
    x := bounds.Dx() / 2
    y := bounds.Dy() / 2
    color := img.At(x, y)
    fmt.Println("Orta piksel:", color)
}
``
/*
👉 Bu örnek **bir görseli açar ve ortadaki pikselin rengini terminale yazdırır.**

---

## 🎨 3. Siyah-Beyaz Görsel (Gray) Oluşturma
*/
``go
package main

import (
    "image"
    "image/color"
    "image/png"
    "os"
)

func main() {
    img := image.NewGray(image.Rect(0, 0, 100, 100))

    // Çapraz çizgi çiz
    for i := 0; i < 100; i++ {
        img.Set(i, i, color.Gray{Y: 255}) // beyaz
    }

    file, _ := os.Create("gray.png")
    defer file.Close()
    png.Encode(file, img)
}
``
/*
👉 Bu kod siyah arka plan üzerinde beyaz çapraz çizgi çizer.

---

# 📌 Özet

* `image` paketi, Go’da resimlerin **oluşturulması, işlenmesi ve kaydedilmesi** için temel araçları sağlar.
* **Arayüzler (`image.Image`)** → resimleri temsil eder.
* **Tipler (`image.RGBA`, `image.Gray`)** → farklı renk modelleri ile çalışır.
* **Fonksiyonlar (`image.NewRGBA`, `image.Rect`)** → yeni resimler oluşturur.
* Diğer paketlerle birlikte kullanılır:

  * `image/color` → renkleri tanımlar
  * `image/png`, `image/jpeg`, `image/gif` → resimleri kaydetmek/okumak için

---

İstersen sana **`image` paketini `image/draw` paketiyle birlikte kullanarak görselleri üzerine çizim yapmayı** da anlatabilirim. İstiyor musun?
EVET
*/
/*
Süper 👍 Şimdi `image` paketini **`image/draw`** paketiyle birlikte nasıl kullanabileceğini detaylıca göstereyim.

`image` paketi tek başına piksellere erişip set etmene izin verir ama karmaşık çizimlerde (resmi kopyalama, üstüne başka bir görsel ekleme, doldurma, çizgi, maskeleme) bu zahmetli olur. İşte burada **`image/draw`** devreye giriyor.

---

# 📌 `image/draw` Paketi Nedir?

`image/draw` paketi, resimler üzerinde **kopyalama, boyama, üst üste bindirme (compositing)** gibi işlemleri kolaylaştırır.

Temel işlevi:
*/
``go
func Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point, op Op)
``
/*
* `dst` → Hedef resim
* `r` → Hedefte çizilecek alan
* `src` → Kaynak resim
* `sp` → Kaynak başlangıç noktası
* `op` → İşlem türü (`draw.Src`, `draw.Over`)

---

# 📌 `draw.Op` Türleri

* `draw.Src` → Kaynak resmi **direkt** hedefin üzerine yazar. (arkaplanı siler)
* `draw.Over` → Kaynak resmi hedefin üzerine **alfa kanalıyla karıştırarak** çizer. (şeffaflık destekler)

---

# 🎨 Örnekler

## 1. Basit Doldurma (`draw.Draw`)

Aşağıdaki kod 200x200 resim oluşturur, arka planı yeşil yapar:
*/
``go
package main

import (
    "image"
    "image/color"
    "image/draw"
    "image/png"
    "os"
)

func main() {
    // Hedef resim (RGBA)
    dst := image.NewRGBA(image.Rect(0, 0, 200, 200))

    // Yeşil renk oluştur
    green := &image.Uniform{color.RGBA{0, 255, 0, 255}}

    // Tüm alanı yeşil ile doldur
    draw.Draw(dst, dst.Bounds(), green, image.Point{}, draw.Src)

    // Dosyaya kaydet
    file, _ := os.Create("green.png")
    defer file.Close()
    png.Encode(file, dst)
}
``
/*
👉 Çalıştırınca **tamamen yeşil arkaplanlı bir resim** oluşur.

---

## 2. İki Görseli Üst Üste Koyma
*/
``go
package main

import (
    "image"
    "image/color"
    "image/draw"
    "image/png"
    "os"
)

func main() {
    // Hedef (beyaz arkaplan)
    dst := image.NewRGBA(image.Rect(0, 0, 300, 300))
    white := &image.Uniform{color.White}
    draw.Draw(dst, dst.Bounds(), white, image.Point{}, draw.Src)

    // Kırmızı kutu
    redBox := image.NewRGBA(image.Rect(0, 0, 100, 100))
    red := &image.Uniform{color.RGBA{255, 0, 0, 255}}
    draw.Draw(redBox, redBox.Bounds(), red, image.Point{}, draw.Src)

    // Mavi kutu
    blueBox := image.NewRGBA(image.Rect(0, 0, 100, 100))
    blue := &image.Uniform{color.RGBA{0, 0, 255, 128}} // yarı saydam
    draw.Draw(blueBox, blueBox.Bounds(), blue, image.Point{}, draw.Src)

    // Kırmızı kutuyu sola, mavi kutuyu sağa çiz
    draw.Draw(dst, image.Rect(50, 50, 150, 150), redBox, image.Point{}, draw.Over)
    draw.Draw(dst, image.Rect(100, 100, 200, 200), blueBox, image.Point{}, draw.Over)

    // Kaydet
    file, _ := os.Create("overlap.png")
    defer file.Close()
    png.Encode(file, dst)
}
``
/*
👉 Sonuç: **Beyaz arkaplan üzerinde kırmızı kare ve üstüne kısmen şeffaf mavi kare** olur.

---

## 3. Görseli Maske ile Çizme

Maske kullanarak bir resmin sadece belirli kısmını çizebilirsin.
*/
``go
package main

import (
    "image"
    "image/color"
    "image/draw"
    "image/png"
    "os"
)

func main() {
    dst := image.NewRGBA(image.Rect(0, 0, 200, 200))
    draw.Draw(dst, dst.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

    // Kaynak kırmızı kutu
    src := image.NewUniform(color.RGBA{255, 0, 0, 255})

    // Maske: merkezde siyah çember olacak şekilde Gray
    mask := image.NewGray(image.Rect(0, 0, 200, 200))
    for y := 0; y < 200; y++ {
        for x := 0; x < 200; x++ {
            dx, dy := x-100, y-100
            if dx*dx+dy*dy < 60*60 { // çemberin içindeyse
                mask.SetGray(x, y, color.Gray{Y: 255})
            }
        }
    }

    // Çizim (Src → Over yerine Op kullanılmaz çünkü maske var)
    draw.DrawMask(dst, dst.Bounds(), src, image.Point{}, mask, image.Point{}, draw.Over)

    file, _ := os.Create("masked.png")
    defer file.Close()
    png.Encode(file, dst)
}
``
/*
👉 Sonuç: **Beyaz arka plan üzerinde kırmızı daire** çıkar.

---

# 📌 Özet

* `image` → Piksel tabanlı resim yapıları sağlar.
* `image/draw` → Resimlere **boyama, üst üste bindirme, maskeleme** gibi yüksek seviye işlemler ekler.
* `draw.Draw` → Kaynaktan hedefe resim kopyalar.
* `draw.DrawMask` → Maske ile çizim yapar.
* `draw.Src` → Direkt çizim (arka planı siler).
* `draw.Over` → Üst üste bindirme (şeffaflıkla).

---

İstersen sana bir sonraki adımda `image/gif`, `image/jpeg`, `image/png` paketlerini de anlatarak **nasıl farklı formatlarda görsel kaydedip okuyabileceğini** göstereyim. Anlatayım mı?
EVET
*/
/*
Harika 👍 Şimdi `image` ve `image/draw`’dan sonra, Go’nun standart kütüphanesindeki **görsel formatı paketlerini** (`image/png`, `image/jpeg`, `image/gif`) tek tek açıklayalım.

Bu paketler aslında **`image` ile oluşturduğun veya açtığın görselleri kaydetmek/okumak için codec** sağlar.

---

# 📌 `image/png` Paketi

PNG (lossless / kayıpsız) formatıyla çalışır.

### Önemli Fonksiyonlar

* `png.Encode(w io.Writer, m image.Image) error` → Görseli PNG formatında yazar.
* `png.Decode(r io.Reader) (image.Image, error)` → PNG görselini okur.
* `png.DecodeConfig(r io.Reader)` → Görselin boyut ve renk modeli bilgisini hızlıca döner (tam decode etmeden).

### Örnek: PNG oluşturma
*/
``go
package main

import (
    "image"
    "image/color"
    "image/png"
    "os"
)

func main() {
    img := image.NewRGBA(image.Rect(0, 0, 100, 100))

    // Çapraz beyaz çizgi
    for i := 0; i < 100; i++ {
        img.Set(i, i, color.White)
    }

    file, _ := os.Create("example.png")
    defer file.Close()
    png.Encode(file, img)
}
``
/*
---

# 📌 `image/jpeg` Paketi

JPEG (lossy / kayıplı, fotoğraflarda yaygın) formatıyla çalışır.

### Önemli Fonksiyonlar

* `jpeg.Encode(w io.Writer, m image.Image, o *jpeg.Options) error`

  * `jpeg.Options{Quality: 1..100}` kaliteyi belirler.
* `jpeg.Decode(r io.Reader) (image.Image, error)`
* `jpeg.DecodeConfig(r io.Reader)` → hızlı okuma

### Örnek: JPEG kaydetme
*/
``go
package main

import (
    "image"
    "image/color"
    "image/jpeg"
    "os"
)

func main() {
    img := image.NewRGBA(image.Rect(0, 0, 200, 200))

    // Mavi arkaplan
    for x := 0; x < 200; x++ {
        for y := 0; y < 200; y++ {
            img.Set(x, y, color.RGBA{0, 0, 255, 255})
        }
    }

    file, _ := os.Create("example.jpg")
    defer file.Close()

    // Kalite 80
    jpeg.Encode(file, img, &jpeg.Options{Quality: 80})
}
``
/*
---

# 📌 `image/gif` Paketi

GIF (256 renk + animasyon desteği) formatıyla çalışır.

### Önemli Fonksiyonlar

* `gif.Encode(w io.Writer, m image.Image, o *gif.Options)` → Tek kare GIF
* `gif.EncodeAll(w io.Writer, g *gif.GIF)` → Çok kareli animasyon
* `gif.Decode(r io.Reader)` → Tek kare
* `gif.DecodeAll(r io.Reader)` → Tüm kareler

### Örnek: Tek kare GIF
*/
``go
package main

import (
    "image"
    "image/color"
    "image/gif"
    "os"
)

func main() {
    img := image.NewPaletted(image.Rect(0, 0, 100, 100), []color.Color{
        color.White,
        color.Black,
    })

    // Basit siyah-beyaz kareler
    for i := 0; i < 100; i++ {
        img.SetColorIndex(i, i, 1)
    }

    file, _ := os.Create("example.gif")
    defer file.Close()

    gif.Encode(file, img, nil)
}
``

//### Örnek: Animasyonlu GIF

``go
package main

import (
    "image"
    "image/color"
    "image/gif"
    "os"
)

func main() {
    // Palet
    palette := []color.Color{color.White, color.Black, color.Red}

    // Kareler
    var frames []*image.Paletted
    var delays []int

    for i := 0; i < 10; i++ {
        img := image.NewPaletted(image.Rect(0, 0, 100, 100), palette)

        // Kırmızı hareket eden nokta
        img.SetColorIndex(10+i*5, 50, 2)

        frames = append(frames, img)
        delays = append(delays, 10) // her kare 100ms
    }

    // Animasyon GIF
    g := &gif.GIF{
        Image: frames,
        Delay: delays,
    }

    file, _ := os.Create("anim.gif")
    defer file.Close()
    gif.EncodeAll(file, g)
}
``
/*
👉 Çalıştırınca **kırmızı noktanın hareket ettiği animasyonlu bir GIF** çıkar.

---

# 📌 Özet

* **`image/png`** → Kayıpsız, grafikler için ideal.
* **`image/jpeg`** → Kayıplı, fotoğraflar için ideal.
* **`image/gif`** → 256 renk, animasyon desteği var.

Tüm bu paketler **`image.Image` arayüzüyle uyumlu** çalışır. Yani:

* Görseli `image` ile oluşturursun
* `draw` ile işlersin
* Sonra `png/jpeg/gif` ile kaydedersin

---

İstersen sana son adımda **gerçek bir uygulama örneği** (örneğin: resmi aç → küçük resim (thumbnail) oluştur → PNG ve JPEG olarak kaydet) gösterebilirim. Gösteriyim mi?
EVET
*/
/*
Süper 👍 Şimdiye kadar `image`, `image/draw`, `image/png/jpeg/gif` paketlerini tek tek gördük.
Şimdi bunların hepsini birleştirip **gerçek bir uygulama örneği** yapalım:

👉 Senaryo:

1. Bir görsel açalım (`input.jpg`).
2. Onu **küçük resim (thumbnail)** boyutuna küçültelim (ör: 100x100).
3. Hem PNG hem JPEG olarak kaydedelim.

---

# 📌 Örnek: Thumbnail Uygulaması
*/
``go
package main

import (
    "fmt"
    "image"
    "image/draw"
    "image/jpeg"
    "image/png"
    "os"
)

func main() {
    // 1. Girdi görselini aç
    inputFile, err := os.Open("input.jpg")
    if err != nil {
        fmt.Println("Hata:", err)
        return
    }
    defer inputFile.Close()

    // Decode et (format otomatik algılanır)
    src, _, err := image.Decode(inputFile)
    if err != nil {
        fmt.Println("Decode hatası:", err)
        return
    }

    // 2. Küçük resim (100x100) oluştur
    thumbRect := image.Rect(0, 0, 100, 100)
    thumb := image.NewRGBA(thumbRect)

    // Oranlı küçültme (en basit yöntem: image/draw ile kopyalama)
    draw.NearestNeighbor.Scale(thumb, thumbRect, src, src.Bounds(), draw.Over, nil)

    // 3a. PNG olarak kaydet
    pngFile, _ := os.Create("thumbnail.png")
    defer pngFile.Close()
    png.Encode(pngFile, thumb)

    // 3b. JPEG olarak kaydet (kalite 80)
    jpegFile, _ := os.Create("thumbnail.jpg")
    defer jpegFile.Close()
    jpeg.Encode(jpegFile, thumb, &jpeg.Options{Quality: 80})

    fmt.Println("Thumbnail başarıyla oluşturuldu: thumbnail.png & thumbnail.jpg")
}
``
/*
---

# 📌 Açıklama

1. **`image.Decode`** → Hangi formatta olursa olsun (`jpeg/png/gif`) resmi açar.
2. **`draw.NearestNeighbor.Scale`** → Kaynağı küçük boyuta ölçekler (piksel bazlı hızlı yöntem).

   * Daha kaliteli yöntemler için `draw.ApproxBiLinear` veya `draw.CatmullRom` kullanılabilir.
3. **`png.Encode`** ve **`jpeg.Encode`** → Aynı görseli farklı formatlarda kaydeder.

---

# 📌 Özet

Bu uygulamada:

* Girdi olarak **JPEG resmi okuduk**,
* `image/draw` ile **küçük resim yaptık**,
* Aynı dosyayı hem **PNG** hem **JPEG** olarak kaydettik.

Böylece `image`, `draw`, `png`, `jpeg` paketlerinin hepsini birlikte kullanmış olduk ✅

---

İstersen bir sonraki adımda sana `golang.org/x/image` paketlerinden de bahsedebilirim (ör. **bmp, tiff, webp, font, colornames** gibi ekstra destekler).
Onları da anlatayım mı?
*/