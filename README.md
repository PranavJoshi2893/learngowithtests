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

---

## Arrays — Sum and SumAll

A program that calculates the sum of integers in a slice and multiple slices.

### Features

* `Sum` takes a slice of integers and returns their sum
* `SumAll` takes multiple slices and returns a slice of sums
* Testing using `TestXxx` with `reflect.DeepEqual` for slice comparison

### Example Usage

```go
Sum([]int{1, 2, 3, 4, 5})        // 15
Sum([]int{1, 2, 3})              // 6
SumAll([]int{1,2,3,4,5}, []int{1,2,3}) // []int{15, 6}
```

### Running Tests

```bash
go test
```
---

## Structs, methods & interfaces -> Shapes — Perimeter and Area

A program that calculates the **perimeter** and **area** of basic geometric shapes using **interfaces** and **structs** in Go.

### Features

* Demonstrates **interfaces** to define common behavior (`Area()`)
* Implements **structs** for different shapes: `Rectangle`, `Circle`, and `Triangle`
* Provides specific methods for:

  * `Perimeter()` — for rectangles
  * `Area()` — for all shapes
* Uses **table-driven tests** to test multiple shapes efficiently
* Employs **subtests** (`t.Run`) for better test readability and organization

### Example Usage

```go
rect := Rectangle{Width: 10, Height: 10}
rect.Perimeter() // 40
rect.Area()      // 100

circle := Circle{Radius: 10}
circle.Area() // 314.1592653589793

triangle := Triangle{Base: 12, Height: 6}
triangle.Area() // 36
```

### Running Tests

```bash
go test
```

### Example Test Output

```bash
=== RUN   TestPerimeter
--- PASS: TestPerimeter (0.00s)
=== RUN   TestArea
=== RUN   TestArea/Rectangle
=== RUN   TestArea/Circle
=== RUN   TestArea/Triangle
--- PASS: TestArea (0.00s)
    --- PASS: TestArea/Rectangle (0.00s)
    --- PASS: TestArea/Circle (0.00s)
    --- PASS: TestArea/Triangle (0.00s)
PASS
ok      shapes  0.001s
```
