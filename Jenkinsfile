pipeline {
	agent any
    stages {
        stage('Test') {
            steps {
                echo 'Hello World JEnkins'
		
                bat 'docker run -p 8091:8091 ms-go'      
		
            }
        }
    }
}
