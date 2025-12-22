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
        echo 'Akses aplikasi di: https://intersection-go.onrender.com/'

        echo 'Pipeline dijeda selama 1 menit agar aplikasi dapat digunakan'
        sleep time: 1, unit: 'MINUTES'

        echo 'Waktu akses selesai'
        echo 'Link aplikasi ditutup: https://intersection-go.onrender.com/'
    }

}
