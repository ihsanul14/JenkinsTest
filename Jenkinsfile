pipeline {
  agent any
  stages {
    stage ('Test') {
      steps {
        echo 'Hello World'
        bat 'kubectl version'
        bat 'docker run -p 8091:8090 -d ms-go'
      }
    }
  }
}
