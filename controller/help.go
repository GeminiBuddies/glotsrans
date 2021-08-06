package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type paramHelp struct {
	Name        string
	Type        string
	Default     string
	Description string
}

type paramsHelp []paramHelp

func (p paramsHelp) String() string {
	sb := &strings.Builder{}

	for _, param := range p {
		defaultString := ""
		if param.Default != "" {
			defaultString = fmt.Sprintf("[default: %7v] ", param.Default)
		}

		sb.WriteString(fmt.Sprintf("    %-8v: %6v, %v%v\n", param.Name, param.Type, defaultString, param.Description))
	}

	return sb.String()
}

type helpItem struct {
	Path        string
	Type        string
	Params      paramsHelp
	Description string
}

type helpItems []helpItem

func (h helpItems) String() string {
	sb := &strings.Builder{}

	for _, item := range h {
		sb.WriteString(fmt.Sprintf("%v -> %v\n\n", item.Path, item.Type))

		if len(item.Params) > 0 {
			sb.WriteString(item.Params.String())
			sb.WriteRune('\n')
		}

		sb.WriteString(fmt.Sprintf("  %v\n\n\n", item.Description))
	}

	return sb.String()
}

var predefinedHelpItems = helpItems{
	{
		Path:        "/int32s",
		Type:        "text/plain",
		Params:      nil,
		Description: "Signed 32-bit integer in decimal form",
	},
	{
		Path:        "/int32u",
		Type:        "text/plain",
		Params:      nil,
		Description: "Unsigned 32-bit integer in decimal form",
	},
	{
		Path:        "/int64s",
		Type:        "text/plain",
		Params:      nil,
		Description: "Signed 64-bit integer in decimal form",
	},
	{
		Path:        "/int64u",
		Type:        "text/plain",
		Params:      nil,
		Description: "Unsigned 64-bit integer in decimal form",
	},
	{
		Path: "/hex",
		Type: "text/plain",
		Params: paramsHelp{
			{Name: "bytes", Type: "int", Default: "16", Description: "Count of bytes desired"},
			{Name: "upper", Type: "bool", Default: "false", Description: "Use A-F instead of a-f"},
		},
		Description: "Bytes (no more than 1048576) in hexadecimal form",
	},
	{
		Path: "/blob",
		Type: "application/octet-stream",
		Params: paramsHelp{
			{Name: "bytes", Type: "int", Default: "16", Description: "Count of bytes desired"},
		},
		Description: "Blob (no larger than 1048576 bytes)",
	},
	{
		Path: "/base64",
		Type: "text/plain",
		Params: paramsHelp{
			{Name: "bytes", Type: "int", Default: "16", Description: "Count of bytes desired"},
		},
		Description: "Bytes (no more than 1048576) in base64 form",
	},
	{
		Path: "/s/uuid",
		Type: "text/plain",
		Params: nil,
		Description: "A uuid (Version 4, Variant 2) in format xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
	},
}

func HelpAll(ctx *gin.Context) {
	ctx.String(200, "%v", predefinedHelpItems)
}
