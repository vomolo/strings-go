package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"go-reloaded/mani"
)

func main() {
	if len(os.Args[1:]) != 2 {
		fmt.Println("Error: length")
		return
	}

	input := os.Args[1]
	output := os.Args[2]
	suf := ".txt"

	if !strings.HasSuffix(input, suf) {
		fmt.Println("Error: only .txt file extensions allowed!!!")
		return
	}

	if !strings.HasSuffix(output, suf) {
		fmt.Println("Error: only .txt file extensions allowed!!!")
		return
	}

	file1, err := os.Open(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file1.Close()

	stat, err := file1.Stat()
	if err != nil {
		fmt.Println("ERROR: Failed to get input file stats!!!")
		return
	}

	if stat.Size() == 0 {
		fmt.Println("ERROR: File is empty!!!")
		return
	}

	file2, err := os.Create(output)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer file2.Close()

	scanner := bufio.NewScanner(file1)
	for scanner.Scan() {
		line := scanner.Text()
		HexBin := mani.ReplaceHexAndBinNumbers(line)
		UpInt := mani.UpInt(HexBin)
		CapInt := mani.CapInt(UpInt)
		LowInt := mani.LowInt(CapInt)
		Vow := mani.Vow(LowInt)
		Pung := mani.Pung(Vow)
		SingleQ := mani.SingleQ(Pung)
		fmt.Fprintln(file2, SingleQ)
		fmt.Println(SingleQ)

		if err := scanner.Err(); err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
}
