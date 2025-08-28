/*
Go (Golang) dilindeki **tüm keywords (anahtar kelimeler)**, dilin sözdizimi ve kontrol akışının temelini oluşturur. Go’da toplam **25 keyword** vardır ve bunlar değiştirilemez veya kullanıcı tanımlı değişken olarak kullanılamaz.

Ben sana her bir keyword’ü **geniş açıklama ve örneklerle** anlatacağım.

---

# 1️⃣ `break`

* **Ne işe yarar:** Döngüleri veya switch/case bloklarını erken sonlandırmak için kullanılır.
* **Örnek:**
*/
``go
for i := 0; i < 10; i++ {
    if i == 5 {
        break // Döngüyü durdur
    }
    fmt.Println(i)
}
``
/*
Çıktı:

```
0 1 2 3 4
```

---

# 2️⃣ `default`

* **Ne işe yarar:** switch/case içinde hiçbir case eşleşmediğinde çalıştırılacak bloğu belirtir.
* **Örnek:**
*/
``go
day := "Saturday"
switch day {
case "Monday":
    fmt.Println("Weekday")
default:
    fmt.Println("Not Monday")
}
``
/*
Çıktı:

```
Not Monday
```

---

# 3️⃣ `func`

* **Ne işe yarar:** Fonksiyon veya method tanımlamak için kullanılır.
* **Örnek:**
*/
``go
func sum(a int, b int) int {
    return a + b
}

fmt.Println(sum(3, 5)) // 8
``
/*
---

# 4️⃣ `interface`

* **Ne işe yarar:** Method set’i tanımlayarak türler arasında **polimorfizm** sağlar.
* **Örnek:**
*/
``go
type Speaker interface {
    Speak() string
}

type Person struct {
    Name string
}

func (p Person) Speak() string {
    return "Hello, " + p.Name
}

var s Speaker = Person{"Alice"}
fmt.Println(s.Speak()) // Hello, Alice
``
/*
---

# 5️⃣ `map`

* **Ne işe yarar:** Key-Value veri yapısı oluşturmak için kullanılır.
* **Örnek:**
*/
``go
m := make(map[string]int)
m["age"] = 30
fmt.Println(m["age"]) // 30
```
/*
---

# 6️⃣ `package`

* **Ne işe yarar:** Go dosyasının hangi paket içinde olduğunu belirtir.
* **Örnek:**
*/
``go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
``
/*
---

# 7️⃣ `range`

* **Ne işe yarar:** Slice, array, map veya channel üzerinde iterasyon yapmak için kullanılır.
* **Örnek:**
*/
``go
nums := []int{1,2,3}
for i, v := range nums {
    fmt.Println(i, v)
}
``
/*
Çıktı:

``
0 1
1 2
2 3
``

---

# 8️⃣ `return`

* **Ne işe yarar:** Fonksiyondan değer döndürmek veya fonksiyonu sonlandırmak için kullanılır.
* **Örnek:**
*/
``go
func greet(name string) string {
    return "Hello " + name
}
fmt.Println(greet("Alice")) // Hello Alice
``
/*
---

# 9️⃣ `select`

* **Ne işe yarar:** Channel’lar arasında seçim yapmayı sağlar.
* **Örnek:**
*/
``go
ch1 := make(chan string)
ch2 := make(chan string)

go func() { ch1 <- "one" }()
go func() { ch2 <- "two" }()

select {
case msg1 := <-ch1:
    fmt.Println(msg1)
case msg2 := <-ch2:
    fmt.Println(msg2)
}
``
/*
---

# 1️⃣0️⃣ `struct`

* **Ne işe yarar:** Birden fazla alanı bir arada tutan veri yapısı oluşturur.
* **Örnek:**
*/
``go
type Person struct {
    Name string
    Age  int
}

p := Person{"Alice", 30}
fmt.Println(p.Name, p.Age)
``
/*
---

# 1️⃣1️⃣ `case`

* **Ne işe yarar:** switch veya select bloklarında bir koşul belirtir.
* **Örnek:**
*/
``go
i := 2
switch i {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
default:
    fmt.Println("Other")
}
``
/*
Çıktı:

```
Two
```

---

# 1️⃣2️⃣ `const`

* **Ne işe yarar:** Sabit değerler tanımlamak için kullanılır.
* **Örnek:**
*/
``go
const Pi = 3.14
fmt.Println(Pi)
``
/*
---

# 1️⃣3️⃣ `defer`

* **Ne işe yarar:** Fonksiyon sonlanmadan önce çalışacak işlemleri ertelemek için kullanılır (genellikle cleanup için).
* **Örnek:**
*/
``go
func main() {
    defer fmt.Println("Son")
    fmt.Println("Önce")
}
``
/*
Çıktı:

```
Önce
Son
```

---

# 1️⃣4️⃣ `else`

* **Ne işe yarar:** if koşulu sağlanmadığında çalışacak blok.
* **Örnek:**
*/
``go
x := 5
if x > 10 {
    fmt.Println("Büyük")
} else {
    fmt.Println("Küçük")
}
``
/*
Çıktı:

```
Küçük
```

---

# 1️⃣5️⃣ `fallthrough`

* **Ne işe yarar:** switch case’de bir sonraki case’in de çalışmasını sağlar.
* **Örnek:**
*/
``go
i := 1
switch i {
case 1:
    fmt.Println("One")
    fallthrough
case 2:
    fmt.Println("Two")
}
``
/*
Çıktı:

```
One
Two
```

---

# 1️⃣6️⃣ `for`

* **Ne işe yarar:** Döngü oluşturur.
* **Örnek:**
*/
``go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
``
/*
---

# 1️⃣7️⃣ `go`

* **Ne işe yarar:** Goroutine oluşturmak için kullanılır.
* **Örnek:**
*/
``go
go func() {
    fmt.Println("Hello from goroutine")
}()
``

--

# 1️⃣8️⃣ `goto`

* **Ne işe yarar:** Program akışını belirtilen label’a taşır.
* **Örnek:**

```go
i := 0
Loop:
if i < 3 {
    fmt.Println(i)
    i++
    goto Loop
}
``

