# Position-Based Selection

Learn how to select data by row positions and leverage original indices for memory efficiency.

## Select Specific Rows

Use `AtRow()` to select rows by their position (1-based index):

```go
type Person struct {
    Name string
    Age  int
}

people := []Person{
    {"Alice", 30},
    {"Bob", 25},
    {"Charlie", 35},
    {"Diana", 28},
    {"Eve", 32},
}

plygo.From(people).AtRow(1, 3, 5).Show()
```

::: tip Result
```
+---------+-----+
|    Name | Age |
+---------+-----+
| Alice   |  30 |
| Charlie |  35 |
| Eve     |  32 |
+---------+-----+
[3 rows × 2 columns]
```
:::

## Select Row Range

Use `RowRange()` to select a continuous range of rows:

```go
people := []Person{
    {"Alice", 30},
    {"Bob", 25},
    {"Charlie", 35},
    {"Diana", 28},
    {"Eve", 32},
}

plygo.From(people).RowRange(2, 4).Show()
```

::: tip Result
```
+---------+-----+
|    Name | Age |
+---------+-----+
| Bob     |  25 |
| Charlie |  35 |
+---------+-----+
[2 rows × 2 columns]
```
:::

## Get Last N Rows

Use `Tail()` to get the last N rows:

```go
people := []Person{
    {"Alice", 30},
    {"Bob", 25},
    {"Charlie", 35},
    {"Diana", 28},
    {"Eve", 32},
}

plygo.From(people).Tail(3).Show()
```

::: tip Result
```
+---------+-----+
|    Name | Age |
+---------+-----+
| Charlie |  35 |
| Diana   |  28 |
| Eve     |  32 |
+---------+-----+
[3 rows × 2 columns]
```
:::

## Random Sample

Use `Sample()` to get a random sample of N rows:

```go
people := []Person{
    {"Alice", 30},
    {"Bob", 25},
    {"Charlie", 35},
    {"Diana", 28},
    {"Eve", 32},
}

plygo.From(people).Sample(2).Show()
```

::: tip Result
```
+-------+-----+
|  Name | Age |
+-------+-----+
| Alice |  30 |
| Bob   |  25 |
+-------+-----+
[2 rows × 2 columns]
```
:::

## Track Original Indices with Which()

Get the original positions of filtered data - perfect for memory efficiency:

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
    {"Monitor", 300, 8},
    {"Webcam", 150, 12},
}

// Find expensive products and get their original positions
indices := plygo.From(products).
    Where("Price").GreaterThan(100.0).
    Which()

fmt.Println("Original indices of expensive products:", indices)
fmt.Println("\nAccessing original data using indices:")
for _, idx := range indices {
    p := products[idx-1] // Which() returns 1-based indices
    fmt.Printf("  Position %d: %s - $%.2f\n", idx, p.Name, p.Price)
}
```

::: tip Result
```
Original indices of expensive products: [1 4 5]

Accessing original data using indices:
  Position 1: Laptop - $1000.00
  Position 4: Monitor - $300.00
  Position 5: Webcam - $150.00
```
:::

## Memory Efficiency with Indices

Store only indices instead of duplicating large data:

```go
type LargeData struct {
    ID          int
    Name        string
    Description string
    Value       float64
}

// Simulate large dataset
data := []LargeData{
    {1, "Item A", "Long description A...", 100.0},
    {2, "Item B", "Long description B...", 200.0},
    {3, "Item C", "Long description C...", 150.0},
    {4, "Item D", "Long description D...", 250.0},
    {5, "Item E", "Long description E...", 180.0},
}

// Store only indices, not the full data
highValueIndices := plygo.From(data).
    Where("Value").GreaterThan(150.0).
    Which()

fmt.Printf("Storing %d indices instead of full records\n", len(highValueIndices))
fmt.Println("Indices:", highValueIndices)

// Access original data when needed
fmt.Println("\nAccessing on demand:")
for _, idx := range highValueIndices {
    item := data[idx-1]
    fmt.Printf("  ID %d: %s = $%.2f\n", item.ID, item.Name, item.Value)
}
```

::: tip Result
```
Storing 3 indices instead of full records
Indices: [2 4 5]

Accessing on demand:
  ID 2: Item B = $200.00
  ID 4: Item D = $250.00
  ID 5: Item E = $180.00
```
:::

## Track Positions Through Pipeline

Use `Positions()` to track row positions after multiple filters:

```go
type Sale struct {
    Date     string
    Product  string
    Amount   float64
    Quantity int
}

sales := []Sale{
    {"2024-01-01", "Laptop", 1000, 2},
    {"2024-01-02", "Mouse", 50, 5},
    {"2024-01-03", "Keyboard", 80, 3},
    {"2024-01-04", "Monitor", 300, 1},
    {"2024-01-05", "Webcam", 150, 4},
}

filtered := plygo.From(sales).
    Where("Amount").GreaterThan(100.0)

positions := filtered.Positions()

fmt.Printf("Found %d high-value sales at positions: %v\n",
    positions.RowCount(), positions.Rows)

fmt.Println("\nOriginal records:")
for _, pos := range positions.Rows {
    sale := sales[pos-1]
    fmt.Printf("  Row %d: %s - %s ($%.2f)\n",
        pos, sale.Date, sale.Product, sale.Amount)
}
```

::: tip Result
```
Found 3 high-value sales at positions: [1 4 5]

Original records:
  Row 1: 2024-01-01 - Laptop ($1000.00)
  Row 4: 2024-01-04 - Monitor ($300.00)
  Row 5: 2024-01-05 - Webcam ($150.00)
```
:::

## Display with Original Row Numbers

Show filtered data with original row numbers preserved:

```go
products := []Product{
    {"Laptop", 1000, 5},
    {"Mouse", 50, 20},
    {"Keyboard", 80, 15},
    {"Monitor", 300, 8},
    {"Webcam", 150, 12},
}

// Show filtered results with original row numbers
plygo.From(products).
    Where("Stock").LessThan(15).
    Show(plygo.WithOriginalIndices(true))
```

::: tip Result
```
+---+---------+---------+-------+
| # |    Name |   Price | Stock |
+---+---------+---------+-------+
| 1 | Laptop  | 1000.00 |     5 |
| 4 | Monitor |  300.00 |     8 |
| 5 | Webcam  |  150.00 |    12 |
+---+---------+---------+-------+
[3 rows × 4 columns]
```
:::

::: tip Memory Efficiency Benefits
Using `Which()` and `Positions()` is especially useful when:
- **Working with large datasets** - Store indices instead of duplicating data
- **Multiple filter combinations** - Test different filters without copying data
- **Reference lookups** - Maintain references to original dataset positions
- **Batch processing** - Process items from original dataset in batches
- **Memory constraints** - Minimize memory footprint by storing only indices

**Example:** With 10,000 records of 1KB each, storing indices (40 bytes) instead of filtered copies can save megabytes of memory.
:::

Next: [Sorting](/basics/sorting)
