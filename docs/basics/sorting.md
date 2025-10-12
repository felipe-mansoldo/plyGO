# Sorting Data

Learn how to sort your data using `OrderBy` and `ThenBy`.

## Sort Ascending

Use `OrderBy()` followed by `Asc()` to sort in ascending order:

```go
type Product struct {
    Name  string
    Price float64
    Stock int
}

products := []Product{
    {"Laptop", 999.99, 15},
    {"Mouse", 29.99, 50},
    {"Keyboard", 79.99, 30},
    {"Monitor", 299.99, 20},
}

plygo.From(products).OrderBy("Price").Asc().Show()
```

::: tip Result
```
+----------+--------+-------+
|     Name |  Price | Stock |
+----------+--------+-------+
| Mouse    |  29.99 |    50 |
| Keyboard |  79.99 |    30 |
| Monitor  | 299.99 |    20 |
| Laptop   | 999.99 |    15 |
+----------+--------+-------+
[4 rows × 3 columns]
```
:::

## Sort Descending

Use `OrderBy()` followed by `Desc()` to sort in descending order:

```go
products := []Product{
    {"Laptop", 999.99, 15},
    {"Mouse", 29.99, 50},
    {"Keyboard", 79.99, 30},
    {"Monitor", 299.99, 20},
}

plygo.From(products).OrderBy("Stock").Desc().Show()
```

::: tip Result
```
+----------+--------+-------+
|     Name |  Price | Stock |
+----------+--------+-------+
| Mouse    |  29.99 |    50 |
| Keyboard |  79.99 |    30 |
| Monitor  | 299.99 |    20 |
| Laptop   | 999.99 |    15 |
+----------+--------+-------+
[4 rows × 3 columns]
```
:::

## Multi-Level Sorting

Use `ThenBy()` to add secondary sorting criteria:

```go
type Employee struct {
    Name       string
    Department string
    Salary     float64
}

employees := []Employee{
    {"Alice", "Engineering", 75000},
    {"Bob", "Engineering", 90000},
    {"Charlie", "Sales", 70000},
    {"Diana", "Sales", 85000},
}

plygo.From(employees).
    OrderBy("Department").Asc().
    ThenBy("Salary").Desc().
    Show()
```

::: tip Result
```
+---------+-------------+----------+
|    Name |  Department |   Salary |
+---------+-------------+----------+
| Bob     | Engineering | 90000.00 |
| Alice   | Engineering | 75000.00 |
| Diana   | Sales       | 85000.00 |
| Charlie | Sales       | 70000.00 |
+---------+-------------+----------+
[4 rows × 3 columns]
```
:::

Next: [Grouping](/basics/grouping)
