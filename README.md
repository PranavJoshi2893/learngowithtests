# Learn Go with Tests

A series of simple Go programs written using **test-driven development (TDD)** to demonstrate basic programming concepts.

---

## Hello World

A program that greets a user in different languages.

### Features

* Takes `name` and `language` as inputs
* Defaults to `"World"` if `name` is empty
* Defaults to English if `language` is empty
* Supports English, Spanish, and French

### Example Usage

```go
Hello("Chris", "English") // "Hello, Chris"
Hello("", "Spanish")      // "Hola, World"
Hello("Jean", "")         // "Hello, Jean"
```

### Running Tests

```bash
go test
```

---

## Integers

A program that adds two integers.

### Features

* Testing using `TestXxx`
* Testing using `Example`
* `Sum` takes 2 ints and returns their sum

### Example Usage

```go
Sum(2, 4) // 6
```

### Running Tests

```bash
go test
```

---

## Iteration

A program that repeats a character 5 times.

### Features

* Testing using `TestXxx`
* Benchmark using `BenchmarkXxx` to check performance and memory usage
* `Repeat` takes a character and returns a string
* Tests include simple loops and the `strings` package to avoid unnecessary memory usage, since strings are immutable

### Example Usage

```go
Repeat("a") // "aaaaa"
```

### Running Tests

```bash
go test -bench=. -benchmem
```
