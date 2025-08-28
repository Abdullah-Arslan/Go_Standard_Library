/*
Şimdi sana **Go’da `crypto/aes` paketini** en ince ayrıntısına kadar anlatayım.

---

# 📘 `crypto/aes` Paketi Nedir?

Go’nun standart kütüphanesindeki **`crypto/aes`** paketi,
**AES (Advanced Encryption Standard)** algoritmasının implementasyonunu sağlar.

AES:

* Blok şifreleme algoritmasıdır.
* 128-bit (16 byte) bloklarla çalışır.
* **Key uzunluğu**: 128, 192 veya 256 bit olabilir (yani 16, 24, 32 byte).
* Tek başına kullanılmaz, genelde **`crypto/cipher`** paketi ile birlikte CBC, CFB, CTR, GCM modlarında kullanılır.

---

# 📦 Temel Fonksiyonlar

1. **`aes.NewCipher(key []byte) (cipher.Block, error)`**

   * Key (16, 24 veya 32 byte) alır.
   * Bir **AES cipher.Block** döner.

2. **`aes.BlockSize`**

   * AES’in sabit blok boyutu (16 byte).

---

# 🔎 Örnekler

---

## 1. AES Blok Şifreleme (Temel Kullanım)
*/
``
package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	key := []byte("examplekey123456") // 16 byte
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	fmt.Println("AES block size:", block.BlockSize()) // 16
}
``
/*
---

## 2. AES-CBC (Cipher Block Chaining)

AES bloklarını zincirleme şifreleme modu.
*/
``
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	key := []byte("examplekey123456") // 16 byte
	plaintext := []byte("Merhaba CBC Mode!")

	block, _ := aes.NewCipher(key)

	// IV (Initialization Vector)
	iv := make([]byte, aes.BlockSize)
	io.ReadFull(rand.Reader, iv)

	// CBC Encrypter
	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)

	fmt.Printf("CBC şifreli: %x\n", ciphertext)
}
``
/*
📌 CBC’nin güvenli olması için **padding** eklenmesi gerekir. (PKCS7 gibi)

---

## 3. AES-CFB (Cipher Feedback)

Akış modunda şifreleme, padding gerekmez.
*/
``
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	key := []byte("examplekey123456")
	plaintext := []byte("Merhaba CFB Mode!")

	block, _ := aes.NewCipher(key)

	iv := make([]byte, aes.BlockSize)
	io.ReadFull(rand.Reader, iv)

	// Şifreleme
	ciphertext := make([]byte, len(plaintext))
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("Şifreli: %x\n", ciphertext)

	// Çözme
	decrypted := make([]byte, len(ciphertext))
	streamDec := cipher.NewCFBDecrypter(block, iv)
	streamDec.XORKeyStream(decrypted, ciphertext)
	fmt.Println("Çözüldü:", string(decrypted))
}
``
/*
---

## 4. AES-CTR (Counter Mode)

Çok hızlı, streaming şifreleme için kullanılır.
*/
``
block, _ := aes.NewCipher(key)
iv := make([]byte, aes.BlockSize)
io.ReadFull(rand.Reader, iv)

