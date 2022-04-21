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

        stage('Run QA Deployment to k8s') {
            when {
                not {
                    branch "main"
                }
            }
            steps {
                build job: 'devops-deploy-k8s',
                parameters: [
                    string(name: 'DEPLOY_TO', value: "qa"),
                    string(name: 'IMAGE_ID', value: $BUILD_ID)
                ]
            }   
        }

        stage('Run PROD Deployment to k8s') {
            when {
                branch "main"
            }
            steps {
                build job: 'devops-deploy-k8s',
                parameters: [
                    string(name: 'DEPLOY_TO', value: "prod"),
                    string(name: 'IMAGE_ID', value: $BUILD_ID)
                ]
            }   
        }
    }
}