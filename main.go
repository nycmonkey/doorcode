package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

var (
	digits uint
	ten    = big.NewInt(10)
)

func main() {
	flag.UintVar(&digits, `n`, 5, `number of digits to output`)
	flag.Parse()
	var b strings.Builder
	for i := 0; i < int(digits); i++ {
		zeroToNine, err := rand.Int(rand.Reader, ten)
		if err != nil {
			panic(err)
		}
		b.WriteString(strconv.Itoa(int(zeroToNine.Int64())))
	}
	fmt.Println(b.String())
}
