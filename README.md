<h1 align="center">RLP Decoder</h1>

## About The Project

RECURSIVE-LENGTH PREFIX (RLP) SERIALIZATION Decoder

https://ethereum.org/en/developers/docs/data-structures-and-encoding/rlp/ 

Decoder works on 
1. Single byte - (Hex [0x00, 0x7f]) (decimal [0, 127])
2. Short string - (Hex [0x80, 0xb7]) (decimal [128, 183])
3. Long string - (Hex [0xb8, 0xbf]) (decimal [184, 191])
4. Short list - (Hex [0xc0, 0xf7]) (decimal [192, 247])
5. Long list - (Hex [0xf8, 0xff]) (decimal [248, 255])

## Prerequisites

Download and install [Golang 1.20](https://go.dev/doc/install) (or higher).  

## How To Use?

1. Navigate to rlpDecoder/:
   ``` 
   cd /path/to/folder/sarva-assignment/rlpDecoder/
   ``` 
2. Get dependencies:
   ``` 
   go mod tidy
   ```
3. Run the app:
   ``` 
   go run rlpDecoder.go 
   ```
