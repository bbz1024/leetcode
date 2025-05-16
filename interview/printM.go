package main

import "fmt"

func printM(threadCnt, num int) {
	threads := make([]chan struct{}, threadCnt)
	finish := make(chan struct{})
	for i := 0; i < len(threads); i++ {
		threads[i] = make(chan struct{})
	}
	// 启动
	for i := 0; i < len(threads); i++ {
		go func(id int) {
			for j := id; j <= num; j += threadCnt {
				// 先进行阻塞，等待被唤醒
				<-threads[j%threadCnt]
				fmt.Printf("id %d, num %d\n", id, j)
				if j == num {
					finish <- struct{}{}
					return
				}
				// 唤醒下一个
				threads[(j+1)%threadCnt] <- struct{}{}
			}
		}(i)
	}
	// 进行唤醒
	threads[0] <- struct{}{}
	<-finish
}
