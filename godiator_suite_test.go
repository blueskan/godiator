package godiator_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGodiator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Godiator Suite")
}
