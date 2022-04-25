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

        stage('Run CLOUD Deployment') {
            when {
                branch "cloud"
            }
            steps {
                build job: 'devops-deploy',
                parameters: [
                    string(name: 'DEPLOY_TO', value: "cloud"),
                    string(name: 'branch', value: env.BRANCH_NAME),
                    string(name: 'host_ip', value: '107.21.71.63')
                ]
            }   
        }

        stage('Run QA Deployment') {
            when {
                branch "qa"
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