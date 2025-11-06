package infra

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateWebsiteEC2InstanceProfile creates the policy, role, and instance
// profile for the EC2 instance.
func CreateWebsiteEC2InstanceProfile(
	ctx *pulumi.Context,
	logGroup *cloudwatch.LogGroup,
) (*iam.InstanceProfile, error) {
	policyDocOutput := createIAMPolicyDoc(ctx, logGroup)

	role, err := createIAMRole(ctx, policyDocOutput)
	if err != nil {
		return nil, err
	}

	return createInstanceProfile(ctx, role)
}

func createIAMPolicyDoc(
	ctx *pulumi.Context,
	l *cloudwatch.LogGroup,
) pulumi.Output {
	return l.Arn.ApplyT(func(arn string) (string, error) {
		policyDoc, err := iam.GetPolicyDocument(
			ctx,
			&iam.GetPolicyDocumentArgs{
				Statements: []iam.GetPolicyDocumentStatement{
					{
						Effect: pulumi.StringRef("Allow"),
						Actions: []string{
							"logs:PutLogEvents",
							"logs:DescribeLogGroups",
							"logs:DescribeLogStreams",
						},
						Resources: []string{
							fmt.Sprintf("%s:*", arn),
						},
					},
					{
						Effect: pulumi.StringRef("Allow"),
						Actions: []string{
							"xray:PutTraceSegments",
							"xray:PutTelemetryRecords",
							"xray:GetSamplingRules",
							"xray:GetSamplingTargets",
							"xray:GetSamplingStatisticSummaries",
						},
						Resources: []string{
							"*",
						},
					},
				},
			},
		)

		if err != nil {
			return "", err
		}

		return policyDoc.Json, nil
	})
}

func createIAMRole(
	ctx *pulumi.Context,
	policyOutput pulumi.Output,
) (*iam.Role, error) {
	return iam.NewRole(
		ctx,
		"website-ec2-role",
		&iam.RoleArgs{
			Description:      pulumi.String("Website instance role"),
			AssumeRolePolicy: policyOutput,
		},
	)
}

func createInstanceProfile(
	ctx *pulumi.Context,
	role *iam.Role,
) (*iam.InstanceProfile, error) {
	return iam.NewInstanceProfile(
		ctx,
		"website-ec2-role",
		&iam.InstanceProfileArgs{
			Name: pulumi.String("website-ec2-role"),
			Role: role.Name,
		},
	)
}
