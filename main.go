package main

import (
	"fmt"
	"intersection-go/intersection"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		result := intersection.FindIntersection([]string{
			"1,3,4,7,13",
			"1,2,4,4,13,15",
		})

		fmt.Fprintf(
			w,
			"hasil intersection antara [1,3,4,7,13] dan [1,2,4,13,15] adalah: %s\n",
			result,
		)
	})

	fmt.Println("server running on :3232")
	err := http.ListenAndServe(":3232", nil)
	if err != nil {
		panic(err)
	}
}
