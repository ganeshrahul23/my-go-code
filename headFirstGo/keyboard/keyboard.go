package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//GetFloat Function is used to Get float64 as input.
// Input to GetFloat function is a prompt you want to display.
func GetFloat(pr ...interface{}) (float64, error) {
	if len(pr) != 0 {
		prompt := fmt.Sprintf("%v", pr[0])
		if prompt != "" {
			fmt.Printf("%s\t", prompt)
		}
	}
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
