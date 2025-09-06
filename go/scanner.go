/*
**Go’nun `text/scanner` paketini** en ince ayrıntısına kadar anlatayım.

---

# 📌 `text/scanner` Paketi Nedir?

`text/scanner` paketi, Go’da **metin (text) tarayıcı (scanner/lexer)** yazmak için kullanılan bir pakettir.

👉 Bu paket ile bir metin dosyası veya string üzerinde **karakter karakter ilerleyip, token’lara (kelimelere, sayılara, sembollere)** ayırabilirsin.

Yani temel olarak bir **lexical analyzer (lexer)** işlevi görür.

Örneğin:

* Kaynak kodu analiz etmek,
* Bir programlama dili için küçük bir parser yazmak,
* Konfigürasyon dosyalarını tokenize etmek için kullanılabilir.

---

# 📦 Paketin İçeriği

### 1. `Scanner` Yapısı

Asıl taramayı yapan yapıdır. İçinde taranacak metin, konum bilgileri, modlar ve hata yönetimi bulunur.

### 2. `Scanner.Init`

Bir `Scanner` başlatmak için kullanılır.

### 3. `Scanner.Scan`

Metindeki bir sonraki token’i döndürür.

### 4. `Scanner.TokenText`

Son bulunan token’in string karşılığını verir.

### 5. `Scanner.Position`

Son okunan token’in pozisyonunu (`line`, `column`, `offset`) verir.

### 6. `Scanner.Mode`

Tarama modunu belirler (ör: **identifier**, **string**, **comment** taranacak mı?).

### 7. `Error` ve `ErrorCount`

Hata durumlarını yönetmek için.

---

# 🔑 Önemli Sabitler (Token Türleri)

`text/scanner` paketinde `rune` tabanlı sabitler var:

* `scanner.Ident` → Identifier (değişken, fonksiyon adı)
* `scanner.Int` → Tamsayı
* `scanner.Float` → Ondalıklı sayı
* `scanner.String` → String literal (`"abc"`)
* `scanner.RawString` → Raw string (`` `abc` ``)
* `scanner.Comment` → Yorum satırı (`//` veya `/* */`)
* `scanner.EOF` → Dosya sonu

---

# 📝 Örnekler

## Örnek 1: Basit Tarama
*/
go
package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	var s scanner.Scanner

	// Kaynak metin
	src := `var x = 42
