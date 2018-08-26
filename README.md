# Gorut
Chilean Rut tool written in Golang. The main objectives of the project is to deliver a tool that helps to:

- Validate the rut format (Currently using this Regex: http://regexlib.com/REDetails.aspx?regexp_id=2444&AspxAutoDetectCookieSupport=1)
- Validate rut with the actual algorithm (Link: https://es.wikipedia.org/wiki/Rol_%C3%9Anico_Tributario)
- Extra methods to separate the numeric and the verifier part

The whole idea is to develop this applying good practices, tests, TDD and a good documentation. I want to encourage you to give comments or to contribute with new features o corrections. Just send me your PR's!

# Testing
You can run the test with the following command: 
```shell
go test ./...
```

# Work in progress
The work is in progress, but I will be pushing features as soon as I can. Let's start just with the Rut format verification :)
