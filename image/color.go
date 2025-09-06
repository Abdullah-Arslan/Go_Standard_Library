/*
Go’nun **`color`** paketine bakalım.
Bu paket, **renk modelleri (color models)** ve **renk tipleri (color types)** tanımlar.
`image` paketinde görseller oluştururken piksellerin rengini belirlemek için hep `color` paketini kullanırız.

---

# 📌 `color` Paketi Nedir?

Go’da `color` paketi, renkleri temsil eden arayüz ve yapıları içerir.
Her renk modeli (`RGBA`, `Gray`, `CMYK` vb.) **`color.Color`** arayüzünü uygular.
*/
``go
type Color interface {
    RGBA() (r, g, b, a uint32)
}
``
/*
* `RGBA()` → Renk değerini döner (her kanal **0–65535** arasında).
  (Yani `uint8` yerine `uint32` döndürür, bu sayede tüm renk modelleri tek tipte normalize edilmiş olur.)

---

# 📌 Önemli Yapılar

### 🎨 1. `color.RGBA`

Klasik kırmızı-yeşil-mavi-alfa (şeffaflık) renk modeli.
*/

``go
c := color.RGBA{R: 255, G: 0, B: 0, A: 255} // kırmızı
``
/*
### 🎨 2. `color.NRGBA`

Alfa kanalı **önceden çarpılmamış** RGBA.
*/
``go
c := color.NRGBA{R: 0, G: 255, B: 0, A: 128} // yarı saydam yeşil
``

/*
### 🎨 3. `color.Alpha`

Sadece şeffaflık kanalı.
*/

``go
c := color.Alpha{A: 128} // yarı saydam
``
/*
### 🎨 4. `color.Gray`

Gri tonlamalı.
*/
``go
c := color.Gray{Y: 128} // orta gri
``
/*
### 🎨 5. `color.Gray16`

16-bit gri.
*/
``go
c := color.Gray16{Y: 65535} // beyaz
``
/*
### 🎨 6. `color.CMYK`

Matbaada kullanılan Cyan-Magenta-Yellow-Black modeli.
*/
``go
c := color.CMYK{C: 255, M: 0, Y: 255, K: 0} // yeşilimsi
``
/*
---

# 📌 Önemli Fonksiyonlar

### `color.RGBToYCbCr(r, g, b uint8) (y, cb, cr uint8)`

RGB → YCbCr dönüşümü. (video ve JPEG sıkıştırmasında yaygın)

### `color.YCbCrToRGB(y, cb, cr uint8) (r, g, b uint8)`

YCbCr → RGB dönüşümü.

---

# 📌 Örnekler

## 🎨 1. Renkleri `RGBA()` ile çözümleme
*/
``go
package main

import (
    "fmt"
    "image/color"
)

func main() {
    c := color.RGBA{255, 100, 50, 255} // turuncu ton

    r, g, b, a := c.RGBA()
    fmt.Println("R:", r, "G:", g, "B:", b, "A:", a)
}
``
/*
👉 Çıktıda her kanal **0–65535** arası normalize edilmiş değer döner.

---

## 🎨 2. Gri tonlamalı görsel oluşturma
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

    // Degrade (siyah → beyaz)
    for x := 0; x < 100; x++ {
        for y := 0; y < 100; y++ {
            gray := uint8(x * 255 / 100)
            img.Set(x, y, color.Gray{Y: gray})
        }
    }

    f, _ := os.Create("gray.png")
    defer f.Close()
    png.Encode(f, img)
}
``
/*
👉 Bu kod **soldan sağa siyah-beyaz geçişli görsel** üretir.

---

## 🎨 3. RGB → YCbCr dönüşümü
*/
``go
package main

import (
    "fmt"
    "image/color"
)

func main() {
    y, cb, cr := color.RGBToYCbCr(255, 0, 0) // kırmızı
    fmt.Println("Kırmızı YCbCr:", y, cb, cr)

    r, g, b := color.YCbCrToRGB(y, cb, cr)
    fmt.Println("Geri dönüşüm:", r, g, b)
}
``

