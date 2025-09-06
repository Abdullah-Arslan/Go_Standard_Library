/*
Go (Golang) dilindeki `fmt` paketi, **formatlı Girdi/Çıktı (I/O)** işlemleri için kullanılan temel bir standart kütüphanedir. İsmi "format" kelimesinden gelir ve hem konsola yazı yazmak (output) hem de kullanıcıdan giriş almak (input) için kullanılır.

---

## 📦 `fmt` Paketi Nedir?

`fmt`, metinleri belirli bir biçimde yazdırmak, biçimlendirmek veya okumak için kullanılır. Özellikle:

* `Print`, `Printf`, `Println` gibi fonksiyonlarla yazdırma,
* `Scan`, `Scanf`, `Scanln` gibi fonksiyonlarla kullanıcıdan veri alma,
* `Sprintf`, `Sprintln` gibi string oluşturma fonksiyonları içerir.

---

## 🔧 Paketin Sunduğu Temel Fonksiyonlar

### 1. **Yazdırma Fonksiyonları**

#### ✅ `fmt.Print()`

Argümanları yazdırır, **boşluk koymaz**.
*/
``go
fmt.Print("Merhaba")
fmt.Print("Dünya")
// Çıktı: MerhabaDünya
``
/*
#### ✅ `fmt.Println()`

Argümanları yazdırır, **boşluk koyar** ve **sonunda yeni satır ekler**.
*/
``go
fmt.Println("Merhaba", "Dünya")
// Çıktı: Merhaba Dünya
``
/*
#### ✅ `fmt.Printf()`

C formatlama sözdizimine benzer, **biçim belirleyerek yazdırır**.
*/
``go
isim := "Ali"
yas := 30
fmt.Printf("İsim: %s, Yaş: %d\n", isim, yas)
// Çıktı: İsim: Ali, Yaş: 30
``
/*
---

### 2. **Biçimlendirilmiş Yazdırma (String Üretme)**

#### ✅ `fmt.Sprintf()`

Formatlı bir yazıyı **string olarak döner**, yazdırmaz.
*/
``go
str := fmt.Sprintf("Fiyat: %.2f TL", 99.99)
fmt.Println(str)
// Çıktı: Fiyat: 99.99 TL
``
/*
#### ✅ `fmt.Sprintln()`

Println gibi ama sonucu string olarak döner.
*/
``go
s := fmt.Sprintln("Merhaba", "Dünya")
fmt.Print(s)
// Çıktı: Merhaba Dünya\n
``
/*
#### ✅ `fmt.Sprint()`

Print gibi ama sonucu string olarak döner.
*/
``go
s := fmt.Sprint("Ali", "Veli")
fmt.Println(s)
// Çıktı: AliVeli
``
/*
---

### 3. **Girdi Alma (Konsoldan Okuma)**

#### ✅ `fmt.Scan()`

Boşluklarla ayrılmış değerleri okur.
*/
``go
var isim string
var yas int
fmt.Print("İsim ve yaş girin: ")
fmt.Scan(&isim, &yas)
fmt.Println("İsim:", isim, "Yaş:", yas)
``
/*
#### ✅ `fmt.Scanln()`

Bir satırdan giriş alır, boşluklarla ayırır.
*/
``go
var ad string
fmt.Print("Adınızı girin: ")
fmt.Scanln(&ad)
fmt.Println("Merhaba", ad)
``
/*
#### ✅ `fmt.Scanf()`

Formatlı giriş alır.
*/
``go
var isim string
var yas int
fmt.Print("İsim ve yaş girin (Ali 30 gibi): ")
fmt.Scanf("%s %d", &isim, &yas)
fmt.Printf("İsim: %s, Yaş: %d\n", isim, yas)
``
/*
---

## 📌 Formatlama Belirteçleri (`Printf`, `Sprintf` vb.)

| Belirteç | Anlamı                     | Örnek                |
| -------- | -------------------------- | -------------------- |
| `%s`     | string                     | `"Merhaba"`          |
| `%d`     | integer (decimal)          | `42`                 |
| `%f`     | float                      | `3.14`               |
| `%.2f`   | virgülden sonra 2 basamak  | `3.14`               |
| `%t`     | boolean                    | `true`               |
| `%v`     | otomatik tür biçimi        | her şey için         |
| `%T`     | değişkenin tipi            | `int`, `string`, vb. |
| `%q`     | string’i çift tırnakla yaz | `"metin"`            |

---

## 🧪 Örnek: Tüm Özellikleri Gösteren Küçük Program
*/
``go
package main

