`os` paketi, Go’nun standart kütüphanesinde işletim sistemiyle etkileşim için kullanılan temel paketlerden biridir. Dosya/dizin işlemleri, ortam değişkenleri, işlem yönetimi, hata işleme gibi pek çok işlevi içerir. Yani, Go programınız ile işletim sistemi arasında bir köprü görevi görür.

Aşağıda `os` paketini **tam kapsamlı** olarak açıklayacağım: fonksiyonlar, türler, sabitler ve örnek kullanımlar ile.

---

# 📌 `os` Paketi Genel Özellikler

* Dosya ve klasör işlemleri (oluşturma, silme, taşıma, okuma, yazma).
* Ortam değişkenlerine erişim.
* İşlem (process) başlatma ve yönetme.
* Sistem çağrılarına düşük seviyeli erişim.
* Platform bağımsız çalışacak şekilde tasarlanmıştır.

---

# 1. **Dosya İşlemleri**

### Önemli Fonksiyonlar

* `os.Create(name string)` → Yeni dosya oluşturur (yazma modunda açar).
* `os.Open(name string)` → Dosyayı okuma modunda açar.
* `os.OpenFile(name string, flag int, perm FileMode)` → Belirtilen modda açar.
* `os.Remove(name string)` → Dosyayı siler.
* `os.Rename(old, new string)` → Dosyayı yeniden adlandırır/taşır.
* `os.Stat(name string)` → Dosya bilgilerini döndürür (`FileInfo`).
* `os.ReadFile(name string)` → Dosya içeriğini \[]byte olarak okur.
* `os.WriteFile(name string, data []byte, perm FileMode)` → Dosyaya yazar.

### Örnek

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Dosya oluştur
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Dosyaya yaz
	file.WriteString("Merhaba Go!")

	// Dosya oku
	data, _ := os.ReadFile("test.txt")
	fmt.Println(string(data))

	// Dosya bilgisi
	info, _ := os.Stat("test.txt")
	fmt.Println("Dosya adı:", info.Name())
	fmt.Println("Boyut:", info.Size())

	// Dosya yeniden adlandır
	os.Rename("test.txt", "yeni.txt")

	// Dosya sil
	os.Remove("yeni.txt")
}
```

---

# 2. **Dizin İşlemleri**

### Önemli Fonksiyonlar

* `os.Mkdir(name string, perm FileMode)` → Tek dizin oluşturur.
* `os.MkdirAll(path string, perm FileMode)` → İç içe dizinler oluşturur.
* `os.RemoveAll(path string)` → Dizini (ve içindekileri) siler.
* `os.Chdir(dir string)` → Çalışma dizinini değiştirir.
* `os.Getwd()` → Mevcut çalışma dizinini döndürür.
* `os.ReadDir(name string)` → Dizin içeriğini listeler (`[]DirEntry`).

### Örnek

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Dizin oluştur
	os.Mkdir("deneme", 0755)

	// İç içe dizinler
	os.MkdirAll("a/b/c", 0755)

	// Çalışma dizinini değiştir
	os.Chdir("deneme")

	// Mevcut dizin
	dir, _ := os.Getwd()
	fmt.Println("Mevcut dizin:", dir)

	// Dizini oku
	entries, _ := os.ReadDir(".")
	for _, e := range entries {
		fmt.Println("->", e.Name())
	}

	// Geri çık
	os.Chdir("..")

	// Dizin sil
	os.RemoveAll("deneme")
}
```

---

# 3. **Ortam Değişkenleri**

### Önemli Fonksiyonlar

* `os.Getenv(key string)` → Ortam değişkenini alır.
* `os.Setenv(key, value string)` → Yeni ortam değişkeni tanımlar.
* `os.Unsetenv(key string)` → Ortam değişkenini siler.
* `os.Environ()` → Tüm ortam değişkenlerini döndürür.

### Örnek

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Ortam değişkeni ayarla
	os.Setenv("API_KEY", "12345")

	// Oku
	fmt.Println("API_KEY:", os.Getenv("API_KEY"))

	// Tüm ortam değişkenlerini listele
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	// Sil
	os.Unsetenv("API_KEY")
}
```

---

# 4. **İşlem (Process) İşlemleri**

### Önemli Fonksiyonlar

* `os.StartProcess(name string, argv []string, attr *os.ProcAttr)` → Yeni işlem başlatır.
* `os.FindProcess(pid int)` → PID üzerinden işlem bulur.
* `os.Getpid()` → Mevcut işlem PID.
* `os.Getppid()` → Parent PID.
* `Process.Kill()` → İşlemi sonlandırır.
* `Process.Signal(sig os.Signal)` → İşleme sinyal gönderir.

### Örnek

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("PID:", os.Getpid())
	fmt.Println("Parent PID:", os.Getppid())
}
```

