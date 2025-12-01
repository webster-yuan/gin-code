package handlers

import (
	pb "gin/pkg/pb/echo"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

// ProtoHandler 处理Protocol Buffers路由
func ProtoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := &pb.EchoResp{
			Label: "test",
			Nums:  []int64{1, 2},
		}
		out, _ := proto.Marshal(resp)
		c.Data(http.StatusOK, "application/x-protobuf", out)
	}
}