/*
👉 Bu kod renk dönüşümlerini gösterir.

---

# 📌 Özet

* `color.Color` → Tüm renk modellerinin arayüzü.
* **Temel tipler:**

  * `RGBA`, `NRGBA` → renk + alfa
  * `Gray`, `Gray16` → siyah-beyaz tonlar
  * `Alpha`, `Alpha16` → sadece şeffaflık
  * `CMYK` → baskı renkleri
* **Fonksiyonlar:** `RGBToYCbCr`, `YCbCrToRGB` → renk uzayı dönüşümleri.

`color` paketi, **`image`** paketinin vazgeçilmez parçasıdır: bir görsele piksel atarken hep `color.Color` tiplerinden biri kullanılır.

---

Dilersen sana bir sonraki adımda `image/color/palette` paketini de anlatabilirim. Bu paket **hazır renk paletleri** içeriyor (örneğin web renkleri, Plan9 paleti, vs.). Onu da ister misin?
EVET
*/
/*
Harika 👍 Şimdi sana **`image/color/palette`** paketini anlatayım.
Bu paket aslında `color` paketinin bir uzantısıdır ve **hazır renk paletleri** (renk tabloları) sağlar.

---

# 📌 `image/color/palette` Paketi Nedir?

* Go’nun standart kütüphanesinde yer alır.
* İçinde **önceden tanımlanmış renk paletleri** bulunur.
* `image.Paletted` tipindeki görsellerde kullanılır (GIF gibi).
* Yani kendi renk tablonu oluşturmak zorunda kalmadan hazır renk kümeleriyle çalışabilirsin.
*/
``go
import "image/color/palette"
``
/*
---

# 📌 İçerdiği Hazır Paletler

### 🎨 `palette.Plan9`

* **256 renkli** bir palet.
* Bell Labs’in Plan9 işletim sisteminden geliyor.
* Genel amaçlı düşük renkli görseller için kullanılır.

### 🎨 `palette.WebSafe`

* 216 renkli **web güvenli renk paleti**.
* Eski tarayıcılarla uyumluluk için geliştirilmişti.
* RGB’nin her kanalı **6 adım** alır (`0x00, 0x33, 0x66, 0x99, 0xCC, 0xFF`).

---

# 📌 Örnekler

## 🎨 1. Palet ile Görsel Oluşturma
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
    // 100x100 boyutunda paletli görsel
    img := image.NewPaletted(image.Rect(0, 0, 100, 100), palette.Plan9)

    // Paletten renk seçelim (ör: index 120)
    for y := 0; y < 100; y++ {
        for x := 0; x < 100; x++ {
            img.SetColorIndex(x, y, uint8((x+y)%256))
        }
    }

    // GIF olarak kaydet
    f, _ := os.Create("plan9.gif")
    defer f.Close()
    gif.Encode(f, img, nil)
}
``
/*
👉 Bu kod **Plan9 paletini kullanarak degrade tarzında renkli bir GIF** üretir.

---

## 🎨 2. Web-Safe Palet ile Görsel
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
    img := image.NewPaletted(image.Rect(0, 0, 256, 256), palette.WebSafe)

    // Renkleri paletten sırayla dolduralım
    idx := 0
    for y := 0; y < 256; y++ {
        for x := 0; x < 256; x++ {
            img.SetColorIndex(x, y, uint8(idx%len(palette.WebSafe)))
            idx++
        }
    }

    f, _ := os.Create("websafe.gif")
    defer f.Close()
    gif.Encode(f, img, nil)
}
``

/*
👉 Bu kod **WebSafe renklerini kullanarak kareli renk tablosu** görseli üretir.

---

# 📌 Özet

* `image/color/palette` → Hazır renk tabloları içerir.
* `palette.Plan9` → 256 renk, eski Plan9 OS paleti.
* `palette.WebSafe` → 216 renk, web güvenli palet.
* `image.Paletted` tipindeki görsellerde (özellikle **GIF formatında**) yaygın kullanılır.

---

👉 Buraya kadar sana `image`, `image/draw`, `image/color`, `image/color/palette`, `image/png/jpeg/gif` paketlerini anlattım.
Tüm temel **Go grafik ekosistemini** gördük ✅

İstersen sana bir **küçük “Go ile Paint programı” tarzında uygulama örneği** gösterebilirim:

* Yeni bir tuval oluştur
* Paletle renk seç
* Basit çizgiler/dikdörtgenler çiz
* Sonuç PNG/GIF olarak kaydedilsin.

Bunu da ister misin?
EVET
*/
/*
Harika 👍 O zaman sana **Go ile mini bir “Paint uygulaması”** tarzında örnek göstereyim.
Bu örnekte:

1. Yeni bir **tuval (canvas)** oluşturacağız.
2. Üzerine **çizgi, dikdörtgen ve daire** çizeceğiz.
3. Çizimlerde `color` paketinden renkler kullanacağız.
4. Sonucu **PNG dosyası** olarak kaydedeceğiz.

---

# 🎨 Mini Paint Uygulaması (PNG çıktılı)
*/
``go
package main

