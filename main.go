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

// copyText opens a specific file (.txt) and copy to the object
// text: string type, is the file name to open
// object: hash.Hash type, is the object in wich will be copied the file contents
func copyText(text string, object hash.Hash) {
	f, err := os.Open(text)
	if err != nil {
		fmt.Println("Error al abrir el archivo")
	} else {
		io.Copy(object, f) // Copy the text contents to object
	}
	f.Close()
}

func main() {
	txtFlag := flag.String("file", "", "File name to encrypt its contents") // Define the file flag
	strFlag := flag.String("str", "", "String to encrypt") // Defin the str flag
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
		sha256.Write([]byte(*strFlag)) // Copy the string to sha256
	}

	fmt.Printf("MD5: %x \n", md5.Sum(nil))      // Print md5
	fmt.Printf("SHA-1: %x \n", sha1.Sum(nil))   // Print sha1
	fmt.Printf("SHA256: %x\n", sha256.Sum(nil)) // Print sha256
}
