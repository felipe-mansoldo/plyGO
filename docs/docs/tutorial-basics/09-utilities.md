---
sidebar_position: 9
---

# Utility Functions

Explore essential utility functions for data manipulation.

## Limit Results

Use `Limit()` to restrict the number of results:

```go
type Product struct {
    Name  string
    Price float64
}

products := []Product{
    {"Laptop", 1000},
    {"Mouse", 50},
    {"Keyboard", 80},
    {"Monitor", 300},
    {"Webcam", 150},
}

plygo.From(products).Limit(3).Show()
```

:::success Result
```
+----------+---------+
|     Name |   Price |
+----------+---------+
| Laptop   | 1000.00 |
| Mouse    |   50.00 |
| Keyboard |   80.00 |
+----------+---------+
[3 rows × 2 columns]
```
:::

## Skip Items

Use `Skip()` to skip the first N items:

```go
products := []Product{
    {"Laptop", 1000},
    {"Mouse", 50},
    {"Keyboard", 80},
    {"Monitor", 300},
    {"Webcam", 150},
}

plygo.From(products).Skip(2).Show()
```

:::success Result
```
+----------+--------+
|     Name |  Price |
+----------+--------+
| Keyboard |  80.00 |
| Monitor  | 300.00 |
| Webcam   | 150.00 |
+----------+--------+
[3 rows × 2 columns]
```
:::

## Remove Duplicates

Use `Distinct()` to keep only unique values for a field:

```go
type Order struct {
    Customer string
    Product  string
    Amount   float64
}

orders := []Order{
    {"Alice", "Laptop", 1000},
    {"Bob", "Mouse", 50},
    {"Alice", "Keyboard", 80},
    {"Charlie", "Monitor", 300},
    {"Bob", "Webcam", 150},
}

plygo.From(orders).Distinct("Customer").Show()
```

:::success Result
```
+----------+---------+---------+
| Customer | Product |  Amount |
+----------+---------+---------+
| Alice    | Laptop  | 1000.00 |
| Bob      | Mouse   |   50.00 |
| Charlie  | Monitor |  300.00 |
+----------+---------+---------+
[3 rows × 3 columns]
```
:::

## Get First or Last Item

Use `First()` or `Last()` to get a single item:

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

first, ok := plygo.From(products).First()
if ok {
    fmt.Printf("First: %s - $%.2f\n", first.Name, first.Price)
}

last, ok := plygo.From(products).Last()
if ok {
    fmt.Printf("Last: %s - $%.2f\n", last.Name, last.Price)
}
```

:::success Result
```
First: Laptop - $1000.00
Last: Keyboard - $80.00
```
:::

## Count Items

Use `Count()` to get the total number of items:

```go
type Product struct {
    Name  string
    Price float64
}

products := []Product{
    {"Laptop", 1000},
    {"Mouse", 50},
    {"Keyboard", 80},
    {"Monitor", 300},
}

total := plygo.From(products).Count()
fmt.Printf("Total products: %d\n", total)

expensive := plygo.From(products).
    Where("Price").GreaterThan(100.0).
    Collect()
fmt.Printf("Expensive products (>$100): %d\n", len(expensive))
```

:::success Result
```
Total products: 4
Expensive products (>$100): 2
```
:::

:::tip Pagination
Combine `Skip()` and `Limit()` for pagination:
```go
// Page 2, with 10 items per page
plygo.From(data).Skip(10).Limit(10)
```
:::

Next: [Show](10-show.md)
