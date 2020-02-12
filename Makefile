OUTFILES := $(patsubst cmd/%.go,bin/%,$(wildcard cmd/*.go))

bin/%: cmd/%.go
	go build -o $@ $<

all: $(OUTFILES)

clean:
	rm -rf $(OUTFILES)
	go mod tidy
