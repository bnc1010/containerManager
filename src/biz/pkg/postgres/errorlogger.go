package postgres

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

var ctx = context.Background()

func imageErrorLoger(err error) {
	hlog.CtxFatalf(ctx, "[Postgres - Image] ", err)
}

func projectErrorLoger(err error) {
	hlog.CtxFatalf(ctx, "[Postgres - Project] ", err)
}

func filesErrorLoger(err error) {
	hlog.CtxFatalf(ctx, "[Postgres - Files] ", err)
}

func datasetErrorLoger(err error) {
	hlog.CtxFatalf(ctx, "[Postgres - Dataset] ", err)
}