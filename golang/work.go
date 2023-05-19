package golang

import (
	"context"
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func Run(ctx context.Context) error {
	a := mat.NewDense(3, 3, []float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	})
	b := mat.NewVecDense(3, []float64{1, 2, 3})

	res := mat.NewVecDense(3, nil)
	res.MulVec(a, b)

	fmt.Printf("calculating .....n")
	fmt.Printf("A = \n%.2f\n", mat.Formatted(a))
	fmt.Printf("b = \n%.2f\n", mat.Formatted(b))
	fmt.Printf("A*b = \n%.2f\n", mat.Formatted(res))

	<-ctx.Done()
	fmt.Printf("stopping work ...\n")

	return ctx.Err()
}
