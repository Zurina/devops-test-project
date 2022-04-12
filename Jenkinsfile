pipeline {
    agent any
    tools {
        go 'go-1.17'
    }
    stages {
        stage('Compile') {
            steps {
                sh 'go build'
            }
        }
    }
}