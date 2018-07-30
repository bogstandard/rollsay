package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {

	// establish random number of sheets
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(8-3) + 3

	out := "\n            " + quote() + "\n" + roll(n)

	fmt.Println(out)
}

func quote() string {
	return "Life rolls on and on.. "
}

func roll(sheets int) string {

	roll := strings.Split(`
	 .--""--.__.._
	(  (__)  )    '--.
	|'--..--'|      <|
	|       :|      /
	|       :|-""-./
	'.__  __;'      
	`, "\n")

	// sheets insert at index 11

	sheet := [5]string{"__..__", "     :", "     :", "     :", "--\"\"--"}

	for n := 0; n < sheets; n++ {
		for p := 0; p < 5; p++ {
			roll[p+1] = roll[p+1][:11] + sheet[p] + roll[p+1][11:]
		}
	}

	return strings.Join(roll, "\n")

}