import (
	"fmt"
)

func main() {
	// Print/Println
	fmt.Print("Merhaba")
	fmt.Println(" Dünya!")

	// Printf
	name := "Ayşe"
	age := 25
	fmt.Printf("İsim: %s, Yaş: %d\n", name, age)

	// Sprintf
	info := fmt.Sprintf("Kullanıcı: %s (%d yaşında)", name, age)
	fmt.Println(info)

	// Giriş alma
	var isim string
	var yas int
	fmt.Print("Adınızı ve yaşınızı girin: ")
	fmt.Scan(&isim, &yas)
	fmt.Printf("Hoşgeldiniz %s, %d yaşındasınız.\n", isim, yas)
}
``
/*
---

## 🔚 Sonuç

`fmt` paketi:

* Go'nun temel I/O aracıdır.
* Yazdırma, string formatlama ve kullanıcıdan veri alma işlerinde kullanılır.
* C'deki `printf`, `scanf` mantığına oldukça benzer.
* Format belirteçleriyle birlikte esnek ve güçlü bir kullanım sunar.

---

İstersen bu bilgileri küçük bir PDF el kitabına da dönüştürebilirim. Ya da tek tek uygulamalı örneklerle bir terminal uygulaması yazabiliriz. Yardımcı olayım mı?
EVET
*/
/*
Harika, evet! Şimdi seninle birlikte Go dilinde `fmt` paketini kullanarak terminalde çalışan küçük bir **etkileşimli uygulama** yazalım. Bu uygulamada:

* Kullanıcıdan isim ve yaş alacağız (`fmt.Scanln`)
* Formatlı bir karşılama mesajı vereceğiz (`fmt.Printf`)
* Kullanıcıdan maaş bilgisini alacağız (`fmt.Scanf`)
* Maaşı %15 zamla göstereceğiz (`fmt.Sprintf`)
* Tüm bilgileri özet olarak yazdıracağız

---

## 👨‍💻 1. Tam Çalışan Uygulama Kodu
*/
``go
package main

import (
	"fmt"
)

func main() {
	var isim string
	var yas int

	fmt.Println("🟢 Basit Kişisel Bilgi Uygulamasına Hoş Geldiniz!")
	fmt.Print("Lütfen adınızı girin: ")
	fmt.Scanln(&isim)

	fmt.Print("Lütfen yaşınızı girin: ")
	fmt.Scanln(&yas)

	fmt.Printf("✅ Merhaba %s, %d yaşındasınız.\n\n", isim, yas)

	// Maaş bilgisi
	var maas float64
	fmt.Print("Lütfen aylık maaşınızı girin (örn. 15000.50): ")
	fmt.Scanf("%f\n", &maas)

	// %15 zam hesapla
	zamliMaas := maas * 1.15
	zamliStr := fmt.Sprintf("💰 Zamlı maaşınız: %.2f TL", zamliMaas)
	fmt.Println(zamliStr)

	// Özet Bilgi
	fmt.Println("\n🔍 Bilgi Özeti:")
	fmt.Printf("İsim: %s\n", isim)
	fmt.Printf("Yaş: %d\n", yas)
	fmt.Printf("Eski Maaş: %.2f TL\n", maas)
	fmt.Printf("Yeni Maaş: %.2f TL\n", zamliMaas)
}
``
/*
---

## 🚀 Nasıl Derlenir ve Çalıştırılır?

1. Yukarıdaki kodu bir dosyaya yapıştır: `main.go`
2. Terminalde bu klasöre gel:
*/
   ``bash
   cd /dosya/yolu/
   ``
