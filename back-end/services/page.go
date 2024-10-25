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
