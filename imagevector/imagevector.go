// SPDX-FileCopyrightText: Contributors to the Gardener project
//
// SPDX-License-Identifier: Apache-2.0

package imagevector

import (
	_ "embed"

	"github.com/gardener/gardener/pkg/utils/imagevector"
	"k8s.io/apimachinery/pkg/util/runtime"
)

var (
	//go:embed images.yaml
	imagesYAML string

	imageVector imagevector.ImageVector
)

func init() {
	var err error

	imageVector, _, err = imagevector.Read([]byte(imagesYAML))
	runtime.Must(err)

	imageVector, _, err = imagevector.WithEnvOverride(imageVector, nil, imagevector.OverrideEnv)
	runtime.Must(err)
}

// ImageVector returns the image vector that contains all images for this extension.
func ImageVector() imagevector.ImageVector {
	return imageVector
}
