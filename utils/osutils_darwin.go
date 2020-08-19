// Copyright 2019 NetApp, Inc. All Rights Reserved.

package utils

import (
	"context"
	"errors"

	. "github.com/netapp/trident/logger"
)

// The Trident build process builds the Trident CLI client for both linux and darwin.
// At compile time golang will type checks the entire code base. Since the CLI is part
// of the Trident code base this file exists to handle darwin specific code.

func getFilesystemSize(ctx context.Context, _ string) (int64, error) {

	Logc(ctx).Debug(">>>> osutils_darwin.getFilesystemSize")
	defer Logc(ctx).Debug("<<<< osutils_darwin.getFilesystemSize")
	return 0, errors.New("getFilesystemSize is not supported for darwin")
}

func GetFilesystemStats(ctx context.Context, _ string) (int64, int64, int64, int64, int64, int64, error) {

	Logc(ctx).Debug(">>>> osutils_darwin.GetFilesystemStats")
	defer Logc(ctx).Debug("<<<< osutils_darwin.GetFilesystemStats")
	return 0, 0, 0, 0, 0, 0, errors.New("GetFilesystemStats is not supported for darwin")
}

func getISCSIDiskSize(ctx context.Context, _ string) (int64, error) {

	Logc(ctx).Debug(">>>> osutils_darwin.getISCSIDiskSize")
	defer Logc(ctx).Debug("<<<< osutils_darwin.getISCSIDiskSize")
	return 0, errors.New("getBlockSize is not supported for darwin")
}
