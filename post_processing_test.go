package PostProcessor

import (
	"fmt"
	"testing"
	"time"
)

func TestPostProcessing_Prepare(t *testing.T) {
	start := time.Now()
	for range 4_000 {
		_, err := Calculate("=SUM[~0,RAND[0,500]]", int64(10_000))
		if err != nil {
			t.Error(err)
		}
	}
	fmt.Println(time.Since(start))
}
