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

type PieChartInfo struct {
	MaxSlicesCanBeShown int32 `json:"maxSlicesCanBeShown,omitempty"`
	VirtualDataPoints []VirtualDataPoint `json:"virtualDataPoints,omitempty"`
	ShowLabelsAndLinesOnPC bool `json:"showLabelsAndLinesOnPC,omitempty"`
	Counters []Counter `json:"counters,omitempty"`
	DataPoints []PieChartDataPoint `json:"dataPoints,omitempty"`
	HideZeroPercentSlices bool `json:"hideZeroPercentSlices,omitempty"`
	GroupRemainingAsOthers bool `json:"groupRemainingAsOthers,omitempty"`
	PieChartItems []PieChartItem `json:"pieChartItems"`
	Title string `json:"title,omitempty"`
}