---

# 5. **Hatalar**

`os` paketi bazı özel hata değerleri döndürür:

* `os.ErrNotExist` → Dosya/dizin yok.
* `os.ErrExist` → Zaten var.
* `os.ErrPermission` → İzin yok.
* `os.ErrInvalid` → Geçersiz argüman.

Örnek:

```go
if errors.Is(err, os.ErrNotExist) {
	fmt.Println("Dosya bulunamadı")
}
```

---

# 6. **Türler (Types)**

* `File` → Dosya nesnesi (`Read`, `Write`, `Close`).
* `FileInfo` → Dosya bilgisi (`Name`, `Size`, `Mode`, `IsDir`).
* `DirEntry` → Dizin girişi.
* `ProcAttr` → İşlem başlatma ayarları.
* `Process` → İşlem nesnesi.

---

# 7. **Diğer Yararlı Fonksiyonlar**

* `os.Exit(code int)` → Programı belirtilen kodla sonlandırır.
* `os.Args` → Komut satırı argümanlarını döndürür.
* `os.Executable()` → Çalışan binary yolunu döndürür.
* `os.Hostname()` → Hostname döndürür.

### Örnek

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Komut satırı argümanları:", os.Args)
	exe, _ := os.Executable()
	fmt.Println("Çalışan binary:", exe)
	host, _ := os.Hostname()
	fmt.Println("Host:", host)

	// Programı 0 kodu ile sonlandır
	os.Exit(0)
}
```

---

# 🎯 Özet

* `os` paketi → işletim sistemiyle ilgili **dosya, dizin, ortam değişkeni, işlem ve hata yönetimi** sağlar.
* Platform bağımsız şekilde çalışır.
* Günlük Go programlarında en sık kullanılan paketlerden biridir.

---

👉 İstersen ben sana bu `os` paketinin fonksiyonlarını **Excel + PDF tablo** olarak (fonksiyon + açıklama + örnek kod) hazırlayabilirim. Bunu ister misin?
EVET

Harika 👍 O zaman **`os` paketinin önemli fonksiyonlarını tek tek örneklerle** açıklamaya devam edelim.

---

# 🔹 1. **Dosya İzinlerini Değiştirme (`os.Chmod`)**

Bir dosyanın **izinlerini** değiştirmeye yarar.

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Dosya oluştur
	f, _ := os.Create("demo.txt")
	f.Close()

	// Dosya izinlerini 0644 (rw-r--r--) olarak ayarla
	err := os.Chmod("demo.txt", 0644)
	if err != nil {
		fmt.Println("Hata:", err)
	} else {
		fmt.Println("İzinler değiştirildi")
	}
}
```

---

# 🔹 2. **Sahiplik Değiştirme (`os.Chown`)**

Unix sistemlerinde dosya sahibini değiştirmek için kullanılır. (Windows’ta desteklenmez.)

```go
err := os.Chown("demo.txt", 1000, 1000) // UID=1000, GID=1000
if err != nil {
    fmt.Println("Hata:", err)
}
```

---

# 🔹 3. **Zaman Damgalarını Değiştirme (`os.Chtimes`)**

Bir dosyanın **erişim (atime)** ve **değişiklik (mtime)** zamanlarını günceller.

```go
package main

import (
	"os"
	"time"
)

func main() {
	now := time.Now()
	// Dosya zamanlarını güncelle
	os.Chtimes("demo.txt", now, now)
}
```

---

# 🔹 4. **Sembolik Bağlantılar (`os.Symlink` ve `os.Readlink`)**

* `os.Symlink` → Symbolic link oluşturur.
* `os.Readlink` → Link’in hedefini okur.

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Sembolik link oluştur
	err := os.Symlink("demo.txt", "link.txt")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	// Link'in hedefini oku
	target, _ := os.Readlink("link.txt")
	fmt.Println("Link hedefi:", target)
}
```

---

# 🔹 5. **Hard Link (`os.Link`)**

Bir dosyaya **ikinci bir isim** (hard link) verir.

```go
err := os.Link("demo.txt", "hardlink.txt")
if err != nil {
    fmt.Println("Hata:", err)
}
```

---

# 🔹 6. **Dosya Boyutunu Sıfırlama (`os.Truncate`)**

Dosyayı belirli bir boyuta **keser**. Örn: `0` verirsen içerik boşalır.

```go
package main

