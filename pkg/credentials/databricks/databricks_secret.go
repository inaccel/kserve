/*
Copyright 2021 The KServe Authors.

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

package databricks

import (
	"github.com/kserve/kserve/pkg/credentials/mlflow"
	corev1 "k8s.io/api/core/v1"
)

const (
	Host              = "DATABRICKS_HOST"
	Insecure          = "DATABRICKS_INSECURE"
	MLflowTrackingURI = "databricks"
	Password          = "DATABRICKS_PASSWORD"
	Token             = "DATABRICKS_TOKEN"
	Username          = "DATABRICKS_USERNAME"
)

var EnvKeys = []string{Host, Insecure, Password, Token, Username}

func BuildSecretEnvs(secret *corev1.Secret) []corev1.EnvVar {
	var envs []corev1.EnvVar
	envs = append(envs, corev1.EnvVar{
		Name:  mlflow.TrackingURI,
		Value: MLflowTrackingURI,
	})
	for _, envKey := range EnvKeys {
		if _, ok := secret.Data[envKey]; ok {
			envs = append(envs, corev1.EnvVar{
				Name: envKey,
				ValueFrom: &corev1.EnvVarSource{
					SecretKeyRef: &corev1.SecretKeySelector{
						Key: envKey,
						LocalObjectReference: corev1.LocalObjectReference{
							Name: secret.Name,
						},
					},
				},
			})
		}
	}
	return envs
}
