package main

import "fmt"
import "os"
import "bufio"
import "encoding/hex"
import "github.com/abadojack/whatlanggo"

func xor(first_operand []byte, second_operand byte) ([]byte){
	length := len(first_operand)
	return_value := make([]byte, length)
	for i := 0; i < length; i++ {
		return_value[i] = first_operand[i] ^ second_operand
	}
	return return_value
}
func main() {
	file, _ := os.Open("/home/user/go/src/cryptopals/1_Set/4_challenge/data/datafile")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		hex_string := scanner.Text()
		hex_data, _ := hex.DecodeString(hex_string)
		//fmt.Printf("%s\n", hex_data)
		for i := 0; i < 150; i++ {
			xored_data := xor(hex_data, byte(i))
			info := whatlanggo.Detect(string(xored_data))
			if info.Lang.String() == "English" && info.Confidence > 0.9{
				fmt.Printf("xored data: %s\n", xored_data)
				fmt.Printf("Info data: %s\n", info.Lang.String())
				fmt.Printf("Info data: %f\n", info.Confidence)
			}
		}
	}
}
