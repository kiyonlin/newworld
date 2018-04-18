package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var rageLimiter = time.Tick(10 * time.Millisecond)

// Fetch get url contents
func Fetch(url string) ([]byte, error) {
	<-rageLimiter
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: status code %d", resp.StatusCode)
	}
	reader := bufio.NewReader(resp.Body)
	e := determineEncoder(reader)
	urf8Reader := transform.NewReader(reader, e.NewDecoder())
	return ioutil.ReadAll(urf8Reader)
}

// 网站编码识别转换
func determineEncoder(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error :%v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
