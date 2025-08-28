//ecdsa.go daki örnek uygulamanın tam hali.

// ed25519_cli.go
// Ed25519 JSON key export/import + signing & verifying (single-file CLI)
//
// Commands:
//   genkey  -out priv.json -pub pub.json [-comment "my key"]
//   sign    -key priv.json -in message.txt -out sig.json [-comment "v1"]
//   verify  -pub pub.json -in message.txt -sig sig.json
//   pub     -key priv.json -out pub.json   # derive public from private JSON
//
// JSON formats (all base64-encoded binary fields):
// Private key JSON:
//   {"type":"ed25519","seed":"...","public":"...","comment":"..."}
// Public key JSON:
//   {"type":"ed25519","public":"...","comment":"..."}
// Signature JSON:
//   {"type":"ed25519","signature":"...","public":"...","comment":"..."}
//
// Usage examples:
//   go run ed25519_cli.go genkey -out priv.json -pub pub.json -comment "test key"
//   go run ed25519_cli.go sign   -key priv.json -in README.md -out sig.json -comment "release v1"
//   go run ed25519_cli.go verify -pub pub.json  -in README.md -sig sig.json
//   go run ed25519_cli.go pub    -key priv.json -out pub.json

package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type PrivJSON struct {
	Type    string `json:"type"`
	Seed    string `json:"seed"`   // base64 (32 bytes)
	Public  string `json:"public"` // base64 (32 bytes)
	Comment string `json:"comment,omitempty"`
}

type PubJSON struct {
	Type    string `json:"type"`
	Public  string `json:"public"` // base64 (32 bytes)
	Comment string `json:"comment,omitempty"`
}

type SigJSON struct {
	Type      string `json:"type"`
	Signature string `json:"signature"` // base64 (64 bytes)
	Public    string `json:"public"`    // base64 (32 bytes) – signer pubkey (optional but handy)
	Comment   string `json:"comment,omitempty"`
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	switch os.Args[1] {
	case "genkey":
		cmdGenKey(os.Args[2:])
	case "sign":
		cmdSign(os.Args[2:])
	case "verify":
		cmdVerify(os.Args[2:])
	case "pub":
		cmdPub(os.Args[2:])
	default:
		usage()
		os.Exit(2)
	}
}

func usage() {
	fmt.Println("Ed25519 CLI")
	fmt.Println("  genkey  -out priv.json -pub pub.json [-comment '...']")
	fmt.Println("  sign    -key priv.json -in file -out sig.json [-comment '...']")
	fmt.Println("  verify  -pub pub.json  -in file -sig sig.json")
	fmt.Println("  pub     -key priv.json -out pub.json")
}

// ---------------- genkey ----------------
func cmdGenKey(args []string) {
	fs := flag.NewFlagSet("genkey", flag.ExitOnError)
	privPath := fs.String("out", "priv.json", "private key JSON output path")
	pubPath := fs.String("pub", "pub.json", "public key JSON output path")
	comment := fs.String("comment", "", "optional comment")
	_ = fs.Parse(args)

	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	check(err)
	seed := priv.Seed() // 32-byte seed

	privJSON := PrivJSON{
		Type:    "ed25519",
		Seed:    b64(seed),
		Public:  b64(pub),
		Comment: *comment,
	}
	pubJSON := PubJSON{Type: "ed25519", Public: b64(pub), Comment: *comment}

	writeJSON(*privPath, privJSON, 0600)
	writeJSON(*pubPath, pubJSON, 0644)
	fmt.Printf("Generated key pair.\nPrivate: %s\nPublic : %s\n", *privPath, *pubPath)
}

// ---------------- sign ----------------
func cmdSign(args []string) {
	fs := flag.NewFlagSet("sign", flag.ExitOnError)
	keyPath := fs.String("key", "priv.json", "private key JSON path")
	inPath := fs.String("in", "", "input file")
	outPath := fs.String("out", "sig.json", "signature JSON output path")
	comment := fs.String("comment", "", "optional comment")
	_ = fs.Parse(args)
	if *inPath == "" {
		fatal(errors.New("-in is required"))
	}

	priv, pub := readPriv(*keyPath)
	data := readFile(*inPath)
	sig := ed25519.Sign(priv, data)

	sigJSON := SigJSON{Type: "ed25519", Signature: b64(sig), Public: b64(pub), Comment: *comment}
	writeJSON(*outPath, sigJSON, 0644)
	fmt.Printf("Signed %s -> %s\n", *inPath, *outPath)
}

// ---------------- verify ----------------
func cmdVerify(args []string) {
	fs := flag.NewFlagSet("verify", flag.ExitOnError)
	pubPath := fs.String("pub", "pub.json", "public key JSON path")
	inPath := fs.String("in", "", "input file")
	sigPath := fs.String("sig", "sig.json", "signature JSON path")
	_ = fs.Parse(args)
	if *inPath == "" {
		fatal(errors.New("-in is required"))
	}

	pub := readPub(*pubPath)
	data := readFile(*inPath)

	var s SigJSON
	readJSON(*sigPath, &s)
	if s.Type != "ed25519" {
		fatal(errors.New("signature type mismatch"))
	}
	sig := mustB64(s.Signature)

	ok := ed25519.Verify(pub, data, sig)
	if ok {
		fmt.Println("Signature: VALID ✅")
	} else {
		fmt.Println("Signature: INVALID ❌")
		os.Exit(1)
	}
}

// ---------------- pub (derive) ----------------
func cmdPub(args []string) {
	fs := flag.NewFlagSet("pub", flag.ExitOnError)
	keyPath := fs.String("key", "priv.json", "private key JSON path")
	outPath := fs.String("out", "pub.json", "public key JSON output path")
	_ = fs.Parse(args)

	_, pub := readPriv(*keyPath)
	pubJSON := PubJSON{Type: "ed25519", Public: b64(pub)}
	writeJSON(*outPath, pubJSON, 0644)
	fmt.Printf("Derived public key -> %s\n", *outPath)
}

// ---------------- helpers ----------------
func readPriv(path string) (ed25519.PrivateKey, ed25519.PublicKey) {
	var p PrivJSON
	readJSON(path, &p)
	if p.Type != "ed25519" {
		fatal(errors.New("private key type mismatch"))
	}
	seed := mustB64(p.Seed)
	priv := ed25519.NewKeyFromSeed(seed)
	return priv, priv.Public().(ed25519.PublicKey)
}

func readPub(path string) ed25519.PublicKey {
	var p PubJSON
	readJSON(path, &p)
	if p.Type != "ed25519" {
		fatal(errors.New("public key type mismatch"))
	}
	return mustB64(p.Public)
}

func writeJSON(path string, v any, perm os.FileMode) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, perm)
	check(err)
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	check(enc.Encode(v))
}

func readJSON(path string, v any) {
	f, err := os.Open(path)
	check(err)
	defer f.Close()
	dec := json.NewDecoder(f)
	check(dec.Decode(v))
}

func readFile(path string) []byte {
	b, err := os.ReadFile(path)
	check(err)
	return b
}

func b64(b []byte) string { return base64.StdEncoding.EncodeToString(b) }
func mustB64(s string) []byte {
	b, err := base64.StdEncoding.DecodeString(s)
	check(err)
	return b
}

func check(err error) {
	if err != nil {
		fatal(err)
	}
}
func fatal(err error) {
	if err == nil {
		return
	}
	_, _ = io.WriteString(os.Stderr, "Error: "+err.Error()+"\n")
	os.Exit(1)
}
