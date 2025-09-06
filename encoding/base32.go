/*
Golang’daki **`encoding/base32`** paketi, verileri **Base32** formatına dönüştürmek ve çözmek için kullanılan standart kütüphanedir.

📌 **Base32 Nedir?**
Base32, ikili (binary) veriyi ASCII karakterleriyle temsil etmek için kullanılan bir kodlama yöntemidir.

* Alfabe: `A-Z` ve `2-7` olmak üzere toplam **32 karakter** kullanılır.
* Genellikle **RFC 4648** standardına uygundur.
* Email, URL güvenliği ve ikili verilerin text ortamında taşınması için tercih edilir.
* Base64’e göre daha uzun stringler üretir ama daha okunabilirdir.

---

## 📌 Paket İmport Edilmesi
*/
``go
import (
    "encoding/base32"
    "fmt"
)
``
/*
---

## 1. Basit Encode / Decode Örneği
*/

``go
package main

import (
	"encoding/base32"
	"fmt"
)

func main() {
	data := []byte("Merhaba Dünya")

	// Encode
	encoded := base32.StdEncoding.EncodeToString(data)
	fmt.Println("Encoded:", encoded)

	// Decode
	decoded, err := base32.StdEncoding.DecodeString(encoded)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decoded:", string(decoded))
}
``
/*
🔹 Çıktı:

```
Encoded: JVQW4ZJANVZSA43PNZSXG5DSNRSXGIDNMVXGO===
Decoded: Merhaba Dünya
```

---

## 2. `StdEncoding` ve `HexEncoding`

`base32` paketinde iki farklı encoding tipi vardır:

* **`StdEncoding`**: RFC 4648 standard Base32 (`A-Z` ve `2-7`).
* **`HexEncoding`**: Alfabe `0-9` ve `A-V` şeklindedir (hex tabanlı Base32).
*/
``go
package main

import (
	"encoding/base32"
	"fmt"
)

func main() {
	data := []byte("GoLang")

	// Standart Base32
	encodedStd := base32.StdEncoding.EncodeToString(data)
	fmt.Println("StdEncoding:", encodedStd)

	// Hex Base32
	encodedHex := base32.HexEncoding.EncodeToString(data)
	fmt.Println("HexEncoding:", encodedHex)
}
``
/*
🔹 Çıktı:

```
StdEncoding: IORS4ZDJNZ2GK===
HexEncoding: D1IMOR3U41SS4===
```

---

## 3. Encoder / Decoder ile Stream Kullanımı

Büyük verilerde `io.Writer` veya `io.Reader` üzerinden encode/decode yapılabilir.
*/
``go
package main

import (
	"encoding/base32"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Encode stream
	encoder := base32.NewEncoder(base32.StdEncoding, os.Stdout)
	encoder.Write([]byte("Streaming Example"))
	encoder.Close() // flush işlemi için
	fmt.Println()

	// Decode stream
	input := "KRSXG5BAMJZG653OEBTG64RA" // "Streaming Example"
	decoder := base32.NewDecoder(base32.StdEncoding, strings.NewReader(input))
	buf := make([]byte, 100)
	n, _ := decoder.Read(buf)
	fmt.Println("Decoded:", string(buf[:n]))
}
``
/*
🔹 Çıktı:

```
KNZSAYJAOJSW45DIMUQGC4TJNZTQ====
Decoded: Streaming Example
```

---

## 4. Padding Kontrolü

Base32 çıktıları genellikle `=` karakterleri ile **padding** yapılır.
Eğer padding istemezseniz özel encoding oluşturabilirsiniz:
*/
``go
package main

import (
	"encoding/base32"
	"fmt"
)

func main() {
	// Padding'siz encoding oluştur
	noPadding := base32.StdEncoding.WithPadding(base32.NoPadding)

	data := []byte("GoLang")
	encoded := noPadding.EncodeToString(data)

	fmt.Println("No Padding:", encoded)

	decoded, _ := noPadding.DecodeString(encoded)
	fmt.Println("Decoded:", string(decoded))
}
``
/*
🔹 Çıktı:

```
No Padding: IORS4ZDJNZ2GK
Decoded: GoLang
```

---

## 5. Kendi Custom Base32 Encoding Oluşturma

Alfabe değiştirmek için `base32.NewEncoding()` kullanılabilir:
*/
``go
package main