stream := cipher.NewCTR(block, iv)
ciphertext := make([]byte, len(plaintext))
stream.XORKeyStream(ciphertext, plaintext)
``
/*
---

## 5. AES-GCM (Galois/Counter Mode) ✅ En Güvenli

AES-GCM, hem **şifreleme** hem **doğrulama (auth)** sağlar.
Modern uygulamalarda en çok kullanılan AES modudur (ör. TLS).
*/
``
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	key := []byte("examplekey123456")
	plaintext := []byte("Merhaba AES-GCM!")

	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)

	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)

	// Şifreleme
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	fmt.Printf("AES-GCM şifreli: %x\n", ciphertext)

	// Çözme
	nonce, enc := ciphertext[:gcm.NonceSize()], ciphertext[gcm.NonceSize():]
	plaintext2, err := gcm.Open(nil, nonce, enc, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Çözüldü:", string(plaintext2))
}
``
/*
---

## 6. AES-EAX, AES-OCB (Standart Kütüphanede Yok)

Go `crypto/aes` sadece temel AES blok şifreyi sağlar.
Ek modlar (EAX, OCB) Go standart kütüphanede **yoktur**, harici kütüphanelerden kullanılır.

---

# ✅ Özet (Cheatsheet)

| Fonksiyon / Sabit        | Açıklama                                   |
| ------------------------ | ------------------------------------------ |
| `aes.NewCipher(key)`     | AES cipher oluşturur (16, 24, 32 byte key) |
| `aes.BlockSize`          | 16 byte sabit blok boyutu                  |
| `cipher.NewCBCEncrypter` | CBC şifreleme                              |
| `cipher.NewCBCDecrypter` | CBC çözme                                  |
| `cipher.NewCFBEncrypter` | CFB şifreleme                              |
| `cipher.NewCFBDecrypter` | CFB çözme                                  |
| `cipher.NewCTR`          | CTR modu                                   |
| `cipher.NewGCM`          | GCM modu (en güvenli)                      |

---

📌 `crypto/aes` paketinde asıl iş **AES blok şifreyi oluşturmak**.
Mod seçimi (CBC, CFB, CTR, GCM) **`crypto/cipher`** paketinden yapılır.

---

👉 İstersen ben sana **AES ile dosya şifreleme/çözme aracı** (CLI mini proje) yazabilirim. Bu gerçek hayatta çok işine yarar. Bunu görmek ister misin?
Evet
*/
/*
O zaman **Go ile AES-GCM kullanarak dosya şifreleme/çözme aracı (CLI mini proje)** yazayım 🚀

Bu araç ile:

* Bir dosyayı AES-GCM ile şifreleyeceğiz.
* Aynı dosyayı AES-GCM ile geri çözeceğiz.
* Key (şifre) kullanıcıdan alınacak.

---

# 📘 AES-GCM Dosya Şifreleme/Çözme Aracı
*/
``
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func encryptFile(filename string, key []byte) error {
	plaintext, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	encFile := filename + ".enc"
	err = os.WriteFile(encFile, ciphertext, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Şifrelendi:", encFile)
	return nil
}

func decryptFile(filename string, key []byte) error {
	ciphertext, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return fmt.Errorf("geçersiz dosya")
	}

	nonce, data := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, data, nil)
	if err != nil {
		return err
	}

	decFile := filename + ".dec"
	err = os.WriteFile(decFile, plaintext, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Çözüldü:", decFile)
	return nil
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Kullanım: go run main.go [encrypt|decrypt] <dosya> <key(16,24,32 byte)>")
		return
	}

	mode := os.Args[1]
	filename := os.Args[2]
	key := []byte(os.Args[3])

	switch mode {
	case "encrypt":
		if err := encryptFile(filename, key); err != nil {
			fmt.Println("Hata:", err)
		}
	case "decrypt":
		if err := decryptFile(filename, key); err != nil {
			fmt.Println("Hata:", err)
		}
	default:
		fmt.Println("Geçersiz işlem. encrypt | decrypt kullanın.")
	}
}
``
/*
---

# ⚡ Kullanım

### 1. Şifreleme

```bash
go run main.go encrypt secret.txt examplekey123456
```

➡ `secret.txt.enc` oluşur.

### 2. Çözme

```bash
go run main.go decrypt secret.txt.enc examplekey123456
```

➡ `secret.txt.enc.dec` oluşur (orijinal dosya).

---

# ✅ Özellikler

* AES-GCM ile **güvenli şifreleme + bütünlük doğrulama** yapıyor.
* Key uzunluğu **16, 24 veya 32 byte** olabilir.
* Hem **şifreleme** hem **çözme** tek programda.

---

👉 İstersen ben bunu biraz daha geliştirip **şifreyi kullanıcıdan gizli şekilde (terminalde görünmeden) alacak** hale getirebilirim. İster misin?
Evet
*/
/*
Harika! aşağıya **parola gizli girilen**, paroladan **AES-256 anahtarı türeten (scrypt)** ve **AES-GCM** ile şifreleyip doğrulayan tam bir CLI verdim.
Dosya biçimi kendi içinde **sihirli başlık + sürüm + salt + nonce + ciphertext** içerir; bu sayede çözmede gerekli bilgileri dosyadan okur.

---

# 🔐 Özellikler

* Parola terminalde **gizli** (echo kapalı) alınır
* **scrypt** ile 32-byte AES-256 anahtarı türetilir (salt dosyada saklanır)
* **AES-GCM**: gizlilik + bütünlük/doğrulama
* Şifrelemede parola **iki kez** sorulur (yanlış yazımı önlemek için)
* Basit, tek dosya, çapraz platform (Windows/macOS/Linux)

> Not: `golang.org/x/term` ve `golang.org/x/crypto/scrypt` kullanır.
> Yüklemek için:
> `go get golang.org/x/term golang.org/x/crypto/scrypt`

---

## `main.go`
*/
``
package main

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/scrypt"
	"golang.org/x/term"
)

