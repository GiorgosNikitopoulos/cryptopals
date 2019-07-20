package main

import "fmt"
import "encoding/hex"
import "encoding/base64"

func hextob64(unencoded []byte) ([]byte){
	encodeddata := make([]byte, base64.StdEncoding.EncodedLen(len(unencoded)))
	base64.StdEncoding.Encode(encodeddata, unencoded)
	return encodeddata
}
func main() {
	hexstring := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	hexdata, _ := hex.DecodeString(hexstring)
	fmt.Printf("%s", hexdata)
	base64data := hextob64(hexdata)
	fmt.Printf("base64 data: %s", base64data)
	fmt.Println("This is the first programme in go.")
}
