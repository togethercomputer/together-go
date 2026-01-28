# Beta

## Jig

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#Deployment">Deployment</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#DeploymentLogs">DeploymentLogs</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigListResponse">BetaJigListResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigDestroyResponse">BetaJigDestroyResponse</a>

Methods:

- <code title="get /deployments/{id}">client.Beta.Jig.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#Deployment">Deployment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /deployments/{id}">client.Beta.Jig.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigUpdateParams">BetaJigUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#Deployment">Deployment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /deployments">client.Beta.Jig.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigListResponse">BetaJigListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /deployments">client.Beta.Jig.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigService.Deploy">Deploy</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigDeployParams">BetaJigDeployParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#Deployment">Deployment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /deployments/{id}">client.Beta.Jig.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigService.Destroy">Destroy</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigDestroyResponse">BetaJigDestroyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /deployments/{id}/logs">client.Beta.Jig.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigService.GetLogs">GetLogs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigGetLogsParams">BetaJigGetLogsParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#DeploymentLogs">DeploymentLogs</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Volumes

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#Volume">Volume</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigVolumeListResponse">BetaJigVolumeListResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigVolumeDeleteResponse">BetaJigVolumeDeleteResponse</a>

Methods:

- <code title="post /deployments/storage/volumes">client.Beta.Jig.Volumes.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigVolumeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigVolumeNewParams">BetaJigVolumeNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#Volume">Volume</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /deployments/storage/volumes/{id}">client.Beta.Jig.Volumes.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigVolumeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#Volume">Volume</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /deployments/storage/volumes/{id}">client.Beta.Jig.Volumes.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigVolumeService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigVolumeUpdateParams">BetaJigVolumeUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#Volume">Volume</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /deployments/storage/volumes">client.Beta.Jig.Volumes.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigVolumeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigVolumeListResponse">BetaJigVolumeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /deployments/storage/volumes/{id}">client.Beta.Jig.Volumes.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigVolumeService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigVolumeDeleteResponse">BetaJigVolumeDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Secrets

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#Secret">Secret</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigSecretListResponse">BetaJigSecretListResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigSecretDeleteResponse">BetaJigSecretDeleteResponse</a>

Methods:

- <code title="post /deployments/secrets">client.Beta.Jig.Secrets.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigSecretService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigSecretNewParams">BetaJigSecretNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#Secret">Secret</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /deployments/secrets/{id}">client.Beta.Jig.Secrets.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigSecretService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#Secret">Secret</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /deployments/secrets/{id}">client.Beta.Jig.Secrets.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigSecretService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigSecretUpdateParams">BetaJigSecretUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#Secret">Secret</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /deployments/secrets">client.Beta.Jig.Secrets.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigSecretService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigSecretListResponse">BetaJigSecretListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /deployments/secrets/{id}">client.Beta.Jig.Secrets.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigSecretService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaJigSecretDeleteResponse">BetaJigSecretDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Clusters

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#Cluster">Cluster</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaClusterListResponse">BetaClusterListResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaClusterDeleteResponse">BetaClusterDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaClusterListRegionsResponse">BetaClusterListRegionsResponse</a>

Methods:

- <code title="post /clusters">client.Beta.Clusters.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaClusterService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaClusterNewParams">BetaClusterNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#Cluster">Cluster</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /clusters/{cluster_id}">client.Beta.Clusters.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaClusterService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, clusterID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#Cluster">Cluster</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /clusters/{cluster_id}">client.Beta.Clusters.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaClusterService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, clusterID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaClusterUpdateParams">BetaClusterUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#Cluster">Cluster</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /clusters">client.Beta.Clusters.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaClusterService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaClusterListResponse">BetaClusterListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /clusters/{cluster_id}">client.Beta.Clusters.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaClusterService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, clusterID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaClusterDeleteResponse">BetaClusterDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /clusters/regions">client.Beta.Clusters.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaClusterService.ListRegions">ListRegions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaClusterListRegionsResponse">BetaClusterListRegionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Storage

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ClusterStorage">ClusterStorage</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaClusterStorageListResponse">BetaClusterStorageListResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaClusterStorageDeleteResponse">BetaClusterStorageDeleteResponse</a>

Methods:

- <code title="post /clusters/storages">client.Beta.Clusters.Storage.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaClusterStorageService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaClusterStorageNewParams">BetaClusterStorageNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ClusterStorage">ClusterStorage</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /clusters/storages/{volume_id}">client.Beta.Clusters.Storage.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaClusterStorageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ClusterStorage">ClusterStorage</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /clusters/storages">client.Beta.Clusters.Storage.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaClusterStorageService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaClusterStorageUpdateParams">BetaClusterStorageUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ClusterStorage">ClusterStorage</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /clusters/storages">client.Beta.Clusters.Storage.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaClusterStorageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaClusterStorageListResponse">BetaClusterStorageListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /clusters/storages/{volume_id}">client.Beta.Clusters.Storage.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaClusterStorageService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BetaClusterStorageDeleteResponse">BetaClusterStorageDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Chat

## Completions

