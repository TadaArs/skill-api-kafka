package skill

import "log"


type SkillHandler struct {
	storager SkillStorager

}

func NewSkillHandler(storager SkillStorager) *SkillHandler {
	return &SkillHandler{storager: storager}
}

func (h *SkillHandler) ExtractMsg(message kafkaMsg) {
	switch message.Action {
	case "INSERT":
		_, err := h.storager.CreateSkill(message.Data)
		if err != nil {
			log.Printf("Failed to insert skill: %v", err)
		}
	// case "Update":
	// 	if _, err := a.storage.EditSkill(message.Data); err != nil {
	// 		log.Printf("Failed to update skill: %v", err)
	// 	}
	// case "UpdateName":
	// 	if _, err := a.storage.EditSkillName(message.Key, message.Data.Name); err != nil {
	// 		log.Printf("Failed to update skill name: %v", err)
	// 	}
	// case "UpdateDescription":
	// 	if _, err := a.storage.EditSkillDescription(message.Key, message.Data.Description); err != nil {
	// 		log.Printf("Failed to update skill description: %v", err)
	// 	}
	// case "UpdateLogo":
	// 	if _, err := a.storage.EditSkillLogo(message.Key, message.Data.Logo); err != nil {
	// 		log.Printf("Failed to update skill logo: %v", err)
	// 	}
	// case "UpdateTags":
	// 	if _, err := a.storage.EditSkillTags(message.Key, message.Data.Tags); err != nil {
	// 		log.Printf("Failed to update skill tags: %v", err)
	// 	}
	// case "DeleteSkill":
	// 	if res := a.storage.DeleteSkill(message.Key); res != "success" {
	// 		log.Printf("Failed to delete skill")
	// 	}
	// default:
	// 	log.Printf("Unknown action: %s", message.Action)
	// }
}
}
