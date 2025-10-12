---
layout: home

hero:
  name: "plyGO"
  text: "Data Manipulation in Go, Simplified"
  tagline: "dplyr-inspired data manipulation with full type safety and zero dependencies"
  image:
    src: /img/img_welcome.png
    alt: plyGO
  actions:
    - theme: brand
      text: Get Started
      link: /guide/getting-started
    - theme: alt
      text: View on GitHub
      link: https://github.com/felipe-mansoldo/plyGO

features:
  - icon: 🚀
    title: Zero Dependencies
    details: Pure Go implementation with no external dependencies. Just add to your project and start using.

  - icon: 🔒
    title: Type Safe
    details: Leverages Go generics for full type safety. Catch errors at compile time, not runtime.

  - icon: 🔄
    title: Fluent API
    details: Chainable methods inspired by R's dplyr. Write readable, maintainable data pipelines.

  - icon: ⚡
    title: High Performance
    details: Optimized for speed and memory efficiency. Handle large datasets with confidence.

  - icon: 📊
    title: Data Focused
    details: Built specifically for data manipulation. Filter, select, group, and transform with ease.

  - icon: 🛠️
    title: Extensible
    details: Create custom helpers and compose operations. Adapt to your specific needs.
---

<style>
.vp-doc h2 {
  margin-top: 48px;
  border-top: 1px solid var(--vp-c-divider);
  padding-top: 24px;
}
</style>

## Quick Start

Install plyGO with a single command:

```bash
go get github.com/mansoldof/plyGO
```

Create your first data pipeline in seconds:

```go
package main

import "github.com/mansoldof/plyGO"

type Person struct {
    Name   string
    Age    int
    Salary float64
}

func main() {
    people := []Person{
        {"Alice", 30, 75000},
        {"Bob", 25, 60000},
        {"Charlie", 35, 90000},
    }

    plygo.From(people).
        Where("Age").GreaterThan(30).
        OrderBy("Salary").Desc().
        Show()
}
```

## Why plyGO?

### Elegant Data Manipulation

plyGO brings the elegance of R's dplyr to Go, making data manipulation intuitive and enjoyable. Chain operations together to build complex data pipelines with minimal code.

```go
// Filter, select, and transform in one pipeline
result := plygo.From(employees).
    Where("Department").Equals("Engineering").
    Select("Name", "Salary").
    OrderBy("Salary").Desc().
    Collect()
```

### Type Safe Operations

Unlike reflection-heavy alternatives, plyGO uses Go generics to provide compile-time type safety while maintaining flexibility.

```go
// Type-safe operations with struct inference
filtered := plygo.From(people).
    Where("Age").GreaterThan(25).
    Where("Salary").LessThan(80000).
    Collect()
```

### Built for Go Developers

Designed with Go's philosophy in mind: simple, efficient, and reliable. No magic, no surprises, just clean code that works.

## Platform Support

- **Pure Go** - Works on all platforms Go supports
- Go 1.18+ required (for generics support)
- Zero external dependencies

## What's Next?

<div style="display: grid; grid-template-columns: repeat(auto-fit, minmax(200px, 1fr)); gap: 16px; margin: 32px 0;">
  <a href="/plyGO/guide/getting-started" style="text-decoration: none;">
    <div style="padding: 20px; border: 1px solid var(--vp-c-divider); border-radius: 8px; text-align: center;">
      <div style="font-size: 32px; margin-bottom: 8px;">📚</div>
      <strong>Read the Guide</strong>
      <p style="margin: 8px 0 0; font-size: 14px; color: var(--vp-c-text-2);">Learn the fundamentals</p>
    </div>
  </a>

  <a href="/plyGO/basics/data-loading" style="text-decoration: none;">
    <div style="padding: 20px; border: 1px solid var(--vp-c-divider); border-radius: 8px; text-align: center;">
      <div style="font-size: 32px; margin-bottom: 8px;">💡</div>
      <strong>Basic Operations</strong>
      <p style="margin: 8px 0 0; font-size: 14px; color: var(--vp-c-text-2);">Master the basics</p>
    </div>
  </a>

  <a href="/plyGO/extras/real-world-examples" style="text-decoration: none;">
    <div style="padding: 20px; border: 1px solid var(--vp-c-divider); border-radius: 8px; text-align: center;">
      <div style="font-size: 32px; margin-bottom: 8px;">📖</div>
      <strong>See Examples</strong>
      <p style="margin: 8px 0 0; font-size: 14px; color: var(--vp-c-text-2);">Real-world use cases</p>
    </div>
  </a>
</div>

## License

MIT License 
