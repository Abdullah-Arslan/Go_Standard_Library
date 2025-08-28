/*

Şimdi sana **Go dilinde `crypto/rc4` paketi**ni en temelden tüm detaylarıyla açıklayayım.

---

## 📌 `crypto/rc4` Paketi Nedir?

* `crypto/rc4` paketi, **Rivest Cipher 4 (RC4)** algoritmasını uygular.
* RC4 bir **stream cipher** (akış şifreleyici) algoritmasıdır.
* Çalışma mantığı:

  * Bir **anahtar (key)** kullanılarak bir pseudo-random keystream üretilir.
  * Bu keystream ile **XOR** işlemi yapılarak şifreleme ve çözme gerçekleşir.
* Aynı algoritma **hem şifreleme hem çözme** için kullanılır.

📌 **Uyarı**: RC4 artık **kriptografik olarak güvenli kabul edilmez**. Güvenlik kritik sistemlerde kullanılmamalıdır. Modern alternatifler: AES, ChaCha20.

---

## 📌 Paket Fonksiyonları

### 1. `rc4.NewCipher(key []byte) (*rc4.Cipher, error)`

* Verilen anahtar ile yeni bir RC4 cipher nesnesi oluşturur.
* Anahtar uzunluğu 1 ile 256 byte arasında olmalıdır.
*/
``go
c, err := rc4.NewCipher([]byte("secretkey"))
if err != nil {
	panic(err)
}
``
/*
---

### 2. `(*Cipher) XORKeyStream(dst, src []byte)`

* RC4 şifreleme / çözme işlemini yapar.
* `src` verisini alır, RC4 keystream ile XOR’lar ve sonucu `dst` içine yazar.
* Aynı fonksiyon hem şifreleme hem de çözme için çalışır.

📌 Önemli: `dst` ve `src` aynı slice olabilir (in-place işlem yapılabilir).

```go
cipher.XORKeyStream(dst, src)
```

---

## 📌 Basit Şifreleme / Çözme Örneği
*/
``go
package main

import (
	"crypto/rc4"
	"fmt"
)

func main() {
	key := []byte("gizli-anahtar")
	plaintext := []byte("Merhaba Dünya!")

	// Şifreleyici oluştur
	c, err := rc4.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Şifreleme
	ciphertext := make([]byte, len(plaintext))
	c.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("Şifreli veri: %x\n", ciphertext)

	// Çözme için yeniden cipher oluştur (aynı key ile)
	c2, _ := rc4.NewCipher(key)
	decrypted := make([]byte, len(ciphertext))
	c2.XORKeyStream(decrypted, ciphertext)

	fmt.Printf("Çözülen veri: %s\n", decrypted)
}
``
/*
### 📝 Çıktı (örnek):

```
Şifreli veri: 8a73b98fd1c6a0e9bcaf51
Çözülen veri: Merhaba Dünya!
```

---

## 📌 In-Place Kullanım

`dst` ve `src` aynı slice olabilir.
*/
``go
c, _ := rc4.NewCipher([]byte("anahtar"))
data := []byte("Hello RC4")

// Şifreleme (in-place)
c.XORKeyStream(data, data)
fmt.Printf("Şifreli: %x\n", data)

// Çözme (yine aynı data üzerinde)
c2, _ := rc4.NewCipher([]byte("anahtar"))
c2.XORKeyStream(data, data)
fmt.Printf("Çözülen: %s\n", data)
``
/*
---

## 📌 RC4 Özellikleri

* **Simetrik şifreleme**: Aynı key hem şifreleme hem çözme için kullanılır.
* **Stream cipher**: Veri akışını byte byte işler.
* **Hızlıdır**, ama modern kriptografi için **güvensizdir**.
* Özellikle TLS gibi protokollerde **yasaklanmıştır**.

---

## 📌 Özet

* `crypto/rc4` paketi, RC4 stream cipher implementasyonudur.
* Anahtar ile `NewCipher` oluşturulur.
* `XORKeyStream` ile şifreleme/çözme yapılır.
* RC4 günümüzde güvenli kabul edilmez, sadece eğitim/uyumluluk amaçlı kullanılmalı.

---

👉 İstersen ben sana **RC4 ile dosya şifreleme/çözme** örneği de yazabilirim (örn: `.txt` dosyasını şifreleyip geri açma). İstiyor musun?
EVET

Harika 👍 O zaman sana **RC4 ile dosya şifreleme ve çözme** örneğini yazayım.

Bu örnekte:

* Bir `.txt` dosyasını okuyacağız.
* RC4 ile şifreleyeceğiz ve `encrypted.txt` dosyasına yazacağız.
* Sonra `encrypted.txt` dosyasını açıp tekrar RC4 ile çözerek `decrypted.txt` dosyasına kaydedeceğiz.

---

## 📌 RC4 Dosya Şifreleme / Çözme Örneği
*/
``go
package main

import (
	"crypto/rc4"
	"fmt"
	"io/ioutil"
	"os"
)

func encryptFile(inputFile, outputFile string, key []byte) error {
	// Dosya oku
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}

	// RC4 şifreleyici oluştur
	c, err := rc4.NewCipher(key)
	if err != nil {
		return err
	}

	// Şifreleme (in-place de olurdu ama ayrı slice yapıyoruz)
	encrypted := make([]byte, len(data))
	c.XORKeyStream(encrypted, data)

	// Dosyaya yaz
	return ioutil.WriteFile(outputFile, encrypted, 0644)
}

