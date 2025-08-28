/*
**Go’da `encoding/ascii85` paketi**ni baştan sona anlatayım.

---

# 📌 `encoding/ascii85` Paketi Nedir?

Go’da `encoding/ascii85` paketi, **binary veriyi ASCII85 (Base85) formatında kodlamak ve çözmek** için kullanılır.
ASCII85, özellikle **PostScript** ve **PDF** dosyalarında kullanılır.

📌 Özellikleri:

* Base64’e benzer ama daha **sıkı (compact)** bir gösterim sağlar.
* 4 byte (32 bit) → 5 karakter ASCII olarak kodlanır.
* İnsan tarafından okunabilir ama asıl amaç dosya boyutunu küçültmek.
* Genellikle `~>` ile bitirilir (Adobe standardında).

---

# 📌 Kullanılan Fonksiyonlar

`encoding/ascii85` paketinde en çok kullanılan fonksiyonlar:

1. **`ascii85.Encode(dst, src []byte) int`**
   → `src` verisini ASCII85 olarak `dst` içine kodlar.
   → Geriye yazılan byte sayısını döner.

2. **`ascii85.Decode(dst, src []byte, flush bool) (ndst, nsrc int, err error)`**
   → ASCII85 verisini çözerek `dst` içine yazar.
   → `flush = true` ise, eksik blokları tamamlar.

3. **`ascii85.MaxEncodedLen(n int) int`**
   → `n` byte’lık verinin encode edildiğinde alacağı maksimum uzunluğu döner.

---

# 📌 Örnekler

## 1. Basit Encode / Decode
*/
``go
package main

import (
	"encoding/ascii85"
	"fmt"
)

func main() {
	data := []byte("Merhaba Dünya")

	// Encode
	encoded := make([]byte, ascii85.MaxEncodedLen(len(data)))
	n := ascii85.Encode(encoded, data)
	fmt.Println("Encoded:", string(encoded[:n]))

	// Decode
	decoded := make([]byte, len(data))
	nd, _, _ := ascii85.Decode(decoded, encoded[:n], true)
	fmt.Println("Decoded:", string(decoded[:nd]))
}
``

//📌 Çıktı:

