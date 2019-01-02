package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"math"
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



/*Baby Step Giant Step*/
func guessKey(p,g,h big.Int){
	var power,y big.Int
	var i,j int32
	buffer := big.NewInt(0).Sub(&p,big.NewInt(1))
	buffer.Sqrt(buffer)
	n := int32(math.Ceil(float64(buffer.Int64())))+1
	m:=make(map[string]int32)
	for i=0;i<n;i++{
		index := squareandmultiply(g,*big.NewInt(int64(i)),p)
		m[index.String()]=i
	}
	p_2:=big.NewInt(0).Sub(&p,big.NewInt(int64(2)))
	power.Mul(big.NewInt(int64(n)),p_2)
	c := squareandmultiply(g,power,p)
	for j=0;j<n;j++{
		big_j := big.NewInt(int64(j))
		buffer:= squareandmultiply(c,*big_j,p)
		buffer1 := big.NewInt(0).Mul(&h,&buffer)
		y.Mod(buffer1,&p)
		if val,ok := m[y.String()];ok{
				r := int32(j*n)
				result := int32(r+val)
				fmt.Println(result)
				return 

		}
	}

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


