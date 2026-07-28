# Changelog

## [0.12.0](https://github.com/togethercomputer/together-go/compare/v0.11.0...v0.12.0) (2026-07-28)


### Features

* add fine-tunes preview endpoint ([751e4c5](https://github.com/togethercomputer/together-go/commit/751e4c512e7092516fd47ac3b6806993588693c0))
* add headlamp add-on to gpu cluster ([96fe8ea](https://github.com/togethercomputer/together-go/commit/96fe8ead31b7412efd92d421e69d147d15c3816b))
* add Slurm Web addon to gpu cluster ([89bbf1e](https://github.com/togethercomputer/together-go/commit/89bbf1eb0a242e6433f12bee0a472b9f193d441d))


### Bug Fixes

* **openapi:** remove rollout abort route ([d7b4717](https://github.com/togethercomputer/together-go/commit/d7b4717927af7c894af1b040f10db326208fab27))
* **openapi:** sync RL model input chunk schema ([747f364](https://github.com/togethercomputer/together-go/commit/747f364a9205aa13744602a29425b715ef31ecd5))
* **openapi:** sync RL sampling contract ([452c599](https://github.com/togethercomputer/together-go/commit/452c5995227ba2ba6333cd4221197ba5367d9c4b))
* Revert server host change for chat completions inference ([01aa25e](https://github.com/togethercomputer/together-go/commit/01aa25ee7e6c9582da16fa363d037fac41ad7ae7))
* **rl:** align CISPO/DRO loss schema names with GRPO/PPO (all-caps acronyms) ([7329d72](https://github.com/togethercomputer/together-go/commit/7329d72480cb42a1f46ec301552572f498cc50a2))
* **stlc:** generate cross-entropy loss params type ([a647049](https://github.com/togethercomputer/together-go/commit/a6470498639cbb3085222f3f1aa365471c3350a1))
* Update types for deployment summary to acknowledge optionality on hardware property ([8b56f2c](https://github.com/togethercomputer/together-go/commit/8b56f2cab31206c80d35be9ca28d1ad9b663b7d5))


### Documentation

* **openapi:** allow combined LoRA target modules ([7afe744](https://github.com/togethercomputer/together-go/commit/7afe7442f1722da23b1412e9c72486c22b89a7b7))
* **openapi:** clarify fine-tune user id creator ([c35b741](https://github.com/togethercomputer/together-go/commit/c35b7410577efd1778208ba3f1c49b829fc20e33))
* **openapi:** clarify RL loss mask optionality ([8d123a5](https://github.com/togethercomputer/together-go/commit/8d123a50cde44e63728945f3e9939de71f56fde7))
* sync endpoint events limit ([01b3e00](https://github.com/togethercomputer/together-go/commit/01b3e00e1f7b4d3563de973682a2346b49ff13f4))


### Refactors

* generate better code ([263bd0b](https://github.com/togethercomputer/together-go/commit/263bd0b4f44cba39b299e8c35ea6054c1a43ee43))

## [0.11.0](https://github.com/togethercomputer/together-go/compare/v0.10.0...v0.11.0) (2026-07-16)


### Features

* add /v1/whoami endpoint to OpenAPI spec ([59a7915](https://github.com/togethercomputer/together-go/commit/59a7915a8ef9deeb9002f238dc42572d83d2ff51))
* expose GPU cluster reserved counts ([00de502](https://github.com/togethercomputer/together-go/commit/00de502d472f984309c0e2690a07270212cbee2d))
* **openapi:** add force stop parameter ([0ee74ae](https://github.com/togethercomputer/together-go/commit/0ee74aeae6c88c2ecb26668b05bdac3f057a69db))
* **openapi:** expose remediation linked alerts ([bbd0246](https://github.com/togethercomputer/together-go/commit/bbd02465fdd100b44fc0f9d82aedfd9238713b36))
* **openapi:** sync passive health check alerts ([cfbd257](https://github.com/togethercomputer/together-go/commit/cfbd257a50f93c9a1821d5215530f881cd91afca))
* **openapi:** sync RL model resources stop API ([562ead3](https://github.com/togethercomputer/together-go/commit/562ead3c2f8581ac805fefba09e2e500a04247ef))
* **openapi:** sync RL Muon optimizer API ([fcce400](https://github.com/togethercomputer/together-go/commit/fcce400d264f6213c7b341d836365f00539d7f9c))
* **openapi:** sync SSH CA cluster config ([7a44c42](https://github.com/togethercomputer/together-go/commit/7a44c42ccb6225360e3dfe609b44e24a71b22934))
* SDK methods for new dedicated models inference ([3f6bb70](https://github.com/togethercomputer/together-go/commit/3f6bb70092b31e7631760a900dcf830a6c13621b))


### Bug Fixes

* **openapi:** remove duplicate supports_full_training key ([ceac93e](https://github.com/togethercomputer/together-go/commit/ceac93e7a92cf25487f3fefba8fa85dfda1f6034))


### Chores

* Add staging CI syncing ([#2](https://github.com/togethercomputer/together-go/issues/2)) ([266398b](https://github.com/togethercomputer/together-go/commit/266398bc616a6121fa6ba345668cf7926e165288))
* Add stlc promote action ([265b4b4](https://github.com/togethercomputer/together-go/commit/265b4b43dc3e060c4846a2827b78cf1c560f7f73))
* export finetune model limits type name preoperly ([5b7c710](https://github.com/togethercomputer/together-go/commit/5b7c710f0fe7b375776ba65f10a9fca76d7f195f))
* fix custom code checkin ([4118b8c](https://github.com/togethercomputer/together-go/commit/4118b8c8796892a3f2fdad2cdf4cd3eaab0d2367))
* fix production repo reference ([86d52a5](https://github.com/togethercomputer/together-go/commit/86d52a54dc0f32baaf1414c591ce530285f951ce))
* fix yaml ([35b390b](https://github.com/togethercomputer/together-go/commit/35b390b98e1b941f1c708589690620a1c10c536b))
* Improve summary docs for remediations ([d24061c](https://github.com/togethercomputer/together-go/commit/d24061ccd556bedbab3ca9be846a972331f5a72f))
* integrate production changes to staging repo ([54a085c](https://github.com/togethercomputer/together-go/commit/54a085c0d33785c133555b14c7161709d22789d0))
* stlc integrate ([7d116fe](https://github.com/togethercomputer/together-go/commit/7d116fe2f7685bbbe982a6fae96d5ca9994ba688))
* sync custom code ([65fed74](https://github.com/togethercomputer/together-go/commit/65fed74d6bfef180673e00e70efcaff756182d5d))
* sync SDKs via stlc ([c976b21](https://github.com/togethercomputer/together-go/commit/c976b216b6e6209ea34fe861116acbe852e72cd8))
* update scripts to use github app ([aeaf98b](https://github.com/togethercomputer/together-go/commit/aeaf98b21a4e64afea88adbb19a4841fe90c8863))


### Documentation

* **adapters:** align endpoint adapter code samples with the SDK ([a717b63](https://github.com/togethercomputer/together-go/commit/a717b63dd512996be6b116e4327f6c2f524ca83f))
* expose fine-tune artifact ids ([3402308](https://github.com/togethercomputer/together-go/commit/3402308161f91b42e3411ee027b746ec1f224674))
