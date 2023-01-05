package sort

import (
	"math/rand"
	"time"
)

// Quick is a very simple quick sort implimentation for integers. It sorts an array of type []int, the first parameter, with the second one which is a compare function.
func Quick(ar []int, cmp func(i, j int) int) {
	if len(ar) > 1 {
		rand.Seed(time.Now().UnixNano())
		last := 0
		piv := rand.Intn(len(ar))
		swap(&ar, 0, piv)
		for i := 1; i < len(ar); i++ {
			if cmp(ar[i], ar[0]) <= 0 {
				last++
				swap(&ar, i, last)
			}
		}
		swap(&ar, 0, last)
		Quick(ar[:last], cmp)
		Quick(ar[last:], cmp)
	}
}

func swap(ar *[]int, i, j int) {
	arr := (*ar)
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