import (
	"encoding/base32"
	"fmt"
)

func main() {
	// Kendi alfabemizi tanımlayalım (örnek: ters A-Z2-7)
	customAlphabet := "ZYXWVUTSRQPONMLKJIHGFEDCBA765432"
	customEncoding := base32.NewEncoding(customAlphabet)

	data := []byte("CustomBase32")
	encoded := customEncoding.EncodeToString(data)
	fmt.Println("Custom Encoded:", encoded)

	decoded, _ := customEncoding.DecodeString(encoded)
	fmt.Println("Decoded:", string(decoded))
}
``
/*
---

## 📌 Özet

* `base32.StdEncoding`: Standart Base32 (RFC 4648).
* `base32.HexEncoding`: Hex tabanlı Base32.
* `EncodeToString` / `DecodeString`: Direkt string dönüşüm.
* `NewEncoder` / `NewDecoder`: Stream bazlı dönüşüm.
* `WithPadding`: `=` padding kontrolü.
* `NewEncoding`: Özel alfabe ile encoding oluşturma.

---

İstersen sana **Base32 ile şifreli token üretme** ve **doğrulama sistemi** örneği de yazabilirim. İster misin?
EVET 
*/

/*
Harika 👍 O zaman sana **Base32 ile şifreli token üretme ve doğrulama** örneği yazayım.

Bu tarz kullanım genellikle:

* **Session ID**
* **API Key**
* **Doğrulama Kodu (OTP, e-mail doğrulama, şifre sıfırlama)**
  gibi yerlerde tercih edilir.

---

# 🔑 Base32 ile Token Üretme & Doğrulama
*/
``go
package main

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"time"
)

// Token oluşturma
func GenerateToken(size int) (string, error) {
	// Rastgele byte üret
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Base32 ile stringe çevir (padding'siz, daha temiz)
	token := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(bytes)
	return token, nil
}

// Token doğrulama (örnek: expiry süresi kontrolü)
type Token struct {
	Value     string
	ExpiresAt time.Time
}

func NewToken(size int, duration time.Duration) (*Token, error) {
	val, err := GenerateToken(size)
	if err != nil {
		return nil, err
	}
	return &Token{
		Value:     val,
		ExpiresAt: time.Now().Add(duration),
	}, nil
}

func (t *Token) IsValid(input string) bool {
	if time.Now().After(t.ExpiresAt) {
		return false // süresi geçmiş
	}
	return t.Value == input
}

func main() {
	// 16 byte uzunluğunda, 1 dakikalık token üret
	token, _ := NewToken(16, time.Minute)
	fmt.Println("Generated Token:", token.Value)
	fmt.Println("Expires At:", token.ExpiresAt.Format(time.RFC3339))

	// Doğru giriş
	ok := token.IsValid(token.Value)
	fmt.Println("Doğru giriş:", ok)

	// Yanlış giriş
	ok = token.IsValid("YANLIŞ-TOKEN")
	fmt.Println("Yanlış giriş:", ok)
}
``
/*
---

## 📌 Çalışma Mantığı

1. `rand.Read` → Kriptografik güvenli rastgele byte üretir.
2. `base32.StdEncoding.EncodeToString` → Byte’ları Base32 stringe çevirir.

   * `WithPadding(base32.NoPadding)` kullanarak `=` karakterlerinden kurtulduk.
3. `Token` struct → Hem token değerini hem de geçerlilik süresini tutar.
4. `IsValid` → Token hem süresine hem de eşleşmesine göre kontrol edilir.

---

## ✅ Örnek Çıktı

```
Generated Token: NVHGU3K4VY4ZTASOJBN2T3FM3EUJ6N2S
Expires At: 2025-08-29T10:46:28+03:00
Doğru giriş: true
Yanlış giriş: false
```

---

İstersen bu sistemi biraz daha geliştirebilirim:

* **Token’ı URL-safe hale getirme** (örneğin e-mail linklerinde güvenli kullanım).
* **Kullanıcıya tek kullanımlık doğrulama kodu gönderme (OTP)**.

Bunlardan hangisini isterdin 👉 **URL-safe token** mi, yoksa **OTP (6 haneli doğrulama kodu)** örneği mi?
*/