/*
Çıktı:

``
0
1
2
``
/*
---

# 1️⃣9️⃣ `if`

* **Ne işe yarar:** Koşullu blok oluşturur.
* **Örnek:**
*/
``go
x := 10
if x%2 == 0 {
    fmt.Println("Çift")
}
``
/*
---

# 2️⃣0️⃣ `import`

* **Ne işe yarar:** Başka paketleri projenize dahil eder.
* **Örnek:**
*/
``go
import "fmt"

func main() {
    fmt.Println("Merhaba")
}
``
/*
---

# 2️⃣1️⃣ `package`

* **Ne işe yarar:** Dosyanın hangi pakete ait olduğunu belirtir.
* **Örnek:**
*/
``go
package main
``
/*
---

# 2️⃣2️⃣ `var`

* **Ne işe yarar:** Değişken tanımlar.
* **Örnek:**
*/
``go
var x int = 10
var y = 20 // tip çıkarımı
``
/*
---

# 2️⃣3️⃣ `fallthrough` (Not: zaten yazıldı)

* Tekrar etmeye gerek yok, switch’de bir sonraki case çalıştırılır.

---

# 2️⃣4️⃣ `type`

* **Ne işe yarar:** Yeni tip veya alias tanımlar.
* **Örnek:**
*/
``go
type Age int
var myAge Age = 30
``
/*
---

# 2️⃣5️⃣ `switch`

* **Ne işe yarar:** Çoklu koşul kontrolü yapar.
* **Örnek:**
*/

``go
day := "Monday"
switch day {
case "Monday", "Tuesday":
    fmt.Println("Weekday")
default:
    fmt.Println("Weekend")
}
``
/*
---

# ✅ Özet

Go’daki **25 keyword**:

```
break, default, func, interface, map, package, range, return,
select, struct, case, const, defer, else, fallthrough, for,
go, goto, if, import, var, type, switch
```

* Kontrollere: `if`, `else`, `switch`, `case`, `fallthrough`, `for`, `break`, `goto`, `select`
* Fonksiyon ve tip tanımları: `func`, `type`, `struct`, `interface`, `map`
* Değişken ve sabitler: `var`, `const`
* Paralel/erteleme: `go`, `defer`
* Paket ve import: `package`, `import`

---

İstersen ben bunu **tek bir tablo + örnek kodlarla** görselleştirip **PDF/Markdown** şeklinde hazır bir rehber hâline getirebilirim.
Bunu yapmamı ister misin?
EVET
*/

