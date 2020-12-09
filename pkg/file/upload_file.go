/**
 * @Author pibing
 * @create 2020/12/9 10:32 AM
 */

package file

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strings"
	"github.com/google/uuid"
	imgType "github.com/shamsher31/goimgtype"
)

//响应结构体
type FileResponse struct {
	Size     int64  `json:"size"`           //文件大小
	Path     string `json:"path"`           //文件相对路径
	FullPath string `json:"full_path"`      //文件全路径
	Name     string `json:"name"`           //文件名称
	Type     string `json:"type"`           //文件类型
}


// @Summary 上传文件公用方法--按type区分单个上传/批量上传/base64上传
// @Param c *gin.Context true
// return  FileResponse,error

func UploadFile(c *gin.Context)(FileResponse,error) {

	tag, _ := c.GetPostForm("type")
	urlPerfix := fmt.Sprintf("http://%s/", c.Request.Host)
	var fileResponse FileResponse
	if tag == "" {
		return fileResponse,errors.New("缺少type标识")
	} else {
		switch tag {
		case "1": // 单图
			files, err := c.FormFile("file")
			if err != nil {
				return fileResponse,errors.New("文件不能为空")
			}
			// 上传文件至指定目录
			guid := uuid.New().String()

			singleFile := "static/files/" + guid + GetExt(files.Filename)
			_ = c.SaveUploadedFile(files, singleFile)
			fileType, _ := imgType.Get(singleFile)
			fileResponse = FileResponse{
				Size:     GetFileSize(singleFile),
				Path:     singleFile,
				FullPath: urlPerfix + singleFile,
				Name:     files.Filename,
				Type:     fileType,
			}
		//case "2": // 多图
		//	files := c.Request.MultipartForm.File["file"]
		//	var multipartFile []FileResponse
		//	for _, f := range files {
		//		guid := uuid.New().String()
		//		multipartFileName := "static/files/" + guid + GetExt(f.Filename)
		//		e := c.SaveUploadedFile(f, multipartFileName)
		//		fileType, _ := imgType.Get(multipartFileName)
		//		if e == nil {
		//			multipartFile = append(multipartFile, FileResponse{
		//				Size:     GetFileSize(multipartFileName),
		//				Path:     multipartFileName,
		//				FullPath: urlPerfix + multipartFileName,
		//				Name:     f.Filename,
		//				Type:     fileType,
		//			})
		//		}
		//	}
		//	return multipartFile,nil     //此处如果放开，需要修改返回类型
		case "3": // base64
			files, _ := c.GetPostForm("file")
			file2list := strings.Split(files, ",")
			ddd, _ := base64.StdEncoding.DecodeString(file2list[1])
			guid := uuid.New().String()
			base64File := "static/files/" + guid + ".jpg"
			_ = ioutil.WriteFile(base64File, ddd, 0666)
			typeStr := strings.Replace(strings.Replace(file2list[0], "data:", "", -1), ";base64", "", -1)
			fileResponse = FileResponse{
				Size:     GetFileSize(base64File),
				Path:     base64File,
				FullPath: urlPerfix + base64File,
				Name:     "",
				Type:     typeStr,
			}
		default:
			return fileResponse,errors.New("type类型错误")
		}

	}
	return fileResponse,nil
}


