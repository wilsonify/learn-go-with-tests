package reporters

import (
	"os/exec"

	"github.com/approvals/go-approval-tests/utils"
)

// NewFrontLoadedReporter creates the default front loaded reporter.
func NewFrontLoadedReporter() Reporter {
	return NewFirstWorkingReporter(
		NewContinuousIntegrationReporter(),
	)
}

// NewDiffReporter creates the default diff reporter.
func NewDiffReporter() Reporter {
	return NewFirstWorkingReporter(
		NewBeyondCompareReporter(),
		NewIntelliJReporter(),
		NewFileMergeReporter(),
		NewVSCodeReporter(),
		NewGoLandReporter(),
		NewRealDiffReporter(),
		NewPrintSupportedDiffProgramsReporter(),
		NewQuietReporter(),
	)
}

func launchProgram(programName, approved string, args ...string) bool {
	if !utils.DoesFileExist(programName) {
		return false
	}

	utils.EnsureExists(approved)

	cmd := exec.Command(programName, args...)
	cmd.Start()
	return true
}
