package util

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/stretchr/testify/mock"
)

type MockDynamoDBClient struct {
	mock.Mock
	dynamodbiface.DynamoDBAPI
}

func (m *MockDynamoDBClient) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	args := m.Called(input)
	return args.Get(0).(*dynamodb.QueryOutput), args.Error(1)
}

func (m *MockDynamoDBClient) BatchGetItem(*dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) BatchGetItemWithContext(aws.Context, *dynamodb.BatchGetItemInput, ...request.Option) (*dynamodb.BatchGetItemOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) BatchGetItemRequest(*dynamodb.BatchGetItemInput) (*request.Request, *dynamodb.BatchGetItemOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) BatchGetItemPages(*dynamodb.BatchGetItemInput, func(*dynamodb.BatchGetItemOutput, bool) bool) error {
	return nil
}
func (m *MockDynamoDBClient) BatchGetItemPagesWithContext(aws.Context, *dynamodb.BatchGetItemInput, func(*dynamodb.BatchGetItemOutput, bool) bool, ...request.Option) error {
	return nil
}
func (m *MockDynamoDBClient) BatchWriteItem(*dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) BatchWriteItemWithContext(aws.Context, *dynamodb.BatchWriteItemInput, ...request.Option) (*dynamodb.BatchWriteItemOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) BatchWriteItemRequest(*dynamodb.BatchWriteItemInput) (*request.Request, *dynamodb.BatchWriteItemOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) CreateBackup(*dynamodb.CreateBackupInput) (*dynamodb.CreateBackupOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) CreateBackupWithContext(aws.Context, *dynamodb.CreateBackupInput, ...request.Option) (*dynamodb.CreateBackupOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) CreateBackupRequest(*dynamodb.CreateBackupInput) (*request.Request, *dynamodb.CreateBackupOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) CreateGlobalTable(*dynamodb.CreateGlobalTableInput) (*dynamodb.CreateGlobalTableOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) CreateGlobalTableWithContext(aws.Context, *dynamodb.CreateGlobalTableInput, ...request.Option) (*dynamodb.CreateGlobalTableOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) CreateGlobalTableRequest(*dynamodb.CreateGlobalTableInput) (*request.Request, *dynamodb.CreateGlobalTableOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) CreateTable(*dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) CreateTableWithContext(aws.Context, *dynamodb.CreateTableInput, ...request.Option) (*dynamodb.CreateTableOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) CreateTableRequest(*dynamodb.CreateTableInput) (*request.Request, *dynamodb.CreateTableOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) DeleteBackup(*dynamodb.DeleteBackupInput) (*dynamodb.DeleteBackupOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) DeleteBackupWithContext(aws.Context, *dynamodb.DeleteBackupInput, ...request.Option) (*dynamodb.DeleteBackupOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) DeleteBackupRequest(*dynamodb.DeleteBackupInput) (*request.Request, *dynamodb.DeleteBackupOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) DeleteItem(*dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) DeleteItemWithContext(aws.Context, *dynamodb.DeleteItemInput, ...request.Option) (*dynamodb.DeleteItemOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) DeleteItemRequest(*dynamodb.DeleteItemInput) (*request.Request, *dynamodb.DeleteItemOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) DeleteTable(*dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) DeleteTableWithContext(aws.Context, *dynamodb.DeleteTableInput, ...request.Option) (*dynamodb.DeleteTableOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) DeleteTableRequest(*dynamodb.DeleteTableInput) (*request.Request, *dynamodb.DeleteTableOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) DescribeBackup(*dynamodb.DescribeBackupInput) (*dynamodb.DescribeBackupOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) DescribeBackupWithContext(aws.Context, *dynamodb.DescribeBackupInput, ...request.Option) (*dynamodb.DescribeBackupOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) DescribeBackupRequest(*dynamodb.DescribeBackupInput) (*request.Request, *dynamodb.DescribeBackupOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) DescribeContinuousBackups(*dynamodb.DescribeContinuousBackupsInput) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) DescribeContinuousBackupsWithContext(aws.Context, *dynamodb.DescribeContinuousBackupsInput, ...request.Option) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) DescribeContinuousBackupsRequest(*dynamodb.DescribeContinuousBackupsInput) (*request.Request, *dynamodb.DescribeContinuousBackupsOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) DescribeGlobalTable(*dynamodb.DescribeGlobalTableInput) (*dynamodb.DescribeGlobalTableOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) DescribeGlobalTableWithContext(aws.Context, *dynamodb.DescribeGlobalTableInput, ...request.Option) (*dynamodb.DescribeGlobalTableOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) DescribeGlobalTableRequest(*dynamodb.DescribeGlobalTableInput) (*request.Request, *dynamodb.DescribeGlobalTableOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) DescribeGlobalTableSettings(*dynamodb.DescribeGlobalTableSettingsInput) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) DescribeGlobalTableSettingsWithContext(aws.Context, *dynamodb.DescribeGlobalTableSettingsInput, ...request.Option) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) DescribeGlobalTableSettingsRequest(*dynamodb.DescribeGlobalTableSettingsInput) (*request.Request, *dynamodb.DescribeGlobalTableSettingsOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) DescribeLimits(*dynamodb.DescribeLimitsInput) (*dynamodb.DescribeLimitsOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) DescribeLimitsWithContext(aws.Context, *dynamodb.DescribeLimitsInput, ...request.Option) (*dynamodb.DescribeLimitsOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) DescribeLimitsRequest(*dynamodb.DescribeLimitsInput) (*request.Request, *dynamodb.DescribeLimitsOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) DescribeTable(*dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) DescribeTableWithContext(aws.Context, *dynamodb.DescribeTableInput, ...request.Option) (*dynamodb.DescribeTableOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) DescribeTableRequest(*dynamodb.DescribeTableInput) (*request.Request, *dynamodb.DescribeTableOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) DescribeTimeToLive(*dynamodb.DescribeTimeToLiveInput) (*dynamodb.DescribeTimeToLiveOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) DescribeTimeToLiveWithContext(aws.Context, *dynamodb.DescribeTimeToLiveInput, ...request.Option) (*dynamodb.DescribeTimeToLiveOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) DescribeTimeToLiveRequest(*dynamodb.DescribeTimeToLiveInput) (*request.Request, *dynamodb.DescribeTimeToLiveOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) GetItemWithContext(aws.Context, *dynamodb.GetItemInput, ...request.Option) (*dynamodb.GetItemOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) GetItemRequest(*dynamodb.GetItemInput) (*request.Request, *dynamodb.GetItemOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) ListBackups(*dynamodb.ListBackupsInput) (*dynamodb.ListBackupsOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) ListBackupsWithContext(aws.Context, *dynamodb.ListBackupsInput, ...request.Option) (*dynamodb.ListBackupsOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) ListBackupsRequest(*dynamodb.ListBackupsInput) (*request.Request, *dynamodb.ListBackupsOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) ListGlobalTables(*dynamodb.ListGlobalTablesInput) (*dynamodb.ListGlobalTablesOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) ListGlobalTablesWithContext(aws.Context, *dynamodb.ListGlobalTablesInput, ...request.Option) (*dynamodb.ListGlobalTablesOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) ListGlobalTablesRequest(*dynamodb.ListGlobalTablesInput) (*request.Request, *dynamodb.ListGlobalTablesOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) ListTables(*dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) ListTablesWithContext(aws.Context, *dynamodb.ListTablesInput, ...request.Option) (*dynamodb.ListTablesOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) ListTablesRequest(*dynamodb.ListTablesInput) (*request.Request, *dynamodb.ListTablesOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) ListTablesPages(*dynamodb.ListTablesInput, func(*dynamodb.ListTablesOutput, bool) bool) error {
	return nil
}
func (m *MockDynamoDBClient) ListTablesPagesWithContext(aws.Context, *dynamodb.ListTablesInput, func(*dynamodb.ListTablesOutput, bool) bool, ...request.Option) error {
	return nil
}
func (m *MockDynamoDBClient) ListTagsOfResource(*dynamodb.ListTagsOfResourceInput) (*dynamodb.ListTagsOfResourceOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) ListTagsOfResourceWithContext(aws.Context, *dynamodb.ListTagsOfResourceInput, ...request.Option) (*dynamodb.ListTagsOfResourceOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) ListTagsOfResourceRequest(*dynamodb.ListTagsOfResourceInput) (*request.Request, *dynamodb.ListTagsOfResourceOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) PutItemWithContext(aws.Context, *dynamodb.PutItemInput, ...request.Option) (*dynamodb.PutItemOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) PutItemRequest(*dynamodb.PutItemInput) (*request.Request, *dynamodb.PutItemOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) QueryWithContext(aws.Context, *dynamodb.QueryInput, ...request.Option) (*dynamodb.QueryOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) QueryRequest(*dynamodb.QueryInput) (*request.Request, *dynamodb.QueryOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) QueryPages(*dynamodb.QueryInput, func(*dynamodb.QueryOutput, bool) bool) error {
	return nil
}
func (m *MockDynamoDBClient) QueryPagesWithContext(aws.Context, *dynamodb.QueryInput, func(*dynamodb.QueryOutput, bool) bool, ...request.Option) error {
	return nil
}
func (m *MockDynamoDBClient) RestoreTableFromBackup(*dynamodb.RestoreTableFromBackupInput) (*dynamodb.RestoreTableFromBackupOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) RestoreTableFromBackupWithContext(aws.Context, *dynamodb.RestoreTableFromBackupInput, ...request.Option) (*dynamodb.RestoreTableFromBackupOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) RestoreTableFromBackupRequest(*dynamodb.RestoreTableFromBackupInput) (*request.Request, *dynamodb.RestoreTableFromBackupOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) RestoreTableToPointInTime(*dynamodb.RestoreTableToPointInTimeInput) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) RestoreTableToPointInTimeWithContext(aws.Context, *dynamodb.RestoreTableToPointInTimeInput, ...request.Option) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) RestoreTableToPointInTimeRequest(*dynamodb.RestoreTableToPointInTimeInput) (*request.Request, *dynamodb.RestoreTableToPointInTimeOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) Scan(*dynamodb.ScanInput) (*dynamodb.ScanOutput, error) { return nil, nil }
func (m *MockDynamoDBClient) ScanWithContext(aws.Context, *dynamodb.ScanInput, ...request.Option) (*dynamodb.ScanOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) ScanRequest(*dynamodb.ScanInput) (*request.Request, *dynamodb.ScanOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) ScanPages(*dynamodb.ScanInput, func(*dynamodb.ScanOutput, bool) bool) error {
	return nil
}
func (m *MockDynamoDBClient) ScanPagesWithContext(aws.Context, *dynamodb.ScanInput, func(*dynamodb.ScanOutput, bool) bool, ...request.Option) error {
	return nil
}
func (m *MockDynamoDBClient) TagResource(*dynamodb.TagResourceInput) (*dynamodb.TagResourceOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) TagResourceWithContext(aws.Context, *dynamodb.TagResourceInput, ...request.Option) (*dynamodb.TagResourceOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) TagResourceRequest(*dynamodb.TagResourceInput) (*request.Request, *dynamodb.TagResourceOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) UntagResource(*dynamodb.UntagResourceInput) (*dynamodb.UntagResourceOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) UntagResourceWithContext(aws.Context, *dynamodb.UntagResourceInput, ...request.Option) (*dynamodb.UntagResourceOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) UntagResourceRequest(*dynamodb.UntagResourceInput) (*request.Request, *dynamodb.UntagResourceOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) UpdateContinuousBackups(*dynamodb.UpdateContinuousBackupsInput) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) UpdateContinuousBackupsWithContext(aws.Context, *dynamodb.UpdateContinuousBackupsInput, ...request.Option) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) UpdateContinuousBackupsRequest(*dynamodb.UpdateContinuousBackupsInput) (*request.Request, *dynamodb.UpdateContinuousBackupsOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) UpdateGlobalTable(*dynamodb.UpdateGlobalTableInput) (*dynamodb.UpdateGlobalTableOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) UpdateGlobalTableWithContext(aws.Context, *dynamodb.UpdateGlobalTableInput, ...request.Option) (*dynamodb.UpdateGlobalTableOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) UpdateGlobalTableRequest(*dynamodb.UpdateGlobalTableInput) (*request.Request, *dynamodb.UpdateGlobalTableOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) UpdateGlobalTableSettings(*dynamodb.UpdateGlobalTableSettingsInput) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) UpdateGlobalTableSettingsWithContext(aws.Context, *dynamodb.UpdateGlobalTableSettingsInput, ...request.Option) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) UpdateGlobalTableSettingsRequest(*dynamodb.UpdateGlobalTableSettingsInput) (*request.Request, *dynamodb.UpdateGlobalTableSettingsOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) UpdateItem(*dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) UpdateItemWithContext(aws.Context, *dynamodb.UpdateItemInput, ...request.Option) (*dynamodb.UpdateItemOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) UpdateItemRequest(*dynamodb.UpdateItemInput) (*request.Request, *dynamodb.UpdateItemOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) UpdateTable(*dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) UpdateTableWithContext(aws.Context, *dynamodb.UpdateTableInput, ...request.Option) (*dynamodb.UpdateTableOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) UpdateTableRequest(*dynamodb.UpdateTableInput) (*request.Request, *dynamodb.UpdateTableOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) UpdateTimeToLive(*dynamodb.UpdateTimeToLiveInput) (*dynamodb.UpdateTimeToLiveOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) UpdateTimeToLiveWithContext(aws.Context, *dynamodb.UpdateTimeToLiveInput, ...request.Option) (*dynamodb.UpdateTimeToLiveOutput, error) {
	return nil, nil
}
func (m *MockDynamoDBClient) UpdateTimeToLiveRequest(*dynamodb.UpdateTimeToLiveInput) (*request.Request, *dynamodb.UpdateTimeToLiveOutput) {
	return nil, nil
}
func (m *MockDynamoDBClient) WaitUntilTableExists(*dynamodb.DescribeTableInput) error { return nil }
func (m *MockDynamoDBClient) WaitUntilTableExistsWithContext(aws.Context, *dynamodb.DescribeTableInput, ...request.WaiterOption) error {
	return nil
}
func (m *MockDynamoDBClient) WaitUntilTableNotExists(*dynamodb.DescribeTableInput) error {
	return nil
}
func (m *MockDynamoDBClient) WaitUntilTableNotExistsWithContext(aws.Context, *dynamodb.DescribeTableInput, ...request.WaiterOption) error {
	return nil
}
