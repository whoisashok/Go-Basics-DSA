package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

func fetchAPIs(ctx context.Context, urls []string) ([]string, error) {
	g, ctx := errgroup.WithContext(ctx)
	results := make([]string, len(urls))

	for i, url := range urls {
		i, url := i, url // Scope it right
		g.Go(func() error {
			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil {
				return err
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			results[i] = string(data)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}
	return results, nil
}

func ErrorGroup() {
	urls := []string{
		"https://api.example.com/user",
		"https://api.example.com/order",
		"https://api.example.com/404", // Boom!
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	results, err := fetchAPIs(ctx, urls)
	if err != nil {
		fmt.Println("Bailed:", err)
		return
	}
	for i, r := range results {
		fmt.Printf("Got %d: %s\n", i, r)
	}
}
