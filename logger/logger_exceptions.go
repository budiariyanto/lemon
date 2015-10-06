package logger

import (
	"github.com/budiariyanto/lemon/exceptions"
)

var LoggerFallbackWarning error = exceptions.Exception{Group: "Logger", Message: "Logger Fallbacked.", Description: "Logger fallbacked to standard output because file could not be opened."}