// Bu bir yorum
fmt.Println("Merhaba")`

	// Scanner’ı başlat
	s.Init(strings.NewReader(src))

	// Tokenleri tara
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %q\n", scanner.TokenString(tok), s.TokenText())
	}
}


//📌 Çıktı:


Ident: "var"
Ident: "x"
=: "="
Int: "42"
Comment: "// Bu bir yorum"
Ident: "fmt"
.: "."
Ident: "Println"
String: "\"Merhaba\""

/*
---

## Örnek 2: Pozisyon Bilgisi ile Tarama
*/
go
package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	var s scanner.Scanner
	src := "a = 123\nb = 45.6"

	s.Init(strings.NewReader(src))

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		pos := s.Pos() // Pozisyon bilgisi
		fmt.Printf("%s (%d:%d): %q\n", scanner.TokenString(tok), pos.Line, pos.Column, s.TokenText())
	}
}


//📌 Çıktı:


Ident (1:1): "a"
= (1:3): "="
Int (1:5): "123"
Ident (2:1): "b"
= (2:3): "="
Float (2:5): "45.6"

/*
---

## Örnek 3: Modları Ayarlama

Scanner’ın hangi tür token’ları algılayacağını `Mode` ile değiştirebiliriz.
*/
go
s.Mode = scanner.ScanIdents | scanner.ScanStrings | scanner.ScanInts

/*
* Sadece **identifiers, string ve integer**’ları algılar.
* Yorumlar veya float sayılar göz ardı edilir.

---

## Örnek 4: Hata Yönetimi

Kendi hata mesajını tanımlayabilirsin:
*/
go
s.Error = func(s *scanner.Scanner, msg string) {
	fmt.Printf("Hata (%s): %s\n", s.Pos(), msg)
}

/*
---

# 🚀 Özet

* `text/scanner` paketi, Go’da **metni token’lara ayırmaya** yarayan basit bir **lexer** paketidir.
* `Scanner.Init` → Başlatma
* `Scanner.Scan` → Sıradaki token’i bulma
* `Scanner.TokenText` → Token’in string karşılığı
* `Scanner.Pos` → Pozisyon bilgisi
* `Scanner.Mode` → Tarama modları (identifier, sayı, string, yorum vs.)
* `scanner.EOF` → Dosya sonunu belirtir

---

👉 İstersen ben sana bu paketi kullanarak **küçük bir matematiksel ifade parser’ı (ör: `2 + 3 * 4`) yazıp sonucu hesaplayan** örnek hazırlayayım mı?

Başka bir Açıkla kısmını

Çok güzel bir soru 👍 Şimdi sana **Go’nun `text/scanner` paketini** en ince ayrıntısına kadar açıklayayım.

---

# 📌 `text/scanner` Paketi Nedir?

`text/scanner` paketi, **metinleri (özellikle kaynak kod benzeri verileri) küçük token’lara (sözcüklere, sayılara, sembollere)** ayırmaya yarar.

👉 Yani bu paket bir **lexical scanner / tokenizer** işlevi görür.
👉 `go/scanner` paketinden farklıdır:

* `go/scanner` → Sadece Go kaynak kodunu tokenize eder.
* `text/scanner` → Genel amaçlıdır, istediğin herhangi bir metni tokenize edebilirsin.

---

# 📦 Paketin Temel Yapıları

## 1. `Scanner` tipi

Tokenizasyonu yapan ana tiptir.
*/
go
type Scanner struct {
    ...
}

/*
Metin üzerinde ilerler, `Scan` fonksiyonu ile sıradaki token’ı döner.

---

## 2. Önemli Metotlar

* **`Init(r io.Reader) *Scanner`**
  Bir `Scanner`’ı başlatır ve okuyacağı kaynağı (`string`, `file`, `bytes.Buffer`) verir.

* **`Scan() rune`**
  Sonraki token’ı döner (`Ident`, `Int`, `Float`, `String`, `Char`, `RawString`, `Comment` veya karakter).

* **`TokenText() string`**
  En son taranan token’ın metnini döner.

* **`Pos() Position`**
  Mevcut token’ın bulunduğu satır, sütun, offset bilgilerini döner.

* **`Error` ve `ErrorCount`**
  Hata durumlarını yakalamaya yarar.

---

## 3. Token Türleri

`text/scanner` aşağıdaki sabitleri kullanır:

* **`scanner.EOF`** → Dosya sonu
* **`scanner.Ident`** → Kimlikler (değişken, fonksiyon adı gibi)
* **`scanner.Int`** → Tamsayı
* **`scanner.Float`** → Ondalık sayı
* **`scanner.String`** → String (`"merhaba"`)
* **`scanner.RawString`** → Raw string (`` `merhaba` ``)
* **`scanner.Char`** → Tek karakter (`'a'`)
* **`scanner.Comment`** → Yorum (`//`, `/* */`)

---

# 📝 Örnekler

## Örnek 1: Basit Scanner Kullanımı
*/
go
package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	var s scanner.Scanner
	src := `var x = 42 // değişken tanımı`
	s.Init(strings.NewReader(src))

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("Token: %-10s | Text: %q\n", scanner.TokenString(tok), s.TokenText())
	}
}


//📌 Çıktı:


Token: Ident      | Text: "var"
Token: Ident      | Text: "x"
Token: =          | Text: "="
Token: Int        | Text: "42"
Token: Comment    | Text: "// değişken tanımı"

---

//## Örnek 2: Sayılar, Stringler ve Identler

go
package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	src := `name := "Ali" 
	age := 30 
	pi := 3.14`

	var s scanner.Scanner
	s.Init(strings.NewReader(src))

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%-8s -> %q\n", scanner.TokenString(tok), s.TokenText())
	}
}


//📌 Çıktı:


Ident    -> "name"
:=       -> ":="
String   -> "\"Ali\""
Ident    -> "age"
:=       -> ":="
Int      -> "30"
Ident    -> "pi"
:=       -> ":="
Float    -> "3.14"


---

//## Örnek 3: Satır ve Konum Bilgisi

go
package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	src := `x := 1
y := 2`

	var s scanner.Scanner
	s.Init(strings.NewReader(src))

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		pos := s.Pos()
		fmt.Printf("Line %d, Col %d: %s (%q)\n", pos.Line, pos.Column, scanner.TokenString(tok), s.TokenText())
	}
}


//📌 Çıktı:

``
Line 1, Col 1: Ident ("x")
Line 1, Col 3: := (":=")
Line 1, Col 6: Int ("1")
Line 2, Col 1: Ident ("y")
Line 2, Col 3: := (":=")
Line 2, Col 6: Int ("2")


---

//## Örnek 4: Hataları Yönetme

