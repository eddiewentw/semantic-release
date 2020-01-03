package flag

import (
	"os"
)

func Parse() *Args {
	args := Args{}

	for _, argument := range os.Args[1:] {
		if args.IsDryRun == false && argument == DRY_RUN_FLAG {
			args.IsDryRun = true
		}

		if args.IsFirstRelease == false && argument == FIRST_RELEASE_FLAG {
			args.IsFirstRelease = true
		}
	}

	return &args
}
