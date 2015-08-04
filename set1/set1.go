/*
This package implements all of the necessary functions for the
Matasano Cryptopals challenges set 1.
*/
package set1


// I will also need to import    "encoding/base64"

import (
    "fmt"
    "encoding/hex"
    "encoding/base64"
)



func HexTo64(str string) string {
    bytes,err := hex.DecodeString(str)
    if err == nil {
        return err.Error()
    }
    
    return 
}