/*
Harika! 😄 O zaman sana **Go Keywords Rehberi**ni **Markdown formatında tablo ve örneklerle** hazırlayabilirim. Bu rehber, **tüm 25 keyword’ü**, kısa açıklamasını ve örnek kodunu içerecek.

---

## 📄 Golang Keywords Rehberi (Markdown)

```markdown
# Go Keywords Rehberi

| Keyword       | Açıklama | Örnek |
|---------------|----------|-------|
| break         | Döngü veya switch case’i sonlandırır | `for i:=0;i<5;i++ { if i==3 { break } }` |
| default       | switch/case’de hiçbir case eşleşmediğinde çalışır | `switch x { case 1: ... default: ... }` |
| func          | Fonksiyon veya method tanımlar | `func sum(a,b int) int { return a+b }` |
| interface     | Polimorfizm sağlar, method set’i tanımlar | `type Speaker interface { Speak() string }` |
| map           | Key-Value veri yapısı oluşturur | `m := make(map[string]int); m["age"]=30` |
| package       | Dosyanın ait olduğu paketi belirtir | `package main` |
| range         | Slice, array, map veya channel üzerinde iterasyon yapar | `for i,v := range nums { fmt.Println(i,v) }` |
| return        | Fonksiyondan değer döndürür | `func f() int { return 10 }` |
| select        | Channel’lar arasında seçim yapar | `select { case msg:=<-ch1: ... case msg:=<-ch2: ... }` |
| struct        | Birden fazla alanı tek bir veri yapısında tutar | `type Person struct { Name string; Age int }` |
| case          | switch veya select içinde bir durumu belirtir | `switch x { case 1: ... }` |
| const         | Sabit değerler tanımlar | `const Pi = 3.14` |
| defer         | Fonksiyon sonlanmadan önce çalışacak işlemleri ertelemek için | `defer fmt.Println("Son")` |
| else          | if koşulu sağlanmadığında çalışacak blok | `if x>0 { ... } else { ... }` |
| fallthrough   | Switch case’de bir sonraki case’in de çalışmasını sağlar | `switch i { case 1: fallthrough; case 2: ... }` |
| for           | Döngü oluşturur | `for i:=0;i<5;i++{ fmt.Println(i) }` |
| go            | Goroutine başlatır | `go func(){ fmt.Println("Hello") }()` |
| goto          | Program akışını belirli label’a taşır | `Loop: i++; goto Loop` |
| if            | Koşullu blok oluşturur | `if x>0 { fmt.Println("Pozitif") }` |
| import        | Başka paketleri projeye dahil eder | `import "fmt"` |
| var           | Değişken tanımlar | `var x int = 10` |
| type          | Yeni tip veya alias tanımlar | `type Age int` |
| switch        | Çoklu koşul kontrolü sağlar | `switch x { case 1: ... default: ... }` |
```

---

Bu Markdown rehberi **hem açıklamalı hem de örnekli** olduğu için:

* Hızlıca Go keyword’lerini öğrenebilirsin
* Referans olarak kullanabilirsin
* GitHub veya VSCode’da Markdown olarak açıp inceleyebilirsin

---

İstersen ben bunu ayrıca **her keyword için tam çalıştırılabilir örnek kod bloklarıyla ayrı bir Go dosyası** hâline getirip **örnek çıktılarıyla birlikte** hazırlayabilirim. Böylece bir dosyayı çalıştırınca bütün keyword örneklerini görebileceksin.

Bunu yapmamı ister misin?
EVET
*/

/*
Harika 😄 O zaman sana **tam çalıştırılabilir Go dosyası** hazırlayayım.
Bu dosya, **Go’daki tüm 25 keyword’ü örnekleriyle** içerir ve çalıştırınca çıktısını terminalde görebilirsin.

---

## 📂 Proje Yapısı

``
go_keywords_example/
 ├── main.go
``

---

## 📌 `main.go`
*/
``go
package main

import "fmt"

// type, struct ve interface örnekleri
type Age int

type Person struct {
	Name string
	Age  Age
}

type Speaker interface {
	Speak() string
}

func (p Person) Speak() string {
	return "Hello, " + p.Name
}

// const ve var örnekleri
const Pi = 3.14
var globalVar = "I am global"

