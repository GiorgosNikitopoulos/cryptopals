package main

import "fmt"
import "encoding/hex"
//import "encoding/base64"

func xor(first_operand []byte, second_operand []byte) ([]byte){
	length := len(first_operand)
	return_value := make([]byte, length)
	for i := 0; i < length; i++ {
		return_value[i] = first_operand[i] ^ second_operand[i]
	}
	return return_value
}
func main() {
	first_hex := "1c0111001f010100061a024b53535009181c"
	second_hex := "686974207468652062756c6c277320657965"
	first_hexdata, _ := hex.DecodeString(first_hex)
	second_hexdata, _ := hex.DecodeString(second_hex)
	fmt.Printf("%s", hexdata)
	xored_data := xor(first_hexdata, second_hexdata, length)
	fmt.Printf("xor data: %s", xored_data)
	fmt.Println("This is the first programme in go.")
}
