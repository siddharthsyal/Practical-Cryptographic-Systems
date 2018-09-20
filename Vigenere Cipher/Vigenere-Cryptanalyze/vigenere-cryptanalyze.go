package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"sort"
	"strconv"
)

func showMessage(){
	fmt.Println("Usage :  vigenere-cryptanalyze <ciphertext file> <key length>")
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

func guessKey(length int, ctext string){
	englishLang := [26]float64 {.08167, .01492, .02792, .04253, .12702, .0228, .02015, .06094, .06966, .0153, .0772, .04025, .02406, .06749, .07507, .01929, .0095, .05987, .06327, .09056, .02758, .00978, .02360, .00150, .01974, .0074}
	key :=make([]string,0,length)

	for i:=0;i<length;i++ { 
		ctextFreq := make([]float64,26,26)
		buff1 := make([]float64,26,26)
		buff2 := make([]float64,26,26)
		for j:=i;j<len(ctext);j=j+length{
			ctextFreq[int(ctext[j])-65]++
		}
		for j:=0;j<26;j++{
			var sum float64 =0
			for k:=0;k<26;k++{
			var modulo int = (k+j)%26 
			sum =sum+(ctextFreq[modulo]*englishLang[k])
			
			}
			buff1[j]=sum
			buff2[j]=sum
		}
		sort.Float64s(buff2)
		for j:=0;j<26;j++{
			if buff1[j]==buff2[25]{
				key = append(key,string(j+65))
		}
	}
	}
		fmt.Println("Guessed key is ",key)		
	} 

func main(){
	if len(os.Args)<3{
	showMessage()
	os.Exit(1)	
	}
		
	ciphertext, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
	fmt.Println("Check the input file")
	os.Exit(1)	
	}
	ctext := string(ciphertext)
	ctext = garbageCollection(ctext)
	len1 := os.Args[2]
	length,err := strconv.Atoi(len1)
	if err != nil {
		fmt.Print(err)
	}
	guessKey(length,ctext)

}
