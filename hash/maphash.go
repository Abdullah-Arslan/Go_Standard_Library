/*
Go’da **`hash/maphash`** paketi, **hızlı, güvenli ve çakışma direnci yüksek hash fonksiyonları** sağlamak için tasarlanmıştır.

Bu paket, özellikle **map anahtarlarını hashlemek**, string/veri parçalarını güvenli şekilde hashlemek için kullanılır.
Go 1.14’ten itibaren standart kütüphaneye dahil edilmiştir ve **hash tabanlı saldırılara karşı güvenlik sağlar**.

---

# 📦 Paket İçe Aktarma
*/
``go
import "hash/maphash"
``
/*
---

# 🔑 Temel Yapılar

### 1. `type Seed struct`

* Hash fonksiyonunun başlangıç noktası için kullanılan rastgele bir değer içerir.
* Aynı `Seed` ile aynı veriye aynı hash çıkar.
* Farklı `Seed` ile aynı veriye farklı hash çıkar (saldırıları önlemek için).

### 2. `type Hash struct`

* Akış bazlı (streaming) hash hesaplaması yapılır.
* `hash.Hash64` arayüzünü uygular.

---

# 🔧 Önemli Fonksiyonlar

### 1. `maphash.MakeSeed() maphash.Seed`

Yeni, rastgele bir seed oluşturur. Her çalıştırmada farklı olabilir.
*/
``go
package main

import (
	"fmt"
	"hash/maphash"
)

func main() {
	seed := maphash.MakeSeed()
	fmt.Println("Yeni seed:", seed)
}
``
/*
---

### 2. `(*Hash).SetSeed(seed maphash.Seed)`

Bir hash objesine seed atar.
*/
``go
func main() {
	var h maphash.Hash
	seed := maphash.MakeSeed()
	h.SetSeed(seed)
}
``
/*
---

### 3. `(*Hash).Write(p []byte)` ve `(*Hash).Sum64()`

Veriyi hash’e yazıp, 64-bit hash sonucunu döner.
*/

``go
func main() {
	var h maphash.Hash
	h.SetSeed(maphash.MakeSeed())

	h.Write([]byte("Merhaba "))
	h.Write([]byte("Golang"))

	fmt.Printf("Hash değeri: %x\n", h.Sum64())
}
``
/*
---

### 4. `maphash.Bytes(seed maphash.Seed, b []byte) uint64`

Tek adımda byte slice’in hash’ini alır.
*/
``go
func main() {
	seed := maphash.MakeSeed()
	data := []byte("Merhaba Dünya")

	hash := maphash.Bytes(seed, data)
	fmt.Printf("Bytes hash: %x\n", hash)
}
``
/*
---

### 5. `maphash.String(seed maphash.Seed, s string) uint64`

Bir string’in hash’ini hızlıca hesaplar.
*/
``go
func main() {
	seed := maphash.MakeSeed()

	hash1 := maphash.String(seed, "OpenAI")
	hash2 := maphash.String(seed, "ChatGPT")

	fmt.Printf("OpenAI hash: %x\n", hash1)
	fmt.Printf("ChatGPT hash: %x\n", hash2)
}
``
/*
---

# 📂 Gerçek Hayat Senaryosu

### ✅ String Anahtarlarını Hashlemek (Map İçin Güvenli)
*/
``go
func main() {
	seed := maphash.MakeSeed()
	strings := []string{"apple", "banana", "orange"}

	for _, s := range strings {
		h := maphash.String(seed, s)
		fmt.Printf("%s -> %x\n", s, h)
	}
}
``

//📌 Çıktı (her çalıştırmada farklı seed nedeniyle farklı olabilir):

``
apple  -> 6a9c6c5c7f11a1c2
banana -> 2d5e4c3f8b9e0a4d
orange -> 7f8b9a12345abcd9
``
/*
---

### ✅ Streaming Hash (Büyük Veriler İçin)

Büyük dosya veya veri akışını parça parça hashlemek için `Hash` kullanılır:
*/
``go
import (
	"fmt"
	"hash/maphash"
	"io"
	"os"
)

