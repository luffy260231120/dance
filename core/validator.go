// Copyright 2017 Manu Martinez-Almeida.  All rights reserved.
// Use of this source Code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"dance/cons"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func validEnum(v validator.FieldLevel) bool {
	value := ""
	enums := strings.Split(v.Param(), "-")
	fvals, fkind, _ := v.ExtractType(v.Field())
	if fkind == reflect.String {
		value = fvals.String()
	} else {
		value = strconv.FormatInt(fvals.Int(), 10)
	}
	for _, enum := range enums {
		if enum == value {
			return true
		}
	}
	return false
}

func validTime(v validator.FieldLevel) bool {
	timeStr := v.Field().String()
	_, err := time.ParseInLocation(cons.FORMAT_TIME, timeStr, time.Local)
	if err != nil {
		return false
	}
	return true
}

const alphaRegexString = "^[a-zA-Z#]{0,1}$"

var alphaRegex = regexp.MustCompile(alphaRegexString)

func validAlphaNull(v validator.FieldLevel) bool {
	return alphaRegex.MatchString(v.Field().String())
}

func init() {
	v := binding.Validator.Engine().(*validator.Validate)
	v.RegisterValidation("enum", validEnum)
	v.RegisterValidation("time", validTime)
	v.RegisterValidation("alphanull", validAlphaNull)
}
