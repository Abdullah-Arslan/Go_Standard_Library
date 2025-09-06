/*
**Go (Golang) `errors` paketi**ni en ince ayrıntısına kadar açıklayayım.

---

# 📌 `errors` Paketi Nedir?

Go’da **`errors`** paketi, hata (error) yönetimi için kullanılan temel standart kütüphanedir.
Bu paket sayesinde:

* Yeni hatalar oluşturabiliriz.
* Hataları sarmalayabilir (wrap) ve zincir halinde tutabiliriz.
* Belirli bir hatayı kontrol edebiliriz (`Is`, `As`).
* Hata mesajlarını kolayca özelleştirebiliriz.

`error` Go’da **yerleşik (built-in) bir interface**’tir:
*/
``go
type error interface {
    Error() string
}
``

/*
Herhangi bir tip bu `Error()` metodunu uygularsa, bir `error` olarak kullanılabilir.

---

# 📌 `errors` Paketindeki Fonksiyonlar

`errors` paketinin içindeki başlıca fonksiyonlar şunlardır:

1. **`errors.New`**
2. **`errors.Unwrap`**
3. **`errors.Is`**
4. **`errors.As`**
5. **`errors.Join`** (Go 1.20+)

Şimdi hepsini detaylı örneklerle anlatalım.

---

## 1. `errors.New`

Yeni bir hata oluşturur.
*/

``go
package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("bir hata oluştu")
	fmt.Println(err) // bir hata oluştu
}
``

/*
> Bu en temel hata oluşturma yöntemidir. Çoğunlukla sabit hata mesajlarında kullanılır.

---

## 2. `errors.Unwrap`

Bir hatanın sarmalanmış (wrapped) halinden **altındaki gerçek hatayı çıkarmak** için kullanılır.
Genellikle `fmt.Errorf` ile `%w` kullanılarak yapılan sarma işleminde işe yarar.
*/

``go
package main

import (
	"errors"
	"fmt"
)

func main() {
	baseErr := errors.New("temel hata")
	wrappedErr := fmt.Errorf("işlem başarısız: %w", baseErr)

	fmt.Println(wrappedErr)                  // işlem başarısız: temel hata
	fmt.Println(errors.Unwrap(wrappedErr))  // temel hata
}
``
/*
---

## 3. `errors.Is`

Bir hatanın belirli bir hata ile aynı olup olmadığını kontrol eder.
**Sarmalanmış (wrapped) hatalarda** bile çalışır.
*/

``go
package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("kayıt bulunamadı")

func getUser(id int) error {
	if id == 0 {
		return fmt.Errorf("getUser hatası: %w", ErrNotFound)
	}
	return nil
}

func main() {
	err := getUser(0)

	if errors.Is(err, ErrNotFound) {
		fmt.Println("Hata: Kayıt bulunamadı")
	} else if err != nil {
		fmt.Println("Başka bir hata:", err)
	}
}
``
/*
---

## 4. `errors.As`

Bir hatanın **tipini dönüştürerek (cast)** kontrol etmeye yarar.
Özellikle özel hata tipleri ile kullanılır.
*/

``go
package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Code int
	Msg  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("kod: %d, mesaj: %s", e.Code, e.Msg)
}

func riskyOperation() error {
	return &MyError{Code: 404, Msg: "veri bulunamadı"}
}

func main() {
	err := riskyOperation()

	var myErr *MyError
	if errors.As(err, &myErr) {
		fmt.Println("Özel hata yakalandı:", myErr.Code, myErr.Msg)
	} else {
		fmt.Println("Bilinmeyen hata:", err)
	}
}
``
/*
---

## 5. `errors.Join` (Go 1.20+)

Birden fazla hatayı **tek bir hata objesinde birleştirmek** için kullanılır.
Sonradan `errors.Is` ve `errors.As` ile bu hataların her biri ayrı ayrı kontrol edilebilir.
*/
``go
package main

import (
	"errors"
	"fmt"
)

var (
	ErrDB   = errors.New("veritabanı hatası")
	ErrConn = errors.New("bağlantı hatası")
)

