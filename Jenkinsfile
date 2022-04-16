#!/usr/bin/env groovy

pipeline {
    agent any

    stages {
        stage('Test') {                
            steps {      
                sh 'go test'
            }            
        }
        // stage('Build') {                
        //     steps {      
        //         sh 'go build -o devops main.go'
        //     }            
        // }
        // stage('Save Artifacts') {    
        //     steps {
        //         archiveArtifacts artifacts: 'devops', onlyIfSuccessful: true
        //     }            
        // }

        stage('DockerHub Auth') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'docker-auth', usernameVariable: 'username', passwordVariable: 'password')]) {
                sh 'docker login --username ${username} --password ${password}'
                }
            }   
        }
        stage('Building Docker image') {
            steps {
                sh 'docker build . --tag immassive/devops:${BUILD_ID}'
            }
        }
        stage('Pushing Docker image') {
            steps {
                sh 'docker push immassive/devops:${BUILD_ID}'
            }
        }
    }
}