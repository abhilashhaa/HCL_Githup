

pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                 
                echo 'Building..'
            }
            Stages{
                stage('PullRequest')
                {
                    steps{
                        
                             githubCreatePullRequest script: this
                    }
        }
        
        stage('ATC') {
            steps {
                echo 'Testing..'
         
                
           
            }
       
       }
   
   
    }
   
}