func main() {
	err := errors.Join(ErrDB, ErrConn)

	fmt.Println("Toplu hata:", err)

	if errors.Is(err, ErrDB) {
		fmt.Println("DB hatası mevcut")
	}
	if errors.Is(err, ErrConn) {
		fmt.Println("Bağlantı hatası mevcut")
	}
}
``
/*
---

# 📌 Özel Hata Tipleri (Custom Errors)

Go’da `errors.New` dışında kendi hata tipimizi tanımlayabiliriz.
*/

``go
package main

import (
	"fmt"
)

type DivideError struct {
	Dividend int
	Divisor  int
}

func (e *DivideError) Error() string {
	return fmt.Sprintf("Bölme hatası: %d / %d", e.Dividend, e.Divisor)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivideError{Dividend: a, Divisor: b}
	}
	return a / b, nil
}

func main() {
	_, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	}
}
``
/*
---

# 📌 Özet

| Fonksiyon       | Açıklama                                           | Örnek Kullanım                |
| --------------- | -------------------------------------------------- | ----------------------------- |
| `errors.New`    | Basit hata oluşturur                               | `err := errors.New("hata")`   |
| `errors.Unwrap` | Sarmalanmış hatanın içindeki gerçek hatayı çıkarır | `errors.Unwrap(err)`          |
| `errors.Is`     | Bir hata belirli bir hata ile aynı mı?             | `errors.Is(err, ErrNotFound)` |
| `errors.As`     | Hatanın tipini dönüştürür                          | `errors.As(err, &myErr)`      |
| `errors.Join`   | Birden fazla hatayı tek hata olarak döner          | `errors.Join(err1, err2)`     |

---

👉 Yani `errors` paketi, Go’da **hata yönetiminin çekirdeği**dür. Basit hatalardan, sarmalanmış hatalara, özel hata tiplerinden, toplu hatalara kadar hepsini yönetmemizi sağlar.

---

İstersen sana **gerçek bir mini proje** hazırlayayım (ör: kullanıcı kayıt sistemi, DB hatası, doğrulama hatası gibi farklı error türleriyle). İstiyor musun?
EVET
*/
/*
Harika! O zaman sana **Go’da `errors` paketini kullanarak mini bir kullanıcı kayıt sistemi** örneği hazırlayayım. Bu projede farklı hata türlerini, sarmalama ve özel hata kullanımını göreceğiz.

---

# 📌 Mini Proje: Kullanıcı Kayıt Sistemi

**Senaryo:**

* Kullanıcı adı ve yaş bilgisi ile kayıt yapılacak.
* Hatalar:

  1. Boş kullanıcı adı (`ErrEmptyUsername`)
  2. Geçersiz yaş (`ErrInvalidAge`)
  3. Kullanıcı zaten kayıtlı (`ErrUserExists`)

Bu hataları özel tipler ve `errors` paketi ile yöneteceğiz.

---

## 1. Kod
*/

``go
package main

import (
	"errors"
	"fmt"
)

// Özel hata tipleri
var (
	ErrEmptyUsername = errors.New("kullanıcı adı boş olamaz")
	ErrInvalidAge    = errors.New("yaş 0'dan büyük olmalı")
	ErrUserExists    = errors.New("kullanıcı zaten kayıtlı")
)

// User yapısı
type User struct {
	Username string
	Age      int
}

// Basit "veritabanı" (map)
var userDB = make(map[string]User)

// Kayıt fonksiyonu
func registerUser(username string, age int) error {
	// Hata kontrolleri
	if username == "" {
		return ErrEmptyUsername
	}
	if age <= 0 {
		return ErrInvalidAge
	}
	if _, exists := userDB[username]; exists {
		return ErrUserExists
	}

	// Kayıt işlemi
	userDB[username] = User{Username: username, Age: age}
	return nil
}

