package main

import (
	"job-portal-project/api/config"
	route "job-portal-project/api/route"
	migration "job-portal-project/generate/sql"
	"os"
)

// @title DMS job-portal-project API
// @version v1
// @license AGPLv3
// @description This is a DMS job-portal-project API Server.

func main() {
	args := os.Args
	env := ""
	if len(args) > 1 {
		env = args[1]
	}

	if env == "migrate" {
		migration.Migrate()
	} else if env == "generate" {
		migration.Generate()
	} else if env == "migg" {
		migration.MigrateGG()
	} else if env == "debug" {
		// config.InitEnvConfigs(false, env)
		// db := config.InitDB()
		// config.InitLogger(db)
		// redis := config.InitRedis()
		// route.CreateHandler(db, env, redis)
	} else {
		config.InitEnvConfigs(false, env)
		db := config.InitDB()
		config.InitLogger(db)
		route.StartRouting(db)
	}
}
