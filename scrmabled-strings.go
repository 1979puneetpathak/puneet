package main

import (
	"flag"
	"fmt"
)

func main() {

	dictionary := flag.String("dictionary", "", "A dictionary file, where each line comprises one " +
		"dictionary word from which you can create your dictionary. E.g. “and”, “bath”, etc, but note the dictionary words " +
		"do not need to be real words.")

	input := flag.String("input", "", "An input file that contains a list of long strings, each on a " +
		"newline, that you will need touse to search for your dictionary words. E.g. “btahand”")

	flag.Parse()


	serch_line, serch_text, err := readInput(*input, *dictionary)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = validateInput(serch_text, serch_line)

	if err != nil {
		fmt.Println(err)
		return
	}

	final_output, err := getMatches(serch_line,serch_text)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(final_output)
}