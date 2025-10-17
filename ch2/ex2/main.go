// SPDX-License-Identifier: MIT

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	feetPerMeter = 3.28084
	meterPerFoot = 0.3048
	lbPerKg      = 2.2046226218
	kgPerLb      = 0.45359237
	separator    = "======================"
)

func fail(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			err := printOutput(strings.TrimSpace(input.Text()))
			fail(err)
		}
		fail(input.Err())

	} else {
		for _, arg := range args {
			err := printOutput(strings.TrimSpace(arg))
			fail(err)
		}
	}
}

func printOutput(arg string) error {
	num, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		return err
	}
	cToF := fmtNum(num*1.8 + 32)
	fToC := fmtNum((num - 32) * 5 / 9)
	fmt.Printf("%s\n\t%g°C = %s°F\n\t%g°F = %s°C\n",
		"Temperature:", num, cToF, num, fToC)

	mToFt := fmtNum(num * feetPerMeter)
	ftToM := fmtNum(num * meterPerFoot)
	fmt.Printf("%s\n\t%gm = %sft\n\t%gft = %sm\n",
		"Length:", num, mToFt, num, ftToM)

	kgToLb := fmtNum(num * lbPerKg)
	lbToKg := fmtNum(num * kgPerLb)
	fmt.Printf("%s\n\t%gkg = %slb\n\t%glb = %skg\n%s\n",
		"Mass:", num, kgToLb, num, lbToKg, separator)
	return nil
}

func fmtNum(num float64) string {
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", num), "0"), ".")
}
