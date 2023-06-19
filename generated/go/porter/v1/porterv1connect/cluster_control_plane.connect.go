// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: porter/v1/cluster_control_plane.proto

package porterv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/porter-dev/api-contracts/generated/go/porter/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ClusterControlPlaneServiceName is the fully-qualified name of the ClusterControlPlaneService
	// service.
	ClusterControlPlaneServiceName = "porter.v1.ClusterControlPlaneService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ClusterControlPlaneServiceQuotaPreflightCheckProcedure is the fully-qualified name of the
	// ClusterControlPlaneService's QuotaPreflightCheck RPC.
	ClusterControlPlaneServiceQuotaPreflightCheckProcedure = "/porter.v1.ClusterControlPlaneService/QuotaPreflightCheck"
	// ClusterControlPlaneServiceCreateAssumeRoleChainProcedure is the fully-qualified name of the
	// ClusterControlPlaneService's CreateAssumeRoleChain RPC.
	ClusterControlPlaneServiceCreateAssumeRoleChainProcedure = "/porter.v1.ClusterControlPlaneService/CreateAssumeRoleChain"
	// ClusterControlPlaneServiceAssumeRoleChainTargetsProcedure is the fully-qualified name of the
	// ClusterControlPlaneService's AssumeRoleChainTargets RPC.
	ClusterControlPlaneServiceAssumeRoleChainTargetsProcedure = "/porter.v1.ClusterControlPlaneService/AssumeRoleChainTargets"
	// ClusterControlPlaneServiceCreateAzureConnectionProcedure is the fully-qualified name of the
	// ClusterControlPlaneService's CreateAzureConnection RPC.
	ClusterControlPlaneServiceCreateAzureConnectionProcedure = "/porter.v1.ClusterControlPlaneService/CreateAzureConnection"
	// ClusterControlPlaneServiceCertificateAuthorityDataProcedure is the fully-qualified name of the
	// ClusterControlPlaneService's CertificateAuthorityData RPC.
	ClusterControlPlaneServiceCertificateAuthorityDataProcedure = "/porter.v1.ClusterControlPlaneService/CertificateAuthorityData"
	// ClusterControlPlaneServiceEKSBearerTokenProcedure is the fully-qualified name of the
	// ClusterControlPlaneService's EKSBearerToken RPC.
	ClusterControlPlaneServiceEKSBearerTokenProcedure = "/porter.v1.ClusterControlPlaneService/EKSBearerToken"
	// ClusterControlPlaneServiceKubeConfigForClusterProcedure is the fully-qualified name of the
	// ClusterControlPlaneService's KubeConfigForCluster RPC.
	ClusterControlPlaneServiceKubeConfigForClusterProcedure = "/porter.v1.ClusterControlPlaneService/KubeConfigForCluster"
	// ClusterControlPlaneServiceUpdateContractProcedure is the fully-qualified name of the
	// ClusterControlPlaneService's UpdateContract RPC.
	ClusterControlPlaneServiceUpdateContractProcedure = "/porter.v1.ClusterControlPlaneService/UpdateContract"
	// ClusterControlPlaneServiceClusterStatusProcedure is the fully-qualified name of the
	// ClusterControlPlaneService's ClusterStatus RPC.
	ClusterControlPlaneServiceClusterStatusProcedure = "/porter.v1.ClusterControlPlaneService/ClusterStatus"
	// ClusterControlPlaneServiceDeleteClusterProcedure is the fully-qualified name of the
	// ClusterControlPlaneService's DeleteCluster RPC.
	ClusterControlPlaneServiceDeleteClusterProcedure = "/porter.v1.ClusterControlPlaneService/DeleteCluster"
	// ClusterControlPlaneServiceECRTokenForRegistryProcedure is the fully-qualified name of the
	// ClusterControlPlaneService's ECRTokenForRegistry RPC.
	ClusterControlPlaneServiceECRTokenForRegistryProcedure = "/porter.v1.ClusterControlPlaneService/ECRTokenForRegistry"
	// ClusterControlPlaneServiceAssumeRoleCredentialsProcedure is the fully-qualified name of the
	// ClusterControlPlaneService's AssumeRoleCredentials RPC.
	ClusterControlPlaneServiceAssumeRoleCredentialsProcedure = "/porter.v1.ClusterControlPlaneService/AssumeRoleCredentials"
	// ClusterControlPlaneServiceTokenForRegistryProcedure is the fully-qualified name of the
	// ClusterControlPlaneService's TokenForRegistry RPC.
	ClusterControlPlaneServiceTokenForRegistryProcedure = "/porter.v1.ClusterControlPlaneService/TokenForRegistry"
	// ClusterControlPlaneServiceDockerConfigFileForRegistryProcedure is the fully-qualified name of the
	// ClusterControlPlaneService's DockerConfigFileForRegistry RPC.
	ClusterControlPlaneServiceDockerConfigFileForRegistryProcedure = "/porter.v1.ClusterControlPlaneService/DockerConfigFileForRegistry"
	// ClusterControlPlaneServiceListRepositoriesForRegistryProcedure is the fully-qualified name of the
	// ClusterControlPlaneService's ListRepositoriesForRegistry RPC.
	ClusterControlPlaneServiceListRepositoriesForRegistryProcedure = "/porter.v1.ClusterControlPlaneService/ListRepositoriesForRegistry"
	// ClusterControlPlaneServiceListImagesForRepositoryProcedure is the fully-qualified name of the
	// ClusterControlPlaneService's ListImagesForRepository RPC.
	ClusterControlPlaneServiceListImagesForRepositoryProcedure = "/porter.v1.ClusterControlPlaneService/ListImagesForRepository"
)

