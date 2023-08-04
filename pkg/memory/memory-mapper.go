package memory

import "fmt"

type Region struct {
	Label  string
	Device Memory
	Start  uint16
	End    uint16
	Remap  bool
}

type MemoryMapper struct {
	Regions []Region
}

func NewMemoryMapper() *MemoryMapper {
	return &MemoryMapper{}
}

func (m *MemoryMapper) Map(label string, device Memory, start uint16, size uint16, remap bool) func() {
	end := start + size - 1
	region := Region{Label: label, Device: device, Start: start, End: end, Remap: remap}
	m.Regions = append([]Region{region}, m.Regions...)

	fmt.Printf("Mapping %s from 0x%X to 0x%X\n", label, start, end)

	return func() {
		for i, r := range m.Regions {
			if r == region {
				m.Regions = append(m.Regions[:i], m.Regions[i+1:]...)
				fmt.Printf("Unmapping %s from 0x%X to 0x%X\n", label, start, end)
				break
			}
		}
	}
}

func (m *MemoryMapper) findRegion(offset uint16) (*Region, error) {
	for _, r := range m.Regions {
		if offset >= r.Start && offset <= r.End {
			return &r, nil
		}
	}

	return &Region{}, fmt.Errorf("No region found for address 0x%X", offset)
}

func (m *MemoryMapper) Read(offset uint16) byte {
	r, err := m.findRegion(offset)
	if err != nil {
		fmt.Printf("Attempted read from address 0x%X: %v", offset, err)
		return 0
	}

	if r.Remap {
		offset -= r.Start
	}

	return r.Device.Read(offset)
}

func (m *MemoryMapper) Read16(offset uint16) uint16 {
	r, err := m.findRegion(offset)
	if err != nil {
		fmt.Printf("Attempted read from address 0x%X: %v", offset, err)
		return 0
	}

	if r.Remap {
		offset -= r.Start
	}

	return r.Device.Read16(offset)
}

func (m *MemoryMapper) Write(offset uint16, value byte) {
	r, err := m.findRegion(offset)
	if err != nil {
		fmt.Printf("Attempted write to address 0x%X: %v", offset, err)
		return
	}

	if r.Remap {
		offset -= r.Start
	}

	r.Device.Write(offset, value)
}

func (m *MemoryMapper) Write16(offset uint16, value uint16) {
	r, err := m.findRegion(offset)
	if err != nil {
		fmt.Printf("Attempted write to address 0x%X: %v", offset, err)
		return
	}

	if r.Remap {
		offset -= r.Start
	}

	r.Device.Write16(offset, value)
}

func (m *MemoryMapper) Load(buffer []byte, offset uint16) {
	r, err := m.findRegion(offset)
	if err != nil {
		fmt.Printf("Attempted load to address 0x%X: %v", offset, err)
		return
	}

	if r.Remap {
		offset -= r.Start
	}

	r.Device.Load(buffer, offset)
}

func (m *MemoryMapper) Slice(start uint16, end uint16) []byte {
	r, err := m.findRegion(start)
	if err != nil {
		fmt.Printf("Attempted slice from address 0x%X: %v", start, err)
		return []byte{}
	}

	if r.Remap {
		start -= r.Start
	}

	return r.Device.Slice(start, end)
}

func (m *MemoryMapper) Reset() {
	for _, r := range m.Regions {
		r.Device.Reset()
	}
}
