//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package objects

import "strings"

// determine which field values not to return
func determineResponseFields(fields []*string) map[string]int {
	fieldsToKeep := map[string]int{"class": 0, "properties": 0, "creationtimeunix": 0, "key": 0, "id": 0}

	if len(fields) > 0 {
		// check if "ALL" option is provided
		for _, field := range fields {
			fieldToKeep := strings.ToLower(*field)
			if fieldToKeep == "all" {
				return fieldsToKeep
			}
		}

		fieldsToKeep = make(map[string]int)
		// iterate over the provided fields
		for _, field := range fields {
			fieldToKeep := strings.ToLower(*field)
			fieldsToKeep[fieldToKeep] = 0
		}
	}

	return fieldsToKeep
}
