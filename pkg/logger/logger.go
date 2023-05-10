/*
 * The Alluxio Open Foundation licenses this work under the Apache License, version 2.0
 * (the "License"). You may not use this work except in compliance with the License, which is
 * available at www.apache.org/licenses/LICENSE-2.0
 *
 * This software is distributed on an "AS IS" basis, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied, as more fully set forth in the License.
 *
 * See the NOTICE file distributed with this work for information regarding copyright ownership.
 */

package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	zapLog      *zap.Logger
	sugarZapLog *zap.SugaredLogger
)

func init() {
	var err error
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderConfig

	zapLog, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	sugarZapLog = zapLog.Sugar()
}

func Infof(format string, args ...interface{}) {
	sugarZapLog.Infof(format, args)
}

func Errorf(format string, args ...interface{}) {
	sugarZapLog.Errorf(format, args)
}

func Fatalf(format string, args ...interface{}) {
	sugarZapLog.Fatalf(format, args)
}
