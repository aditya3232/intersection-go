node {
    properties([
        pipelineTriggers([
            pollSCM('H/2 * * * *')
        ])
    ])

    stage('Checkout') {
        checkout scm
    }

    stage('Test') {
        docker.image('golang:1.25-alpine').inside {
            sh '''
                export GOCACHE=/tmp/go-build
                export GOPATH=/tmp/go

                go version
                go test ./... -v -json > test-report.json
            '''
        }
    }

    stage('Build') {
        docker.image('golang:1.25-alpine').inside {
            sh '''
                export GOCACHE=/tmp/go-build
                export GOPATH=/tmp/go

                go build -o intersection-app
            '''
        }
    }
}
