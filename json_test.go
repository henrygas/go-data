package main

import (
	"io"
	"testing"
	"time"
)

func TestJsonDecoder(t *testing.T) {
	post, err := JsonDecoder()
	if err != nil && err != io.EOF {
		t.Error(err)
	}
	if post.Id != 1 {
		t.Error("Wrong id, was expecting 1 but got", post.Id)
	}
	if post.Content != "Hello World" {
		t.Error("Wrong content, was expecting 'Hello World', but got", post.Content)
	}
}

func BenchmarkJsonDecoder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = JsonDecoder()
	}
}

func TestParallel_1(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second)
}

func TestParallel_2(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
}

func TestParallel_3(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}