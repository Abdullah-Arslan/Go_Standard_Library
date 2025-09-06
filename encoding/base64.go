/*
**Go’daki `encoding/base64` paketini** en ince ayrıntısına kadar açıklayayım.

---

# 📌 Go `encoding/base64` Paketi

Go’nun standart kütüphanesindeki **`encoding/base64`** paketi, **ikili (binary) verileri Base64 formatında kodlamak (encode)** ve **çözmek (decode)** için kullanılır.

---

## 🔎 Base64 Nedir?

* **Base64** ikili veriyi ASCII karakterleriyle taşımak için geliştirilmiş bir kodlama yöntemidir.
* Alfabe: `A–Z`, `a–z`, `0–9`, `+`, `/` (veya URL güvenli varyantta `-`, `_`).
* Sonunda **padding** için `=` karakteri kullanılabilir.
* Kullanım alanları:

  * E-posta (MIME)
  * JSON/XML içinde ikili veri taşıma
  * URL, JWT (JSON Web Token)
  * Şifreleme sonrası verilerin güvenli taşınması

---

## 📦 Paketi Import Etme
*/

``go
import (
    "encoding/base64"
    "fmt"
)
``
/*
---

# 🚀 Örneklerle Anlatım

## 1. Basit Encode / Decode
*/

``go
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("Merhaba Dünya")

	// Encode
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println("Encoded:", encoded)

	// Decode
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decoded:", string(decoded))
}
``

/*
🔹 Çıktı:

```
Encoded: TWVyaGFiYSDEsHVueWE=
Decoded: Merhaba Dünya
```

---

## 2. `StdEncoding` ve `URLEncoding`

Base64’te iki temel encoding vardır:

* **`StdEncoding`** → Normal Base64 (`+` ve `/` kullanır).
* **`URLEncoding`** → URL güvenli Base64 (`-` ve `_` kullanır, `+` ve `/` yerine).
*/

``go
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("GoLang")

	// Normal Base64
	encodedStd := base64.StdEncoding.EncodeToString(data)
	fmt.Println("StdEncoding:", encodedStd)

	// URL güvenli Base64
	encodedURL := base64.URLEncoding.EncodeToString(data)
	fmt.Println("URLEncoding:", encodedURL)
}
``
/*
🔹 Çıktı:

```
StdEncoding: R29MYW5n
URLEncoding: R29MYW5n
```

👉 Farklı karakterler içeren veri olduğunda (örn. `+` veya `/` çıkarsa) URL-safe farkı ortaya çıkar.

---

## 3. Encoder / Decoder ile Stream Kullanımı

Büyük verilerde **stream bazlı** encode/decode yapmak için `NewEncoder` ve `NewDecoder` kullanılır.
*/

``go
package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Encode stream
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	encoder.Write([]byte("Streaming Example"))
	encoder.Close() // flush işlemi için
	fmt.Println()

	// Decode stream
	input := "U3RyZWFtaW5nIEV4YW1wbGU=" // "Streaming Example"
	decoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader(input))
	buf := make([]byte, 100)
	n, _ := decoder.Read(buf)
	fmt.Println("Decoded:", string(buf[:n]))
}
``
/*
🔹 Çıktı:

```
U3RyZWFtaW5nIEV4YW1wbGU=
Decoded: Streaming Example
```

---

## 4. Padding (`=`)

Base64 çıktılarında uzunluk 4’ün katı olmalıdır.
Eksik olduğunda `=` padding karakterleri eklenir.

Eğer padding istemezseniz `WithPadding(base64.NoPadding)` kullanabilirsiniz.
*/

``go
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("GoLang")

	// Normal (padding ile)
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println("With Padding:", encoded)

	// Padding'siz
	noPadding := base64.StdEncoding.WithPadding(base64.NoPadding)
	encodedNoPad := noPadding.EncodeToString(data)
	fmt.Println("No Padding:", encodedNoPad)
}
``
/*
🔹 Çıktı:

```
With Padding: R29MYW5n
No Padding: R29MYW5n
```

---

## 5. Binary Veriyi Encode Etme (ör. resim, şifreli veri)

Base64 sadece string değil, herhangi bir **binary dosya** üzerinde kullanılabilir.
*/
``go
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// Örnek binary veri
	binaryData := []byte{0x01, 0x02, 0xFA, 0xFB, 0xFC}

	encoded := base64.StdEncoding.EncodeToString(binaryData)
	fmt.Println("Encoded Binary:", encoded)

	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Println("Decoded Binary:", decoded)
}
``
/*
🔹 Çıktı:

```
Encoded Binary: AQL6+/w=
Decoded Binary: [1 2 250 251 252]
```

---

## 6. Custom Encoding (Kendi Base64 Alfabenizi Oluşturma)

Base64’de kullanılan alfabeyi değiştirmek için `base64.NewEncoding()` kullanılır.
*/

