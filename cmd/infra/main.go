package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/waduhek/website/internal/infra"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		websiteVPC, websitePublicSubnet, err := infra.CreateVPCComponents(ctx)
		if err != nil {
			return err
		}

		websiteSG, err := infra.CreateSecurityGroup(ctx, websiteVPC)
		if err != nil {
			return err
		}

		websiteInstance, err := infra.CreateEC2Instance(
			ctx,
			websitePublicSubnet,
			websiteSG,
		)
		if err != nil {
			return err
		}

		ctx.Export("websiteVPC", websiteVPC.ID())
		ctx.Export("websitePublicIP", websiteInstance.PublicIp)

		return nil
	})
}
