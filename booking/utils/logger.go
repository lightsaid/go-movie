package utils

import "go.uber.org/zap"

func NewLogger(paths ...string) (*zap.Logger, error) {
	// 配置 logger
	cfg := zap.NewDevelopmentConfig()
	if len(paths) > 0 {
		cfg.OutputPaths = paths
	} else {
		cfg.OutputPaths = []string{
			"stderr",
		}
	}

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	// zap.S()和zap.L()函数是zap提供全局安全访问logger方式
	// 替换zap默认的全局logger
	_ = zap.ReplaceGlobals(logger)

	return logger, nil
}
