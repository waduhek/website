package infra

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateLogGroup creates the log group and log stream for uploading logs.
func CreateLogGroup(ctx *pulumi.Context) (*cloudwatch.LogGroup, error) {
	logGroup, err := cloudwatch.NewLogGroup(
		ctx,
		"website-logs",
		&cloudwatch.LogGroupArgs{
			Name:            pulumi.String("website-logs"),
			RetentionInDays: pulumi.IntPtr(7),
		},
	)
	if err != nil {
		return nil, err
	}

	_, err = cloudwatch.NewLogStream(
		ctx,
		"website-logs-stream",
		&cloudwatch.LogStreamArgs{
			Name:         pulumi.String("website-logs-stream"),
			LogGroupName: logGroup.Name,
		},
	)
	if err != nil {
		return nil, err
	}

	_, err = cloudwatch.NewLogStream(
		ctx,
		"website-metrics-stream",
		&cloudwatch.LogStreamArgs{
			Name:         pulumi.String("website-metrics-stream"),
			LogGroupName: logGroup.Name,
		},
	)

	return logGroup, nil
}
