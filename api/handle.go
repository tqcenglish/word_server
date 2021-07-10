package api

import (
	"fmt"
	"github/tqcenglish/word-english/app/model"
	"github/tqcenglish/word-english/configs"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//DetailWord 查看单词详情
/**
 * @api {post} /modbus/config 查看单词详情
 * @apiName DetailWord
 * @apiGroup modbus
 *
 * @apiParam {Object[]} dis 更新数据
 * @apiParam {string} dis.name name
 * @apiParam {string} dis.low  low url
 * @apiParam {string} dis.high high url
 *
 * @apiSuccess {String} status success.
 */
func DetailWord(ctx *gin.Context) {
	word := ctx.Param("word")

	//查找文件
	path := fmt.Sprintf("%s/out/%s/%s", configs.BasicPath, string(word[0]), word)

	wordDir, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			logrus.Errorf("path %s not exist", path)
			ctx.JSON(http.StatusOK, gin.H{"status": "failure", "message": "not exist"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "failure", "message": err.Error()})
		return
	}
	if !wordDir.IsDir() {
		ctx.JSON(http.StatusOK, gin.H{"status": "failure", "message": "not dir"})
		return
	}

	// 遍历文件
	ctx.JSON(http.StatusOK, gin.H{"status": "success",
		"data": []string{
			fmt.Sprintf("%s/%s/0.jpg", string(word[0]), word),
			fmt.Sprintf("%s/%s/2.jpg", string(word[0]), word),
			fmt.Sprintf("%s/%s/3.jpg", string(word[0]), word),
			fmt.Sprintf("%s/%s/4.jpg", string(word[0]), word),
		},
		"voice": fmt.Sprintf("%s/%s/%s.mp3", string(word[0]), word, word),
	})
	return

}

//GetAllWords 获取所有
/**
 * @api {get} /modbus/config 获取所有
 * @apiName GetAllWords
 * @apiGroup modbus
 *
 * @apiSuccess {Ojbect}  data 数据
 * @apiSuccess {Object[]}  data.dis 数据
 * @apiSuccess {string}  data.dis.name name
 * @apiSuccess {string}  data.dis.low low url
 * @apiSuccess {string}  data.dis.high high url
 * @apiSuccess {string}  data.status success
 */
func GetAllWords(ctx *gin.Context) {
	data, _ := configs.LoadValue()
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": data})
}

//PlayMP3 测试发送
/**
 * @api {post} /modbus/play-mp3 测试发送
 * @apiName PlayMP3
 * @apiGroup modbus
 *
 * @apiParam {String} url 调用此 url
 *
 * @apiSuccess {String} status 调用结果
 */
func PlayMP3(ctx *gin.Context) {
	var param model.PlayMP3Param

	if err := ctx.ShouldBindJSON(&param); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}

	// 请求 url
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
