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

		logGroup, err := infra.CreateLogGroup(ctx)
		if err != nil {
			return err
		}

		ec2InstProfile, err := infra.CreateWebsiteEC2InstanceProfile(
			ctx, logGroup,
		)
		if err != nil {
			return err
		}

		websiteInstance, err := infra.CreateEC2Instance(
			ctx,
			websitePublicSubnet,
			websiteSG,
			ec2InstProfile,
		)
		if err != nil {
			return err
		}

		ctx.Export("websiteVPC", websiteVPC.ID())
		ctx.Export("websitePublicIP", websiteInstance.PublicIp)
		ctx.Export("websiteIPv6", websiteInstance.Ipv6Addresses)

		return nil
	})
}
