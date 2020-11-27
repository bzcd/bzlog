# bzlog

## install

```
go get -u github.com/bzcd/bzlog
```

## use

### options

```
l := bzlog.New(&bzlog.Options{
    Output: "stdout",  // default stdout
    TimeFormat: "2006-01-02 15:0:05", // time format
    Level: "info",  // log level
})

l.WithFields("key", "hello", "value", "world").Info("WELCOME")

// Output
// bzlog.OutputStdout (default)
// bzlog.OutputStderr
// bzlog.OutputFile  # write to file <FileName & Size>

// TimeFormat
// iso8601 (default)
// {"level":"INFO","ts":"2020-11-27T13:54:17.175+0800","caller":"bzlog/bzlog_test.go:64","msg":"WELCOME","key":"hello","value":"world"}
// epoch_time_encoder, epoch_time
// {"level":"INFO","ts":1606456601.699569,"caller":"bzlog/bzlog_test.go:65","msg":"WELCOME","key":"hello","value":"world"}
// epoch_millis_time_encoder, epoch_millis_time
// {"level":"INFO","ts":1606456622140.622,"caller":"bzlog/bzlog_test.go:65","msg":"WELCOME","key":"hello","value":"world"}
// rfc3339
{"level":"INFO","ts":"2020-11-27T13:57:25+08:00","caller":"bzlog/bzlog_test.go:65","msg":"WELCOME","key":"hello","value":"world"}
// rfc3339nano
// {"level":"INFO","ts":"2020-11-27T13:57:47.368494+08:00","caller":"bzlog/bzlog_test.go:65","msg":"WELCOME","key":"hello","value":"world"}
// 2006-01-02 15:04:05 (user custom)
// {"level":"INFO","ts":"2020-11-27 13:59:13","caller":"bzlog/bzlog_test.go:65","msg":"WELCOME","key":"hello","value":"world"}

// Format
// bzlog.FormatJSON (default)
// {"level":"INFO","ts":"2020-11-27 13:0:41","caller":"bzlog/bzlog_test.go:65","msg":"WELCOME","key":"hello","value":"world"}
// bzlog.FormatText
// 2020-11-27 13:0:18	INFO	bzlog/bzlog_test.go:66	WELCOME	{"key": "hello", "value": "world"}

// Level
// bzlog.LevelDebug (default)
// bzlog.LevelDebug
// bzlog.LevelInfo
// bzlog.LevelWarn
// bzlog.LevelError
// bzlog.LevelFatal
// bzlog.LevelPanic

// Output == bzlog.Outputfile
// FileName # file name
// Size     # file size

// NoAddCaller
// false (default)
// true
// {"level":"INFO","ts":"2020-11-27 14:03:38","msg":"WELCOME","key":"hello","value":"world"}
```

### add log

```
# WithMap
kvs := map[string]interface{} {
    "action": "push",
    "value": 123,
}
l.WithMap(kvs).Info("WELCOME")

// {"level":"INFO","ts":"2020-11-27 14:07:05","caller":"bzlog/bzlog_test.go:68","msg":"WELCOME","action":"push","value":123}

# WithFields
l.WithFields("action", "push", "value", 123).Info("WELCOME")
// {"level":"INFO","ts":"2020-11-27 14:08:41","caller":"bzlog/bzlog_test.go:64","msg":"WELCOME","action":"push","value":123}
```

### custom your logger

```
# SetDefaultLog
var lg *zap.Logger
# lg := zap.New???
bzlog.SetDefaultLog(lg)

# AddCallerSkip
type LocalLog struct {
    l bzlog.Logger
    // ...
}

var ll = &LocalLog{
    l: bzlog.New(...).AddCallerSkip(1),
    // ...
}

func (l *LocalLog) Info(msg string) {
    // do something
    l = l.l
    // l = l.WithFields(...)
    l.Info(msg)
}
```
