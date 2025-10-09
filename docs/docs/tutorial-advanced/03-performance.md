---
sidebar_position: 3
---

# Performance Optimization

Learn techniques to optimize plyGO pipelines for better performance.

## Minimize Collect() Calls

Each `Collect()` creates a new slice. Chain operations to reduce intermediate copies:

```go
type Record struct {
    ID    int
    Value float64
}

func generateData(n int) []Record {
    data := make([]Record, n)
    for i := 0; i < n; i++ {
        data[i] = Record{ID: i, Value: float64(i * 100)}
    }
    return data
}

data := generateData(10000)

// ❌ Bad: Multiple Collect() calls (creates 3 intermediate slices)
step1 := plygo.From(data).Where("Value").GreaterThan(500.0).Collect()
step2 := plygo.From(step1).Where("Value").LessThan(500000.0).Collect()
result1 := plygo.From(step2).OrderBy("Value").Desc().Collect()

// ✅ Good: Chain operations, single Collect()
result2 := plygo.From(data).
    Where("Value").GreaterThan(500.0).
    Where("Value").LessThan(500000.0).
    OrderBy("Value").Desc().
    Collect()
```

:::success Result
```
Multiple Collect(): 5.1ms
Single Collect(): 4.2ms
Results match: true
```
:::

## Use Index-Based Filtering

For large records, store indices instead of full data copies:

```go
type LargeRecord struct {
    ID          int
    Data        [100]byte  // Simulate large data
    Value       float64
    Description string
}

// Create dataset
data := make([]LargeRecord, 1000)
for i := 0; i < 1000; i++ {
    data[i] = LargeRecord{
        ID:          i,
        Value:       float64(i * 10),
        Description: fmt.Sprintf("Record %d", i),
    }
}

// Memory-efficient: Store only indices
indices := plygo.From(data).
    Where("Value").GreaterThan(5000.0).
    Which()

fmt.Printf("Filtered %d records\n", len(indices))
fmt.Printf("Storing %d indices instead of full records\n", len(indices))
fmt.Printf("Memory saved: ~%d bytes per record\n", 100+16+8)

// Access original data when needed
fmt.Println("\nFirst 3 filtered records:")
for i := 0; i < 3 && i < len(indices); i++ {
    rec := data[indices[i]-1]
    fmt.Printf("  ID %d: Value=%.2f\n", rec.ID, rec.Value)
}
```

:::success Result
```
Filtered 499 records
Storing 499 indices instead of full records
Memory saved: ~124 bytes per record

First 3 filtered records:
  ID 501: Value=5010.00
  ID 502: Value=5020.00
  ID 503: Value=5030.00
```
:::

## Filter Before Expensive Operations

Apply cheap filters first to reduce the dataset before expensive operations:

```go
type Product struct {
    Name     string
    Price    float64
    Category string
    InStock  bool
}

func expensiveTransform(p Product) Product {
    time.Sleep(100 * time.Microsecond)
    p.Name = strings.ToUpper(p.Name)
    return p
}

products := make([]Product, 100)
for i := 0; i < 100; i++ {
    products[i] = Product{
        Name:     fmt.Sprintf("Product%d", i),
        Price:    float64(i * 10),
        Category: "Electronics",
        InStock:  i%3 == 0,
    }
}

// ❌ Bad: Transform all (100 items), then filter
result1 := plygo.From(products).
    Transform(expensiveTransform).
    Where("InStock").IsTrue().
    Collect()

// ✅ Good: Filter first (to ~34 items), then transform
result2 := plygo.From(products).
    Where("InStock").IsTrue().
    Transform(expensiveTransform).
    Collect()
```

:::success Result
```
Transform then filter: 59.0ms (100 transforms)
Filter then transform: 20.3ms (34 transforms)
Speedup: 2.90x
```
:::

## Limit Results Early

When you only need top N results, slice after sorting:

```go
type Sale struct {
    ID     int
    Amount float64
    Region string
}

sales := make([]Sale, 10000)
for i := 0; i < 10000; i++ {
    sales[i] = Sale{
        ID:     i,
        Amount: float64(i * 100),
        Region: "North",
    }
}

// Sort and get top 10
topSales := plygo.From(sales).
    OrderBy("Amount").Desc().
    Collect()

result := topSales[:10]

fmt.Printf("Sorted all %d records\n", len(sales))
fmt.Printf("Showing top %d\n", len(result))
fmt.Printf("Top sale: ID=%d, Amount=%.2f\n", 
    result[0].ID, result[0].Amount)
```

:::success Result
```
Sorted all 10000 records
Showing top 10
Top sale: ID=9999, Amount=999900.00
```
:::

:::tip Performance Best Practices

**Memory Optimization:**
1. **Use `Which()`** for large records - store indices instead of copies
2. **Chain operations** - minimize `Collect()` calls
3. **Slice results** - use `[:n]` after sorting instead of creating new pipelines

**Execution Optimization:**
4. **Filter early** - apply cheap filters before expensive transforms
5. **Order matters** - filter → transform → sort → limit
6. **Avoid re-pipelines** - don't wrap `Collect()` results unnecessarily

**When to use `Collect()`:**
- When you need to branch the pipeline
- When you need to reuse filtered data
- At the end of your pipeline
- NOT between every operation
:::

:::warning Common Performance Pitfalls

**❌ Avoid these patterns:**
```go
// Multiple unnecessary Collect() calls
data1 := plygo.From(data).Where(...).Collect()
data2 := plygo.From(data1).Transform(...).Collect()
data3 := plygo.From(data2).OrderBy(...).Collect()

// Re-wrapping collected data
result := plygo.From(plygo.From(data).Collect()).Collect()

// Transforming before filtering
plygo.From(data).
    Transform(expensiveFunc).  // Transforms ALL items
    Where("field").Equals(x)   // Then filters
```

**✅ Use these instead:**
```go
// Single pipeline chain
result := plygo.From(data).
    Where(...).
    Transform(...).
    OrderBy(...).
    Collect()

// Filter before transform
result := plygo.From(data).
    Where("field").Equals(x).  // Filter first
    Transform(expensiveFunc).  // Transform fewer items
    Collect()
```
:::

Next: [Concurrency](04-concurrency.md)
