# Go Programlama Dili - Operatörler

Bu doküman Go dilindeki tüm operatörleri **örneklerle** açıklamaktadır.

------------------------------------------------------------------------

## 1. Aritmetik Operatörler

``` go
x, y := 10, 3
fmt.Println(x + y) // 13
fmt.Println(x - y) // 7
fmt.Println(x * y) // 30
fmt.Println(x / y) // 3
fmt.Println(x % y) // 1
```

------------------------------------------------------------------------

## 2. Atama Operatörleri

``` go
x := 5
x += 3
fmt.Println(x) // 8

x -= 2
fmt.Println(x) // 6

x *= 4
fmt.Println(x) // 24

x /= 6
fmt.Println(x) // 4

x %= 3
fmt.Println(x) // 1
```

------------------------------------------------------------------------

## 3. Karşılaştırma Operatörleri

``` go
x, y := 10, 20
fmt.Println(x == y) // false
fmt.Println(x != y) // true
fmt.Println(x > y)  // false
fmt.Println(x < y)  // true
fmt.Println(x >= y) // false
fmt.Println(x <= y) // true
```

------------------------------------------------------------------------

## 4. Mantıksal Operatörler

``` go
a, b := true, false
fmt.Println(a && b) // false
fmt.Println(a || b) // true
fmt.Println(!a)     // false
```

------------------------------------------------------------------------

## 5. Bit Düzeyinde Operatörler

``` go
x, y := 6, 3 // 6 = 110, 3 = 011

fmt.Println(x & y)   // 2 (010)
fmt.Println(x | y)   // 7 (111)
fmt.Println(x ^ y)   // 5 (101)
fmt.Println(x &^ y)  // 4 (100)
fmt.Println(x << 1)  // 12 (1100)
fmt.Println(x >> 1)  // 3 (011)
```

------------------------------------------------------------------------

## 6. Pointer Operatörleri

``` go
x := 42
p := &x
fmt.Println(p)  // adres
fmt.Println(*p) // 42
```

------------------------------------------------------------------------

## 7. Özel Operatörler

### Kanal (`<-`)

``` go
ch := make(chan int)
go func() { ch <- 10 }()
val := <-ch
fmt.Println(val) // 10
```

### Slice genişletme (`...`)

``` go
nums := []int{1, 2, 3, 4}
fmt.Println(sum(nums...)) // 10
```

------------------------------------------------------------------------

## 8. ++ ve -- Operatörleri

``` go
x := 5
x++
fmt.Println(x) // 6

x--
fmt.Println(x) // 5
```
