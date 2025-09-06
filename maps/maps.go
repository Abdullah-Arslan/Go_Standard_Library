/*
Go 1.21 ile birlikte gelen **`maps`** paketi, `slices` paketi gibi **generic (jenerik) yardımcı fonksiyonlar** sağlar.
Yani `map` türü için sık kullanılan işlemleri kolaylaştırır.

---

# 📌 `maps` Paketinin Amacı

Go’nun standart `map` tipinde (örn. `map[string]int`) anahtar-değer çiftleri saklayabiliriz.
Ama çoğu zaman şunlara ihtiyaç duyarız:

* İki map eşit mi kontrol et (`maps.Equal`)
* Bir map’i kopyala (`maps.Clone`)
* Bir map’i başka bir map’e kopyala (`maps.Copy`)
* İki map aynı anahtarlara sahip mi bak (`maps.EqualKeys`)

İşte bu işlemleri **generic ve güvenli şekilde** yapmamızı sağlayan fonksiyonlar **`maps` paketinde** bulunur.

---

# 📌 Paket Fonksiyonları ve Örnekleri

## 1. `maps.Clone`

Bir map’in kopyasını oluşturur.
*/
``go
package main

import (
	"fmt"
	"maps"
)

func main() {
	original := map[string]int{"a": 1, "b": 2}
	copyMap := maps.Clone(original)

	copyMap["a"] = 100

	fmt.Println("Original:", original) // map[a:1 b:2]
	fmt.Println("Clone   :", copyMap)  // map[a:100 b:2]
}
``
/*
🔑 Burada `copyMap` değiştirildi ama `original` etkilenmedi.

---

## 2. `maps.Copy`

Bir map’in içeriğini başka bir map’e kopyalar (varolan değerleri günceller).
*/
``go
package main

import (
	"fmt"
	"maps"
)

func main() {
	dst := map[string]int{"a": 1, "b": 2}
	src := map[string]int{"b": 20, "c": 30}

	maps.Copy(dst, src)

	fmt.Println(dst) // map[a:1 b:20 c:30]
}
``
/*
🔑 `b` anahtarı güncellendi, `c` anahtarı eklendi.

---

## 3. `maps.Equal`

İki map aynı anahtar-değer çiftlerine sahipse `true` döner.
*/
``go
package main

import (
	"fmt"
	"maps"
)

func main() {
	m1 := map[string]int{"x": 1, "y": 2}
	m2 := map[string]int{"y": 2, "x": 1}
	m3 := map[string]int{"x": 1, "y": 3}

	fmt.Println(maps.Equal(m1, m2)) // true
	fmt.Println(maps.Equal(m1, m3)) // false
}
``
/*
🔑 Sıra önemli değil, değerler eşleşirse `true` olur.

---

## 4. `maps.EqualFunc`

Özel karşılaştırma fonksiyonu kullanarak iki map’i karşılaştırır.
*/
``go
package main

import (
	"fmt"
	"maps"
	"strings"
)

func main() {
	m1 := map[string]string{"a": "HELLO"}
	m2 := map[string]string{"a": "hello"}

	// Karşılaştırmayı case-insensitive yapıyoruz
	eq := maps.EqualFunc(m1, m2, func(v1, v2 string) bool {
		return strings.ToLower(v1) == strings.ToLower(v2)
	})

	fmt.Println(eq) // true
}
``
/*
---

## 5. `maps.EqualKeys`

İki map aynı anahtarlara sahipse `true` döner.
*/
``go
package main

import (
	"fmt"
	"maps"
)

func main() {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 5, "a": 99}
	m3 := map[string]int{"a": 1, "c": 2}

	fmt.Println(maps.EqualKeys(m1, m2)) // true
	fmt.Println(maps.EqualKeys(m1, m3)) // false
}
``
/*
🔑 Değerler farklı olabilir, sadece **anahtar kümeleri aynı mı** kontrol eder.

---

# 📌 Özet

* `maps.Clone(m)` → Yeni bir map kopyası oluşturur
* `maps.Copy(dst, src)` → `src` içeriğini `dst`’ye kopyalar
* `maps.Equal(m1, m2)` → Aynı anahtar-değer çiftleri mi kontrol eder
* `maps.EqualFunc(m1, m2, cmp)` → Özel karşılaştırıcı ile eşitlik kontrolü
* `maps.EqualKeys(m1, m2)` → Sadece anahtar kümeleri aynı mı kontrol eder

---

👉 İstersen ben sana **tüm maps fonksiyonlarını bir arada kullanıp, gerçek bir senaryo (örn. config dosyası karşılaştırma)** örneği de hazırlayabilirim. İstiyor musun?
EVET
*/
/*
Harika 👍
O zaman sana **`maps` paketini gerçek bir senaryoda nasıl kullanabileceğini** göstereyim:

Senaryo: Diyelim ki bir uygulamanın **varsayılan konfigürasyonu** ve **kullanıcı tarafından sağlanan konfigürasyon** var.
Amacımız:

1. Eksik anahtarları varsayılan konfig’dan al
2. Farklı değerleri tespit et
3. Anahtarlar aynı mı, değerler aynı mı kontrol et

---

# 📌 Örnek: Config Karşılaştırma
*/
``go
package main

