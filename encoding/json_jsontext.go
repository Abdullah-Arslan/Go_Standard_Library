/*
Go dilindeki `encoding/json/jsontext` paketine dair en güncel ve kapsamlı bilgiyi şimdi seninle paylaşacağım — hem teorik altyapısını hem de örneklerle kullanımını detaylıca anlatacağım.

---

## `encoding/json/jsontext` Paketi Nedir?

* `jsontext`, Go’nun deneysel bir parçası olan **sentaks düzeyinde JSON işlemleri** sunan bir pakettir.
* Yani **JSON metnini (syntax)** doğrudan işlemek ve dönüştürmek için tasarlanmıştır; veriyi Go struct’larına veya tiplere dönüştürmek (semantic layer) yerine **ham JSON yapısını** kontrol etmeye odaklanır. ([Go Packages][1], [golang.bg][2])
* Bu paket, Go’nun `encoding/json/v2` (yeni versiyon) sisteminin **alt katmanını** oluşturur ve asıl `json` paketine göre daha az bağımlılıkla çalışır, hatta yansıma (`reflect`) kullanmaz — bu sayede daha **compact binary boyutu** ve **düşük bağımlılık** sunar ([GitHub][3]).
* Unutmamak gerek: Bu paket **experimental (deneyseldir)**. Go 1 uyumluluk garantisine dahil değildir ve kullanılabilmesi için `GOEXPERIMENT=jsonv2` ortam değişkeninin aktif olması gerekir ([Go Packages][4]).

---

## Temel Yapılar ve Kavramlar

### 1. `Token` ve `Value`

* **`Token`**: JSON’daki temel yapı taşlarını temsil eder: `null`, `true`/`false`, string, number, `{`, `}`, `[`, `]` gibi.
* **`Value`**: Bir JSON değerini (literal, string, number, nesne veya dizi) ham `[]byte` formatında tutar. Örneğin `{...}` veya `[1,2,3]` gibi toplu değerler `Value` olarak alınabilir ([GitHub][3]).

### 2. `Decoder` ve `Encoder`

* **`Decoder`**: `NewDecoder(io.Reader, ...Options)` ile oluşturulur.

  * `ReadToken()` ile sıradaki JSON bileşenini `Token` olarak okur.
  * `ReadValue()` ile tamamını `Value` olarak alabilir.
  * `PeekKind()` ile okunacak tokenın türü önceden tespit edilebilir.
  * `SkipValue()` ile istemediğin bir JSON değerini atlayabilirsin ([GitHub][3]).
* **`Encoder`**: `NewEncoder(io.Writer, ...Options)` ile oluşturulur.

  * `WriteToken(Token)` ya da `WriteValue(Value)` ile JSON yapılarını yazmanı sağlar ([GitHub][3]).

### 3. `Kind` Türü

* JSON için token ya da value türünü temsil eden bir enum’dur. Örneğin:

  * `'n'` → null
  * `'t'` → true
  * `'"'` → string
  * `'{'`, `'}'` → object begin/end
  * `'['`, `']'` → array begin/end ([GitHub][3]).

---

## Neden `jsontext` Kullanılır?

1. **Minimum binary boyutu** (özellikle `reflect` kullanılmaz).
2. JSON’u Go struct’ına dönüştürmeden **syntax seviyesinde işlemek** (örneğin, JSON’u dönüştürme veya filtreleme).
3. TinyGo, WASI gibi ortamlar için uygun — hafif bağımlılık.

Bu paket, `encoding/json/v2` içindeki semantic işlemlerin (marshal/unmarshal) alt altyapısını oluşturur ([GitHub][3]).

---

## Örnek Kullanım: String Değiştiren Basit Processor

Aşağıdaki örnekte, JSON metnindeki `"Golang"` kelimesini `"Go"` ile değiştiriyoruz:
*/
``go
package main

import (
    "bytes"
    "fmt"
    "io"
    "strings"

    "encoding/json/jsontext"
)

