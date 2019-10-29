module github.com/pulumi/pulumi-okta

go 1.12

require (
	github.com/Azure/go-autorest/autorest/azure/auth v0.4.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.2.0 // indirect
	github.com/articulate/terraform-provider-okta v0.0.0-20191023132502-355ea4d61f43
	github.com/docker/docker v1.13.1 // indirect
	github.com/gliderlabs/ssh v0.1.3 // indirect
	github.com/google/btree v1.0.0 // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.2.0
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.2.0
	github.com/pulumi/pulumi-terraform-bridge v1.2.0
	github.com/stretchr/testify v1.4.0
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
	github.com/okta/okta-sdk-golang => github.com/articulate/okta-sdk-golang v1.1.1
)
