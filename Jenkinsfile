pipeline {
    agent any
    
    // Posisikan environment di sini agar dibaca global oleh agent
    environment {
        GOCACHE = "${env.WORKSPACE}/.go-cache"
        GOPATH = "${env.WORKSPACE}/.go-path"
    }

    stages {
        stage('build') {
            steps {
                sh 'go mod download'
            }
        }
        
        stage('test') {
            steps {
                sh 'chmod +x ./jenkins/test.sh'
                sh './jenkins/test.sh'
            }
        }
    }
}