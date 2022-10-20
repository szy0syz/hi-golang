package main

import "fmt"

func main() {
	str := "Go is a beautiful language"
	fmt.Printf("The length of str is: %d\n", len(str))

	for pos, char := range str {
		fmt.Printf("Char on position %d is: %c \n", pos, char)
	}
	fmt.Println()
	str2 := "Chinese: 汉语"
	fmt.Printf("The length of str2 is %d\n", len(str2))
	for pos, char := range str2 {
		fmt.Printf("Char %c starts at byte position %d\n", char, pos)
	}
	fmt.Println()
	fmt.Println("index int(rune) rune    char bytes")
	for index, rune := range str2 {
		fmt.Printf("%-2d      %d      %U '%c' % X\n", index, rune, rune, rune, []byte(string(rune)))
	}
}

/*
index int(rune) rune    char bytes
0       67      U+0043 'C' 43
1       104      U+0068 'h' 68
2       105      U+0069 'i' 69
3       110      U+006E 'n' 6E
4       101      U+0065 'e' 65
5       115      U+0073 's' 73
6       101      U+0065 'e' 65
7       58      U+003A ':' 3A
8       32      U+0020 ' ' 20
9       27721      U+6C49 '汉' E6 B1 89
12      35821      U+8BED '语' E8 AF AD
*/
