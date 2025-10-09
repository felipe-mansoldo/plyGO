---
sidebar_position: 4
---

# Selecting Fields

Learn how to select specific fields or columns from your data.

## Select by Field Name

Use `Select()` to choose specific fields by name:

```go
type Employee struct {
    Name       string
    Age        int
    Department string
    Salary     float64
}

employees := []Employee{
    {"Alice", 30, "Engineering", 75000},
    {"Bob", 25, "Marketing", 60000},
    {"Charlie", 35, "Engineering", 90000},
    {"Diana", 28, "Sales", 70000},
}

plygo.From(employees).Select("Name", "Salary").Show()
```

:::success Result
```
+---------+----------+
|    Name |   Salary |
+---------+----------+
| Alice   | 75000.00 |
| Bob     | 60000.00 |
| Charlie | 90000.00 |
| Diana   | 70000.00 |
+---------+----------+
[4 rows × 2 columns]
```
:::

## Select by Column Index

Use `AtCol()` to select columns by their position (1-based index):

```go
employees := []Employee{
    {"Alice", 30, "Engineering", 75000},
    {"Bob", 25, "Marketing", 60000},
    {"Charlie", 35, "Engineering", 90000},
}

plygo.From(employees).AtCol(1, 4).Show()
```

:::success Result
```
+---------+----------+
|    Name |   Salary |
+---------+----------+
| Alice   | 75000.00 |
| Bob     | 60000.00 |
| Charlie | 90000.00 |
+---------+----------+
[3 rows × 2 columns]
```
:::

## Select Column Range

Use `ColRange()` to select a range of columns:

```go
employees := []Employee{
    {"Alice", 30, "Engineering", 75000},
    {"Bob", 25, "Marketing", 60000},
}

plygo.From(employees).ColRange(2, 4).Show()
```

:::success Result
```
+-----+-------------+
| Age |  Department |
+-----+-------------+
|  30 | Engineering |
|  25 | Marketing   |
+-----+-------------+
[2 rows × 2 columns]
```
:::

Next: [Position-Based Selection](05-positions.md)