// ClusterControlPlaneServiceClient is a client for the porter.v1.ClusterControlPlaneService
// service.
type ClusterControlPlaneServiceClient interface {
	// QuotaPreflightCheck checks if the target account and region has sufficient resources (EIP addresses and VPCs) to provision a new cluster
	QuotaPreflightCheck(context.Context, *connect_go.Request[v1.QuotaPreflightCheckRequest]) (*connect_go.Response[v1.QuotaPreflightCheckResponse], error)
	// CreateAssumeRoleChain creates a new assume role chain for a given project and checks if the target assumed role has sufficient permissions
	CreateAssumeRoleChain(context.Context, *connect_go.Request[v1.CreateAssumeRoleChainRequest]) (*connect_go.Response[v1.CreateAssumeRoleChainResponse], error)
	// AssumeRoleChainTargets gets the final destination target_arns for a given project
	AssumeRoleChainTargets(context.Context, *connect_go.Request[v1.AssumeRoleChainTargetsRequest]) (*connect_go.Response[v1.AssumeRoleChainTargetsResponse], error)
	// CreateAzureConnection stores the new azure service principal credentials and creates the azure cluster identity
	CreateAzureConnection(context.Context, *connect_go.Request[v1.CreateAzureConnectionRequest]) (*connect_go.Response[v1.CreateAzureConnectionResponse], error)
	// CertificateAuthorityData gets the certificate authority data for a customer cluster
	CertificateAuthorityData(context.Context, *connect_go.Request[v1.CertificateAuthorityDataRequest]) (*connect_go.Response[v1.CertificateAuthorityDataResponse], error)
	// EKSBearerToken gets a bearer token for programatic access to an EKS cluster's kubernetes API
	EKSBearerToken(context.Context, *connect_go.Request[v1.EKSBearerTokenRequest]) (*connect_go.Response[v1.EKSBearerTokenResponse], error)
	// KubeConfigForCluster gets a valid kubeconfig from the management cluster, for a given workload cluster
	KubeConfigForCluster(context.Context, *connect_go.Request[v1.KubeConfigForClusterRequest]) (*connect_go.Response[v1.KubeConfigForClusterResponse], error)
	// UpdateContract takes in a Porter Contract, actioning upon it where necessary
	UpdateContract(context.Context, *connect_go.Request[v1.UpdateContractRequest]) (*connect_go.Response[v1.UpdateContractResponse], error)
	// ClusterStatus returns the status of a given workload cluster
	ClusterStatus(context.Context, *connect_go.Request[v1.ClusterStatusRequest]) (*connect_go.Response[v1.ClusterStatusResponse], error)
	// DeleteCluster uninstalls system components from a given workload cluster before deleting it.
	// This should ultimately be wrapped into UpdateContract
	DeleteCluster(context.Context, *connect_go.Request[v1.DeleteClusterRequest]) (*connect_go.Response[v1.DeleteClusterResponse], error)
	// ECRTokenForRegistry returns a docker-compatible token for accessing a given ECR registry
	ECRTokenForRegistry(context.Context, *connect_go.Request[v1.ECRTokenForRegistryRequest]) (*connect_go.Response[v1.ECRTokenForRegistryResponse], error)
	// AssumeRoleCredentials should be used vary sparingly, and ONLY for replacing AWS Integrations which have no workaround on the Porter API.
	// This endpoint returns temporary AWS credentials for a given AWS Account ID, and should not be expanded further to allow specifc role selection without being tied to a project and cluster
	AssumeRoleCredentials(context.Context, *connect_go.Request[v1.AssumeRoleCredentialsRequest]) (*connect_go.Response[v1.AssumeRoleCredentialsResponse], error)
	// TokenForRegistry returns a docker-compatible token for accessing a given registry
	TokenForRegistry(context.Context, *connect_go.Request[v1.TokenForRegistryRequest]) (*connect_go.Response[v1.TokenForRegistryResponse], error)
	// DockerConfigFileForRegistry returns a docker-compatible token for accessing a given registry
	DockerConfigFileForRegistry(context.Context, *connect_go.Request[v1.DockerConfigFileForRegistryRequest]) (*connect_go.Response[v1.DockerConfigFileForRegistryResponse], error)
	// ListRepositoriesForRegistry lists the repositories for a given registry, provided it is in the scope of the project id
	ListRepositoriesForRegistry(context.Context, *connect_go.Request[v1.ListRepositoriesForRegistryRequest]) (*connect_go.Response[v1.ListRepositoriesForRegistryResponse], error)
	// ListImagesForRepository lists the repositories for a given registry, provided it is in the scope of the project id
	ListImagesForRepository(context.Context, *connect_go.Request[v1.ListImagesForRepositoryRequest]) (*connect_go.Response[v1.ListImagesForRepositoryResponse], error)
}

