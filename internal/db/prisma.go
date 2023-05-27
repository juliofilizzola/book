package internal

import (
	"api/prisma/db"
)

func PrismaClientDB() (*db.PrismaClient, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return nil, err
	}
	return client, nil
}
