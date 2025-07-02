package model

type CONFIG struct {
	LOG LOG `json:"LOG"`
}

type LOG struct {
	FILE_NAME         string `json:"filename"`
	SYMINK_NAME       string `json:"symlinkName"`
	ADD_SOURCE        bool   `json:"addSource"`
	MULTI_WRITER      bool   `json:"multiwriter"`
	LEVEL             int    `json:"level"`
	TIME_FORMAT       string `json:"timeFormat"`
	RETENTION_IN_DAYS int64  `json:"retentionInDays"`
	ROTATE_IN_HOUR    int64  `json:"rotateInHour"`
}

// Error Logger
type LOGGER_LINE_ERROR struct {
	TYPE          string
	SERVICE       string
	MESSAGE       string
	LOCATION      string
	ERROR_MESSAGE string
}


// Info Logger
type LOGGER_LINE_INFO struct {
	TYPE          string
	SERVICE       string
	MESSAGE       string
	LOCATION      string
}