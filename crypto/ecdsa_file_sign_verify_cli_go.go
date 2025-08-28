//ecdsa.go örnek tam hali

// ecdsa_cli.go
// Simple ECDSA file signer & verifier
// Commands:
//   genkey   -curve [P256|P384|P521] -priv priv.pem -pub pub.pem
//   sign     -key priv.pem -in file -out sig.json
//   verify   -key pub.pem -in file -sig sig.json
//
// Usage examples:
//   go run ecdsa_cli.go genkey -curve P256 -priv priv.pem -pub pub.pem
//   go run ecdsa_cli.go sign   -key priv.pem -in README.md -out sig.json
//   go run ecdsa_cli.go verify -key pub.pem  -in README.md -sig sig.json

package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
	"io"
	"math/big"
	"os"
	"strings"
)

type Signature struct {
	R     string `json:"r"`     // hex
	S     string `json:"s"`     // hex
	Curve string `json:"curve"` // P256, P384, P521
	Hash  string `json:"hash"`  // SHA-256, SHA-384, SHA-512
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
	default:
		usage()
		os.Exit(2)
	}
}

func usage() {
	fmt.Println("ECDSA CLI")
	fmt.Println("  genkey  -curve [P256|P384|P521] -priv priv.pem -pub pub.pem")
	fmt.Println("  sign    -key priv.pem -in file -out sig.json")
	fmt.Println("  verify  -key pub.pem  -in file -sig sig.json")
}

// ------------------------ genkey ------------------------
func cmdGenKey(args []string) {
	fs := flag.NewFlagSet("genkey", flag.ExitOnError)
	curveName := fs.String("curve", "P256", "elliptic curve: P256|P384|P521")
	privPath := fs.String("priv", "priv.pem", "private key output path")
	pubPath := fs.String("pub", "pub.pem", "public key output path")
	_ = fs.Parse(args)

	curve, err := parseCurve(*curveName)
	if err != nil {
		fatal(err)
	}

	priv, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		fatal(err)
	}

	if err := writeECPrivateKeyPEM(*privPath, priv); err != nil {
		fatal(err)
	}
	if err := writeECPublicKeyPEM(*pubPath, &priv.PublicKey); err != nil {
		fatal(err)
	}

	fmt.Printf("Generated %s key pair.\nPrivate: %s\nPublic : %s\n", *curveName, *privPath, *pubPath)
}

// ------------------------ sign ------------------------
func cmdSign(args []string) {
	fs := flag.NewFlagSet("sign", flag.ExitOnError)
	keyPath := fs.String("key", "priv.pem", "EC private key (PEM)")
	inPath := fs.String("in", "", "input file to sign")
	outPath := fs.String("out", "sig.json", "signature output path (JSON)")
	_ = fs.Parse(args)

	if *inPath == "" {
		fatal(errors.New("-in is required"))
	}

	priv, curveName, hashName, err := loadECPrivateKey(*keyPath)
	if err != nil {
		fatal(err)
	}

	data, err := os.ReadFile(*inPath)
	if err != nil {
		fatal(err)
	}

	hash := computeHash(hashName, data)
	r, s, err := ecdsa.Sign(rand.Reader, priv, hash)
	if err != nil {
		fatal(err)
	}

	sig := Signature{
		R:     hex.EncodeToString(r.Bytes()),
		S:     hex.EncodeToString(s.Bytes()),
		Curve: curveName,
		Hash:  hashName,
	}
	b, _ := json.MarshalIndent(sig, "", "  ")
	if err := os.WriteFile(*outPath, b, 0644); err != nil {
		fatal(err)
	}
	fmt.Printf("Signed %s -> %s\n", *inPath, *outPath)
}

