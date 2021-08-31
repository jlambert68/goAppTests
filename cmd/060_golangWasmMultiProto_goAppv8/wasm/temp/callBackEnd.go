package temp

import (
	"time"
)

func callBackEnd() string {

	t := time.Now()
	return t.String()

}
