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

    stage('Manual Approval') {
        input message: 'Lanjutkan ke tahap Deploy?', ok: 'Proceed'
    }

    stage('Deploy') {
        echo 'Aplikasi berhasil di-deploy'
        echo 'Menjeda pipeline selama 1 menit agar aplikasi dapat digunakan'
        sleep time: 1, unit: 'MINUTES'
        echo '1 menit selesai, pipeline sukses'
    }

}
