/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// FirstFile opens the first file of given dir for reading.
// If successful, methods on the returned file can be used for reading;
// the associated file descriptor has mode os.O_RDONLY.
// If there is an error, it will be of type *os.PathError.
func FirstFile(dir string) (io.Reader, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		return os.Open(dir + "/" + f.Name())
	}
	return nil, fmt.Errorf("empty directory: %s", dir)
}