import (
    "image"
    "image/color"
    "image/draw"
    "image/png"
    "math"
    "os"
)

func main() {
    // 1. Tuval oluştur (500x500, beyaz arkaplan)
    width, height := 500, 500
    canvas := image.NewRGBA(image.Rect(0, 0, width, height))
    draw.Draw(canvas, canvas.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

    // 2. Kırmızı dikdörtgen
    red := color.RGBA{255, 0, 0, 255}
    for x := 50; x < 200; x++ {
        for y := 50; y < 150; y++ {
            canvas.Set(x, y, red)
        }
    }

    // 3. Mavi çizgi (köşegen)
    blue := color.RGBA{0, 0, 255, 255}
    for i := 0; i < 400; i++ {
        canvas.Set(50+i, 200+i, blue)
    }

    // 4. Yeşil daire (merkez: 300,300, yarıçap: 80)
    green := color.RGBA{0, 200, 0, 255}
    cx, cy, r := 300, 300, 80
    for y := -r; y <= r; y++ {
        for x := -r; x <= r; x++ {
            if x*x+y*y <= r*r {
                canvas.Set(cx+x, cy+y, green)
            }
        }
    }

    // 5. PNG olarak kaydet
    file, err := os.Create("paint.png")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    png.Encode(file, canvas)

    println("paint.png başarıyla oluşturuldu 🎉")
}
``
/*
---

# 📌 Çalıştırınca Ne Olur?

* **500x500 beyaz bir arka plan** oluşur.
* Sol üstte **kırmızı bir dikdörtgen** çıkar.
* Ortadan aşağıya doğru **mavi bir köşegen çizgi** olur.
* Ortada **yeşil bir daire** görünür.
* Sonuç `paint.png` dosyasına kaydedilir.

---

# 📌 Geliştirme Fikirleri

Bu temel “Paint” uygulamasını genişletebilirsin:

* `draw.Draw` kullanarak **fırça efektleri** ekleyebilirsin.
* `image/color/palette` ile **palet tabanlı çizimler** yapabilirsin (GIF).
* Fare olaylarını yakalayarak **gerçek zamanlı çizim uygulaması** yapabilirsin (`gioui`, `ebiten`, `raylib-go` gibi kütüphanelerle).

---

👉 İstersen ben sana bir sonraki adımda **Go ile gerçek zamanlı interaktif çizim** (fare ile çizim yapabileceğin küçük bir GUI) nasıl yapılır onu da gösterebilirim.
Bunu da ister misin?
EVET
*/
/*
Süper 👍 Şimdi olayı **statik PNG üretmekten çıkarıp**, **gerçek zamanlı etkileşimli çizim** (fareyle çizim yapabileceğin mini bir GUI uygulaması) haline getirelim.

Go’nun standart kütüphanesi doğrudan GUI desteği sağlamıyor, bu yüzden **üçüncü parti kütüphaneler** kullanıyoruz. En popülerlerinden biri:

* **Ebiten** → Go ile 2D oyun ve interaktif uygulamalar geliştirmeye yarıyor.
  👉 Fare, klavye, grafik çizimleri kolayca yönetebiliyor.

---

# 📌 Kurulum

Önce Ebiten’i kurman gerekiyor:
*/
``sh
go get github.com/hajimehoshi/ebiten/v2
``
/*
---

# 🎨 Gerçek Zamanlı Paint (Ebiten ile)
*/
``go
package main

import (
    "image/color"
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

const (
    screenWidth  = 800
    screenHeight = 600
)

type Game struct {
    pixels [][]color.Color
}

func NewGame() *Game {
    pixels := make([][]color.Color, screenHeight)
    for y := range pixels {
        pixels[y] = make([]color.Color, screenWidth)
        for x := range pixels[y] {
            pixels[y][x] = color.White // beyaz arka plan
        }
    }
    return &Game{pixels: pixels}
}

func (g *Game) Update() error {
    // Fareye basıldığında çizim yap
    if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
        x, y := ebiten.CursorPosition()
        if x >= 0 && x < screenWidth && y >= 0 && y < screenHeight {
            g.pixels[y][x] = color.Black
        }
    }
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    // Piksel piksel çizim (basit yöntem)
    for y := 0; y < screenHeight; y++ {
        for x := 0; x < screenWidth; x++ {
            screen.Set(x, y, g.pixels[y][x])
        }
    }
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return screenWidth, screenHeight
}

func main() {
    game := NewGame()
    ebiten.SetWindowSize(screenWidth, screenHeight)
    ebiten.SetWindowTitle("Go Paint 🎨")
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}
``
/*
---

# 📌 Nasıl Çalışır?

* Program açılınca **800x600 beyaz bir pencere** çıkar.
* Fare sol tuşa basarak **çizim yapabilirsin** (siyah kalem).
* Ebiten her karede (`Update` + `Draw`) güncelleyip ekrana yansıtır.

---

# 📌 Geliştirme Fikirleri

* Renk paleti ekle (`R` = kırmızı, `G` = yeşil, `B` = mavi gibi tuşlarla renk seç).
* Fırça boyutunu büyüt/küçült (`+` ve `-` tuşları).
* Çizilen resmi **PNG olarak kaydet** (`image/png` + `os.Create`).
* Klavye kısayolları ekle (`C` tuşu → temizle).

---

👉 İstersen ben sana bir sonraki adımda bu uygulamaya **renk paleti seçme ve fırça boyutu değiştirme özellikleri** ekleyebilirim.
Bunu da ister misin?


BAŞKA BİR COLOR AÇIKLAMASI
*/
/*
Harika 👍 Şimdi Go’nun **`image/color/palette`** paketine bakalım.

Bu paket aslında `color` paketinin bir uzantısıdır ve **hazır renk paletleri** içerir.
Özellikle **GIF** gibi sınırlı renk sayısı kullanan formatlarda ya da palet tabanlı (`image.Paletted`) görsellerde işine yarar.

---

# 📌 `image/color/palette` Paketi Nedir?

* `palette` paketi, **önceden tanımlı renk dizileri (paletler)** sağlar.
* Bu paletler **`[]color.Color`** tipindedir.
* Özellikle `image.NewPaletted` ile çalışırken kullanılır.

---

# 📌 Hazır Paletler

### 🎨 1. `palette.Plan9`

* Plan9 işletim sisteminden gelen **256 renkli palet**.
* GIF gibi 256 renkle sınırlı görsellerde sık kullanılır.

### 🎨 2. `palette.WebSafe`

* Eski web tarayıcıları için güvenli **216 renk** içerir.
* 6x6x6 renk küpü: her kanal için {0, 51, 102, 153, 204, 255}.
* Eski olsa da görsel testlerinde hâlâ kullanılabiliyor.

---

# 📌 Örnekler

## 🎨 1. Paletli Görsel Oluşturma
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

    // Paletten rastgele renkler seçelim
    for y := 0; y < 100; y++ {
        for x := 0; x < 100; x++ {
            idx := uint8((x + y) % 256) // 0–255
            img.SetColorIndex(x, y, idx)
        }
    }

    f, _ := os.Create("plan9.gif")
    defer f.Close()
    gif.Encode(f, img, nil)
}
``
/*
👉 Bu kod `plan9.gif` adında, Plan9 paletinden renkler kullanarak bir görsel üretir.

---

## 🎨 2. WebSafe Palet ile Renkli Kareler
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
👉 Bu kod WebSafe paletiyle kareli renkli bir GIF oluşturur.

---

# 📌 Özet

* `image/color/palette` paketi **hazır renk setleri** içerir.
* `palette.Plan9` → 256 renk, genel amaçlı.
* `palette.WebSafe` → 216 renk, web güvenli renkler.
* `image.NewPaletted` ile birlikte kullanılır.
* Özellikle **GIF animasyonları ve düşük renkli görseller** için faydalıdır.

---

İstersen sana bir sonraki adımda **Go’da renk manipülasyon tekniklerini** (ör: parlaklık değiştirme, ters çevirme, sepya filtresi, vb.) anlatabilirim. Böylece `color` paketini gerçek dünyada nasıl uygulayacağını öğrenmiş olursun. Anlatayım mı?
EVET
*/
/*
Süper 👍 Artık `color` ve `palette` paketlerini öğrendik. Şimdi işin eğlenceli tarafına geçelim:
Go’da **renk manipülasyonu teknikleri**.

Bunlar sayesinde görseller üzerinde **filtreler** uygulayabilir, renklerle oynayabiliriz.

---

# 📌 Genel Mantık

1. Görseli aç (`image.Decode`) veya oluştur (`image.NewRGBA`).
2. Her pikseli `At(x, y)` ile oku.
3. `color.Color` → `RGBA()` ile renk kanallarını al.
4. İstediğin renk dönüşümünü uygula.
5. Yeni görselde `Set(x, y, newColor)` ile yaz.

---

# 🎨 Örnek Filtreler

## 1. **Ters Çevirme (Negative)**

Her kanalı `255 - değer` yaparsak resim negatife dönüşür.
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
    file, _ := os.Open("input.jpg")
    defer file.Close()
    img, _, _ := image.Decode(file)

    bounds := img.Bounds()
    newImg := image.NewRGBA(bounds)

    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            r, g, b, a := img.At(x, y).RGBA()
            newColor := color.RGBA{
                R: uint8(255 - r/257),
                G: uint8(255 - g/257),
                B: uint8(255 - b/257),
                A: uint8(a / 257),
            }
            newImg.Set(x, y, newColor)
        }
    }

    out, _ := os.Create("negative.jpg")
    defer out.Close()
    jpeg.Encode(out, newImg, nil)
}
``
/*
👉 Çalıştırınca resmin negatif versiyonu çıkar.

---

## 2. **Gri Tonlama (Grayscale)**

Her piksel için ortalama (`(R+G+B)/3`) veya ağırlıklı ortalama kullanılır.
*/
``go
grayVal := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
``

//Örnek:

``go
newColor := color.Gray{Y: uint8(grayVal / 257)}
``
/*
---

## 3. **Parlaklık Ayarı**

Her pikselin `R, G, B` değerine `+delta` ekle.
*/
``go
nr := clamp(int(r/257)+30, 0, 255)
ng := clamp(int(g/257)+30, 0, 255)
nb := clamp(int(b/257)+30, 0, 255)
``

//👉 `clamp` fonksiyonu:

``go
func clamp(v, min, max int) uint8 {
    if v < min {
        return uint8(min)
    }
    if v > max {
        return uint8(max)
    }
    return uint8(v)
}
``
/*
---

## 4. **Sepya Filtresi**

Klasik fotoğraf efekti.
Formül:
*/
``go
tr := 0.393*r + 0.769*g + 0.189*b
tg := 0.349*r + 0.686*g + 0.168*b
tb := 0.272*r + 0.534*g + 0.131*b
``
/*
Sonra `clamp` ile 0–255 aralığına alırsın.

---

## 5. **Siyah-Beyaz (Threshold)**

Bir eşik değeri (`128` gibi) belirle. Ortalama > 128 ise beyaz, aksi halde siyah yap.
*/
``go
avg := (r + g + b) / 3
if avg > 128 {
    newColor = color.RGBA{255, 255, 255, 255}
} else {
    newColor = color.RGBA{0, 0, 0, 255}
}
``
/*
---

# 📌 Daha Gelişmiş Yöntemler

* **Kontrast artırma/azaltma**
* **Renk doygunluğu (saturation)** değiştirme
* **Kanal karıştırma (RGB → farklı tonlama)**
* **Blur, sharpen** gibi efektler (`golang.org/x/image/draw` veya `fogleman/gg` gibi kütüphanelerle yapılır).

---

# 📌 Özet

* Görselleri pikseller bazında işleyerek filtre uygulayabilirsin.
* `RGBA()` ile renk değerlerini al → matematiksel dönüşüm uygula → `Set` ile yeni görsele yaz.
* Basit filtreler: **negative, grayscale, brightness, sepia, threshold**.
* Daha karmaşık efektler için ek kütüphaneler (örn. `golang.org/x/image`) kullanılabilir.

---

İstersen sana bu filtreleri tek bir programda menü şeklinde sunup, kullanıcı seçimine göre uygulayan **mini bir resim editörü** örneği yazabilirim. İster misin?
EVET
*/
/*
Harika! 🎉 O zaman Go’da tamamen **kendi mini resim editörümüzü** yapalım.
Program, bir görseli açacak, kullanıcıya **filtre seçenekleri** sunacak ve seçilen filtreyi uygulayıp sonucu kaydedecek.

---

# 📌 Mini Resim Editörü Örneği
*/
``go
package main

import (
    "fmt"
    "image"
    "image/color"
    "image/jpeg"
    "image/png"
    "os"
    "strings"
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

// Sepya filtre
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

// Threshold (siyah-beyaz)
func threshold(img image.Image, t uint8) *image.RGBA {
    bounds := img.Bounds()
    newImg := image.NewRGBA(bounds)
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            r, g, b, a := img.At(x, y).RGBA()
            avg := uint8((r/257 + g/257 + b/257) / 3)
            var val uint8
            if avg > t {
                val = 255
            } else {
                val = 0
            }
            newImg.Set(x, y, color.RGBA{val, val, val, uint8(a / 257)})
        }
    }
    return newImg
}

