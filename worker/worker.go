package worker

import (
	"net/http"
	"bytes"
	"io"
	"encoding/json"
)

type WorkerGroup struct {
	urls    []string
	lines   chan []byte
	OutChan chan string
}

func NewWorkerGroup(num int, urls  []string) *WorkerGroup {
	workers := WorkerGroup{urls:urls, lines:make(chan []byte, 100)}
	workers.initWorkers(num)
	return &workers
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
	for str := range this.lines {
		this.processLine(str)
	}
}

func (this*WorkerGroup) processLine(line []byte) string {
	ret := make([]*Result, len(this.urls))
	for _, url := range this.urls {
		resp, _ := http.Post(url, "application/json", bytes.NewBufferString(string(line)))
		body := make([]byte, resp.ContentLength)
		io.ReadFull(resp.Body, body)
		res := NewResult(url, resp.StatusCode, string(body))
		res.url = url
		ret = append(ret, res)
	}
	bytes, _ := json.Marshal(ret)
	return string(bytes)
}

type Result struct {
	url     string
	status  int
	content string
}

func NewResult(url string, status int, content string) *Result {
	return &Result{url:url, status:status,content:content}
}