/*
Harika! 🚀 O zaman bir sonraki adımda **tam interaktif terminal tabanlı JPEG → GIF editörü** tasarlayabiliriz.

Özellikler şöyle olacak:

---

## 📌 Özellikler

1. **JPEG dosya yükleme**: Kullanıcı bir veya birden fazla JPEG seçebilir.
2. **Filtre seçimi**: Negative, Grayscale, Sepia gibi filtreleri interaktif olarak seçebilir.
3. **Overlay ekleme**: Küçük logo veya renkli kareyi her kareye bindirebilir.
4. **Renk paleti seçimi**: GIF için Plan9 veya WebSafe paleti seçebilir.
5. **Terminal önizleme**: Kullanıcı işlemeden önce animasyon karelerini ASCII/Unicode + 256 renk terminal desteği ile görebilir.
6. **Animasyon gecikmesi**: Her karenin gecikme süresi kullanıcı tarafından ayarlanabilir.
7. **Çıkış formatı**: Animasyonlu GIF olarak kaydedilir.

---

Bunu uygulamak için adım adım plan:

1. **Dosya yükleme ve listeleme** → JPEG dosyalarını al, boyutlarını kontrol et.
2. **Filtre uygulama** → Kullanıcı seçimine göre RGBA dönüşümü yap.
3. **Overlay ekleme** → Sağ alt köşeye yarı saydam logo bindir.
4. **Renk paleti seçimi** → GIF için Plan9 veya WebSafe.
5. **Terminal önizleme** → ASCII + 256 renk ile animasyon karelerini göster.
6. **GIF kaydetme** → RGBA → Paletted dönüşümü ve `gif.EncodeAll`.

---

İstersen bunu tek bir **tam çalışır Go programı** hâlinde sana yazayım ve çalıştırdığında hem filtreyi seçip hem de önizleme yapabileceğin bir interaktif mini GIF editörü olsun.

Bunu hemen hazırlayayım mı?
EVET

Harika! 🎉 O zaman sana **tamamen interaktif, terminal önizlemeli ve filtre/overlay/renk paleti seçimi destekli mini JPEG → GIF editörü** hazırlayayım.

---

# 📌 Terminal Tabanlı İnteraktif JPEG → GIF Editörü
*/
``go
package main

