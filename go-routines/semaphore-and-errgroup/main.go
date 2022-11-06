package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

func CorrectVersion() error {
	g, ctx := errgroup.WithContext(context.Background())
	const (
		maxWorkers = 5
	)

	g.Go(func() error {
		sem := semaphore.NewWeighted(maxWorkers)

		for i := 0; i < 50; i++ {
			k := i
			if err := sem.Acquire(ctx, 1); err != nil {
				return err
			}

			g.Go(func() error {
				defer sem.Release(1)
				time.Sleep(1 * time.Second)
				fmt.Println(k)

				if k == 5 {
					return errors.New("[yu")
				}
				return nil
			})
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

func IncorrectVersion() error {
	g, ctx := errgroup.WithContext(context.Background())
	const (
		maxWorkers = 5
	)

	sem := semaphore.NewWeighted(maxWorkers)

	for i := 0; i < 20; i++ {
		k := i
		if err := sem.Acquire(ctx, 1); err != nil {
			return err
		}

		g.Go(func() error {
			defer sem.Release(1)
			time.Sleep(1 * time.Second)
			fmt.Println(k)

			if k == 5 {
				return errors.New("[yu")
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}
}

func main() {
	CorrectVersion()
	// IncorrectVersion()
}
