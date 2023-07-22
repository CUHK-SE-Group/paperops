package main

import (
	"encoding/xml"
	"fmt"
	"github.com/go-co-op/gocron"
	"io"
	"net/http"
	"time"
)

// 定义返回数据的结构体
type Entry struct {
	Title     string `xml:"title" json:"title"`
	Published string `xml:"published" json:"published"`
	Summary   string `xml:"summary" json:"summary"`
}

type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Entries []Entry  `xml:"entry"`
}

func fetchPaper() {
	// 设置请求URL
	url := "http://export.arxiv.org/api/query?search_query=cat:cs.*&sortBy=submittedDate&sortOrder=descending&max_results=10"

	// 发起HTTP请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP request error:", err)
		return
	}
	defer resp.Body.Close()

	// 读取返回的数据
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read response body error:", err)
		return
	}

	// 解析XML数据
	var feed Feed
	err = xml.Unmarshal(data, &feed)
	if err != nil {
		fmt.Println("XML unmarshal error:", err)
		return
	}
	cli := NewProducer([]string{"localhost:19092"}, "paper-trend")
	for _, entry := range feed.Entries {
		fmt.Println(entry)
		err := cli.SendMessage(entry)
		if err != nil {
			fmt.Println("Failed to produce record:", err)
			return
		}
	}
	cli.Flush()
	fmt.Println("Messages sent to Kafka successfully.")
}

func main() {
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(1).Day().At("00:00").Do(fetchPaper)
	if err != nil {
		fmt.Printf("Error creating scheduler: %v\n", err)
		return
	}

	s.StartAsync()
	stop := make(chan struct{})
	<-stop
}
