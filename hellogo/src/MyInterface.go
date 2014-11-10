package main
import (
	"fmt"
	"math"
	"strings"
	"io"
)
type Abser interface {
	Abs() float64
}
func main(){
	fmt.Println("hell")
	var a Abser
	f := MyFloat(-12345)
	v := Vertex{3, 4}

	a = f  // a MyFloat 实现了 Abser
	fmt.Println(a.Abs())

	a = &v // a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	//a = v
	fmt.Println(a.Abs())


	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	for  {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