//3. Çalıştırmak için:

   ``bash
   go run main.go
   ``
/*
---

## 🧪 Örnek Çalışma Ekranı
*/
``
🟢 Basit Kişisel Bilgi Uygulamasına Hoş Geldiniz!
Lütfen adınızı girin: Ali
Lütfen yaşınızı girin: 30
✅ Merhaba Ali, 30 yaşındasınız.

Lütfen aylık maaşınızı girin (örn. 15000.50): 20000
💰 Zamlı maaşınız: 23000.00 TL

🔍 Bilgi Özeti:
İsim: Ali
Yaş: 30
Eski Maaş: 20000.00 TL
Yeni Maaş: 23000.00 TL
``
/*
---

## 🎯 Öğrenilen `fmt` Fonksiyonları

| Fonksiyon     | Açıklama                               |
| ------------- | -------------------------------------- |
| `fmt.Print`   | Yazı yazar, satır atlamaz              |
| `fmt.Println` | Yazı yazar, satır atlar                |
| `fmt.Printf`  | Formatlı çıktı sağlar                  |
| `fmt.Scanln`  | Kullanıcıdan giriş alır (satır bazlı)  |
| `fmt.Scanf`   | Formatlı giriş alır                    |
| `fmt.Sprintf` | Formatlı yazıyı string olarak döndürür |

---

İstersen bu uygulamayı genişletip:

* Kullanıcıya bir menü sunabiliriz
* Maaş zam oranını kullanıcı belirleyebilir
* Kullanıcı bilgilerini dosyaya yazabiliriz

Hangisini yapmak istersin?
Kullanıcıya bir menü sunabiliriz

Süper! Şimdi uygulamamıza bir **terminal tabanlı menü** ekleyelim. Bu menü sayesinde kullanıcı birden fazla işlem yapabilecek.

---

## 🎯 Hedef

Kullanıcıya şu menüyü sunacağız:
*/
``
📋 Menü:
1. Kişisel Bilgileri Gir
2. Maaş Hesapla (%15 zam)
3. Bilgileri Göster
4. Çıkış
``
/*
---

## 👨‍💻 Tam Çalışan Menü Tabanlı Uygulama (Go + fmt)
*/
``go
package main

import (
	"fmt"
)

func main() {
	var isim string
	var yas int
	var maas float64
	var bilgiGirilmis bool = false

	for {
		// Menü
		fmt.Println("\n📋 Menü:")
		fmt.Println("1. Kişisel Bilgileri Gir")
		fmt.Println("2. Maaş Hesapla (%15 zam)")
		fmt.Println("3. Bilgileri Göster")
		fmt.Println("4. Çıkış")
		fmt.Print("Bir seçenek girin (1-4): ")

		var secim int
		fmt.Scanln(&secim)

		switch secim {
		case 1:
			fmt.Print("İsminizi girin: ")
			fmt.Scanln(&isim)

			fmt.Print("Yaşınızı girin: ")
			fmt.Scanln(&yas)

			fmt.Print("Maaşınızı girin: ")
			fmt.Scanln(&maas)

			bilgiGirilmis = true
			fmt.Println("✅ Bilgiler kaydedildi.")
		case 2:
			if !bilgiGirilmis {
				fmt.Println("⚠️ Önce kişisel bilgileri girmeniz gerekiyor.")
				continue
			}
			zamli := maas * 1.15
			fmt.Printf("💰 Zamlı maaş: %.2f TL (%%15 zam)\n", zamli)
		case 3:
			if !bilgiGirilmis {
				fmt.Println("⚠️ Henüz bilgi girilmedi.")
				continue
			}
			fmt.Println("🔍 Bilgi Özeti:")
			fmt.Printf("İsim: %s\n", isim)
			fmt.Printf("Yaş: %d\n", yas)
			fmt.Printf("Maaş: %.2f TL\n", maas)
		case 4:
			fmt.Println("👋 Programdan çıkılıyor. Görüşmek üzere!")
			return
		default:
			fmt.Println("❌ Geçersiz seçenek. Lütfen 1-4 arasında bir değer girin.")
		}
	}
}
``
/*
---

## 🧪 Örnek Terminal Kullanımı
*/
``
📋 Menü:
1. Kişisel Bilgileri Gir
2. Maaş Hesapla (%15 zam)
3. Bilgileri Göster
4. Çıkış
Bir seçenek girin (1-4): 2
⚠️ Önce kişisel bilgileri girmeniz gerekiyor.

