package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"log/slog"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	"backend/app/model"
)


func CHECK_ENV() bool {
	LOAD_ENV()
	if os.Getenv("POSTGRES_URL") == "" {
		LOAD_ENV()
	}

	return os.Getenv("POSTGRES_URL") != ""
}

func LOAD_ENV() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error loading .env file")
		}
	}
}

func GET_CONFIG() (model.CONFIG, error, string) {
	var config model.CONFIG
	data, err := ioutil.ReadFile(os.Getenv("BACKEND_CONFIG"))
	if err != nil {
		slog.Debug("GET_CONFIG()", "err", "Failed To Get Backend Config")
		return config, err, GetLocation()
	}
	if err := json.Unmarshal(data, &config); err != nil {
		slog.Debug("GET_CONFIG()", "err", "Error unmarshalling JSON : "+fmt.Sprintf(err.Error()))
		return config, err, GetLocation()
	}
	return config, nil, GetLocation()
}

func GetLocation() string {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return "unknown file and line number"
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func PRINT_LOG(TYPE, INFO, LOCATION, ERROR_MESSAGE string) {
	if TYPE == "ERROR" {

		detailLog := model.LOGGER_LINE_ERROR{
			TYPE:          TYPE,
			SERVICE:       "BACKEND SERVICE",
			MESSAGE:       INFO,
			LOCATION:      LOCATION,
			ERROR_MESSAGE: ERROR_MESSAGE,
		}
		conv, err := json.Marshal(detailLog)
		if err != nil {
			log.Println("Failed Create Log")
		} else {
			log.Println(string(conv))
		}
	}

	if TYPE == "INFO" {
		detailLog := model.LOGGER_LINE_INFO{
			TYPE:          TYPE,
			SERVICE:       "BACKEND SERVICE",
			MESSAGE:       INFO,
			LOCATION:      LOCATION,
		}
		conv, err := json.Marshal(detailLog)
		if err != nil {
			log.Println("Failed Create Log")
		} else {
			log.Println(string(conv))
		}
	}
}

// Disable Trace Method
func DISABLE_TRACE_METHOD() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodTrace {
			c.AbortWithStatus(http.StatusMethodNotAllowed)
			return
		}
		c.Next()
	}
}
