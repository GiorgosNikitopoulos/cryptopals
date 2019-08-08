package main

import "fmt"
import "encoding/hex"

func xor(first_operand []byte, second_operand []byte) ([]byte){
	length := len(first_operand)
	return_value := make([]byte, length)
	for i := 0; i < length; i++ {
		return_value[i] = first_operand[i] ^ second_operand[i % 3]
	}
	return return_value
}
func main() {
	ice_string := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"
	ice_bytes := []byte(ice_string)
	key_bytes := []byte(key)
	output_bytes := xor(ice_bytes, key_bytes)
	encoded_str := hex.EncodeToString(output_bytes)
	fmt.Printf("%s\n", encoded_str)

}
