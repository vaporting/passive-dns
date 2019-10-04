package dispatcher

import (
	"passive-dns/hunter"

	"passive-dns/types"

	"encoding/json"

	"reflect"

	"strings"
	//"fmt"
)

// HunterDispatcher dispatchs the element entry to the right handler
type HunterDispatcher struct {
	hunter     map[string]hunter.IHunter
	tempHunter *hunter.SourceIPHunter
}

// Init intializes the HunterDispatcher
func (dispatcher *HunterDispatcher) Init() {
	dispatcher.hunter = make(map[string]hunter.IHunter)
	sourceIPHunter := hunter.NewSourceIPHunter()
	for _, sourceType := range sourceIPHunter.SourceTypes {
		dispatcher.hunter[sourceType] = sourceIPHunter
	}
	dispatcher.tempHunter = hunter.NewSourceIPHunter()
}

// Hunt hunt the targets from sources
func (dispatcher *HunterDispatcher) Hunt(huntingSources types.HuntingSources) (string, error) {
	results := make(map[string]interface{})
	var err error = nil
	e := reflect.ValueOf(&huntingSources).Elem()
	// dispatch sources to correspond hunter
	for i := 0; i < e.NumField(); i++ {
		sourceType := strings.ToLower(e.Type().Field(i).Name)
		sources := e.Field(i).Interface().([]string)
		result := []byte{}
		if hunter, ok := dispatcher.hunter[sourceType]; ok {
			result, err = hunter.Hunt(sources)
		}
		dat := make(map[string]interface{})
		json.Unmarshal(result, &dat)
		results[sourceType] = dat
	}
	jsonBytes, err := json.Marshal(results)
	// fmt.Println("json: ", string(jsonBytes))
	return string(jsonBytes), err
}
