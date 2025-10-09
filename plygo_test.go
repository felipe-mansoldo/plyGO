package plygo

import (
	"testing"
)

type Person struct {
	Name   string
	Age    int
	City   string
	Salary float64
	Active bool
}

func TestBasicFiltering(t *testing.T) {
	people := []Person{
		{"Alice", 30, "NYC", 75000, true},
		{"Bob", 25, "LA", 60000, true},
		{"Charlie", 35, "NYC", 90000, false},
		{"Diana", 28, "Chicago", 70000, true},
		{"Eve", 32, "LA", 85000, true},
	}

	result := From(people).
		Where("Age").GreaterThan(30).
		Where("Active").IsTrue().
		Collect()

	if len(result) != 1 {
		t.Errorf("Expected 1 result, got %d", len(result))
	}
	if result[0].Name != "Eve" {
		t.Errorf("Expected Eve, got %s", result[0].Name)
	}
}

func TestAndConditions(t *testing.T) {
	people := []Person{
		{"Alice", 30, "NYC", 75000, true},
		{"Bob", 25, "LA", 60000, true},
		{"Charlie", 35, "NYC", 90000, false},
	}

	result := From(people).
		Where("Age").GreaterThan(30).And("Active").IsTrue().
		Collect()

	if len(result) != 0 {
		t.Errorf("Expected 0 results, got %d", len(result))
	}
}

func TestOrConditions(t *testing.T) {
	people := []Person{
		{"Alice", 30, "NYC", 75000, true},
		{"Bob", 25, "LA", 60000, true},
		{"Charlie", 35, "Chicago", 90000, false},
	}

	result := From(people).
		Where("City").Equals("NYC").Or("City").Equals("LA").
		Collect()

	if len(result) != 2 {
		t.Errorf("Expected 2 results, got %d", len(result))
	}
}

func TestOneOf(t *testing.T) {
	people := []Person{
		{"Alice", 30, "NYC", 75000, true},
		{"Bob", 25, "LA", 60000, true},
		{"Charlie", 35, "Chicago", 90000, false},
		{"Diana", 28, "Miami", 70000, true},
	}

	result := From(people).
		Where("City").OneOf("NYC", "LA", "Chicago").
		Collect()

	if len(result) != 3 {
		t.Errorf("Expected 3 results, got %d", len(result))
	}
}

func TestSelect(t *testing.T) {
	people := []Person{
		{"Alice", 30, "NYC", 75000, true},
		{"Bob", 25, "LA", 60000, true},
	}

	result := From(people).
		Where("Age").GreaterThan(25).
		Select("Name", "Salary").
		Collect()

	if len(result) != 1 {
		t.Errorf("Expected 1 result, got %d", len(result))
	}
	if result[0]["Name"] != "Alice" {
		t.Errorf("Expected Alice, got %v", result[0]["Name"])
	}
	if _, ok := result[0]["Age"]; ok {
		t.Error("Age should not be in selected fields")
	}
}

func TestOrderBy(t *testing.T) {
	people := []Person{
		{"Alice", 30, "NYC", 75000, true},
		{"Bob", 25, "LA", 60000, true},
		{"Charlie", 35, "Chicago", 90000, false},
	}

	result := From(people).
		OrderBy("Age").Desc().
		Collect()

	if result[0].Name != "Charlie" {
		t.Errorf("Expected Charlie first, got %s", result[0].Name)
	}
	if result[2].Name != "Bob" {
		t.Errorf("Expected Bob last, got %s", result[2].Name)
	}
}

func TestGroupBy(t *testing.T) {
	people := []Person{
		{"Alice", 30, "NYC", 75000, true},
		{"Bob", 25, "LA", 60000, true},
		{"Charlie", 35, "NYC", 90000, false},
		{"Diana", 28, "LA", 70000, true},
	}

	result := From(people).
		GroupBy("City").
		Count()

	if result["NYC"] != 2 {
		t.Errorf("Expected 2 people in NYC, got %d", result["NYC"])
	}
	if result["LA"] != 2 {
		t.Errorf("Expected 2 people in LA, got %d", result["LA"])
	}
}

