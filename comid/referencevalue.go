// Copyright 2021 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package comid

import "fmt"
import "github.com/jraman567/sev-snp-measure-go/ovmf"

type ReferenceValue struct {
	_                struct{}             `cbor:",toarray"`
	Environment      Environment          `json:"environment"`
	Measurements     Measurements         `json:"measurements"`
	OvmfMetadata	 ovmf.MetadataWrapper `json:"ovmfmetadata"`
}

func (o ReferenceValue) Valid() error {
	if err := o.Environment.Valid(); err != nil {
		return fmt.Errorf("environment validation failed: %w", err)
	}

	if o.Measurements.Empty() && len(o.OvmfMetadata.OVMFHash) == 0 {
		return fmt.Errorf("Both measurements and ovmfmetadata fields can't be empty")
	}

	if err := o.Measurements.Valid(); err != nil {
		return fmt.Errorf("measurements validation failed: %w", err)
	}

	return nil
}
