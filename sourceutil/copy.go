/*

Copyright (c) 2022 - Present. Blend Labs, Inc. All rights reserved
Use of this source code is governed by a MIT license that can be found in the LICENSE file.

*/

package sourceutil

import (
	"context"
	"io"
	"os"
)

// Copy copies a file from a source to a destination.
func Copy(ctx context.Context, destination, source string) error {
	// Debugf(ctx, "copying %s to %s", source, destination)
	sourceReader, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceReader.Close()
	destinationWriter, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destinationWriter.Close()
	_, err = io.Copy(destinationWriter, sourceReader)
	if err != nil {
		return err
	}
	return nil
}
