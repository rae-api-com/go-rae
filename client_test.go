package rae

import (
	"context"
	"testing"
)

func TestRandom(t *testing.T) {
	t.Skip()

	t.Log(GetRandom(context.Background(), "dev"))
}

func TestDaily(t *testing.T) {
	t.Skip()

	t.Log(GetDaily(context.Background(), "dev"))
}
