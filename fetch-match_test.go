package main

import (
	"reflect"
	"testing"
)

type testInputResult struct {
	inputText []string
	dictionary []string
}

func TestResult(t *testing.T)  {
	var inputR testInputResult
	expectedResult := make(map[string]int)

	inputR.dictionary = []string{"abc","klmn","oepep"}
	inputR.inputText = []string{"jdhabcjfkshdfkmlnswdf"}

	//expected result
	expectedResult["Case #1"] = 2

	receivedResult, err := getMatches(inputR.inputText, inputR.dictionary)

	if err != nil {
		t.Errorf("getMatches Failed, expected %v got %v", expectedResult, err)
	} else if reflect.DeepEqual(expectedResult, receivedResult){
		t.Errorf("getMatches Success, expected %v got %v", expectedResult, receivedResult)
	} else {
		t.Errorf("getMatches Failed, expected %v got %v", expectedResult, receivedResult)
	}
}