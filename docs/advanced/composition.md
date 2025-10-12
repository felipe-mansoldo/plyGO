# Pipeline Composition

Learn advanced techniques for composing and reusing plyGO pipelines.

## Chaining Operations

The simplest form of composition is chaining multiple operations inline:

```go
type Product struct {
    Name     string
    Price    float64
    Category string
    InStock  bool
}

products := []Product{
    {"Laptop", 1000, "Electronics", true},
    {"Mouse", 25, "Electronics", false},
    {"Desk", 300, "Furniture", true},
    {"Chair", 150, "Furniture", true},
    {"Keyboard", 75, "Electronics", true},
}

// Compose filters inline
plygo.From(products).
    Where("InStock").IsTrue().
    Where("Price").GreaterThan(100.0).
    Show()
```

::: tip Result
```
+--------+---------+-------------+---------+
|   Name |   Price |    Category | InStock |
+--------+---------+-------------+---------+
| Laptop | 1000.00 | Electronics | true    |
| Desk   |  300.00 | Furniture   | true    |
| Chair  |  150.00 | Furniture   | true    |
+--------+---------+-------------+---------+
[3 rows × 4 columns]
```
:::

## Storing Intermediate Results

Store intermediate results to reuse filtered data:

```go
type Sale struct {
    Date     string
    Product  string
    Amount   float64
    Region   string
}

sales := []Sale{
    {"2024-01-01", "Laptop", 1000, "North"},
    {"2024-01-02", "Monitor", 300, "South"},
    {"2024-01-03", "Server", 2000, "North"},
    {"2024-01-04", "Router", 150, "East"},
    {"2024-01-05", "Laptop", 900, "North"},
}

// Store intermediate pipeline
highValueSales := plygo.From(sales).
    Where("Amount").GreaterThan(500.0).
    Collect()

// Reuse filtered data
northSales := plygo.From(highValueSales).
    Where("Region").Equals("North").
    OrderBy("Amount").Desc().
    Collect()

plygo.From(northSales).Show()
```

::: tip Result
```
+------------+---------+---------+--------+
|       Date | Product |  Amount | Region |
+------------+---------+---------+--------+
| 2024-01-03 | Server  | 2000.00 | North  |
| 2024-01-01 | Laptop  | 1000.00 | North  |
| 2024-01-05 | Laptop  |  900.00 | North  |
+------------+---------+---------+--------+
[3 rows × 4 columns]
```
:::

## Composable Helper Functions

Create reusable functions that return slices:

```go
type Order struct {
    ID       int
    Customer string
    Amount   float64
    Status   string
}

// Helper functions that work with slices
func getCompleted(orders []Order) []Order {
    return plygo.From(orders).
        Where("Status").Equals("completed").
        Collect()
}

func getLargeOrders(orders []Order) []Order {
    return plygo.From(orders).
        Where("Amount").GreaterThan(1000.0).
        Collect()
}

func sortByAmount(orders []Order) []Order {
    return plygo.From(orders).
        OrderBy("Amount").Desc().
        Collect()
}

orders := []Order{
    {1, "Alice", 1500, "completed"},
    {2, "Bob", 800, "pending"},
    {3, "Charlie", 2000, "completed"},
    {4, "Diana", 500, "completed"},
    {5, "Eve", 1200, "completed"},
}

// Compose functions
result := sortByAmount(getLargeOrders(getCompleted(orders)))
plygo.From(result).Show()
```

::: tip Result
```
+----+----------+---------+-----------+
| ID | Customer |  Amount |    Status |
+----+----------+---------+-----------+
|  3 | Charlie  | 2000.00 | completed |
|  1 | Alice    | 1500.00 | completed |
|  5 | Eve      | 1200.00 | completed |
+----+----------+---------+-----------+
[3 rows × 4 columns]
```
:::

## Complex Multi-Stage Pipelines

Combine multiple transformations and filters:

```go
type Employee struct {
    Name       string
    Department string
    Salary     float64
    YearsExp   int
}

employees := []Employee{
    {"alice smith", "engineering", 70000, 3},
    {"bob jones", "sales", 60000, 5},
    {"charlie brown", "engineering", 85000, 7},
    {"diana prince", "marketing", 75000, 4},
}

// Complex multi-stage pipeline
plygo.From(employees).
    Where("Department").Equals("engineering").
    Transform(func(e Employee) Employee {
        e.Name = strings.Title(e.Name)
        return e
    }).
    Where("YearsExp").GreaterThan(4).
    Transform(func(e Employee) Employee {
        e.Salary = e.Salary * 1.1  // 10% bonus
        return e
    }).
    OrderBy("Salary").Desc().
    Show()
```

::: tip Result
```
+---------------+-------------+----------+----------+
|          Name |  Department |   Salary | YearsExp |
+---------------+-------------+----------+----------+
| Charlie Brown | engineering | 93500.00 |        7 |
+---------------+-------------+----------+----------+
[1 rows × 4 columns]
```
:::

::: tip Composition Best Practices
1. **Use `Collect()`** to materialize intermediate results when you need to branch or reuse data
2. **Create helper functions** that accept and return slices for maximum flexibility
3. **Keep pipelines focused** - each function should do one thing well
4. **Name functions clearly** - descriptive names make composition readable
5. **Consider memory** - `Collect()` creates a new slice, use judiciously with large datasets
:::

::: warning Performance Considerations
- Each `Collect()` creates a new slice copy
- For very large datasets, minimize intermediate `Collect()` calls
- Consider using `Which()` for index-based composition (see Positions tutorial)
- Chain operations inline when you don't need to reuse intermediate results
:::

Next: [Custom Helpers](/advanced/custom-helpers)
