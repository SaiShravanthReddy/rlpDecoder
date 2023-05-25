package main

import (
	"encoding/hex"
	"fmt"
)

type RLPData struct {
	Type  string
	Value interface{}
}

func rlpDecode(byteSlice []byte) ([]RLPData, error) {
	var result []RLPData

	for i := 0; i < len(byteSlice); {
		indexByte := byteSlice[i]
		data := RLPData{}

		switch {
		// Single byte - (Hex [0x00, 0x7f]) (decimal [0, 127])
		case indexByte < 0x80:
			data.Type = "Byte"
			data.Value = indexByte
			result = append(result, data)
			i++

		// Short string - (Hex [0x80, 0xb7]) (decimal [128, 183])
		case indexByte < 0xb8:
			length := int(indexByte - 0x80)
			data.Type = "Short String"
			data.Value = string(byteSlice[i+1 : i+length+1])
			result = append(result, data)
			i += length + 1

		// Long string - (Hex [0xb8, 0xbf]) (decimal [184, 191])
		case indexByte < 0xc0:
			length := int(indexByte - 0xb7)
			data.Type = "Long String"
			data.Value = string(byteSlice[i+1 : i+length+1])
			result = append(result, data)
			i += length + 1

		// Short list - (Hex [0xc0, 0xf7]) (decimal [192, 247])
		case indexByte < 0xf8:
			length := int(indexByte - 0xc0)
			subData, err := rlpDecode(byteSlice[i+1 : i+length+1])
			if err != nil {
				return nil, err
			}
			data.Type = "Short List"
			data.Value = subData
			result = append(result, data)
			i += length + 1

		// Long list - (Hex [0xf8, 0xff]) (decimal [248, 255])
		default:
			length := int(indexByte - 0xf7)
			subData, err := rlpDecode(byteSlice[i+1 : i+length+1])
			if err != nil {
				return nil, err
			}
			data.Type = "Long List"
			data.Value = subData
			result = append(result, data)
			i += length + 1
		}
	}

	return result, nil
}

func stringToHex(hexString string) {
	byteSlice, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Println("Error:", err)
	}

	decodedData, err := rlpDecode(byteSlice)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Decoded byteSlice:", decodedData)
}

func main() {
	input1 := "ed90416e616e746861204b726973686e616e8d526168756c204c656e6b616c618d47616e65736820507261736164"
	input2 := "e5922034342e38313538393735343033373334319132302e3435343733343334343535353435"

	stringToHex(input1)
	stringToHex(input2)

}
