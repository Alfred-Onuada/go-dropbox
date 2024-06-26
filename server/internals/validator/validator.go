package validator

import (
	"fmt"
	"reflect"
	"strings"
)

func Validate(rules map[string]string, params map[string]interface{}) (error) {

	for key, rule := range rules {
		value,exists := params[key]
		if !exists {
			return fmt.Errorf("%s Does not exist in params",rules[key])
		}
		conds := strings.Split(rule, "|")
		if len(conds) < 2 {
			return fmt.Errorf("Atleast 2 Arguments are required containing the type and nullability ")
		}
		nullable := conds[0]
		valtype := conds[1]
		err := handleNullability(nullable,key,value)
		if err != nil {
			return err
		}
		err = handleTypeChecking(valtype,key,value)
		if err != nil{
			return err
		}
		//TODO Add Additional functionalities for min max and other additional checks (if i have time)
		
	}
	return nil

}

func handleNullability( rule string, key string,value interface{}) (error){
	switch v :=value.(type) {
	case string:
		if rule == "required"{
			if(v == ""){
				 return fmt.Errorf("%s is required",key) 
			}
		}
		if rule == "nullable"{
			return nil
		}	
		
	}
	return nil
}

func handleTypeChecking(valtype, key string, value interface{}) error {
    valueType := reflect.TypeOf(value).Kind() 

    expectedTypes := map[string]reflect.Kind{
        "string":  reflect.String,
        "int":     reflect.Int,
        "float":   reflect.Float64,
        "bool":    reflect.Bool,
    }

    expectedType, ok := expectedTypes[valtype]
    if !ok {
        return fmt.Errorf("unsupported type: %s", valtype)
    }

    if valueType != expectedType {
        return fmt.Errorf("expected %s for %s, got %s", valtype, key, valueType)
    }

    return nil
}