pipeline {
    agent any

    tools {
        go 'go-1.22'
    }

    stages{
        stage('Dev'){
            steps{
                git 'https://github.com/saltcodes/ezyevent-api.git'
                sh 'go build cmd/ezyevent-api/main.go'
            }
        }
    }
}