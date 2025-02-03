package client
import "github.com/cloudwego/hertz/pkg/app/server"
import "filemanager/handler"
func main(){
	h:=server.Default(server.WithHostPorts("127.0.0.1:8080"))
	h.POST("/upload",handler.UploadFile)
	h.GET("/download",handler.DownloadFile)
	h.DELETE("/delete",handler.DeleteFile)
	h.Spin()
}