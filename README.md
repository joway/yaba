# YaBa

![GitHub release](https://img.shields.io/github/tag/joway/yaba.svg?label=release)
[![Go Report Card](https://goreportcard.com/badge/github.com/joway/yaba)](https://goreportcard.com/report/github.com/joway/yaba)
[![codecov](https://codecov.io/gh/joway/yaba/branch/master/graph/badge.svg)](https://codecov.io/gh/joway/yaba)
[![CircleCI](https://circleci.com/gh/joway/yaba.svg?style=shield)](https://circleci.com/gh/joway/yaba)

## Usage

```go
import "github.com/jowat/yaba"

seg := yaba.NewSegmenter()
seg.LoadStopWords("")
seg.LoadDictionary("")

text := "我来到了北京清华大学旁边的北京大学旁边的清华大学"
tokens := seg.Segment(text)

for _, token := range tokens {
    fmt.Println(token.Text)
}
```

output:

```
我
来到
了
北京
清华大学
旁边
的
北京大学
旁边
的
清华大学
```