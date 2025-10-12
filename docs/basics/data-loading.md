# Data Loading

Learn how to load data into plyGO pipelines.

## From Slices

```go
numbers := []int{1, 2, 3, 4, 5}
plygo.From(numbers).Show()
```

::: tip Result
```
+-------+
| Value |
+-------+
|     1 |
|     2 |
|     3 |
|     4 |
|     5 |
+-------+
[5 rows × 1 columns]
```
:::

## From Structs

```go
type Person struct {
    Name string
    Age  int
}

people := []Person{
    {"Alice", 30},
    {"Bob", 25},
}

plygo.From(people).Show()
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

Next: [Filtering](/basics/filtering)
