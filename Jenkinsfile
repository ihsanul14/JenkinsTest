pipeline {
  agent {
    kubernetes {
      yaml """
        apiVersion: v1
        kind: Pod
        metadata:
        labels:
          component: ci
        spec:
          serviceAccountName: cd-jenkins
          containers:
          - name: node
            image: node:12.6.0
            command:
            - cat
            tty: true
          - name: gcloud
            image: google/cloud-sdk:latest
            command:
            - cat
            tty: true
          - name: helm
            image: alpine/helm:2.14.0
            command:
            - cat
            tty: true
          - name: jnlp
            image: mfahry/bri-jnlp-slave:1.7
        """
    }
  }
  stages {
        stage('Test') {
            steps {
                echo 'Hello World'
                sh 'ENV Port = $PORT'
            }
        }
    }
}
