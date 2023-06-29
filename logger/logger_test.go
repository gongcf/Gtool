package logger

import (
	"testing"
)

func TestSetLevel(t *testing.T) {
	Log.SetLevel("trace")
}

func TestTrace(t *testing.T) {
	Log.SetLevel("trace")
	Glogger.Trace("trace")
	Glogger.SetLevel("off")
	Glogger.Trace("trace")
}

func TestTracef(t *testing.T) {
	Glogger.SetLevel("trace")
	Glogger.Tracef("tracef")
	Glogger.SetLevel("off")
	Glogger.Tracef("tracef")
}

func TestDebug(t *testing.T) {
	Glogger.SetLevel("debug")
	Glogger.Debug("debug")
	Glogger.SetLevel("off")
	Glogger.Debug("debug")
}

func TestDebugf(t *testing.T) {
	Glogger.SetLevel("debug")
	Glogger.Debugf("debugf")
	Glogger.SetLevel("off")
	Glogger.Debug("debug")
}

func TestInfo(t *testing.T) {
	Glogger.SetLevel("info")
	Glogger.Info("info")
	Glogger.SetLevel("off")
	Glogger.Info("info")
}

func TestInfof(t *testing.T) {
	Glogger.SetLevel("info")
	Glogger.Infof("infof")
	Glogger.SetLevel("off")
	Glogger.Infof("infof")
}

func TestWarn(t *testing.T) {
	Glogger.SetLevel("warn")
	Glogger.Warn("warn")
	Glogger.SetLevel("off")
	Glogger.Warn("warn")
}

func TestWarnf(t *testing.T) {
	Glogger.SetLevel("warn")
	Glogger.Warnf("warnf")
	Glogger.SetLevel("off")
	Glogger.Warnf("warnf")
}

func TestError(t *testing.T) {
	Glogger.SetLevel("error")
	Glogger.Error("error")
	Glogger.SetLevel("off")
	Glogger.Error("error")
}

func TestErrorf(t *testing.T) {
	Glogger.SetLevel("error")
	Glogger.Errorf("errorf")
	Glogger.SetLevel("off")
	Glogger.Errorf("errorf")
}

func TestGetLevel(t *testing.T) {
	if getLevel("trace") != Trace {
		t.FailNow()

		return
	}

	if getLevel("debug") != Debug {
		t.FailNow()

		return
	}

	if getLevel("info") != Info {
		t.FailNow()

		return
	}

	if getLevel("warn") != Warn {
		t.FailNow()

		return
	}

	if getLevel("error") != Error {
		t.FailNow()

		return
	}
}

func TestLoggerSetLevel(t *testing.T) {
	Glogger.SetLevel("trace")

	if Glogger.level != Trace {
		t.FailNow()

		return
	}
}

func TestIsTraceEnabled(t *testing.T) {
	Glogger.SetLevel("trace")

	if !Glogger.IsTraceEnabled() {
		t.FailNow()

		return
	}
}

func TestIsDebugEnabled(t *testing.T) {
	Glogger.SetLevel("debug")

	if !Glogger.IsDebugEnabled() {
		t.FailNow()

		return
	}
}

func TestIsWarnEnabled(t *testing.T) {
	Glogger.SetLevel("warn")

	if !Glogger.IsWarnEnabled() {
		t.FailNow()

		return
	}
}
