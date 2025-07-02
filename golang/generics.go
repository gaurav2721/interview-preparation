Working with **generics in Go (since Go 1.18)** allows you to write reusable and type-safe code that works with **any type** — without sacrificing performance or writing repetitive boilerplate.

---

## ✅ What are Generics?

Generics let you write **functions**, **types**, and **methods** that operate on a **type parameter** — instead of a specific type like `int`, `string`, or `float64`.

---

## 🧱 Basic Syntax

### 1. **Generic Function**

```go
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}
```

Usage:

```go
PrintSlice([]int{1, 2, 3})
PrintSlice([]string{"a", "b", "c"})
```

* `T` is a **type parameter**.
* `[T any]` means `T` can be **any type**.

---

### 2. **Generic Struct**

```go
type Box[T any] struct {
	Value T
}

func (b *Box[T]) Get() T {
	return b.Value
}
```

Usage:

```go
intBox := Box[int]{Value: 10}
strBox := Box[string]{Value: "hello"}

fmt.Println(intBox.Get()) // 10
fmt.Println(strBox.Get()) // hello
```

---

### 3. **Generic Interface (with constraints)**

```go
type Adder[T any] interface {
	Add(a, b T) T
}
```

You can create more specific constraints (e.g., only numeric types) using **type sets** (see below).

---

### 4. **Type Constraints**

You can restrict what types are allowed using constraints.

```go
type Number interface {
	~int | ~float64
}

func Sum[T Number](a, b T) T {
	return a + b
}
```

* `~int` allows `int` and types **derived from int**.
* `|` is a union — "either or".

---

### 5. **Generic Method on a Generic Struct**

```go
type Pair[T any] struct {
	A, B T
}

func (p *Pair[T]) Swap() {
	p.A, p.B = p.B, p.A
}
```

---

## 📌 Built-in Constraints (from `constraints` package)

Go provides standard type constraints:

```go
import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
```

* `constraints.Ordered` includes `int`, `float`, and `string`.

---

## 🛠 When to Use Generics

| Use Case                        | Benefit                        |
| ------------------------------- | ------------------------------ |
| Collections (e.g. `Stack[T]`)   | Type-safe reusable structures  |
| Algorithms (e.g. `Min`, `Sort`) | No duplication for each type   |
| Middleware, Services, Caches    | Generic interfaces and structs |

---

## ❗ Limitations

* You **can’t reflect on** the type parameter at runtime.
* No operator overloading, so constraints are necessary for `+`, `<`, etc.
* Type inference may not always work — you might need to specify `[int]` explicitly.

---

## ✅ Summary

| Feature          | Example                              |               |
| ---------------- | ------------------------------------ | ------------- |
| Generic function | `func Foo[T any](x T) {}`            |               |
| Generic struct   | `type Box[T any] struct { Value T }` |               |
| Type constraints | `func Sum[T Number](a, b T) T`       |               |
| Union types      | \`type Number interface { \~int      | \~float64 }\` |
| Standard lib     | `constraints.Ordered`                |               |

---

Let me know if you want:

* A generic stack/queue
* A generic repository pattern
* Generics in interfaces or methods on interfaces

I can show working examples!
