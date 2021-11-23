// Copyright 2017 Manu Martinez-Almeida.  All rights reserved.
// Use of this source Code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
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

func validTimestamp(v validator.FieldLevel) bool {
	tims := v.Field().Int()
	nows := time.Now().Unix()
	if nows-tims >= 24*3600 || tims-nows >= 24*3600 {
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
	v.RegisterValidation("timestamp", validTimestamp)
	v.RegisterValidation("alphanull", validAlphaNull)
}
