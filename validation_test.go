package main

import "testing"

type testInput struct {
	inputText []string
	dictionary []string
}

func TestValidation(t *testing.T)  {

	//check for empty input
	serch_line, serch_text, err := readInput("", "") //error expected

	if err == nil {
		t.Errorf("validateInput Failed, expected error message got %v and %v", serch_line, serch_text)
	} else {
		t.Errorf("validateInput Success, error message received : %v", err)
	}

	//check for empty dictionary file
	var input testInput
	err = validateInput(input.dictionary, input.inputText)

	if err == nil {
		t.Errorf("validateInput Failed, expected error message got %v", err)
	} else {
		t.Errorf("validateInput Success, error message received : %v", err)
	}

	//check for empty input file
	input.inputText = append(input.inputText, "hfnvbdmbvbdmxvmxcvbmxcnvbmnnmsdv")

	err = validateInput(input.dictionary, input.inputText)

	if err == nil {
		t.Errorf("validateInput Failed, expected error message got %v", err)
	} else {
		t.Errorf("validateInput Success, error message received : %v", err)
	}

	//check for ambiguous text in dictionary file
	input.dictionary = []string{"axdf","dsdd","axdf"}

	err = validateInput(input.dictionary, input.inputText)

	if err == nil {
		t.Errorf("validateInput Failed, expected error message got %v", err)
	} else {
		t.Errorf("validateInput Success, error message received : %v", err)
	}

	//check for word length in dictionary file, word length should be between 2 - 105
	input.dictionary = input.dictionary[:2]
	input.dictionary = append(input.dictionary, "bndbhbjdbjbdnvcbnmbmvbxmbbcnmbmbcvmcvmnbvmnbmnbvbmnfdfdbffddbkjjkuhjkmnvdbmnbnmvdsxcvbmbmcxvndbhbjdbjbdnvcbnmbmvbxmbbcnmbmbcvmcvmnbvmnbmnbvbmnfdfdbffddbkjjkuhjkmnvdbmnbnmvdsxcvbmbmcxvndbhbjdbjbdnvcbnmbmvbxmbbcnmbmbcvmcvmnbvmnbmnbvbmnfdfdbffddbkjjkuhjkmnvdbmnbnmvdsxcvbmbmcxvndbhbjdbjbdnvcbnmbmvbxmbbcnmbmbcvmcvmnbvmnbmnbvbmnfdfdbffddbkjjkuhjkmnvdbmnbnmvdsxcvbmbmcxv")

	err = validateInput(input.dictionary, input.inputText)

	if err == nil {
		t.Errorf("validateInput Failed, expected error message got %v", err)
	} else {
		t.Errorf("validateInput Success, error message received : %v", err)
	}

	//check for total length of words in dictionary, it should not be more than 105
	input.dictionary = input.dictionary[:2]
	input.dictionary = append(input.dictionary, "dsfdfbdfvhjvsdgjgjhvgdjgjhvdghjghjbdvbjbjsdvbjhbghjdvsbhjbhjsdvbhjsdhjhjsdfhjhj")
	input.dictionary = append(input.dictionary, "jndxvdvbnmbjhbdvjhbjhbjkdvbjkbdvjkbjkbnjdsvbnjkbjkvdbkbkdnvknbjknvjknmcxvmbnjknk")
	input.dictionary = append(input.dictionary, "cmvncbvmbmnbauidshiubdsmvsgvjhvcsdgcvushjbkciskhcibkjsuigvcueisauwiihpaasdwoew")


	err = validateInput(input.dictionary, input.inputText)

	if err == nil {
		t.Errorf("validateInput Failed, expected error message got %v", err)
	} else {
		t.Errorf("validateInput Success, error message received : %v", err)
	}

}
