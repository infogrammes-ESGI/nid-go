BUILD_DIR=build
BINARY_NAME=nid-go
BINARY_FILE_PATH=${BUILD_DIR}/${BINARY_NAME}

run: build
	${BINARY_FILE_PATH}

build main.go: build_rule_parser
	mkdir -p build
	go build -o ${BUILD_DIR} ./...

build_rule_parser cmd/rule-parser/rule-parser.y:
	goyacc -o ./cmd/ruleparser/ruleparser.go -p RuleParser ./cmd/ruleparser/ruleparser.y

clean:
	go clean
	rm y.output
	rm -r ${BUILD_DIR}
