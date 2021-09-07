/*
Copyright 2021 Tim Seagren.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// HarborProvisionerSpec defines the desired state of HarborProvisioner
type HarborProvisionerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of HarborProvisioner. Edit HarborProvisioner_types.go to remove/update
	Aws    HarborProvisionerAwsSpec    `json:"aws,omitempty"`
	Harbor HarborProvisionerHarborSpec `json:"harbor,omitempty"`
}
type HarborProvisionerAwsSpec struct {
	Region string `json:"region,omitempty"`
}
type HarborProvisionerHarborSpec struct {
	Url                       string                        `json:"url,omitempty"`
	ProjectsEndpint           string                        `json:"projectsEndpoint,omitempty"`
	UserEndpoint              string                        `json:"usersEndpoint,omitempty"`
	RobotsEndpoint            string                        `json:"robotsEndpoint,omitempty"`
	ConifugrationEndpoint     string                        `json:"configurationEndpoint,omitempty"`
	RetentionPoliciesEndpoint string                        `json:"retentionPoliciesEndpoint,omitempty"`
	Username                  string                        `json:"username,omitempty"`
	Password                  string                        `json:"password,omitempty"`
	AwsSecretName             string                        `json:"awsSecretName,omitempty"`
	HarborProjects            []HarborProjectsSpec          `json:"projects,omitempty"`
	Robots                    []HarborRobotSpec             `json:"robots,omitempty"`
	RetentionPolicies         []HarborRetentionPoliciesSpec `json:"retentionPolicies,omitempty"`
	//   # oidc updated values
	//   oidc:
	//     auth_mode: "oidc_auth"
	//     oidc_client_id: "harbor"
	//     oidc_endpoint: "https://login.dso.mil/auth/realms/baby-yoda"
	//     oidc_groups_claim: "group-simple"
	//     oidc_name: "P1 SSO"
	//     oidc_scope: "openid,profile,email,offline_access"
}

type HarborRetentionPoliciesSpec struct {
	Algorithm string                             `json:"algorithm,omitempty"`
	Id        string                             `json:"id,omitempty"`
	Rules     []HarborRetentionPoliciesRulesSpec `json:"rules,omitemtpy"`
	Scope     HarborRetentionPoliciesScopeSpec   `json:"scope,omitempty"`
}

type HarborRetentionPoliciesScopeSpec struct {
	Trigger HarborRetentionPoliciesScopeTriggerSpec `json:"trigger,omitempty"`
}

type HarborRetentionPoliciesScopeTriggerSpec struct {
	Kind     string                                          `json:"kind,omitempty"`
	Settings HarborRetentionPoliciesScopeTriggerSettingsSpec `json:"settings,omitempty"`
}

type HarborRetentionPoliciesScopeTriggerSettingsSpec struct {
	Cron string `json:"cron,omitempty"`
}

type HarborRetentionPoliciesScopeSpec struct {
	Level string `json:"level,omitempty"`
	Ref   string `json:"ref,omitempty"`
}
type HarborRetentionPoliciesRulesSpec struct {
	Action         string                                     `json:"action,omitempty"`
	ScopeSelectors HarborRetentionPoliciesRulesScopeSelectors `json:"scopeSelectors,omitempty"`
	TagSelectors   HarborRetentionPoliciesRulesTagSelectors   `json:"tagSelectors,omitempty"`
	//           tag_selectors:
	//             - decoration: "matches"
	//               extras: '{"untagged":false}'
	//               kind: "doublestar"
	//               pattern: "**"
	//           template: "always"
	//         - action: retain
	//           scope_selectors:
	//             repository:
	//               - decoration: "repoMatches"
	//                 kind: "doublestar"
	//                 pattern: "**"
	//           tag_selectors:
	//             - decoration: "excludes"
	//               extras: '{"untagged":true}'
	//               kind: "doublestar"
	//               pattern: "**"
	//           template: "always"
}

type HarborRetentionPoliciesRulesScopeSelectors struct {
	//             repository:
	//               - decoration: "repoMatches"
	//                 kind: "doublestar"
	//                 pattern: "**"
}

type HarborRobotSpec struct {
	Name        string                       `json:"name,omitempty"`
	Level       string                       `json:"level,omitempty"`
	Duration    int                          `json:"duration,omitempty"`
	Description string                       `json:"description,omitempty"`
	permission  []HarborRobotPermissionsSpec `json:"permissions,omitempty"`
}

type HarborRobotPermissionsSpec struct {
	Access    []HarborRobotPermissionsAccessSpec `json:"access,omitempty"`
	Kind      string                             `json:"kind,omitempty"`
	Namespace string                             `json:"namespace,omitempty"`
}

type HarborRobotPermissionsAccessSpec struct {
	Action   string `json:"action,omitempty"`
	Resource string `json:"resource,omitempty"`
	Effect   string `json:"effect,omitempty"`
}

type HarborProjectsSpec struct {
	ProjectName string                     `json:"projectName,omitempty"`
	MetaData    HarborProjectsMetadataSpec `json:"metadata,omitempty"`
}

type HarborProjectsMetadataSpec struct {
	Public                bool   `json:"public,omitempty"`
	AutoScan              bool   `json:"autoScan,omitempty"`
	EnableContentTrust    bool   `json:"enableContentTrust,omitempty"`
	PreventVul            bool   `json:"preventVul,omitempty"`
	RetentionId           int    `json:"retentionId,omitempty"`
	ReuseSysCveWhistelist bool   `json:"reuseSysCveWhitelist,omitempty"`
	Severity              string `json:"severity,omitempty"`
}

// HarborProvisionerStatus defines the observed state of HarborProvisioner
type HarborProvisionerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// HarborProvisioner is the Schema for the harborprovisioners API
type HarborProvisioner struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HarborProvisionerSpec   `json:"spec,omitempty"`
	Status HarborProvisionerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HarborProvisionerList contains a list of HarborProvisioner
type HarborProvisionerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HarborProvisioner `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HarborProvisioner{}, &HarborProvisionerList{})
}
