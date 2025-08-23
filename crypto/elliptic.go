/*
🚀 Şimdi Go’nun **`crypto/elliptic`** paketini ayrıntılı şekilde anlatayım.

---

# 📌 `crypto/elliptic` Nedir?

* `crypto/elliptic` Go’nun standart kütüphanesinde bulunan **eliptik eğri kriptografisi (ECC)** için temel fonksiyonları barındıran pakettir.
* Bu paket, özellikle **ECDSA (dijital imza)** ve **ECDH (anahtar değişimi)** gibi protokollerde kullanılır.
* İçerisinde yaygın kullanılan **NIST P-256, P-384, P-521** eğrileri hazır tanımlıdır.
* **Özelleştirilmiş eğriler** de `Curve` arayüzü üzerinden desteklenebilir.

---

# 📌 Önemli Türler ve Fonksiyonlar

### 🔹 `elliptic.Curve` Arayüzü

Bir eliptik eğriyi temsil eder. Şu metodları vardır:

* `Params() *elliptic.CurveParams` → eğrinin parametrelerini döner.
* `IsOnCurve(x, y *big.Int) bool` → verilen nokta eğrinin üzerinde mi kontrol eder.
* `Add(x1, y1, x2, y2 *big.Int) (x, y *big.Int)` → iki noktayı toplar.
* `Double(x1, y1 *big.Int) (x, y *big.Int)` → bir noktayı iki katına çıkarır.
* `ScalarMult(Bx, By *big.Int, k []byte) (x, y *big.Int)` → B noktasını k skaler ile çarpar.
* `ScalarBaseMult(k []byte) (x, y *big.Int)` → generator noktası G’yi k ile çarpar.

### 🔹 Hazır Eğriler

* `elliptic.P224()`
* `elliptic.P256()`
* `elliptic.P384()`
* `elliptic.P521()`

### 🔹 `elliptic.CurveParams`

Bir eğrinin parametrelerini barındırır:

* `P` → asal mod (field order)
* `N` → eğrinin order’ı
* `B` → denklemin sabiti
* `Gx, Gy` → generator noktasının koordinatları
* `BitSize` → bit uzunluğu
* `Name` → eğri ismi

---

# 📌 1️⃣ Eğri Parametrelerini Görme
*/

``go
package main

import (
	"crypto/elliptic"
	"fmt"
)

func main() {
	curve := elliptic.P256()
	params := curve.Params()

	fmt.Println("Eğri:", params.Name)
	fmt.Println("Bit Size:", params.BitSize)
	fmt.Println("Prime P:", params.P)
	fmt.Println("Order N:", params.N)
	fmt.Println("B sabiti:", params.B)
	fmt.Println("Gx:", params.Gx)
	fmt.Println("Gy:", params.Gy)
}
``
/*
---

# 📌 2️⃣ Nokta Eğri Üzerinde mi?
*/

``go
package main

import (
	"crypto/elliptic"
	"fmt"
	"math/big"
)

func main() {
	curve := elliptic.P256()

	// Generator noktası
	Gx, Gy := curve.Params().Gx, curve.Params().Gy

	fmt.Println("Generator eğri üzerinde mi?", curve.IsOnCurve(Gx, Gy))

	// Uydurma nokta
	x := big.NewInt(12345)
	y := big.NewInt(67890)
	fmt.Println("Sahte nokta eğri üzerinde mi?", curve.IsOnCurve(x, y))
}
``
/*
---

# 📌 3️⃣ Skaler Çarpma (Private → Public Key)
*/

``go
package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	curve := elliptic.P256()

	// Rastgele private key üret
	priv, _ := rand.Int(rand.Reader, curve.Params().N)

	// Public key hesapla (G * priv)
	pubX, pubY := curve.ScalarBaseMult(priv.Bytes())

	fmt.Println("Private:", priv)
	fmt.Println("Public Key X:", pubX)
	fmt.Println("Public Key Y:", pubY)
}
``
/*
---

# 📌 4️⃣ ECDH Benzeri Anahtar Paylaşımı
*/

