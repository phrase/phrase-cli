package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/phrase/phrase-cli/cmd/internal/paths"
	"github.com/phrase/phrase-cli/cmd/internal/print"
	"github.com/phrase/phrase-cli/cmd/internal/prompt"
	"github.com/phrase/phrase-cli/cmd/internal/shared"
	"github.com/phrase/phrase-cli/cmd/internal/spinner"
	"github.com/phrase/phrase-go/v3"
	"gopkg.in/yaml.v2"
)

// The steps for a successful project initialization.
const (
	StepAskForToken   = "ask for token"
	StepSelectProject = "select project"
	StepSelectFormat  = "select format"
	StepConfigSources = "config sources"
	StepConfigTargets = "config targets"
	StepWriteConfig   = "write configuration file"
	StepFinished      = "finished"
)

var nextStep = map[string]string{
	StepAskForToken:   StepSelectProject,
	StepSelectProject: StepSelectFormat,
	StepSelectFormat:  StepConfigSources,
	StepConfigSources: StepConfigTargets,
	StepConfigTargets: StepWriteConfig,
	StepWriteConfig:   StepFinished,
}

type stepFunc func(*InitCommand) error

var stepFuncs = map[string]stepFunc{
	StepAskForToken:   (*InitCommand).askForToken,
	StepSelectProject: (*InitCommand).selectProject,
	StepSelectFormat:  (*InitCommand).selectFormat,
	StepConfigSources: (*InitCommand).configureSources,
	StepConfigTargets: (*InitCommand).configureTargets,
	StepWriteConfig:   (*InitCommand).writeConfig,
}

// structs that can be marshalled to YAML to create a valid configuration file

type ConfigYAML struct {
	Host        string                            `yaml:"host,omitempty"`
	AccessToken string                            `yaml:"access_token,omitempty"`
	ProjectID   string                            `yaml:"project_id"`
	FileFormat  string                            `yaml:"file_format,omitempty"`
	PerPage     int                               `yaml:"per_page,omitempty"`
	Defaults    map[string]map[string]interface{} `yaml:"defaults,omitempty"`
	Push        PushYAML                          `yaml:"push,omitempty"`
	Pull        PullYAML                          `yaml:"pull,omitempty"`
}

type PushYAML struct {
	Sources []SourcesYAML `yaml:"sources,omitempty"`
}

type PullYAML struct {
	Targets []TargetsYAML `yaml:"targets,omitempty"`
}

type SourcesYAML struct {
	File   string                 `yaml:"file,omitempty"`
	Params map[string]interface{} `yaml:"params,omitempty"`
}

type TargetsYAML SourcesYAML

// the actual command

type InitCommand struct {
	phrase.Config

	client     *phrase.APIClient
	YAML       ConfigYAML
	FileFormat *phrase.Format
}

func (cmd *InitCommand) Run() error {
	// keep host if specified in config file or as command line parameter
	if cmd.Config.Host != "" {
		cmd.YAML.Host = cmd.Config.Host
	}

	step := StepSelectProject
	if cmd.Config.Token == "" {
		step = StepAskForToken
	} else {
		cmd.setToken(cmd.Config.Token)
	}

	for step != StepFinished {
		err := stepFuncs[step](cmd)
		if err != nil {
			return err
		}

		fmt.Println()

		step = nextStep[step]
	}

	return nil
}

func (cmd *InitCommand) askForToken() error {
	print.PhraseLogo()
	fmt.Println("phrase.com API Client Setup")
	fmt.Println()

	token := ""
	for {
		err := prompt.P("Please enter your API access token (you can generate one in your profile at phrase.com):", &token)
		if err != nil {
			continue
		}

		token = strings.ToLower(token)
		success, err := regexp.MatchString("^[0-9a-f]{64}$", token)
		if err != nil {
			continue
		}

		if !success {
			print.Failure("Invalid access token! A valid access token is 64 characters long and contains only a-f, 0-9.")
			continue
		}

		break
	}

	cmd.setToken(token)
	return nil
}

func (cmd *InitCommand) setToken(token string) {
	cmd.YAML.AccessToken = token
	cmd.Credentials.Token = token
	Config = &cmd.Config
	client := newClient()
	cmd.client = client
}

func (cmd *InitCommand) selectProject() error {
	taskResult := make(chan []phrase.Project, 1)
	taskErr := make(chan error, 1)

	Config = &cmd.Config
	client := newClient()

	fmt.Print("Loading projects... ")
	spinner.While(func() {
		projects, _, err := Projects(client)
		taskResult <- projects
		taskErr <- err
	})
	fmt.Println()

	projects := <-taskResult
	if err := <-taskErr; err != nil {
		if strings.Contains(err.Error(), "401") {
			return fmt.Errorf("%s is not a valid access token. It may be revoked or missing the read or write scope. Please create a new token and try again.", cmd.Credentials.Token)
		}
		return err
	}

	if len(projects) == 0 {
		fmt.Println("Since you don't have any projects yet, a new one will be created.")
		return cmd.newProject()
	}

	for i, project := range projects {
		fmt.Printf("%2d: %s (Id: %s)\n", i+1, project.Name, project.Id)
	}
	fmt.Printf("%2d: Create new project\n", len(projects)+1)

	selection := 0
	var err error
	for {
		err = prompt.P(fmt.Sprintf("Select project: (%v-%v)", 1, len(projects)+1), &selection)
		if err != nil {
			continue
		}

		if selection < 1 || selection > len(projects)+1 {
			print.Failure("Please select a project from the list by specifying its position in the list, e.g. 2 for the second project.")
			continue
		}

		break
	}

	if selection == len(projects)+1 {
		return cmd.newProject()
	}

	print.Success("Using project %v", projects[selection-1].Name)

	cmd.YAML.ProjectID = projects[selection-1].Id
	cmd.DefaultFileFormat = projects[selection-1].MainFormat

	return nil
}

