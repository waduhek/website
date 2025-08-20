package infra

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateEC2Instance creates a new EC2 instance and associates a key pair to it.
func CreateEC2Instance(
	ctx *pulumi.Context,
	subnet *ec2.Subnet,
	sg *ec2.SecurityGroup,
) (*ec2.Instance, error) {
	ec2KeyPair, err := createKeyPair(ctx)
	if err != nil {
		return nil, err
	}

	return createInstance(ctx, subnet, sg, ec2KeyPair)
}

func createKeyPair(ctx *pulumi.Context) (*ec2.KeyPair, error) {
	return ec2.NewKeyPair(
		ctx,
		"ryans-public-ssh-key",
		&ec2.KeyPairArgs{
			KeyName:   pulumi.String("ryans-public-ssh-key"),
			PublicKey: pulumi.String("ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIL11BYUiOu1vO4kecR8P/gxHUDjXeWljOC4/uT0EmA6N openpgp:0xB6108B0A"),
		},
	)
}

func createInstance(
	ctx *pulumi.Context,
	subnet *ec2.Subnet,
	sg *ec2.SecurityGroup,
	keyPair *ec2.KeyPair,
) (*ec2.Instance, error) {
	return ec2.NewInstance(
		ctx,
		"website-ec2",
		&ec2.InstanceArgs{
			Ami:                      pulumi.String("ami-0b09627181c8d5778"), // Amazon Linux 2023 kernel 6.1
			InstanceType:             ec2.InstanceType_T2_Micro,
			SubnetId:                 subnet.ID(),
			AssociatePublicIpAddress: pulumi.Bool(true),
			Ipv6AddressCount:         pulumi.Int(1),
			VpcSecurityGroupIds: pulumi.StringArray{
				sg.ID(),
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
			KeyName: keyPair.KeyName,
		},
	)
}
