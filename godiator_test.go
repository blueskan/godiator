package godiator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "godiator"
)

var _ = Describe("Godiator", func() {
	Context("When publish event in incorrect form, missing:name", func() {
		var publishResult error

		BeforeEach(func() {
			godiator := NewGodiator()

			publishResult = godiator.Publish(Event{})
		})

		It("should return error", func() {
			Expect(publishResult).To(HaveOccurred())
		})
	})

	Context("When publish event but there is no handler for this event", func() {
		var publishResult error

		BeforeEach(func() {
			godiator := NewGodiator()

			godiator.Subscribe("user:subscription", func(event Event) {})
			publishResult = godiator.Publish(Event{
				Name: "user:unsubscription",
			})
		})

		It("should return error", func() {
			Expect(publishResult).To(HaveOccurred())
		})
	})

	Context("When publishing event and there are 2 handlers for same event", func() {
		var (
			handlingUserRegistrationCount int
			handlingUserPurchaseCount     int
			publishResult                 error
		)

		BeforeEach(func() {
			godiator := NewGodiator()

			handlingUserRegistrationCount = 0
			handlingUserPurchaseCount = 0

			godiator.
				Subscribe("user:registration", func(event Event) {
					handlingUserRegistrationCount++
				}).
				Subscribe("user:registration", func(event Event) {
					handlingUserRegistrationCount++
				}).
				Subscribe("user:purchase", func(event Event) {
					handlingUserPurchaseCount++
				})

			publishResult = godiator.Publish(Event{
				Name: "user:registration",
			})
		})

		It("should not return error", func() {
			Expect(publishResult).ToNot(HaveOccurred())
		})

		It("should called twice", func() {
			Expect(handlingUserRegistrationCount).Should(Equal(2))
		})

		It("should not call irrelevant handler", func() {
			Expect(handlingUserPurchaseCount).Should(Equal(0))
		})
	})
})
