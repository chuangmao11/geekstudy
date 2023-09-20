package main

func DeferClourseLoopV1() {
	for i := 0; i < 10; i++ {
		defer func() {
			println(i)
		}()

	}
}

func DeferClourseLoopV2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			println(val)
		}(i)

	}
}

func DeferClourseLoopV3() {
	for i := 0; i < 10; i++ {
		j := 1
		defer func() {
			println(j)
		}()

	}
}
