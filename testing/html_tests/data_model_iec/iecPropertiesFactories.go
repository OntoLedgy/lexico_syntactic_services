package data_model_iec

import (
	"net/url"
	"path"
	"sync"
)

type IecPropertiesFactory struct {
	PropertiesRegister map[string]*IecProperty
	mu                 sync.Mutex
}

func NewIecPropertiesFactory() *IecPropertiesFactory {
	return &IecPropertiesFactory{

		PropertiesRegister: make(map[string]*IecProperty),
	}
}

func (iecPropertiesFactory *IecPropertiesFactory) GetIecProperty(propertyUrl string) (*IecProperty, error) {

	propertyID, _ := iecPropertiesFactory.extractPropertyID(propertyUrl)

	iecPropertiesFactory.mu.Lock()
	defer iecPropertiesFactory.mu.Unlock()

	// Check if the IecClass is already loaded in memory
	if iecProperty, ok := iecPropertiesFactory.PropertiesRegister[propertyID]; ok {
		return iecProperty, nil
	}

	// If not, load the IecClass by scraping the URL
	iecProperty := iecPropertiesFactory.NewIecProperty(propertyID, propertyUrl)

	iecPropertiesFactory.PropertiesRegister[propertyID] = iecProperty

	return iecProperty, nil
}

func (iecPropertiesFactory *IecPropertiesFactory) NewIecProperty(propertyID, url string) *IecProperty {

	property := &IecProperty{
		PropertyId:  propertyID,
		PropertyUrl: url,
	}

	property.scrapePropertyPage()

	return property
}

func (iecPropertiesFactory *IecPropertiesFactory) extractPropertyID(propertyLink string) (string, error) {
	parsedUrl, err := url.Parse(propertyLink)
	if err != nil {
		return "", err
	}

	propertyID := path.Base(parsedUrl.Path)

	return propertyID, nil
}
