/**
 * Copyright 2022 MobiledgeX, Inc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

variable "environ_tag" {
  description = "Name to tag instances created by this workspace"
  default     = "qa"
}

variable "azure_location" {
  description = "Name of the Azure resource group for the cluster"
  default     = "West US 2"
}

variable "azure_terraform_service_principal_id" {
  description = "Azure service principal client ID"
  type        = string
}

variable "azure_terraform_service_principal_secret" {
  description = "Azure service principal client secret"
  type        = string
}

variable "azure_subscription_id" {
  description = "Azure subscription ID"
  type        = string
}

variable "azure_tenant_id" {
  description = "Azure tenant ID"
  type        = string
}

variable "cloudflare_account_email" {
  description = "Cloudflare account email"
  type        = string
}

variable "cloudflare_account_api_token" {
  description = "Cloudflare account API token"
  type        = string
}

variable "cluster_name" {
  default = "mexplat-qa"
}

variable "resource_group_name" {
  default = "mexplat-qa-rg"
}

variable "address_space" {
  default = "172.30.0.0/24"
}

variable "azure_vm_size" {
  default = "Standard_DS1_v2"
}

variable "gcp_project" {
  default = "still-entity-201400"
}

variable "gcp_zone" {
  default = "us-west2-a"
}

variable "gitlab_instance_name" {
  default = "gitlab-qa"
}

variable "console_instance_name" {
  default = "console-qa"
}

variable "vault_b_instance_name" {
  default = "vault-qa-b"
}

variable "vault_b_gcp_zone" {
  default = "europe-west3-a"
}

// DNS entries

variable "vault_vm_domain_name" {
  description = "Vault domain name"
  default     = "vault-qa"
}

variable "gitlab_domain_name" {
  description = "Gitlab domain name"
  default     = "gitlab-qa.mobiledgex.net"
}

variable "gitlab_docker_domain_name" {
  description = "Gitlab docker repo domain name"
  default     = "docker-qa.mobiledgex.net"
}

variable "console_domain_name" {
  description = "Console domain name"
  default     = "console-qa"
}

variable "console_vnc_domain_name" {
  description = "Console VNC domain name"
  default     = "console-qa-vnc"
}

variable "notifyroot_domain_name" {
  description = "Notifyroot service domain name"
  default     = "notifyroot-qa"
}

variable "jaeger_domain_name" {
  default = "jaeger-qa"
}

variable "esproxy_domain_name" {
  default = "events-qa.es.mobiledgex.net"
}

variable "alertmanager_domain_name" {
  default = "alertmanager-qa.mobiledgex.net"
}

variable "crm_vm_domain_name" {
  description = "CRM VM domain name"
  default        = "crm-qa.mobiledgex.net"
}

variable "postgres_domain_name" {
  description = "Postgres domain name"
  default     = "postgres-qa.mobiledgex.net"
}

variable "vault_a_domain_name" {
  default = "vault-qa-a.mobiledgex.net"
}

variable "vault_b_domain_name" {
  default = "vault-qa-b.mobiledgex.net"
}

variable "vault_c_domain_name" {
  default     = "vault-qa-c.mobiledgex.net"
}

variable "kafka_instance_name" {
  default     = "kafka-qa"
}

variable "kafka_domain_name" {
  default     = "kafka-qa.mobiledgex.net"
}

variable "ssh_public_key_file" {
  description = "SSH public key file for the ansible account"
  type        = string
  default     = "~/.mobiledgex/id_rsa_mex.pub"
}

