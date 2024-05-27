package main

import (
	"github.com/adrianomr/investments/src/application/controllers"
	jobs "github.com/adrianomr/investments/src/application/jobs"
	"github.com/colibri-project-io/colibri-sdk-go"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/database/sqlDB"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restserver"
	"time"
)

func init() {
	colibri.InitializeApp()
	//storage.Initialize() // uncomment if you use storage
	//cacheDB.Initialize() // uncomment if you use cache
	sqlDB.Initialize() // uncomment if you use sql database
	//messaging.Initialize() // uncomment if you use messaging
	// investments
}

func main() {
	executeJobs()
	restserver.AddRoutes(controllers.NewCdbController().Routes())
	restserver.ListenAndServe()
}

func executeJobs() {
	jobsList := []jobs.Job{jobs.NewCdbsUpdateJob()}
	for _, job := range jobsList {
		go executeJob(job)
	}
}

func executeJob(job jobs.Job) {
	for {
		job.Execute()
		time.Sleep(job.ExecuteAfter())
	}
}
