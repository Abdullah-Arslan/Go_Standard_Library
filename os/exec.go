/*
`exec` paketi Go’nun **`os/exec`** paketidir ve işletim sistemi üzerinde **dış komutları / programları çalıştırmak** için kullanılır. Bu paket ile terminalde çalıştırdığınız komutları Go içinden de çalıştırabilir, çıktısını alabilir, hata yönetimi yapabilir ve stdin/stdout/stderr akışlarını kontrol edebilirsiniz.

Aşağıda paketi detaylıca tüm yönleriyle örneklerle açıklıyorum:

---

## 1. Temel Kullanım: `exec.Command`

`exec.Command` fonksiyonu çalıştırmak istediğiniz komutu ve argümanlarını alır, size bir `*Cmd` nesnesi döner.
*/
``go
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("echo", "Merhaba Dünya")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Println(string(output))
}
``

//📌 Çıktı:

``
Merhaba Dünya
``
/*
---

## 2. Komutun Hata Çıkışını Yakalama (`CombinedOutput`)

`CombinedOutput()` ile hem standart çıktıyı hem de hata çıktısını birleştirebilirsiniz.
*/

``go
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "/olmayan-klasor")
	output, err := cmd.CombinedOutput()
	fmt.Println("Çıktı:", string(output))
	fmt.Println("Hata:", err)
}
``
/*
---

## 3. Çalışan Komutun Doğrudan Terminale Yazması

`Run()` metodu komutu çalıştırır ama çıktıyı döndürmez. Eğer `cmd.Stdout` ve `cmd.Stderr` alanlarını bağlarsanız çıktıyı direkt terminalde görebilirsiniz.
*/

``go
package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-l")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
``
/*
---

## 4. Komutun Girdisine Veri Yazmak (`StdinPipe`)

Bir programa veri göndermek için `StdinPipe()` kullanılabilir.
*/

``go
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("grep", "go")

	stdin, _ := cmd.StdinPipe()
	go func() {
		defer stdin.Close()
		stdin.Write([]byte("golang\njava\npython\ngo\n"))
	}()

	output, _ := cmd.CombinedOutput()
	fmt.Println(string(output))
}
``

//📌 Çıktı:

``
golang
go
``
/*
---

## 5. Çalışma Dizini Belirlemek (`cmd.Dir`)
*/
``go
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls")
	cmd.Dir = "/tmp" // /tmp dizininde çalıştır
	out, _ := cmd.Output()
	fmt.Println(string(out))
}
``
/*
---

## 6. Ortam Değişkenleri Belirlemek (`cmd.Env`)
?7
``go
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("printenv", "MY_VAR")
	cmd.Env = append(cmd.Env, "MY_VAR=Selamlar")
	out, _ := cmd.Output()
	fmt.Println(string(out))
}
``
/*
---

## 7. Asenkron Çalıştırma (`Start` ve `Wait`)

`Start()` komutu başlatır ama beklemez, `Wait()` ile sonlanmasını bekleyebilirsiniz.
*/
``go
package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	cmd := exec.Command("sleep", "2")

	err := cmd.Start()
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	fmt.Println("Komut başlatıldı, bekleniyor...")
	cmd.Wait()
	fmt.Println("Komut bitti!")
}
``
/*
---

## 8. Komutu Zaman Aşımı ile Çalıştırmak (`context` ile)

Go’nun `context` paketi ile birlikte `exec.CommandContext` kullanarak zaman aşımı koyabilirsiniz.
*/

``go
package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "sleep", "5")

	err := cmd.Run()
	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Komut zaman aşımına uğradı!")
	} else {
		fmt.Println("Komut tamamlandı:", err)
	}
}
``
/*
---

## 9. Birden Fazla Komutu Pipe ile Bağlamak

Go’da `cmd.StdoutPipe()` ve `cmd.Stdin` ile komutları birbirine bağlayabilirsiniz. Örneğin `echo "golang" | grep go`:
*/

``go
package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	echoCmd := exec.Command("echo", "golang")
	grepCmd := exec.Command("grep", "go")

	pipe, _ := echoCmd.StdoutPipe()
	grepCmd.Stdin = pipe

	out, _ := grepCmd.StdoutPipe()

	_ = echoCmd.Start()
	_ = grepCmd.Start()

	bytes, _ := io.ReadAll(out)
	_ = echoCmd.Wait()
	_ = grepCmd.Wait()

	fmt.Println(string(bytes))
}
``

//📌 Çıktı:

