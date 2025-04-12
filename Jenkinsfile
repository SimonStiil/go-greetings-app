properties([disableConcurrentBuilds(), buildDiscarder(logRotator(artifactDaysToKeepStr: '5', artifactNumToKeepStr: '5', daysToKeepStr: '5', numToKeepStr: '5'))])

@Library('pipeline-library')
import static dk.stiil.pipeline.Constants.*

podTemplate(yaml: '''
    apiVersion: v1
    kind: Pod
    spec:
      containers:
      - name: kaniko
        image: gcr.io/kaniko-project/executor:debug
        command:
        - sleep
        args: 
        - 99d
        volumeMounts:
        - name: kaniko-secret
          mountPath: /kaniko/.docker
      - name: k8s
        image: alpine/k8s:1.23.14
        command:
        - sleep
        args: 
        - 99d
      serviceAccountName: jenkins-tester
      restartPolicy: Never
      volumes:
      - name: kaniko-secret
        secret:
          secretName: dockercred
          items:
          - key: .dockerconfigjson
            path: config.json
''') {
  node(POD_LABEL) {
    LinkedHashMap<String,Object> scmInfo
    stage('checkout SCM') {  
      checkout scm
      scmInfo = scmGetInfo()
    }
    stage('Build Docker Image') {
      container('kaniko') {
        withEnv(["GIT_REPO_NAME=${scmInfo.repoName}"]) {
          sh '''
            /kaniko/executor --force --context `pwd` --destination registry.stiil.dk/jenkins/$GIT_REPO_NAME:$BRANCH_NAME
          '''
        }
      }
    }
    stage('Use Service account') {
      container('k8s') {
        setupKubernetesSA()
      }
    }
    stage('Deploy container') {
      container('k8s') {
        withEnv([
          "GIT_ORG_NAME=${scmInfo.orgName}",
          "GIT_REPO_NAME=${scmInfo.repoName}",
          "GIT_COMMIT=${scmInfo.commit}",
          "GIT_TAG=${scmInfo.tag}",
          "GIT_URL=${scmInfo.url}",
          "DEV_NAMESPACE=${DEV_NAMESPACE}"]) {
          sh '''
            helm upgrade -i -n $DEV_NAMESPACE $GIT_REPO_NAME-$BRANCH_NAME --set git.url=$GIT_URL \
            --set git.org=$GIT_ORG_NAME --set git.repo=$GIT_REPO_NAME --set git.commit=$GIT_COMMIT \
            --set git.tag=$GIT_TAG --set git.branch=$BRANCH_NAME ./base_chart
          '''
        }
      }
    }
  }
  stage('is service online'){
    waitForServiceToComeOnline serviceBaseURL: "https://${GIT_REPO_NAME}-${BRANCH_NAME}.k3s.stiil.dk/aktuator/health", maxRetries: 20
  }
}