# Custom Helper Functions

Learn how to build reusable helper functions for common data operations.

## Custom Filter Helpers

Create domain-specific filter functions:

```go
type Product struct {
    Name     string
    Price    float64
    Category string
}

// Custom helper: Filter by price range
func InPriceRange(data []Product, min, max float64) []Product {
    return plygo.From(data).
        Where("Price").GreaterThan(min).
        Where("Price").LessThan(max).
        Collect()
}

// Custom helper: Filter by category prefix
func InCategory(data []Product, prefix string) []Product {
    filtered := []Product{}
    for _, p := range data {
        if strings.HasPrefix(strings.ToLower(p.Category), strings.ToLower(prefix)) {
            filtered = append(filtered, p)
        }
    }
    return filtered
}

products := []Product{
    {"Laptop", 1000, "Electronics"},
    {"Mouse", 50, "Electronics"},
    {"Desk", 300, "Furniture"},
    {"Monitor", 250, "Electronics"},
    {"Chair", 150, "Furniture"},
}

// Use custom helpers
result := InPriceRange(InCategory(products, "Electr"), 100.0, 500.0)
plygo.From(result).Show()
```

::: tip Result
```
+---------+--------+-------------+
|    Name |  Price |    Category |
+---------+--------+-------------+
| Monitor | 250.00 | Electronics |
+---------+--------+-------------+
[1 rows × 3 columns]
```
:::

## Custom Transformation Helpers

Build helpers for common transformations:

```go
type Employee struct {
    Name   string
    Salary float64
    Title  string
}

// Apply percentage increase
func ApplyRaise(employees []Employee, percent float64) []Employee {
    return plygo.From(employees).
        Transform(func(e Employee) Employee {
            e.Salary = e.Salary * (1 + percent/100)
            return e
        }).
        Collect()
}

// Normalize names to title case
func NormalizeNames(employees []Employee) []Employee {
    return plygo.From(employees).
        Transform(func(e Employee) Employee {
            e.Name = strings.Title(strings.ToLower(e.Name))
            return e
        }).
        Collect()
}

employees := []Employee{
    {"alice SMITH", 50000, "Engineer"},
    {"BOB jones", 60000, "Manager"},
    {"charlie BROWN", 55000, "Designer"},
}

// Chain custom helpers
result := ApplyRaise(NormalizeNames(employees), 10)
plygo.From(result).Show()
```

::: tip Result
```
+---------------+----------+----------+
|          Name |   Salary |    Title |
+---------------+----------+----------+
| Alice Smith   | 55000.00 | Engineer |
| Bob Jones     | 66000.00 | Manager  |
| Charlie Brown | 60500.00 | Designer |
+---------------+----------+----------+
[3 rows × 3 columns]
```
:::

## Custom Aggregation Helpers

Create helpers for calculations and summaries:

```go
type Sale struct {
    Product  string
    Amount   float64
    Quantity int
}

// Custom helper: Calculate total revenue
func TotalRevenue(sales []Sale) float64 {
    total := 0.0
    for _, s := range sales {
        total += s.Amount
    }
    return total
}

// Custom helper: Average sale amount
func AverageSale(sales []Sale) float64 {
    if len(sales) == 0 {
        return 0
    }
    return TotalRevenue(sales) / float64(len(sales))
}

// Custom helper: Top N products
func TopProducts(sales []Sale, n int) []Sale {
    sorted := plygo.From(sales).
        OrderBy("Amount").Desc().
        Collect()
    if len(sorted) > n {
        return sorted[:n]
    }
    return sorted
}

sales := []Sale{
    {"Laptop", 1000, 2},
    {"Mouse", 50, 10},
    {"Keyboard", 80, 5},
    {"Monitor", 300, 3},
    {"Webcam", 150, 4},
}

// Use custom aggregation helpers
fmt.Printf("Total Revenue: $%.2f\n", TotalRevenue(sales))
fmt.Printf("Average Sale: $%.2f\n", AverageSale(sales))
fmt.Println("\nTop 3 Products:")
plygo.From(TopProducts(sales, 3)).Show()
```

