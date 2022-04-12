pipeline {
    agent { docker { image 'golang:1.17.5-alpine' } }
    stages {
        stage('Compile') {
            steps {
                sh 'GOCACHE=/tmp/'
                sh 'GOCACHE=off'
                sh 'go build'
            }
        }
    }
}