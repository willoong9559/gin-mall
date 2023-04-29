package global

import (
	"gorm.io/gorm"

	"github.com/willoong9559/gin-mall/pkg/logger"
)

var (
	Logger *logger.Logger
	DB     *gorm.DB
)
