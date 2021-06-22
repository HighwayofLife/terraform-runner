def registryCredentialId = "buildpipeline-acr-cred"

pipeline {
  agent {
    kubernetes {
      cloud 'build-pipeline'
      defaultcontainer 'kaniko'
      inheritFrom 'jnlp'
      yaml '''
apiVersion: v1
kind: Pod
metadata:
  labels:
    app-build: terraform-runner-api
spec:
  containers:
    - name: go
      image: buildpipeline.azurecr.io/build-images/go-build:1.16.5
      command:
        - cat
      tty: true
    - name: kaniko
      image: buildpipeline.azurecr.io/build-images/kaniko:v1.5.2 
      command:
        - /busybox/cat
      tty: true
'''
    } // kubernetes
  } // agent

  options {
    ansiColor('xterm')
  }

  environment {
    PLUGIN_REPO     = "terraform"
    PLUGIN_APP_NAME = "terraform-runner-api"
    PLUGIN_REGISTRY = "buildpipeline.azurecr.io"
  }

  stages {
    stage('go test') {
      container('go') {
        sh 'go test --cover -v ./cmd/...'
      }
    } // stage(go test)

    stage('build test') {
      when {
        not {
          branch "master"
        }
      }
      container('kaniko') {
        // Builds but don't push the image (no .tags)
        sh 'kaniko-build'
        sh 'kaniko-versions'
        sh 'echo "No Tags" > VERSIONS'
      } // container(kaniko)
    } // stage(Build test)

    stage('Build and Push') {
      when {
        branch "master"
      }
      container('kaniko') {
        // Build and push the image
        withCredentials([usernamePassword(
                          credentialsId: registryCredentialId,
                          passwordVariable: "PLUGIN_PASSWORD",
                          usernameVariable: "PLUGIN_USERNAME")]) {
          // Build and push to registry
          sh 'create-tags'
          sh 'kaniko-build'
          sh 'kaniko-versions'
        } // withCredentials
      } // container(kaniko)
    } // stage(build and push)
  } // stages
} // pipeline