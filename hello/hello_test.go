package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {

		// Helper 声明辅助函数，测试失败报告行号只会出现在调用处
		t.Helper()

		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	//	t.Run() 对不同场景进行子测试
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Gopher", "")
		want := "Hello, Gopher"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Chinese", func(t *testing.T) {
		got := Hello("World", "Chinese")
		want := "你好, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("World", "Spanish")
		want := "Hola, World"
		assertCorrectMessage(t, got, want)
	})
}
