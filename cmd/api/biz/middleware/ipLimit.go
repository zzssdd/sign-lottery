package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"sign-lottery/dao/cache"
	"sign-lottery/pkg/constants"
)

func IpLimitMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		clientIp := c.ClientIP()
		Cache := cache.NewCache().Ip
		err := Cache.IncrIpCount(ctx, clientIp)
		if err != nil {
			c.Abort()
			return
		}
		count, err := Cache.GetIpCount(ctx, clientIp)
		if err != nil {
			c.Abort()
			return
		}
		if count > constants.IpLimitCount {
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}
