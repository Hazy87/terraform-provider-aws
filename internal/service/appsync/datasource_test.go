package appsync_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/appsync"
	"github.com/hashicorp/aws-sdk-go-base/tfawserr"
	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfappsync "github.com/hashicorp/terraform-provider-aws/internal/service/appsync"
)

func testAccAppSyncDataSource_basic(t *testing.T) {
	rName := fmt.Sprintf("tfacctest%d", sdkacctest.RandInt())
	resourceName := "aws_appsync_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t); acctest.PreCheckPartitionHasService(appsync.EndpointsID, t) },
		ErrorCheck:   acctest.ErrorCheck(t, appsync.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: testAccCheckDestroyDataSource,
		Steps: []resource.TestStep{
			{
				Config: testAccAppsyncDatasourceConfig_Type_None(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExistsDataSource(resourceName),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "appsync", regexp.MustCompile(fmt.Sprintf("apis/.+/datasources/%s", rName))),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttr(resourceName, "dynamodb_config.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "elasticsearch_config.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "http_config.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "lambda_config.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "type", "NONE"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAppSyncDataSource_description(t *testing.T) {
	rName := fmt.Sprintf("tfacctest%d", sdkacctest.RandInt())
	resourceName := "aws_appsync_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t); acctest.PreCheckPartitionHasService(appsync.EndpointsID, t) },
		ErrorCheck:   acctest.ErrorCheck(t, appsync.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: testAccCheckDestroyDataSource,
		Steps: []resource.TestStep{
			{
				Config: testAccAppsyncDatasourceConfig_Description(rName, "description1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExistsDataSource(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "description1"),
				),
			},
			{
				Config: testAccAppsyncDatasourceConfig_Description(rName, "description2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExistsDataSource(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAppSyncDataSource_DynamoDB_region(t *testing.T) {
	rName := fmt.Sprintf("tfacctest%d", sdkacctest.RandInt())
	resourceName := "aws_appsync_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t); acctest.PreCheckPartitionHasService(appsync.EndpointsID, t) },
		ErrorCheck:   acctest.ErrorCheck(t, appsync.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: testAccCheckDestroyDataSource,
		Steps: []resource.TestStep{
			{
				Config: testAccAppsyncDatasourceConfig_DynamoDBConfig_Region(rName, acctest.Region()),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExistsDataSource(resourceName),
					resource.TestCheckResourceAttr(resourceName, "dynamodb_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "dynamodb_config.0.region", acctest.Region()),
				),
			},
			{
				Config: testAccAppsyncDatasourceConfig_Type_DynamoDB(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExistsDataSource(resourceName),
					resource.TestCheckResourceAttr(resourceName, "dynamodb_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "dynamodb_config.0.region", acctest.Region()),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAppSyncDataSource_DynamoDB_useCallerCredentials(t *testing.T) {
	rName := fmt.Sprintf("tfacctest%d", sdkacctest.RandInt())
	resourceName := "aws_appsync_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t); acctest.PreCheckPartitionHasService(appsync.EndpointsID, t) },
		ErrorCheck:   acctest.ErrorCheck(t, appsync.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: testAccCheckDestroyDataSource,
		Steps: []resource.TestStep{
			{
				Config: testAccAppsyncDatasourceConfig_DynamoDBConfig_UseCallerCredentials(rName, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExistsDataSource(resourceName),
					resource.TestCheckResourceAttr(resourceName, "dynamodb_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "dynamodb_config.0.use_caller_credentials", "true"),
				),
			},
			{
				Config: testAccAppsyncDatasourceConfig_DynamoDBConfig_UseCallerCredentials(rName, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExistsDataSource(resourceName),
					resource.TestCheckResourceAttr(resourceName, "dynamodb_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "dynamodb_config.0.use_caller_credentials", "false"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAppSyncDataSource_Elasticsearch_region(t *testing.T) {
	rName := fmt.Sprintf("tfacctest%d", sdkacctest.RandInt())
	resourceName := "aws_appsync_datasource.test"

	// Keep this test Parallel as it takes considerably longer to run than any non-Elasticsearch tests.
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t); acctest.PreCheckPartitionHasService(appsync.EndpointsID, t) },
		ErrorCheck:   acctest.ErrorCheck(t, appsync.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: testAccCheckDestroyDataSource,
		Steps: []resource.TestStep{
			{
				Config: testAccAppsyncDatasourceConfig_ElasticsearchConfig_Region(rName, acctest.Region()),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExistsDataSource(resourceName),
					resource.TestCheckResourceAttr(resourceName, "elasticsearch_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "elasticsearch_config.0.region", acctest.Region()),
				),
			},
			{
				Config: testAccAppsyncDatasourceConfig_Type_Elasticsearch(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExistsDataSource(resourceName),
					resource.TestCheckResourceAttr(resourceName, "elasticsearch_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "elasticsearch_config.0.region", acctest.Region()),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAppSyncDataSource_HTTP_endpoint(t *testing.T) {
	rName := fmt.Sprintf("tfacctest%d", sdkacctest.RandInt())
	resourceName := "aws_appsync_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t); acctest.PreCheckPartitionHasService(appsync.EndpointsID, t) },
		ErrorCheck:   acctest.ErrorCheck(t, appsync.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: testAccCheckDestroyDataSource,
		Steps: []resource.TestStep{
			{
				Config: testAccAppsyncDatasourceConfig_HTTPConfig_Endpoint(rName, "http://example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExistsDataSource(resourceName),
					resource.TestCheckResourceAttr(resourceName, "http_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "http_config.0.endpoint", "http://example.com"),
					resource.TestCheckResourceAttr(resourceName, "type", "HTTP"),
				),
			},
			{
				Config: testAccAppsyncDatasourceConfig_HTTPConfig_Endpoint(rName, "http://example.org"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExistsDataSource(resourceName),
					resource.TestCheckResourceAttr(resourceName, "http_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "http_config.0.endpoint", "http://example.org"),
					resource.TestCheckResourceAttr(resourceName, "type", "HTTP"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAppSyncDataSource_type(t *testing.T) {
	rName := fmt.Sprintf("tfacctest%d", sdkacctest.RandInt())
	resourceName := "aws_appsync_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t); acctest.PreCheckPartitionHasService(appsync.EndpointsID, t) },
		ErrorCheck:   acctest.ErrorCheck(t, appsync.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: testAccCheckDestroyDataSource,
		Steps: []resource.TestStep{
			{
				Config: testAccAppsyncDatasourceConfig_Type_None(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExistsDataSource(resourceName),
					resource.TestCheckResourceAttr(resourceName, "type", "NONE"),
				),
			},
			{
				Config: testAccAppsyncDatasourceConfig_Type_HTTP(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExistsDataSource(resourceName),
					resource.TestCheckResourceAttr(resourceName, "type", "HTTP"),
				),
			},
		},
	})
}

func testAccAppSyncDataSource_Type_dynamoDB(t *testing.T) {
	rName := fmt.Sprintf("tfacctest%d", sdkacctest.RandInt())
	dynamodbTableResourceName := "aws_dynamodb_table.test"
	iamRoleResourceName := "aws_iam_role.test"
	resourceName := "aws_appsync_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t); acctest.PreCheckPartitionHasService(appsync.EndpointsID, t) },
		ErrorCheck:   acctest.ErrorCheck(t, appsync.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: testAccCheckDestroyDataSource,
		Steps: []resource.TestStep{
			{
				Config: testAccAppsyncDatasourceConfig_Type_DynamoDB(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExistsDataSource(resourceName),
					resource.TestCheckResourceAttr(resourceName, "dynamodb_config.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "dynamodb_config.0.table_name", dynamodbTableResourceName, "name"),
					resource.TestCheckResourceAttr(resourceName, "dynamodb_config.0.region", acctest.Region()),
					resource.TestCheckResourceAttrPair(resourceName, "service_role_arn", iamRoleResourceName, "arn"),
					resource.TestCheckResourceAttr(resourceName, "type", "AMAZON_DYNAMODB"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAppSyncDataSource_Type_elasticSearch(t *testing.T) {
	rName := fmt.Sprintf("tfacctest%d", sdkacctest.RandInt())
	iamRoleResourceName := "aws_iam_role.test"
	resourceName := "aws_appsync_datasource.test"

	// Keep this test Parallel as it takes considerably longer to run than any non-Elasticsearch tests.
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t); acctest.PreCheckPartitionHasService(appsync.EndpointsID, t) },
		ErrorCheck:   acctest.ErrorCheck(t, appsync.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: testAccCheckDestroyDataSource,
		Steps: []resource.TestStep{
			{
				Config: testAccAppsyncDatasourceConfig_Type_Elasticsearch(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExistsDataSource(resourceName),
					resource.TestCheckResourceAttr(resourceName, "elasticsearch_config.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "elasticsearch_config.0.endpoint"),
					resource.TestCheckResourceAttr(resourceName, "elasticsearch_config.0.region", acctest.Region()),
					resource.TestCheckResourceAttrPair(resourceName, "service_role_arn", iamRoleResourceName, "arn"),
					resource.TestCheckResourceAttr(resourceName, "type", "AMAZON_ELASTICSEARCH"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAppSyncDataSource_Type_http(t *testing.T) {
	rName := fmt.Sprintf("tfacctest%d", sdkacctest.RandInt())
	resourceName := "aws_appsync_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t); acctest.PreCheckPartitionHasService(appsync.EndpointsID, t) },
		ErrorCheck:   acctest.ErrorCheck(t, appsync.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: testAccCheckDestroyDataSource,
		Steps: []resource.TestStep{
			{
				Config: testAccAppsyncDatasourceConfig_Type_HTTP(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExistsDataSource(resourceName),
					resource.TestCheckResourceAttr(resourceName, "http_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "http_config.0.endpoint", "http://example.com"),
					resource.TestCheckResourceAttr(resourceName, "type", "HTTP"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAppSyncDataSource_Type_lambda(t *testing.T) {
	rName := fmt.Sprintf("tfacctest%d", sdkacctest.RandInt())
	iamRoleResourceName := "aws_iam_role.test"
	lambdaFunctionResourceName := "aws_lambda_function.test"
	resourceName := "aws_appsync_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t); acctest.PreCheckPartitionHasService(appsync.EndpointsID, t) },
		ErrorCheck:   acctest.ErrorCheck(t, appsync.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: testAccCheckDestroyDataSource,
		Steps: []resource.TestStep{
			{
				Config: testAccAppsyncDatasourceConfig_Type_Lambda(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExistsDataSource(resourceName),
					resource.TestCheckResourceAttr(resourceName, "lambda_config.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "lambda_config.0.function_arn", lambdaFunctionResourceName, "arn"),
					resource.TestCheckResourceAttrPair(resourceName, "service_role_arn", iamRoleResourceName, "arn"),
					resource.TestCheckResourceAttr(resourceName, "type", "AWS_LAMBDA"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAppSyncDataSource_Type_none(t *testing.T) {
	rName := fmt.Sprintf("tfacctest%d", sdkacctest.RandInt())
	resourceName := "aws_appsync_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t); acctest.PreCheckPartitionHasService(appsync.EndpointsID, t) },
		ErrorCheck:   acctest.ErrorCheck(t, appsync.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: testAccCheckDestroyDataSource,
		Steps: []resource.TestStep{
			{
				Config: testAccAppsyncDatasourceConfig_Type_None(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExistsDataSource(resourceName),
					resource.TestCheckResourceAttr(resourceName, "type", "NONE"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckDestroyDataSource(s *terraform.State) error {
	conn := acctest.Provider.Meta().(*conns.AWSClient).AppSyncConn
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_appsync_datasource" {
			continue
		}

		apiID, name, err := tfappsync.DecodeID(rs.Primary.ID)

		if err != nil {
			return err
		}

		input := &appsync.GetDataSourceInput{
			ApiId: aws.String(apiID),
			Name:  aws.String(name),
		}

		_, err = conn.GetDataSource(input)
		if err != nil {
			if tfawserr.ErrCodeEquals(err, appsync.ErrCodeNotFoundException) {
				return nil
			}
			return err
		}
	}
	return nil
}

func testAccCheckExistsDataSource(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("Resource has no ID: %s", name)
		}

		apiID, name, err := tfappsync.DecodeID(rs.Primary.ID)

		if err != nil {
			return err
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).AppSyncConn

		input := &appsync.GetDataSourceInput{
			ApiId: aws.String(apiID),
			Name:  aws.String(name),
		}

		_, err = conn.GetDataSource(input)

		return err
	}
}

func testAccAppsyncDatasourceConfig_base_DynamoDB(rName string) string {
	return fmt.Sprintf(`
resource "aws_dynamodb_table" "test" {
  hash_key       = "UserId"
  name           = %q
  read_capacity  = 1
  write_capacity = 1

  attribute {
    name = "UserId"
    type = "S"
  }
}

resource "aws_iam_role" "test" {
  name = %q

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "appsync.amazonaws.com"
      },
      "Effect": "Allow"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "test" {
  role = aws_iam_role.test.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "dynamodb:*"
      ],
      "Effect": "Allow",
      "Resource": [
        "${aws_dynamodb_table.test.arn}"
      ]
    }
  ]
}
EOF
}
`, rName, rName)
}

func testAccAppsyncDatasourceConfig_base_Elasticsearch(rName string) string {
	return fmt.Sprintf(`
resource "aws_elasticsearch_domain" "test" {
  domain_name = %q

  ebs_options {
    ebs_enabled = true
    volume_size = 10
  }
}

resource "aws_iam_role" "test" {
  name = %q

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "appsync.amazonaws.com"
      },
      "Effect": "Allow"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "test" {
  role = aws_iam_role.test.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "es:*"
      ],
      "Effect": "Allow",
      "Resource": [
        "${aws_elasticsearch_domain.test.arn}"
      ]
    }
  ]
}
EOF
}
`, rName, rName)
}

func testAccAppsyncDatasourceConfig_base_Lambda(rName string) string {
	return fmt.Sprintf(`
resource "aws_iam_role" "lambda" {
  name = "%slambda"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow"
    }
  ]
}
EOF
}

resource "aws_lambda_function" "test" {
  filename      = "test-fixtures/lambdatest.zip"
  function_name = %q
  handler       = "exports.test"
  role          = aws_iam_role.lambda.arn
  runtime       = "nodejs12.x"
}

resource "aws_iam_role" "test" {
  name = %q

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "appsync.amazonaws.com"
      },
      "Effect": "Allow"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "test" {
  role = aws_iam_role.test.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "lambda:*"
      ],
      "Effect": "Allow",
      "Resource": [
        "${aws_lambda_function.test.arn}"
      ]
    }
  ]
}
EOF
}
`, rName, rName, rName)
}

func testAccAppsyncDatasourceConfig_Description(rName, description string) string {
	return fmt.Sprintf(`
resource "aws_appsync_graphql_api" "test" {
  authentication_type = "API_KEY"
  name                = %q
}

resource "aws_appsync_datasource" "test" {
  api_id      = aws_appsync_graphql_api.test.id
  description = %q
  name        = %q
  type        = "HTTP"

  http_config {
    endpoint = "http://example.com"
  }
}
`, rName, description, rName)
}

func testAccAppsyncDatasourceConfig_DynamoDBConfig_Region(rName, region string) string {
	return testAccAppsyncDatasourceConfig_base_DynamoDB(rName) + fmt.Sprintf(`
resource "aws_appsync_graphql_api" "test" {
  authentication_type = "API_KEY"
  name                = %q
}

resource "aws_appsync_datasource" "test" {
  api_id           = aws_appsync_graphql_api.test.id
  name             = %q
  service_role_arn = aws_iam_role.test.arn
  type             = "AMAZON_DYNAMODB"

  dynamodb_config {
    region     = %q
    table_name = aws_dynamodb_table.test.name
  }
}
`, rName, rName, region)
}

func testAccAppsyncDatasourceConfig_DynamoDBConfig_UseCallerCredentials(rName string, useCallerCredentials bool) string {
	return testAccAppsyncDatasourceConfig_base_DynamoDB(rName) + fmt.Sprintf(`
resource "aws_appsync_graphql_api" "test" {
  authentication_type = "AWS_IAM"
  name                = %q
}

resource "aws_appsync_datasource" "test" {
  api_id           = aws_appsync_graphql_api.test.id
  name             = %q
  service_role_arn = aws_iam_role.test.arn
  type             = "AMAZON_DYNAMODB"

  dynamodb_config {
    table_name             = aws_dynamodb_table.test.name
    use_caller_credentials = %t
  }
}
`, rName, rName, useCallerCredentials)
}

func testAccAppsyncDatasourceConfig_ElasticsearchConfig_Region(rName, region string) string {
	return testAccAppsyncDatasourceConfig_base_Elasticsearch(rName) + fmt.Sprintf(`
resource "aws_appsync_graphql_api" "test" {
  authentication_type = "API_KEY"
  name                = %q
}

resource "aws_appsync_datasource" "test" {
  api_id           = aws_appsync_graphql_api.test.id
  name             = %q
  service_role_arn = aws_iam_role.test.arn
  type             = "AMAZON_ELASTICSEARCH"

  elasticsearch_config {
    endpoint = "https://${aws_elasticsearch_domain.test.endpoint}"
    region   = %q
  }
}
`, rName, rName, region)
}

func testAccAppsyncDatasourceConfig_HTTPConfig_Endpoint(rName, endpoint string) string {
	return fmt.Sprintf(`
resource "aws_appsync_graphql_api" "test" {
  authentication_type = "API_KEY"
  name                = %q
}

resource "aws_appsync_datasource" "test" {
  api_id = aws_appsync_graphql_api.test.id
  name   = %q
  type   = "HTTP"

  http_config {
    endpoint = %q
  }
}
`, rName, rName, endpoint)
}

func testAccAppsyncDatasourceConfig_Type_DynamoDB(rName string) string {
	return testAccAppsyncDatasourceConfig_base_DynamoDB(rName) + fmt.Sprintf(`
resource "aws_appsync_graphql_api" "test" {
  authentication_type = "API_KEY"
  name                = %q
}

resource "aws_appsync_datasource" "test" {
  api_id           = aws_appsync_graphql_api.test.id
  name             = %q
  service_role_arn = aws_iam_role.test.arn
  type             = "AMAZON_DYNAMODB"

  dynamodb_config {
    table_name = aws_dynamodb_table.test.name
  }
}
`, rName, rName)
}

func testAccAppsyncDatasourceConfig_Type_Elasticsearch(rName string) string {
	return testAccAppsyncDatasourceConfig_base_Elasticsearch(rName) + fmt.Sprintf(`
resource "aws_appsync_graphql_api" "test" {
  authentication_type = "API_KEY"
  name                = %q
}

resource "aws_appsync_datasource" "test" {
  api_id           = aws_appsync_graphql_api.test.id
  name             = %q
  service_role_arn = aws_iam_role.test.arn
  type             = "AMAZON_ELASTICSEARCH"

  elasticsearch_config {
    endpoint = "https://${aws_elasticsearch_domain.test.endpoint}"
  }
}
`, rName, rName)
}

func testAccAppsyncDatasourceConfig_Type_HTTP(rName string) string {
	return fmt.Sprintf(`
resource "aws_appsync_graphql_api" "test" {
  authentication_type = "API_KEY"
  name                = %q
}

resource "aws_appsync_datasource" "test" {
  api_id = aws_appsync_graphql_api.test.id
  name   = %q
  type   = "HTTP"

  http_config {
    endpoint = "http://example.com"
  }
}
`, rName, rName)
}

func testAccAppsyncDatasourceConfig_Type_Lambda(rName string) string {
	return testAccAppsyncDatasourceConfig_base_Lambda(rName) + fmt.Sprintf(`
resource "aws_appsync_graphql_api" "test" {
  authentication_type = "API_KEY"
  name                = %q
}

resource "aws_appsync_datasource" "test" {
  api_id           = aws_appsync_graphql_api.test.id
  name             = %q
  service_role_arn = aws_iam_role.test.arn
  type             = "AWS_LAMBDA"

  lambda_config {
    function_arn = aws_lambda_function.test.arn
  }
}
`, rName, rName)
}

func testAccAppsyncDatasourceConfig_Type_None(rName string) string {
	return fmt.Sprintf(`
resource "aws_appsync_graphql_api" "test" {
  authentication_type = "API_KEY"
  name                = %q
}

resource "aws_appsync_datasource" "test" {
  api_id = aws_appsync_graphql_api.test.id
  name   = %q
  type   = "NONE"
}
`, rName, rName)
}
