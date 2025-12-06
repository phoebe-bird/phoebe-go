# Changelog

## 0.1.0-alpha.12 (2025-12-06)

Full Changelog: [v0.1.0-alpha.11...v0.1.0-alpha.12](https://github.com/phoebe-bird/phoebe-go/compare/v0.1.0-alpha.11...v0.1.0-alpha.12)

### Bug Fixes

* **mcp:** correct code tool API endpoint ([83c2db3](https://github.com/phoebe-bird/phoebe-go/commit/83c2db30d5e1c8e970f18a6d8b741f0c3e1226ad))


### Chores

* elide duplicate aliases ([5412db1](https://github.com/phoebe-bird/phoebe-go/commit/5412db109804bd1900e62d743576dae5dc92f88a))
* **internal:** codegen related update ([8b9d35d](https://github.com/phoebe-bird/phoebe-go/commit/8b9d35d2a11958141d9261067fea8dda59aa9973))

## 0.1.0-alpha.11 (2025-10-08)

Full Changelog: [v0.1.0-alpha.10...v0.1.0-alpha.11](https://github.com/phoebe-bird/phoebe-go/compare/v0.1.0-alpha.10...v0.1.0-alpha.11)

### Features

* **client:** support optional json html escaping ([63dc659](https://github.com/phoebe-bird/phoebe-go/commit/63dc6598151cf469f0927251d2004d89512c11e8))


### Bug Fixes

* close body before retrying ([0b46486](https://github.com/phoebe-bird/phoebe-go/commit/0b464866bdeb7cb584dbbc57af84922b3aba710e))
* use slices.Concat instead of sometimes modifying r.Options ([9f64eea](https://github.com/phoebe-bird/phoebe-go/commit/9f64eeaade594b581b917ad84b4306fb4fdb924e))


### Chores

* bump minimum go version to 1.22 ([5c541a8](https://github.com/phoebe-bird/phoebe-go/commit/5c541a80ac94135f66eda9494d5034baf87603d1))
* do not install brew dependencies in ./scripts/bootstrap by default ([7119ddc](https://github.com/phoebe-bird/phoebe-go/commit/7119ddcb776046d73621459b39ef3db332a8982a))
* **internal:** codegen related update ([65e8f5e](https://github.com/phoebe-bird/phoebe-go/commit/65e8f5e1dd1b6dbd19a94ff6c9214a94a3875ccf))
* **internal:** update comment in script ([16dd698](https://github.com/phoebe-bird/phoebe-go/commit/16dd698779a0dc3616b8c23ff9121ae8a3f7663c))
* update @stainless-api/prism-cli to v5.15.0 ([c5a8690](https://github.com/phoebe-bird/phoebe-go/commit/c5a8690b64ab9dd48bf88b98d84d12a9669a612f))
* update more docs for 1.22 ([9ac1b9b](https://github.com/phoebe-bird/phoebe-go/commit/9ac1b9b8cb7f4c34b5d4ff28c1f06abc13c2774f))

## 0.1.0-alpha.10 (2025-07-22)

Full Changelog: [v0.1.0-alpha.9...v0.1.0-alpha.10](https://github.com/phoebe-bird/phoebe-go/compare/v0.1.0-alpha.9...v0.1.0-alpha.10)

### Bug Fixes

* **client:** process custom base url ahead of time ([51e8cc6](https://github.com/phoebe-bird/phoebe-go/commit/51e8cc6b7a82368a83266a1c43c4f689b65cd69d))


### Chores

* **ci:** only run for pushes and fork pull requests ([c576996](https://github.com/phoebe-bird/phoebe-go/commit/c576996d175efffa2a84a070cbd7b7f79d6ceeda))
* **internal:** fix lint script for tests ([f4144bc](https://github.com/phoebe-bird/phoebe-go/commit/f4144bcc82809e7ccc9e5a9c460b9da19e20f970))
* lint tests ([aec3490](https://github.com/phoebe-bird/phoebe-go/commit/aec3490aa0eb9a04a7acbe745ba9fc7941a7c984))
* lint tests in subpackages ([f12c91f](https://github.com/phoebe-bird/phoebe-go/commit/f12c91fdd90e0f2c43f0be0921000ecce85e7b23))

## 0.1.0-alpha.9 (2025-06-28)

Full Changelog: [v0.1.0-alpha.8...v0.1.0-alpha.9](https://github.com/phoebe-bird/phoebe-go/compare/v0.1.0-alpha.8...v0.1.0-alpha.9)

### Features

* **client:** add debug log helper ([bf2b23a](https://github.com/phoebe-bird/phoebe-go/commit/bf2b23a9a2484b1cf9375ec72c7d3197fcd4cbbd))
* **client:** add support for endpoint-specific base URLs in python ([4e43c63](https://github.com/phoebe-bird/phoebe-go/commit/4e43c6374ab88e96a2361542fa916cfcfe10ded7))


### Bug Fixes

* don't try to deserialize as json when ResponseBodyInto is []byte ([060e226](https://github.com/phoebe-bird/phoebe-go/commit/060e226f4a51cfb347b6ddc269db5e99b76c2e4f))


### Chores

* **ci:** enable for pull requests ([3367448](https://github.com/phoebe-bird/phoebe-go/commit/33674489f7d78540428c8a783e7505388b303e8c))
* **docs:** grammar improvements ([8d6526c](https://github.com/phoebe-bird/phoebe-go/commit/8d6526cf4aa5a2eb0c1862bdfe10d9b0a950b409))
* improve devcontainer setup ([4ec61da](https://github.com/phoebe-bird/phoebe-go/commit/4ec61da1b852b3a94ef54bbea8efcf7cdc9fd2fe))
* make go mod tidy continue on error ([19bca49](https://github.com/phoebe-bird/phoebe-go/commit/19bca499ff93549ca78d96a4301232945ee9a06b))

## 0.1.0-alpha.8 (2025-05-07)

Full Changelog: [v0.1.0-alpha.7...v0.1.0-alpha.8](https://github.com/phoebe-bird/phoebe-go/compare/v0.1.0-alpha.7...v0.1.0-alpha.8)

### Features

* **client:** add support for reading base URL from environment variable ([f67494c](https://github.com/phoebe-bird/phoebe-go/commit/f67494cf48669813b51f6d079f6718c88d30626c))
* **client:** support custom http clients ([#47](https://github.com/phoebe-bird/phoebe-go/issues/47)) ([b77ae0e](https://github.com/phoebe-bird/phoebe-go/commit/b77ae0e3dd3a2ea5d400dd29fb401864e0248740))


### Bug Fixes

* **client:** clean up reader resources ([1e9f142](https://github.com/phoebe-bird/phoebe-go/commit/1e9f14249f327c917dfa3e43de292c72f4847ec4))
* **client:** correctly update body in WithJSONSet ([670e890](https://github.com/phoebe-bird/phoebe-go/commit/670e8904823a8f16c56cac873e3c445b9b0db2c7))
* handle empty bodies in WithJSONSet ([406fefe](https://github.com/phoebe-bird/phoebe-go/commit/406fefed024ce8861dd8930762a322934aae051c))


### Chores

* **ci:** add timeout thresholds for CI jobs ([0cefe8b](https://github.com/phoebe-bird/phoebe-go/commit/0cefe8be441875fc67834dad43327014ea50453c))
* **ci:** only use depot for staging repos ([aba413a](https://github.com/phoebe-bird/phoebe-go/commit/aba413a1cdbb40b9d9161b3f4ec66ef4876b23ab))
* **docs:** document pre-request options ([fb354a6](https://github.com/phoebe-bird/phoebe-go/commit/fb354a6cc6b0d238e18b4126742415c61be49022))
* **internal:** codegen related update ([3824665](https://github.com/phoebe-bird/phoebe-go/commit/3824665a72a7dd9d84ccc803390f60836127d30b))
* **internal:** expand CI branch coverage ([93f0ad0](https://github.com/phoebe-bird/phoebe-go/commit/93f0ad04ffe6574a27a4a3972a1b5502e6b12af3))
* **internal:** reduce CI branch coverage ([7ec6263](https://github.com/phoebe-bird/phoebe-go/commit/7ec6263947a764af02068ae1d1d4fcb86d78afda))


### Documentation

* update documentation links to be more uniform ([21d4f72](https://github.com/phoebe-bird/phoebe-go/commit/21d4f72732350137d269c0114ca906c2c7df18a8))

## 0.1.0-alpha.7 (2025-04-03)

Full Changelog: [v0.1.0-alpha.6...v0.1.0-alpha.7](https://github.com/phoebe-bird/phoebe-go/compare/v0.1.0-alpha.6...v0.1.0-alpha.7)

### Bug Fixes

* **client:** return error on bad custom url instead of panic ([#45](https://github.com/phoebe-bird/phoebe-go/issues/45)) ([d60b6dc](https://github.com/phoebe-bird/phoebe-go/commit/d60b6dcb55c19ec28f36bbb1c7339df49cac3013))


### Chores

* fix typos ([#43](https://github.com/phoebe-bird/phoebe-go/issues/43)) ([4825a72](https://github.com/phoebe-bird/phoebe-go/commit/4825a72992c7ba48f99363dd3443a05ca51c3ecd))

## 0.1.0-alpha.6 (2025-03-26)

Full Changelog: [v0.1.0-alpha.5...v0.1.0-alpha.6](https://github.com/phoebe-bird/phoebe-go/compare/v0.1.0-alpha.5...v0.1.0-alpha.6)

### Features

* **client:** allow custom baseurls without trailing slash ([#34](https://github.com/phoebe-bird/phoebe-go/issues/34)) ([ebe2ecf](https://github.com/phoebe-bird/phoebe-go/commit/ebe2ecf8641ac6475989f960345bec02f4600ac9))
* **client:** improve default client options support ([#37](https://github.com/phoebe-bird/phoebe-go/issues/37)) ([96e15fd](https://github.com/phoebe-bird/phoebe-go/commit/96e15fda795f088e025c6cbdb069c295afe1ead6))


### Bug Fixes

* **test:** return early after test failure ([#41](https://github.com/phoebe-bird/phoebe-go/issues/41)) ([7c5be5c](https://github.com/phoebe-bird/phoebe-go/commit/7c5be5caa15b695f115d1ce3a9fe82d5139d6faf))


### Chores

* add request options to client tests ([#40](https://github.com/phoebe-bird/phoebe-go/issues/40)) ([5a369dd](https://github.com/phoebe-bird/phoebe-go/commit/5a369dd3481315dd97690d2253509da164c6fe57))
* **docs:** improve security documentation ([#39](https://github.com/phoebe-bird/phoebe-go/issues/39)) ([ec2e4c1](https://github.com/phoebe-bird/phoebe-go/commit/ec2e4c167ecc6e343eab90f9dcd896b5aa6451bb))
* **internal:** codegen related update ([#32](https://github.com/phoebe-bird/phoebe-go/issues/32)) ([7f2cb99](https://github.com/phoebe-bird/phoebe-go/commit/7f2cb9909324d2225b30d29e2ab968b51a8a15c5))
* **internal:** codegen related update ([#35](https://github.com/phoebe-bird/phoebe-go/issues/35)) ([b7a680e](https://github.com/phoebe-bird/phoebe-go/commit/b7a680eeae6d93cd62556a6042f105f55ccc9b55))
* **internal:** remove extra empty newlines ([#38](https://github.com/phoebe-bird/phoebe-go/issues/38)) ([73e31ed](https://github.com/phoebe-bird/phoebe-go/commit/73e31eda4deef36a0c320a3f8a558a14aaac8f43))


### Refactors

* tidy up dependencies ([#36](https://github.com/phoebe-bird/phoebe-go/issues/36)) ([16815a1](https://github.com/phoebe-bird/phoebe-go/commit/16815a1a4fbb3be9fa4d9242001f4d81ec9c141f))

## 0.1.0-alpha.5 (2025-02-01)

Full Changelog: [v0.1.0-alpha.4...v0.1.0-alpha.5](https://github.com/phoebe-bird/phoebe-go/compare/v0.1.0-alpha.4...v0.1.0-alpha.5)

### Bug Fixes

* fix unicode encoding for json ([#28](https://github.com/phoebe-bird/phoebe-go/issues/28)) ([276534f](https://github.com/phoebe-bird/phoebe-go/commit/276534f829f813e5a046a2079221b7d57548c87a))

## 0.1.0-alpha.4 (2025-01-29)

Full Changelog: [v0.1.0-alpha.3...v0.1.0-alpha.4](https://github.com/phoebe-bird/phoebe-go/compare/v0.1.0-alpha.3...v0.1.0-alpha.4)

### Bug Fixes

* fix apijson.Port for embedded structs ([#24](https://github.com/phoebe-bird/phoebe-go/issues/24)) ([95355e3](https://github.com/phoebe-bird/phoebe-go/commit/95355e3890c4d9365db944f676c0bb0774b4927a))
* fix apijson.Port for embedded structs ([#25](https://github.com/phoebe-bird/phoebe-go/issues/25)) ([aca76ce](https://github.com/phoebe-bird/phoebe-go/commit/aca76cebb3a8d0a497d7f9e92129656f4ee7998f))


### Chores

* bump license year ([#22](https://github.com/phoebe-bird/phoebe-go/issues/22)) ([51abcdf](https://github.com/phoebe-bird/phoebe-go/commit/51abcdf3b4d7b5f504d9eb65f25a39a5357b605b))
* **internal:** codegen related update ([#19](https://github.com/phoebe-bird/phoebe-go/issues/19)) ([1937370](https://github.com/phoebe-bird/phoebe-go/commit/193737065aaffb9282bf768abd13e85900f5b5de))
* **internal:** codegen related update ([#21](https://github.com/phoebe-bird/phoebe-go/issues/21)) ([39b5291](https://github.com/phoebe-bird/phoebe-go/commit/39b5291c45312a0fa34489c81d823230b8f9a1e9))
* **internal:** codegen related update ([#23](https://github.com/phoebe-bird/phoebe-go/issues/23)) ([1044bec](https://github.com/phoebe-bird/phoebe-go/commit/1044bec365c34746c7426e3204b4859f3b9794c8))
* refactor client tests ([#26](https://github.com/phoebe-bird/phoebe-go/issues/26)) ([fdb33b1](https://github.com/phoebe-bird/phoebe-go/commit/fdb33b1d759e11df6938a8e8a1723be470d9f76e))

## 0.1.0-alpha.3 (2024-11-12)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/phoebe-bird/phoebe-go/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Bug Fixes

* **api:** get repo clean ([#13](https://github.com/phoebe-bird/phoebe-go/issues/13)) ([0a7a20e](https://github.com/phoebe-bird/phoebe-go/commit/0a7a20e6d5d23c12d864bae7d5940248ea732feb))


### Chores

* rebuild project due to codegen change ([#16](https://github.com/phoebe-bird/phoebe-go/issues/16)) ([7ce37c8](https://github.com/phoebe-bird/phoebe-go/commit/7ce37c8deb76388d9d1b9684113bb56018986a72))
* rebuild project due to codegen change ([#17](https://github.com/phoebe-bird/phoebe-go/issues/17)) ([7a3d5fc](https://github.com/phoebe-bird/phoebe-go/commit/7a3d5fcc1f4d2dc89dcac24980ad9fb29cdc86d6))

## 0.1.0-alpha.2 (2024-07-07)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/phoebe-bird/phoebe-go/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Features

* **api:** add docs to openapi spec ([#11](https://github.com/phoebe-bird/phoebe-go/issues/11)) ([4919966](https://github.com/phoebe-bird/phoebe-go/commit/491996613ea742452effdcaf71559a9cb8764d33))
* **api:** fix adjacent region ([#9](https://github.com/phoebe-bird/phoebe-go/issues/9)) ([6069a29](https://github.com/phoebe-bird/phoebe-go/commit/6069a29d06680517bb85780b77b0d29b027da8d0))
* **docs:** update contact email ([#12](https://github.com/phoebe-bird/phoebe-go/issues/12)) ([dc35b21](https://github.com/phoebe-bird/phoebe-go/commit/dc35b21f7b73ed5d69bc3e67ceced7e8916b56b3))

## 0.1.0-alpha.1 (2024-07-07)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/phoebe-bird/phoebe-go/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* **api:** manual updates ([2f10d1c](https://github.com/phoebe-bird/phoebe-go/commit/2f10d1c506b76ac4f6950ea1fe112f9799a93e09))
* **api:** manual updates ([053b4b6](https://github.com/phoebe-bird/phoebe-go/commit/053b4b62603c63d937d5edc370023fb9e1b3141b))
* **api:** manual updates ([30e6118](https://github.com/phoebe-bird/phoebe-go/commit/30e61186d3bf277ad304fcb9f49d19f95ef565a5))
* **api:** manual updates ([0d7593d](https://github.com/phoebe-bird/phoebe-go/commit/0d7593d64fe01a365de7f31280f00147782c114e))
* **api:** manual updates ([#3](https://github.com/phoebe-bird/phoebe-go/issues/3)) ([e0ca83b](https://github.com/phoebe-bird/phoebe-go/commit/e0ca83bc08cc1d429155d94eb28310b1850b6b28))
* **api:** manual updates ([#4](https://github.com/phoebe-bird/phoebe-go/issues/4)) ([d803536](https://github.com/phoebe-bird/phoebe-go/commit/d803536ff44dadcbbce7f5c6c0c077e1e63a9c56))
* **api:** rename package back ([#6](https://github.com/phoebe-bird/phoebe-go/issues/6)) ([9a0a6e3](https://github.com/phoebe-bird/phoebe-go/commit/9a0a6e3b76346f5f4b2cd5da11a5db589dba238c))


### Chores

* configure new SDK language ([45da456](https://github.com/phoebe-bird/phoebe-go/commit/45da456f275e972a7d4ea434c0f54d3b6f3fbb0d))
* go live ([#1](https://github.com/phoebe-bird/phoebe-go/issues/1)) ([46f02c2](https://github.com/phoebe-bird/phoebe-go/commit/46f02c2612f8c44005ed20de53e9e943cc598fc8))
