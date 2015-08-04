/*
This package implements all of the necessary functions for the
Matasano Cryptopals challenges set 1.
*/
package set1


// I will also need to import    "encoding/base64"

import (
    "encoding/hex"
    "encoding/base64"
	"errors"
)


//Challenge 1
func HexTo64(str string) (string,error) {
    bytes,err := hex.DecodeString(str)
    if err != nil {
        return "",err
    }
    ret := base64.StdEncoding.EncodeToString(bytes)    
    return ret,nil
}

//Challenge 2
func FixedXOR(buf1 string, buf2 string) (string,error) {
	//Convert the strings to bytes, thne error check
	hex1,err1 := hex.DecodeString(buf1)
	hex2,err2 := hex.DecodeString(buf2)

	if err1 != nil{
		return "", err1
	}
	if err2 != nil{
		return "", err2
	}
	if len(hex1) != len(hex2){
		return "",errors.New("String lengths not equal")
	}

	for i := range(hex1){
		hex2[i] ^= hex1[i]
	}
	return hex.EncodeToString(hex2),nil

}



