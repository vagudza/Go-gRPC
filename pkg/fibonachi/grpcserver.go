package fibonachi

import (
	"context"
	"fibonachi_gRPC/pkg/api"
	"fmt"
	"strings"
)

// GRPCServer ...
type GRPCServer struct{}

// Calculating a number from the Fibonacci sequence
func fibonachi() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}

// Returns a slice of a sequence of numbers from the Fibonacci series from x to y
func getFibonachiSlice(x, y int) []int {
	f := fibonachi()
	var result []int
	for i := 0; i <= y; i++ {
		value := f()
		if i >= x && i <= y {
			result = append(result, value)
		}
	}
	return result
}

// Returns Fibonacci slice in attribute "result"
func (s *GRPCServer) Fibonachi(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	x := int(req.GetX())
	y := int(req.GetY())
	slice := getFibonachiSlice(x, y)
	result := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slice)), ", "), "[]")
	//fmt.Println(x, y, slice, result)

	return &api.AddResponse{Result: string(result)}, nil
}
