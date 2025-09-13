/*
Harika bir konuya geldik 👍 Go’nun **`regexp`** paketi, **düzenli ifadeler (regular expressions)** ile metin üzerinde arama, eşleme (match), değiştirme (replace), bölme (split) gibi güçlü işlemler yapmanı sağlar.

Ben sana şimdi bu paketi **tüm önemli fonksiyonları ve tipleriyle**, **örneklerle** ve **gerçek kullanım senaryolarıyla** açıklayacağım.

---

# 📦 Go `regexp` Paketi

## 1. Temel Kavramlar

* Go’nun `regexp` paketi **RE2** motorunu kullanır (Google’ın hızlı ve güvenli regex motoru).
* PCRE’nin (`perl`, `python` regex) çoğu özelliğini destekler ama bazılarını (ör. lookbehind) desteklemez.
* Regex ifadeleri **string** olarak yazılır:
*/
  ``go
  r := regexp.MustCompile("[a-z]+")
  ``
/*
---

## 2. Önemli Fonksiyonlar ve Metotlar

### ✅ `regexp.Compile` / `MustCompile`

Regex ifadesini derler.

* `Compile` → hata döndürür.
* `MustCompile` → hata durumunda panik atar (çoğunlukla tercih edilen).
*/
``go
r1, err := regexp.Compile("a+b")
fmt.Println(r1.MatchString("aaab")) // true
fmt.Println(err) // nil

r2 := regexp.MustCompile("[0-9]+")
fmt.Println(r2.MatchString("1234")) // true
``
/*
---

### ✅ `MatchString` – Basit Eşleşme
*/
``go
r := regexp.MustCompile("^go")
fmt.Println(r.MatchString("golang")) // true
fmt.Println(r.MatchString("python")) // false
``
/*
---

### ✅ `FindString` – İlk Eşleşmeyi Döndür
*/
``go
r := regexp.MustCompile("[0-9]+")
s := "abc123def456"
fmt.Println(r.FindString(s)) // 123
``
/*
---

### ✅ `FindAllString` – Tüm Eşleşmeleri Döndür
*/
``go
r := regexp.MustCompile("[0-9]+")
s := "abc123def456"
fmt.Println(r.FindAllString(s, -1)) // [123 456]
``
/*
📌 İkinci parametre limit → `-1` = hepsini bul.

---

### ✅ `FindStringIndex`

Eşleşmenin başlangıç ve bitiş indexlerini döndürür.
*/
``go
r := regexp.MustCompile("def")
fmt.Println(r.FindStringIndex("abc123def456")) // [6 9]
``
/*
---

### ✅ `FindAllStringIndex`

Tüm eşleşmelerin indexlerini döndürür.
*/
``go
r := regexp.MustCompile("[0-9]+")
fmt.Println(r.FindAllStringIndex("a1b22c333", -1))
// [[1 2] [3 5] [6 9]]
``
/*
---

### ✅ `FindStringSubmatch` – Grup Yakalama

Regex içinde parantez `( )` gruplarını döndürür.
*/
``go
r := regexp.MustCompile(`(\w+)@(\w+)\.(\w+)`)
s := "mail: test@example.com"
match := r.FindStringSubmatch(s)
fmt.Println(match) // ["test@example.com" "test" "example" "com"]
``
/*
---

### ✅ `FindAllStringSubmatch`

Birden fazla eşleşmede grup sonuçlarını alır.
*/
``go
r := regexp.MustCompile(`(\d+)-(\w+)`)
s := "1-one 2-two 3-three"
fmt.Println(r.FindAllStringSubmatch(s, -1))
``

//📌 Çıktı:

