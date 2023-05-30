pipeline {
  environment {
    image = "docker.io/srikantsarwa/gotest" + ":$BUILD_NUMBER"
    registryCredential = "registry-credentials"
    repository = 'https://github.com/ssarwa/go-hello-world.git'
    api_endpoint = 'https://us2.app.sysdig.com/'
    myimage = ''
  }
  agent any
  stages {
    stage('Cloning Git') {
      steps {
        git branch: 'main', url: repository
      }
    }
    stage('Building image') {
      steps{
        script {
          sh "docker build -t srikantsarwa/gotest ."
        }
      }
    }
    stage('Scanning Image') {
        steps {
            sysdigImageScan engineCredentialsId: 'sysdig-secure-api-token', imageName: "docker://" + image, engineURL: api_endpoint
        }
    }
    stage('Deploy Image') {
      steps{
        script {
          docker.withRegistry('', registryCredential) {
            myimage.push()
            myimage.push('latest')
          }
        }
      }
    }
  }
}