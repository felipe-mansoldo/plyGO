# Welcome to plyGO

**Data Manipulation in Go, Simplified**

plyGO brings elegant, R-inspired data manipulation to Go with full type safety and zero dependencies.

## Quick Example

```go
plygo.From(people).
    Where("Age").GreaterThan(30).
    OrderBy("Salary").Desc().
    Show()
```

See [Getting Started](./getting-started.md) to begin!
