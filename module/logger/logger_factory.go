/*
 Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
   SPDX-License-Identifier: Apache-2.0
*/

package logger

import (
	"strings"
	"sync"

	"go.uber.org/zap"
)

const (
	// output system.log
	ModuleDefault         = "[DEFAULT]"
	ModuleServer          = "[SERVER]"
	ModuleCli             = "[CLI]"
	ModuleStorage         = "[STORAGE]"
	ModuleP2P             = "[P2P]"
	ModuleNet             = "[NET]"
	ModuleHandler         = "[HANDLER]"
	ModuleRouter          = "[ROUTER]"
	ModuleAdapter         = "[ADAPTER]"
	ModuleProver          = "[PROVER]"
	ModuleInnerListener   = "[INNER_LISTENER]"
	ModuleWebListener     = "[WEB_LISTENER]"
	ModuleChannelListener = "[CHANNEL_LISTENER]"
	ModuleGrpcListener    = "[GRPC_LISTENER]"
	ModuleTransactionMgr  = "[TRANSACTION_MGR]"

	defaultLogPath = "./logs/default.log" // TODO release struct need this path
)

var (
	defaultLogConfig *Config
	loggers          = make(map[string]*zap.SugaredLogger)
	loggerMutex      sync.Mutex
)

// InitLogConfig set the config of logger module, called in initialization of config module
func InitLogConfig(config []*LogModuleConfig) {
	// 初始化loggers
	for _, logModuleConfig := range config {
		logPrintName := logPrintName(logModuleConfig.ModuleName)
		config := &Config{
			Module:       logPrintName,
			LogPath:      logModuleConfig.FilePath,
			LogLevel:     GetLogLevel(logModuleConfig.LogLevel),
			MaxAge:       logModuleConfig.MaxAge,
			RotationTime: logModuleConfig.RotationTime,
			JsonFormat:   false,
			ShowLine:     true,
			LogInConsole: logModuleConfig.LogInConsole,
			ShowColor:    logModuleConfig.ShowColor,
		}
		logger, _ := InitSugarLogger(config)
		loggers[logPrintName] = logger
	}
	// 最后添加"ModuleDefault"
	if _, exist := loggers[ModuleDefault]; !exist {
		// 创建默认的logger
		loggers[ModuleDefault] = getLogDefaultModuleConfig()
	}
}

// GetLogger return the instance of SugaredLogger
func GetLogger(name string) *zap.SugaredLogger {
	loggerMutex.Lock()
	defer loggerMutex.Unlock()
	logHeader := name
	logger, ok := loggers[logHeader]
	if !ok {
		logger = getLogModuleConfig(name)
		loggers[name] = logger
	}
	return logger
}

func getLogDefaultModuleConfig() *zap.SugaredLogger {
	if defaultLogConfig == nil {
		defaultLogConfig = &Config{
			Module:       ModuleDefault,
			LogPath:      defaultLogPath,
			LogLevel:     LEVEL_INFO,
			MaxAge:       DEFAULT_MAX_AGE,
			RotationTime: DEFAULT_ROTATION_TIME,
			JsonFormat:   false,
			ShowLine:     true,
			LogInConsole: true,
			ShowColor:    true,
		}
		logger, _ := InitSugarLogger(defaultLogConfig)
		return logger
	} else {
		logger, _ := InitSugarLogger(defaultLogConfig)
		return logger
	}
}

func getLogModuleConfig(moduleName string) *zap.SugaredLogger {
	innerLogConfig := &Config{
		Module:       moduleName,
		LogPath:      defaultLogPath,
		LogLevel:     LEVEL_INFO,
		MaxAge:       DEFAULT_MAX_AGE,
		RotationTime: DEFAULT_ROTATION_TIME,
		JsonFormat:   false,
		ShowLine:     true,
		LogInConsole: true,
		ShowColor:    true,
	}
	logger, _ := InitSugarLogger(innerLogConfig)
	return logger
}

func logPrintName(moduleName string) string {
	if moduleName == "" {
		return ModuleDefault
	}
	return "[" + strings.ToUpper(moduleName) + "]"
}
