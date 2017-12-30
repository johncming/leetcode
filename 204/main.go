package leetcode

func countPrimes(n int) int {
	cnt := 0
	var primes []int
	for i := 1; i < n; i++ {
		if isPrime(i, primes) {
			primes = append(primes, i)
			cnt++
		}
	}
	return cnt
}

func isPrime(n int, primes []int) bool {
	if n < 2 {
		return false
	}

	for _, p := range primes {
		if p*p > n {
			break
		}
		if n%p == 0 {
			return false
		}
	}

	return true
}
