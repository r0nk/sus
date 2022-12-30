package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	//flag.BoolVar(&quietMode, "q", false, "quiet mode (no output at all)")
	//flag.Parse()

	//fn := flag.Arg(0)

	line_counts := make(map[string]int)
	var lines []string

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		line := sc.Text()
		if line_counts[line] == 0 {
			lines = append(lines, line)
		}
		line_counts[line] += 1
	}

	sort.Slice(lines, func(i, j int) bool {
		return line_counts[lines[i]] > line_counts[lines[j]]
	})

	for _, line := range lines {
		fmt.Printf("%d %s\n", line_counts[line], line)
	}
}
