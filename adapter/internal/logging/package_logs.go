/*
 *  Copyright (c) 2020, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

// Package logging holds the implementation for adapter logs.
package logging

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/wso2/adapter/config"
)

func logLevelMapper(pkgLevel string) logrus.Level {
	logLevel := defaultLogLevel
	switch pkgLevel {
	case warnLevel:
		logLevel = logrus.WarnLevel
	case debugLevel:
		logLevel = logrus.DebugLevel
	case errorLevel:
		logLevel = logrus.ErrorLevel
	case infoLevel:
		logLevel = logrus.InfoLevel
	case fatalLevel:
		logLevel = logrus.FatalLevel
	case panicLevel:
		logLevel = logrus.PanicLevel
	}

	return logLevel
}

// InitPackageLogger initialises the package loggers for given package name.
// If the package log level is defined in the log_config.toml file, it override the
// root log level.
func InitPackageLogger(pkgName string) *logrus.Logger {

	pkgLogLevel := defaultLogLevel //default log level
	isPackegeLevelDefined := false

	logger := logrus.New()
	logger.SetReportCaller(true)

	formatter := loggerFormat()
	logger.SetFormatter(formatter)

	logConf, errReadConfig := config.ReadLogConfigs()
	if errReadConfig != nil {
		logger.Error("Error loading log configuration. ", errReadConfig)
	}

	// Create the log file if doesn't exist. And append to it if it already exists.
	_, err := os.OpenFile(logConf.Logfile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		// Cannot open log file. Logging to stderr
		logger.Error("failed to open logfile", err)
		logger.SetOutput(os.Stdout)
	} else {
		// log output set to stdout and file
		multiWriter := io.MultiWriter(os.Stdout, setLogRotation(logConf.Logfile))
		logger.SetOutput(multiWriter)
	}

	for _, pkg := range logConf.Pkg {
		if pkg.Name == pkgName {
			pkgLogLevel = logLevelMapper(pkg.LogLevel)
			isPackegeLevelDefined = true
			break
		}
	}

	if !isPackegeLevelDefined {
		pkgLogLevel = logLevelMapper(logConf.LogLevel)
	}

	logger.SetLevel(pkgLogLevel)
	return logger
}
