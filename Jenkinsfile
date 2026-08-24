pipeline {
    agent {
        docker {
            image 'golang:1.26.5'
            args '-p 3000:3000'
        }
    }

    stages {
        stage('build') {
            steps {
                sh 'go mod download'
            }
        }
    }
}