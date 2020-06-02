pipeline {
	agent {dockerfile true}
    stages {
        stage('Test') {
            steps {
                echo 'Hello World JEnkins'
		bat 'docker build . -t ms-go-app'
                bat 'docker run -p 8091:8091 ms-go'      
		
            }
        }
    }
}
