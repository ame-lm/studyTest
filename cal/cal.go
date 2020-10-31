package cal

type IntCalculatorImpl struct{}

type IntCalculator interface {
	Add(int, int) int
	Sub(int, int) int
}
