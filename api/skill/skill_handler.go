package skill

import (
	"encoding/json"
	"net/http"
	"os"
	"skill/errs"
	"skill/response"

	"github.com/gin-gonic/gin"
)

type SkillHandler struct {
	storager SkillStorager
	producer Producer
}

func NewSkillHandler(storager SkillStorager, producer Producer) *SkillHandler {
	return &SkillHandler{storager: storager}
}

func (h *SkillHandler) GetSkillByKey(ctx *gin.Context) {
	key := ctx.Param("key")
	skill, err := h.storager.GetSkillByKey(key)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, http.StatusOK, skill)
}

func (h *SkillHandler) CreateSkill(ctx *gin.Context) {
	skill := Skill{}
	err := ctx.BindJSON(&skill)
	if err != nil {
		response.Error(ctx, errs.NewError(http.StatusBadRequest, "Can't bind payload"))
	}
	strMsg, err := json.Marshal(&skill)
	if err != nil {
		response.Error(ctx, errs.NewError(http.StatusInternalServerError, "Can't convert payload struct"))
	}
	h.producer.Publish(os.Getenv("CREATE_TOPIC"), strMsg)
}
