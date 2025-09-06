/*
Go’nun **`image/png`** paketini baştan sona detaylı şekilde açıklayalım ve örneklerle gösterelim.

---

# 📌 `image/png` Paketi Nedir?

* `image/png` paketi, Go’nun standart kütüphanesinde **PNG formatındaki görselleri okuma ve yazma** işlemleri için kullanılır.
* PNG (Portable Network Graphics), **kayıpsız sıkıştırma** kullanan yaygın bir resim formatıdır.
* Şeffaflık (alpha channel) desteği vardır.
* `image`, `image/color` ve `image/draw` paketleri ile birlikte çalışır.

---

# 📌 Temel Fonksiyonlar

1. **`png.Decode(r io.Reader) (image.Image, error)`**

   * PNG dosyasını okur ve `image.Image` olarak döner.

2. **`png.Encode(w io.Writer, m image.Image) error`**

   * `image.Image` nesnesini PNG formatında yazar.
   * PNG’de kalite ayarı yoktur çünkü kayıpsızdır.

---

# 📌 Örnekler

## 1️⃣ PNG Dosyası Okuma
*/
``go
package main

import (
    "fmt"
    "image/png"
    "os"
)

func main() {
    f, err := os.Open("example.png")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    img, err := png.Decode(f)
    if err != nil {
        panic(err)
    }

    fmt.Println("Görsel boyutu:", img.Bounds())
}
``
/*
---

## 2️⃣ PNG Dosyası Oluşturma
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
    // 200x200 beyaz kare
    img := image.NewRGBA(image.Rect(0, 0, 200, 200))

    for y := 0; y < 200; y++ {
        for x := 0; x < 200; x++ {
            img.Set(x, y, color.RGBA{255, 255, 255, 255})
        }
    }

    f, err := os.Create("output.png")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    png.Encode(f, img)
}
``
/*
---

## 3️⃣ PNG + Şeffaflık Örneği
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
    img := image.NewRGBA(image.Rect(0, 0, 200, 200))

    for y := 0; y < 200; y++ {
        for x := 0; x < 200; x++ {
            // Yarı saydam kırmızı
            img.Set(x, y, color.RGBA{255, 0, 0, 128})
        }
    }

    f, _ := os.Create("transparent.png")
    defer f.Close()
    png.Encode(f, img)
}
``
/*
---

# 📌 Özet

1. `image/png` **PNG görselleri okuma ve yazma** için kullanılır.
2. Kaybı olmayan bir formattır (lossless).
3. Alpha kanalı sayesinde **şeffaflık** destekler.
4. Diğer `image` paketleri ile birlikte kullanarak **filtre, overlay, boyutlandırma** işlemleri yapılabilir.

---

İstersen bir sonraki adımda **JPEG → PNG dönüşümü + filtre + overlay + terminal önizleme** örneğini canlı olarak gösterelim.

Bunu yapalım mı?
EVET
*/
/*
Harika! 🚀 O zaman şimdi **JPEG → PNG dönüşümü + filtre + overlay + terminal önizleme** örneğini hazırlayalım.

Bu örnek, önceki GIF editörü mantığını PNG formatına uygulayacak ve terminalde önizleme yapacak.

---

# 📌 Örnek Program: JPEG → PNG + Filtre + Overlay + Terminal Önizleme
*/
``go
package main

import (
    "fmt"
    "image"
    "image/color"
    "image/draw"
    "image/jpeg"
    "image/png"
    "os"
    "time"
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
        return grayscale(img)
    default:
        bounds := img.Bounds()
        rgba := image.NewRGBA(bounds)
        draw.Draw(rgba, bounds, img, image.Point{}, draw.Src)
        return rgba
    }
}

// ---------------- Terminal Önizleme ----------------
func rgbToTermColor(c color.Color) int {
    r, g, b, _ := c.RGBA()
    r8 := int(r / 257)
    g8 := int(g / 257)
    b8 := int(b / 257)
    rTerm := r8 * 5 / 255
    gTerm := g8 * 5 / 255
    bTerm := b8 * 5 / 255
    return 16 + 36*rTerm + 6*gTerm + bTerm
}

func previewTerminal(img image.Image) {
    bounds := img.Bounds()
    for y := bounds.Min.Y; y < bounds.Max.Y; y += 2 {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            c := img.At(x, y)
            colorCode := rgbToTermColor(c)
            fmt.Printf("\x1b[48;5;%dm  \x1b[0m", colorCode)
        }
        fmt.Println()
    }
}