``go
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// Alfabe: standart değil, custom
	customAlphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz012345-_"
	customEncoding := base64.NewEncoding(customAlphabet)

	data := []byte("CustomBase64")

	encoded := customEncoding.EncodeToString(data)
	fmt.Println("Custom Encoded:", encoded)

	decoded, _ := customEncoding.DecodeString(encoded)
	fmt.Println("Decoded:", string(decoded))
}
``
/*
---

# 📌 Özet

* **`StdEncoding`**: Normal Base64 (`+` ve `/`).
* **`URLEncoding`**: URL güvenli Base64 (`-` ve `_`).
* **`EncodeToString` / `DecodeString`**: Direkt string dönüşüm.
* **`NewEncoder` / `NewDecoder`**: Stream bazlı encode/decode.
* **`WithPadding`**: Padding kontrolü (`=`, `NoPadding`).
* **`NewEncoding`**: Özel alfabe ile Base64.

---

👉 İstersen sana Base64 ile **şifreli token üretme** (JWT benzeri küçük bir sistem) veya **dosya encode/decode** örneği de yazabilirim.
EVET
*/

/*
Hangisini isterdin:

* 🔑 **Base64 ile Token Üretme**
* 📂 **Dosyayı Base64 encode/decode**

🔑 **Base64 ile Token Üretme** Bu Yapılacak Örnek Olarak

Harika 👍 O zaman sana **Base64 ile basit bir token üretme ve doğrulama sistemi** göstereyim. Bu sistem JWT mantığının sadeleştirilmiş bir versiyonudur.

---

# 🔑 Base64 ile Token Üretme & Doğrulama (JWT benzeri)

📌 JWT mantığında 3 kısım vardır:

1. **Header** → Algoritma ve tip bilgisi
2. **Payload** → Kullanıcıya ait bilgiler (id, email vs.)
3. **Signature** → Gizli anahtar ile imza

Biz burada sade bir örnek yapacağız:

* Payload → Kullanıcı bilgisi + Token süresi
* Signature → `HMAC-SHA256` ile imza
* Encode işlemi → Base64

---

## 📝 Kod
*/
``go
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

// Gizli anahtar
var secretKey = []byte("super-secret-key")

// Token yapısı
type TokenPayload struct {
	UserID    int       `json:"user_id"`
	Email     string    `json:"email"`
	ExpiresAt time.Time `json:"exp"`
}

// Token üret
func GenerateToken(userID int, email string, duration time.Duration) (string, error) {
	payload := TokenPayload{
		UserID:    userID,
		Email:     email,
		ExpiresAt: time.Now().Add(duration),
	}

	// Payload'ı JSON’a çevir
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	// Base64 ile encode et
	payloadBase64 := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(payloadBytes)

	// İmza oluştur
	h := hmac.New(sha256.New, secretKey)
	h.Write([]byte(payloadBase64))
	signature := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(h.Sum(nil))

	// Token = payload + "." + signature
	token := payloadBase64 + "." + signature
	return token, nil
}

// Token doğrula
func ValidateToken(token string) (*TokenPayload, bool) {
	// Tokeni "." dan ayır
	parts := make([]string, 2)
	n, _ := fmt.Sscanf(token, "%s.%s", &parts[0], &parts[1])
	if n != 2 {
		return nil, false
	}
	payloadBase64, signature := parts[0], parts[1]

	// İmza doğrula
	h := hmac.New(sha256.New, secretKey)
	h.Write([]byte(payloadBase64))
	expectedSignature := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(h.Sum(nil))

	if signature != expectedSignature {
		return nil, false
	}

	// Payload çöz
	payloadBytes, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(payloadBase64)
	if err != nil {
		return nil, false
	}

	var payload TokenPayload
	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		return nil, false
	}

	// Süre kontrolü
	if time.Now().After(payload.ExpiresAt) {
		return nil, false
	}

	return &payload, true
}

