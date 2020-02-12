OUTFILES := $(patsubst cmd/%.go,bin/%,$(wildcard cmd/*.go))

bin:
	mkdir bin

bin/%: cmd/%.go
	go build -o $@ $<

all: bin $(OUTFILES)

clean:
	rm -rf $(OUTFILES)
