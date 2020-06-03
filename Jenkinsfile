pipeline {
  agent {
    kubernetes {
      yaml """
apiVersion: v1
kind: Pod
metadata:
  labels:
    some-label: some-label-value
spec:
  containers:
  - name: ms-go-v1
    image: ihsanul14/microsergo:tesgo
    command:
    - cat
    tty: true
"""
    }
  }
  stages {
    stage('Run maven') {
      steps {
        container('ms-go-v1') {
          sh 'echo Hello World'
        }
      }
    }
  }
}
