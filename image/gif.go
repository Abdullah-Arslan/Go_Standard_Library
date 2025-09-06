/*
Go’nun **`image/gif`** paketini baştan sona detaylı şekilde açıklayalım ve örneklerle gösterelim.

---

# 📌 `image/gif` Paketi Nedir?

* `image/gif` paketi, Go’nun standart kütüphanesinde **GIF formatındaki görselleri okuma, oluşturma ve kaydetme** işlemleri için kullanılır.
* Özellikle **animasyonlu GIF** veya **düşük renkli (paletli) görseller** için uygundur.
* GIF’ler, palet tabanlı (256 renk) görseller olduğu için `image/color/palette` ve `image/draw` paketleri ile birlikte sık kullanılır.

---

# 📌 Temel Fonksiyonlar

1. **`gif.Decode(r io.Reader) (image.Image, error)`**

   * Tek kare GIF dosyasını okur.
   * Dönen değer: `image.Image` (RGBA veya Paletted).

2. **`gif.DecodeAll(r io.Reader) (*gif.GIF, error)`**

   * Çok kareli GIF dosyalarını okur.
   * Dönen değer: `*gif.GIF` → `Image []image.Paletted` ve `Delay []int`.

3. **`gif.Encode(w io.Writer, m image.Image, o *gif.Options) error`**

   * Tek kare GIF oluşturur ve kaydeder.

4. **`gif.EncodeAll(w io.Writer, g *gif.GIF) error`**

   * Çok kareli GIF oluşturur ve kaydeder.

---

# 📌 Önemli Tipler

* **`gif.GIF` struct**
*7
``go
type GIF struct {
    Image []*image.Paletted // Her kare
    Delay []int             // Her karenin gecikme süresi (1/100 s)
    LoopCount int           // Animasyonun kaç kez dönmesi
    Disposal []byte         // Her kare sonrası işlem türü
}
``
/*
* **`image.Paletted`** → Palet tabanlı görsel.
* **`image/color/palette`** → GIF renk paletlerini oluşturmak için kullanılır.

---

# 📌 Örnekler

## 1️⃣ Tek Kare GIF Oluşturma
*/
``go
package main

import (
    "image"
    "image/color"
    "image/color/palette"
    "image/gif"
    "os"
)

func main() {
    img := image.NewPaletted(image.Rect(0, 0, 100, 100), palette.Plan9)

    // Renkleri ata
    for y := 0; y < 100; y++ {
        for x := 0; x < 100; x++ {
            img.SetColorIndex(x, y, uint8((x + y) % 256))
        }
    }

    f, _ := os.Create("single.gif")
    defer f.Close()
    gif.Encode(f, img, nil)
}
``
/*
---

## 2️⃣ Animasyonlu GIF Oluşturma
*/
``go
package main

import (
    "image"
    "image/color/palette"
    "image/gif"
    "os"
)

func main() {
    anim := &gif.GIF{}
    frames := 20
    size := 100

    for f := 0; f < frames; f++ {
        img := image.NewPaletted(image.Rect(0, 0, size, size), palette.Plan9)

        for y := 0; y < size; y++ {
            for x := 0; x < size; x++ {
                img.SetColorIndex(x, y, uint8((x + y + f*5) % 256))
            }
        }

        anim.Image = append(anim.Image, img)
        anim.Delay = append(anim.Delay, 5) // 50ms
    }

    out, _ := os.Create("anim.gif")
    defer out.Close()
    gif.EncodeAll(out, anim)
}
``
/*
---

## 3️⃣ Var olan GIF’i Okuma
*/
``go
package main

import (
    "fmt"
    "image/gif"
    "os"
)