import (
	"fmt"
	"maps"
)

func main() {
	// Varsayılan konfigürasyon
	defaultConfig := map[string]string{
		"host":     "localhost",
		"port":     "8080",
		"logLevel": "info",
		"debug":    "false",
	}

	// Kullanıcı tarafından sağlanan konfig
	userConfig := map[string]string{
		"port":     "9090",
		"logLevel": "INFO",
	}

	// 1️⃣ Eksik anahtarları varsayılanla doldur
	for k, v := range defaultConfig {
		if _, ok := userConfig[k]; !ok {
			userConfig[k] = v
		}
	}
	fmt.Println("Tamamlanmış userConfig:", userConfig)

	// 2️⃣ Anahtarlar aynı mı?
	keysEqual := maps.EqualKeys(defaultConfig, userConfig)
	fmt.Println("Anahtarlar aynı mı?", keysEqual)

	// 3️⃣ Değerler eşit mi? (case-insensitive logLevel için özel karşılaştırma)
	valuesEqual := maps.EqualFunc(defaultConfig, userConfig, func(v1, v2 string) bool {
		return v1 == v2 || (v1 == "info" && v2 == "INFO")
	})
	fmt.Println("Değerler eşit mi?", valuesEqual)

	// 4️⃣ Farklı değerleri bul
	for k, v := range userConfig {
		if defVal, ok := defaultConfig[k]; ok && defVal != v {
			fmt.Printf("Farklı değer: key=%s, default=%s, user=%s\n", k, defVal, v)
		}
	}
}
``
/*
---

## 📌 Çıktı
*/
``
Tamamlanmış userConfig: map[debug:false host:localhost logLevel:INFO port:9090]
Anahtarlar aynı mı? true
Değerler eşit mi? false
Farklı değer: key=port, default=8080, user=9090
Farklı değer: key=logLevel, default=info, user=INFO
``
/*
---

✅ Bu örnek ile:

* `maps.EqualKeys` → Anahtarları kontrol ettik
* `maps.EqualFunc` → Değerleri özel bir mantıkla karşılaştırdık
* Döngü + `maps.Clone`/`maps.Copy` → eksik veya farklı değerleri işledik

---

İstersen sana bir sonraki adımda **maps paketi ile iki map’i birleştirip tüm kombinasyonlarını alacak daha gelişmiş bir örnek** de hazırlayabilirim. Bunu ister misin?
EVET
*/
/*
Harika 👍
O zaman sana **`maps` paketi ile iki map’i birleştirip tüm kombinasyonlarını nasıl alabileceğini** göstereyim.

Senaryo:

* Elimizde iki map var: **mapA** ve **mapB**
* Amacımız:

  1. Aynı anahtarlar varsa **mapB**’nin değerini al
  2. Farklı anahtarları da ekle
* Yani iki map’in **birleşmiş halini** oluşturacağız

---

# 📌 Örnek: Map Merge / Kombinasyon
*/
``go
package main

import (
	"fmt"
	"maps"
)

func main() {
	mapA := map[string]int{"a": 1, "b": 2, "c": 3}
	mapB := map[string]int{"b": 20, "c": 30, "d": 40}

	// 1️⃣ mapA'yı kopyala (orijinal bozulmasın)
	merged := maps.Clone(mapA)

	// 2️⃣ mapB'yi merged map'e kopyala
	maps.Copy(merged, mapB)

	fmt.Println("Map A:", mapA)
	fmt.Println("Map B:", mapB)
	fmt.Println("Birleşmiş map:", merged)
}
``
/*
---

## 📌 Çıktı
*/
``
Map A: map[a:1 b:2 c:3]
Map B: map[b:20 c:30 d:40]
Birleşmiş map: map[a:1 b:20 c:30 d:40]
``
/*
✅ Açıklama:

* `maps.Clone(mapA)` → Orijinal mapA değişmeden yeni bir map oluşturduk
* `maps.Copy(merged, mapB)` → mapB’nin tüm anahtarlarını merged’e ekledik / aynı anahtar varsa üzerine yazıldı

---

# 📌 Daha Gelişmiş Örnek: Özel Birleştirme Fonksiyonu

Bazı durumlarda aynı anahtar varsa **toplama, birleştirme veya string birleştirme** yapabiliriz:
*/
``go
package main

