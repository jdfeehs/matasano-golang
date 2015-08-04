package main

import(
    "matasano-golang/set1"
    "fmt"
)



func main() {

	fmt.Println("-------------------------------------")
	fmt.Println("Test 1")
	fmt.Println("-------------------------------------")
    str := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    ret1,err1 := set1.HexTo64(str)

    if err1 == nil{
		fmt.Println(ret1)
	} else {
		fmt.Println(err1)
	}

	fmt.Println("-------------------------------------")
	fmt.Println("Test 2")
	fmt.Println("-------------------------------------")
	str1 := "1c0111001f010100061a024b53535009181c"
	str2 := "686974207468652062756c6c277320657965"

	ret2,err2 := set1.FixedXOR(str1,str2)
	if err2 == nil{
		fmt.Println(ret2)
	} else {
		fmt.Println(err2)
	}
}
