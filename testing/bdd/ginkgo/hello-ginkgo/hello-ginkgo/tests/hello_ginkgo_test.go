package tests

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "hello-ginkgo/src"
)

var _ = Describe("Hello Function", func() {
    Context("When called", func() {
        It("should return a greeting message", func() {
            greeting := main.Hello()
            Expect(greeting).To(Equal("Hello, Ginkgo!"))
        })
    })
})