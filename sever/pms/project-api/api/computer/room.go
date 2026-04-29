package computer

import "github.com/gin-gonic/gin"

type HandlerComputer struct {
}

func New() *HandlerComputer {
	return &HandlerComputer{}
}

func (c HandlerComputer) computer(context *gin.Context) {

}