const (
	magic   = "AESG" // dosya sihri
	version = 1

	saltSize  = 16            // scrypt salt
	nonceSize = 12            // GCM nonce
	keyLen    = 32            // AES-256
)

var (
	// scrypt parametreleri: interaktif kullanım için makul
	scryptN = 1 << 15 // 32768
	scryptR = 8
	scryptP = 1
)

func deriveKeyFromPass(pass, salt []byte) ([]byte, error) {
	return scrypt.Key(pass, salt, scryptN, scryptR, scryptP, keyLen)
}

func readPassword(prompt string) ([]byte, error) {
	fmt.Fprint(os.Stderr, prompt)
	pw, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Fprintln(os.Stderr)
	return pw, err
}

func confirmPassword() ([]byte, error) {
	pw1, err := readPassword("Parola: ")
	if err != nil {
		return nil, err
	}
	pw2, err := readPassword("Parola (tekrar): ")
	if err != nil {
		zero(pw1)
		return nil, err
	}
	if !bytes.Equal(pw1, pw2) {
		zero(pw1); zero(pw2)
		return nil, errors.New("parolalar eşleşmiyor")
	}
	zero(pw2)
	return pw1, nil
}

func zero(b []byte) {
	for i := range b {
		b[i] = 0
	}
}

func encryptFile(inPath, outPath string) error {
	// parola al
	pw, err := confirmPassword()
	if err != nil {
		return err
	}
	defer zero(pw)

	// dosya oku
	plain, err := os.ReadFile(inPath)
	if err != nil {
		return err
	}

	// salt ve key üret
	salt := make([]byte, saltSize)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return err
	}
	key, err := deriveKeyFromPass(pw, salt)
	if err != nil {
		return err
	}
	defer zero(key)

	// AES-GCM
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	gcm, err := cipher.NewGCMWithNonceSize(block, nonceSize)
	if err != nil {
		return err
	}
	nonce := make([]byte, nonceSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	ciphertext := gcm.Seal(nil, nonce, plain, nil)

	// çıktı: | magic(4) | version(1) | salt(16) | nonce(12) | ciphertext |
	var buf bytes.Buffer
	buf.WriteString(magic)
	buf.WriteByte(byte(version))
	buf.Write(salt)
	buf.Write(nonce)
	buf.Write(ciphertext)

	// diske yaz
	if outPath == "" {
		outPath = inPath + ".enc"
	}
	if err := os.WriteFile(outPath, buf.Bytes(), 0644); err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Şifrelendi →", outPath)
	return nil
}

