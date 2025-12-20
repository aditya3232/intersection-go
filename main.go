package main

import (
	"fmt"
	"intersection-go/intersection"
)

func main() {
	result := intersection.FindIntersection([]string{"1,3,4,7,13", "1,2,4,13,15"})
	fmt.Println("hasil intersection antara [\"1,3,4,7,13\", \"1,2,4,13,15\"] adalah: ", result)
}
