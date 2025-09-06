/*
Go’nun **`go/token` paketindeki Operators & Delimiters** kısmını **tek tek** açıklayayım.

Go dilinde operatörler ve ayırıcılar (operators & delimiters), kaynak kodun en küçük sözdizimsel parçalarıdır.
Bunlar `token.Token` tipinde sabitler olarak tanımlanır ve parser tarafından tanınır.

---

# 📦 Operators (Operatörler)

Operatörler, Go’da aritmetik, karşılaştırma, mantıksal işlemler gibi görevler için kullanılır.
*/

### 🔹 Aritmetik Operatörler

| Token | Sembol | Açıklama                    | Örnek                |
| ----- | ------ | --------------------------- | -------------------- |
| `ADD` | `+`    | Toplama, string birleştirme | `3 + 2`, `"a" + "b"` |
| `SUB` | `-`    | Çıkarma, negatif işaret     | `10 - 3`, `-x`       |
| `MUL` | `*`    | Çarpma, pointer dereference | `4 * 5`, `*ptr`      |
| `QUO` | `/`    | Bölme                       | `10 / 2`             |
| `REM` | `%`    | Mod (kalan)                 | `10 % 3`             |

---

### 🔹 Artırma / Azaltma

| Token | Sembol | Açıklama                                       | Örnek |
| ----- | ------ | ---------------------------------------------- | ----- |
| `INC` | `++`   | 1 artırma (sadece ifade olarak kullanılabilir) | `x++` |
| `DEC` | `--`   | 1 azaltma                                      | `y--` |

> 📌 Not: Go’da `++` ve `--` sadece **statement** olarak kullanılabilir, `z = x++` gibi bir kullanım yoktur.

---

### 🔹 Karşılaştırma Operatörleri

| Token | Sembol | Açıklama            | Örnek    |
| ----- | ------ | ------------------- | -------- |
| `EQL` | `==`   | Eşit mi?            | `a == b` |
| `NEQ` | `!=`   | Eşit değil mi?      | `a != b` |
| `LSS` | `<`    | Küçük mü?           | `x < y`  |
| `LEQ` | `<=`   | Küçük veya eşit mi? | `x <= y` |
| `GTR` | `>`    | Büyük mü?           | `x > y`  |
| `GEQ` | `>=`   | Büyük veya eşit mi? | `x >= y` |

---

### 🔹 Mantıksal Operatörler

