package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
)



func squareandmultiply (a,b,c big.Int)(big.Int){//result = a^b mod c
	var p,r big.Int
	x := a
	y := b
	n := c
	var result big.Int
	p.Set(&y)
	r.Set(&x)
	buffer := big.NewInt(1)
	result.Set(buffer)
	for p.BitLen()>0{
		if p.Bit(0)!=0 {
			result.Mul(&result,&r)
			result.Mod(&result,&n)
		}
		p.Rsh(&p,1)
		r.Mul(&r,&r)
		r.Mod(&r,&n)
	}

	return result
}


func textSeperator(b []byte)(*big.Int,*big.Int,*big.Int){
	var p,g,g_a big.Int
	file := string(b)
	file_string := strings.Split(file,",")
	file_string_p := strings.Split(file_string[0],"(")
	file_string_g :=file_string[1]
	file_string_g_a := strings.Split(file_string[len(file_string)-1],")")
	p.SetString(file_string_p[1],10)
	g.SetString(file_string_g,10)
	g_a.SetString(file_string_g_a[0],10)
	return &p,&g,&g_a
}

func exponentiation(num *big.Int,exp *big.Int,result *big.Int){
	var p,r big.Int
	p.Set(exp)
	r.Set(num)
	buffer := big.NewInt(1)
	result.Set(buffer)
	for p.BitLen()>0{
		if p.Bit(0)!=0 {
			result.Mul(result,&r)

		}
		p.Rsh(&p,1)
		r.Mul(&r,&r)
	}
}

func guessKey(p,g,h big.Int){
	var result,p_copy,power1,power2 big.Int
	p_copy.Set(&p)
	p_copy.Add(&p_copy,big.NewInt(1))
	half := new(big.Int).Div(&p_copy,big.NewInt(2))
	result = squareandmultiply(g,*half,p)
	if result.Cmp(&h)==0{
		fmt.Println(half)
		os.Exit(1)
	}
	power1.Sub(half,big.NewInt(1))
	power2.Add(half,big.NewInt(1))
	for {
		if power1.Cmp(big.NewInt(0))==1{
			result = squareandmultiply(g,power1,p)
			if result.Cmp(&h)==0{
				fmt.Println(&power1)
				os.Exit(1)
			}
			power1.Sub(&power1,big.NewInt(1))
		}
		if power2.Cmp(&p)==-1{
			result = squareandmultiply(g,power2,p)
			if result.Cmp(&h)==0{
				fmt.Println(&power2)
				os.Exit(1)
			}
			power2.Add(&power2,big.NewInt(1))
		}
	}
	fmt.Println("Error")

}

func main(){
	if len(os.Args)!=2{
		fmt.Println("Usage Error")
		os.Exit(1)
	}
	filename := os.Args[1]
	file_content,err := os.Open(filename)
	if err!=nil{
		fmt.Println("Input File Error")
		os.Exit(1)
	}
	defer file_content.Close()
	file_content_string,_ := ioutil.ReadAll(file_content)
	p,g,h := textSeperator([]byte(file_content_string))
	guessKey(*p,*g,*h)
}


