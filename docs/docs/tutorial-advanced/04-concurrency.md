---
sidebar_position: 4
---

# Concurrency Patterns

Learn how to use Go's concurrency features with plyGO for parallel data processing.

## Parallel Processing with Goroutines

Process independent tasks in parallel, then analyze results with plyGO:

```go
type Task struct {
    ID     int
    Result float64
}

func processTask(id int) Task {
    time.Sleep(10 * time.Millisecond)  // Simulate work
    return Task{ID: id, Result: float64(id * 100)}
}

ids := make([]int, 20)
for i := 0; i < 20; i++ {
    ids[i] = i
}

// Sequential processing
start := time.Now()
sequential := make([]Task, len(ids))
for i, id := range ids {
    sequential[i] = processTask(id)
}
seqTime := time.Since(start)

// Parallel processing
start = time.Now()
results := make([]Task, len(ids))
var wg sync.WaitGroup
for i, id := range ids {
    wg.Add(1)
    go func(idx, taskID int) {
        defer wg.Done()
        results[idx] = processTask(taskID)
    }(i, id)
}
wg.Wait()
parTime := time.Since(start)

fmt.Printf("Sequential: %v\n", seqTime)
fmt.Printf("Parallel: %v\n", parTime)
fmt.Printf("Speedup: %.2fx\n", float64(seqTime)/float64(parTime))

// Process results with plyGO
filtered := plygo.From(results).
    Where("Result").GreaterThan(1000.0).
    OrderBy("Result").Desc().
    Collect()

top5 := filtered[:5]
plygo.From(top5).Show()
```

:::success Result
```
Sequential: 205.1ms
Parallel: 10.3ms
Speedup: 19.90x
+----+---------+
| ID |  Result |
+----+---------+
| 19 | 1900.00 |
| 18 | 1800.00 |
| 17 | 1700.00 |
| 16 | 1600.00 |
| 15 | 1500.00 |
+----+---------+
[5 rows × 2 columns]
```
:::

## Worker Pool Pattern

Use a worker pool for controlled concurrency:

```go
type Job struct {
    ID    int
    Value float64
}

type Result struct {
    JobID  int
    Output float64
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        time.Sleep(5 * time.Millisecond)  // Simulate work
        results <- Result{
            JobID:  job.ID,
            Output: job.Value * 2,
        }
    }
}

numJobs := 20
numWorkers := 4

jobs := make(chan Job, numJobs)
results := make(chan Result, numJobs)

// Start workers
var wg sync.WaitGroup
for w := 1; w <= numWorkers; w++ {
    wg.Add(1)
    go worker(w, jobs, results, &wg)
}

// Send jobs
start := time.Now()
for j := 1; j <= numJobs; j++ {
    jobs <- Job{ID: j, Value: float64(j * 10)}
}
close(jobs)

// Wait for workers
go func() {
    wg.Wait()
    close(results)
}()

// Collect results
var allResults []Result
for r := range results {
    allResults = append(allResults, r)
}
duration := time.Since(start)

fmt.Printf("Processed %d jobs with %d workers in %v\n", 
    numJobs, numWorkers, duration)

// Analyze with plyGO
sorted := plygo.From(allResults).
    OrderBy("Output").Desc().
    Collect()

top5 := sorted[:5]
fmt.Println("\nTop 5 results:")
plygo.From(top5).Show()
```

:::success Result
```
Processed 20 jobs with 4 workers in 26.0ms

Top 5 results:
+-------+--------+
| JobID | Output |
+-------+--------+
|    20 | 400.00 |
|    19 | 380.00 |
|    18 | 360.00 |
|    17 | 340.00 |
|    16 | 320.00 |
+-------+--------+
[5 rows × 2 columns]
```
:::

## Concurrent Data Collection

Fetch from multiple sources concurrently:

