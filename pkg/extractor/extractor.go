package extractor

import (
	"fmt"
	"log"
	"time"
)

type ExtractRule struct {
	Id              string
	Name            string
	DeviceType      []string
	Created         time.Time
	LastModified    time.Time
	Example         string
	TokenParsedBy   []TokenParsedBy
	TokenToMetadata []TokenToMetadata
}

type TokenParsedBy struct {
	Name       string
	Separators []string
}

type ValueDescriptor struct {
	Id           string
	Created      int64
	Modified     int64
	Origin       int
	Name         string
	Description  string
	Min          string
	Max          string
	Type         string
	UomLabel     string
	DefaultValue string
	Formatting   string
	Labels       []string
}

type TokenToMetadata struct {
	ValueDescriptor ValueDescriptor
	IsExtract       bool
}

type Reading struct {
	device string
	ReadingData []ReadingData
}

type ReadingData struct {
	Name       string
	Value string
}

type ExtractionError struct {
	filePath string
	cause string
}

func (e ExtractionError) Error() string {
	return fmt.Sprintf("An error occurred: %s, filePath: %s", e.cause, e.filePath)
}

func Extract(filePath string, rule ExtractRule) (Reading, error) {
	log.Println("Input filePath: ", filePath)

	var result Reading
	if filePath == "" {
		return result, ExtractionError{filePath, "filePath is empty"}
	}
	//TODO: Implementing the logic of extraction
	return result, nil
}