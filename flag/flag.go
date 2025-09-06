/*
Go’nun **`flag` paketini** en ince ayrıntısına kadar açıklayayım. Bu paket Go’nun standart kütüphanesinde bulunur ve **komut satırı argümanlarını (CLI flags)** kolayca okumamızı sağlar.

---

# 📖 Go `flag` Paketi

`import "flag"`

Go uygulamalarında, program çalıştırılırken verilen **komut satırı parametrelerini** okumak için kullanılır.
Örnek:
*/

``bash
go run main.go -port=8080 -debug=true -name=Ali
``
/*
---

## 📌 Genel Kullanım

* Program başlamadan önce flag’ler tanımlanır (`flag.String`, `flag.Int`, vs.).
* Ardından `flag.Parse()` çağrılır → komut satırındaki argümanlar işlenir.
* Değerler pointer olarak döner.

---

## 📌 Flag Türleri

### 1. String
*/

``go
name := flag.String("name", "default", "Kullanıcı adı")
``
/*
* İlk parametre: flag ismi (`-name`)
* İkincisi: varsayılan değer (`default`)
* Üçüncüsü: açıklama (help çıktısı için)

---

### 2. Int
*/

``go
port := flag.Int("port", 8080, "Sunucu portu")
``
/*
---

### 3. Bool
*/

``go
debug := flag.Bool("debug", false, "Debug modu")
``
/*
---

### 4. Float64
*/

``go
ratio := flag.Float64("ratio", 0.75, "Oran değeri")
``
/*
---

## 📌 Örnek Kullanım
*/
``go
package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "Guest", "Kullanıcı adı")
	port := flag.Int("port", 8080, "Sunucu portu")
	debug := flag.Bool("debug", false, "Debug modu")

	flag.Parse() // Argümanları işle

	fmt.Println("Name:", *name)
	fmt.Println("Port:", *port)
	fmt.Println("Debug:", *debug)
}
``

//Çalıştırma:

``bash
go run main.go -name=Ali -port=9090 -debug=true
``
/*
Çıktı:

```
Name: Ali
Port: 9090
Debug: true
```

---

## 📌 Özel Kullanım

### `flag.Var` ile Kendi Türünü Tanımla
*/

``go
type list []string

func (l *list) String() string {
	return fmt.Sprint(*l)
}

func (l *list) Set(value string) error {
	*l = append(*l, value)
	return nil
}

func main() {
	var langs list
	flag.Var(&langs, "lang", "Programlama dilleri")

	flag.Parse()
	fmt.Println("Languages:", langs)
}
``

//Çalıştırma:

``bash
go run main.go -lang=Go -lang=Python -lang=Rust
``
/*
Çıktı:

```
Languages: [Go Python Rust]
```

---

## 📌 Diğer Önemli Fonksiyonlar

* **flag.Args()** → Parse sonrası kalan “positional arguments” listesi
* **flag.Arg(i)** → Belirli index’teki positional argüman
* **flag.NArg()** → Positional argüman sayısı
* **flag.NFlag()** → Set edilmiş flag sayısı

Örnek:
*/
``go
flag.Parse()
fmt.Println("Kalan args:", flag.Args())
fmt.Println("Arg[0]:", flag.Arg(0))
``
/*
---

## 📌 Kullanıcıya Yardım (Help)

Go otomatik olarak `-h` veya `-help` parametresini destekler.
*/

``bash
go run main.go -h
``

//Örnek çıktı:

``
  -debug
        Debug modu
  -name string
        Kullanıcı adı (default "Guest")
  -port int
        Sunucu portu (default 8080)
``
/*
---

## 📌 flag.CommandLine

Varsayılan `FlagSet`. Yani `flag.String`, `flag.Int` gibi fonksiyonlar aslında global `flag.CommandLine` üzerine kayıt yapar.
Kendi `FlagSet` oluşturabilirsin:
*/

``go
fs := flag.NewFlagSet("myapp", flag.ExitOnError)
port := fs.Int("port", 8080, "Port numarası")
fs.Parse(os.Args[1:])
``

/*
---

## 📌 Özel FlagSet Kullanımı
*/

``go
fs := flag.NewFlagSet("server", flag.ContinueOnError)
host := fs.String("host", "localhost", "Host adresi")
_ = fs.Parse([]string{"-host", "127.0.0.1"})
fmt.Println("Host:", *host)
``
/*
---

## 📌 Özet

* `flag` paketi Go’da komut satırı argümanlarını okumak için kullanılır.
* Hazır türler: `String`, `Int`, `Bool`, `Float64`.
* `flag.Var` ile kendi tipini tanımlayabilirsin.
* `flag.Args`, `flag.NArg`, `flag.NFlag` gibi yardımcı fonksiyonlar bulunur.
* `FlagSet` kullanarak özel flag grupları oluşturabilirsin.

---

İstersen sana bunun da **expvar için yaptığım gibi bir “flag cheatsheet” PDF** hazırlayayım. İstiyor musun?
*/