func main() {
    const input = `{
        "title": "Golang version 1 is released",
        "text": "This is about Golang!"
    }`

    in := strings.NewReader(input)
    dec := jsontext.NewDecoder(in)
    out := new(bytes.Buffer)
    enc := jsontext.NewEncoder(out, jsontext.Multiline(true)) // daha okunabilir
    for {
        tok, err := dec.ReadToken()
        if err == io.EOF {
            break
        }
        if err != nil {
            panic(err)
        }
        if tok.Kind() == jsontext.KindString {
            if string(tok.String()) == "Golang version 1 is released" {
                tok = jsontext.NewStringToken("Go version 1 is released")
            }
        }
        enc.WriteToken(tok)
    }
    fmt.Println(out.String())
}
``
/*
**Ne yapıyor?**

* JSON’u token token okuyup,
* Sadece string token’leri kontrol ediyor,
* Belirlediğiniz string’e uygun token’ı değiştiriyor,
* Ardından yeni stream’e yazıyor.

Bu tarz işlemler traditional `encoding/json` paketinin sağlayamadığı **sentaks seviyesinde dönüşümleri** mümkün kılar.

---

## Özet Tablosu

| Katman              | Paket / Bağlantı          | Açıklama                                                     |
| ------------------- | ------------------------- | ------------------------------------------------------------ |
| Söz Dizimi (syntax) | `jsontext`                | JSON’u token/value olarak okuyup yazmaya odaklı, yansımasız. |
| Anlam (semantic)    | `encoding/json/v2` (json) | Go değerleriyle çalışır; reflect ile marshal/unmarshal.      |

---

## Kapanış ve Tavsiye

* `jsontext`, Go’da JSON’u ham haliyle işlemek isteyen ileri düzey kullanıcılar için mükemmel bir araç.
* Ancak henüz deneyseldir ve Go’nun resmi 1.x sürümlerinde **varsayılan kullanılacak paket değildir**.
* Eğer JSON verisini Go struct’ına dönüştürme (marshal/unmarshal) ihtiyacın varsa, hâlâ `encoding/json` ya da `encoding/json/v2` paketlerini kullanman daha doğru olur.

[1]: https://pkg.go.dev/github.com/go-json-experiment/json/jsontext?utm_source=chatgpt.com "jsontext package - github.com/go-json-experiment/json/jsontext"
[2]: https://golang.bg/pkg/encoding/json/jsontext/?utm_source=chatgpt.com "jsontext - Go Documentation Server"
[3]: https://github.com/golang/go/discussions/63397?utm_source=chatgpt.com "encoding/json/v2 · golang go · Discussion #63397"
[4]: https://pkg.go.dev/encoding/json/jsontext?utm_source=chatgpt.com "encoding/json/jsontext - Go Packages"

İstersen bu konu üzerine daha da derinlemesine gidebiliriz—örneğin bir kullan-case geliştirip adım adım çalışabiliriz ya da diğer deneysel JSON paketlerine (örneğin `jsonseq`) göz atabiliriz. Devam edelim mi?
EVET
*/

/*
Harika! 👍 O zaman bir adım öteye geçelim ve **`jsontext` paketi ile gerçek bir kullanım senaryosu** oluşturalım.

Senaryo: **JSON log dosyasında bazı alanları filtreleme ve değiştirme**

* Büyük bir JSON log dizisi var (`[{...},{...},...]`)
* Bazı kullanıcı alanlarını maskeliyoruz (`username`)
* Sadece token seviyesinde çalışıyoruz, Go struct’ı kullanmıyoruz → performans kazanıyoruz.

---

## 📝 Örnek Kod: JSON Log Maskleme
*/
``go
package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"encoding/json/jsontext"
)

func main() {
	// Örnek JSON log
	const logs = `[
		{"username":"alice","action":"login"},
		{"username":"bob","action":"logout"},
		{"username":"charlie","action":"login"}
	]`

	in := strings.NewReader(logs)
	dec := jsontext.NewDecoder(in)
	out := new(bytes.Buffer)
	enc := jsontext.NewEncoder(out, jsontext.Multiline(true)) // okunabilir output

	for {
		tok, err := dec.ReadToken()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		// username alanını maskleme
		if tok.Kind() == jsontext.KindString {
			if string(tok.String()) == "alice" || string(tok.String()) == "bob" || string(tok.String()) == "charlie" {
				tok = jsontext.NewStringToken("*****")
			}
		}

		enc.WriteToken(tok)
	}

	fmt.Println(out.String())
}
``
/*
---

## 🔹 Çıktı
*/
``json
[
  {
    "username": "*****",
    "action": "login"
  },
  {
    "username": "*****",
    "action": "logout"
  },
  {
    "username": "*****",
    "action": "login"
  }
]
``
/*
---

### 📌 Açıklamalar

1. `jsontext.NewDecoder` → JSON’u token token okur, büyük JSON dizileri için hafıza dostudur.
2. `tok.Kind()` ile her tokenın türünü (string, number, object start/end, array start/end) kontrol ediyoruz.
3. `jsontext.NewStringToken` → string token’ı değiştirmek için kullanılır.
4. `jsontext.NewEncoder` → değişen tokenları tekrar JSON stream olarak yazar.
5. Go struct’ına dönüştürmeye gerek yok → **daha hızlı ve hafızada az yer kaplar**.

---

### ✅ Avantajları

* JSON büyük boyutlu olsa bile streaming olarak işleyebiliriz.
* Performans kritik sistemlerde `reflect` kullanılmaz → hızlıdır.
* Token seviyesinde kontrol → sadece gerekli alanları değiştirir veya atlar.
* `SkipValue()` ile istemediğiniz alanları tamamen geçebilirsiniz.

---

İstersen bir sonraki adımda sana **`jsontext` ile JSON içindeki belirli alanları tamamen atlayarak filtreleme** örneğini de gösterebilirim. Bu, log masking ve güvenlik filtrelerinde çok kullanışlıdır. Bunu da ekleyeyim mi?
*/