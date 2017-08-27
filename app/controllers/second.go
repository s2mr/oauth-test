package controllers

import "github.com/revel/revel"

type Second struct {
	*revel.Controller
	RenderArgs map[string]interface{} // Args passed to the template.
}

func (s Second) Index() revel.Result {

	//l := "yeah"
	s.RenderArgs["a"] = "hoge"

	s.Render()

	return s.Render()
}
