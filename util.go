/*
 * Author: zsh
 * Create: 2020/2/14
 */

package ezlog

import (
	"bytes"
	"errors"
	"os"
)

// Check if regular file exists, regular file is not
// ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice | ModeCharDevice | ModeIrregular
func isFileExist(fileName string) (bool, error) {
	stat, err := os.Stat(fileName)
	if err == nil {
		if stat.Mode() & os.ModeType == 0 {
			return true, nil
		}
		return false, errors.New(fileName + " exists but not regular file")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Create a directory
func mkdir(dir string) error {
	_, err := os.Stat(dir)
	if !(err == nil || os.IsExist(err)) {
		if err := os.Mkdir(dir, 0775); err != nil {
			if os.IsPermission(err) {
				return err
			}
		}
	}
	return nil;
}

// Cheap integer to fixed-width decimal ASCII. Give a negative width to avoid zero-padding.
// if wid < 0 , Width is determined by numbers
func itoa(buf *bytes.Buffer, i int, wid int) {
	var u uint = uint(i)
	if u== 0 && wid <= 1 {
		buf.WriteByte('0')
		return
	}
	// Assemble decimal in reverse order.
	var b [32]byte
	bp := len(b)
	for ; u > 0 || wid > 0; u /= 10 {
		bp--
		wid--
		b[bp] = byte(u%10) + '0'
	}
	for bp < len(b) {
		buf.WriteByte(b[bp])
		bp++
	}
}