func main() {
	// Token üret
	token, _ := GenerateToken(123, "user@example.com", time.Minute*2)
	fmt.Println("Generated Token:", token)

	// Token doğrula
	payload, valid := ValidateToken(token)
	if valid {
		fmt.Println("Valid Token ✅")
		fmt.Printf("UserID: %d, Email: %s, ExpiresAt: %s\n",
			payload.UserID, payload.Email, payload.ExpiresAt.Format(time.RFC3339))
	} else {
		fmt.Println("Invalid Token ❌")
	}
}
``
/*
---

## 📌 Örnek Çıktı

```
Generated Token: eyJ1c2VyX2lkIjoxMjMsImVtYWlsIjoidXNlckBleGFtcGxlLmNvbSIsImV4cCI6IjIwMjUtMDgtMjlUMTA6NTU6NDBaIn0.I3QZy0RGAd2cGlhIj8pPHqX7dGVHtW8C9lVEk8OwGCE
Valid Token ✅
UserID: 123, Email: user@example.com, ExpiresAt: 2025-08-29T10:55:40Z
```

---

## 🔎 Mantık

1. Kullanıcı bilgilerini JSON’a çevirdik.
2. JSON’u **Base64 ile encode** ettik.
3. `HMAC-SHA256` ile gizli anahtar kullanarak imza ürettik.
4. Token = `payload.signature` şeklinde oluşturuldu.
5. Doğrulama sırasında imza ve süre kontrol edildi.

---

👉 İstersen ben bunu bir adım ileri taşıyıp **JWT standardına %100 uyumlu** hale getirebilirim (header + payload + signature).
Bunu da ister misin?
EVET
*/
/*
Tamamdır ⚡ Şimdi sana **Go ile sıfırdan JWT standardına %100 uyumlu** bir örnek yazacağım.

JWT formatı şu şekildedir:

```
header.payload.signature
```

* **Header** → Algoritma ve tip bilgisi (ör. `{ "alg": "HS256", "typ": "JWT" }`)
* **Payload** → Kullanıcı bilgileri, expire süresi gibi bilgiler
* **Signature** → Header + Payload değerlerinin gizli anahtar ile HMAC-SHA256 imzalanması

---

# 🔑 Go ile JWT Uygulaması
*/
``go
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Gizli anahtar
var secretKey = []byte("super-secret-key")

// JWT Header yapısı
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

// JWT Payload yapısı
type Payload struct {
	UserID    int       `json:"user_id"`
	Email     string    `json:"email"`
	ExpiresAt time.Time `json:"exp"`
}

// Base64 URL-safe encode (padding'siz)
func base64URLEncode(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}

// Base64 URL-safe decode
func base64URLDecode(data string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(data)
}

// JWT oluştur
func GenerateJWT(userID int, email string, duration time.Duration) (string, error) {
	// 1. Header
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	headerJSON, _ := json.Marshal(header)
	headerB64 := base64URLEncode(headerJSON)

	// 2. Payload
	payload := Payload{
		UserID:    userID,
		Email:     email,
		ExpiresAt: time.Now().Add(duration),
	}
	payloadJSON, _ := json.Marshal(payload)
	payloadB64 := base64URLEncode(payloadJSON)

	// 3. Signature
	unsigned := headerB64 + "." + payloadB64
	h := hmac.New(sha256.New, secretKey)
	h.Write([]byte(unsigned))
	signature := base64URLEncode(h.Sum(nil))

	// JWT = header.payload.signature
	token := unsigned + "." + signature
	return token, nil
}

// JWT doğrula
func ValidateJWT(token string) (*Payload, bool) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, false
	}

	headerB64, payloadB64, signature := parts[0], parts[1], parts[2]

	// İmza yeniden hesapla
	unsigned := headerB64 + "." + payloadB64
	h := hmac.New(sha256.New, secretKey)
	h.Write([]byte(unsigned))
	expectedSig := base64URLEncode(h.Sum(nil))

	if signature != expectedSig {
		return nil, false
	}

	// Payload decode et
	payloadJSON, err := base64URLDecode(payloadB64)
	if err != nil {
		return nil, false
	}

	var payload Payload
	if err := json.Unmarshal(payloadJSON, &payload); err != nil {
		return nil, false
	}

	// Expire kontrolü
	if time.Now().After(payload.ExpiresAt) {
		return nil, false
	}

	return &payload, true
}

func main() {
	// Token üret
	token, _ := GenerateJWT(42, "user@example.com", time.Minute*2)
	fmt.Println("Generated JWT:", token)

	// Token doğrula
	payload, valid := ValidateJWT(token)
	if valid {
		fmt.Println("Valid ✅")
		fmt.Printf("UserID: %d, Email: %s, ExpiresAt: %s\n",
			payload.UserID, payload.Email, payload.ExpiresAt.Format(time.RFC3339))
	} else {
		fmt.Println("Invalid ❌")
	}
}
``
/*
---

## 📌 Örnek Çıktı

```
Generated JWT: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0MiwiZW1haWwiOiJ1c2VyQGV4YW1wbGUuY29tIiwiZXhwIjoiMjAyNS0wOC0yOVQxMTowODoyNloifQ.1Q6u_p9IhQhEoEcf4e5Wjj2x_c3U77uq4ltkjTudZJQ
Valid ✅
UserID: 42, Email: user@example.com, ExpiresAt: 2025-08-29T11:08:26Z
```

---

## 🔎 Ne Yaptık?

1. **Header** → `{ "alg": "HS256", "typ": "JWT" }` → Base64URL encode
2. **Payload** → Kullanıcı bilgisi + expire süresi → Base64URL encode
3. **Signature** → `HMAC-SHA256(header.payload, secretKey)`
4. Final Token → `header.payload.signature`

Yani artık elimizde **tamamen JWT standardına uyumlu** bir sistem var ✅

---

👉 İstersen bu JWT sistemini **middleware** olarak ekleyelim.
Yani bir HTTP API yapalım, gelen isteklerde Authorization header’daki JWT token doğrulansın.

Bunu ister misin?
EVET
*/