func FileHash(filename string, seed maphash.Seed) (uint64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var h maphash.Hash
	h.SetSeed(seed)

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if n > 0 {
			h.Write(buf[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}
	}
	return h.Sum64(), nil
}

func main() {
	seed := maphash.MakeSeed()
	hash, _ := FileHash("test.txt", seed)
	fmt.Printf("Dosya hash değeri: %x\n", hash)
}
``
/*
---

# 📝 Özet

* **`maphash`**, hızlı ve güvenli hash üretmek için kullanılır.
* Özellikle Go **map anahtarları için güvenli hashleme** algoritmalarına dayanır.
* Fonksiyonlar:

  * `MakeSeed()` → Rastgele seed üretir.
  * `Bytes(seed, data)` → \[]byte hash’ler.
  * `String(seed, str)` → string hash’ler.
  * `Hash` struct → Streaming hash hesaplama.
* **Her çalıştırmada farklı seed** = saldırılara karşı güvenlik.

---

👉 İstersen sana hem `fnv` hem de `maphash` kullanarak **aynı string dizisini hashleyip karşılaştıran bir örnek** yazayım. Böylece hız ve farklılıkları görebilirsin. İstiyor musun?
EVET
*/
/*
Harika 👍 O zaman şimdi sana **aynı string listesi** için hem **`fnv`** hem de **`maphash`** kullanarak hash hesaplayan bir örnek yazacağım.

Bu sayede:

* **`fnv`** → deterministik (her zaman aynı sonucu verir)
* **`maphash`** → güvenli ve hızlı, ama `seed` değişirse sonuçlar da değişir

farkını görebileceksin. 🚀

---

## 📂 Örnek: `fnv_vs_maphash.go`
*/
``go
package main

import (
	"fmt"
	"hash/fnv"
	"hash/maphash"
)

func main() {
	// Karşılaştırma için string listesi
	words := []string{"apple", "banana", "orange", "grape", "melon"}

	// FNV Hash
	fmt.Println("🔹 FNV-1a 64-bit:")
	for _, w := range words {
		h := fnv.New64a()
		h.Write([]byte(w))
		fmt.Printf("%s -> %x\n", w, h.Sum64())
	}

	fmt.Println()

	// maphash Hash (tek seed kullanıyoruz)
	seed := maphash.MakeSeed()
	fmt.Println("🔹 maphash 64-bit (tek seed ile):")
	for _, w := range words {
		h := maphash.String(seed, w)
		fmt.Printf("%s -> %x\n", w, h)
	}

	fmt.Println()

	// maphash Hash (her seferinde farklı seed)
	fmt.Println("🔹 maphash 64-bit (her defasında farklı seed):")
	for _, w := range words {
		h := maphash.String(maphash.MakeSeed(), w)
		fmt.Printf("%s -> %x\n", w, h)
	}
}
``
/*
---

## 📌 Örnek Çıktı
*/
``
🔹 FNV-1a 64-bit:
apple -> 4f9f2cab3f92d9e5
banana -> 08b6e6a51707a7a6
orange -> 9ec9f59a2a53bfee
grape -> 53ccaa0c99b2ff37
melon -> 5b95bdb7a52cc2cb

🔹 maphash 64-bit (tek seed ile):
apple -> 7c2e3d4459e7c3a9
banana -> 6fa46b2d93bb07b0
orange -> 1a5f3ed79ca88afc
grape -> 4209f96c4c1a5ed1
melon -> 3c9bb41af07ce9e7

🔹 maphash 64-bit (her defasında farklı seed):
apple -> 8a3e5b4f2a9c1d93
banana -> 29eab65ff0d9c2e1
orange -> 6b4d91f7a8d12c99
grape -> d1e2f8c33b7a5c42
melon -> 71aefc23d9b84f5e
``
/*
---

# 🔍 Yorum

* **FNV** → her zaman sabit sonuçlar üretir, tekrar edilebilir testler için ideal.
* **maphash (tek seed)** → aynı çalıştırmada tekrar edilebilir, ama seed değişirse sonuç değişir.
* **maphash (farklı seed)** → her çağrıda farklı sonuç, bu da saldırılara karşı daha güvenli.

---

👉 İstersen ben sana küçük bir **benchmark** örneği de yazayım. Böylece `fnv` ve `maphash` hızlarını karşılaştırabilirsin. İster misin?
*/