```go
type DataSource struct {
    Name string
    Data []float64
}

func fetchFromSource(name string, delay time.Duration) DataSource {
    time.Sleep(delay)  // Simulate network delay
    data := make([]float64, 10)
    for i := 0; i < 10; i++ {
        data[i] = float64(i * 100)
    }
    return DataSource{Name: name, Data: data}
}

sources := []struct {
    name  string
    delay time.Duration
}{
    {"Source-A", 50 * time.Millisecond},
    {"Source-B", 30 * time.Millisecond},
    {"Source-C", 40 * time.Millisecond},
}

start := time.Now()
var wg sync.WaitGroup
resultChan := make(chan DataSource, len(sources))

// Fetch concurrently
for _, src := range sources {
    wg.Add(1)
    go func(name string, delay time.Duration) {
        defer wg.Done()
        resultChan <- fetchFromSource(name, delay)
    }(src.name, src.delay)
}

// Wait and close
go func() {
    wg.Wait()
    close(resultChan)
}()

// Collect all results
var allData []DataSource
for ds := range resultChan {
    allData = append(allData, ds)
}
duration := time.Since(start)

fmt.Printf("Fetched from %d sources in %v\n", len(allData), duration)
fmt.Printf("Total data points: %d\n", len(allData)*10)

// Process with plyGO
plygo.From(allData).Show()
```

:::success Result
```
Fetched from 3 sources in 50.6ms
Total data points: 30
+----------+--------------------------------+
|     Name |                           Data |
+----------+--------------------------------+
| Source-B | [0 100 200 300 400 500 600 ... |
| Source-C | [0 100 200 300 400 500 600 ... |
| Source-A | [0 100 200 300 400 500 600 ... |
+----------+--------------------------------+
[3 rows × 2 columns]
```
:::

## Pipeline with Concurrent Stages

Create a concurrent pipeline with multiple stages:

```go
type Record struct {
    ID        int
    Processed bool
    Value     float64
}

func generateRecords(n int) <-chan Record {
    out := make(chan Record)
    go func() {
        for i := 0; i < n; i++ {
            out <- Record{ID: i, Value: float64(i)}
            time.Sleep(2 * time.Millisecond)
        }
        close(out)
    }()
    return out
}

func processRecords(in <-chan Record) <-chan Record {
    out := make(chan Record)
    go func() {
        for r := range in {
            r.Processed = true
            r.Value = r.Value * 2
            out <- r
        }
        close(out)
    }()
    return out
}

start := time.Now()

// Create pipeline
records := generateRecords(10)
processed := processRecords(records)

// Collect results
var results []Record
for r := range processed {
    results = append(results, r)
}

duration := time.Since(start)

fmt.Printf("Processed %d records in %v\n", len(results), duration)

// Analyze with plyGO
plygo.From(results).
    Where("Value").GreaterThan(10.0).
    OrderBy("Value").Desc().
    Show()
```

:::success Result
```
Processed 10 records in 22.2ms
+----+-----------+-------+
| ID | Processed | Value |
+----+-----------+-------+
|  9 | true      | 18.00 |
|  8 | true      | 16.00 |
|  7 | true      | 14.00 |
|  6 | true      | 12.00 |
+----+-----------+-------+
[4 rows × 3 columns]
```
:::

:::tip Concurrency Best Practices

**When to Use Concurrency:**
1. **I/O-bound operations** - Network calls, file reads, database queries
2. **CPU-intensive tasks** - Independent calculations, data processing
3. **Multiple data sources** - Parallel fetching from different APIs/databases
4. **Batch processing** - Processing many independent items

**Pattern Selection:**
- **Simple parallelism**: Use goroutines + WaitGroup for independent tasks
- **Limited resources**: Use worker pool to control max concurrent operations
- **Streaming**: Use channels for pipeline-style processing
- **Fan-out/Fan-in**: Combine multiple results from parallel workers

**plyGO Integration:**
- Process concurrent results after collection
- Use plyGO for filtering, sorting, and aggregating parallel results
- Keep plyGO operations sequential for thread safety
:::

:::warning Concurrency Safety

**Thread Safety:**
- plyGO pipelines are **not thread-safe** - don't share them across goroutines
- **Collect results first**, then process with plyGO
- Use channels or mutexes for goroutine synchronization

**❌ Don't do this:**
```go
// BAD: Sharing pipeline across goroutines
pipe := plygo.From(data)
go func() { pipe.Where(...) }()  // UNSAFE!
go func() { pipe.OrderBy(...) }() // UNSAFE!
```

**✅ Do this instead:**
```go
// GOOD: Collect concurrent results, then process
var results []Item
var mu sync.Mutex
for _, item := range items {
    go func(i Item) {
        processed := process(i)
        mu.Lock()
        results = append(results, processed)
        mu.Unlock()
    }(item)
}
// Then use plyGO on collected results
plygo.From(results).Where(...).Show()
```
:::

Next: [Large Data Handling](05-large-data.md)