// Kayıt işlemini yönet
func main() {
	users := []User{
		{"Alice", 25},
		{"Bob", -5},       // Hatalı yaş
		{"", 30},          // Boş kullanıcı adı
		{"Alice", 40},     // Zaten kayıtlı
		{"Charlie", 22},   // Başarılı kayıt
	}

	for _, u := range users {
		err := registerUser(u.Username, u.Age)
		if err != nil {
			// Hata tipine göre mesaj
			if errors.Is(err, ErrEmptyUsername) {
				fmt.Printf("Hata: Kullanıcı adı boş! (User: %+v)\n", u)
			} else if errors.Is(err, ErrInvalidAge) {
				fmt.Printf("Hata: Geçersiz yaş! (User: %+v)\n", u)
			} else if errors.Is(err, ErrUserExists) {
				fmt.Printf("Hata: Kullanıcı zaten var! (User: %+v)\n", u)
			} else {
				fmt.Printf("Bilinmeyen hata: %v\n", err)
			}
		} else {
			fmt.Printf("Başarıyla kayıt edildi: %+v\n", u)
		}
	}

	fmt.Println("\nKayıtlı kullanıcılar:")
	for _, u := range userDB {
		fmt.Printf("- %+v\n", u)
	}
}
``
/*
---

## 2. Çıktı Örneği

```
Başarıyla kayıt edildi: {Username:Alice Age:25}
Hata: Geçersiz yaş! (User: {Username:Bob Age:-5})
Hata: Kullanıcı adı boş! (User: {Username: Age:30})
Hata: Kullanıcı zaten var! (User: {Username:Alice Age:40})
Başarıyla kayıt edildi: {Username:Charlie Age:22}

Kayıtlı kullanıcılar:
- {Username:Alice Age:25}
- {Username:Charlie Age:22}
```

---

## 3. Açıklamalar

1. **Hata yönetimi**

   * `errors.Is` ile hangi hatayla karşılaştığımızı kontrol ediyoruz.
   * Bu sayede farklı hata tiplerine farklı işlemler uygulayabiliriz.

2. **Basit "veritabanı"**

   * Go’da map kullanarak kullanıcıları depoladık.
   * Gerçek projelerde bu DB yerine MySQL, PostgreSQL veya MongoDB kullanabilirsin.

3. **Özel hata mesajları**

   * Hata mesajları kullanıcı dostu ve tip kontrollü.

4. **Genişletilebilirlik**

   * Yeni hata tipleri veya kullanıcı doğrulama kuralları kolayca eklenebilir.

---

İstersen ben bunu **bir adım ileri taşıyıp**, `errors.Wrap` mantığıyla **hataları sarmalayarak hangi fonksiyonda oluştuğunu da gösterecek** bir sürümünü hazırlayabilirim. Bu, Go’da profesyonel hata yönetimi pratiğine çok yakın olur.

Bunu yapmamı ister misin?
EVET
*/
/*
Harika! O zaman projeyi bir adım ileri taşıyalım ve **hataları sarmalayarak (wrap) hangi fonksiyonda oluştuğunu** gösterecek hale getirelim. Bu, Go 1.13+ ile gelen `%w` ve `errors.Unwrap`/`errors.Is` mantığını kullanacak.

---

# 📌 Geliştirilmiş Kullanıcı Kayıt Sistemi (Hataları Sarmalama)

**Özellikler:**

* Hatalar hangi fonksiyonda oluştuğunu içerir.
* Sarmalanmış hataları çözümleyerek (`errors.Unwrap`) alt hata türüne erişiriz.
* Hata mesajı zincirini yazdırabiliriz.

---

## 1. Kod
*/
``go
package main

import (
	"errors"
	"fmt"
)

// Özel hata tipleri
var (
	ErrEmptyUsername = errors.New("kullanıcı adı boş olamaz")
	ErrInvalidAge    = errors.New("yaş 0'dan büyük olmalı")
	ErrUserExists    = errors.New("kullanıcı zaten kayıtlı")
)

// User yapısı
type User struct {
	Username string
	Age      int
}

// Basit "veritabanı" (map)
var userDB = make(map[string]User)

// Fonksiyonlar
func validateUser(u User) error {
	if u.Username == "" {
		return fmt.Errorf("validateUser: %w", ErrEmptyUsername)
	}
	if u.Age <= 0 {
		return fmt.Errorf("validateUser: %w", ErrInvalidAge)
	}
	return nil
}