| Token  | Sembol | Açıklama        | Örnek            |                |          |   |          |
| ------ | ------ | --------------- | ---------------- | -------------- | -------- | - | -------- |
| `LAND` | `&&`   | Mantıksal VE    | `x > 0 && y > 0` |                |          |   |          |
| `LOR`  | \`     |                 | \`               | Mantıksal VEYA | \`x == 1 |   | y == 2\` |
| `NOT`  | `!`    | Mantıksal DEĞİL | `!done`          |                |          |   |          |

---

### 🔹 Bit Operatörleri

| Token     | Sembol | Açıklama                | Örnek          |          |
| --------- | ------ | ----------------------- | -------------- | -------- |
| `AND`     | `&`    | Bitwise AND             | `a & b`        |          |
| `OR`      | \`     | \`                      | Bitwise OR     | `a \| b` |
| `XOR`     | `^`    | Bitwise XOR             | `a ^ b`        |          |
| `SHL`     | `<<`   | Bit kaydırma sola       | `1 << 3` (`8`) |          |
| `SHR`     | `>>`   | Bit kaydırma sağa       | `8 >> 2` (`2`) |          |
| `AND_NOT` | `&^`   | Bit temizleme (AND NOT) | `a &^ b`       |          |

---

### 🔹 Atama Operatörleri

| Token    | Sembol | Açıklama       | Örnek     |
| -------- | ------ | -------------- | --------- |
| `ASSIGN` | `=`    | Atama          | `x = 5`   |
| `DEFINE` | `:=`   | Kısa tanımlama | `x := 10` |

📌 Ayrıca tüm aritmetik ve bitwise operatörlerin kısaltılmış atamaları vardır:
`+=`, `-=`, `*=`, `/=`, `%=` … `&=`, `|=`, `^=`, `<<=`, `>>=`, `&^=`

---

# 📦 Delimiters (Ayırıcılar)

Go’da kodun yapısını belirleyen sembollerdir.

| Token    | Sembol | Açıklama                                           | Örnek                 |
| -------- | ------ | -------------------------------------------------- | --------------------- |
| `LPAREN` | `(`    | Parantez aç                                        | `f(x, y)`             |
| `RPAREN` | `)`    | Parantez kapat                                     | `f(x, y)`             |
| `LBRACK` | `[`    | Köşeli parantez aç                                 | `arr[0]`              |
| `RBRACK` | `]`    | Köşeli parantez kapat                              | `arr[0]`              |
| `LBRACE` | `{`    | Süslü parantez aç                                  | `func main() { ... }` |
| `RBRACE` | `}`    | Süslü parantez kapat                               | `func main() { ... }` |
| `COMMA`  | `,`    | Virgül                                             | `a, b := 1, 2`        |
| `SEMI`   | `;`    | Noktalı virgül (Go’da çoğu zaman otomatik eklenir) | `x := 1; y := 2`      |
| `COLON`  | `:`    | İki nokta (label, `:=`)                            | `label:`              |

/*
---

# 📌 Özet

* **Aritmetik**: `+ - * / %`
* **Karşılaştırma**: `== != < <= > >=`
* **Mantıksal**: `&& || !`
* **Bitwise**: `& | ^ << >> &^`
* **Atama**: `= :=` ve kısaltılmış halleri (`+=`, `&=`, vb.)
* **Artırma/Azaltma**: `++ --`
* **Delimiters**: `() [] {} , ; :`

---

👉 İstersen ben sana **her bir operatör için küçük kod örnekleri** yazabilirim. Mesela `&^` (bit temizleme) biraz karışık olur, onu örnekle açıklayabilirim. İstiyor musun?
EVET
*/
/*
Harika 😄 O zaman Go’daki bazı **karmaşık veya ilginç operatörler** için küçük örnekler gösterelim, böylece gerçek kullanımını daha iyi anlayabilirsin.

---

# 1️⃣ `&^` → Bit Temizleme (AND NOT)

`a &^ b` ifadesi, **`a` nın bitlerini `b` ile temizler**.

* Yani: `a &^ b = a & (^b)`
*/

``go
package main

import "fmt"

func main() {
	a := 14       // 1110
	b := 5        // 0101
	c := a &^ b   // 1110 &^ 0101 = 1010 (10)
	fmt.Println(c)
}
``

//📌 Çıktı:

``
10
``

/*
---

# 2️⃣ `<<` ve `>>` → Bit Kaydırma
*/

``go
package main

import "fmt"

func main() {
	x := 3          // 0011
	left := x << 2  // 1100 = 12
	right := x >> 1 // 0001 = 1
	fmt.Println(left, right)
}
``

//📌 Çıktı:

``
12 1
``
/*
---

# 3️⃣ `+=`, `-=`, `*=`, `/=` → Kısaltılmış Atama
*/

``go
package main

import "fmt"

func main() {
	x := 5
	x += 3   // x = x + 3 → 8
	x *= 2   // x = x * 2 → 16
	fmt.Println(x)
}
``

//📌 Çıktı:

``
16
``

/*
---

# 4️⃣ `++` ve `--` → Artırma / Azaltma
*/

``go
package main

import "fmt"

func main() {
	i := 1
	i++      // 2
	i--      // 1
	fmt.Println(i)
}
``

//📌 Çıktı:

``
1
``

/*
> ⚠️ Not: Go’da `x = i++` gibi ifadeler yoktur, sadece ayrı statement olarak kullanılır.

---

# 5️⃣ `&&`, `||`, `!` → Mantıksal Operatörler
*/

``go
package main

import "fmt"

func main() {
	a := true
	b := false
	fmt.Println(a && b) // false
	fmt.Println(a || b) // true
	fmt.Println(!a)     // false
}
``
/*
---

# 6️⃣ `, ; :` → Ayırıcılar (Delimiters)
*/

``go
package main

import "fmt"

func main() {
	x, y := 1, 2   // virgül ile birden fazla değişken tanımı
	fmt.Println(x, y)

label:
	for i := 0; i < 1; i++ { // ; ile for statement
		fmt.Println("Label örneği")
		break label // : ile label kullanımı
	}
}
``

📌 Çıktı:

``
1 2
Label örneği
``

---

# 🔹 Özet

* `&^` → bit temizleme
* `<<` / `>>` → bit kaydırma
* `+=`, `-=`, `*=`, `/=` → kısaltılmış atama
* `++` / `--` → artırma / azaltma
* `&&`, `||`, `!` → mantıksal operatörler
* `, ; :` → ayırıcılar (virgül, noktalı virgül, label)

/*
---

İstersen ben sana **tüm Operators & Delimiters için tek bir tabloda hem token hem örnek** olacak şekilde **tam bir referans tablosu** hazırlayabilirim, böylece Go operatörlerini bir bakışta görebilirsin. Bunu yapayım mı?
EVET
*/
/*
Harika 😄 O zaman sana Go’daki **tüm Operators & Delimiters** için hem **token adı**, hem **sembol**, hem **küçük örnek** içeren **tam bir referans tablosu** hazırladım.

---
*/
# 🟢 Go Operators & Delimiters – Tam Referans Tablosu

| Token            | Sembol | Açıklama                    | Küçük Örnek        | Çıktı          |        |       |         |        |
| ---------------- | ------ | --------------------------- | ------------------ | -------------- | ------ | ----- | ------- | ------ |
| `ADD`            | `+`    | Toplama, string birleştirme | `3 + 2`            | `5`            |        |       |         |        |
| `SUB`            | `-`    | Çıkarma, negatif işaret     | `10 - 3`           | `7`            |        |       |         |        |
| `MUL`            | `*`    | Çarpma, pointer dereference | `4 * 5`            | `20`           |        |       |         |        |
| `QUO`            | `/`    | Bölme                       | `10 / 2`           | `5`            |        |       |         |        |
| `REM`            | `%`    | Mod (kalan)                 | `10 % 3`           | `1`            |        |       |         |        |
| `INC`            | `++`   | 1 artırma                   | `i := 1; i++`      | `i=2`          |        |       |         |        |
| `DEC`            | `--`   | 1 azaltma                   | `i := 2; i--`      | `i=1`          |        |       |         |        |
| `EQL`            | `==`   | Eşit mi?                    | `3 == 3`           | `true`         |        |       |         |        |
| `NEQ`            | `!=`   | Eşit değil mi?              | `3 != 4`           | `true`         |        |       |         |        |
| `LSS`            | `<`    | Küçük mü?                   | `3 < 5`            | `true`         |        |       |         |        |
| `LEQ`            | `<=`   | Küçük veya eşit             | `3 <= 3`           | `true`         |        |       |         |        |
| `GTR`            | `>`    | Büyük mü?                   | `5 > 2`            | `true`         |        |       |         |        |
| `GEQ`            | `>=`   | Büyük veya eşit             | `5 >= 5`           | `true`         |        |       |         |        |
| `LAND`           | `&&`   | Mantıksal VE                | `true && false`    | `false`        |        |       |         |        |
| `LOR`            | \`     |                             | \`                 | Mantıksal VEYA | \`true |       | false\` | `true` |
| `NOT`            | `!`    | Mantıksal DEĞİL             | `!true`            | `false`        |        |       |         |        |
| `AND`            | `&`    | Bitwise AND                 | `6 & 3`            | `2`            |        |       |         |        |
| `OR`             | \`     | \`                          | Bitwise OR         | \`6            | 3\`    | `7`   |         |        |
| `XOR`            | `^`    | Bitwise XOR                 | `6 ^ 3`            | `5`            |        |       |         |        |
| `SHL`            | `<<`   | Bit kaydırma sola           | `1 << 3`           | `8`            |        |       |         |        |
| `SHR`            | `>>`   | Bit kaydırma sağa           | `8 >> 2`           | `2`            |        |       |         |        |
| `AND_NOT`        | `&^`   | Bit temizleme               | `14 &^ 5`          | `10`           |        |       |         |        |
| `ASSIGN`         | `=`    | Atama                       | `x = 5`            | `x=5`          |        |       |         |        |
| `DEFINE`         | `:=`   | Kısa değişken tanımlama     | `x := 10`          | `x=10`         |        |       |         |        |
| `ADD_ASSIGN`     | `+=`   | Toplama + atama             | `x := 5; x += 3`   | `x=8`          |        |       |         |        |
| `SUB_ASSIGN`     | `-=`   | Çıkarma + atama             | `x := 5; x -= 2`   | `x=3`          |        |       |         |        |
| `MUL_ASSIGN`     | `*=`   | Çarpma + atama              | `x := 4; x *= 2`   | `x=8`          |        |       |         |        |
| `QUO_ASSIGN`     | `/=`   | Bölme + atama               | `x := 8; x /= 2`   | `x=4`          |        |       |         |        |
| `REM_ASSIGN`     | `%=`   | Mod + atama                 | `x := 7; x %= 3`   | `x=1`          |        |       |         |        |
| `AND_ASSIGN`     | `&=`   | Bitwise AND + atama         | `x := 6; x &= 3`   | `x=2`          |        |       |         |        |
| `OR_ASSIGN`      | \`     | =\`                         | Bitwise OR + atama | \`x := 6; x    | = 3\`  | `x=7` |         |        |
| `XOR_ASSIGN`     | `^=`   | Bitwise XOR + atama         | `x := 6; x ^= 3`   | `x=5`          |        |       |         |        |
| `SHL_ASSIGN`     | `<<=`  | Bit shift left + atama      | `x := 1; x <<= 3`  | `x=8`          |        |       |         |        |
| `SHR_ASSIGN`     | `>>=`  | Bit shift right + atama     | `x := 8; x >>= 2`  | `x=2`          |        |       |         |        |
| `AND_NOT_ASSIGN` | `&^=`  | Bit temizleme + atama       | `x := 14; x &^= 5` | `x=10`         |        |       |         |        |

---

# 📦 Delimiters (Ayırıcılar)

| Token    | Sembol | Açıklama              | Örnek            | Çıktı |
| -------- | ------ | --------------------- | ---------------- | ----- |
| `LPAREN` | `(`    | Parantez aç           | `f(x)`           | —     |
| `RPAREN` | `)`    | Parantez kapat        | `f(x)`           | —     |
| `LBRACK` | `[`    | Köşeli parantez aç    | `arr[0]`         | —     |
| `RBRACK` | `]`    | Köşeli parantez kapat | `arr[0]`         | —     |
| `LBRACE` | `{`    | Süslü parantez aç     | `func() {}`      | —     |
| `RBRACE` | `}`    | Süslü parantez kapat  | `func() {}`      | —     |
| `COMMA`  | `,`    | Virgül                | `a, b := 1, 2`   | —     |
| `SEMI`   | `;`    | Noktalı virgül        | `x := 1; y := 2` | —     |
| `COLON`  | `:`    | İki nokta (label, :=) | `label: x := 5`  | —     |

/*
---

Bu tablo sayesinde artık **Go’daki tüm operatörler ve ayırıcıları** bir bakışta görebilir ve küçük örneklerle kullanımlarını anlayabilirsin ✅

---

Eğer istersen ben bunu bir adım ileri götürüp **her operatör için Go playground’da çalıştırılabilir minik örnek kodu otomatik oluşturacak bir Go programı** da yazabilirim.

Bunu yapayım mı?
EVET
*/

/*
Harika 😄 O zaman sana bir **Go programı** hazırlayacağım. Bu program:

* Tüm **operators & delimiters** listesini kullanacak
* Her biri için küçük bir **örnek kod snippet’i** oluşturacak
* Çalıştırılabilir olacak ve çıktı gösterecek

Yani bir nevi **otomatik Go operatör örnek üreteci** olacak 🚀

---

# 📂 Kod – Go Operator & Delimiter Örnek Üreteci
*/
``go
package main

import "fmt"

func main() {
	fmt.Println("📌 Aritmetik Operatörler")
	a, b := 10, 3
	fmt.Printf("ADD: %d + %d = %d\n", a, b, a+b)
	fmt.Printf("SUB: %d - %d = %d\n", a, b, a-b)
	fmt.Printf("MUL: %d * %d = %d\n", a, b, a*b)
	fmt.Printf("QUO: %d / %d = %d\n", a, b, a/b)
	fmt.Printf("REM: %d %% %d = %d\n", a, b, a%b)

	fmt.Println("\n📌 Artırma / Azaltma")
	i := 5
	fmt.Println("i =", i)
	i++
	fmt.Println("i++ =", i)
	i--
	fmt.Println("i-- =", i)

	fmt.Println("\n📌 Karşılaştırma Operatörleri")
	fmt.Printf("EQL: %d == %d -> %v\n", a, b, a==b)
	fmt.Printf("NEQ: %d != %d -> %v\n", a, b, a!=b)
	fmt.Printf("LSS: %d < %d -> %v\n", a, b, a<b)
	fmt.Printf("LEQ: %d <= %d -> %v\n", a, b, a<=b)
	fmt.Printf("GTR: %d > %d -> %v\n", a, b, a>b)
	fmt.Printf("GEQ: %d >= %d -> %v\n", a, b, a>=b)

	fmt.Println("\n📌 Mantıksal Operatörler")
	x, y := true, false
	fmt.Printf("LAND: %v && %v -> %v\n", x, y, x&&y)
	fmt.Printf("LOR: %v || %v -> %v\n", x, y, x||y)
	fmt.Printf("NOT: !%v -> %v\n", x, !x)

	fmt.Println("\n📌 Bit Operatörleri")
	c, d := 6, 3
	fmt.Printf("AND: %d & %d = %d\n", c, d, c&d)
	fmt.Printf("OR: %d | %d = %d\n", c, d, c|d)
	fmt.Printf("XOR: %d ^ %d = %d\n", c, d, c^d)
	fmt.Printf("SHL: %d << 1 = %d\n", c, c<<1)
	fmt.Printf("SHR: %d >> 1 = %d\n", c, c>>1)
	fmt.Printf("AND_NOT: %d &^ %d = %d\n", 14, 5, 14&^5)

	fmt.Println("\n📌 Atama Operatörleri")
	v := 5
	fmt.Println("v =", v)
	v += 3
	fmt.Println("v += 3 ->", v)
	v -= 2
	fmt.Println("v -= 2 ->", v)
	v *= 2
	fmt.Println("v *= 2 ->", v)
	v /= 2
	fmt.Println("v /= 2 ->", v)
	v %= 3
	fmt.Println("v %= 3 ->", v)
	v &= 3
	fmt.Println("v &= 3 ->", v)
	v |= 1
	fmt.Println("v |= 1 ->", v)
	v ^= 2
	fmt.Println("v ^= 2 ->", v)
	v <<= 1
	fmt.Println("v <<= 1 ->", v)
	v >>= 1
	fmt.Println("v >>= 1 ->", v)
	v &^= 1
	fmt.Println("v &^= 1 ->", v)

	fmt.Println("\n📌 Delimiters (Ayırıcılar)")
	arr := []int{1,2,3}
	fmt.Println("arr:", arr)
	fmt.Println("Parantez kullanımı:", (1 + 2) * 3)
	fmt.Println("Süslü parantez: if true { fmt.Println(\"Hello\") }")
	fmt.Println("Virgül kullanımı: a,b := 1,2 ->", 1, 2)
	fmt.Println("Noktalı virgül ve label örneği:")
	label:
	for i := 0; i < 1; i++ {
		fmt.Println("Label çalıştı")
		break label
	}
}
``
/*
---

# 📌 Açıklama

1. **Aritmetik operatörler**: `+ - * / %`
2. **Artırma / Azaltma**: `++ --`
3. **Karşılaştırma**: `== != < <= > >=`
4. **Mantıksal**: `&& || !`
5. **Bitwise**: `& | ^ << >> &^`
6. **Atama**: `= += -= *= /= %= &= |= ^= <<= >>= &^=`
7. **Delimiters**: `() [] {}` , `, ; :`

Bu programı çalıştırırsan, **her operatör ve delimiter için örnek çıktıyı** görebilirsin ✅

---

İstersen ben bunu bir adım daha ileri götürüp **her operatörü tek bir fonksiyon içinde otomatik test edip JSON çıktısı üreten bir araç** da yapabilirim. Böylece programatik olarak referans tablosu ve test seti de olur.

Bunu yapayım mı?
*/