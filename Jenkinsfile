#!/usr/bin/env groovy

pipeline {
    agent { docker { image 'golang' } }    

    stages {
        stage('Build') {                
            steps {      
                sh 'cd ${GOPATH}/src'
                sh 'mkdir -p ${GOPATH}/src/devops'

                sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/devops'
                sh 'go clean -cache'
                sh 'go build'
            }            
        }
    }
}