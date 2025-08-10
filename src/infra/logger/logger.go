package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/labstack/gommon/log"
	"github.com/rs/zerolog"
)

const logSeparator = " "

type EchoZLogger struct {
	zLogger *zerolog.Logger
}

func NewEchoZLogger(zLogger *zerolog.Logger) EchoZLogger {
	return EchoZLogger{zLogger: zLogger}
}

func (e EchoZLogger) GetZLogger() *zerolog.Logger {
	return e.zLogger
}

func (e EchoZLogger) Output() io.Writer {
	panic("not implemented")
}

func (e EchoZLogger) SetOutput(w io.Writer) {
	panic("not implemented")
}

func (e EchoZLogger) Prefix() string {
	panic("not implemented")
}

func (e EchoZLogger) SetPrefix(p string) {
	panic("not implemented")
}

func (e EchoZLogger) Level() log.Lvl {
	return convertLevelToLog(e.zLogger.GetLevel())
}

func (e EchoZLogger) SetLevel(v log.Lvl) {
	e.zLogger.Level(convertLevelToZerolog(v))
}

func (e EchoZLogger) SetHeader(h string) {
	panic("not implemented")
}

func (e EchoZLogger) Print(i ...interface{}) {
	e.zLogger.Print(i...)
}

func (e EchoZLogger) Printf(format string, args ...interface{}) {
	e.zLogger.Printf(format, args...)
}

func (e EchoZLogger) Printj(j log.JSON) {
	e.zLogger.Print(jsonToString(j))
}

func (e EchoZLogger) Debug(i ...interface{}) {
	e.zLogger.Debug().Msg(buildMessage(logSeparator, i...))
}

func (e EchoZLogger) Debugf(format string, args ...interface{}) {
	e.zLogger.Debug().Msgf(format, args...)
}

func (e EchoZLogger) Debugj(j log.JSON) {
	e.zLogger.Debug().Msg(jsonToString(j))
}

func (e EchoZLogger) Info(i ...interface{}) {
	e.zLogger.Info().Msg(buildMessage(logSeparator, i...))
}

func (e EchoZLogger) Infof(format string, args ...interface{}) {
	e.zLogger.Info().Msgf(format, args...)
}

func (e EchoZLogger) Infoj(j log.JSON) {
	e.zLogger.Info().Msg(jsonToString(j))
}

func (e EchoZLogger) Warn(i ...interface{}) {
	e.zLogger.Warn().Msg(buildMessage(logSeparator, i...))
}

func (e EchoZLogger) Warnf(format string, args ...interface{}) {
	e.zLogger.Warn().Msgf(format, args...)
}

func (e EchoZLogger) Warnj(j log.JSON) {
	e.zLogger.Warn().Msg(jsonToString(j))
}

func (e EchoZLogger) Error(i ...interface{}) {
	e.zLogger.Error().Msg(buildMessage(logSeparator, i...))
}

func (e EchoZLogger) Errorf(format string, args ...interface{}) {
	e.zLogger.Error().Msgf(format, args...)
}

func (e EchoZLogger) Errorj(j log.JSON) {
	e.zLogger.Error().Msg(jsonToString(j))
}

func (e EchoZLogger) Fatal(i ...interface{}) {
	e.zLogger.Fatal().Msg(buildMessage(logSeparator, i...))
}

func (e EchoZLogger) Fatalj(j log.JSON) {
	e.zLogger.Fatal().Msg(jsonToString(j))
}

func (e EchoZLogger) Fatalf(format string, args ...interface{}) {
	e.zLogger.Fatal().Msgf(format, args...)
}
func (e EchoZLogger) Panic(i ...interface{}) {
	e.zLogger.Panic().Msg(buildMessage(logSeparator, i...))
}

func (e EchoZLogger) Panicj(j log.JSON) {
	e.zLogger.Panic().Msg(jsonToString(j))
}

func (e EchoZLogger) Panicf(format string, args ...interface{}) {
	e.zLogger.Panic().Msgf(format, args...)
}

func convertLevelToZerolog(lvl log.Lvl) zerolog.Level {
	switch lvl {
	case log.DEBUG:
		return zerolog.DebugLevel
	case log.INFO:
		return zerolog.InfoLevel
	case log.WARN:
		return zerolog.WarnLevel
	case log.ERROR:
		return zerolog.ErrorLevel
	case log.OFF:
		return zerolog.Disabled
	default:
		return zerolog.InfoLevel
	}
}

func convertLevelToLog(lvl zerolog.Level) log.Lvl {
	switch lvl {
	case zerolog.DebugLevel:
		return log.DEBUG
	case zerolog.InfoLevel:
		return log.INFO
	case zerolog.WarnLevel:
		return log.WARN
	case zerolog.ErrorLevel:
		return log.ERROR
	case zerolog.Disabled:
		return log.OFF
	default:
		return log.INFO
	}
}

func jsonToString(j log.JSON) string {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func buildMessage(separator string, itrs ...interface{}) string {
	m := make([]string, len(itrs))
	for i, itr := range itrs {
		m[i] = fmt.Sprintf("%v", itr)
	}
	return strings.Join(m, separator)
}
