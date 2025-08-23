/*
Go (Golang) `crypto/hmac` paketini** en baştan, tüm fonksiyonlarını ve kullanım senaryolarını örneklerle açıklayayım.

---

# 🔑 `crypto/hmac` Paketi Nedir?

Go’daki [`crypto/hmac`](https://pkg.go.dev/crypto/hmac) paketi, **HMAC (Hash-based Message Authentication Code)** algoritmasının uygulanmasını sağlar.

HMAC, bir mesajın bütünlüğünü ve kimliğini doğrulamak için bir **hash fonksiyonu (örn: SHA256, SHA512)** ve **gizli anahtar (secret key)** kullanır.
Amaç:

* Mesajın değiştirilmediğini garanti etmek ✅
* Mesajın gerçekten gizli anahtarı bilen biri tarafından üretildiğini ispatlamak ✅

---

# 📦 Paketin İçeriği

`hmac` paketinde çok fazla fonksiyon yoktur. Temel olarak şu özellikleri sağlar:

1. **`func New(h func() hash.Hash, key []byte) hash.Hash`**
   → Yeni bir HMAC nesnesi oluşturur.

2. **`func Equal(mac1, mac2 []byte) bool`**
   → İki MAC değerini **zaman sabitliğinde** (constant-time) karşılaştırır.
   (Böylece timing attack engellenir.)

---

# 🔧 Kullanım Örnekleri

### 1. Basit HMAC oluşturma (SHA256 ile)
*/
``go
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main() {
	// Gizli anahtar (secret key)
	key := []byte("mysecretkey")

	// Mesaj
	message := []byte("Merhaba Dünya")

	// HMAC oluştur (SHA256 kullanarak)
	h := hmac.New(sha256.New, key)
	h.Write(message)

	// MAC değeri (hash çıktısı)
	mac := h.Sum(nil)

	fmt.Printf("HMAC (SHA256): %x\n", mac)
}
``
/*
📌 Çıktı (her zaman aynı mesaj ve key için aynı olur):

```
HMAC (SHA256): e2f7aebc1f9f54e4b10b168dc39d6be37d05f1d7cd7cf6edb7e94a2dfc56e79b
```

---

### 2. HMAC doğrulama

Mesajı alan taraf, gelen mesaj + MAC değerini tekrar hesaplayarak doğrular.
*/
``go
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main() {
	key := []byte("mysecretkey")
	message := []byte("Merhaba Dünya")

	// Gönderilen tarafın ürettiği HMAC
	h := hmac.New(sha256.New, key)
	h.Write(message)
	expectedMAC := h.Sum(nil)

	// Alıcı tarafın hesapladığı HMAC
	h2 := hmac.New(sha256.New, key)
	h2.Write(message)
	calculatedMAC := h2.Sum(nil)

	// Güvenli karşılaştırma
	if hmac.Equal(expectedMAC, calculatedMAC) {
		fmt.Println("✅ Mesaj doğrulandı, değiştirilmemiş.")
	} else {
		fmt.Println("❌ Mesaj veya MAC yanlış!")
	}
}
``
/*
---

### 3. Yanlış anahtar kullanıldığında
*/
``go
wrongKey := []byte("wrongkey")
h := hmac.New(sha256.New, wrongKey)
h.Write(message)
fakeMAC := h.Sum(nil)

if hmac.Equal(expectedMAC, fakeMAC) {
    fmt.Println("Doğrulama başarılı (yanlışlıkla)!")
} else {
    fmt.Println("Doğrulama başarısız (anahtar yanlış).")
}
``
/*
---
Çıktı:

``
Doğrulama başarısız (anahtar yanlış).
``


### 4. Farklı Hash Algoritmalarıyla HMAC

`hmac.New` fonksiyonuna farklı hash fonksiyonları verilebilir:
*/
``go
import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha512"
)
``

//Kullanım:

``go
hmac.New(md5.New, key)
hmac.New(sha1.New, key)
hmac.New(sha512.New, key)
``
/*
---

### 5. HMAC’i Fonksiyonlaştırma

Genelde pratik kullanım için HMAC üretimi fonksiyon haline getirilir:
*/
``go
func ComputeHMAC(message, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(message)
	return h.Sum(nil)
}

