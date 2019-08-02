package extractor

import (
	"fmt"
	"strings"
	"testing"
)

func TestExtractEmptyPath(t *testing.T) {
	emptyFilePath := ""
	//var expected Reading
	var rule ExtractRule

	_, err := Extract(emptyFilePath, rule)
	if err == nil {
		t.Errorf("Unexpected results, must have error return")
	}

	expectedErrorMessage := "filePath is empty"
	if !strings.Contains(err.Error(), expectedErrorMessage) {
		t.Errorf("Unexpected error message: %v, and expected: %v.", err.Error(), expectedErrorMessage)
	}
}

func TestExtractNormalPath(t *testing.T) {
	filePath := ""
	var expected Reading
	expected.device = "device1"
	expected.ReadingData = []ReadingData {
		ReadingData{
			Name: "123",
			Value: "test",
		},
	}

	fmt.Println(expected)

	var rule ExtractRule
	result, err := Extract(filePath, rule)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	fmt.Println(result)
	//if result != expected {
	//	t.Errorf("Unexpected results, must have error return")
	//}
}