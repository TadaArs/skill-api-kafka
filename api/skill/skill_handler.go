package skill

import (

	"net/http"
	"skill/errs"
	"skill/response"

	
	"github.com/gin-gonic/gin"
)

type SkillHandler struct {
	storager SkillStorager
}

func NewSkillHandler(storager SkillStorager) (*SkillHandler){
	return &SkillHandler{storager: storager}
}

func (h *SkillHandler) GetSkillByKey(ctx *gin.Context){
	key := ctx.Param("key")
	skill, err := h.storager.GetSkillByKey(key)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, http.StatusOK, skill)
}

func (h *SkillHandler) CreateSkill(ctx *gin.Context){
	skill := Skill{}
	err := ctx.BindJSON(&skill)
	if err != nil {
		response.Error(ctx, errs.NewError(http.StatusBadRequest, "Can't bind payload"))
	}
	// strMsg, err := json.Marshal(&skill)
	// if err != nil {
	// 	response.Error(ctx, errs.NewError(http.StatusInternalServerError, "Can't convert payload struct"))
	// }
	// print(string(strMsg))
	// msg := &sarama.ProducerMessage{Topic: "Define-Topic", Value: sarama.StringEncoder(string(strMsg))}
	
}