func main() {
    var inputPath, outputPath, choice string
    fmt.Print("Resim dosya yolu (örn: input.jpg): ")
    fmt.Scanln(&inputPath)

    // Görseli aç
    file, err := os.Open(inputPath)
    if err != nil {
        fmt.Println("Dosya açılamadı:", err)
        return
    }
    defer file.Close()

    img, format, err := image.Decode(file)
    if err != nil {
        fmt.Println("Görsel okunamadı:", err)
        return
    }
    fmt.Println("Görsel formatı:", format)

    // Filtre seçenekleri
    fmt.Println("Filtre seçenekleri:")
    fmt.Println("1 - Negative (Negatif)")
    fmt.Println("2 - Grayscale (Gri tonlama)")
    fmt.Println("3 - Sepia")
    fmt.Println("4 - Threshold (Siyah-Beyaz)")
    fmt.Print("Seçiminiz (1-4): ")
    fmt.Scanln(&choice)

    var newImg *image.RGBA
    switch choice {
    case "1":
        newImg = negative(img)
    case "2":
        newImg = grayscale(img)
    case "3":
        newImg = sepia(img)
    case "4":
        newImg = threshold(img, 128)
    default:
        fmt.Println("Geçersiz seçim!")
        return
    }

    fmt.Print("Kaydedilecek dosya adı (örn: output.png): ")
    fmt.Scanln(&outputPath)

    // Dosyayı uzantısına göre kaydet
    f, err := os.Create(outputPath)
    if err != nil {
        fmt.Println("Dosya oluşturulamadı:", err)
        return
    }
    defer f.Close()

    ext := strings.ToLower(outputPath[strings.LastIndex(outputPath, ".")+1:])
    if ext == "jpg" || ext == "jpeg" {
        jpeg.Encode(f, newImg, &jpeg.Options{Quality: 90})
    } else {
        png.Encode(f, newImg)
    }

    fmt.Println("Filtre uygulandı ve kaydedildi:", outputPath)
}
``
/*
---

# 📌 Kullanım

1. Programı çalıştır (`go run main.go`).
2. Açmak istediğin resmin yolunu gir (`input.jpg`).
3. Filtre seç (`1-4`).
4. Kaydedilecek dosya adını gir (`output.png` veya `output.jpg`).
5. İşlem tamam! 🎉

---

# 📌 Özellikler

* Negative, Grayscale, Sepia, Threshold filtreleri hazır.
* Hem PNG hem JPEG kaydı destekli.
* `color.RGBA` ve `RGBA()` mantığını tüm filtrelerde kullanıyor.
* İstediğin zaman **yeni filtreler eklemek kolay** (ör. parlaklık, kontrast).

---

İstersen bir sonraki adımda **bu programı geliştirip bir “menü tabanlı terminal mini editör” + animasyon desteği** de ekleyebiliriz. Böylece GIF animasyonlarına filtre uygulamak da mümkün olur.

Bunu yapalım mı?
EVET
*/
/*
Harika! 🎉 O zaman bir sonraki adımda **mini terminal tabanlı resim editörünü geliştirelim ve GIF animasyon desteği ekleyelim**.

Böylece hem tek kare görselleri hem de **animasyonlu GIF’leri** filtreleyebileceğiz.

---

# 📌 Geliştirilmiş Mini Resim Editörü + GIF Desteği

Özellikler:

1. Tek kare (PNG/JPEG) ve çok kareli GIF’leri açabilir.
2. Mevcut filtreleri (Negative, Grayscale, Sepia, Threshold) uygular.
3. Yeni GIF dosyası oluşturur ve filtre uygulanmış kareleri kaydeder.

---
*/
``go
package main

