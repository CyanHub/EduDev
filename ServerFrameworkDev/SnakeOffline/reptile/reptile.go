package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// 设置代理URL
	proxyURL, err := url.Parse("http://127.0.0.1:7890")
	if err != nil {
		fmt.Println("解析代理URL失败:", err)
		return
	}

	// 创建一个自定义的Transport，并设置代理
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	// 创建一个自定义的Client，并设置Transport
	client := &http.Client{
		Transport: transport,
	}

	// 目标URL
	url := "https://en.wikipedia.org/wiki/Python_(programming_language)"

	// 发送HTTP GET请求
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	// 将响应内容转换为字符串
	html := string(body)

	// 使用goquery解析HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println("解析HTML失败:", err)
		return
	}

	// 提取标题
	title := doc.Find("#firstHeading").Text()
	fmt.Println("标题:", title)

	// 提取第一段内容
	firstParagraph := doc.Find("p").First().Text()
	fmt.Println("第一段内容:", firstParagraph)
}
