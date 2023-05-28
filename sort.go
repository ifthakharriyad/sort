package sort

import (
	"math"
	"math/rand"
	"time"

	"github.com/oleiade/lane"
)

// MergeSort is a simple merge sort implimentation for integers.
// It sorts an slice of type []int, the first parameter.
func MergeSort(s []int) {
	if len(s) > 1 {
		mid := int(math.Floor((float64(len(s)) / 2)))
		MergeSort(s[:mid])
		MergeSort(s[mid:])
		Merge(s, mid)
	}

}

// Merge function merges two slices into sorted order.
func Merge(s []int, mid int) {
	// Initializing two queues for containing two slices s1 and s2.
	q1 := lane.NewQueue()
	q2 := lane.NewQueue()
	// Enqueueing queues
	for _, item := range s[:mid] {
		q1.Enqueue(item)
	}
	for _, item := range s[mid:] {
		q2.Enqueue(item)
	}
	i := 0
	for !(q1.Empty() || q2.Empty()) {
		if q1.Head().(int) <= q2.Head().(int) {
			s[i] = q1.Dequeue().(int)
		} else {
			s[i] = q2.Dequeue().(int)
		}
		i++
	}
	for !q1.Empty() {
		s[i] = q1.Dequeue().(int)
		i++
	}
	for !q2.Empty() {
		s[i] = q2.Dequeue().(int)
		i++
	}
}

// QuickSort is a very simple quick sort implimentation for integers. It sorts an slice of type []int, the first parameter, with the second one which is a compare function.
func QuickSort(s []int, cmp func(i, j int) int) {
	if len(s) > 1 {
		rand.Seed(time.Now().UnixNano())
		last := 0
		piv := rand.Intn(len(s))
		swap(&s, 0, piv)
		for i := 1; i < len(s); i++ {
			if cmp(s[i], s[0]) <= 0 {
				last++
				swap(&s, i, last)
			}
		}
		swap(&s, 0, last)
		QuickSort(s[:last], cmp)
		QuickSort(s[last:], cmp)
	}
}

func swap(s *[]int, i, j int) {
	slc := *s
	temp := slc[i]
	slc[i] = slc[j]
	slc[j] = temp
}
