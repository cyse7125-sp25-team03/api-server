node {

    stage('Clone repository') {
        checkout scm
    }

    stage('Semantic Release') {
            withCredentials([usernamePassword(credentialsId: 'github-pat', usernameVariable: 'GITHUB_USERNAME', passwordVariable: 'GITHUB_TOKEN')]) {
                script {
                    sh "npx semantic-release"
                }
            }    
        }
    stage('Build and Push multi-platform image') {
        withCredentials([usernamePassword(credentialsId: 'docker-pat', usernameVariable: 'DOCKER_USERNAME', passwordVariable: 'DOCKER_TOKEN')]) {
            script{
                // Get the latest version
                def LATEST_TAG = sh(script: "git describe --tags --abbrev=0", returnStdout: true).trim()
                // Login to Docker
                sh """
                    docker login -u ${DOCKER_USERNAME} -p ${DOCKER_TOKEN}
                """
                
                // Setup buildx
                sh """
                    docker buildx create --use --name builder || docker buildx use builder
                    docker buildx inspect --bootstrap
                """
                
                // Build and push multi-platform image
                sh """
                    docker buildx build \\
                        --platform linux/amd64,linux/arm64 \\
                        -t roarceus/api-server:${LATEST_TAG} \\
                        -t roarceus/api-server:latest \\
                        --push .
                """
            }
        }
    }

    stage('Tag Repository') {
        withCredentials([usernamePassword(credentialsId: 'github-pat', usernameVariable: 'GITHUB_USERNAME', passwordVariable: 'GITHUB_TOKEN')]) {
            sh """
                git config user.email "jenkins@csyeteam03.xyz"
                git config user.name "Automated Release Bot"
                git tag -a ${NEW_VERSION} -m "Release ${NEW_VERSION}"
                git push https://${GITHUB_USERNAME}:${GITHUB_TOKEN}@github.com/cyse7125-sp25-team03/api-server.git ${NEW_VERSION}
            """
        }
    }
}