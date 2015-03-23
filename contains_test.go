package goproof

import (
  "testing"
)

////////////////////////////////////////////////////////////////
func TestContains(t *testing.T) {

  list := []string{"a", "b", "c"}

  if !contains(list, "b") {
    t.Error("list should contain b")
  }

  if contains(list, "d") {
    t.Error("list should NOT contain d")
  }

}