func checkUserExists(u User) error {
	if _, exists := userDB[u.Username]; exists {
		return fmt.Errorf("checkUserExists: %w", ErrUserExists)
	}
	return nil
}

func registerUser(u User) error {
	// Kullanıcı doğrulama
	if err := validateUser(u); err != nil {
		return fmt.Errorf("registerUser: %w", err)
	}

	// Mevcut kullanıcı kontrolü
	if err := checkUserExists(u); err != nil {
		return fmt.Errorf("registerUser: %w", err)
	}

	// Kayıt işlemi
	userDB[u.Username] = u
	return nil
}

// Hata zincirini yazdıran yardımcı fonksiyon
func printErrorChain(err error) {
	fmt.Println("Hata zinciri:")
	for err != nil {
		fmt.Println(" -", err)
		err = errors.Unwrap(err)
	}
}

func main() {
	users := []User{
		{"Alice", 25},
		{"Bob", -5},      // Hatalı yaş
		{"", 30},         // Boş kullanıcı adı
		{"Alice", 40},    // Zaten kayıtlı
		{"Charlie", 22},  // Başarılı kayıt
	}

	for _, u := range users {
		err := registerUser(u)
		if err != nil {
			// Hata tipine göre mesaj
			if errors.Is(err, ErrEmptyUsername) {
				fmt.Printf("Hata (kullanıcı adı boş): %+v\n", u)
			} else if errors.Is(err, ErrInvalidAge) {
				fmt.Printf("Hata (geçersiz yaş): %+v\n", u)
			} else if errors.Is(err, ErrUserExists) {
				fmt.Printf("Hata (kullanıcı zaten var): %+v\n", u)
			} else {
				fmt.Printf("Bilinmeyen hata: %+v\n", err)
			}

			printErrorChain(err)
			fmt.Println()
		} else {
			fmt.Printf("Başarıyla kayıt edildi: %+v\n", u)
		}
	}

	fmt.Println("\nKayıtlı kullanıcılar:")
	for _, u := range userDB {
		fmt.Printf("- %+v\n", u)
	}
}
``
/*
---

## 2. Çıktı Örneği

```
Başarıyla kayıt edildi: {Username:Alice Age:25}
Hata (geçersiz yaş): {Username:Bob Age:-5}
Hata zinciri:
 - registerUser: validateUser: yaş 0'dan büyük olmalı
 - validateUser: yaş 0'dan büyük olmalı

Hata (kullanıcı adı boş): {Username: Age:30}
Hata zinciri:
 - registerUser: validateUser: kullanıcı adı boş olamaz
 - validateUser: kullanıcı adı boş olamaz

Hata (kullanıcı zaten var): {Username:Alice Age:40}
Hata zinciri:
 - registerUser: checkUserExists: kullanıcı zaten kayıtlı
 - checkUserExists: kullanıcı zaten kayıtlı

Başarıyla kayıt edildi: {Username:Charlie Age:22}

Kayıtlı kullanıcılar:
- {Username:Alice Age:25}
- {Username:Charlie Age:22}
```

---

## 3. Açıklamalar

1. **Hataları sarmalama**

   * `fmt.Errorf("context: %w", err)` ile hatayı sarmalıyoruz.
   * Böylece hangi fonksiyonda oluştuğu da hata mesajında görünüyor.

2. **Hata zinciri**

   * `errors.Unwrap` kullanarak tüm zinciri yazdırabiliyoruz.
   * Bu, özellikle büyük projelerde **hata kaynağını hızlıca bulmak** için çok faydalı.

3. **`errors.Is` ile kontrol**

   * Sarmalanmış hatalarda bile alt hata tipini kontrol edebiliyoruz.
   * Bu sayede farklı hatalara özel işlemler yapılabiliyor.

---

Bu yapı, Go projelerinde **profesyonel hata yönetimi için standart bir yaklaşımdır**.
İstersen bir sonraki adımda bunu **`errors.Join` ile birden fazla hata toplamak** ve tek seferde yönetmek için de geliştirebiliriz.

Bunu da ekleyelim mi?
*/