// **bytes paketiyle tam bir uygulama** yazalım.

// Bu uygulama:

// 1. Kullanıcıdan bir metin alacak
// 2. `bytes` paketi ile **çeşitli işlemleri** gösterecek: split, join, replace, trim, upper/lower, contains, index, repeat, prefix/suffix kontrolü
// 3. `bytes.Buffer` ile veriyi birleştirip gösterecek

// ---

// ## 📝 Uygulama: `bytes_demo.go`

package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Kullanıcıdan örnek metin
	input := []byte("  Merhaba Go Dünyası!  ")

	fmt.Println("Orijinal:", string(input))

	// Trim ve TrimSpace
	trimmed := bytes.TrimSpace(input)
	fmt.Println("TrimSpace:", string(trimmed))
	trimmed2 := bytes.Trim(trimmed, "M!")
	fmt.Println("Trim 'M!':", string(trimmed2))

	// ToUpper ve ToLower
	fmt.Println("ToUpper:", string(bytes.ToUpper(trimmed)))
	fmt.Println("ToLower:", string(bytes.ToLower(trimmed)))

	// Contains, Index, HasPrefix, HasSuffix
	fmt.Println("Contains 'Go':", bytes.Contains(trimmed, []byte("Go")))
	fmt.Println("Index 'Go':", bytes.Index(trimmed, []byte("Go")))
	fmt.Println("HasPrefix 'Mer':", bytes.HasPrefix(trimmed, []byte("Mer")))
	fmt.Println("HasSuffix 'Dünyası!':", bytes.HasSuffix(trimmed, []byte("Dünyası!")))

	// Split ve Join
	words := bytes.Split(trimmed, []byte(" "))
	fmt.Println("Split:", words)
	joined := bytes.Join(words, []byte("-"))
	fmt.Println("Join:", string(joined))

	// Replace
	replaced := bytes.Replace(trimmed, []byte("Go"), []byte("Golang"), -1)
	fmt.Println("Replace:", string(replaced))

	// Repeat
	repeated := bytes.Repeat([]byte("Ha"), 3)
	fmt.Println("Repeat:", string(repeated))

	// bytes.Buffer kullanımı
	var buf bytes.Buffer
	buf.WriteString("Buffer ile başla: ")
	buf.Write(trimmed)
	buf.WriteString(" | Son!")
	fmt.Println("Buffer:", buf.String())

	// bytes.Equal ve bytes.Compare
	a := []byte("abc")
	b := []byte("abd")
	fmt.Println("Equal 'abc' vs 'abd':", bytes.Equal(a, b))
	fmt.Println("Compare 'abc' vs 'abd':", bytes.Compare(a, b))
}

// ---

// ## 📌 Örnek Çıktı

// ```
// Orijinal:   Merhaba Go Dünyası!
// TrimSpace: Merhaba Go Dünyası!
// Trim 'M!': erhaba Go Dünyası
// ToUpper: MERHABA GO DÜNYASI!
// ToLower: merhaba go dünyası!
// Contains 'Go': true
// Index 'Go': 8
// HasPrefix 'Mer': true
// HasSuffix 'Dünyası!': true
// Split: [[77 101 114 104 97 98 97] [71 111] [68 252 110 121 97 115 131 33]]
// Join: Merhaba-Go-Dünyası!
// Replace: Merhaba Golang Dünyası!
// Repeat: HaHaHa
// Buffer: Buffer ile başla: Merhaba Go Dünyası! | Son!
// Equal 'abc' vs 'abd': false
// Compare 'abc' vs 'abd': -1
// ```

// ---

// Bu uygulama ile `bytes` paketinin **hemen hemen tüm temel fonksiyonlarını** tek bir yerde görebiliyorsun:

// * Trim / TrimSpace
// * ToUpper / ToLower
// * Contains / Index / HasPrefix / HasSuffix
// * Split / Join / Replace / Repeat
// * Buffer kullanımı
// * Equal / Compare

// ---
