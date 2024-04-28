package main

import (
	"context"
	"log"
	"testing"
)

func TestDel(t *testing.T) {
	r := NewRedis()
	ctx := context.Background()
	err := r.Del(ctx, "bucunzi").Err()
	if err != nil {
		log.Fatalln(err)
	}
}
