package web

import (
	"go_web/util"
	"net/http"
)

func Index(w http.ResponseWriter,r *http.Request){
	r.ParseMultipartForm(1024*1024*3)




	index_template:=util.Get_tmeplate_path()+"index.html"
	util.Render_template(w,index_template,nil)

}
