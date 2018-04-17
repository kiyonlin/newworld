package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

// 网站编码识别转换
func determineEncoder(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

const regexpstr = `<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+"[^>]*>[^<]+</a>`

const regexpstr1 = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func printCityList(contents []byte) {
	re := regexp.MustCompile(regexpstr1)
	amtch := re.FindAllSubmatch(contents, -1)

	for _, m := range amtch {
		fmt.Printf("%s\n %s \n", m[1], m[2])
	}
}

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code ", resp.StatusCode)
		return
	}

	e := determineEncoder(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	result, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	printCityList(result)
}
