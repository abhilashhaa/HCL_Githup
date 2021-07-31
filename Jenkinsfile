pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                 echo 'Building..'
                githubCreatePullRequest script: this
                host: 'https://hcluks4hana.hcldigilabs.com:8001/',
                client: '200',
                abapCredentialsId: 'ABAPUserPasswordCredentialsId',
                repository: 'HCL-Githup',
                remoteRepositoryURL: "https://github.com/abhilashhaa/HCL_Githup.git",
               
                gctsDeploy(
                 script: this,
                host: 'https://hclutl1909.hcldigilabs.com:8001/',
                client: '200',
                abapCredentialsId: 'ABAPUserPasswordCredentialsId',
                repository: 'HCL-Githup',
                remoteRepositoryURL: "https://github.com/abhilashhaa/HCL_Githup.git",
                role: 'SOURCE',
                vSID: 'FEF',
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
                
            }
        }
    }
}
