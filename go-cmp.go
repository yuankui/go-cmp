package main

import (
	"os"
	"bufio"
	"fmt"
	"github.com/yuankui/go-cmp/worker"
)

/**
TODO
1. 完成请求
2. 完成数据对比
3. 完成性能数据沉淀
 */

func main() {
	urls := os.Args[1:]
	reader := bufio.NewReader(os.Stdin)

	group := worker.NewWorkerGroup(10, urls, )
	for {
		line, _, err := reader.ReadLine()

		if err != nil {
			println(err)
			os.Exit(1)
		}

		group.AddTask(line)
	}

	for line := range group.OutChan {
		fmt.Println(line)
	}
}
