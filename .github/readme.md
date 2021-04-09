# Firestore Manager

#### Features
- Being able to import local json documents to firestore
- Being able to export json documents from firestore to local
- Being able to delete a certain document from firestore

---
### Usage
```
NAME:
   firestore-manager - A new cli application

USAGE:
   firestore-manager [global options] command [command options] [arguments...]

COMMANDS:
   create, c  create a document on firebase
   set, s     updates a document on firebase
   delete, d  deletes a document on firebase
   read, r    read a document from firebase
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

---
### Installation

Lets generate GCP credentials from `https://cloud.google.com/iam/docs/creating-managing-service-accounts`, Once you have your service account, your ready.

#### Unix
```
export GOOGLE_APPLICATION_CREDENTIALS=[PATH]
```

#### Windows
```
$env:GOOGLE_APPLICATION_CREDENTIALS=[PATH]
```

Download the latest `firestore-manager` binary
```
go get -v github.com/complexorganizations/firestore-manager
```

Using `firestore-manager` binary
```
firestore-manager read -auth [PATH] --collection contacts --document [DOCUMENT_ID] --file [PATH]
```