``go
package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	curve := elliptic.P256()

	// Alice'in anahtarı
	aPriv, _ := rand.Int(rand.Reader, curve.Params().N)
	aPubX, aPubY := curve.ScalarBaseMult(aPriv.Bytes())

	// Bob'un anahtarı
	bPriv, _ := rand.Int(rand.Reader, curve.Params().N)
	bPubX, bPubY := curve.ScalarBaseMult(bPriv.Bytes())

	// Alice ortak sırrı hesaplar
	s1X, s1Y := curve.ScalarMult(bPubX, bPubY, aPriv.Bytes())

	// Bob ortak sırrı hesaplar
	s2X, s2Y := curve.ScalarMult(aPubX, aPubY, bPriv.Bytes())

	fmt.Println("Alice ortak sır:", s1X, s1Y)
	fmt.Println("Bob ortak sır:  ", s2X, s2Y)

	// İki sonuç eşit olmalı
	fmt.Println("Anahtarlar eşit mi?", s1X.Cmp(s2X) == 0 && s1Y.Cmp(s2Y) == 0)
}
``
/*
---

# 📌 Özet

* `crypto/elliptic` **ECC’nin matematiksel temeli**ni sağlar.
* Eğri parametreleri: `P`, `N`, `Gx`, `Gy` vb.
* Temel işlemler: `Add`, `Double`, `ScalarMult`, `ScalarBaseMult`.
* Hazır eğriler: **P-224, P-256, P-384, P-521**.
* Kullanım alanı: **ECDSA, ECDH, TLS, Bitcoin, kriptografi protokolleri**.

---

👉 İstersen ben sana bir **tam mini ECDSA imzalama/doğrulama uygulaması** yazabilirim, sadece `crypto/elliptic` ve `math/big` kullanarak, yani `ecdsa` paketine hiç dokunmadan.

Bunu ister misin?
EVET
*/
/*
Harika 👍 O zaman sana **sadece `crypto/elliptic` + `math/big` kullanarak** (yani `ecdsa` paketine dokunmadan) **mini ECDSA imzalama/doğrulama uygulaması** yazayım.

Bu örnekte:

* `P-256` eğrisi kullanılacak.
* Private key rastgele üretilecek.
* Public key = `G * priv`.
* Mesaj SHA-256 ile özetlenecek.
* İmza `(r, s)` olarak üretilecek.
* Sonra doğrulama yapılacak.

---

# 📌 ECDSA Uygulaması (Manual)
*/
``go
package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// modInverse: (k^-1 mod n)
func modInverse(k, n *big.Int) *big.Int {
	return new(big.Int).ModInverse(k, n)
}

