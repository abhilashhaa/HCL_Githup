// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/piperenv"
	"github.com/SAP/jenkins-library/pkg/splunk"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/spf13/cobra"
)

type protecodeExecuteScanOptions struct {
	ExcludeCVEs                 string `json:"excludeCVEs,omitempty"`
	FailOnSevereVulnerabilities bool   `json:"failOnSevereVulnerabilities,omitempty"`
	ScanImage                   string `json:"scanImage,omitempty"`
	DockerRegistryURL           string `json:"dockerRegistryUrl,omitempty"`
	DockerConfigJSON            string `json:"dockerConfigJSON,omitempty"`
	CleanupMode                 string `json:"cleanupMode,omitempty"`
	FilePath                    string `json:"filePath,omitempty"`
	IncludeLayers               bool   `json:"includeLayers,omitempty"`
	TimeoutMinutes              string `json:"timeoutMinutes,omitempty"`
	ServerURL                   string `json:"serverUrl,omitempty"`
	ReportFileName              string `json:"reportFileName,omitempty"`
	FetchURL                    string `json:"fetchUrl,omitempty"`
	Group                       string `json:"group,omitempty"`
	VerifyOnly                  bool   `json:"verifyOnly,omitempty"`
	Username                    string `json:"username,omitempty"`
	Password                    string `json:"password,omitempty"`
	Version                     string `json:"version,omitempty"`
	PullRequestName             string `json:"pullRequestName,omitempty"`
}

type protecodeExecuteScanInflux struct {
	step_data struct {
		fields struct {
			protecode bool
		}
		tags struct {
		}
	}
	protecode_data struct {
		fields struct {
			excluded_vulnerabilities   int
			historical_vulnerabilities int
			major_vulnerabilities      int
			minor_vulnerabilities      int
			triaged_vulnerabilities    int
			vulnerabilities            int
		}
		tags struct {
		}
	}
}

func (i *protecodeExecuteScanInflux) persist(path, resourceName string) {
	measurementContent := []struct {
		measurement string
		valType     string
		name        string
		value       interface{}
	}{
		{valType: config.InfluxField, measurement: "step_data", name: "protecode", value: i.step_data.fields.protecode},
		{valType: config.InfluxField, measurement: "protecode_data", name: "excluded_vulnerabilities", value: i.protecode_data.fields.excluded_vulnerabilities},
		{valType: config.InfluxField, measurement: "protecode_data", name: "historical_vulnerabilities", value: i.protecode_data.fields.historical_vulnerabilities},
		{valType: config.InfluxField, measurement: "protecode_data", name: "major_vulnerabilities", value: i.protecode_data.fields.major_vulnerabilities},
		{valType: config.InfluxField, measurement: "protecode_data", name: "minor_vulnerabilities", value: i.protecode_data.fields.minor_vulnerabilities},
		{valType: config.InfluxField, measurement: "protecode_data", name: "triaged_vulnerabilities", value: i.protecode_data.fields.triaged_vulnerabilities},
		{valType: config.InfluxField, measurement: "protecode_data", name: "vulnerabilities", value: i.protecode_data.fields.vulnerabilities},
	}

	errCount := 0
	for _, metric := range measurementContent {
		err := piperenv.SetResourceParameter(path, resourceName, filepath.Join(metric.measurement, fmt.Sprintf("%vs", metric.valType), metric.name), metric.value)
		if err != nil {
			log.Entry().WithError(err).Error("Error persisting influx environment.")
			errCount++
		}
	}
	if errCount > 0 {
		log.Entry().Fatal("failed to persist Influx environment")
	}
}

