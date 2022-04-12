#!/usr/bin/env groovy

pipeline {
    agent any

    stages {
        stage('Test') {                
            steps {      
                sh 'go test'
            }            
        }
        stage('Build') {                
            steps {      
                sh 'go build -o devops main.go'
            }            
        }
        stage('Save Artifacts') {                
            archiveArtifacts artifacts: 'devops', onlyIfSuccessful: true
        }
    }
}