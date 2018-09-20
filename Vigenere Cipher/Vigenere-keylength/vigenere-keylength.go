package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"sort"
)

func showMessage(){
	fmt.Println("Usage :  vigenere-keylength <ciphertext file>")
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
func HCF (num1 int, num2 int) int {
	var small int
	var large int
	var buff int
	if num1>num2{
		small = num1
		large = num2	
	}else{
		small = num1
		large = num2	
	}
	for small !=0 {
		buff = (large % small)
		large = small
		small = buff	
	}
	return large	
	}
func keyLen (ctext string){
	var result int
	freqCount := make([]int,100,100)
	for i:=1;i<100;i++ {
		for j:=0; j<len(ctext);j++{
			if((j+i)>=len(ctext)){
				break			
			}
			if(string(ctext[j])==string(ctext[j+i])){
				freqCount[i]++			
			}			
		}	
	}
		
	sortedFreqCount := make([]int,100, 100)
	for j:=0;j<100;j++{
		sortedFreqCount[j] = freqCount[j]
	}
	sort.Ints(sortedFreqCount)	
	largeIndex := make([]int,4,4)
	for i:=0;i<4;i++ {
		for j:=0;j<len(freqCount);j++ {
			if freqCount[j] == sortedFreqCount[99-i]{
				largeIndex[i] = j 
				break			
			}		
		}	
	}	
	result = HCF(largeIndex[0],largeIndex[1])
	for i:=2;i<len(largeIndex);i++ {
		result = HCF(result, largeIndex[i]) 
	}
	if result == 1 {
		sort.Ints(largeIndex)
		result = largeIndex[0]+1	
	}
	fmt.Println("Guessed Key Length = ",result)	
	} 

func main(){
	if len(os.Args)<2{
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
	keyLen(ctext)
	}