// ProtecodeExecuteScanCommand Protecode is an Open Source Vulnerability Scanner that is capable of scanning binaries. It can be used to scan docker images but is supports many other programming languages especially those of the C family.
func ProtecodeExecuteScanCommand() *cobra.Command {
	const STEP_NAME = "protecodeExecuteScan"

	metadata := protecodeExecuteScanMetadata()
	var stepConfig protecodeExecuteScanOptions
	var startTime time.Time
	var influx protecodeExecuteScanInflux
	var logCollector *log.CollectorHook

	var createProtecodeExecuteScanCmd = &cobra.Command{
		Use:   STEP_NAME,
		Short: "Protecode is an Open Source Vulnerability Scanner that is capable of scanning binaries. It can be used to scan docker images but is supports many other programming languages especially those of the C family.",
		Long: `Protecode is an Open Source Vulnerability Scanner that is capable of scanning binaries. It can be used to scan docker images but is supports many other programming languages especially those of the C family.

!!! hint "Auditing findings (Triaging)"
    Triaging is now supported by the Protecode backend and also Piper does consider this information during the analysis of the scan results though product versions are not supported by Protecode. Therefore please make sure that the ` + "`" + `fileName` + "`" + ` you are providing does either contain a stable version or that it does not contain one at all. By ensuring that you are able to triage CVEs globally on the upload file's name without affecting any other artifacts scanned in the same Protecode group and as such triaged vulnerabilities will be considered during the next scan and will not fail the build anymore.`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			startTime = time.Now()
			log.SetStepName(STEP_NAME)
			log.SetVerbose(GeneralConfig.Verbose)

			path, _ := os.Getwd()
			fatalHook := &log.FatalHook{CorrelationID: GeneralConfig.CorrelationID, Path: path}
			log.RegisterHook(fatalHook)

			err := PrepareConfig(cmd, &metadata, STEP_NAME, &stepConfig, config.OpenPiperFile)
			if err != nil {
				log.SetErrorCategory(log.ErrorConfiguration)
				return err
			}
			log.RegisterSecret(stepConfig.DockerConfigJSON)
			log.RegisterSecret(stepConfig.Username)
			log.RegisterSecret(stepConfig.Password)

			if len(GeneralConfig.HookConfig.SentryConfig.Dsn) > 0 {
				sentryHook := log.NewSentryHook(GeneralConfig.HookConfig.SentryConfig.Dsn, GeneralConfig.CorrelationID)
				log.RegisterHook(&sentryHook)
			}

			if len(GeneralConfig.HookConfig.SplunkConfig.Dsn) > 0 {
				logCollector = &log.CollectorHook{CorrelationID: GeneralConfig.CorrelationID}
				log.RegisterHook(logCollector)
			}

			return nil
		},
		Run: func(_ *cobra.Command, _ []string) {
			telemetryData := telemetry.CustomData{}
			telemetryData.ErrorCode = "1"
			handler := func() {
				config.RemoveVaultSecretFiles()
				influx.persist(GeneralConfig.EnvRootPath, "influx")
				telemetryData.Duration = fmt.Sprintf("%v", time.Since(startTime).Milliseconds())
				telemetryData.ErrorCategory = log.GetErrorCategory().String()
				telemetry.Send(&telemetryData)
				if len(GeneralConfig.HookConfig.SplunkConfig.Dsn) > 0 {
					splunk.Send(&telemetryData, logCollector)
				}
			}
			log.DeferExitHandler(handler)
			defer handler()
			telemetry.Initialize(GeneralConfig.NoTelemetry, STEP_NAME)
			if len(GeneralConfig.HookConfig.SplunkConfig.Dsn) > 0 {
				splunk.Initialize(GeneralConfig.CorrelationID,
					GeneralConfig.HookConfig.SplunkConfig.Dsn,
					GeneralConfig.HookConfig.SplunkConfig.Token,
					GeneralConfig.HookConfig.SplunkConfig.Index,
					GeneralConfig.HookConfig.SplunkConfig.SendLogs)
			}
			protecodeExecuteScan(stepConfig, &telemetryData, &influx)
			telemetryData.ErrorCode = "0"
			log.Entry().Info("SUCCESS")
		},
	}

	addProtecodeExecuteScanFlags(createProtecodeExecuteScanCmd, &stepConfig)
	return createProtecodeExecuteScanCmd
}

