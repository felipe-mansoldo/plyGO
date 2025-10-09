---
sidebar_position: 10
---

# Display Styles

Learn how to display your data in different formats using `Show()` options.

## Default Style

The basic `Show()` displays data in a formatted table:

```go
type Product struct {
    Name  string
    Price float64
    Stock int
}

products := []Product{
    {"Laptop", 999.99, 15},
    {"Mouse", 29.99, 50},
    {"Keyboard", 79.99, 30},
}

plygo.From(products).Show()
```

:::success Result
```
+----------+--------+-------+
|     Name |  Price | Stock |
+----------+--------+-------+
| Laptop   | 999.99 |    15 |
| Mouse    |  29.99 |    50 |
| Keyboard |  79.99 |    30 |
+----------+--------+-------+
[3 rows × 3 columns]
```
:::

## Markdown Style

Use `WithStyle("markdown")` for markdown-formatted tables:

```go
products := []Product{
    {"Laptop", 999.99, 15},
    {"Mouse", 29.99, 50},
    {"Keyboard", 79.99, 30},
}

plygo.From(products).Show(plygo.WithStyle("markdown"))
```

:::success Result
```
|----------|--------|-------|
|     Name |  Price | Stock |
|----------|--------|-------|
| Laptop   | 999.99 |    15 |
| Mouse    |  29.99 |    50 |
| Keyboard |  79.99 |    30 |
|----------|--------|-------|
[3 rows × 3 columns]
```
:::

## Custom Options

Combine multiple options for customized output:

```go
products := []Product{
    {"Laptop", 999.99, 15},
    {"Mouse", 29.99, 50},
    {"Keyboard", 79.99, 30},
}

plygo.From(products).Show(
    plygo.WithTitle("Product Inventory"),
    plygo.WithRowNumbers(true),
    plygo.WithFloatPrecision(2),
)
```

:::success Result
```

        Product Inventory
+---+----------+--------+-------+
| # |     Name |  Price | Stock |
+---+----------+--------+-------+
| 1 | Laptop   | 999.99 |    15 |
| 2 | Mouse    |  29.99 |    50 |
| 3 | Keyboard |  79.99 |    30 |
+---+----------+--------+-------+
[3 rows × 4 columns]
```
:::

## Available Options

Here are all available `Show()` options:

| Option | Description | Example |
|--------|-------------|---------|
| `WithStyle(style)` | Set table style | `"compact"`, `"markdown"`, `"simple"`, `"csv"` |
| `WithTitle(title)` | Add a title above the table | `"Sales Report"` |
| `WithRowNumbers(bool)` | Show row numbers | `true` or `false` |
| `WithOriginalIndices(bool)` | Show original indices | `true` or `false` |
| `WithFloatPrecision(n)` | Set decimal places for floats | `2` |
| `WithMaxRows(n)` | Limit displayed rows | `100` |
| `WithMaxColWidth(n)` | Limit column width | `30` |
| `WithMaxWidth(n)` | Limit total table width | `120` |

:::tip Multiple Options
You can combine multiple options in a single `Show()` call:
```go
.Show(
    plygo.WithTitle("Report"),
    plygo.WithStyle("markdown"),
    plygo.WithFloatPrecision(2),
    plygo.WithRowNumbers(true),
)
```
:::

## Style Examples

### Compact Style
```go
plygo.From(data).Show(plygo.WithStyle("compact"))
```

### Simple Style
```go
plygo.From(data).Show(plygo.WithStyle("simple"))
```

### CSV Style
```go
plygo.From(data).Show(plygo.WithStyle("csv"))
```

Next: [Error Handling](error-handling.md)
