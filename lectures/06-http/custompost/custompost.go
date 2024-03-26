package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	ctx := context.Background()
	body := bytes.NewBufferString("All your base are belong to us")
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://myapi.com/create", body)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("X-Source", "Zero Wing")
	repr, err := httputil.DumpRequestOut(req, true)
	if err == nil {
		fmt.Println(string(repr))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
}
