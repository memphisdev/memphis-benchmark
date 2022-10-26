def gitBranch = env.BRANCH_NAME
def imageName = "memphis-benchmark"
def gitURL = "git@github.com:Memphisdev/memphis-benchmark.git"
def repoUrlPrefix = "memphisos"

node {
  git credentialsId: 'main-github', url: gitURL, branch: 'master'
  def versionTag = readFile "./version.conf"
	
  try{
	  
    stage('Login to Docker Hub') {
      withCredentials([usernamePassword(credentialsId: 'docker-hub', usernameVariable: 'DOCKER_HUB_CREDS_USR', passwordVariable: 'DOCKER_HUB_CREDS_PSW')]) {
        sh 'docker login -u $DOCKER_HUB_CREDS_USR -p $DOCKER_HUB_CREDS_PSW'
      }
    }
	  
    stage('Build and push docker image to Docker Hub') {
      sh "docker buildx build --push --tag ${repoUrlPrefix}/${imageName}:${versionTag} --tag ${repoUrlPrefix}/${imageName} --platform linux/amd64,linux/arm64 ." 
    }

    stage('Install terraform'){
      sh 'sudo yum install -y yum-utils'
      sh 'sudo yum-config-manager --add-repo https://rpm.releases.hashicorp.com/AmazonLinux/hashicorp.repo'
      sh 'sudo yum -y install terraform'
    }

    stage('Deploy new K8s+Memphis cluster'){
      dir ('memphis-terraform'){
        git credentialsId: 'main-github', url: 'git@github.com:memphisdev/memphis-terraform.git', branch: 'benchmark'
	git credentialsId: 'main-github', url: 'git@github.com:memphisdev/memphis-k8s.git', branch: 'benchmark'
        sh 'cd ./AWS/EKS/ && make infra'
      }
    }
    
	  
    stage('Deploy memphis benchmark tool'){
      dir ('memphis-benchmark'){
        git credentialsId: 'main-github', url: 'git@github.com:memphisdev/memphis-benchmark.git', branch: 'master'
        sh 'kubectl create ns memphis-benchmark'
	sh 'kubectl apply -f deployment.yaml'
      }
    }
	  
    notifySuccessful()
	  
  } catch (e) {
      currentBuild.result = "FAILED"
      cleanWs()
      notifyFailed()
      throw e
  }
}

def notifySuccessful() {
  emailext (
      subject: "SUCCESSFUL: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]'",
      body: """<p>SUCCESSFUL: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]':</p>
        <p>Check console output at &QUOT;<a href='${env.BUILD_URL}'>${env.JOB_NAME} [${env.BUILD_NUMBER}]</a>&QUOT;</p>""",
      recipientProviders: [[$class: 'DevelopersRecipientProvider']]
    )
}

def notifyFailed() {
  emailext (
      subject: "FAILED: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]'",
      body: """<p>FAILED: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]':</p>
        <p>Check console output at &QUOT;<a href='${env.BUILD_URL}'>${env.JOB_NAME} [${env.BUILD_NUMBER}]</a>&QUOT;</p>""",
      recipientProviders: [[$class: 'DevelopersRecipientProvider']]
    )
}
