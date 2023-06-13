package main

import (
	"reflect"
	"testing"
)

func TestRadix16Tree_InsertAndSearch(t *testing.T) {
	t.Run("Insert and search a value", func(t *testing.T) {
		tree := &Radix16Tree{root: NewRadix16Node()}
		tree.Insert("apple", 5)
		value, found := tree.Search("apple")
		if !found {
			t.Error("Expected apple to be found in the tree")
		}
		if value != 5 {
			t.Errorf("Expected value of 5 to be found for key apple, got %v", value)
		}
	})

	t.Run("Search a non-existent value", func(t *testing.T) {
		tree := &Radix16Tree{root: NewRadix16Node()}
		value, found := tree.Search("banana")
		if found {
			t.Error("Expected banana to not be found in the tree")
		}
		if value != nil {
			t.Errorf("Expected nil value to be returned for non-existent key banana, got %v", value)
		}
	})
}

func TestNewRadix16Node(t *testing.T) {
	tests := []struct {
		name string
		want *Radix16Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRadix16Node(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRadix16Node() = %v, want %v", got, tt.want)
			}
		})
	}
}
