package main

func main() {

	/*
		rune is Go's solution for handling Unicode characters effectively in a world of global text and diverse languages.
	*/

	// rune string
	// s := "สวัสดีครับ"
	// fmt.Printf("%s, length: %d\n", s, len(s))

	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("%x ", s[i])
	// }

	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("i: %d, value %x ", i, s[i])
	// }

	// fmt.Println()

	// fmt.Printf("%s rune length: %d\n", s, utf8.RuneCountInString(s))

	// for i, w := 0, 0; i < len(s); i += w {
	// 	runeValue, width := utf8.DecodeRuneInString(s[i:])
	// 	fmt.Printf("%#U starts at %d\n", runeValue, i)
	// 	w = width
	// }

	/*-------- normal string --------- */

	// s := "hello"
	// fmt.Printf("%s, length: %d\n", s, len(s))

	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("%x ", s[i])
	// }

	// fmt.Println()

	// fmt.Printf("%s rune length: %d\n", s, utf8.RuneCountInString(s))

	// for i, w := 0, 0; i < len(s); i += w {
	// 	runeValue, width := utf8.DecodeRuneInString(s[i:])
	// 	fmt.Printf("%#U starts at %d\n", runeValue, i)
	// 	w = width
	// }

	/* -------- rune emoji -------- */

	// s := "golang is 👍"
	// runes := []rune(s)
	// for i, r := range runes {
	// 	if r == '👍' {
	// 		runes[i] = '👌'
	// 	}
	// }
	// fmt.Println(string(runes)) // Output: golang is 👌

	/* ------ use emoji with string ------*/
	// s := "golang is 👍"
	// runes := []string(s)
	// for i, r := range runes {
	// 	if r == '👍' {
	// 		runes[i] = '👌'
	// 	}
	// }
	// fmt.Println(string(runes)) // Output: golang is 👌

}
