package log

import (
    "fmt"
    "os"

    "github.com/lowellmower/audmon/pkg/config"

    "github.com/moogar0880/venom"
    "github.com/sirupsen/logrus"
)

// Logger defines a set of methods for writing application logs. Derived from and
// inspired by logrus.Entry.
type Logger interface {
    Debug(args ...interface{})
    Debugf(format string, args ...interface{})
    Debugln(args ...interface{})
    Error(args ...interface{})
    Errorf(format string, args ...interface{})
    Errorln(args ...interface{})
    Fatal(args ...interface{})
    Fatalf(format string, args ...interface{})
    Fatalln(args ...interface{})
    Info(args ...interface{})
    Infof(format string, args ...interface{})
    Infoln(args ...interface{})
    Panic(args ...interface{})
    Panicf(format string, args ...interface{})
    Panicln(args ...interface{})
    Print(args ...interface{})
    Printf(format string, args ...interface{})
    Println(args ...interface{})
    Warn(args ...interface{})
    Warnf(format string, args ...interface{})
    Warning(args ...interface{})
    Warningf(format string, args ...interface{})
    Warningln(args ...interface{})
    Warnln(args ...interface{})
}

// global log for the daemon
var Log *logrus.Logger

// init sets up our loggers for our application. As the application expands, we
// ideally want to be able to isolate logs and should make a logging mechanism
// per service or properly structure the daemon log to isolate parts of logs.
func init() {
    if Log == nil {
        Log = newLogrusLogger(config.AppConf.Registry)
    }
}

// newLogrusLogger takes a pointer to a venom config and returns a pointer to an
// instance of logrun.Logger. The configuraiton values for logging should always
// be a second level JSON struct in the configuration file.
func newLogrusLogger(cfg *venom.Venom) *logrus.Logger {
    l := logrus.New()
    if _, exist := cfg.Find("json_logs"); exist {
        l.Formatter = new(logrus.JSONFormatter)
    }

    logFile := cfg.GetString("log.file")
    fmt.Printf("LOG FILE: %s\n", logFile)
    if logFile != "" {
        file, err := os.Create(logFile)
        if err != nil {
            panic("could not create log file: " + err.Error())
        }
        l.Out = file
    } else {
        l.Out = os.Stdout
    }

    switch cfg.GetString("log.level") {
    case "info":
        l.Level = logrus.InfoLevel
    case "warning":
        l.Level = logrus.WarnLevel
    case "error":
        l.Level = logrus.ErrorLevel
    case "trace":
        l.Level = logrus.TraceLevel
    default:
        l.Level = logrus.DebugLevel
    }

    l.ReportCaller = cfg.GetBool("log.report_caller")
    return l
}

// Fields is a map string interface to define fields in the structured log
type Fields map[string]interface{}

// With allow us to define fields in out structured logs
func (f Fields) With(k string, v interface{}) Fields {
    f[k] = v
    return f
}

// WithFields allow us to define fields in out structured logs
func (f Fields) WithFields(f2 Fields) Fields {
    for k, v := range f2 {
        f[k] = v
    }
    return f
}

// WithFields allow us to define fields in out structured logs
func WithFields(fields Fields) Logger {
    return Log.WithFields(logrus.Fields(fields))
}

// Debug package-level convenience method.
func Debug(args ...interface{}) {
    Log.Debug(args...)
}

// Debugf package-level convenience method.
func Debugf(format string, args ...interface{}) {
    Log.Debugf(format, args...)
}

// Debugln package-level convenience method.
func Debugln(args ...interface{}) {
    Log.Debugln(args...)
}

// Error package-level convenience method.
func Error(args ...interface{}) {
    Log.Error(args...)
}

// Errorf package-level convenience method.
func Errorf(format string, args ...interface{}) {
    Log.Errorf(format, args...)
}

// Errorln package-level convenience method.
func Errorln(args ...interface{}) {
    Log.Errorln(args...)
}

// Fatal package-level convenience method.
func Fatal(args ...interface{}) {
    Log.Fatal(args...)
}

// Fatalf package-level convenience method.
func Fatalf(format string, args ...interface{}) {
    Log.Fatalf(format, args...)
}

// Fatalln package-level convenience method.
func Fatalln(args ...interface{}) {
    Log.Fatalln(args...)
}

// Info package-level convenience method.
func Info(args ...interface{}) {
    Log.Info(args...)
}

// Infof package-level convenience method.
func Infof(format string, args ...interface{}) {
    Log.Infof(format, args...)
}

// Infoln package-level convenience method.
func Infoln(args ...interface{}) {
    Log.Infoln(args...)
}

// Panic package-level convenience method.
func Panic(args ...interface{}) {
    Log.Panic(args...)
}

// Panicf package-level convenience method.
func Panicf(format string, args ...interface{}) {
    Log.Panicf(format, args...)
}

// Panicln package-level convenience method.
func Panicln(args ...interface{}) {
    Log.Panicln(args...)
}

// Print package-level convenience method.
func Print(args ...interface{}) {
    Log.Print(args...)
}

// Printf package-level convenience method.
func Printf(format string, args ...interface{}) {
    Log.Printf(format, args...)
}

// Println package-level convenience method.
func Println(args ...interface{}) {
    Log.Println(args...)
}

// Warn package-level convenience method.
func Warn(args ...interface{}) {
    Log.Warn(args...)
}

// Warnf package-level convenience method.
func Warnf(format string, args ...interface{}) {
    Log.Warnf(format, args...)
}

// Warning package-level convenience method.
func Warning(args ...interface{}) {
    Log.Warning(args...)
}

// Warningf package-level convenience method.
func Warningf(format string, args ...interface{}) {
    Log.Warningf(format, args...)
}

// Warningln package-level convenience method.
func Warningln(args ...interface{}) {
    Log.Warningln(args...)
}

// Warnln package-level convenience method.
func Warnln(args ...interface{}) {
    Log.Warnln(args...)
}
