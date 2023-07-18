package config

import (
	"log"
	"momentum-core/utils"
	"os"

	"github.com/google/uuid"
)

const TRACE_PREFIX = "TRACE  :"
const INFO_PREFIX = "INFO   :"
const WARN_PREFIX = "WARNING:"
const ERROR_PREFIX = "ERROR  :"

type OneFileLoggerClient struct {
	ILoggerClient

	traceLog *log.Logger
	infoLog  *log.Logger
	warnLog  *log.Logger
	errorLog *log.Logger
}

func NewLogger(path string) (ILoggerClient, error) {

	logger := new(OneFileLoggerClient)

	logPath := utils.BuildPath(path, "log.txt")
	writer, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	logger.traceLog = log.New(writer, TRACE_PREFIX, log.Ldate|log.Ltime) // |log.Lshortfile
	logger.infoLog = log.New(writer, INFO_PREFIX, log.Ldate|log.Ltime)   // |log.Lshortfile
	logger.warnLog = log.New(writer, WARN_PREFIX, log.Ldate|log.Ltime)   // |log.Lshortfile
	logger.errorLog = log.New(writer, ERROR_PREFIX, log.Ldate|log.Ltime) // |log.Lshortfile

	logger.traceLog.SetOutput(writer)
	logger.infoLog.SetOutput(writer)
	logger.warnLog.SetOutput(writer)
	logger.errorLog.SetOutput(writer)

	return logger, nil
}

func (logger *OneFileLoggerClient) LogTrace(msg string, traceId string) {
	logger.traceLog.Println(traceId, msg)
}

func (logger *OneFileLoggerClient) LogInfo(msg string, traceId string) {
	logger.infoLog.Println(traceId, msg)

}

func (logger *OneFileLoggerClient) LogWarning(msg string, err error, traceId string) {
	logger.warnLog.Println(traceId, msg, err.Error())

}

func (logger *OneFileLoggerClient) LogError(msg string, err error, traceId string) {
	logger.errorLog.Println(traceId, msg, err.Error())
}

func (logger *OneFileLoggerClient) TraceId() string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		logger.errorLog.Println("", err, "NO TRACE ID")
	}
	return uuid.String()
}
