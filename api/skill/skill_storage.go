package skill

import (
	"skill/errs"
	"database/sql"
	"log"
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
	UpdateSkillNameByKey(key string, name string) (*Skill, error)
	UpdateSkillDescriptionByKey(key string, description string) (*Skill, error)
	UpdateSkillLogoByKey(key string, logo string) (*Skill, error)
	UpdateSkillTagsByKey(key string, tags []string) (*Skill, error)
	DeleteSkill(key string) error
}

func NewSkillStorage(db *sql.DB) *SkillStorage {
	return &SkillStorage{db: db}
}


func Scan(record *sql.Row,skill *Skill) error{
	return record.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, pq.Array(&skill.Tags))
}

func (s *SkillStorage) GetSkills() ([]Skill, error) {
	skills := []Skill{}
	query := "SELECT key, name, description, logo, tags FROM skills"
	records, err := s.db.Query(query)
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
	query := "SELECT key, name, description, logo, tags FROM skills WHERE key=$1"
	record := s.db.QueryRow(query, key)
	err := Scan(record, &skill)
	if err != nil {
		return &Skill{}, errs.NewError(http.StatusNotFound, "Skill not found")
	}
	return &skill, nil
}

func (s *SkillStorage) CreateSkill(skill Skill) (*Skill, error) {
	createSkill := Skill{}
	query := "INSERT INTO skills (key, name, description, logo, tags) VALUES ($1, $2, $3, $4, $5) RETURNING key, name, description, logo, tags"
	record := s.db.QueryRow(query, skill.Key, skill.Name, skill.Description, skill.Logo, pq.Array(skill.Tags))
	err := Scan(record, &createSkill)
	if err != nil {
		return &Skill{}, errs.NewError(http.StatusConflict, "Skill already exists")
	}
	return &createSkill, nil
}

func (s *SkillStorage) UpdateSkill(skill Skill, key string) (*Skill, error) {
	updateSkill := Skill{}
	query := "UPDATE skills SET name=$1, description=$2, logo=$3, tags=$4 WHERE key=$5 RETURNING key, name, description, logo, tags"
	record := s.db.QueryRow(query, skill.Name, skill.Description, skill.Logo, pq.Array(skill.Tags), key)
	err := Scan(record, &updateSkill)
	if err != nil {
		log.Print(err)
		return &Skill{}, errs.NewError(http.StatusConflict, "not be able to update skil")
	}
	return &updateSkill, nil
}

func (s *SkillStorage) UpdateSkillNameByKey(key string, name string) (*Skill, error){
	updateSkill := Skill{}
	query := "UPDATE skills SET name=$1 WHERE key=$2 RETURNING key, name, description, logo, tags"
	record := s.db.QueryRow(query,name,key)
	err := Scan(record,&updateSkill)
	if err != nil {
		return &Skill{}, errs.NewError(http.StatusInternalServerError, "not be able to update skill name")
	}
	return &updateSkill, nil
}

func (s *SkillStorage)	UpdateSkillDescriptionByKey(key string, description string) (*Skill, error){
	updateSkill := Skill{}
	query := "UPDATE skills SET description=$1 WHERE key=$2 RETURNING key, name, description, logo, tags"
	record := s.db.QueryRow(query, description, key)
	err := Scan(record,&updateSkill)
	if err != nil {
		return &Skill{}, errs.NewError(http.StatusInternalServerError, "not be able to update skill description")
	}
	return &updateSkill, nil
}

func (s *SkillStorage)	UpdateSkillLogoByKey(key string, logo string) (*Skill, error){
	updatedSkill := Skill{}
	query := "UPDATE skills SET logo=$1 WHERE key=$2 RETURNING key, name, description, logo, tags"
	record := s.db.QueryRow(query, logo, key)
	err := Scan(record, &updatedSkill)
	if err != nil {
		return &Skill{}, errs.NewError(http.StatusInternalServerError, "not be able to update skill logo")
	}
	return &updatedSkill, nil
}

func (s *SkillStorage) UpdateSkillTagsByKey(key string, tags []string) (*Skill, error){
	updatedSkill := Skill{}
	query := "UPDATE skills SET tags=$1 WHERE key=$2 RETURNING key, name, description, logo, tags"
	record := s.db.QueryRow(query, pq.Array(tags), key)
	err := Scan(record, &updatedSkill)
	if err != nil {
		return &Skill{}, errs.NewError(http.StatusInternalServerError, "not be able to update skill tags")
	}
	return &updatedSkill, nil
}
func (s *SkillStorage) DeleteSkill(key string) error {
	query := "DELETE FROM skills WHERE key=$1"
	result, err := s.db.Exec(query, key)
	if err != nil {
		return errs.NewError(http.StatusInternalServerError,"not be able to delete skill")
	}

	resultRows, err := result.RowsAffected()
	if err != nil {
		return errs.NewError(http.StatusInternalServerError,"not be able to delete skill")
	}
	if resultRows == 0 {
		return errs.NewError(http.StatusInternalServerError,"not be able to delete skill")
	}

	return nil
}