// NewClusterControlPlaneServiceClient constructs a client for the
// porter.v1.ClusterControlPlaneService service. By default, it uses the Connect protocol with the
// binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests. To use the
// gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewClusterControlPlaneServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ClusterControlPlaneServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &clusterControlPlaneServiceClient{
		quotaPreflightCheck: connect_go.NewClient[v1.QuotaPreflightCheckRequest, v1.QuotaPreflightCheckResponse](
			httpClient,
			baseURL+ClusterControlPlaneServiceQuotaPreflightCheckProcedure,
			opts...,
		),
		createAssumeRoleChain: connect_go.NewClient[v1.CreateAssumeRoleChainRequest, v1.CreateAssumeRoleChainResponse](
			httpClient,
			baseURL+ClusterControlPlaneServiceCreateAssumeRoleChainProcedure,
			opts...,
		),
		assumeRoleChainTargets: connect_go.NewClient[v1.AssumeRoleChainTargetsRequest, v1.AssumeRoleChainTargetsResponse](
			httpClient,
			baseURL+ClusterControlPlaneServiceAssumeRoleChainTargetsProcedure,
			opts...,
		),
		createAzureConnection: connect_go.NewClient[v1.CreateAzureConnectionRequest, v1.CreateAzureConnectionResponse](
			httpClient,
			baseURL+ClusterControlPlaneServiceCreateAzureConnectionProcedure,
			opts...,
		),
		certificateAuthorityData: connect_go.NewClient[v1.CertificateAuthorityDataRequest, v1.CertificateAuthorityDataResponse](
			httpClient,
			baseURL+ClusterControlPlaneServiceCertificateAuthorityDataProcedure,
			opts...,
		),
		eKSBearerToken: connect_go.NewClient[v1.EKSBearerTokenRequest, v1.EKSBearerTokenResponse](
			httpClient,
			baseURL+ClusterControlPlaneServiceEKSBearerTokenProcedure,
			opts...,
		),
		kubeConfigForCluster: connect_go.NewClient[v1.KubeConfigForClusterRequest, v1.KubeConfigForClusterResponse](
			httpClient,
			baseURL+ClusterControlPlaneServiceKubeConfigForClusterProcedure,
			opts...,
		),
		updateContract: connect_go.NewClient[v1.UpdateContractRequest, v1.UpdateContractResponse](
			httpClient,
			baseURL+ClusterControlPlaneServiceUpdateContractProcedure,
			opts...,
		),
		clusterStatus: connect_go.NewClient[v1.ClusterStatusRequest, v1.ClusterStatusResponse](
			httpClient,
			baseURL+ClusterControlPlaneServiceClusterStatusProcedure,
			opts...,
		),
		deleteCluster: connect_go.NewClient[v1.DeleteClusterRequest, v1.DeleteClusterResponse](
			httpClient,
			baseURL+ClusterControlPlaneServiceDeleteClusterProcedure,
			opts...,
		),
		eCRTokenForRegistry: connect_go.NewClient[v1.ECRTokenForRegistryRequest, v1.ECRTokenForRegistryResponse](
			httpClient,
			baseURL+ClusterControlPlaneServiceECRTokenForRegistryProcedure,
			opts...,
		),
		assumeRoleCredentials: connect_go.NewClient[v1.AssumeRoleCredentialsRequest, v1.AssumeRoleCredentialsResponse](
			httpClient,
			baseURL+ClusterControlPlaneServiceAssumeRoleCredentialsProcedure,
			opts...,
		),
		tokenForRegistry: connect_go.NewClient[v1.TokenForRegistryRequest, v1.TokenForRegistryResponse](
			httpClient,
			baseURL+ClusterControlPlaneServiceTokenForRegistryProcedure,
			opts...,
		),
		dockerConfigFileForRegistry: connect_go.NewClient[v1.DockerConfigFileForRegistryRequest, v1.DockerConfigFileForRegistryResponse](
			httpClient,
			baseURL+ClusterControlPlaneServiceDockerConfigFileForRegistryProcedure,
			opts...,
		),
		listRepositoriesForRegistry: connect_go.NewClient[v1.ListRepositoriesForRegistryRequest, v1.ListRepositoriesForRegistryResponse](
			httpClient,
			baseURL+ClusterControlPlaneServiceListRepositoriesForRegistryProcedure,
			opts...,
		),
		listImagesForRepository: connect_go.NewClient[v1.ListImagesForRepositoryRequest, v1.ListImagesForRepositoryResponse](
			httpClient,
			baseURL+ClusterControlPlaneServiceListImagesForRepositoryProcedure,
			opts...,
		),
	}
}

