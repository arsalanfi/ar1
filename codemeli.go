package main

import (
	"fmt"
)

func main() {
	var nID Ncode.NationalID

	nID = "01/2303040"

	fmt.Printf("my national id %s is valid : %v", nID, nID.IsNationalIDvalid(nID))
}
