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
    stage('checkout SCM') {  
      checkout scm
    }
    stage('Build Docker Image') {
      container('kaniko') {
        sh '''
          /kaniko/executor --force --context `pwd` --destination registry.stiil.dk/jenkins/go-greetings-app:$BRANCH_NAME
        '''
      }
    }
    stage('Use Service account') {
      container('k8s') {
        sh '''
          export SAPATH=/var/run/secrets/kubernetes.io/serviceaccount 
          kubectl config set-cluster cfc --server=https://kubernetes.default --certificate-authority=$SAPATH/ca.crt
          set +x
          token=$(cat $SAPATH/token)
          kubectl config set-credentials cfc --token=${token}
          set -x
          kubectl config set-context cfc --cluster=cfc --user=cfc
          kubectl config use-context cfc
        '''
      }
    }
    stage('Deploy container') {
      container('k8s') {
        sh '''
          helm upgrade -i -n jenkins-dev go-greetings-app-jenkins ./base_chart
        '''
      }
    }
  }
  stage('is service online'){
    httpRequest url: 'https://go-greetings-app-jenkins.k3s.stiil.dk/aktuator/health',
            acceptType: 'APPLICATION_JSON',
            contentType: 'APPLICATION_JSON',
            validResponseCodes: "200"
  }
}