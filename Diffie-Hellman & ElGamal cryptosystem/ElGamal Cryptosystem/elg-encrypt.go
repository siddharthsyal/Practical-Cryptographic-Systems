package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
	"encoding/hex"
)

func squareandmultiply (x *big.Int,y *big.Int,n *big.Int,result *big.Int){//result = x^y mod n
	var p,r big.Int
	p.Set(y)
	r.Set(x)
	buffer := big.NewInt(1)
	result.Set(buffer)
	for p.BitLen()>0{
		if p.Bit(0)!=0 {
			result.Mul(result,&r)
			result.Mod(result,n)
		}
		p.Rsh(&p,1)
		r.Mul(&r,&r)
		r.Mod(&r,n)
	}
}


func encrypt(key []byte, plainText string)([]byte){
	cipherBlock_buffer, err := aes.NewCipher(key)
	if err !=nil{
		fmt.Println("AES encryption error")
		os.Exit(1	)
	}
	cipherBlock, _ := cipher.NewGCMWithNonceSize(cipherBlock_buffer,16)
	iv := make([]byte,16)
	_,_ = rand.Read(iv)
	cipherText := cipherBlock.Seal(nil,iv,[]byte(plainText),nil)
	result := make([]byte,len(cipherText)+len(iv))
	result = iv
	for j:=0;j<len(cipherText);j++{
		result = append(result,cipherText[j])		
	}
	
	return result
}

func concatinateandHash (g_a,g_b,g_ab big.Int)([]byte){
	g_a_string := g_a.String()
	g_b_string := g_b.String()
	g_ab_string := g_ab.String()
	buffer := g_a_string+" "+g_b_string+" "+g_ab_string
	hash := sha256.Sum256([]byte(buffer))
	return hash[:]
}

func writetomyfile(g_b big.Int,input[]byte,filename string){
	start := "("
	end := ")"
	comma := ","
	g_b_string := g_b.String()
	final_string :=  start+g_b_string+comma+string(input)+end
	final_byte := []byte(final_string)
	err := ioutil.WriteFile(filename,final_byte,0644)
	if err !=nil{
		panic(err)
	}
}

func textSeperator(b []byte)(*big.Int,*big.Int,*big.Int){
	var p,g,g_a big.Int
	file := string(b)
	file = strings.Replace(file, " ", "", -1)
	file_string := strings.Split(file,",")
	file_string_p := strings.Split(file_string[0],"(")
	file_string_g :=file_string[1]
	file_string_g_a := strings.Split(file_string[len(file_string)-1],")")
	p.SetString(file_string_p[1],10)
	g.SetString(file_string_g,10)
	g_a.SetString(file_string_g_a[0],10)
	return &p,&g,&g_a
}

func main(){
	var g_b,g_ab big.Int
	if len(os.Args)!=4{
		fmt.Println("Usage Error")
		os.Exit(1)
	}
	plainText := os.Args[1]
	public_key := os.Args[2]
	cipherText_file := os.Args[3]
	file_content,err := os.Open(public_key)
	if err!=nil{
		fmt.Println("Input File Error")
		os.Exit(1)
	}
	defer file_content.Close()
	file_content_string,_ := ioutil.ReadAll(file_content)
	p,g,g_a := textSeperator([]byte(file_content_string))
	upper_limit := big.NewInt(0).Sub(p,big.NewInt(2))
	b , _ := rand.Int(rand.Reader,upper_limit)
	squareandmultiply(g,b,p,&g_b)// Calculating g^b
	squareandmultiply(g_a,b,p,&g_ab)//Calculating g^ab
	k := concatinateandHash(*g_a,g_b,g_ab)
	ciphterText :=  encrypt(k,plainText)
	cipherText_hex := make([]byte,hex.EncodedLen(len(ciphterText)))
	hex.Encode(cipherText_hex,ciphterText)
	writetomyfile(g_b,cipherText_hex,cipherText_file)
}

