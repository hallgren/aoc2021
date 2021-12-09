package main

import (
	"testing"
)

func TestLowest(t *testing.T) {
	if !Lowest(Point{0, 0, 0}, []Point{Point{Value: 1}, Point{Value: 2}, Point{Value: 3}}) {
		t.Fatal("0 not lowest")
	}
	if Lowest(Point{Value: 1}, []Point{Point{Value: 1}, Point{Value: 2}, Point{Value: 3}}) {
		t.Fatal("1 lowest")
	}
	if Lowest(Point{Value: 4}, []Point{Point{Value: 1}, Point{Value: 2}, Point{Value: 3}}) {
		t.Fatal("4 lowest")
	}
}

func TestValidPoint(t *testing.T) {
	m := Matrix{}
	for i := 0; i <= 5; i++ {
		r := Row{}
		for j := 0; j <= 5; j++ {
			p := Point{i, j, 0}
			r.Points = append(r.Points, p)
		}
		m.Rows = append(m.Rows, r)
	}

	if !ValidPoint(Point{0, 0, 0}, m) {
		t.Fatal("0,0 should be valid")
	}
	if ValidPoint(Point{-1, 0, 0}, m) {
		t.Fatal("-1,0 should be valid")
	}
	if ValidPoint(Point{0, -1, 0}, m) {
		t.Fatal("0,-1 should be valid")
	}
	if !ValidPoint(Point{5, 0, 0}, m) {
		t.Fatal("5,0 should not be valid")
	}
	if !ValidPoint(Point{0, 5, 0}, m) {
		t.Fatal("0,5 should not be valid")
	}
}

func TestNearPoints(t *testing.T) {
	m := Matrix{}
	for i := 0; i <= 5; i++ {
		r := Row{}
		for j := 0; j <= 5; j++ {
			p := Point{i, j, 0}
			r.Points = append(r.Points, p)
		}
		m.Rows = append(m.Rows, r)
	}

	points := ValidPoints(Point{0, 0, 1}, m)
	if len(points) != 2 {
		t.Fatalf("expected 2 got %d", len(points))
	}
	points = ValidPoints(Point{5, 5, 1}, m)
	if len(points) != 2 {
		t.Fatalf("expected 2 got %d", len(points))
	}
	points = ValidPoints(Point{2, 2, 1}, m)
	if len(points) != 4 {
		t.Fatalf("expected 4 got %d", len(points))
	}
}

func TestMax(t *testing.T) {
	m := Matrix{}
	for i := 0; i < 5; i++ {
		r := Row{}
		for j := 0; j < 5; j++ {
			p := Point{i, j, 0}
			r.Points = append(r.Points, p)
		}
		m.Rows = append(m.Rows, r)
	}

	if m.MaxX() != 5 {
		t.Fatalf("expected 5 got %d", m.MaxX())
	}
	if m.MaxY() != 5 {
		t.Fatalf("expected 5 got %d", m.MaxX())
	}
}
