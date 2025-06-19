pipeline {
  agent { label 'vps-agent1' }
  
  environment {
    GIT_COMMIT_SHORT = sh(script: "git rev-parse --short HEAD", returnStdout: true).trim()
    DOCKER_HUB_USER = "m1kky8"
    DOCKER_REPO = "osi-bot"
    DOCKERHUB_CREDS = "580b959d-d40a-422f-a3d7-cf11b2ec7a4c"

  }

  stages {
    stage('Build') {
      steps {
        script {
          def branch = env.BRANCH_NAME
          def imageTag = "${DOCKER_HUB_USER}/${DOCKER_REPO}:${GIT_COMMIT_SHORT}"
          def latestTag = "${DOCKER_HUB_USER}/${DOCKER_REPO}:latest"

          withCredentials([usernamePassword(
            credentialsId: DOCKERHUB_CREDS,
            usernameVariable: 'DOCKER_USER',
            passwordVariable: 'DOCKER_PASS'
          )]) {
            echo "Logging in to Docker Hub..."
            sh 'echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin'
            
            echo "Building Docker image ${imageTag}..."
            sh "docker build -t ${imageTag} ."

            echo "Tagging as latest..."
            sh "docker tag ${imageTag} ${latestTag}"

            echo "Pushing Docker image ${imageTag}..."
            sh "docker push ${imageTag}"

            echo "Pushing latest tag"
            sh "docker push ${latestTag}"
          }
        }
      }
    }
  }
}
