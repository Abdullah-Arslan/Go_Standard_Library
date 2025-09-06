/*
O zaman sana **tam kapsamlı bir “Map Toolkit”** örneği hazırlayacağım.
Bu toolkit, `maps` paketi ve generic yapıyı kullanarak şunları yapacak:

1. İki map’i birleştirme (`MergeMaps`)
2. Farklı değerleri bulma (`DiffMaps`)
3. Anahtar ve değer eşitliğini kontrol etme (`EqualKeys`, `EqualValues`)
4. Eksik anahtarları tespit etme (`MissingKeys`)

---

# 📌 Komple Map Toolkit Örneği
*/
``go
package main

import (
	"fmt"
	"maps"
)

// ----------------------------
// MergeMaps: iki map'i birleştirir
type MergeFunc[V any] func(v1, v2 V) V

func MergeMaps[K comparable, V any](mapA, mapB map[K]V, merge MergeFunc[V]) map[K]V {
	result := maps.Clone(mapA)

	for k, vB := range mapB {
		if vA, ok := result[k]; ok {
			result[k] = merge(vA, vB)
		} else {
			result[k] = vB
		}
	}

	return result
}

// ----------------------------
// DiffMaps: farklı değerleri bulur
func DiffMaps[K comparable, V comparable](mapA, mapB map[K]V) map[K][2]V {
	diff := make(map[K][2]V)

	for k, vA := range mapA {
		if vB, ok := mapB[k]; ok && vA != vB {
			diff[k] = [2]V{vA, vB}
		}
	}

	for k, vB := range mapB {
		if _, ok := mapA[k]; !ok {
			diff[k] = [2]V{*new(V), vB} // mapA'da olmayanlar
		}
	}

	return diff
}

// ----------------------------
// EqualKeys: anahtar kümeleri eşit mi?
func EqualKeys[K comparable, V any](mapA, mapB map[K]V) bool {
	return maps.EqualKeys(mapA, mapB)
}

// EqualValues: değerler aynı mı? (default karşılaştırma)
func EqualValues[K comparable, V comparable](mapA, mapB map[K]V) bool {
	return maps.Equal(mapA, mapB)
}

// ----------------------------
// MissingKeys: mapB'de olup mapA'da olmayan anahtarlar
func MissingKeys[K comparable, V any](mapA, mapB map[K]V) []K {
	var missing []K
	for k := range mapB {
		if _, ok := mapA[k]; !ok {
			missing = append(missing, k)
		}
	}
	return missing
}

