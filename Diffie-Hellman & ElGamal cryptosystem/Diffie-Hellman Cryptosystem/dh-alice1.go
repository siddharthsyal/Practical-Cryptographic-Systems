package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"io/ioutil"
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

func millerRand(lower *big.Int,upper *big.Int,result *big.Int){
	var buffer big.Int
	buffer.Sub(upper,lower)
	g,_ := rand.Int(rand.Reader,&buffer)
	result.Add(lower,g)
}

func isPrime(n big.Int, iterations int) bool{

	big_1 :=big.NewInt(1)
	big__2:=big.NewInt(2)
	if ((n.Cmp(big_1)!=1)||(n.Cmp(big.NewInt(4))==0)){
		return false
	}
	if (n.Cmp(big.NewInt(3))!=1){
		return true
	}
	d:=big.NewInt(0)
	d.Sub(&n,big_1)
	mod_result := big.NewInt(0)
	mod_result.Mod(d,big__2)
	for mod_result.Cmp(big.NewInt(0))==0{
		d.Div(d,big__2)
		mod_result.Mod(d,big__2)
	}
	for i:=0;i<iterations;i++{
		if  miller(*d,n)==false{
			return false

		}
	}
	return true
}

func miller(d ,n big.Int) bool{
	var a,n_2,x,n_1 big.Int
	big_2 := big.NewInt(2)
	big_1 := big.NewInt(1)
	n_1.Sub(&n,big_1)
	n_2.Sub(&n,big_2)
	millerRand(big_2,&n_2,&a)
	squareandmultiply(&a,&d,&n,&x)
	if (big_1.Cmp(&x)==0||n_1.Cmp(&x)==0){
		return true
	}
	for d.Cmp(&n_1)!=0{
		x.Mul(&x,&x)
		x.Mod(&x,&n)
		d.Mul(&d,big_2)
		if big_1.Cmp(&x)==0{
			return false
		}
		if n_1.Cmp(&x)==0{
			return true
		}

	}
	return false
}

/*According to Algorithm 4.80 from HAC*/
func generatorTest(g,p,order *big.Int,factors map[string]*big.Int)bool{
	var buffer big.Int
	for _,val := range factors{
		buffer1 :=big.NewInt(0).Div(order,val)
		squareandmultiply(g,buffer1,p,&buffer)
		if buffer.Cmp(big.NewInt(1))==0{
			return false
		}
	}
	return true
}

/*Note := This func gives Generator(G) of order p-1. For better security, G should be of order Q while using Schnorr primes*/
func generateG(p,q *big.Int) (*big.Int){
	order := big.NewInt(0).Sub(p,big.NewInt(1))
	r:= big.NewInt(0).Div(order,q)
	factors := prime_factor(r)
	factors[q.String()] = q
	i := big.NewInt(2)
	for order.Cmp(i)>0{
		if generatorTest(i,p,order,factors){
			return i
		}
		i = new(big.Int).Add(i, big.NewInt(1))
	}
	return order
}

func prime_factor(a *big.Int) map[string]*big.Int {
	m := make(map[string]*big.Int)

	j:= big.NewInt(2)
	for a.Cmp(big.NewInt(1))==1{
		if big.NewInt(0).Mod(a, j).Cmp(big.NewInt(0)) == 0{
			m[j.String()] = j
			for big.NewInt(0).Mod(a, j).Cmp(big.NewInt(0)) == 0{
				a = big.NewInt(0).Div(a, j)
			}
		}else{
			j = big.NewInt(0).Add(j, big.NewInt(1))
		}
	}
	return m
}

func generatePQ()(big.Int){
	var upper_limit  big.Int
	exponentiation(big.NewInt(2),big.NewInt(1016),&upper_limit)
	q,_ := rand.Int(rand.Reader,&upper_limit)
	for isPrime(*q,5)==false{//Miller Rabin prime number test for Q
		buffer_q,err := rand.Int(rand.Reader,&upper_limit)
		q.Set(buffer_q)
		if err!=nil{
			fmt.Println(err)
			os.Exit(1)
		}
	}
	return *q
}

func generateP()(big.Int,big.Int){
	var buffer,upper_limit, lower_limit big.Int
	i:= int64(2)
	j:=0
	exponentiation(big.NewInt(2),big.NewInt(1024),&upper_limit)
	exponentiation(big.NewInt(2),big.NewInt(1022),&lower_limit)
	big_1:=big.NewInt(1)
	q:=generatePQ()
	buffer.Mul(&q,big.NewInt(i))
	p := big.NewInt(0).Add(&buffer,big_1)
	for {
		i++
		j++
		buffer.Mul(&q,big.NewInt(i))
		p_buff := big.NewInt(0).Add(&buffer,big_1)
		p = p_buff
		if ((isPrime(*p,5)==true)&&(p.Cmp(&upper_limit) == -1)&&(p.Cmp(&lower_limit)==1)){
			break
		}
		if j==512{
			q=generatePQ()
			i = int64(2)
			j=1
		}
	}
	return *p,q
}

func writetomyfile(p,g,a big.Int,filename string){
	start := "("
	end := ")"
	comma := ","
	p_string := p.String()
	g_string := g.String()
	a_string := a.String()
	final_string :=  start+p_string+comma+g_string+comma+a_string+end
	final_byte := []byte(final_string)
	err := ioutil.WriteFile(filename,final_byte,0644)
	if err !=nil{
		panic(err)
	}
}

func writetoBob(p,g, g_a big.Int, filename string){
	start := "("
	end := ")"
	comma := ","
	p_string := p.String()
	g_string := g.String()
	g_a_string := g_a.String()
	final_string :=  start+p_string+comma+g_string+comma+g_a_string+end
	final_byte := []byte(final_string)
	err := ioutil.WriteFile(filename,final_byte,0644)
	if err !=nil{
		panic(err)
	}
}

func main() {
	if len(os.Args)!=3{
		fmt.Println("Usage Error")
		os.Exit(1)
	}
	filename_bob:=os.Args[1]
	filename_secret := os.Args[2]
	p,q:=generateP()
	g := generateG(&p,&q)
	var g_a big.Int
	upper_limit := big.NewInt(0).Sub(&p,big.NewInt(1))
	a,_ := rand.Int(rand.Reader,upper_limit)
	squareandmultiply(g,a,&p,&g_a)
	writetoBob(p,*g,g_a,filename_bob)
	writetomyfile(p,*g,*a,filename_secret)
}
