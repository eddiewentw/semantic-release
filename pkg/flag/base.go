package flag

const DRY_RUN_FLAG = "--dry-run"
const FIRST_RELEASE_FLAG = "--first-release"

type Args struct {
	IsDryRun       bool
	IsFirstRelease bool
}
