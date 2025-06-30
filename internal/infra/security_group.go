package infra

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/vpc"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateSecurityGroup allows the HTTP, HTTPS, and SSH ingress access to the
// provided VPC and allows all egress traffic for both, IPv4 and IPv6 requests.
func CreateSecurityGroup(
	ctx *pulumi.Context,
	websiteVPC *ec2.Vpc,
) (*ec2.SecurityGroup, error) {
	sg, err := ec2.NewSecurityGroup(
		ctx,
		"website-sg",
		&ec2.SecurityGroupArgs{
			VpcId: websiteVPC.ID(),
		},
	)
	if err != nil {
		return nil, err
	}

	err = addHTTPIngressRulesToSecurityGroup(ctx, sg)
	if err != nil {
		return nil, err
	}

	err = addHTTPSIngressRulesToSecurityGroup(ctx, sg)
	if err != nil {
		return nil, err
	}

	err = addSSHIngressRuleToSecurityGroup(ctx, sg)
	if err != nil {
		return nil, err
	}

	err = addEgressRulesToSecurityGroup(ctx, sg)
	if err != nil {
		return nil, err
	}

	return sg, nil
}

func addHTTPIngressRulesToSecurityGroup(
	ctx *pulumi.Context,
	sg *ec2.SecurityGroup,
) error {
	_, err := vpc.NewSecurityGroupIngressRule(
		ctx,
		"ingress-allow-ipv4-http",
		&vpc.SecurityGroupIngressRuleArgs{
			SecurityGroupId: sg.ID(),
			CidrIpv4:        pulumi.String("0.0.0.0/0"),
			IpProtocol:      pulumi.String("tcp"),
			FromPort:        pulumi.Int(80),
			ToPort:          pulumi.Int(80),
		},
	)
	if err != nil {
		return err
	}

	_, err = vpc.NewSecurityGroupIngressRule(
		ctx,
		"ingress-allow-ipv6-http",
		&vpc.SecurityGroupIngressRuleArgs{
			SecurityGroupId: sg.ID(),
			CidrIpv6:        pulumi.String("::0/0"),
			IpProtocol:      pulumi.String("tcp"),
			FromPort:        pulumi.Int(80),
			ToPort:          pulumi.Int(80),
		},
	)

	return err
}

func addHTTPSIngressRulesToSecurityGroup(
	ctx *pulumi.Context,
	sg *ec2.SecurityGroup,
) error {
	_, err := vpc.NewSecurityGroupIngressRule(
		ctx,
		"ingress-allow-ipv4-https",
		&vpc.SecurityGroupIngressRuleArgs{
			SecurityGroupId: sg.ID(),
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
			SecurityGroupId: sg.ID(),
			CidrIpv6:        pulumi.String("::0/0"),
			IpProtocol:      pulumi.String("tcp"),
			FromPort:        pulumi.Int(443),
			ToPort:          pulumi.Int(443),
		},
	)

	return err
}

func addSSHIngressRuleToSecurityGroup(
	ctx *pulumi.Context,
	sg *ec2.SecurityGroup,
) error {
	_, err := vpc.NewSecurityGroupIngressRule(
		ctx,
		"ingress-allow-ssh-all",
		&vpc.SecurityGroupIngressRuleArgs{
			SecurityGroupId: sg.ID(),
			CidrIpv4:        pulumi.String("0.0.0.0/0"),
			IpProtocol:      pulumi.String("tcp"),
			FromPort:        pulumi.Int(22),
			ToPort:          pulumi.Int(22),
		},
	)

	return err
}

func addEgressRulesToSecurityGroup(
	ctx *pulumi.Context,
	sg *ec2.SecurityGroup,
) error {
	_, err := vpc.NewSecurityGroupEgressRule(
		ctx,
		"egress-allow-ipv4-all",
		&vpc.SecurityGroupEgressRuleArgs{
			SecurityGroupId: sg.ID(),
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
			SecurityGroupId: sg.ID(),
			CidrIpv6:        pulumi.String("::0/0"),
			IpProtocol:      pulumi.String("-1"),
		},
	)

	return err
}
