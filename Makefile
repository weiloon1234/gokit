.PHONY: entity-update help

entity-update:
	GOPROXY=direct go get -u github.com/weiloon1234/gokit-base-entity@latest

# Help target
help:
	@echo "Available targets:"
	@echo "  entity-update        To update to latest gokit-base-entity"