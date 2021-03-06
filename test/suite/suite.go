package suite

import (
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"

	nuclio "github.com/nuclio/nuclio-sdk"
	"github.com/nuclio/nuclio/pkg/cmdrunner"
	nucliozap "github.com/nuclio/nuclio/pkg/zap"
)

type NuclioTestSuite struct {
	suite.Suite

	Logger         nuclio.Logger
	Cmd            cmdrunner.CmdRunner
	NuclioRootPath string
}

func (suite *NuclioTestSuite) gitRoot() string {
	out, err := suite.Cmd.Run(nil, "git rev-parse --show-toplevel")
	suite.Require().NoError(err, "Can't create command runner")
	return strings.TrimSpace(out)
}

func (suite *NuclioTestSuite) SetupSuite() {
	zap, err := nucliozap.NewNuclioZapTest("nuclio-test")
	suite.Require().NoError(err, "Can't create logger")
	suite.Logger = zap
	cmd, err := cmdrunner.NewShellRunner(suite.Logger)
	suite.Require().NoError(err, "Can't create command runner")
	suite.Cmd = cmd
	suite.NuclioRootPath = suite.gitRoot()
}

// Run takes a testing suite and runs all of the tests attached to it.
// Here so you won't need to import "github.com/stretchr/testify/suite" as well
func Run(t *testing.T, testSuite suite.TestingSuite) {
	suite.Run(t, testSuite)
}
