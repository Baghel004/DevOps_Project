pipeline {
  agent any

  environment {
    SONAR_TOKEN = credentials('sonar-token-id')
  }

  stages {
    stage('Checkout') {
      steps { checkout scm }
    }

    stage('Info') {
      steps {
        bat 'where go'
        bat 'go version'
      }
    }

    stage('Test') {
      steps {
        bat 'go test ./... -coverprofile=coverage.out'
      }
    }

    stage('SonarQube Analysis') {
    steps {
        withSonarQubeEnv('SonarQube') {
            bat """
            %SONAR_SCANNER_HOME%\\bin\\sonar-scanner ^
              -Dsonar.projectKey=DevOps_Project ^
              -Dsonar.sources=. ^
              -Dsonar.host.url=%SONAR_HOST_URL% ^
              -Dsonar.login=%SONAR_TOKEN%
            """
        }
    }
}

  }
}
