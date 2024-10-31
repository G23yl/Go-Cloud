// 提供页面数据

package services

import (
	"disk/database"
	"disk/model/request"
	"disk/model/response"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 返回总数据：各类文件数量，仓库总大小和目前大小
func Index(ctx *gin.Context) {
	// 获取用户id
	userID, ok := ctx.Get("userID")
	if !ok {
		response.UnauthorizedError(ctx, response.ErrorUnauthorized)
		return
	}
	// 查询仓库容量和大小
	storeID, capacity, size := database.GetStoreInfo(userID.(uint))
	// 各类文件的数量和总文件数
	categoryNum, filesNum := database.GetFilesNum(storeID)
	// 文件夹数量
	foldersNum := database.GetFoldersNum(storeID)
	response.Success(ctx, "获取成功", gin.H{
		"filesNum":   filesNum,
		"foldersNum": foldersNum,
		"storeDetail": gin.H{
			"capacity": capacity,
			"size":     size,
		},
		"categoryNum": categoryNum,
	})
}

// docs文件
func Docs(ctx *gin.Context) {
	// 获取用户id
	userID, ok := ctx.Get("userID")
	if !ok {
		response.UnauthorizedError(ctx, response.ErrorUnauthorized)
		return
	}
	// 获取仓库ID
	storeID, _, _ := database.GetStoreInfo(userID.(uint))
	// 获取docs文件
	primitiveDocs := database.GetCategoryFiles(storeID, 1)
	docs := make([]response.CategoryRes, len(primitiveDocs))
	for idx, doc := range primitiveDocs {
		docs[idx] = response.CategoryRes{
			FileID:      doc.ID,
			FileName:    doc.FileName,
			FileSize:    doc.FileSize,
			UpdateTime:  doc.UpdatedAt,
			DownloadNum: doc.DownloadCount,
			FilePath:    doc.FilePath,
		}
	}
	response.Success(ctx, "获取成功", docs)
}

// Images文件
func Images(ctx *gin.Context) {
	// 获取用户id
	userID, ok := ctx.Get("userID")
	if !ok {
		response.UnauthorizedError(ctx, response.ErrorUnauthorized)
		return
	}
	// 获取仓库ID
	storeID, _, _ := database.GetStoreInfo(userID.(uint))
	// 获取images文件
	primitiveImages := database.GetCategoryFiles(storeID, 2)
	images := make([]response.CategoryRes, len(primitiveImages))
	for idx, image := range primitiveImages {
		images[idx] = response.CategoryRes{
			FileID:      image.ID,
			FileName:    image.FileName,
			FileSize:    image.FileSize,
			UpdateTime:  image.UpdatedAt,
			DownloadNum: image.DownloadCount,
			FilePath:    image.FilePath,
		}
	}
	response.Success(ctx, "获取成功", images)
}

// Videos文件
func Videos(ctx *gin.Context) {
	// 获取用户id
	userID, ok := ctx.Get("userID")
	if !ok {
		response.UnauthorizedError(ctx, response.ErrorUnauthorized)
		return
	}
	// 获取仓库ID
	storeID, _, _ := database.GetStoreInfo(userID.(uint))
	// 获取videos文件
	primitiveVideos := database.GetCategoryFiles(storeID, 3)
	videos := make([]response.CategoryRes, len(primitiveVideos))
	for idx, video := range primitiveVideos {
		videos[idx] = response.CategoryRes{
			FileID:      video.ID,
			FileName:    video.FileName,
			FileSize:    video.FileSize,
			UpdateTime:  video.UpdatedAt,
			DownloadNum: video.DownloadCount,
			FilePath:    video.FilePath,
		}
	}
	response.Success(ctx, "获取成功", videos)
}

// Audios文件
func Audios(ctx *gin.Context) {
	// 获取用户id
	userID, ok := ctx.Get("userID")
	if !ok {
		response.UnauthorizedError(ctx, response.ErrorUnauthorized)
		return
	}
	// 获取仓库ID
	storeID, _, _ := database.GetStoreInfo(userID.(uint))
	// 获取audios文件
	primitiveAudios := database.GetCategoryFiles(storeID, 4)
	audios := make([]response.CategoryRes, len(primitiveAudios))
	for idx, audio := range primitiveAudios {
		audios[idx] = response.CategoryRes{
			FileID:      audio.ID,
			FileName:    audio.FileName,
			FileSize:    audio.FileSize,
			UpdateTime:  audio.UpdatedAt,
			DownloadNum: audio.DownloadCount,
			FilePath:    audio.FilePath,
		}
	}
	response.Success(ctx, "获取成功", audios)
}

// Others文件
func Others(ctx *gin.Context) {
	// 获取用户id
	userID, ok := ctx.Get("userID")
	if !ok {
		response.UnauthorizedError(ctx, response.ErrorUnauthorized)
		return
	}
	// 获取仓库ID
	storeID, _, _ := database.GetStoreInfo(userID.(uint))
	// 获取others文件
	primitiveOthers := database.GetCategoryFiles(storeID, 5)
	others := make([]response.CategoryRes, len(primitiveOthers))
	for idx, other := range primitiveOthers {
		others[idx] = response.CategoryRes{
			FileID:      other.ID,
			FileName:    other.FileName,
			FileSize:    other.FileSize,
			UpdateTime:  other.UpdatedAt,
			DownloadNum: other.DownloadCount,
			FilePath:    other.FilePath,
		}
	}
	response.Success(ctx, "获取成功", others)
}

// 上传文件
func Upload(ctx *gin.Context) {
	// 获取用户id
	userID, ok := ctx.Get("userID")
	if !ok {
		response.UnauthorizedError(ctx, response.ErrorUnauthorized)
		return
	}
	// 获取仓库ID
	storeID, capacity, size := database.GetStoreInfo(userID.(uint))
	// 获取文件
	file, err := ctx.FormFile("file")
	if err != nil {
		response.RequestError(ctx, response.ErrorFileUpload)
		return
	}
	// 判断是否超过仓库容量，超过直接返回
	if capacity < size+file.Size {
		response.RequestError(ctx, response.ErrorStoreFull)
		return
	}
	// 获取文件的路由
	var fileInfo request.UploadReq
	if err := ctx.ShouldBind(&fileInfo); err != nil {
		response.RequestError(ctx, response.ErrorParams)
		return
	}
	// 获取文件要保存的位置并判断有没有重复文件，如果有，就不允许上传
	// 查询数据库中是否有filename相同且filepath也相同的文件
	if exist := database.CheckFileExists(storeID, file.Filename, fileInfo.FilePath); exist {
		response.RequestError(ctx, response.ErrorFileExist)
		return
	}
	// 没有呢
	// 保存文件
	dir, _ := os.Getwd()
	// 需要给每个用户都分配一个空间
	localFilePath := fmt.Sprintf("%s/static/local_store/user_%d%s", dir, userID.(uint), fileInfo.FilePath)
	if fileInfo.FilePath != "/" {
		localFilePath += "/"
	}
	//FIXME 感觉有个bug在这里，如果路径对应的文件夹不存在的话应该要报错的，但是只要用户不输入保存位置就不会出现这个bug，以后再说
	err = ctx.SaveUploadedFile(file, localFilePath+file.Filename)
	if err != nil {
		response.RequestError(ctx, response.ErrorFileUpload)
		return
	}
	// 存在数据库中
	err = database.CreateFile(storeID, fileInfo.FilePath, file.Filename, file.Size)
	if err != nil {
		response.RequestError(ctx, response.ErrorFileUpload)
		return
	}
	response.Success(ctx, "上传成功", nil)
}

// 获取对应路由下的文件和文件夹
func GetInPathFiles(ctx *gin.Context) {
	// 采用query参数，类似于?path=xxx
	path := ctx.Query("path")
	//查询所有对应path的文件和文件夹
	files, folders := database.GetInPathFiles(path)
	// 把查询到的结果整理一下返回
	res := make([]response.FFRes, len(files)+len(folders))
	for idx, f := range folders {
		res[idx] = response.FFRes{
			FileID:     f.ID,
			FileName:   f.FolderName,
			UpdateTime: f.UpdatedAt,
			Type:       "dir",
			FileType:   0,
			FilePath:   f.FilePath,
		}
	}
	for idx, f := range files {
		res[idx+len(folders)] = response.FFRes{
			FileID:      f.ID,
			FileName:    f.FileName,
			UpdateTime:  f.UpdatedAt,
			Type:        "file",
			FileType:    f.Type,
			FileSize:    f.FileSize,
			DownloadNum: f.DownloadCount,
			FilePath:    f.FilePath,
		}
	}
	response.Success(ctx, "获取成功", res)
}

// 删除文件
func DeleteFile(ctx *gin.Context) {
	// 获取用户id
	userID, ok := ctx.Get("userID")
	if !ok {
		response.UnauthorizedError(ctx, response.ErrorUnauthorized)
		return
	}
	// 获取仓库ID
	storeID, _, _ := database.GetStoreInfo(userID.(uint))
	// 获取删除文件id
	fileID := ctx.Query("fileID")
	IFileID, _ := strconv.Atoi(fileID)
	filePath := ctx.Query("filePath")
	fileName := ctx.Query("fileName")
	// 删除文件时，应该先删除本地文件，然后删除数据库信息
	// 如果先删除数据库成功，但是后删除本地文件失败，那么用户看不到文件，就不能再重试删除文件了
	// 如果先删除本地文件成功，但是后删除数据库失败，那么用户依然看得到文件，可以重试删除文件
	// 删除本地文件
	dir, _ := os.Getwd()
	localFilePath := fmt.Sprintf("%s/static/local_store/user_%d%s", dir, userID.(uint), filePath)
	if filePath != "/" {
		localFilePath += "/"
	}
	fmt.Println(localFilePath)
	err := os.Remove(localFilePath + fileName)
	if err != nil {
		fmt.Println("111")
		response.ServerError(ctx, response.ErrorDeleteFile)
		return
	}
	// 删除数据库
	err = database.DeleteFile(storeID, uint(IFileID))
	if err != nil {
		response.ServerError(ctx, response.ErrorDeleteFile)
		return
	}
	// 删除成功
	response.Success(ctx, "删除成功", nil)
}
