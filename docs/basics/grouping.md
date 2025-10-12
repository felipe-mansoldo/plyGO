# Grouping & Aggregation

Learn how to group data and perform aggregations like count, sum, and average.

## Count by Group

Use `GroupBy()` with `Count()` to count items in each group:

```go
type Sale struct {
    Product  string
    Category string
    Amount   float64
    Quantity int
}

sales := []Sale{
    {"Laptop", "Electronics", 999.99, 2},
    {"Mouse", "Electronics", 29.99, 5},
    {"Desk", "Furniture", 299.99, 1},
    {"Chair", "Furniture", 199.99, 3},
    {"Keyboard", "Electronics", 79.99, 4},
}

result := plygo.From(sales).GroupBy("Category").Count()
for category, count := range result {
    fmt.Printf("%v: %d\n", category, count)
}
```

::: tip Result
```
Electronics: 3
Furniture: 2
```
:::

## Sum by Group

Use `GroupBy()` with `Sum()` to sum values in each group:

```go
sales := []Sale{
    {"Laptop", "Electronics", 999.99, 2},
    {"Mouse", "Electronics", 29.99, 5},
    {"Desk", "Furniture", 299.99, 1},
    {"Chair", "Furniture", 199.99, 3},
    {"Keyboard", "Electronics", 79.99, 4},
}

result := plygo.From(sales).GroupBy("Category").Sum("Quantity")
for category, total := range result {
    fmt.Printf("%v: %.0f\n", category, total)
}
```

::: tip Result
```
Electronics: 11
Furniture: 4
```
:::

## Average by Group

Use `GroupBy()` with `Avg()` to calculate averages:

```go
sales := []Sale{
    {"Laptop", "Electronics", 999.99, 2},
    {"Mouse", "Electronics", 29.99, 5},
    {"Desk", "Furniture", 299.99, 1},
    {"Chair", "Furniture", 199.99, 3},
    {"Keyboard", "Electronics", 79.99, 4},
}

result := plygo.From(sales).GroupBy("Category").Avg("Amount")
for category, avg := range result {
    fmt.Printf("%v: %.2f\n", category, avg)
}
```

::: tip Result
```
Electronics: 369.99
Furniture: 249.99
```
:::

::: tip Available Aggregations
GroupBy supports these aggregation functions:
- `Count()` - Count items in each group
- `Sum(field)` - Sum numeric field values
- `Avg(field)` - Average of numeric field values
- `Min(field)` - Minimum value in each group
- `Max(field)` - Maximum value in each group
:::

Next: [Transformation](/basics/transformation)
