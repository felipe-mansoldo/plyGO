# Error Handling

Best practices for handling errors and edge cases in plyGO pipelines.

## Type Safety

plyGO uses Go's type system to catch errors at compile time:

```go
type Person struct {
    Name string
    Age  int
}

people := []Person{
    {"Alice", 30},
    {"Bob", 25},
}

// This won't compile - type mismatch
// plygo.From(people).Where("Age").Equals("30") // string instead of int

// Correct - matching types
plygo.From(people).Where("Age").Equals(30).Show()
```

::: tip Result
```
+-------+-----+
|  Name | Age |
+-------+-----+
| Alice |  30 |
+-------+-----+
[1 rows × 2 columns]
```
:::

## Handling Empty Results

Always check for empty results when needed:

```go
type Person struct {
    Name   string
    Age    int
    Salary float64
}

people := []Person{
    {"Alice", 30, 75000},
    {"Bob", 25, 60000},
    {"Charlie", 35, 90000},
}

result := plygo.From(people).
    Where("Age").GreaterThan(1000).
    Collect()

if len(result) == 0 {
    fmt.Println("No results found")
} else {
    fmt.Printf("Found %d results\n", len(result))
}
```

::: tip Result
```
No results found
```
:::

## Using First() and Last() Safely

Both `First()` and `Last()` return a boolean to indicate if a value exists:

```go
type Product struct {
    Name  string
    Price float64
}

products := []Product{
    {"Laptop", 1000},
}

// Safe usage with ok pattern
first, ok := plygo.From(products).
    Where("Price").GreaterThan(2000).
    First()

if ok {
    fmt.Printf("Found: %s\n", first.Name)
} else {
    fmt.Println("No product found matching criteria")
}
```

::: tip Result
```
No product found matching criteria
```
:::

## Validation Before Processing

Validate your data before applying operations:

```go
type Employee struct {
    Name   string
    Salary float64
}

employees := []Employee{
    {"Alice", 50000},
    {"Bob", 60000},
}

count := plygo.From(employees).Count()
if count == 0 {
    fmt.Println("No employees to process")
} else {
    fmt.Printf("Processing %d employees\n", count)
    plygo.From(employees).
        Where("Salary").GreaterThan(55000.0).
        Show()
}
```

::: tip Result
```
Processing 2 employees
+------+----------+
| Name |   Salary |
+------+----------+
| Bob  | 60000.00 |
+------+----------+
[1 rows × 2 columns]
```
:::

::: tip Best Practices
1. **Use type-safe comparisons** - Let the compiler catch type errors
2. **Check empty results** - Use `Count()` or check `len()` of `Collect()`
3. **Use ok pattern** - Always check the boolean return from `First()` and `Last()`
4. **Validate field names** - Ensure field names match struct fields exactly
5. **Handle nil safely** - Check for empty datasets before processing
:::

::: warning Common Pitfalls
- Using string values when comparing numeric fields
- Not checking if `First()` or `Last()` returns valid data
- Assuming results exist without validation
- Case-sensitive field names (use exact struct field names)
:::

Next: [Advanced Topics](/advanced/composition)
