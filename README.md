Got it! Here’s a concise, clean version of your README:

---

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
Hello("Jean", "")          // "Hello, Jean"
```

### Running Tests

```bash
go test
```

---

Do you want me to also make a **one-line description for the repo header**?
