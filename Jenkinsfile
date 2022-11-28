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
        volumeMounts:
        - name: kubeconfig-secret
          mountPath: /.kube/config
          subPath: config
        env:
        - name: KUBECONFIG
          value: "/.kube/config"
      restartPolicy: Never
      volumes:
      - name: kaniko-secret
        secret:
            secretName: dockercred
            items:
            - key: .dockerconfigjson
              path: config.json
      - name: kubeconfig-secret
        secret:
            secretName: master-kubeconfig
            items:
            - key: config
              path: config
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
    stage('Deploy container') {
      container('k8s') {
        sh '''
          helm upgrade -i -n jenkins-dev go-greetings-app-jenkins ./base_chart
        '''
      }
    }
  }
}