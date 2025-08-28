// Go dilinde **`bufio`** paketi, giriş/çıkış (I/O) işlemlerini **buffer** (tampon) kullanarak daha verimli hale getiren bir pakettir.
// Normalde `os.File`, `net.Conn` gibi yapılarla doğrudan okuma/yazma yaptığında her işlem işletim sistemi çağrısı gerektirir (maliyetli olur). 
// `bufio`, araya bir tampon koyarak I/O işlemlerini hızlandırır.


//TAMPON NEDİR?
// Go programlama dilinde bufio.Writer, tamponlu (buffered) yazma işlemleri yapmak için kullanılan bir yapıdır. 
// "Tamponlu yazma" ifadesi, verilerin doğrudan hedefe (örneğin bir dosyaya veya bir ağ bağlantısına) yazılmak 
// yerine önce bir bellek tamponuna (buffer) yazılması, ardından bu tamponun dolması ya da açıkça boşaltılması durumunda hedefe yazılması anlamına gelir.

// Bellek tamponu (buffer), geçici olarak veri saklamak için kullanılan bir RAM alanıdır. 
// Veri bir yerden bir yere gönderilirken (örneğin klavyeden programa, programdan dosyaya, 
// programdan ağa) genellikle hemen işlenmez ya da gönderilmez; önce bu tampon bellekte toplanır, sonra işlenir ya da gönderilir.

// Daha Basit Tanımla:
// Bir buffer, verilerin geçici olarak beklediği bir ara durak gibidir.
// Özet:
// Buffer (tampon), verilerin geçici olarak bellekte saklandığı, performans ve verimlilik için kullanılan bir araçtır.



// ---

// ## 📦 `bufio` Paketindeki Temel Yapılar

// ### 1. **`bufio.Reader`**

// * Veri kaynaklarından (dosya, bağlantı vb.) tamponlu okuma yapar.
// * Küçük küçük okumalar yerine, belleğe büyük bir blok alır, sonra programın isteğine göre parçalar halinde verir.
// * Kullanımı: `bufio.NewReader(r io.Reader) *Reader`

// #### Önemli Metodlar:

// * `Read(p []byte) (n int, err error)` → Verilen byte dilimini doldurur.
// * `ReadByte() (byte, error)` → Tek byte okur.
// * `ReadString(delim byte) (string, error)` → Belirli bir ayraç karakterine kadar okur.
// * `ReadLine() (line []byte, isPrefix bool, err error)` → Satır okur.

// #### Örnek:

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("ornek.txt")
	defer file.Close()

	reader := bufio.NewReader(file)

	// İlk satırı oku
	line, _, _ := reader.ReadLine()
	fmt.Println("Satır:", string(line))

	// Noktaya kadar oku
	str, _ := reader.ReadString('.')
	fmt.Println("Noktaya kadar:", str)
}
```

---

### 2. **`bufio.Writer`**

* Veri kaynağına tamponlu yazma yapar.
* Verileri önce belleğe yazar, tampon dolunca veya `Flush()` çağrılınca dosyaya/bağlantıya yazar.
* Kullanımı: `bufio.NewWriter(w io.Writer) *Writer`

#### Önemli Metodlar:

* `Write(p []byte) (n int, err error)` → Byte dilimini yazar.
* `WriteString(s string) (int, error)` → String yazar.
* `WriteByte(c byte) error` → Tek byte yazar.
* `Flush() error` → Tampondaki veriyi hedefe gönderir.

#### Örnek:

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Create("cikti.txt")
	defer file.Close()

	writer := bufio.NewWriter(file)

	writer.WriteString("Merhaba bufio!\n")
	writer.WriteString("Bu yazı tamponda bekliyor...\n")

	// Eğer Flush() çağrılmazsa dosyaya yazılmayabilir
	writer.Flush()

	fmt.Println("Yazma tamamlandı.")
}
```

---

### 3. **`bufio.Scanner`**

* Satır satır veya özel kurallarla okuma yapmayı kolaylaştırır.
* Kullanımı: `bufio.NewScanner(r io.Reader) *Scanner`

#### Önemli Metodlar:

* `Scan() bool` → Bir sonraki token’ı okur (true/false döner).
* `Text() string` → Okunan veriyi string olarak döner.
* `Bytes() []byte` → Okunan veriyi byte dilimi olarak döner.
* `Split(splitFunc SplitFunc)` → Ayrıştırma kuralını değiştirir (örn. kelime kelime, satır satır).

#### Örnek (satır satır okuma):

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("ornek.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println("Satır:", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Hata:", err)
	}
}
```

#### Örnek (kelime kelime okuma):

```go
scanner := bufio.NewScanner(file)
scanner.Split(bufio.ScanWords)

for scanner.Scan() {
    fmt.Println("Kelime:", scanner.Text())
}
```

---

### 4. **`bufio.ReadWriter`**

* Hem `Reader` hem `Writer` özelliklerini birleştirir.
* Kullanımı: `bufio.NewReadWriter(r *Reader, w *Writer) *ReadWriter`

#### Örnek:

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Create("rw.txt")
	defer file.Close()

	rw := bufio.NewReadWriter(bufio.NewReader(file), bufio.NewWriter(file))

	rw.WriteString("Merhaba Dünya!\n")
	rw.Flush()

	fmt.Println("Yazıldı.")
}
```

---

## 🚀 Özet

* **`Reader`** → Okuma (satır, byte, string)
* **`Writer`** → Yazma (string, byte, buffer ile)
* **`Scanner`** → Daha kolay okuma (satır satır, kelime kelime)
* **`ReadWriter`** → Hem okuma hem yazma

---

👉 İstersen sana **tam bir uygulama** yapayım: Dosyadan satır satır okuyup, her satırın başına numara ekleyerek yeni bir dosyaya yazan bir program. İster misin?
