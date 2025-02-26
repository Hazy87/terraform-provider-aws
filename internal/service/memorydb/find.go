package memorydb

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/memorydb"
	"github.com/hashicorp/aws-sdk-go-base/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

func FindSubnetGroupByName(ctx context.Context, conn *memorydb.MemoryDB, name string) (*memorydb.SubnetGroup, error) {
	input := memorydb.DescribeSubnetGroupsInput{
		SubnetGroupName: aws.String(name),
	}

	output, err := conn.DescribeSubnetGroupsWithContext(ctx, &input)

	if tfawserr.ErrCodeEquals(err, memorydb.ErrCodeSubnetGroupNotFoundFault) {
		return nil, &resource.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	if output == nil || len(output.SubnetGroups) == 0 || output.SubnetGroups[0] == nil {
		return nil, tfresource.NewEmptyResultError(input)
	}

	if count := len(output.SubnetGroups); count > 1 {
		return nil, tfresource.NewTooManyResultsError(count, input)
	}

	return output.SubnetGroups[0], nil
}

func FindUserByName(ctx context.Context, conn *memorydb.MemoryDB, name string) (*memorydb.User, error) {
	input := memorydb.DescribeUsersInput{
		UserName: aws.String(name),
	}

	output, err := conn.DescribeUsersWithContext(ctx, &input)

	if tfawserr.ErrCodeEquals(err, memorydb.ErrCodeUserNotFoundFault) {
		return nil, &resource.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	if output == nil || len(output.Users) == 0 || output.Users[0] == nil {
		return nil, tfresource.NewEmptyResultError(input)
	}

	if count := len(output.Users); count > 1 {
		return nil, tfresource.NewTooManyResultsError(count, input)
	}

	return output.Users[0], nil
}
