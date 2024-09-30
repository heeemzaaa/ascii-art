package fs

import (
	"fmt"
	"os"
	"strings"
)

func Read_file(s string) []string {
	file, err := os.ReadFile("art/" + s + ".txt")
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

func Middle(slice []string) bool {
	before := false
	after := false
	middle := true
	i := 0
	for ; i < len(slice); i++ {
		if slice[i] == "" {
			break
		}
	}
	for j := i - 1; j >= 0; j++ {
		before = false
		if slice[j] != "" {
			before = true
			break
		}
	}

	for k := i + 1; k < len(slice); k++ {
		before = false
		if slice[k] != "" {
			after = true
			break
		}
	}
	if before && after {
		middle = true
	} else {
		middle = false
	}
	return middle
}

func EmptyCase(slice []string) []string {
	empty := true
	for i := 0; i < len(slice); i++ {
		if slice[i] != "" {
			empty = false
			break
		}
	}
	if empty {
		slice = slice[1:]
	}
	return slice
}

func CleanSlice(slice []string) []string {
	check := true
	for i := 0; i < len(slice); i++ {
		if slice[i] != "" {
			check = false
		} else if slice[i] == "" {
			if Middle(slice) && !check {
				check = true
				slice = append(slice[:i], slice[i+1:]...)
			} else {
				i++
			}
		}
	}
	return slice
}

func PrintAscii(sl []string, file []string) string {
	result := ""
	holder := 0
	slice := EmptyCase(sl)
	for i := 0; i < len(slice); i++ {
		if slice[i] == "" {
			result += "\n"
		} else {
			for j := 0; j < 8; j++ {
				for k := 0; k < len(slice[i]); k++ {
					holder = (int(slice[i][k])-32)*9 + j
					result += file[holder]
				}
				result += "\n"
			}
		}
	}
	result = result[:len(result)-1]
	return result
}

func Is_ascii(s string) string {
	var result string
	slice := []rune(s)
	for i := 0; i < len(slice); i++ {
		if slice[i] == 10 || slice[i] == 13 {
			result += string(slice[i])
		} else if slice[i] >= 32 && slice[i] <= 126 {
			result += string(slice[i])
		}
	}
	return result
}

func FinalPrint(text string, banner string) string {
	name := ""
	if banner == "thinkertoy" || banner == "standard" || banner == "shadow" {
		name = banner
	} else {
		fmt.Println("incorrect banner")
		return ""
	}
	file := Read_file(name)
	if file == nil {
		return ""
	}
	line := text
	ret := Is_ascii(line)
	if len(line) < 1 {
		return ""
	}
	finalResult := ""
	splitted_line := strings.Split(ret, "\\n")
	cleaned := CleanSlice(splitted_line)
	finalResult = PrintAscii(cleaned, file[1:])
	return finalResult
}
