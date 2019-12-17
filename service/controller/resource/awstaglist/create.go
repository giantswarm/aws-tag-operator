package test

import (
	"context"
	"fmt"
)

func (r *Resource) EnsureCreated(ctx context.Context, obj interface{}) error {
	fmt.Println("YOLO")
	return nil
}
