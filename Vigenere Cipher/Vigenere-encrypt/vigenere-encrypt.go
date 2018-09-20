package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"io"
	"strings"
)

func showMessage(){
	fmt.Println("Vigenere Cipher tool for encryption");
	fmt.Println("Usage :  vigenere-encrypt <encryption key> <plaintext file>")
}

func garbageCollection(buff string) string {
	out := []rune{}
	for _,r := range buff {
		if r>=65 && r<=90{
			out = append(out,r)		
		} else if r>=97 && r<=122 {
			out = append(out,r-32)		
		}	
	}
		return string(out)	
	}

func encrypt (p rune, k rune) rune {
	return (((p+k)%26)+'A')
	}

func makeCipherText (key string, plaintext string) string {
	output := make([]rune,0,len(plaintext))
	for i,r := range plaintext {
		output = append(output,encrypt(r,rune(key[i%len(key)])))	
	}
	return string(output)	
	} 

func outputFile(cipherText string) {
	_, err := os.Stat("ciphertext.txt")
	if err == nil {
		fmt.Println("File already exists")
		os.Exit(1)	
	}
	fileinfo, err := os.Create("ciphertext.txt")
	if err != nil {
		fmt.Println("File cannot be created")
		os.Exit(1)	
	}
	defer fileinfo.Close()
	_,err = io.Copy(fileinfo, strings.NewReader(cipherText))
	_,err = io.Copy(fileinfo, strings.NewReader("\n"))	
	}

func main(){
	var ptext string
	if len(os.Args)<3{
	showMessage()
	os.Exit(1)	
	}
	
	if len(os.Args[1])>32{
	fmt.Println("Key Length Overflow")
	os.Exit(1)	
	}
	
	plaintext, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
	fmt.Println("Check the input file")
	os.Exit(1)	
	}
	ptext = string(plaintext)
	ptext = garbageCollection(ptext)
	key := garbageCollection(os.Args[1])
	cipherText := makeCipherText (key,ptext)
	fmt.Println("file contents %s", ptext)
	fmt.Println("key", key)	
	fmt.Println ("Ciphertext",cipherText)
	outputFile(cipherText)
}
