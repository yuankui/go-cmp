package main

import (
	"os"
	"bufio"
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
	requestChan := make(chan string, 300)

	NewWorkerGroup(10, urls, )
	for {
		line, _, err := reader.ReadLine()

		if err != nil {
			println(err)
			os.Exit(1)
		}

		requestChan <- line
	}

}

type WorkerGroup struct {
	urls  []string
	lines chan []byte
}

func NewWorkerGroup(num int, urls chan string) {
	workers := WorkerGroup{urls:urls, lines:make(chan []byte, 100)}
	workers.initWorkers(num)
}

func (this*WorkerGroup) AddTask(line []byte) {
	this.lines <- line
}

func (this*WorkerGroup) initWorkers(n int) {
	for i := 0; i < n; i++ {
		go this.worker()
	}
}

func (this*WorkerGroup) worker() {
	for {
		for str := range this.lines {
			this.processLine(str)
		}
	}
}
func (this*WorkerGroup) processLine(line []byte) string {
	for url := range this.urls {
		println(url)
	}

	return ""
}
