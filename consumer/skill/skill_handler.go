package skill

type SkillHandler struct {
	storager SkillStorager
}

func NewSkillHandler(storager SkillStorager) *SkillHandler {
	return &SkillHandler{storager: storager}
}

func (h *SkillHandler) ExtractMsg(message kafkaMsg) {
	switch message.Action {
	case "INSERT":
		h.CreateSkill(message)
	case "UPDATE":
		h.UpdateSkill(message)
	case "PATCH-NAME":
		h.UpdateSkillNameByKey(message)
	case "PATCH-DESCRIPTION":
		h.UpdateSkillDescriptionByKey(message)
	case "PATCH-LOGO":
		h.UpdateSkillLogoByKey(message)
	case "PATCH-TAGS":
		h.UpdateSkillTagsByKey(message)
	case "DELETE":
		h.DeleteSkill(message)
	}

}

func (h *SkillHandler) CreateSkill(message kafkaMsg) error {
	_, err := h.storager.CreateSkill(message.Data)
	if err != nil {
		return err
	}
	return nil
}

func (h *SkillHandler) UpdateSkill(message kafkaMsg) error {
	_, err := h.storager.UpdateSkill(message.Data, message.Key)
	if err != nil {
		return err
	}
	return nil
}

func (h *SkillHandler) UpdateSkillNameByKey(message kafkaMsg) error {
	_, err := h.storager.UpdateSkillNameByKey(message.Key, message.Data.Name)
	if err != nil {
		return err
	}
	return nil
}

func (h *SkillHandler) UpdateSkillDescriptionByKey(message kafkaMsg) error {
	_, err := h.storager.UpdateSkillDescriptionByKey(message.Key, message.Data.Description)
	if err != nil {
		return err
	}
	return nil
}

func (h *SkillHandler) UpdateSkillLogoByKey(message kafkaMsg) error {
	_, err := h.storager.UpdateSkillLogoByKey(message.Key, message.Data.Logo)
	if err != nil {
		return err
	}
	return nil
}

func (h *SkillHandler) UpdateSkillTagsByKey(message kafkaMsg) error {
	_, err := h.storager.UpdateSkillTagsByKey(message.Key, message.Data.Tags)
	if err != nil {
		return err
	}
	return nil
}

func (h *SkillHandler) DeleteSkill(message kafkaMsg) error {
	err := h.storager.DeleteSkill(message.Key)
	if err != nil {
		return err
	}
	return nil
}
