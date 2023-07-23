package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:3005", nil)
	if err != nil {
		fmt.Println(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
