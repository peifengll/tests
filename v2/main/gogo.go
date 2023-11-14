package main

import (
	"fmt"
	tempconv "github.com/peifengll/tests/v2/tempconv"
)

func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))

}
