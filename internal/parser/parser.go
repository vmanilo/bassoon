package parser

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"bassoon/internal/app/model"
)

func ParseFile(ctx context.Context, filename string) (<-chan *model.Port, error) {
	file, err := os.Open(filepath.Clean(filename))
	if err != nil {
		return nil, err
	}

	const chanSize = 10
	pipe := make(chan *model.Port, chanSize)

	go func() {
		defer func() {
			if err := file.Close(); err != nil {
				log.Printf("parser: failed to close file: %v\n", err)
			}

			close(pipe)
		}()

		scanner := bufio.NewScanner(file)
		scanner.Split(scanPorts)

		for scanner.Scan() {
			select {
			case <-ctx.Done():
				log.Println("parser: ctx.Done - interrupt parsing file")

				return
			default:
			}

			str := scanner.Text()

			port := parsePort(str)
			if port != nil {
				pipe <- port
			}

			// just to see progress
			fmt.Print(".") //nolint:forbidigo
		}
	}()

	return pipe, nil
}

func parsePort(str string) *model.Port {
	i := strings.Index(str, ":")
	portID := strings.Trim(str[:i], `"`)

	var port model.Port
	if err := json.Unmarshal([]byte(str[i+1:]), &port); err != nil {
		// skip this data
		return nil
	}

	port.ID = portID

	return &port
}

func scanPorts(data []byte, atEOF bool) (int, []byte, error) {
	const minObjLength = 10

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	start := bytes.IndexByte(data, '"')
	end := bytes.IndexByte(data, '}')

	if start >= 0 && end >= 0 && end-start > minObjLength {
		return end + 1, data[start : end+1], nil
	}

	// Request more data.
	return 0, nil, nil
}
