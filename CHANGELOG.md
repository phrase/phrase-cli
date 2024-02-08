# Changelog

## [2.22.0](https://github.com/phrase/openapi/compare/cli-v2.21.2...cli-v2.22.0) (2024-02-08)


### Features

* add query param for properties ([#542](https://github.com/phrase/openapi/issues/542)) ([b4e12d0](https://github.com/phrase/openapi/commit/b4e12d04fd2916351f9201e1e6de504143ecc9aa))

## [2.21.2](https://github.com/phrase/openapi/compare/cli-v2.21.1...cli-v2.21.2) (2024-02-05)


### Bug Fixes

* **API:** allow nullable value for job's due_date ([#534](https://github.com/phrase/openapi/issues/534)) ([38b51b5](https://github.com/phrase/openapi/commit/38b51b51095394f8ce769873140038abba628514))

## [2.21.1](https://github.com/phrase/openapi/compare/cli-v2.21.0...cli-v2.21.1) (2024-02-01)


### Bug Fixes

* **API:** Adjust documentation of QPS endpoint ([#525](https://github.com/phrase/openapi/issues/525)) ([4b4f1ac](https://github.com/phrase/openapi/commit/4b4f1acf28fbd13b3d16c37162cdccfa05c38ffa))

## [2.21.0](https://github.com/phrase/openapi/compare/cli-v2.20.0...cli-v2.21.0) (2024-01-17)


### Features

* **api:** Add QPS endpoint and documentation ([#521](https://github.com/phrase/openapi/issues/521)) ([d22c558](https://github.com/phrase/openapi/commit/d22c558adfbb7fcd13759e388c038744914e42fa))

## [2.20.0](https://github.com/phrase/openapi/compare/cli-v2.19.3...cli-v2.20.0) (2024-01-10)


### Features

* **CLI:** Optional cleanup after push [TSI-2234] ([#518](https://github.com/phrase/openapi/issues/518)) ([f3730a1](https://github.com/phrase/openapi/commit/f3730a1ab8e1377aa5f4f2793b0bd72c3dab1664))

## [2.19.3](https://github.com/phrase/openapi/compare/cli-v2.19.2...cli-v2.19.3) (2024-01-02)


### Bug Fixes

* **CLI:** retrigger release for GO ([fca7edb](https://github.com/phrase/openapi/commit/fca7edb2bfecaeb2ed92c5f50acb8d820ef94cb0))
* **GO:** fix type mismatch error ([fca7edb](https://github.com/phrase/openapi/commit/fca7edb2bfecaeb2ed92c5f50acb8d820ef94cb0))
* **GO:** fix type mismatch for nested args ([#515](https://github.com/phrase/openapi/issues/515)) ([fca7edb](https://github.com/phrase/openapi/commit/fca7edb2bfecaeb2ed92c5f50acb8d820ef94cb0))

## [2.19.2](https://github.com/phrase/openapi/compare/cli-v2.19.1...cli-v2.19.2) (2024-01-02)


### Bug Fixes

* **API:** Create Custom Metadata endpoint fix [TSI-2222] ([#499](https://github.com/phrase/openapi/issues/499)) ([ce2ed94](https://github.com/phrase/openapi/commit/ce2ed9488e111fb5d9bc3810a78c47d23553c8b7))

## [2.19.1](https://github.com/phrase/openapi/compare/cli-v2.19.0...cli-v2.19.1) (2023-12-21)


### Bug Fixes

* **CLI:** trigger new release ([#508](https://github.com/phrase/openapi/issues/508)) ([162eae3](https://github.com/phrase/openapi/commit/162eae3d941684e3708091ec7aeee816ca06e3d5))

## [2.19.0](https://github.com/phrase/openapi/compare/cli-v2.18.0...cli-v2.19.0) (2023-12-15)


### Features

* **CLI:** Improve push error display [TSI-1736] ([#491](https://github.com/phrase/openapi/issues/491)) ([555c017](https://github.com/phrase/openapi/commit/555c01712eec53aaf25fbe075359c8411a64eb7c))

## [2.18.0](https://github.com/phrase/openapi/compare/cli-v2.17.0...cli-v2.18.0) (2023-12-14)


### Features

* **CLI:** On push, print upload URL [TSI-1735] ([#485](https://github.com/phrase/openapi/issues/485)) ([cbdc8ed](https://github.com/phrase/openapi/commit/cbdc8ed12217cb775d41faa662b9890050ad5a4e))

## [2.17.0](https://github.com/phrase/openapi/compare/cli-v2.16.0...cli-v2.17.0) (2023-12-13)


### Features

* **API:** add Custom Metadata endpoints ([#474](https://github.com/phrase/openapi/issues/474)) ([d407d8b](https://github.com/phrase/openapi/commit/d407d8be5ccddec1afde14a12804a7a616f77d7a))
* **API:** Add custom_metadata_filters param to locale download endpoint [TSI-2174] ([#478](https://github.com/phrase/openapi/issues/478)) ([3623478](https://github.com/phrase/openapi/commit/3623478fc1518b457ab018b5630a693081637d6e))
* **API:** Add url field to uploads ([#481](https://github.com/phrase/openapi/issues/481)) ([7332a84](https://github.com/phrase/openapi/commit/7332a84f9958346f2fb28dee4b0353519ef466d5))


### Bug Fixes

* **CLI:** fix required parameter handling ([#488](https://github.com/phrase/openapi/issues/488)) ([3d0412d](https://github.com/phrase/openapi/commit/3d0412df3c40b19cf8b12d5105e730990fd137b5))

## [2.16.0](https://github.com/phrase/openapi/compare/cli-v2.15.0...cli-v2.16.0) (2023-11-28)


### Features

* Add reports locales endpoint to API [TSS-2439] ([#465](https://github.com/phrase/openapi/issues/465)) ([e03aa9f](https://github.com/phrase/openapi/commit/e03aa9f49f031517b36db715fe70e8e0b65a438b))

## [2.15.0](https://github.com/phrase/openapi/compare/cli-v2.14.0...cli-v2.15.0) (2023-11-03)


### Features

* [TSI-2083] enable format_options argument for java-client   ([#426](https://github.com/phrase/openapi/issues/426)) ([faa8cb3](https://github.com/phrase/openapi/commit/faa8cb353ba9f1030b9f7cfd46b894b4d6d26e70))
* **CLI:** On pull error, print the API response ([#463](https://github.com/phrase/openapi/issues/463)) ([b0b6ed1](https://github.com/phrase/openapi/commit/b0b6ed1b928f2c64fb618de76d88742e5f0cee8c))
* Update openapi-generator to v7 ([#418](https://github.com/phrase/openapi/issues/418)) ([524626f](https://github.com/phrase/openapi/commit/524626f5e914bfef6025d0e1c2cbc7a728d08f56))

## [2.14.0](https://github.com/phrase/openapi/compare/cli-v2.13.0...cli-v2.14.0) (2023-10-23)


### Features

* **API:** Add order param to comment list endpoints ([#441](https://github.com/phrase/openapi/issues/441)) ([441c9c4](https://github.com/phrase/openapi/commit/441c9c46169f8c5ac4e71ade09a95dab136314ef))

## [2.13.0](https://github.com/phrase/openapi/compare/cli-v2.12.0...cli-v2.13.0) (2023-10-13)


### Features

* **API:** Implement figma attachments endpoints ([#415](https://github.com/phrase/openapi/issues/415)) ([970e612](https://github.com/phrase/openapi/commit/970e612fda620ca882a221ef541036b8d200b675))

## [2.12.0](https://github.com/phrase/openapi/compare/cli-v2.11.0...cli-v2.12.0) (2023-09-14)


### Features

* **CLI:** use phrase-go v2.14 ([#413](https://github.com/phrase/openapi/issues/413)) ([bf96057](https://github.com/phrase/openapi/commit/bf96057e8e0fffde405a65df5e2c993696080c58))
* Optionally tag only affected keys on upload [TSI-2041] ([#412](https://github.com/phrase/openapi/issues/412)) ([e8f958e](https://github.com/phrase/openapi/commit/e8f958e91469c2542f44ab68469c933688958383))
* **TSI-1946:** Add reviewed_at to translations ([#396](https://github.com/phrase/openapi/issues/396)) ([3e663d9](https://github.com/phrase/openapi/commit/3e663d971a99a816f0165dd6653a9a1e8a87c95e))

## [2.11.0](https://github.com/phrase/openapi/compare/cli-v2.10.0...cli-v2.11.0) (2023-08-28)


### Features

* **API:** Document new query parameters ([#393](https://github.com/phrase/openapi/issues/393)) ([770515a](https://github.com/phrase/openapi/commit/770515a9628122955bb3919405babf9392684eb9))

## [2.10.0](https://github.com/phrase/openapi/compare/cli-v2.9.1...cli-v2.10.0) (2023-08-24)


### Features

* **API:** Introduce comment replies endpoints ([#383](https://github.com/phrase/openapi/issues/383)) ([71351ac](https://github.com/phrase/openapi/commit/71351ac285f4f49976092e176c77b09f3485eb65))

## [2.9.1](https://github.com/phrase/openapi/compare/cli-v2.9.0...cli-v2.9.1) (2023-08-24)


### Bug Fixes

* **cli:** Bump go library version ([#384](https://github.com/phrase/openapi/issues/384)) ([ab2f655](https://github.com/phrase/openapi/commit/ab2f6556184bc7171c00c3f1b81908aec0e57a04))

## [2.9.0](https://github.com/phrase/openapi/compare/cli-v2.8.4...cli-v2.9.0) (2023-08-22)


### Features

* **TSE-950:** Document comment_reactions endpoints ([#380](https://github.com/phrase/openapi/issues/380)) ([f230244](https://github.com/phrase/openapi/commit/f230244e6e9c069b18edc4c35dd5e290fd14793b))

## [2.8.4](https://github.com/phrase/openapi/compare/cli-v2.8.3...cli-v2.8.4) (2023-07-31)


### Bug Fixes

* **schemas:** Fix gitlab_sync type mismatch ([#373](https://github.com/phrase/openapi/issues/373)) ([1cb1f65](https://github.com/phrase/openapi/commit/1cb1f650598c68afee6e2cd7c3c4ede1484aba35))

## [2.8.3](https://github.com/phrase/openapi/compare/cli-v2.8.2...cli-v2.8.3) (2023-07-28)


### Bug Fixes

* Fix gitlab_sync history status type mismatch ([#363](https://github.com/phrase/openapi/issues/363)) ([ebcaa4e](https://github.com/phrase/openapi/commit/ebcaa4e5dfcb2f73559a56c78b0f2512ca798375))
