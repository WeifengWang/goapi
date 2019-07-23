package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type LogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w LogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func PrintReqAndRes(ctx *gin.Context) {
	blw := &LogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
	ctx.Writer = blw
	fmt.Printf("request body is:%s\n", string(getArgs(ctx)))
	ctx.Next()
	fmt.Printf("response body is:%s\n", strings.TrimSpace(blw.body.String()))
}


//get request params
func getArgs(c *gin.Context) []byte {
	if c.ContentType() == "multipart/form-data" {
		c.Request.ParseMultipartForm(1024*1024*10)
	} else {
		c.Request.ParseForm()
	}
	args, _ := json.Marshal(c.Request.Form)
	return args
}
