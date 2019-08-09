package cell_fix_getters

type FixProcessorsFactory struct {
}

func (fix_processor_factory *FixProcessorsFactory) Create() *fixProcessors {

	fix_processor := new(fixProcessors)

	return fix_processor
}
