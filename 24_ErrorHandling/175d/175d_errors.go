package main

import (
	"fmt"
	"log"
)

type matherror struct {
	lat  string
	long string
	err  error
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func (n matherror) Error() string {
	return fmt.Sprintf("a math issue occured: %v %v %v", n.lat, n.long, n.err)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		me := fmt.Errorf("Math issue, sqrt of negative problem not possible %v", f)
		return 0, matherror{"50.2222 N", "99.99 W", me}
	}
	return 42, nil
}