func CheckHMAC(message, key, mac []byte) bool {
	expected := ComputeHMAC(message, key)
	return hmac.Equal(mac, expected)
}
``

//Kullanımı:

``go
mac := ComputeHMAC([]byte("Hello"), []byte("secret"))
ok := CheckHMAC([]byte("Hello"), []byte("secret"), mac)
fmt.Println("Doğrulama:", ok) // true
``
/*
---

# 📌 Özet

* `crypto/hmac` → HMAC algoritmasını uygular.
* `New(hashFunc, key)` → HMAC hesaplama için kullanılır.
* `Equal(mac1, mac2)` → Güvenli karşılaştırma yapar.
* SHA256, SHA1, SHA512 gibi farklı hash fonksiyonlarıyla kullanılabilir.
* Mesajların bütünlüğünü ve kimliğini doğrulamak için kritik öneme sahiptir.

---

İstersen sana **HMAC ile token üretme ve doğrulama (JWT benzeri küçük bir örnek)** de gösterebilirim. İster misin?
EVET
*/

/*
Süper 👍 O zaman sana **HMAC kullanarak basit bir token üretme ve doğrulama (JWT benzeri mantık)** örneğini adım adım göstereyim.

JWT’lerde mantık şu:

1. **Payload (veri)** oluşturulur (örneğin kullanıcı adı, tarih, id).
2. Bu payload’ın HMAC-SHA256 ile imzası alınır.
3. Payload + imza birlikte token olarak gönderilir.
4. Karşı taraf aynı işlemi yaparak imzayı doğrular.

---

# 🛠 Basit HMAC Token Örneği
*/
``go
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
	"time"
)

// HMAC üretme fonksiyonu
func signMessage(message, key []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write(message)
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

// Token oluşturma
func createToken(username string, secret []byte) string {
	// Basit payload: kullanıcı adı + zaman damgası
	payload := fmt.Sprintf("%s|%d", username, time.Now().Unix())

	// Payload’u base64 ile encode et
	encodedPayload := base64.URLEncoding.EncodeToString([]byte(payload))

	// İmza oluştur
	signature := signMessage([]byte(encodedPayload), secret)

	// Token = payload + "." + signature
	return fmt.Sprintf("%s.%s", encodedPayload, signature)
}

// Token doğrulama
func verifyToken(token string, secret []byte) bool {
	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return false
	}

	encodedPayload := parts[0]
	signature := parts[1]

	// Beklenen imzayı hesapla
	expectedSig := signMessage([]byte(encodedPayload), secret)

	// Güvenli karşılaştırma
	return hmac.Equal([]byte(signature), []byte(expectedSig))
}

func main() {
	secretKey := []byte("supersecretkey")

	// Kullanıcı için token üretelim
	token := createToken("ali", secretKey)
	fmt.Println("Üretilen Token:", token)

	// Doğrulama
	if verifyToken(token, secretKey) {
		fmt.Println("✅ Token geçerli!")
	} else {
		fmt.Println("❌ Token geçersiz!")
	}

	// Yanlış secret ile doğrulamayı deneyelim
	wrongKey := []byte("fakekey")
	if verifyToken(token, wrongKey) {
		fmt.Println("Doğrulama başarılı (yanlışlıkla)!")
	} else {
		fmt.Println("Doğrulama başarısız (yanlış secret).")
	}
}
``
/*
---

# 📌 Çalışma Mantığı

1. `createToken("ali", secretKey)` →
   Payload = `"ali|1724400000"` (timestamp ile birlikte).
   Base64 → `"YWxpfDE3MjQ0MDAwMDA="`
   İmza → `"qJz8...=="`
   Token =

   ```
   YWxpfDE3MjQ0MDAwMDA=.qJz8...==
   ```

2. `verifyToken(token, secretKey)` →
   Payload’u alır, yeniden imza hesaplar.
   İmza eşleşiyorsa ✅ geçerli.

3. Yanlış secret kullanılırsa ❌ doğrulama başarısız olur.

---

# 🎯 Özet

* Bu örnek **JWT mantığının sadeleştirilmiş hali**.
* Gerçek JWT’de JSON kullanılır (`header.payload.signature`).
* Biz burada sadece payload + HMAC imzasını yaptık.
* Bu yöntem ile **token tabanlı oturum açma** gibi sistemler kurabilirsin.

---

İstersen bu örneği biraz daha geliştirip:
👉 Payload’u JSON yapalım (ör: `{ "user": "ali", "exp": 1724400000 }`)
👉 Token’a süre koyup (expiry) doğrularken kontrol edelim.

Bunu da görmek ister misin?
EVET
*/

