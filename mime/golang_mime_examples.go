/*
İstersen şimdi ben bu **tüm örnekleri tek bir bütün halinde çalıştırılabilir Go dosyası** hâline getirip, doğrudan kopyalayıp çalıştırabileceğin bir set hazırlayabilirim.

Bunu yapmamı ister misin?
EVET
*/
//Golang_Mime_Examples Go Örneği

package main

import (
	"fmt"
	"mime"
)

func main() {
	// 1. TypeByExtension
	fmt.Println("--- TypeByExtension ---")
	fmt.Println(mime.TypeByExtension(".txt"))  // text/plain; charset=utf-8
	fmt.Println(mime.TypeByExtension(".html")) // text/html; charset=utf-8
	fmt.Println(mime.TypeByExtension(".png"))  // image/png

	// 2. ExtensionsByType
	fmt.Println("--- ExtensionsByType ---")
	exts, err := mime.ExtensionsByType("text/html")
	if err != nil {
		fmt.Println("Hata:", err)
	} else {
		fmt.Println(exts) // [.html .htm]
	}

	// 3. AddExtensionType
	fmt.Println("--- AddExtensionType ---")
	mime.AddExtensionType(".md", "text/markdown")
	fmt.Println(mime.TypeByExtension(".md")) // text/markdown

	// 4. FormatMediaType
	fmt.Println("--- FormatMediaType ---")
	params := map[string]string{"charset": "utf-8"}
	fmt.Println(mime.FormatMediaType("text/plain", params)) // text/plain; charset=utf-8

	// 5. ParseMediaType
	fmt.Println("--- ParseMediaType ---")
	typ, params, err := mime.ParseMediaType("text/plain; charset=utf-8")
	if err != nil {
		fmt.Println("Hata:", err)
	} else {
		fmt.Println("Tip:", typ)
		fmt.Println("Parametreler:", params)
	}

	// 6. Custom MIME type .log
	fmt.Println("--- Custom MIME Type ---")
	mime.AddExtensionType(".log", "text/log")
	fmt.Println(mime.TypeByExtension(".log"))
	exts, _ = mime.ExtensionsByType("text/log")
	fmt.Println(exts)

	// 7. FormatMediaType with multiple params
	fmt.Println("--- FormatMediaType with multiple params ---")
	params2 := map[string]string{"charset": "utf-8", "version": "1.0"}
	fmt.Println(mime.FormatMediaType("application/json", params2)) // application/json; charset=utf-8; version=1.0

	// 8. ParseMediaType with multiple params
	fmt.Println("--- ParseMediaType with multiple params ---")
	typ2, params2Parsed, err := mime.ParseMediaType("application/json; charset=utf-8; version=1.0")
	if err != nil {
		fmt.Println("Hata:", err)
	} else {
		fmt.Println("Tip:", typ2)
		fmt.Println("Parametreler:", params2Parsed)
	}
}

/*

Tamamdır! Tüm `mime` paketi örneklerini tek bir **çalıştırılabilir Go dosyası** hâline getirdim. Artık bu dosyayı doğrudan çalıştırabilir ve tüm örnekleri tek seferde görebilirsin.

Dosya adı: **Golang Mime Examples**.
*/
