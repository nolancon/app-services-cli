package kafka

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

// Templates
const (
	mockedProperties = `
## Credentials for Kafka cluster: 'serviceapi' 
## Generated by rhmas cli
user=wtrocki-kafka-service
password=d0b8122f-8dfb-46b7-b68a-f5cc4e25d000
`

	mockedQuarkus = `
## Credentials for Kafka cluster: 'serviceapi' 
## Generated by rhmas cli
kafka.sasl.jaas.config=org.apache.kafka.common.security.plain.PlainLoginModule required username="wtrocki-kafka-service" password="d237abba-d098-4e3c-8ee3-badd453d6245";
kafka.sasl.mechanism=PLAIN
kafka.security.protocol=SASL_SSL
kafka.ssl.protocol=TLSv1.2
`

	mockedJSON = `
{ 
	"name":"service-api",
	"user":"wtrocki-kafka-service", 
	"password":"d0b8122f-8dfb-46b7-b68a-f5cc4e25d000" 
}
`
)

var outputFlagValue string

// NewGetCommand gets a new command for getting kafkas.
func NewCredentialsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "credentials",
		Short: "Generate credentials to connect to your cluster",
		Long:  "Generate credentials to connect your application to the Kafka cluster",
		Run:   runCredentials,
	}
	cmd.Flags().StringVarP(&outputFlagValue, "output", "o", "properties", "Format of the config [quarkus, properties, json]")
	return cmd
}

func runCredentials(cmd *cobra.Command, _ []string) {
	var propertyFormat string
	var fileName string

	if outputFlagValue == "properties" {
		propertyFormat = mockedProperties
		fileName = "kafka.properties"
	} else if outputFlagValue == "quarkus" {
		propertyFormat = mockedQuarkus
		fileName = "kafka.properties"
	} else if outputFlagValue == "json" {
		propertyFormat = mockedJSON
		fileName = "credentials.json"
	}

	dataToWrite := []byte(propertyFormat)
	err := ioutil.WriteFile(fileName, dataToWrite, 0644)
	if err != nil {
		fmt.Println("Error when saving file:", err)
	} else {
		fmt.Println("Successfully saved credentials to", fileName)
	}
}
