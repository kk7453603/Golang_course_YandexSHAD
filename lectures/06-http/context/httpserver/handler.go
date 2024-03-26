package httpserver

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fd, _ := os.Open("core.c")
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		if err := ctx.Err(); err != nil {
			fmt.Println(ctx.Err())
			return
		}
		_, _ = w.Write(scanner.Bytes())
	}
}
