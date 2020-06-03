pipeline {
  agent {
    kubernetes {
      yamlFile 'ms-go.yaml'
    }
  }
  stages {
    stage ('Test') {
      steps {
        echo 'Hello World'
        bat 'docker run -p 8091:8090 -d ms-go'
      }
    }
  }
}
