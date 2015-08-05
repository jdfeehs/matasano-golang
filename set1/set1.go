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
	"fmt"
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
	//Convert the strings to bytes, then error check
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

//Challenge 3

func DecodeXOR(str string) int {
	numSolutions := 0
	for i:=0;i<255; i++ {
		bytes, err := XORChar(str,byte(i))
		if err != nil{
			return -1
		}
		ret := TestEnglish(bytes)
		if ret != "" {
			//We have a potential match!
			//fmt.Println("------")
			//fmt.Println(hex.EncodeToString(bytes))
			fmt.Println(ret)
			numSolutions ++
		}
	}
	return numSolutions
}

/***********************************************************
 * This function will XOR a string by a constant byte
 * It returns the hex string of the decoded bytes
 ***********************************************************/
func XORChar(str string, key byte) ([]byte,error) {
	bytes,err := hex.DecodeString(str)
	if err != nil {
		return nil,err
	}
	for i := range(bytes) {
		bytes[i] ^= key
	}
	return bytes,nil
}

/**********************************************************
 * This function will test if a byte array, when decoded
 * into ASCII, appears to be English.
 *
 * It does this by checking that all of the bytes are in
 * a normal range for ASCII English characters
 *
 * If it appears to be valid English, it returns the string
 * that the byte slice represents. Otherwise, it returns an
 * empty string
 **********************************************************/
func TestEnglish(bytes []byte) string {
	for i := range(bytes){
		if bytes[i] != 0 && bytes[i] < 32 || bytes[i] > 126 {
		//if bytes[i] > 177 {
			//It isn't a valid ASCII character, so return the empty string
			return ""
		}
	}
	//If it reaches here, all of the bytes checked out
	return string(bytes)
}
