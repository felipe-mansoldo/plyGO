---
sidebar_position: 1
---

# Getting Started

Install plyGO and create your first pipeline.

```bash
go get github.com/felipe-mansoldo/plyGO
```

## Your First Program

```go
package main

import "github.com/felipe-mansoldo/plyGO"

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
    
    plygo.From(people).Show()
}
```

:::success Result
```
+---------+-----+----------+
|    Name | Age |   Salary |
+---------+-----+----------+
| Alice   |  30 | 75000.00 |
| Bob     |  25 | 60000.00 |
| Charlie |  35 | 90000.00 |
+---------+-----+----------+
[3 rows × 3 columns]
```
:::

See the [Tutorial Basics](02-data-loading.md) to learn more!
