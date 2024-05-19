package repositories

import (
	"github.com/colibri-project-io/colibri-sdk-go/pkg/base/logging"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/base/test"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/database/sqlDB"
	"os"
	"testing"
)

var (
	pc       *test.PostgresContainer
	basePath string
)

func TestMain(m *testing.M) {
	if err := os.Setenv("MIGRATION_SOURCE_URL", "../../../migrations"); err != nil {
		logging.Fatal("could not set MIGRATION_SOURCE_URL: %v", err)
	}

	test.InitializeSqlDBTest()

	basePath = test.MountAbsolutPath("../../../development-environment/database/tests-dataset/")
	pc = test.UsePostgresContainer()

	sqlDB.Initialize()
	m.Run()
}