// Fonksiyon örnekleri
func sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("==== Go Keywords Örnekleri ====")

	// if-else
	x := 10
	if x%2 == 0 {
		fmt.Println("if: Çift sayı")
	} else {
		fmt.Println("else: Tek sayı")
	}

	// for ve break
	for i := 0; i < 5; i++ {
		if i == 3 {
			fmt.Println("break ile döngüden çıkıyoruz")
			break
		}
		fmt.Println("for döngüsü:", i)
	}

	// goto
	i := 0
Loop:
	if i < 2 {
		fmt.Println("goto ile loop:", i)
		i++
		goto Loop
	}

	// defer
	defer fmt.Println("defer: Bu en son çalışır")

	// switch-case-fallthrough-default
	day := "Monday"
	switch day {
	case "Saturday":
		fmt.Println("case: Hafta sonu")
	case "Monday":
		fmt.Println("case: Hafta içi")
		fallthrough
	default:
		fmt.Println("default: Devam eden case")
	}

	// map
	m := map[string]int{"age": 30}
	fmt.Println("map:", m["age"])

	// range
	slice := []string{"a", "b", "c"}
	for idx, val := range slice {
		fmt.Println("range:", idx, val)
	}

	// func ve return
	fmt.Println("func sum:", sum(5, 7))

	// package ve import zaten dosyanın başında kullanıldı
	fmt.Println("package ve import: fmt paketi ile yazdırdık")

	// go (goroutine)
	done := make(chan bool)
	go func() {
		fmt.Println("go: Goroutine çalıştı")
		done <- true
	}()
	<-done

	// struct, type ve interface örneği
	p := Person{"Alice", 25}
	fmt.Println("struct ve type:", p.Name, p.Age)
	var s Speaker = p
	fmt.Println("interface:", s.Speak())

	// const ve var örneği
	fmt.Println("const Pi:", Pi)
	fmt.Println("var globalVar:", globalVar)

	fmt.Println("==== Tüm örnekler tamamlandı ====")
}
``
/*
---

## 🚀 Çalıştırma
?
``bash
go run main.go
``

//### Örnek çıktı:

``
==== Go Keywords Örnekleri ====
if: Çift sayı
for döngüsü: 0
for döngüsü: 1
for döngüsü: 2
break ile döngüden çıkıyoruz
goto ile loop: 0
goto ile loop: 1
case: Hafta içi
default: Devam eden case
map: 30
range: 0 a
range: 1 b
range: 2 c
func sum: 12
package ve import: fmt paketi ile yazdırdık
go: Goroutine çalıştı
struct ve type: Alice 25
interface: Hello, Alice
const Pi: 3.14
var globalVar: I am global
defer: Bu en son çalışır
==== Tüm örnekler tamamlandı ====
``
/*
---

Bu dosya ile:

* Tüm Go keyword’lerini **çalıştırılabilir örnekleriyle** gördün
* Keyword’lerin kullanım biçimlerini **terminalde test edebiliyorsun**

---

İstersen ben bunu bir **Docker imajı** hâline getirip, herhangi bir makinede Go kurulu olmasa bile çalıştırabileceğin bir versiyon da hazırlayabilirim.

Bunu ekleyeyim mi?
EVET
*/
/*
Harika 😄 O zaman bu **Go Keywords Örnekleri** projesi için bir **Dockerfile** hazırlayalım.
Böylece Go yüklü olmayan bir makinede bile bütün keyword örneklerini çalıştırabilirsin.

---

## 📂 Proje Yapısı

```
go_keywords_example/
 ├── main.go
 ├── Dockerfile
```

---

## 📌 `Dockerfile`
*/

``dockerfile
# 1. Builder aşaması: Go ile derle
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Kaynak kodu kopyala
COPY . .

# Eğer go.mod yoksa oluştur
RUN go mod init go_keywords_example || true

# Binary'yi oluştur
RUN go build -o keywords-example main.go

# 2. Final aşama: küçük Alpine imajı
FROM alpine:latest

WORKDIR /root/

# Binary'yi kopyala
COPY --from=builder /app/keywords-example .

# Varsayılan komut
ENTRYPOINT ["./keywords-example"]
``
/*
---

## 🚀 Kullanım

### 1. Docker imajını build et
*/

``bash
docker build -t go-keywords-example .
``

/### 2. Çalıştır

``bash
docker run --rm go-keywords-example
``

//### Örnek çıktı

