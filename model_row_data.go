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

type RowData struct {
	Cells []CellData `json:"cells,omitempty"`
	DeviceId int32 `json:"deviceId,omitempty"`
	DeviceDisplayName string `json:"deviceDisplayName,omitempty"`
}
