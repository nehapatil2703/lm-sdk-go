/*
 * logicmonitor_sdk
 *
 * LogicMonitor is a SaaS-based performance monitoring platform that provides full visibility into complex, hybrid infrastructures, offering granular performance monitoring and actionable data and insights. logicmonitor_sdk enables you to manage your LogicMonitor account programmatically.
 *
 * API version: 1.0.0
 * Contact: sdk@logicmonitor.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package logicmonitor

type NextUpgradeInfo struct {
	UpgradeTime string `json:"upgradeTime,omitempty"`
	Stable bool `json:"stable,omitempty"`
	MajorVersion int32 `json:"majorVersion,omitempty"`
	MinorVersion int32 `json:"minorVersion,omitempty"`
	Mandatory bool `json:"mandatory,omitempty"`
	UpgradeTimeEpoch int64 `json:"upgradeTimeEpoch,omitempty"`
}
