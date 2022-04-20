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
            steps {
                archiveArtifacts artifacts: 'devops', onlyIfSuccessful: true
            }            
        }

        stage('Run QA Deployment') {
            when {
                not {
                    branch "main"
                }
            }
            steps {
                build job: 'devops-deploy',
                parameters: [
                    string(name: 'DEPLOY_TO', value: "qa"),
                    string(name: 'branch', value: env.BRANCH_NAME),
                    string(name: 'host_ip', value: '10.10.50.4')
                ]
            }   
        }

        stage('Run PROD Deployment') {
            when {
                branch "main"
            }
            steps {
                build job: 'devops-deploy',
                parameters: [
                    string(name: 'DEPLOY_TO', value: "prod"), 
                    string(name: 'branch', value: env.BRANCH_NAME),
                    string(name: 'host_ip', value: '10.10.50.3')
                ]
            }   
        }
    }
}