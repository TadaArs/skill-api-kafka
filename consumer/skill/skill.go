package skill

type Skill struct {
	Key string `json:"key"`
	Name string `json:"name" default:""`
	Description string `json:"description" default:""`
	Logo string `json:"logo" default:""`
	Tags []string `json:"tags" default:"{}"`
}

type SkillCreateRequest struct {
	Key string `json:"key"`
	Name string `json:"name"`
	Description string `json:"description"`
	Logo string `json:"logo"`
	Tags []string `json:"tags"`
}

type SkillUpdateRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Logo string `json:"logo"`
	Tags []string `json:"tags"`
}

