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
	// Context *gin.Context `json:"context,omitempty"`
}

func NewSkillHandler(storager SkillStorager, producer Producer) *SkillHandler {
	return &SkillHandler{storager: storager, producer: producer}
}

func (h *SkillHandler) GetSkills(ctx *gin.Context){
	skills, err := h.storager.GetSkills()
	if err != nil {
		response.Error(ctx, err)
		return
	}
	
	response.Success(ctx, http.StatusOK, skills)
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
		response.Error(ctx, errs.NewError(http.StatusConflict, "Can't convert payload struct"))
	}
	h.producer.Publish(os.Getenv("TOPIC"), strMsg)
}

func (h *SkillHandler) UpdateSkill(ctx *gin.Context) {
	skill := Skill{}
	key := ctx.Param("key")
	err := ctx.BindJSON(&skill)
	if err != nil {
		response.Error(ctx, errs.NewError(http.StatusBadRequest, "Can't bind payload"))
	}
	msg := kafkaMsg{
		Action: "UPDATE",
		Key: key,
		Data: skill,
	}
	strMsg, err := json.Marshal(&msg)
	if err != nil {
		response.Error(ctx, errs.NewError(http.StatusConflict, "Can't convert payload struct"))
	}
	h.producer.Publish(os.Getenv("TOPIC"), strMsg)
}

func (h *SkillHandler) UpdateSkillNameByKey(ctx *gin.Context){
	req := NameUpdateRequest{}
	key := ctx.Param("key")
	err := ctx.BindJSON(&req)
	if err != nil {
		response.Error(ctx, errs.NewError(http.StatusBadRequest, "Can't bind payload"))
		return
	}
	msg := kafkaMsg{
		Action: "PATCH-NAME",
		Key: key,
		Data: req,
	}
	strMsg, err := json.Marshal(&msg)
	if err != nil {
		response.Error(ctx, errs.NewError(http.StatusConflict, "Can't convert payload struct"))
	}
	h.producer.Publish(os.Getenv("TOPIC"), strMsg)
}

func (h *SkillHandler) UpdateSkillDescriptionByKey(ctx *gin.Context){
	req := DescriptionUpdateRequest{}
	key := ctx.Param("key")
	err := ctx.BindJSON(&req)
	if err != nil {
		response.Error(ctx, errs.NewError(http.StatusBadRequest, "Can't bind payload"))
		return
	}
	msg := kafkaMsg{
		Action: "PATCH-DESCRIPTION",
		Key: key,
		Data: req,
	}
	strMsg, err := json.Marshal(&msg)
	if err != nil {
		response.Error(ctx, errs.NewError(http.StatusConflict, "Can't convert payload struct"))
	}
	h.producer.Publish(os.Getenv("TOPIC"), strMsg)
}

func (h *SkillHandler) UpdateSkillLogoByKey(ctx *gin.Context){
	req := LogoUpdateRequest{}
	key := ctx.Param("key")
	err := ctx.BindJSON(&req)
	if err != nil {
		response.Error(ctx, errs.NewError(http.StatusBadRequest, "Can't bind payload"))
		return
	}
	msg := kafkaMsg{
		Action: "PATCH-LOGO",
		Key: key,
		Data: req,

	}
	strMsg, err := json.Marshal(&msg)
	if err != nil {
		response.Error(ctx, errs.NewError(http.StatusConflict, "Can't convert payload struct"))
	}
	h.producer.Publish(os.Getenv("TOPIC"), strMsg)
}

func (h *SkillHandler) UpdateSkillTagsByKey(ctx *gin.Context){
	req := TagsUpdateRequest{}
	key := ctx.Param("key")
	err := ctx.BindJSON(&req)
	if err != nil {
		response.Error(ctx, errs.NewError(http.StatusBadRequest, "Can't bind payload"))
		return
	}
	msg := kafkaMsg{
		Action: "PATCH-TAGS",
		Key: key,
		Data: req,
	}
	strMsg, err := json.Marshal(&msg)
	if err != nil {
		response.Error(ctx, errs.NewError(http.StatusConflict, "Can't convert payload struct"))
	}
	h.producer.Publish(os.Getenv("TOPIC"), strMsg)
}

func (h *SkillHandler) DeleteSkill(ctx *gin.Context){
	key := ctx.Param("key")
	msg := kafkaMsg{
		Action: "DELETE",
		Key: key,
	}
	strMsg, err := json.Marshal(&msg)
	if err != nil {
		response.Error(ctx, errs.NewError(http.StatusConflict, "Can't convert payload struct"))
	}
	h.producer.Publish(os.Getenv("TOPIC"), strMsg)
}