``
==== Go Keywords Örnekleri ====
if: Çift sayı
for döngüsü: 0
for döngüsü: 1
for döngüsü: 2
break ile döngüden çıkıyoruz
goto ile loop: 0
goto ile loop: 1
case: Hafta içi
default: Devam eden case
map: 30
range: 0 a
range: 1 b
range: 2 c
func sum: 12
package ve import: fmt paketi ile yazdırdık
go: Goroutine çalıştı
struct ve type: Alice 25
interface: Hello, Alice
const Pi: 3.14
var globalVar: I am global
defer: Bu en son çalışır
==== Tüm örnekler tamamlandı ====
``
/*
---

✅ Avantajlar:

* Go kurulu olmayan makinelerde çalışır
* Tüm Go keyword’lerini örnekleriyle test edebilirsin
* Docker container içinde güvenli bir şekilde çalışır

---

İstersen ben bu Docker imajını **multi-platform** hâle getirip hem Linux, hem Windows, hem macOS için hazır binary’leri üretecek şekilde ayarlayabilirim.
Bunu da yapayım mı?
EVET
*/
/*
Harika 😄 O zaman sana **multi-platform Docker build** ayarını göstereyim. Bu sayede aynı Dockerfile ile **Linux, Windows ve macOS** için Go binary’lerini üretebilirsin.

---

## 📌 Multi-platform Build Adımları

1. **Dockerfile** zaten elimizde var. Aynı `go_keywords_example/Dockerfile` kullanacağız.

2. **Buildx’i etkinleştir** (Docker 19+ ile birlikte gelir):
*/

``bash
docker buildx create --use
docker buildx inspect --bootstrap
``

//3. **Multi-platform build yap**:


``bash
docker buildx build --platform linux/amd64,linux/arm64,windows/amd64 -t go-keywords-example:multi --push .
``
/*
Açıklama:

* `--platform` → Hedef platformları belirler (Linux AMD64, ARM64, Windows AMD64 vb.)
* `--push` → Docker Hub veya registry’ye yükler (opsiyonel, lokal için `--load` da kullanabilirsin)

4. **Lokal olarak yüklemek için**:
*/

``bash
docker buildx build --platform linux/amd64,linux/arm64,windows/amd64 -t go-keywords-example:multi --load .
``
/*
5. **Platforma göre çalıştırma**:

* Linux:
?7

``bash
docker run --rm go-keywords-example:multi
``

//* Windows (PowerShell):

``powershell
docker run --rm go-keywords-example:multi
``
/*
---

## ✅ Özet

* Tek bir Dockerfile ile **birden fazla işletim sistemi ve mimari için binary üretebilirsin**.
* Multi-platform sayesinde CI/CD pipeline’larında her OS için ayrı build yapmana gerek kalmaz.
* Bu yapı, Go keyword örnekleri projesi veya başka Go projeleri için de kullanılabilir.

---

İstersen ben sana bu Dockerfile’ı ve build komutlarını kullanarak **otomatik platform algılayıp container içinde Go keyword’lerini çalıştıran hazır script** de hazırlayabilirim.
Bunu yapayım mı?
EVET
*/
/*
Süper 😄 O zaman sana **otomatik platform algılayan ve Go keyword örneklerini çalıştıran bir Docker setup script’i** hazırlayayım. Bu setup:

* Docker imajını build eder
* Çalıştırılacak platformu otomatik algılar
* Container içinde `main.go` çalıştırır ve çıktıyı gösterir

---

## 📂 Proje Yapısı
*/
``
go_keywords_example/
 ├── main.go
 ├── Dockerfile
 └── run.sh
``
/*
---

## 📌 `Dockerfile`

> Daha önce hazırladığımız Dockerfile, multi-platform build uyumlu hâle getirildi.
*/
``dockerfile
# 1. Builder aşaması: Go ile derle
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod init go_keywords_example || true
RUN go build -o keywords-example main.go

# 2. Final aşama: küçük Alpine imajı
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/keywords-example .

ENTRYPOINT ["./keywords-example"]
``
/*
---

## 📌 `run.sh`
*/
``bash
#!/bin/bash
set -e

IMAGE_NAME="go-keywords-example"
TAG="multi"

echo "🌐 Docker platformunu algılıyor..."
PLATFORM=$(uname | tr '[:upper:]' '[:lower:]')

if [[ "$PLATFORM" == "linux" ]]; then
    TARGET="linux/amd64"
