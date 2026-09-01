package main

import (
	"fmt"
	"math"
)

func main() {
	// ==== CONSTANTS ==== \\
	// ==== Basic Mathmetical Functions ==== \\
	// p := math.Pi
	// e := math.E
	// ph := math.Phi

	// sqrt2 := math.Sqrt2
	// sqrtE := math.SqrtE
	// sqrtPi := math.SqrtPi
	// sqrtPhi := math.SqrtPhi
	// ln2 := math.Ln2
	// ln10 := math.Ln10

	// log2E := math.Log2E
	// log10E := math.Log10E

	// ==== Basic Mathmetical Functions ==== \\

	// ==== ABSOLUTE VALUE ==== \\
	// func Abs(x float64) float64
	fmt.Printf("\n\n==== ABSOLUTE VALUE ====\n")
	fmt.Println(math.Abs(-3.14)) // will return 3.14

	// ==== SQUARE ROOT ==== \\
	// func Sqrt(x float64)
	fmt.Printf("\n\n==== SQUARE ROOT ====\n")
	fmt.Println(math.Sqrt(16)) // 4

	// ==== POWER ==== \\
	// func Powe(x, y float64) float64
	fmt.Printf("\n\n==== POWER ====\n")
	fmt.Println(math.Pow(2, 3)) // 2 * 2 * 2

	// ==== EXPONENTIAL ==== \\
	// func Exp(x float64) float64
	fmt.Printf("\n\n==== EXPONENTIAL ====\n")
	fmt.Println(math.Exp(1)) // 2.718281828459045

	// ==== LOGARITHMS ==== \\
	// func Log(x float64) float64 // Natural logarithm
	// func Log10(x float64) float64 // Base-10 logarithm
	fmt.Printf("\n\n==== LOGARITHMS ====\n")
	fmt.Println(math.Log(math.E)) // 1
	fmt.Println(math.Log10(100))  // 2

	// ==== TRIGONOMETRIC FUNCTIONS ==== \\
	// func Sin(x float64) float64
	// func Cos(x float64) float64
	// func Tan(x float64) float64
	fmt.Printf("\n\n==== TRIGONOMETRIC FUNCTIONS ==== \n")
	fmt.Println(math.Sin(math.Pi / 2)) // 1
	fmt.Println(math.Cos(math.Pi))     // -1
	fmt.Println(math.Tan(math.Pi / 4)) // 1

	// ==== INVERSE TRIGONOMETRIC FUNCTIONS ==== \\
	// func Asin(x float64) float64
	// func Acos(x float64) float64
	// func Atan(x float64) float64
	// func Atan2(y, x float64) float64
	fmt.Printf("\n\n==== INVERSE TRIGONOMETRIC FUNCTIONS ==== \n")
	fmt.Println(math.Asin(1))     // 1.5707963267948966
	fmt.Println(math.Acos(0))     // 1.5707963267948966
	fmt.Println(math.Atan(1))     // 0.7853981633974483
	fmt.Println(math.Atan2(1, 1)) // 0.7853981633974483

	fmt.Printf("\n\n==== HYPERBOLIC FUNCTIONS ==== \n")
	// func Sinh(x float64) float64
	// func Cosh(x float64) float64
	// func Tanh(x float64) float64
	fmt.Println(math.Sinh(1)) // Output: 1.1752011936438014
	fmt.Println(math.Cosh(1)) // Output: 1.5430806348152437
	fmt.Println(math.Tanh(1)) // Output: 0.7615941559557649

	fmt.Printf("\n\n==== SPECIAL FUNCTIONS ==== \n")
	// func Gamma(x float64) float64
	fmt.Println(math.Gamma(0.5)) // Output: 1.772453850905516

	// func Erf(x float64) float64
	// func Erfc(x float64) float64
	fmt.Println(math.Erf(1))  // Output: 0.8427007929497149
	fmt.Println(math.Erfc(1)) // Output: 0.15729920705028513

	fmt.Printf("\n\n==== Rounding Functions ==== \n")

	// func Ceil(x float64) float64
	fmt.Println(math.Ceil(1.3)) // Output: 2

	// func Floor(x float64) float64
	fmt.Println(math.Floor(1.7)) // Output: 1

	// func Round(x float64) float64
	fmt.Println(math.Round(1.5)) // Output: 2

	fmt.Printf("\n\n==== Utility Functions ==== \n")
	// func Min(x, y float64) float64
	// func Max(x, y float64) float64

	fmt.Printf("\n\n==== Utility Min ==== \n")
	fmt.Println(math.Min(1, 2)) // Output: 1
	fmt.Printf("\n\n==== Utility Max ==== \n")
	fmt.Println(math.Max(1, 2)) // Output: 2

	fmt.Printf("\n\n==== Utility Hypot ==== \n")
	// func Hypot(p, q float64) float64
	fmt.Println(math.Hypot(3, 4)) // Output: 5

}
