module vertual_world_system

replace github.com/RuiHirano/vertual_world_system/src/simulator => ./src/simulator

replace github.com/RuiHirano/vertual_world_system/src/monitor => ./src/monitor

replace github.com/RuiHirano/vertual_world_system/src/util => ./src/util

go 1.13

require (
	github.com/RuiHirano/vertual_world_system/src/monitor v0.0.0-00010101000000-000000000000 // indirect
	github.com/RuiHirano/vertual_world_system/src/simulator v0.0.0-00010101000000-000000000000 // indirect
	github.com/RuiHirano/vertual_world_system/src/util v0.0.0-00010101000000-000000000000 // indirect
)
