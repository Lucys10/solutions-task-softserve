package reader

import (
	"bufio"
	"os"
)

func ReadData() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}
