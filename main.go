package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go web Server!")
}
func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("서버가 9100포트에서 시작되었습니다.")

	err := http.ListenAndServe(":9100", nil)
	if err != nil {
		fmt.Println("서버 실행 실패: ", err)
	}
}
