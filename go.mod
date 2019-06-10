module github.com/ian-howell/rootCmd

go 1.12

require (
	github.com/ian-howell/barCommand v0.0.0-00010101000000-000000000000
	github.com/ian-howell/fooCommand v0.0.0-20190610162611-3a1132d8ab32
	github.com/spf13/cobra v0.0.5
)

replace github.com/ian-howell/barCommand => github.com/ian-howell/fooCommand v0.0.0-20190610162807-c271ada0b97c
