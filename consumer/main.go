package main

import (
	"consumer/database"
	"consumer/skill"
	"os"
)

func main() {
	db := database.Connect()
	defer db.Close()
	storage := skill.NewSkillStorage(db)

	consumer := skill.NewConsumer(os.Getenv("TOPIC"), *storage)

	defer consumer.Close()
	consumer.Consume(os.Getenv("TOPIC"))
}