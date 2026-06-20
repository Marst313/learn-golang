package main

import (
	"fmt"
	"strings"
)

/* import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
) */

func main() {

	/* str := "Hello Go!"
	fmt.Println(str)

	str1 := "Hello"
	str2 := "World"

	result := str1 + " " + str2
	fmt.Println(result)

	fmt.Println(str[0])
	fmt.Println(str[0:6])

	// String Conversion
	num := 18
	str3 := strconv.Itoa(num) // conver number to string

	fmt.Println(len(str3))
	// fmt.Println(len(num)) // ! cant do this because number cant use len() methods

	// string splitting
	fruit := "apple, orange, banana"
	fruit2 := "apple-orange-banana"
	parts := strings.Split(fruit, ",")
	parts2 := strings.Split(fruit2, "-")

	fmt.Println(parts)
	fmt.Println(parts2)

	countries := []string{"Germany", "France", "Italy"}
	joined := strings.Join(countries, ", ")
	fmt.Println(joined)

	fmt.Println(strings.Contains(str, "Go?"))

	replaced := strings.Replace(str, "Go", "Universe", 1)
	fmt.Println(replaced)

	strwspace := " Hello Everyone ! "
	fmt.Println(strings.TrimSpace(strwspace))

	fmt.Println(strings.ToLower(strwspace))
	fmt.Println(strings.ToUpper(strwspace))
	fmt.Println(strings.Repeat("foo ", 3))
	fmt.Println(strings.Count("Hello", "H"))
	fmt.Println(strings.HasPrefix("Hello", "lo"))
	fmt.Println(strings.HasSuffix("Hello", "lo"))

	str5 := "Hel1lo, 123 Go 11!"
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str5, -1)
	fmt.Println(matches)

	str6 := "Hello, 世界"
	fmt.Println(utf8.RuneCountInString(str6)) */

	// ? STRING BUILDER
	var builder strings.Builder

	// ? Write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World!")

	// ? Convert builder to a string
	result := builder.String()
	fmt.Println(result)

	// ? Using writerune to add a character
	builder.WriteRune(' ')
	builder.WriteString("How are you ?")

	result = builder.String()
	fmt.Println(result)

	// ? Reset the builder
	builder.Reset()
	builder.WriteString("Strating fresh!")
	result = builder.String()
	fmt.Println(result)

}
