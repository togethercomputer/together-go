# together

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#RerankResponse">RerankResponse</a>

Methods:

- <code title="post /rerank">client.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#TogetherService.Rerank">Rerank</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#RerankParams">RerankParams</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#RerankResponse">RerankResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

- <code title="post /chat/completions">client.Chat.Completions.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ChatCompletionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ChatCompletionNewParams">ChatCompletionNewParams</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ChatCompletion">ChatCompletion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

- <code title="post /completions">client.Completions.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#CompletionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#CompletionNewParams">CompletionNewParams</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#Completion">Completion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Embeddings

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#Embedding">Embedding</a>

Methods:

- <code title="post /embeddings">client.Embeddings.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EmbeddingService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EmbeddingNewParams">EmbeddingNewParams</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#Embedding">Embedding</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Files

Params Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FilePurpose">FilePurpose</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileType">FileType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FilePurpose">FilePurpose</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileType">FileType</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileGetResponse">FileGetResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileListResponse">FileListResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileDeleteResponse">FileDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileUploadResponse">FileUploadResponse</a>

Methods:

- <code title="get /files/{id}">client.Files.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileGetResponse">FileGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /files">client.Files.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileListResponse">FileListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /files/{id}">client.Files.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileDeleteResponse">FileDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /files/{id}/content">client.Files.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileService.Content">Content</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /files/upload">client.Files.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileService.Upload">Upload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileUploadParams">FileUploadParams</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileUploadResponse">FileUploadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# FineTune

Params Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#CosineLrSchedulerArgsParam">CosineLrSchedulerArgsParam</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FullTrainingTypeParam">FullTrainingTypeParam</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#LinearLrSchedulerArgsParam">LinearLrSchedulerArgsParam</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#LoRaTrainingTypeParam">LoRaTrainingTypeParam</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#LrSchedulerParam">LrSchedulerParam</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#TrainingMethodDpoParam">TrainingMethodDpoParam</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#TrainingMethodSftParam">TrainingMethodSftParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#CosineLrSchedulerArgs">CosineLrSchedulerArgs</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTune">FineTune</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneEvent">FineTuneEvent</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FullTrainingType">FullTrainingType</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#LinearLrSchedulerArgs">LinearLrSchedulerArgs</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#LoRaTrainingType">LoRaTrainingType</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#LrScheduler">LrScheduler</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#TrainingMethodDpo">TrainingMethodDpo</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#TrainingMethodSft">TrainingMethodSft</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneNewResponse">FineTuneNewResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneListResponse">FineTuneListResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneDeleteResponse">FineTuneDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneCancelResponse">FineTuneCancelResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneDownloadResponse">FineTuneDownloadResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneListEventsResponse">FineTuneListEventsResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneGetCheckpointsResponse">FineTuneGetCheckpointsResponse</a>

Methods:

- <code title="post /fine-tunes">client.FineTune.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneNewParams">FineTuneNewParams</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneNewResponse">FineTuneNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /fine-tunes/{id}">client.FineTune.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTune">FineTune</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /fine-tunes">client.FineTune.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneListResponse">FineTuneListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /fine-tunes/{id}">client.FineTune.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneDeleteParams">FineTuneDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneDeleteResponse">FineTuneDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /fine-tunes/{id}/cancel">client.FineTune.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneCancelResponse">FineTuneCancelResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /finetune/download">client.FineTune.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneService.Download">Download</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneDownloadParams">FineTuneDownloadParams</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneDownloadResponse">FineTuneDownloadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /fine-tunes/{id}/events">client.FineTune.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneService.ListEvents">ListEvents</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneListEventsResponse">FineTuneListEventsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /fine-tunes/{id}/checkpoints">client.FineTune.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneService.GetCheckpoints">GetCheckpoints</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneGetCheckpointsResponse">FineTuneGetCheckpointsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# CodeInterpreter

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ExecuteResponseUnion">ExecuteResponseUnion</a>

Methods:

- <code title="post /tci/execute">client.CodeInterpreter.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#CodeInterpreterService.Execute">Execute</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#CodeInterpreterExecuteParams">CodeInterpreterExecuteParams</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ExecuteResponseUnion">ExecuteResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Sessions

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#SessionListResponse">SessionListResponse</a>

Methods:

- <code title="get /tci/sessions">client.CodeInterpreter.Sessions.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#CodeInterpreterSessionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#SessionListResponse">SessionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Images

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ImageDataB64">ImageDataB64</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ImageDataURL">ImageDataURL</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ImageFile">ImageFile</a>