📋 Menü:
1. Kişisel Bilgileri Gir
2. Maaş Hesapla (%15 zam)
3. Bilgileri Göster
4. Çıkış
Bir seçenek girin (1-4): 1
İsminizi girin: Zeynep
Yaşınızı girin: 28
Maaşınızı girin: 18000
✅ Bilgiler kaydedildi.

📋 Menü:
1. Kişisel Bilgileri Gir
2. Maaş Hesapla (%15 zam)
3. Bilgileri Göster
4. Çıkış
Bir seçenek girin (1-4): 2
💰 Zamlı maaş: 20700.00 TL (%15 zam)
``
/*
---

## 🔍 Kod Açıklaması

| Özellik                 | Açıklama                                    |
| ----------------------- | ------------------------------------------- |
| `for {}`                | Sonsuz döngü (kullanıcı çıkana kadar)       |
| `switch secim`          | Menü seçimlerini yöneten yapı               |
| `fmt.Scanln(&degisken)` | Kullanıcıdan input alma                     |
| `bilgiGirilmis`         | Giriş yapıldı mı kontrolü için boolean flag |
| `return`                | Programdan çıkmak için                      |

---

## ➕ Geliştirme Fikirleri

İstersen bu menüyü daha da geliştirebiliriz:

* ✅ Zam oranını kullanıcı belirlesin (örneğin %15 değil, istediği oran)
* 💾 Kullanıcı bilgilerini dosyaya kaydedelim (JSON/TXT)
* 🔄 Birden fazla kullanıcı eklenebilsin (dizi/map ile)
* 🧮 Ek vergi hesaplaması, net maaş gibi finansal işlemler

---

Yukarıdaki fikirlerden birini yapmak ister misin? Hangi yönde ilerleyelim?
Harika tercih! Şimdi Go dilinde kullanıcı bilgilerini bir dosyaya kaydedeceğimiz bir sistem oluşturacağız. Bu işlem için şu iki adımı takip edeceğiz:

---

## 🛠️ Hedefimiz

* Kullanıcı bilgilerini `struct` içinde tutmak
* JSON formatında `.json` dosyasına yazmak ve dosyadan tekrar okuyabilmek
* Menüye şu seçenekleri eklemek:

  * `5. Bilgileri Kaydet (JSON)`
  * `6. Bilgileri Yükle (JSON)`

---

## ✅ Gerekli Kütüphaneler

* `encoding/json`: JSON'a çevirme ve JSON'dan okuma
* `os`: Dosya açma/yazma işlemleri

---

## 👨‍💻 Tam Menü + JSON Kaydet/Yükle Uygulaması
*/
``go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Kullanıcıyı tanımlayan struct
type Kullanici struct {
	Isim string  `json:"isim"`
	Yas  int     `json:"yas"`
	Maas float64 `json:"maas"`
}