func (cmd *InitCommand) newProject() error {
	params := phrase.ProjectCreateParameters{
		Name: "",
	}

	for {
		err := prompt.P("Enter the name of the new project:", &params.Name)
		if err == nil {
			break
		}
	}

	var details *phrase.ProjectDetails

	data, _, err := cmd.client.ProjectsApi.ProjectCreate(Auth, params, &phrase.ProjectCreateOpts{})
	if err != nil {
		if strings.Contains(err.Error(), "401") {
			return fmt.Errorf("Your access token %s is not valid for the 'write' scope. Please create a new Access Token with read and write scope.", cmd.Credentials.Token)
		}
		return err
	}

	details = &data

	print.Success("Using project %v", details.Name)

	cmd.YAML.ProjectID = details.Id

	return nil
}

func (cmd *InitCommand) selectFormat() error {
	formats, _, err := cmd.client.FormatsApi.FormatsList(Auth)
	if err != nil {
		return err
	}

	// ensure that the default file format from the config file is a valid format
	for _, format := range formats {
		if format.ApiName == cmd.DefaultFileFormat {
			cmd.FileFormat = &format
			break
		}
	}

	for i, format := range formats {
		fmt.Printf("%2d: %s - %s, file extension: %s\n", i+1, format.ApiName, format.Name, format.Extension)
	}

	promptText := fmt.Sprintf("Select the format to use for language files you download from Phrase Strings (%v-%v", 1, len(formats))
	if cmd.FileFormat != nil && cmd.FileFormat.Name != "" {
		promptText += fmt.Sprintf(" or leave blank to use the default, %s)", cmd.FileFormat.Name)
	}
	promptText += "):"

	selection := 0
	for {
		err = prompt.P(promptText, &selection)
		if err != nil {
			if cmd.FileFormat != nil && cmd.FileFormat.Name != "" {
				break
			}

			continue
		}

		if selection < 1 || selection > len(formats) {
			print.Failure("Please select a format from the list by specifying the number in front of it.")
			continue
		}

		cmd.FileFormat = &formats[selection-1]
		break
	}

	print.Success("Using format %v", cmd.FileFormat.Name)

	return nil
}

func (cmd *InitCommand) configureSources() error {
	fmt.Println("Enter the path to the language file you want to upload to Phrase.")
	fmt.Printf("For documentation, see %s#push\n", shared.DocsConfigUrl)

	pushPath := ""
	for {
		err := prompt.WithDefault("Source file path:", &pushPath, cmd.FileFormat.DefaultFile)
		if err != nil {
			return err
		}

		err = paths.Validate(pushPath, cmd.FileFormat.ApiName, cmd.FileFormat.Extension)
		if err != nil {
			print.Failure(err.Error())
		} else {
			break
		}
	}

	sourceYAML := SourcesYAML{
		File: pushPath,
		Params: map[string]interface{}{
			"file_format": cmd.FileFormat.ApiName,
		},
	}

	cmd.YAML.Push.Sources = append(cmd.YAML.Push.Sources, sourceYAML)

	return nil
}

func (cmd *InitCommand) configureTargets() error {
	fmt.Println("Enter the path to which to download language files from Phrase.")
	fmt.Printf("For documentation, see %s#pull\n", shared.DocsConfigUrl)

	pullPath := ""
	for {
		err := prompt.WithDefault("Target file path:", &pullPath, cmd.FileFormat.DefaultFile)
		if err != nil {
			return err
		}

		err = paths.Validate(pullPath, cmd.FileFormat.ApiName, cmd.FileFormat.Extension)
		if err != nil {
			print.Failure(err.Error())
		} else {
			break
		}
	}

	targetYAML := TargetsYAML{
		File: pullPath,
		Params: map[string]interface{}{
			"file_format": cmd.FileFormat.ApiName,
		},
	}

	cmd.YAML.Pull.Targets = append(cmd.YAML.Pull.Targets, targetYAML)

	return nil
}

func (cmd *InitCommand) writeConfig() error {
	wrapper := struct {
		Config ConfigYAML `yaml:"phrase"`
	}{
		Config: cmd.YAML,
	}

	yamlBytes, err := yaml.Marshal(wrapper)
	if err != nil {
		return err
	}

	filename := ".phrase.yml"
	err = ioutil.WriteFile(filename, yamlBytes, 0644)
	if err != nil {
		return err
	}

	print.Success("We created the following configuration file for you: " + filename)

	fmt.Println()
	fmt.Println(string(yamlBytes))

	print.Success("For advanced configuration options, take a look at the documentation: " + shared.DocsConfigUrl)
	print.Success("You can now use the push & pull commands in your workflow:")
	fmt.Println()
	fmt.Println("$ phrase push")
	fmt.Println("$ phrase pull")
	fmt.Println()

	pushNow := ""
	_ = prompt.WithDefault("Do you want to upload your locales now for the first time? (y/n)", &pushNow, "y")
	if pushNow == "y" {
		err = firstPush()
		if err != nil {
			return err
		}
	}

	print.Success("Project initialization completed!")

	return nil
}

func firstPush() error {
	config, err := phrase.ReadConfig("")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(2)
	}
	cmd := &PushCommand{Config: *config}
	return cmd.Run()
}