``
golang
``
/*
---

## 10. Özet – `exec` Paketinin Sağladıkları

* **`Command`** → Yeni bir komut oluşturur
* **`Run`** → Çalıştırır, bekler
* **`Output`** → Çıktıyı döndürür
* **`CombinedOutput`** → Çıktı + hata birleşik döner
* **`Start` & `Wait`** → Asenkron çalıştırma
* **`StdinPipe`, `StdoutPipe`, `StderrPipe`** → Veri alışverişi
* **`Dir`** → Çalışma dizini
* **`Env`** → Ortam değişkenleri
* **`CommandContext`** → Zaman aşımı ve iptal desteği

---

Tamamdır 👍 Sadece burada detaylı açıklamalar ve örnek kodlarla anlattığım haliyle bırakıyorum.

Yani `exec` paketiyle artık:

* Dış programları nasıl çalıştıracağını,
* Çıktılarını nasıl alacağını,
* Ortam değişkenleri / çalışma dizinini nasıl değiştireceğini,
* Komutları asenkron çalıştırmayı,
* Komutları birbirine pipe ile bağlamayı,
* Zaman aşımı ile komut çalıştırmayı

tam kapsamlı görmüş oldun ✅

İstersen bir sonraki adımda sana gerçek **mini proje örneği** hazırlayabilirim:
Mesela küçük bir “komut çalıştırıcı terminal arayüzü” (Go ile yazılmış basit bir shell).

👉 Böyle bir proje ister misin, yoksa paket açıklamaları bu şekilde kalsın mı?
EVET HAZIRLA
*/
/*
Süper 👍 O zaman sana **`exec` paketiyle küçük bir mini proje** hazırlayayım:
Bir **Go Shell** (küçük komut çalıştırıcı terminal arayüzü).

---

# 🔹 Mini Proje: Go Shell (exec paketi ile)

Bu program:

* Kullanıcıdan komut alır (`ls`, `echo`, `pwd` gibi).
* `exit` yazılana kadar çalışmaya devam eder.
* `exec.Command` ile komutu işletim sisteminde çalıştırır.
* Çıktıyı ekrana basar.

---

## 📄 Kod
*/
``go
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("🔹 Go Shell'e hoş geldiniz!")
	fmt.Println("Komut girin (çıkmak için 'exit' yazın)")

	for {
		fmt.Print("go-shell> ")

		// Kullanıcıdan girdi al
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Çıkış kontrolü
		if input == "exit" {
			fmt.Println("Go Shell kapatılıyor...")
			break
		}

		// Boş komut girildiyse devam et
		if input == "" {
			continue
		}

		// Komutu parçala
		args := strings.Split(input, " ")
		cmd := exec.Command(args[0], args[1:]...)

		// Çıktıları terminale bağla
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Komutu çalıştır
		err := cmd.Run()
		if err != nil {
			fmt.Println("Hata:", err)
		}
	}
}
``
/*
---

## 🖥 Kullanım Örneği

Programı çalıştırdığında:
*/
``
🔹 Go Shell e hoş geldiniz!
Komut girin (çıkmak için 'exit' yazın)
go-shell> pwd
/home/kullanici

go-shell> ls
main.go
notlar.txt

go-shell> echo Merhaba Golang
Merhaba Golang

go-shell> exit
Go Shell kapatılıyor...
``
/*
---

Bu şekilde kendi **küçük terminalini** Go ile yazmış oldun 🚀
Hatta bu yapıyı geliştirip:

* **zaman aşımı desteği** (uzun süren komutları iptal etme),
* **pipe desteği** (`ls | grep go`),
* **ortam değişkenleri** ayarlama

gibi özellikler ekleyebilirsin.

👉 İstiyor musun, sana bu Go Shell’i **pipe destekli (ör: `ls | grep go`) gelişmiş sürüm** olarak da yazayım mı?
EVET
*/
/*
Harika 🚀 O zaman şimdi sana **`exec` paketiyle pipe destekli Go Shell** yazıyorum.
Bu sürümde:

* `ls | grep go` gibi **pipe** (|) operatörünü destekler.
* Birden fazla komutu pipe ile bağlayıp arka arkaya çalıştırabilirsin.
* Çıktılar sırayla aktarılır.

---

# 🔹 Gelişmiş Go Shell (Pipe Destekli)
*/

