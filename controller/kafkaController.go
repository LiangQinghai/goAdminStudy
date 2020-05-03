package controller

import (
	"github.com/gin-gonic/gin"
	"goAdminStudy/log"
	"goAdminStudy/models"
	"goAdminStudy/utils"
)

type KafkaController struct {
}

// add kafka config
// @Summery add kafka config
// @Description this is desc
// @Accept json
// @Produce json
// @Param name body models.KafkaConfig true "Name of the kafka config"
// @Success 200 {object} models.KafkaConfig
// @Header 200 {string} Token "token"
// @Router /adm/kafka/save [post]
func (k *KafkaController) Add(c *gin.Context) {
	kafkaConfig := models.KafkaConfig{}
	if err := c.ShouldBindJSON(&kafkaConfig); err != nil {
		panic(err)
	}
	if err := kafkaConfig.Save(); err != nil {
		log.Errorf("%v", err)
		utils.ErrorResult(err.Error())
	}
	utils.SuccessResultWithContext(kafkaConfig, c)
}

// @Summery 获取列表
// @Accept json
// @Produce json
// @Param data body utils.Page true "分页参数"
// @Success 200 {object} models.KafkaConfig
// @Router /adm/kafka/list [get]
func (k *KafkaController) List(ctx *gin.Context) {

	page := utils.NewDefaultPage()
	if err := ctx.ShouldBindJSON(&page); err != nil {
		panic(err)
	}

	kafkaConfig := models.KafkaConfig{}
	if kafkaConfigs, err := kafkaConfig.GetByPage(&page); err != nil {
		panic(err)
	} else {
		utils.SuccessPageResult(&page, kafkaConfigs, ctx)
	}

}
