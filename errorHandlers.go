package crunchyTools

import (
	"fmt"
)

// HasError Checks if the first arg is set, you must provide some context and set the silent mode. However you can choose to run a function if a error is detected, this function
// will inject a string parameter with the detail of the error
func HasError(err error, context string, silent bool, lastFunction ...func(data interface{})) error {
	if err != nil {
		logger := FetchLogger()
		errorMessage := fmt.Sprintf("Error in %s: %s\n", context, err)
		if len(lastFunction) > 0 {
			lastFunction[0](err.Error())
		}
		if !silent {
			logger.Err.Fatalf(errorMessage)
		}
		logger.Warn.Print(errorMessage)
		return err
	}
	return nil
}

// Check array of error and if found, will retour true and the error message
func HasErrorInSlice(errors []error) (bool, string) {
	for _, err := range errors {
		if err != nil {
			return true, err.Error()
		}
	}
	return false, ""
}
