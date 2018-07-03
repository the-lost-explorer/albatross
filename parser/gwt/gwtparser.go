/*
 *   Copyright 2018 Amey Parundekar

 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at

 *       http://www.apache.org/licenses/LICENSE-2.0

 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package gwt

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Gwtrpcjson Custom JSON Object identifier
type Gwtrpcjson struct {
	GWTRPCVersion           int
	GWTRPCFlag              int
	GWTRPCStringTablelength int
	GWTRPCStringTable       []string
	GWTRPCPayload           []int
}

// Parse GWT and convert to JSON. RETURNS: []byte
func Parse(requestRPC string) []byte {
	//Split tokenStream
	var tokenStream []string
	tokenStream = strings.Split(requestRPC, "|")

	//Store necessary information

	/* gwtRPCVersion : The GWT RPC version of the request.
	 * latest:version 7 as of 2018
	 */
	gwtRPCVersion, err := strconv.Atoi(tokenStream[0])
	if err != nil {
		panic(fmt.Sprintf("There was a parsing error. The token stream does not contain a valid GWT version."))
	}

	/* getRPCFlag : The GWT RPC request flag value.
	 * 0 is false and anything else is true
	 */
	gwtRPCFlag, err := strconv.Atoi(tokenStream[1])
	if err != nil {
		panic(fmt.Sprintf("There was a parsing error. The token stream does not contain a valid Flag value."))
	}

	/* stringTableLength : The GWT RPC request string table length.
	 * String table is the next entitity after this in the tokenStream.
	 */
	stringTableLength, err := strconv.Atoi(tokenStream[2])
	if err != nil {
		panic(fmt.Sprintf("There was a parsing error. The token stream does not contain valid number of entities in the string table."))
	}

	/* stringTable : The GWT RPC request string table.
	 * Starts from the 4th token of the tokenStream and has a length equal to the 3rd element of tokenStream.
	 */
	stringTable := tokenStream[3 : stringTableLength+3]

	/* payloadString : Middleware function for the payload store.
	 */
	payloadString := tokenStream[stringTableLength+3:]

	/* payload : The payload which are the integer numbers after the string table in the RPC tokenStream
	 * that come after the string table and may or may not refer to the token stream value.
	 */
	var payload = []int{}

	for _, i := range payloadString {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		payload = append(payload, j)
	}

	//gwtRPCJson : Custom data type to be converted to json
	gwtRPCJson := &Gwtrpcjson{
		gwtRPCVersion,
		gwtRPCFlag,
		stringTableLength,
		stringTable,
		payload}

	//gwtRPCconvertedToJSON : JSON converted RPC in []byte format.
	gwtRPCconvertedToJSON, err := json.Marshal(gwtRPCJson)
	if err != nil {
		panic(fmt.Sprintf("There was a parsing error. Could not convert to JSON."))
	}

	return gwtRPCconvertedToJSON
}
