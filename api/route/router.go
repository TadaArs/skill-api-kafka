package router

import (
	"github.com/gin-gonic/gin"
	"skill/skill"
)


func NewRouter(skillHandler *skill.SkillHandler) *gin.Engine {

	router := gin.Default()

	r := router.Group("/api/v1")
	r.GET("/skills/:key", skillHandler.GetSkillByKey)
    r.GET("/skills", skillHandler.GetSkills)
 	r.POST("/skills",skillHandler.CreateSkill)
	r.PUT("/skills/:key", skillHandler.UpdateSkill)
	r.PATCH("/skills/:key/actions/name", skillHandler.UpdateSkillNameByKey)
	r.PATCH("/skills/:key/actions/description", skillHandler.UpdateSkillDescriptionByKey)
	r.PATCH("/skills/:key/actions/logo", skillHandler.UpdateSkillLogoByKey)
	r.PATCH("/skills/:key/actions/tags", skillHandler.UpdateSkillTagsByKey)
	r.DELETE("/skills/:key", skillHandler.DeleteSkill)

	return router
}
