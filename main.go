package main

import (
	"fmt"
	"intersection-go/intersection"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		result := intersection.FindIntersection(
			[]string{"1,3,4,7,13", "1,2,4,13,15"},
		)
		fmt.Fprintf(w, "Hasil intersection: %s\n", result)
	})

	fmt.Println("Server running on :3232")
	err := http.ListenAndServe(":3232", nil)
	if err != nil {
		fmt.Println(err)
	}
}
