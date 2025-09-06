/*
Go’da **`draw` paketi** özellikle **görseller üzerinde çizim ve kopyalama işlemleri** için kullanılır. Bunu baştan sona örneklerle açıklayalım.

---

# 📌 `image/draw` Paketi Nedir?

* `image/draw` paketi, Go’nun standart kütüphanesinde yer alır.
* Temel amaç: **bir görüntüden diğerine çizim ve kopyalama yapmak**.
* `image` paketinde oluşturduğun görselleri (RGBA, Paletted vb.) manipüle etmek için kullanılır.
* Sık kullanılan fonksiyonlar:

1. **`Draw(dst Image, r Rectangle, src Image, sp Point, op Op)`**

   * `dst` → Hedef görsel
   * `r` → Hedef dikdörtgen
   * `src` → Kaynak görsel
   * `sp` → Kaynak görselin başlangıç noktası
   * `op` → İşlem türü (`Over`, `Src`)

2. **`Copy(dst Image, r Rectangle, src Image, sp Point)`**

   * `Draw`’ın basitleştirilmiş versiyonu
   * Kaynak görseli hedefe birebir kopyalar

---

# 📌 Önemli Tipler

### 1. `draw.Image`

* `image.Image` arayüzünü genişletir.
* `Set` metodu ile pikselleri değiştirebilir.

### 2. `draw.Op`

* Çizim işleminin tipini belirtir:

  * `draw.Over` → Kaynak, hedefin üzerine çizilir (alfa kanalı ile birleşir)
  * `draw.Src` → Kaynak, hedefi tamamen değiştirir

---

# 📌 Örnekler

## 1. Basit Kopyalama
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
    // Kaynak görsel (kırmızı)
    src := image.NewRGBA(image.Rect(0, 0, 100, 100))
    draw.Draw(src, src.Bounds(), &image.Uniform{color.RGBA{255, 0, 0, 255}}, image.Point{}, draw.Src)

    // Hedef görsel (beyaz)
    dst := image.NewRGBA(image.Rect(0, 0, 200, 200))
    draw.Draw(dst, dst.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

    // Kaynağı hedefe kopyala
    draw.Draw(dst, image.Rect(50, 50, 150, 150), src, image.Point{}, draw.Over)

    f, _ := os.Create("draw_example.png")
    defer f.Close()
    png.Encode(f, dst)
}
``
/*
👉 Çıktı: Beyaz arka plan üzerine 100x100 kırmızı kare (50,50) noktasında.

---

## 2. Alfa Kanallı Çizim
*/
``go
aColor := color.RGBA{0, 255, 0, 128} // yarı saydam yeşil
draw.Draw(dst, image.Rect(75, 75, 175, 175), &image.Uniform{aColor}, image.Point{}, draw.Over)
``
/*
* Kaynak yarı saydam ise, hedefin üzerine **blend** (karıştırma) yapılır.
* `draw.Over` bunu sağlar.
* `draw.Src` kullanılırsa hedef tamamen değiştirilir.

---

## 3. Bir Görseli Diğerine Yerleştirme

* Örneğin bir resmin sol üst kısmını başka bir resme kopyalamak:
*/
``go
src := image.NewRGBA(image.Rect(0, 0, 50, 50))
dst := image.NewRGBA(image.Rect(0, 0, 200, 200))

