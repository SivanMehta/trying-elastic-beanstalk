.PHONY:
serve:
	go run application.go

application.zip:
	zip -r application.zip .