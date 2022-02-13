package list

import "errors"

var (
	ErrReportBug       = errors.New("logically unsound: report to github.com/xavloose/go-containers")
	ErrWrongDataType   = errors.New("unsupported data type")
	ErrOutOfBounds     = errors.New("index given is out of bounds")
	ErrCursorNotOnNode = errors.New("cursor is not on a node")
)
