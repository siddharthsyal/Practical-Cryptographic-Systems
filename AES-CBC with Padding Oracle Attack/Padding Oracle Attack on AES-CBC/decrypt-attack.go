package main

import (
  "io/ioutil"
  "fmt"
  "os"
  "os/exec"
  "strings"
  "crypto/rand"
)

func pseudoDecrypt (iv []byte, cipherText[] byte)[]byte{
		cipherText = append(iv,cipherText...)	
		buffer := make([]byte, len(cipherText))
		copy(buffer,cipherText)
		blocks := len(cipherText)/16
		for i:=blocks-1;i>0;i--{
			copy(cipherText[len(cipherText)-32:len(cipherText)],buffer[i*16-16:i*16+16])
			newBlock:= Decrypt(cipherText)
			copy(buffer[i*16:i*16+16],newBlock)		
		}
		return buffer[16:]	
	}

// Decrypt method does the actual bruteforcing to achieve the plaintext
	
func Decrypt(block []byte)[]byte{
		org_cipherText := make([]byte ,16)
		copy(org_cipherText,block[len(block)-32:len(block)-16])
		recoveredText := make([]byte,16)
		buff_state := make([]byte,16)
		mod_cipherText := block[len(block)-32 : len(block)-16]
		_,err:=rand.Read(mod_cipherText)
		if err!=nil{
			panic(err)		
		}
		for i:=15;i>=0;i-- {
			pad  :=byte(16-i)	
			for j:=i+1;j<16;j++ {
				mod_cipherText[j] = pad ^ buff_state[j]
			}
			for k:= 0x00;k<0x100;{
				mod_cipherText[i] = byte(k)	
				ioutil.WriteFile("temp.txt",block,0644)	
				out, err1 := exec.Command("./decrypt-test","-i","temp.txt").CombinedOutput()
				if err1 != nil{
					panic(err1)//Haults the execution if the above command did not work well				
				}
				if !strings.Contains(string(out), "INVALID PADDING"){
					break;				
				}
				k++
			}
			buff_state[i] =  pad ^ mod_cipherText[i]		
		}
		  for i := range buff_state {
    			recoveredText[i]= buff_state[i] ^ org_cipherText[i]
		}
		return recoveredText	
	}	

func main() {
	if len(os.Args) !=3{
		fmt.Println("Usage Error")
		os.Exit(1)
	}
	cipherTextFile := os.Args[2]
	raw_data,_ := ioutil.ReadFile(cipherTextFile)
	IV,cipherText := raw_data[:16],raw_data[16:]
	if len(cipherText)%16 != 0{
		fmt.Println("File Input Data Error")
		os.Exit(1)
	}
	plainText_res:=pseudoDecrypt(IV,cipherText)//Plaintext with MAC and Padding
//	fmt.Println("Len",len(plainText_res))
	padLen := int(plainText_res[len(plainText_res)-1])
	rem := len(plainText_res) - padLen - 32
//	fmt.Println("rem ",rem)
	plainText := plainText_res[0:rem]// plaintext with MAC and padding removed
	fmt.Print(string(plainText))
	return
}
