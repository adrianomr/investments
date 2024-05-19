package controllers

import (
	"github.com/colibri-project-io/colibri-sdk-go/pkg/base/test"
	"testing"
)

func TestMain(m *testing.M) {
	test.InitializeBaseTest()

	m.Run()
}
