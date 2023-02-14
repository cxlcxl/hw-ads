package vars

type V struct {
	Date    string   `json:"date"`
	Version string   `json:"v"`
	Fix     []string `json:"fixs"`
}

var VersionInfo = []V{
	{
		"2023-02-15", "v1.0.4",
		[]string{
			"1. 应用列表展示投放账户关联并添加筛选功能",
			"2. 账户添加/修改填写信息调整",
			"3. 账户添加生成相关 Token 数据",
			"4. 账户放开重新认证 Token 信息",
			"5. 报表应用筛选优化等功能更新...",
		},
	},
	{
		"2023-02-14", "v1.0.3",
		[]string{
			"1. 添加了版本更新信息",
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
