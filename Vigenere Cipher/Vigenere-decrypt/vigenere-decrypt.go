package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"io"
	"strings"
)

func showMessage(){
	fmt.Println("Usage :  vigenere-decrypt <dencryption key> <ciphertext file>")
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

func decrypt (c rune, k rune) rune {
	return (((c-k+26)%26)+'A')
	}

func makePlainText (key string, cipherText string) string {
	output := make([]rune,0,len(cipherText))
	for i,r := range cipherText {
		output = append(output,decrypt(r,rune(key[i%len(key)])))	
	}
	return string(output)	
	} 

func outputFile(plainText string) {
	_, err := os.Stat("recovered_plaintext.txt")
	if err == nil {
		fmt.Println("File already exists")
		os.Exit(1)	
	}
	fileinfo, err := os.Create("recovered_plaintext.txt")
	if err != nil {
		fmt.Println("File cannot be created")
		os.Exit(1)	
	}
	defer fileinfo.Close()
	_,err = io.Copy(fileinfo, strings.NewReader(plainText))
	_,err = io.Copy(fileinfo, strings.NewReader("\n"))	
	}

func main(){
	var ctext string
	if len(os.Args)<3{
	showMessage()
	os.Exit(1)	
	}
	
	if len(os.Args[1])>32{
	fmt.Println("Key Length Overflow")
	os.Exit(1)	
	}
	
	ciphertext, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
	fmt.Println("Check the input file")
	os.Exit(1)	
	}
	ctext = string(ciphertext)
	ctext = garbageCollection(ctext)
	key := garbageCollection(os.Args[1])
	plainText := makePlainText (key,ctext)
	outputFile(plainText)
}
