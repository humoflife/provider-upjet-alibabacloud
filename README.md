# provider-upjet-alibabacloud

## Build

Run the following to build the provider locally.

```
make submodules
make generate
```

## Authentication

Provider credentials are configured with a `ProviderConfig`. Static credentials
and STS session credentials are read from a Kubernetes `Secret`; AssumeRole and
AssumeRoleWithOIDC options are configured as structured fields on the
`ProviderConfig` spec.

Do not put `assume_role` or `assume_role_with_oidc` blocks in the credentials
secret. The credentials secret should contain only credential material such as
`access_key`, `secret_key`, `security_token`, and `region`.

### Static Credentials

The default static AK/SK example uses `credentials.source: Secret`:

```bash
kubectl apply -f examples/providerconfig/v1beta1/secret.yaml.tmpl
kubectl apply -f examples/providerconfig/v1beta1/providerconfig.yaml
```

The credentials secret uses JSON in the `credentials` key:

```json
{
  "access_key": "...",
  "secret_key": "...",
  "region": "cn-hangzhou"
}
```

`region_id` is still accepted as a compatibility fallback, but new examples
should use `region`.

### STS Session Credentials

STS session credentials are supported by adding `security_token` to the secret:

```bash
kubectl apply -f examples/providerconfig/v1beta1/secret-sts-token.yaml.tmpl
kubectl apply -f examples/providerconfig/v1beta1/providerconfig-sts-token.yaml
```

### AssumeRole

For AssumeRole, use a base credential secret for the caller identity and put
the role parameters in `spec.assumeRole`:

```bash
kubectl apply -f examples/providerconfig/v1beta1/secret.yaml.tmpl
kubectl apply -f examples/providerconfig/v1beta1/providerconfig-assume-role.yaml
```

```yaml
apiVersion: alibabacloud.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: assume-role
spec:
  credentials:
    source: Secret
    secretRef:
      name: example-creds
      namespace: crossplane-system
      key: credentials
  assumeRole:
    roleARN: acs:ram::<account-id>:role/<role-name>
    sessionName: crossplane-assume-role
    sessionExpiration: 3600
```

### AssumeRoleWithOIDC

For AssumeRoleWithOIDC, configure the role, OIDC provider, and token source in
`spec.assumeRoleWithOIDC`:

```bash
kubectl apply -f examples/providerconfig/v1beta1/providerconfig-assume-role-with-oidc.yaml
```

```yaml
apiVersion: alibabacloud.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: assume-role-with-oidc
spec:
  credentials:
    source: None
  assumeRoleWithOIDC:
    roleARN: acs:ram::<account-id>:role/<role-name>
    oidcProviderARN: acs:ram::<account-id>:oidc-provider/<provider-name>
    oidcTokenFile: /var/run/secrets/ack.alibabacloud.com/rrsa-tokens/token
    roleSessionName: crossplane-oidc
    sessionExpiration: 3600
```

When running in ACK with RRSA enabled, use the token file path injected by RRSA.
For local testing, create a projected ServiceAccount token and point
`oidcTokenFile` at that local file.

After applying a `ProviderConfig`, managed resources can reference it with
`spec.providerConfigRef.name`. If omitted, resources use the `default`
`ProviderConfig`.

## Test

Add an environment variable to set the credentials for the target Alibaba
account for the tests as follows and then run `make e2e`.

```
export UPTEST_CLOUD_CREDENTIALS='{
    "access_key": "...",
    "secret_key": "...",
    "region": "us-west-1"
}'
```

## Submit PR

- `make reviewable` before submitting a new PR
- git commit -s -m "sign every commit"

## Release New Provider Version

### Determine Version

Identify the version to be released by increasing the minor version by one. For example, if the provider's latest version is v1.1.0, the new version will be v1.2.0.

According to the semantic versioning specification, a version number is represented as MAJOR.MINOR.PATCH. For 1.2.0 : MAJOR=1, MINOR=2, PATCH=0 

### Create Release Branch

From the GitHub UI, create a new branch from the main branch with the name release-<major>-<minor><patch>.

To cut the release v1.2.0, we will name our branch release-1.2.0.

### Build Release Candidate

GitHub should automatically trigger a `CI` workflow run on the newly created branch and produce a package.  You can check it from the GitHub UI by clicking `Actions => CI`.

If it does not, you can manually run the GitHub workflow named CI on the release branch to produce a package.

### Cut The Release

Tag the release branch with the version by running the GitHub workflow named `Tag` on the release branch.

### Publish The Providers

Build and push the family packages using the `Publish Provider Packages` Github Actions workflow. To do this, you need to provide the values of the following parameters:

- subpackages (to be built individually, e.g. config ram): config ack ackone alb alidns cdn cloudmonitorservice ecs fcv3 kms messageservice oss polardb privatelink quotas ram slb tair vpc
- size (Number of smaller provider packages to build and push with each build job): 30
- concurrency (Number of parallel package builds within each build job): 1
- version (Version string to use while publishing the packages,e.g. v1.2.0): v1.2.0
- go-version (Go version to use if building needs to be done): 1.24

Your release build will be published once the `Publish Provider Packages` job if
releasing a family of providers succeeds. Check their availability in the
Upbound marketplace [here](https://marketplace.upbound.io/providers/crossplane-contrib/provider-family-alibabacloud).

### Add Release Notes

Go [here](https://github.com/crossplane-contrib/provider-upjet-alibabacloud) and
click on releases on the left side. 

On the releases page, click on "Draft New Release".
- As target select your release branch that you created above
- Select the corresponding release tag
- Use your version as Release Title, e.g. v1.2.0
- Click "Generate release notes"
- Click "Publish release"
