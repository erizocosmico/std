package log15util

import (
	"os"

	"gopkg.in/inconshreveable/log15.v2"
)

func SetRootLogger(file string, debug bool) error {
	var w = os.Stdout
	if file != "" {
		var err error
		w, err = os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
	}

	var format = log15.LogfmtFormat()
	if w == os.Stdout {
		format = log15.TerminalFormat()
	}

	var maxLvl = log15.LvlInfo
	if debug {
		maxLvl = log15.LvlDebug
	}

	log15.Root().SetHandler(log15.LvlFilterHandler(maxLvl, log15.StreamHandler(w, format)))
	return nil
}
