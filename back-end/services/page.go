// 提供页面数据

package services

import (
	"disk/database"
	"disk/model/response"

	"github.com/gin-gonic/gin"
)

// 返回总数据：各类文件数量，仓库总大小和目前大小
func Index(ctx *gin.Context) {
	// 获取用户id
	userID, ok := ctx.Get("userID")
	if !ok {
		response.UnauthorizedError(ctx, response.ErrorUnauthorized)
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
			FileUrl:     doc.FilePath,
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
	}
	// 获取仓库ID
	storeID, _, _ := database.GetStoreInfo(userID.(uint))
	// 获取images文件
	primitiveImages := database.GetCategoryFiles(storeID, 1)
	images := make([]response.CategoryRes, len(primitiveImages))
	for idx, image := range primitiveImages {
		images[idx] = response.CategoryRes{
			FileID:      image.ID,
			FileName:    image.FileName,
			FileSize:    image.FileSize,
			UpdateTime:  image.UpdatedAt,
			DownloadNum: image.DownloadCount,
			FileUrl:     image.FilePath,
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
	}
	// 获取仓库ID
	storeID, _, _ := database.GetStoreInfo(userID.(uint))
	// 获取videos文件
	primitiveVideos := database.GetCategoryFiles(storeID, 1)
	videos := make([]response.CategoryRes, len(primitiveVideos))
	for idx, video := range primitiveVideos {
		videos[idx] = response.CategoryRes{
			FileID:      video.ID,
			FileName:    video.FileName,
			FileSize:    video.FileSize,
			UpdateTime:  video.UpdatedAt,
			DownloadNum: video.DownloadCount,
			FileUrl:     video.FilePath,
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
			FileUrl:     audio.FilePath,
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
			FileUrl:     other.FilePath,
		}
	}
	response.Success(ctx, "获取成功", others)
}
