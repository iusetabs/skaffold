/*
Copyright 2020 The Skaffold Authors

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

package manifest

import (
	"fmt"

	sErrors "github.com/GoogleContainerTools/skaffold/pkg/skaffold/errors"
	"github.com/GoogleContainerTools/skaffold/proto/v1"
)

func replaceImageErr(err error) error {
	return sErrors.NewError(err,
		&proto.ActionableErr{
			Message: fmt.Sprintf("replacing images in manifest: %s", err),
			ErrCode: proto.StatusCode_DEPLOY_REPLACE_IMAGE_ERR,
		})
}

func transformManifestErr(err error) error {
	return sErrors.NewError(err,
		&proto.ActionableErr{
			Message: fmt.Sprintf("unable to transform manifests: %s", err),
			ErrCode: proto.StatusCode_DEPLOY_TRANSFORM_MANIFEST_ERR,
		})
}

func labelSettingErr(err error) error {
	return sErrors.NewError(err,
		&proto.ActionableErr{
			Message: fmt.Sprintf("setting labels in manifests: %s", err),
			ErrCode: proto.StatusCode_DEPLOY_SET_LABEL_ERR,
		})
}

func parseImagesInManifestErr(err error) error {
	if err == nil {
		return err
	}
	return sErrors.NewError(err,
		&proto.ActionableErr{
			Message: fmt.Sprintf("parsing images in manifests: %s", err),
			ErrCode: proto.StatusCode_DEPLOY_PARSE_MANIFEST_IMAGES_ERR,
		})
}

func writeErr(err error) error {
	return sErrors.NewError(err,
		&proto.ActionableErr{
			Message: err.Error(),
			ErrCode: proto.StatusCode_DEPLOY_MANIFEST_WRITE_ERR,
		})
}

func nsSettingErr(err error) error {
	return sErrors.NewError(err,
		&proto.ActionableErr{
			Message: fmt.Sprintf("setting namespace in manifests: %s", err),
			ErrCode: proto.StatusCode_RENDER_SET_NAMESPACE_ERR,
		})
}

func nsAlreadySetErr() error {
	err := fmt.Errorf("namespace field already set in the manifests")
	return sErrors.NewError(err,
		&proto.ActionableErr{
			Message: err.Error(),
			ErrCode: proto.StatusCode_RENDER_NAMESPACE_ALREADY_SET_ERR,
			Suggestions: []*proto.Suggestion{
				{
					SuggestionCode: proto.SuggestionCode_REMOVE_NAMESPACE_FROM_MANIFESTS,
					Action:         "remove/unset 'namespace' field in manifests and try again",
				},
			},
		})
}