func TestWhereSome(t *testing.T) {
	people := []Person{
		{"Alice", 30, "NYC", 75000, true},
		{"Bob", 25, "LA", 60000, true},
		{"Charlie", 35, "Chicago", 90000, false},
	}

	result := From(people).
		WhereSome(
			W[Person]("City").Equals("NYC"),
			W[Person]("City").Equals("LA"),
		).
		Collect()

	if len(result) != 2 {
		t.Errorf("Expected 2 results, got %d", len(result))
	}
}

func TestWhereEvery(t *testing.T) {
	people := []Person{
		{"Alice", 30, "NYC", 75000, true},
		{"Bob", 25, "LA", 60000, true},
		{"Charlie", 35, "NYC", 90000, false},
	}

	result := From(people).
		WhereEvery(
			W[Person]("Age").GreaterThan(25),
			W[Person]("Active").IsTrue(),
		).
		Collect()

	if len(result) != 1 {
		t.Errorf("Expected 1 result, got %d", len(result))
	}
	if result[0].Name != "Alice" {
		t.Errorf("Expected Alice, got %s", result[0].Name)
	}
}

func TestLimit(t *testing.T) {
	people := []Person{
		{"Alice", 30, "NYC", 75000, true},
		{"Bob", 25, "LA", 60000, true},
		{"Charlie", 35, "Chicago", 90000, false},
	}

	result := From(people).
		Limit(2).
		Collect()

	if len(result) != 2 {
		t.Errorf("Expected 2 results, got %d", len(result))
	}
}

func TestDistinct(t *testing.T) {
	people := []Person{
		{"Alice", 30, "NYC", 75000, true},
		{"Bob", 25, "LA", 60000, true},
		{"Charlie", 35, "NYC", 90000, false},
		{"Diana", 28, "LA", 70000, true},
	}

	result := From(people).
		Distinct("City").
		Collect()

	if len(result) != 2 {
		t.Errorf("Expected 2 distinct cities, got %d", len(result))
	}
}

func TestTransform(t *testing.T) {
	people := []Person{
		{"Alice", 30, "NYC", 75000, true},
		{"Bob", 25, "LA", 60000, true},
	}

	result := From(people).
		Transform(func(p Person) Person {
			p.Salary *= 1.1
			return p
		}).
		Collect()

	if result[0].Salary != 82500 {
		t.Errorf("Expected 82500, got %f", result[0].Salary)
	}
}

func TestComplexPipeline(t *testing.T) {
	people := []Person{
		{"Alice", 30, "NYC", 75000, true},
		{"Bob", 25, "LA", 60000, true},
		{"Charlie", 35, "NYC", 90000, false},
		{"Diana", 28, "Chicago", 70000, true},
		{"Eve", 32, "LA", 85000, true},
	}

	result := From(people).
		Where("Active").IsTrue().
		Where("Age").GreaterThan(25).
		Where("City").Equals("NYC").Or("City").Equals("LA").
		OrderBy("Salary").Desc().
		Select("Name", "City", "Salary").
		Collect()

	if len(result) != 2 {
		t.Errorf("Expected 2 results, got %d", len(result))
	}
}

func TestBetween(t *testing.T) {
	people := []Person{
		{"Alice", 30, "NYC", 75000, true},
		{"Bob", 25, "LA", 60000, true},
		{"Charlie", 35, "NYC", 90000, false},
	}

	result := From(people).
		Where("Age").Between(26, 32).
		Collect()

	if len(result) != 1 {
		t.Errorf("Expected 1 result, got %d", len(result))
	}
	if result[0].Name != "Alice" {
		t.Errorf("Expected Alice, got %s", result[0].Name)
	}
}

func TestStringOperations(t *testing.T) {
	people := []Person{
		{"Alice Smith", 30, "NYC", 75000, true},
		{"Bob Johnson", 25, "LA", 60000, true},
		{"Alice Brown", 35, "NYC", 90000, false},
	}

	result := From(people).
		Where("Name").StartsWith("Alice").
		Collect()

	if len(result) != 2 {
		t.Errorf("Expected 2 results, got %d", len(result))
	}

	result2 := From(people).
		Where("Name").Contains("Smith").
		Collect()

	if len(result2) != 1 {
		t.Errorf("Expected 1 result, got %d", len(result2))
	}
}
