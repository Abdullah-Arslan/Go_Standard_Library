/*
Go’nun **`image/jpeg`** paketini baştan sona detaylı şekilde açıklayalım ve örneklerle gösterelim.

---

# 📌 `image/jpeg` Paketi Nedir?

* `image/jpeg` paketi, Go’nun standart kütüphanesinde **JPEG formatındaki görselleri okuma ve yazma** işlemleri için kullanılır.
* JPEG, kayıplı sıkıştırma kullanan yaygın bir resim formatıdır.
* Paket, tek kareli görseller için uygundur (animasyonlu JPEG yoktur).
* `image` ve `image/color` paketleri ile birlikte çalışır.

---

# 📌 Temel Fonksiyonlar

1. **`jpeg.Decode(r io.Reader) (image.Image, error)`**

   * JPEG dosyasını okur ve `image.Image` olarak döner.

2. **`jpeg.Encode(w io.Writer, m image.Image, o *jpeg.Options) error`**

   * `image.Image` nesnesini JPEG formatında yazar.
   * `Options` ile kalite (0-100) ayarlanabilir.

---

# 📌 Önemli Tipler

* **`jpeg.Options` struct**
*/
``go
type Options struct {
    Quality int // 1-100 arası kalite değeri
}
``
/*
* `image.Image` → Görselin tüm piksel verilerini tutar.
* JPEG kayıplı olduğu için yüksek kalite (örn. 90-100) daha iyi sonuç verir, dosya boyutu artar.

---

# 📌 Örnekler

## 1️⃣ JPEG Dosyası Okuma
*/
``go
package main

import (
    "fmt"
    "image/jpeg"
    "os"
)

func main() {
    f, err := os.Open("example.jpg")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    img, err := jpeg.Decode(f)
    if err != nil {
        panic(err)
    }

    fmt.Println("Görsel boyutu:", img.Bounds())
}
``

/*
---

## 2️⃣ JPEG Dosyası Oluşturma
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
    // 200x200 beyaz kare
    img := image.NewRGBA(image.Rect(0, 0, 200, 200))

    for y := 0; y < 200; y++ {
        for x := 0; x < 200; x++ {
            img.Set(x, y, color.RGBA{255, 255, 255, 255})
        }
    }

    f, err := os.Create("output.jpg")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // Kaliteyi ayarla (örn. 90)
    opt := &jpeg.Options{Quality: 90}
    jpeg.Encode(f, img, opt)
}
``
/*
---

## 3️⃣ JPEG + Filtre Örneği (Grayscale)
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
    f, _ := os.Open("example.jpg")
    defer f.Close()
    img, _ := jpeg.Decode(f)

    bounds := img.Bounds()
    grayImg := image.NewGray(bounds)

    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            r, g, b, _ := img.At(x, y).RGBA()
            gray := uint8((299*int(r/257) + 587*int(g/257) + 114*int(b/257)) / 1000)
            grayImg.Set(x, y, color.Gray{Y: gray})
        }
    }

    out, _ := os.Create("grayscale.jpg")
    defer out.Close()
    jpeg.Encode(out, grayImg, &jpeg.Options{Quality: 90})
}
``
/*
---

# 📌 Özet

1. `image/jpeg` **JPEG görselleri okuma ve yazma** için kullanılır.
2. Tek kareli görseller için uygundur (animasyon desteklenmez).
3. Kalite `jpeg.Options{Quality: ...}` ile ayarlanabilir.
4. Diğer `image` paketleri ile birlikte kullanarak **filtre, overlay, boyutlandırma** işlemleri yapılabilir.

---

İstersen bir sonraki adımda **`jpeg` + `draw` + `palette` kullanarak JPEG’den GIF veya PNG’ye dönüşüm ve filtre uygulama** örneği yapabiliriz.

Bunu da yapalım mı?
EVET
*/
/*
Harika! 🚀 O zaman şimdi **JPEG → GIF/PNG dönüşümü + filtre ve overlay uygulama** örneğini yapalım.

Senaryo:

* Kullanıcı bir JPEG görsel yükleyecek.
* Üzerine **filtre** (Negative, Grayscale, Sepia) uygulanacak.
* Opsiyonel olarak **overlay** (küçük renkli kare veya logo) eklenecek.
* Sonuç **PNG veya GIF** olarak kaydedilecek.

---

# 📌 Örnek Program: JPEG → PNG/GIF + Filtre + Overlay
*/
``go
package main

import (
    "fmt"
    "image"
    "image/color"
    "image/draw"
    "image/gif"
    "image/jpeg"
    "image/png"
    "os"
)