func main() {
    f, _ := os.Open("anim.gif")
    defer f.Close()

    g, _ := gif.DecodeAll(f)

    fmt.Println("Kare sayısı:", len(g.Image))
    fmt.Println("Her karenin gecikme süresi (1/100 s):", g.Delay)
}
``
/*
---

# 📌 Özet

1. `image/gif` **GIF okuma ve yazma** için kullanılır.
2. Tek kare veya çok kareli GIF’ler için farklı fonksiyonlar vardır (`Encode` vs `EncodeAll`).
3. **Palet tabanlı görseller** (`image.Paletted`) ile çalışır.
4. Animasyon oluşturmak için `gif.GIF.Image` ve `gif.GIF.Delay` kullanılır.
5. `image/color/palette` ve `image/draw` ile birlikte **filtreler, overlay, logo** eklemek mümkün.

---

İstersen bir sonraki adımda sana **`gif` + `draw` + `palette` paketlerini birleştirip animasyonlu GIF’e filtre ve overlay uygulayan tam bir örnek program** da gösterebilirim.
Bunu yapalım mı?
EVET


Harika! 🎉 O zaman şimdi **tam bir örnek program yapalım**:

* Mevcut GIF’e veya yeni animasyonlu GIF’e
* **Filtre (Negative, Grayscale, Sepia)** ve
* **Overlay (logo / renkli kare)** uygulayacağız.

Böylece **`gif` + `draw` + `palette`** paketlerini birleştirmiş olacağız.

---

# 📌 Örnek: Animasyonlu GIF + Filtre + Overlay
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
    // Kullanıcıdan giriş
    var choice string
    fmt.Println("Filtre Seçenekleri:")
    fmt.Println("1 - Negative")
    fmt.Println("2 - Grayscale")
    fmt.Println("3 - Sepia")
    fmt.Print("Seçiminiz: ")
    fmt.Scanln(&choice)

    // GIF oluşturulacak boyut ve kare sayısı
    const size = 100
    const frames = 20

    anim := &gif.GIF{}

    // Overlay (küçük yarı saydam kare / logo)
    overlay := image.NewRGBA(image.Rect(0, 0, 30, 30))
    draw.Draw(overlay, overlay.Bounds(), &image.Uniform{color.RGBA{255, 0, 0, 128}}, image.Point{}, draw.Src)

    // Her kareyi oluştur
    for f := 0; f < frames; f++ {
        // Palet tabanlı animasyon kareleri
        palImg := image.NewPaletted(image.Rect(0, 0, size, size), palette.Plan9)

        // Renkleri kaydır
        for y := 0; y < size; y++ {
            for x := 0; x < size; x++ {
                palImg.SetColorIndex(x, y, uint8((x+y+f*5)%256))
            }
        }

        // Filtre uygula
        rgba := applyFilter(palImg, choice)

        // Overlay bindir
        offset := image.Pt(size-30, size-30) // sağ alt köşe
        draw.Draw(rgba, overlay.Bounds().Add(offset), overlay, image.Point{}, draw.Over)

        // RGBA → Paletted dönüşümü
        newPaletted := image.NewPaletted(rgba.Bounds(), palette.Plan9)
        draw.FloydSteinberg.Draw(newPaletted, rgba.Bounds(), rgba, image.Point{})

        anim.Image = append(anim.Image, newPaletted)
        anim.Delay = append(anim.Delay, 5)
    }

    // Yeni GIF'i kaydet
    outFile, _ := os.Create("final_anim.gif")
    defer outFile.Close()
    gif.EncodeAll(outFile, anim)

    fmt.Println("Animasyon oluşturuldu: final_anim.gif")
}
``
/*
---

# 📌 Özellikler

1. **Animasyonlu GIF** üretiyor (20 kare).
2. Kullanıcı seçimine göre **Negative, Grayscale veya Sepia** filtreleri uygulanıyor.
3. Sağ alt köşeye **yarı saydam overlay** bindiriliyor (logo gibi düşünülebilir).
4. RGBA → Paletted dönüşümü ile GIF uyumluluğu sağlanıyor.
5. Tüm işlemler **`gif` + `draw` + `palette`** paketleri ile yapılmıştır.

---

İstersen bir sonraki adımda bunu **terminal tabanlı interaktif mini GIF editörü** hâline getirip:

* Kullanıcı girişli GIF yükleme
* Filtre seçimi
* Overlay ekleme
* Kare sayısı ve hız ayarı

gibi özellikler ekleyebiliriz.

Bunu da yapalım mı?
EVET
*/
/*
Harika! 🎉 O zaman adım adım **terminal tabanlı interaktif mini GIF editörü** yapalım.

Özellikler:

1. Kullanıcı GIF yükleyebilir veya yeni animasyon oluşturabilir.
2. Filtre seçimi: Negative, Grayscale, Sepia.
3. Overlay ekleyebilir (logo/küçük kare).
4. Kare sayısı ve animasyon hızı ayarlanabilir.
5. Sonuç `final_anim.gif` olarak kaydedilir.

---

# 📌 Terminal Tabanlı Mini GIF Editörü
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
    "os"
)

// ---------------- Filtre Fonksiyonları ----------------
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

// ---------------- Ana Program ----------------
func main() {
    var choice, inputPath string
    var frames, delay int

    fmt.Print("Mevcut GIF yüklemek ister misiniz? (evet/hayır): ")
    var load string
    fmt.Scanln(&load)

    var anim *gif.GIF

    if load == "evet" {
        fmt.Print("GIF dosya yolu: ")
        fmt.Scanln(&inputPath)
        f, err := os.Open(inputPath)
        if err != nil {
            panic(err)
        }
        defer f.Close()
        g, err := gif.DecodeAll(f)
        if err != nil {
            panic(err)
        }
        anim = g
        frames = len(anim.Image)
        delay = 5
        fmt.Println("GIF yüklendi:", frames, "kare")
    } else {
        // Yeni animasyon oluştur
        fmt.Print("Kare sayısı: ")
        fmt.Scanln(&frames)
        fmt.Print("Her karenin gecikmesi (1/100 s): ")
        fmt.Scanln(&delay)
        anim = &gif.GIF{}
        size := 100
        for f := 0; f < frames; f++ {
            img := image.NewPaletted(image.Rect(0, 0, size, size), palette.Plan9)
            for y := 0; y < size; y++ {
                for x := 0; x < size; x++ {
                    img.SetColorIndex(x, y, uint8((x+y+f*5)%256))
                }
            }
            anim.Image = append(anim.Image, img)
            anim.Delay = append(anim.Delay, delay)
        }
    }

    // Filtre seçimi
    fmt.Println("Filtre Seçenekleri:")
    fmt.Println("1 - Negative")
    fmt.Println("2 - Grayscale")
    fmt.Println("3 - Sepia")
    fmt.Print("Seçiminiz: ")
    fmt.Scanln(&choice)

    // Overlay oluştur
    overlay := image.NewRGBA(image.Rect(0, 0, 30, 30))
    draw.Draw(overlay, overlay.Bounds(), &image.Uniform{color.RGBA{255, 0, 0, 128}}, image.Point{}, draw.Src)

    // Her kareyi işleme
    for i, frame := range anim.Image {
        rgba := image.NewRGBA(frame.Bounds())
        draw.Draw(rgba, rgba.Bounds(), frame, image.Point{}, draw.Src)

        // Filtre uygula
        filtered := applyFilter(rgba, choice)

        // Overlay bindir (sağ alt köşe)
        offset := image.Pt(rgba.Bounds().Dx()-30, rgba.Bounds().Dy()-30)
        draw.Draw(filtered, overlay.Bounds().Add(offset), overlay, image.Point{}, draw.Over)

        // RGBA → Paletted
        newPal := image.NewPaletted(filtered.Bounds(), palette.Plan9)
        draw.FloydSteinberg.Draw(newPal, filtered.Bounds(), filtered, image.Point{})
        anim.Image[i] = newPal
    }

    // Kaydet
    outFile, _ := os.Create("interactive_anim.gif")
    defer outFile.Close()
    gif.EncodeAll(outFile, anim)

    fmt.Println("GIF oluşturuldu: interactive_anim.gif")
}
``
/*
---

# 📌 Özellikler

1. **Kullanıcı GIF yükleyebilir** veya yeni animasyon oluşturabilir.
2. **Filtre seçimi**: Negative / Grayscale / Sepia.
3. **Overlay (yarı saydam logo/kare)** uygulanır.
4. **Kare sayısı ve hız** kullanıcı tarafından belirlenir.
5. Sonuç **`interactive_anim.gif`** olarak kaydedilir.

---

Bu mini editörle artık **GIF animasyonlarına filtre ve overlay eklemek** terminal üzerinden tamamen interaktif hâle geldi.

İstersen bir sonraki adımda bunu **terminal önizleme + renk paleti seçimi** ekleyerek daha profesyonel hâle getirebiliriz.

Bunu da yapalım mı?
EVET
*/
/*
Harika! 🎉 O zaman önce **terminal önizleme fonksiyonunu** göstereyim.

Terminal önizleme, GIF karelerini veya herhangi bir görseli **ASCII veya Unicode blok karakterleri** ve terminal renk kodlarıyla gösterecek.

---

# 📌 Örnek: Terminal Önizleme Fonksiyonu
*/
``go
package main