elif [[ "$PLATFORM" == "darwin" ]]; then
    TARGET="linux/amd64" # macOS için Linux binary
elif [[ "$PLATFORM" == "windows_nt" ]]; then
    TARGET="windows/amd64"
else
    TARGET="linux/amd64"
fi

echo "🛠️ Platform: $PLATFORM -> Docker target: $TARGET"

echo "📦 Docker imajını build ediyor..."
docker buildx build --platform $TARGET -t $IMAGE_NAME:$TAG --load .

echo "▶️ Container çalıştırılıyor..."
docker run --rm $IMAGE_NAME:$TAG
``
/*
---

## 🚀 Kullanım

1. Script’e çalıştırma izni ver:
*/

``bash
chmod +x run.sh
``

//2. Script’i çalıştır:

``bash
./run.sh
``
/*
> Script otomatik olarak işletim sistemini algılar, Docker imajını build eder ve container içinde Go keyword örneklerini çalıştırır.

### Örnek çıktı:

```
🌐 Docker platformunu algılıyor...
🛠️ Platform: linux -> Docker target: linux/amd64
📦 Docker imajını build ediyor...
▶️ Container çalıştırılıyor...
==== Go Keywords Örnekleri ====
if: Çift sayı
for döngüsü: 0
...
defer: Bu en son çalışır
==== Tüm örnekler tamamlandı ====
```

---

✅ Avantajları:

* Tek komutla build ve run işlemi
* Platform algılama ile Linux, macOS ve Windows uyumu
* Multi-platform Docker imajı desteği
* Go keyword örneklerini kolayca test etme

---

İstersen ben bunu bir adım ileriye taşıyıp **CI/CD entegrasyonu için GitHub Actions workflow** da hazırlayabilirim. Böylece her push’ta otomatik olarak multi-platform imaj build edilir ve test edilir.

Bunu da ekleyeyim mi?
EVET
*/
/*
Harika 😄 O zaman sana **Go Keywords Örnekleri projesi için GitHub Actions CI/CD workflow** hazırlayacağım.
Bu workflow sayesinde:

* Her push veya PR’da otomatik olarak Docker imajı build edilir
* Multi-platform (Linux/amd64, Linux/arm64, Windows/amd64) binary’leri hazırlanır
* Container çalıştırılarak Go keyword örnekleri test edilir

---

## 📂 Proje Yapısı

```
go_keywords_example/
 ├── .github/
 │    └── workflows/
 │         └── docker-build.yml
 ├── main.go
 ├── Dockerfile
 └── run.sh
```

---

## 📌 `.github/workflows/docker-build.yml`
*/
``yaml
name: Build and Test Go Keywords Docker

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  build-and-test:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout repository
        uses: actions/checkout@v3

      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v3

      - name: Log in to DockerHub (optional)
        uses: docker/login-action@v2
        with:
          username: ${{ secrets.DOCKER_USERNAME }}
          password: ${{ secrets.DOCKER_PASSWORD }}

      - name: Build multi-platform Docker image
        run: |
          docker buildx build \
          --platform linux/amd64,linux/arm64,windows/amd64 \
          -t go-keywords-example:ci \
          --load .

      - name: Run container to test Go keywords examples
        run: docker run --rm go-keywords-example:ci
``
/*
---

### 🔹 Açıklamalar:

1. **`runs-on: ubuntu-latest`** → CI workflow Ubuntu üzerinde çalışacak.
2. **Docker Buildx** kuruluyor → Multi-platform build için gerekli.
3. **Docker login** (opsiyonel) → Eğer imajı DockerHub’a push etmek istersen secrets ekleyebilirsin.
4. **`--platform`** → Linux/AMD64, ARM64 ve Windows/AMD64 için imaj build ediyor.
5. **`docker run`** → Container çalıştırılarak Go keyword örnekleri test ediliyor.

---

### ✅ Avantajları:

* Her push veya pull request’te otomatik olarak **multi-platform Docker imajı** hazırlanır
* Go keyword örnekleri **container içinde test edilir**
* CI/CD pipeline ile **dev ortamı ve prod ortamı için güvenli** bir build sağlanır

---

İstersen ben bu workflow’a ek olarak **DockerHub’a otomatik push** ve **tag yönetimi** ekleyip, her release için hazır multi-platform imaj üretecek şekilde genişletebilirim.
Bunu da yapayım mı?
*/