/*
Go’nun **`fmt` paketi**, en çok kullanılan paketlerden biridir. Bu paket **formatlı giriş/çıkış (I/O) işlemleri** için kullanılır. C dilindeki `printf` ailesine çok benzer. Konsola yazı yazmak, string formatlamak, input almak gibi işlerin temelini oluşturur.

Şimdi sana **`fmt` paketini tüm fonksiyonlarıyla** (print, scan, format, hata yönetimi) açıklayayım ve örneklerle göstereyim.

---

# 📖 Go `fmt` Paketi

`import "fmt"`

* Yazdırma (output) fonksiyonları: `Print`, `Printf`, `Println`, `Fprint`, `Sprint`, ...
* Okuma (input) fonksiyonları: `Scan`, `Scanf`, `Scanln`, `Fscan`, `Sscan`, ...
* Format string’lerde **`%` ile başlayan verb’ler** kullanılır.

---

## 📌 Yazdırma Fonksiyonları

### 1. `Print`

Yan yana verilen argümanları yazdırır, otomatik boşluk koyar.
*/
``go
fmt.Print("Merhaba", "Dünya") 
// Çıktı: MerhabaDünya
``
/*
### 2. `Println`

Her argüman arasına boşluk koyar ve sonunda `\n` ekler.
*/
``go
fmt.Println("Merhaba", "Dünya")
// Çıktı: Merhaba Dünya
``
/*
### 3. `Printf`

Format string kullanarak yazdırır.
*/
``go
name := "Ali"
age := 30
fmt.Printf("Ad: %s, Yaş: %d\n", name, age)
// Çıktı: Ad: Ali, Yaş: 30
``
/*
### 4. `Sprint`, `Sprintf`, `Sprintln`

String döner (ekrana yazmaz).
*/
``go
msg := fmt.Sprintf("Toplam: %d", 42)
fmt.Println(msg)
``
/*
### 5. `Fprint`, `Fprintf`, `Fprintln`

Bir `io.Writer` (ör. dosya) içine yazar.
*/
``go
file, _ := os.Create("out.txt")
defer file.Close()
fmt.Fprintln(file, "Dosyaya yazıldı")
``
/*
---

## 📌 Okuma Fonksiyonları

### 1. `Scan`

Kullanıcıdan boşluklarla ayrılmış input alır.
*/
``go
var name string
var age int
fmt.Scan(&name, &age)
// Input: "Ali 30"
// Çıktı: name="Ali", age=30
``
/*
### 2. `Scanf`

Formatlı input alır.
*/
``go
var name string
var age int
fmt.Scanf("%s %d", &name, &age)
``
/*
### 3. `Scanln`

Boşluklarla ayırır ama satır sonuna kadar okur.
*/
``go
var city string
fmt.Scanln(&city)
``
/*
### 4. `Fscan`, `Fscanf`, `Fscanln`

`io.Reader` üzerinden okur (örneğin dosya).
*/
``go
file, _ := os.Open("data.txt")
defer file.Close()
var n int
fmt.Fscan(file, &n)
``
/*
### 5. `Sscan`, `Sscanf`, `Sscanln`

String içinden okur.
*/
``go
var a, b int
fmt.Sscanf("10 20", "%d %d", &a, &b)
``
/*
---

## 📌 Format Verb’leri (En Çok Kullanılanlar)

### Genel:

* `%v` → Varsayılan format (her şey için)
* `%+v` → Struct alan isimleriyle
* `%#v` → Go syntax’ına uygun gösterim
* `%T` → Tipini yazdırır
*/
``go
type User struct{ Name string; Age int }
u := User{"Ali", 30}
fmt.Printf("%v\n", u)   // {Ali 30}
fmt.Printf("%+v\n", u)  // {Name:Ali Age:30}
fmt.Printf("%#v\n", u)  // main.User{Name:"Ali", Age:30}
fmt.Printf("%T\n", u)   // main.User
``
/*
### String:

* `%s` → Normal string
* `%q` → Çift tırnaklı string
* `%x` → Hex
*/
``go
s := "Go"
fmt.Printf("%s\n", s)  // Go
fmt.Printf("%q\n", s)  // "Go"
``
/*
### Sayılar:

* `%d` → Ondalık (int)
* `%b` → Binary
* `%x` → Hexadecimal
* `%o` → Octal
* `%f` → Ondalık (float)
* `%e` → Bilimsel gösterim
*/
``go
n := 42
pi := 3.14159
fmt.Printf("%d %b %x\n", n, n, n)     // 42 101010 2a
fmt.Printf("%f %.2f %e\n", pi, pi, pi) // 3.141590 3.14 3.141590e+00
``
/*
### Mantıksal:

* `%t` → `true`/`false`

---

## 📌 Hata Yönetimi

`fmt.Errorf` ile formatlı hata oluşturulur.
*/
``go
err := fmt.Errorf("kullanıcı bulunamadı: %s", "Ali")
fmt.Println(err) // kullanıcı bulunamadı: Ali
``

//Go 1.13+ ile `%w` kullanarak hata sarma (wrap) yapılabilir:

``go
baseErr := errors.New("veritabanı hatası")
err := fmt.Errorf("işlem başarısız: %w", baseErr)
``
/*
---

## 📌 Özel Writer Kullanımı
*/

``go
buf := &bytes.Buffer{}
fmt.Fprintf(buf, "Merhaba %s", "Go")
fmt.Println(buf.String()) // Merhaba Go
``
/*
---

## 📌 Özet

* **Yazdırma:** `Print`, `Printf`, `Println`, `Sprint`, `Fprint`
* **Okuma:** `Scan`, `Scanf`, `Scanln`, `Sscan`, `Fscan`
* **Format Verb’leri:** `%v`, `%s`, `%d`, `%f`, `%T` …
* **Hata yönetimi:** `fmt.Errorf` ve `%w`

---

Bütün bunları sana **expvar ve flag için yaptığım gibi bir PDF cheatsheet** halinde özetlememi ister misin?
*/