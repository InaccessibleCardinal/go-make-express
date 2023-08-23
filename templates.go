package main

type Template struct {
	Name string
	Path string
}

func GetTemplates() map[string]Template {
	return map[string]Template{
		"fnBare": {Name: "fnBare", Path: "templates\\express-template-fn-bare"},
		"fnNt":   {Name: "fnNt", Path: "templates\\express-template-fn-nt"},
		"fnVan":  {Name: "fnVan", Path: "templates\\express-template-vanilla"},
		"oop":    {Name: "oop", Path: "templates\\express-template-oop"},
	}
}