Methods:

- <code title="post /images/generations">client.Images.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ImageService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ImageNewParams">ImageNewParams</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ImageFile">ImageFile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Videos

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#VideoJob">VideoJob</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#VideoNewResponse">VideoNewResponse</a>

Methods:

- <code title="post /videos">client.Videos.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#VideoService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#VideoNewParams">VideoNewParams</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#VideoNewResponse">VideoNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /videos/{id}">client.Videos.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#VideoService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#VideoJob">VideoJob</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Audio

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AudioSpeechStreamChunk">AudioSpeechStreamChunk</a>

Methods:

- <code title="post /audio/speech">client.Audio.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AudioService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AudioNewParams">AudioNewParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Transcriptions

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AudioTranscriptionNewResponseUnion">AudioTranscriptionNewResponseUnion</a>

Methods:

- <code title="post /audio/transcriptions">client.Audio.Transcriptions.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AudioTranscriptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AudioTranscriptionNewParams">AudioTranscriptionNewParams</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AudioTranscriptionNewResponseUnion">AudioTranscriptionNewResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Translations

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AudioTranslationNewResponseUnion">AudioTranslationNewResponseUnion</a>

Methods:

- <code title="post /audio/translations">client.Audio.Translations.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AudioTranslationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AudioTranslationNewParams">AudioTranslationNewParams</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AudioTranslationNewResponseUnion">AudioTranslationNewResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Models

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ModelListResponse">ModelListResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ModelUploadResponse">ModelUploadResponse</a>

Methods:

- <code title="get /models">client.Models.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ModelListResponse">ModelListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /models">client.Models.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ModelService.Upload">Upload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ModelUploadParams">ModelUploadParams</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ModelUploadResponse">ModelUploadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Jobs

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#JobGetResponse">JobGetResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#JobListResponse">JobListResponse</a>

Methods:

- <code title="get /jobs/{jobId}">client.Jobs.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#JobService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#JobGetResponse">JobGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /jobs">client.Jobs.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#JobService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#JobListResponse">JobListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Endpoints

Params Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#AutoscalingParam">AutoscalingParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#Autoscaling">Autoscaling</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointNewResponse">EndpointNewResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointGetResponse">EndpointGetResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointUpdateResponse">EndpointUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointListResponse">EndpointListResponse</a>

Methods:

- <code title="post /endpoints">client.Endpoints.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointNewParams">EndpointNewParams</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointNewResponse">EndpointNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /endpoints/{endpointId}">client.Endpoints.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpointID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointGetResponse">EndpointGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /endpoints/{endpointId}">client.Endpoints.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpointID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointUpdateParams">EndpointUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointUpdateResponse">EndpointUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /endpoints">client.Endpoints.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointListParams">EndpointListParams</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointListResponse">EndpointListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /endpoints/{endpointId}">client.Endpoints.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EndpointService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpointID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Hardware

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#HardwareListResponse">HardwareListResponse</a>

Methods:

- <code title="get /hardware">client.Hardware.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#HardwareService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#HardwareListParams">HardwareListParams</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#HardwareListResponse">HardwareListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Batches

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BatchNewResponse">BatchNewResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BatchGetResponse">BatchGetResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BatchListResponse">BatchListResponse</a>

Methods:

- <code title="post /batches">client.Batches.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BatchService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BatchNewParams">BatchNewParams</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BatchNewResponse">BatchNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /batches/{id}">client.Batches.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BatchService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BatchGetResponse">BatchGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /batches">client.Batches.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BatchService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#BatchListResponse">BatchListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Evals

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EvalGetResponse">EvalGetResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EvalListResponse">EvalListResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EvalGetAllowedModelsResponse">EvalGetAllowedModelsResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EvalGetStatusResponse">EvalGetStatusResponse</a>

Methods:

- <code title="get /evaluation/{id}">client.Evals.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EvalService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EvalGetResponse">EvalGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /evaluations">client.Evals.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EvalService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EvalListParams">EvalListParams</a>) ([]<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EvalListResponse">EvalListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /evaluations/model-list">client.Evals.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EvalService.GetAllowedModels">GetAllowedModels</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EvalGetAllowedModelsResponse">EvalGetAllowedModelsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /evaluation/{id}/status">client.Evals.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EvalService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#EvalGetStatusResponse">EvalGetStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
