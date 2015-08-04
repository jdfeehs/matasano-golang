/*
This package implements all of the necessary functions for the
Matasano Cryptopals challenges set 1.
*/
package set1

import (
    "fmt"
    "encoding/base64"
    "encoding/hex"

)



func HexTo64(str string) string {
    bytes,err := hex.DecodeString(str)
    if err == nil {
        return err.Error()
    }
    return string
}