``
[[1-one 1 one] [2-two 2 two] [3-three 3 three]]
``
/*
---

### ✅ `ReplaceAllString`

Eşleşmeleri başka bir string ile değiştir.
*/
``go
r := regexp.MustCompile("[0-9]+")
s := "abc123def456"
fmt.Println(r.ReplaceAllString(s, "#")) 
// abc#def#
``
/*
---

### ✅ `ReplaceAllStringFunc`

Eşleşmeleri bir fonksiyonla dönüştür.
*/
``go
r := regexp.MustCompile("[0-9]+")
s := "abc123def456"
result := r.ReplaceAllStringFunc(s, func(m string) string {
	return "[" + m + "]"
})
fmt.Println(result)
// abc[123]def[456]
``
/*
---

### ✅ `Split`

Regex’e göre string’i böler.
*/
``go
r := regexp.MustCompile("[,; ]+")
fmt.Println(r.Split("a,b; c  d", -1))
// [a b c d]
``
/*
---

## 3. Flags (Opsiyonlar)

* `(?i)` → case-insensitive
* `(?m)` → multi-line
* `(?s)` → dot (.) newline eşleşsin

Örnek:
*/
``go
r := regexp.MustCompile(`(?i)go`)
fmt.Println(r.MatchString("GOlang")) // true
``
/*
---

## 4. Kullanışlı Örnekler

### 📌 Email Doğrulama
*/
``go
r := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
fmt.Println(r.MatchString("test@example.com")) // true
fmt.Println(r.MatchString("wrong@com"))        // false
``
/*
---

### 📌 Telefon Numarası (Türkiye formatı)
*/
``go
r := regexp.MustCompile(`^\+90\d{10}$`)
fmt.Println(r.MatchString("+905321234567")) // true
fmt.Println(r.MatchString("5321234567"))    // false
``
/*
---

### 📌 HTML Tag Çekme
*/
``go
r := regexp.MustCompile(`<([a-z]+)>.*?</\1>`)
fmt.Println(r.MatchString("<b>bold</b>"))   // true
fmt.Println(r.MatchString("<i>italic</b>")) // false
``
/*
---

## 5. Özet – `regexp` ile Yapabileceklerin

✅ String içinde arama (`MatchString`, `FindString`)
✅ Tüm eşleşmeleri bulma (`FindAll*`)
✅ Gruplama ve parantezlerle yakalama (`Submatch`)
✅ Değiştirme (`ReplaceAllString`, `ReplaceAllStringFunc`)
✅ Bölme (`Split`)
✅ Flags (`(?i)`, `(?m)`, `(?s)`)

---

🔑 **Sonuç:**
Go’nun `regexp` paketi, metin işleme, log analizi, form doğrulama, parser yazma gibi yerlerde çok güçlüdür. JSON, XML veya HTML’den veri çekmek için değil, **pattern tabanlı eşleşme** için idealdir.

---

👉 İstersen sana bunun üstüne **küçük bir “Regex Tester CLI” programı** yazabilirim (komut satırından regex girip test eden). Böylece kendi regexlerini hızlıca deneyebilirsin.

Bunu da ister misin?
EVET
*/
/*
Süper 👍 O zaman sana **Go ile küçük bir Regex Tester CLI** hazırlayayım.
Bu programla terminalden regex deseni ve test string’i gireceksin → eşleşme sonuçlarını gösterecek.

---

# 📦 Regex Tester CLI (Go)

## Kod
*/
``go
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("🔎 Go Regex Tester")
	fmt.Println("-------------------")

	// Regex deseni al
	fmt.Print("Regex desenini gir: ")
	pattern, _ := reader.ReadString('\n')
	pattern = pattern[:len(pattern)-1] // newline sil

	// Regex derle
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("❌ Regex hatalı:", err)
		return
	}

	// Test string al
	fmt.Print("Test edilecek string gir: ")
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]

	// Basit match
	match := re.MatchString(text)
	fmt.Println("Match sonucu:", match)

	// İlk eşleşme
	first := re.FindString(text)
	fmt.Println("İlk eşleşme:", first)

	// Tüm eşleşmeler
	all := re.FindAllString(text, -1)
	fmt.Println("Tüm eşleşmeler:", all)

	// Submatch (gruplar)
	sub := re.FindStringSubmatch(text)
	if len(sub) > 0 {
		fmt.Println("Gruplar:", sub[1:])
	}
}
``
/*
---

## Kullanım

1. Kaydet: `regex_tester.go`

2. Çalıştır:
*/
   ``bash
   go run regex_tester.go
   ``
