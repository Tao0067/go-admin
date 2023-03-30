package sessions

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SessionsCookieName  = "session_id"
	SessionsContextName = "AUTH"
)

func SessonMiddlerware() gin.HandlerFunc {
	return func(context *gin.Context) {
		session := sessions.Default(context)
		data := session.Get(SessionsContextName)
		if data == nil {
			context.Redirect(http.StatusMovedPermanently, "/login")
			return
		}
		context.Next()
	}
}

func Save(context *gin.Context, data string) {
	session := sessions.Default(context)
	session.Set(SessionsContextName, data)
	session.Save()
}

func getUuid(context *gin.Context) string {
	session := sessions.Default(context)
	data := session.Get(SessionsContextName).(string)
	return data
}
