package main

import (
	"testing"
)

func TestPositiveSum(t *testing.T){
	want:=40
	a:=30
	b:=10
	result := sum(a,b)
	if (result != want) {
		t.Fatalf(`Sum must be %v, but got %v`, want, result)
	}
}

func TestNegativeSum(t *testing.T){
	want:=-40
	a:=-30
	b:=-10
	result := sum(a,b)
	if (result != want) {
		t.Fatalf(`Sum must be %v, but got %v`, want, result)
	}
}

func TestMistSum(t *testing.T){
	want:=-20
	a:=-30
	b:=10
	result := sum(a,b)
	if (result != want) {
		t.Fatalf(`Sum must be %v, but got %v`, want, result)
	}
}

func TestInvertMistSum(t *testing.T){
	want:=20
	a:=30
	b:=-10
	result := sum(a,b)
	if (result != want) {
		t.Fatalf(`Sum must be %v, but got %v`, want, result)
	}
}