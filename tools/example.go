package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/lifeng1992/sego"
)

var (
	text = flag.String("text", "中国互联网历史上最大的一笔并购案", "要分词的文本")
)

func main() {
	flag.Parse()

	var seg sego.Segmenter
	b, err := ioutil.ReadFile("./dictionary.txt")
	if err != nil {
		log.Fatal(err)
	}
	seg.LoadDictionary(b)

	segments := seg.Segment([]byte(*text))
	fmt.Println(sego.SegmentsToString(segments, true))
}
