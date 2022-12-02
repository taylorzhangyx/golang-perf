package main

import (
	"flag"
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"time"

	"taylorzhangyx.com/golang-perf/logic"
)

type Func func(interface{})

func main() {
	var p int64 = 0
	flag.Int64Var(&p, "loop", 10000000, "use general loop to test performance")

	funcs := []Func{
		logic.ListLoop,
		logic.ListLoopWithSize,
		logic.MapLoop,
		logic.MapLoopWithSize,
	}
	startTs := time.Now()
	wg := sync.WaitGroup{}
	for _, f := range funcs {
		wg.Add(1)
		go func(f Func) {
			defer wg.Done()
			MeasureTime(f, p)
		}(f)
	}
	wg.Wait()
	println("finished all time cost %s", time.Since(startTs).String())
}

func MeasureTime(f Func, p interface{}) {
	startTs := time.Now()
	defer func() {
		endTs := time.Since(startTs)
		funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		println(fmt.Sprintf("%s time cost: %v", funcName, endTs.String()))
	}()
	f(p)
}
