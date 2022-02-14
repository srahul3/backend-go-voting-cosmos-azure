# backend-go-voting-cosmos-azure

## Creating AZURE_CREDENTIALS using CLI

### Prerequisit 
A terminal with azure cli.

Login to the azure using cli command
az login

### Create the Azure Service Principal token using command
az ad sp create-for-rbac

The resultant JSON is the token and value to this input.

## Creating AZURE_CREDENTIALS using CLI

### Prerequisit 
Follow the steps metioned under section 'Creating AZURE_CREDENTIALS using CLI'

### Aquiring the config
az aks get-credentials --resource-group <AZURE_AKS_RESOURCE_GROUP> --name <AZURE_AKS_NAME>
The above command will save the token in a config file in your file-system.

The JSON content of this is the token and value to this input.

### Creating secret value MONGODB_CONNECTION_STRING

The connection string must be wrapped in below JSON format and used.
{
	"MONGODB_CONNECTION_STRING": "<mongo-db-connection-string>"
}


