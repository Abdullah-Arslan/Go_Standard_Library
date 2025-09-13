Go dilinde **operatörler (operators)** birkaç gruba ayrılır. İşte tüm listeyi gruplara ayrılmış şekilde veriyorum:

---

## 1. Aritmetik Operatörler

Matematiksel işlemler için kullanılır:

| Operatör | Açıklama               | Örnek (x=10, y=3) | Sonuç |
| -------- | ---------------------- | ----------------- | ----- |
| `+`      | Toplama                | `x + y`           | 13    |
| `-`      | Çıkarma                | `x - y`           | 7     |
| `*`      | Çarpma                 | `x * y`           | 30    |
| `/`      | Bölme (tam sayı bölme) | `x / y`           | 3     |
| `%`      | Mod (kalan)            | `x % y`           | 1     |

---

## 2. Atama Operatörleri

Değişkenlere değer atamak için:

| Operatör | Açıklama          | Örnek (`x=5`)    | Sonuç |       |     |
| -------- | ----------------- | ---------------- | ----- | ----- | --- |
| `=`      | Basit atama       | `x = 10`         | x=10  |       |     |
| `+=`     | Toplayarak atama  | `x += 3`         | x=8   |       |     |
| `-=`     | Çıkararak atama   | `x -= 2`         | x=3   |       |     |
| `*=`     | Çarparak atama    | `x *= 4`         | x=20  |       |     |
| `/=`     | Bölerek atama     | `x /= 5`         | x=1   |       |     |
| `%=`     | Mod alarak atama  | `x %= 3`         | x=2   |       |     |
| `&=`     | Bitwise AND atama | `x &= 2`         | ...   |       |     |
| \`       | =\`               | Bitwise OR atama | \`x   | = 2\` | ... |
| `^=`     | Bitwise XOR atama | `x ^= 2`         | ...   |       |     |
| `<<=`    | Bit kaydırma sola | `x <<= 1`        | ...   |       |     |
| `>>=`    | Bit kaydırma sağa | `x >>= 1`        | ...   |       |     |

---

## 3. Karşılaştırma (İlişkisel) Operatörler

Sonuç **true** veya **false** döner:

| Operatör | Açıklama      | Örnek (`x=10, y=20`) | Sonuç |
| -------- | ------------- | -------------------- | ----- |
| `==`     | Eşit mi       | `x == y`             | false |
| `!=`     | Eşit değil mi | `x != y`             | true  |
| `>`      | Büyüktür      | `x > y`              | false |
| `<`      | Küçüktür      | `x < y`              | true  |
| `>=`     | Büyük eşit    | `x >= y`             | false |
| `<=`     | Küçük eşit    | `x <= y`             | true  |

---

## 4. Mantıksal Operatörler

Boolean (true/false) üzerinde işlem yapar:

| Operatör | Açıklama      | Örnek (`x=true, y=false`) | Sonuç        |     |   |     |      |
| -------- | ------------- | ------------------------- | ------------ | --- | - | --- | ---- |
| `&&`     | Mantıksal AND | `x && y`                  | false        |     |   |     |      |
| \`       |               | \`                        | Mantıksal OR | \`x |   | y\` | true |
| `!`      | Mantıksal NOT | `!x`                      | false        |     |   |     |      |

---

## 5. Bit Düzeyinde (Bitwise) Operatörler

Sayıların ikilik (binary) temsilleri üzerinde işlem yapar:

| Operatör | Açıklama            | Örnek (`x=6 (110), y=3 (011)`) | Sonuç     |     |         |
| -------- | ------------------- | ------------------------------ | --------- | --- | ------- |
| `&`      | AND                 | `x & y`                        | 2 (010)   |     |         |
| \`       | \`                  | OR                             | \`x       | y\` | 7 (111) |
| `^`      | XOR                 | `x ^ y`                        | 5 (101)   |     |         |
| `&^`     | AND NOT (bit clear) | `x &^ y`                       | 4 (100)   |     |         |
| `<<`     | Sola kaydırma       | `x << 1`                       | 12 (1100) |     |         |
| `>>`     | Sağa kaydırma       | `x >> 1`                       | 3 (011)   |     |         |

---

## 6. Adres Operatörleri

Bellek adresleri üzerinde işlem yapar:

| Operatör | Açıklama                        |
| -------- | ------------------------------- |
| `&`      | Bir değişkenin adresini alır    |
| `*`      | Pointer üzerinden değere erişir |

---

## 7. Özel Operatörler

Go diline özgü bazı operatörler:

| Operatör | Açıklama                                |
| -------- | --------------------------------------- |
| `<-`     | Kanal (channel) gönderme/alma operatörü |
| `...`    | Slice genişletme (variadic arg.)        |

---

✅ Go’da **++ ve --** operatörleri **postfix** (sonda) olarak vardır (`i++`, `i--`), **prefix** (örn. `++i`) yoktur.

---

İstersen sana her operatör için **küçük çalıştırılabilir Go örnekleri** de yazabilirim. Onları da hazırlamamı ister misin?