Params Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ChatCompletionStructuredMessageImageURLParam">ChatCompletionStructuredMessageImageURLParam</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ChatCompletionStructuredMessageTextParam">ChatCompletionStructuredMessageTextParam</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ChatCompletionStructuredMessageVideoURLParam">ChatCompletionStructuredMessageVideoURLParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ChatCompletion">ChatCompletion</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ChatCompletionChunk">ChatCompletionChunk</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ChatCompletionUsage">ChatCompletionUsage</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ChatCompletionWarning">ChatCompletionWarning</a>

Methods:

- <code title="post /chat/completions">client.Chat.Completions.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ChatCompletionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ChatCompletionNewParams">ChatCompletionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ChatCompletion">ChatCompletion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Completions

Params Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ToolChoiceParam">ToolChoiceParam</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ToolsParam">ToolsParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#Completion">Completion</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#CompletionChunk">CompletionChunk</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#LogProbs">LogProbs</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ToolChoice">ToolChoice</a>

Methods:

- <code title="post /completions">client.Completions.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#CompletionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#CompletionNewParams">CompletionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#Completion">Completion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Embeddings

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#Embedding">Embedding</a>

Methods:

- <code title="post /embeddings">client.Embeddings.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EmbeddingService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EmbeddingNewParams">EmbeddingNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#Embedding">Embedding</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Files

Params Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FilePurpose">FilePurpose</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileType">FileType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileList">FileList</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FilePurpose">FilePurpose</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileResponse">FileResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileType">FileType</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileDeleteResponse">FileDeleteResponse</a>

Methods:

- <code title="get /files/{id}">client.Files.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileResponse">FileResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /files">client.Files.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileList">FileList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /files/{id}">client.Files.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileDeleteResponse">FileDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /files/{id}/content">client.Files.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileService.Content">Content</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /files/upload">client.Files.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileService.Upload">Upload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileUploadParams">FileUploadParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileResponse">FileResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# FineTuning

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FinetuneEvent">FinetuneEvent</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FinetuneEventType">FinetuneEventType</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FinetuneResponse">FinetuneResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningNewResponse">FineTuningNewResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningListResponse">FineTuningListResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningDeleteResponse">FineTuningDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningCancelResponse">FineTuningCancelResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningEstimatePriceResponse">FineTuningEstimatePriceResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningListCheckpointsResponse">FineTuningListCheckpointsResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningListEventsResponse">FineTuningListEventsResponse</a>

Methods:

