// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: syncql

package syncql

import (
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/verror"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Error definitions

var (
	ErrBadFieldInWhere                 = verror.Register("v.io/v23/query/syncql.BadFieldInWhere", verror.NoRetry, "{1:}{2:} [{3}]Where field must be 'k' or 'v[{.<ident>}...]'.")
	ErrBoolInvalidExpression           = verror.Register("v.io/v23/query/syncql.BoolInvalidExpression", verror.NoRetry, "{1:}{2:} [{3}]Boolean operands may only be used in equals and not equals expressions.")
	ErrCheckOfUnknownStatementType     = verror.Register("v.io/v23/query/syncql.CheckOfUnknownStatementType", verror.NoRetry, "{1:}{2:} [{3}]Cannot semantically check unknown statement type.")
	ErrCouldNotConvert                 = verror.Register("v.io/v23/query/syncql.CouldNotConvert", verror.NoRetry, "{1:}{2:} [{3}]Could not convert {4} to {5}.")
	ErrDotNotationDisallowedForKey     = verror.Register("v.io/v23/query/syncql.DotNotationDisallowedForKey", verror.NoRetry, "{1:}{2:} [{3}]Dot notation may not be used on a key field.")
	ErrExecOfUnknownStatementType      = verror.Register("v.io/v23/query/syncql.ExecOfUnknownStatementType", verror.NoRetry, "{1:}{2:} [{3}]Cannot execute unknown statement type: {4}.")
	ErrExpected                        = verror.Register("v.io/v23/query/syncql.Expected", verror.NoRetry, "{1:}{2:} [{3}]Expected '{4}'.")
	ErrExpectedFrom                    = verror.Register("v.io/v23/query/syncql.ExpectedFrom", verror.NoRetry, "{1:}{2:} [{3}]Expected 'from', found {4}.")
	ErrExpectedIdentifier              = verror.Register("v.io/v23/query/syncql.ExpectedIdentifier", verror.NoRetry, "{1:}{2:} [{3}]Expected identifier, found {4}.")
	ErrExpectedOperand                 = verror.Register("v.io/v23/query/syncql.ExpectedOperand", verror.NoRetry, "{1:}{2:} [{3}]Expected operand, found {4}.")
	ErrExpectedOperator                = verror.Register("v.io/v23/query/syncql.ExpectedOperator", verror.NoRetry, "{1:}{2:} [{3}]Expected operator, found {4}.")
	ErrFunctionArgCount                = verror.Register("v.io/v23/query/syncql.FunctionArgCount", verror.NoRetry, "{1:}{2:} [{3}]Function '{4}' expects {5} args, found: {6}.")
	ErrFunctionAtLeastArgCount         = verror.Register("v.io/v23/query/syncql.FunctionAtLeastArgCount", verror.NoRetry, "{1:}{2:} [{3}]Function '{4}' expects at least {5} args, found: {6}.")
	ErrFunctionTypeInvalidArg          = verror.Register("v.io/v23/query/syncql.FunctionTypeInvalidArg", verror.NoRetry, "{1:}{2:} [{3}]Function 'Type()' cannot get type of argument -- expecting object.")
	ErrFunctionLenInvalidArg           = verror.Register("v.io/v23/query/syncql.FunctionLenInvalidArg", verror.NoRetry, "{1:}{2:} [{3}]Function 'Len()' expects array, list, set, map, string or nil.")
	ErrFunctionArgBad                  = verror.Register("v.io/v23/query/syncql.FunctionArgBad", verror.NoRetry, "{1:}{2:} [{3}]Function '{4}' arg '{5}' could not be resolved.")
	ErrFunctionNotFound                = verror.Register("v.io/v23/query/syncql.FunctionNotFound", verror.NoRetry, "{1:}{2:} [{3}]Function '{4}' not found.")
	ErrArgMustBeField                  = verror.Register("v.io/v23/query/syncql.ArgMustBeField", verror.NoRetry, "{1:}{2:} [{3}]Argument must be a value field (i.e., must begin with 'v').")
	ErrBigIntConversionError           = verror.Register("v.io/v23/query/syncql.BigIntConversionError", verror.NoRetry, "{1:}{2:} [{3}]Can't convert to BigInt: {4}.")
	ErrBigRatConversionError           = verror.Register("v.io/v23/query/syncql.BigRatConversionError", verror.NoRetry, "{1:}{2:} [{3}]Can't convert to BigRat: {4}.")
	ErrBoolConversionError             = verror.Register("v.io/v23/query/syncql.BoolConversionError", verror.NoRetry, "{1:}{2:} [{3}]Can't convert to Bool: {4}.")
	ErrComplexConversionError          = verror.Register("v.io/v23/query/syncql.ComplexConversionError", verror.NoRetry, "{1:}{2:} [{3}]Can't convert to Complex: {4}.")
	ErrUintConversionError             = verror.Register("v.io/v23/query/syncql.UintConversionError", verror.NoRetry, "{1:}{2:} [{3}]Can't convert to Uint: {4}.")
	ErrTimeConversionError             = verror.Register("v.io/v23/query/syncql.TimeConversionError", verror.NoRetry, "{1:}{2:} [{3}]Can't convert to time: {4}.")
	ErrLocationConversionError         = verror.Register("v.io/v23/query/syncql.LocationConversionError", verror.NoRetry, "{1:}{2:} [{3}]Can't convert to location: {4}.")
	ErrStringConversionError           = verror.Register("v.io/v23/query/syncql.StringConversionError", verror.NoRetry, "{1:}{2:} [{3}]Can't convert to string: {4}.")
	ErrFloatConversionError            = verror.Register("v.io/v23/query/syncql.FloatConversionError", verror.NoRetry, "{1:}{2:} [{3}]Can't convert to float: {4}.")
	ErrIntConversionError              = verror.Register("v.io/v23/query/syncql.IntConversionError", verror.NoRetry, "{1:}{2:} [{3}]Can't convert to int: {4}.")
	ErrIsIsNotRequireLhsValue          = verror.Register("v.io/v23/query/syncql.IsIsNotRequireLhsValue", verror.NoRetry, "{1:}{2:} [{3}]'Is/is not' expressions require left operand to be a value operand.")
	ErrIsIsNotRequireRhsNil            = verror.Register("v.io/v23/query/syncql.IsIsNotRequireRhsNil", verror.NoRetry, "{1:}{2:} [{3}]'Is/is not' expressions require right operand to be nil.")
	ErrInvalidLikePattern              = verror.Register("v.io/v23/query/syncql.InvalidLikePattern", verror.NoRetry, "{1:}{2:} [{3}]Invalid like pattern: {4}.")
	ErrInvalidSelectField              = verror.Register("v.io/v23/query/syncql.InvalidSelectField", verror.NoRetry, "{1:}{2:} [{3}]Select field must be 'k' or 'v[{.<ident>}...]'.")
	ErrKeyExpressionLiteral            = verror.Register("v.io/v23/query/syncql.KeyExpressionLiteral", verror.NoRetry, "{1:}{2:} [{3}]Key (i.e., 'k') compares against literals must be string literal.")
	ErrKeyValueStreamError             = verror.Register("v.io/v23/query/syncql.KeyValueStreamError", verror.NoRetry, "{1:}{2:} [{3}]KeyValueStream error: {4}.")
	ErrLikeExpressionsRequireRhsString = verror.Register("v.io/v23/query/syncql.LikeExpressionsRequireRhsString", verror.NoRetry, "{1:}{2:} [{3}]Like expressions require right operand of type <string-literal>.")
	ErrLimitMustBeGt0                  = verror.Register("v.io/v23/query/syncql.LimitMustBeGt0", verror.NoRetry, "{1:}{2:} [{3}]Limit must be > 0.")
	ErrMaxStatementLenExceeded         = verror.Register("v.io/v23/query/syncql.MaxStatementLenExceeded", verror.NoRetry, "{1:}{2:} [{3}]Maximum length of statements is {4}; found {5}.")
	ErrNoStatementFound                = verror.Register("v.io/v23/query/syncql.NoStatementFound", verror.NoRetry, "{1:}{2:} [{3}]No statement found.")
	ErrOffsetMustBeGe0                 = verror.Register("v.io/v23/query/syncql.OffsetMustBeGe0", verror.NoRetry, "{1:}{2:} [{3}]Offset must be > 0.")
	ErrScanError                       = verror.Register("v.io/v23/query/syncql.ScanError", verror.NoRetry, "{1:}{2:} [{3}]Scan error: {4}.")
	ErrTableCantAccess                 = verror.Register("v.io/v23/query/syncql.TableCantAccess", verror.NoRetry, "{1:}{2:} [{3}]Table {4} does not exist (or cannot be accessed): {5}.")
	ErrUnexpected                      = verror.Register("v.io/v23/query/syncql.Unexpected", verror.NoRetry, "{1:}{2:} [{3}]Unexpected: {4}.")
	ErrUnexpectedEndOfStatement        = verror.Register("v.io/v23/query/syncql.UnexpectedEndOfStatement", verror.NoRetry, "{1:}{2:} [{3}]Unexpected end of statement.")
	ErrUnknownIdentifier               = verror.Register("v.io/v23/query/syncql.UnknownIdentifier", verror.NoRetry, "{1:}{2:} [{3}]Unknown identifier: {4}.")
	ErrInvalidEscapeChar               = verror.Register("v.io/v23/query/syncql.InvalidEscapeChar", verror.NoRetry, "{1:}{2:} [{3}]'{4}' is not a valid escape character.")
	ErrDidYouMeanLowercaseK            = verror.Register("v.io/v23/query/syncql.DidYouMeanLowercaseK", verror.NoRetry, "{1:}{2:} [{3}]Did you mean: 'k'?")
	ErrDidYouMeanLowercaseV            = verror.Register("v.io/v23/query/syncql.DidYouMeanLowercaseV", verror.NoRetry, "{1:}{2:} [{3}]Did you mean: 'v'?")
	ErrDidYouMeanFunction              = verror.Register("v.io/v23/query/syncql.DidYouMeanFunction", verror.NoRetry, "{1:}{2:} [{3}]Did you mean: '{4}'?")
	ErrNotEnoughParamValuesSpecified   = verror.Register("v.io/v23/query/syncql.NotEnoughParamValuesSpecified", verror.NoRetry, "{1:}{2:} [{3}]Not enough parameter values specified.")
	ErrTooManyParamValuesSpecified     = verror.Register("v.io/v23/query/syncql.TooManyParamValuesSpecified", verror.NoRetry, "{1:}{2:} [{3}]Too many parameter values specified.")
	ErrPreparedStatementNotFound       = verror.Register("v.io/v23/query/syncql.PreparedStatementNotFound", verror.NoRetry, "{1:}{2:} [0]Prepared statement not found.")
	ErrIndexKindNotSupported           = verror.Register("v.io/v23/query/syncql.IndexKindNotSupported", verror.NoRetry, "{1:}{2:} [{3}]Index kind {4} of field {5} on table {6} not supported.")
	ErrInvalidIndexField               = verror.Register("v.io/v23/query/syncql.InvalidIndexField", verror.NoRetry, "{1:}{2:} [{3}]Invalid index field {4} returned by table {5}.")
	ErrNotWritable                     = verror.Register("v.io/v23/query/syncql.NotWritable", verror.NoRetry, "{1:}{2:} [0]Can't write to table {3} (not supported on batch/connection).")
	ErrOperationNotSupported           = verror.Register("v.io/v23/query/syncql.OperationNotSupported", verror.NoRetry, "{1:}{2:} [0]{3} not supported.")
)

// NewErrBadFieldInWhere returns an error with the ErrBadFieldInWhere ID.
func NewErrBadFieldInWhere(ctx *context.T, off int64) error {
	return verror.New(ErrBadFieldInWhere, ctx, off)
}

// NewErrBoolInvalidExpression returns an error with the ErrBoolInvalidExpression ID.
func NewErrBoolInvalidExpression(ctx *context.T, off int64) error {
	return verror.New(ErrBoolInvalidExpression, ctx, off)
}

// NewErrCheckOfUnknownStatementType returns an error with the ErrCheckOfUnknownStatementType ID.
func NewErrCheckOfUnknownStatementType(ctx *context.T, off int64) error {
	return verror.New(ErrCheckOfUnknownStatementType, ctx, off)
}

// NewErrCouldNotConvert returns an error with the ErrCouldNotConvert ID.
func NewErrCouldNotConvert(ctx *context.T, off int64, from string, to string) error {
	return verror.New(ErrCouldNotConvert, ctx, off, from, to)
}

// NewErrDotNotationDisallowedForKey returns an error with the ErrDotNotationDisallowedForKey ID.
func NewErrDotNotationDisallowedForKey(ctx *context.T, off int64) error {
	return verror.New(ErrDotNotationDisallowedForKey, ctx, off)
}

// NewErrExecOfUnknownStatementType returns an error with the ErrExecOfUnknownStatementType ID.
func NewErrExecOfUnknownStatementType(ctx *context.T, off int64, statementType string) error {
	return verror.New(ErrExecOfUnknownStatementType, ctx, off, statementType)
}

// NewErrExpected returns an error with the ErrExpected ID.
func NewErrExpected(ctx *context.T, off int64, expected string) error {
	return verror.New(ErrExpected, ctx, off, expected)
}

// NewErrExpectedFrom returns an error with the ErrExpectedFrom ID.
func NewErrExpectedFrom(ctx *context.T, off int64, found string) error {
	return verror.New(ErrExpectedFrom, ctx, off, found)
}

// NewErrExpectedIdentifier returns an error with the ErrExpectedIdentifier ID.
func NewErrExpectedIdentifier(ctx *context.T, off int64, found string) error {
	return verror.New(ErrExpectedIdentifier, ctx, off, found)
}

// NewErrExpectedOperand returns an error with the ErrExpectedOperand ID.
func NewErrExpectedOperand(ctx *context.T, off int64, found string) error {
	return verror.New(ErrExpectedOperand, ctx, off, found)
}

// NewErrExpectedOperator returns an error with the ErrExpectedOperator ID.
func NewErrExpectedOperator(ctx *context.T, off int64, found string) error {
	return verror.New(ErrExpectedOperator, ctx, off, found)
}

// NewErrFunctionArgCount returns an error with the ErrFunctionArgCount ID.
func NewErrFunctionArgCount(ctx *context.T, off int64, name string, expected int64, found int64) error {
	return verror.New(ErrFunctionArgCount, ctx, off, name, expected, found)
}

// NewErrFunctionAtLeastArgCount returns an error with the ErrFunctionAtLeastArgCount ID.
func NewErrFunctionAtLeastArgCount(ctx *context.T, off int64, name string, expected int64, found int64) error {
	return verror.New(ErrFunctionAtLeastArgCount, ctx, off, name, expected, found)
}

// NewErrFunctionTypeInvalidArg returns an error with the ErrFunctionTypeInvalidArg ID.
func NewErrFunctionTypeInvalidArg(ctx *context.T, off int64) error {
	return verror.New(ErrFunctionTypeInvalidArg, ctx, off)
}

// NewErrFunctionLenInvalidArg returns an error with the ErrFunctionLenInvalidArg ID.
func NewErrFunctionLenInvalidArg(ctx *context.T, off int64) error {
	return verror.New(ErrFunctionLenInvalidArg, ctx, off)
}

// NewErrFunctionArgBad returns an error with the ErrFunctionArgBad ID.
func NewErrFunctionArgBad(ctx *context.T, off int64, funcName string, argName string) error {
	return verror.New(ErrFunctionArgBad, ctx, off, funcName, argName)
}

// NewErrFunctionNotFound returns an error with the ErrFunctionNotFound ID.
func NewErrFunctionNotFound(ctx *context.T, off int64, name string) error {
	return verror.New(ErrFunctionNotFound, ctx, off, name)
}

// NewErrArgMustBeField returns an error with the ErrArgMustBeField ID.
func NewErrArgMustBeField(ctx *context.T, off int64) error {
	return verror.New(ErrArgMustBeField, ctx, off)
}

// NewErrBigIntConversionError returns an error with the ErrBigIntConversionError ID.
func NewErrBigIntConversionError(ctx *context.T, off int64, err error) error {
	return verror.New(ErrBigIntConversionError, ctx, off, err)
}

// NewErrBigRatConversionError returns an error with the ErrBigRatConversionError ID.
func NewErrBigRatConversionError(ctx *context.T, off int64, err error) error {
	return verror.New(ErrBigRatConversionError, ctx, off, err)
}

// NewErrBoolConversionError returns an error with the ErrBoolConversionError ID.
func NewErrBoolConversionError(ctx *context.T, off int64, err error) error {
	return verror.New(ErrBoolConversionError, ctx, off, err)
}

// NewErrComplexConversionError returns an error with the ErrComplexConversionError ID.
func NewErrComplexConversionError(ctx *context.T, off int64, err error) error {
	return verror.New(ErrComplexConversionError, ctx, off, err)
}

// NewErrUintConversionError returns an error with the ErrUintConversionError ID.
func NewErrUintConversionError(ctx *context.T, off int64, err error) error {
	return verror.New(ErrUintConversionError, ctx, off, err)
}

// NewErrTimeConversionError returns an error with the ErrTimeConversionError ID.
func NewErrTimeConversionError(ctx *context.T, off int64, err error) error {
	return verror.New(ErrTimeConversionError, ctx, off, err)
}

// NewErrLocationConversionError returns an error with the ErrLocationConversionError ID.
func NewErrLocationConversionError(ctx *context.T, off int64, err error) error {
	return verror.New(ErrLocationConversionError, ctx, off, err)
}

// NewErrStringConversionError returns an error with the ErrStringConversionError ID.
func NewErrStringConversionError(ctx *context.T, off int64, err error) error {
	return verror.New(ErrStringConversionError, ctx, off, err)
}

// NewErrFloatConversionError returns an error with the ErrFloatConversionError ID.
func NewErrFloatConversionError(ctx *context.T, off int64, err error) error {
	return verror.New(ErrFloatConversionError, ctx, off, err)
}

// NewErrIntConversionError returns an error with the ErrIntConversionError ID.
func NewErrIntConversionError(ctx *context.T, off int64, err error) error {
	return verror.New(ErrIntConversionError, ctx, off, err)
}

// NewErrIsIsNotRequireLhsValue returns an error with the ErrIsIsNotRequireLhsValue ID.
func NewErrIsIsNotRequireLhsValue(ctx *context.T, off int64) error {
	return verror.New(ErrIsIsNotRequireLhsValue, ctx, off)
}

// NewErrIsIsNotRequireRhsNil returns an error with the ErrIsIsNotRequireRhsNil ID.
func NewErrIsIsNotRequireRhsNil(ctx *context.T, off int64) error {
	return verror.New(ErrIsIsNotRequireRhsNil, ctx, off)
}

// NewErrInvalidLikePattern returns an error with the ErrInvalidLikePattern ID.
func NewErrInvalidLikePattern(ctx *context.T, off int64, err error) error {
	return verror.New(ErrInvalidLikePattern, ctx, off, err)
}

// NewErrInvalidSelectField returns an error with the ErrInvalidSelectField ID.
func NewErrInvalidSelectField(ctx *context.T, off int64) error {
	return verror.New(ErrInvalidSelectField, ctx, off)
}

// NewErrKeyExpressionLiteral returns an error with the ErrKeyExpressionLiteral ID.
func NewErrKeyExpressionLiteral(ctx *context.T, off int64) error {
	return verror.New(ErrKeyExpressionLiteral, ctx, off)
}

// NewErrKeyValueStreamError returns an error with the ErrKeyValueStreamError ID.
func NewErrKeyValueStreamError(ctx *context.T, off int64, err error) error {
	return verror.New(ErrKeyValueStreamError, ctx, off, err)
}

// NewErrLikeExpressionsRequireRhsString returns an error with the ErrLikeExpressionsRequireRhsString ID.
func NewErrLikeExpressionsRequireRhsString(ctx *context.T, off int64) error {
	return verror.New(ErrLikeExpressionsRequireRhsString, ctx, off)
}

// NewErrLimitMustBeGt0 returns an error with the ErrLimitMustBeGt0 ID.
func NewErrLimitMustBeGt0(ctx *context.T, off int64) error {
	return verror.New(ErrLimitMustBeGt0, ctx, off)
}

// NewErrMaxStatementLenExceeded returns an error with the ErrMaxStatementLenExceeded ID.
func NewErrMaxStatementLenExceeded(ctx *context.T, off int64, max int64, found int64) error {
	return verror.New(ErrMaxStatementLenExceeded, ctx, off, max, found)
}

// NewErrNoStatementFound returns an error with the ErrNoStatementFound ID.
func NewErrNoStatementFound(ctx *context.T, off int64) error {
	return verror.New(ErrNoStatementFound, ctx, off)
}

// NewErrOffsetMustBeGe0 returns an error with the ErrOffsetMustBeGe0 ID.
func NewErrOffsetMustBeGe0(ctx *context.T, off int64) error {
	return verror.New(ErrOffsetMustBeGe0, ctx, off)
}

// NewErrScanError returns an error with the ErrScanError ID.
func NewErrScanError(ctx *context.T, off int64, err error) error {
	return verror.New(ErrScanError, ctx, off, err)
}

// NewErrTableCantAccess returns an error with the ErrTableCantAccess ID.
func NewErrTableCantAccess(ctx *context.T, off int64, table string, err error) error {
	return verror.New(ErrTableCantAccess, ctx, off, table, err)
}

// NewErrUnexpected returns an error with the ErrUnexpected ID.
func NewErrUnexpected(ctx *context.T, off int64, found string) error {
	return verror.New(ErrUnexpected, ctx, off, found)
}

// NewErrUnexpectedEndOfStatement returns an error with the ErrUnexpectedEndOfStatement ID.
func NewErrUnexpectedEndOfStatement(ctx *context.T, off int64) error {
	return verror.New(ErrUnexpectedEndOfStatement, ctx, off)
}

// NewErrUnknownIdentifier returns an error with the ErrUnknownIdentifier ID.
func NewErrUnknownIdentifier(ctx *context.T, off int64, found string) error {
	return verror.New(ErrUnknownIdentifier, ctx, off, found)
}

// NewErrInvalidEscapeChar returns an error with the ErrInvalidEscapeChar ID.
func NewErrInvalidEscapeChar(ctx *context.T, off int64, escChar string) error {
	return verror.New(ErrInvalidEscapeChar, ctx, off, escChar)
}

// NewErrDidYouMeanLowercaseK returns an error with the ErrDidYouMeanLowercaseK ID.
func NewErrDidYouMeanLowercaseK(ctx *context.T, off int64) error {
	return verror.New(ErrDidYouMeanLowercaseK, ctx, off)
}

// NewErrDidYouMeanLowercaseV returns an error with the ErrDidYouMeanLowercaseV ID.
func NewErrDidYouMeanLowercaseV(ctx *context.T, off int64) error {
	return verror.New(ErrDidYouMeanLowercaseV, ctx, off)
}

// NewErrDidYouMeanFunction returns an error with the ErrDidYouMeanFunction ID.
func NewErrDidYouMeanFunction(ctx *context.T, off int64, correctName string) error {
	return verror.New(ErrDidYouMeanFunction, ctx, off, correctName)
}

// NewErrNotEnoughParamValuesSpecified returns an error with the ErrNotEnoughParamValuesSpecified ID.
func NewErrNotEnoughParamValuesSpecified(ctx *context.T, off int64) error {
	return verror.New(ErrNotEnoughParamValuesSpecified, ctx, off)
}

// NewErrTooManyParamValuesSpecified returns an error with the ErrTooManyParamValuesSpecified ID.
func NewErrTooManyParamValuesSpecified(ctx *context.T, off int64) error {
	return verror.New(ErrTooManyParamValuesSpecified, ctx, off)
}

// NewErrPreparedStatementNotFound returns an error with the ErrPreparedStatementNotFound ID.
func NewErrPreparedStatementNotFound(ctx *context.T) error {
	return verror.New(ErrPreparedStatementNotFound, ctx)
}

// NewErrIndexKindNotSupported returns an error with the ErrIndexKindNotSupported ID.
func NewErrIndexKindNotSupported(ctx *context.T, off int64, kind string, fieldName string, table string) error {
	return verror.New(ErrIndexKindNotSupported, ctx, off, kind, fieldName, table)
}

// NewErrInvalidIndexField returns an error with the ErrInvalidIndexField ID.
func NewErrInvalidIndexField(ctx *context.T, off int64, fieldName string, table string) error {
	return verror.New(ErrInvalidIndexField, ctx, off, fieldName, table)
}

// NewErrNotWritable returns an error with the ErrNotWritable ID.
func NewErrNotWritable(ctx *context.T, table string) error {
	return verror.New(ErrNotWritable, ctx, table)
}

// NewErrOperationNotSupported returns an error with the ErrOperationNotSupported ID.
func NewErrOperationNotSupported(ctx *context.T, operation string) error {
	return verror.New(ErrOperationNotSupported, ctx, operation)
}

var __VDLInitCalled bool

// __VDLInit performs vdl initialization.  It is safe to call multiple times.
// If you have an init ordering issue, just insert the following line verbatim
// into your source files in this package, right after the "package foo" clause:
//
//    var _ = __VDLInit()
//
// The purpose of this function is to ensure that vdl initialization occurs in
// the right order, and very early in the init sequence.  In particular, vdl
// registration and package variable initialization needs to occur before
// functions like vdl.TypeOf will work properly.
//
// This function returns a dummy value, so that it can be used to initialize the
// first var in the file, to take advantage of Go's defined init order.
func __VDLInit() struct{} {
	if __VDLInitCalled {
		return struct{}{}
	}
	__VDLInitCalled = true

	// Set error format strings.
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrBadFieldInWhere.ID), "{1:}{2:} [{3}]Where field must be 'k' or 'v[{.<ident>}...]'.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrBoolInvalidExpression.ID), "{1:}{2:} [{3}]Boolean operands may only be used in equals and not equals expressions.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrCheckOfUnknownStatementType.ID), "{1:}{2:} [{3}]Cannot semantically check unknown statement type.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrCouldNotConvert.ID), "{1:}{2:} [{3}]Could not convert {4} to {5}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrDotNotationDisallowedForKey.ID), "{1:}{2:} [{3}]Dot notation may not be used on a key field.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrExecOfUnknownStatementType.ID), "{1:}{2:} [{3}]Cannot execute unknown statement type: {4}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrExpected.ID), "{1:}{2:} [{3}]Expected '{4}'.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrExpectedFrom.ID), "{1:}{2:} [{3}]Expected 'from', found {4}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrExpectedIdentifier.ID), "{1:}{2:} [{3}]Expected identifier, found {4}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrExpectedOperand.ID), "{1:}{2:} [{3}]Expected operand, found {4}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrExpectedOperator.ID), "{1:}{2:} [{3}]Expected operator, found {4}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrFunctionArgCount.ID), "{1:}{2:} [{3}]Function '{4}' expects {5} args, found: {6}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrFunctionAtLeastArgCount.ID), "{1:}{2:} [{3}]Function '{4}' expects at least {5} args, found: {6}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrFunctionTypeInvalidArg.ID), "{1:}{2:} [{3}]Function 'Type()' cannot get type of argument -- expecting object.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrFunctionLenInvalidArg.ID), "{1:}{2:} [{3}]Function 'Len()' expects array, list, set, map, string or nil.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrFunctionArgBad.ID), "{1:}{2:} [{3}]Function '{4}' arg '{5}' could not be resolved.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrFunctionNotFound.ID), "{1:}{2:} [{3}]Function '{4}' not found.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrArgMustBeField.ID), "{1:}{2:} [{3}]Argument must be a value field (i.e., must begin with 'v').")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrBigIntConversionError.ID), "{1:}{2:} [{3}]Can't convert to BigInt: {4}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrBigRatConversionError.ID), "{1:}{2:} [{3}]Can't convert to BigRat: {4}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrBoolConversionError.ID), "{1:}{2:} [{3}]Can't convert to Bool: {4}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrComplexConversionError.ID), "{1:}{2:} [{3}]Can't convert to Complex: {4}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrUintConversionError.ID), "{1:}{2:} [{3}]Can't convert to Uint: {4}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrTimeConversionError.ID), "{1:}{2:} [{3}]Can't convert to time: {4}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrLocationConversionError.ID), "{1:}{2:} [{3}]Can't convert to location: {4}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrStringConversionError.ID), "{1:}{2:} [{3}]Can't convert to string: {4}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrFloatConversionError.ID), "{1:}{2:} [{3}]Can't convert to float: {4}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrIntConversionError.ID), "{1:}{2:} [{3}]Can't convert to int: {4}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrIsIsNotRequireLhsValue.ID), "{1:}{2:} [{3}]'Is/is not' expressions require left operand to be a value operand.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrIsIsNotRequireRhsNil.ID), "{1:}{2:} [{3}]'Is/is not' expressions require right operand to be nil.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrInvalidLikePattern.ID), "{1:}{2:} [{3}]Invalid like pattern: {4}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrInvalidSelectField.ID), "{1:}{2:} [{3}]Select field must be 'k' or 'v[{.<ident>}...]'.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrKeyExpressionLiteral.ID), "{1:}{2:} [{3}]Key (i.e., 'k') compares against literals must be string literal.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrKeyValueStreamError.ID), "{1:}{2:} [{3}]KeyValueStream error: {4}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrLikeExpressionsRequireRhsString.ID), "{1:}{2:} [{3}]Like expressions require right operand of type <string-literal>.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrLimitMustBeGt0.ID), "{1:}{2:} [{3}]Limit must be > 0.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrMaxStatementLenExceeded.ID), "{1:}{2:} [{3}]Maximum length of statements is {4}; found {5}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNoStatementFound.ID), "{1:}{2:} [{3}]No statement found.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrOffsetMustBeGe0.ID), "{1:}{2:} [{3}]Offset must be > 0.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrScanError.ID), "{1:}{2:} [{3}]Scan error: {4}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrTableCantAccess.ID), "{1:}{2:} [{3}]Table {4} does not exist (or cannot be accessed): {5}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrUnexpected.ID), "{1:}{2:} [{3}]Unexpected: {4}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrUnexpectedEndOfStatement.ID), "{1:}{2:} [{3}]Unexpected end of statement.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrUnknownIdentifier.ID), "{1:}{2:} [{3}]Unknown identifier: {4}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrInvalidEscapeChar.ID), "{1:}{2:} [{3}]'{4}' is not a valid escape character.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrDidYouMeanLowercaseK.ID), "{1:}{2:} [{3}]Did you mean: 'k'?")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrDidYouMeanLowercaseV.ID), "{1:}{2:} [{3}]Did you mean: 'v'?")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrDidYouMeanFunction.ID), "{1:}{2:} [{3}]Did you mean: '{4}'?")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNotEnoughParamValuesSpecified.ID), "{1:}{2:} [{3}]Not enough parameter values specified.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrTooManyParamValuesSpecified.ID), "{1:}{2:} [{3}]Too many parameter values specified.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrPreparedStatementNotFound.ID), "{1:}{2:} [0]Prepared statement not found.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrIndexKindNotSupported.ID), "{1:}{2:} [{3}]Index kind {4} of field {5} on table {6} not supported.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrInvalidIndexField.ID), "{1:}{2:} [{3}]Invalid index field {4} returned by table {5}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNotWritable.ID), "{1:}{2:} [0]Can't write to table {3} (not supported on batch/connection).")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrOperationNotSupported.ID), "{1:}{2:} [0]{3} not supported.")

	return struct{}{}
}
