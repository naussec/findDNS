/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"net"
	"github.com/spf13/cobra"
)


// getIPCmd represents the getIP command
var getIPCmd = &cobra.Command{
	Use:   "getIP",
	Short: "Provides IP of given web address",
	Long: `Search for IP address and Name Servers by webaddress.
	For example: google.com -getIP (will provide you with googles IP Address)`,
	Run: func(cmd *cobra.Command, args []string) {
		iprecords, _ := net.LookupIP(webAddress)
		for _, ip := range iprecords {
		fmt.Println(webAddress, "IP Address is", ip)
	},
}

func init() {
	rootCmd.AddCommand(getIPCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getIPCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getIPCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
