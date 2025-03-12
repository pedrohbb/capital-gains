package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/pedrohbb/capital-gains/internal/domain/operations"
)

type OpsTaxesService interface {
	ProcessOperationTaxes(operationsList []operations.Operation) operations.ChainOperations
}

type ProcessOperationsTaxes struct {
	svc OpsTaxesService
}

func (pot *ProcessOperationsTaxes) Run() {
	var opsOutput []operations.ChainOperations

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		bLine := scanner.Bytes()
		emptLine := len(string(bLine)) == 0

		var stockLineOps []operations.Operation
		if err := json.Unmarshal(bLine, &stockLineOps); err != nil && !emptLine {
			log.Fatalln("Error decoding input:", err)
		}

		opsOutput = append(opsOutput, pot.svc.ProcessOperationTaxes(stockLineOps))
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("Error reading input:", err)
	}

	for _, r := range opsOutput {
		resJSON, err := json.Marshal(r.Taxes)
		if err != nil {
			log.Fatalln("Error marshalling result:", err)

			return
		}

		fmt.Println(string(resJSON))
	}
}

func NewCliExec(svc OpsTaxesService) *ProcessOperationsTaxes {
	return &ProcessOperationsTaxes{svc: svc}
}
