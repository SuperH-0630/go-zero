package builderx

import (
	"github.com/wuntsong-org/go-zero-plus/core/stores/builder"
)

// Deprecated: Use github.com/wuntsong-org/go-zero-plus/core/stores/builder.RawFieldNames instead.
func FieldNames(in any) []string {
	return builder.RawFieldNames(in)
}

// Deprecated: Use github.com/wuntsong-org/go-zero-plus/core/stores/builder.RawFieldNames instead.
func RawFieldNames(in any, postgresSql ...bool) []string {
	return builder.RawFieldNames(in, postgresSql...)
}

// Deprecated: Use github.com/wuntsong-org/go-zero-plus/core/stores/builderx.PostgreSqlJoin instead.
func PostgreSqlJoin(elems []string) string {
	return builder.PostgreSqlJoin(elems)
}