func decryptFile(inputFile, outputFile string, key []byte) error {
	// Dosya oku
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}

	// RC4 çözme (aynı anahtar)
	c, err := rc4.NewCipher(key)
	if err != nil {
		return err
	}

	decrypted := make([]byte, len(data))
	c.XORKeyStream(decrypted, data)

	return ioutil.WriteFile(outputFile, decrypted, 0644)
}

func main() {
	key := []byte("super-secret-key")

	// Örnek plaintext dosyası oluştur
	_ = ioutil.WriteFile("plain.txt", []byte("Merhaba! Bu RC4 ile şifrelenmiş bir mesajdır."), 0644)

	// Şifreleme
	if err := encryptFile("plain.txt", "encrypted.txt", key); err != nil {
		fmt.Println("Şifreleme hatası:", err)
		os.Exit(1)
	}
	fmt.Println("plain.txt dosyası encrypted.txt olarak şifrelendi.")

	// Çözme
	if err := decryptFile("encrypted.txt", "decrypted.txt", key); err != nil {
		fmt.Println("Çözme hatası:", err)
		os.Exit(1)
	}
	fmt.Println("encrypted.txt dosyası decrypted.txt olarak çözüldü.")
}
``
/*
---

## 📌 Çalışma Adımları

1. `plain.txt` oluşturulur → içine mesaj yazılır.
2. `encrypted.txt` → RC4 ile şifrelenmiş binary içerik. (Düz okunamaz)
3. `decrypted.txt` → Orijinal mesaj geri gelir.

### 📝 Örnek Çalışma

```
plain.txt       → "Merhaba! Bu RC4 ile şifrelenmiş bir mesajdır."
encrypted.txt   → (binary/gibberish, okunmaz)
decrypted.txt   → "Merhaba! Bu RC4 ile şifrelenmiş bir mesajdır."
```

---

👉 İstersen ben bu örneği biraz daha geliştirip **komut satırından kullanılabilen bir RC4 dosya şifreleyici/çözücü** haline de getirebilirim (`go run main.go -e file.txt -o out.enc -k secret`).

Bunu da ister misin?
EVET
*/

/*
Süper 👍 O zaman sana **komut satırından kullanılabilen RC4 dosya şifreleyici/çözücü** yazayım.

Bu programı şu şekilde kullanabileceksin:

* **Şifreleme**
*/
``bash
go run main.go -mode=encrypt -in=plain.txt -out=secret.enc -key=superkey
``

//* **Çözme**

``bash
go run main.go -mode=decrypt -in=secret.enc -out=decrypted.txt -key=superkey
``
/*
---

## 📌 RC4 Command-Line File Encryptor
*/
``go
package main

import (
	"crypto/rc4"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// RC4 ile dosya işleme (hem şifreleme hem çözme)
func processFile(mode, inFile, outFile string, key []byte) error {
	// Girdi dosyasını oku
	data, err := ioutil.ReadFile(inFile)
	if err != nil {
		return fmt.Errorf("dosya okunamadı: %v", err)
	}

	// RC4 cipher oluştur
	c, err := rc4.NewCipher(key)
	if err != nil {
		return fmt.Errorf("cipher oluşturulamadı: %v", err)
	}

	// İşlem (XOR)
	result := make([]byte, len(data))
	c.XORKeyStream(result, data)

	// Çıktı dosyasına yaz
	err = ioutil.WriteFile(outFile, result, 0644)
	if err != nil {
		return fmt.Errorf("dosya yazılamadı: %v", err)
	}

	if mode == "encrypt" {
		fmt.Printf("✅ '%s' dosyası şifrelendi → '%s'\n", inFile, outFile)
	} else {
		fmt.Printf("✅ '%s' dosyası çözüldü → '%s'\n", inFile, outFile)
	}
	return nil
}

func main() {
	// Komut satırı parametreleri
	mode := flag.String("mode", "encrypt", "İşlem modu: encrypt | decrypt")
	inFile := flag.String("in", "", "Girdi dosyası")
	outFile := flag.String("out", "", "Çıktı dosyası")
	key := flag.String("key", "", "Şifreleme anahtarı")

	flag.Parse()

	// Parametre kontrolü
	if *inFile == "" || *outFile == "" || *key == "" {
		fmt.Println("Kullanım: go run main.go -mode=encrypt|decrypt -in=girdi.txt -out=çıktı.txt -key=anahtar")
		os.Exit(1)
	}

	// İşlemi yap
	err := processFile(*mode, *inFile, *outFile, []byte(*key))
	if err != nil {
		fmt.Println("Hata:", err)
		os.Exit(1)
	}
}
``
/*
---

## 📌 Kullanım Örneği

1. `plain.txt` oluştur (örnek mesaj yaz):

```
Merhaba! Bu bir RC4 şifreleme testi.
```
*/
//2. **Şifrele**

``bash
go run main.go -mode=encrypt -in=plain.txt -out=secret.enc -key=superkey
``

3. **Çöz**

``bash
go run main.go -mode=decrypt -in=secret.enc -out=decrypted.txt -key=superkey
``
/*
4. `decrypted.txt` içeriği orijinal mesajla aynı olur. ✅

---

👉 İstersen ben bu programı biraz daha geliştirip **base64 çıktı opsiyonu** da ekleyebilirim (örn: şifreli dosya binary değil, okunabilir ASCII olsun). İster misin?
*/