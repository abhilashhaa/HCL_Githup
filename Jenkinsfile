
@Library('piper-library-os')_ 

node () 
{
  
    stage('DeployCommit') 
    gctsDeploy(
  script: this,
  host: 'https://hclutl1909.hcldigilabs.com:8001',
  abapCredentialsId: 'ABAPUserPasswordCredentialsId',
  repository: 'HCL-Githup',
  remoteRepositoryURL: "https://github.com/abhilashhaa/HCL_Githup.git",
  role: 'TARGET',
  vSID: 'FEF',
  commit: 'c630005',
  rollback: false,
  
)  
  
  stage('RunUnitTest') 
    gctsExecuteABAPUnitTests(
      script: this,
      host: 'hcluks4hana.hcldigilabs.com:8001',
      client: '200',
      abapCredentialsId: 'AbapSystem',
      repository: 'OpenSAP'
  )
       
  
    stage('RollbackCommit') 
    gctsRollback(
      script: this,
      host: "hcluks4hana.hcldigilabs.com:8001",
      client: "200",
      abapCredentialsId: 'AbapSystem',
      repository: "OpenSAP"
  )
}
