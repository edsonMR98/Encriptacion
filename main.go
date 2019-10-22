package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
)

func copyText(texto string, object hash.Hash) {
	f, err := os.Open(texto)
	if err != nil {
		fmt.Println("Error al abrir el archivo")
	} else {
		io.Copy(object, f) // Copy the text content to md5 object
	}
	f.Close()
}

func main() {
	txtFlag := flag.String("file", "", "File name to encrypt")
	strFlag := flag.String("str", "", "String to encrypt")
	flag.Parse()

	md5 := md5.New()
	sha1 := sha1.New()
	sha256 := sha256.New()

	if *txtFlag != "" {
		copyText(*txtFlag, md5)
		copyText(*txtFlag, sha1)
		copyText(*txtFlag, sha256)		
	} else if *strFlag != "" {
		io.WriteString(md5, *strFlag)  // Copy the string to md5
		io.WriteString(sha1, *strFlag) // Copy the string to sha1
		sha256.Write([]byte(*strFlag)) // Copy the striing to sha256
	}

	fmt.Printf("MD5: %x \n", md5.Sum(nil))      // Print md5
	fmt.Printf("SHA-1: %x \n", sha1.Sum(nil))   // Print sha1
	fmt.Printf("SHA256: %x\n", sha256.Sum(nil)) // Print sha256
}
