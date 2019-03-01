package main

import "testing"

func TestParseArg(t *testing.T) {
  args               := []string{"test1", "test2"}
  expectedArg        := "test1"
  defaultExpectedArg := "config/config.stg.json"

  // Correct arg position
  correctPos  := 0
  argVal      := parseArg(args, correctPos, defaultExpectedArg)
  if argVal != expectedArg {
     t.Errorf("Return value was incorrect, got: %s, want: %s.", argVal, expectedArg)
  }

  // Incorrect arg position
  incorrectPos := 2
  defaultVal := parseArg(args, incorrectPos, defaultExpectedArg)
  if defaultVal != defaultExpectedArg {
     t.Errorf("Return value was incorrect, got: %s, want: %s.", defaultVal, defaultExpectedArg)
  }
}
