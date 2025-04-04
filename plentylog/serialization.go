package plentylog

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// textSerialization serializes a log record into a human-readable string format.
func textSerialization(l Record) string {
	var buffer bytes.Buffer

	buffer.WriteString(l.Timestamp.Format("2006-01-02 15:04:05"))
	buffer.WriteString(" ")
	buffer.WriteString(string(l.Level))
	buffer.WriteString(" \"")
	buffer.WriteString(l.Message)
	buffer.WriteString("\" ")
	if l.TransactionID != "" {
		buffer.WriteString("transaction id: ")
		buffer.WriteString(l.TransactionID)
		buffer.WriteString(", ")
	}

	index := 0
	for key, value := range l.Metadata {
		buffer.WriteString(key)
		buffer.WriteString(": ")
		buffer.WriteString(fmt.Sprintf("%v", value))
		if index < len(l.Metadata)-1 {
			buffer.WriteString(", ")
		}
		index++
	}

	return buffer.String()
}

// jsonSerialization serializes a log record into a JSON format.
func jsonSerialization(l Record) (*string, error) {
	jsonBytes, err := json.Marshal(l)
	if err != nil {
		return nil, err
	}

	jsonString := string(jsonBytes)

	return &jsonString, nil
}