func addProtecodeExecuteScanFlags(cmd *cobra.Command, stepConfig *protecodeExecuteScanOptions) {
	cmd.Flags().StringVar(&stepConfig.ExcludeCVEs, "excludeCVEs", ``, "DEPRECATED: Do use triaging within the Protecode UI instead")
	cmd.Flags().BoolVar(&stepConfig.FailOnSevereVulnerabilities, "failOnSevereVulnerabilities", true, "Whether to fail the job on severe vulnerabilties or not")
	cmd.Flags().StringVar(&stepConfig.ScanImage, "scanImage", os.Getenv("PIPER_scanImage"), "The reference to the docker image to scan with Protecode")
	cmd.Flags().StringVar(&stepConfig.DockerRegistryURL, "dockerRegistryUrl", os.Getenv("PIPER_dockerRegistryUrl"), "The reference to the docker registry to scan with Protecode")
	cmd.Flags().StringVar(&stepConfig.DockerConfigJSON, "dockerConfigJSON", os.Getenv("PIPER_dockerConfigJSON"), "Path to the file `.docker/config.json` - this is typically provided by your CI/CD system. You can find more details about the Docker credentials in the [Docker documentation](https://docs.docker.com/engine/reference/commandline/login/).")
	cmd.Flags().StringVar(&stepConfig.CleanupMode, "cleanupMode", `binary`, "Decides which parts are removed from the Protecode backend after the scan")
	cmd.Flags().StringVar(&stepConfig.FilePath, "filePath", os.Getenv("PIPER_filePath"), "The path to the file from local workspace to scan with Protecode")
	cmd.Flags().BoolVar(&stepConfig.IncludeLayers, "includeLayers", false, "Flag if the docker layers should be included")
	cmd.Flags().StringVar(&stepConfig.TimeoutMinutes, "timeoutMinutes", `60`, "The timeout to wait for the scan to finish")
	cmd.Flags().StringVar(&stepConfig.ServerURL, "serverUrl", os.Getenv("PIPER_serverUrl"), "The URL to the Protecode backend")
	cmd.Flags().StringVar(&stepConfig.ReportFileName, "reportFileName", `protecode_report.pdf`, "The file name of the report to be created")
	cmd.Flags().StringVar(&stepConfig.FetchURL, "fetchUrl", os.Getenv("PIPER_fetchUrl"), "The URL to fetch the file or image to scan with Protecode.")
	cmd.Flags().StringVar(&stepConfig.Group, "group", os.Getenv("PIPER_group"), "The Protecode group ID of your team")
	cmd.Flags().BoolVar(&stepConfig.VerifyOnly, "verifyOnly", false, "Whether the step shall only apply verification checks or whether it does a full scan and check cycle")
	cmd.Flags().StringVar(&stepConfig.Username, "username", os.Getenv("PIPER_username"), "User which is used for the protecode scan")
	cmd.Flags().StringVar(&stepConfig.Password, "password", os.Getenv("PIPER_password"), "Password which is used for the user")
	cmd.Flags().StringVar(&stepConfig.Version, "version", os.Getenv("PIPER_version"), "The version of the artifact to allow identification in protecode backend")
	cmd.Flags().StringVar(&stepConfig.PullRequestName, "pullRequestName", os.Getenv("PIPER_pullRequestName"), "The name of the pull request")

	cmd.MarkFlagRequired("serverUrl")
	cmd.MarkFlagRequired("group")
	cmd.MarkFlagRequired("username")
	cmd.MarkFlagRequired("password")
}

