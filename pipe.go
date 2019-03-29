package cli

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// IsChained returns true if the cli tool is the target of a command chain.
func IsChained() (bool, error) {
	stdin, err := os.Stdin.Stat()
	if err != nil {
		return false, err
	}

	return stdin.Mode()&os.ModeCharDevice == 0, nil
}

// PipeIn reads the output of the command chain
func PipeIn() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	var builder strings.Builder
	for {
		input, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		_, err = builder.WriteRune(input)
		if err != nil {
			return "", err
		}
	}

	return builder.String(), nil
}

// PipeInIfChained reads the output of the command chain if the cli tool is piped in.
//
// A boolean value will be returned indicating whether or not the current binary was piped in.
// This methods an alias for IsChained and PipeIn methods combined.
func PipeInIfChained() (string, bool, error) {
	chained, err := IsChained()
	if err != nil {
		return "", false, err
	}

	if !chained {
		return "", chained, nil
	}

	output, err := PipeIn()
	if err != nil {
		return "", chained, err
	}
	return output, chained, nil
}