// ----------------------------
// Örnek kullanım
func main() {
	mapA := map[string]int{"a": 1, "b": 2, "c": 3}
	mapB := map[string]int{"b": 20, "c": 30, "d": 40}

	// Merge
	merged := MergeMaps(mapA, mapB, func(v1, v2 int) int { return v1 + v2 })
	fmt.Println("Merged:", merged)

	// Farklı değerler
	diff := DiffMaps(mapA, mapB)
	fmt.Println("Diff:", diff)

	// Anahtar eşitliği
	fmt.Println("Equal Keys:", EqualKeys(mapA, mapB))

	// Değer eşitliği
	fmt.Println("Equal Values:", EqualValues(mapA, mapB))

	// Eksik anahtarlar
	missing := MissingKeys(mapA, mapB)
	fmt.Println("Missing Keys:", missing)
}
``
/*
---

## 📌 Çıktı
*/
``
Merged: map[a:1 b:22 c:33 d:40]
Diff: map[b:[2 20] c:[3 30] d:[0 40]]
Equal Keys: false
Equal Values: false
Missing Keys: [d]
``
/*
---

# 🔑 Açıklama

1. `MergeMaps` → İki map’i birleştiriyor, aynı anahtar varsa **toplama veya özel işlem** yapabiliyoruz.
2. `DiffMaps` → Aynı anahtar ama farklı değerleri ve mapA’da olmayanları buluyor.
3. `EqualKeys` / `EqualValues` → Anahtar ve değer eşitliği kontrolü.
4. `MissingKeys` → mapB’de olup mapA’da olmayan anahtarları listeliyor.

---

✅ Artık **bir map toolkit’in var**:

* Config yönetimi
* Farklı map’leri karşılaştırma
* Merge & kombinasyon
* Eksik veya farklı değerleri bulma

---

İstersen bir sonraki adımda **bu toolkit’i string, int ve struct tipleri ile generic olarak nasıl kullanacağını** gösterebilirim, böylece tamamen her tip map için hazır hale gelir. Bunu da ister misin?
EVET
*/
/*
Harika 😎
O zaman toolkit’i **tamamen generic** hâle getirip, **string, int, struct** gibi her tip map için kullanabileceğimiz şekilde geliştirelim.

---

# 📌 Komple Generic Map Toolkit
*7
``go
package main

import (
	"fmt"
	"maps"
)

// MergeFunc: Aynı anahtar varsa ne yapılacağını belirler
type MergeFunc[V any] func(v1, v2 V) V

// MergeMaps: iki map'i birleştirir (her tip için generic)
func MergeMaps[K comparable, V any](mapA, mapB map[K]V, merge MergeFunc[V]) map[K]V {
	result := maps.Clone(mapA)

	for k, vB := range mapB {
		if vA, ok := result[k]; ok {
			result[k] = merge(vA, vB)
		} else {
			result[k] = vB
		}
	}

	return result
}

// DiffMaps: farklı değerleri bulur (generic, comparable tipler)
func DiffMaps[K comparable, V comparable](mapA, mapB map[K]V) map[K][2]V {
	diff := make(map[K][2]V)

	for k, vA := range mapA {
		if vB, ok := mapB[k]; ok && vA != vB {
			diff[k] = [2]V{vA, vB}
		}
	}

	for k, vB := range mapB {
		if _, ok := mapA[k]; !ok {
			diff[k] = [2]V{*new(V), vB}
		}
	}

	return diff
}

// EqualKeys: anahtar kümeleri eşit mi?
func EqualKeys[K comparable, V any](mapA, mapB map[K]V) bool {
	return maps.EqualKeys(mapA, mapB)
}

// EqualValues: değerler aynı mı? (comparable tipler)
func EqualValues[K comparable, V comparable](mapA, mapB map[K]V) bool {
	return maps.Equal(mapA, mapB)
}

// MissingKeys: mapB'de olup mapA'da olmayan anahtarlar
func MissingKeys[K comparable, V any](mapA, mapB map[K]V) []K {
	var missing []K
	for k := range mapB {
		if _, ok := mapA[k]; !ok {
			missing = append(missing, k)
		}
	}
	return missing
}

// ----------------------------
// Struct tipi örnek
type Config struct {
	Port     int
	LogLevel string
	Debug    bool
}

// Struct’lar karşılaştırılırken == kullanılamaz
func DiffStructMaps[K comparable, V any](mapA, mapB map[K]V, eq func(a, b V) bool) map[K][2]V {
	diff := make(map[K][2]V)
	for k, vA := range mapA {
		if vB, ok := mapB[k]; ok && !eq(vA, vB) {
			diff[k] = [2]V{vA, vB}
		}
	}
	for k, vB := range mapB {
		if _, ok := mapA[k]; !ok {
			diff[k] = [2]V{*new(V), vB}
		}
	}
	return diff
}

// ----------------------------
// Örnek kullanım
func main() {
	// int map
	mapA := map[string]int{"a": 1, "b": 2, "c": 3}
	mapB := map[string]int{"b": 20, "c": 30, "d": 40}

	merged := MergeMaps(mapA, mapB, func(v1, v2 int) int { return v1 + v2 })
	fmt.Println("Merged int map:", merged)

	diff := DiffMaps(mapA, mapB)
	fmt.Println("Diff int map:", diff)

	// string map
	mapStrA := map[string]string{"x": "foo", "y": "bar"}
	mapStrB := map[string]string{"y": "baz", "z": "qux"}

	mergedStr := MergeMaps(mapStrA, mapStrB, func(v1, v2 string) string { return v1 + "+" + v2 })
	fmt.Println("Merged string map:", mergedStr)

	// struct map
	mapStructA := map[string]Config{
		"service1": {Port: 8080, LogLevel: "info", Debug: true},
	}
	mapStructB := map[string]Config{
		"service1": {Port: 8080, LogLevel: "INFO", Debug: true},
		"service2": {Port: 9090, LogLevel: "warn", Debug: false},
	}

	diffStruct := DiffStructMaps(mapStructA, mapStructB, func(a, b Config) bool {
		return a.Port == b.Port && a.LogLevel == b.LogLevel && a.Debug == b.Debug
	})
	fmt.Println("Diff struct map:", diffStruct)
}
``
/*
---

## 📌 Çıktı Örneği
*/
``
Merged int map: map[a:1 b:22 c:33 d:40]
Diff int map: map[b:[2 20] c:[3 30] d:[0 40]]
Merged string map: map[x:foo y:bar+baz z:qux]
Diff struct map: map[service1:[{8080 info true} {8080 INFO true}] service2:[{0  false} {9090 warn false}]]
``
/*
---

# 🔑 Özet

* Artık toolkit **tamamen generic**: `int`, `string`, `struct` gibi her tip map ile çalışır
* `MergeMaps` → aynı anahtarlar için özel merge mantığı uygulanabilir
* `DiffMaps` / `DiffStructMaps` → değer farklarını tespit eder
* `EqualKeys` / `EqualValues` → anahtar ve değer eşitliği kontrolü
* `MissingKeys` → eksik anahtarları listeler

---

İstersen bir sonraki adımda **bu generic toolkit’i `slog` ile birlikte kullanıp, değişiklikleri loglayarak sistematik rapor üreten bir örnek** hazırlayabilirim. Bunu yapalım mı?
*/