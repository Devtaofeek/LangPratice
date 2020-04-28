package main

func sum(a[]int, c chan int)  {

	total := 0

	for v,_ := range a{
		total += v
	}
	c <- total
}