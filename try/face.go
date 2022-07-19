package try

import (
	"fmt"
)

func Face(name string) string {
	message := fmt.Sprintf("        ****        \n      ********      \n      ** ** **      \n     **  **  **     \n     **********     \n     **%v**     \n     **********     \n       ******       \n", name)
	return message
}
