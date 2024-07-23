package skill

import (
	"consumer/errs"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/lib/pq"
)

type SkillStorage struct {
	db *sql.DB
}


type SkillStorager interface {
	GetSkills() ([]Skill, error)
	GetSkillByKey(key string) (*Skill, error)
	CreateSkill(skill Skill) (*Skill, error)
	UpdateSkill(skill Skill, key string) (*Skill, error)
	DeleteSkill(key string) (*Skill, error)
}

func NewSkillStorage(db *sql.DB) *SkillStorage {
	return &SkillStorage{db: db}
}

func (s *SkillStorage) GetSkills() ([]Skill, error) {
	skills := []Skill{}
	records, err := s.db.Query("SELECT key, name, description, logo, tags FROM skills")
	if err != nil {
		return []Skill{}, errs.NewError(http.StatusInternalServerError, err.Error())
	}
	for records.Next() {
		skill := Skill{}
		err := records.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, pq.Array(&skill.Tags))
		if err != nil {
			return []Skill{}, errs.NewError(http.StatusInternalServerError, err.Error())
		}
		skills = append(skills, skill)
	}
	return skills, nil

}

func (s *SkillStorage) GetSkillByKey(key string) (*Skill, error) {
	skill := Skill{}
	record := s.db.QueryRow("SELECT key, name, description, logo, tags FROM skills WHERE key=$1", key)
	err := record.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, pq.Array(&skill.Tags))
	if err != nil {
		return &Skill{}, errs.NewError(http.StatusNotFound, "Skill not found")
	}
	return &skill, nil
}

func (s *SkillStorage) CreateSkill(skill Skill) (*Skill, error) {
	fmt.Println(skill)
	createdSkill := Skill{}
	record := s.db.QueryRow("INSERT INTO skills (key, name, description, logo, tags) VALUES ($1, $2, $3, $4, $5) RETURNING key, name, description, logo, tags", skill.Key, skill.Name, skill.Description, skill.Logo, pq.Array(skill.Tags))
	err := record.Scan(&createdSkill.Key, &createdSkill.Name, &createdSkill.Description, &createdSkill.Logo, pq.Array(&createdSkill.Tags))
	if err != nil {
		return &Skill{}, errs.NewError(http.StatusConflict, "Skill already exists")
	}
	return &createdSkill, nil
}

// DeleteSkill implements SkillStorager.
func (s *SkillStorage) DeleteSkill(key string) (*Skill, error) {
	panic("unimplemented")
}

// UpdateSkill implements SkillStorager.
func (s *SkillStorage) UpdateSkill(skill Skill, key string) (*Skill, error) {
	panic("unimplemented")
}