- <code title="post /fine-tunes">client.FineTuning.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningNewParams">FineTuningNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningNewResponse">FineTuningNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /fine-tunes/{id}">client.FineTuning.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FinetuneResponse">FinetuneResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /fine-tunes">client.FineTuning.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningListResponse">FineTuningListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /fine-tunes/{id}">client.FineTuning.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningDeleteParams">FineTuningDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningDeleteResponse">FineTuningDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /fine-tunes/{id}/cancel">client.FineTuning.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningCancelResponse">FineTuningCancelResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /finetune/download">client.FineTuning.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningService.Content">Content</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningContentParams">FineTuningContentParams</a>) (\*http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /fine-tunes/estimate-price">client.FineTuning.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningService.EstimatePrice">EstimatePrice</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningEstimatePriceParams">FineTuningEstimatePriceParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningEstimatePriceResponse">FineTuningEstimatePriceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /fine-tunes/{id}/checkpoints">client.FineTuning.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningService.ListCheckpoints">ListCheckpoints</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningListCheckpointsResponse">FineTuningListCheckpointsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /fine-tunes/{id}/events">client.FineTuning.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningService.ListEvents">ListEvents</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuningListEventsResponse">FineTuningListEventsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# CodeInterpreter

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ExecuteResponseUnion">ExecuteResponseUnion</a>

Methods:

- <code title="post /tci/execute">client.CodeInterpreter.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#CodeInterpreterService.Execute">Execute</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#CodeInterpreterExecuteParams">CodeInterpreterExecuteParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ExecuteResponseUnion">ExecuteResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Sessions

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#SessionListResponse">SessionListResponse</a>

Methods:

- <code title="get /tci/sessions">client.CodeInterpreter.Sessions.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#CodeInterpreterSessionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#SessionListResponse">SessionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Images

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ImageDataB64">ImageDataB64</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ImageDataURL">ImageDataURL</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ImageFile">ImageFile</a>

Methods:

- <code title="post /images/generations">client.Images.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ImageService.Generate">Generate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ImageGenerateParams">ImageGenerateParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ImageFile">ImageFile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Videos

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#VideoJob">VideoJob</a>

Methods:

- <code title="post /videos">client.Videos.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#VideoService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#VideoNewParams">VideoNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#VideoJob">VideoJob</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /videos/{id}">client.Videos.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#VideoService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#VideoJob">VideoJob</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Audio

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AudioSpeechStreamChunk">AudioSpeechStreamChunk</a>

## Speech

Methods:

- <code title="post /audio/speech">client.Audio.Speech.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AudioSpeechService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AudioSpeechNewParams">AudioSpeechNewParams</a>) (\*http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Voices

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AudioVoiceListResponse">AudioVoiceListResponse</a>

Methods:

- <code title="get /voices">client.Audio.Voices.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AudioVoiceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AudioVoiceListResponse">AudioVoiceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Transcriptions

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AudioTranscriptionNewResponseUnion">AudioTranscriptionNewResponseUnion</a>

Methods:

- <code title="post /audio/transcriptions">client.Audio.Transcriptions.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AudioTranscriptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AudioTranscriptionNewParams">AudioTranscriptionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AudioTranscriptionNewResponseUnion">AudioTranscriptionNewResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Translations

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AudioTranslationNewResponseUnion">AudioTranslationNewResponseUnion</a>

Methods:

- <code title="post /audio/translations">client.Audio.Translations.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AudioTranslationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AudioTranslationNewParams">AudioTranslationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AudioTranslationNewResponseUnion">AudioTranslationNewResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Models

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ModelObject">ModelObject</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ModelUploadResponse">ModelUploadResponse</a>

Methods:

- <code title="get /models">client.Models.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ModelListParams">ModelListParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ModelObject">ModelObject</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /models">client.Models.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ModelService.Upload">Upload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ModelUploadParams">ModelUploadParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ModelUploadResponse">ModelUploadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Jobs

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#JobGetResponse">JobGetResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#JobListResponse">JobListResponse</a>

Methods:

- <code title="get /jobs/{jobId}">client.Jobs.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#JobService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#JobGetResponse">JobGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /jobs">client.Jobs.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#JobService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#JobListResponse">JobListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Endpoints

Params Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AutoscalingParam">AutoscalingParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#Autoscaling">Autoscaling</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#DedicatedEndpoint">DedicatedEndpoint</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointListResponse">EndpointListResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointListAvzonesResponse">EndpointListAvzonesResponse</a>

Methods:

- <code title="post /endpoints">client.Endpoints.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointNewParams">EndpointNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#DedicatedEndpoint">DedicatedEndpoint</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /endpoints/{endpointId}">client.Endpoints.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpointID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#DedicatedEndpoint">DedicatedEndpoint</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /endpoints/{endpointId}">client.Endpoints.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpointID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointUpdateParams">EndpointUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#DedicatedEndpoint">DedicatedEndpoint</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /endpoints">client.Endpoints.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointListParams">EndpointListParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointListResponse">EndpointListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /endpoints/{endpointId}">client.Endpoints.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpointID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /clusters/availability-zones">client.Endpoints.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointService.ListAvzones">ListAvzones</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointListAvzonesResponse">EndpointListAvzonesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Hardware

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#HardwareListResponse">HardwareListResponse</a>

Methods:

- <code title="get /hardware">client.Hardware.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#HardwareService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#HardwareListParams">HardwareListParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#HardwareListResponse">HardwareListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Rerank

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#RerankNewResponse">RerankNewResponse</a>

Methods:

- <code title="post /rerank">client.Rerank.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#RerankService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#RerankNewParams">RerankNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#RerankNewResponse">RerankNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Batches

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BatchJob">BatchJob</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BatchNewResponse">BatchNewResponse</a>

Methods:

- <code title="post /batches">client.Batches.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BatchService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BatchNewParams">BatchNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BatchNewResponse">BatchNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /batches/{id}">client.Batches.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BatchService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BatchJob">BatchJob</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /batches">client.Batches.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BatchService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*[]<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BatchJob">BatchJob</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /batches/{id}/cancel">client.Batches.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BatchService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BatchJob">BatchJob</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Evals

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EvaluationJob">EvaluationJob</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EvalNewResponse">EvalNewResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EvalStatusResponse">EvalStatusResponse</a>

Methods:

- <code title="post /evaluation">client.Evals.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EvalService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EvalNewParams">EvalNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EvalNewResponse">EvalNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /evaluation/{id}">client.Evals.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EvalService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EvaluationJob">EvaluationJob</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /evaluation">client.Evals.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EvalService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EvalListParams">EvalListParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EvaluationJob">EvaluationJob</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /evaluation/{id}/status">client.Evals.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EvalService.Status">Status</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EvalStatusResponse">EvalStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Queue

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#QueueCancelResponse">QueueCancelResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#QueueGetMetricsResponse">QueueGetMetricsResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#QueueGetStatusResponse">QueueGetStatusResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#QueueSubmitResponse">QueueSubmitResponse</a>

Methods:

- <code title="post /queue/cancel">client.Queue.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#QueueService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#QueueCancelParams">QueueCancelParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#QueueCancelResponse">QueueCancelResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /queue/metrics">client.Queue.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#QueueService.GetMetrics">GetMetrics</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#QueueGetMetricsParams">QueueGetMetricsParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#QueueGetMetricsResponse">QueueGetMetricsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /queue/status">client.Queue.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#QueueService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#QueueGetStatusParams">QueueGetStatusParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#QueueGetStatusResponse">QueueGetStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /queue/submit">client.Queue.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#QueueService.Submit">Submit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#QueueSubmitParams">QueueSubmitParams</a>) (\*<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#QueueSubmitResponse">QueueSubmitResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
