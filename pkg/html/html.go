package html

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

// Checks if the user is logged in and has a session, if not err is not nil
//func session(writer http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
//	cookie, err := request.Cookie("_cookie")
//	if err == nil {
//		sess = models.Session{Uuid: cookie.Value}
//		if ok, _ := sess.Check(); !ok {
//			err = errors.New("Invalid session")
//		}
//	}
//	return
//}

// 生成 HTML 模板
func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("layout").Funcs(funcMap)
	templates := template.Must(t.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

func LoadTemplates() multitemplate.Renderer {
	templatesDir := "web/views"
	render := multitemplate.NewRenderer()
	//
	layouts, err := filepath.Glob(templatesDir + "/layout/*.html")

	if err != nil {
		panic(err.Error())
	}
	includes, err := filepath.Glob(templatesDir + "/template/*.html")
	if err != nil {
		panic(err.Error())
	}

	for _, include := range includes {
		//layoutCopy := make([]string, len(layouts))
		//copy(layoutCopy, layouts)
		//files := append(layoutCopy, include)
		//dirSlice := strings.Split(include, string(filepath.Separator))
		//fileName := strings.Join(dirSlice[len(dirSlice)-2:], "/")
		//config.Logger.Infof("fileName:%s", fileName)
		////render.AddFromString(fileName, template2.GlobalTemplateFun, files...)
		//render.AddFromFiles(fileName, files...)
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		render.AddFromFiles(filepath.Base(include), files...)
	}

	return render
}

// version
func Version() string {
	return "0.1"
}

// 异常处理统一重定向到错误页面
func errorMessage(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

// 日期格式化
func formatDate(t time.Time) string {
	datetime := "2006-01-02 15:04:05"
	return t.Format(datetime)
}
