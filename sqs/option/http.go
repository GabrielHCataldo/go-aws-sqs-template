package option

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/smithy-go/logging"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type HttpClient struct {
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this sqsClient. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error
	// The optional application specific identifier appended to the User-Agent header.
	AppID string
	// This endpoint will be given as input to an EndpointResolverV2. It is used for
	// providing a custom base endpoint that is subject to modifications by the
	// processing EndpointResolverV2.
	BaseEndpoint *string
	// Configures the events that will be sent to the configured logger.
	ClientLogMode aws.ClientLogMode
	// The credentials object to use when signing requests.
	Credentials aws.CredentialsProvider
	// The configuration DefaultsMode that the SDK should use when constructing the
	// clients initial default settings.
	DefaultsMode aws.DefaultsMode
	// Allows you to disable the client's validation of response message checksums.
	// Enabled by default. Used by SendMessage, SendMessageBatch, and ReceiveMessage.
	DisableMessageChecksumValidation bool
	// The endpoint options to be used when attempting to resolve an endpoint.
	EndpointOptions sqs.EndpointResolverOptions
	// Resolves the endpoint used for a particular service operation. This should be
	// used over the deprecated EndpointResolver.
	EndpointResolverV2 sqs.EndpointResolverV2
	// Signature Version 4 (SigV4) Signer
	HTTPSignerV4 sqs.HTTPSignerV4
	// The logger writer interface to write logging messages to.
	Logger logging.Logger
	// The region to send requests to. (Required)
	Region string
	// RetryMaxAttempts specifies the maximum number attempts an API client will call
	// an operation that fails with a retryable error. A value of 0 is ignored, and
	// will not be used to configure the API client created default retryer, or modify
	// per operation call's retry max attempts. If specified in an operation call's
	// functional options with a value that is different from the constructed client's
	// Options, the Client's Retryer will be wrapped to use the operation's specific
	// RetryMaxAttempts value.
	RetryMaxAttempts int
	// RetryMode specifies the retry mode the API client will be created with, if
	// Retryer option is not also specified. When creating a new API Clients this
	// member will only be used if the Retryer Options member is nil. This value will
	// be ignored if Retryer is not nil. Currently, does not support per operation call
	// overrides, may in the future.
	RetryMode aws.RetryMode
	// Retryer guides how HTTP requests should be retried in case of recoverable
	// failures. When nil the API client will use a default retryer. The kind of
	// default retry created by the API client can be changed with the RetryMode
	// option.
	Retryer aws.Retryer
	// The RuntimeEnvironment configuration, only populated if the DefaultsMode is set
	// to DefaultsModeAuto and is initialized using config.LoadDefaultConfig . You
	// should not populate this structure programmatically, or rely on the values here
	// within your applications.
	RuntimeEnvironment aws.RuntimeEnvironment
	// The initial DefaultsMode used when the client options were constructed. If the
	// DefaultsMode was set to aws.DefaultsModeAuto this will store what the resolved
	// value was at that point in time. Currently, does not support per operation call
	// overrides, may in the future.
	resolvedDefaultsMode aws.DefaultsMode
	// The HTTP client to invoke API calls with. Defaults to client's default HTTP
	// implementation if nil.
	HTTPClient sqs.HTTPClient
	// The auth scheme resolver which determines how to authenticate for each
	// operation.
	AuthSchemeResolver sqs.AuthSchemeResolver
	// The list of auth schemes supported by the sqsClient.
	AuthSchemes []smithyhttp.AuthScheme
}

func FuncByHttpClient(opt *HttpClient) func(options *sqs.Options) {
	return func(options *sqs.Options) {
		if opt == nil {
			return
		}
		options = &sqs.Options{
			APIOptions:                       opt.APIOptions,
			AppID:                            opt.AppID,
			BaseEndpoint:                     opt.BaseEndpoint,
			ClientLogMode:                    opt.ClientLogMode,
			Credentials:                      opt.Credentials,
			DefaultsMode:                     opt.DefaultsMode,
			DisableMessageChecksumValidation: opt.DisableMessageChecksumValidation,
			EndpointOptions:                  opt.EndpointOptions,
			EndpointResolverV2:               opt.EndpointResolverV2,
			HTTPSignerV4:                     opt.HTTPSignerV4,
			Logger:                           opt.Logger,
			Region:                           opt.Region,
			RetryMaxAttempts:                 opt.RetryMaxAttempts,
			RetryMode:                        opt.RetryMode,
			Retryer:                          opt.Retryer,
			RuntimeEnvironment:               opt.RuntimeEnvironment,
			HTTPClient:                       opt.HTTPClient,
			AuthSchemeResolver:               opt.AuthSchemeResolver,
			AuthSchemes:                      opt.AuthSchemes,
		}
	}
}
