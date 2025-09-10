package writer

import (
	"github.com/ForceCLI/force-md/internal"
)

// WriteToFile writes the given metadata object to a file using proper XML formatting
func WriteToFile(t interface{}, fileName string) error {
	return internal.WriteToFile(t, fileName)
}

// Marshal serializes the given metadata object to XML bytes
func Marshal(t interface{}) ([]byte, error) {
	return internal.Marshal(t)
}
