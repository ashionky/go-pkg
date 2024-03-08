package asynctask

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const contextTimeOut = 10

/**
 * 异步任务设计需求，通过创建新的context，将原来的context里面的追踪的值复制过来，然后设定过期时间，以及请求完成之后销毁
 */

func AddTrace(pctx context.Context, ctx context.Context) context.Context {
	md, _ := metadata.FromOutgoingContext(pctx)
	ctx = metadata.NewOutgoingContext(ctx, md)
	//添加上下文信息
	ctx = context.WithValue(ctx, "traceId", pctx.Value("traceId"))
	ctx = context.WithValue(ctx, "caller", pctx.Value("caller"))
	return ctx
}