import (
    "fmt"
    "image"
    "image/color"
    "image/gif"
    "image/jpeg"
    "image/png"
    "os"
    "strings"
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

// Sepya filtre
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

// Threshold (siyah-beyaz)
func threshold(img image.Image, t uint8) *image.RGBA {
    bounds := img.Bounds()
    newImg := image.NewRGBA(bounds)
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            r, g, b, a := img.At(x, y).RGBA()
            avg := uint8((r/257 + g/257 + b/257) / 3)
            var val uint8
            if avg > t {
                val = 255
            } else {
                val = 0
            }
            newImg.Set(x, y, color.RGBA{val, val, val, uint8(a / 257)})
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
    case "4":
        return threshold(img, 128)
    default:
        return nil
    }
}

func main() {
    var inputPath, outputPath, choice string
    fmt.Print("Resim dosya yolu (örn: input.jpg veya anim.gif): ")
    fmt.Scanln(&inputPath)

    file, err := os.Open(inputPath)
    if err != nil {
        fmt.Println("Dosya açılamadı:", err)
        return
    }
    defer file.Close()

    // Format kontrol
    img, format, err := image.Decode(file)
    if err != nil {
        fmt.Println("Görsel okunamadı:", err)
        return
    }
    fmt.Println("Görsel formatı:", format)

    // Filtre seçenekleri
    fmt.Println("Filtre seçenekleri:")
    fmt.Println("1 - Negative (Negatif)")
    fmt.Println("2 - Grayscale (Gri tonlama)")
    fmt.Println("3 - Sepia")
    fmt.Println("4 - Threshold (Siyah-Beyaz)")
    fmt.Print("Seçiminiz (1-4): ")
    fmt.Scanln(&choice)

    fmt.Print("Kaydedilecek dosya adı (örn: output.png veya output.gif): ")
    fmt.Scanln(&outputPath)

    ext := strings.ToLower(outputPath[strings.LastIndex(outputPath, ".")+1:])

    if ext == "gif" {
        // GIF animasyon işleme
        file.Seek(0, 0) // dosyayı başa al
        g, err := gif.DecodeAll(file)
        if err != nil {
            fmt.Println("GIF okunamadı:", err)
            return
        }

        // Filtreyi her kareye uygula
        for i, frame := range g.Image {
            rgba := applyFilter(frame, choice)
            if rgba != nil {
                newPaletted := image.NewPaletted(rgba.Bounds(), frame.Palette)
                for y := 0; y < rgba.Bounds().Dy(); y++ {
                    for x := 0; x < rgba.Bounds().Dx(); x++ {
                        newPaletted.Set(x, y, rgba.At(x, y))
                    }
                }
                g.Image[i] = newPaletted
            }
        }

        out, _ := os.Create(outputPath)
        defer out.Close()
        gif.EncodeAll(out, g)
    } else {
        // Tek kare (PNG/JPEG)
        newImg := applyFilter(img, choice)
        if newImg == nil {
            fmt.Println("Filtre uygulanamadı!")
            return
        }

        out, _ := os.Create(outputPath)
        defer out.Close()

        if ext == "jpg" || ext == "jpeg" {
            jpeg.Encode(out, newImg, &jpeg.Options{Quality: 90})
        } else {
            png.Encode(out, newImg)
        }
    }

    fmt.Println("Filtre uygulandı ve kaydedildi:", outputPath)
}
``
/*
---

# 📌 Özellikler

1. **Tek kare resimler (PNG/JPEG)** ve **GIF animasyonlar** destekleniyor.
2. GIF’de her kare ayrı ayrı filtreleniyor.
3. Aynı filtreler: Negative, Grayscale, Sepia, Threshold.
4. Çıktıyı kullanıcı seçebiliyor: PNG, JPEG veya GIF.

---

İstersen bir sonraki adımda **bu mini editörü geliştirebilir ve terminalde gerçek zamanlı menü + renk önizleme + daha fazla filtreler** ekleyebiliriz.
Bunu da yapalım mı?
*/


