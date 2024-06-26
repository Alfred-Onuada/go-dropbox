package validator

import (
	"fmt"
)

func Validate(rules map[string]string, params map[string]interface{}) (error) {

	for key, rule := range rules {
		value,exists := params[key]
		if !exists {
			return fmt.Errorf("%s Does not exist in params",rules[key])
		}
		switch v :=value.(type) {
		case string:
			if rule == "required"{
				fmt.Printf("%s %s %s %s\n",key,rule,key,value);
				if(v == ""){
					 return fmt.Errorf("%s is required",key) 
				}
			}
			
			
		}
	}
	return nil

}