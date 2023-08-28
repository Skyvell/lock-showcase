package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/skyvell/locksv2/s3lock"
)

const bucketName = "versioningbucketcrossbreed"
const key = "key-201"

func main() {
	ctx := context.Background()
	config, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic("Could not load config.")
	}

	// Construct 10 lock instances that will compete for the lock.
	competingLocks := []*s3lock.S3Lock{}
	for i := 0; i < 10; i++ {
		competingLocks = append(competingLocks, s3lock.NewS3Lock(config, fmt.Sprintf("Lock%v", i), bucketName, key, 50*time.Second))
	}

	// Use WaitGroup to wait for all Goroutines to complete.
	var wg sync.WaitGroup
	for _, lock := range competingLocks {
		wg.Add(1)
		go func(l *s3lock.S3Lock) {
			// Decrement counter when Goroutine completes.
			defer wg.Done()

			err := l.AcquireLockWithRetry(ctx, time.Second*60)
			if err == nil {
				// Execute code here when lock is acquired.
				fmt.Printf("Lock: %s aquired lock. Executing:\n", l.LockName)
				fmt.Printf("Sleeping for 1 second. \n\n")
				time.Sleep(time.Second * 1)
				l.ReleaseLock(ctx)
			}
		}(lock)
	}
	wg.Wait()
}
