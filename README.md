# 🚀 Learning Go (Golang)

Welcome! This repository is a starting point for learning the Go programming language. Go (or Golang) is a statically typed, compiled language designed for simplicity, performance, and reliability.

---

## 📦 Getting Started

Make sure you have Go installed:

```bash
go version
```

If not, download it from: https://go.dev/dl/

---

## 🛠️ Initialize a Project

Create a new Go module:

```bash
go mod init your-module-name
```

This will generate a `go.mod` file to manage dependencies.

---

## ▶️ Run Your Program

To run your Go application:

```bash
go run .
```

This compiles and runs all `.go` files in the current directory.

---

## 🧹 Manage Dependencies

Clean up and sync your dependencies:

```bash
go mod tidy
```

This removes unused dependencies and adds missing ones.

---

## 🧪 Run Tests

Execute all tests in your project:

```bash
go test ./...
```

Or just in the current directory:

```bash
go test
```

---

## 🏗️ Build Your Program

Compile your Go code into a binary:

```bash
go build
```

For a detailed guide on compiling and installing Go programs, check out:
👉 https://go.dev/doc/tutorial/compile-install

---

## 📖 Explore Go Commands

Get help for Go commands:

```bash
go help
```

You can also explore specific commands:

```bash
go help build
go help test
```

---

## 📚 Recommended Learning Path

1. Learn basic syntax (variables, functions, control flow)
2. Understand packages and modules
3. Work with structs and interfaces
4. Explore concurrency (goroutines & channels)
5. Build small CLI or API projects

---

## 🧠 Tips

* Keep your code simple and readable
* Use `gofmt` to format your code:

  ```bash
  go fmt ./...
  ```
* Follow Go conventions and idioms

---

## 📁 Example Project Structure

```
.
├── go.mod
├── main.go
└── README.md
```

---

## 🙌 Resources

* Official Docs: https://go.dev/doc/
* Tour of Go: https://go.dev/tour/
* Effective Go: https://go.dev/doc/effective_go

---

Happy coding! 💡