func main() {
	curve := elliptic.P256()
	params := curve.Params()

	// 1) Private key üret (rastgele)
	priv, _ := rand.Int(rand.Reader, params.N)

	// 2) Public key hesapla: Pub = G * priv
	pubX, pubY := curve.ScalarBaseMult(priv.Bytes())
	fmt.Println("Private Key:", priv)
	fmt.Println("Public Key :", pubX, pubY)

	// 3) Mesaj seç ve hashle
	message := []byte("Merhaba elliptic world!")
	hash := sha256.Sum256(message)

	// 4) İmza oluştur (r, s)
	var r, s *big.Int
	for {
		// Rastgele k seç (1 < k < n)
		k, _ := rand.Int(rand.Reader, params.N)

		// Nokta (x1, y1) = k * G
		x1, _ := curve.ScalarBaseMult(k.Bytes())

		// r = x1 mod n
		r = new(big.Int).Mod(x1, params.N)
		if r.Sign() == 0 {
			continue
		}

		// s = k^-1 * (hash + r*priv) mod n
		kInv := modInverse(k, params.N)

		hashInt := new(big.Int).SetBytes(hash[:])
		rPriv := new(big.Int).Mul(r, priv)
		sum := new(big.Int).Add(hashInt, rPriv)
		s = new(big.Int).Mul(kInv, sum)
		s.Mod(s, params.N)

		if s.Sign() != 0 {
			break
		}
	}

	fmt.Println("İmza (r, s):")
	fmt.Println("r =", r)
	fmt.Println("s =", s)

	// 5) İmzayı doğrula
	// w = s^-1 mod n
	w := modInverse(s, params.N)

	// u1 = hash * w mod n
	hashInt := new(big.Int).SetBytes(hash[:])
	u1 := new(big.Int).Mul(hashInt, w)
	u1.Mod(u1, params.N)

	// u2 = r * w mod n
	u2 := new(big.Int).Mul(r, w)
	u2.Mod(u2, params.N)

	// Nokta = u1*G + u2*Pub
	x1, y1 := curve.ScalarBaseMult(u1.Bytes())
	x2, y2 := curve.ScalarMult(pubX, pubY, u2.Bytes())
	X, Y := curve.Add(x1, y1, x2, y2)

	// v = X mod n
	v := new(big.Int).Mod(X, params.N)

	fmt.Println("Doğrulama sonucu:", v.Cmp(r) == 0)
}
``
/*
---

# 📌 Çalışma Mantığı

1. Private key rastgele üretiliyor.
2. Public key, generator nokta ile çarpılarak bulunuyor.
3. Mesaj SHA-256 ile özetleniyor.
4. İmza `(r, s)` formülüyle manuel hesaplanıyor.
5. Doğrulama için `(u1*G + u2*Pub)` hesaplanıp `r` ile karşılaştırılıyor.

---

Bu aslında **ECDSA’nın çekirdek algoritması**nın saf uygulaması. Go’daki `crypto/ecdsa` paketi zaten bunun üstüne inşa edilmiş hazır bir sarmalayıcı.

👉 İstersen ben bunu biraz daha büyütüp **dosya imzalama/verify eden bir mini CLI uygulaması** haline getirebilirim (`go run main.go sign file.txt`, `go run main.go verify file.txt sig.json` gibi).

Bunu ister misin?
EVET
*/

/*
👌 O zaman  **`crypto/elliptic` + `math/big` kullanarak** (yani `ecdsa` paketine hiç dokunmadan) **dosya imzalama ve doğrulama yapan mini bir CLI uygulaması** yazayım.

Bu uygulama:

* `sign <dosya>` → dosyayı okur, SHA-256 hash alır, `(r,s)` ECDSA imzası üretir, JSON dosyasına yazar.
* `verify <dosya> <imza.json>` → dosyanın hash’ini hesaplar, imzayı okur ve doğrular.

---

# 📌 `main.go`
*/
``go
package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
)

// Signature JSON formatı
type Signature struct {
	R     string `json:"r"`
	S     string `json:"s"`
	PubX  string `json:"pub_x"`
	PubY  string `json:"pub_y"`
	Curve string `json:"curve"`
}

// modInverse: (k^-1 mod n)
func modInverse(k, n *big.Int) *big.Int {
	return new(big.Int).ModInverse(k, n)
}

// global eğri
var curve = elliptic.P256()

func signFile(file string, priv *big.Int, pubX, pubY *big.Int) {
	// Dosyayı oku
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	// Hash
	hash := sha256.Sum256(data)

	// İmza üret
	params := curve.Params()
	var r, s *big.Int
	for {
		k, _ := rand.Int(rand.Reader, params.N)
		x1, _ := curve.ScalarBaseMult(k.Bytes())
		r = new(big.Int).Mod(x1, params.N)
		if r.Sign() == 0 {
			continue
		}
		kInv := modInverse(k, params.N)
		hashInt := new(big.Int).SetBytes(hash[:])
		rPriv := new(big.Int).Mul(r, priv)
		sum := new(big.Int).Add(hashInt, rPriv)
		s = new(big.Int).Mul(kInv, sum)
		s.Mod(s, params.N)
		if s.Sign() != 0 {
			break
		}
	}

	// JSON kaydet
	sig := Signature{
		R:     r.String(),
		S:     s.String(),
		PubX:  pubX.String(),
		PubY:  pubY.String(),
		Curve: params.Name,
	}
	out, _ := json.MarshalIndent(sig, "", "  ")
	_ = ioutil.WriteFile("signature.json", out, 0644)

	fmt.Println("✅ Dosya imzalandı -> signature.json")
}