/*
Harika 👍 şimdi bir adım ileri gidelim:
Token’ı **JSON payload** ile yapalım ve içine **expire (geçerlilik süresi)** koyalım. Böylece JWT’ye çok daha yakın bir örnek elde edeceğiz.

---

# 🛠 HMAC ile JSON Payload + Expiry Token
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

// Payload yapısı (JSON olacak)
type Payload struct {
	User string `json:"user"`
	Exp  int64  `json:"exp"` // geçerlilik süresi (timestamp)
}

// İmza oluşturma (HMAC-SHA256)
func signMessage(message, key []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write(message)
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

// Token oluşturma
func createToken(username string, secret []byte, duration time.Duration) string {
	// Geçerlilik süresi ekle
	payload := Payload{
		User: username,
		Exp:  time.Now().Add(duration).Unix(),
	}

	// Payload’u JSON’a çevir
	jsonPayload, _ := json.Marshal(payload)

	// Base64 encode
	encodedPayload := base64.URLEncoding.EncodeToString(jsonPayload)

	// İmza oluştur
	signature := signMessage([]byte(encodedPayload), secret)

	// Token formatı: payload.signature
	return fmt.Sprintf("%s.%s", encodedPayload, signature)
}

// Token doğrulama
func verifyToken(token string, secret []byte) (bool, string) {
	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return false, "Hatalı token formatı"
	}

	encodedPayload := parts[0]
	signature := parts[1]

	// Beklenen imzayı hesapla
	expectedSig := signMessage([]byte(encodedPayload), secret)
	if !hmac.Equal([]byte(signature), []byte(expectedSig)) {
		return false, "İmza geçersiz"
	}

	// Payload’u decode et
	payloadBytes, err := base64.URLEncoding.DecodeString(encodedPayload)
	if err != nil {
		return false, "Payload decode edilemedi"
	}

	var payload Payload
	err = json.Unmarshal(payloadBytes, &payload)
	if err != nil {
		return false, "JSON parse hatası"
	}

	// Süre kontrolü
	if time.Now().Unix() > payload.Exp {
		return false, "Token süresi dolmuş"
	}

	return true, fmt.Sprintf("Kullanıcı: %s", payload.User)
}

func main() {
	secretKey := []byte("supersecretkey")

	// 5 saniyelik token üretelim
	token := createToken("ali", secretKey, 5*time.Second)
	fmt.Println("Üretilen Token:", token)

	// Hemen doğrula
	ok, msg := verifyToken(token, secretKey)
	fmt.Println("Doğrulama:", ok, "-", msg)

	// 6 saniye bekleyelim (token expire olacak)
	time.Sleep(6 * time.Second)

	ok, msg = verifyToken(token, secretKey)
	fmt.Println("Doğrulama (6sn sonra):", ok, "-", msg)
}
``
/*
---

# 📌 Çalışma Mantığı

1. **Token oluşturma:**
*/
   ``json
   {
     "user": "ali",
     "exp": 1724400050
   }
   ``
/*
   → Base64 encode edilir + HMAC imzalanır.

2. **Doğrulama:**

   * İmza doğru mu?
   * JSON parse edildi mi?
   * `exp` zamanı dolmuş mu?

3. **Süre dolduysa**:

   ```
   Doğrulama (6sn sonra): false - Token süresi dolmuş
   ```

---

# 🎯 Özet

* Artık payload **JSON** formatında.
* Token’a **expiry süresi** eklendi.
* `verifyToken` hem imzayı hem de geçerlilik süresini kontrol ediyor.

Bu haliyle sistem, **JWT’nin basitleştirilmiş versiyonu** oldu 🚀

---

İstersen bir sonraki adımda sana:
👉 `header.payload.signature` formatını **tam bir JWT’ye** dönüştürüp gösterebilirim.
JWT’de header kısmında `"alg": "HS256", "typ": "JWT"` gibi bilgiler oluyor.

Bunu da görmek ister misin?
*/