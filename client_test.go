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

func TestSearch(t *testing.T) {
	res, err := GetSearch(context.Background(), "dev", "perro")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Hits: %d", len(res))
	for _, r := range res {
		t.Log(r.Doc.Word)
		t.Log(r.WordEntry())
	}
}
