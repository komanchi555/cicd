package main

import "testing"

func TestAddA(t *testing.T) {
 tests := []struct {
  input  string
  output string
 }{
  {"hello", "helloa"},
  {"", "a"},
  {"a", "aa"},
 }

 for _, test := range tests {
  t.Run("testing "+test.input, func(t *testing.T) {
   result := addA(test.input)
   if result != test.output {
    t.Errorf("Expected %s, got %s", test.output, result)
   }
  })
 }
}
