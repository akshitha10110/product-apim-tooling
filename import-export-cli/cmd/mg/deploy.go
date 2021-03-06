/*
*  Copyright (c) WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
*
*  WSO2 Inc. licenses this file to you under the Apache License,
*  Version 2.0 (the "License"); you may not use this file except
*  in compliance with the License.
*  You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing,
* software distributed under the License is distributed on an
* "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
* KIND, either express or implied.  See the License for the
* specific language governing permissions and limitations
* under the License.
 */

package mg

import (
	"github.com/spf13/cobra"
	"github.com/wso2/product-apim-tooling/import-export-cli/utils"
)

const (
	deployCmdLiteral   = "deploy"
	deployCmdShortDesc = "Deploy an API (apictl project) in Microgateway"
	deployCmdLongDesc  = "Deploy an API (apictl project) in Microgateway by " +
		"specifying the microgateway adapter environment."
)

const deployCmdExamples = utils.ProjectName + " " + mgCmdLiteral + " " +
	deployCmdLiteral + " " + apiCmdLiteral + " -e dev " +
	"-f petstore" +

	"\n\nNote: The flags --environment (-e), --file (-f) are mandatory. " +
	"The user needs to be logged in to use this command."

// DeployCmd represents the deploy command
var DeployCmd = &cobra.Command{
	Use:     deployCmdLiteral,
	Short:   deployCmdShortDesc,
	Long:    deployCmdLongDesc,
	Example: deployCmdExamples,
	Run: func(cmd *cobra.Command, args []string) {
		utils.Logln(utils.LogPrefixInfo + deployCmdLiteral + " called")
	},
}

// init using Cobra
func init() {
	MgCmd.AddCommand(DeployCmd)
}
