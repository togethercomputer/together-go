# Chat

## Completions

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ChatCompletion">ChatCompletion</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ChatCompletionChunk">ChatCompletionChunk</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ChatCompletionUsage">ChatCompletionUsage</a>

Methods:

- <code title="post /chat/completions">client.Chat.Completions.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ChatCompletionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ChatCompletionNewParams">ChatCompletionNewParams</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ChatCompletion">ChatCompletion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Completions

Params Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ToolChoiceParam">ToolChoiceParam</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ToolsParam">ToolsParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#Completion">Completion</a>
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

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileGetResponse">FileGetResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileListResponse">FileListResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileDeleteResponse">FileDeleteResponse</a>

Methods:

- <code title="get /files/{id}">client.Files.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileGetResponse">FileGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /files">client.Files.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileListResponse">FileListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /files/{id}">client.Files.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileDeleteResponse">FileDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /files/{id}/content">client.Files.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FileService.Content">Content</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# FineTune

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTune">FineTune</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneEvent">FineTuneEvent</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneListResponse">FineTuneListResponse</a>
- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneDownloadResponse">FineTuneDownloadResponse</a>

Methods:

- <code title="post /fine-tunes">client.FineTune.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneNewParams">FineTuneNewParams</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTune">FineTune</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /fine-tunes/{id}">client.FineTune.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTune">FineTune</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /fine-tunes">client.FineTune.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneListResponse">FineTuneListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /fine-tunes/{id}/cancel">client.FineTune.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTune">FineTune</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /finetune/download">client.FineTune.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneService.Download">Download</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneDownloadParams">FineTuneDownloadParams</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneDownloadResponse">FineTuneDownloadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /fine-tunes/{id}/events">client.FineTune.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneService.ListEvents">ListEvents</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#FineTuneEvent">FineTuneEvent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Images

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ImageFile">ImageFile</a>

Methods:

- <code title="post /images/generations">client.Images.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ImageService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ImageNewParams">ImageNewParams</a>) (<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ImageFile">ImageFile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Models

Response Types:

- <a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ModelListResponse">ModelListResponse</a>

Methods:

- <code title="get /models">client.Models.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/togethercomputer/together-go">together</a>.<a href="https://pkg.go.dev/github.com/togethercomputer/together-go#ModelListResponse">ModelListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
