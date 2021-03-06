package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"log"
)

// package-wide shared variables pointing to collections in mongodb
var (
	colAccounts    *mgo.Collection
	colTags        *mgo.Collection
	colDocuments   *mgo.Collection
	colPages       *mgo.Collection
	colAnnotations *mgo.Collection
	// colComments *mgo.Collection
	gridFS *mgo.GridFS
)

// initPersistency sets up database connection and initializes col* variables
// it also ensures indexes are existing and will give a fatal error when that fails.
func initPersistency() {
	// dial to localhost mongoDB instance
	mgoConn, err := mgo.Dial(flags.MongoHost)
	if err != nil {
		log.Fatalf("fatal error while dialing mgo connection: %s\n", err)
	}

	// get "nulpunt" database
	dbNulpunt := mgoConn.DB(fmt.Sprintf("nulpunt-%s", flags.Environment))

	// get gridfs
	gridFS = dbNulpunt.GridFS("fs")

	// get "accounts" collection
	colAccounts = dbNulpunt.C("accounts")

	// ensure that key "username" is unique for collection "users".
	err = colAccounts.EnsureIndex(mgo.Index{
		Key:    []string{"username"},
		Unique: true,
	})
	if err != nil {
		log.Fatalf("fatal error when ensuring index on accounts.username: %s\n", err)
	}

	// get "tags" collection
	colTags = dbNulpunt.C("tags")

	// ensure that key "tag" is unique for collection "tags".
	err = colTags.EnsureIndex(mgo.Index{
		Key:    []string{"tag"},
		Unique: true,
	})
	if err != nil {
		log.Fatalf("fatal error when ensuring index on tag.tag: %s\n", err)
	}

	// get "Documents" collection
	colDocuments = dbNulpunt.C("documents")

	// get "Pages" collection
	colPages = dbNulpunt.C("pages")

	// get "Annotations" collection
	colAnnotations = dbNulpunt.C("annotations")

	// get "Comments" collection
	// colComments = dbNulpunt.C("comments")
}
