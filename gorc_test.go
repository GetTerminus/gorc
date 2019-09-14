package gorc_test

import (
	"github.com/GetTerminus/gorc"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gorc", func() {
	var cut gorc.Gorc

	BeforeEach(func() {
		cut = gorc.NewGorc()
	})

	Describe("NewGorc", func() {
		It("should make a new gorc", func() {
			Expect(cut).NotTo(BeNil())
		})
	})

	Describe("IncBy", func() {
		BeforeEach(func() {
			cut.IncBy(4)
		})

		It("should be at 4", func() {
			Expect(cut.GetCount()).To(Equal(4))
		})

		Describe("DecBy", func() {
			BeforeEach(func() {
				cut.DecBy(3)
			})

			It("should be at 1", func() {
				Expect(cut.GetCount()).To(Equal(1))
			})
		})
	})

	Describe("SetWaitMillis", func() {

	})

	Describe("WaitLow", func() {

	})

	Describe("WaitHigh", func() {

	})
})
