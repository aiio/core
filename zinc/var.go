package zinc

type searchType struct {
	AllDocuments string
	Wildcard     string
	Fuzzy        string
	Term         string
	DateRange    string
	MatchAll     string
	Match        string
	MatchPhrase  string
	MultiPhrase  string
	Prefix       string
	QueryString  string
}

var SearchType = searchType{
	AllDocuments: "alldocuments", // 所有文件
	Wildcard:     "wildcard",     // 通配符
	Fuzzy:        "fuzzy",        // 模糊
	Term:         "term",         // 学期
	DateRange:    "daterange",    // 日期范围
	MatchAll:     "matchall",     // 匹配全部
	Match:        "match",        // 匹配
	MatchPhrase:  "matchphrase",  // 匹配短语
	MultiPhrase:  "multiphrase",  // 多词组
	Prefix:       "prefix",       // 字首
	QueryString:  "querystring",  // 请求参数
}
