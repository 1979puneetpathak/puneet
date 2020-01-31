package main

import (
	"flag"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

func main() {

	dictionary := flag.String("dictionary", "", "A dictionary file, where each line comprises one " +
		"dictionary word from which you can create your dictionary. E.g. “and”, “bath”, etc, but note the dictionary words " +
		"do not need to be real words.")

	input := flag.String("input", "", "An input file that contains a list of long strings, each on a " +
		"newline, that you will need touse to search for your dictionary words. E.g. “btahand”")

	flag.Parse()


	serch_line, serch_text, err := validateInput(*input, *dictionary)

	if err != nil {
		fmt.Println(err)
		return
	}

	final_output := make(map[string]int)

	case_num := 0

	for _,line := range serch_line {

		case_num += 1

		match := 0

		//fmt.Println(line)

		sclice_line := []rune(line)

		//char_match_start_end := []string{}

		for _,text := range serch_text {
			//fmt.Println(text)

			//fmt.Println(len(text))

			//fmt.Println(text[:1])

			text_last_char := text[len(text)-1:len(text)]

			re := regexp.MustCompile(text[:1])

			//fmt.Println(re.FindAllStringSubmatchIndex(line, -1))

			match_sclice := re.FindAllStringSubmatchIndex(line, -1)

			check_unique_map := make(map[string]bool)

			for _,msclice := range match_sclice {
				//fmt.Println(msclice[0])

				char_match_start := string(sclice_line[msclice[0]:msclice[0]+len(text)])

				//match last char of input text and extracted text
				if text_last_char == char_match_start[len(char_match_start)-1:len(char_match_start)] {

					//char_match_start_end = append(char_match_start_end, char_match_start)

					is_match := false

					if len(text) > 3 {
						input_text_middle_char := SortStringByCharacter(text[1:len(text)-1])

						extracted_text_middle_char := SortStringByCharacter(char_match_start[1:len(char_match_start)-1])

						if input_text_middle_char == extracted_text_middle_char {
							is_match = true
						}
					} else if char_match_start == text {
						is_match = true
					}

					if _,ok := check_unique_map[char_match_start]; ok == false && is_match {
						//fmt.Println(char_match_start)
						check_unique_map[char_match_start] = true
						match += 1
					}
				}
			}
		}

		final_output["Case #" + strconv.Itoa(case_num)] = match

	}

	fmt.Println(final_output)
}

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	r := StringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}