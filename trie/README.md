# gotrie

`gotrie` is a generic Trie (prefix tree) implementation in Go, which supports storing key-value pairs with optional expiry durations.

## Features

- Generic Type Support: Store values of any type in the Trie.
- Optional Expiry: Insert key-value pairs with an optional expiry duration, allowing automatic invalidation of outdated entries.
- Efficient Operations: Fast insert, find, and remove operations.
- Zero-value Handling: Properly handles Go zero values, ensuring accurate and predictable behavior.

## Installation

To install gotrie, use the following command:

```sh
go get github.com/binaek/gotrie
```

## Usage

Below is an example of how to use gotrie in your Go project:

```go
package main

import (
    "fmt"
    "time"

    "github.com/binaek/gotrie"

)

func main() {
    // Create a new Trie
    trie := gotrie.NewTree[string]()

    // Insert key-value pairs
    oldValue, replaced := trie.Insert("hello", "world")
    fmt.Println("Old Value:", oldValue, "Replaced:", replaced)

    // Insert key-value pairs with expiry
    expiryDuration := 5 * time.Second
    oldValue, replaced = trie.InsertWithExpiry("temporary", "data", expiryDuration)
    fmt.Println("Old Value:", oldValue, "Replaced:", replaced)

    // Retrieve values
    value, found := trie.Find("hello")
    fmt.Println("Value:", value, "Found:", found)

    // Wait for expiry
    time.Sleep(6 * time.Second)
    value, found = trie.Find("temporary")
    fmt.Println("Value:", value, "Found:", found) // Should indicate not found

    // Remove values
    oldValue, removed := trie.Remove("hello")
    fmt.Println("Old Value:", oldValue, "Removed:", removed)
}
```

## API

### `func NewTree[T any]() Tree[T]`

Creates and returns a new Trie instance.

### `func (t *Tree[T]) Insert(key string, value T) (oldValue T, replaced bool)`

Inserts a key-value pair into the Trie. Returns the old value (if any) and a boolean indicating if a value was replaced.

### `func (t *Tree[T]) InsertWithExpiry(key string, value T, expiry time.Duration) (oldValue T, replaced bool)`

Inserts a key-value pair into the Trie with an expiry duration. Returns the old value (if any) and a boolean indicating if a value was replaced.

### `func (t *Tree[T]) Find(key string) (value T, found bool)`

Retrieves the value associated with the given key. Returns the value and a boolean indicating if the key was found.

### `func (t *Tree[T]) Remove(key string) (oldValue T, removed bool)`

Deletes the key-value pair from the Trie. Returns the old value (if any) and a boolean indicating if a value was removed.

## Advantages

### Type Safety

Leverages Go's generics to provide type-safe operations.

### Memory Efficient

Uses a compact node structure to minimize memory usage.

### Expiry Handling

Built-in support for expiring entries makes it suitable for caching and time-sensitive data storage.

### Ease of Use

Simple and intuitive API, with operations mirroring common map operations.

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT). See the [LICENSE](./LICENSE) file for details.

Feel free to contribute to this project by opening issues or submitting pull requests on GitHub.