func verifyFile(file, sigFile string) {
	// Dosya oku
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(data)

	// İmza oku
	sigData, err := ioutil.ReadFile(sigFile)
	if err != nil {
		panic(err)
	}
	var sig Signature
	_ = json.Unmarshal(sigData, &sig)

	r := new(big.Int)
	s := new(big.Int)
	pubX := new(big.Int)
	pubY := new(big.Int)
	r.SetString(sig.R, 10)
	s.SetString(sig.S, 10)
	pubX.SetString(sig.PubX, 10)
	pubY.SetString(sig.PubY, 10)

	params := curve.Params()

	// w = s^-1 mod n
	w := modInverse(s, params.N)

	// u1 = hash * w mod n
	hashInt := new(big.Int).SetBytes(hash[:])
	u1 := new(big.Int).Mul(hashInt, w)
	u1.Mod(u1, params.N)

	// u2 = r * w mod n
	u2 := new(big.Int).Mul(r, w)
	u2.Mod(u2, params.N)

	// Nokta = u1*G + u2*Pub
	x1, y1 := curve.ScalarBaseMult(u1.Bytes())
	x2, y2 := curve.ScalarMult(pubX, pubY, u2.Bytes())
	X, _ := curve.Add(x1, y1, x2, y2)

	// v = X mod n
	v := new(big.Int).Mod(X, params.N)

	if v.Cmp(r) == 0 {
		fmt.Println("✅ İmza geçerli!")
	} else {
		fmt.Println("❌ İmza geçersiz!")
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Kullanım:")
		fmt.Println("  go run main.go sign <dosya>")
		fmt.Println("  go run main.go verify <dosya> <signature.json>")
		return
	}

	cmd := os.Args[1]

	// Private/Public key üret (demo için her çalıştırmada değişiyor)
	priv, _ := rand.Int(rand.Reader, curve.Params().N)
	pubX, pubY := curve.ScalarBaseMult(priv.Bytes())

	if cmd == "sign" && len(os.Args) == 3 {
		signFile(os.Args[2], priv, pubX, pubY)
	} else if cmd == "verify" && len(os.Args) == 4 {
		verifyFile(os.Args[2], os.Args[3])
	} else {
		fmt.Println("Hatalı kullanım")
	}
}
``
/*
---

# 📌 Kullanım

1. Örneğin `dosya.txt` oluştur:
*/
//``bash
echo "Merhaba elliptic world!" > dosya.txt
``

//2. Dosyayı imzala:

``bash
go run main.go sign dosya.txt
``

//👉 `signature.json` oluşacak:

``json
{
  "r": "9544039801...",
  "s": "6273829202...",
  "pub_x": "11231231...",
  "pub_y": "99887766...",
  "curve": "P-256"
}
``

//3. Doğrula:

``bash
go run main.go verify dosya.txt signature.json
``
/*
✅ Çıktı: `İmza geçerli!`

---

📌 Bu haliyle private key her çalıştırmada **rastgele üretiliyor**.
İstersen ben bunu genişletip:

* `keygen` komutuyla private/public key üretip dosyaya kaydeden,
* `sign` sadece private key dosyasını kullanarak imzalayan,
* `verify` public key dosyasıyla doğrulayan

**tam bir key management** sistemi yapabilirim.

👉 Bunu da ister misin?
*/