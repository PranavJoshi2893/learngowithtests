# Learn Go with Tests — Hello World

A simple Go program that greets a user in different languages, written using **test-driven development (TDD)**.

### Features

* Takes `name` and `language` as inputs.
* Defaults to `"World"` if `name` is empty.
* Defaults to English if `language` is empty.
* Supports English, Spanish, and French.

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

# Learn Go with Tests — Integers

A simple Go program that adds two integers and returns an int, written using **test-driven development (TDD)**.

### Features

* Testing using `TestXxx`
* Testing using `Example`
* `Sum` takes 2 ints and returns their addition

### Example Usage

```go
Sum(2, 4) // 6
```

### Running Tests

```bash
go test
```