func main() {
	var kullanici Kullanici
	var bilgiGirilmis bool = false

	for {
		fmt.Println("\n📋 Menü:")
		fmt.Println("1. Kişisel Bilgileri Gir")
		fmt.Println("2. Maaş Hesapla (%15 zam)")
		fmt.Println("3. Bilgileri Göster")
		fmt.Println("4. Çıkış")
		fmt.Println("5. Bilgileri Kaydet (JSON)")
		fmt.Println("6. Bilgileri Yükle (JSON)")
		fmt.Print("Bir seçenek girin (1-6): ")

		var secim int
		fmt.Scanln(&secim)

		switch secim {
		case 1:
			fmt.Print("İsminizi girin: ")
			fmt.Scanln(&kullanici.Isim)

			fmt.Print("Yaşınızı girin: ")
			fmt.Scanln(&kullanici.Yas)

			fmt.Print("Maaşınızı girin: ")
			fmt.Scanln(&kullanici.Maas)

			bilgiGirilmis = true
			fmt.Println("✅ Bilgiler kaydedildi.")
		case 2:
			if !bilgiGirilmis {
				fmt.Println("⚠️ Önce kişisel bilgileri girmeniz gerekiyor.")
				continue
			}
			zamli := kullanici.Maas * 1.15
			fmt.Printf("💰 Zamlı maaş: %.2f TL (%%15 zam)\n", zamli)
		case 3:
			if !bilgiGirilmis {
				fmt.Println("⚠️ Henüz bilgi girilmedi.")
				continue
			}
			fmt.Println("🔍 Bilgi Özeti:")
			fmt.Printf("İsim: %s\n", kullanici.Isim)
			fmt.Printf("Yaş: %d\n", kullanici.Yas)
			fmt.Printf("Maaş: %.2f TL\n", kullanici.Maas)
		case 4:
			fmt.Println("👋 Programdan çıkılıyor. Görüşmek üzere!")
			return
		case 5:
			if !bilgiGirilmis {
				fmt.Println("⚠️ Kaydetmek için önce bilgi girin.")
				continue
			}
			err := bilgileriKaydetJSON(kullanici)
			if err != nil {
				fmt.Println("❌ Kaydetme başarısız:", err)
			} else {
				fmt.Println("💾 Bilgiler 'kullanici.json' dosyasına kaydedildi.")
			}
		case 6:
			loaded, err := bilgileriYukleJSON()
			if err != nil {
				fmt.Println("❌ Yükleme başarısız:", err)
			} else {
				kullanici = loaded
				bilgiGirilmis = true
				fmt.Println("📂 Bilgiler başarıyla yüklendi.")
			}
		default:
			fmt.Println("❌ Geçersiz seçim. Lütfen 1-6 arasında bir değer girin.")
		}
	}
}

// JSON dosyasına yazma fonksiyonu
func bilgileriKaydetJSON(k Kullanici) error {
	file, err := os.Create("kullanici.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(k)
}

// JSON dosyasından okuma fonksiyonu
func bilgileriYukleJSON() (Kullanici, error) {
	file, err := os.Open("kullanici.json")
	if err != nil {
		return Kullanici{}, err
	}
	defer file.Close()

	var k Kullanici
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&k)
	return k, err
}
``
/*
---

## 🧪 Çalışma Örneği

1. **Bilgi gir**
2. **Kaydet**
3. Programı kapat
4. **Yeniden aç ve "Yükle" seçeneğini seç**
*/
``bash
go run main.go
``

``
📋 Menü:
1. Kişisel Bilgileri Gir
2. Maaş Hesapla (%15 zam)
3. Bilgileri Göster
4. Çıkış
5. Bilgileri Kaydet (JSON)
6. Bilgileri Yükle (JSON)
Bir seçenek girin (1-6): 5
💾 Bilgiler 'kullanici.json' dosyasına kaydedildi.
``
/*
Dosya sisteminde:

```
📁 bulunduğun klasör/
    └── kullanici.json
```

İçeriği:
*/
``json
{
  "isim": "Ali",
  "yas": 30,
  "maas": 20000
}
``
/*
---

## ✅ Neler Öğrendik?

* `struct` kullanımı
* JSON encode/decode (`json.NewEncoder`, `json.NewDecoder`)
* Dosya oluşturma/açma (`os.Create`, `os.Open`)
* `defer file.Close()` ile dosyayı kapatma

---

## 🔄 Sonraki Geliştirme Fikirleri

* 📂 Farklı kullanıcıları liste halinde kaydet (slice of struct)
* 📁 Her kullanıcı için ayrı dosya (örneğin: `ali.json`, `ayse.json`)
* 🧾 TXT formatında kaydetme (isteğe bağlı)

İstersen şimdi JSON yerine **düz TXT dosyasına** kayıt yapmayı da gösterebilirim. İlgini çeker mi?
Farklı kullanıcıları liste halinde kaydet (slice of struct)

Mükemmel! Şimdi `Go` dilinde birden fazla kullanıcıyı yönetmek için:

> 🔁 Kullanıcıları bir `slice` (dilim) içinde tutacağız, yani `[]Kullanici` yapısı kullanacağız.

Ve bu kullanıcı listesini:

* **JSON dosyasına kaydedeceğiz**
* **JSON dosyasından geri yükleyeceğiz**
* **Yeni kullanıcı ekleyeceğiz**
* **Tüm kullanıcıları listeleyeceğiz**
*/
/*
---

## 🧱 Temel Yapılar

### 🧍 Kullanıcı Tanımı (`struct`)
*/
``go
type Kullanici struct {
    Isim string  `json:"isim"`
    Yas  int     `json:"yas"`
    Maas float64 `json:"maas"`
}
``

