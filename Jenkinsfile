pipeline {
    agent any

    stages {
        stage('Test') {                
            steps {      
                sh 'go test'
            }            
        }

        stage('DockerHub Auth') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'docker-auth', usernameVariable: 'username', passwordVariable: 'password')]) {
                sh 'echo ${password} | docker login --username ${username} --password-stdin'
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