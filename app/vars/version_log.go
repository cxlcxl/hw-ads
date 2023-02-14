package vars

type V struct {
	Date    string   `json:"date"`
	Version string   `json:"v"`
	Fixs    []string `json:"fixs"`
}

var VersionInfo = []V{
	{
		"2023-02-14", "v1.0.1",
		[]string{
			"Fix: 综合报表收入字段分账户查看去重",
		},
	},
	{
		"2023-02-13", "v1.0.0",
		[]string{
			"1. 系统发布",
			"2. 包含综合、变现报表、应用数据、账户数据等",
		},
	},
}
