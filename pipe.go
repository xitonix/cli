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
		return false, nil
	}

	return stdin.Mode()&os.ModeCharDevice == 0 && stdin.Size() > 0, nil
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
// This is an alias for IsChained and PipeIn.
func PipeInIfChained() (string, error) {
	chained, err := IsChained();
	if err != nil {
		return "", err
	}

	if !chained {
		return "", nil
	}

	output, err := PipeIn()
	if err != nil {
		return "", err
	}
	return output, nil
}