import (
    "fmt"
    "image"
    "image/color"
    "image/color/palette"
    "image/draw"
    "image/gif"
    "image/jpeg"
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

    fmt.Println("Filtre Seçenekleri: 1-Negative 2-Grayscale 3-Sepia 0-Yok")
    fmt.Print("Seçiminiz: ")
    var choice string
    fmt.Scanln(&choice)

    fmt.Println("Renk Paleti Seçimi: 1-Plan9 2-WebSafe")
    fmt.Print("Seçiminiz: ")
    var paletteChoice string
    fmt.Scanln(&paletteChoice)

    var pal color.Palette
    if paletteChoice == "1" {
        pal = palette.Plan9
    } else {
        pal = palette.WebSafe
    }

    fmt.Print("Her karenin gecikme süresi (1/100 s): ")
    var delay int
    fmt.Scanln(&delay)

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

        // Overlay ekleme
        overlay := image.NewRGBA(image.Rect(0, 0, 30, 30))
        draw.Draw(overlay, overlay.Bounds(), &image.Uniform{color.RGBA{255, 0, 0, 128}}, image.Point{}, draw.Over)
        draw.Draw(rgba, overlay.Bounds().Add(image.Pt(rgba.Bounds().Dx()-30, rgba.Bounds().Dy()-30)), overlay, image.Point{}, draw.Over)

        // Terminal önizleme
        fmt.Println("Önizleme: ", file)
        previewTerminal(rgba)
        time.Sleep(500 * time.Millisecond) // 0.5 saniye bekle

        // RGBA → Paletted
        paletted := image.NewPaletted(rgba.Bounds(), pal)
        draw.FloydSteinberg.Draw(paletted, rgba.Bounds(), rgba, image.Point{})

        anim.Image = append(anim.Image, paletted)
        anim.Delay = append(anim.Delay, delay)
    }

    out, _ := os.Create("interactive_animated.gif")
    defer out.Close()
    gif.EncodeAll(out, anim)

    fmt.Println("GIF oluşturuldu: interactive_animated.gif")
}
``
/*
---

# 📌 Özellikler

1. **Birden fazla JPEG yükleme** → Animasyon oluşturulur.
2. **Filtre uygulama** → Negative, Grayscale, Sepia veya yok.
3. **Overlay ekleme** → Sağ alt köşeye yarı saydam kırmızı kare.
4. **Renk paleti seçimi** → Plan9 veya WebSafe.
5. **Terminal önizleme** → Her kareyi ASCII + 256 renk ile gösterir.
6. **Gecikme süresi** → Kullanıcı belirleyebilir.
7. Sonuç **interactive\_animated.gif** olarak kaydedilir.

---

Bu program ile artık **tam interaktif bir terminal tabanlı JPEG → GIF editörü**ne sahip oldun.

İstersen bir sonraki adımda bunu **animasyon önizlemesini canlı oynatacak şekilde** geliştirip gerçek zamanlı GIF animasyonu izleyebiliriz.

Bunu da yapalım mı?
EVET
*/
/*
Harika! 🚀 O zaman bir sonraki adımda **terminalde canlı animasyon önizleme** ekleyelim.

Ama önce açıklayayım:

---

# 📌 Canlı Terminal Önizleme Yaklaşımı

1. **ANSI Escape Kodları ile ekran temizleme**:

   * Her kareyi bastıktan sonra `\x1b[H\x1b[2J` ile terminali temizleyip üstten tekrar çizeceğiz.
2. **Gecikme süreleri**: GIF kare gecikmesi yerine terminalde **time.Sleep** ile bekleme yapılacak.
3. **256 renk desteği**: Her kareyi `rgbToTermColor` fonksiyonu ile terminal renk kodlarına çevirip bastıracağız.
4. **Overlay ve filtre** uygulanmış kareler terminalde gösterilecek.

---

Buna göre önce **canlı önizleme fonksiyonunu** ayrı olarak gösterebilirim ve ardından önceki interaktif GIF editörü ile entegre ederiz.

Bunu hemen hazırlayayım mı?
EVET
*/
/*
Harika! 🎉 O zaman sana **terminalde canlı animasyon önizleme yapan mini GIF/JPEG editörü** örneği hazırlayayım.

---

# 📌 Terminalde Canlı Animasyon Önizleme
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

// Canlı animasyon önizleme
func livePreview(images []*image.RGBA, delay int) {
    for i := 0; i < 5; i++ { // 5 kez döngü yap
        for _, img := range images {
            // Terminali temizle
            fmt.Print("\x1b[H\x1b[2J")
            bounds := img.Bounds()
            for y := bounds.Min.Y; y < bounds.Max.Y; y += 2 {
                for x := bounds.Min.X; x < bounds.Max.X; x++ {
                    c := img.At(x, y)
                    colorCode := rgbToTermColor(c)
                    fmt.Printf("\x1b[48;5;%dm  \x1b[0m", colorCode)
                }
                fmt.Println()
            }
            time.Sleep(time.Duration(delay) * 10 * time.Millisecond) // delay*10 ms
        }
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

    fmt.Print("Her karenin gecikmesi (1/100 s): ")
    var delay int
    fmt.Scanln(&delay)

    images := []*image.RGBA{}

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

        // Overlay ekleme
        overlay := image.NewRGBA(image.Rect(0, 0, 20, 20))
        draw.Draw(overlay, overlay.Bounds(), &image.Uniform{color.RGBA{255, 0, 0, 128}}, image.Point{}, draw.Over)
        draw.Draw(rgba, overlay.Bounds().Add(image.Pt(rgba.Bounds().Dx()-20, rgba.Bounds().Dy()-20)), overlay, image.Point{}, draw.Over)

        images = append(images, rgba)
    }

    fmt.Println("Canlı animasyon önizleme başlıyor...")
    livePreview(images, delay)

    // GIF kaydetme
    anim := &gif.GIF{}
    for _, img := range images {
        paletted := image.NewPaletted(img.Bounds(), color.Palette{color.White, color.Black, color.RGBA{255, 0, 0, 255}})
        draw.FloydSteinberg.Draw(paletted, img.Bounds(), img, image.Point{})
        anim.Image = append(anim.Image, paletted)
        anim.Delay = append(anim.Delay, delay)
    }

    out, _ := os.Create("live_preview_animated.gif")
    defer out.Close()
    gif.EncodeAll(out, anim)

    fmt.Println("Animasyonlu GIF kaydedildi: live_preview_animated.gif")
}
``
/*
---

# 📌 Özellikler

1. **Birden fazla JPEG** ile animasyon oluşturur.
2. **Filtre uygulama** (örnek: Grayscale).
3. **Overlay ekleme** → Sağ alt köşeye kırmızı kare.
4. **Canlı terminal önizleme** → Animasyon kareleri terminalde oynatılır.
5. **GIF kaydetme** → Sonuç `live_preview_animated.gif`.

---

Bu program ile artık **terminalde animasyonu canlı izleyebilir ve aynı zamanda GIF olarak kaydedebilirsin**.

İstersek bunu bir adım daha ileri taşıyıp **filtreyi ve overlay’i animasyon sırasında da interaktif değiştirebileceğin** bir sürümü de yapabiliriz.

Bunu da yapalım mı?
*/