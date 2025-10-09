---
sidebar_position: 3
---

# Filtering

Master the Where clause to filter your data.

## Basic Conditions

```go
type Person struct {
    Name   string
    Age    int
    City   string
    Active bool
}

people := []Person{
    {"Alice", 30, "NYC", true},
    {"Bob", 25, "LA", false},
    {"Charlie", 35, "NYC", true},
    {"Diana", 28, "Chicago", true},
}

plygo.From(people).
    Where("Age").GreaterThan(30).
    Show()
```

:::success Result
```
+---------+-----+------+--------+
|    Name | Age | City | Active |
+---------+-----+------+--------+
| Charlie |  35 | NYC  | true   |
+---------+-----+------+--------+
[1 rows × 4 columns]
```
:::

## Chaining AND

```go
plygo.From(people).
    Where("Age").GreaterThan(25).
    Where("Active").IsTrue().
    Show()
```

:::success Result
```
+---------+-----+---------+--------+
|    Name | Age |    City | Active |
+---------+-----+---------+--------+
| Alice   |  30 | NYC     | true   |
| Charlie |  35 | NYC     | true   |
| Diana   |  28 | Chicago | true   |
+---------+-----+---------+--------+
[3 rows × 4 columns]
```
:::

## OR Conditions

```go
plygo.From(people).
    Where("City").Equals("NYC").Or("City").Equals("LA").
    Show()
```

:::success Result
```
+---------+-----+------+--------+
|    Name | Age | City | Active |
+---------+-----+------+--------+
| Alice   |  30 | NYC  | true   |
| Bob     |  25 | LA   | false  |
| Charlie |  35 | NYC  | true   |
+---------+-----+------+--------+
[3 rows × 4 columns]
```
:::

Next: [Selecting Fields](04-selecting.md)
