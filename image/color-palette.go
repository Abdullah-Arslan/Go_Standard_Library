/*
Go’nun **`image/color/palette`** paketini baştan sona anlatayım.

---

# 📌 `image/color/palette` Paketi Nedir?

* `image/color/palette`, **hazır renk paletleri** sağlayan küçük bir standart kütüphane paketidir.
* Palet: Belirli sayıda renk içeren bir listedir (`[]color.Color`).
* Özellikle **GIF** gibi renk sayısının sınırlı olduğu formatlarda veya **palet tabanlı görsellerde (`image.Paletted`)** kullanılır.
* Paletler:

  * `palette.Plan9` → 256 renk
  * `palette.WebSafe` → 216 renk

---

# 📌 Sağlanan Paletler

### 🎨 1. `palette.Plan9`

* Plan 9 işletim sisteminden alınan 256 renkli palettir.
* Renkler RGB küpünde eşit aralıklarla seçilmiştir.
* 0–255 arası indekslenir.

### 🎨 2. `palette.WebSafe`

* Eski web tarayıcılarında “güvenli” kabul edilen 216 renkten oluşur.
* RGB’nin her kanalı şu değerleri alır: `{0, 51, 102, 153, 204, 255}`.
* Bu sayede farklı tarayıcılarda aynı görünüm sağlanır (eskiden önemliydi).

---

# 📌 Kullanım Örnekleri

## 1. Plan9 Paleti ile Basit Görsel
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
    // Plan9 paletiyle 100x100 görsel
    img := image.NewPaletted(image.Rect(0, 0, 100, 100), palette.Plan9)

    // Her piksele paletten bir renk ata
    for y := 0; y < 100; y++ {
        for x := 0; x < 100; x++ {
            idx := uint8((x + y) % 256) // 0–255 arası
            img.SetColorIndex(x, y, idx)
        }
    }

    f, _ := os.Create("plan9.gif")
    defer f.Close()
    gif.Encode(f, img, nil)
}
``
/*
👉 `plan9.gif` çıktısında 256 renk paletinden üretilmiş bir görsel oluşur.

---

## 2. WebSafe Paleti ile Kareli Görsel
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
    img := image.NewPaletted(image.Rect(0, 0, 300, 300), palette.WebSafe)

    // Kareler çizelim
    size := 50
    idx := uint8(0)
    for y := 0; y < 300; y++ {
        for x := 0; x < 300; x++ {
            img.SetColorIndex(x, y, idx)
            if (x/size+y/size)%2 == 0 {
                idx = (idx + 1) % uint8(len(palette.WebSafe))
            }
        }
    }

    f, _ := os.Create("websafe.gif")
    defer f.Close()
    gif.Encode(f, img, nil)
}
``
/*
👉 Çıktı: WebSafe paletinden renklerle kareli bir GIF.

---

# 📌 Özet

* `image/color/palette` paketi hazır renk paletleri sunar.
* **Plan9 (256 renk)** ve **WebSafe (216 renk)** paletlerini içerir.
* `image.NewPaletted` ile birlikte kullanılır.
* Özellikle **GIF ve düşük renkli görseller** için idealdir.

---

İstersen sana bu paletleri kullanarak **animasyonlu GIF** üretme örneğini de gösterebilirim (ör: dönen renkli kareler). Bunu ister misin?
EVET
*/
/*
Harika! 🎉 O zaman **`image/color/palette` kullanarak animasyonlu GIF** oluşturalım.

Senaryo:

* WebSafe veya Plan9 paletini kullanacağız.
* Renkler dönecek veya kayacak şekilde kareler üreteceğiz.
* Sonuçta animasyonlu bir GIF elde edeceğiz.

---

# 📌 Örnek: Dönen Renkli GIF (Plan9 Paleti)
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
    const size = 100      // Kare boyutu
    const frames = 20     // Toplam kare sayısı

    anim := &gif.GIF{}

    for f := 0; f < frames; f++ {
        img := image.NewPaletted(image.Rect(0, 0, size, size), palette.Plan9)

        // Renkleri kaydır
        for y := 0; y < size; y++ {
            for x := 0; x < size; x++ {
                idx := uint8((x + y + f*5) % 256) // her karede renk kayar
                img.SetColorIndex(x, y, idx)
            }
        }

        anim.Image = append(anim.Image, img)
        anim.Delay = append(anim.Delay, 5) // her kare 50ms
    }

    f, _ := os.Create("anim_plan9.gif")
    defer f.Close()
    gif.EncodeAll(f, anim)
}
``
/*
---

# 📌 Açıklama

1. **`image.NewPaletted`** → Plan9 paleti ile paletli görsel oluşturuyoruz.
2. **Renk kaydırma** → `x + y + f*5` ile her karede renkleri kaydırıyoruz.
3. **`gif.GIF`** → Çok kareli animasyon oluşturuyoruz.
4. **`Delay`** → Her kare arasındaki süreyi belirliyoruz (50ms = 5).
5. **`gif.EncodeAll`** → Tüm kareleri birleştirip GIF olarak kaydediyoruz.

---

# 📌 WebSafe Paleti ile Animasyon
*/
``go
img := image.NewPaletted(image.Rect(0, 0, size, size), palette.WebSafe)
``
/*
* Aynı mantıkla kullanılır, sadece renk paleti farklıdır.

---

# 📌 Sonuç

* GIF’de renkler her karede kayarak animasyon oluşturur.
* `palette.Plan9` veya `palette.WebSafe` kullanarak **renk kontrolü** sağlanır.
* Bu yöntemle **basit animasyonlar** veya **efektler** hızlıca üretilebilir.

---

İstersen bir sonraki adımda bu **animasyonlu GIF’e filtre uygulamayı** (negative, grayscale, sepia gibi) de gösterebilirim.
Bunu yapalım mı?
EVET
*/
/*
Harika! 🎉 O zaman şimdi **animasyonlu GIF’e filtre uygulama** örneğini yapalım.

Senaryo:

* Önce Plan9 veya WebSafe paletli animasyon GIF oluşturuyoruz.
* Ardından her kareye **Negative, Grayscale veya Sepia** filtrelerinden birini uyguluyoruz.

---

# 📌 Örnek: Animasyonlu GIF + Filtre Uygulama
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

// clamp: 0-255 aralığına sıkıştır
func clamp(v int) uint8 {
    if v < 0 {
        return 0
    }
    if v > 255 {
        return 255
    }
    return uint8(v)
}

// Negative filtre
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

// Grayscale filtre
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

// Sepia filtre
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

func main() {
    const size = 100
    const frames = 20

    var choice string
    fmt.Println("Filtre Seçenekleri:")
    fmt.Println("1 - Negative")
    fmt.Println("2 - Grayscale")
    fmt.Println("3 - Sepia")
    fmt.Print("Seçiminiz: ")
    fmt.Scanln(&choice)

    anim := &gif.GIF{}

    // GIF karelerini oluştur
    for f := 0; f < frames; f++ {
        palImg := image.NewPaletted(image.Rect(0, 0, size, size), palette.Plan9)
        for y := 0; y < size; y++ {
            for x := 0; x < size; x++ {
                idx := uint8((x + y + f*5) % 256)
                palImg.SetColorIndex(x, y, idx)
            }
        }

        // Filtre uygula
        rgba := applyFilter(palImg, choice)
        newPaletted := image.NewPaletted(rgba.Bounds(), palette.Plan9)
        for y := 0; y < size; y++ {
            for x := 0; x < size; x++ {
                newPaletted.Set(x, y, rgba.At(x, y))
            }
        }

        anim.Image = append(anim.Image, newPaletted)
        anim.Delay = append(anim.Delay, 5)
    }

    outFile, _ := os.Create("filtered_anim.gif")
    defer outFile.Close()
    gif.EncodeAll(outFile, anim)

    fmt.Println("Filtre uygulanmış animasyon oluşturuldu: filtered_anim.gif")
}
``
/*
---

# 📌 Açıklama

1. **Her kare Plan9 paleti ile oluşturulur.**
2. **applyFilter** fonksiyonu seçilen filtreyi uygular (Negative, Grayscale, Sepia).
3. **Paletli GIF** oluşturmak için RGBA → Paletted dönüşümü yapılır.
4. Her kare animasyona eklenir ve `Delay` ile animasyon süresi ayarlanır.
5. Sonuç: **filtered\_anim.gif**

---

Bu yöntemle artık:

* Hem **animasyonlu GIF oluşturabilir**,
* Hem de **tüm karelerine filtre uygulayabilirsin**.

---

İstersen bir sonraki adımda bu programı **terminal tabanlı menü + kullanıcı girişli kare sayısı ve renk paleti seçimi** ile daha interaktif hâle getirebilirim.
Bunu da yapalım mı?
*/