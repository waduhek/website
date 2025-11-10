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
	assumeRolePolicy, err := createAssumeRolePolicy(ctx)
	if err != nil {
		return nil, err
	}

	inlinePolicy := createIAMPolicyDoc(ctx, logGroup)

	role, err := createIAMRole(ctx, assumeRolePolicy, inlinePolicy)
	if err != nil {
		return nil, err
	}

	return createInstanceProfile(ctx, role)
}

func createAssumeRolePolicy(ctx *pulumi.Context) (string, error) {
	doc, err := iam.GetPolicyDocument(
		ctx,
		&iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"ec2.amazonaws.com",
							},
						},
					},
					Actions: []string{
						"sts:AssumeRole",
					},
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return doc.Json, nil
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
							fmt.Sprintf("%s:*:*", arn),
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
	assumeRolePolicy string,
	inlinePolicy pulumi.Output,
) (*iam.Role, error) {
	role, err := iam.NewRole(
		ctx,
		"website-ec2-role",
		&iam.RoleArgs{
			Description:      pulumi.String("Website instance role"),
			AssumeRolePolicy: pulumi.String(assumeRolePolicy),
		},
	)
	if err != nil {
		return nil, err
	}

	_, err = iam.NewRolePolicy(
		ctx,
		"website-role-policy",
		&iam.RolePolicyArgs{
			Name:   pulumi.String("website-role-policy"),
			Role:   role.ID(),
			Policy: inlinePolicy,
		},
	)
	if err != nil {
		return nil, err
	}

	return role, err
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
