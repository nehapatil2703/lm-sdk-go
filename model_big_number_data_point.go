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

type BigNumberDataPoint struct {
	DataSourceId int32 `json:"dataSourceId,omitempty"`
	DataPointName string `json:"dataPointName,omitempty"`
	DataPointId int32 `json:"dataPointId,omitempty"`
	InstanceName string `json:"instanceName"`
	DataSourceFullName string `json:"dataSourceFullName,omitempty"`
	Name string `json:"name"`
	AggregateFunction string `json:"aggregateFunction,omitempty"`
	DeviceGroupFullPath string `json:"deviceGroupFullPath"`
	DeviceDisplayName string `json:"deviceDisplayName"`
}
