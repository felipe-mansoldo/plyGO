---
sidebar_position: 8
---

# Transformation

Learn how to transform data using the `Transform` function.

## Basic Transformation

Apply a function to each item in the pipeline:

```go
type Product struct {
    Name  string
    Price float64
}

products := []Product{
    {"Laptop", 1000},
    {"Mouse", 50},
    {"Keyboard", 80},
}

plygo.From(products).
    Transform(func(p Product) Product {
        p.Price = p.Price * 0.9  // 10% discount
        return p
    }).
    Show()
```

:::success Result
```
+----------+--------+
|     Name |  Price |
+----------+--------+
| Laptop   | 900.00 |
| Mouse    |  45.00 |
| Keyboard |  72.00 |
+----------+--------+
[3 rows × 2 columns]
```
:::

## Conditional Transformation

Apply different transformations based on conditions:

```go
type Employee struct {
    Name   string
    Salary float64
}

employees := []Employee{
    {"alice", 50000},
    {"bob", 60000},
    {"charlie", 55000},
}

plygo.From(employees).
    Transform(func(e Employee) Employee {
        e.Name = strings.Title(e.Name)
        if e.Salary < 55000 {
            e.Salary = e.Salary * 1.1  // 10% raise
        }
        return e
    }).
    Show()
```

:::success Result
```
+---------+----------+
|    Name |   Salary |
+---------+----------+
| Alice   | 55000.00 |
| Bob     | 60000.00 |
| Charlie | 55000.00 |
+---------+----------+
[3 rows × 2 columns]
```
:::

## Chained Transformation

Combine Transform with other pipeline operations:

```go
type Product struct {
    Name  string
    Price float64
    Stock int
}

products := []Product{
    {"Laptop", 1000, 5},
    {"Mouse", 50, 20},
    {"Keyboard", 80, 15},
}

plygo.From(products).
    Where("Stock").GreaterThan(10).
    Transform(func(p Product) Product {
        p.Price = p.Price * 0.85  // 15% discount for high stock
        return p
    }).
    Show()
```

:::success Result
```
+----------+-------+-------+
|     Name | Price | Stock |
+----------+-------+-------+
| Mouse    | 42.50 |    20 |
| Keyboard | 68.00 |    15 |
+----------+-------+-------+
[2 rows × 3 columns]
```
:::

:::tip Use Cases
Transform is perfect for:
- Applying discounts or markups
- Formatting text fields
- Calculating derived values
- Data normalization
- Currency conversion
:::

Next: [Utilities](09-utilities.md)