// retrieve step metadata
func protecodeExecuteScanMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:        "protecodeExecuteScan",
			Aliases:     []config.Alias{},
			Description: "Protecode is an Open Source Vulnerability Scanner that is capable of scanning binaries. It can be used to scan docker images but is supports many other programming languages especially those of the C family.",
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Secrets: []config.StepSecrets{
					{Name: "protecodeCredentialsId", Description: "Jenkins 'Username with password' credentials ID containing username and password to authenticate to the Protecode system.", Type: "jenkins"},
					{Name: "dockerConfigJsonCredentialsId", Description: "Jenkins 'Secret file' credentials ID containing Docker config.json (with registry credential(s)). You can create it like explained in the Docker Success Center in the article about [how to generate a new auth in the config.json file](https://success.docker.com/article/generate-new-auth-in-config-json-file).", Type: "jenkins", Aliases: []config.Alias{{Name: "dockerCredentialsId", Deprecated: true}}},
				},
				Parameters: []config.StepParameters{
					{
						Name:        "excludeCVEs",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",

						Mandatory: false,
						Default:   ``,
						Aliases:   []config.Alias{{Name: "protecodeExcludeCVEs"}},
					},
					{
						Name:        "failOnSevereVulnerabilities",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",

						Mandatory: false,
						Default:   true,
						Aliases:   []config.Alias{{Name: "protecodeFailOnSevereVulnerabilities"}},
					},
					{
						Name: "scanImage",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "commonPipelineEnvironment",
								Param: "container/imageNameTag",
							},
						},
						Scope: []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:  "string",

						Mandatory: false,
						Default:   os.Getenv("PIPER_scanImage"),
						Aliases:   []config.Alias{{Name: "dockerImage"}},
					},
					{
						Name: "dockerRegistryUrl",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "commonPipelineEnvironment",
								Param: "container/registryUrl",
							},
						},
						Scope: []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:  "string",

						Mandatory: false,
						Default:   os.Getenv("PIPER_dockerRegistryUrl"),
						Aliases:   []config.Alias{},
					},
					{
						Name: "dockerConfigJSON",
						ResourceRef: []config.ResourceReference{
							{
								Name: "dockerConfigJsonCredentialsId",
								Type: "secret",
							},

							{
								Name:  "",
								Paths: []string{"$(vaultPath)/docker-config", "$(vaultBasePath)/$(vaultPipelineName)/docker-config", "$(vaultBasePath)/GROUP-SECRETS/docker-config"},
								Type:  "vaultSecretFile",
							},
						},
						Scope: []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:  "string",

						Mandatory: false,
						Default:   os.Getenv("PIPER_dockerConfigJSON"),
						Aliases:   []config.Alias{},
					},
					{
						Name:        "cleanupMode",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",

						Mandatory: false,
						Default:   `binary`,
						Aliases:   []config.Alias{},
					},
					{
						Name:        "filePath",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",

						Mandatory: false,
						Default:   os.Getenv("PIPER_filePath"),
						Aliases:   []config.Alias{},
					},
					{
						Name:        "includeLayers",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",

						Mandatory: false,
						Default:   false,
						Aliases:   []config.Alias{},
					},
					{
						Name:        "timeoutMinutes",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",

						Mandatory: false,
						Default:   `60`,
						Aliases:   []config.Alias{{Name: "protecodeTimeoutMinutes"}},
					},
					{
						Name:        "serverUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",

						Mandatory: true,
						Default:   os.Getenv("PIPER_serverUrl"),
						Aliases:   []config.Alias{{Name: "protecodeServerUrl"}},
					},
					{
						Name:        "reportFileName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",

						Mandatory: false,
						Default:   `protecode_report.pdf`,
						Aliases:   []config.Alias{},
					},
					{
						Name:        "fetchUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",

						Mandatory: false,
						Default:   os.Getenv("PIPER_fetchUrl"),
						Aliases:   []config.Alias{},
					},
					{
						Name:        "group",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",

						Mandatory: true,
						Default:   os.Getenv("PIPER_group"),
						Aliases:   []config.Alias{{Name: "protecodeGroup"}},
					},
					{
						Name:        "verifyOnly",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",

						Mandatory: false,
						Default:   false,
						Aliases:   []config.Alias{{Name: "reuseExisting"}},
					},
					{
						Name: "username",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "protecodeCredentialsId",
								Param: "username",
								Type:  "secret",
							},

							{
								Name:  "",
								Paths: []string{"$(vaultPath)/protecode", "$(vaultBasePath)/$(vaultPipelineName)/protecode", "$(vaultBasePath)/GROUP-SECRETS/protecode"},
								Type:  "vaultSecret",
							},
						},
						Scope: []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:  "string",

						Mandatory: true,
						Default:   os.Getenv("PIPER_username"),
						Aliases:   []config.Alias{{Name: "user"}},
					},
					{
						Name: "password",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "protecodeCredentialsId",
								Param: "password",
								Type:  "secret",
							},

							{
								Name:  "",
								Paths: []string{"$(vaultPath)/protecode", "$(vaultBasePath)/$(vaultPipelineName)/protecode", "$(vaultBasePath)/GROUP-SECRETS/protecode"},
								Type:  "vaultSecret",
							},
						},
						Scope: []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:  "string",

						Mandatory: true,
						Default:   os.Getenv("PIPER_password"),
						Aliases:   []config.Alias{},
					},
					{
						Name: "version",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "commonPipelineEnvironment",
								Param: "artifactVersion",
							},
						},
						Scope: []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:  "string",

						Mandatory: false,
						Default:   os.Getenv("PIPER_version"),
						Aliases:   []config.Alias{{Name: "artifactVersion"}},
					},
					{
						Name:        "pullRequestName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",

						Mandatory: false,
						Default:   os.Getenv("PIPER_pullRequestName"),
						Aliases:   []config.Alias{},
					},
				},
			},
			Outputs: config.StepOutputs{
				Resources: []config.StepResources{
					{
						Name: "influx",
						Type: "influx",
						Parameters: []map[string]interface{}{
							{"Name": "step_data"}, {"fields": []map[string]string{{"name": "protecode"}}},
							{"Name": "protecode_data"}, {"fields": []map[string]string{{"name": "excluded_vulnerabilities"}, {"name": "historical_vulnerabilities"}, {"name": "major_vulnerabilities"}, {"name": "minor_vulnerabilities"}, {"name": "triaged_vulnerabilities"}, {"name": "vulnerabilities"}}},
						},
					},
				},
			},
		},
	}
	return theMetaData
}
