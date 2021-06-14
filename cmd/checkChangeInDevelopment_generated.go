// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/splunk"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/spf13/cobra"
)

type checkChangeInDevelopmentOptions struct {
	Endpoint                       string   `json:"endpoint,omitempty"`
	Username                       string   `json:"username,omitempty"`
	Password                       string   `json:"password,omitempty"`
	ChangeDocumentID               string   `json:"changeDocumentId,omitempty"`
	FailIfStatusIsNotInDevelopment bool     `json:"failIfStatusIsNotInDevelopment,omitempty"`
	ClientOpts                     []string `json:"clientOpts,omitempty"`
}

// CheckChangeInDevelopmentCommand Checks if a certain change is in status 'in development'
func CheckChangeInDevelopmentCommand() *cobra.Command {
	const STEP_NAME = "checkChangeInDevelopment"

	metadata := checkChangeInDevelopmentMetadata()
	var stepConfig checkChangeInDevelopmentOptions
	var startTime time.Time
	var logCollector *log.CollectorHook

	var createCheckChangeInDevelopmentCmd = &cobra.Command{
		Use:   STEP_NAME,
		Short: "Checks if a certain change is in status 'in development'",
		Long:  `"Checks if a certain change is in status 'in development'"`,
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
			checkChangeInDevelopment(stepConfig, &telemetryData)
			telemetryData.ErrorCode = "0"
			log.Entry().Info("SUCCESS")
		},
	}

	addCheckChangeInDevelopmentFlags(createCheckChangeInDevelopmentCmd, &stepConfig)
	return createCheckChangeInDevelopmentCmd
}

func addCheckChangeInDevelopmentFlags(cmd *cobra.Command, stepConfig *checkChangeInDevelopmentOptions) {
	cmd.Flags().StringVar(&stepConfig.Endpoint, "endpoint", os.Getenv("PIPER_endpoint"), "The service endpoint")
	cmd.Flags().StringVar(&stepConfig.Username, "username", os.Getenv("PIPER_username"), "The user")
	cmd.Flags().StringVar(&stepConfig.Password, "password", os.Getenv("PIPER_password"), "The password")
	cmd.Flags().StringVar(&stepConfig.ChangeDocumentID, "changeDocumentId", os.Getenv("PIPER_changeDocumentId"), "The change document which should be checked for the status")
	cmd.Flags().BoolVar(&stepConfig.FailIfStatusIsNotInDevelopment, "failIfStatusIsNotInDevelopment", true, "lets the build fail in case the change is not in status 'in developent'. Otherwise a warning is emitted to the log")
	cmd.Flags().StringSliceVar(&stepConfig.ClientOpts, "clientOpts", []string{}, "additional options passed to cm client, e.g. for troubleshooting")

	cmd.MarkFlagRequired("endpoint")
	cmd.MarkFlagRequired("username")
	cmd.MarkFlagRequired("password")
	cmd.MarkFlagRequired("changeDocumentId")
}

// retrieve step metadata
func checkChangeInDevelopmentMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:        "checkChangeInDevelopment",
			Aliases:     []config.Alias{},
			Description: "Checks if a certain change is in status 'in development'",
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Parameters: []config.StepParameters{
					{
						Name:        "endpoint",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS", "GENERAL"},
						Type:        "string",

						Mandatory: true,
						Default:   os.Getenv("PIPER_endpoint"),
						Aliases:   []config.Alias{{Name: "changeManagement/endpoint"}},
					},
					{
						Name:        "username",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS", "GENERAL"},
						Type:        "string",

						Mandatory: true,
						Default:   os.Getenv("PIPER_username"),
						Aliases:   []config.Alias{},
					},
					{
						Name:        "password",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "GENERAL"},
						Type:        "string",

						Mandatory: true,
						Default:   os.Getenv("PIPER_password"),
						Aliases:   []config.Alias{},
					},
					{
						Name:        "changeDocumentId",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",

						Mandatory: true,
						Default:   os.Getenv("PIPER_changeDocumentId"),
						Aliases:   []config.Alias{},
					},
					{
						Name:        "failIfStatusIsNotInDevelopment",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",

						Mandatory: false,
						Default:   true,
						Aliases:   []config.Alias{},
					},
					{
						Name:        "clientOpts",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",

						Mandatory: false,
						Default:   []string{},
						Aliases:   []config.Alias{},
					},
				},
			},
		},
	}
	return theMetaData
}