go
package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	var s scanner.Scanner
	s.Init(strings.NewReader("\"bozuk string"))

	// Hata callback’i
	s.Error = func(s *scanner.Scanner, msg string) {
		fmt.Printf("Hata: %s, Konum: %v\n", msg, s.Pos())
	}

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Println(scanner.TokenString(tok), s.TokenText())
	}
}


//📌 Çıktı:

`
Hata: literal not terminated, Konum: 1:14
``
/*
---

# 🚀 Özet

* `text/scanner` → Metinleri **token’lara ayıran** genel amaçlı bir tokenizer’dır.
* **`Scanner`** tipi üzerinden çalışır.
* **`Scan()`** → sıradaki token’ı döner.
* **`TokenText()`** → token’ın gerçek metnini verir.
* **Token türleri** → `Ident`, `Int`, `Float`, `String`, `Comment`, vb.
* **Kullanım alanları** → derleyici yazımı, özel betik dili parser’ı, statik analiz araçları.

---

👉 İstersen ben sana bu paketle **küçük bir basit hesap makinesi (lexer + evaluator)** örneği yazabilirim.
Yani `"3 + 5 * 2" gibi bir ifadeyi okuyup sonucu hesaplayalım. İster misin?
EVET
*/
/*
Harika 👌 Şimdi text/scanner paketini kullanarak **küçük bir hesap makinesi** (lexer + evaluator) yazalım.

Bu program:

1. Kullanıcıdan "3 + 5 * 2 - 4" gibi bir ifade alacak
2. scanner.Scanner ile token’lara ayıracak
3. Basit bir **çarpma/çıkartma/çarpma/bölme** önceliği olan hesap makinesi çalıştıracak
4. Sonucu ekrana yazacak

---

# 📂 Kod – Mini Hesap Makinesi
*/
`go
package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

// Parser yapısı
type Parser struct {
	s   scanner.Scanner
	tok rune
}

// Bir sonraki token’ı oku
func (p *Parser) next() {
	p.tok = p.s.Scan()
}

// İfade -> Term { (+|-) Term }
func (p *Parser) parseExpr() int {
	result := p.parseTerm()
	for p.tok == '+' || p.tok == '-' {
		op := p.tok
		p.next()
		val := p.parseTerm()
		if op == '+' {
			result += val
		} else {
			result -= val
		}
	}
	return result
}

// Term -> Factor { (*|/) Factor }
func (p *Parser) parseTerm() int {
	result := p.parseFactor()
	for p.tok == '*' || p.tok == '/' {
		op := p.tok
		p.next()
		val := p.parseFactor()
		if op == '*' {
			result *= val
		} else {
			result /= val
		}
	}
	return result
}

// Factor -> Number | (Expr)
func (p *Parser) parseFactor() int {
	switch p.tok {
	case scanner.Int:
		val := 0
		fmt.Sscanf(p.s.TokenText(), "%d", &val)
		p.next()
		return val
	case '(':
		p.next()
		val := p.parseExpr()
		if p.tok != ')' {
			panic(") eksik")
		}
		p.next()
		return val
	default:
		panic(fmt.Sprintf("Beklenmeyen token: %q", p.s.TokenText()))
	}
}

func eval(input string) int {
	var p Parser
	p.s.Init(strings.NewReader(input))
	p.s.Whitespace = 1<<'\t' | 1<<'\n' | 1<<' ' // boşlukları yok say
	p.next()
	result := p.parseExpr()
	if p.tok != scanner.EOF {
		panic("Fazladan token var")
	}
	return result
}

func main() {
	tests := []string{
		"3 + 5 * 2",
		"(3 + 5) * 2",
		"10 - 4 / 2",
		"8 / 2 * (2 + 2)",
	}

	for _, t := range tests {
		fmt.Printf("%s = %d\n", t, eval(t))
	}
}

/*
---

# 📌 Çıktı

Program çalıştırıldığında:
*/

``
3 + 5 * 2 = 13
(3 + 5) * 2 = 16
10 - 4 / 2 = 8
8 / 2 * (2 + 2) = 16
``
/*
---

# 🚀 Özet

* `scanner.Scanner` ile metni **token’lara ayırdık**
* Klasik **recursive descent parser** yazdık:

  * `Expr` → toplama/çıkarma
  * `Term` → çarpma/bölme
  * `Factor` → sayılar ve parantezler
* Böylece temel aritmetik ifadeleri çözümleyip hesapladık ✅

---

👉 İstersen ben bunu **geliştirip değişken atamaları (ör: `x = 5; x * 2`) ve fonksiyonlar (ör: `max(3,5)`)** destekleyen mini bir dil haline getirebilirim. İster misin?
*/