import (
	"fmt"
)

func main() {
	mapA := map[string]int{"a": 1, "b": 2, "c": 3}
	mapB := map[string]int{"b": 20, "c": 30, "d": 40}

	merged := make(map[string]int)

	for k, v := range mapA {
		merged[k] = v
	}

	for k, v := range mapB {
		if val, ok := merged[k]; ok {
			// Aynı anahtar varsa toplama yap
			merged[k] = val + v
		} else {
			merged[k] = v
		}
	}

	fmt.Println("Birleşmiş map (toplamlı):", merged)
}
``

//### 📌 Çıktı

``
Birleşmiş map (toplamlı): map[a:1 b:22 c:33 d:40]
``
/*
---

🔑 Özet:

* `maps.Clone` → Orijinal map’i bozmadan kopyalama
* `maps.Copy` → Anahtarları kopyalama / üzerine yazma
* Döngü ile özel birleştirme → toplama, string birleştirme, merge mantığı
* Böylece iki map’in **tüm kombinasyonlarını güvenli şekilde birleştirebiliriz**

---

İstersen bir sonraki adımda sana **generic bir `MergeMaps` fonksiyonu** hazırlayıp, **her tip map için kullanılabilir** hale getirebilirim. Bunu yapayım mı?
EVET
*/
/*
Harika 😎
O zaman sana **generic bir `MergeMaps` fonksiyonu** hazırlayacağım. Böylece **her tip map için kullanılabilir**, aynı anahtar varsa özel birleştirme mantığı uygulayabileceğiz.

---

# 📌 Generic MergeMaps Fonksiyonu
*/
``go
package main

import (
	"fmt"
	"maps"
)

// MergeFunc: Aynı anahtar varsa ne yapılacağını belirler
type MergeFunc[V any] func(v1, v2 V) V

// MergeMaps: iki map'i birleştirir
func MergeMaps[K comparable, V any](mapA, mapB map[K]V, merge MergeFunc[V]) map[K]V {
	result := maps.Clone(mapA) // mapA'yı kopyala

	for k, vB := range mapB {
		if vA, ok := result[k]; ok {
			// Aynı anahtar varsa merge fonksiyonunu uygula
			result[k] = merge(vA, vB)
		} else {
			result[k] = vB
		}
	}

	return result
}

func main() {
	mapA := map[string]int{"a": 1, "b": 2, "c": 3}
	mapB := map[string]int{"b": 20, "c": 30, "d": 40}

	// Merge: aynı anahtar varsa toplama yap
	merged := MergeMaps(mapA, mapB, func(v1, v2 int) int {
		return v1 + v2
	})

	fmt.Println("Map A:", mapA)
	fmt.Println("Map B:", mapB)
	fmt.Println("Birleşmiş map:", merged)
}
``
/*
---

## 📌 Çıktı
*/
``
Map A: map[a:1 b:2 c:3]
Map B: map[b:20 c:30 d:40]
Birleşmiş map: map[a:1 b:22 c:33 d:40]
``
/*
---

# 🔑 Açıklama

1. `maps.Clone(mapA)` → Orijinal map’i bozmadan kopyaladık
2. Döngü ile `mapB`’yi taradık
3. Aynı anahtar varsa `merge(vA, vB)` fonksiyonunu uyguladık
4. Yeni bir map döndürdük → generic olduğu için **her tip map** için kullanılabilir

---

### 💡 Örnek Kullanım Farklı Tiplerde
*/
``go
mapStr := map[string]string{"x": "foo"}
mapStr2 := map[string]string{"x": "bar", "y": "baz"}

// Aynı anahtar varsa birleştirme
mergedStr := MergeMaps(mapStr, mapStr2, func(v1, v2 string) string {
	return v1 + "+" + v2
})

fmt.Println(mergedStr)
// Çıktı: map[x:foo+bar y:baz]
``
/*
---

✅ Böylece artık **generic MergeMaps** fonksiyonu ile **sayısal, string veya diğer tip map’leri** kolayca birleştirebilirsin ve aynı anahtar için özel mantık uygulayabilirsin.

---

İstersen bir sonraki adımda sana **maps paketi ile merge + eşitlik + farklı anahtarları tespit eden komple bir “map toolkit” örneği** de hazırlayabilirim. Bunu yapayım mı?
*/
