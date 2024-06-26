
package util

import "log"

// CheckErr logs the error if it is not nil
func CheckErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