// ------------------------ verify ------------------------
func cmdVerify(args []string) {
	fs := flag.NewFlagSet("verify", flag.ExitOnError)
	keyPath := fs.String("key", "pub.pem", "EC public key (PEM)")
	inPath := fs.String("in", "", "input file to verify")
	sigPath := fs.String("sig", "sig.json", "signature file (JSON)")
	_ = fs.Parse(args)

	if *inPath == "" {
		fatal(errors.New("-in is required"))
	}

	pub, err := loadECPublicKey(*keyPath)
	if err != nil {
		fatal(err)
	}

	var sig Signature
	b, err := os.ReadFile(*sigPath)
	if err != nil {
		fatal(err)
	}
	if err := json.Unmarshal(b, &sig); err != nil {
		fatal(err)
	}

	data, err := os.ReadFile(*inPath)
	if err != nil {
		fatal(err)
	}
	hash := computeHash(sig.Hash, data)

	rBytes, err := hex.DecodeString(sig.R)
	if err != nil {
		fatal(err)
	}
	sBytes, err := hex.DecodeString(sig.S)
	if err != nil {
		fatal(err)
	}
	r := new(big.Int).SetBytes(rBytes)
	s := new(big.Int).SetBytes(sBytes)

	ok := ecdsa.Verify(pub, hash, r, s)
	if ok {
		fmt.Println("Signature: VALID ✅")
	} else {
		fmt.Println("Signature: INVALID ❌")
		os.Exit(1)
	}
}

// ------------------------ helpers ------------------------
func parseCurve(name string) (elliptic.Curve, error) {
	switch strings.ToUpper(name) {
	case "P256", "NISTP256", "SECP256R1":
		return elliptic.P256(), nil
	case "P384", "NISTP384", "SECP384R1":
		return elliptic.P384(), nil
	case "P521", "NISTP521", "SECP521R1":
		return elliptic.P521(), nil
	default:
		return nil, fmt.Errorf("unknown curve: %s", name)
	}
}

func defaultHashForCurve(curve elliptic.Curve) string {
	switch curve {
	case elliptic.P256():
		return "SHA-256"
	case elliptic.P384():
		return "SHA-384"
	case elliptic.P521():
		return "SHA-512"
	default:
		return "SHA-256"
	}
}

func computeHash(name string, data []byte) []byte {
	switch strings.ToUpper(name) {
	case "SHA-256":
		sum := sha256.Sum256(data)
		return sum[:]
	case "SHA-384":
		sum := sha512.Sum384(data)
		return sum[:]
	case "SHA-512":
		sum := sha512.Sum512(data)
		return sum[:]
	default:
		// fallback
		sum := sha256.Sum256(data)
		return sum[:]
	}
}

func writeECPrivateKeyPEM(path string, priv *ecdsa.PrivateKey) error {
	der, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		return err
	}
	block := &pem.Block{Type: "EC PRIVATE KEY", Bytes: der}
	return os.WriteFile(path, pem.EncodeToMemory(block), 0600)
}

func writeECPublicKeyPEM(path string, pub *ecdsa.PublicKey) error {
	der, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return err
	}
	block := &pem.Block{Type: "PUBLIC KEY", Bytes: der}
	return os.WriteFile(path, pem.EncodeToMemory(block), 0644)
}

func loadECPrivateKey(path string) (*ecdsa.PrivateKey, string, string, error) {
	pemBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, "", "", err
	}
	block, _ := pem.Decode(pemBytes)
	if block == nil || block.Type != "EC PRIVATE KEY" {
		return nil, "", "", errors.New("invalid EC private key PEM")
	}
	priv, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, "", "", err
	}
	curveName := curveString(priv.Curve)
	hashName := defaultHashForCurve(priv.Curve)
	return priv, curveName, hashName, nil
}

func loadECPublicKey(path string) (*ecdsa.PublicKey, error) {
	pemBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(pemBytes)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("invalid public key PEM")
	}
	pubAny, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub, ok := pubAny.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("not an ECDSA public key")
	}
	return pub, nil
}

func curveString(c elliptic.Curve) string {
	switch c {
	case elliptic.P256():
		return "P256"
	case elliptic.P384():
		return "P384"
	case elliptic.P521():
		return "P521"
	default:
		return "Unknown"
	}
}

func fatal(err error) {
	if err == nil {
		return
	}
	_, _ = io.WriteString(os.Stderr, "Error: "+err.Error()+"\n")
	os.Exit(1)
}
