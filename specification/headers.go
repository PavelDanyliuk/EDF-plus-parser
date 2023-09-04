package specification

type Header struct {
	IsArray  bool
	Position int
	Size     int
}

var HeadersSpecification = map[string]Header{
	"version": {
		IsArray:  false,
		Position: 0,
		Size:     8,
	},
	"patient_id": {
		IsArray:  false,
		Position: 1,
		Size:     80,
	},
	"record_id": {
		IsArray:  false,
		Position: 2,
		Size:     80,
	},
	"start_date": {
		IsArray:  false,
		Position: 3,
		Size:     8,
	},
	"start_time": {
		IsArray:  false,
		Position: 4,
		Size:     8,
	},
	"header_Size": {
		IsArray:  false,
		Position: 5,
		Size:     8,
	},
	"reserved1": {
		IsArray:  false,
		Position: 6,
		Size:     44,
	},
	"number_of_records": {
		IsArray:  false,
		Position: 7,
		Size:     8,
	},
	"record_duration": {
		IsArray:  false,
		Position: 8,
		Size:     8,
	},
	"number_of_signals": {
		IsArray:  false,
		Position: 9,
		Size:     4,
	},
	"labels": {
		IsArray:  true,
		Position: 10,
		Size:     16,
	},
	"transducer_types": {
		IsArray:  true,
		Position: 11,
		Size:     80,
	},
	"physical_dimensions": {
		IsArray:  true,
		Position: 12,
		Size:     8,
	},
	"physical_minimums": {
		IsArray:  true,
		Position: 13,
		Size:     8,
	},
	"physical_maximums": {
		IsArray:  true,
		Position: 14,
		Size:     8,
	},
	"digital_minimums": {
		IsArray:  true,
		Position: 15,
		Size:     8,
	},
	"digital_maximums": {
		IsArray:  true,
		Position: 16,
		Size:     8,
	},
	"prefiltering": {
		IsArray:  true,
		Position: 17,
		Size:     80,
	},
	"samples_per_record": {
		IsArray:  true,
		Position: 18,
		Size:     8,
	},
	"reserved2": {
		IsArray:  true,
		Position: 19,
		Size:     32,
	},
}

var HeadersOrder = []string{
	"version",
	"patient_id",
	"record_id",
	"start_date",
	"start_time",
	"header_Size",
	"reserved1",
	"number_of_records",
	"record_duration",
	"number_of_signals",
	"labels",
	"transducer_types",
	"physical_dimensions",
	"physical_minimums",
	"physical_maximums",
	"digital_minimums",
	"digital_maximums",
	"prefiltering",
	"samples_per_record",
	"reserved2",
}

type ParsedHeaders map[string][]string
