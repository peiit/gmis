package types

type AmisSiteResp struct {
	Status int          `json:"status"`
	Msg    string       `json:"msg"`
	Data   AmisSiteData `json:"data"`
}
type AmisSiteSchema struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type AmisSiteData struct {
	Pages []AmisSitePage `json:"pages"`
}

// https://github.com/baidu/amis/blob/master/packages/amis/src/renderers/App.tsx interface AppPage
type AmisSitePage struct {
	Label         string         `json:"label"`                   //菜单文字
	Icon          string         `json:"icon,omitempty"`          //菜单图标，比如： fa fa-file
	URL           string         `json:"url,omitempty"`           //路由规则。比如：/banner/:id。当地址以 / 打头，则不继承上层的路径，否则将集成父级页面的路径。
	Redirect      string         `json:"redirect,omitempty"`      //当match url 时跳转到目标地址.没有配置 schema 和 shcemaApi  时有效.
	Rewrite       string         `json:"rewrite,omitempty"`       //当match url 转成渲染目标地址的页面.没有配置 schema 和 shcemaApi  时有效.
	IsDefaultPage bool           `json:"isDefaultPage,omitempty"` //不要出现多个，如果出现多个只有第一个有用。在路由找不到的时候作为默认页面。
	Link          string         `json:"link,omitempty"`          //单纯的地址。可以设置外部链接。
	Children      []AmisSitePage `json:"children,omitempty"`      //支持多层级。
	Visible       bool           `json:"visible,omitempty"`       //是否在导航中可见，适合于那种需要携带参数才显示的页面。比如具体某个数据的编辑页面。
	// ClassName     string    `json:"className,omitempty"`     //菜单上的类名 string or string dict

	// 默认是自动，即：自己选中或者有孩子节点选中则展开。
	//如果配置成 always 或者配置成 true 则永远展开。
	// 如果配置成 false 则永远不展开。
	Expanded string `json:"expanded,omitempty"`

	//二选一，如果配置了 url 一定要配置。否则不知道如何渲染。
	Schema    *AmisSiteSchema `json:"schema,omitempty"`
	SchemaAPI string          `json:"schemaApi,omitempty"`
}
