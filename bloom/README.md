# bloom

`bloom` is a Go library that provides a Bloom Filter implementation using a Trie (prefix tree) under the hood for efficient probabilistic membership testing.

## Features

- Probabilistic Membership Testing: Efficiently checks if an element is possibly in a set.
- Customizable Hash Functions: Supports the use of different hash functions for better distribution and collision handling (`fnv.New64()` and `fnv.New64a()` by default).
- Thread-Safe: Safe for concurrent use with internal locking mechanisms.

## Installation

To install gotrie, use the following command:

```sh
go get github.com/binaek/gocoll
```

## Advantages

### Efficient

Uses a Trie structure for storing hashed values, ensuring efficient membership testing.

### Customizable

Supports various hash functions for better collision handling and distribution.

### Thread-Safe

Designed to be safe for concurrent use, making it suitable for multi-threaded applications.

## Example:

```go
package main

import (
    "crypto/sha256"
    "crypto/sha512"
    "fmt"
    "hash"

    "github.com/binaek/gocoll/bloom"
)

func main() {
    // Create a new Bloom Filter with custom hashers
    // By default bloom uses fnv.New64() and fnv.New64a()
    hashers := []hash.Hash{
        sha256.New(),
        sha512.New(),
    }
    bf := bloom.NewBloomFilter(bloom.WithHashers(hashers))

    // Add tokens to the Bloom Filter
    bf.Add("apple")
    bf.Add("banana")

    // Test for membership
    fmt.Println("Contains 'apple':", bf.Test("apple"))   // Should be true
    fmt.Println("Contains 'banana':", bf.Test("banana")) // Should be true
    fmt.Println("Contains 'grape':", bf.Test("grape"))   // Should be false
}
```
