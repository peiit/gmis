package controllers

import (
	"gmis/lib/types"

	beego "github.com/beego/beego/v2/server/web"
)

type NavController struct {
	beego.Controller
}

// Get Get
func (c *NavController) Get() {
	resp := types.AmisSiteResp{}
	data := types.AmisSiteData{}
	pages := make([]types.AmisSitePage, 0, 100)

	linkpage := types.AmisSitePage{
		Icon:  "fa fa-cloud-sun-rain",
		Label: "Amis Label",
		Link:  "https://www.baidu.com/",
	}

	schemaPage := types.AmisSitePage{
		Icon:      "fa fa-cloud",
		Label:     "Amis Schema Page",
		URL:       "/schemapage",
		SchemaAPI: "/_page",
	}

	page := types.AmisSitePage{
		Label:         "Amis",
		IsDefaultPage: true,
		Children:      []types.AmisSitePage{linkpage, schemaPage},
	}

	page2 := types.AmisSitePage{
		Label:         "Amis page2",
		IsDefaultPage: true,
		Children:      []types.AmisSitePage{linkpage},
	}

	pages = []types.AmisSitePage{page, page2}
	data.Pages = pages
	resp.Data = data

	c.Data["json"] = &resp
	c.ServeJSON()
}
