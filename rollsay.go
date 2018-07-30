package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var quote string

	// prefer direct input via args over piped in
	if len(os.Args[1:]) > 0 {
		quote = strings.Join(os.Args[1:], " ")
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		quote = scanner.Text()
	}

	out := "\n" + roll(quote)

	fmt.Println(out)
}

func roll(quote string) string {

	sheets := (len(quote) / 6) + 1

	roll := strings.Split(`
	 .--""--.__.._
	(  (__)  )    '--.
	|'--..--'|      <|
	|       :|      /
	|       :|-""-./
	'.__  __;'      
	`, "\n")

	// sheets insert at index 11
	sheet := [5]string{"__..__", "     ⠃", "      ", "     ⠰", "--\"\"--"}

	for n := 0; n < sheets; n++ {
		for p := 0; p < 5; p++ {
			roll[p+1] = roll[p+1][:11] + sheet[p] + roll[p+1][11:]
		}
	}

	roll[3] = roll[3][:16] + quote + roll[3][16+len(quote):]

	return strings.Join(roll, "\n")

}