// ---------------- Ana Program ----------------
func main() {
    var inputPath string
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

    fmt.Println("Filtre seçimi: 1-Grayscale 0-Yok")
    var choice string
    fmt.Scanln(&choice)

    filtered := applyFilter(img, choice)

    // Overlay ekleme (sağ alt köşeye kırmızı kare)
    overlay := image.NewRGBA(image.Rect(0, 0, 30, 30))
    draw.Draw(overlay, overlay.Bounds(), &image.Uniform{color.RGBA{255, 0, 0, 128}}, image.Point{}, draw.Over)
    draw.Draw(filtered, overlay.Bounds().Add(image.Pt(filtered.Bounds().Dx()-30, filtered.Bounds().Dy()-30)), overlay, image.Point{}, draw.Over)

    fmt.Println("Terminal önizleme:")
    previewTerminal(filtered)
    time.Sleep(2 * time.Second) // 2 saniye bekle

    out, err := os.Create("output.png")
    if err != nil {
        panic(err)
    }
    defer out.Close()

    png.Encode(out, filtered)
    fmt.Println("PNG dosyası kaydedildi: output.png")
}
``
/*
---

# 📌 Açıklama

1. **JPEG okuma** → `jpeg.Decode`.
2. **Filtre uygulama** → Grayscale veya yok.
3. **Overlay ekleme** → Sağ alt köşeye yarı saydam kırmızı kare.
4. **Terminal önizleme** → ASCII + 256 renk ile görseli gösterir.
5. **PNG kaydetme** → `png.Encode` ile kaydedilir.

---

Bu örnek sayesinde artık:

* JPEG’i PNG’ye dönüştürebilir,
* Filtre ve overlay ekleyebilir,
* Terminalde önizleyebilir,
* PNG olarak kaydedebilirsin.

---

İstersek bir sonraki adımda **birden fazla JPEG’den PNG serisi veya animasyonlu GIF üretip terminalde önizleme** de ekleyebiliriz.

Bunu yapalım mı?
EVET
*/
/*
Harika! 🚀 O zaman bir sonraki adımda **birden fazla JPEG’den PNG serisi oluşturma + filtre + overlay + terminal önizleme** örneğini yapalım.

Senaryo:

* Kullanıcı birden fazla JPEG dosyası girecek.
* Her kareye filtre uygulanacak (Grayscale veya yok).
* Overlay eklenebilecek (sağ alt köşe kırmızı kare).
* Terminalde her kare **önizleme** olarak gösterilecek.
* Her kare ayrı PNG olarak kaydedilecek (örn. `frame1.png`, `frame2.png`, …).

---

# 📌 Örnek Program: JPEG Serisinden PNG + Filtre + Overlay + Terminal Önizleme
*/
``go
package main

import (
    "fmt"
    "image"
    "image/color"
    "image/draw"
    "image/jpeg"
    "image/png"
    "os"
    "time"
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
        return grayscale(img)
    default:
        bounds := img.Bounds()
        rgba := image.NewRGBA(bounds)
        draw.Draw(rgba, bounds, img, image.Point{}, draw.Src)
        return rgba
    }
}

// ---------------- Terminal Önizleme ----------------
func rgbToTermColor(c color.Color) int {
    r, g, b, _ := c.RGBA()
    r8 := int(r / 257)
    g8 := int(g / 257)
    b8 := int(b / 257)
    rTerm := r8 * 5 / 255
    gTerm := g8 * 5 / 255
    bTerm := b8 * 5 / 255
    return 16 + 36*rTerm + 6*gTerm + bTerm
}

func previewTerminal(img image.Image) {
    bounds := img.Bounds()
    for y := bounds.Min.Y; y < bounds.Max.Y; y += 2 {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            c := img.At(x, y)
            colorCode := rgbToTermColor(c)
            fmt.Printf("\x1b[48;5;%dm  \x1b[0m", colorCode)
        }
        fmt.Println()
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

    fmt.Println("Filtre seçimi: 1-Grayscale 0-Yok")
    var choice string
    fmt.Scanln(&choice)

    for i, file := range jpegFiles {
        f, err := os.Open(file)
        if err != nil {
            panic(err)
        }
        img, err := jpeg.Decode(f)
        f.Close()
        if err != nil {
            panic(err)
        }

        filtered := applyFilter(img, choice)

        // Overlay ekleme (sağ alt köşe kırmızı kare)
        overlay := image.NewRGBA(image.Rect(0, 0, 30, 30))
        draw.Draw(overlay, overlay.Bounds(), &image.Uniform{color.RGBA{255, 0, 0, 128}}, image.Point{}, draw.Over)
        draw.Draw(filtered, overlay.Bounds().Add(image.Pt(filtered.Bounds().Dx()-30, filtered.Bounds().Dy()-30)), overlay, image.Point{}, draw.Over)

        fmt.Printf("Önizleme: %s\n", file)
        previewTerminal(filtered)
        time.Sleep(1 * time.Second)

        outName := fmt.Sprintf("frame%d.png", i+1)
        outFile, err := os.Create(outName)
        if err != nil {
            panic(err)
        }
        png.Encode(outFile, filtered)
        outFile.Close()

        fmt.Printf("PNG kaydedildi: %s\n", outName)
    }

    fmt.Println("Tüm PNG kareler oluşturuldu ve önizlendi.")
}
``
/*
---

# 📌 Özellikler

1. **Birden fazla JPEG** ile işlem yapılabilir.
2. **Filtre uygulanabilir** → Grayscale veya yok.
3. **Overlay eklenebilir** → Sağ alt köşeye kırmızı kare.
4. **Terminal önizleme** → Her kareyi ASCII + 256 renk ile gör.
5. **Her kare ayrı PNG** olarak kaydedilir (`frame1.png`, `frame2.png`, …).

---

İstersen bir sonraki adımda bunu **animasyonlu GIF’e dönüştürüp, terminalde canlı oynatma** ekleyerek tam bir GIF editörü haline getirebiliriz.

Bunu da yapalım mı?
*/