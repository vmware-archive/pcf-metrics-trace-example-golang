package acceptance_test

import (
	"fmt"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"github.com/pivotal-cf/pcf-metrics-trace-example-golang/testhelpers"
)

var _ = Describe("Endpoints", func() {
	var (
		port           uint16
		paymentsServer *gexec.Session
		ordersServer   *gexec.Session
		env            []string
	)

	BeforeEach(func() {
		port = testhelpers.GetOpenPort()
		env = []string{
			fmt.Sprintf("PORT=%d", port),
		}
	})

	Describe("Payments", func() {
		BeforeEach(func() {
			paymentsServer = testhelpers.StartGoProcess("github.com/pivotal-cf/pcf-metrics-trace-example-golang/payments", env)
		})

		AfterEach(func() {
			paymentsServer.Terminate()
		})

		It("returns 200", func() {
			resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/charge-card", port))
			Expect(err).ToNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
		})
	})

	Describe("Orders", func() {
		var (
			ordersPort uint16
		)
		BeforeEach(func() {
			ordersPort = testhelpers.GetOpenPort()
			paymentsPort := testhelpers.GetOpenPort()

			ordersEnv := []string{
				fmt.Sprintf("PORT=%d", ordersPort),
				fmt.Sprintf("PAYMENTS_HOST=127.0.0.1:%d", paymentsPort),
			}

			paymentsEnv := []string{
				fmt.Sprintf("PORT=%d", paymentsPort),
			}

			ordersServer = testhelpers.StartGoProcess("github.com/pivotal-cf/pcf-metrics-trace-example-golang/orders", ordersEnv)
			paymentsServer = testhelpers.StartGoProcess("github.com/pivotal-cf/pcf-metrics-trace-example-golang/payments", paymentsEnv)
		})

		AfterEach(func() {
			ordersServer.Terminate()
			paymentsServer.Terminate()
		})

		It("returns 200", func() {
			resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/process-order", ordersPort))
			Expect(err).ToNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
		})
	})

	Describe("Shopping Cart", func() {
		var (
			shoppingCartPort uint16

			shoppingCartServer *gexec.Session
			paymentsServer     *gexec.Session
			ordersServer       *gexec.Session
		)
		BeforeEach(func() {
			shoppingCartPort = testhelpers.GetOpenPort()
			ordersPort := testhelpers.GetOpenPort()
			paymentsPort := testhelpers.GetOpenPort()

			shoppingCartEnv := []string{
				fmt.Sprintf("PORT=%d", shoppingCartPort),
				fmt.Sprintf("ORDERS_HOST=127.0.0.1:%d", ordersPort),
			}

			ordersEnv := []string{
				fmt.Sprintf("PORT=%d", ordersPort),
				fmt.Sprintf("PAYMENTS_HOST=127.0.0.1:%d", paymentsPort),
			}

			paymentsEnv := []string{
				fmt.Sprintf("PORT=%d", paymentsPort),
			}

			shoppingCartServer = testhelpers.StartGoProcess("github.com/pivotal-cf/pcf-metrics-trace-example-golang/shopping_cart", shoppingCartEnv)
			ordersServer = testhelpers.StartGoProcess("github.com/pivotal-cf/pcf-metrics-trace-example-golang/orders", ordersEnv)
			paymentsServer = testhelpers.StartGoProcess("github.com/pivotal-cf/pcf-metrics-trace-example-golang/payments", paymentsEnv)
		})

		AfterEach(func() {
			shoppingCartServer.Terminate()
			ordersServer.Terminate()
			paymentsServer.Terminate()
		})

		It("returns 200", func() {
			resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/checkout", shoppingCartPort))
			Expect(err).ToNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
		})
	})
})
