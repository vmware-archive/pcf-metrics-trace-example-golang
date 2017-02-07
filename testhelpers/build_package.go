package testhelpers

import (
	"os/exec"
	"time"

	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func StartGoProcess(importPath string, env []string, args ...string) *gexec.Session {
	commandPath, err := gexec.Build(importPath, "-race")
	Expect(err).ToNot(HaveOccurred())

	command := exec.Command(commandPath, args...)
	command.Env = env
	return runCommand(command)
}

func runCommand(cmd *exec.Cmd) *gexec.Session {
	var session *gexec.Session
	Eventually(func() error {
		var err error
		session, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		return err
	}).Should(BeNil())

	Consistently(session.Exited).ShouldNot(BeClosed())

	time.Sleep(time.Second)

	return session
}