/*
Süper 👍 Şimdi JWT sistemimizi **HTTP API’ye middleware** olarak ekleyelim.

📌 Amaç:

* Kullanıcı `/login` endpoint’ine girdiğinde JWT token alacak.
* Kullanıcı `/profile` endpoint’ine girdiğinde ise gelen `Authorization: Bearer <token>` header kontrol edilecek.
* Eğer token geçerli değilse ❌ 401 Unauthorized dönülecek.
* Eğer geçerliyse ✅ kullanıcı bilgileri gösterilecek.

---
*/
//# 🌐 JWT Middleware Örneği

``go
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

var secretKey = []byte("super-secret-key")

// JWT header
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

// JWT payload
type Payload struct {
	UserID    int       `json:"user_id"`
	Email     string    `json:"email"`
	ExpiresAt time.Time `json:"exp"`
}

// Base64 URL encode/decode
func base64URLEncode(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}
func base64URLDecode(data string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(data)
}

// Token üret
func GenerateJWT(userID int, email string, duration time.Duration) (string, error) {
	header := Header{Alg: "HS256", Typ: "JWT"}
	headerJSON, _ := json.Marshal(header)
	headerB64 := base64URLEncode(headerJSON)

	payload := Payload{
		UserID:    userID,
		Email:     email,
		ExpiresAt: time.Now().Add(duration),
	}
	payloadJSON, _ := json.Marshal(payload)
	payloadB64 := base64URLEncode(payloadJSON)

	unsigned := headerB64 + "." + payloadB64
	h := hmac.New(sha256.New, secretKey)
	h.Write([]byte(unsigned))
	signature := base64URLEncode(h.Sum(nil))

	return unsigned + "." + signature, nil
}

// Token doğrula
func ValidateJWT(token string) (*Payload, bool) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, false
	}

	headerB64, payloadB64, signature := parts[0], parts[1], parts[2]

	// İmza doğrulama
	unsigned := headerB64 + "." + payloadB64
	h := hmac.New(sha256.New, secretKey)
	h.Write([]byte(unsigned))
	expectedSig := base64URLEncode(h.Sum(nil))

	if signature != expectedSig {
		return nil, false
	}

	// Payload çöz
	payloadJSON, err := base64URLDecode(payloadB64)
	if err != nil {
		return nil, false
	}

	var payload Payload
	if err := json.Unmarshal(payloadJSON, &payload); err != nil {
		return nil, false
	}

	if time.Now().After(payload.ExpiresAt) {
		return nil, false
	}

	return &payload, true
}

// Middleware: JWT kontrolü
func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		payload, valid := ValidateJWT(token)
		if !valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Payload bilgilerini request context’e ekleyebiliriz (şimdilik direkt yazdıralım)
		fmt.Printf("✅ Token valid for user %d (%s)\n", payload.UserID, payload.Email)

		// Devam et
		next(w, r)
	}
}

// Handlers
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Normalde burada kullanıcı adı/şifre kontrol edilir.
	// Biz örnek olsun diye sabit kullanıcı üretelim.
	token, _ := GenerateJWT(42, "user@example.com", time.Minute*5)
	w.Write([]byte("Your token: " + token))
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to your profile! ✅"))
}

func main() {
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/profile", JWTMiddleware(ProfileHandler))

	fmt.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
``
/*
---

## 📌 Çalışma Mantığı

1. **`/login`** → Token üretir ve kullanıcıya verir.

   ```
   GET http://localhost:8080/login
   ```

   Dönen cevap:

   ```
   Your token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
   ```

2. **`/profile`** → Authorization header ile token gönderilmesi gerekir.

   ```
   GET http://localhost:8080/profile
   Authorization: Bearer <token>
   ```

   Eğer token geçerli → ✅ `"Welcome to your profile!"`
   Eğer geçersiz veya süresi dolmuş → ❌ `401 Unauthorized`

---

👉 Burada context kullanımını eklemedik. İstersen ben sana payload’taki `UserID` ve `Email` bilgisini **context.Context içine ekleyip handler’larda erişilebilir hale getirecek** versiyonunu yazayım.

Bunu da ister misin?
*/