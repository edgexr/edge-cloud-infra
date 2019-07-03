variable "environ_tag" {
  description = "Name to tag instances created by this workspace"
  default     = "mexdemo"
}

variable "azure_location" {
  description = "Name of the Azure resource group for the cluster"
  default     = "Central US"
}

variable "azure_eu_location" {
  description = "Name of the Azure resource group for the EU cluster"
  default     = "West Europe"
}

variable "azure_terraform_service_principal_id" {
  description = "Azure service principal client ID"
  type        = "string"
}

variable "azure_terraform_service_principal_secret" {
  description = "Azure service principal client secret"
  type        = "string"
}

variable "azure_subscription_id" {
  description = "Azure subscription ID"
  type        = "string"
}

variable "azure_tenant_id" {
  description = "Azure tenant ID"
  type        = "string"
}

variable "cloudflare_account_email" {
  description = "Cloudflare account email"
  type        = "string"
}

variable "cloudflare_account_api_token" {
  description = "Cloudflare account API token"
  type        = "string"
}

variable "cluster_name" {
  default     = "mexdemo2-cluster"
}

variable "eu_cluster_name" {
  default     = "mexdemo-eu-cluster"
}

variable "resource_group_name" {
  default     = "mexdemo2-resource-group"
}

variable "eu_resource_group_name" {
  default     = "mexdemo-eu-resource-group"
}

variable "azure_vm_size" {
  default     = "Standard_DS1_v2"
}

variable "gcp_project" {
  default     = "still-entity-201400"
}

variable "gcp_zone" {
  default     = "us-west2-a"
}