package server

import (
	"context"
	"fmt"

	"github.com/hetznercloud/hcloud-go/hcloud"
	"github.com/pkg/errors"

	infrav1 "sigs.k8s.io/cluster-api-provider-hetzner/api/v1alpha3"
)

func (s *Service) findImageIDBySpec(ctx context.Context, spec *infrav1.HetznerImageSpec) (*infrav1.HetznerImageID, error) {
	if spec == nil {
		return nil, errors.New("no image specified")
	}

	// TODO filter with label selector
	images, err := s.scope.HetznerClient().ListImages(ctx, hcloud.ImageListOpts{})
	if err != nil {
		return nil, fmt.Errorf("error listing images: %w", err)
	}

	for _, image := range images {
		imageID := infrav1.HetznerImageID(image.ID)

		// match by ID
		if specID := spec.ID; specID != nil && *specID == imageID {
			return &imageID, nil
		}

		// match by name
		if specName := spec.Name; specName != nil && *specName == image.Name {
			return &imageID, nil
		}

		// match by description
		if specName := spec.Name; specName != nil && *specName == image.Description {
			return &imageID, nil
		}

	}

	return nil, errors.New("no matching image found")
}
