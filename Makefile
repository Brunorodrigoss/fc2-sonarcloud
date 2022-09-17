start:
	docker run -d --name sonarqube -e SONAR_ES_BOOTSTRAP_CHECKS_DISABLE=true -p 9000:9000 sonarqube:latest

go-sonar-scanner:
	go test --coverprofile=coverage.out
	sonar-scanner