// ------------- Filtreler ----------------
func clamp(v int) uint8 {
    if v < 0 {
        return 0
    }
    if v > 255 {
        return 255
    }
    return uint8(v)
}

func negative(img image.Image) *image.RGBA {
    bounds := img.Bounds()
    newImg := image.NewRGBA(bounds)
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            r, g, b, a := img.At(x, y).RGBA()
            newImg.Set(x, y, color.RGBA{
                R: 255 - uint8(r/257),
                G: 255 - uint8(g/257),
                B: 255 - uint8(b/257),
                A: uint8(a / 257),
            })
        }
    }
    return newImg
}

func grayscale(img image.Image) *image.RGBA {
    bounds := img.Bounds()
    newImg := image.NewRGBA(bounds)
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            r, g, b, a := img.At(x, y).RGBA()
            gray := uint8((299*int(r/257) + 587*int(g/257) + 114*int(b/257)) / 1000)
            newImg.Set(x, y, color.RGBA{gray, gray, gray, uint8(a / 257)})
        }
    }
    return newImg
}

func sepia(img image.Image) *image.RGBA {
    bounds := img.Bounds()
    newImg := image.NewRGBA(bounds)
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            r, g, b, a := img.At(x, y).RGBA()
            tr := clamp(int(0.393*float64(r/257) + 0.769*float64(g/257) + 0.189*float64(b/257)))
            tg := clamp(int(0.349*float64(r/257) + 0.686*float64(g/257) + 0.168*float64(b/257)))
            tb := clamp(int(0.272*float64(r/257) + 0.534*float64(g/257) + 0.131*float64(b/257)))
            newImg.Set(x, y, color.RGBA{tr, tg, tb, uint8(a / 257)})
        }
    }
    return newImg
}

func applyFilter(img image.Image, choice string) *image.RGBA {
    switch choice {
    case "1":
        return negative(img)
    case "2":
        return grayscale(img)
    case "3":
        return sepia(img)
    default:
        return nil
    }
}

// ------------- Ana Program ----------------
func main() {
    var inputPath, format, choice string
    fmt.Print("JPEG dosya yolu: ")
    fmt.Scanln(&inputPath)

    f, err := os.Open(inputPath)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    img, err := jpeg.Decode(f)
    if err != nil {
        panic(err)
    }

    fmt.Println("Filtre Seçenekleri: 1-Negative 2-Grayscale 3-Sepia")
    fmt.Print("Seçiminiz: ")
    fmt.Scanln(&choice)

    filtered := applyFilter(img, choice)

    // Overlay ekleme (opsiyonel)
    overlay := image.NewRGBA(image.Rect(0, 0, 50, 50))
    draw.Draw(overlay, overlay.Bounds(), &image.Uniform{color.RGBA{255, 0, 0, 128}}, image.Point{}, draw.Over)
    draw.Draw(filtered, overlay.Bounds().Add(image.Pt(filtered.Bounds().Dx()-50, filtered.Bounds().Dy()-50)), overlay, image.Point{}, draw.Over)

    fmt.Print("Kaydedilecek format (png/gif): ")
    fmt.Scanln(&format)

    if format == "png" {
        out, _ := os.Create("output.png")
        defer out.Close()
        png.Encode(out, filtered)
        fmt.Println("PNG dosyası kaydedildi: output.png")
    } else if format == "gif" {
        // Palet tabanlı GIF için
        paletted := image.NewPaletted(filtered.Bounds(), color.Palette{color.White, color.Black, color.RGBA{255, 0, 0, 255}})
        draw.FloydSteinberg.Draw(paletted, filtered.Bounds(), filtered, image.Point{})
        out, _ := os.Create("output.gif")
        defer out.Close()
        gif.Encode(out, paletted, nil)
        fmt.Println("GIF dosyası kaydedildi: output.gif")
    } else {
        fmt.Println("Desteklenmeyen format!")
    }
}
``
/*
---

# 📌 Açıklama

1. **JPEG okuma** → `jpeg.Decode`.
2. **Filtreler** → Negative, Grayscale, Sepia uygulanabilir.
3. **Overlay** → Sağ alt köşeye yarı saydam kare bindirilir.
4. **Dönüştürme ve kaydetme** → `png.Encode` veya `gif.Encode` ile format seçimi yapılabilir.
5. GIF için **palet tabanlı** dönüşüm yapılır (`draw.FloydSteinberg`).

---

Bu program sayesinde artık **JPEG görselleri filtreleyip overlay ekleyerek PNG veya GIF olarak kaydedebiliyorsun**.

İstersen bir sonraki adımda bunu **animasyonlu JPEG serilerini GIF’e dönüştüren bir mini program** hâline getirip birden fazla JPEG’ten animasyonlu GIF oluşturabiliriz.

Bunu da yapalım mı?
EVET
*/
/*
Harika! 🎉 O zaman şimdi **birden fazla JPEG’den animasyonlu GIF oluşturma** örneği yapalım.

Senaryo:

* Kullanıcı bir klasörde veya liste halinde JPEG dosyaları verecek.
* Her JPEG’e opsiyonel **filtre** uygulanacak.
* Sonuçta **animasyonlu GIF** üretilecek.

---

# 📌 Örnek Program: JPEG Serisinden Animasyonlu GIF
*/
``go
package main

