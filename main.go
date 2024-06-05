package main

import (
  "context"
  "github.com/cloudwego/hertz/pkg/app"
  "github.com/cloudwego/hertz/pkg/app/server"
  "github.com/cloudwego/hertz/pkg/common/hlog"
  "github.com/cloudwego/hertz/pkg/common/utils"
  "io"
  "log"
  "os"
  "strings"
)

func init() {
  f, err := os.OpenFile("/data/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
  if err != nil {
    log.Fatalln(err)
  }
  fileWriter := io.MultiWriter(f, os.Stdout)
  log.SetOutput(fileWriter)
  log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
}
func main() {
  h := server.Default()
  h.POST("/log", PostLogHandler)
  h.Spin()
}

func PostLogHandler(ctx context.Context, reqCtx *app.RequestContext) {
  var reqVo struct {
    Level string                   `json:"level"`
    Args  []map[string]interface{} `json:"args"`
  }
  token := os.Getenv("TOKEN")
  authorization := string(reqCtx.GetHeader("Authorization"))

  if !strings.HasSuffix(authorization, token) {
    reqCtx.JSON(403, utils.H{"error": "Unauthorized"})
    return
  }

  err := reqCtx.BindJSON(&reqVo)
  if err != nil {
    hlog.Error(err)
    reqCtx.JSON(400, utils.H{"error": "Invalid request payload"})
    return
  }
  go func() {
    log.Printf("%s: %v", reqVo.Level, reqVo.Args)
  }()
}
