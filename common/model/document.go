package model

type Document struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Documents []Document

func (d Documents) Len() int {
	return len(d)
}

func (d Documents) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d Documents) Less(i, j int) bool {
	if d[i].Id < d[j].Id {
		return true
	}
	if d[i].Id > d[j].Id {
		return false
	}
	if d[i].Name < d[j].Name {
		return true
	}
	if d[i].Name > d[j].Name {
		return false
	}
	return true
}