// clusterControlPlaneServiceClient implements ClusterControlPlaneServiceClient.
type clusterControlPlaneServiceClient struct {
	quotaPreflightCheck         *connect_go.Client[v1.QuotaPreflightCheckRequest, v1.QuotaPreflightCheckResponse]
	createAssumeRoleChain       *connect_go.Client[v1.CreateAssumeRoleChainRequest, v1.CreateAssumeRoleChainResponse]
	assumeRoleChainTargets      *connect_go.Client[v1.AssumeRoleChainTargetsRequest, v1.AssumeRoleChainTargetsResponse]
	createAzureConnection       *connect_go.Client[v1.CreateAzureConnectionRequest, v1.CreateAzureConnectionResponse]
	certificateAuthorityData    *connect_go.Client[v1.CertificateAuthorityDataRequest, v1.CertificateAuthorityDataResponse]
	eKSBearerToken              *connect_go.Client[v1.EKSBearerTokenRequest, v1.EKSBearerTokenResponse]
	kubeConfigForCluster        *connect_go.Client[v1.KubeConfigForClusterRequest, v1.KubeConfigForClusterResponse]
	updateContract              *connect_go.Client[v1.UpdateContractRequest, v1.UpdateContractResponse]
	clusterStatus               *connect_go.Client[v1.ClusterStatusRequest, v1.ClusterStatusResponse]
	deleteCluster               *connect_go.Client[v1.DeleteClusterRequest, v1.DeleteClusterResponse]
	eCRTokenForRegistry         *connect_go.Client[v1.ECRTokenForRegistryRequest, v1.ECRTokenForRegistryResponse]
	assumeRoleCredentials       *connect_go.Client[v1.AssumeRoleCredentialsRequest, v1.AssumeRoleCredentialsResponse]
	tokenForRegistry            *connect_go.Client[v1.TokenForRegistryRequest, v1.TokenForRegistryResponse]
	dockerConfigFileForRegistry *connect_go.Client[v1.DockerConfigFileForRegistryRequest, v1.DockerConfigFileForRegistryResponse]
	listRepositoriesForRegistry *connect_go.Client[v1.ListRepositoriesForRegistryRequest, v1.ListRepositoriesForRegistryResponse]
	listImagesForRepository     *connect_go.Client[v1.ListImagesForRepositoryRequest, v1.ListImagesForRepositoryResponse]
}

// QuotaPreflightCheck calls porter.v1.ClusterControlPlaneService.QuotaPreflightCheck.
func (c *clusterControlPlaneServiceClient) QuotaPreflightCheck(ctx context.Context, req *connect_go.Request[v1.QuotaPreflightCheckRequest]) (*connect_go.Response[v1.QuotaPreflightCheckResponse], error) {
	return c.quotaPreflightCheck.CallUnary(ctx, req)
}

// CreateAssumeRoleChain calls porter.v1.ClusterControlPlaneService.CreateAssumeRoleChain.
func (c *clusterControlPlaneServiceClient) CreateAssumeRoleChain(ctx context.Context, req *connect_go.Request[v1.CreateAssumeRoleChainRequest]) (*connect_go.Response[v1.CreateAssumeRoleChainResponse], error) {
	return c.createAssumeRoleChain.CallUnary(ctx, req)
}

// AssumeRoleChainTargets calls porter.v1.ClusterControlPlaneService.AssumeRoleChainTargets.
func (c *clusterControlPlaneServiceClient) AssumeRoleChainTargets(ctx context.Context, req *connect_go.Request[v1.AssumeRoleChainTargetsRequest]) (*connect_go.Response[v1.AssumeRoleChainTargetsResponse], error) {
	return c.assumeRoleChainTargets.CallUnary(ctx, req)
}

// CreateAzureConnection calls porter.v1.ClusterControlPlaneService.CreateAzureConnection.
func (c *clusterControlPlaneServiceClient) CreateAzureConnection(ctx context.Context, req *connect_go.Request[v1.CreateAzureConnectionRequest]) (*connect_go.Response[v1.CreateAzureConnectionResponse], error) {
	return c.createAzureConnection.CallUnary(ctx, req)
}

// CertificateAuthorityData calls porter.v1.ClusterControlPlaneService.CertificateAuthorityData.
func (c *clusterControlPlaneServiceClient) CertificateAuthorityData(ctx context.Context, req *connect_go.Request[v1.CertificateAuthorityDataRequest]) (*connect_go.Response[v1.CertificateAuthorityDataResponse], error) {
	return c.certificateAuthorityData.CallUnary(ctx, req)
}

