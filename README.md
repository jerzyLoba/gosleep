## README.md

# `gosleep` - A Robust Concurrency Sleep Package

`gosleep` is a highly efficient, state-of-the-art Go package that addresses the complexity of concurrent sleep operations. With `gosleep`, we introduce a revolutionary approach to concurrent programming that utilizes Go’s goroutines and channels to execute sleep operations in parallel, providing unparalleled sleep management within your Go applications.

![I sleep](https://i.kym-cdn.com/photos/images/newsfeed/001/225/928/a11.jpg)

## Overview

The `SleepConcurrently` function is a key feature of the `gosleep` package. It allows for the concurrent execution of a sleep operation with the use of goroutines, channels, and `sync.WaitGroup` synchronization mechanisms. This advanced approach to sleeping will ensure that your Go application sleeps concurrently while maintaining complete control over execution flow.

## Why use gosleep?

- **Thread-Safe**: By leveraging Go’s goroutines, we ensure that your application’s sleeping operations are thread-safe and properly synchronized.
- **Performance-Centric**: Sleep while maintaining full concurrency without blocking other operations.
- **Minimalistic Design**: The function is simple, efficient, and purpose-built to sleep concurrently.

## Key Features

- **Concurrency-First Sleep**: Achieve true concurrency by sleeping inside a goroutine, with advanced synchronization provided by `sync.WaitGroup`.
- **Channel-Based Done Notification**: Utilize a channel to signal completion of the sleep operation. This modern approach minimizes overhead and enhances clarity in managing completion status.
- **Unparalleled Performance**: With Go’s powerful concurrency model, this function guarantees that your application will remain responsive, even while taking naps.

## Installation

To install `gosleep`, use the following command:

```sh
go get github.com/jerzyLoba/gosleep
```

## Usage

```go
import (
	"fmt"
	"github.com/jerzyLoba/gosleep"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	done := make(chan bool)

	// Initiating a concurrent sleep operation
	wg.Add(1)
	gosleep.SleepConcurrently(2*time.Second, done, &wg)

	// Wait for the sleep operation to complete
	wg.Wait()

	// Use the done channel to confirm completion
	select {
	case <-done:
		fmt.Println("Sleep completed successfully!")
	}
}
```
