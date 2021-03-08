# Firestore Manager

#### Features
- Being able to import local json documents to firestore
- Being able to export json documents from firestore to local
- Being able to delete a certain document from firestore

---
### Installation

Lets generate GCP credentials from `https://cloud.google.com/iam/docs/creating-managing-service-accounts`, Once you have your service account, your ready.

#### Unix
```
export GOOGLE_APPLICATION_CREDENTIALS="[PATH]"
```

#### Windows
```
$env:GOOGLE_APPLICATION_CREDENTIALS="[PATH]"
```

Download the latest `firestore-manager` binary
```
go get -v github.com/complexorganizations/firestore-manager
```

Using `firestore-manager` binary
```
./firestore-manager read -auth [PATH] --collection contacts --document [DOCUMENT_ID] --file [PATH]
```
