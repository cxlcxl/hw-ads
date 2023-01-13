package servicemarketing

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
)

// Dictionary 定向规则
type Dictionary struct {
	Id       string        `json:"id"`
	Pid      string        `json:"pid"`
	Label    string        `json:"label"`
	Value    string        `json:"value"`
	Children []*Dictionary `json:"children"`
}

func FormatAppInterest(src []*Dictionary) (dist []*Dictionary) {
	dist = make([]*Dictionary, 0)
	for s, s2 := range vars.AppInterest {
		children := make([]*Dictionary, 0)
		for _, dictionary := range src {
			if dictionary.Pid == s2 {
				children = append(children, &Dictionary{
					Id:       dictionary.Id,
					Pid:      s,
					Label:    dictionary.Label,
					Value:    dictionary.Value,
					Children: nil,
				})
			}
		}
		dist = append(dist, &Dictionary{
			Id:       s,
			Pid:      "0",
			Label:    s2,
			Value:    s,
			Children: children,
		})
	}
	return
}

func FormatAppCategory(src []*Dictionary) (dist []*Dictionary) {
	dist = make([]*Dictionary, 0)
	for _, dictionary := range src {
		if dictionary.Pid == "0" {
			dist = append(dist, dictionary)
		}
	}
	return
}

func FormatMediaAppCategory(src []*Dictionary) (dist []*Dictionary) {
	return DictionaryToTree(src, "0")
}

func FormatCarrier(src []*Dictionary, countries []*model.OverseasRegion) (dist []*Dictionary) {
	tmp := make(map[string][]*Dictionary)
	for _, dictionary := range src {
		if _, ok := tmp[dictionary.Pid]; !ok {
			tmp[dictionary.Pid] = make([]*Dictionary, 0)
		}
		tmp[dictionary.Pid] = append(tmp[dictionary.Pid], dictionary)
	}
	for _, country := range countries {
		children, ok := tmp[country.CCode]
		if !ok {
			continue
		}
		dist = append(dist, &Dictionary{
			Id:       country.CId,
			Pid:      "0",
			Label:    country.CName,
			Value:    country.CId,
			Children: children,
		})
	}
	return
}

func FormatAudience(src []*Dictionary) (dist []*Dictionary) {
	dist = make([]*Dictionary, 0)
	tmp := make(map[string]struct{})
	for _, dictionary := range src {
		if _, ok := tmp[dictionary.Id]; !ok {
			dist = append(dist, dictionary)
		}
		tmp[dictionary.Id] = struct{}{}
	}
	return
}

// DictionaryToTree 数组转树
func DictionaryToTree(origin []*Dictionary, pid string) (d []*Dictionary) {
	for _, v := range origin {
		if v.Pid == pid {
			v.Children = DictionaryToTree(origin, v.Id)
			d = append(d, v)
		}
	}
	return d
}
