# sego 分词库  go get github.com/huichen/sego

```
/* 
 *  分词
 */ 
func initSego()  {
	//第一步: 载入词典
	var segmenter sego.Segmenter
	segmenter.LoadDictionary("./dictionary/dictionary.txt")

	//第二步： 分词
	sego.SegmentsToString([]byte("普通模式分词"),false)  //普通模式
	sego.SegmentsToString([]byte("普通模式分词"),true)  //搜索模式
}

```

## ElasticSearch 搜索


```
/* 
 * 全文搜索
 */ 
 type ElasticSearchData struct{
 
 

 }



```