// EKSBearerToken calls porter.v1.ClusterControlPlaneService.EKSBearerToken.
func (c *clusterControlPlaneServiceClient) EKSBearerToken(ctx context.Context, req *connect_go.Request[v1.EKSBearerTokenRequest]) (*connect_go.Response[v1.EKSBearerTokenResponse], error) {
	return c.eKSBearerToken.CallUnary(ctx, req)
}

// KubeConfigForCluster calls porter.v1.ClusterControlPlaneService.KubeConfigForCluster.
func (c *clusterControlPlaneServiceClient) KubeConfigForCluster(ctx context.Context, req *connect_go.Request[v1.KubeConfigForClusterRequest]) (*connect_go.Response[v1.KubeConfigForClusterResponse], error) {
	return c.kubeConfigForCluster.CallUnary(ctx, req)
}

// UpdateContract calls porter.v1.ClusterControlPlaneService.UpdateContract.
func (c *clusterControlPlaneServiceClient) UpdateContract(ctx context.Context, req *connect_go.Request[v1.UpdateContractRequest]) (*connect_go.Response[v1.UpdateContractResponse], error) {
	return c.updateContract.CallUnary(ctx, req)
}

// ClusterStatus calls porter.v1.ClusterControlPlaneService.ClusterStatus.
func (c *clusterControlPlaneServiceClient) ClusterStatus(ctx context.Context, req *connect_go.Request[v1.ClusterStatusRequest]) (*connect_go.Response[v1.ClusterStatusResponse], error) {
	return c.clusterStatus.CallUnary(ctx, req)
}

// DeleteCluster calls porter.v1.ClusterControlPlaneService.DeleteCluster.
func (c *clusterControlPlaneServiceClient) DeleteCluster(ctx context.Context, req *connect_go.Request[v1.DeleteClusterRequest]) (*connect_go.Response[v1.DeleteClusterResponse], error) {
	return c.deleteCluster.CallUnary(ctx, req)
}

// ECRTokenForRegistry calls porter.v1.ClusterControlPlaneService.ECRTokenForRegistry.
func (c *clusterControlPlaneServiceClient) ECRTokenForRegistry(ctx context.Context, req *connect_go.Request[v1.ECRTokenForRegistryRequest]) (*connect_go.Response[v1.ECRTokenForRegistryResponse], error) {
	return c.eCRTokenForRegistry.CallUnary(ctx, req)
}

// AssumeRoleCredentials calls porter.v1.ClusterControlPlaneService.AssumeRoleCredentials.
func (c *clusterControlPlaneServiceClient) AssumeRoleCredentials(ctx context.Context, req *connect_go.Request[v1.AssumeRoleCredentialsRequest]) (*connect_go.Response[v1.AssumeRoleCredentialsResponse], error) {
	return c.assumeRoleCredentials.CallUnary(ctx, req)
}

// TokenForRegistry calls porter.v1.ClusterControlPlaneService.TokenForRegistry.
func (c *clusterControlPlaneServiceClient) TokenForRegistry(ctx context.Context, req *connect_go.Request[v1.TokenForRegistryRequest]) (*connect_go.Response[v1.TokenForRegistryResponse], error) {
	return c.tokenForRegistry.CallUnary(ctx, req)
}

// DockerConfigFileForRegistry calls
// porter.v1.ClusterControlPlaneService.DockerConfigFileForRegistry.
func (c *clusterControlPlaneServiceClient) DockerConfigFileForRegistry(ctx context.Context, req *connect_go.Request[v1.DockerConfigFileForRegistryRequest]) (*connect_go.Response[v1.DockerConfigFileForRegistryResponse], error) {
	return c.dockerConfigFileForRegistry.CallUnary(ctx, req)
}

// ListRepositoriesForRegistry calls
// porter.v1.ClusterControlPlaneService.ListRepositoriesForRegistry.
func (c *clusterControlPlaneServiceClient) ListRepositoriesForRegistry(ctx context.Context, req *connect_go.Request[v1.ListRepositoriesForRegistryRequest]) (*connect_go.Response[v1.ListRepositoriesForRegistryResponse], error) {
	return c.listRepositoriesForRegistry.CallUnary(ctx, req)
}

// ListImagesForRepository calls porter.v1.ClusterControlPlaneService.ListImagesForRepository.
func (c *clusterControlPlaneServiceClient) ListImagesForRepository(ctx context.Context, req *connect_go.Request[v1.ListImagesForRepositoryRequest]) (*connect_go.Response[v1.ListImagesForRepositoryResponse], error) {
	return c.listImagesForRepository.CallUnary(ctx, req)
}

