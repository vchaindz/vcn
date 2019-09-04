/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package meta

import (
	"log"
	"os"
)

// Stage is the type for all possible stage values
type Stage int64

// Allowed stage values
const (
	StageProduction Stage = 0
	StageStaging    Stage = 1
	StageTest       Stage = 2
)

// StageEnvironment returns the current Stage value
func StageEnvironment() Stage {
	switch os.Getenv("STAGE") {
	case "PRODUCTION":
		return StageProduction
	case "STAGING":
		return StageStaging
	case "TEST":
		return StageTest
	default:
		return StageProduction
	}
}

// StageName returns the name of the given stage as string
func StageName(stage Stage) string {
	switch stage {
	case StageProduction:
		return "PRODUCTION"
	case StageStaging:
		return "STAGING"
	case StageTest:
		return "TEST"
	default:
		log.Fatal("unsupported stage", stage)
		return ""
	}
}
