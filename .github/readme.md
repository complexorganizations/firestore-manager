# Firestore (Export|Import|Delete)

#### Features
- Being able to import json documents to firestore
- Being able to delete a certain document from firestore
- Being able to import any size document to firestore


#### Usage
```
NAME:
   firestore-export-import - A new cli application

USAGE:
   firestore-export-import [global options] command [command options] [arguments...]

COMMANDS:
   create, c  create a document on firebase
   set, s     updates a document on firebase
   delete, d  deletes a document on firebase
   read, r    read a document from firebase
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
---
./firestore-export-import delete -auth auth_file.json --collection collection --document document
```
