package printer

import (
	"fmt"

	"github.com/goodelias/otusgo-basic/hw02_fix_app/types"
)

func PrintStaff(staff []types.Employee) {
	for _, v := range staff {
		fmt.Println(v.String())
	}
}
