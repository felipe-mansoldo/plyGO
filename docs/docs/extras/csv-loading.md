---
sidebar_position: 1
---

# CSV File Loading

Learn simple and practical patterns for loading CSV files and processing them with plyGO.

## Basic CSV Loading

The simplest way to load a CSV file into a struct slice:

```go
type Employee struct {
    Name   string
    Age    int
    City   string
    Salary float64
}

func loadCSV(filename string) ([]Employee, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }
    
    var employees []Employee
    for i, record := range records {
        if i == 0 {
            continue  // Skip header
        }
        
        age, _ := strconv.Atoi(record[1])
        salary, _ := strconv.ParseFloat(record[3], 64)
        
        employees = append(employees, Employee{
            Name:   record[0],
            Age:    age,
            City:   record[2],
            Salary: salary,
        })
    }
    
    return employees, nil
}

employees, err := loadCSV("employees.csv")
if err != nil {
    fmt.Println("Error:", err)
    return
}

fmt.Printf("Loaded %d employees\n", len(employees))
plygo.From(employees).Show()
```

**employees.csv:**
```csv
Name,Age,City,Salary
Alice,30,New York,75000
Bob,25,London,60000
Charlie,35,Tokyo,90000
Diana,28,Paris,70000
Eve,32,Berlin,80000
```

:::success Result
```
Loaded 5 employees
+---------+-----+----------+----------+
|    Name | Age |     City |   Salary |
+---------+-----+----------+----------+
| Alice   |  30 | New York | 75000.00 |
| Bob     |  25 | London   | 60000.00 |
| Charlie |  35 | Tokyo    | 90000.00 |
| Diana   |  28 | Paris    | 70000.00 |
| Eve     |  32 | Berlin   | 80000.00 |
+---------+-----+----------+----------+
[5 rows × 4 columns]
```
:::

## Filter While Loading

For large files, filter data as you read to save memory:

```go
func loadAndFilter(filename string, minSalary float64) ([]Employee, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    reader := csv.NewReader(file)
    
    // Skip header
    _, err = reader.Read()
    if err != nil {
        return nil, err
    }
    
    var employees []Employee
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            return nil, err
        }
        
        age, _ := strconv.Atoi(record[1])
        salary, _ := strconv.ParseFloat(record[3], 64)
        
        // Filter while loading - only keep high salaries
        if salary >= minSalary {
            employees = append(employees, Employee{
                Name:   record[0],
                Age:    age,
                City:   record[2],
                Salary: salary,
            })
        }
    }
    
    return employees, nil
}

employees, _ := loadAndFilter("employees.csv", 70000)

fmt.Printf("Loaded %d high-salary employees\n", len(employees))
plygo.From(employees).Show()
```

:::success Result
```
Loaded 4 high-salary employees
+---------+-----+----------+----------+
|    Name | Age |     City |   Salary |
+---------+-----+----------+----------+
| Alice   |  30 | New York | 75000.00 |
| Charlie |  35 | Tokyo    | 90000.00 |
| Diana   |  28 | Paris    | 70000.00 |
| Eve     |  32 | Berlin   | 80000.00 |
+---------+-----+----------+----------+
[4 rows × 4 columns]
```
:::

## Load and Process with plyGO

Load all data, then use plyGO for complex filtering and sorting:

```go
type Product struct {
    Product  string
    Price    float64
    Stock    int
    Category string
}

func loadProducts(filename string) ([]Product, error) {
    file, _ := os.Open(filename)
    defer file.Close()
    
    reader := csv.NewReader(file)
    records, _ := reader.ReadAll()
    
    var products []Product
    for i, rec := range records {
        if i == 0 { continue }
        
        price, _ := strconv.ParseFloat(rec[1], 64)
        stock, _ := strconv.Atoi(rec[2])
        
        products = append(products, Product{
            Product: rec[0], Price: price, 
            Stock: stock, Category: rec[3],
        })
    }
    
    return products, nil
}

products, _ := loadProducts("products.csv")

fmt.Println("Electronics in stock:")
plygo.From(products).
    Where("Category").Equals("Electronics").
    Where("Stock").GreaterThan(15).
    OrderBy("Price").Desc().
    Show()
```

**products.csv:**
```csv
Product,Price,Stock,Category
Laptop,999.99,15,Electronics
Mouse,29.99,50,Electronics
Desk,299.99,10,Furniture
Chair,199.99,25,Furniture
Monitor,349.99,20,Electronics
```

:::success Result
```
Electronics in stock:
+---------+--------+-------+-------------+
| Product |  Price | Stock |    Category |
+---------+--------+-------+-------------+
| Monitor | 349.99 |    20 | Electronics |
| Mouse   |  29.99 |    50 | Electronics |
+---------+--------+-------+-------------+
[2 rows × 4 columns]
```
:::

