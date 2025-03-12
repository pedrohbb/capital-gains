package main

import (
	"github.com/pedrohbb/capital-gains/internal/app/cli"
	"github.com/pedrohbb/capital-gains/internal/domain/operations"
	. "github.com/pedrohbb/capital-gains/pkg"
)

func main() {
	taxBucket := Float64(20000.00)
	taxRate := Float64(0.20)

	service := operations.NewService(taxBucket, taxRate)
	cli := cli.NewCliExec(service)
	cli.Run()
}

