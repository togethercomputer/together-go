# Changelog

## 0.6.0 (2026-01-31)

Full Changelog: [v0.5.0...v0.6.0](https://github.com/togethercomputer/together-go/compare/v0.5.0...v0.6.0)

### Features

* Add chat completion support for reasoning.enabled ([9d821b1](https://github.com/togethercomputer/together-go/commit/9d821b19e7eab5361ca9bb62d13279bfe5cb7ddf))
* **api:** Add API for listing deployments ([7ae9e4d](https://github.com/togethercomputer/together-go/commit/7ae9e4d53e7f35c116411bebcd28f851fa8a1e3b))
* **api:** Add beta sdks for jig features ([f80bbfb](https://github.com/togethercomputer/together-go/commit/f80bbfb5bed41d6c08d36c772e14ee7b11fa1060))
* **api:** Move jobs apis to nest under model uploads per their use case ([b3eec48](https://github.com/togethercomputer/together-go/commit/b3eec48cd07506f2a1668b598b5424d3dba55b88))
* **api:** Move queue out of jig namespace ([a8e9ed4](https://github.com/togethercomputer/together-go/commit/a8e9ed46061dd0f65fe67c7b43f20071b85d8b34))
* **api:** Update Jig types and add retrieve_logs api to jig ([bb63107](https://github.com/togethercomputer/together-go/commit/bb631076f433d47b2f9cabd86e39b9e7ed2f2172))
* **client:** add a convenient param.SetJSON helper ([b8673c7](https://github.com/togethercomputer/together-go/commit/b8673c714b3a42bad6bdc8182a29141d026683e9))
* internal: Add code samples to deployments features ([c0f96ad](https://github.com/togethercomputer/together-go/commit/c0f96ad3aac9815b8ea88392c6a75aceb3aa7dab))
* internal: Update to new cluster api routing ([553f81b](https://github.com/togethercomputer/together-go/commit/553f81b2f1b1aa2cae645debc49ba4f22cb2bf9c))
* move byoc features under /deployments route ([0b9b31e](https://github.com/togethercomputer/together-go/commit/0b9b31e611f82bf02bf4fddfc3a85f09b018b547))


### Chores

* **api:** Improve type names for jig volumes and logs ([b18ce77](https://github.com/togethercomputer/together-go/commit/b18ce7728007f6d3b71ff1251fc49c2cffaa06f6))
* **api:** move hardware listing feature under endpoints resource. ([1b438c0](https://github.com/togethercomputer/together-go/commit/1b438c0fafa3c9f25dc44e42a00ca98b9e21f9e4))
* **api:** Move Queue SDK methods into Jig namespace ([21a6462](https://github.com/togethercomputer/together-go/commit/21a6462c289096dc18e75a84b6e55595d7d41dd9))
* **api:** Rename jig queue apis ([76aef2c](https://github.com/togethercomputer/together-go/commit/76aef2cb7bd3fd59de5700f1f7a356cb551878d4))
* **client:** improve example values ([0d7c1e5](https://github.com/togethercomputer/together-go/commit/0d7c1e5988caef53a71570193c2bb7ba5021d783))


### Documentation

* **axle-queue:** added axle-queue endpoints ([f4809ad](https://github.com/togethercomputer/together-go/commit/f4809adf111711752d2ed781fd696b567e51b46f))

## 0.5.0 (2026-01-21)

Full Changelog: [v0.4.0...v0.5.0](https://github.com/togethercomputer/together-go/compare/v0.4.0...v0.5.0)

### Features

* [wip] ([d677d1f](https://github.com/togethercomputer/together-go/commit/d677d1f8649a16ea8d9572c6e5cf5c78862681c8))
* Add started_at timestamp to fix time estimation ([774f5af](https://github.com/togethercomputer/together-go/commit/774f5af4676523508c208b84d598a593ff02eaa6))
* **client:** add a StreamError type to access raw events from sse streams ([e1bbd04](https://github.com/togethercomputer/together-go/commit/e1bbd04960e7983c589fdc9e486f69c1278dad82))


### Bug Fixes

* **client:** retain streaming when user sets request body ([7824476](https://github.com/togethercomputer/together-go/commit/7824476d76e43d2c96eab88c7668b12a2d3b71ac))
* **docs:** add missing pointer prefix to api.md return types ([22ac5e5](https://github.com/togethercomputer/together-go/commit/22ac5e589d1a4c39894b60143647898c6b42b7a1))


### Chores

* Add code samples and descriptions to instant cluster apis ([24bd713](https://github.com/togethercomputer/together-go/commit/24bd713e6caa21c1a3b350af72f6d8424fc67b3c))
* **api:** Remove APIs that were accidentally added in the wrong namespace ([051991e](https://github.com/togethercomputer/together-go/commit/051991e96f317b5b39ee6965551b3079f28b66eb))
* **internal:** update `actions/checkout` version ([4c69157](https://github.com/togethercomputer/together-go/commit/4c691577cd2a00a307db2e8e9c902ae3a0b69f51))
* Mark disable_prompt_cache as deprecated for endpoint creation ([d086fd1](https://github.com/togethercomputer/together-go/commit/d086fd1914a460e6c364c3b85ee1c129f972b8a7))
* Update cluster apis to reflect their new response shape ([ff8489a](https://github.com/togethercomputer/together-go/commit/ff8489acd3c071c1fef43f70908420aef88c6da2))

## 0.4.0 (2026-01-06)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/togethercomputer/together-go/compare/v0.3.0...v0.4.0)

### Features

* Add compliance and chat_template_kwargs to chat completions spec ([46c9a08](https://github.com/togethercomputer/together-go/commit/46c9a08159c5dd224befc67b311924ee3dc895f5))
* Support VLM finetuning ([5368df6](https://github.com/togethercomputer/together-go/commit/5368df638a2af2cb5d75dc78bb0cc3f078b78d59))
* VLM Support update ([8979fd6](https://github.com/togethercomputer/together-go/commit/8979fd638b1fad9437ae75c8bcbaae4ede154e2e))


### Chores

* add float64 to valid types for RegisterFieldValidator ([934b198](https://github.com/togethercomputer/together-go/commit/934b1989093957486d579f4b183055ed4c67a8c9))
* Add Instant Clusters to OpenAPI spec ([94d3cbd](https://github.com/togethercomputer/together-go/commit/94d3cbd49b592448e028cae8cc1d20ecb7f2306c))
* **internal:** codegen related update ([310bb8b](https://github.com/togethercomputer/together-go/commit/310bb8b88fb2db3f83995ea20b99c656d500c5b2))
* **internal:** use different example values for some enums ([397c345](https://github.com/togethercomputer/together-go/commit/397c345369b0218bd9df8d1c19b55ad0925086b7))

## 0.3.0 (2025-12-16)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/togethercomputer/together-go/compare/v0.2.0...v0.3.0)

### Features

* **files:** add support for string alternative to file upload type ([07c6a44](https://github.com/togethercomputer/together-go/commit/07c6a44da0e74c3331698b517cc5e26aec5461e4))

## 0.2.0 (2025-12-15)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/togethercomputer/together-go/compare/v0.1.0...v0.2.0)

### Features

* **api:** Add fine_tuning.estimate_price api ([b6e4290](https://github.com/togethercomputer/together-go/commit/b6e4290ad67ad0ef0fb13c1dea0e0c10cb4b759d))
* **api:** api update ([be27aa0](https://github.com/togethercomputer/together-go/commit/be27aa07d615dae3f1d5134f25713af1cbc6f6e9))
* **api:** api update ([4431915](https://github.com/togethercomputer/together-go/commit/443191564fd2b856e8ad024f06d5cc20954dc32f))
* **api:** api update ([53b002f](https://github.com/togethercomputer/together-go/commit/53b002f2b43da5c0feb8622783bdef9150b07b81))
* **api:** api update ([68e2d56](https://github.com/togethercomputer/together-go/commit/68e2d565b2db089b39dfa66b7f99e42aa430c157))
* **encoder:** support bracket encoding form-data object members ([9388a38](https://github.com/togethercomputer/together-go/commit/9388a380b64304e89df132c10edf376072c1dc0e))


### Bug Fixes

* **mcp:** correct code tool API endpoint ([895a603](https://github.com/togethercomputer/together-go/commit/895a60327fc7a1143124988e2cb11afbf555b2da))
* rename param to avoid collision ([0e29655](https://github.com/togethercomputer/together-go/commit/0e29655600600ea2f8852e4ed8fc59dfe6a8170a))


### Chores

* elide duplicate aliases ([b75b29f](https://github.com/togethercomputer/together-go/commit/b75b29f189639c8e6f7bf2e0518fe3569e96d98d))
* **internal:** codegen related update ([0d6799b](https://github.com/togethercomputer/together-go/commit/0d6799b6e759b37a402a2b52cf6bdd75afc634e7))

## 0.1.0 (2025-12-05)

Full Changelog: [v0.1.0-alpha.3...v0.1.0](https://github.com/togethercomputer/together-go/compare/v0.1.0-alpha.3...v0.1.0)

### Features

* **api:** api update ([347416a](https://github.com/togethercomputer/together-go/commit/347416ad531f8b85153407c59cb094bf7debf6d8))

## 0.1.0-alpha.3 (2025-12-03)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/togethercomputer/together-go/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### ⚠ BREAKING CHANGES

* **api:** Change call signature for `audio.create` to `audio.speech.create` to match spec with python library and add space for future APIs
* **api:** Update method signature for reranking to `rerank.create()`
* **api:** Change Fine Tuning method name from `download()` to `content()` to align with other namespaces
* **api:** For the TS SDK the `images.create` is now `images.generate`
* **api:** Access to the api for listing checkpoints has changed its name to `list_checkpoints`
* **api:** Access to fine tuning APIs namespace has changed from `fine_tune` to `fine_tuning`
* **api:** The default max retries for api calls has changed from 5 to 2. This may result in more frequent non-200 responses.

### Features

* **api:** Add audio.voices.list sdk ([0aacf5d](https://github.com/togethercomputer/together-go/commit/0aacf5d6e75b326f74921d3c3dd2e88ed1f32b96))
* **api:** Add batches.cancel API ([a8e2951](https://github.com/togethercomputer/together-go/commit/a8e29513556c09d803000fcdadb1a4604736f036))
* **api:** Add endpoints.list_avzones ([c07fe49](https://github.com/togethercomputer/together-go/commit/c07fe49d9ea1202facfd0c4c6019e9119b636297))
* **api:** Add fine_tune.delete API ([dceaf80](https://github.com/togethercomputer/together-go/commit/dceaf802fb93e5ef9e2fb7ed04b80ff4a6dbacec))
* **api:** api update ([c2f758c](https://github.com/togethercomputer/together-go/commit/c2f758c8719cdb7a71b79edbdb53bef06c11df46))
* **api:** api update ([b26c51b](https://github.com/togethercomputer/together-go/commit/b26c51ba708df330cf423e7e8718a8b1f1c4e62d))
* **api:** api update ([be1a06c](https://github.com/togethercomputer/together-go/commit/be1a06c5a1fdbf6106d5bb288b2d1568aec85d2b))
* **api:** api update ([638ebc7](https://github.com/togethercomputer/together-go/commit/638ebc7e0e5ce88228a4ffb6f6f1bd904bb09f8c))
* **api:** api update ([5169921](https://github.com/togethercomputer/together-go/commit/51699214da2ce174e13634af5ca11ac9936cbd13))
* **api:** api update ([48737a3](https://github.com/togethercomputer/together-go/commit/48737a323b7fb5f5272e5c208c7078084c8507ab))
* **api:** api update ([5699d1a](https://github.com/togethercomputer/together-go/commit/5699d1ae3e6e40db16dc6ee1eb136c399e660d1d))
* **api:** api update ([a51a5f3](https://github.com/togethercomputer/together-go/commit/a51a5f3c92661e87dcc2af8b86d709a0c755016e))
* **api:** api update ([2c9ca33](https://github.com/togethercomputer/together-go/commit/2c9ca331f6de3549c4dedb5a78dbc50609d49cf7))
* **api:** api update ([e53f0e7](https://github.com/togethercomputer/together-go/commit/e53f0e79fcf213f40e6537c07f18629d09d9302d))
* **api:** api update ([c64fc6d](https://github.com/togethercomputer/together-go/commit/c64fc6db4e868f65c32d836636b91b65f03b0e8a))
* **api:** Change fine tuning download method to `.create` ([faeb0f8](https://github.com/togethercomputer/together-go/commit/faeb0f8c877a78cf8c77395bf48f7dd55efc131b))
* **api:** Change image creation signature to `images.generate` ([5ec73f1](https://github.com/togethercomputer/together-go/commit/5ec73f1b1f6ac90d705d2a4d03b21a607922eff7))
* **api:** Change rerank method signature ([15519be](https://github.com/togethercomputer/together-go/commit/15519be8442800c3ddbfe1fa808f4d21027a91e1))
* **api:** Change the default max retries from 5 to 2 ([becb776](https://github.com/togethercomputer/together-go/commit/becb77688fe7899159c09a34df6061c0a46716f1))
* **api:** Change TTS call signature ([f906b2e](https://github.com/togethercomputer/together-go/commit/f906b2e51c2a043cdf1bc6e5d057d25787dc9bfb))
* **api:** Fix internal references for VideoJob spec ([bbf9a21](https://github.com/togethercomputer/together-go/commit/bbf9a21463c634a490c389d407f1f07fc098dc7a))
* **api:** manual updates ([651e447](https://github.com/togethercomputer/together-go/commit/651e4473bdc16002ec4fc12374b6d445eaf70bdd))
* **api:** Update Eval APIs ([a2baaa3](https://github.com/togethercomputer/together-go/commit/a2baaa3a41e18d0d9356ffd0f46ae5558a5048a4))


### Bug Fixes

* **client:** correctly specify Accept header with */* instead of empty ([6d504e0](https://github.com/togethercomputer/together-go/commit/6d504e008ed150736cf1e3498d603f9bd9782418))
* remove invalid cast ([72d5d52](https://github.com/togethercomputer/together-go/commit/72d5d521cd2fa88dbf31edb4e5dd9fc9fd195ed4))


### Chores

* **api:** Cleanup some exported types ([aade2f0](https://github.com/togethercomputer/together-go/commit/aade2f06b3c4c705e08fdec6dd4847aec4c8409a))
* **api:** Remove API that is not intended to be public. ([df90a15](https://github.com/togethercomputer/together-go/commit/df90a159de5d770e1e9049d9a9dc7b7befe66921))
* bump gjson version ([704f413](https://github.com/togethercomputer/together-go/commit/704f413869dd157060e6cf96d9f1cfb6aa9cc424))
* fix integration tests ([#106](https://github.com/togethercomputer/together-go/issues/106)) ([8f2d317](https://github.com/togethercomputer/together-go/commit/8f2d317d03fcfff195ae96bfde59d06f39380c17))
* **internal:** grammar fix (it's -&gt; its) ([97b3fc5](https://github.com/togethercomputer/together-go/commit/97b3fc5382b8e4d36340516f06b6c60c7daeb95e))


### Styles

* **api:** Change fine tuning method `retrieve_checkpoints` to `list_checkpoints` ([7e12276](https://github.com/togethercomputer/together-go/commit/7e1227697c594718cedaa70e55be4b7ffd08e836))
* **api:** Change fine tuning namespace to `fine_tuning` ([cfb8297](https://github.com/togethercomputer/together-go/commit/cfb82976c6e84cba4cab24fbe07616fa7d8e561f))

## 0.1.0-alpha.2 (2025-10-30)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/togethercomputer/together-go/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Features

* **api:** api update ([c9508ec](https://github.com/togethercomputer/together-go/commit/c9508ec4f8682c98cccd68e369e651586610d093))
* **api:** api update ([1a6b2c3](https://github.com/togethercomputer/together-go/commit/1a6b2c37351e49d18836a0e4dd26c8a24d8a68a9))

## 0.1.0-alpha.1 (2025-10-21)

Full Changelog: [v0.0.1-alpha.1...v0.1.0-alpha.1](https://github.com/togethercomputer/together-go/compare/v0.0.1-alpha.1...v0.1.0-alpha.1)

### Features

* add SKIP_BREW env var to ./scripts/bootstrap ([#82](https://github.com/togethercomputer/together-go/issues/82)) ([c67fa65](https://github.com/togethercomputer/together-go/commit/c67fa6541c972fa9cee8159ca7f93710cc3df749))
* **api:** add audio create method ([#59](https://github.com/togethercomputer/together-go/issues/59)) ([7dcaec8](https://github.com/togethercomputer/together-go/commit/7dcaec82d9a87dc1b1ef4c77a6c99e687d0f5c7b))
* **api:** add batch api to config ([8b2dab0](https://github.com/togethercomputer/together-go/commit/8b2dab076a494a2506b028d4dafcfc1957d12f63))
* **api:** add evals api to config ([9036aa3](https://github.com/togethercomputer/together-go/commit/9036aa347846d4e026563d22a412ecbf6de9bdce))
* **api:** Add file_type and file_purpose ([ebecf25](https://github.com/togethercomputer/together-go/commit/ebecf250557aea905c39c0797eefb1986a0082d9))
* **api:** add files/upload apu support and switch upload_file method over to use it. ([be1a90e](https://github.com/togethercomputer/together-go/commit/be1a90e9dfaf2269356a74d49154947ff5bf25d7))
* **api:** add models for chat completion structured message types ([#50](https://github.com/togethercomputer/together-go/issues/50)) ([5b71d34](https://github.com/togethercomputer/together-go/commit/5b71d34a3c72c9883a855c1facbc97c24c1cadd5))
* **api:** add tci resources ([1afafee](https://github.com/togethercomputer/together-go/commit/1afafee520bd63db80879d97d4a1c8df13c72939))
* **api:** Add Video APIs ([e3905be](https://github.com/togethercomputer/together-go/commit/e3905bee00459ea29225e55fed78b43c585017f5))
* **api:** adding audio APIs ([7cb450e](https://github.com/togethercomputer/together-go/commit/7cb450e5e11d25d396bb48c77ca7289291406b9f))
* **api:** address diagnostic issues in audio api, correct openapi issue in images api, disambiguate a response in finetune api, enable automated testing on finetune and images ([1be944d](https://github.com/togethercomputer/together-go/commit/1be944d2570130414acf3e23eef7d0d3912123d9))
* **api:** adds unspecified endpoints ([162f10c](https://github.com/togethercomputer/together-go/commit/162f10c4af98dbc77291a8f98848e977b4c72d05))
* **api:** api update ([7151eb3](https://github.com/togethercomputer/together-go/commit/7151eb37250e2985365ff40abd4ee9e2780eb1d4))
* **api:** api update ([2613a22](https://github.com/togethercomputer/together-go/commit/2613a22ea88917c9d1097efa8f4f5307cfde14b2))
* **api:** api update ([fcb7987](https://github.com/togethercomputer/together-go/commit/fcb7987e8308b9717a1bcfb22bca017491ab3a1c))
* **api:** api update ([6209f26](https://github.com/togethercomputer/together-go/commit/6209f268e7ef4c66c4be774d6f6d1b6055e428b9))
* **api:** api update ([bfeb6d7](https://github.com/togethercomputer/together-go/commit/bfeb6d71d28d41b12f62622c4f6acc63c4d723fe))
* **api:** api update ([eb3812f](https://github.com/togethercomputer/together-go/commit/eb3812fa3648afb5e922cbcfab8edf1c0363fa05))
* **api:** api update ([d494d58](https://github.com/togethercomputer/together-go/commit/d494d58855714695a915213bfa5e077254437c45))
* **api:** api update ([a9679b6](https://github.com/togethercomputer/together-go/commit/a9679b6de64d8e6ccecd2dea77450a5be8fa4089))
* **api:** api update ([3c34676](https://github.com/togethercomputer/together-go/commit/3c34676e848b6db1df0e59718777710a88f289a5))
* **api:** api update ([7c7bb40](https://github.com/togethercomputer/together-go/commit/7c7bb405bfe7f90a5295c79c92022599440f8789))
* **api:** api update ([e1cfe52](https://github.com/togethercomputer/together-go/commit/e1cfe5260c860178fd5ccf591b8ed64794e4ff99))
* **api:** api update ([3a516e3](https://github.com/togethercomputer/together-go/commit/3a516e36a5642ad090a7bf76e5d4ed03fafd0525))
* **api:** api update ([35df411](https://github.com/togethercomputer/together-go/commit/35df411660fc506ab66e022536f3d656fd13234d))
* **api:** api update ([3c4d345](https://github.com/togethercomputer/together-go/commit/3c4d34501c888be8563f4fd6bdc1025c7ec44e65))
* **api:** api update ([6e212af](https://github.com/togethercomputer/together-go/commit/6e212af45999e8210d3861e7473ea3027a2d16a3))
* **api:** api update ([c4dcfd1](https://github.com/togethercomputer/together-go/commit/c4dcfd1f903173fd52c66be828c4c80fcd20eee2))
* **api:** api update ([b0bec16](https://github.com/togethercomputer/together-go/commit/b0bec16c8b67a5785623790a4ba0eccf049f54d4))
* **api:** api update ([41c8ee1](https://github.com/togethercomputer/together-go/commit/41c8ee101527d9ee8a1be08c7ea6224f0b7a699b))
* **api:** api update ([57547b1](https://github.com/togethercomputer/together-go/commit/57547b1d18bcaa7229e5b3fe6a5c29899db64ff1))
* **api:** api update ([da325b4](https://github.com/togethercomputer/together-go/commit/da325b4a643284dd617e017d4423f8de437c92b0))
* **api:** api update ([641bca8](https://github.com/togethercomputer/together-go/commit/641bca810be50e9297326c4a66ca81f9d3f5263c))
* **api:** api update ([862f389](https://github.com/togethercomputer/together-go/commit/862f389ef399400ea4d92327c373ea1300ebc50e))
* **api:** api update ([e33e547](https://github.com/togethercomputer/together-go/commit/e33e54797cf1652942bdaf759139dd7d85bda3e7))
* **api:** api update ([947f64c](https://github.com/togethercomputer/together-go/commit/947f64ca05e3175e47bb76c1f6b25fe9e3782262))
* **api:** api update ([128beb1](https://github.com/togethercomputer/together-go/commit/128beb1237d45dd7285fd0274ac189af57509667))
* **api:** api update ([62e53ec](https://github.com/togethercomputer/together-go/commit/62e53ec6228de6cf6446a89664db297b9489463f))
* **api:** api update ([9443342](https://github.com/togethercomputer/together-go/commit/9443342a9ecead016600b895af933121fcd6827c))
* **api:** api update ([a857087](https://github.com/togethercomputer/together-go/commit/a85708793ec113caf647b86d8b12116347390a87))
* **api:** api update ([09ea398](https://github.com/togethercomputer/together-go/commit/09ea398ef901f01466fce41cf191fa5a96afbbb9))
* **api:** api update ([3b587e1](https://github.com/togethercomputer/together-go/commit/3b587e1c3d717fbf3b57368091df32a71c7a09fe))
* **api:** api update ([#32](https://github.com/togethercomputer/together-go/issues/32)) ([702036a](https://github.com/togethercomputer/together-go/commit/702036a7acd27155ee29a456f920c1c52394a0a9))
* **api:** api update ([#33](https://github.com/togethercomputer/together-go/issues/33)) ([808507d](https://github.com/togethercomputer/together-go/commit/808507db434e30b567414b8ca62eaee2f7b1b33f))
* **api:** api update ([#34](https://github.com/togethercomputer/together-go/issues/34)) ([ba315f5](https://github.com/togethercomputer/together-go/commit/ba315f56c399edcc51ef0ac9c16a26e7b1888375))
* **api:** api update ([#35](https://github.com/togethercomputer/together-go/issues/35)) ([6019a6f](https://github.com/togethercomputer/together-go/commit/6019a6f90d28ab91f247442df80805b7c90e8a27))
* **api:** api update ([#36](https://github.com/togethercomputer/together-go/issues/36)) ([6110d6e](https://github.com/togethercomputer/together-go/commit/6110d6e9e0ca4c649478c3ced391aea5bf3225e9))
* **api:** api update ([#37](https://github.com/togethercomputer/together-go/issues/37)) ([bfd0cee](https://github.com/togethercomputer/together-go/commit/bfd0cee27d841a3f406274770d9c0f69950dfab9))
* **api:** api update ([#38](https://github.com/togethercomputer/together-go/issues/38)) ([62dd3d7](https://github.com/togethercomputer/together-go/commit/62dd3d7f9e56ad2caedb5adb73dc7ffacdf8a088))
* **api:** api update ([#43](https://github.com/togethercomputer/together-go/issues/43)) ([d8a9b55](https://github.com/togethercomputer/together-go/commit/d8a9b55ab9d644035c64039e498eeebfd856a4ed))
* **api:** api update ([#45](https://github.com/togethercomputer/together-go/issues/45)) ([f0a3b59](https://github.com/togethercomputer/together-go/commit/f0a3b598cb71731d98a3309221d1ac18c6798a1c))
* **api:** api update ([#46](https://github.com/togethercomputer/together-go/issues/46)) ([c28a6de](https://github.com/togethercomputer/together-go/commit/c28a6de9935754796419e35d47d24180901f846b))
* **api:** api update ([#47](https://github.com/togethercomputer/together-go/issues/47)) ([1976ac7](https://github.com/togethercomputer/together-go/commit/1976ac73b96813e09ca5fc4f8fe91269bcb53a28))
* **api:** api update ([#48](https://github.com/togethercomputer/together-go/issues/48)) ([afcd02e](https://github.com/togethercomputer/together-go/commit/afcd02ea954f5f9c3185fb4f2cd3afce6b9e64fe))
* **api:** api update ([#49](https://github.com/togethercomputer/together-go/issues/49)) ([4cab411](https://github.com/togethercomputer/together-go/commit/4cab411a9e84569c3b79748fc8135dbb6e13d651))
* **api:** api update ([#60](https://github.com/togethercomputer/together-go/issues/60)) ([cc1cdb3](https://github.com/togethercomputer/together-go/commit/cc1cdb3984182c8c1e7bc355441d82eff7c5900a))
* **api:** api update ([#65](https://github.com/togethercomputer/together-go/issues/65)) ([ebf6be5](https://github.com/togethercomputer/together-go/commit/ebf6be570c2fd59e9768dcc00dbb666c991914a5))
* **api:** api update ([#71](https://github.com/togethercomputer/together-go/issues/71)) ([a28a155](https://github.com/togethercomputer/together-go/commit/a28a155b09908951e1816fcd5b5d7ff97ac6ad5b))
* **api:** api update ([#74](https://github.com/togethercomputer/together-go/issues/74)) ([c8af2c2](https://github.com/togethercomputer/together-go/commit/c8af2c2b4aa909e0ebe4d1b6c7eeb12abd522d72))
* **api:** api update ([#76](https://github.com/togethercomputer/together-go/issues/76)) ([e0074c9](https://github.com/togethercomputer/together-go/commit/e0074c987aeeb2b365e1191eaa5d046767dbd452))
* **api:** api update ([#79](https://github.com/togethercomputer/together-go/issues/79)) ([3567145](https://github.com/togethercomputer/together-go/commit/35671455f95299495016d8dd8bcce8e5059847f9))
* **api:** api update ([#80](https://github.com/togethercomputer/together-go/issues/80)) ([b76eea1](https://github.com/togethercomputer/together-go/commit/b76eea11bfffbb6e1d499f699a900ca8997dab89))
* **api:** api update ([#85](https://github.com/togethercomputer/together-go/issues/85)) ([3df0e39](https://github.com/togethercomputer/together-go/commit/3df0e39d2198f0142e3bca06064d4666b620b820))
* **api:** api update ([#86](https://github.com/togethercomputer/together-go/issues/86)) ([5df18cb](https://github.com/togethercomputer/together-go/commit/5df18cb1695d3f2e74811d23c1d12dd3dbabd90d))
* **api:** api update ([#90](https://github.com/togethercomputer/together-go/issues/90)) ([a93098c](https://github.com/togethercomputer/together-go/commit/a93098c8549a4f302e4af80fae1c283af41e8303))
* **api:** api update ([#95](https://github.com/togethercomputer/together-go/issues/95)) ([703d4b8](https://github.com/togethercomputer/together-go/commit/703d4b86ea3916e9e32e1661f9d6b3d8545ea9fc))
* **api:** api update ([#99](https://github.com/togethercomputer/together-go/issues/99)) ([168fd0e](https://github.com/togethercomputer/together-go/commit/168fd0e0529b2b17776789928573623466711b69))
* **api:** Formatting fixes, some lint fixes ([8b1a344](https://github.com/togethercomputer/together-go/commit/8b1a3442d954269d0e37d36685d321010b56dffb))
* **api:** get test_code_interpreter passing ([594e65a](https://github.com/togethercomputer/together-go/commit/594e65afa5b3cd8eaf7e57fa76efba80a4f267b9))
* **api:** manual updates ([99f4880](https://github.com/togethercomputer/together-go/commit/99f48809fb77544e6e36a3ac78708d166840b33e))
* **api:** manual updates ([e3f45e3](https://github.com/togethercomputer/together-go/commit/e3f45e31e3fb271cd533174e14d84cf69729e990))
* **api:** manual updates ([de64c64](https://github.com/togethercomputer/together-go/commit/de64c641c1769d682826b23f31f26ce4ae420202))
* **api:** manual updates ([03aa5bb](https://github.com/togethercomputer/together-go/commit/03aa5bb54f1f83575ccda0217beb82002534ff74))
* **api:** manual updates ([6a602e4](https://github.com/togethercomputer/together-go/commit/6a602e4b2e5d6039ad93cc5021354ef9cb99cb3c))
* **api:** manual updates ([5181a43](https://github.com/togethercomputer/together-go/commit/5181a43dcd9622a212d59a026b346f20ee386711))
* **api:** manual updates ([#15](https://github.com/togethercomputer/together-go/issues/15)) ([e96f5f3](https://github.com/togethercomputer/together-go/commit/e96f5f3691107983735873349f4594f2d9e1d438))
* **api:** manual updates ([#9](https://github.com/togethercomputer/together-go/issues/9)) ([7f6085b](https://github.com/togethercomputer/together-go/commit/7f6085b9138d4354ce942c73ba9b8b2ca9524b3d))
* **api:** OpenAPI spec update via Stainless API ([#10](https://github.com/togethercomputer/together-go/issues/10)) ([8785874](https://github.com/togethercomputer/together-go/commit/87858748c80d12b102d05eb9ccebb77171f347ad))
* **api:** OpenAPI spec update via Stainless API ([#11](https://github.com/togethercomputer/together-go/issues/11)) ([3d8a2a8](https://github.com/togethercomputer/together-go/commit/3d8a2a85759db4a692f9bd8472846aea16ab63c2))
* **api:** OpenAPI spec update via Stainless API ([#12](https://github.com/togethercomputer/together-go/issues/12)) ([67ce61b](https://github.com/togethercomputer/together-go/commit/67ce61ba8d57e86be4d545ae5e11c828eda3dc44))
* **api:** OpenAPI spec update via Stainless API ([#13](https://github.com/togethercomputer/together-go/issues/13)) ([3e256d6](https://github.com/togethercomputer/together-go/commit/3e256d6c0b5f081c76c6d866d6c8f56f6f2eefe8))
* **api:** OpenAPI spec update via Stainless API ([#14](https://github.com/togethercomputer/together-go/issues/14)) ([a551a1c](https://github.com/togethercomputer/together-go/commit/a551a1cf884e072a704076efdcb483e45b67d8b6))
* **api:** OpenAPI spec update via Stainless API ([#20](https://github.com/togethercomputer/together-go/issues/20)) ([65860e7](https://github.com/togethercomputer/together-go/commit/65860e73a24cc6d2add4437b4b85753b83adc623))
* **api:** OpenAPI spec update via Stainless API ([#22](https://github.com/togethercomputer/together-go/issues/22)) ([9438559](https://github.com/togethercomputer/together-go/commit/943855908662362dda93bae3816bf729d25b4a90))
* **api:** OpenAPI spec update via Stainless API ([#23](https://github.com/togethercomputer/together-go/issues/23)) ([0975968](https://github.com/togethercomputer/together-go/commit/097596895ef968204c188986331c98891ee98e7b))
* **api:** OpenAPI spec update via Stainless API ([#25](https://github.com/togethercomputer/together-go/issues/25)) ([53a6e5c](https://github.com/togethercomputer/together-go/commit/53a6e5c4f694c7bdb6470ecdb622f40b2fb17e74))
* **api:** OpenAPI spec update via Stainless API ([#26](https://github.com/togethercomputer/together-go/issues/26)) ([a0fe177](https://github.com/togethercomputer/together-go/commit/a0fe177dcaaed3cb1d4499be3ec2e280211df6c9))
* **api:** OpenAPI spec update via Stainless API ([#27](https://github.com/togethercomputer/together-go/issues/27)) ([2966f79](https://github.com/togethercomputer/together-go/commit/2966f7930a5793137473a3546032e4827155a4cc))
* **api:** OpenAPI spec update via Stainless API ([#29](https://github.com/togethercomputer/together-go/issues/29)) ([b6124a4](https://github.com/togethercomputer/together-go/commit/b6124a4baf354c1b82a7390431acd8ab692184fb))
* **api:** OpenAPI spec update via Stainless API ([#4](https://github.com/togethercomputer/together-go/issues/4)) ([de76795](https://github.com/togethercomputer/together-go/commit/de76795dfbe74d608d3b976c50d032a731316c5b))
* **api:** OpenAPI spec update via Stainless API ([#6](https://github.com/togethercomputer/together-go/issues/6)) ([ed5ed53](https://github.com/togethercomputer/together-go/commit/ed5ed5323787ed0ae3d8b897aa54b8277311bb20))
* **api:** OpenAPI spec update via Stainless API ([#7](https://github.com/togethercomputer/together-go/issues/7)) ([54889c4](https://github.com/togethercomputer/together-go/commit/54889c4867d1313191f13893386c93f61f75b4e5))
* **api:** removed streaming from translation/transcription endpoints ([de822f4](https://github.com/togethercomputer/together-go/commit/de822f4de0bc42959d151a61e7371309d9644730))
* **api:** Rename evaluation sdks to evals ([46d2fe0](https://github.com/togethercomputer/together-go/commit/46d2fe07b77119a6b0f27960b206cbdb71098a99))
* **api:** update spec / config to remove remaining codegen warnings ([388034e](https://github.com/togethercomputer/together-go/commit/388034e975f781b6f5d89f63912ff53ced874405))
* **api:** Update spec and config to get all tests except code-interpolation an fine_tune unit tests working. ([58f8450](https://github.com/togethercomputer/together-go/commit/58f8450581ebd3959dd6ba34c94f4248091eb8a6))
* **client:** accept RFC6838 JSON content types ([#83](https://github.com/togethercomputer/together-go/issues/83)) ([6ac248d](https://github.com/togethercomputer/together-go/commit/6ac248d2afaed4b5cfba5be119b575c84b5dfc4a))
* **client:** add debug log helper ([25351e3](https://github.com/togethercomputer/together-go/commit/25351e3b4d10bd007f46e33e1c7a75c22f06e52b))
* **client:** add dynamic streaming buffer to handle large lines ([0adad00](https://github.com/togethercomputer/together-go/commit/0adad0079eaa95d47320d666ce2494a5fb7d4597))
* **client:** add support for endpoint-specific base URLs in python ([3da11c8](https://github.com/togethercomputer/together-go/commit/3da11c8c4bc0eb4e32f1dae9f2c5f2f35cc30101))
* **client:** add support for reading base URL from environment variable ([ee16864](https://github.com/togethercomputer/together-go/commit/ee16864be13c16dcb734cf2f04709c62aa5bff02))
* **client:** allow custom baseurls without trailing slash ([#81](https://github.com/togethercomputer/together-go/issues/81)) ([4770789](https://github.com/togethercomputer/together-go/commit/4770789bad55a4093b2530de346f7b57d377b39b))
* **client:** expand max streaming buffer size ([bf3a03e](https://github.com/togethercomputer/together-go/commit/bf3a03e6ab352f29f63fb7569cf5961491e974ad))
* **client:** improve default client options support ([#87](https://github.com/togethercomputer/together-go/issues/87)) ([a134d38](https://github.com/togethercomputer/together-go/commit/a134d389d5a8e278641ec2918a55258cec31b3cc))
* **client:** send `X-Stainless-Timeout` header ([#67](https://github.com/togethercomputer/together-go/issues/67)) ([5b9cba4](https://github.com/togethercomputer/together-go/commit/5b9cba46e3c1d85b286160ad18e4821391209757))
* **client:** support custom http clients ([#100](https://github.com/togethercomputer/together-go/issues/100)) ([0f30cd9](https://github.com/togethercomputer/together-go/commit/0f30cd9aeb3cbc8c01672bd4d2ccd26973943892))
* **client:** support optional json html escaping ([ab26ec5](https://github.com/togethercomputer/together-go/commit/ab26ec5e6ea3095a9c4285508f48fc1e6539455e))
* support deprecated markers ([#58](https://github.com/togethercomputer/together-go/issues/58)) ([d0fad57](https://github.com/togethercomputer/together-go/commit/d0fad573346391201e0559b8f8b200bfb653ae21))


### Bug Fixes

* bugfix for setting JSON keys with special characters ([a880cae](https://github.com/togethercomputer/together-go/commit/a880cae12944343a90065d9127295875d5255f35))
* **client:** clean up reader resources ([e2121d5](https://github.com/togethercomputer/together-go/commit/e2121d5b0e79d1995459b7bdbda46d99c5c13b11))
* **client:** correct type to enum ([#89](https://github.com/togethercomputer/together-go/issues/89)) ([5d4f153](https://github.com/togethercomputer/together-go/commit/5d4f15338c27362060a4353011d9ac8fc50e6d8f))
* **client:** correctly update body in WithJSONSet ([099102c](https://github.com/togethercomputer/together-go/commit/099102c299b47a4c438d108d144b6e56c4bb223c))
* **client:** don't truncate manually specified filenames ([#75](https://github.com/togethercomputer/together-go/issues/75)) ([72cdfc1](https://github.com/togethercomputer/together-go/commit/72cdfc1f02601b7255644a7f8e212e6215d14c5f))
* **client:** fix auth via Bearer token ([#8](https://github.com/togethercomputer/together-go/issues/8)) ([fd4ed4e](https://github.com/togethercomputer/together-go/commit/fd4ed4ea0fb8e367bba7b8c534e7dbc0818c29f5))
* **client:** increase max stream buffer size ([7123a09](https://github.com/togethercomputer/together-go/commit/7123a09f2d7a015558ea7cf51aa86681f3ae5e56))
* **client:** process custom base url ahead of time ([65a477e](https://github.com/togethercomputer/together-go/commit/65a477e57715bc801c7ff1157c26d202884988ac))
* **client:** return error on bad custom url instead of panic ([#98](https://github.com/togethercomputer/together-go/issues/98)) ([9b78b47](https://github.com/togethercomputer/together-go/commit/9b78b4727d8def333fbcc527df9184f0aefaef5b))
* **client:** unmarshal stream events into fresh memory ([#97](https://github.com/togethercomputer/together-go/issues/97)) ([c40239d](https://github.com/togethercomputer/together-go/commit/c40239d244dcfa8dbe3e2c8ae49e6956d3cc2a8a))
* **client:** use scanner for streaming ([86b8659](https://github.com/togethercomputer/together-go/commit/86b8659f410f7f7dfa65acaad795e2b75df9cd3c))
* close body before retrying ([ed0c1eb](https://github.com/togethercomputer/together-go/commit/ed0c1ebc7ccace7ef243a376d5c0f2cbf05d1858))
* deserialization of struct unions that implement json.Unmarshaler ([#18](https://github.com/togethercomputer/together-go/issues/18)) ([08bf08b](https://github.com/togethercomputer/together-go/commit/08bf08b629b1989ebaf360e4dacf41fe37a11574))
* do not call path.Base on ContentType ([#73](https://github.com/togethercomputer/together-go/issues/73)) ([9d91dc0](https://github.com/togethercomputer/together-go/commit/9d91dc033c9815e89a1271fbb299d5a90212bde5))
* don't try to deserialize as json when ResponseBodyInto is []byte ([892786b](https://github.com/togethercomputer/together-go/commit/892786bd64e1f30ebc11c47db354f8d7981c7e25))
* fix apijson.Port for embedded structs ([#55](https://github.com/togethercomputer/together-go/issues/55)) ([926697d](https://github.com/togethercomputer/together-go/commit/926697d9ad31e27a50c16f42d1131ef43573f40b))
* fix apijson.Port for embedded structs ([#57](https://github.com/togethercomputer/together-go/issues/57)) ([03c6cf9](https://github.com/togethercomputer/together-go/commit/03c6cf9b26052042e53845a06a909544065667a9))
* fix early cancel when RequestTimeout is provided for streaming requests ([#72](https://github.com/togethercomputer/together-go/issues/72)) ([7a3ab38](https://github.com/togethercomputer/together-go/commit/7a3ab38f113e5515473568ac9ffb774abb2c9425))
* fix unicode encoding for json ([#63](https://github.com/togethercomputer/together-go/issues/63)) ([2abbd4e](https://github.com/togethercomputer/together-go/commit/2abbd4ee27747e561ca0e10c099c7781f064b435))
* handle empty bodies in WithJSONSet ([3285ea3](https://github.com/togethercomputer/together-go/commit/3285ea3f6aa2806c3742c6662858ba1b8311e1f0))
* remove null from release please manifest ([2aba684](https://github.com/togethercomputer/together-go/commit/2aba684529bba40ac80b1ca6c58f9f6b16cdce20))
* **requestconfig:** copy over more fields when cloning ([#28](https://github.com/togethercomputer/together-go/issues/28)) ([d327497](https://github.com/togethercomputer/together-go/commit/d32749788a265a1552462f22a52394d177789d05))
* skip invalid fine-tune test ([#96](https://github.com/togethercomputer/together-go/issues/96)) ([4db5dc7](https://github.com/togethercomputer/together-go/commit/4db5dc79d81cf45d5a5c527f3aaa72a9db1df0fb))
* **stream:** ensure .Close() doesn't panic ([#66](https://github.com/togethercomputer/together-go/issues/66)) ([698370a](https://github.com/togethercomputer/together-go/commit/698370a64c3ee1035d8ba9e166fc5e175d26afe6))
* **test:** return early after test failure ([#93](https://github.com/togethercomputer/together-go/issues/93)) ([29ad02d](https://github.com/togethercomputer/together-go/commit/29ad02de097206b4e7488acf10b6b7cee3c1f964))
* **tests:** correctly skip create fine tune tests ([#101](https://github.com/togethercomputer/together-go/issues/101)) ([26861ca](https://github.com/togethercomputer/together-go/commit/26861cac50ad538930c7860216007fa1f92e3694))
* **tests:** skip invalid test ([#61](https://github.com/togethercomputer/together-go/issues/61)) ([3d32240](https://github.com/togethercomputer/together-go/commit/3d3224071d382c32144ad019436627f76a7a5854))
* update stream error handling ([#70](https://github.com/togethercomputer/together-go/issues/70)) ([146526d](https://github.com/togethercomputer/together-go/commit/146526d12cbdf545c5c2c684624e8ba17db10529))
* use release please annotations on more places ([8711e90](https://github.com/togethercomputer/together-go/commit/8711e90df8e8fcb85ee68bb2df8d596c986e2a14))
* use slices.Concat instead of sometimes modifying r.Options ([1461e1b](https://github.com/togethercomputer/together-go/commit/1461e1b7e1be56de8ae5cec2967137e56b34e739))


### Chores

* add request options to client tests ([#92](https://github.com/togethercomputer/together-go/issues/92)) ([31eddb9](https://github.com/togethercomputer/together-go/commit/31eddb96ebd539f6924c10213b71a6890fb8c30e))
* add UnionUnmarshaler for responses that are interfaces ([#69](https://github.com/togethercomputer/together-go/issues/69)) ([fcb5a8d](https://github.com/togethercomputer/together-go/commit/fcb5a8d7d74659d61236b67f3476fda2e47f0139))
* **api:** re-enable audio unit tests ([726adba](https://github.com/togethercomputer/together-go/commit/726adba9c9a00b796805f6ecff2996672d155d1c))
* bump Go to v1.21 ([#19](https://github.com/togethercomputer/together-go/issues/19)) ([5042a00](https://github.com/togethercomputer/together-go/commit/5042a00c400c63a6167045ac946be807a0294f9f))
* bump minimum go version to 1.22 ([abe802b](https://github.com/togethercomputer/together-go/commit/abe802bb9316404de5274d97cd724e7354065e48))
* **ci:** add timeout thresholds for CI jobs ([684b37e](https://github.com/togethercomputer/together-go/commit/684b37ef4e5b6fa04d517c03e1a5bab3bd3cdb00))
* **ci:** bump prism mock server version ([#17](https://github.com/togethercomputer/together-go/issues/17)) ([1054d43](https://github.com/togethercomputer/together-go/commit/1054d43afab641d41b69704da0925d01fd9b2995))
* **ci:** enable for pull requests ([97b1c71](https://github.com/togethercomputer/together-go/commit/97b1c719f37051d5fbd9b283229791d08bb75e14))
* **ci:** only run for pushes and fork pull requests ([c8d642e](https://github.com/togethercomputer/together-go/commit/c8d642ecc7cdc0830ee38f4c26a091eee38a8d8f))
* **ci:** only use depot for staging repos ([3def4b6](https://github.com/togethercomputer/together-go/commit/3def4b669d4c78e799a16e7af9c4ef33fbbbf3e7))
* **ci:** run on more branches and use depot runners ([8ed4f7f](https://github.com/togethercomputer/together-go/commit/8ed4f7f28422056ac7c85a715bd7e0b1d75bd62b))
* do not install brew dependencies in ./scripts/bootstrap by default ([635d52d](https://github.com/togethercomputer/together-go/commit/635d52d5f507371ed83ff3372081d50cd9a7e4cb))
* **docs:** add docstring explaining streaming pattern ([#68](https://github.com/togethercomputer/together-go/issues/68)) ([b819cd2](https://github.com/togethercomputer/together-go/commit/b819cd2eeeceaefd66e31d96a46a3423b91883c0))
* **docs:** document pre-request options ([75903e0](https://github.com/togethercomputer/together-go/commit/75903e06a2bacbcb6f2f68a530f96c5612f972b1))
* **docs:** grammar improvements ([5ff5557](https://github.com/togethercomputer/together-go/commit/5ff5557ac052563d32d21dd2abbb56e567dcf732))
* **docs:** improve security documentation ([#91](https://github.com/togethercomputer/together-go/issues/91)) ([ad5ea37](https://github.com/togethercomputer/together-go/commit/ad5ea3745eea9f2b0a76eb68f3ed0b1c359831dd))
* **examples:** minor formatting changes ([#21](https://github.com/togethercomputer/together-go/issues/21)) ([f0be048](https://github.com/togethercomputer/together-go/commit/f0be0485928606c04ad69eb7c3f4a209503848a1))
* fix typos ([#94](https://github.com/togethercomputer/together-go/issues/94)) ([ba10a6a](https://github.com/togethercomputer/together-go/commit/ba10a6abc9a15b7f516008ff3c3bb8ac2114e752))
* improve devcontainer setup ([a4e3f3b](https://github.com/togethercomputer/together-go/commit/a4e3f3b201a2f108c0ccabe7831b4272b37225f5))
* **internal:** codegen related update ([1ed14ac](https://github.com/togethercomputer/together-go/commit/1ed14ac5a618b006e3a523900087a4c2994c76ec))
* **internal:** codegen related update ([#16](https://github.com/togethercomputer/together-go/issues/16)) ([1159109](https://github.com/togethercomputer/together-go/commit/11591093f9a3f58a998c89b0975a488f8dc13ea1))
* **internal:** codegen related update ([#24](https://github.com/togethercomputer/together-go/issues/24)) ([e007726](https://github.com/togethercomputer/together-go/commit/e0077267129e0585f2c405025794abbe12e84d24))
* **internal:** codegen related update ([#30](https://github.com/togethercomputer/together-go/issues/30)) ([60239e2](https://github.com/togethercomputer/together-go/commit/60239e28685f714ab268c9926d0750830355301b))
* **internal:** codegen related update ([#51](https://github.com/togethercomputer/together-go/issues/51)) ([26db20a](https://github.com/togethercomputer/together-go/commit/26db20a78f18b3d23630f255a7deeb9c1c366ecb))
* **internal:** codegen related update ([#52](https://github.com/togethercomputer/together-go/issues/52)) ([f3ae938](https://github.com/togethercomputer/together-go/commit/f3ae938f54db2cf0f763f1e65d12e415fadbd566))
* **internal:** codegen related update ([#53](https://github.com/togethercomputer/together-go/issues/53)) ([234c94b](https://github.com/togethercomputer/together-go/commit/234c94b18c44d2a5b40eb513ab452fa7816df76d))
* **internal:** codegen related update ([#54](https://github.com/togethercomputer/together-go/issues/54)) ([fc5f63b](https://github.com/togethercomputer/together-go/commit/fc5f63bc3560deb448bfc4739653969370815832))
* **internal:** codegen related update ([#56](https://github.com/togethercomputer/together-go/issues/56)) ([e01ff9b](https://github.com/togethercomputer/together-go/commit/e01ff9b75ae894255f8419f9b2c50e85ebe0c70f))
* **internal:** codegen related update ([#62](https://github.com/togethercomputer/together-go/issues/62)) ([c81ae0b](https://github.com/togethercomputer/together-go/commit/c81ae0b6681009692aadac42eeffb785d83afe23))
* **internal:** expand CI branch coverage ([#102](https://github.com/togethercomputer/together-go/issues/102)) ([ec17a31](https://github.com/togethercomputer/together-go/commit/ec17a316f950c33c69922c0172ee64a7efdc2074))
* **internal:** fix devcontainers setup ([#77](https://github.com/togethercomputer/together-go/issues/77)) ([65dde36](https://github.com/togethercomputer/together-go/commit/65dde3673772a4f31d335e180a8d513ed4de55bf))
* **internal:** fix lint script for tests ([0c992dd](https://github.com/togethercomputer/together-go/commit/0c992dd375546264f822759d3b29082010b42ef9))
* **internal:** reduce CI branch coverage ([af0a83d](https://github.com/togethercomputer/together-go/commit/af0a83d6b2028ebf809ce32c50ecb21d058def14))
* **internal:** remove extra empty newlines ([#88](https://github.com/togethercomputer/together-go/issues/88)) ([de182e6](https://github.com/togethercomputer/together-go/commit/de182e6fd2c1bf0266600b33d311a171998ca43f))
* **internal:** update comment in script ([a0d23ee](https://github.com/togethercomputer/together-go/commit/a0d23ee9559f571a9abd01fcc245668da75fea85))
* **internal:** update test skipping reason ([47d73d9](https://github.com/togethercomputer/together-go/commit/47d73d95da98f7d954a36119fb532b7dec58300f))
* lint tests ([4664b94](https://github.com/togethercomputer/together-go/commit/4664b94fddba8a74a7ea314bbde8f5a9536866f4))
* lint tests in subpackages ([742e27e](https://github.com/togethercomputer/together-go/commit/742e27e26e5f416fd1e369297e12a7c1dceda44f))
* make go mod tidy continue on error ([5e3ad8b](https://github.com/togethercomputer/together-go/commit/5e3ad8b045fa7b83da718c85ffe4066cd1c43e0a))
* rebuild project due to codegen change ([#39](https://github.com/togethercomputer/together-go/issues/39)) ([a926b78](https://github.com/togethercomputer/together-go/commit/a926b7851067ac13f15b4e73c6690d0dede5ac35))
* rebuild project due to codegen change ([#40](https://github.com/togethercomputer/together-go/issues/40)) ([2a910da](https://github.com/togethercomputer/together-go/commit/2a910dac5993fefebbe2458b3f725fd53ed5ff7c))
* rebuild project due to codegen change ([#41](https://github.com/togethercomputer/together-go/issues/41)) ([246009d](https://github.com/togethercomputer/together-go/commit/246009d1269d11260bff67eece802d18d876ea82))
* rebuild project due to codegen change ([#42](https://github.com/togethercomputer/together-go/issues/42)) ([3f01cf1](https://github.com/togethercomputer/together-go/commit/3f01cf16a23b4f97c7179039d5912f311f794191))
* rebuild project due to codegen change ([#44](https://github.com/togethercomputer/together-go/issues/44)) ([18ef904](https://github.com/togethercomputer/together-go/commit/18ef90400e637711a2d6bb6a353b5b0054fc6952))
* update @stainless-api/prism-cli to v5.15.0 ([f5daf38](https://github.com/togethercomputer/together-go/commit/f5daf3897a989b799e4a22e78e2ac39f5e66e08f))
* update more docs for 1.22 ([efd4e9a](https://github.com/togethercomputer/together-go/commit/efd4e9a98047c17fa001759585a3bedfe90ae60b))


### Documentation

* document raw responses ([#64](https://github.com/togethercomputer/together-go/issues/64)) ([d1eab51](https://github.com/togethercomputer/together-go/commit/d1eab510c000dbcf8363ab1f398f43dd016fcb9f))
* update documentation links to be more uniform ([bd8ed3d](https://github.com/togethercomputer/together-go/commit/bd8ed3d1a04242b1b836e35e0f53f012249f15fb))
* update URLs from stainlessapi.com to stainless.com ([#78](https://github.com/togethercomputer/together-go/issues/78)) ([5e7cf56](https://github.com/togethercomputer/together-go/commit/5e7cf562630a82a7577bf8356c095b3111180270))


### Refactors

* tidy up dependencies ([#84](https://github.com/togethercomputer/together-go/issues/84)) ([13c63e7](https://github.com/togethercomputer/together-go/commit/13c63e7027bb76d98c0d1f46ca7fac8af5ae00b0))

## 0.0.1-alpha.1 (2024-06-13)

Full Changelog: [v0.0.1-alpha.0...v0.0.1-alpha.1](https://github.com/togethercomputer/together-go/compare/v0.0.1-alpha.0...v0.0.1-alpha.1)

### Chores

* go live ([#1](https://github.com/togethercomputer/together-go/issues/1)) ([5a50db0](https://github.com/togethercomputer/together-go/commit/5a50db03e8dc0c1fbcb2f037a4f35ccbc23366fc))
