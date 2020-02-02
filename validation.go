package main

import (
	"bufio"
	"errors"
	"os"
)

func readInput(input string, dictionary string) ([]string,[]string,error) {
	input_file, err := os.Open(input)

	if err != nil {
		return nil,nil,errors.New("INVALID INPUT FILE SPECIFIED IN INPUT PARAMETER")
	}

	dictionary_file, err := os.Open(dictionary)

	if err != nil {
		return nil,nil,errors.New("INVALID DICTIONARY FILE SPECIFIED IN DICTIONARY PARAMETER")
	}

	defer input_file.Close()

	defer dictionary_file.Close()

	scanner := bufio.NewScanner(input_file)

	scanner.Split(bufio.ScanLines)

	input_text := []string{}

	for scanner.Scan() {
		input_text = append(input_text, scanner.Text())
	}

	scanner = bufio.NewScanner(dictionary_file)

	scanner.Split(bufio.ScanLines)

	serch_line := []string{}

	for scanner.Scan() {
		serch_line = append(serch_line, scanner.Text())
	}

	return serch_line,input_text,nil
}

func validateInput(input_text []string, serch_line []string) error  {

	if len(serch_line) <= 0 {
		return errors.New("DICTIONARY FILE CANNOT BE EMPTY")
	}

	if len(input_text) <= 0 {
		return errors.New("INPUT FILE CANNOT BE EMPTY")
	}

	input_text_map := make(map[string]bool)

	all_text_length := 0

	for _, text := range input_text {
		if _,ok := input_text_map[text]; ok == true {
			return errors.New("DICTIONARY FILE HAS AMBIGUOUS WORDS")
		}

		input_text_map[text] = true

		if text_len := len(text); text_len < 2 || text_len > 105 {
			return errors.New("WORD LENGTH IN DICTIONARY SHOULD BE BETWEEN 2 AND 105")
		}

		all_text_length += len(text)
	}

	if all_text_length > 105 {
		return errors.New("TOTAL LENGTH OF WORDS IN DICTIONARY SHOULD NOT EXCEED 105")
	}

	return nil

}