/*
3. Örnek test:
*/
``
🔎 Go Regex Tester
-------------------
Regex desenini gir: (\d+)-(\w+)
Test edilecek string gir: 1-one 2-two 3-three
Match sonucu: true
İlk eşleşme: 1-one
Tüm eşleşmeler: [1-one 2-two 3-three]
Gruplar: [1 one]
``
/*
---

## Özellikler

* Hatalı regex girersen uyarır.
* İlk eşleşme ve tüm eşleşmeleri gösterir.
* Grupları (`()` içindekileri) yakalar.
* CLI üzerinden istediğin kadar deneyebilirsin.

---

👉 İstersen bunu daha da geliştirip:

* **sonsuz döngü** içinde çalıştırabiliriz (her seferinde yeni regex ve string sormadan, sadece string girip denersin).
* **renkli çıktı** (örn. eşleşen kısımları sarıyla vurgulama).
* **flag parametreleri** (örn. `-pattern` ve `-text` ile doğrudan terminalden test).

Bu özelliklerden hangisini görmek istersin?
EVET
*/
/*
Harika 🚀 O zaman Regex Tester CLI’yi **daha gelişmiş** hale getirelim:

* Sonsuz döngüde çalışsın → program kapanmadan tekrar tekrar test yapılabilsin.
* Terminal parametreleriyle (`-pattern` ve `-text`) doğrudan regex testi yapılabilsin.
* Eşleşen kısımlar **renkli** (sarı) vurgulansın.

---

# 📦 Gelişmiş Regex Tester CLI

## Kod
*/
``go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const yellow = "\033[33m"
const reset = "\033[0m"

func highlightMatches(text string, re *regexp.Regexp) string {
	// Eşleşmeleri renklendir
	return re.ReplaceAllStringFunc(text, func(m string) string {
		return yellow + m + reset
	})
}

func main() {
	// Komut satırı argümanları
	patternFlag := flag.String("pattern", "", "Regex deseni")
	textFlag := flag.String("text", "", "Test edilecek string")
	flag.Parse()

	// Eğer flag ile verilmişse direkt test yap
	if *patternFlag != "" && *textFlag != "" {
		runRegexTest(*patternFlag, *textFlag)
		return
	}

	// Aksi halde interaktif mod
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("🔎 Go Regex Tester (CTRL+C ile çık)")

	for {
		// Regex deseni al
		fmt.Print("\nRegex desenini gir: ")
		pattern, _ := reader.ReadString('\n')
		pattern = strings.TrimSpace(pattern)

		if pattern == "" {
			fmt.Println("❌ Desen boş olamaz")
			continue
		}

		// Regex derle
		re, err := regexp.Compile(pattern)
		if err != nil {
			fmt.Println("❌ Regex hatalı:", err)
			continue
		}

		// Test string al
		fmt.Print("Test edilecek string gir: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		runRegexTestWithCompiled(re, text)
	}
}

func runRegexTest(pattern, text string) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("❌ Regex hatalı:", err)
		return
	}
	runRegexTestWithCompiled(re, text)
}

func runRegexTestWithCompiled(re *regexp.Regexp, text string) {
	fmt.Println("Match sonucu:", re.MatchString(text))
	fmt.Println("İlk eşleşme:", re.FindString(text))
	fmt.Println("Tüm eşleşmeler:", re.FindAllString(text, -1))
	sub := re.FindStringSubmatch(text)
	if len(sub) > 1 {
		fmt.Println("Gruplar:", sub[1:])
	}
	fmt.Println("Renkli gösterim:", highlightMatches(text, re))
}
``
/*
---

## Kullanım

### 1. Normal çalıştır (interaktif mod)
*/
``bash
go run regex_tester.go
``

``
🔎 Go Regex Tester (CTRL+C ile çık)

Regex desenini gir: (\d+)-(\w+)
Test edilecek string gir: 1-one 2-two 3-three
Match sonucu: true
İlk eşleşme: 1-one
Tüm eşleşmeler: [1-one 2-two 3-three]
Gruplar: [1 one]
Renkli gösterim: [33m1-one[0m [33m2-two[0m [33m3-three[0m
``
/*
📌 Terminalde eşleşmeler **sarı renkle** vurgulanacak.

---

### 2. Parametre ile direkt test
*/
``bash
go run regex_tester.go -pattern="go" -text="golang is cool"
``

``
Match sonucu: true
İlk eşleşme: go
Tüm eşleşmeler: [go]
Renkli gösterim: [33mgo[0mlang is cool
``
/*
---

🔑 Artık elinde **tam özellikli Regex Tester** var:

* İnteraktif kullanım
* Tek komutla test
* Renkli eşleşme vurgusu

---

👉 İstersen ben buna bir **history özelliği** ekleyebilirim (önceki regex ve string girişlerini kaydedip ok tuşlarıyla tekrar kullanma).
Bu şekilde VS Code veya terminalde regex’lerini tekrar tekrar denemek çok kolay olur.

Bunu da ister misin?
*/