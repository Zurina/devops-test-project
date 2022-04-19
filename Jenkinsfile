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
                    branch: "main"
                }
            }
            steps {
                build job: 'devops-deploy', parameters: [string(name: 'DEPLOY_TO', value: "qa")]
            }   
        }

        stage('Run PROD Deployment') {
            when {
                    branch: "main"
            }
            steps {
                build job: 'devops-deploy', parameters: [string(name: 'DEPLOY_TO', value: "prod")]
            }   
        }
    }
}