``go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func runPipeline(commands []string) error {
	var prevStdout io.ReadCloser

	for i, cmdStr := range commands {
		// Komutu parçala
		args := strings.Fields(strings.TrimSpace(cmdStr))
		if len(args) == 0 {
			continue
		}

		cmd := exec.Command(args[0], args[1:]...)

		// Eğer önceki komuttan pipe varsa stdin olarak bağla
		if prevStdout != nil {
			cmd.Stdin = prevStdout
		}

		// Son komut mu?
		if i == len(commands)-1 {
			// Son komutun çıktısı ekrana yazılsın
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
		} else {
			// Sonraki komuta aktarılacak pipe oluştur
			stdout, _ := cmd.StdoutPipe()
			prevStdout = stdout
		}

		// Çalıştır
		if err := cmd.Start(); err != nil {
			return err
		}

		// Son komutsa bekle
		if i == len(commands)-1 {
			if err := cmd.Wait(); err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("🔹 Gelişmiş Go Shell (pipe destekli)")
	fmt.Println("Komut girin (çıkmak için 'exit' yazın)")

	for {
		fmt.Print("go-shell> ")

		// Kullanıcıdan komut al
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Çıkış
		if input == "exit" {
			fmt.Println("Go Shell kapatılıyor...")
			break
		}

		if input == "" {
			continue
		}

		// Pipe'lara göre parçala
		commands := strings.Split(input, "|")

		// Çalıştır
		if err := runPipeline(commands); err != nil {
			fmt.Println("Hata:", err)
		}
	}
}
``
/*
---

## 🖥 Kullanım Örneği
*/

``
go-shell> ls | grep go
main.go
example.go

go-shell> echo "golang\njava\npython" | grep go
golang

go-shell> ps aux | grep go
kullanici   1234   0.0  ...  ./goshell
``
/*
---

👉 Bu haliyle artık senin shell’in **Unix tarzı pipe operatörünü** destekliyor.
Bunu biraz daha geliştirip:

* **zaman aşımı** (`command & timeout` gibi),
* **arka planda çalıştırma (& operatörü)**,
* **komut geçmişi (history)**

gibi özellikler de ekleyebilirim.

İster misin, sana bu shell’i **arka plan işlemleri (&) destekli** hale de getireyim mi?
EVET
*/
/*
Süper 👍 Şimdi mevcut **pipe destekli Go Shell**’i biraz daha ileri taşıyalım:
Bu sürümde **arka plan işlemleri (&)** desteklenecek.

Yani:

* `sleep 5 &` → komut **arka planda** çalışacak, shell kilitlenmeyecek.
* Normal komutlar (`ls`, `echo`, `pwd`) eskisi gibi hemen çalışacak.
* Pipe (`|`) ve arka plan (`&`) beraber de kullanılabilecek.

---

# 🔹 Gelişmiş Go Shell (Pipe + Arka Plan İşlemleri)
*/
``go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
)

// Pipe destekli çalıştırma
func runPipeline(commands []string, background bool, wg *sync.WaitGroup) {
	defer func() {
		if background && wg != nil {
			wg.Done()
		}
	}()

	var prevStdout io.ReadCloser
	var processes []*exec.Cmd

	for i, cmdStr := range commands {
		// Komutu parçala
		args := strings.Fields(strings.TrimSpace(cmdStr))
		if len(args) == 0 {
			continue
		}

		cmd := exec.Command(args[0], args[1:]...)

		// Önceki komuttan pipe varsa stdin'e bağla
		if prevStdout != nil {
			cmd.Stdin = prevStdout
		}

		// Son komut mu?
		if i == len(commands)-1 {
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
		} else {
			stdout, _ := cmd.StdoutPipe()
			prevStdout = stdout
		}

		// Komutu başlat
		if err := cmd.Start(); err != nil {
			fmt.Println("Hata:", err)
			return
		}

		processes = append(processes, cmd)
	}

	// Eğer arka planda değilse hepsini bekle
	if !background {
		for _, cmd := range processes {
			cmd.Wait()
		}
	} else {
		fmt.Println("🔹 Arka planda çalışıyor:", strings.Join(commands, " | "))
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var wg sync.WaitGroup

	fmt.Println("🔹 Gelişmiş Go Shell (pipe + background &)")
	fmt.Println("Komut girin (çıkmak için 'exit' yazın)")

	for {
		fmt.Print("go-shell> ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Çıkış
		if input == "exit" {
			fmt.Println("Go Shell kapatılıyor...")
			break
		}

		if input == "" {
			continue
		}

		// Arka plan kontrolü
		background := false
		if strings.HasSuffix(input, "&") {
			background = true
			input = strings.TrimSpace(strings.TrimSuffix(input, "&"))
		}

		// Pipe'lara göre parçala
		commands := strings.Split(input, "|")

		if background {
			wg.Add(1)
			go runPipeline(commands, true, &wg)
		} else {
			runPipeline(commands, false, nil)
		}
	}

	// Tüm arka plan işler bitsin
	wg.Wait()
}
``
/*
---

## 🖥 Kullanım Örneği
*/
``
go-shell> sleep 5 &
🔹 Arka planda çalışıyor: sleep 5

go-shell> echo "golang\njava\npython" | grep go
golang

go-shell> ls | grep go &
🔹 Arka planda çalışıyor: ls | grep go
``
/*
---

Artık shell’in:
✅ Pipe (`|`)
✅ Arka plan işlemleri (`&`)
destekliyor 🎉

---

👉 İstersen ben buna bir sonraki adımda **zaman aşımı desteği (timeout)** ekleyebilirim.
Örn: `timeout 3 sleep 10` → 3 saniye sonra işlemi öldürür.

Bunu da ekleyeyim mi?
*/