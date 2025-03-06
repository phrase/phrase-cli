# Changelog

## [2.39.0](https://github.com/phrase/openapi/compare/cli-v2.38.0...cli-v2.39.0) (2025-03-06)


### Features

* **CLI:** Support cleanup over multiple sources #STRINGS-1378 ([#810](https://github.com/phrase/openapi/issues/810)) ([e78ab70](https://github.com/phrase/openapi/commit/e78ab704ef451b1be161011204f6d402201408a1))

## [2.38.0](https://github.com/phrase/openapi/compare/cli-v2.37.0...cli-v2.38.0) (2025-03-05)


### Features

* **API:** Include roles in account response ([#811](https://github.com/phrase/openapi/issues/811)) ([dc27ee5](https://github.com/phrase/openapi/commit/dc27ee5117762222b6e1e6abb639f8e00c6a9101))

## [2.37.0](https://github.com/phrase/openapi/compare/cli-v2.36.0...cli-v2.37.0) (2025-02-25)


### Features

* **API:** add updated_since filter to job list #STRINGS-1555 ([#799](https://github.com/phrase/openapi/issues/799)) ([dc9b6ed](https://github.com/phrase/openapi/commit/dc9b6ed12e013231d397820449086c87fea2f8ba))


### Bug Fixes

* **CLI:** Revert viper library version #STRINGS-1487 ([#809](https://github.com/phrase/openapi/issues/809)) ([8b68918](https://github.com/phrase/openapi/commit/8b689182b5e23bf7fee447db9229c5aba0417c7d))

## [2.36.0](https://github.com/phrase/openapi/compare/cli-v2.35.6...cli-v2.36.0) (2025-02-17)


### Features

* **API:** Add locale_ids param to synchronous download endpoint [STRINGS-1492] ([#780](https://github.com/phrase/openapi/issues/780)) ([47186a4](https://github.com/phrase/openapi/commit/47186a44fc8c0b8e466636acf3d49413b1f29f30))
* **API:** Add source last updated at information on job details ([#777](https://github.com/phrase/openapi/issues/777)) ([c9b8423](https://github.com/phrase/openapi/commit/c9b8423766b4138980d0553502b3e18ca524f34e))
* **API:** document tags attribute of an upload #STRINGS-1221 ([#790](https://github.com/phrase/openapi/issues/790)) ([fff505b](https://github.com/phrase/openapi/commit/fff505bdff35a0033fee06e505c42fe794c88562))

## [2.35.6](https://github.com/phrase/openapi/compare/cli-v2.35.5...cli-v2.35.6) (2025-01-29)


### Bug Fixes

* **API:** pass translation_key_ids when removing keys from job ([#771](https://github.com/phrase/openapi/issues/771)) ([f670e27](https://github.com/phrase/openapi/commit/f670e2763b1112fefd1812109b3c09def42b7bd2))
* **cli:** CLI do not exit with 0 on pull when branch does not exist ([#768](https://github.com/phrase/openapi/issues/768)) ([e7d8c43](https://github.com/phrase/openapi/commit/e7d8c435dafec065ed8fb657e8879bf843ee3a48))
* **CLI:** Undo version bump ([#770](https://github.com/phrase/openapi/issues/770)) ([643aba1](https://github.com/phrase/openapi/commit/643aba13c79c6350f7416af3f5a78837bc937055))

## [2.35.5](https://github.com/phrase/openapi/compare/cli-v2.35.4...cli-v2.35.5) (2025-01-06)


### Bug Fixes

* **CLI:** Adjust operationId for quality_performance_score #STRINGS-1104 ([#721](https://github.com/phrase/openapi/issues/721)) ([7aa3b9b](https://github.com/phrase/openapi/commit/7aa3b9b508d1d24a4af7f4977b1a2fead8bfda78))

## [2.35.4](https://github.com/phrase/openapi/compare/cli-v2.35.3...cli-v2.35.4) (2024-12-20)


### Bug Fixes

* **API:** Repo Sync Event errors field type #STRINGS-1074 ([#756](https://github.com/phrase/openapi/issues/756)) ([c7670e0](https://github.com/phrase/openapi/commit/c7670e04810f95359d72ba6346b5f626bfb77b6f))

## [2.35.3](https://github.com/phrase/openapi/compare/cli-v2.35.2...cli-v2.35.3) (2024-12-20)


### Bug Fixes

* **API:** Repo Sync schema fixes #STRINGS-1074 ([#748](https://github.com/phrase/openapi/issues/748)) ([033be10](https://github.com/phrase/openapi/commit/033be1003fe01b5115de1f8ba2336d32b4862bfd))

## [2.35.2](https://github.com/phrase/openapi/compare/cli-v2.35.1...cli-v2.35.2) (2024-12-19)


### Bug Fixes

* **CLI:** Bump go version ([#745](https://github.com/phrase/openapi/issues/745)) ([7210e8a](https://github.com/phrase/openapi/commit/7210e8ae8f9f8cb04bca535658f65e30d1ca4831))

## [2.35.1](https://github.com/phrase/openapi/compare/cli-v2.35.0...cli-v2.35.1) (2024-12-19)


### Bug Fixes

* **CLI:** Bump dependencies' versions ([#742](https://github.com/phrase/openapi/issues/742)) ([d78d620](https://github.com/phrase/openapi/commit/d78d6209eed6bddda05260d81567fcaffd9d637b))

## [2.35.0](https://github.com/phrase/openapi/compare/cli-v2.34.1...cli-2.25.0) (2024-12-18)


### ⚠ BREAKING CHANGES

* Remove old Git sync endpoints. Replaced with new repo sync ([#735](https://github.com/phrase/openapi/issues/735))

### Features

* **API:** Add 'default_encoding' documentation ([#733](https://github.com/phrase/openapi/issues/733)) ([0139c51](https://github.com/phrase/openapi/commit/0139c51da747fbe7bc9929bcf3534aad7f22f39a))
* Remove old Git sync endpoints. Replaced with new repo sync ([#735](https://github.com/phrase/openapi/issues/735)) ([c3bd8ec](https://github.com/phrase/openapi/commit/c3bd8eccaabcfa1b1066ea4438971ac59833af46))


### Bug Fixes

* **API:** Add missing branch parameter to job comment endpoints #STRINGS-988 ([#724](https://github.com/phrase/openapi/issues/724)) ([64d399c](https://github.com/phrase/openapi/commit/64d399ced0980ac2a48366f91110047287a0c590))

## [2.34.1](https://github.com/phrase/openapi/compare/cli-v2.34.0...cli-v2.34.1) (2024-11-27)


### Bug Fixes

* **CLI:** crash when target/params section is empty #STRINGS-921 ([#722](https://github.com/phrase/openapi/issues/722)) ([03032c4](https://github.com/phrase/openapi/commit/03032c43be427a948f7e49cb7b7ff257bcf41821))

## [2.34.0](https://github.com/phrase/openapi/compare/cli-v2.33.1...cli-v2.34.0) (2024-11-25)


### Features

* **API:** Add Pagination header to POST search endpoints [[#457](https://github.com/phrase/openapi/issues/457)] ([#706](https://github.com/phrase/openapi/issues/706)) ([9a79fa3](https://github.com/phrase/openapi/commit/9a79fa31bb3b9d58272fa2f4e82d72d0d44a93a0))
* **API:** autotranslate param in key creation [STRINGS-786] ([#713](https://github.com/phrase/openapi/issues/713)) ([581d0ff](https://github.com/phrase/openapi/commit/581d0ff5f1d06757e5ddd9603b78fc8d435d68ee))
* **CLI:** Support branch in pull config #STRINGS-538 ([#701](https://github.com/phrase/openapi/issues/701)) ([30ae809](https://github.com/phrase/openapi/commit/30ae809b38ead0d0d019362b67a8b604ac15fe5f))


### Bug Fixes

* **API:** Comment creation schema fix #STRINGS-866 ([#718](https://github.com/phrase/openapi/issues/718)) ([e201d13](https://github.com/phrase/openapi/commit/e201d1360c89698dd8d3642cc28f89dd0e50a1fb))

## [2.33.1](https://github.com/phrase/openapi/compare/cli-v2.33.0...cli-v2.33.1) (2024-10-10)


### Bug Fixes

* **cli:** Adapt to formats API fix ([#699](https://github.com/phrase/openapi/issues/699)) ([3363de7](https://github.com/phrase/openapi/commit/3363de7b1f9564dad363932c3964a24b87887e7d))

## [2.33.0](https://github.com/phrase/openapi/compare/cli-v2.32.0...cli-v2.33.0) (2024-10-02)


### Features

* Add translation key prefixes for upload and download ([#687](https://github.com/phrase/openapi/issues/687)) ([9c9c959](https://github.com/phrase/openapi/commit/9c9c959830631bcac8beaf1de30ab31755ac1ee5))


### Bug Fixes

* **API:** Format list is not paginated and authenticated #STRINGS-458 ([#690](https://github.com/phrase/openapi/issues/690)) ([25e90f4](https://github.com/phrase/openapi/commit/25e90f46513e70cf328be80c36ae785cead05851))

## [2.32.0](https://github.com/phrase/openapi/compare/cli-v2.31.1...cli-v2.32.0) (2024-09-09)


### Features

* Add update_translations_on_source_match ([#670](https://github.com/phrase/openapi/issues/670)) ([11003ac](https://github.com/phrase/openapi/commit/11003ace7353bf99893482ca4aa32214abf3e581))

## [2.31.1](https://github.com/phrase/openapi/compare/cli-v2.31.0...cli-v2.31.1) (2024-07-15)


### Bug Fixes

* **CLI:** Pull should not crash on empty files ([#655](https://github.com/phrase/openapi/issues/655)) ([3e3b33c](https://github.com/phrase/openapi/commit/3e3b33cf0e5dd173c2596b23f5a815692f9ee865))

## [2.31.0](https://github.com/phrase/openapi/compare/cli-v2.30.0...cli-v2.31.0) (2024-07-03)


### Features

* **CLI:** Add option for async download [TSI-2515] ([#649](https://github.com/phrase/openapi/issues/649)) ([976353a](https://github.com/phrase/openapi/commit/976353aa639310dd8bad45cc090aff4768b520f1))

## [2.30.0](https://github.com/phrase/openapi/compare/cli-v2.29.0...cli-v2.30.0) (2024-07-02)


### Features

* add repo sync events show endpoint ([#641](https://github.com/phrase/openapi/issues/641)) ([e1d9cfb](https://github.com/phrase/openapi/commit/e1d9cfb23e079fea2d9e5475dff9a4137f1f0154))

## [2.29.0](https://github.com/phrase/openapi/compare/cli-v2.28.1...cli-v2.29.0) (2024-06-25)


### Features

* **API:** Async downloads [TSI-2515] ([#642](https://github.com/phrase/openapi/issues/642)) ([6fcab5d](https://github.com/phrase/openapi/commit/6fcab5d4719f64e8e5dd49c327dc9348b384de4c))

## [2.28.1](https://github.com/phrase/openapi/compare/cli-v2.28.0...cli-v2.28.1) (2024-06-18)


### Bug Fixes

* add app_min_version and app_max_version param to releases ([#633](https://github.com/phrase/openapi/issues/633)) ([b384301](https://github.com/phrase/openapi/commit/b3843012460ace4c1d34c4373e5158595466adcb))

## [2.28.0](https://github.com/phrase/openapi/compare/cli-v2.27.1...cli-v2.28.0) (2024-06-12)


### Features

* **API:** Add OTA Release Triggers API [TSI-2485] ([#622](https://github.com/phrase/openapi/issues/622)) ([8cb91dc](https://github.com/phrase/openapi/commit/8cb91dcce2c19ca700cf9d0713fa74f28ad59434))

## [2.27.1](https://github.com/phrase/openapi/compare/cli-v2.27.0...cli-v2.27.1) (2024-05-31)


### Bug Fixes

* **CLI:** use id instead of repo_sync_id as param ([#618](https://github.com/phrase/openapi/issues/618)) ([7a1a0d9](https://github.com/phrase/openapi/commit/7a1a0d9115262dc7fa3065ee7cc76bec7a381ddb))
* **PHP:** Fix deserializing custom metadata values in key response ([#607](https://github.com/phrase/openapi/issues/607)) ([b6eeeba](https://github.com/phrase/openapi/commit/b6eeeba223e3eabec268a8f3d372afcb6abd09dd))

## [2.27.0](https://github.com/phrase/openapi/compare/cli-v2.26.0...cli-v2.27.0) (2024-04-29)


### Features

* **API:** Add 'update_translation_keys' for Uploads [TSI-2292] ([#578](https://github.com/phrase/openapi/issues/578)) ([4492ec0](https://github.com/phrase/openapi/commit/4492ec0a7c62f9de9ab1c1125071615bddcc26ce))

## [2.26.0](https://github.com/phrase/openapi/compare/cli-v2.25.0...cli-v2.26.0) (2024-04-23)


### Bug Fixes

* add missing required params ([#571](https://github.com/phrase/openapi/issues/571))

### Code Refactoring

* add missing required params ([#571](https://github.com/phrase/openapi/issues/571)) ([d810e9e](https://github.com/phrase/openapi/commit/d810e9ebc767e14ba9e56106de8c5774d9d6d178))

## [2.25.0](https://github.com/phrase/openapi/compare/cli-v2.24.0...cli-v2.25.0) (2024-04-22)


### Features

* Add linked-parent to translation details ([#570](https://github.com/phrase/openapi/issues/570)) ([2c6f432](https://github.com/phrase/openapi/commit/2c6f43253e24b670b71ac810c85dce0759c29403))

## [2.24.0](https://github.com/phrase/openapi/compare/cli-v2.23.2...cli-v2.24.0) (2024-04-17)


### Features

* **API:** Add Repo Sync [TSI-1923] ([#569](https://github.com/phrase/openapi/issues/569)) ([0bd1756](https://github.com/phrase/openapi/commit/0bd17562018cb045ff41cc1ff5008b9419a0ed12))

## [2.23.2](https://github.com/phrase/openapi/compare/cli-v2.23.1...cli-v2.23.2) (2024-04-10)


### Bug Fixes

* **CLI:** Update protobuf library ([#567](https://github.com/phrase/openapi/issues/567)) ([a19bb4b](https://github.com/phrase/openapi/commit/a19bb4ba90a994d32e55206bcde4eba72d2eec6a))

## [2.23.1](https://github.com/phrase/openapi/compare/cli-v2.23.0...cli-v2.23.1) (2024-04-04)


### Bug Fixes

* (API) Add mandatory params to linked keys endpoints ([#564](https://github.com/phrase/openapi/issues/564)) ([08d9846](https://github.com/phrase/openapi/commit/08d9846bc224d349e2ade9abf28d733afb1e8be3))

## [2.23.0](https://github.com/phrase/openapi/compare/cli-v2.22.0...cli-v2.23.0) (2024-04-03)


### Features

* **API:** add Linked Keys endpoints ([#555](https://github.com/phrase/openapi/issues/555)) ([4935dac](https://github.com/phrase/openapi/commit/4935dac58c787eaade2f1f65ce649f466b5e3a60))

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
