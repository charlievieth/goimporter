package goimporter

import (
	"fmt"
	"go/build"
	"testing"
)

var packageNames = []string{
	"archive/tar",
	"archive/zip",
	"bufio",
	"bytes",
	"crypto/aes",
	"crypto/cipher",
	"crypto/des",
	"crypto/dsa",
	"crypto/ecdsa",
	"crypto/elliptic",
	"crypto/hmac",
	"crypto/md5",
	"crypto/rand",
	"crypto/rc4",
	"crypto/rsa",
	"crypto/sha1",
	"crypto/sha256",
	"crypto/sha512",
	"crypto/subtle",
	"crypto/tls",
	"crypto/x509",
	"encoding/ascii85",
	"encoding/asn1",
	"encoding/base32",
	"encoding/base64",
	"encoding/binary",
	"encoding/csv",
	"encoding/gob",
	"encoding/hex",
	"encoding/json",
	"encoding/pem",
	"encoding/xml",
	"go/ast",
	"go/build",
	"go/scanner",
	"go/token",
	"net",
	"net/http",
	"strings",
}

var context = DefaultContext()

func init() {
	n := 0
	var last error
	for i, path := range packageNames {
		_, last = build.Import(path, ".", build.FindOnly|build.AllowBinary)
		if last == nil {
			packageNames[n] = packageNames[i]
			n++
		}
	}
	if n == 0 {
		panic(fmt.Sprintf("Error importing packages for test: %s", last))
	}
	packageNames = packageNames[:n]
}

func TestCache(t *testing.T) {
	c := NewCache(len(packageNames))
	m := c.Importer(context)
	for _, path := range packageNames {
		if _, err := m.Import(path); err != nil {
			t.Fatalf("TestCache (%s): %s", path, err)
		}
	}
	for _, path := range packageNames {
		if _, ok := c.c.Get(context, path); !ok {
			t.Errorf("Cache miss: %s", path)
		}
		if _, err := m.Import(path); err != nil {
			t.Fatalf("TestCache (%s): %s", path, err)
		}
	}
}

func BenchmarkCache(b *testing.B) {
	m := NewCache(len(packageNames)).Importer(DefaultContext())
	b.ResetTimer()
	n := 0
	for i := 0; i < b.N; i++ {
		if _, err := m.Import(packageNames[n%len(packageNames)]); err != nil {
			b.Fatal(err)
		}
		n++
	}
}
