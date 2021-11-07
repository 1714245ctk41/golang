package main
import("fmt")
func main(){
	ii := []int{1,2,3,3,45,6,76,7,8,8,3}
	s := sum(ii...)
	
	fmt.Println("all number ", s)
	
	s2 := even(sum, ii...)
	fmt.Println("even numbers", s2)
}


func sum( xi ...int) int{
	
	total := 0
	for _, v := range xi{
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int{
	var yi []int
	for _, v := range vi{
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}