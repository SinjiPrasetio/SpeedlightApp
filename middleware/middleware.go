package middleware

import (
	"App/data"

	"github.com/SinjiPrasetio/speedlight"
)

type Middleware struct {
	App    *speedlight.Speedlight
	Models data.Models
}
