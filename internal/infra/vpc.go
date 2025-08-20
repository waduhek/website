package infra

import (
	"net"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateVPCComponents sets up a VPC with a single public subnet.
func CreateVPCComponents(ctx *pulumi.Context) (*ec2.Vpc, *ec2.Subnet, error) {
	websiteVPC, err := ec2.NewVpc(
		ctx,
		"website-vpc",
		&ec2.VpcArgs{
			CidrBlock:                    pulumi.String("10.0.0.0/16"),
			AssignGeneratedIpv6CidrBlock: pulumi.Bool(true),
		},
	)
	if err != nil {
		return nil, nil, err
	}

	websiteVPCPublicSubnet, err := createPublicSubnet(ctx, websiteVPC)
	if err != nil {
		return nil, nil, err
	}

	websiteIGW, err := createInternetGateway(ctx, websiteVPC)
	if err != nil {
		return nil, nil, err
	}

	websiteRT, err := createRouteTable(ctx, websiteVPC, websiteIGW)
	if err != nil {
		return nil, nil, err
	}

	err = associateRouteTableToSubnet(ctx, websiteRT, websiteVPCPublicSubnet)
	if err != nil {
		return nil, nil, err
	}

	return websiteVPC, websiteVPCPublicSubnet, nil
}

func createPublicSubnet(
	ctx *pulumi.Context,
	vpc *ec2.Vpc,
) (*ec2.Subnet, error) {
	return ec2.NewSubnet(
		ctx,
		"website-vpc-public-1",
		&ec2.SubnetArgs{
			VpcId:                       vpc.ID(),
			CidrBlock:                   pulumi.String("10.0.0.0/20"),
			Ipv6CidrBlock:               getPublicSubnetIPv6(vpc.Ipv6CidrBlock),
			AssignIpv6AddressOnCreation: pulumi.Bool(true),
			AvailabilityZone:            pulumi.String("ap-south-1a"),
		},
	)
}

func createInternetGateway(
	ctx *pulumi.Context,
	vpc *ec2.Vpc,
) (*ec2.InternetGateway, error) {
	return ec2.NewInternetGateway(
		ctx,
		"website-igw",
		&ec2.InternetGatewayArgs{
			VpcId: vpc.ID(),
		},
	)
}

func createRouteTable(
	ctx *pulumi.Context,
	vpc *ec2.Vpc,
	igw *ec2.InternetGateway,
) (*ec2.RouteTable, error) {
	return ec2.NewRouteTable(
		ctx,
		"website-rt",
		&ec2.RouteTableArgs{
			VpcId: vpc.ID(),
			Routes: ec2.RouteTableRouteArray{
				&ec2.RouteTableRouteArgs{
					CidrBlock: pulumi.String("0.0.0.0/0"),
					GatewayId: igw.ID(),
				},
				&ec2.RouteTableRouteArgs{
					Ipv6CidrBlock: pulumi.String("::0/0"),
					GatewayId:     igw.ID(),
				},
			},
		},
	)
}

func associateRouteTableToSubnet(
	ctx *pulumi.Context,
	rt *ec2.RouteTable,
	subnet *ec2.Subnet,
) error {
	_, err := ec2.NewRouteTableAssociation(
		ctx,
		"website-rt-association",
		&ec2.RouteTableAssociationArgs{
			RouteTableId: rt.ID(),
			SubnetId:     subnet.ID(),
		},
	)

	return err
}

func getPublicSubnetIPv6(vpcIPv6Block pulumi.StringOutput) pulumi.StringOutput {
	return vpcIPv6Block.ApplyT(func(block string) string {
		_, ipNet, _ := net.ParseCIDR(block)
		ipNet.Mask = net.CIDRMask(64, 128)

		return ipNet.String()
	}).(pulumi.StringOutput)
}