//### 🧾 Kullanıcı Listesi

``go
var kullanicilar []Kullanici
``
/*
---

## 🧑‍💻 Tam Menü Tabanlı Kod (Çoklu Kullanıcı + JSON Kayıt)
*/
``go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Kullanıcıyı tanımlayan struct
type Kullanici struct {
	Isim string  `json:"isim"`
	Yas  int     `json:"yas"`
	Maas float64 `json:"maas"`
}

var kullanicilar []Kullanici

func main() {
	for {
		fmt.Println("\n📋 Menü:")
		fmt.Println("1. Yeni Kullanıcı Ekle")
		fmt.Println("2. Tüm Kullanıcıları Listele")
		fmt.Println("3. Kullanıcıları JSON Olarak Kaydet")
		fmt.Println("4. JSON Dosyasından Kullanıcıları Yükle")
		fmt.Println("5. Çıkış")
		fmt.Print("Bir seçenek girin (1-5): ")

		var secim int
		fmt.Scanln(&secim)

		switch secim {
		case 1:
			var k Kullanici
			fmt.Print("İsim: ")
			fmt.Scanln(&k.Isim)

			fmt.Print("Yaş: ")
			fmt.Scanln(&k.Yas)

			fmt.Print("Maaş: ")
			fmt.Scanln(&k.Maas)

			kullanicilar = append(kullanicilar, k)
			fmt.Println("✅ Kullanıcı eklendi.")
		case 2:
			if len(kullanicilar) == 0 {
				fmt.Println("⚠️ Henüz kullanıcı eklenmedi.")
			} else {
				fmt.Println("👥 Kullanıcı Listesi:")
				for i, k := range kullanicilar {
					fmt.Printf("%d. %s | Yaş: %d | Maaş: %.2f TL\n", i+1, k.Isim, k.Yas, k.Maas)
				}
			}
		case 3:
			err := jsonKaydet(kullanicilar)
			if err != nil {
				fmt.Println("❌ Kaydetme hatası:", err)
			} else {
				fmt.Println("💾 Kullanıcılar 'kullanicilar.json' dosyasına kaydedildi.")
			}
		case 4:
			yuklenen, err := jsonYukle()
			if err != nil {
				fmt.Println("❌ Yükleme hatası:", err)
			} else {
				kullanicilar = yuklenen
				fmt.Printf("📂 %d kullanıcı yüklendi.\n", len(kullanicilar))
			}
		case 5:
			fmt.Println("👋 Çıkış yapılıyor...")
			return
		default:
			fmt.Println("❌ Geçersiz seçim. Lütfen 1-5 arasında bir değer girin.")
		}
	}
}

