package main

import (
	"fmt"
	"math/big"
	"strings"
)

var orderOfMagnitudes = []string{"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion", "sextillion"}

var teens = map[string]string{
	"10": "ten",
	"11": "eleven",
	"12": "twelve",
	"13": "thirteen",
	"14": "fourteen",
	"15": "fifteen",
	"16": "sixteen",
	"17": "seventeen",
	"18": "eighteen",
	"19": "nineteen",
}

var tens = map[string]string{
	"2": "twenty",
	"3": "thirty",
	"4": "fourty",
	"5": "fifty",
	"6": "sixty",
	"7": "seventy",
	"8": "eighty",
	"9": "ninety",
}

var ones = map[string]string{
	"1": "one",
	"2": "two",
	"3": "three",
	"4": "four",
	"5": "five",
	"6": "six",
	"7": "seven",
	"8": "eight",
	"9": "nine",
}

func convertTripple(tri string) string {
	tri = fmt.Sprintf("%03s", tri)
	if tri == "000" {
		return ""
	}
	builder := []string{}

	// hundreds place
	if tri[:1] != "0" {
		builder = []string{fmt.Sprintf("%s hundred", ones[tri[:1]])}
	}

	// teens or (tens and ones)
	if tri[1:2] == "1" {
		builder = append([]string{fmt.Sprintf("%s", teens[tri[1:]])}, builder...)
	} else {
		// tens
		t, okt := tens[tri[1:2]]
		if okt {
			builder = append(builder, fmt.Sprintf("%s", t))
		}
		// ones
		o, oko := ones[tri[2:]]
		if oko {
			builder = append(builder, fmt.Sprintf("%s", o))
		}
	}

	return strings.Join(builder, " ")
}

func convert(i *big.Int) string {
	if i.Cmp(big.NewInt(0)) == 0 {
		return "zero"
	}

	s := i.String()
	numWordList := []string{}

	j := 0
	for len(s) > 0 {
		min := 3
		if len(s) < 3 {
			min = len(s)
		}
		idx := len(s) - min
		numWords := convertTripple(s[idx:])
		if len(numWords) > 0 {
			numWordList = append([]string{fmt.Sprintf("%s %s", numWords, orderOfMagnitudes[j])}, numWordList...)
		}
		s = s[:idx]
		j++
	}
	return strings.Join(numWordList, " ")
}

func main() {
	sextillion, thousand, zero := &big.Int{}, &big.Int{}, &big.Int{}
	sextillion.UnmarshalText([]byte("1014045600000000000123"))
	fmt.Printf("%s\n", convert(sextillion))
	thousand.UnmarshalText([]byte("1000"))
	fmt.Printf("%s\n", convert(thousand))
	zero.UnmarshalText([]byte("0"))
	fmt.Printf("%s\n", convert(zero))
}
