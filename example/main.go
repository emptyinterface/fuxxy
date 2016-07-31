package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/emptyinterface/fuxxy"
)

func main() {

	file, err := os.Open("/usr/share/dict/words")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var words []string
	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		words = append(words, sc.Text())
	}

	http.Handle("/", http.FileServer(http.Dir("./web")))

	http.HandleFunc("/autocomplete", func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		for _, word := range fuxxy.Matches(words, req.FormValue("input")) {
			fmt.Fprintln(w, word)
		}
		fmt.Println("rendered in", time.Since(start))
	})

	fmt.Println(len(words), "words ready")

	log.Fatal(http.ListenAndServe(":8080", nil))

}