import (
    "fmt"
    "image"
    "image/color"
    "image/draw"
    "image/gif"
    "image/jpeg"
    "os"
)

// ---------------- Filtreler ----------------
func clamp(v int) uint8 {
    if v < 0 {
        return 0
    }
    if v > 255 {
        return 255
    }
    return uint8(v)
}

func grayscale(img image.Image) *image.RGBA {
    bounds := img.Bounds()
    newImg := image.NewRGBA(bounds)
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            r, g, b, a := img.At(x, y).RGBA()
            gray := uint8((299*int(r/257) + 587*int(g/257) + 114*int(b/257)) / 1000)
            newImg.Set(x, y, color.RGBA{gray, gray, gray, uint8(a / 257)})
        }
    }
    return newImg
}

func applyFilter(img image.Image, choice string) *image.RGBA {
    switch choice {
    case "1":
        return grayscale(img) // örnek olarak sadece grayscale
    default:
        bounds := img.Bounds()
        rgba := image.NewRGBA(bounds)
        draw.Draw(rgba, bounds, img, image.Point{}, draw.Src)
        return rgba
    }
}

// ---------------- Ana Program ----------------
func main() {
    var n int
    fmt.Print("Kaç JPEG kullanacaksınız? ")
    fmt.Scanln(&n)

    jpegFiles := make([]string, n)
    for i := 0; i < n; i++ {
        fmt.Printf("JPEG dosya %d: ", i+1)
        fmt.Scanln(&jpegFiles[i])
    }

    fmt.Println("Filtre seçimi (1- Grayscale, 0- Yok): ")
    var choice string
    fmt.Scanln(&choice)

    anim := &gif.GIF{}

    for _, file := range jpegFiles {
        f, err := os.Open(file)
        if err != nil {
            panic(err)
        }
        img, err := jpeg.Decode(f)
        f.Close()
        if err != nil {
            panic(err)
        }

        rgba := applyFilter(img, choice)

        // Overlay ekleme (küçük kırmızı kare)
        overlay := image.NewRGBA(image.Rect(0, 0, 30, 30))
        draw.Draw(overlay, overlay.Bounds(), &image.Uniform{color.RGBA{255, 0, 0, 128}}, image.Point{}, draw.Over)
        draw.Draw(rgba, overlay.Bounds().Add(image.Pt(rgba.Bounds().Dx()-30, rgba.Bounds().Dy()-30)), overlay, image.Point{}, draw.Over)

        // RGBA → Paletted
        pal := image.NewPaletted(rgba.Bounds(), color.Palette{color.White, color.Black, color.RGBA{255, 0, 0, 255}})
        draw.FloydSteinberg.Draw(pal, rgba.Bounds(), rgba, image.Point{})

        anim.Image = append(anim.Image, pal)
        anim.Delay = append(anim.Delay, 10) // 100ms
    }

    out, err := os.Create("animated_from_jpeg.gif")
    if err != nil {
        panic(err)
    }
    defer out.Close()

    gif.EncodeAll(out, anim)
    fmt.Println("Animasyonlu GIF oluşturuldu: animated_from_jpeg.gif")
}
``
/*
---

# 📌 Açıklama

1. Kullanıcıdan **JPEG dosya listesi** alınıyor.
2. İsteğe bağlı **grayscale filtre** uygulanıyor (dilediğiniz başka filtreyi de ekleyebilirsiniz).
3. Her kareye **overlay** ekleniyor (küçük yarı saydam kırmızı kare).
4. **RGBA → Paletted** dönüşümü yapılıyor (GIF uyumluluğu).
5. Animasyonun gecikmesi `anim.Delay` ile ayarlanıyor (10 = 100ms).
6. Sonuç `animated_from_jpeg.gif` olarak kaydediliyor.

---

Bu yöntemle artık **birden fazla JPEG’den animasyonlu GIF oluşturabilir, filtre ve overlay ekleyebilirsiniz**.

İstersen bir sonraki adımda bunu **terminal önizleme + renk paleti seçimi + interaktif filtre seçimi** ile tamamen interaktif hâle getirebiliriz.

Bunu da yapalım mı?
EVET
*/

