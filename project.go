package main

type Project struct {
}

func (cp *Project) list() []string {
	return []string{
		"test",
		"demo",
		"lol",
	}
}