::: tip Result
```
Total Revenue: $1580.00
Average Sale: $316.00

Top 3 Products:
+---------+---------+----------+
| Product |  Amount | Quantity |
+---------+---------+----------+
| Laptop  | 1000.00 |        2 |
| Monitor |  300.00 |        3 |
| Webcam  |  150.00 |        4 |
+---------+---------+----------+
[3 rows × 3 columns]
```
:::

## Domain-Specific Helper Packages

Organize helpers into domain-specific packages:

```go
type Transaction struct {
    ID     int
    Amount float64
    Date   string
    Type   string
}

// Domain helpers for financial data
func CreditTransactions(txns []Transaction) []Transaction {
    return plygo.From(txns).
        Where("Type").Equals("credit").
        Collect()
}

func DebitTransactions(txns []Transaction) []Transaction {
    return plygo.From(txns).
        Where("Type").Equals("debit").
        Collect()
}

func HighValueTransactions(txns []Transaction, threshold float64) []Transaction {
    return plygo.From(txns).
        Where("Amount").GreaterThan(threshold).
        OrderBy("Amount").Desc().
        Collect()
}

transactions := []Transaction{
    {1, 500, "2024-01-01", "credit"},
    {2, 200, "2024-01-02", "debit"},
    {3, 1500, "2024-01-03", "credit"},
    {4, 300, "2024-01-04", "debit"},
    {5, 100, "2024-01-05", "credit"},
}

// Use domain helpers
result := HighValueTransactions(CreditTransactions(transactions), 400.0)
plygo.From(result).Show()
```

::: tip Result
```
+----+---------+------------+--------+
| ID |  Amount |       Date |   Type |
+----+---------+------------+--------+
|  3 | 1500.00 | 2024-01-03 | credit |
|  1 |  500.00 | 2024-01-01 | credit |
+----+---------+------------+--------+
[2 rows × 4 columns]
```
:::

## Validation Helpers

Build helpers for data quality checks:

```go
type Record struct {
    ID    int
    Value float64
    Valid bool
}

// Validation helper
func ValidRecords(records []Record) []Record {
    return plygo.From(records).
        Where("Valid").IsTrue().
        Collect()
}

// Non-zero values
func NonZeroValues(records []Record) []Record {
    return plygo.From(records).
        Where("Value").GreaterThan(0.0).
        Collect()
}

// Quality check helper
func QualityCheck(records []Record) ([]Record, int, int) {
    valid := ValidRecords(records)
    nonZero := NonZeroValues(valid)
    return nonZero, len(valid), len(nonZero)
}

records := []Record{
    {1, 100, true},
    {2, 0, true},
    {3, 200, false},
    {4, 150, true},
    {5, 0, false},
}

clean, validCount, finalCount := QualityCheck(records)

fmt.Printf("Valid records: %d\n", validCount)
fmt.Printf("Non-zero records: %d\n", finalCount)
fmt.Println("\nClean data:")
plygo.From(clean).Show()
```

::: tip Result
```
Valid records: 3
Non-zero records: 2

Clean data:
+----+--------+-------+
| ID |  Value | Valid |
+----+--------+-------+
|  1 | 100.00 | true  |
|  4 | 150.00 | true  |
+----+--------+-------+
[2 rows × 3 columns]
```
:::

::: tip Helper Function Design Patterns
1. **Accept and return slices** - Makes helpers composable
2. **Single responsibility** - Each helper does one thing well
3. **Descriptive names** - Name after what they do, not how
4. **Add parameters** - Make helpers flexible with configuration
5. **Return multiple values** - Include metadata when useful (counts, errors, etc.)
6. **Package by domain** - Group related helpers together
:::

::: tip Real-World Use Cases
- **E-commerce**: `ActiveProducts()`, `InStockItems()`, `DiscountedProducts()`
- **Analytics**: `DailyMetrics()`, `TopPerformers()`, `GrowthRate()`
- **Finance**: `CreditTransactions()`, `HighValueOrders()`, `MonthlyTotal()`
- **Logging**: `ErrorLogs()`, `RecentEvents()`, `CriticalAlerts()`
- **User Management**: `ActiveUsers()`, `PremiumAccounts()`, `RecentSignups()`
:::

Next: [Performance Optimization](/advanced/performance)
