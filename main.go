package main

import (
	"log"
	"github.com/Th8rHammer/kbot/cmd" // Переконайтеся, що шлях відповідає вашому go mod init
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error executing root command: %v", err)
	}
}