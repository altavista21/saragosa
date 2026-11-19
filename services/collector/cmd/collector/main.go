package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/altavista21/saragosa/services/collector/internal/eve"
)

type summary struct {
	Processed int            `json:"processed"`
	Rejected  int            `json:"rejected"`
	ByType    map[string]int `json:"by_type"`
}

func main() {
	result := summary{ByType: make(map[string]int)}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 64*1024), 4*1024*1024)

	for scanner.Scan() {
		event, err := eve.ParseLine(scanner.Bytes())
		if err != nil {
			result.Rejected++
			fmt.Fprintf(os.Stderr, "collector: rejected malformed EVE line: %v\n", err)
			continue
		}
		result.Processed++
		result.ByType[event.EventType]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "collector: input read failed: %v\n", err)
		os.Exit(1)
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(result); err != nil {
		fmt.Fprintf(os.Stderr, "collector: write summary: %v\n", err)
		os.Exit(1)
	}
}
