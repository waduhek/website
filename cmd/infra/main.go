package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/vpc"
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
				VpcId:            websiteVPC.ID(),
				CidrBlock:        pulumi.String("10.0.0.0/20"),
				AvailabilityZone: pulumi.String("ap-south-1a"),
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

		websiteRT, err := ec2.NewRouteTable(
			ctx,
			"website-rt",
			&ec2.RouteTableArgs{
				VpcId: websiteVPC.ID(),
				Routes: ec2.RouteTableRouteArray{
					&ec2.RouteTableRouteArgs{
						CidrBlock: pulumi.String("0.0.0.0/0"),
						GatewayId: websiteIGW.ID(),
					},
					&ec2.RouteTableRouteArgs{
						Ipv6CidrBlock: pulumi.String("::0/0"),
						GatewayId:     websiteIGW.ID(),
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
		if err != nil {
			return err
		}

		websiteSecGroup, err := ec2.NewSecurityGroup(
			ctx,
			"website-sg",
			&ec2.SecurityGroupArgs{
				VpcId: websiteVPC.ID(),
			},
		)
		if err != nil {
			return err
		}

		_, err = vpc.NewSecurityGroupIngressRule(
			ctx,
			"ingress-allow-ipv4-https",
			&vpc.SecurityGroupIngressRuleArgs{
				SecurityGroupId: websiteSecGroup.ID(),
				CidrIpv4:        pulumi.String("0.0.0.0/0"),
				IpProtocol:      pulumi.String("tcp"),
				FromPort:        pulumi.Int(443),
				ToPort:          pulumi.Int(443),
			},
		)
		if err != nil {
			return err
		}

		_, err = vpc.NewSecurityGroupIngressRule(
			ctx,
			"ingress-allow-ipv6-https",
			&vpc.SecurityGroupIngressRuleArgs{
				SecurityGroupId: websiteSecGroup.ID(),
				CidrIpv6:        pulumi.String("::0/0"),
				IpProtocol:      pulumi.String("tcp"),
				FromPort:        pulumi.Int(443),
				ToPort:          pulumi.Int(443),
			},
		)
		if err != nil {
			return err
		}

		_, err = vpc.NewSecurityGroupIngressRule(
			ctx,
			"ingress-allow-ssh-all",
			&vpc.SecurityGroupIngressRuleArgs{
				SecurityGroupId: websiteSecGroup.ID(),
				CidrIpv4:        pulumi.String("0.0.0.0/0"),
				IpProtocol:      pulumi.String("tcp"),
				FromPort:        pulumi.Int(22),
				ToPort:          pulumi.Int(22),
			},
		)
		if err != nil {
			return err
		}

		_, err = vpc.NewSecurityGroupEgressRule(
			ctx,
			"egress-allow-ipv4-all",
			&vpc.SecurityGroupEgressRuleArgs{
				SecurityGroupId: websiteSecGroup.ID(),
				CidrIpv4:        pulumi.String("0.0.0.0/0"),
				IpProtocol:      pulumi.String("-1"),
			},
		)
		if err != nil {
			return err
		}

		_, err = vpc.NewSecurityGroupEgressRule(
			ctx,
			"egress-allow-ipv6-all",
			&vpc.SecurityGroupEgressRuleArgs{
				SecurityGroupId: websiteSecGroup.ID(),
				CidrIpv6:        pulumi.String("::0/0"),
				IpProtocol:      pulumi.String("-1"),
			},
		)
		if err != nil {
			return err
		}

		// EC2 setup

		ec2KeyPair, err := ec2.NewKeyPair(
			ctx,
			"ryans-public-ssh-key",
			&ec2.KeyPairArgs{
				KeyName:   pulumi.String("ryans-public-ssh-key"),
				PublicKey: pulumi.String("ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIL11BYUiOu1vO4kecR8P/gxHUDjXeWljOC4/uT0EmA6N openpgp:0xB6108B0A"),
			},
		)

		websiteInstance, err := ec2.NewInstance(
			ctx,
			"website-ec2",
			&ec2.InstanceArgs{
				Ami:                      pulumi.String("ami-0b09627181c8d5778"), // Amazon Linux 2023 kernel 6.1
				InstanceType:             ec2.InstanceType_T2_Micro,
				SubnetId:                 websiteVPCPublicSubnet.ID(),
				AssociatePublicIpAddress: pulumi.Bool(true),
				VpcSecurityGroupIds: pulumi.StringArray{
					websiteSecGroup.ID(),
				},
				UserData: pulumi.String(`
#!/usr/bin/env bash
sudo dnf --refresh update -y
sudo dnf install docker -y
sudo systemctl start docker
sudo systemctl enable docker
sudo usermod -aG docker $USER
newgrp docker
				`),
				KeyName: ec2KeyPair.KeyName,
			},
		)
		if err != nil {
			return err
		}

		ctx.Export("websiteVPC", websiteVPC.ID())
		ctx.Export("websitePublicIP", websiteInstance.PublicIp)

		return nil
	})
}
