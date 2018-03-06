package views

import (
	"github.com/gin-contrib/multitemplate"
)

var prefixViewFolder = "app/views/"

// http://blog.questionable.services/article/approximating-html-template-inheritance/
func GetMultiRender() multitemplate.Render {
	r := multitemplate.New()
	r.AddFromFiles("crypto_index", prefixViewFolder+"layouts/base.html", prefixViewFolder+"crypto/index.html")
	r.AddFromFiles("ark_index", prefixViewFolder+"layouts/base.html", prefixViewFolder+"explorer/ark_indx.html")
	r.AddFromFiles("eth_index", prefixViewFolder+"layouts/base.html", prefixViewFolder+"explorer/eth_indx.html")
	r.AddFromFiles("voters_index", prefixViewFolder+"layouts/base.html", prefixViewFolder+"delgate/voters_index.html")
	r.AddFromFiles("voters_profit_index", prefixViewFolder+"layouts/base.html", prefixViewFolder+"delgate/voters_profit_index.html")
	return r
}
