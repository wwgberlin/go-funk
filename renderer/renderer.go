package renderer

import (
	"io"
	"time"
)

type RenderFunc func(io.Writer, []int, time.Duration, Config)