// ClusterControlPlaneServiceHandler is an implementation of the
// porter.v1.ClusterControlPlaneService service.
type ClusterControlPlaneServiceHandler interface {
	// QuotaPreflightCheck checks if the target account and region has sufficient resources (EIP addresses and VPCs) to provision a new cluster
	QuotaPreflightCheck(context.Context, *connect_go.Request[v1.QuotaPreflightCheckRequest]) (*connect_go.Response[v1.QuotaPreflightCheckResponse], error)
	// CreateAssumeRoleChain creates a new assume role chain for a given project and checks if the target assumed role has sufficient permissions
	CreateAssumeRoleChain(context.Context, *connect_go.Request[v1.CreateAssumeRoleChainRequest]) (*connect_go.Response[v1.CreateAssumeRoleChainResponse], error)
	// AssumeRoleChainTargets gets the final destination target_arns for a given project
	AssumeRoleChainTargets(context.Context, *connect_go.Request[v1.AssumeRoleChainTargetsRequest]) (*connect_go.Response[v1.AssumeRoleChainTargetsResponse], error)
	// CreateAzureConnection stores the new azure service principal credentials and creates the azure cluster identity
	CreateAzureConnection(context.Context, *connect_go.Request[v1.CreateAzureConnectionRequest]) (*connect_go.Response[v1.CreateAzureConnectionResponse], error)
	// CertificateAuthorityData gets the certificate authority data for a customer cluster
	CertificateAuthorityData(context.Context, *connect_go.Request[v1.CertificateAuthorityDataRequest]) (*connect_go.Response[v1.CertificateAuthorityDataResponse], error)
	// EKSBearerToken gets a bearer token for programatic access to an EKS cluster's kubernetes API
	EKSBearerToken(context.Context, *connect_go.Request[v1.EKSBearerTokenRequest]) (*connect_go.Response[v1.EKSBearerTokenResponse], error)
	// KubeConfigForCluster gets a valid kubeconfig from the management cluster, for a given workload cluster
	KubeConfigForCluster(context.Context, *connect_go.Request[v1.KubeConfigForClusterRequest]) (*connect_go.Response[v1.KubeConfigForClusterResponse], error)
	// UpdateContract takes in a Porter Contract, actioning upon it where necessary
	UpdateContract(context.Context, *connect_go.Request[v1.UpdateContractRequest]) (*connect_go.Response[v1.UpdateContractResponse], error)
	// ClusterStatus returns the status of a given workload cluster
	ClusterStatus(context.Context, *connect_go.Request[v1.ClusterStatusRequest]) (*connect_go.Response[v1.ClusterStatusResponse], error)
	// DeleteCluster uninstalls system components from a given workload cluster before deleting it.
	// This should ultimately be wrapped into UpdateContract
	DeleteCluster(context.Context, *connect_go.Request[v1.DeleteClusterRequest]) (*connect_go.Response[v1.DeleteClusterResponse], error)
	// ECRTokenForRegistry returns a docker-compatible token for accessing a given ECR registry
	ECRTokenForRegistry(context.Context, *connect_go.Request[v1.ECRTokenForRegistryRequest]) (*connect_go.Response[v1.ECRTokenForRegistryResponse], error)
	// AssumeRoleCredentials should be used vary sparingly, and ONLY for replacing AWS Integrations which have no workaround on the Porter API.
	// This endpoint returns temporary AWS credentials for a given AWS Account ID, and should not be expanded further to allow specifc role selection without being tied to a project and cluster
	AssumeRoleCredentials(context.Context, *connect_go.Request[v1.AssumeRoleCredentialsRequest]) (*connect_go.Response[v1.AssumeRoleCredentialsResponse], error)
	// TokenForRegistry returns a docker-compatible token for accessing a given registry
	TokenForRegistry(context.Context, *connect_go.Request[v1.TokenForRegistryRequest]) (*connect_go.Response[v1.TokenForRegistryResponse], error)
	// DockerConfigFileForRegistry returns a docker-compatible token for accessing a given registry
	DockerConfigFileForRegistry(context.Context, *connect_go.Request[v1.DockerConfigFileForRegistryRequest]) (*connect_go.Response[v1.DockerConfigFileForRegistryResponse], error)
	// ListRepositoriesForRegistry lists the repositories for a given registry, provided it is in the scope of the project id
	ListRepositoriesForRegistry(context.Context, *connect_go.Request[v1.ListRepositoriesForRegistryRequest]) (*connect_go.Response[v1.ListRepositoriesForRegistryResponse], error)
	// ListImagesForRepository lists the repositories for a given registry, provided it is in the scope of the project id
	ListImagesForRepository(context.Context, *connect_go.Request[v1.ListImagesForRepositoryRequest]) (*connect_go.Response[v1.ListImagesForRepositoryResponse], error)
}

// NewClusterControlPlaneServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewClusterControlPlaneServiceHandler(svc ClusterControlPlaneServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(ClusterControlPlaneServiceQuotaPreflightCheckProcedure, connect_go.NewUnaryHandler(
		ClusterControlPlaneServiceQuotaPreflightCheckProcedure,
		svc.QuotaPreflightCheck,
		opts...,
	))
	mux.Handle(ClusterControlPlaneServiceCreateAssumeRoleChainProcedure, connect_go.NewUnaryHandler(
		ClusterControlPlaneServiceCreateAssumeRoleChainProcedure,
		svc.CreateAssumeRoleChain,
		opts...,
	))
	mux.Handle(ClusterControlPlaneServiceAssumeRoleChainTargetsProcedure, connect_go.NewUnaryHandler(
		ClusterControlPlaneServiceAssumeRoleChainTargetsProcedure,
		svc.AssumeRoleChainTargets,
		opts...,
	))
	mux.Handle(ClusterControlPlaneServiceCreateAzureConnectionProcedure, connect_go.NewUnaryHandler(
		ClusterControlPlaneServiceCreateAzureConnectionProcedure,
		svc.CreateAzureConnection,
		opts...,
	))
	mux.Handle(ClusterControlPlaneServiceCertificateAuthorityDataProcedure, connect_go.NewUnaryHandler(
		ClusterControlPlaneServiceCertificateAuthorityDataProcedure,
		svc.CertificateAuthorityData,
		opts...,
	))
	mux.Handle(ClusterControlPlaneServiceEKSBearerTokenProcedure, connect_go.NewUnaryHandler(
		ClusterControlPlaneServiceEKSBearerTokenProcedure,
		svc.EKSBearerToken,
		opts...,
	))
	mux.Handle(ClusterControlPlaneServiceKubeConfigForClusterProcedure, connect_go.NewUnaryHandler(
		ClusterControlPlaneServiceKubeConfigForClusterProcedure,
		svc.KubeConfigForCluster,
		opts...,
	))
	mux.Handle(ClusterControlPlaneServiceUpdateContractProcedure, connect_go.NewUnaryHandler(
		ClusterControlPlaneServiceUpdateContractProcedure,
		svc.UpdateContract,
		opts...,
	))
	mux.Handle(ClusterControlPlaneServiceClusterStatusProcedure, connect_go.NewUnaryHandler(
		ClusterControlPlaneServiceClusterStatusProcedure,
		svc.ClusterStatus,
		opts...,
	))
	mux.Handle(ClusterControlPlaneServiceDeleteClusterProcedure, connect_go.NewUnaryHandler(
		ClusterControlPlaneServiceDeleteClusterProcedure,
		svc.DeleteCluster,
		opts...,
	))
	mux.Handle(ClusterControlPlaneServiceECRTokenForRegistryProcedure, connect_go.NewUnaryHandler(
		ClusterControlPlaneServiceECRTokenForRegistryProcedure,
		svc.ECRTokenForRegistry,
		opts...,
	))
	mux.Handle(ClusterControlPlaneServiceAssumeRoleCredentialsProcedure, connect_go.NewUnaryHandler(
		ClusterControlPlaneServiceAssumeRoleCredentialsProcedure,
		svc.AssumeRoleCredentials,
		opts...,
	))
	mux.Handle(ClusterControlPlaneServiceTokenForRegistryProcedure, connect_go.NewUnaryHandler(
		ClusterControlPlaneServiceTokenForRegistryProcedure,
		svc.TokenForRegistry,
		opts...,
	))
	mux.Handle(ClusterControlPlaneServiceDockerConfigFileForRegistryProcedure, connect_go.NewUnaryHandler(
		ClusterControlPlaneServiceDockerConfigFileForRegistryProcedure,
		svc.DockerConfigFileForRegistry,
		opts...,
	))
	mux.Handle(ClusterControlPlaneServiceListRepositoriesForRegistryProcedure, connect_go.NewUnaryHandler(
		ClusterControlPlaneServiceListRepositoriesForRegistryProcedure,
		svc.ListRepositoriesForRegistry,
		opts...,
	))
	mux.Handle(ClusterControlPlaneServiceListImagesForRepositoryProcedure, connect_go.NewUnaryHandler(
		ClusterControlPlaneServiceListImagesForRepositoryProcedure,
		svc.ListImagesForRepository,
		opts...,
	))
	return "/porter.v1.ClusterControlPlaneService/", mux
}

// UnimplementedClusterControlPlaneServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedClusterControlPlaneServiceHandler struct{}

func (UnimplementedClusterControlPlaneServiceHandler) QuotaPreflightCheck(context.Context, *connect_go.Request[v1.QuotaPreflightCheckRequest]) (*connect_go.Response[v1.QuotaPreflightCheckResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("porter.v1.ClusterControlPlaneService.QuotaPreflightCheck is not implemented"))
}

