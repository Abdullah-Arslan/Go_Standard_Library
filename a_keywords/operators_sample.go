/*
Harika 👍 O zaman sana **tüm operatörler** için küçük ve çalıştırılabilir **Go örnekleri** hazırlayayım. Gruplara göre yazacağım, böylece daha kolay okunur olacak.

---

# 1. Aritmetik Operatörler
*/
``go
package main

import "fmt"

func main() {
    x, y := 10, 3

    fmt.Println("x + y =", x+y) // 13
    fmt.Println("x - y =", x-y) // 7
    fmt.Println("x * y =", x*y) // 30
    fmt.Println("x / y =", x/y) // 3 (tam sayı bölme!)
    fmt.Println("x % y =", x%y) // 1
}
``
/*
---

# 2. Atama Operatörleri
*/
``go
package main

import "fmt"

func main() {
    x := 5
    x += 3 // x = x + 3
    fmt.Println("x += 3 ->", x) // 8

    x -= 2
    fmt.Println("x -= 2 ->", x) // 6

    x *= 4
    fmt.Println("x *= 4 ->", x) // 24

    x /= 6
    fmt.Println("x /= 6 ->", x) // 4

    x %= 3
    fmt.Println("x %= 3 ->", x) // 1
}
``
/*
---

# 3. Karşılaştırma Operatörleri
*/
``go
package main

import "fmt"

func main() {
    x, y := 10, 20

    fmt.Println("x == y:", x == y) // false
    fmt.Println("x != y:", x != y) // true
    fmt.Println("x > y:", x > y)   // false
    fmt.Println("x < y:", x < y)   // true
    fmt.Println("x >= y:", x >= y) // false
    fmt.Println("x <= y:", x <= y) // true
}
``
/*
---

# 4. Mantıksal Operatörler
*/
``go
package main

import "fmt"

func main() {
    a, b := true, false

    fmt.Println("a && b:", a && b) // false
    fmt.Println("a || b:", a || b) // true
    fmt.Println("!a:", !a)         // false
}
``
/*
---

# 5. Bit Düzeyinde Operatörler
*/
``go
package main

import "fmt"

func main() {
    x, y := 6, 3 // 6 = 110, 3 = 011 (binary)

    fmt.Println("x & y =", x&y)   // 2 (010)
    fmt.Println("x | y =", x|y)   // 7 (111)
    fmt.Println("x ^ y =", x^y)   // 5 (101)
    fmt.Println("x &^ y =", x&^y) // 4 (100)

    fmt.Println("x << 1 =", x<<1) // 12 (1100)
    fmt.Println("x >> 1 =", x>>1) // 3 (011)
}
``
/*
---

# 6. Adres (Pointer) Operatörleri
*/
``go
package main

import "fmt"

func main() {
    x := 42
    p := &x // x’in adresi

    fmt.Println("x =", x)   // 42
    fmt.Println("p =", p)   // adres (ör. 0xc0000140a8)
    fmt.Println("*p =", *p) // 42 (adresteki değer)

    *p = 100
    fmt.Println("x =", x) // 100 (pointer üzerinden değişti)
}
``
/*
---

# 7. Özel Operatörler

### Kanal (`<-`)
*/
``go
package main

import "fmt"

func main() {
    ch := make(chan int)

    go func() {
        ch <- 10 // kanala veri gönder
    }()

    val := <-ch // kanaldan veri al
    fmt.Println("Kanaldan alınan değer:", val) // 10
}
``

//### Slice genişletme (`...`)

``go
package main

import "fmt"

func toplam(nums ...int) int {
    sum := 0
    for _, n := range nums {
        sum += n
    }
    return sum
}

func main() {
    sayilar := []int{1, 2, 3, 4}
    fmt.Println("Toplam:", toplam(sayilar...)) // 10
}
``
/*
---

# 8. ++ ve -- (Sadece postfix!)
*/
``go
package main

import "fmt"

func main() {
    x := 5
    x++ // x = x + 1
    fmt.Println("x++ ->", x) // 6

    x-- // x = x - 1
    fmt.Println("x-- ->", x) // 5
}
``
/*
---

👉 Böylece Go’daki **tüm operatörler** için örnekleri tek tek görmüş oldun.

İstersen ben bunları **tablo + örnek kod** halinde **Excel ve PDF** olarak da hazırlayabilirim, ki offline çalışırken daha rahat bakabilirsin. Bunu ister misin?
*/