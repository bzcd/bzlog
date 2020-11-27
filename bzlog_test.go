package bzlog_test

import (
	"testing"

	"github.com/bzcd/bzlog"
)

func TestBzLog(t *testing.T) {
	out(&bzlog.Options{})
	out(&bzlog.Options{
		Format:     bzlog.FormatText,
		TimeFormat: "2006-01-02 15:04:05",
	})
	out(&bzlog.Options{
		Output:     "file",
		TimeFormat: "epoch_time",
		Level:      "info",

		FileName: "1.log",
		Size:     1,
	})
}

func TestDefault(t *testing.T) {
	out2(bzlog.GetDefaultLog())

	l := bzlog.New(&bzlog.Options{
		Format: bzlog.FormatText,
	})

	bzlog.SetDefaultLog(l.WithFields("app_name", "bzlog"))
	out2(bzlog.GetDefaultLog())
}

func out(opt *bzlog.Options) {
	l := bzlog.New(opt)
	out2(l)
}

func out2(l bzlog.Logger) {
	l.Debug("debug")
	l.Info("info")
	l.Warn("warn")
	l.Error("error")

	l0 := l.WithFields("wf0", "va0")
	l1 := l.WithFields("wf1", "va1")

	l1.Info("l1.Info")
	l0.Info("l0.Info")
	l.Info("l.Info")

	l.WithFields("p0", "v0", "p1", 1, "p2", 2, 3).
		Error("ok")
}

func TestF(t *testing.T) {
	l := bzlog.New(&bzlog.Options{
		Output:     "stdout", // default stdout
		Level:      "info",   // log level
		TimeFormat: "2006-01-02 15:04:05",
	})
	l.WithFields("action", "push", "value", 123).Info("WELCOME")
}

func TestFormatSize(t *testing.T) {
	const defaultSize = 50 * 1024 * 1024

	_formatSize(t, "1024", 1024)
	_formatSize(t, "", defaultSize)
	_formatSize(t, "1024a", defaultSize)
	_formatSize(t, "1024k", 1024*1024)
	_formatSize(t, "1024K", 1024*1024)
	_formatSize(t, "1024m", 1024*1024*1024)
	_formatSize(t, "1024M", 1024*1024*1024)
	_formatSize(t, "1g", 1024*1024*1024)
	_formatSize(t, "1G", 1024*1024*1024)
	_formatSize(t, "G", 1024*1024*1024)
	_formatSize(t, "1kG", defaultSize)
	_formatSize(t, "1 G", defaultSize)
	_formatSize(t, "1 ", 1)
}

func _formatSize(t *testing.T, sizeString string, cmp int) {
	val := bzlog.FormatSize(sizeString)
	if val != cmp {
		t.Fatal("call FormatSize Faield, give", sizeString, ", get", val, ", want", cmp)
	}
}
