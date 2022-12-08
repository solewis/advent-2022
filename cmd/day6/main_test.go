package main

import "testing"

func TestCharsToFirstPacketMarker(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		input := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
		result := charsToFirstMarker(input, 4)
		if result != 7 {
			t.Errorf("Expected 7 but was %d", result)
		}
	})
	t.Run("example 2", func(t *testing.T) {
		input := "bvwbjplbgvbhsrlpgdmjqwftvncz"
		result := charsToFirstMarker(input, 4)
		if result != 5 {
			t.Errorf("Expected 5 but was %d", result)
		}
	})

	t.Run("example 3", func(t *testing.T) {
		input := "nppdvjthqldpwncqszvftbrmjlhg"
		result := charsToFirstMarker(input, 4)
		if result != 6 {
			t.Errorf("Expected 6 but was %d", result)
		}
	})

	t.Run("example 4", func(t *testing.T) {
		input := "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
		result := charsToFirstMarker(input, 4)
		if result != 10 {
			t.Errorf("Expected 10 but was %d", result)
		}
	})

	t.Run("example 5", func(t *testing.T) {
		input := "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"
		result := charsToFirstMarker(input, 4)
		if result != 11 {
			t.Errorf("Expected 11 but was %d", result)
		}
	})
}

func TestCharsToFirstMessageMarker(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		input := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
		result := charsToFirstMarker(input, 14)
		if result != 19 {
			t.Errorf("Expected 19 but was %d", result)
		}
	})

	t.Run("example 2", func(t *testing.T) {
		input := "bvwbjplbgvbhsrlpgdmjqwftvncz"
		result := charsToFirstMarker(input, 14)
		if result != 23 {
			t.Errorf("Expected 23 but was %d", result)
		}
	})

	t.Run("example 3", func(t *testing.T) {
		input := "nppdvjthqldpwncqszvftbrmjlhg"
		result := charsToFirstMarker(input, 14)
		if result != 23 {
			t.Errorf("Expected 23 but was %d", result)
		}
	})

	t.Run("example 4", func(t *testing.T) {
		input := "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
		result := charsToFirstMarker(input, 14)
		if result != 29 {
			t.Errorf("Expected 29 but was %d", result)
		}
	})

	t.Run("example 5", func(t *testing.T) {
		input := "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"
		result := charsToFirstMarker(input, 14)
		if result != 26 {
			t.Errorf("Expected 26 but was %d", result)
		}
	})
}
