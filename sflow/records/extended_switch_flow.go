package records

import (
	"encoding/binary"
	"fmt"
	"io"
)

// ExtendedSwitchFlow is an extended switch flow record.
type ExtendedSwitchFlow struct {
	SourceVlan          uint32
	SourcePriority      uint32
	DestinationVlan     uint32
	DestinationPriority uint32
}

func (f ExtendedSwitchFlow) String() string {
	type X ExtendedSwitchFlow
	x := X(f)
	return fmt.Sprintf("ExtendedSwitchFlow: %+v", x)
}

// RecordName returns the Name of this flow record
func (f ExtendedSwitchFlow) RecordName() string {
	return "ExtendedSwitchFlow"
}

// RecordType returns the ID of the sflow flow record
func (f ExtendedSwitchFlow) RecordType() int {
	return TypeExtendedSwitchFlowRecord
}

func (f ExtendedSwitchFlow) calculateBinarySize() int {
	return binary.Size(f)
}

func (f ExtendedSwitchFlow) Encode(w io.Writer) error {
	var err error

	err = binary.Write(w, binary.BigEndian, uint32(f.RecordType()))
	if err != nil {
		return err
	}

	encodedRecordLength := f.calculateBinarySize()

	err = binary.Write(w, binary.BigEndian, uint32(encodedRecordLength))
	if err != nil {
		return err
	}

	return binary.Write(w, binary.BigEndian, f)
}