func (UnimplementedClusterControlPlaneServiceHandler) CreateAssumeRoleChain(context.Context, *connect_go.Request[v1.CreateAssumeRoleChainRequest]) (*connect_go.Response[v1.CreateAssumeRoleChainResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("porter.v1.ClusterControlPlaneService.CreateAssumeRoleChain is not implemented"))
}

func (UnimplementedClusterControlPlaneServiceHandler) AssumeRoleChainTargets(context.Context, *connect_go.Request[v1.AssumeRoleChainTargetsRequest]) (*connect_go.Response[v1.AssumeRoleChainTargetsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("porter.v1.ClusterControlPlaneService.AssumeRoleChainTargets is not implemented"))
}

func (UnimplementedClusterControlPlaneServiceHandler) CreateAzureConnection(context.Context, *connect_go.Request[v1.CreateAzureConnectionRequest]) (*connect_go.Response[v1.CreateAzureConnectionResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("porter.v1.ClusterControlPlaneService.CreateAzureConnection is not implemented"))
}

func (UnimplementedClusterControlPlaneServiceHandler) CertificateAuthorityData(context.Context, *connect_go.Request[v1.CertificateAuthorityDataRequest]) (*connect_go.Response[v1.CertificateAuthorityDataResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("porter.v1.ClusterControlPlaneService.CertificateAuthorityData is not implemented"))
}

func (UnimplementedClusterControlPlaneServiceHandler) EKSBearerToken(context.Context, *connect_go.Request[v1.EKSBearerTokenRequest]) (*connect_go.Response[v1.EKSBearerTokenResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("porter.v1.ClusterControlPlaneService.EKSBearerToken is not implemented"))
}

func (UnimplementedClusterControlPlaneServiceHandler) KubeConfigForCluster(context.Context, *connect_go.Request[v1.KubeConfigForClusterRequest]) (*connect_go.Response[v1.KubeConfigForClusterResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("porter.v1.ClusterControlPlaneService.KubeConfigForCluster is not implemented"))
}

func (UnimplementedClusterControlPlaneServiceHandler) UpdateContract(context.Context, *connect_go.Request[v1.UpdateContractRequest]) (*connect_go.Response[v1.UpdateContractResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("porter.v1.ClusterControlPlaneService.UpdateContract is not implemented"))
}

func (UnimplementedClusterControlPlaneServiceHandler) ClusterStatus(context.Context, *connect_go.Request[v1.ClusterStatusRequest]) (*connect_go.Response[v1.ClusterStatusResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("porter.v1.ClusterControlPlaneService.ClusterStatus is not implemented"))
}

func (UnimplementedClusterControlPlaneServiceHandler) DeleteCluster(context.Context, *connect_go.Request[v1.DeleteClusterRequest]) (*connect_go.Response[v1.DeleteClusterResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("porter.v1.ClusterControlPlaneService.DeleteCluster is not implemented"))
}

func (UnimplementedClusterControlPlaneServiceHandler) ECRTokenForRegistry(context.Context, *connect_go.Request[v1.ECRTokenForRegistryRequest]) (*connect_go.Response[v1.ECRTokenForRegistryResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("porter.v1.ClusterControlPlaneService.ECRTokenForRegistry is not implemented"))
}

func (UnimplementedClusterControlPlaneServiceHandler) AssumeRoleCredentials(context.Context, *connect_go.Request[v1.AssumeRoleCredentialsRequest]) (*connect_go.Response[v1.AssumeRoleCredentialsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("porter.v1.ClusterControlPlaneService.AssumeRoleCredentials is not implemented"))
}

func (UnimplementedClusterControlPlaneServiceHandler) TokenForRegistry(context.Context, *connect_go.Request[v1.TokenForRegistryRequest]) (*connect_go.Response[v1.TokenForRegistryResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("porter.v1.ClusterControlPlaneService.TokenForRegistry is not implemented"))
}

func (UnimplementedClusterControlPlaneServiceHandler) DockerConfigFileForRegistry(context.Context, *connect_go.Request[v1.DockerConfigFileForRegistryRequest]) (*connect_go.Response[v1.DockerConfigFileForRegistryResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("porter.v1.ClusterControlPlaneService.DockerConfigFileForRegistry is not implemented"))
}

func (UnimplementedClusterControlPlaneServiceHandler) ListRepositoriesForRegistry(context.Context, *connect_go.Request[v1.ListRepositoriesForRegistryRequest]) (*connect_go.Response[v1.ListRepositoriesForRegistryResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("porter.v1.ClusterControlPlaneService.ListRepositoriesForRegistry is not implemented"))
}

func (UnimplementedClusterControlPlaneServiceHandler) ListImagesForRepository(context.Context, *connect_go.Request[v1.ListImagesForRepositoryRequest]) (*connect_go.Response[v1.ListImagesForRepositoryResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("porter.v1.ClusterControlPlaneService.ListImagesForRepository is not implemented"))
}
