# sego 分词库 

```
//第一步: 载入词典
var segmenter sego.Segmenter
segmenter.LoadDictionary("github.com/huichen/sego/data/dictionary.txt")
	
//第二步： 分词
sego.SegmentsToString("需要分词的文字",false)  //普通模式
sego.SegmentsToString("需要分词的文字",true)  //搜索模式
```

## 