/*
chapter6 is my practice pkg
*/
package chapter6

// Average returns the average
func Average(s []int) int {
	total := 0
	for _, i := range s {
		total += i
	}
	return (total / len(s))
}