import (
    "fmt"
    "image"
    "image/color"
)

// RGB → 256 renk terminal koduna dönüşüm
func rgbToTermColor(c color.Color) int {
    r, g, b, _ := c.RGBA()
    r8 := int(r / 257)
    g8 := int(g / 257)
    b8 := int(b / 257)

    // Basit 6x6x6 renk küpü (16-231 arası)
    rTerm := r8 * 5 / 255
    gTerm := g8 * 5 / 255
    bTerm := b8 * 5 / 255

    return 16 + 36*rTerm + 6*gTerm + bTerm
}

// Görseli terminalde ASCII ile göster
func previewTerminal(img image.Image) {
    bounds := img.Bounds()
    for y := bounds.Min.Y; y < bounds.Max.Y; y += 2 { // 2 satır = 1 blok yüksekliği
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            c := img.At(x, y)
            colorCode := rgbToTermColor(c)
            fmt.Printf("\x1b[48;5;%dm  \x1b[0m", colorCode) // arka plan rengi
        }
        fmt.Println()
    }
}

func main() {
    // Örnek: 8x8 renkli blok
    img := image.NewRGBA(image.Rect(0, 0, 16, 16))
    for y := 0; y < 16; y++ {
        for x := 0; x < 16; x++ {
            img.Set(x, y, color.RGBA{uint8(x * 16), uint8(y * 16), 128, 255})
        }
    }

    previewTerminal(img)
}
``
/*
---

# 📌 Açıklama

1. **`rgbToTermColor`** → 24-bit RGB → 256 renk terminal koduna dönüştürür.
2. **`\x1b[48;5;%dm`** → ANSI escape kodu ile terminal arka plan rengini ayarlar.
3. **`2 satır = 1 blok`** → Görseli daha kare oranlı göstermek için satır atlıyoruz.
4. `previewTerminal(img)` ile her görseli terminalde önizleyebilirsin.

---

Bu yöntemi GIF karelerine uygularsak:

* **Kullanıcı animasyonun bir veya birkaç karesini terminalde görebilir.**
* **Filtre ve overlay uygulandıktan sonra önizleme yapılabilir.**

---

İstersen bir sonraki adımda bunu **önceki interaktif GIF editörü ile birleştirip terminal önizlemeli mini GIF editörü** yapalım.
Bunu da yapalım mı?
*/