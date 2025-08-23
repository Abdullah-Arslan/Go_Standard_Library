// GoLand'da builtin terimi, Go dilinin standart kütüphanesinde yer alan ve dil tarafından doğrudan desteklenen yerleşik fonksiyonları ifade eder.
// Bu fonksiyonlar özel olarak Go derleyicisi tarafından tanınır ve çoğunlukla temel işlemler için kullanılır.
// Go dilinde builtin paketinden doğrudan import yapılmaz çünkü bunlar dilin içinde gömülü olarak gelir.

// GoLand’da builtin (built-in) dediğimiz şey, Go dilinde doğrudan kullanılabilen ve ekstra import gerektirmeyen fonksiyonlar, tipler ve sabitlerdir.
// Yani Go diliyle beraber gelen temel işlevlerdir.

// Go dilinde builtin paketi, bu yerleşik fonksiyonları ve tipleri içerir.
// GoLand IDE'sinde de bu built-in fonksiyonları kolayca kullanabilir ve hata almadan kod yazabilirsiniz.

// Builtin Fonksiyonlar Nedir?

// Go dilinde aşağıdaki gibi temel fonksiyonlar builtin fonksiyonlardır:

// len()

// cap()

// append()

// copy()

// close()

// delete()

// panic()

// recover()

// new()

// make()

// print()

// println()

// Bu fonksiyonlar dilin temel yapısını ve çalışma mantığını destekler ve çoğu zaman programcının ekstra bir kütüphane eklemesine gerek kalmaz.


// Örneklerle Açıklama
// 1. len()

// Bir dizinin, string'in veya slice'ın uzunluğunu döner.
package main

import "fmt"

func main() {
    s := "Merhaba Go"
    fmt.Println(len(s)) // Çıktı: 10 (karakter sayısı)
    
    arr := []int{1, 2, 3, 4}
    fmt.Println(len(arr)) // Çıktı: 4
}


//2. cap()

//Bir slice veya kanalın kapasitesini döner.
package main

import "fmt"

func main() {
    s := make([]int, 3, 5)
    fmt.Println(len(s)) // 3 -> uzunluk
    fmt.Println(cap(s)) // 5 -> kapasite
}


//3. append()

//Bir slice’a eleman eklemek için kullanılır.
package main

import "fmt"

func main() {
    s := []int{1, 2, 3}
    s = append(s, 4, 5)
    fmt.Println(s) // Çıktı: [1 2 3 4 5]
}


//4. copy()

//Bir slice’ı başka bir slice’a kopyalamak için kullanılır.
package main

import "fmt"

func main() {
    src := []int{1, 2, 3}
    dst := make([]int, 3)
    copy(dst, src)
    fmt.Println(dst) // Çıktı: [1 2 3]
}

//5. delete()

//Map’ten anahtar (bilgi) silmek için kullanılır.
package main

import "fmt"

func main() {
    m := map[string]int{"a": 1, "b": 2}
    delete(m, "a")
    fmt.Println(m) // Çıktı: map[b:2]
}


//6. panic() ve recover()

//Hata fırlatmak ve yakalamak için kullanılır.

package main

import "fmt"

func mayPanic() {
    panic("bir hata oluştu")
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    
    mayPanic()
    fmt.Println("Bu satır çalışmaz.")
}


// 7. make() ve new()

// make() slice, map ve kanal oluşturmak için kullanılır.

// new() ise tipin sıfır değerine sahip pointer’ını döner.
package main

import "fmt"

func main() {
    s := make([]int, 3) // slice oluşturur
    fmt.Println(s)      // [0 0 0]

    p := new(int)       // int tipinde pointer oluşturur, değeri 0
    fmt.Println(*p)     // 0
}


// GoLand’da Builtin Fonksiyonları Görmek

// GoLand IDE'sinde kod yazarken, builtin fonksiyonlar otomatik tamamlamaya (autocomplete) dahil olur ve dokümantasyonlarını görmek için:

// Fonksiyonun üzerine gelip Ctrl+Q (Windows/Linux) ya da F1 (Mac) ile detaylı açıklamayı görebilirsiniz.

// Ayrıca, builtin fonksiyonlar için doğrudan Go Documentation penceresinde bilgi alabilirsiniz.

// Özet

// Builtin fonksiyonlar, Go diline gömülü temel fonksiyonlardır.

// Kodda herhangi bir import yapmadan doğrudan kullanılabilirler.

// Diziler, slice’lar, map’ler, pointer’lar, kanallar ve hata yönetimi için önemli fonksiyonlar içerir.

// GoLand IDE, bu fonksiyonları otomatik tamamlar ve belgelerini kolayca gösterir.
