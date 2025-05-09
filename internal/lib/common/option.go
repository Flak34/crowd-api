package common

import "github.com/huandu/go-sqlbuilder"

type SelectOption func(sb *sqlbuilder.SelectBuilder) *sqlbuilder.SelectBuilder
