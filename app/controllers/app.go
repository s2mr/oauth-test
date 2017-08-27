package controllers

import (
	"github.com/revel/revel"
	//"golang.org/x/net/context"
	//"golang.org/x/oauth2/google"
	//"google.golang.org/api/oauth2/v1"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	// Use oauth2.NoContext if there isn't a good context to pass in.
	//ctx := context.Background()
	//
	//client, err := google.DefaultClient(ctx, compute.ComputeScope)
	//if err != nil {
	//	//...
	//}
	//computeService, err := compute.New(client)
	//if err != nil {
	//	//...
	//}
	//
	//ts, err := google.DefaultTokenSource(ctx, scope1, scope2, ...)
	//if err != nil {
	//	//...
	//}
	//client := oauth2.NewClient(ctx, ts)



	return c.Render()
}
