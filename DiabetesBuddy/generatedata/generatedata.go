package generatedata

import (
	"math/rand"
	"time"
)

func randomGen() {

	var chanint chan int
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate 100 random floats between 0 and 20
	for i := 0; i < 100; i++ {
		x := rand.Float64() * 20
		go func() {
			chanint <- int(x)
		}()
	}

}
