# Output Display Style Reference

This document provides templates for all 8 output display styles tested in the documentation.
Copy and paste the one you prefer when adding examples to documentation pages.

---

## Option 1: Simple Code Block
**Best for:** Clean, minimal display without decoration

**Template:**
```markdown
```
[paste your output here]
```
```

**Example:**
```markdown
```
+---------+-----+----------+
|    Name | Age |   Salary |
+---------+-----+----------+
| Alice   |  30 | 75000.00 |
| Bob     |  25 | 60000.00 |
| Charlie |  35 | 90000.00 |
+---------+-----+----------+
[3 rows × 3 columns]
```
```

---

## Option 2: Highlighted Console Output
**Best for:** Terminal/console outputs, commands

**Template:**
```markdown
```console
[paste your output here]
```
```

**Example:**
```markdown
```console
+---------+-----+----------+
|    Name | Age |   Salary |
+---------+-----+----------+
| Alice   |  30 | 75000.00 |
| Bob     |  25 | 60000.00 |
| Charlie |  35 | 90000.00 |
+---------+-----+----------+
[3 rows × 3 columns]
```
```

---

## Option 3: Tip Admonition
**Best for:** Helpful outputs, tips, suggestions

**Template:**
```markdown
:::tip Output
```
[paste your output here]
```
:::
```

**Example:**
```markdown
:::tip Output
```
+---------+-----+----------+
|    Name | Age |   Salary |
+---------+-----+----------+
| Alice   |  30 | 75000.00 |
| Bob     |  25 | 60000.00 |
| Charlie |  35 | 90000.00 |
+---------+-----+----------+
[3 rows × 3 columns]
```
:::
```

---

## Option 4: Info Admonition ⭐ RECOMMENDED
**Best for:** General program outputs, results
**Why recommended:** Professional, clear, non-intrusive blue theme

**Template:**
```markdown
:::info Program Output
```
[paste your output here]
```
:::
```

**Example:**
```markdown
:::info Program Output
```
+---------+-----+----------+
|    Name | Age |   Salary |
+---------+-----+----------+
| Alice   |  30 | 75000.00 |
| Bob     |  25 | 60000.00 |
| Charlie |  35 | 90000.00 |
+---------+-----+----------+
[3 rows × 3 columns]
```
:::
```

---

## Option 5: Success Style
**Best for:** Successful operations, positive results

**Template:**
```markdown
:::success Result
```
[paste your output here]
```
:::
```

**Example:**
```markdown
:::success Result
```
+---------+-----+----------+
|    Name | Age |   Salary |
+---------+-----+----------+
| Alice   |  30 | 75000.00 |
| Bob     |  25 | 60000.00 |
| Charlie |  35 | 90000.00 |
+---------+-----+----------+
[3 rows × 3 columns]
```
:::
```

---

## Option 6: Collapsible Details
**Best for:** Long outputs, optional information

**Template:**
```markdown
<details>
<summary>Click to see output</summary>

```
[paste your output here]
```

</details>
```

**Example:**
```markdown
<details>
<summary>Click to see output</summary>

```
+---------+-----+----------+
|    Name | Age |   Salary |
+---------+-----+----------+
| Alice   |  30 | 75000.00 |
| Bob     |  25 | 60000.00 |
| Charlie |  35 | 90000.00 |
+---------+-----+----------+
[3 rows × 3 columns]
```

</details>
```

---

## Option 7: With Explanation Header
**Best for:** Simple, straightforward outputs with minimal decoration

**Template:**
```markdown
**Output:**
```
[paste your output here]
```
```

**Example:**
```markdown
**Output:**
```
+---------+-----+----------+
|    Name | Age |   Salary |
+---------+-----+----------+
| Alice   |  30 | 75000.00 |
| Bob     |  25 | 60000.00 |
| Charlie |  35 | 90000.00 |
+---------+-----+----------+
[3 rows × 3 columns]
```
```

---

## Option 8: Table in Blockquote
**Best for:** Quoted or referenced outputs

**Template:**
```markdown
> **Output:**
> ```
> [paste your output here]
> ```
```

**Example:**
```markdown
> **Output:**
> ```
> +---------+-----+----------+
> |    Name | Age |   Salary |
> +---------+-----+----------+
> | Alice   |  30 | 75000.00 |
> | Bob     |  25 | 60000.00 |
> | Charlie |  35 | 90000.00 |
> +---------+-----+----------+
> [3 rows × 3 columns]
> ```
```

---

## Quick Decision Guide

| Style | Use Case | Visual Impact |
|-------|----------|---------------|
| Option 1 - Simple | Default, clean | Minimal |
| Option 2 - Console | CLI outputs | Low |
| Option 3 - Tip | Helpful hints | Medium (green) |
| **Option 4 - Info** ⭐ | **General outputs** | **Medium (blue)** ⭐ |
| Option 5 - Success | Positive results | Medium (green) |
| Option 6 - Collapsible | Long/optional | Hidden by default |
| Option 7 - Header | Simple with label | Minimal |
| Option 8 - Blockquote | Referenced output | Low |

---

## Recommendation

**For consistent documentation:**
- Use **Option 4 (Info Admonition)** for most program outputs
- Use **Option 3 (Tip)** for important tips or best practices
- Use **Option 6 (Collapsible)** for very long outputs
- Use **Option 7 (Header)** for inline, quick examples

This creates a professional, consistent look throughout the documentation.

---

## Example Usage in Documentation

```markdown
## Example: Filtering Data

```go
result := plygo.From(people).
    Where("Age").GreaterThan(25).
    Collect()
```

:::info Program Output
```
+---------+-----+----------+
|    Name | Age |   Salary |
+---------+-----+----------+
| Alice   |  30 | 75000.00 |
| Charlie |  35 | 90000.00 |
+---------+-----+----------+
[2 rows × 3 columns]
```
:::
```

---

**All styles are tested and verified to work with Docusaurus 3.9.1!**
