package gorc_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGorc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gorc Suite")
}
