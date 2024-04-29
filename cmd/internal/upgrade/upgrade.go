//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package upgrade

import (
	"errors"
	"github.com/energye/energy/v2/cmd/internal/command"
)

// Upgrade 升级当前版本
func Upgrade(c *command.Config) (err error) {
	//wd, _ := os.Getwd()
	//gomod := filepath.Join(wd, "go.mod")
	//data, _ := ioutil.ReadFile(gomod)
	//fl, err := modfile.Parse(gomod, data, nil)
	return errors.New("no added features")
}
