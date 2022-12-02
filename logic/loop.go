package logic

func ListLoop(pCycle interface{}) {
	cycle := pCycle.(int64)
	l := make([]int64, 0, 0)
	for i := int64(0); i < cycle; i++ {
		l = append(l, i)
	}
	return
}

func ListLoopWithSize(pCycle interface{}) {
	cycle := pCycle.(int64)
	l := make([]int64, 0, cycle)
	for i := int64(0); i < cycle; i++ {
		l = append(l, i)
	}
	return
}

func MapLoop(pCycle interface{}) {
	cycle := pCycle.(int64)
	m := make(map[int64]bool)
	for i := int64(0); i < cycle; i++ {
		m[i] = true
	}
	return
}

func MapLoopWithSize(pCycle interface{}) {
	cycle := pCycle.(int64)
	m := make(map[int64]bool, cycle)
	for i := int64(0); i < cycle; i++ {
		m[i] = true
	}
	return
}
