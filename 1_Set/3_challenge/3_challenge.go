package main

import "fmt"
import "encoding/hex"
import "github.com/abadojack/whatlanggo"
//import "encoding/base64"

func xor(first_operand []byte, second_operand byte) ([]byte){
	length := len(first_operand)
	return_value := make([]byte, length)
	for i := 0; i < length; i++ {
		return_value[i] = first_operand[i] ^ second_operand
	}
	return return_value
}
func main() {
	hex_string := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	hex_data, _ := hex.DecodeString(hex_string)
	fmt.Printf("%s\n", hex_data)
	for i := 0; i < 150; i++ {
		xored_data := xor(hex_data, byte(i))
		info := whatlanggo.Detect(string(xored_data))
		if info.Lang.String() == "English"{
			fmt.Printf("xored data: %s\n", xored_data)
			fmt.Printf("Info data: %s\n", info.Lang.String())
			fmt.Printf("Info data: %f\n", info.Confidence)
		}
	}	
}
