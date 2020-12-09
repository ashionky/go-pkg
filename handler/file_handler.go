/**
 * @Author pibing
 * @create 2020/12/9 11:05 AM
 */

package handler

import (
	"github.com/gin-gonic/gin"
	"go-pkg/constant"
	"go-pkg/pkg/file"
	"go-pkg/pkg/vo"
)

func UploadFile(c *gin.Context)  {
	res := vo.GetDefaultResult()
	data,err:=file.UploadFile(c)
	if err != nil {
		vo.SendFailure(c,constant.InternalError,res)
		return
	}
	res.Data=data
	vo.SendSuccess(c,res)
}
