package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		websiteVPC, err := ec2.NewVpc(
			ctx,
			"website-vpc",
			&ec2.VpcArgs{
				CidrBlock:                    pulumi.String("10.0.0.0/16"),
				AssignGeneratedIpv6CidrBlock: pulumi.Bool(true),
			},
		)
		if err != nil {
			return err
		}

		websiteVPCPublicSubnet, err := ec2.NewSubnet(
			ctx,
			"website-vpc-public-1",
			&ec2.SubnetArgs{
				VpcId:     websiteVPC.ID(),
				CidrBlock: pulumi.String("10.0.0.0/20"),
			},
		)
		if err != nil {
			return err
		}

		websiteIGW, err := ec2.NewInternetGateway(
			ctx,
			"website-igw",
			&ec2.InternetGatewayArgs{
				VpcId: websiteVPC.ID(),
			},
		)
		if err != nil {
			return err
		}

		_, err = ec2.NewInternetGatewayAttachment(
			ctx,
			"website-igw-attachment",
			&ec2.InternetGatewayAttachmentArgs{
				VpcId:             websiteVPC.ID(),
				InternetGatewayId: websiteIGW.ID(),
			},
		)
		if err != nil {
			return err
		}

		websiteRT, err := ec2.NewRouteTable(
			ctx,
			"website-rt",
			&ec2.RouteTableArgs{
				VpcId: websiteVPC.ID(),
				Routes: ec2.RouteTableRouteArray{
					&ec2.RouteTableRouteArgs{
						CidrBlock: websiteVPCPublicSubnet.CidrBlock,
						GatewayId: websiteIGW.ID(),
					},
				},
			},
		)
		if err != nil {
			return err
		}

		_, err = ec2.NewRouteTableAssociation(
			ctx,
			"website-rt-association",
			&ec2.RouteTableAssociationArgs{
				RouteTableId: websiteRT.ID(),
				SubnetId:     websiteVPCPublicSubnet.ID(),
			},
		)

		return nil
	})
}
