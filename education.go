package main

import (
	"fmt"
)

// ===============================
// Educational & Defensive Analysis
// ===============================

// AnalyzeTraffic explains latency and status behavior.
// This function is OFFLINE and EDUCATIONAL.
func AnalyzeTraffic(latencies []int, statuses []int) {
	fmt.Println("\n=== Traffic Behavior Analysis ===")

	if len(latencies) == 0 || len(statuses) == 0 {
		fmt.Println("No data provided.")
		return
	}

	sum := 0
	min := latencies[0]
	max := latencies[0]

	for _, l := range latencies {
		sum += l
		if l < min {
			min = l
		}
		if l > max {
			max = l
		}
	}

	fmt.Println("Total events observed :", len(statuses))
	fmt.Println("Average latency (ms)  :", sum/len(latencies))
	fmt.Println("Min latency (ms)      :", min)
	fmt.Println("Max latency (ms)      :", max)

	statusMap := map[int]int{}
	for _, s := range statuses {
		statusMap[s]++
	}

	fmt.Println("\nHTTP Status Distribution:")
	for code, count := range statusMap {
		fmt.Printf("  %d -> %d\n", code, count)
	}

	if statusMap[429] > 0 {
		fmt.Println("\n• Rate limiting detected (HTTP 429)")
	}
	if statusMap[403] > 0 {
		fmt.Println("• Blocking detected (HTTP 403)")
	}
	if statusMap[500] > 0 {
		fmt.Println("• Possible backend saturation (HTTP 5xx)")
	}
}

// ExplainDefenses gives defensive interpretation.
func ExplainDefenses(statuses []int) {
	fmt.Println("\n=== Defensive Interpretation ===")

	for _, s := range statuses {
		if s == 429 {
			fmt.Println("Likely Defense: Rate Limiter / API Gateway")
			return
		}
		if s == 403 {
			fmt.Println("Likely Defense: WAF / Firewall")
			return
		}
	}

	fmt.Println("No visible defensive reaction detected.")
}

// ExplainPatterns teaches traffic patterns.
func ExplainPatterns() {
	fmt.Println("\n=== Common Traffic Patterns ===")
	fmt.Println("• Burst  : Sudden spike, often triggers rate limiting.")
	fmt.Println("• Ramp   : Gradual increase, tests adaptive defenses.")
	fmt.Println("• Steady : Constant pressure, exposes capacity limits.")
}

// EthicalWarning prints usage warning.
func EthicalWarning() {
	fmt.Println("\n=== Ethical & Legal Warning ===")
	fmt.Println("This project is for EDUCATIONAL and DEFENSIVE RESEARCH only.")
	fmt.Println("Run traffic-generation code ONLY on systems you own")
	fmt.Println("or where you have explicit written permission.")
	fmt.Println("Unauthorized use may be illegal.")
}

// -------------------------------
// Safe Demo (no networking)
// -------------------------------
func main() {
	EthicalWarning()
	ExplainPatterns()

	latencies := []int{120, 140, 160, 90, 95, 80}
	statuses := []int{200, 200, 200, 429, 429, 403}

	AnalyzeTraffic(latencies, statuses)
	ExplainDefenses(statuses)
}
