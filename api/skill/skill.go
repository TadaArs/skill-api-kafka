package skill

type Skill struct {
	Key string `json:"key"`
	Name string `json:"name" default:""`
	Description string `json:"description" default:""`
	Logo string `json:"logo" default:""`
	Tags []string `json:"tags" default:"{}"`
}

type NameUpdateRequest struct {
	Name string `json:"name"`
}

type DescriptionUpdateRequest struct {
	Description string `json:"description"`
}

type LogoUpdateRequest struct {
	Logo string `json:"logo"`
}

type TagsUpdateRequest struct {
	Tags []string `json:"tags"`
}

