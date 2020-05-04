module github.com/phrase/phrase-cli

go 1.14

require (
	github.com/antihax/optional v1.0.0
	github.com/coreos/go-semver v0.2.0
	github.com/daviddengcn/go-colortext v0.0.0-20171126034257-17e75f6184bc
	github.com/phrase/phrase-go v1.0.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.6.3
	golang.org/dl v0.0.0-20200414231856-f86334ee252a // indirect
	gopkg.in/yaml.v2 v2.2.8
)

replace github.com/phrase/phrase-go => ../phrase-go
