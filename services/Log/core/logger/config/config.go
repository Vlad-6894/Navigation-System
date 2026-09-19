package log_core_logger_config

type LoggerConfig interface {
	GetLevel() string
	GetFolder() string
}
