# CLI v2 client for phrase

The Phrase Client is available for all major platforms and lets you access all API endpoints as well as easily sync your locale files between your source code and Phrase.

Phrase is a translation management platform for software projects. You can collaborate on language file translation with your team or order translations through our platform. The API allows you to import locale files, download locale files, tag keys or interact in other ways with the localization data stored in Phrase for your account.

This CLI client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.

- API version: 2.0.0
- Package version: 2.4.12

- Build package: org.openapitools.codegen.languages.GoClientCodegen

For more information, please visit [https://developers.phrase.com/api/](https://developers.phrase.com/api/)


## Quick Start

This quick start will guide you through the basic steps to get up and running with the Phrase Client.

#### 1. Install

[Download and install](https://phrase.com/cli) the client for your platform. See our [detailed installation guide](https://help.phrase.com/help/installation-1) for more information.

##### Homebrew

If you use homebrew, we have provided a tap to make installation easier on Mac OS X:

        brew tap phrase/brewed
        brew install phrase

The tap is linked to our Formula collection and will be updated, when you call `brew update` as well.

#### 2. Init

Initialize your project by executing the `init` command. This lets you define your preferred locale file format, source files and more.

    $ cd /path/to/project
    $ phrase init

#### 3. Upload your locale files

Use the `push` command to upload your locale files from your defined [sources](https://help.phrase.com/help/configuration):

    $ phrase push

#### 4. Download your locale files

Use the `pull` command to download the most recent locale files back into your project according to your [targets](https://help.phrase.com/help/configuration):

    $ phrase pull

#### 5. Docker

You can also use `phrase` through a docker image, without installing the cli on your computer.

    $ docker run --rm phrase/phrase-cli:2.0.15 help

Some commands are interactive and require the `-it` flag.

    $ docker run -it phrase/phrase-cli:2.0.15 init

You can also use the docker image as base for more complex images.

#### 6. More

To see a list of all available commands, simply execute:

    $ phrase

To see all supported options for a command, simple use the `--help` flag:

    $ phrase locales list --help

See our [detailed guides](https://help.phrase.com/help/phrase-in-your-terminal) for in-depth instructions on how to use the Phrase Client.

## Contributing

This tool and it's source code are auto-generated from templates that run against a API specification file. Therefore we can not accept any pull requests in this repository. Please use the GitHub Issue Tracker to report bugs.

## Further reading
* [Phrase Client Download Page](https://phrase.com/cli)

## Licenses

phrase-client is licensed under MIT license. (see [LICENSE](LICENSE))

Parts of phrase-client use third party libraries which are vendored and licensed under different licenses

## Author

support@phrase.com

## Get help / support

Please contact [support@phrase.com](mailto:support@phrase.com?subject=[GitHub]%20phrase-cli) and we can take more direct action toward finding a solution.