## Handle Errors and Missing Data

Add validation to handle malformed CSV files:

```go
type Employee struct {
    Name   string
    Age    int
    City   string
    Salary float64
    Valid  bool
}

func loadWithValidation(filename string) ([]Employee, error) {
    file, _ := os.Open(filename)
    defer file.Close()
    
    reader := csv.NewReader(file)
    records, _ := reader.ReadAll()
    
    var employees []Employee
    for i, rec := range records {
        if i == 0 { continue }
        
        // Validate record length
        if len(rec) < 4 {
            employees = append(employees, Employee{Valid: false})
            continue
        }
        
        age, err1 := strconv.Atoi(rec[1])
        salary, err2 := strconv.ParseFloat(rec[3], 64)
        
        // Mark as invalid if parsing fails
        valid := err1 == nil && err2 == nil
        
        employees = append(employees, Employee{
            Name:   rec[0],
            Age:    age,
            City:   rec[2],
            Salary: salary,
            Valid:  valid,
        })
    }
    
    return employees, nil
}

employees, _ := loadWithValidation("employees.csv")

// Show only valid records
valid := plygo.From(employees).
    Where("Valid").IsTrue().
    Collect()

fmt.Printf("Valid records: %d of %d\n", len(valid), len(employees))
plygo.From(valid).Select("Name", "Age", "City", "Salary").Show()
```

:::success Result
```
Valid records: 5 of 5
+---------+-----+----------+----------+
|    Name | Age |     City |   Salary |
+---------+-----+----------+----------+
| Alice   |  30 | New York | 75000.00 |
| Bob     |  25 | London   | 60000.00 |
| Charlie |  35 | Tokyo    | 90000.00 |
| Diana   |  28 | Paris    | 70000.00 |
| Eve     |  32 | Berlin   | 80000.00 |
+---------+-----+----------+----------+
[5 rows × 4 columns]
```
:::

## Reusable CSV Loader

Create a generic CSV loader for any struct type:

```go
// Generic CSV loader with custom parser
func LoadCSV[T any](filename string, parser func([]string) T) ([]T, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }
    
    var items []T
    for i, record := range records {
        if i == 0 {
            continue  // Skip header
        }
        items = append(items, parser(record))
    }
    
    return items, nil
}

// Use the generic loader
parseEmployee := func(rec []string) Employee {
    age, _ := strconv.Atoi(rec[1])
    salary, _ := strconv.ParseFloat(rec[3], 64)
    return Employee{
        Name: rec[0], Age: age, City: rec[2], Salary: salary,
    }
}

employees, _ := LoadCSV("employees.csv", parseEmployee)

fmt.Printf("Loaded %d employees\n", len(employees))

// Sort and show top 3
sorted := plygo.From(employees).
    OrderBy("Salary").Desc().
    Collect()

top3 := sorted[:3]
plygo.From(top3).Show()
```

:::success Result
```
Loaded 5 employees
+---------+-----+----------+----------+
|    Name | Age |     City |   Salary |
+---------+-----+----------+----------+
| Charlie |  35 | Tokyo    | 90000.00 |
| Eve     |  32 | Berlin   | 80000.00 |
| Alice   |  30 | New York | 75000.00 |
+---------+-----+----------+----------+
[3 rows × 4 columns]
```
:::

:::tip CSV Loading Best Practices

**File Handling:**
1. **Always close files** - Use `defer file.Close()`
2. **Check errors** - Handle file open and read errors
3. **Skip headers** - Start iteration from index 1 or read first line separately
4. **Use buffering** - `csv.NewReader` is already buffered

**Data Parsing:**
5. **Validate field count** - Check `len(record)` before accessing
6. **Handle parse errors** - Check errors from `strconv` functions
7. **Use appropriate types** - Match CSV data to struct field types
8. **Trim whitespace** - Use `strings.TrimSpace()` if needed

**Performance:**
9. **Pre-allocate slices** - Use `make([]T, 0, estimatedSize)` when size is known
10. **Filter while loading** - For large files, filter during read
11. **Stream for huge files** - Use `reader.Read()` loop instead of `ReadAll()`
12. **Consider chunking** - Process in batches for very large files
:::

:::tip Common CSV Patterns

**Simple load and process:**
```go
data, _ := loadCSV("file.csv")
result := plygo.From(data).Where(...).Collect()
```

**Stream and filter:**
```go
data, _ := loadAndFilter("file.csv", filterFunc)
plygo.From(data).Show()
```

**Load, validate, and clean:**
```go
data, _ := loadWithValidation("file.csv")
clean := plygo.From(data).Where("Valid").IsTrue().Collect()
```

**Generic loader:**
```go
data, _ := LoadCSV("file.csv", parserFunc)
plygo.From(data).OrderBy(...).Show()
```
:::

Next: [Real-World Examples](real-world-examples.md)
