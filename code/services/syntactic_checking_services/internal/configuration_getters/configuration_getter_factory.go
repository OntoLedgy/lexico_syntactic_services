package configuration_getters

type ConfigurationGetterFactories struct{}

func (ConfigurationGetterFactories) Create() *configurationGetters {

	configuration_getter :=
		new(
			configurationGetters)

	return configuration_getter
}
