package controllers

import beego "github.com/beego/beego/v2/server/web"

type PageController struct {
	beego.Controller
}

// Get Get
func (c *PageController) Get() {
	// {
	// 	"type": "page",
	// 	"title": "Hello world",
	// 	"body": [
	// 	  {
	// 		"type": "icon-picker",
	// 		"label": "图标",
	// 		"name": "icon",
	// 		"id": "u:5b20517a78c4"
	// 	  }
	// 	],
	// 	"id": "u:9e47b57f4f2e"
	//   }
	resp := make(map[string]interface{}, 0)
	resp["type"] = "page"
	resp["title"] = "Hello world"

	body := make(map[string]string, 0)
	body["type"] = "icon-picker"
	body["label"] = "图标"
	body["name"] = "icon"

	resp["body"] = []interface{}{body}

	c.Data["json"] = &resp
	c.ServeJSON()
}
