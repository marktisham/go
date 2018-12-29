package controllers

import (
	"GoMVC/app/models"
	"github.com/revel/revel"
)

type MyStuff struct {
	*revel.Controller
}

func (c MyStuff) MarkStuff() revel.Result {

	mystuff:=models.MyModel{}
	mystuff.Message="hello from my model!"
	mystuff.Value=123

	mystuff2:=models.MyModel{"Message 2",123}

	return c.Render(mystuff,mystuff2)
}