draw.Draw(dst, image.Rect(0, 0, 50, 50), src, image.Point{0, 0}, draw.Src)
``
/*
---

# 📌 Özet

1. **`image/draw`**, Go’da **görselleri birleştirme ve çizme** paketidir.
2. **Temel fonksiyon:** `draw.Draw(dst, r, src, sp, op)`

   * `op` ile alfa kanalı işlemlerini kontrol edersin (`Over` veya `Src`).
3. **Kullanım alanları:**

   * Filtre uygulama sonrası görüntüyü birleştirme
   * Animasyonlu GIF karelerini birleştirme
   * Maske, overlay veya watermark ekleme

---

İstersen bir sonraki adımda **`draw` paketini kullanarak iki görseli üst üste bindiren ve alfa kanalıyla karıştıran gerçek bir örnek program** da yapabiliriz.
Bunu gösterelim mi?
EVET


Harika! 🎉 O zaman şimdi Go’da **`draw` paketi kullanarak iki görseli üst üste bindirme ve alfa kanalı ile karıştırma** örneği yapalım.

---

# 📌 Örnek: Alfa Kanallı Görsel Bindirme
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
    // Hedef görsel: beyaz arka plan 200x200
    dst := image.NewRGBA(image.Rect(0, 0, 200, 200))
    draw.Draw(dst, dst.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

    // Kaynak görsel: yarı saydam kırmızı kare 100x100
    src := image.NewRGBA(image.Rect(0, 0, 100, 100))
    redAlpha := color.RGBA{255, 0, 0, 128} // alfa = 128 → yarı saydam
    draw.Draw(src, src.Bounds(), &image.Uniform{redAlpha}, image.Point{}, draw.Src)

    // Kaynağı hedefin üzerine bindir (50,50 konumuna)
    draw.Draw(dst, image.Rect(50, 50, 150, 150), src, image.Point{0, 0}, draw.Over)

    // Bir başka kaynak: yarı saydam yeşil kare 100x100
    src2 := image.NewRGBA(image.Rect(0, 0, 100, 100))
    greenAlpha := color.RGBA{0, 255, 0, 128} // yarı saydam
    draw.Draw(src2, src2.Bounds(), &image.Uniform{greenAlpha}, image.Point{}, draw.Src)

    // Bu kareyi üstüne bindir (80,80 konumuna)
    draw.Draw(dst, image.Rect(80, 80, 180, 180), src2, image.Point{0, 0}, draw.Over)

    // Sonucu kaydet
    f, _ := os.Create("draw_overlay.png")
    defer f.Close()
    png.Encode(f, dst)
}
``
/*
---

# 📌 Açıklama

1. **`dst`** → Beyaz arka plan, üzerine çizim yapılacak.
2. **`src` ve `src2`** → Yarı saydam renkli kareler.
3. **`draw.Over`** → Alfa kanalı ile **blend** işlemi yapar.
4. Karelerin konumu `image.Rect(x0, y0, x1, y1)` ile belirlenir.
5. Sonuç: **iki renkli kare üst üste geldiğinde, alfa kanalı sayesinde altındaki kare görünür**.

---

# 📌 Özet

* `draw.Draw` ile görselleri **istediğin konuma bindirebilir** ve **alfa kanalıyla karıştırabilirsin**.
* `draw.Src` → Hedefi tamamen değiştirir (blend yok).
* `draw.Over` → Kaynak, hedefin üzerine alfa ile bindirilir.
* Bu teknik **overlay, watermark, filtre sonrası birleştirme** gibi işlemlerde çok kullanışlıdır.

---

İstersen bir sonraki adımda bu tekniği kullanarak **animasyonlu GIF’lere overlay (ör: watermark veya logo) ekleyen bir program** da gösterebilirim.
Bunu da yapalım mı?
EVET
*/
/*
Harika! 🎉 O zaman Go’da **animasyonlu GIF’lere overlay (örneğin logo veya watermark) ekleme** örneğini yapalım.

Senaryo:

* Mevcut bir GIF animasyonu var.
* Her kareye **küçük bir logo veya renkli kare** bindiriyoruz.
* Sonuçta **overlay eklenmiş yeni GIF** oluşturuyoruz.

---

# 📌 Örnek: GIF’e Overlay Ekleme
*/
``go
package main

import (
    "image"
    "image/color"
    "image/draw"
    "image/gif"
    "image/png"
    "os"
)

func main() {
    // 1️⃣ Mevcut GIF'i aç
    f, err := os.Open("input_anim.gif")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    g, err := gif.DecodeAll(f)
    if err != nil {
        panic(err)
    }

    // 2️⃣ Overlay görseli (küçük renkli kare / logo)
    overlay := image.NewRGBA(image.Rect(0, 0, 30, 30))
    draw.Draw(overlay, overlay.Bounds(), &image.Uniform{color.RGBA{255, 0, 0, 128}}, image.Point{}, draw.Src)

    // 3️⃣ Her kareye overlay uygula
    for i, frame := range g.Image {
        rgba := image.NewRGBA(frame.Bounds())
        draw.Draw(rgba, rgba.Bounds(), frame, image.Point{}, draw.Src)

        // Overlay'yi bindir (ör: sağ alt köşeye)
        offset := image.Pt(frame.Bounds().Dx()-30, frame.Bounds().Dy()-30)
        draw.Draw(rgba, overlay.Bounds().Add(offset), overlay, image.Point{}, draw.Over)

        // Paletli kareye dönüştür
        newFrame := image.NewPaletted(rgba.Bounds(), frame.Palette)
        draw.FloydSteinberg.Draw(newFrame, rgba.Bounds(), rgba, image.Point{})
        g.Image[i] = newFrame
    }

    // 4️⃣ Yeni GIF'i kaydet
    out, err := os.Create("overlay_anim.gif")
    if err != nil {
        panic(err)
    }
    defer out.Close()

    gif.EncodeAll(out, g)
    println("Overlay uygulanmış GIF kaydedildi: overlay_anim.gif")
}
``
/*
---

# 📌 Açıklama

1. **`gif.DecodeAll`** → Mevcut GIF’in tüm karelerini okur.
2. **Overlay görseli** → Yarı saydam kırmızı kare oluşturduk (`color.RGBA{255,0,0,128}`).
3. **`draw.Over`** → Alfa kanalı ile karelerin üzerine bindirilir.
4. **`draw.FloydSteinberg.Draw`** → RGBA kareyi paletli kareye dönüştürür (GIF için gerekli).
5. **Kaydetme** → `gif.EncodeAll` ile yeni GIF oluşturulur.

---

# 📌 Sonuç

* Tüm karelere logo veya watermark bindirilir.
* Alfa kanalı sayesinde **orijinal animasyon görünümü korunur**.
* Bu yöntemle:

  * Filtre + overlay kombinasyonu yapılabilir.
  * GIF animasyonlar üzerinde profesyonel efektler eklenebilir.

---

İstersen bir sonraki adımda **filtre + overlay kombinasyonunu tek programda birleştirip, kullanıcı menüsü ile interaktif GIF editörü** haline getirebiliriz.
Bunu da yapalım mı?
*/