func decryptFile(inPath, outPath string) error {
	enc, err := os.ReadFile(inPath)
	if err != nil {
		return err
	}
	// başlık kontrol
	if len(enc) < 4+1+saltSize+nonceSize {
		return errors.New("geçersiz dosya: kısa")
	}
	if string(enc[:4]) != magic {
		return errors.New("geçersiz dosya: magic uyuşmuyor")
	}
	ver := enc[4]
	if ver != version {
		return fmt.Errorf("desteklenmeyen sürüm: %d", ver)
	}
	offset := 5
	salt := enc[offset : offset+saltSize]
	offset += saltSize
	nonce := enc[offset : offset+nonceSize]
	offset += nonceSize
	ciphertext := enc[offset:]
	if len(ciphertext) == 0 {
		return errors.New("geçersiz dosya: şifreli veri yok")
	}

	// parola al
	pw, err := readPassword("Parola: ")
	if err != nil {
		return err
	}
	defer zero(pw)

	key, err := deriveKeyFromPass(pw, salt)
	if err != nil {
		return err
	}
	defer zero(key)

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	gcm, err := cipher.NewGCMWithNonceSize(block, nonceSize)
	if err != nil {
		return err
	}

	plain, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return errors.New("çözme/kimlik doğrulama başarısız (parola yanlış veya dosya bozuk)")
	}

	if outPath == "" {
		// ".enc" sonekini düşürmeye çalış
		if len(inPath) > 4 && inPath[len(inPath)-4:] == ".enc" {
			outPath = inPath[:len(inPath)-4]
		} else {
			outPath = inPath + ".dec"
		}
	}
	if err := os.WriteFile(outPath, plain, 0644); err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Çözüldü →", outPath)
	return nil
}

func usage() {
	fmt.Fprintf(os.Stderr, `Kullanım:
  %s encrypt <girdi_dosyası> [çıktı_dosyası]
  %s decrypt <girdi_dosyası> [çıktı_dosyası]

Açıklama:
  - Şifreleme: parola iki kez istenir (gizli).
  - Çözme: parola bir kez istenir (gizli).
  - Anahtar, scrypt ile paroladan türetilir (N=32768, r=8, p=1, salt=16).
  
`, os.Args[0], os.Args[0])
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() < 2 {
		usage()
		return
	}

	cmd := flag.Arg(0)
	in := flag.Arg(1)
	out := ""
	if flag.NArg() >= 3 {
		out = flag.Arg(2)
	}

	switch cmd {
	case "encrypt":
		if err := encryptFile(in, out); err != nil {
			fmt.Fprintln(os.Stderr, "Hata:", err)
			os.Exit(1)
		}
	case "decrypt":
		if err := decryptFile(in, out); err != nil {
			fmt.Fprintln(os.Stderr, "Hata:", err)
			os.Exit(1)
		}
	default:
		usage()
		os.Exit(2)
	}
}

// Küçük yardımcı: okunabilir tam sayı yaz/oku (şimdilik kullanılmıyor ama geleceğe hazır)
func writeU32(w io.Writer, v uint32) {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, v)
	w.Write(b)
}
``
/*
---

## 🚀 Çalıştırma

```bash
# bağımlılıkları ekle
go get golang.org/x/term golang.org/x/crypto/scrypt

# derle
go build -o aesc aesc.go  # (dosya adın farklıysa uyarlayabilirsin)

# şifrele
./aesc encrypt secret.txt          # parola 2 kez istenir → secret.txt.enc

# çöz
./aesc decrypt secret.txt.enc      # parola istenir → secret.txt
```

---

## Notlar & İpuçları

* **Parola gücü** önemli: uzun ve rastgele bir cümle/parola tercih et.
* scrypt parametreleri (N,r,p) makinene göre **arttırılabilir** (daha güvenli, daha yavaş).
* Çıktı biçimi versiyonlu olduğundan gelecekte parametre/gcm tag uzunluğu vs. değiştirmek kolay.
* Bellekte gerçek “güvenli silme” garanti edilemez, ama `zero()` ile byte slice’ları en azından temizliyoruz.

İstersen bu aracı:

* **STDIN/STDOUT** üzerinden pipe ile çalışacak (streaming) hale,
* **parolayı dosyadan/çevre değişkeninden** alma,
* **argon2id** kullanacak şekilde
  geliştirebilirim. Hangisini istersin?
*/