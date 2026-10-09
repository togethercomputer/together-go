# Changelog

## [0.13.0](https://github.com/togethercomputer/together-go/compare/v0.12.0...v0.13.0) (2026-10-09)


### ⚠ BREAKING CHANGES

* **rl:** collapse forward into forward-backward

### Features

* add canonical NVIDIA version selector ([1ca46d2](https://github.com/togethercomputer/together-go/commit/1ca46d280030ca626d39d6f3d06bde7cc0578c13))
* add deployment capacity type ([a6f2589](https://github.com/togethercomputer/together-go/commit/a6f258968fde5bebb54150b8fdddf19e1a1094cd))
* add DPPO loss to RL OpenAPI ([09487c5](https://github.com/togethercomputer/together-go/commit/09487c52c234685b6ef9838adf24cd3f4f413128))
* add rl checkpoint list endpoint ([115bc68](https://github.com/togethercomputer/together-go/commit/115bc6854e481d7335d653ace411adb9788cd7eb))
* add RL model resource queue position ([6013cbb](https://github.com/togethercomputer/together-go/commit/6013cbbfcd73ff942022018a5310a0ff5f8735c1))
* Add Rollouts SDK code for endpoint deployments management ([39214e5](https://github.com/togethercomputer/together-go/commit/39214e584202ed0757211f489e15a342904a7fcc))
* **clusters:** Promote GPU clusters off beta! ([0924a58](https://github.com/togethercomputer/together-go/commit/0924a58b797c0324ba5750dbf5ac4aa7d845793a))
* **containers:** expose deployment model mounts field ([e4e64ea](https://github.com/togethercomputer/together-go/commit/e4e64eacff5035a7f52f7dc8cfa0b2a5b17cc027))
* **deployments:** add deployment compliance policy for hipaa requirements ([6d34197](https://github.com/togethercomputer/together-go/commit/6d341978e7e07f59c00707e3a582af6f6f9d57d8))
* **deployments:** Add inactiveTimeout to deployments ([#644](https://github.com/togethercomputer/together-go/issues/644)) ([3b49f97](https://github.com/togethercomputer/together-go/commit/3b49f9742e0945cd392980cbd0e07596b3855602))
* **deployments:** expose autoscaling policies and reset semantics (MLE-7727) ([151b622](https://github.com/togethercomputer/together-go/commit/151b62202ee9124ec5f3a6ad23a4c412165cf8ba))
* document shaping calibration files ([ba9cb5f](https://github.com/togethercomputer/together-go/commit/ba9cb5feb88e45b05d0483d300a9baa70b456a82))
* **endpoint:** Add adapter validation status fields ([9d86afa](https://github.com/togethercomputer/together-go/commit/9d86afa203a85f1f93a8d72a13fbff2e22907718))
* **endpoints:** add supported model adapter mode field ([18c22f2](https://github.com/togethercomputer/together-go/commit/18c22f2c84ca2ea471d70a44dc8409f184ec0532))
* **endpoints:** expose deployment concurrency limit response ([3e8937f](https://github.com/togethercomputer/together-go/commit/3e8937f12e4e71be526d82cd80214c0e07efe0b0))
* **endpoints:** List enum values for endpoint scaling metric values ([ff0435e](https://github.com/togethercomputer/together-go/commit/ff0435ee3b5837ad6d100e394a9c1985edf2bf2f))
* expose RL checkpoint registry artifacts ([04957a3](https://github.com/togethercomputer/together-go/commit/04957a3fe49ab23eef2130e6f3c77d413ba458ea))
* expose RL checkpoint registry ids in OpenAPI ([21d2485](https://github.com/togethercomputer/together-go/commit/21d248507a5d0bc7087aa43a0f92969d27c800ed))
* expose RL GPU configurations in OpenAPI ([c347bc7](https://github.com/togethercomputer/together-go/commit/c347bc712b798576c17161b43d3997659e3f5f90))
* expose RL session policy state ([b4a94ee](https://github.com/togethercomputer/together-go/commit/b4a94ee018adff53a936d5566f5d89c11d3880ea))
* **fine-tuning:** sync tokenized dataset download OpenAPI ([1868d94](https://github.com/togethercomputer/together-go/commit/1868d94e3c21bddd2bb4f9932a38ee66e2c1139e))
* **jig:** add deployment revision and rollback APIs ([afa0a0a](https://github.com/togethercomputer/together-go/commit/afa0a0a9473549753e731d585dda20517e707314))
* **jig:** add S3 volume origins ([f5a740e](https://github.com/togethercomputer/together-go/commit/f5a740e6d67c205a0bb057aeada4f0a9d1a68e81))
* **openapi:** add B300 GPU cluster type ([1bd70a4](https://github.com/togethercomputer/together-go/commit/1bd70a4c5f4ac08cb461c4af3d63ca3af071f893))
* **openapi:** add RL LoRA train_unembed ([cf8dac5](https://github.com/togethercomputer/together-go/commit/cf8dac5f4901aa8f5e4f18b6243af7bc0741c062))
* **openapi:** add RL prompt cache hit tokens ([ad60dc8](https://github.com/togethercomputer/together-go/commit/ad60dc8dcdfd675e51d8140847a2f92d846002cb))
* **openapi:** add RL routing replay fields ([b6dfcc0](https://github.com/togethercomputer/together-go/commit/b6dfcc05cab58a92bc35c06a860a77fb0c7589b3))
* **openapi:** add RL weights sync operation ([10798ca](https://github.com/togethercomputer/together-go/commit/10798ca3cbb69e33546a9d8bd7d5922cad40cc2f))
* **openapi:** expose deployment adapter id ([e229df9](https://github.com/togethercomputer/together-go/commit/e229df90760e9471bec9cad13acc68710295995a))
* **openapi:** sync reserved endpoint enums ([f1ac92b](https://github.com/togethercomputer/together-go/commit/f1ac92b4950f931f469554b0173a7ff067ac2cbe))
* **projects:** add public list projects endpoint ([2735a3e](https://github.com/togethercomputer/together-go/commit/2735a3e18d87630867de3ba18ca4bb9474423ab0))
* **rl:** add base_weights_ref to model resources (MOSH-5159) ([f60f960](https://github.com/togethercomputer/together-go/commit/f60f96097eac8f055ba3e599062db4f55df8a3d1))
* **rl:** add forward-backward loss function outputs ([f6a2ba4](https://github.com/togethercomputer/together-go/commit/f6a2ba48bfc63f5598b86023b3a83ec8ba6d269a))
* **rl:** Add SDK methods for the new Reinforcement Learning product ([8cac559](https://github.com/togethercomputer/together-go/commit/8cac559c353ac8ef9428061d2793ac90fe11eb89))
* **rl:** collapse forward into forward-backward ([ce3335f](https://github.com/togethercomputer/together-go/commit/ce3335f4c75bf95f7d01d4ccbc5f046f45d243d6))
* **rl:** expose checkpoint type filter ([b1033a5](https://github.com/togethercomputer/together-go/commit/b1033a5d4e9ac210bebe03542b55e4f276dedd1c))
* **rl:** remove resume_from_hf_checkpoint from training sessions (MOSH-5276) ([60bb889](https://github.com/togethercomputer/together-go/commit/60bb889fc6ebc06ebbf8e5c8c604f5737d463a09))
* **rollout:** add PreviewRolloutDefaults endpoint ([e04d032](https://github.com/togethercomputer/together-go/commit/e04d0328ea2cbc0379e889784d40bdb9efe10577))
* **supported-models:** expose serverless pricing ([dc682d6](https://github.com/togethercomputer/together-go/commit/dc682d6377f15586219737c5c8da571cc484a0be))
* sync passive health check alert ordering ([f8f824a](https://github.com/togethercomputer/together-go/commit/f8f824a8ed4411a851e424cf8334e16be2e8879f))
* sync rollout partial capacity status ([ffca4d8](https://github.com/togethercomputer/together-go/commit/ffca4d8a84ed7c0b96e78ec69bfd2b4dd3c6a242))
* sync rollout pausing OpenAPI spec ([3156f3a](https://github.com/togethercomputer/together-go/commit/3156f3a0b33de97c03c51fe8c91e833939386c12))
* sync rollout preview landing band fields ([6164cd4](https://github.com/togethercomputer/together-go/commit/6164cd45a06bc9134c27209dd71766be1b3e7242))
* sync rollout step states in OpenAPI ([39c22f2](https://github.com/togethercomputer/together-go/commit/39c22f2a7c2f02dd783ee1ea97d3f5c63177c643))
* sync shared volume cluster pinning ([780e821](https://github.com/togethercomputer/together-go/commit/780e8217866679a6f35c27f32a3276547f27de90))


### Bug Fixes

* **clusters:** Mark id/os as required properties ([b14421c](https://github.com/togethercomputer/together-go/commit/b14421cf71ef4d8f05922278ecf1cc5f57f243c9))
* **endpoints:** Drop unsupported `enableLora` parameter from deployment schemas ([d75ede1](https://github.com/togethercomputer/together-go/commit/d75ede11b355f461c5ba1b293e1cca52f6c87651))
* **openapi:** add request examples for CreateRollout so generated samples are valid ([8b74cc3](https://github.com/togethercomputer/together-go/commit/8b74cc378b21d3bebcc6e3c4da4f43c7e8ad0114))
* **openapi:** add RL non-finite loss error code ([a53bf95](https://github.com/togethercomputer/together-go/commit/a53bf9506f98fa46de2041a43d7fd00c513b4938))
* **openapi:** align RL optimizer config names ([5f5c456](https://github.com/togethercomputer/together-go/commit/5f5c456f329afbe137629c2bc0aaa8138e0845b7))
* **openapi:** remove RL training sample policy segments ([bdd6d17](https://github.com/togethercomputer/together-go/commit/bdd6d178c1d63542490b6c516ac70ccd52c7ea42))
* **openapi:** sync RL prompt logprobs schema ([5d71e1b](https://github.com/togethercomputer/together-go/commit/5d71e1b187ae8c9e674faa2f62d5a67c3e22c5ef))
* **openapi:** sync RL Tinker parameter renames ([17cc6c6](https://github.com/togethercomputer/together-go/commit/17cc6c69181a51919bfa1174f25cc18c72787b3b))
* **openapi:** sync rollout gate optional values ([f3ddb35](https://github.com/togethercomputer/together-go/commit/f3ddb35379c881e9cf4834ba670f6d60f9b7e067))
* **stlc:** re-anchor custom-code tracking to staging main; seal partial builds ([aebc295](https://github.com/togethercomputer/together-go/commit/aebc2950ff3168059d072b08e505688ba32f0e9b))
* sync adapter attachment ids in OpenAPI ([e291079](https://github.com/togethercomputer/together-go/commit/e291079f505bc139f39dffde0182e5defdfd00d5))
* sync RL checkpoint OpenAPI schemas ([54a893b](https://github.com/togethercomputer/together-go/commit/54a893b81f241f7d7526bf57ebd0d9db7ef1ca61))
* use valid metric in rollout preview examples ([5a3c32b](https://github.com/togethercomputer/together-go/commit/5a3c32b795fdb8d68d337b20ed11631ec640a401))


### Chores

* Remove generated beta code for clusters and integrate SDK level shim ([211377f](https://github.com/togethercomputer/together-go/commit/211377fd9e2a36ca1a581f7d12d47594949af58d))
* update CI scripts for syncing ([#17](https://github.com/togethercomputer/together-go/issues/17)) ([944ab82](https://github.com/togethercomputer/together-go/commit/944ab82b65352d93be82388e4dc29573f8067af4))


### Documentation

* **audio:** correct direct-upload limit to 80 MB in descriptions ([bc25990](https://github.com/togethercomputer/together-go/commit/bc25990f0f1490fd0b7e9d2ab2f205a60d93d78c))
* clarify deployment desired replicas ownership ([81d29cb](https://github.com/togethercomputer/together-go/commit/81d29cb0de3c3ca980b4e57018ad0af3b7d91e75))
* **endpoints:** Update rolling rollout description ([54f0d4f](https://github.com/togethercomputer/together-go/commit/54f0d4f9d2e42596bb415c1f1fee754bfc46c280))
* **files:** document file upload filename limits ([b15ed02](https://github.com/togethercomputer/together-go/commit/b15ed02e00cf0a999d88243837d9420f053fc5c9))
* **finetuning:** document fine-tune teardown response ([752d08e](https://github.com/togethercomputer/together-go/commit/752d08e4ad9d75d986f7603d4d4ba79b162f46d3))
* **openapi:** clarify expert LoRA merge output ([cf66e1b](https://github.com/togethercomputer/together-go/commit/cf66e1b48be625b4b6aa60746b6bbe9d41e872d1))
* **openapi:** sync endpoint active rollout docs ([70b4156](https://github.com/togethercomputer/together-go/commit/70b41569444dd81292e1d98e48dfa058f93d536a))
* **openapi:** sync rollout create defaults ([f19638a](https://github.com/togethercomputer/together-go/commit/f19638af9b89313c18043478dc2db9b2541f3f18))
* **openapi:** sync rollout landing floor descriptions ([87e8c4d](https://github.com/togethercomputer/together-go/commit/87e8c4d368cbbaa2eb761fd678178d2d9ebb9037))
* **rl:** document model resource quota rejection error code ([fb97095](https://github.com/togethercomputer/together-go/commit/fb97095b8bd7c67c49633babb4ad62e3c8db46a6))
* sync rollout autoscaling preview docs ([26a3ca4](https://github.com/togethercomputer/together-go/commit/26a3ca4804bbdcb4947c17b5208ed4ec1f12de26))
* sync rollout final target semantics ([c204723](https://github.com/togethercomputer/together-go/commit/c204723e8749ce4edaee89e758974d0d398a4892))
* sync rollout landing ceiling OpenAPI description ([0575bfa](https://github.com/togethercomputer/together-go/commit/0575bfa308781d3b765b1571a4cf3d67914c6e1c))
* sync rollout metric gate OpenAPI docs ([5b93e71](https://github.com/togethercomputer/together-go/commit/5b93e714c7bc0a01fe5dfb5c857b246a26dd1857))
* sync rollout OpenAPI defaults ([02f07da](https://github.com/togethercomputer/together-go/commit/02f07da256f629f63967b74181119e37afbefb0d))
* sync shadow target rollout guard ([b63d5c1](https://github.com/togethercomputer/together-go/commit/b63d5c1014555669a248b6d5cd49ec77b742e9ff))
* sync supported model profile name ([7d6c579](https://github.com/togethercomputer/together-go/commit/7d6c57996fd3ac3e2fa607780de0973de779bfb2))
