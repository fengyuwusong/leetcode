package main

func fib(n int) int {
	memo := make([]int, n+1)
	return helper(memo, n)
}

func helper(memo []int, n int) int {
	if n == 0 || n == 1 {
		return n
	}
	if memo[n] != 0 {
		return memo[n]
	}

	memo[n] = helper(memo, n-1) + helper(memo, n-2)

	return memo[n]
}

func main() {
	println(fib(3))
}