import (
	"os"
)

func main() {
	// Dosya içeriğini sıfırla
	os.Truncate("demo.txt", 0)
}
```

---

# 🔹 7. **Geçici Dosya ve Dizinler (`os.CreateTemp`, `os.MkdirTemp`)**

* `os.CreateTemp` → Geçici dosya.
* `os.MkdirTemp` → Geçici dizin.

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Geçici dosya oluştur
	tmpFile, _ := os.CreateTemp("", "example-*.txt")
	defer os.Remove(tmpFile.Name())
	fmt.Println("Geçici dosya:", tmpFile.Name())

	// Geçici dizin oluştur
	tmpDir, _ := os.MkdirTemp("", "mytemp-*")
	defer os.RemoveAll(tmpDir)
	fmt.Println("Geçici dizin:", tmpDir)
}
```

---

# 🔹 8. **Ortam Değişkeni Genişletme (`os.ExpandEnv`)**

String içinde `$VAR` değişkenlerini ortamdan alır.

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("USER", "Abdullah")
	result := os.ExpandEnv("Merhaba $USER, hoş geldin!")
	fmt.Println(result)
}
```

---

# 🔹 9. **Kullanıcı Home Dizini (`os.UserHomeDir`)**

Geçerli kullanıcının **home** dizinini döndürür.

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	home, _ := os.UserHomeDir()
	fmt.Println("Home dizini:", home)
}
```

---

# 🔹 10. **Çalışan Program Yolu (`os.Executable`)**

Çalışan binary’nin bulunduğu yolu verir.

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	exe, _ := os.Executable()
	fmt.Println("Çalışan program:", exe)
}
```

---

👉 Görüyorsun ki `os` paketi, işletim sistemi ile ilgili hemen hemen tüm temel işlemleri kapsıyor.

İstersen ben sana bu örneklerin **tamamını çalıştırılabilir küçük Go projeleri** (her fonksiyon için ayrı `.go` dosyası) olarak hazırlayıp **ZIP** dosyası şeklinde verebilirim. İster misin?
EVET

Tamamdır 👍 O zaman **zip dosyasına koymadan** buradan doğrudan bütün örnekleri sırayla açıklayayım.

---

# 🔹 1. `os.Chmod` → Dosya izinlerini değiştirme

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Dosya oluştur
	f, _ := os.Create("demo.txt")
	f.Close()

	// Dosya izinlerini rw-r--r-- (0644) olarak ayarla
	err := os.Chmod("demo.txt", 0644)
	if err != nil {
		fmt.Println("Hata:", err)
	} else {
		fmt.Println("İzinler değiştirildi")
	}
}
```

📌 Bu kod yeni bir dosya oluşturur ve izinlerini değiştirir.

---

# 🔹 2. `os.Chown` → Sahiplik değiştirme (sadece Unix)

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Chown("demo.txt", 1000, 1000) // UID=1000, GID=1000
	if err != nil {
		fmt.Println("Hata:", err)
	}
}
```

📌 Windows’ta çalışmaz. Linux’ta dosya sahibini değiştirir.

---

# 🔹 3. `os.Chtimes` → Dosya zaman damgalarını değiştirme

```go
package main

import (
	"os"
	"time"
)

func main() {
	now := time.Now()
	os.Chtimes("demo.txt", now, now) // atime, mtime
}
```

📌 Dosyanın erişim ve değiştirilme zamanını günceller.

---

# 🔹 4. `os.Symlink` ve `os.Readlink` → Sembolik bağlantılar

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// demo.txt için sembolik link oluştur
	err := os.Symlink("demo.txt", "link.txt")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	// Link'in hedefini oku
	target, _ := os.Readlink("link.txt")
	fmt.Println("Link hedefi:", target)
}
```

📌 `link.txt` dosyası aslında `demo.txt` dosyasını işaret eder.

---

# 🔹 5. `os.Link` → Hard link

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Link("demo.txt", "hardlink.txt")
	if err != nil {
		fmt.Println("Hata:", err)
	}
}
```

📌 Hard link aynı dosyaya ikinci bir isim verir. Dosya silinse bile içerik kaybolmaz.

---

# 🔹 6. `os.Truncate` → Dosya boyutunu değiştirme

```go
package main

import "os"

