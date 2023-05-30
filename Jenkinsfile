pipeline {

    parameters { 
        string(name: 'DOCKER_REPOSITORY', defaultValue: 'srikantsarwa/gotest', description: 'Name of the image to be built (e.g.: sysdiglabs/dummy-vuln-app)') 
        string(name: 'GIT_REPOSITORY', defaultValue: 'https://github.com/ssarwa/go-hello-world.git', description: 'Name of the repository with the Dockerfile to be built (e.g.: https://github.com/sysdiglabs/secure-inline-scan-examples.git)') 
        string(name: 'SYSDIG_ENDPOINT', defaultValue: 'https://us2.app.sysdig.com/', description: 'The appropriate Sysdig vulnerability scanning endpoint depending on your region, see https://docs.sysdig.com/en/docs/administration/saas-regions-and-ip-ranges (e.g.: https://github.com/sysdiglabs/secure-inline-scan-examples.git)') 
    }

  agent any
  stages {
    stage('Cloning Git') {
      steps {
        git branch: 'main', url: "${params.GIT_REPOSITORY}"
      }
    }
    stage('Build Image') {
      steps {
          sh "docker build -f ./Dockerfile -t ${DOCKER_REPOSITORY} ."
        }
    }
    stage('Scan image') {
      steps {
        withCredentials([usernamePassword(credentialsId: 'sysdig-secure-api-token', passwordVariable: 'SECURE_API_TOKEN', usernameVariable: '')]) {
          sh '''
            VERSION=$(curl -L -s https://download.sysdig.com/scanning/sysdig-cli-scanner/latest_version.txt)
            curl -LO "https://download.sysdig.com/scanning/bin/sysdig-cli-scanner/${VERSION}/linux/amd64/sysdig-cli-scanner"
            chmod +x ./sysdig-cli-scanner
            ./sysdig-cli-scanner --apiurl ${SYSDIG_ENDPOINT} docker://${DOCKER_REPOSITORY}
            '''
        }
      }
    }
    stage('Push Image') {

      steps {
        withCredentials([usernamePassword(credentialsId: 'registry-credentials', passwordVariable: 'password', usernameVariable: 'username')]){
                    sh '''
                        docker login -u ${username} -p ${password}
                        docker push ${DOCKER_REPOSITORY}
                    '''
                }        }
    }
    }
}