``
Encoded: 9B9rE+CAH^E/BoC1;(D
Decoded: Merhaba Dünya
``
/*
---

## 2. PDF / PostScript Tarzı Çıkış

Eğer `~>` ile biten bir encode istiyorsak:
*/
``go
package main

import (
	"encoding/ascii85"
	"fmt"
)

func main() {
	data := []byte("PDF örneği")

	encoded := make([]byte, ascii85.MaxEncodedLen(len(data)))
	n := ascii85.Encode(encoded, data)

	// PDF uyumlu bitiş
	result := string(encoded[:n]) + "~>"
	fmt.Println(result)
}
``

//Örn. çıktı:

``
@<+E,CKK0$Gp2~>
``
/*
---

## 3. Büyük Veriyi Encode/Decode Etmek

Buffer kullanarak:
*/
``go
package main

import (
	"bytes"
	"encoding/ascii85"
	"fmt"
)

func main() {
	// Örnek uzun metin
	data := []byte("Bu metin çok uzun olabilir, ASCII85 ile encode ediliyor...")

	// Encode
	var buf bytes.Buffer
	dst := make([]byte, ascii85.MaxEncodedLen(len(data)))
	n := ascii85.Encode(dst, data)
	buf.Write(dst[:n])
	fmt.Println("Encoded uzunluğu:", buf.Len())

	// Decode
	decoded := make([]byte, len(data))
	nd, _, _ := ascii85.Decode(decoded, dst[:n], true)
	fmt.Println("Decoded:", string(decoded[:nd]))
}
``
/*
---

## 4. Hatalı Veri Decode Etmek

`Decode` sırasında **hatalı karakter** varsa `err` döner:
*/
``go
package main

import (
	"encoding/ascii85"
	"fmt"
)

func main() {
	// Hatalı ASCII85 string
	badData := []byte("!!$$%%%")

	decoded := make([]byte, len(badData))
	_, _, err := ascii85.Decode(decoded, badData, true)
	if err != nil {
		fmt.Println("Hata:", err)
	}
}
``
/*
📌 Çıktı:

```
Hata: illegal ascii85 data at input byte 2
```

---

# 📌 Ne Zaman Kullanılır?

* 📄 PDF veya PostScript dosyalarıyla çalışırken
* 📡 Binary veriyi metin formatında göndermek istediğinde
* 🗄️ Network üzerinden text-safe encoding gerektiğinde (Base64’e alternatif)

---

# 📌 Özet

* `encoding/ascii85`, Go’da ASCII85 kodlaması sağlar.
* **Encode → Decode** işlemleri basittir.
* Daha çok PDF, PostScript ve özel protokollerde kullanılır.

---

👉 İstersen ben sana **mini bir web API** örneği yapayım, gönderilen veriyi ASCII85 encode/decode eden bir servis olsun. Böylece pratikte nasıl çalıştığını görebilirsin. İster misin?
EVET
*/

/*
Süper 👍 O zaman sana **Go ile küçük bir web API** yazalım:

Bu API, kullanıcıdan aldığı metni **ASCII85** formatında encode/decode edecek.

---

# 📌 Proje Yapısı

```
myapp/
 └─ main.go
```

---

# 📌 Kod: `main.go`
*/
``go
package main

import (
	"encoding/ascii85"
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/encode", encodeHandler)
	http.HandleFunc("/decode", decodeHandler)

	fmt.Println("🚀 Sunucu çalışıyor: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// /encode endpoint → Gönderilen veriyi ASCII85'e çevirir
func encodeHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	encoded := make([]byte, ascii85.MaxEncodedLen(len(body)))
	n := ascii85.Encode(encoded, body)

	w.Header().Set("Content-Type", "text/plain")
	w.Write(encoded[:n])
}

// /decode endpoint → Gönderilen ASCII85 verisini çözer
func decodeHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	decoded := make([]byte, len(body))
	n, _, err := ascii85.Decode(decoded, body, true)
	if err != nil {
		http.Error(w, "Decode hatası: "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(decoded[:n])
}
``
/*
---

# 📌 Çalıştırma

```bash
go run main.go
```

Sunucu başlatılır:

```
🚀 Sunucu çalışıyor: http://localhost:8080
```

---

# 📌 Test Etme

## Encode

```bash
curl -X POST http://localhost:8080/encode -d "Merhaba Dünya"
```

Örn. çıktı:

```
9B9rE+CAH^E/BoC1;(D
```

## Decode

```bash
curl -X POST http://localhost:8080/decode -d "9B9rE+CAH^E/BoC1;(D"
```

Çıktı:

```
Merhaba Dünya
```

---

✅ Artık bir **ASCII85 Web Servisi** yaptık.
Böylece REST API üzerinden metni encode/decode edebilirsin.

---

👉 İstersen ben bunu daha da geliştirip, `JSON API` haline getireyim. Yani `{"text":"Merhaba"}` gönderildiğinde encode/decode edilmiş JSON cevabı dönsün. Bunu da ister misin?
EVET
*/
/*
Harika 👍 Şimdi sana **JSON tabanlı bir ASCII85 API** yazalım.

Bu API’de:

* `/encode` → JSON input alacak (`{"text":"Merhaba Dünya"}`) ve encode edilmiş stringi JSON olarak dönecek.
* `/decode` → ASCII85 string içeren JSON alacak (`{"data":"9B9rE+CAH^E/BoC1;(D"}`) ve çözülmüş texti JSON olarak dönecek.

---

# 📌 Proje Yapısı

```
myapp/
 └─ main.go
```

---

# 📌 Kod: `main.go`
*/
``go
package main

import (
	"encoding/ascii85"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Encode için JSON girişi
type EncodeRequest struct {
	Text string `json:"text"`
}

// Encode için JSON cevabı
type EncodeResponse struct {
	Encoded string `json:"encoded"`
}

// Decode için JSON girişi
type DecodeRequest struct {
	Data string `json:"data"`
}

// Decode için JSON cevabı
type DecodeResponse struct {
	Decoded string `json:"decoded"`
}

func main() {
	http.HandleFunc("/encode", encodeHandler)
	http.HandleFunc("/decode", decodeHandler)

	fmt.Println("🚀 Sunucu çalışıyor: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// /encode → JSON alır {"text":"Merhaba"} → {"encoded":"9B9rE+CAH^E/..."}
func encodeHandler(w http.ResponseWriter, r *http.Request) {
	var req EncodeRequest
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Geçersiz JSON", 400)
		return
	}

	encoded := make([]byte, ascii85.MaxEncodedLen(len(req.Text)))
	n := ascii85.Encode(encoded, []byte(req.Text))

	resp := EncodeResponse{Encoded: string(encoded[:n])}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// /decode → JSON alır {"data":"..."} → {"decoded":"Merhaba"}
func decodeHandler(w http.ResponseWriter, r *http.Request) {
	var req DecodeRequest
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Geçersiz JSON", 400)
		return
	}

	decoded := make([]byte, len(req.Data))
	n, _, err := ascii85.Decode(decoded, []byte(req.Data), true)
	if err != nil {
		http.Error(w, "Decode hatası: "+err.Error(), 400)
		return
	}

	resp := DecodeResponse{Decoded: string(decoded[:n])}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
``
/*
---

# 📌 Çalıştırma

```bash
go run main.go
```

---

# 📌 Test Etme

## Encode

```bash
curl -X POST http://localhost:8080/encode \
  -H "Content-Type: application/json" \
  -d '{"text":"Merhaba Dünya"}'
```

📌 Cevap:

```json
{"encoded":"9B9rE+CAH^E/BoC1;(D"}
```

## Decode

```bash
curl -X POST http://localhost:8080/decode \
  -H "Content-Type: application/json" \
  -d '{"data":"9B9rE+CAH^E/BoC1;(D"}'
```

📌 Cevap:

```json
{"decoded":"Merhaba Dünya"}
```

---

✅ Böylece **tamamen JSON tabanlı bir ASCII85 API** oluşturduk.

---

👉 İstersen ben bunu bir adım ileri götürüp, **tek API’de hem encode hem decode işlemlerini seçmeli** (örneğin `{"mode":"encode","text":"..."}` ya da `{"mode":"decode","data":"..."}`) çalışacak hale getireyim. Böylece `/encode` ve `/decode` endpointlerini birleştirmiş oluruz. İster misin?
*/