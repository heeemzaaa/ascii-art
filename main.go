package main

import (
	"fmt"
	"os"
	"strings"
)

func read_file(s string) []string {
	file, err := os.ReadFile(s + ".txt")
	if err != nil {
		fmt.Println("Ascii file not found")
		return nil
	}
	ret := strings.Split(string(file), "\n")
	for i := 0; i < len(ret) && s == "thinkertoy"; i++ {
		ret[i] = strings.ReplaceAll(ret[i], "\r", "")
	}
	return ret
}

func count_next_line(line string) []int {
	var ret []int
	j := 0
	ret = append(ret, 0)
	for i := 0; i < len(line); i++ {
		if i+1 < len(line) && line[i] == '\\' {
			if line[i+1] == 'n' {
				ret[j]++
			}
			i++
			if i+1 < len(line) && line[i+1] != '\\' {
				ret = append(ret, 0)
				j++
			}
		}
	}
	return ret
}

func print_art(file []string, splitted_line []string, lines_count []int) {
	holder := 0
	i := 0
	for ; i < len(splitted_line); i++ {
		for j := 0; j < 8; j++ {
			for k := 0; k < len(splitted_line[i]); k++ {
				holder = (int(splitted_line[i][k])-32)*9 + j
				fmt.Printf("%s", file[holder])
			}
			fmt.Println()
		}
		for ; len(lines_count) > 0 && lines_count[i] > 1; lines_count[i]-- {
			fmt.Println()
		}
	}
	i--
	if i >= 0 && i < len(lines_count) {
		for ; lines_count[i] > 0; lines_count[i]-- {
			fmt.Println()
		}
	}
}

func check_if_empty(s []string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != "" {
			return true
		}
	}
	return false
}

func cleaned_split(s []string, lines_count []int) ([]string, []int) {
	var ret []string
	i := 0
	if s[0] == "" {
		if !check_if_empty(s) {
			i++
		}
		for ; i < len(s) && s[i] == ""; i++ {
			fmt.Println()
		}
		if len(lines_count) > 1 {
			lines_count = lines_count[1:]
		}
	}
	for ; i < len(s); i++ {
		if s[i] != "" {
			ret = append(ret, s[i])
		}
	}
	return ret, lines_count
}

func is_ascii(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < 32 || s[i] > 126 {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Argument Error")
		return
	}
	file := read_file(os.Args[2])
	if file == nil {
		return
	}
	line := os.Args[1]
	if !is_ascii(line) {
		fmt.Println("Non Ascii character found")
		return
	}
	if len(line) < 1 {
		return
	}
	lines_count := count_next_line(line)
	splitted_line := strings.Split(line, "\\n")
	splitted_line, lines_count = cleaned_split(splitted_line, lines_count)
	print_art(file[1:], splitted_line, lines_count)
}