// JSON'a kaydetme fonksiyonu
func jsonKaydet(liste []Kullanici) error {
	file, err := os.Create("kullanicilar.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(liste)
}

// JSON'dan okuma fonksiyonu
func jsonYukle() ([]Kullanici, error) {
	file, err := os.Open("kullanicilar.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var liste []Kullanici
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&liste)
	return liste, err
}
``
/*
---

## 🧪 Örnek Terminal Kullanımı
*/
``
📋 Menü:
1. Yeni Kullanıcı Ekle
2. Tüm Kullanıcıları Listele
3. Kullanıcıları JSON Olarak Kaydet
4. JSON Dosyasından Kullanıcıları Yükle
5. Çıkış
Bir seçenek girin (1-5): 1
İsim: Ali
Yaş: 30
Maaş: 15000
✅ Kullanıcı eklendi.

📋 Menü:
Bir seçenek girin (1-5): 2
👥 Kullanıcı Listesi:
1. Ali | Yaş: 30 | Maaş: 15000.00 TL
``
/*
---

## 📁 JSON Dosyası Örneği (`kullanicilar.json`)
*/
``json
[
  {
    "isim": "Ali",
    "yas": 30,
    "maas": 15000
  },
  {
    "isim": "Ayşe",
    "yas": 28,
    "maas": 18000
  }
]
``
/*
---

## 🧠 Öğrenilenler

| Özellik                    | Açıklama                       |
| -------------------------- | ------------------------------ |
| `[]Kullanici`              | Kullanıcı listesi (slice)      |
| `append(slice, eleman)`    | Slice'e yeni eleman ekleme     |
| `json.Encoder` / `Decoder` | JSON’a yazma ve JSON’dan okuma |
| `os.Create`, `os.Open`     | Dosya işlemleri                |

---

## 🔧 Ekstra Geliştirmeler (istersen yapabiliriz)

* 🗑️ Kullanıcı silme (index ile)
* 🔍 Belirli kullanıcıyı arama
* 🔢 Kullanıcıya ID atama
* 📄 Her kullanıcıyı ayrı dosyaya yazma (`ali.json`, `ayse.json` gibi)
* 🔐 Basit şifre korumalı giriş sistemi

---

İstersen bunlardan birini şimdi yapalım. Hangisi ilgini çeker?
* 🔐 Basit şifre korumalı giriş sistemi
*/
/*
Süper! Şimdi kullanıcılar için **basit bir şifre korumalı giriş sistemi** ekleyelim.

---

## 🔐 Hedefimiz

* Kullanıcılar:

  * **İsim**
  * **Yaş**
  * **Maaş**
  * **Şifre** (🆕 yeni alan)
* Uygulama açıldığında:

  * Kullanıcılar JSON dosyasından yüklenir
  * Kullanıcı isterse:

    * ✅ Yeni kullanıcı ekler (şifreyle)
    * 🔐 Var olan bir kullanıcı olarak giriş yapar (şifre doğrulamasıyla)
    * 👁️ Kendi bilgilerini görüntüler

---
*/
//## 🧱 Genişletilmiş Struct

```go
type Kullanici struct {
	Isim   string  `json:"isim"`
	Yas    int     `json:"yas"`
	Maas   float64 `json:"maas"`
	Sifre  string  `json:"sifre"`
}
```

---

## 👨‍💻 Tam Kod: Şifreli Giriş Sistemli Uygulama

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Kullanici struct {
	Isim  string  `json:"isim"`
	Yas   int     `json:"yas"`
	Maas  float64 `json:"maas"`
	Sifre string  `json:"sifre"`
}

var kullanicilar []Kullanici
var aktifKullanici *Kullanici // Oturum açan kullanıcıyı tutar

func main() {

	//JSON'dan kullanıcıları otomatik yükle
	
	yuklenen, err := jsonYukle()
	if err == nil {
		kullanicilar = yuklenen
	}

	for {
		fmt.Println("\n📋 Menü:")
		fmt.Println("1. Yeni Kullanıcı Oluştur")
		fmt.Println("2. Giriş Yap")
		fmt.Println("3. Aktif Kullanıcı Bilgileri")
		fmt.Println("4. Kullanıcıları Kaydet (JSON)")
		fmt.Println("5. Çıkış")
		fmt.Print("Bir seçenek girin (1-5): ")

		var secim int
		fmt.Scanln(&secim)

		switch secim {
		case 1:
			kullanici := kullaniciOlustur()
			kullanicilar = append(kullanicilar, kullanici)
			fmt.Println("✅ Kullanıcı başarıyla oluşturuldu.")
		case 2:
			girisYap()
		case 3:
			if aktifKullanici == nil {
				fmt.Println("❌ Giriş yapılmamış.")
			} else {
				k := aktifKullanici
				fmt.Println("🔐 Aktif Kullanıcı:")
				fmt.Printf("İsim: %s | Yaş: %d | Maaş: %.2f TL\n", k.Isim, k.Yas, k.Maas)
			}
		case 4:
			err := jsonKaydet(kullanicilar)
			if err != nil {
				fmt.Println("❌ Kaydetme hatası:", err)
			} else {
				fmt.Println("💾 Kullanıcılar kaydedildi (kullanicilar.json).")
			}
		case 5:
			fmt.Println("👋 Çıkış yapılıyor...")
			return
		default:
			fmt.Println("❌ Geçersiz seçim.")
		}
	}
}


// Yeni kullanıcı oluşturur
func kullaniciOlustur() Kullanici {
	var k Kullanici
	fmt.Print("İsim: ")
	fmt.Scanln(&k.Isim)

	fmt.Print("Yaş: ")
	fmt.Scanln(&k.Yas)

	fmt.Print("Maaş: ")
	fmt.Scanln(&k.Maas)

	fmt.Print("Şifre belirleyin: ")
	fmt.Scanln(&k.Sifre)

	return k
}

// Şifre doğrulamalı giriş
func girisYap() {
	var isim, sifre string
	fmt.Print("İsminizi girin: ")
	fmt.Scanln(&isim)

	fmt.Print("Şifrenizi girin: ")
	fmt.Scanln(&sifre)

	for i := range kullanicilar {
		if kullanicilar[i].Isim == isim && kullanicilar[i].Sifre == sifre {
			aktifKullanici = &kullanicilar[i]
			fmt.Printf("✅ Hoş geldin, %s!\n", isim)
			return
		}
	}
	fmt.Println("❌ Hatalı giriş. Şifre ya da kullanıcı adı yanlış.")
}

// Kullanıcıları JSON olarak kaydet
func jsonKaydet(liste []Kullanici) error {
	file, err := os.Create("kullanicilar.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(liste)
}

// JSON'dan kullanıcıları yükle
func jsonYukle() ([]Kullanici, error) {
	file, err := os.Open("kullanicilar.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var liste []Kullanici
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&liste)
	return liste, err
}
``
/*
---

## 🔐 Örnek Akış
*/
``
📋 Menü:
1. Yeni Kullanıcı Oluştur
2. Giriş Yap
3. Aktif Kullanıcı Bilgileri
4. Kullanıcıları Kaydet (JSON)
5. Çıkış
Bir seçenek girin (1-5): 1
İsim: Ayşe
Yaş: 27
Maaş: 20000
Şifre belirleyin: 1234
✅ Kullanıcı başarıyla oluşturuldu.

📋 Menü:
Bir seçenek girin (1-5): 2
İsminizi girin: Ayşe
Şifrenizi girin: 1234
✅ Hoş geldin, Ayşe!

📋 Menü:
Bir seçenek girin (1-5): 3
🔐 Aktif Kullanıcı:
İsim: Ayşe | Yaş: 27 | Maaş: 20000.00 TL
``
/*
---

## 🧠 Öğrenilenler

| Konsept               | Açıklama                                      |
| --------------------- | --------------------------------------------- |
| `struct` içinde şifre | Güvenlik için kullanıcıya ait şifre saklama   |
| `*Kullanici`          | Pointer ile aktif kullanıcı referansı         |
| `slice` ile arama     | Döngü ile kullanıcı & şifre eşleşmesini bulma |
| JSON encode/decode    | Liste halinde veri saklama                    |

---

## ⚠️ Güvenlik Notu

Bu örnekte şifreler **düz metin** olarak saklanıyor. Gerçek bir uygulamada:

* Şifreler **hash** (örn. bcrypt) ile saklanmalı
* Terminal şifre girişi **gizli** yapılmalı (örn. `golang.org/x/term` ile)

İstersen bir sonraki adım olarak:
👉 "Şifreleri hash ile güvenli saklama" özelliği ekleyebiliriz.

Devam etmek ister misin?
*/