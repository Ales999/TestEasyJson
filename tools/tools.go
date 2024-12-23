//go:build tools
// +build tools

package tools

import (
	_ "github.com/mailru/easyjson"
	_ "github.com/mailru/easyjson/gen"
	_ "github.com/mailru/easyjson/jlexer"
	_ "github.com/mailru/easyjson/jwriter"
)
