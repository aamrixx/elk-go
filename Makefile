build:
	go build elk_main.go elk_lexer.go elk_parser.go elk_interpreter.go elk_utils.go elk_types.go
	mv elk_main elk

run:
	$(./elk test.elk)

clean:
	rm elk