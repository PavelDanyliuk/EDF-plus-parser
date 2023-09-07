package specification

type Header struct {
	Name string
	Size int
}

var HeadersStringLike = []Header{
	{
		Name: "version",
		Size: 8,
	},
	{
		Name: "patient_id",
		Size: 80,
	},
	{
		Name: "record_id",
		Size: 80,
	},
	{
		Name: "start_date",
		Size: 8,
	},
	{
		Name: "start_time",
		Size: 8,
	},
	{
		Name: "header_Size",
		Size: 8,
	},
	{
		Name: "reserved1",
		Size: 44,
	},
	{
		Name: "number_of_records",
		Size: 8,
	},
	{
		Name: "record_duration",
		Size: 8,
	},
	{
		Name: "number_of_signals",
		Size: 4,
	},
}

var HeadersArrayLike = []Header{
	{
		Name: "labels",
		Size: 16,
	},
	{
		Name: "transducer_types",
		Size: 80,
	},
	{
		Name: "physical_dimensions",
		Size: 8,
	},
	{
		Name: "physical_minimums",
		Size: 8,
	},
	{
		Name: "physical_maximums",
		Size: 8,
	},
	{
		Name: "digital_minimums",
		Size: 8,
	},
	{
		Name: "digital_maximums",
		Size: 8,
	},
	{
		Name: "prefiltering",
		Size: 80,
	},
	{
		Name: "samples_per_record",
		Size: 8,
	},
	{
		Name: "reserved2",
		Size: 32,
	},
}