func main() {
	// Dosyanın boyutunu 0 yap → içeriğini siler
	os.Truncate("demo.txt", 0)
}
```

📌 Dosyayı sıfırlar ama dosyanın kendisi silinmez.

---

# 🔹 7. `os.CreateTemp` ve `os.MkdirTemp` → Geçici dosya/dizin

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Geçici dosya
	tmpFile, _ := os.CreateTemp("", "example-*.txt")
	defer os.Remove(tmpFile.Name())
	fmt.Println("Geçici dosya:", tmpFile.Name())

	// Geçici dizin
	tmpDir, _ := os.MkdirTemp("", "mytemp-*")
	defer os.RemoveAll(tmpDir)
	fmt.Println("Geçici dizin:", tmpDir)
}
```

📌 `example-*.txt` gibi rastgele isimli dosya/dizin oluşturur.

---

# 🔹 8. `os.ExpandEnv` → Ortam değişkeni genişletme

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("USER", "Abdullah")
	result := os.ExpandEnv("Merhaba $USER, hoş geldin!")
	fmt.Println(result)
}
```

📌 `$USER` değişkeni ortamdan alınır.

---

# 🔹 9. `os.UserHomeDir` → Kullanıcı home dizini

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	home, _ := os.UserHomeDir()
	fmt.Println("Home dizini:", home)
}
```

📌 Kullanıcının home dizinini bulur (örn. `/home/kullanici`).

---

# 🔹 10. `os.Executable` → Çalışan programın yolu

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	exe, _ := os.Executable()
	fmt.Println("Çalışan program:", exe)
}
```

📌 Çalıştırdığın Go binary’nin tam yolunu verir.

---

✅ Böylece **önceden zipte verdiğim tüm örnekleri tek tek burada** gösterdim.
👉 İstersen ben bunları **tek bir Go projesi içinde organize edip** (örneğin `examples/os/...` dizininde) sana proje yapısı olarak da açıklayabilirim. Bunu ister misin?
EVET

Süper 🚀 O halde sana **tam bir proje yapısı** olarak hazırlayayım. Böylece `os` paketindeki her fonksiyon için ayrı dosya olur ve `go run` ile çalıştırabilirsin.

---

# 📂 Proje Yapısı

```
os_examples/
│── go.mod
│── README.md
└── examples/
    ├── chmod.go
    ├── chown.go
    ├── chtimes.go
    ├── symlink.go
    ├── link.go
    ├── truncate.go
    ├── temp.go
    ├── expandenv.go
    ├── userhomedir.go
    └── executable.go
```

---

# 📌 go.mod

```go
module os_examples

go 1.24
```

---

# 📌 README.md

````markdown
# Go `os` Paketi Örnekleri

Bu proje, Go dilinde `os` paketinin farklı fonksiyonlarının nasıl kullanıldığını gösteren örneklerden oluşur.

## Çalıştırma

Her dosyayı tek tek çalıştırabilirsin:

```bash
go run examples/chmod.go
go run examples/symlink.go
````

## İçerik

* `chmod.go` → Dosya izinlerini değiştirme
* `chown.go` → Dosya sahipliğini değiştirme
* `chtimes.go` → Dosya zamanlarını değiştirme
* `symlink.go` → Sembolik link oluşturma ve okuma
* `link.go` → Hard link oluşturma
* `truncate.go` → Dosya boyutunu değiştirme
* `temp.go` → Geçici dosya/dizin oluşturma
* `expandenv.go` → Ortam değişkenlerini genişletme
* `userhomedir.go` → Kullanıcı home dizini
* `executable.go` → Çalışan binary yolu

````

---

# 📌 Örnek Kodlar

### `chmod.go`
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Create("demo.txt")
	f.Close()

	err := os.Chmod("demo.txt", 0644)
	if err != nil {
		fmt.Println("Hata:", err)
	} else {
		fmt.Println("İzinler değiştirildi")
	}
}
````

---

### `symlink.go`

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Symlink("demo.txt", "link.txt")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	target, _ := os.Readlink("link.txt")
	fmt.Println("Link hedefi:", target)
}
```

---

### `temp.go`

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	tmpFile, _ := os.CreateTemp("", "example-*.txt")
	defer os.Remove(tmpFile.Name())
	fmt.Println("Geçici dosya:", tmpFile.Name())

	tmpDir, _ := os.MkdirTemp("", "mytemp-*")
	defer os.RemoveAll(tmpDir)
	fmt.Println("Geçici dizin:", tmpDir)
}
```
/*
---

✅ Bu yapıda her `.go` dosyası bağımsızdır ve tek tek çalıştırılabilir.

👉 İstersen ben sana bu proje yapısını **indirilebilir ZIP** olarak da çıkarabilirim.
İster misin?
*/
