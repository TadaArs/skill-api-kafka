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

type kafkaMsg struct {
	Action string `json:"action"`
	Key    string `json:"key,omitempty"`
	Data   any    `json:"data,omitempty"`
}

func NewSkillHandler(storager SkillStorager, producer Producer) *SkillHandler {
	return &SkillHandler{storager: storager, producer: producer}
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
	msg := kafkaMsg{
		Action: "INSERT",
		Data: skill,
	}
	strMsg, err := json.Marshal(&msg)
	if err != nil {
		response.Error(ctx, errs.NewError(http.StatusInternalServerError, "Can't convert payload struct"))
	}
	h.producer.Publish(os.Getenv("TOPIC"), strMsg)
}
